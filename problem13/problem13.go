package problem13

func max(a,b int) int {
	if a >b {return a}
	return b
}

func problem13(s string, k int) int {
	if len(s) <= k {
		return len(s)
	}
	counter := k
	stringLen := 0
	start := 0
	end := 0
	var Map = make(map[rune]int)
	var r = []rune(s)
	for end < len(r) {
		//fmt.Printf("1:%d\t",end)
		if _, ok := Map[r[end]]; !ok {
			counter--
		}
		Map[r[end]]++
		end++
		//fmt.Printf("2:%d\t",end)
		for (counter<0) {
			if val := Map[r[start]]; val == 1 {
				delete(Map, r[start])
				counter++
			} else {
				Map[r[start]]--
			}
			start++
		}
		stringLen = max(stringLen, end-start)
		//fmt.Printf("3:%d\n",end)
	}
	return stringLen
}

func main() {

}
