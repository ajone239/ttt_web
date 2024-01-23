package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ajone239/ttt_web/game"
)

var globalGame = game.NewGame()

func main() {
	// Static files
	fs := http.FileServer(http.Dir("./frontend/build"))
	http.Handle("/", fs)

	// API routes
	api := http.NewServeMux()
	api.HandleFunc("/api/new_game", newBoardHandler)
	api.HandleFunc("/api/get_board", getBoardHandler)
	api.HandleFunc("/api/play_move", playMoveHandler)
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
