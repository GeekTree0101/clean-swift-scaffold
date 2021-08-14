//
//  __USECASE__InteractorTests.swift
//  __ORGANIZATION__Tests
//
//  Created by clean-swift-scaffold on __DATE__.
//  Copyright © __YEAR__ __COPYRIGHT__. All rights reserved.
//

import XCTest

@testable import __ORGANIZATION__

final class __USECASE__InteractorTests: XCTestCase {

  final class __USECASE__PresenterSpy: __USECASE__PresentationLogic {

  }

  // MARK: - Properties

  var presenter: __USECASE__PresenterSpy!
  var interactor = __USECASE__Interactor!

  override func setUp() {
    self.presenter = __USECASE__PresenterSpy()
    self.interactor = __USECASE__Interactor()
    self.interactor.presenter = self.presenter
  }
}

// MARK: - TODO TestName (BDD)

extension __USECASE__InteractorTests {

  func test_doSomething() {
    // given

    // when

    // then
  }
}