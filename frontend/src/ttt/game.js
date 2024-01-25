import { useState } from "react";

import { Board } from "./board"

export function Game() {

  const [xIsNext, setXIsNext] = useState(true);
  const [squares, setSquares] = useState(Array(9).fill(null))
  const [currentMove, setCurrentMove] = useState(0);

  const handlePlay = (nextMove) => {
    play_move(nextMove)
      .then(() => get_board())
      .then(board => {
        setXIsNext(!xIsNext)
        setCurrentMove(currentMove + 1)
        setSquares(board)
        console.log("test1")
      })
    console.log("test2")

  }

  return (
    <div className="game">
      <div className="game-board">
        <Board xIsNext={xIsNext} currentMove={currentMove} squares={squares} onPlay={handlePlay} />
      </div>
      <div className="clear-button">
        <button onClick={() => clear_game(setSquares, setXIsNext)}>
          Clear
        </button>
      </div>
    </div>
  );
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
