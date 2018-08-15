package problem13

import ("fmt"
		"testing"
)

func TestProblem13(t *testing.T) {
	kVals := []int{1,2,3,4}
	s := []string{"abcba","abcbbbca"}
	for _, k := range kVals {
			fmt.Printf("K : %d\n",k)
		for _, val := range s {
			fmt.Printf("The length of longest string in %s with utmost %d distinct characters is : %d\n",val,k,problem13(val,k))
		}
	}
}
