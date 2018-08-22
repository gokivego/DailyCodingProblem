package problem20

import ("fmt"
        "testing"
)

func makeLLFromArray(arr []int) *ListNode {
    if len(arr) == 0 {return nil}
    var head *ListNode = &ListNode{}
    temp := head
    for _ , data := range arr {
        temp.next = &ListNode{Val:data, next:nil}
        temp = temp.next
    }
    return head.next
}

func joinLL(headA, headB *ListNode, joinIndex int) {
    // We join headA's joinIndex number to the tail of headB
    // A = 1->2->3->4->5 , 4 will be the intersection point
    // We pass the joinIndex and the user should pass a valid index
    lenA := getLen(headA)
    if joinIndex >= lenA {return}
    for (headB.next != nil) {
        // We reach the end of the B linkedlist with this
        headB = headB.next
    }
    for (joinIndex>0) {
        headA = headA.next
        joinIndex--
    }
    headB.next = headA
}

func printLL(head *ListNode) {
    for (head != nil) {
        fmt.Printf("%d->",head.Val)
        head = head.next
    }
    fmt.Println("nil")
}

func TestProblem20(t *testing.T) {
    headA := makeLLFromArray([]int{1,2,3,4,5})
    headB := makeLLFromArray([]int{5,7,8,9,10})
    joinLL(headA, headB, 3)
    fmt.Printf("LinkedListA:")
    printLL(headA)
    fmt.Printf("LinkedListB:")
    printLL(headB)
    fmt.Printf("The intersection node of the two linkedlists is %v\n",problem20(headA, headB))
}
