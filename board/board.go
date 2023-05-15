package board

import (
	"fmt"
	"math/rand"
  "os/exec"
  "os"
	"time"
)

func Display(board [][]int) {
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()
  for i := 0; i < 4; i ++ {
    Print_ver_line()
    for j := 0; j < 4; j ++ {
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
  for i := 0; i < 44; i ++ {
   fmt.Print("-") 
  }
  fmt.Print("\n")
}

func Generate(board [][]int) {
  rand.Seed(time.Now().Unix())
  index := make([][2]int, 0) 
  for i := 0; i < 4; i ++ {
    for j := 0; j < 4; j ++ {
      if board[i][j] == 0 {
        index = append(index, [2]int{i,j})
      }
    }
  }
  rand_index := rand.Int() % len(index)
  nx,ny := index[rand_index][0], index[rand_index][1]
  board[nx][ny] = 2
}

func Move_up(board [][]int) {
  for i := 3; i > 0; i -- {
    for j := 3; j > 0; j -- {
      if board[i][j] != 0 && board[i - 1][j] == 0 {
        board[i - 1][j], board[i][j] = board[i][j], board[i - 1][j] 
      }
    }
  }
  Generate(board)
}

func Move_down(board [][]int) {
  for i := 0; i < 3; i ++ {
    for j := 0; j < 3; j ++ {
      if board[i][j] != 0 && board[i + 1][j] == 0 {
        board[i][j], board[i + 1][j] = board[i + 1][j], board[i][j]  
      }
    }
  }
  Generate(board)
}
func Move_right(board [][]int) {
  for i := 0; i < 4; i ++ {
    for j := 0; j < 3; j ++ {
      if board[i][j] != 0 && board[i][j + 1] == 0 {
        board[i][j], board[i][j + 1] = board[i][j + 1], board[i][j]  
      }
    }
  }
  Generate(board)
}
func Move_left(board [][]int) {
  for i := 3; i >= 0; i -- {
    for j := 3; j > 0; j -- {
      if board[i][j] != 0 && board[i][j - 1] == 0 {
        board[i][j], board[i][j - 1] = board[i][j - 1], board[i][j]  
      }
    }
  }
  Generate(board)
}
