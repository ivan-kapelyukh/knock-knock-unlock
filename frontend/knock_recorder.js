class KnockRecorder {
  constructor() {
    this._knocks = [];
  }

  start() {
    document.addEventListener("keypress", this._keyPressHandler);
  }

  stop() {
    document.removeEventListener("keypress", this._keyPressHandler);

    return this._processed_knocks();
  }

  _processed_knocks() {
    const res = [];

    if (this._knocks.length == 0) {
      return [];
    }

    let last_knock = this._knocks[0];
    for (let i = 1; i < res.length; i++) {
      let knock_interval = this._knocks[i] - last_knock;
      res.push(knock_interval);
      last_knock = this._knocks[i];
    }

    return res;
  }

  _keyPressHandler(e) {
    this._knocks.push(Date.now().getTime());
  }
}
