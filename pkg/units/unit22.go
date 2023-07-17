package units

import (
	"fmt"
	"math/big"
)

func Calc() {
	a := big.NewInt(2)
	b := big.NewInt(20)

	// Умножение
	multiplyResult := new(big.Int).Mul(a, b)
	fmt.Println("Результат умножения:", multiplyResult)

	// Деление
	divideResult := new(big.Int).Div(a, b)
	fmt.Println("Результат деления:", divideResult)

	// Сложение
	addResult := new(big.Int).Add(a, b)
	fmt.Println("Результат сложения:", addResult)

	// Вычитание
	subtractResult := new(big.Int).Sub(a, b)
	fmt.Println("Результат вычитания:", subtractResult)
}
