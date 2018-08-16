package problem9

func max(a,b int) int {
	if a>b {return a}
	return b
}

// Solution in O(N) time and constant space
func problem9(sum_arr []int) int {
	if len(sum_arr) == 0 {return 0}
	var dp = make([]int, 3)
	for _, val := range sum_arr {
		dp[2] = max(val+dp[0], dp[1])
		dp[0] = dp[1]
		dp[1] = dp[2]
	}
	return dp[2]
}

func main() {}
