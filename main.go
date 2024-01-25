package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ajone239/ttt_web/game"
	"github.com/ajone239/ttt_web/game/bot"
)

var player1 game.Player = &game.HumanPlayer{}
var player2 game.Player = bot.NewBotPlayer(2, false)
var globalGame = game.NewGame(&player1, &player2)

func main() {
	// Static files
	fs := http.FileServer(http.Dir("./frontend/build"))
	http.Handle("/", fs)

	// API routes
	api := http.NewServeMux()
	api.HandleFunc("/api/new_game", newBoardHandler)
	api.HandleFunc("/api/get_board", getBoardHandler)
	api.HandleFunc("/api/play_move", playMoveHandler)
	api.HandleFunc("/api/board_result", boardResultHandler)
	api.HandleFunc("/api/get_bot_move", getBotMoveHandler)
	http.Handle("/api/", api)

	http.ListenAndServe(":8080", nil)
}

func newBoardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a hit new board")

	globalGame.ClearGame()
}

func getBoardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a hit get board")

	board := globalGame.GetBoard()

	fmt.Println(board)

	json.NewEncoder(w).Encode(board)
}

type Move struct {
	Move int `json:"move"`
}

func playMoveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a hit play move")
	decoder := json.NewDecoder(r.Body)

	var data Move
	err := decoder.Decode(&data)

	if err != nil {
		fmt.Println(err)
		panic("failed")
	}

	i := data.Move / 3
	j := data.Move % 3

	globalGame.PlayMove(i, j)
}

func getBotMoveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a get bot move")

	board := globalGame.GetBoard()

	if board.IsFull() {
		json.NewEncoder(w).Encode(nil)
		return
	}

	i, j := globalGame.GetMove()

	moveData := Move{
		Move: (i * 3) + j,
	}

	json.NewEncoder(w).Encode(moveData)
}

type BoardResult struct {
	Terminal bool        `json:"terminal"`
	Result   game.Square `json:"result"`
}

func boardResultHandler(w http.ResponseWriter, r *http.Request) {

	board := globalGame.GetBoard()

	result := board.CheckForWin()

	// Terminal if:
	// - the result isn't blank: i.e. noughts or crosses won
	// - the board is full: i.e. a draw or a win with a full board
	terminal := result != game.Blank || board.IsFull()

	boardResult := BoardResult{
		Terminal: terminal,
		Result:   result,
	}

	json.NewEncoder(w).Encode(boardResult)
}
