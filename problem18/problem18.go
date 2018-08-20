package problem18

//import ("fmt")

func max(a,b int) int {
    if a>b {return a}
    return b
}

func maxInSlidingWindow(arr []int,k int) []int {
    var maxValInSlidingWindow []int
    // If the length of the array is less than k then we return the max of the array
    if len(arr) < k {
        for _, val := range arr {
            if len(maxValInSlidingWindow) == 0 {
                maxValInSlidingWindow = append(maxValInSlidingWindow,val)
            } else {
                maxValInSlidingWindow[0] = max(maxValInSlidingWindow[0], val)
            }
        }
        //fmt.Println(maxValInSlidingWindow)
        return maxValInSlidingWindow
    }

    // Use a window to hold the indices.
    arrSize := len(arr)
    var window []int
    // We add the appropriate indices to the first window
    // for every new element we remove all the elements from the back of
    // the window that are less than the current element
    // We also check if the first element of the window falls outside
    // the window range
    for i:=0;i<k;i++ {
        for (len(window) > 0 && arr[window[len(window)-1]] <= arr[i]) {
            //fmt.Println(window)
            window = window[:len(window)-1]
        }
        window = append(window, i)
    }
    maxValInSlidingWindow = append(maxValInSlidingWindow,arr[window[0]])

    //fmt.Println(window)
    
    for i:=1;i<arrSize-k+1;i++ {
        for (len(window)>0 && arr[window[len(window)-1]] <= arr[i+k-1]) {
            //fmt.Println(window)
            window = window[:len(window)-1]
            // If the first element of window, index is less than i remove it
        }
        window = append(window, i+k-1)
        //fmt.Println(window)
        if window[0] < i {
            window = window[1:]
        }
        //fmt.Println(window)
        maxValInSlidingWindow = append(maxValInSlidingWindow,arr[window[0]])
    }
    //fmt.Println(maxValInSlidingWindow)
    return maxValInSlidingWindow
}

func main() {
}


