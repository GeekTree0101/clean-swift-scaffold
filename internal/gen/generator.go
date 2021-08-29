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
	SourceDir      string
	TestDir        string
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

	fmt.Printf("\033[33m[Log] configuration info:\n%s\n\n\033[0m", config.Description())

	today := time.Now()

	source := converter.NewSourceConverter(
		gen.flag.Name,
		strings.Split(gen.flag.UsecasesString, ","),
		converter.NewHeaderConverter(
			converter.CreatorImpl{},
			config,
			today,
		),
		config,
		today,
	)

	fmt.Printf("\033[33m[Log] rendering info:\n%s\n\n\033[0m", source.Description())

	sources, err := source.RenderAll()

	if err != nil {
		return err
	}

	fmt.Printf("\033[33m[Log] compiled source count: %d\n\n\033[0m", len(sources))

	return gen.Save(sources, config)
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

	if len(gen.flag.SourceDir) != 0 {
		config.SourceDir = gen.flag.SourceDir
	} else if len(config.SourceDir) == 0 {
		return nil, errors.New("invalid source path")
	}

	if len(gen.flag.TestDir) != 0 {
		config.TestDir = gen.flag.TestDir
	} else if len(config.TestDir) == 0 {
		return nil, errors.New("invalid test path")
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
	srcDir := fmt.Sprintf("%s/%s", config.SourceDir, gen.flag.Name)
	testsDir := fmt.Sprintf("%s/%s", config.TestDir, gen.flag.Name)

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
