package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestParseFlags is a test function for parseFlags
func TestParseFlags(t *testing.T) {
	testArgs := []string{
		"-o", "1,3,5", "-m", "2,5,7", "-p", "5,10,20",
	}

	expectedConfig := Config{
		oArray: []float64{1, 3, 5},
		mArray: []float64{2, 5, 7},
		pArray: []float64{5, 10, 20},
	}

	var actualConfigPtr *Config

	actualConfigPtr, _ = parseFlags(actualConfigPtr, testArgs)

	assert.Equal(
		t, actualConfigPtr.oArray, expectedConfig.oArray,
		"oArray is not same with expected",
	)
	assert.Equal(
		t, actualConfigPtr.mArray, expectedConfig.mArray,
		"mArray is not same with expected",
	)
	assert.Equal(
		t, actualConfigPtr.pArray, expectedConfig.pArray,
		"pArray is not same with expected",
	)
}
