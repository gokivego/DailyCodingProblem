/**
 * Definition of a singly-linked list.
 * type ListNode struct {
 *     Val int;
 *     next *ListNode;
 * };
**/

package problem20

type ListNode struct {
    Val int;
    next *ListNode;
};

func getLen(head *ListNode) int {
    length := 0;
    for (head != nil) {
        head = head.next
        length++
    }
    return length
}

func problem20Helper(headA, headB *ListNode, lenA, lenB int) *ListNode {
    // We receive input so that headB is longer
    tempA := headA
    tempB := headB
    diff := lenB - lenA
    // Advance the longer linked list diff moves ahead
    for (diff >0) {
        tempB = tempB.next
        diff--
    }
    // Now both the lists are aligned in length
    // We need to check if they intersect
    // They intersect if the value from that point till end is the same
    for (tempA != nil && tempB != nil) {
        if tempA == tempB {
            return tempA
        } 
        tempA = tempA.next
        tempB = tempB.next
    }
    return nil
}

func problem20(headA, headB *ListNode) *ListNode {
    lenA := getLen(headA)
    lenB := getLen(headB)
    if lenA > lenB {
        return problem20Helper(headB, headA , lenB, lenA)
    }
    return problem20Helper(headA, headB, lenA, lenB)
}

func main() {}
