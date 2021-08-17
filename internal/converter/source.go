package converter

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/Geektree0101/clean-swift-scaffold/internal/model"
)

type SourceConverter struct {
	sceneName string
	usecases  []string
	header    *HeaderConverter
	config    *model.Config
	date      time.Time
}

func NewSourceConverter(
	sceneName string,
	usecases []string,
	header *HeaderConverter,
	config *model.Config,
	date time.Time,
) *SourceConverter {

	return &SourceConverter{
		sceneName: sceneName,
		usecases:  usecases,
		header:    header,
		config:    config,
		date:      date,
	}
}

func (c *SourceConverter) RenderAll() ([]model.Source, error) {
	interactorByte, err := ioutil.ReadFile(fmt.Sprintf("%s/src/Interactor.swift", c.config.TemplatePath))

	if err != nil {
		return nil, err
	}

	presenterByte, err := ioutil.ReadFile(fmt.Sprintf("%s/src/Presenter.swift", c.config.TemplatePath))

	if err != nil {
		return nil, err
	}

	displayByte, err := ioutil.ReadFile(fmt.Sprintf("%s/src/ViewController.swift", c.config.TemplatePath))

	if err != nil {
		return nil, err
	}

	routerByte, err := ioutil.ReadFile(fmt.Sprintf("%s/src/Router.swift", c.config.TemplatePath))

	if err != nil {
		return nil, err
	}

	modelByte, err := ioutil.ReadFile(fmt.Sprintf("%s/src/Model.swift", c.config.TemplatePath))

	if err != nil {
		return nil, err
	}

	interactorTestsByte, err := ioutil.ReadFile(fmt.Sprintf("%s/test/Interactor.swift", c.config.TemplatePath))

	if err != nil {
		return nil, err
	}

	presenterTestsByte, err := ioutil.ReadFile(fmt.Sprintf("%s/test/Presenter.swift", c.config.TemplatePath))

	if err != nil {
		return nil, err
	}

	displayTestsByte, err := ioutil.ReadFile(fmt.Sprintf("%s/test/ViewController.swift", c.config.TemplatePath))

	if err != nil {
		return nil, err
	}

	sourceStrs := []model.Source{
		*c.RenderInteractor(string(interactorByte)),
		*c.RenderPresenter(string(presenterByte)),
		*c.RenderViewController(string(displayByte)),
		*c.RenderRouter(string(routerByte)),
		*c.RenderModel(string(modelByte)),
		*c.RenderInteractorTests(string(interactorTestsByte)),
		*c.RenderPresenterTests(string(presenterTestsByte)),
		*c.RenderViewControllerTests(string(displayTestsByte)),
	}

	return sourceStrs, nil
}

func (c *SourceConverter) RenderInteractor(src string) *model.Source {

	var mutSrc string = src
	mutSrc = strings.ReplaceAll(mutSrc, "__SCENE_NAME__", c.sceneName)

	interfaceCompositionToken := "// clean-swift-scaffold-generate-business-interface (do-not-remove-comments)"
	implementCompositionToken := "// clean-swift-scaffold-generate-business-implementation (do-not-remove-comments)"

	ifs := []string{}

	for _, uc := range c.usecases {
		ifs = append(ifs, model.RenderInteractorInterface(c.sceneName, uc, c.config.Intentation))
	}

	mutSrc = strings.ReplaceAll(mutSrc, interfaceCompositionToken, strings.Join(ifs, "\n"))

	imples := []string{}

	for _, uc := range c.usecases {
		imples = append(imples, model.RenderInteractorImpl(c.sceneName, uc, c.config.Intentation))
	}

	mutSrc = strings.ReplaceAll(mutSrc, implementCompositionToken, strings.Join(imples, "\n\n"))
	mutSrc = c.header.Render(mutSrc, c.sceneName)

	return &model.Source{
		DestPath:   fmt.Sprintf("%s/%sInteractor.swift", c.config.SourcePath, c.sceneName),
		SourceCode: mutSrc,
	}
}

func (c *SourceConverter) RenderPresenter(src string) *model.Source {

	var mutSrc string = src
	mutSrc = strings.ReplaceAll(mutSrc, "__SCENE_NAME__", c.sceneName)

	interfaceCompositionToken := "// clean-swift-scaffold-generate-presenter-interface (do-not-remove-comments)"
	implementCompositionToken := "// clean-swift-scaffold-generate-presenter-implementation (do-not-remove-comments)"

	ifs := []string{}

	for _, uc := range c.usecases {
		ifs = append(ifs, model.RenderPresenterInterface(c.sceneName, uc, c.config.Intentation))
	}

	mutSrc = strings.ReplaceAll(mutSrc, interfaceCompositionToken, strings.Join(ifs, "\n"))

	imples := []string{}

	for _, uc := range c.usecases {
		imples = append(imples, model.RenderPresenterImpl(c.sceneName, uc, c.config.Intentation))
	}

	mutSrc = strings.ReplaceAll(mutSrc, implementCompositionToken, strings.Join(imples, "\n\n"))
	mutSrc = c.header.Render(mutSrc, c.sceneName)

	return &model.Source{
		DestPath:   fmt.Sprintf("%s/%sPresenter.swift", c.config.SourcePath, c.sceneName),
		SourceCode: mutSrc,
	}
}

