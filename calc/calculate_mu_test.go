package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests for "GetMu"

func TestGetMuSingleSuccessCase(t *testing.T) {
	// test case - single case success
	// mu = (O + 4N + P) / 6
	actual, err := GetMu(2, 3, 10)
	assert.Nil(t, err, "Error occurred.")
	expected := float64(4)
	assert.Equal(
		t, expected, actual,
		ERRMSG_SINGLEVALUE,
	)
}

func TestGetMuSingleFailureCase(t *testing.T) {
	// test case - single case success
	// mu = (O + 4N + P) / 6
	actual, actualError := GetMu(3, 2, 5)
	if assert.Error(t, actualError) {
		expectedErrorMsg := "This condition should be met: o < m < p"
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

// Tests for "GetMuMultiple"

func TestGetMuMultipleSuccessCase(t *testing.T) {
	// test case - mulitple cases success
	oArray := []float64{1, 3, 2, 5, 10, 100}
	mArray := []float64{3, 5, 3, 5.5, 50, 300}
	pArray := []float64{7, 12, 4, 5.9, 200, 1000}

	var err error

	actualMuArray, actualTotalMu, err := GetMuMultiple(oArray, mArray, pArray)
	assert.Nil(t, err, ERRMSG_UNEXPECTED)

	expectedMuArray := []float64{
		float64(20) / float64(6),
		float64(35) / float64(6),
		float64(18) / float64(6),
		float64(32.9) / float64(6),
		float64(410) / float64(6),
		float64(2300) / float64(6),
	}
	expectedTotalMu := float64(2815.9) / float64(6)

	assert.Equal(
		t, expectedMuArray, actualMuArray,
		ERRMSG_MULTIPLEVALUES,
	)
	assert.InDelta(
		t, expectedTotalMu, actualTotalMu, FLOATDELTA,
		ERRMSG_SINGLEVALUE,
	)
}

func TestGetMuDifferentLengthError(t *testing.T) {
	// failure due to different input length
	oArray := []float64{1, 3, 2, 5, 10, 100}
	mArray := []float64{3, 5, 5.5, 50, 300}
	pArray := []float64{7, 12, 100, 4, 5.9, 200, 1000}

	actualMuArray, actualTotalMu, actualError := GetMuMultiple(
		oArray, mArray, pArray,
	)
	if assert.Error(t, actualError) {
		expectedErrorMsg := "The lengths of input arrays should be same"
		assert.EqualError(
			t, actualError, expectedErrorMsg,
			ERRMSG_UNEXPECTED,
		)
	}
	expectedMuArray := []float64{-1}
	expectedTotalMu := float64(-1)
	assert.Equal(
		t, expectedMuArray, actualMuArray,
		ERRMSG_MULTIPLEVALUES,
	)
	assert.Equal(
		t, expectedTotalMu, actualTotalMu,
		ERRMSG_SINGLEVALUE,
	)
}

func TestGetMuValueSizeError(t *testing.T) {
	// failure due to size of input
	oArray := []float64{1, 3, 2, 5, 10, 100}
	mArray := []float64{10, 16, 3, 5.5, 500, 300}
	pArray := []float64{7, 12, 4, 5.9, 200, 1000}

	actualMuArray, actualTotalMu, actualError := GetMuMultiple(
		oArray, mArray, pArray,
	)

	if assert.Error(t, actualError) {
		expectedErrorMsg := "This condition should be met: o < m < p"
		assert.EqualError(
			t, actualError, expectedErrorMsg,
			ERRMSG_UNEXPECTED,
		)
	}
	expectedMuArray := []float64{-1}
	expectedTotalMu := float64(-1)
	assert.Equal(
		t, expectedMuArray, actualMuArray,
		ERRMSG_MULTIPLEVALUES,
	)
	assert.Equal(
		t, expectedTotalMu, actualTotalMu,
		ERRMSG_SINGLEVALUE,
	)
}
