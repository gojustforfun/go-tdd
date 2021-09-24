package template_test

import (
	"testing"

	. "github.com/gojustforfun/go-tdd/template"
	"github.com/stretchr/testify/assert"
)

func TestTemplate(t *testing.T) {
	template := NewTemplate("Hello, ${name}")
	template.Set("name", "Reader")
	assert.Equal(t, "Hello, Reader", template.Evaluate())
}
