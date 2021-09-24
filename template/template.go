package template

type Template struct {
}

func NewTemplate(content string) *Template {
	return &Template{}
}

func (t *Template) Set(name string, value string) {

}

func (t *Template) Evaluate() string {
	return ""
}
