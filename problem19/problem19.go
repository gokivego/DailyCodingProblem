package problem19

import ("math")

func problem19(costs [][]int) int {
	N := len(costs)
	if N == 0 {return 0}
	K := len(costs[0])
	preMin := 0
	preSecondMin := 0
	preMinIndex := -1
	for i:=0;i<N;i++ {
		currentMin := math.MaxInt32
		currentSecondMin := math.MaxInt32
		currentMinIndex := 0
		for j:=0;j<K;j++ {
			if (preMinIndex == j) {
				costs[i][j] += preSecondMin
			} else {
				costs[i][j] += preMin
			}
			if (currentMin > costs[i][j]) {
				currentSecondMin = currentMin
				currentMin = costs[i][j]
				currentMinIndex = j
			} else if (currentSecondMin > costs[i][j]) {
				currentSecondMin = costs[i][j]
			}
		}
		preMin = currentMin
		preSecondMin = currentSecondMin
		preMinIndex = currentMinIndex
	}
	result := math.MaxInt32
	for j:=0;j<K;j++ {
		if costs[K-1][j] < result {
			result = costs[K-1][j]
		}
	}
	return result
}

func main() {} 