package board

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

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

func Start_game(board [][]int) {
	for i := 0; i < 2; i++ {
		Generate(board)
	} // generate 2 number in random place
	result := true
	for result && !find_2048(board) {
		Display(board)
		result = Read_input(board)
	}
}

func find_2048(board [][]int) bool {
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if board[i][j] == 2048 {
				return true
			}
		}
	}
	return false
}

func Display(board [][]int) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	for i := 0; i < 4; i++ {
		Print_ver_line()
		for j := 0; j < 4; j++ {
			// every number occupies 10 space
			if board[i][j] == 0 {
				fmt.Printf("%10s", "")
			} else {
				fmt.Printf("%7d%3s", board[i][j], "")
			}
			Print_ver_line()
		}
		fmt.Print("\n")
		Print_hor_line()
	}
}

func Print_ver_line() {
	fmt.Print("|")
}
func Print_hor_line() {
	for i := 0; i < 44; i++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
}

func Generate(board [][]int) {
	rand.Seed(time.Now().Unix())
	index := make([][2]int, 0)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if board[i][j] == 0 {
				index = append(index, [2]int{i, j})
			}
		}
	}
	rand_index := rand.Int() % len(index)
	nx, ny := index[rand_index][0], index[rand_index][1]
	board[nx][ny] = 2
}

func Move_up(board [][]int) {
	for i := 3; i > 0; i-- {
		for j := 3; j >= 0; j-- {
			if board[i][j] != 0 && board[i-1][j] == 0 {
				board[i-1][j], board[i][j] = board[i][j], board[i-1][j]
			} else if board[i][j] != 0 && board[i-1][j] == board[i][j] {
				board[i-1][j], board[i][j] = board[i][j]+board[i-1][j], 0
			}
		}
	}
	Generate(board)
}

func Move_down(board [][]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j <= 3; j++ {
			if board[i][j] != 0 && board[i+1][j] == 0 {
				board[i][j], board[i+1][j] = board[i+1][j], board[i][j]
			} else if board[i][j] != 0 && board[i+1][j] == board[i][j] {
				board[i+1][j], board[i][j] = board[i][j]+board[i+1][j], 0
			}
		}
	}
	Generate(board)
}
func Move_right(board [][]int) {
	for i := 0; i <= 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] != 0 && board[i][j+1] == 0 {
				board[i][j], board[i][j+1] = board[i][j+1], board[i][j]
			} else if board[i][j] != 0 && board[i][j+1] == board[i][j] {
				board[i][j+1], board[i][j] = board[i][j]+board[i][j+1], 0
			}
		}
	}
	Generate(board)
}
func Move_left(board [][]int) {
	for i := 3; i >= 0; i-- {
		for j := 3; j > 0; j-- {
			if board[i][j] != 0 && board[i][j-1] == 0 {
				board[i][j], board[i][j-1] = board[i][j-1], board[i][j]
			} else if board[i][j] != 0 && board[i][j-1] == board[i][j] {
				board[i][j-1], board[i][j] = board[i][j]+board[i][j-1], 0
			}

		}
	}
	Generate(board)
}
func Read_input(board [][]int) bool {
	key, err := _get_key()
	if err != nil {
		fmt.Println(err.Error())
	}
	switch key {
	case UP:
		Move_up(board)
	case DOWN:
		Move_down(board)
	case RIGHT:
		Move_right(board)
	case LEFT:
		Move_left(board)
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
