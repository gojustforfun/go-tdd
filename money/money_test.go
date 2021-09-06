package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.times(2)
	assert.Equal(t, 10, product.amount)
	product = five.times(3)
	assert.Equal(t, 15, product.amount)
}

func TestEquality(t *testing.T) {
	assert.Equal(t, NewDollar(5), NewDollar(5))
	assert.NotEqual(t, NewDollar(5), NewDollar(6))
}
