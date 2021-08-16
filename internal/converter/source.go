package converter

import (
	"fmt"
	"strings"
	"time"

	"github.com/Geektree0101/clean-swift-scaffold/internal/model"
)

type SourceConverter struct {
	sceneName  string
	usecases   []string
	sourcePath string
	testPath   string
	date       time.Time
	header     *HeaderConverter
}

func NewSourceConverter(
	sceneName string,
	usecases []string,
	sourcePath string,
	testPath string,
	date time.Time,
	header *HeaderConverter) *SourceConverter {

	return &SourceConverter{
		sceneName:  sceneName,
		usecases:   usecases,
		sourcePath: sourcePath,
		testPath:   testPath,
		date:       date,
		header:     header,
	}
}

func (c *SourceConverter) RenderAll() ([]model.Source, error) {
	// TODO:

	return []model.Source{}, nil
}

func (c *SourceConverter) RenderInteractor(src string) (*model.Source, error) {

	return nil, nil
}

func (c *SourceConverter) RenderPresenter(src string) (*model.Source, error) {

	return nil, nil
}

func (c *SourceConverter) RenderViewController(src string) (*model.Source, error) {

	return nil, nil
}

func (c *SourceConverter) RenderModel(src string) *model.Source {

	var mutSrc string = src
	mutSrc = strings.ReplaceAll(mutSrc, "__SCENE_NAME__", c.sceneName)
	compositionToken := "// clean-swift-scaffold-generate-dto (do-not-remove-comments)"

	imples := []string{}

	for _, uc := range c.usecases {
		mutTemplate := model.UsecaseTemplate
		mutTemplate = strings.ReplaceAll(mutTemplate, "__USECASE__", uc)
		imples = append(imples, mutTemplate)
	}

	mutSrc = strings.ReplaceAll(mutSrc, compositionToken, strings.Join(imples, "\n"))
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
