package main


type ListNode struct {
	Val int
	Next *ListNode
}

func main(){
	head := &ListNode{
		Val: -1,
		Next: &ListNode{
			Val: -7,
			Next: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: -4,
				},
			},
		},
	}
	res := detectCycle(head)
	if res != nil {
		println(res.Val)
	}
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func detectCycle(head *ListNode) *ListNode {
    tmp := head
    nodeNap := make(map[int]*ListNode,)

	count := 0
    for tmp != nil && tmp.Next != nil{
        nodeNap[tmp.Val] = tmp
        tmp = tmp.Next
		count++

        if index,ok := nodeNap[tmp.Val]; ok {
            return index
        }
		println(count)
    }
    return nil
}
   
   
   
   
   
   
   
   
   
   
   
   
   
   