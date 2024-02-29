package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	minimum = "minimum"

	average = "average"

	maximum = "maximum"
)

func TestAverage(t *testing.T) {

	var averageExpected int = 3
	var minimumExpected int = 2
	var maximumExpected int = 5
	//send Operation minimun and grades[] and minimumExpected to the function operation and compare the result with the expected result
	minFunc, err := operation(minimum)
	assert.Equal(t, minimumExpected, minFunc(2, 3, 3, 4, 10, 2, 4, 5)) // Pass the array as a slice
	assert.Equal(t, nil, err)
	//send Operation average and grades[] and averageExpected to the function operation and compare the result with the expected result
	averageFunc, err := operation(average)
	assert.Equal(t, averageExpected, averageFunc(1, 2, 3, 4, 5))
	assert.Equal(t, nil, err)
	//send Operation maximum and grades[] and maximumExpected to the function operation and compare the result with the expected result
	maxFunc, err := operation(maximum)
	assert.Equal(t, maximumExpected, maxFunc(1, 2, 3, 4, 5))
	assert.Equal(t, nil, err)
}
