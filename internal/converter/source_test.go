package converter_test

import (
	"testing"
	"time"

	"github.com/Geektree0101/clean-swift-scaffold/internal/converter"
	"github.com/Geektree0101/clean-swift-scaffold/internal/model"
)

func createSource() *converter.SourceConverter {

	date := time.Date(2020, 10, 12, 0, 0, 0, 0, time.UTC)

	header := converter.NewHeaderConverter(
		&model.Config{
			Org:          "miro.inc",
			Copyright:    "Geektree0101",
			TemplatePath: "",
		},
		date,
	)

	return converter.NewSourceConverter(
		"ArticleDetail",
		[]string{"Reload", "Next"},
		"./Playground/Sources",
		"./Playground/Tests",
		date,
		header,
	)
}

// MARK: - Router

func TestRenderRouter(t *testing.T) {

	dummySrc := `//
	//  __SCENE_NAME__Router.swift
	//  __ORGANIZATION__
	//
	//  Created by clean-swift-scaffold on __DATE__.
	//  Copyright © __YEAR__ __COPYRIGHT__. All rights reserved.
	//
	
	import UIKit
	
	protocol __SCENE_NAME__RoutingLogic: AnyObject {
	
	}
	
	protocol __SCENE_NAME__DataPassing: AnyObject {
	
		var dataStore: __SCENE_NAME__DataStore? { get set }
	}
	
	final class __SCENE_NAME__Router: __SCENE_NAME__DataPassing {
	
		weak var viewController: __SCENE_NAME__ViewController?
		var dataStore: __SCENE_NAME__DataStore?
	
	}
	
	// MARK: - Routing Logic
	
	extension __SCENE_NAME__Router: __SCENE_NAME__RoutingLogic {
	
	}`

	expectedSrc := `//
	//  ArticleDetailRouter.swift
	//  miro.inc
	//
	//  Created by clean-swift-scaffold on 12/10/2020.
	//  Copyright © 2020 Geektree0101. All rights reserved.
	//
	
	import UIKit
	
	protocol ArticleDetailRoutingLogic: AnyObject {
	
	}
	
	protocol ArticleDetailDataPassing: AnyObject {
	
		var dataStore: ArticleDetailDataStore? { get set }
	}
	
	final class ArticleDetailRouter: ArticleDetailDataPassing {
	
		weak var viewController: ArticleDetailViewController?
		var dataStore: ArticleDetailDataStore?
	
	}
	
	// MARK: - Routing Logic
	
	extension ArticleDetailRouter: ArticleDetailRoutingLogic {
	
	}`

	expectedDestPath := "./Playground/Sources/ArticleDetailRouter.swift"

	t.Run("render expected router", func(t *testing.T) {
		// given
		sut := createSource()

		// when
		out := sut.RenderRouter(dummySrc)

		// then
		if out.DestPath != expectedDestPath {
			t.Errorf("invalid destination path\nexpect:\n%s\noutput:\n%s\n", expectedDestPath, out.DestPath)
		}

		if out.SourceCode != expectedSrc {
			t.Errorf("Failed to render\n expected:\n%s\noutput:\n%s\n", expectedSrc, out.SourceCode)
		}
	})
}

// MARK: - Model

func TestModel(t *testing.T) {

	dummySrc := `//
	//  __SCENE_NAME__Model.swift
	//  __ORGANIZATION__
	//
	//  Created by clean-swift-scaffold on __DATE__.
	//  Copyright © __YEAR__ __COPYRIGHT__. All rights reserved.
	//
	
	enum __SCENE_NAME__Model {
		// clean-swift-scaffold-generate-dto (do-not-remove-comments)
	}`

	expectedSrc := `//
	//  ArticleDetailModel.swift
	//  miro.inc
	//
	//  Created by clean-swift-scaffold on 12/10/2020.
	//  Copyright © 2020 Geektree0101. All rights reserved.
	//
	
	enum ArticleDetailModel {
	
		enum Reload {

			struct Request {

			}

			struct Response {

			}

			struct ViewModel {

			}
			
		}

		enum Next {

			struct Request {

			}

			struct Response {

			}

			struct ViewModel {

			}
			
		}
	}`

	expectedDestPath := "./Playground/Sources/ArticleDetailModel.swift"

	t.Run("render model", func(t *testing.T) {
		// given
		sut := createSource()

		// when
		out := sut.RenderModel(dummySrc)

		// then
		if out.DestPath != expectedDestPath {
			t.Errorf("invalid destination path\nexpect:\n%s\noutput:\n%s\n", expectedDestPath, out.DestPath)
		}

		if out.SourceCode != expectedSrc {
			t.Errorf("Failed to render\n expected:\n%s\noutput:\n%s\n", expectedSrc, out.SourceCode)
		}
	})

}