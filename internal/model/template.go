package model

import (
	"fmt"
	"strings"
)

func RenderUsecaseTemplate(usecase string, indentation int) string {

	mutStr := ""
	// Open
	mutStr += fmt.Sprintf("%senum %s {\n\n", whiteSpace(indentation), usecase)

	// Request
	mutStr += fmt.Sprintf("%sstruct Request {\n\n", whiteSpace(indentation*2))
	mutStr += fmt.Sprintf("%s}\n\n", whiteSpace(indentation*2))

	// Response
	mutStr += fmt.Sprintf("%sstruct Response {\n\n", whiteSpace(indentation*2))
	mutStr += fmt.Sprintf("%s}\n\n", whiteSpace(indentation*2))

	// ViewModel
	mutStr += fmt.Sprintf("%sstruct ViewModel {\n\n", whiteSpace(indentation*2))
	mutStr += fmt.Sprintf("%s}\n\n", whiteSpace(indentation*2))

	// Close
	mutStr += fmt.Sprintf("%s}", whiteSpace(indentation))

	return mutStr
}

func RenderDisplayInterface(sceneName string, usecase string, indentation int) string {

	return fmt.Sprintf(
		"%sfunc display%s(viewModel: %s.%s.ViewModel)",
		whiteSpace(indentation),
		usecase,
		sceneName,
		usecase,
	)
}

func RenderDisplayImpl(sceneName string, usecase string, indentation int) string {

	mutStr := ""

	// Open
	mutStr += fmt.Sprintf(
		"%sfunc display%s(viewModel: %s.%s.ViewModel) {\n\n",
		whiteSpace(indentation),
		usecase,
		sceneName,
		usecase,
	)

	// Close
	mutStr += fmt.Sprintf("%s}\n\n", whiteSpace(indentation))

	return mutStr
}

func whiteSpace(i int) string {
	whitespaces := make([]string, i)
	for i := range whitespaces {
		whitespaces[i] = " "
	}
	return strings.Join(whitespaces, "")
}

const (

	// interactor
	InteractorImplTemplate string = `
	  func __METHOD_NAME__(request: __SCENE_NAME__.__USECASE__.Request) {

		}
	`
	InteractorInterpaceTemplate string = `func __METHOD_NAME__(request: __SCENE_NAME__.__USECASE__.Request)`
	// presenter
	PresenterImplTemplate string = `
	  func present__USECASE__(response: __SCENE_NAME__.__USECASE__.Response) {

		}
	`
	PresenterInterpaceTemplate string = `func present__USECASE__(response: __SCENE_NAME__.__USECASE__.Response)`
)
