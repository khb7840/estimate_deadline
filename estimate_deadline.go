package main

import (
	"fmt"
	"os"

	"github.com/khb7840/estimate_deadline/calc"
)

func printOutput(muArray, stdArray []float64, totalMu, totalStd float64) {
	for i, mu := range muArray {
		std := stdArray[i]
		fmt.Printf("Task %d - mean of expected time: %0.2f\n", i+1, mu)
		fmt.Printf("Task %d - std of expected time: %0.2f\n", i+1, std)
	}
	fmt.Printf("Total tasks - mean of expected time: %0.2f\n", totalMu)
	fmt.Printf("Total task - std of expected time: %0.2f\n", totalStd)
}

func main() {

	var mainConfig *Config

	mainConfig, _ = parseFlags(mainConfig, os.Args[1:])

	muArray, totalMu, err := calc.GetMuMultiple(
		mainConfig.oArray, mainConfig.mArray,
		mainConfig.pArray,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	stdArray, totalStd, err := calc.GetStdMultiple(
		mainConfig.oArray, mainConfig.pArray,
	)
	// print output
	printOutput(muArray, stdArray, totalMu, totalStd)
}
