package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSalary(t *testing.T) {
	//Test when salary when is smaller than 50000. Expected 0 and shows if it passed the test
	assert.Equal(t, 0.0, discount(40000))
}
