package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSalary(t *testing.T) {
	var catA, catB, catC = "A", "B", "C"
	var minA, minB, minC = 60, 60, 60
	var salA, salB, salC float32
	salA = 50*float32(minA) + 0.50*50*float32(minA)
	salB = 25*float32(minB) + 0.2*25*float32(minB)
	salC = 16.66 * float32(minC)
	//Test if catA  with minA is equal to salA
	testA := salary(catA, minA)
	assert.Equal(t, salA, testA)
	//Test if catB  with minB is equal to salB
	testB := salary(catB, minB)
	assert.Equal(t, salB, testB)
	//Test if catC  with minC is equal to salC
	testC := salary(catC, minC)
	assert.Equal(t, salC, testC)

}
