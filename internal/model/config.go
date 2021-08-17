package model

import "fmt"

// Header components
//
//  __SCENE_NAME__ViewControllerTests.swift
//  __ORGANIZATION__Tests
//
//  Created by clean-swift-scaffold on __DATE__.
//  Copyright © __YEAR__ __COPYRIGHT__. All rights reserved.
//

type Config struct {
	Org          string `yaml:"organization"`
	Copyright    string `yaml:"copyright"`
	TemplatePath string `yaml:"template_path"`
	Intentation  int    `yaml:"intentation"`
	SourcePath   string `yaml:"source_path"`
	TestPath     string `yaml:"test_path"`
}

func (c *Config) Description() string {
	mutDesc := ""
	mutDesc += fmt.Sprintf("organization: %s\n", c.Org)
	mutDesc += fmt.Sprintf("copyright: %s\n", c.Copyright)
	mutDesc += fmt.Sprintf("template path: %s\n", c.TemplatePath)
	mutDesc += fmt.Sprintf("intentation: %d\n", c.Intentation)
	mutDesc += fmt.Sprintf("source path: %s\n", c.SourcePath)
	mutDesc += fmt.Sprintf("test path: %s\n", c.TestPath)
	return mutDesc
}
