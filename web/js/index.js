import { KnockRecorder } from "./knock_recorder.js";
import { login } from "./client.js";
import Waves from "./waves.js";

let waves = new Waves();

let knockRecorder = new KnockRecorder();
knockRecorder.add_knock_listener(waves.knock.bind(waves));

console.log("Loading...");

document.getElementById("start").addEventListener("click", function() {
  console.log("Recording...");
  document.getElementById("username").style.display = "none";
  document.getElementById("waves-view").style.display = "flex";
  waves.play();
  knockRecorder.start();
});

document.getElementById("stop").addEventListener("click", function() {
  let knocks = knockRecorder.stop();
  console.log(knocks);
  login("test3", knocks).then(d => console.log(d.text()));
});
