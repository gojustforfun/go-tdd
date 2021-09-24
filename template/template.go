package template

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrUnmatchedVariable = errors.New("unmatched variable")
)

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

func (t *Template) Evaluate() (string, error) {
	text := t.text
	for name, value := range t.variables {
		text = strings.ReplaceAll(text, "${"+name+"}", value)
	}
	if t.hasUnmatchedVariables(text) {
		return "", ErrUnmatchedVariable
	}
	return text, nil
}

func (t *Template) hasUnmatchedVariables(text string) bool {
	matched, _ := regexp.MatchString(`(\$\{(\d|\w)+\})+`, text)
	return matched
}
