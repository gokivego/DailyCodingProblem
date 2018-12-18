package problem2

func LeftProduct(arr []int) []int {
    lp := make([]int, len(arr))
    lp[0] = 1;
    for i := 1; i < len(arr); i++ {
        lp[i] = lp[i-1]*arr[i-1];
    }
    return lp;
}

func ProductExceptSelf(arr []int) []int {
    lp := LeftProduct(arr)
    pes := make([]int, len(arr))
    rp := 1;
    for i := len(arr)-1; i>=0; i-- {
        pes[i] = rp*lp[i];
        rp *= arr[i];
    }
    return pes;
}
