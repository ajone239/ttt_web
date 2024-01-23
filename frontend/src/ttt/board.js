import { Square } from "./square";

export function Board({ xIsNext, squares, onPlay }) {

  const winner = calculateWinner(squares);
  let status;
  if (winner) {
    status = "Winner: " + winner;
  } else {
    status = "Next player: " + (xIsNext ? "X" : "O");
  }

  const handleClick = (i) => {
    if (squares[i] || calculateWinner(squares)) {
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
      {board}
    </>
  );
}

function calculateWinner(squares) {
  const lines = [
    [0, 1, 2],
    [3, 4, 5],
    [6, 7, 8],
    [0, 3, 6],
    [1, 4, 7],
    [2, 5, 8],
    [0, 4, 8],
    [2, 4, 6]
  ];
  for (let i = 0; i < lines.length; i++) {
    const [a, b, c] = lines[i];
    if (squares[a] && squares[a] === squares[b] && squares[a] === squares[c]) {
      return squares[a];
    }
  }
  return null;
}
