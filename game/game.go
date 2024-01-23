package game

type GameResult int

const (
	CrossWin  GameResult = iota
	Draw      GameResult = iota
	NoughtWin GameResult = iota
)

type Game struct {
	board    *Board
	whosTurn Square
}

func NewGame() *Game {
	return &Game{
		board:    NewBoard(),
		whosTurn: Cross,
	}
}

func (g *Game) ClearGame() {
	g.board = NewBoard()
	g.whosTurn = Cross
}

func (g *Game) GetBoard() Board {
	return *g.board
}

func (g *Game) PlayMove(i, j int) {
	g.board.MakeMove(i, j, g.whosTurn)
	g.SwitchPlayer()
}

// Is it player 1's turn?
func (g *Game) isPlayer1() bool {
	return g.whosTurn == Cross
}

// Switch the player
func (g *Game) SwitchPlayer() {
	g.whosTurn = g.whosTurn.Switch()
}
