package problem2

import ("testing"
        "reflect"
)

type testStruct struct {
    arr []int
    res []int
}

func TestProductExceptSelf(t *testing.T) {
    arrays := []testStruct {
            testStruct{[]int{1,2,3,4,5},[]int{120,60,40,30,24}},
            testStruct{[]int{3,2,1},[]int{2,3,6}},
    }

    for _, arr := range arrays {
        product := ProductExceptSelf(arr.arr)
        if !reflect.DeepEqual(product, arr.res) {
            t.Errorf("The result is supposed to be %v but got result %v", arr.res, product)
        }
    }
}
