package template

type Template struct {
	variableValue string
}

func NewTemplate(content string) *Template {
	return &Template{}
}

func (t *Template) Set(name string, value string) {
	t.variableValue = value
}

func (t *Template) Evaluate() string {
	return "Hello, " + t.variableValue
}
