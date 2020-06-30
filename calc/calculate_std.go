package calc

import (
	"math"

	"github.com/pkg/errors"
)

// GetStd calculates std (standard deviation) from two values; o and p
// This condition should be met: o < p
// Params:
//   o: A float64 value. Expected days from the most optimistic scenario
//   p: A float64 value. Expected days from the most pessimistic scenario
// Returns:
//   std: A float64 value. Standard deviation of expected days
//   error: An error.
func GetStd(o float64, p float64) (float64, error) {
	// check the size of value
	if o >= p {
		err := errors.New("This condition should be met: o < p")
		return -1, err
	}
	std := (p - o) / 6
	return std, nil
}

// GetStdMultiple calculates std (standard deviation) from two arrays.
// This condition should be met: oArray[i]  < pArray[i]
// Params:
//   oArray: A float64 array. Expected days from the most optimistic scenario
//   pArray: A float64 array. Expected days from the most pessimistic scenario
// Returns:
//   stdArray: A float64 array. Array with standard deviation of expected days
//   totalStd: A float64 value. Standard deviation of total expected days
//   error: An error.
func GetStdMultiple(oArray []float64, pArray []float64) (
	[]float64, float64, error) {
	// check input array length
	if len(oArray) != len(pArray) {
		err := errors.New("The lengths of input arrays should be same")
		return []float64{-1}, -1, err
	}

	// declare output
	stdArray := make([]float64, len(oArray))
	totalVariance := float64(0)

	// iterate
	for i := range oArray {
		// Get Std
		tempStd, err := GetStd(oArray[i], pArray[i])
		if err != nil {
			return []float64{-1}, -1, err
		}
		stdArray[i] = tempStd
		totalVariance += (tempStd * tempStd)
	}

	// return
	return stdArray, math.Sqrt(totalVariance), nil
}
