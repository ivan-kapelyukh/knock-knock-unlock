import p5 from "p5";
import "p5/lib/addons/p5.sound";

class Block {
  constructor(p, x) {
    this.p = p;
    this.height = 0;
    this.x = x;
  }

  draw(height) {
    let y = (this.p.windowHeight - height) / 2;

    this.p.rect(this.x, y, 8, height, 5);
  }
}

export default class Waves {
  constructor() {
    this._noise = new p5.Noise();
    this._noise.amp(0.01);

    this._knock = new p5.Envelope();
    this._knock.setADSR(0.01, 0.01, 0, 0);
    this._knock.setRange(1, 0);

    this._knockOsc = new p5.Oscillator(400, "sine");
    this._knockOsc.amp(this._knock);

    this._fft = new p5.FFT();
  }

  _sketch(p) {
    let blocks = [];
    let heights = [];
    let distance_between_blocks = 18;

    for (let x = 0; x < p.windowWidth; x += distance_between_blocks) {
      blocks.push(new Block(p, x));
      heights.push(0);
    }

    p.setup = () => {
      p.createCanvas(p.windowWidth, p.windowHeight);
      p.noStroke();
      p.fill(0);

      this._noise.start();
      this._knockOsc.start();
    };

    p.draw = () => {
      p.background(255);

      let spectrum = this._fft.analyze();

      for (let i = 0; i < blocks.length; i++) {
        let block = blocks[i];

        const maxHeight = 250;

        let height = 0;
        if (i < heights.length) {
          height = p.map(spectrum[i], 0, 255, 0, maxHeight);
        }

        let rgb = this.getColor(height / maxHeight);

        p.fill(rgb[0], rgb[1], rgb[2]);
        block.draw(height);
      }
    };
  }

  getColor(weight) {
    const color1 = [138, 35, 135];
    const color2 = [233, 64, 87];
    const color3 = [242, 113, 33];

    if (weight < 0.5) {
      let w1 = weight * 2;
      let w2 = 1 - w1;
      let rgb = [
        Math.round(color1[0] * w1 + color2[0] * w2),
        Math.round(color1[1] * w1 + color2[1] * w2),
        Math.round(color1[2] * w1 + color2[2] * w2)
      ];

      return rgb;
    } else {
      let w1 = (weight - 0.5) * 2;
      let w2 = 1 - w1;
      let rgb = [
        Math.round(color2[0] * w1 + color3[0] * w2),
        Math.round(color2[1] * w1 + color3[1] * w2),
        Math.round(color2[2] * w1 + color3[2] * w2)
      ];

      return rgb;
    }
  }

  play() {
    new p5(this._sketch.bind(this), window.document.getElementById("waves"));
  }

  knock() {
    this._knock.play();
  }
}
