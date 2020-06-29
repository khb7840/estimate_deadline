package main

import (
	"math"

	"github.com/pkg/errors"
)

/*
 PERT implementation
 This script is to estimate the deadline for a task (or a project)
 with trivariate input
	1) "O" for most optimistic case
	2) "M" for most likely case
	3) "P" for most pessimisitic case
 Usage with source code:
	go run estimate_deadline.go -o 1,2,3 -m 2,5,7 -p 4,9,12
	(the input lengths for "o", "m", and "p" should be same)
 Output:
	1) mean and standard deviation for individual tasks & all tasks
	2) 1-std interval & 2-std interval
 Formula:
	mu = (O + 4M + P) / 6
	std = (P - O) / 6
	total_mu = sum(mu)
	total_std = sqrt(sum(std^2))
 PERT presumes that this approximates a beta distribution.
 This makes sense since the minimum duration for a task
 is often much more certain than the maximum.
*/

// CalcMu calculates mu (mean) from three values; o, m, and p
// This condition should be met: o < m < p
// Params:
//   o: A float64 value. Expected days from the most optimistic scenario
//   m: A float64 value. Expected days from the most likely scenario
//   p: A float64 value. Expected days from the most pessimistic scenario
// Returns:
//   mu: A float64 value. Expected average days
//   error: An error.
func CalcMu(o float64, m float64, p float64) (float64, error) {
	// check the size of value
	if o >= m || m >= p || o >= p {
		err := errors.New("This condition should be met: o < m < p")
		return -1, err
	}
	// return
	mu := (o + 4*m + p) / 6
	return mu, nil
}

// CalcStd calculates std (standard deviation) from two values; o and p
// This condition should be met: o < p
// Params:
//   o: A float64 value. Expected days from the most optimistic scenario
//   p: A float64 value. Expected days from the most pessimistic scenario
// Returns:
//   std: A float64 value. Standard deviation of expected days
//   error: An error.
func CalcStd(o float64, p float64) (float64, error) {
	// check the size of value
	if o >= p {
		err := errors.New("This condition should be met: o < p")
		return -1, err
	}
	std := (p - o) / 6
	return std, nil
}

// CalcMuMultiple calculates mu (mean) from three arrays.
// This condition should be met: oArray[i] < mArray[i] < pArray[i]
// Params:
//   oArray: A float64 array. Expected days from the most optimistic scenario
//   mArray: A float64 array. Expected days from the most likely scenario
//   pArray: A float64 array. Expected days from the most pessimistic scenario
// Returns:
//   muArray: A float64 array. Expected average days calculated
//   totalMu: A float64 value. Sum of average days
//   error: An error.
func CalcMuMultiple(oArray []float64, mArray []float64, pArray []float64) (
	[]float64, float64, error) {
	// check input array length
	if len(oArray) != len(mArray) || len(oArray) != len(pArray) {
		err := errors.New("The lengths of input arrays should be same")
		return []float64{-1}, -1, err
	}
	// declare output
	muArray := make([]float64, len(oArray))
	totalMu := float64(0)
	// iterate
	for i := range oArray {
		// Calc Mu
		tempMu, err := CalcMu(oArray[i], mArray[i], pArray[i])
		if err != nil {
			return []float64{-1}, -1, err
		}
		muArray[i] = tempMu
		totalMu += tempMu
	}
	// return
	return muArray, totalMu, nil
}

// CalcStdMultiple calculates std (standard deviation) from two arrays.
// This condition should be met: oArray[i]  < pArray[i]
// Params:
//   oArray: A float64 array. Expected days from the most optimistic scenario
//   pArray: A float64 array. Expected days from the most pessimistic scenario
// Returns:
//   stdArray: A float64 array. Array with standard deviation of expected days
//   totalStd: A float64 value. Standard deviation of total expected days
//   error: An error.
func CalcStdMultiple(oArray []float64, pArray []float64) (
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
		// Calc Std
		tempStd, err := CalcStd(oArray[i], pArray[i])
		if err != nil {
			return []float64{-1}, -1, err
		}
		stdArray[i] = tempStd
		totalVariance += (tempStd * tempStd)
	}

	// return
	return stdArray, math.Sqrt(totalVariance), nil
}

// CalcInterval calculates interval with three values; mu, std, and z
// Negative minimum will be replaced with zero.
// Params:
//   mu: A float64 value. Expected average days
//   std: A float64 value. Standard deviation of expected days
//   z: A float64 value.
// Returns:
//   interval: A float64 array with two values; minVal & maxVal
func CalcInterval(mu float64, std float64, z float64) [2]float64 {
	// calculate
	minVal := mu - (std * z)
	if minVal < 0 {
		minVal = float64(0)
	}
	maxVal := mu + (std * z)
	// return
	return [2]float64{minVal, maxVal}
}

/*
func main() {

}
*/
