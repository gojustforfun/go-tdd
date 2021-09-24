package template

import "strings"

type Template struct {
	variables map[string]string
	text      string
}

func NewTemplate(text string) *Template {
	return &Template{text: text, variables: make(map[string]string)}
}

func (t *Template) Set(name string, value string) {
	t.variables[name] = value
}

func (t *Template) Evaluate() string {
	text := t.text
	for name, value := range t.variables {
		text = strings.ReplaceAll(text, "${"+name+"}", value)
	}
	return text
}
