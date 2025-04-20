package strivercp

import "fmt"

func bitland() int {
	x := 0

	var tc int
	fmt.Scan(&tc)

	for range tc {
		var statement string
		fmt.Scan(&statement)

		for _, val := range statement {
			if val == '+' {
				x++
				break
			}
			if val == '-' {
				x--
				break
			}
		}
	}

	return x
}