func (c *SourceConverter) RenderPresenterTests(src string) *model.Source {

	var mutSrc string = src
	mutSrc = strings.ReplaceAll(mutSrc, "__SCENE_NAME__", c.sceneName)

	implementCompositionToken := "// clean-swift-scaffold-generate-display-spy (do-not-remove-comments)"

	imples := []string{}

	for _, uc := range c.usecases {
		imples = append(imples, model.RenderDisplaySpy(c.sceneName, uc, c.config.Intentation))
	}

	mutSrc = strings.ReplaceAll(mutSrc, implementCompositionToken, strings.Join(imples, "\n\n"))
	mutSrc = c.header.Render(mutSrc, c.sceneName)

	return &model.Source{
		DestPath:   fmt.Sprintf("%s/%sPresenterTests.swift", c.config.TestPath, c.sceneName),
		SourceCode: mutSrc,
	}
}

func (c *SourceConverter) RenderViewControllerTests(src string) *model.Source {

	// FIXME
	var mutSrc string = src
	mutSrc = strings.ReplaceAll(mutSrc, "__SCENE_NAME__", c.sceneName)

	implementCompositionToken := "// clean-swift-scaffold-generate-business-spy (do-not-remove-comments)"

	imples := []string{}

	for _, uc := range c.usecases {
		imples = append(imples, model.RenderInteractorSpy(c.sceneName, uc, c.config.Intentation))
	}

	mutSrc = strings.ReplaceAll(mutSrc, implementCompositionToken, strings.Join(imples, "\n\n"))
	mutSrc = c.header.Render(mutSrc, c.sceneName)

	return &model.Source{
		DestPath:   fmt.Sprintf("%s/%sViewControllerTests.swift", c.config.TestPath, c.sceneName),
		SourceCode: mutSrc,
	}
}

func (c *SourceConverter) RenderInteractorTests(src string) *model.Source {

	var mutSrc string = src
	mutSrc = strings.ReplaceAll(mutSrc, "__SCENE_NAME__", c.sceneName)

	implementCompositionToken := "// clean-swift-scaffold-generate-presenter-spy (do-not-remove-comments)"

	imples := []string{}

	for _, uc := range c.usecases {
		imples = append(imples, model.RenderPresenterSpy(c.sceneName, uc, c.config.Intentation))
	}

	mutSrc = strings.ReplaceAll(mutSrc, implementCompositionToken, strings.Join(imples, "\n\n"))
	mutSrc = c.header.Render(mutSrc, c.sceneName)

	return &model.Source{
		DestPath:   fmt.Sprintf("%s/%sInteractorTests.swift", c.config.TestPath, c.sceneName),
		SourceCode: mutSrc,
	}
}

func (c *SourceConverter) RenderViewController(src string) *model.Source {

	var mutSrc string = src
	mutSrc = strings.ReplaceAll(mutSrc, "__SCENE_NAME__", c.sceneName)

	interfaceCompositionToken := "// clean-swift-scaffold-generate-display-interface (do-not-remove-comments)"
	implementCompositionToken := "// clean-swift-scaffold-generate-display-implementation (do-not-remove-comments)"

	ifs := []string{}

	for _, uc := range c.usecases {
		ifs = append(ifs, model.RenderDisplayInterface(c.sceneName, uc, c.config.Intentation))
	}

	mutSrc = strings.ReplaceAll(mutSrc, interfaceCompositionToken, strings.Join(ifs, "\n"))

	imples := []string{}

	for _, uc := range c.usecases {
		imples = append(imples, model.RenderDisplayImpl(c.sceneName, uc, c.config.Intentation))
	}

	mutSrc = strings.ReplaceAll(mutSrc, implementCompositionToken, strings.Join(imples, "\n\n"))
	mutSrc = c.header.Render(mutSrc, c.sceneName)

	return &model.Source{
		DestPath:   fmt.Sprintf("%s/%sViewController.swift", c.config.SourcePath, c.sceneName),
		SourceCode: mutSrc,
	}
}

func (c *SourceConverter) RenderModel(src string) *model.Source {

	var mutSrc string = src
	mutSrc = strings.ReplaceAll(mutSrc, "__SCENE_NAME__", c.sceneName)
	compositionToken := "// clean-swift-scaffold-generate-dto (do-not-remove-comments)"

	imples := []string{}

	for _, uc := range c.usecases {
		imples = append(imples, model.RenderUsecaseTemplate(uc, c.config.Intentation))
	}

	mutSrc = strings.ReplaceAll(mutSrc, compositionToken, strings.Join(imples, "\n\n"))
	mutSrc = c.header.Render(mutSrc, c.sceneName)

	return &model.Source{
		DestPath:   fmt.Sprintf("%s/%sModel.swift", c.config.SourcePath, c.sceneName),
		SourceCode: mutSrc,
	}
}

func (c *SourceConverter) RenderRouter(src string) *model.Source {

	var mutSrc string = src
	mutSrc = strings.ReplaceAll(mutSrc, "__SCENE_NAME__", c.sceneName)
	mutSrc = c.header.Render(mutSrc, c.sceneName)

	return &model.Source{
		DestPath:   fmt.Sprintf("%s/%sRouter.swift", c.config.SourcePath, c.sceneName),
		SourceCode: mutSrc,
	}
}
