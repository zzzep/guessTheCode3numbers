package guess3

import (
	"testing"
)

func ExampleGuess() {
	Try()
	//Output:
	//Escolha um número, mantenha em sua mente
	//Agora é hora das dicas
	//Digite números com apenas o que indica na tela
	//Um número de 3 digitos com: um número correto e no lugar certo
	//Um número de 3 digitos com: um número correto mas no lugar errado
	//Um número de 3 digitos com: dois números corretos mas no lugar errado
	//Um número de 3 digitos com: nenhum número correto
	//Um número de 3 digitos com: um número correto mas no lugar errado
	//Número correto: 459
}

func BenchmarkTry(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Try()
	}
}

func TestRightPlaceWithNoValidNumber(t *testing.T) {
	numberTried := "170"
	hint := hint{Number: "289", RightPlace: 0, WrongPlace: 0, Message: "Unit Test"}
	response := RightPlace(hint, numberTried)
	if response == false {
		t.Error("Fail to Check Right Place", numberTried, hint)
	}
}

func TestRightPlaceWithOneValidNumber(t *testing.T) {
	numberTried := "179"
	hint := hint{Number: "289", RightPlace: 1, WrongPlace: 0, Message: "Unit Test"}
	response := RightPlace(hint, numberTried)
	if response == false {
		t.Error("Fail to Check Right Place", numberTried, hint)
	}
}

func TestRightPlaceWithTwoValidNumber(t *testing.T) {
	numberTried := "189"
	hint := hint{Number: "289", RightPlace: 2, WrongPlace: 0, Message: "Unit Test"}
	response := RightPlace(hint, numberTried)
	if response == false {
		t.Error("Fail to Check Right Place", numberTried, hint)
	}
}

func TestRightPlaceWithThreeValidNumber(t *testing.T) {
	numberTried := "289"
	hint := hint{Number: "289", RightPlace: 3, WrongPlace: 0, Message: "Unit Test"}
	response := RightPlace(hint, numberTried)
	if response == false {
		t.Error("Fail to Check Right Place", numberTried, hint)
	}
}

func TestRightPlaceWithFourValidNumber(t *testing.T) {
	numberTried := "289"
	hint := hint{Number: "289", RightPlace: 4, WrongPlace: 0, Message: "Unit Test"}
	response := RightPlace(hint, numberTried)
	if response == true {
		t.Error("Fail to Check fail in Right Place", numberTried, hint)
	}
}

func TestWrongPlaceWithNoneValidNumber(t *testing.T) {
	numberTried := "000"
	hint := hint{Number: "289", RightPlace: 0, WrongPlace: 0, Message: "Unit Test"}
	response := WrongPlace(hint, numberTried)
	if response == false {
		t.Error("Fail to Check Wrong Place", numberTried, hint)
	}
}

func TestWrongPlaceWithOneValidNumber(t *testing.T) {
	numberTried := "002"
	hint := hint{Number: "289", RightPlace: 0, WrongPlace: 1, Message: "Unit Test"}
	response := WrongPlace(hint, numberTried)
	if response == false {
		t.Error("Fail to Check Wrong Place", numberTried, hint)
	}
}

func TestWrongPlaceWithTwoValidNumber(t *testing.T) {
	numberTried := "092"
	hint := hint{Number: "289", RightPlace: 0, WrongPlace: 2, Message: "Unit Test"}
	response := WrongPlace(hint, numberTried)
	if response == false {
		t.Error("Fail to Check Wrong Place", numberTried, hint)
	}
}

func TestWrongPlaceWithThreeValidNumber(t *testing.T) {
	numberTried := "892"
	hint := hint{Number: "289", RightPlace: 0, WrongPlace: 3, Message: "Unit Test"}
	response := WrongPlace(hint, numberTried)
	if response == false {
		t.Error("Fail to Check Wrong Place", numberTried, hint)
	}
}

func TestWrongPlaceWithFourValidNumber(t *testing.T) {
	numberTried := "092"
	hint := hint{Number: "289", RightPlace: 0, WrongPlace: 4, Message: "Unit Test"}
	response := WrongPlace(hint, numberTried)
	if response == true {
		t.Error("Fail to Check Fail in Wrong Place", numberTried, hint)
	}
}

func TestPadNumberWith3Zeros(t *testing.T) {
	onlyZero := padNumberWith3Zeros(0)
	if onlyZero != "000" {
		t.Error("fail to fill with zero", onlyZero)
	}
	oneDigit := padNumberWith3Zeros(1)
	if oneDigit != "001" {
		t.Error("fail to fill with zero", oneDigit)
	}
	twoDigit := padNumberWith3Zeros(10)
	if twoDigit != "010" {
		t.Error("fail to fill with zero", twoDigit)
	}
	threeDigit := padNumberWith3Zeros(100)
	if threeDigit != "100" {
		t.Error("fail to fill with zero", threeDigit)
	}
	fourDigit := padNumberWith3Zeros(9999)
	if fourDigit != "9999" {
		t.Error("fail to fill with zero", fourDigit)
	}
}

func TestGetHint(t *testing.T) {
	hints := getHints()
	for _, h := range hints {
		if h.Number == "" {
			t.Error("Number not Found")
		}
	}
}
