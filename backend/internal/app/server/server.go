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
	port  int
	mongo *mongo.Mongo
}

func New(port int, mongo *mongo.Mongo) Server {
	return Server{port, mongo}
}

func (s *Server) Serve() error {
	r := mux.NewRouter()
	r.HandleFunc("/register", s.registerHandler)
	r.HandleFunc("/login", s.loginHandler)
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
		s.status(w, http.StatusForbidden)
		return
	}

	// FIXME: this is a dummy length check just to validate *something*
	valid := 0
	for _, knock := range user.Knocks {
		if len(knock) == len(login.Knock) {
			valid++
		}
	}

	if valid < len(user.Knocks)/2 {
		s.status(w, http.StatusForbidden)
		return
	}

	fmt.Fprintf(w, "Found user: %v", user)
}
