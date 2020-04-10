package main

import (
	"fmt"
	"strings"
)

type hint struct {
	Number     string
	RightPlace int8
	WrongPlace int8
}

func main() {
	numbers := [5]hint{
		{"289", 1, 0},
		{"215", 0, 1},
		{"942", 0, 2},
		{"738", 0, 0},
		{"784", 0, 1},
	}

	corrects := 0

	for n := 0; n <= 999; n++ {
		nc := padNumberWith3Zeros(n)
		corrects = 0

		for hintsCount := 0; hintsCount < len(numbers); hintsCount++ {
			if RightPlace(numbers[hintsCount], nc) == false {
				break
			}

			if WrongPlace(numbers[hintsCount], nc) == false {
				break
			}

			corrects++
		}

		if corrects == len(numbers) {
			fmt.Println("Número correto: ", nc)
		}
	}
}

func WrongPlace(n hint, h string) bool {
	var c int8 = 0

	for i, x := range n.Number {
		if strings.Contains(h, string(x)) && h[i] != n.Number[i] {
			c++
		}
	}

	return c == n.WrongPlace
}

func RightPlace(n hint, h string) bool {
	var c int8 = 0

	for i, x := range n.Number {
		if int8(x) == int8(h[i]) {
			c++
		}
	}

	return c == n.RightPlace
}

func padNumberWith3Zeros(value int) string {
	return fmt.Sprintf("%03d", value)
}
