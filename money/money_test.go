package money_test

import (
	"testing"

	. "github.com/gojustforfun/go-tdd/money"
	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	assert.Equal(t, NewDollar(10), five.Times(2))
	assert.Equal(t, NewDollar(15), five.Times(3))
}

func TestEquality(t *testing.T) {
	assert.Equal(t, NewDollar(5), NewDollar(5))
	assert.NotEqual(t, NewDollar(5), NewDollar(6))
}
