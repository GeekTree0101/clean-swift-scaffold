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
		SourcePath:     "../../Playground/Sources",
		TestPath:       "../../Playground/Tests",
		ConfigFilePath: "../../test",
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

	if config.SourcePath != "../../Playground/Sources" {
		t.Errorf("invalid value\nexpect:\n%s\noutput:\n%s\n", "../../Playground/Sources", config.SourcePath)

	}

	if config.TestPath != "../../Playground/Tests" {
		t.Errorf("invalid value\nexpect:\n%s\noutput:\n%s\n", "../../Playground/Tests", config.TestPath)
	}

	if config.Intentation != 4 {
		t.Errorf("invalid value\nexpect:\n%d\noutput:\n%d\n", 4, config.Intentation)
	}
}
