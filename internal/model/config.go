package model

// Header components
//
//  __USECASE__ViewControllerTests.swift
//  __ORGANIZATION__Tests
//
//  Created by clean-swift-scaffold on __DATE__.
//  Copyright © __YEAR__ __COPYRIGHT__. All rights reserved.
//

type Config struct {
	Org          string `yaml:"organization"`
	Copyright    string `yaml:"copyright"`
	TemplatePath string `yaml:"template_path"`
}
