package problem18

import ("fmt"
        "testing"
)

func TestProblem18(t *testing.T) {
    k := 3
    arrays := [][]int{{1,2,3,4,5,4,3,2,1,1,1},{-4,2,-5,1,-1,6},{1,3,-1,-3,5,3,6,7}}
    for _, array := range arrays {
        fmt.Printf("The sliding window maximums in the array %v for a k: %d are %v\n",array,k,maxInSlidingWindow(array,k))
    }
}
