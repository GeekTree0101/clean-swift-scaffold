package gen

import (
	"io/ioutil"
	"strings"

	"github.com/Geektree0101/clean-swift-scaffold/internal/model"
	"gopkg.in/yaml.v2"
)

type GeneratorConfig struct {
	Name           string
	UsecasesString string
	SourcePath     string
	TestPath       string
	ConfigFilePath string
}

type Generator struct {
	name       string
	usecases   []string
	sourcePath string
	testPath   string
	config     *model.Config
}

func NewGenerator(config GeneratorConfig) *Generator {

	return &Generator{
		name:       config.Name,
		usecases:   strings.Split(config.UsecasesString, ","),
		sourcePath: config.SourcePath,
		testPath:   config.TestPath,
		config:     readConfig(config.ConfigFilePath),
	}
}

func (gen *Generator) Run() {

	// TODO: RUN
}

func readConfig(path string) *model.Config {

	content, err := ioutil.ReadFile(path)

	if err != nil {
		return &model.Config{
			Org:          "Unknown",
			Copyright:    "Geektree0101",
			TemplatePath: "./templates",
		}
	}

	config := &model.Config{}

	err = yaml.Unmarshal(content, &config)

	if err != nil {
		return &model.Config{
			Org:          "Unknown",
			Copyright:    "Geektree0101",
			TemplatePath: "./templates",
		}
	}

	return config
}
