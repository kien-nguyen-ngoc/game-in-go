package play

import (
	"../board/cell"
)


// checkState determines the state of the Cell for the next tick of the game.
func CheckState(c *cell.Cell, cells [][]*cell.Cell) {
	c.Alive = c.AliveNext
	c.AliveNext = c.Alive

	liveCount := LiveNeighbors(c, cells)
	if *c.Alive {
		// 1. Any live Cell with fewer than two live neighbours dies, as if caused by underpopulation.
		if liveCount < 2 {
			*c.AliveNext = false
		}

		// 2. Any live Cell with two or three live neighbours lives on to the next generation.
		if liveCount == 2 || liveCount == 3 {
			*c.AliveNext = true
		}

		// 3. Any live Cell with more than three live neighbours dies, as if by overpopulation.
		if liveCount > 3 {
			*c.AliveNext = false
		}
	} else {
		// 4. Any dead Cell with exactly three live neighbours becomes a live Cell, as if by reproduction.
		if liveCount == 3 {
			*c.AliveNext = true
		}
	}
}

// liveNeighbors returns the number of live neighbors for a Cell.
func LiveNeighbors(c *cell.Cell, cells [][]*cell.Cell) int {
	var liveCount int
	add := func(x, y int) {
		// If we're at an edge, check the other side of the board.
		if x == len(cells) {
			x = 0
		} else if x == -1 {
			x = len(cells) - 1
		}
		if y == len(cells[x]) {
			y = 0
		} else if y == -1 {
			y = len(cells[x]) - 1
		}

		if *cells[x][y].Alive {
			liveCount++
		}
	}

	add(c.X-1, c.Y)   // To the left
	add(c.X+1, c.Y)   // To the right
	add(c.X, c.Y+1)   // up
	add(c.X, c.Y-1)   // down
	add(c.X-1, c.Y+1) // top-left
	add(c.X+1, c.Y+1) // top-right
	add(c.X-1, c.Y-1) // bottom-left
	add(c.X+1, c.Y-1) // bottom-right

	return liveCount
}
