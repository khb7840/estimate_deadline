package calc

// FLOATDELTA is a float constant for delta in testing.
// When comparing floats, use assert.InDelta.
// Differences less than FLOATDELTA will be ignored.
const FLOATDELTA = 0.000001

// Error messages

// ERRMSG_MULTIPLEVALUES - for checking multiple values
const ERRMSG_MULTIPLEVALUES = "Return values are not same with expected values"

// ERRMSG_SINGLEVALUE - for checking single value
const ERRMSG_SINGLEVALUE = "Return value is not same with expected value"

// ERRMSG_UNEXPECTED - unexpected error
const ERRMSG_UNEXPECTED = "Unexpected error raised"
