package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := []int{}
	p := l1
	q := l2
	sum := 0
	for p != nil || q != nil {
		if p != nil {
			sum += p.Val
			p = p.Next
		}
		if q != nil {
			sum += q.Val
			q = q.Next
		}
		res = append(res, sum%10)
		sum = sum / 10
	}

	if sum > 0 {
		res = append(res, sum)
	}

	return toList(res)
}

func toList(nums []int) *ListNode {
	var res, p *ListNode
	for _, num := range nums {
		q := &ListNode{}
		q.Next = nil
		q.Val = num
		if p != nil {
			p.Next = q
		}
		if res == nil {
			res = q
		}
		p = q
	}
	return res
}

func toArray(l *ListNode) []int {
	p := l
	res := []int{}
	for p != nil {
		res = append(res, p.Val)
		p = p.Next
	}

	return res
}

func main() {
	l1 := []int{2, 4, 3}
	l2 := []int{5, 6, 4}

	fmt.Println(toArray(toList(l1)))
	fmt.Println(toArray(addTwoNumbers(toList(l1), toList(l2))))
}
