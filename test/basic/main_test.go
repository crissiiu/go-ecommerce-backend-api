package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T) {
	var (
		input  = 1
		output = 2
	)

	actual := AddOne(input)
	if actual != output {
		t.Errorf("AddOne(%d), input(%d), actual(%d)", input, output, actual)
	}
}

func TestAddTwo(t *testing.T) {
	assert.Equal(t, AddOne(2), 3, "AddOne Shoule be 3")
}
