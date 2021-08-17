package model

import (
	"fmt"
	"strings"
	"unicode"
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
		"%sfunc display%s(viewModel: %sModel.%s.ViewModel)",
		whiteSpace(indentation),
		usecase,
		sceneName,
		usecase,
	)
}

func RenderDisplayImpl(sceneName string, usecase string, indentation int) string {

	mutStr := ""
	mutStr += RenderDisplayInterface(sceneName, usecase, indentation)
	mutStr += openAndClose(indentation)

	return mutStr
}

func RenderPresenterInterface(sceneName string, usecase string, indentation int) string {

	return fmt.Sprintf(
		"%sfunc present%s(response: %sModel.%s.Response)",
		whiteSpace(indentation),
		usecase,
		sceneName,
		usecase,
	)
}

func RenderPresenterImpl(sceneName string, usecase string, indentation int) string {

	mutStr := ""
	mutStr += RenderPresenterInterface(sceneName, usecase, indentation)
	mutStr += openAndClose(indentation)

	return mutStr
}

func RenderInteractorInterface(sceneName string, usecase string, indentation int) string {

	return fmt.Sprintf(
		"%sfunc %s(request: %sModel.%s.Request)",
		whiteSpace(indentation),
		lowerFirstChar(usecase),
		sceneName,
		usecase,
	)
}

func RenderInteractorImpl(sceneName string, usecase string, indentation int) string {

	mutStr := ""
	mutStr += RenderInteractorInterface(sceneName, usecase, indentation)
	mutStr += openAndClose(indentation)

	return mutStr
}

func openAndClose(indentation int) string {
	/**
	example
	{

	}
	**/
	mutStr := ""
	mutStr += " {"
	mutStr += "\n\n"
	mutStr += fmt.Sprintf("%s}", whiteSpace(indentation))
	return mutStr
}

func lowerFirstChar(str string) string {
	/**
	example
	ReloadData -> reloadData
	**/
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return str
}

func whiteSpace(i int) string {
	whitespaces := make([]string, i)
	for i := range whitespaces {
		whitespaces[i] = " "
	}
	return strings.Join(whitespaces, "")
}
