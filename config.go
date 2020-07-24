package main

import (
	"flag"
	"strconv"
	"strings"
)

// Config is a struct with three float arrays
type Config struct {
	oArray []float64
	mArray []float64
	pArray []float64
}

func parseFlags(configPtr *Config, args []string) (*Config, error) {
	var config Config

	// Let's parse flags here
	oStringPtr := flag.String(
		"o", "",
		"Expected time with most optimistic scenario",
	) // flag.* return a pointer variable
	mStringPtr := flag.String(
		"m", "",
		"Expected time with most likely scenario",
	)
	pStringPtr := flag.String(
		"p", "",
		"Expected time with most pessimistic scenario",
	)

	flag.CommandLine.Parse(args)

	// Split string with comma
	oStrings := strings.Split(*oStringPtr, ",")
	mStrings := strings.Split(*mStringPtr, ",")
	pStrings := strings.Split(*pStringPtr, ",")

	for i, oChar := range oStrings {
		// get values with index
		mChar := mStrings[i]
		pChar := pStrings[i]
		// parse into float64
		oVal, _ := strconv.ParseFloat(oChar, 64)
		mVal, _ := strconv.ParseFloat(mChar, 64)
		pVal, _ := strconv.ParseFloat(pChar, 64)
		// append values
		config.oArray = append(config.oArray, oVal)
		config.mArray = append(config.mArray, mVal)
		config.pArray = append(config.pArray, pVal)
	}
	configPtr = &config

	return configPtr, nil
}
