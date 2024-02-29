package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	/*var grades [5]int = [5]int{1, 2, 3, 4, 5}
	var averageExpected float64 = 3.0
	//send grades[] and averageExpected to the function average and compare the result with the expected result
	average, err := average(grades[:])
	assert.Equal(t, averageExpected, average)
	assert.Equal(t, "", err)*/
	var grades [5]int = [5]int{1, 2, 3, 4, -5}
	//var averageExpected float64 = 0.0
	var errExpected string = "negative number"
	average, err := average(grades[:])
	//assert.Equal(t, averageExpected, average)
	if assert.Equal(t, errExpected, err) {
		fmt.Println("There is a negative number", average)
	}
}
