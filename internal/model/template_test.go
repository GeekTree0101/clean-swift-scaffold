package model_test

import (
	"testing"

	"github.com/Geektree0101/clean-swift-scaffold/internal/model"
)

func TestRenderModelTemplate(t *testing.T) {

	expectedOutput := "  enum ArticleDetail {\n\n    struct Request {\n\n    }\n\n    struct Response {\n\n    }\n\n    struct ViewModel {\n\n    }\n\n  }"

	t.Run("return expected model", func(t *testing.T) {
		// given
		usecase := "ArticleDetail"
		indentation := 2

		// when
		out := model.RenderUsecaseTemplate(usecase, indentation)

		// then
		if out != expectedOutput {
			t.Errorf("invalid output\nexpect:\n%s\n\noutput:\n%s\n", expectedOutput, out)
		}
	})

}
