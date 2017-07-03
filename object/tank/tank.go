package tank

import "../cell"

type Tank struct {
	Cells []cell.Cell

	X     int
	Y     int
	Lives int

	bullet *Bullet
}

type Bullet struct {
	point *cell.Cell

	X int
	Y int
}

func New(boardRows, boardColumns int) *Tank {
	tank := new(Tank)

	tank.bullet = nil

	tank.Cells = make([]cell.Cell, 12)

	tank.Cells[0] = *cell.NewCell(0, 0, boardColumns, boardRows)
	tank.Cells[1] = *cell.NewCell(1, 0, boardColumns, boardRows)
	tank.Cells[2] = *cell.NewCell(2, 0, boardColumns, boardRows)
	tank.Cells[3] = *cell.NewCell(3, 0, boardColumns, boardRows)
	tank.Cells[4] = *cell.NewCell(4, 0, boardColumns, boardRows)
	tank.Cells[5] = *cell.NewCell(5, 0, boardColumns, boardRows)
	tank.Cells[6] = *cell.NewCell(6, 0, boardColumns, boardRows)
	tank.Cells[7] = *cell.NewCell(7, 0, boardColumns, boardRows)
	tank.Cells[8] = *cell.NewCell(8, 0, boardColumns, boardRows)
	tank.Cells[9] = *cell.NewCell(9, 0, boardColumns, boardRows)
	tank.Cells[10] = *cell.NewCell(10, 0, boardColumns, boardRows)
	tank.Cells[11] = *cell.NewCell(11, 0, boardColumns, boardRows)

	tank.SetAlive(true)
	if tank == nil {
		panic("tank is nil")
	}

	if tank == nil {
		panic("")
	}

	return tank
}

func (tank *Tank) SetAlive(alive bool) {
	for _, c := range tank.Cells {
		*c.Alive = alive
	}
}

func (tank *Tank) Move(direction int) {

}

func (tank *Tank) Rotate(direction int) {

}

func (tank *Tank) fire() {

}

func (tank *Tank) Draw() {
	for _, c := range tank.Cells {
		c.Draw()
	}
}
