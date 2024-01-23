import { useState } from "react";

import { Board } from "./board"

export function Game() {

  const [xIsNext, setXIsNext] = useState(true);
  const [squares, setSquares] = useState(Array(9).fill(null))
  const [currentMove, setCurrentMove] = useState(0);

  const handlePlay = (nextMove) => {
    f = async () => {
      await play_move(nextMove)

      const board = await get_board()

      setSquares(board)

      setCurrentMove(currentMove + 1);
      setXIsNext(!xIsNext);
    }
    f()
  }


  return (
    <div className="game">
      <div className="game-board">
        <Board xIsNext={xIsNext} squares={squares} onPlay={handlePlay} />
      </div>
      <div className="">
        <button onClick={() => clear_game(setSquares)}>
          Clear
        </button>
      </div>
    </div>
  );
}

async function play_move(move) {
  await fetch('/api/play_move', {
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

async function get_board() {
  const response = await fetch('/api/get_board')
  const data = await response.json()
  console.log(data)

  const board = data.board.flat().map((s) => {
    if (s == 0) {
      return null
    } else if (s == 1) {
      return "O"
    } else if (s == 2) {
      return "X"
    }
  })

  console.log(board)

  return board
}

async function clear_game(setSquares) {
  await fetch('/api/new_game')
    .catch(err => console.log(err));

  const board = await get_board();

  setSquares(board)
}
