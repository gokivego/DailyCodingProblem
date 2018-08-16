package problem9

import ("fmt"
		"testing"
)

func TestProblem9(t *testing.T) {
	sum_arr := [][]int{{2,4,6,2,5}}
	for _, val := range sum_arr {
		fmt.Printf("Maximum sum of non-adjacent numbers in the array %v is %d\n",val, problem9(val))
	}
}
