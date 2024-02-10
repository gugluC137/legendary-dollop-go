package code

func Sort012(arrP *[]int) {
	lengthOfArr := len(*arrP)
	pos := [3]int{0, 0, lengthOfArr - 1}

	for pos[1] <= pos[2] {
		switch (*arrP)[pos[1]] {
		case 0:
			swap(arrP, pos[0], pos[1])
			pos[0]++
			pos[1]++
		case 1:
			pos[1]++
		case 2:
			swap(arrP, pos[1], pos[2])
			pos[2]--
		}
	}
}

func swap(ar *[]int, i, j int) {
	temp := (*ar)[i]
	(*ar)[i] = (*ar)[j]
	(*ar)[j] = temp
}
