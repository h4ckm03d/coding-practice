package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// Direction to crawler
var Direction = [][]int{
	{-1, 0},  // top
	{-1, 1},  // top right
	{0, 1},   // right
	{1, 1},   // right bottom
	{1, 0},   // bottom
	{1, -1},  // bottom left
	{0, -1},  // left
	{-1, -1}, // top left
}

// ClickResult to represent click result
type ClickResult int

const (
	// Clicked means cell already clicked
	Clicked ClickResult = iota
	// Beware represent selected cell nearby mine
	Beware
	// Mine represent selected cell is mine that means you lose
	Mine
	// Empty return empty cell when select
	Empty
)

// Cell contains information of each block on the board
type Cell struct {
	Value   int
	Mine    bool
	Clicked bool
	Safe    bool
}

// Render cell
func (c *Cell) Render() string {
	switch {
	case c.Clicked == false:
		return "[#]"
	case c.Mine == true:
		return "[M]"
	case c.Value == 0:
		return "[ ]"
	default:
		return fmt.Sprintf("[%d]", c.Value)
	}
}

// Board is cell container
type Board struct {
	Width   int
	Height  int
	MineNum int
	grid    []*Cell
}

// Get grid from given position
func (b *Board) Get(column, row int) *Cell {

	if b.grid == nil {
		return nil
	}

	if column < 0 || column >= b.Width || row < 0 || row >= b.Height {
		return nil
	}

	return b.grid[row*b.Width+column]
}

// Init board value
func (b *Board) Init() (err error) {

	if b.MineNum > b.Height*b.Width-1 {
		err = errors.New("number of mines cannot exceed number of cells")
		return err
	}
	// init board with empty cells
	mines := b.genMines()
	//loop through all cells in the slice, keeping track of index
	//if mines at index is true, set cell.mine to true
	for index := range b.grid {
		b.grid[index] = &Cell{
			Mine:    mines[index],
			Clicked: false,
			Value:   0,
		}
	}

	//go through each element in each row and check if
	for r := 0; r < b.Height; r++ {
		for c := 0; c < b.Width; c++ {
			cell := b.Get(c, r)
			if cell.Mine {
				for _, d := range Direction {
					if neighbour := b.Get(c+d[0], r+d[1]); neighbour != nil {
						neighbour.Value++
					}
				}
			}
		}
	}
	return err
}

func (b *Board) genMines() []bool {
	mineIndexList := make([]bool, (b.Height * b.Width))
	i := 0
	for i < b.MineNum {
		index := rand.Intn(b.Height * b.Width)
		if mineIndexList[index] == false {
			mineIndexList[index] = true
			i++
		}
	}
	return mineIndexList
}

// RenderAll to show all board
func (b *Board) RenderAll() {

	for row := 0; row < b.Height; row++ {
		for col := 0; col < b.Width; col++ {
			cell := b.Get(col, row)
			cell.Clicked = true
			fmt.Printf(cell.Render())
		}
		fmt.Printf("\n")
	}
}

// NewBoard create new instance of board
func NewBoard(width, height, mine int) *Board {
	return &Board{
		width,
		height,
		mine,
		make([]*Cell, width*height),
	}
}

func main() {
	var n int
	flag.IntVar(&n, "n", 5, "board size")
	flag.Parse()
	fmt.Println(n)
	rand.Seed(time.Now().Unix())
	b := NewBoard(n, n, n)
	b.Init()
	b.RenderAll()
}
