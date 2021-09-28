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

func (t *Template) Evaluate() (text string, err error) {
	defer func() {
		if r, ok := recover().(string); ok {
			text, err = "", fmt.Errorf("%w : %s", ErrMissingValueForVariable, r)
		}
	}()
	return t.replaceVariables(), nil
}

func (t *Template) replaceVariables() string {
	re := regexp.MustCompile(`\$\{\w+\}`)
	return re.ReplaceAllStringFunc(t.text, func(variable string) string {
		name := strings.Trim(variable, "${}")
		if _, ok := t.variables[name]; !ok {
			panic(variable)
		}
		return t.variables[name]
	})
}
