// https://leetcode.com/problems/cherry-pickup-ii/
package hard

type curPosition struct {
	x  int
	y  int
	mx int
	my int
}

type Memo [][][]int

func (this *curPosition) isValid() bool {
	if this.x >= this.mx {
		return false
	}
	if this.y >= this.my {
		return false
	}
	if this.x < 0 {
		return false
	}
	if this.y < 0 {
		return false
	}
	return true
}

func (this curPosition) goLeft() curPosition {
	this.x -= 1
	this.y += 1
	return this
}

func (this curPosition) goRight() curPosition {
	this.x += 1
	this.y += 1
	return this
}

func (this curPosition) goCenter() curPosition {
	this.y += 1
	return this
}

func (this *curPosition) isEqual(sec *curPosition) bool {
	return this.x == sec.x && this.y == sec.y
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func (this *Memo) curValue(r1 *curPosition, r2 *curPosition) int {
	return (*this)[r1.y][r1.x][r2.x]
}

func (this *Memo) setValue(r1 *curPosition, r2 *curPosition, v int) {
	(*this)[r1.y][r1.x][r2.x] = v
	return
}

func cherryPickupHelper(grid [][]int, dp *Memo, r1 curPosition, r2 curPosition) int {
	if !(&r1).isValid() || !(&r2).isValid() {
		return 0
	}
	curCherries := dp.curValue(&r1, &r2)
	if curCherries != -1 {
		return curCherries
	}
	curCherries = 0
	if (&r1).isEqual(&r2) {
		curCherries = grid[r1.y][r1.x]
	} else {
		curCherries = grid[r1.y][r1.x] + grid[r2.y][r2.x]
	}
	m := cherryPickupHelper(grid, dp, r1.goLeft(), r2.goLeft())
	m = max(m, cherryPickupHelper(grid, dp, r1.goLeft(), r2.goRight()))
	m = max(m, cherryPickupHelper(grid, dp, r1.goLeft(), r2.goCenter()))
	m = max(m, cherryPickupHelper(grid, dp, r1.goRight(), r2.goLeft()))
	m = max(m, cherryPickupHelper(grid, dp, r1.goRight(), r2.goRight()))
	m = max(m, cherryPickupHelper(grid, dp, r1.goRight(), r2.goCenter()))
	m = max(m, cherryPickupHelper(grid, dp, r1.goCenter(), r2.goLeft()))
	m = max(m, cherryPickupHelper(grid, dp, r1.goCenter(), r2.goRight()))
	m = max(m, cherryPickupHelper(grid, dp, r1.goCenter(), r2.goCenter()))
	curCherries += m
	dp.setValue(&r1, &r2, curCherries)
	return curCherries
}

func cherryPickup(grid [][]int) int {
	nr := len(grid)
	nc := len(grid[0])
	dp := make(Memo, nr)
	for i := range dp {
		dp[i] = make([][]int, nc)
		for j := range dp[i] {
			dp[i][j] = make([]int, nc)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	r1 := curPosition{x: 0, y: 0, mx: nc, my: nr}
	r2 := curPosition{x: (nc - 1), y: 0, mx: nc, my: nr}
	return cherryPickupHelper(grid, &dp, r1, r2)
}
