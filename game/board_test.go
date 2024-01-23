package game

import "testing"

/*
 * Tests
 */

func TestNewBoard(t *testing.T) {
	board := NewBoard()
	if board == nil {
		t.Error("Board is nil")
	}
}

func TestCheckSquaresForWin(t *testing.T) {
	var s [BOARD_SIDE_LENGTH]square
	s = [BOARD_SIDE_LENGTH]square{blank, blank, blank}
	if checkSquaresForWin(s) != blank {
		t.Error("blank board should not win")
	}

	s = [BOARD_SIDE_LENGTH]square{blank, blank, nought}
	if checkSquaresForWin(s) != blank {
		t.Error("blank board should not win")
	}

	s = [BOARD_SIDE_LENGTH]square{cross, cross, cross}
	if checkSquaresForWin(s) != cross {
		t.Error("cross should win")
	}

	s = [BOARD_SIDE_LENGTH]square{nought, nought, nought}
	if checkSquaresForWin(s) != nought {
		t.Error("nought should win")
	}
}

func TestCheckRows(t *testing.T) {
	board := NewBoard()
	board.board = [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]square{
		{blank, blank, blank},
		{blank, blank, blank},
		{blank, blank, blank},
	}
	if board.checkRows() != blank {
		t.Error("blank board should not win")
	}

  // Check middle row
	board.board = [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]square{
		{blank, blank, blank},
		{cross, cross, cross},
		{blank, blank, blank},
	}
	if board.checkRows() != cross {
		t.Error("cross should win")
	}

	board.board = [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]square{
		{blank, blank, blank},
		{nought, nought, nought},
		{blank, blank, blank},
	}
	if board.checkRows() != nought {
		t.Error("nought should win")
	}

  // Check top row
  board.board = [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]square{
    {cross, cross, cross},
    {blank, blank, blank},
    {blank, blank, blank},
  }
  if board.checkRows() != cross {
    t.Error("cross should win")
  }
  // Check bottom row
  board.board = [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]square{
    {blank, blank, blank},
    {blank, blank, blank},
    {nought, nought, nought},
  }
  if board.checkRows() != nought {
    t.Error("nought should win")
  }
}

func TestCheckColumns(t *testing.T) {
	board := NewBoard()
	board.board = [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]square{
		{blank, blank, blank},
		{blank, blank, blank},
		{blank, blank, blank},
	}
	if board.checkColumns() != blank {
		t.Error("blank board should not win")
	}

  // Check middle column
	board.board = [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]square{
		{blank, cross, blank},
		{blank, cross, blank},
		{blank, cross, blank},
	}
	if board.checkColumns() != cross {
		t.Error("cross should win")
	}

	board.board = [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]square{
		{blank, nought, blank},
		{blank, nought, blank},
		{blank, nought, blank},
	}
	if board.checkColumns() != nought {
		t.Error("nought should win")
	}

  // Check left column
  board.board = [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]square{
    {cross, blank, blank},
    {cross, blank, blank},
    {cross, blank, blank},
  }
  if board.checkColumns() != cross {
    t.Error("cross should win")
  }
  // Check right column
  board.board = [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]square{
    {blank, blank, nought},
    {blank, blank, nought},
    {blank, blank, nought},
  }
  if board.checkColumns() != nought {
    t.Error("nought should win")
  }
}

func TestCheckDiagonals(t *testing.T) {
	board := NewBoard()
	board.board = [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]square{
		{blank, blank, blank},
		{blank, blank, blank},
		{blank, blank, blank},
	}
	if board.checkDiagonols() != blank {
		t.Error("blank board should not win")
	}

	board.board = [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]square{
		{cross, blank, blank},
		{blank, cross, blank},
		{blank, blank, cross},
	}
	if board.checkDiagonols() != cross {
		t.Error("cross should win")
	}

	board.board = [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]square{
		{blank, blank, nought},
		{blank, nought, blank},
		{nought, blank, blank},
	}
	if board.checkDiagonols() != nought {
		t.Error("nought should win")
	}
}

func TestCheckForWin(t *testing.T) {
	board := NewBoard()
	board.board = [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]square{
		{blank, blank, blank},
		{blank, blank, blank},
		{blank, blank, blank},
	}
	if board.CheckForWin() != blank {
		t.Error("blank board should not win")
	}

	board.board = [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]square{
		{cross, blank, blank},
		{blank, cross, blank},
		{blank, blank, cross},
	}
	if board.CheckForWin() != cross {
		t.Error("cross should win")
	}

	board.board = [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]square{
		{blank, blank, nought},
		{blank, nought, blank},
		{nought, blank, blank},
	}
	if board.CheckForWin() != nought {
		t.Error("nought should win")
	}
}

func TestEvaluate(t *testing.T) {
	board := NewBoard()
  // Check nought
	board.board = [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]square{
		{nought, nought, nought},
		{blank, blank, blank},
		{blank, blank, blank},
	}
	if board.Evaluate() != -1 {
		t.Error("Expected -1, got ", board.Evaluate())
	}

  // Check cross
  board.board = [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]square{
    {cross, cross, cross},
    {blank, blank, blank},
    {blank, blank, blank},
  }
  if board.Evaluate() != 1 {
    t.Error("Expected 1, got ", board.Evaluate())
  }

  // Check draw
  board.board = [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]square{
    {cross, nought, cross},
    {cross, nought, nought},
    {nought, cross, cross},
  }
  if board.Evaluate() != 0 {
    t.Error("Expected 0, got ", board.Evaluate())
  }

  // Test this board
  // XOX
  // OOX
  // OXX
  board.board = [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]square{
    {cross, nought, cross},
    {nought, cross, cross},
    {nought, cross, cross},
  }
  if board.Evaluate() != 1 {
    t.Error("Expected 1, got ", board.Evaluate())
  }

  // Flip the board over the diagonal
  // XOO
  // OOX
  // XXX
  board.board = [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]square{
    {cross, nought, nought},
    {nought, cross, cross},
    {cross, cross, cross},
  }
  if board.Evaluate() != 1 {
    t.Error("Expected 1, got ", board.Evaluate())
  }
}
