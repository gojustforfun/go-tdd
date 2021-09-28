package template

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	ErrMissingValueForVariable = errors.New("missing value for variable")
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
	text := t.replaceVariables()
	if names := t.FindUnreplacedVariables(text); len(names) != 0 {
		return "", fmt.Errorf("%w no value for %s", ErrMissingValueForVariable, strings.Join(names, " "))
	}
	return text, nil
}

func (t *Template) replaceVariables() string {
	text := t.text
	for name, value := range t.variables {
		text = strings.ReplaceAll(text, "${"+name+"}", value)
	}
	return text
}

// FindUnreplacedVariables return names of unreplaced variable
func (t *Template) FindUnreplacedVariables(text string) []string {
	return regexp.MustCompile(`\$\{.+\}`).FindStringSubmatch(text)
}
