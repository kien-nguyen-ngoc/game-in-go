package object

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"math/rand"
	"time"
)

const (
	WindowWidth  = 250
	WindowHeight = 250

	rows    = 100
	columns = 100
)

type Board struct {
	Cells [][]*Cell
	Tanks []Tank

	Window  *glfw.Window
	Program uint32
}

func (board *Board) GetBoardWidth() int {
	return WindowWidth
}

func (board *Board) GetBoardHeight() int {
	return WindowHeight
}
func (board *Board) GetBoardRows() int {
	return rows
}
func (board *Board) GetBoardColumns() int {
	return columns
}

func (board *Board) DestroyTank(destroyId int) {
	tanks := make([]Tank, 0)
	for _, tank := range board.Tanks {
		if tank.Id != destroyId {
			tanks = append(tanks, tank)
		}
	}
	board.Tanks = tanks
}

func (board *Board) MakeTanks(num_of_tank int) {
	tanks := make([]Tank, 0)
	for i := 0; i < num_of_tank; i++ {
		tanks = append(tanks, *board.New(board.GetBoardRows(), board.GetBoardColumns(), i, false))
	}
	board.Tanks = tanks
}

func (board *Board) MakeCells() {
	rand.Seed(time.Now().UnixNano())

	cells := make([][]*Cell, rows, rows)
	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			c := board.NewCell(x, y, columns, rows)

			*c.Alive = rand.Float64() < 0.15
			c.AliveNext = c.Alive

			cells[x] = append(cells[x], c)
		}
	}
	board.Cells = cells
}

func (board *Board) Draw(cells [][]*Cell, tanks []Tank) {
	gl.UseProgram(board.Program)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	//DrawCell(cells, Window, Program)
	DrawTank(tanks, board.Window, board.Program)

	glfw.PollEvents()
	board.Window.SwapBuffers()
}

func DrawCell(cells [][]*Cell, window *glfw.Window, program uint32) {
	for x := range cells {
		for _, c := range cells[x] {
			c.Draw()
		}
	}
}

func DrawTank(tanks []Tank, window *glfw.Window, program uint32) {
	for _, tank := range tanks {
		tank.Draw()
	}
}
