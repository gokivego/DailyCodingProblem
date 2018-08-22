package problem19

import ("fmt"
		"testing"
)

func TestProblem19(t *testing.T) {
	costs := [][]int{{1,2,4},{3,2,1},{2,4,5}}
	/*
	for _, rows := range costs {
		for _, cols := range rows {
			fmt.Printf("%d\t",cols)
		}
		fmt.Printf("\n")
	}
	*/
	fmt.Printf("The minimum cost to paint %d houses using costs %v is %d\n",len(costs),costs,problem19(costs))
}