package board

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"math/rand"
	"time"

	"../cell"
	"../tank"
)

const (
	WindowWidth  = 250
	WindowHeight = 250

	rows    = 100
	columns = 100
)

type Board struct {
	Cells [][]cell.Cell
	tanks []tank.Tank
}

func (board *Board)GetBoardWidth() int {
	return WindowWidth
}

func (board *Board)GetBoardHeight() int {
	return WindowHeight
}
func (board *Board)GetBoardRows() int {
	return rows
}
func (board *Board)GetBoardColumns() int {
	return columns
}

func MakeCells() [][]*cell.Cell {
	rand.Seed(time.Now().UnixNano())

	cells := make([][]*cell.Cell, rows, rows)
	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			c := cell.NewCell(x, y, columns, rows)

			*c.Alive = rand.Float64() < 0.15
			c.AliveNext = c.Alive

			cells[x] = append(cells[x], c)
		}
	}

	return cells
}

func Draw(window *glfw.Window, program uint32, cells [][]*cell.Cell, tank *tank.Tank) {
	gl.UseProgram(program)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	//DrawCell(cells, window, program)
	DrawTank(tank, window, program)

	glfw.PollEvents()
	window.SwapBuffers()
}

func DrawCell(cells [][]*cell.Cell, window *glfw.Window, program uint32) {
	for x := range cells {
		for _, c := range cells[x] {
			c.Draw()
		}
	}
}

func DrawTank(tank *tank.Tank, window *glfw.Window, program uint32) {
	tank.Draw()
}
