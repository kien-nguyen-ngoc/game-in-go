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
	Barriers    []Barrier

	//Window  *glfw.Window
	//Program uint32
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
		tanks = append(tanks, *board.NewTank(board.GetBoardRows(), board.GetBoardColumns(), isEnemy))
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

func (board *Board) MakeBarrier(num_of_barrier int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	barriers := make([]Barrier, 0)
	for i := 0; i < num_of_barrier; i++ {
		x := int(r1.Int31n(100))
		y := int(r1.Int31n(100))
		height := int(r1.Int31n(10))
		width := int(r1.Int31n(20))
		barrier := board.NewBarrier(height, width, x, y)
		barriers = append(barriers, *barrier)
	}
	board.Barriers = barriers
}

func (board *Board) Draw(Window  *glfw.Window, Program uint32) {
	gl.UseProgram(Program)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	board.drawTanks()
	board.drawBarriers()

	glfw.PollEvents()
	Window.SwapBuffers()
}

func (board *Board) drawTanks() {
	for _, tank := range append(board.EnemyTanks, board.PlayerTanks...) {
		tank.Draw()
		for i := 0; i < len(tank.Bullets); i++ {
			tank.Bullets[i].Draw()
		}
	}
}

func (boad *Board) drawBarriers() {
	for _, barrier := range boad.Barriers {
		barrier.Draw()
	}
}


func (board *Board)OnKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	}
	board.ControlTank(w,key,scancode,action,mods)
}