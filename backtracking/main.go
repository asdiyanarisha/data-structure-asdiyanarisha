package main

import "fmt"

func findElementsWithSum(arr [10]int, combinations [19]int, size int, k int, addValue int, l int, m int) int {
	var num int = 0
	if addValue > k {
		return -1
	}

	fmt.Println("addValue", addValue) // debug
	fmt.Println("recur m", m)
	fmt.Println("comb ", combinations)
	fmt.Println("l ", l)

	if addValue == k {
		fmt.Println("Result:")
		num = num + 1
		var p int = 0
		for p = 0; p < m; p++ {
			fmt.Printf("%d,", arr[combinations[p]])
		}
		fmt.Println(" end result")
	}
	var i int
	for i = l; i < size; i++ {
		combinations[m] = l
		fmt.Println("i = ", i)
		findElementsWithSum(arr, combinations, size, k, addValue+arr[i], l, m+1)
		l = l + 1
	}
	return num
}

func main() {
	var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	var addedSum int = 18
	var combinations [19]int
	findElementsWithSum(arr, combinations, 10, addedSum, 0, 0, 0)
	//fmt.Println(check)//var check2 bool = findElement(arr,9)
	//fmt.Println(check2)
}
