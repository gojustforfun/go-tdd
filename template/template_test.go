package template_test

import (
	"testing"

	. "github.com/gojustforfun/go-tdd/template"
	"github.com/stretchr/testify/assert"
)

func TestTemplate_OneVariable(t *testing.T) {
	testcases := map[string]struct {
		content string
		name    string
		value   string
		want    string
	}{
		"simple": {
			content: "Hello, ${name}",
			name:    "name",
			value:   "Reader",
			want:    "Hello, Reader",
		},
		"different content with different value of same variable": {
			content: "Hi, ${name}",
			name:    "name",
			value:   "Go Developer",
			want:    "Hi, Go Developer",
		},
	}
	for desc, tt := range testcases {
		t.Run(desc, func(t *testing.T) {
			tr := NewTemplate(tt.content)
			tr.Set(tt.name, tt.value)
			assert.Equal(t, tt.want, tr.Evaluate())
		})
	}
}

func TestTemplate_MultipleVariables(t *testing.T) {
	template := NewTemplate("${one}, ${two}, ${three}")
	template.Set("one", "1")
	template.Set("two", "2")
	template.Set("three", "3")
	assert.Equal(t, "1, 2, 3", template.Evaluate())
}
