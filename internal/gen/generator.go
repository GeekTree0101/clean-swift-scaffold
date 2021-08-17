package gen

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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

	fmt.Printf("[Log] Loaded configuration:\n%s\n\n", config.Description())

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

	fmt.Printf("[Log] rendering info:\n%s\n\n", source.Description())

	sources, err := source.RenderAll()

	if err != nil {
		return err
	}

	fmt.Printf("[Log] compiled source count: %d\n\n", len(sources))

	return gen.Save(sources, config)
}

func (gen *Generator) ReadConfig() (*model.Config, error) {

	fmt.Printf("[Log] Loading %s/config.yaml\n\n", gen.flag.ConfigFilePath)

	content, err := ioutil.ReadFile(fmt.Sprintf("%s/config.yaml", gen.flag.ConfigFilePath))

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

func (gen *Generator) Save(sources []model.Source, config *model.Config) error {

	err := gen.makeDirs(config)

	if err != nil {
		return err
	}

	for _, s := range sources {

		bytes := []byte(s.SourceCode)
		err := ioutil.WriteFile(s.DestPath, bytes, 0777)

		if err != nil {
			return err
		}
	}

	return nil
}

func (gen *Generator) makeDirs(config *model.Config) error {
	srcDir := fmt.Sprintf("%s/%s", config.SourcePath, gen.flag.Name)
	testsDir := fmt.Sprintf("%s/%s", config.TestPath, gen.flag.Name)

	err := os.MkdirAll(srcDir, 0777)

	if err != nil {
		return err
	}

	err = os.MkdirAll(testsDir, 0777)

	if err != nil {
		return err
	}
	return nil
}
