package main

import (
  "fmt"
  "2048/board"
  "github.com/eiannone/keyboard"
)
const (
  UP keyboard.Key = iota
  DOWN
  LEFT 
  RIGHT 
  QUIT
  ERROR_KEY
)
func Read_input(b [][]int) bool {
  key, err := _get_key()
  if err != nil {
    fmt.Println(err.Error())
  }
  switch key {
  case UP:
    board.Move_up(b)
  case DOWN:
    board.Move_down(b)
  case RIGHT:
    board.Move_right(b)
  case LEFT:
    board.Move_left(b)
  case QUIT:
    fmt.Println("Game over!Bye~")
    return false
  }
  return true
}
func _get_key() (keyboard.Key, error) {
	char, key, err := keyboard.GetSingleKey()
	if err != nil {
		return ERROR_KEY, err
	}
	//fmt.Printf("You pressed: %c, key %X\r\n", char, key)
	if int(char) == 0 {
		switch key {
		case keyboard.KeyArrowUp:
			return UP, nil
		case keyboard.KeyArrowDown:
			return DOWN, nil
		case keyboard.KeyArrowLeft:
			return LEFT, nil
		case keyboard.KeyArrowRight:
			return RIGHT, nil
		case keyboard.KeyEsc:
			return QUIT, nil
		default:
			return ERROR_KEY, nil
		}
	} else {
		switch char {
		case 119:
			return UP, nil
		case 97:
			return LEFT, nil
		case 115:
			return DOWN, nil
		case 100:
			return RIGHT, nil
		default:
			return ERROR_KEY, nil
		}
	}
}
func main() {
  b := make([][]int, 4)
  for i := 0; i < 4; i ++ {
    b[i] = make([]int, 4)
  }
  fmt.Println("Let's start it!")
  for i := 0 ; i < 2; i ++ {
    board.Generate(b)
  }
  result := true
  for result == true {
    board.Display(b)
    result = Read_input(b)
  }
  // board.Display(b)
  // fmt.Println("")
  //
  // board.Move_up(b)
  // board.Display(b)
  // fmt.Println("")
  //
  // board.Move_left(b)
  // board.Display(b)
  // fmt.Println("")
  //
  // board.Move_right(b)
  // board.Display(b)
  // fmt.Println("")
  // 
  // board.Move_down(b)
  // board.Display(b)
  // fmt.Println("")
}
