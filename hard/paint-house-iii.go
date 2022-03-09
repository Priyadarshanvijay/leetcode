// https://leetcode.com/problems/paint-house-iii
package hard

type Memo [101][101][21]int

func (this *Memo) initialize(value int) {
	for i := range *this {
		for j := range (*this)[i] {
			for k := range (*this)[i][j] {
				(*this)[i][j][k] = -1
			}
		}
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func solve(curIndex int, remainingNeighborhoods int, prevPaint int, houses *[]int, cost *[][]int, DP *Memo) int {
	impossible := 1000000000
	if remainingNeighborhoods < 0 {
		// We created more neighborhoods than required
		return impossible
	}
	if curIndex == len(*houses) {
		// Painted all the houses
		if remainingNeighborhoods == 0 {
			// Created the exact amount of neighborhoods!!
			return 0
		}
		// Created less neighborhoods than required
		return impossible
	}
	if (*DP)[curIndex][remainingNeighborhoods][prevPaint] != -1 {
		// We have already solved this subproblem
		return (*DP)[curIndex][remainingNeighborhoods][prevPaint]
	}
	if (*houses)[curIndex] != 0 {
		// Current house was already painted
		neighborhoodsToCreate := remainingNeighborhoods
		if prevPaint != (*houses)[curIndex] {
			// Previous paint that we did does not matches the current house, we created one more neighborhood
			neighborhoodsToCreate -= 1
		}
		(*DP)[curIndex][remainingNeighborhoods][prevPaint] = solve(curIndex+1, neighborhoodsToCreate, (*houses)[curIndex], houses, cost, DP)
		return (*DP)[curIndex][remainingNeighborhoods][prevPaint]
	} else {
		// Current House Isn't painted yet
		(*DP)[curIndex][remainingNeighborhoods][prevPaint] = impossible // Initialize current cost with max possible value
		for i := 0; i < len((*cost)[curIndex]); i++ {
			// iterate over all paints
			neighborhoodsToCreate := remainingNeighborhoods
			if prevPaint != i+1 {
				// We're creating new neighborhood
				neighborhoodsToCreate -= 1
			}
			// Take minimum of current and previous cost
			(*DP)[curIndex][remainingNeighborhoods][prevPaint] = min((*DP)[curIndex][remainingNeighborhoods][prevPaint], (*cost)[curIndex][i]+solve(curIndex+1, neighborhoodsToCreate, i+1, houses, cost, DP))
		}
		return (*DP)[curIndex][remainingNeighborhoods][prevPaint]
	}
}

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	impossible := 1000000000
	var DP Memo
	(&DP).initialize(-1)
	ans := solve(0, target, 0, &houses, &cost, &DP)
	if ans == impossible {
		return -1
	}
	return ans
}
