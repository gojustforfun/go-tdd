package template_test

import (
	"testing"

	. "github.com/gojustforfun/go-tdd/template"
	"github.com/stretchr/testify/suite"
)

func TestTemplateTestSuite(t *testing.T) {
	suite.Run(t, new(TemplateTestSuite))
}

type TemplateTestSuite struct {
	suite.Suite
	template *Template
}

func (s *TemplateTestSuite) SetupTest() {
	s.template = NewTemplate("${one}, ${two}, ${three}")
	s.template.Set("one", "1")
	s.template.Set("two", "2")
	s.template.Set("three", "3")
}

func (s *TemplateTestSuite) Test_Multiple_Variables() {
	s.assertTemplateEvaluateTo("1, 2, 3")
}

func (s *TemplateTestSuite) Test_Ignore_Unknown_Variables() {
	s.template.Set("unknown_four", "4")
	s.template.Set("unknown_five", "5")
	s.assertTemplateEvaluateTo("1, 2, 3")
}

func (s *TemplateTestSuite) Test_Multiple_Evaluate_Calls() {

	s.Run("Before Modify Values", func() {
		s.assertTemplateEvaluateTo("1, 2, 3")
		s.assertTemplateEvaluateTo("1, 2, 3")
	})

	s.Run("After Modify Values", func() {
		s.template.Set("one", "3")
		s.template.Set("two", "2")
		s.template.Set("three", "1")
		s.template.Set("six", "6")
		s.assertTemplateEvaluateTo("3, 2, 1")
		s.assertTemplateEvaluateTo("3, 2, 1")
	})

}

func (s *TemplateTestSuite) Test_Different_Template_Text() {
	s.template = NewTemplate("Hi, ${name}")
	s.template.Set("name", "Gopher")
	s.assertTemplateEvaluateTo("Hi, Gopher")
}

func (s *TemplateTestSuite) Test_No_Variable_In_Template_Text() {

	s.template = NewTemplate("Go")
	s.assertTemplateEvaluateTo("Go")

	s.Run("Set Multiple Variables", func() {
		s.template.Set("unknown", "java")
		s.template.Set("seven", "python")
		s.assertTemplateEvaluateTo("Go")
	})
}

func (s *TemplateTestSuite) Test_Unreplaced_Variables_Should_Return_Error() {
	s.template = NewTemplate("${foo} ${bar}")
	s.assertTemplateEvaluateReturnError(ErrMissingValueForVariable)
}

func (s *TemplateTestSuite) Test_Variables_Get_Processed_Just_Once() {
	s.template.Set("one", "${one}")
	s.template.Set("two", "${three}")
	s.template.Set("three", "${two}")
	s.assertTemplateEvaluateTo("${one}, ${three}, ${two}")
}

func (s *TemplateTestSuite) assertTemplateEvaluateTo(expected string) {
	actual, err := s.template.Evaluate()
	s.NoError(err)
	s.Equal(expected, actual)
}

func (s *TemplateTestSuite) assertTemplateEvaluateReturnError(target error) {
	actual, err := s.template.Evaluate()
	s.Zero(actual)
	s.Error(err)
	s.ErrorIs(err, target)
}
