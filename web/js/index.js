import { KnockRecorder } from "./knock_recorder.js";
import { login } from "./client.js";

let knock_recorder;

document.addEventListener("DOMContentLoaded", function(event) {
  document.getElementById("start").addEventListener("click", function() {
    startRecording();
  });

  document.getElementById("stop").addEventListener("click", function() {
    stopRecording();
  });
});

function startRecording() {
  console.log("Recording...");
  knock_recorder = new KnockRecorder();
  knock_recorder.start();
}

function stopRecording() {
  if (knock_recorder === undefined) {
    console.log("Knock recorder not started!");
    return;
  }
  let knocks = knock_recorder.stop();

  console.log(knocks);
  login("test1", knocks);

  delete knock_recorder;
}
