package converter_test

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/Geektree0101/clean-swift-scaffold/internal/converter"
	"github.com/Geektree0101/clean-swift-scaffold/internal/model"
)

func createSource() *converter.SourceConverter {

	date := time.Date(2020, 10, 12, 0, 0, 0, 0, time.UTC)

	header := converter.NewHeaderConverter(
		&model.Config{
			Org:          "miro.inc",
			Copyright:    "Geektree0101",
			TemplatePath: "",
			Intentation:  2,
		},
		date,
	)

	return converter.NewSourceConverter(
		"ArticleDetail",
		[]string{"Reload", "Next"},
		"./Playground/Sources",
		"./Playground/Tests",
		date,
		2,
		header,
	)
}

// MARK: - Router

func TestRenderRouter(t *testing.T) {

	templateData, err := ioutil.ReadFile("../../templates/src/Router.swift")
	expectedData, err := ioutil.ReadFile("../../test/ArticleDetailRouter.swift")

	if err != nil {
		panic(err.Error())
	}

	dummySrc := string(templateData)
	expectedSrc := string(expectedData)

	expectedDestPath := "./Playground/Sources/ArticleDetailRouter.swift"

	t.Run("render expected router", func(t *testing.T) {
		// given
		sut := createSource()

		// when
		out := sut.RenderRouter(dummySrc)

		// then
		if out.DestPath != expectedDestPath {
			t.Errorf("invalid destination path\nexpect:\n%s\noutput:\n%s\n", expectedDestPath, out.DestPath)
		}

		if out.SourceCode != expectedSrc {
			t.Errorf("Failed to render\n expected:\n%s\noutput:\n%s\n", expectedSrc, out.SourceCode)
		}
	})
}

// MARK: - Model

func TestModel(t *testing.T) {

	templateData, err := ioutil.ReadFile("../../templates/src/Model.swift")
	expectedData, err := ioutil.ReadFile("../../test/ArticleDetailModel.swift")

	if err != nil {
		panic(err.Error())
	}

	dummySrc := string(templateData)
	expectedSrc := string(expectedData)

	expectedDestPath := "./Playground/Sources/ArticleDetailModel.swift"

	t.Run("render model", func(t *testing.T) {
		// given
		sut := createSource()

		// when
		out := sut.RenderModel(dummySrc)

		// then
		if out.DestPath != expectedDestPath {
			t.Errorf("invalid destination path\nexpect:\n%s\noutput:\n%s\n", expectedDestPath, out.DestPath)
		}

		if out.SourceCode != expectedSrc {
			t.Errorf("Failed to render\n expected:\n%s\noutput:\n%s\n", expectedSrc, out.SourceCode)
		}
	})

}

// MARK: - ViewController

func TestViewController(t *testing.T) {

	templateData, err := ioutil.ReadFile("../../templates/src/ViewController.swift")
	expectedData, err := ioutil.ReadFile("../../test/ArticleDetailViewController.swift")

	if err != nil {
		panic(err.Error())
	}

	dummySrc := string(templateData)
	expectedSrc := string(expectedData)

	expectedDestPath := "./Playground/Sources/ArticleDetailViewController.swift"

	t.Run("render view controller", func(t *testing.T) {
		// given
		sut := createSource()

		// when
		out := sut.RenderViewController(dummySrc)

		// then
		if out.DestPath != expectedDestPath {
			t.Errorf("invalid destination path\nexpect:\n%s\noutput:\n%s\n", expectedDestPath, out.DestPath)
		}

		if out.SourceCode != expectedSrc {
			t.Errorf("Failed to render\n expected:\n%s\noutput:\n%s\n", expectedSrc, out.SourceCode)
		}
	})

}

// MARK: - Presenter

func TestPresenter(t *testing.T) {

	templateData, err := ioutil.ReadFile("../../templates/src/Presenter.swift")
	expectedData, err := ioutil.ReadFile("../../test/ArticleDetailPresenter.swift")

	if err != nil {
		panic(err.Error())
	}

	dummySrc := string(templateData)
	expectedSrc := string(expectedData)

	expectedDestPath := "./Playground/Sources/ArticleDetailPresenter.swift"

	t.Run("render presenter", func(t *testing.T) {
		// given
		sut := createSource()

		// when
		out := sut.RenderPresenter(dummySrc)

		// then
		if out.DestPath != expectedDestPath {
			t.Errorf("invalid destination path\nexpect:\n%s\noutput:\n%s\n", expectedDestPath, out.DestPath)
		}

		if out.SourceCode != expectedSrc {
			t.Errorf("Failed to render\n expected:\n%s\noutput:\n%s\n", expectedSrc, out.SourceCode)
		}
	})

}
