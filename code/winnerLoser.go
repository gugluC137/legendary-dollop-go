package code

import (
	"slices"
)

func findWinners(matches [][]int) [][]int {
	mapOfTeams := make(map[int]int)

	for _, v := range matches {

		if _, found0 := mapOfTeams[v[0]]; !found0 {
			mapOfTeams[v[0]] = 0
		}

		if _, found1 := mapOfTeams[v[1]]; found1 {
			mapOfTeams[v[1]]++
		} else {
			mapOfTeams[v[1]] = 1
		}
	}

	ans := make([][]int, 2)
	ans[0] = make([]int, 0)
	ans[1] = make([]int, 0)

	for k, v := range mapOfTeams {
		if v == 0 {
			ans[0] = append(ans[0], k)
		} else if v == 1 {
			ans[1] = append(ans[1], k)
		}
	}

	slices.Sort(ans[0])
	slices.Sort(ans[1])

	return ans

}
