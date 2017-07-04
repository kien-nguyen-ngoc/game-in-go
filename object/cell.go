package object

import (
	"github.com/go-gl/gl/v2.1/gl"

	"../util"
)

var (
	Square = []float32{
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,

		-0.5, 0.5, 0,
		0.5, 0.5, 0,
		0.5, -0.5, 0,
	}
)

type Cell struct {
	GameBoard *Board

	Drawable uint32

	Alive     *bool
	AliveNext *bool

	X int
	Y int
}

func (c *Cell) Draw() {
	if !*c.Alive {
		return
	}

	gl.BindVertexArray(c.Drawable)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(Square)/3))
}

func (board *Board) NewCell(x, y, columns, rows int) *Cell {
	points := make([]float32, len(Square), len(Square))
	copy(points, Square)

	for i := 0; i < len(points); i++ {
		var position float32
		var size float32
		switch i % 3 {
		case 0:
			size = 1.0 / float32(columns)
			position = float32(x) * size
		case 1:
			size = 1.0 / float32(rows)
			position = float32(y) * size
		default:
			continue
		}

		if points[i] < 0 {
			points[i] = (position * 2) - 1
		} else {
			points[i] = ((position + size) * 2) - 1
		}
	}

	return &Cell{
		GameBoard: board,
		Drawable:  util.MakeVao(points),

		X: x,
		Y: y,

		Alive:     new(bool),
		AliveNext: new(bool),
	}
}
