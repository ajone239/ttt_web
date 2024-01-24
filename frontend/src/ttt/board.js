import { useState } from "react";

import { Square } from "./square";

export function Board({ xIsNext, player, squares, onPlay }) {

  const [winner, setWinner] = useState(null)

  calculateWinner(setWinner)

  let status;
  if (winner) {
    if (winner == "Draw") {
      status = "Draw!"
    } else {
      status = "Winner: " + winner;
    }
  } else {
    status = "Next player: " + (xIsNext ? "X" : "O");
  }

  let inputString;
  if (winner) {
    inputString = "Game over"
  } else if (player == "human") {
    inputString = "Waiting on human input..."
  } else if (player == "bot") {
    inputString = "Bot is thinking"
  } else {
    inputString = "idk man..."
  }

  const handleClick = (i) => {
    calculateWinner(setWinner)

    if (player == "bot" ||
      squares[i] ||
      winner) {
      return;
    }

    onPlay(i);
  }

  var board = [];

  for (var i = 0; i < 3; i++) {
    var row = [];
    for (var j = 0; j < 3; j++) {
      const idx = (i * 3) + j;
      row.push((
        <Square value={squares[idx]} onSquareClick={() => handleClick(idx)} />
      ));
    }
    board.push((
      <div className="board-row">
        {row}
      </div>
    ));
  }

  return (
    <>
      <div className="status">{status}</div>
      <div className="input-str">{inputString}</div>
      {board}
    </>
  );
}

function calculateWinner(setWinner) {
  fetch('/api/board_result')
    .then(response => response.json())
    .then(data => {
      // {terminal: false, result: 0}
      if (!data.terminal) {
        return null
      }

      const s = data.result

      if (s == 0) {
        return "Draw"
      } else if (s == 1) {
        return "O"
      } else if (s == 2) {
        return "X"
      }
    }
    )
    .then(res => setWinner(res))
    .catch(err => console.log(err));
}
