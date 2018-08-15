// Test module for Problem12 

package problem12

import ("fmt"
		"testing")

func TestProblem12(t *testing.T) {
	possibleWays := []int{1,2}
	arr := []int{5,10,15,20,25}
	for _, val := range arr {
		fmt.Printf("The number of ways of getting to %d is %d ways\n",val,problem12(val,possibleWays))
	}
}
