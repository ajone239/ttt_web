package game

type GameResult int

const (
	CrossWin  GameResult = iota
	Draw      GameResult = iota
	NoughtWin GameResult = iota
)

type Game struct {
	player1  *Player
	player2  *Player
	board    *Board
	whosTurn Square
}

func NewGame(p1, p2 *Player) *Game {
	return &Game{
		player1:  p1,
		player2:  p2,
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

func (g *Game) GetMove() (int, int) {
	var player *Player

	if g.whosTurn == Cross {
		player = g.player1
	} else {
		player = g.player2
	}
	return (*player).GetMove(g.board)
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
