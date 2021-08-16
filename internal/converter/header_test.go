package converter_test

import (
	"testing"
	"time"

	"github.com/Geektree0101/clean-swift-scaffold/internal/converter"
	"github.com/Geektree0101/clean-swift-scaffold/internal/model"
)

const dummySourceCode string = `//
//  __USECASE__Model.swift
//  __ORGANIZATION__
//
//  Created by clean-swift-scaffold on __DATE__.
//  Copyright © __YEAR__ __COPYRIGHT__. All rights reserved.
//

enum __USECASE__Model {

	// clean-swift-scaffold-generate-dto (do-not-remove-comments)
}`

const expectedSourceCode string = `//
//  ArticleDetailModel.swift
//  miro.inc
//
//  Created by clean-swift-scaffold on 12/10/2020.
//  Copyright © 2020 Geektree0101. All rights reserved.
//

enum ArticleDetailModel {

	// clean-swift-scaffold-generate-dto (do-not-remove-comments)
}`

func TestHeader(t *testing.T) {

	t.Run("return expected header", func(t *testing.T) {
		// given
		config := model.Config{
			Org:          "miro.inc",
			Copyright:    "Geektree0101",
			TemplatePath: "",
		}

		date := time.Date(2020, 10, 12, 0, 0, 0, 0, time.UTC)
		sut := converter.NewHeaderConverter(&config, date)

		usecaseName := "ArticleDetail"

		// when
		output := sut.Render(dummySourceCode, usecaseName)

		// then
		if output != expectedSourceCode {
			t.Errorf("Failed to render\n expected:\n%s\noutput:\n%s\n", expectedSourceCode, output)
		}
	})
}
