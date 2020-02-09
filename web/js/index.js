import { KnockRecorder } from "./knock_recorder.js";
import { login } from "./client.js";
import Waves from "./waves.js";

let waves = new Waves();

let knockRecorder = new KnockRecorder();
knockRecorder.add_knock_listener(waves.knock.bind(waves));

document.getElementById("username-input").value = "";
document.getElementById("username-input").focus();

document.getElementById("start").addEventListener("submit", function(e) {
  e.preventDefault();
  document.getElementById("username-input").blur();
  console.log("Recording...");
  document.getElementById("username").classList.add("hide");
  document.getElementById("waves-view").style.display = "flex";
  waves.play();
  knockRecorder.start();

  document.addEventListener("keypress", keyHandler);
});

function keyHandler(e) {
  if (e.keyCode != 13) {
    return;
  }

  document.removeEventListener("keypress", keyHandler);

  // Enter pressed
  let knocks = knockRecorder.stop();
  console.log(knocks);

  let username = document.getElementById("username-input").value;

  console.log("Username", username);

  let logginResp = login(username, knocks);

  console.log("logginResp", logginResp);

  logginResp
    .then(d => {
      if (d.status != 200) {
        throw "not 200!";
      }

      console.log("logginResp", logginResp);
      console.log("Return data", d);

      document.getElementById("waves-view").style.display = "none";
      document.getElementById("waves").style.display = "none";

      d.text().then(text => {
        console.log("text", text);
        if (text.includes("REGISTERED")) {
          console.log("Registered!");
          document.getElementById("registered").style.display = "flex";
        } else {
          console.log("Logged in!");
          document.getElementById("success").style.display = "flex";
        }
      });
    })
    .catch(d => {
      console.log("logginResp", logginResp);
      console.log("Failed!", d);
      document.getElementById("waves-view").style.display = "none";
      document.getElementById("waves").style.display = "none";
      document.getElementById("failure").style.display = "flex";
    });
}
