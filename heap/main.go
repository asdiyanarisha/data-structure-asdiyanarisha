package main

import (
	"container/heap"
	"fmt"
)

// integerHeap a type
type IntegerHeap []int

// IntegerHeap method - gets the length of integerHeap
func (iHeap IntegerHeap) Len() int { return len(iHeap) }

func (iHeap IntegerHeap) Less(i, j int) bool {
	fmt.Println("Call Less", iHeap[i], iHeap[j])
	return iHeap[i] < iHeap[j]
}

func (iHeap IntegerHeap) Swap(i, j int) {
	fmt.Println("Call Swap", iHeap[i], iHeap[j])
	iHeap[i], iHeap[j] = iHeap[j], iHeap[i]
}

func (iHeap *IntegerHeap) Push(heapIntF interface{}) {
	fmt.Println("Call Push", heapIntF)
	*iHeap = append(*iHeap, heapIntF.(int))
}

func (iHeap *IntegerHeap) Pop() interface{} {
	fmt.Println("Call Pop")
	var (
		n    int
		x1   int
		prev = *iHeap
	)

	n = len(prev)
	x1 = prev[n-1]
	*iHeap = prev[0 : n-1]

	return x1
}

func main() {
	fmt.Println("hello world")

	var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5}

	heap.Init(intHeap)

	heap.Push(intHeap, 2)

	fmt.Printf("minimum: %d\n", (*intHeap)[0])
	fmt.Printf("maximam: %d\n", (*intHeap)[intHeap.Len()-1])

	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}

	fmt.Println("finish", intHeap)
}
