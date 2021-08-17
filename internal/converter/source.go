package converter

import (
	"fmt"
	"strings"
	"time"

	"github.com/Geektree0101/clean-swift-scaffold/internal/model"
)

type SourceConverter struct {
	sceneName   string
	usecases    []string
	sourcePath  string
	testPath    string
	date        time.Time
	indentation int
	header      *HeaderConverter
}

func NewSourceConverter(
	sceneName string,
	usecases []string,
	sourcePath string,
	testPath string,
	date time.Time,
	indentation int,
	header *HeaderConverter) *SourceConverter {

	return &SourceConverter{
		sceneName:   sceneName,
		usecases:    usecases,
		sourcePath:  sourcePath,
		testPath:    testPath,
		date:        date,
		indentation: indentation,
		header:      header,
	}
}

func (c *SourceConverter) RenderAll() ([]model.Source, error) {
	// TODO:

	return []model.Source{}, nil
}

func (c *SourceConverter) RenderInteractor(src string) *model.Source {

	return nil
}

func (c *SourceConverter) RenderPresenter(src string) *model.Source {

	return nil
}

func (c *SourceConverter) RenderViewController(src string) *model.Source {

	var mutSrc string = src
	mutSrc = strings.ReplaceAll(mutSrc, "__SCENE_NAME__", c.sceneName)

	interfaceCompositionToken := "// clean-swift-scaffold-generate-display-interface (do-not-remove-comments)"
	implementCompositionToken := "// clean-swift-scaffold-generate-display-implementation (do-not-remove-comments)"

	ifs := []string{}

	for _, uc := range c.usecases {
		ifs = append(ifs, model.RenderDisplayInterface(c.sceneName, uc, c.indentation))
	}

	mutSrc = strings.ReplaceAll(mutSrc, interfaceCompositionToken, strings.Join(ifs, "\n"))

	imples := []string{}

	for _, uc := range c.usecases {
		imples = append(imples, model.RenderDisplayImpl(c.sceneName, uc, c.indentation))
	}

	mutSrc = strings.ReplaceAll(mutSrc, implementCompositionToken, strings.Join(imples, "\n\n"))
	mutSrc = c.header.Render(mutSrc, c.sceneName)

	return &model.Source{
		DestPath:   fmt.Sprintf("%s/%sViewController.swift", c.sourcePath, c.sceneName),
		SourceCode: mutSrc,
	}
}

func (c *SourceConverter) RenderModel(src string) *model.Source {

	var mutSrc string = src
	mutSrc = strings.ReplaceAll(mutSrc, "__SCENE_NAME__", c.sceneName)
	compositionToken := "// clean-swift-scaffold-generate-dto (do-not-remove-comments)"

	imples := []string{}

	for _, uc := range c.usecases {
		imples = append(imples, model.RenderUsecaseTemplate(uc, c.indentation))
	}

	mutSrc = strings.ReplaceAll(mutSrc, compositionToken, strings.Join(imples, "\n\n"))
	mutSrc = c.header.Render(mutSrc, c.sceneName)

	return &model.Source{
		DestPath:   fmt.Sprintf("%s/%sModel.swift", c.sourcePath, c.sceneName),
		SourceCode: mutSrc,
	}
}

func (c *SourceConverter) RenderRouter(src string) *model.Source {

	var mutSrc string = src
	mutSrc = strings.ReplaceAll(mutSrc, "__SCENE_NAME__", c.sceneName)
	mutSrc = c.header.Render(mutSrc, c.sceneName)

	return &model.Source{
		DestPath:   fmt.Sprintf("%s/%sRouter.swift", c.sourcePath, c.sceneName),
		SourceCode: mutSrc,
	}
}
