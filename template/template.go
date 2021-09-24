package template

import "strings"

type Template struct {
	variableValue string
	content       string
}

func NewTemplate(content string) *Template {
	return &Template{content: content}
}

func (t *Template) Set(name string, value string) {
	t.variableValue = value
}

func (t *Template) Evaluate() string {
	return strings.ReplaceAll(t.content, "${name}", t.variableValue)
}
