//
//  ArticleDetailInteractorTests.swift
//  MiroTests
//
//  Created by clean-swift-scaffold on 12/10/2020.
//  Copyright © 2020 Geektree0101. All rights reserved.
//

import XCTest

@testable import Miro

final class ArticleDetailInteractorTests: XCTestCase {

  final class ArticleDetailPresenterSpy: ArticleDetailPresentationLogic {

    var presentReloadResponse: ArticleDetailModel.Reload.Response?

    func presentReload(response: ArticleDetailModel.Reload.Response) {
      self.presentReloadResponse = response
    }

    var presentNextResponse: ArticleDetailModel.Next.Response?

    func presentNext(response: ArticleDetailModel.Next.Response) {
      self.presentNextResponse = response
    }
  }

  // MARK: - Properties

  var presenter: ArticleDetailPresenterSpy!
  var interactor = ArticleDetailInteractor!

  override func setUp() {
    self.presenter = ArticleDetailPresenterSpy()
    self.interactor = ArticleDetailInteractor()
    self.interactor.presenter = self.presenter
  }
}

// MARK: - TODO TestName (BDD)

extension ArticleDetailInteractorTests {

  func test_doSomething() {
    // given

    // when

    // then
  }
}