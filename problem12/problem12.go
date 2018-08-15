package problem12

func problem12(N int, possibleWays []int) int {
	var dp = make([]int,N+1)
	dp[0] = 1
	i := 1
	for (i<N+1) {
		for _,val := range possibleWays {
			if (i-val >= 0) {
				dp[i] += dp[i-val]
			}
		}
		i++
	}
	return dp[N]
}

func main() {
}
