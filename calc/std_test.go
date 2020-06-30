package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests for "GetStd"

func TestGetStdSingleCaseSuccess(t *testing.T) {
	// test case - single case success
	// std = (P - O) / 6
	actual, err := GetStd(2, 10)
	assert.Nil(t, err, "Error occurred.")
	expected := float64(8) / float64(6)
	assert.Equal(
		t, expected, actual,
		ERRMSG_SINGLEVALUE,
	)
}

func TestGetStdSingleFailureCase(t *testing.T) {
	// test case - single case success
	// std = (P - O) / 6
	actual, actualError := GetStd(30, 1)
	if assert.Error(t, actualError) {
		expectedErrorMsg := "This condition should be met: o < p"
		assert.EqualError(
			t, actualError, expectedErrorMsg,
			ERRMSG_UNEXPECTED,
		)
	}
	expected := float64(-1)
	assert.Equal(
		t, expected, actual,
		ERRMSG_SINGLEVALUE,
	)
}

// Tests for "GetStdMultiple"

func TestGetStdMultipleCases(t *testing.T) {
	// test case - mulitple cases success
	oArray := []float64{1, 3, 2, 5, 10, 100}
	pArray := []float64{7, 12, 4, 5.9, 200, 1000}

	var err error

	actualStdArray, actualTotalStd, err := GetStdMultiple(oArray, pArray)
	assert.Nil(t, err, ERRMSG_UNEXPECTED)

	expectedStdArray := []float64{
		float64(6) / float64(6),
		float64(9) / float64(6),
		float64(2) / float64(6),
		float64(0.9) / float64(6),
		float64(190) / float64(6),
		float64(900) / float64(6),
	}
	expectedTotalStd := float64(153.3171921)

	assert.InDeltaSlice(
		t, expectedStdArray, actualStdArray, FLOATDELTA,
		ERRMSG_MULTIPLEVALUES,
	)

	assert.InDelta(
		t, expectedTotalStd, actualTotalStd, FLOATDELTA,
		ERRMSG_SINGLEVALUE,
	)
}

func TestGetStdDifferentLengthError(t *testing.T) {
	// failure due to different input length
	oArray := []float64{1, 3, 2, 5, 10, 100}
	pArray := []float64{7, 12, 100, 4, 5.9, 200, 1000}

	actualStdArray, actualTotalStd, actualError := GetStdMultiple(
		oArray, pArray,
	)
	if assert.Error(t, actualError) {
		expectedErrorMsg := "The lengths of input arrays should be same"
		assert.EqualError(
			t, actualError, expectedErrorMsg,
			ERRMSG_UNEXPECTED,
		)
	}
	expectedStdArray := []float64{-1}
	expectedTotalStd := float64(-1)
	assert.Equal(
		t, expectedStdArray, actualStdArray,
		ERRMSG_MULTIPLEVALUES,
	)
	assert.Equal(
		t, expectedTotalStd, actualTotalStd,
		ERRMSG_SINGLEVALUE,
	)
}

func TestGetStdValueSizeError(t *testing.T) {
	// failure due to size of input
	oArray := []float64{1, 3, 2, 6, 10, 100}
	pArray := []float64{7, 12, 4, 4, 200, 1000}

	actualStdArray, actualTotalStd, actualError := GetStdMultiple(
		oArray, pArray,
	)

	if assert.Error(t, actualError) {
		expectedErrorMsg := "This condition should be met: o < p"
		assert.EqualError(
			t, actualError, expectedErrorMsg,
			ERRMSG_UNEXPECTED,
		)
	}
	expectedStdArray := []float64{-1}
	expectedTotalStd := float64(-1)
	assert.Equal(
		t, expectedStdArray, actualStdArray,
		ERRMSG_MULTIPLEVALUES,
	)
	assert.Equal(
		t, expectedTotalStd, actualTotalStd,
		ERRMSG_SINGLEVALUE,
	)
}
