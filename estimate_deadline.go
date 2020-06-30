package main

import (
	"fmt"

	"github.com/khb7840/estimate_deadline/calc"
)

func main() {
	a, _ := calc.GetMu(1, 5, 10)
	fmt.Print(a)
}
