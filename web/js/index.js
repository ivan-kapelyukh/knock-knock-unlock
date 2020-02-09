import "core-js/stable";
import "regenerator-runtime/runtime";
import { KnockRecorder } from "./knock_recorder.js";
import { login } from "./client.js";
import Waves from "./waves.js";

let waves = new Waves();
waves.play();

let knockRecorder = new KnockRecorder();
knockRecorder.add_knock_listener(waves.knock);

document.getElementById("start").addEventListener("click", function() {
  console.log("Recording...");
  knockRecorder.start();
});

document.getElementById("stop").addEventListener("click", function() {
  let knocks = knockRecorder.stop();
  console.log(knocks);
  login("test3", knocks).then(d => console.log(d));
});
