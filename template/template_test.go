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
		"same content with different value of same variable name": {
			content: "Hello, ${name}",
			name:    "name",
			value:   "Go Developer",
			want:    "Hello, Go Developer",
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
