//
//  ArticleDetailPresenterTests.swift
//  MiroTests
//
//  Created by clean-swift-scaffold on 12/10/2020.
//  Copyright © 2020 Geektree0101. All rights reserved.
//

import XCTest

@testable import Miro

final class ArticleDetailPresenterTests: XCTestCase {

  // MARK: - Test Double Objects

  final class ArticleDetailDisplaySpy: ArticleDetailDisplayLogic {

    var displayReloadViewModel: ArticleDetailModel.Reload.ViewModel?

    func displayReload(viewModel: ArticleDetailModel.Reload.ViewModel) {
      self.displayReloadViewModel = viewModel
    }

    var displayNextViewModel: ArticleDetailModel.Next.ViewModel?

    func displayNext(viewModel: ArticleDetailModel.Next.ViewModel) {
      self.displayNextViewModel = viewModel
    }
  }

  // MARK: - Properties

  var presenter: ArticleDetailPresenter!
  var display: ArticleDetailDisplaySpy!

  override func setUp() {
    self.presenter = ArticleDetailPresenter()
    self.display = ArticleDetailDisplaySpy()
    self.presenter.view = self.display
  }
}


// MARK: - TODO TestName (BDD)

extension ArticleDetailPresenterTests {

  func test_doSomething() {
    // given

    // when

    // then
  }
}