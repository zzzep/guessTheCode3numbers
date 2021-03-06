//Copyright ₢
//Author: Giuseppe Fechio
//Email: giuseppe.fechio@gmail.com
package guess3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type hint struct {
	Number     string
	RightPlace int8
	WrongPlace int8
	Message    string
}

const maxDigits = 3

func Try() {
	hints := getHints()

	corrects := 0

	for n := 0; n <= 999; n++ {
		nc := padNumberWith3Zeros(n)
		corrects = 0

		for hintsCount := 0; hintsCount < len(hints); hintsCount++ {
			if RightPlace(hints[hintsCount], nc) == false {
				break
			}

			if WrongPlace(hints[hintsCount], nc) == false {
				break
			}

			corrects++
		}

		if corrects == len(hints) {
			fmt.Println("Número correto:", nc)
		}
	}
}

func getHints() [5]hint {
	hints := [5]hint{
		{"289", 1, 0, "Um número de 3 digitos com: um número correto e no lugar certo"},
		{"215", 0, 1, "Um número de 3 digitos com: um número correto mas no lugar errado"},
		{"942", 0, 2, "Um número de 3 digitos com: dois números corretos mas no lugar errado"},
		{"738", 0, 0, "Um número de 3 digitos com: nenhum número correto"},
		{"784", 0, 1, "Um número de 3 digitos com: um número correto mas no lugar errado"},
	}

	fmt.Println("Escolha um número, mantenha em sua mente")
	time.Sleep(3)
	fmt.Println("Agora é hora das dicas\nDigite números com apenas o que indica na tela")

	for i, h := range hints {
		fmt.Println(h.Message)
		buf := bufio.NewReader(os.Stdin)
		input, _ := buf.ReadString('\n')

		numberTried := leftPad(input, "0", 3)
		if numberTried != "" {
			hints[i].Number = numberTried
		}
	}

	return hints
}

func WrongPlace(n hint, h string) bool {
	var c int8 = 0

	for i, x := range n.Number {
		if i >= maxDigits {
			break
		}
		if strings.Contains(h, string(x)) && h[i] != n.Number[i] {
			c++
		}
	}

	return c == n.WrongPlace
}

func RightPlace(n hint, h string) bool {
	var c int8 = 0

	for i, x := range n.Number {
		if i >= 3 {
			break
		}
		if int8(x) == int8(h[i]) {
			c++
		}
	}

	return c == n.RightPlace
}

func padNumberWith3Zeros(value int) string {
	return fmt.Sprintf("%03d", value)
}

func leftPad(s string, p string, l int) string {
	return (strings.Repeat(p, l) + s)[l:]
}
