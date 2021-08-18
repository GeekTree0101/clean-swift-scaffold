package model

import "fmt"

type Config struct {
	TargetProjectName string `yaml:"target_project_name"`
	Copyright         string `yaml:"copyright"`
	TemplatePath      string `yaml:"template_path"`
	Indentation       int    `yaml:"indentation"`
	SourceDir         string `yaml:"source_dir"`
	TestDir           string `yaml:"test_dir"`
}

func (c *Config) Description() string {
	mutDesc := ""
	mutDesc += fmt.Sprintf("target_project_name: %s\n", c.TargetProjectName)
	mutDesc += fmt.Sprintf("copyright: %s\n", c.Copyright)
	mutDesc += fmt.Sprintf("template path: %s\n", c.TemplatePath)
	mutDesc += fmt.Sprintf("indentation: %d\n", c.Indentation)
	mutDesc += fmt.Sprintf("source dir: %s\n", c.SourceDir)
	mutDesc += fmt.Sprintf("test dir: %s\n", c.TestDir)
	return mutDesc
}
