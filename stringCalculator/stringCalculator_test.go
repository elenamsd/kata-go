package stringCalculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 3, Add("1,2"), "Test of 1,2 add must be 3")
}
