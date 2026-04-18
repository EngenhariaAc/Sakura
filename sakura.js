class Sakura {
  constructor(color, bloomSeason) {
    this.color = color;
    this.bloomSeason = bloomSeason;
  }

  bloom() {
    console.log(`The ${this.color} sakura blooms beautifully in ${this.bloomSeason}. 🌸`);
  }
}

module.exports = Sakura;
