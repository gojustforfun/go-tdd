package template_test

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	template = NewTemplate("Hello, ${name}")
	template.Set("name", "Reader")
	assertEqual(t, "Hello, Reader", template.Evaluate())
}
