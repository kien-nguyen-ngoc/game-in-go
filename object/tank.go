package object

import (
	"math/rand"
	"time"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Tank struct {
	GameBoard *Board
	Id        int64
	Cells     []Cell

	Lives *int

	Bullets     []Bullet
	BulletLimit int

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
	Id   int64
	Tank *Tank
	Cell *Cell

	X       int
	Y       int
	XDirect int
	YDirect int
}

func (board *Board) NewTank(boardRows, boardColumns int, isEnemy bool) *Tank {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	tank := new(Tank)
	tank.Id = time.Now().UnixNano()
	tank.Bullets = nil
	tank.Lives = new(int)
	*tank.Lives = int(r1.Int31n(5)) + 1
	tank.IsEnemy = isEnemy
	tank.BulletLimit = 5
	tank.GameBoard = board

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

func (tank *Tank) MoveToPosition(x, y int) {
	for i := 0; i < len(tank.Cells); i++ {
		tank.Cells[i] = *tank.GameBoard.NewCell(x-tank.Cells[i].X, y-tank.Cells[i].Y,
			tank.GameBoard.GetBoardColumns(), tank.GameBoard.GetBoardRows())
	}
	tank.SetAlive(*tank.Lives >= 0, tank.IsEnemy)
}

func (tank *Tank) MoveForward() {
	gun_point := tank.Cells[tank.GUN_1]
	bottom_point := tank.Cells[tank.BOTTOM_1]
	moveX_step := (gun_point.X - bottom_point.X) / (len(tank.Cells)/3 - 1)
	moveY_step := (gun_point.Y - bottom_point.Y) / (len(tank.Cells)/3 - 1)

	if gun_point.X == tank.GameBoard.GetBoardColumns()-1 || (gun_point.X == 0 && moveX_step < 0) {
		moveX_step = 0
	}
	if gun_point.Y == tank.GameBoard.GetBoardRows()-1 || (gun_point.Y == 0 && moveY_step < 0) {
		moveY_step = 0
	}

	for i := 0; i < len(tank.Cells); i++ {
		tank.Cells[i] = *tank.GameBoard.NewCell(tank.Cells[i].X+moveX_step, tank.Cells[i].Y+moveY_step,
			tank.GameBoard.GetBoardColumns(), tank.GameBoard.GetBoardRows())
	}
	tank.SetAlive(*tank.Lives >= 0, tank.IsEnemy)
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

	tank.SetAlive(*tank.Lives >= 0, tank.IsEnemy)
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

	tank.SetAlive(*tank.Lives >= 0, tank.IsEnemy)
}

func (tank *Tank) Rotate180() {
	tank.RotateRight()
	tank.RotateRight()
}

func (tank *Tank) Fire() {
	if len(tank.Bullets) > tank.BulletLimit {
		return
	}
	game_board := tank.GameBoard
	bullet := new(Bullet)
	bullet.X = tank.Cells[tank.GUN_1].X
	bullet.Y = tank.Cells[tank.GUN_1].Y
	bullet.Cell = game_board.NewCell(bullet.X, bullet.Y, game_board.GetBoardColumns(), game_board.GetBoardRows())
	*bullet.Cell.Alive = true
	*bullet.Cell = tank.Cells[tank.GUN_1]
	bullet.XDirect = (tank.Cells[tank.GUN_1].X - tank.Cells[tank.BOTTOM_1].X) / (len(tank.Cells)/3 - 1)
	bullet.YDirect = (tank.Cells[tank.GUN_1].Y - tank.Cells[tank.BOTTOM_1].Y) / (len(tank.Cells)/3 - 1)
	bullet.Tank = tank
	bullet.Id = time.Now().UnixNano()

	tank.Bullets = append(tank.Bullets, *bullet)
}

func (tank *Tank) Draw() {
	for _, c := range tank.Cells {
		c.Draw()
	}
}

func (bullet *Bullet) Draw() {
	tank := bullet.Tank
	game_board := tank.GameBoard
	bullet.Cell.Draw()

	x := bullet.Cell.X
	y := bullet.Cell.Y

	x += bullet.XDirect
	y += bullet.YDirect

	if x < 0 || y < 0 || x >= game_board.GetBoardColumns() || y >= game_board.GetBoardRows() {
		bullet.Remove()
	} else {
		bullet.Cell = game_board.NewCell(x, y, game_board.GetBoardColumns(), game_board.GetBoardRows())
		*bullet.Cell.Alive = true
		bullet.X = x
		bullet.Y = y
	}

	if tank.IsEnemy {
		for i := 0; i < len(game_board.PlayerTanks); i++ {
			player := game_board.PlayerTanks[i]
			player.checkHit(bullet)
		}
	} else {
		for i := 0; i < len(game_board.EnemyTanks); i++ {
			enemy := game_board.EnemyTanks[i]
			enemy.checkHit(bullet)
		}
	}
}

func (tank *Tank) checkHit(bullet *Bullet) {
	for i := 0; i < len(tank.Cells); i++ {
		cell := tank.Cells[i]
		if cell.X == bullet.X && cell.Y == bullet.Y {
			*tank.Lives -= 1

			bullet.Remove()

			break
		}
	}
	if *tank.Lives < 0 {
		tank.SetAlive(false, tank.IsEnemy)
		tank.GameBoard.DestroyTank(tank.Id)
	}
}

func (bullet *Bullet) Remove() {
	bullets := make([]Bullet, 0)
	for _, b := range bullet.Tank.Bullets {
		if bullet.Id != b.Id {
			bullets = append(bullets, b)
		}
	}
	bullet.Tank.Bullets = bullets
}


func (board *Board)ControlTank(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyW && action == glfw.Press {
		board.PlayerTanks[0].MoveForward()
	} else if key == glfw.KeyS && action == glfw.Press {
		board.PlayerTanks[0].Rotate180()
	} else if key == glfw.KeyA && action == glfw.Press {
		board.PlayerTanks[0].RotateLeft()
	} else if key == glfw.KeyD && action == glfw.Press {
		board.PlayerTanks[0].RotateRight()
	}  else if key == glfw.KeySpace && action == glfw.Press {
		board.PlayerTanks[0].Fire()
	}
}
