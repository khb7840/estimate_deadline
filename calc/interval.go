package calc

// GetInterval calculates interval with three values; mu, std, and z
// Negative minimum will be replaced with zero.
// Params:
//   mu: A float64 value. Expected average days
//   std: A float64 value. Standard deviation of expected days
//   z: A float64 value.
// Returns:
//   interval: A float64 array with two values; minVal & maxVal
func GetInterval(mu float64, std float64, z float64) [2]float64 {
	// calculate
	minVal := mu - (std * z)
	if minVal < 0 {
		minVal = float64(0)
	}
	maxVal := mu + (std * z)
	// return
	return [2]float64{minVal, maxVal}
}
