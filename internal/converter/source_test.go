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

	config := &model.Config{
		Org:          "Miro",
		Copyright:    "Geektree0101",
		TemplatePath: "../../templates",
		SourceDir:    "./Playground/Sources",
		TestDir:      "./Playground/Tests",
		Indentation:  2,
	}

	header := converter.NewHeaderConverter(
		config,
		date,
	)

	return converter.NewSourceConverter(
		"ArticleDetail",
		[]string{"Reload", "Next"},
		header,
		config,
		date,
	)
}

// MARK: - Render All

func TestRenderAll(t *testing.T) {

	t.Run("render all", func(t *testing.T) {
		// given
		sut := createSource()

		// when
		out, err := sut.RenderAll()

		// then
		if err != nil {
			t.Error(err.Error())
			return
		}

		if len(out) == 0 {
			t.Errorf("unexpected rendering, output: %d", len(out))
		}
	})
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

	expectedDestPath := "./Playground/Sources/ArticleDetail/ArticleDetailRouter.swift"

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

	expectedDestPath := "./Playground/Sources/ArticleDetail/ArticleDetailModel.swift"

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

	expectedDestPath := "./Playground/Sources/ArticleDetail/ArticleDetailViewController.swift"

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

	expectedDestPath := "./Playground/Sources/ArticleDetail/ArticleDetailPresenter.swift"

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

// MARK: - Interactor

func TestInteractor(t *testing.T) {

	templateData, err := ioutil.ReadFile("../../templates/src/Interactor.swift")
	expectedData, err := ioutil.ReadFile("../../test/ArticleDetailInteractor.swift")

	if err != nil {
		panic(err.Error())
	}

	dummySrc := string(templateData)
	expectedSrc := string(expectedData)

	expectedDestPath := "./Playground/Sources/ArticleDetail/ArticleDetailInteractor.swift"

	t.Run("return expected interactor", func(t *testing.T) {
		// given
		sut := createSource()

		// when
		out := sut.RenderInteractor(dummySrc)

		// then
		if out.DestPath != expectedDestPath {
			t.Errorf("invalid destination path\nexpect:\n%s\noutput:\n%s\n", expectedDestPath, out.DestPath)
		}

		if out.SourceCode != expectedSrc {
			t.Errorf("Failed to render\n expected:\n%s\noutput:\n%s\n", expectedSrc, out.SourceCode)
		}
	})
}

// MARK: - PresenterTests

func TestPresenterTests(t *testing.T) {

	templateData, err := ioutil.ReadFile("../../templates/test/Presenter.swift")
	expectedData, err := ioutil.ReadFile("../../test/ArticleDetailPresenterTests.swift")

	if err != nil {
		panic(err.Error())
	}

	dummySrc := string(templateData)
	expectedSrc := string(expectedData)

	expectedDestPath := "./Playground/Tests/ArticleDetail/ArticleDetailPresenterTests.swift"

	t.Run("render presenter tests", func(t *testing.T) {
		// given
		sut := createSource()

		// when
		out := sut.RenderPresenterTests(dummySrc)

		// then
		if out.DestPath != expectedDestPath {
			t.Errorf("invalid destination path\nexpect:\n%s\noutput:\n%s\n", expectedDestPath, out.DestPath)
		}

		if out.SourceCode != expectedSrc {
			t.Errorf("Failed to render\n expected:\n%s\noutput:\n%s\n", expectedSrc, out.SourceCode)
		}
	})

}

// MARK: - InteractorTests

func TestInteractorTests(t *testing.T) {

	templateData, err := ioutil.ReadFile("../../templates/test/Interactor.swift")
	expectedData, err := ioutil.ReadFile("../../test/ArticleDetailInteractorTests.swift")

	if err != nil {
		panic(err.Error())
	}

	dummySrc := string(templateData)
	expectedSrc := string(expectedData)

	expectedDestPath := "./Playground/Tests/ArticleDetail/ArticleDetailInteractorTests.swift"

	t.Run("render interactor tests", func(t *testing.T) {
		// given
		sut := createSource()

		// when
		out := sut.RenderInteractorTests(dummySrc)

		// then
		if out.DestPath != expectedDestPath {
			t.Errorf("invalid destination path\nexpect:\n%s\noutput:\n%s\n", expectedDestPath, out.DestPath)
		}

		if out.SourceCode != expectedSrc {
			t.Errorf("Failed to render\n expected:\n%s\noutput:\n%s\n", expectedSrc, out.SourceCode)
		}
	})

}

// MARK: - DisplayTests

func TestDisplayTests(t *testing.T) {

	templateData, err := ioutil.ReadFile("../../templates/test/ViewController.swift")
	expectedData, err := ioutil.ReadFile("../../test/ArticleDetailViewControllerTests.swift")

	if err != nil {
		panic(err.Error())
	}

	dummySrc := string(templateData)
	expectedSrc := string(expectedData)

	expectedDestPath := "./Playground/Tests/ArticleDetail/ArticleDetailViewControllerTests.swift"

	t.Run("render view controller tests", func(t *testing.T) {
		// given
		sut := createSource()

		// when
		out := sut.RenderViewControllerTests(dummySrc)

		// then
		if out.DestPath != expectedDestPath {
			t.Errorf("invalid destination path\nexpect:\n%s\noutput:\n%s\n", expectedDestPath, out.DestPath)
		}

		if out.SourceCode != expectedSrc {
			t.Errorf("Failed to render\n expected:\n%s\noutput:\n%s\n", expectedSrc, out.SourceCode)
		}
	})
}
