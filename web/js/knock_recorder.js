export class KnockRecorder {
  constructor() {
    this._knocks = [];
    this._listeners = [];
  }

  add_knock_listener(listener) {
    this._listeners.push(listener);
  }

  start() {
    this._knocks = [];
    document.addEventListener("keypress", () => this._keyPressHandler());
  }

  stop() {
    document.removeEventListener("keypress", () => this._keyPressHandler());
    return this._processed_knocks();
  }

  _processed_knocks() {
    const res = [];

    if (this._knocks.length == 0) {
      return [];
    }

    let last_knock = this._knocks[0];
    for (let i = 1; i < this._knocks.length; i++) {
      let knock_interval = this._knocks[i] - last_knock;
      res.push(knock_interval);
      last_knock = this._knocks[i];
    }

    return res;
  }

  _keyPressHandler() {
    this._knocks.push(Date.now());

    for (let listener of this._listeners) {
      listener();
    }
  }
}