package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// When comparing floats, use assert.InDelta.
// Differences less than FLOATDELTA will be ignored.
const FLOATDELTA = 0.000001

// Error messages
const ERRMSG_MULTIPLEVALUES = "Return values are not same with expected values"
const ERRMSG_SINGLEVALUE = "Return value is not same with expected value"
const ERRMSG_UNEXPECTED = "Unexpected error raised"

// Tests for "CalcMu"

func TestCalcMuSingleSuccessCase(t *testing.T) {
	// test case - single case success
	// mu = (O + 4N + P) / 6
	actual, err := CalcMu(2, 3, 10)
	assert.Nil(t, err, "Error occurred.")
	expected := float64(4)
	assert.Equal(
		t, expected, actual,
		ERRMSG_SINGLEVALUE,
	)
}

func TestCalcMuSingleFailureCase(t *testing.T) {
	// test case - single case success
	// mu = (O + 4N + P) / 6
	actual, actualError := CalcMu(3, 2, 5)
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

// Tests for "CalcMuMultiple"

func TestCalcMuMultipleSuccessCase(t *testing.T) {
	// test case - mulitple cases success
	oArray := []float64{1, 3, 2, 5, 10, 100}
	mArray := []float64{3, 5, 3, 5.5, 50, 300}
	pArray := []float64{7, 12, 4, 5.9, 200, 1000}

	var err error

	actualMuArray, actualTotalMu, err := CalcMuMultiple(oArray, mArray, pArray)
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

func TestCalcMuDifferentLengthError(t *testing.T) {
	// failure due to different input length
	oArray := []float64{1, 3, 2, 5, 10, 100}
	mArray := []float64{3, 5, 5.5, 50, 300}
	pArray := []float64{7, 12, 100, 4, 5.9, 200, 1000}

	actualMuArray, actualTotalMu, actualError := CalcMuMultiple(
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

func TestCalcMuValueSizeError(t *testing.T) {
	// failure due to size of input
	oArray := []float64{1, 3, 2, 5, 10, 100}
	mArray := []float64{10, 16, 3, 5.5, 500, 300}
	pArray := []float64{7, 12, 4, 5.9, 200, 1000}

	actualMuArray, actualTotalMu, actualError := CalcMuMultiple(
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

// Tests for "CalcStd"

func TestCalcStdSingleCaseSuccess(t *testing.T) {
	// test case - single case success
	// std = (P - O) / 6
	actual, err := CalcStd(2, 10)
	assert.Nil(t, err, "Error occurred.")
	expected := float64(8) / float64(6)
	assert.Equal(
		t, expected, actual,
		ERRMSG_SINGLEVALUE,
	)
}

func TestCalcStdSingleFailureCase(t *testing.T) {
	// test case - single case success
	// std = (P - O) / 6
	actual, actualError := CalcStd(30, 1)
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

// Tests for "CalcStdMultiple"

func TestCalcStdMultipleCases(t *testing.T) {
	// test case - mulitple cases success
	oArray := []float64{1, 3, 2, 5, 10, 100}
	pArray := []float64{7, 12, 4, 5.9, 200, 1000}

	var err error

	actualStdArray, actualTotalStd, err := CalcStdMultiple(oArray, pArray)
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

func TestCalcStdDifferentLengthError(t *testing.T) {
	// failure due to different input length
	oArray := []float64{1, 3, 2, 5, 10, 100}
	pArray := []float64{7, 12, 100, 4, 5.9, 200, 1000}

	actualStdArray, actualTotalStd, actualError := CalcStdMultiple(
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

func TestCalcStdValueSizeError(t *testing.T) {
	// failure due to size of input
	oArray := []float64{1, 3, 2, 6, 10, 100}
	pArray := []float64{7, 12, 4, 4, 200, 1000}

	actualStdArray, actualTotalStd, actualError := CalcStdMultiple(
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

func TestCalcInterval(t *testing.T) {
	actualOneStdInterval := CalcInterval(5, 4, 1)
	expectedOneStdInterval := [2]float64{1, 9}
	actualTwoStdInterval := CalcInterval(5, 4, 2)
	expectedTwoStdInterval := [2]float64{0, 13}
	assert.Equal(
		t, expectedOneStdInterval, actualOneStdInterval,
		ERRMSG_MULTIPLEVALUES,
	)
	assert.Equal(
		t, expectedTwoStdInterval, actualTwoStdInterval,
		ERRMSG_MULTIPLEVALUES,
	)
}

/*
func TestParseFlags(t *testing.T) {
	var tests = []struct {
		args []string
		oArray []float64
		mArray []float64
		pArray []float64
	}{
		{[]string{"-o","1","-m","3","-p","7"},
		 []float64{1}, []float64{3}, []float64{7}},
		{[]string{"-o","1,3,10","-m","3,5,50","-p","7,12,200"},
		 []float64{1,3,10}, []float64{3,5,50}, []float64{7,12,200}},
	}
}
*/
