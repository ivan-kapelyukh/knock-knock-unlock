package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ivan-kapelyukh/knock-knock-unlock/internal/app/server/mongo"
)

type Server struct {
	port      int
	staticDir string
	mongo     *mongo.Mongo
}

func New(port int, staticDir string, mongo *mongo.Mongo) Server {
	return Server{port, staticDir, mongo}
}

func (s *Server) Serve() error {
	r := mux.NewRouter()
	r.HandleFunc("/register", s.registerHandler)
	r.HandleFunc("/login", s.loginHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(s.staticDir)))
	http.Handle("/", r)

	fmt.Printf("Listening on localhost:%v...\n", s.port)
	return http.ListenAndServe(fmt.Sprintf(":%v", s.port), nil)
}

func (s *Server) status(w http.ResponseWriter, code int) {
	body := fmt.Sprintf("%v %v", code, http.StatusText(code))
	http.Error(w, body, code)
}

func (s *Server) registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		s.status(w, http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var user mongo.User
	err := decoder.Decode(&user)

	if err != nil {
		log.Printf("Error decoding `%v` into mongo.User: %v\n", r.Body, err)
		s.status(w, http.StatusBadRequest)
		return
	}

	err = user.Validate()

	if err != nil {
		s.status(w, http.StatusBadRequest)
		return
	}

	err = s.mongo.Insert(user)

	if err != nil {
		log.Printf("Error storing %v into kku.users: %v\n", user, err)

		// TODO: this might also be an internal server error
		s.status(w, http.StatusBadRequest)
		return
	}

	s.status(w, http.StatusOK)
}

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		s.status(w, http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var login Login
	err := decoder.Decode(&login)

	if err != nil {
		log.Printf("Error decoding `%v` into server.Login: %v\n", r.Body, err)
		s.status(w, http.StatusBadRequest)
		return
	}

	err = login.Validate()

	if err != nil {
		s.status(w, http.StatusBadRequest)
		return
	}

	user, err := s.mongo.Find(login.Username)

	if err != nil {
		// FIXME: assuming the user does not exist
		user := mongo.User{login.Username, [][]int{login.Knock, login.Knock, login.Knock}}

		err = s.mongo.Insert(user)

		if err != nil {
			log.Printf("Error storing %v into kku.users: %v\n", user, err)

			// TODO: this might also be an internal server error
			s.status(w, http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "200 REGISTERED")
		return
	}

	if !MatchesMajority(login.Knock, user.Knocks) {
		s.status(w, http.StatusForbidden)
		return
	}

	s.status(w, http.StatusOK)
}
