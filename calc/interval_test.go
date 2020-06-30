package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInterval(t *testing.T) {
	actualOneStdInterval := GetInterval(5, 4, 1)
	expectedOneStdInterval := [2]float64{1, 9}
	actualTwoStdInterval := GetInterval(5, 4, 2)
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
