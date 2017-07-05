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

	Rows    = 100
	Columns = 100
)

type Board struct {
	Cells       [][]*Cell
	EnemyTanks  []Tank
	PlayerTanks []Tank

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
	return Rows
}
func (board *Board) GetBoardColumns() int {
	return Columns
}

func (board *Board) DestroyTank(destroyId int64) {
	tanks := make([]Tank, 0)
	for _, tank := range board.EnemyTanks {
		if tank.Id != destroyId {
			tanks = append(tanks, tank)
		}
	}
	board.EnemyTanks = tanks
}

func (board *Board) makeTanks(num_of_tank int, isEnemy bool) []Tank {
	tanks := make([]Tank, 0)
	for i := 0; i < num_of_tank; i++ {
		tanks = append(tanks, *board.New(board.GetBoardRows(), board.GetBoardColumns(), isEnemy))
	}
	return tanks
}

func (board *Board) MakeEnemyTanks(num_of_tank int) {
	board.EnemyTanks = board.makeTanks(num_of_tank, true)
}

func (board *Board) MakePlayerTanks(num_of_tank int) {
	board.PlayerTanks = board.makeTanks(num_of_tank, false)
}

func (board *Board) MakeCells() {
	rand.Seed(time.Now().UnixNano())

	cells := make([][]*Cell, Rows, Rows)
	for x := 0; x < Rows; x++ {
		for y := 0; y < Columns; y++ {
			c := board.NewCell(x, y, Columns, Rows)

			*c.Alive = rand.Float64() < 0.15
			c.AliveNext = c.Alive

			cells[x] = append(cells[x], c)
		}
	}
	board.Cells = cells
}

func (board *Board) Draw() {
	gl.UseProgram(board.Program)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	//DrawCell(cells, board.Window, board.Program)
	DrawTank(board.EnemyTanks, board.Window, board.Program)
	DrawTank(board.PlayerTanks, board.Window, board.Program)

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
		for i := 0; i < len(tank.Bullet); i++ {
			tank.Bullet[i].Draw()
		}
	}
}
