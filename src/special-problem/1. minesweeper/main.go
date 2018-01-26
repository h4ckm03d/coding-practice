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

// Click to get result
func (c *Cell) Click() ClickResult {
	switch {
	case c.Mine == true:
		return Mine
	case c.Clicked == true:
		return Clicked
	case c.Value == 0:
		return Empty
	case c.Value > 0:
		return Beware
	}
	return Clicked
}

// Board is cell container
type Board struct {
	Width   int
	Height  int
	MineNum int
	grid    []*Cell
}

func (b *Board) checkPosition(r, c int) bool {
	if r < 0 || r >= b.Height {
		return false
	}

	if c < 0 || c >= b.Width {
		return false
	}
	return true
}

// Get grid from given position
func (b *Board) Get(row, column int) *Cell {
	if b.grid == nil {
		return nil
	}

	if !b.checkPosition(row, column) {
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
			cell := b.Get(r, c)
			if cell != nil && cell.Mine {
				for _, d := range Direction {
					if neighbour := b.Get(r+d[0], c+d[1]); neighbour != nil {
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
			cell := b.Get(row, col)
			cell.Clicked = true
			fmt.Printf(cell.Render())
		}
		fmt.Printf("\n")
	}
}

// Render Board
func (b *Board) Render() {
	for row := 0; row < b.Height; row++ {
		for col := 0; col < b.Width; col++ {
			cell := b.Get(row, col)
			fmt.Printf(cell.Render())
		}
		fmt.Printf("\n")
	}
}

func (b *Board) UpdateValue(r, c int) int {
	openedBoardNum := 0

	for _, v := range Direction {
		if ok, val := b.revealCell(r+v[0], c+v[1]); ok && val > 0 {
			openedBoardNum++
		} else if val == 0 {
			b.Get(r+v[0], c+v[1]).Clicked = false
		}
	}

	return openedBoardNum
}

// RevealCell to show cell by given row column
func (b *Board) revealCell(r, c int) (bool, int) {
	cell := b.Get(r, c)
	if cell == nil ||
		cell.Mine ||
		cell.Clicked {
		return false, -1
	}
	cell.Clicked = true
	return true, cell.Value
}

func (b *Board) CrawlerEmpty(r, c int) int {
	openedBoardNum := 0
	b.revealCell(r, c)
	for _, v := range Direction {
		ok, val := b.revealCell(r+v[0], c+v[1])
		crawler := false
		if ok && val == 0 {
			crawler = true
		}
		ta, tb := v[0], v[1]
		for crawler {
			ta += v[0]
			tb += v[1]
			ok, v := b.revealCell(r+ta, c+tb)
			if ok && v == 0 {
				crawler = true
				openedBoardNum++
			} else {
				crawler = false
			}
		}
	}

	return openedBoardNum
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

func mainloop(b *Board) {
	squaresToReveal := (b.Width * b.Height) - b.MineNum

	var WinCondition = true
	fmt.Println("MineSweeper Ala-ala")
	for squaresToReveal > 0 && WinCondition {
		b.Render()
		fmt.Print("> ")
		var r, c int
		_, err := fmt.Scanf("%d %d", &r, &c)

		if err != nil {
			fmt.Println(err)
			continue
		}
		if b.checkPosition(r, c) {
			cell := b.Get(r, c)
			switch cell.Click() {
			case Mine:
				cell.Clicked = true
				WinCondition = false
			case Beware:
				fmt.Println("beware nearby mines")
				revealedNeighbour := b.UpdateValue(r, c) // get nearby number
				squaresToReveal -= revealedNeighbour + 1
				cell.Clicked = true
			case Empty:
				revealedNeighbour := b.CrawlerEmpty(r, c)
				fmt.Println("empty")
				squaresToReveal -= revealedNeighbour + 1
				cell.Clicked = true
			case Clicked:
				fmt.Println("you've already clicked that position")
			}
		}
	}
	b.RenderAll()
	if WinCondition {
		fmt.Println("You Won")
	} else {
		fmt.Println("You Lose")
	}
	return
}

func main() {
	var n int
	flag.IntVar(&n, "n", 5, "board size")

	flag.Parse()
	rand.Seed(time.Now().Unix())
	b := NewBoard(n, n, n)
	b.Init()
	mainloop(b)
}
