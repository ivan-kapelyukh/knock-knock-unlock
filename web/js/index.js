import { KnockRecorder } from "./knock_recorder.js";
import { login } from "./client.js";
import Waves from "./waves.js";

let waves = new Waves();

let knockRecorder = new KnockRecorder();
knockRecorder.add_knock_listener(waves.knock.bind(waves));

console.log("Loading...");

document.getElementById("start").addEventListener("submit", function(e) {
  e.preventDefault();
  console.log("Recording...");
  document.getElementById("username").style.display = "none";
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

  let username = document.getElementById("username").value;

  console.log("Username", username);

  login(username, knocks)
    .then(d => {
      console.log("Return data", d);

      document.getElementById("waves-view").style.display = "none";
      document.getElementById("waves").style.display = "none";

      if (d.includes("registered")) {
        console.log("Registered!");
        document.getElementById("registered").style.display = "flex";
      } else {
        console.log("Logged in!");
        document.getElementById("success").style.display = "flex";
      }
    })
    .catch(() => {
      console.log("Failed!");
      document.getElementById("waves-view").style.display = "none";
      document.getElementById("waves").style.display = "none";
      document.getElementById("failure").style.display = "flex";
    });
}
