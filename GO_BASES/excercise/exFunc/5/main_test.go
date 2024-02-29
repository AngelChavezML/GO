package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDog(t *testing.T) {
	animalDog, msg := animal("dog")
	assert.Equal(t, "dog", msg)
	var expected float64 = 100
	assert.Equal(t, expected, animalDog(10))
}
func TestCat(t *testing.T) {
	animalCat, msg := animal("cat")
	assert.Equal(t, "cat", msg)
	var expected float64 = 50
	assert.Equal(t, expected, animalCat(10))
}
func TestHamster(t *testing.T) {
	animalHamster, msg := animal("hamster")
	assert.Equal(t, "hamster", msg)
	var expected float64 = 2.5
	assert.Equal(t, expected, animalHamster(10))
}
func TestTarantula(t *testing.T) {
	animalTarantula, msg := animal("tarantula")
	assert.Equal(t, "tarantula", msg)
	var expected float64 = 1.5
	assert.Equal(t, expected, animalTarantula(10))
}
