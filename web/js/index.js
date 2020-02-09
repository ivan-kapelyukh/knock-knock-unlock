import { KnockRecorder } from "./knock_recorder.js";
import { login } from "./client.js";
import Record from "./record.js";
import Waves from "./waves.js";

let waves = new Waves();
let record = new Record();

let knockRecorder = new KnockRecorder();
knockRecorder.add_knock_listener(waves.knock.bind(waves));

console.log("Loading...");

document.getElementById("start").addEventListener("click", function() {
  console.log("Recording...");
  waves.play();
  knockRecorder.start();

  record.start();
});

document.getElementById("stop").addEventListener("click", function() {
  let knocks = knockRecorder.stop();
  console.log(knocks);
  login("test3", knocks).then(d => console.log(d.text()));

  record.stop();
});
