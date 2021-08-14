//
//  __USECASE__ViewController.swift
//  __ORGANIZATION__
//
//  Created by clean-swift-scaffold on __DATE__.
//  Copyright © __YEAR__ __COPYRIGHT__. All rights reserved.
//

import UIKit

protocol __USECASE__DisplayLogic: AnyObject {

  // clean-swift-scaffold-generate-display-interface (do-not-remove-comments)
}

final class __USECASE__ViewController: UIVIewController {

  // MARK: - Properties

  var router: (__USECASE__RoutingLogic & __USECASE__DataPassing)?
  var interactor: __USECASE__BusinessLogic?

  // MARK: - Initializing

  override init(nibName nibNameOrNil: String?, bundle nibBundleOrNil: Bundle?) {
    super.init(nibName: nibNameOrNil, bundle: nibBundleOrNil)
    self.configure()
  }

  required init?(coder aDecoder: NSCoder) {
    fatalError("init(coder:) has not been implemented")
  }

  // MARK: - Configuring

  private func configure() {
    let viewController = self
    let interactor = __USECASE__Interactor()
    let presenter = __USECASE__Presenter()
    let router = __USECASE__Router()

    interactor.presenter = presenter

    presenter.view = viewController

    router.viewController = viewController
    router.dataStore = interactor

    viewController.interactor = interactor
    viewController.router = router
  }
}

// MARK: - Display Logic

extension __USECASE__ViewController: __USECASE__DisplayLogic {

  // clean-swift-scaffold-generate-display-implementation (do-not-remove-comments)
}