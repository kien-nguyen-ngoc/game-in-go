package object

import "time"

type Barrier struct {
	GameBoard *Board
	Cells     []Cell
	Id        int64
	Remain    *int
}

func (board *Board) NewBarrier(height, width, x, y int) *Barrier {
	barrier := new(Barrier)
	barrier.Id = time.Now().UnixNano()
	barrier.Remain = new(int)
	barrier.GameBoard = board
	barrier.Cells = make([]Cell, 0)

	lossX := x - (board.GetBoardColumns() - height)
	lossY := y - (board.GetBoardRows() - width)
	if lossX > 0 {
		x = x - lossX
	}
	if lossY > 0 {
		y = y - lossY
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			cell := *board.NewCell(x+i, y+j, board.GetBoardColumns(), board.GetBoardRows())
			*cell.Alive = true
			barrier.Cells = append(barrier.Cells, cell)
		}
	}

	return barrier
}

func (barrier *Barrier) Draw() {
	game_board := barrier.GameBoard
	barrier.checkHits(game_board.EnemyTanks)
	barrier.checkHits(game_board.PlayerTanks)

	for _, c := range barrier.Cells {
		c.Draw()
	}
}

func (barrier *Barrier) remove() {
	game_board := barrier.GameBoard
	barriers := make([]Barrier, 0)
	for _, b := range game_board.Barriers {
		if barrier.Id != b.Id {
			barriers = append(barriers, b)
		}
	}
	game_board.Barriers = barriers
}

func (barrier *Barrier) checkHits(tanks []Tank) {
	for i := 0; i < len(tanks); i++ {
		tank := tanks[i]
		for j := 0; j < len(tank.Bullets); j++ {
			barrier.checkHit(&tank.Bullets[j])
		}
	}
}
func (barrier *Barrier) checkHit(bullet *Bullet) {
	for i := 0; i < len(barrier.Cells); i++ {
		cell := barrier.Cells[i]
		if *cell.Alive && cell.X == bullet.X && cell.Y == bullet.Y {
			*cell.Alive = false
			bullet.Remove()
			*barrier.Remain -= 1

			if *barrier.Remain == 0 {
				barrier.remove()
			}
		}
	}
}
