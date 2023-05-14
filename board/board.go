package board

import (
	"fmt"
	"math/rand"
	"time"
	// "strconv"
)
type Board interface{
  Display()
}
type board struct {
  board [][]int
}
var b board
func Display() {
  b.board = generate()
  for i := 0; i < 4; i ++ {
    for j := 0; j < 4; j ++ {
      fmt.Print(" ")
      if b.board[i][j] == 0 {
        fmt.Print(" ")
      } else {
        fmt.Print(b.board[i][j])
      }
    }
    fmt.Print("\n")
  }
}

func generate() [][]int {
  nums := make([]int, 0)
  nums = append(nums, 0)
  for i := 2; i <= 2048; i *= 2{
    nums = append(nums, i)
  }

  rand.Seed(time.Now().Unix())
  matrix := make([][]int, 4)
  for i := 0; i < 4; i ++ {
    matrix[i] = make([]int, 4)
  }
  for i := 0; i < 4; i ++ {
    for j := 0; j < 4; j ++ {
      matrix[i][j] = nums[rand.Int()%len(nums)]
    }
  }
  return matrix
}
