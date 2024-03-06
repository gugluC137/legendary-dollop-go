package customsort

import "github.com/gugluC137/legendary-dollop-go/code/utils"

type Bubble struct{}

func (*Bubble) Sort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	for sorted := false; !sorted; {
		sorted = true
		for i := 1; i < len(arr); i++ {
			if arr[i] < arr[i-1] {
				utils.SwapValuesIntArray(&arr, i, i-1)
				sorted = false
			}
		}
	}

	return arr
}
