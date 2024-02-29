package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSalary(t *testing.T) {
	//Test when salary when is smaller than 50000. Expected 0 and shows if it passed the test
	//assert.Equal(t, 0.0, discount(40000))
	//Test when salary when is greater than 50000 and smaller than 150000. Expected 8670 and shows if it passed the test
	assert.Equal(t, 8670.0, discount(51000))
}
