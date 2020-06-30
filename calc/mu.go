package calc

import "github.com/pkg/errors"

// GetMu calculates mu (mean) from three values; o, m, and p
// This condition should be met: o < m < p
// Params:
//   o: A float64 value. Expected days from the most optimistic scenario
//   m: A float64 value. Expected days from the most likely scenario
//   p: A float64 value. Expected days from the most pessimistic scenario
// Returns:
//   mu: A float64 value. Expected average days
//   error: An error.
func GetMu(o float64, m float64, p float64) (float64, error) {
	// check the size of value
	if o >= m || m >= p || o >= p {
		err := errors.New("This condition should be met: o < m < p")
		return -1, err
	}
	// return
	mu := (o + 4*m + p) / 6
	return mu, nil
}

// GetMuMultiple calculates mu (mean) from three arrays.
// This condition should be met: oArray[i] < mArray[i] < pArray[i]
// Params:
//   oArray: A float64 array. Expected days from the most optimistic scenario
//   mArray: A float64 array. Expected days from the most likely scenario
//   pArray: A float64 array. Expected days from the most pessimistic scenario
// Returns:
//   muArray: A float64 array. Expected average days calculated
//   totalMu: A float64 value. Sum of average days
//   error: An error.
func GetMuMultiple(oArray []float64, mArray []float64, pArray []float64) (
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
		// Get Mu
		tempMu, err := GetMu(oArray[i], mArray[i], pArray[i])
		if err != nil {
			return []float64{-1}, -1, err
		}
		muArray[i] = tempMu
		totalMu += tempMu
	}
	// return
	return muArray, totalMu, nil
}
