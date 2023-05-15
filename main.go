package main

import (
	"2048/board"
	"fmt"
)

func main() {
	b := make([][]int, 4)
	for i := 0; i < 4; i++ {
		b[i] = make([]int, 4)
	} // init board
	fmt.Println("Let's start it!Press enter to start")
	fmt.Println("Use W A S D to move up, move left, move down and move right ")
  start := ""
  fmt.Scanln(&start)
	// start game
	board.Start_game(b)
  fmt.Println("You win!")
}
