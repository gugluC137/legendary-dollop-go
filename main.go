package main

import (
	"fmt"

	"github.com/gugluC137/legendary-dollop-go/code"
	sort "github.com/gugluC137/legendary-dollop-go/code/customsort"
)

func main() {
	fmt.Println("Hello, World!!")

	code.PrintN(3)

	matches := [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}

	ans := code.FindWinners(matches)
	fmt.Println(ans)

	arr := []int{3, 4, 2, 1, 7, 0}
	fmt.Println(arr)
	arr = sort.SortFunc(arr, new(sort.Bubble))
	fmt.Println(arr)

}
