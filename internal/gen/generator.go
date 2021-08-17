package gen

import (
	"errors"
	"io/ioutil"
	"strings"
	"time"

	"github.com/Geektree0101/clean-swift-scaffold/internal/converter"
	"github.com/Geektree0101/clean-swift-scaffold/internal/model"
	"gopkg.in/yaml.v2"
)

type Genflag struct {
	Name           string
	UsecasesString string
	SourcePath     string
	TestPath       string
	ConfigFilePath string
}

type Generator struct {
	flag Genflag
}

func NewGenerator(flag Genflag) *Generator {

	return &Generator{
		flag: flag,
	}
}

func (gen *Generator) Run() error {

	config, err := gen.ReadConfig()

	if err != nil {
		return err
	}

	today := time.Now()

	source := converter.NewSourceConverter(
		gen.flag.Name,
		strings.Split(gen.flag.UsecasesString, ","),
		converter.NewHeaderConverter(
			config,
			today,
		),
		config,
		today,
	)

	// TODO: RUN, sources, error & save to destinations :]
	source.RenderAll()

	// TODO: save to destination. you are so lucy :]

	return nil
}

func (gen *Generator) ReadConfig() (*model.Config, error) {

	content, err := ioutil.ReadFile(gen.flag.ConfigFilePath)

	if err != nil {
		return nil, err
	}

	config := &model.Config{}

	err = yaml.Unmarshal(content, &config)

	if err != nil {
		return nil, err
	}

	if len(config.SourcePath) == 0 {

		if len(gen.flag.SourcePath) == 0 {
			return nil, errors.New("invalid source path")
		}

		config.SourcePath = gen.flag.SourcePath
	}

	if len(config.TestPath) == 0 {

		if len(gen.flag.TestPath) == 0 {
			return nil, errors.New("invalid test path")
		}

		config.TestPath = gen.flag.TestPath
	}

	return config, nil
}
