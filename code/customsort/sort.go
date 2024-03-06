package customsort

type CustomSort interface {
	Sort(arr []int) []int
}

func SortFunc(arr []int, f CustomSort) []int {
	return f.Sort(arr)
}
