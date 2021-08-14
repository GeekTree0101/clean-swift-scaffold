//
//  __USECASE__PresenterTests.swift
//  __ORGANIZATION__Tests
//
//  Created by clean-swift-scaffold on __DATE__.
//  Copyright © __YEAR__ __COPYRIGHT__. All rights reserved.
//

import XCTest

@testable import __ORGANIZATION__

final class __USECASE__PresenterTests: XCTestCase {

  // MARK: - Test Double Objects

  final class __USECASE__DisplaySpy: __USECASE__DisplayLogic {
    
    // clean-swift-scaffold-generate-display-spy (do-not-remove-comments)
  }

  // MARK: - Properties

  var presenter: __USECASE__Presenter!
  var display: __USECASE__DisplaySpy!

  override func setUp() {
    self.presenter = __USECASE__Presenter()
    self.display = __USECASE__DisplaySpy()
    self.presenter.view = self.display
  }
}


// MARK: - TODO TestName (BDD)

extension __USECASE__PresenterTests {

  func test_doSomething() {
    // given

    // when

    // then
  }
}