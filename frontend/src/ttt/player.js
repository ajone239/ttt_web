export class Player {
  constructor(playerNumber, handlePlay) {
    this.turn = playerNumber
    this.handlePlay = handlePlay
  }
}

export class HumanPlayer extends Player {
  notify() { return }
}

export class BotPlayer extends Player {
  notify() {
    this.get_bot_move(this.handlePlay)
  }

  get_bot_move(handlePlay) {
    fetch('/api/get_bot_move')
      .then(response => response.json())
      .then(data => {
        console.log(data)
        if (!data) {
          return
        }
        if (!Object.keys(data).includes("move")) {
          return
        }
        handlePlay(data.move)
      })
      .catch(err => console.log(err))
  }
}
