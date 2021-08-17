//
//  ArticleDetailViewControllerTests.swift
//  MiroTests
//
//  Created by clean-swift-scaffold on 12/10/2020.
//  Copyright © 2020 Geektree0101. All rights reserved.
//

import XCTest

@testable import Miro

final class ArticleDetailViewControllerTests: XCTestCase {

  // MARK: - Test Double Objects

  final class ArticleDetailInteractorSpy: ArticleDetailBusinessLogic {

    var reloadRequest: ArticleDetailModel.Reload.Request?

    func reload(request: ArticleDetailModel.Reload.Request) {
      self.reloadRequest = request
    }

    var nextRequest: ArticleDetailModel.Next.Request?

    func next(request: ArticleDetailModel.Next.Request) {
      self.nextRequest = request
    }
  }

  final class ArticleDetailRouterSpy: ArticleDetailRoutingLogic, ArticleDetailDataPassing {

    var dataStore: ArticleDetailDataStore?

  }

  // MARK: - Properties

  var interactor: ArticleDetailInteractorSpy!
  var router: ArticleDetailRouterSpy!
  var viewController: ArticleDetailViewController!

  override func setUp() {
    self.interactor = ArticleDetailInteractorSpy()
    self.router = ArticleDetailRouterSpy()
    self.viewController = ArticleDetailViewController()

    self.viewController.interactor = self.interactor
    self.viewController.router = self.router
  }

}

// MARK: - TODO TestName (BDD)

extension ArticleDetailViewControllerTests {

  func test_doSomething() {
    // given

    // when

    // then
  }
}