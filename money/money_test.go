package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	five.times(2)
	assert.Equal(t, 10, five.amount)
	five.times(3)
	assert.Equal(t, 15, five.amount)
}
