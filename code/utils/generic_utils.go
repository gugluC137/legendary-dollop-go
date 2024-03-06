package utils

func SwapValuesIntArray(arr *[]int, idx1, idx2 int) {
	temp := (*arr)[idx1]
	(*arr)[idx1] = (*arr)[idx2]
	(*arr)[idx2] = temp
}
