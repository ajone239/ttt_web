import { useState } from "react";

import { Board } from "./board"

export function Game() {

  const [xIsNext, setXIsNext] = useState(true);
  const [squares, setSquares] = useState(Array(9).fill(null))

  const player1 = "human"
  const player2 = "bot"

  const player = xIsNext ? player1 : player2

  const handlePlay = (nextMove) => {
    play_move(nextMove)
      .then(() => get_board())
      .then(board => setSquares(board))

    setXIsNext(!xIsNext);
  }

  if (player == "bot") {
    get_bot_move(handlePlay)
  }


  return (
    <div className="game">
      <div className="game-board">
        <Board xIsNext={xIsNext} player={player} squares={squares} onPlay={handlePlay} />
      </div>
      <div className="clear-button">
        <button onClick={() => clear_game(setSquares, setXIsNext)}>
          Clear
        </button>
      </div>
    </div>
  );
}

function get_bot_move(handlePlay) {
  fetch('/api/get_bot_move')
    .then(response => response.json())
    .then(data => handlePlay(data.move))
    .catch(err => console.log(err))
}

function play_move(move) {
  return fetch('/api/play_move', {
    method: "POST",
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      move: move
    }),
  }).catch(err => console.log(err))
}

function get_board() {
  return fetch('/api/get_board')
    .then(response => response.json())
    .then(data => {
      const board = data.board.flat().map((s) => {
        if (s == 0) {
          return null
        } else if (s == 1) {
          return "O"
        } else if (s == 2) {
          return "X"
        }
      })

      return board
    }).catch(err => console.log(err))
}

function clear_game(setSquares, setXIsNext) {
  return fetch('/api/new_game')
    .then(() => {
      const board = get_board();

      setSquares(board)
      setXIsNext(true)
    })
    .catch(err => console.log(err));
}
