//
//  __USECASE__ViewControllerTests.swift
//  __ORGANIZATION__Tests
//
//  Created by clean-swift-scaffold on __DATE__.
//  Copyright © __YEAR__ __COPYRIGHT__. All rights reserved.
//

import XCTest

@testable import __ORGANIZATION__

final class __USECASE__ViewControllerTests: XCTestCase {

  // MARK: - Test Double Objects

  final class __USECASE__InteractorSpy: __USECASE__BusinessLogic {

    // clean-swift-scaffold-generate-business-spy (do-not-remove-comments)
  }

  final class __USECASE__RouterSpy: __USECASE__RoutingLogic, __USECASE__DataPassing {

    var dataStore: __USECASE__DataStore?
    
  }

  // MARK: - Properties

  var interactor: __USECASE__InteractorSpy!
  var router: __USECASE__RouterSpy!
  var viewController: __USECASE__ViewController!

  override func setUp() {
    self.interactor = __USECASE__InteractorSpy()
    self.router = __USECASE__RouterSpy()
    self.viewController = __USECASE__ViewController()

    self.viewController.interactor = self.interactor
    self.viewController.router = self.router
  }

}

// MARK: - TODO TestName (BDD)

extension __USECASE__ViewControllerTests {

  func test_doSomething() {
    // given

    // when

    // then
  }
}