package template

import "strings"

type Template struct {
	variables     map[string]string
	variableValue string
	content       string
}

func NewTemplate(content string) *Template {
	return &Template{content: content, variables: make(map[string]string)}
}

func (t *Template) Set(name string, value string) {
	t.variableValue = value
	t.variables[name] = value
}

func (t *Template) Evaluate() string {
	text := t.content
	for name, value := range t.variables {
		text = strings.ReplaceAll(text, "${"+name+"}", value)
	}
	return text
}
