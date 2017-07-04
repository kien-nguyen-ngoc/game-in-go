package object

import (
	"log"
	"math/rand"
	"time"
)

//var (
//	BOTTOM_0 = 0
//	BOTTOM_1 = 1
//	BOTTOM_2 = 2
//	MID_0    = 3
//	MID_1    = 4
//	MID_2    = 5
//	GUN_0    = 6
//	GUN_1    = 7
//	GUN_2    = 8
//)

type Tank struct {
	GameBoard *Board
	Id        int
	Cells     []Cell

	X     int
	Y     int
	Lives int

	Bullet  []Bullet
	IsEnemy bool

	BOTTOM_0 int
	BOTTOM_1 int
	BOTTOM_2 int
	MID_0    int
	MID_1    int
	MID_2    int
	GUN_0    int
	GUN_1    int
	GUN_2    int
}

type Bullet struct {
	point *Cell

	X int
	Y int
}

func (board *Board) New(boardRows, boardColumns int, Id int, isEnemy bool) *Tank {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	tank := new(Tank)
	tank.Id = Id
	tank.Bullet = nil
	tank.Lives = int(r1.Int31n(5)) + 1

	tank.BOTTOM_0 = 0
	tank.BOTTOM_1 = 1
	tank.BOTTOM_2 = 2
	tank.MID_0 = 3
	tank.MID_1 = 4
	tank.MID_2 = 5
	tank.GUN_0 = 6
	tank.GUN_1 = 7
	tank.GUN_2 = 8

	tank.Cells = make([]Cell, 9)

	tank.Cells[tank.BOTTOM_0] = *board.NewCell(0, 0, boardColumns, boardRows)
	tank.Cells[tank.BOTTOM_1] = *board.NewCell(1, 0, boardColumns, boardRows) // player or enemy
	tank.Cells[tank.BOTTOM_2] = *board.NewCell(2, 0, boardColumns, boardRows)
	tank.Cells[tank.MID_0] = *board.NewCell(0, 1, boardColumns, boardRows)
	tank.Cells[tank.MID_1] = *board.NewCell(1, 1, boardColumns, boardRows)
	tank.Cells[tank.MID_2] = *board.NewCell(2, 1, boardColumns, boardRows)
	tank.Cells[tank.GUN_0] = *board.NewCell(0, 2, boardColumns, boardRows) // hidden
	tank.Cells[tank.GUN_1] = *board.NewCell(1, 2, boardColumns, boardRows)
	tank.Cells[tank.GUN_2] = *board.NewCell(2, 2, boardColumns, boardRows) // hidden

	tank.SetAlive(true, tank.IsEnemy)

	return tank
}

func (tank *Tank) SetAlive(alive bool, isEnemy bool) {
	for _, c := range tank.Cells {
		*c.Alive = alive
	}
	*tank.Cells[tank.GUN_0].Alive = false
	*tank.Cells[tank.GUN_2].Alive = false

	if isEnemy {
		*tank.Cells[tank.BOTTOM_1].Alive = false
	}

}

func (tank *Tank) MoveForward() {
	isEnemy := tank.IsEnemy
	gun_point := tank.Cells[tank.GUN_1]
	bottom_point := tank.Cells[tank.BOTTOM_1]
	moveX_step := (gun_point.X - bottom_point.X) / 3
	moveY_step := (gun_point.Y - bottom_point.Y) / 3

	log.Print(gun_point.X)

	for i := 0; i < len(tank.Cells); i++ {
		tank.Cells[i] = *tank.GameBoard.NewCell(tank.Cells[i].X+moveX_step, tank.Cells[i].Y+moveY_step,
			tank.GameBoard.GetBoardColumns(), tank.GameBoard.GetBoardRows())
	}
	tank.SetAlive(tank.Lives >= 0, isEnemy)
}

func (tank *Tank) RotateRight() {
	tmp_BOTTOM_0 := tank.BOTTOM_0
	tmp_BOTTOM_1 := tank.BOTTOM_1
	tmp_BOTTOM_2 := tank.BOTTOM_2
	tmp_MID_0 := tank.MID_0
	tmp_MID_1 := tank.MID_1
	tmp_MID_2 := tank.MID_2
	tmp_GUN_0 := tank.GUN_0
	tmp_GUN_1 := tank.GUN_1
	tmp_GUN_2 := tank.GUN_2

	tank.BOTTOM_0 = tmp_GUN_0
	tank.BOTTOM_1 = tmp_MID_0
	tank.BOTTOM_2 = tmp_BOTTOM_0
	tank.MID_0 = tmp_GUN_1
	tank.MID_1 = tmp_MID_1
	tank.MID_2 = tmp_BOTTOM_1
	tank.GUN_0 = tmp_GUN_2
	tank.GUN_1 = tmp_MID_2
	tank.GUN_2 = tmp_BOTTOM_2

	tank.SetAlive(tank.Lives >= 0, tank.IsEnemy)
}

func (tank *Tank) RotateLeft() {
	tmp_BOTTOM_0 := tank.BOTTOM_0
	tmp_BOTTOM_1 := tank.BOTTOM_1
	tmp_BOTTOM_2 := tank.BOTTOM_2
	tmp_MID_0 := tank.MID_0
	tmp_MID_1 := tank.MID_1
	tmp_MID_2 := tank.MID_2
	tmp_GUN_0 := tank.GUN_0
	tmp_GUN_1 := tank.GUN_1
	tmp_GUN_2 := tank.GUN_2

	tank.BOTTOM_0 = tmp_BOTTOM_2
	tank.BOTTOM_1 = tmp_MID_2
	tank.BOTTOM_2 = tmp_GUN_2
	tank.MID_0 = tmp_BOTTOM_1
	tank.MID_1 = tmp_MID_1
	tank.MID_2 = tmp_GUN_1
	tank.GUN_0 = tmp_BOTTOM_0
	tank.GUN_1 = tmp_MID_0
	tank.GUN_2 = tmp_GUN_0

	tank.SetAlive(tank.Lives >= 0, tank.IsEnemy)
}

func (tank *Tank) Rotate180() {
	tank.RotateRight()
	tank.RotateRight()
}

func (tank *Tank) fire() {
	bullet := new(Bullet)
	bullet.X = tank.Cells[tank.GUN_1].X
	bullet.Y = tank.Cells[tank.GUN_1].Y
	bullet.point = new(Cell)
	*bullet.point = tank.Cells[tank.GUN_1]

	tank.Bullet = append(tank.Bullet, *bullet)
}

func (tank *Tank) Draw() {
	for _, c := range tank.Cells {
		c.Draw()
	}
}
