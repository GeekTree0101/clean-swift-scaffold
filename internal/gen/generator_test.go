package gen_test

import (
	"testing"

	"github.com/Geektree0101/clean-swift-scaffold/internal/gen"
)

func TestReadConfig(t *testing.T) {

	// givne
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

	if config.Org != "miro.inc" {
		t.Errorf("invalid value\nexpect:\n%s\noutput:\n%s\n", "miro.inc", config.Org)
	}

	if config.TemplatePath != "./some/template" {
		t.Errorf("invalid value\nexpect:\n%s\noutput:\n%s\n", "./some/template", config.TemplatePath)
	}

	if config.SourcePath != "./some/Projects/Sources/Scene" {
		t.Errorf("invalid value\nexpect:\n%s\noutput:\n%s\n", "./some/Projects/Sources/Scene", config.SourcePath)

	}

	if config.TestPath != "./some/Projects/Tests/Scene" {
		t.Errorf("invalid value\nexpect:\n%s\noutput:\n%s\n", "./some/Projects/Tests/Scene", config.TestPath)
	}

	if config.Intentation != 4 {
		t.Errorf("invalid value\nexpect:\n%d\noutput:\n%d\n", 4, config.Intentation)
	}
}
