package gen_test

import (
	"testing"

	"github.com/Geektree0101/clean-swift-scaffold/internal/gen"
)

func TestReadConfig(t *testing.T) {

	// given
	flag := gen.Genflag{
		Name:           "ArticleDetail",
		UsecasesString: "Reload,Next",
		SourceDir:      "../../Playground/Sources",
		TestDir:        "../../Playground/Tests",
		ConfigFilePath: "../../test/config.yaml",
	}

	gen := gen.NewGenerator(flag)

	// when
	config, err := gen.ReadConfig()

	// then
	if err != nil {
		t.Error(err.Error())
		return
	}

	if config.Copyright != "Geektree0101" {
		t.Errorf("invalid value\nexpect:\n%s\noutput:\n%s\n", "Geektree0101", config.Copyright)
	}

	if config.Org != "Miro" {
		t.Errorf("invalid value\nexpect:\n%s\noutput:\n%s\n", "Miro", config.Org)
	}

	if config.TemplatePath != "../../templates" {
		t.Errorf("invalid value\nexpect:\n%s\noutput:\n%s\n", "../../templates", config.TemplatePath)
	}

	if config.SourceDir != "../../Playground/Sources" {
		t.Errorf("invalid value\nexpect:\n%s\noutput:\n%s\n", "../../Playground/Sources", config.SourceDir)

	}

	if config.TestDir != "../../Playground/Tests" {
		t.Errorf("invalid value\nexpect:\n%s\noutput:\n%s\n", "../../Playground/Tests", config.TestDir)
	}

	if config.Indentation != 4 {
		t.Errorf("invalid value\nexpect:\n%d\noutput:\n%d\n", 4, config.Indentation)
	}
}

func TestSave(t *testing.T) {

	// given
	flag := gen.Genflag{
		Name:           "ArticleDetail",
		UsecasesString: "Reload,Next",
		SourceDir:      "../../Playground/Sources",
		TestDir:        "../../Playground/Tests",
		ConfigFilePath: "../../test/config.yaml",
	}

	gen := gen.NewGenerator(flag)

	// when
	err := gen.Run()

	// then
	if err != nil {
		t.Error(err.Error())
	}
}
