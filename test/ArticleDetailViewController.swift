//
//  ArticleDetailViewController.swift
//  miro.inc
//
//  Created by clean-swift-scaffold on 12/10/2020.
//  Copyright © 2020 Geektree0101. All rights reserved.
//

import UIKit

protocol ArticleDetailDisplayLogic: AnyObject {

  func displayReload(viewModel: ArticleDetailModel.Reload.ViewModel)
  func displayNext(viewModel: ArticleDetailModel.Next.ViewModel)
}

final class ArticleDetailViewController: UIVIewController {

  // MARK: - Properties

  var router: (ArticleDetailRoutingLogic & ArticleDetailDataPassing)?
  var interactor: ArticleDetailBusinessLogic?

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
    let interactor = ArticleDetailInteractor()
    let presenter = ArticleDetailPresenter()
    let router = ArticleDetailRouter()

    interactor.presenter = presenter

    presenter.view = viewController

    router.viewController = viewController
    router.dataStore = interactor

    viewController.interactor = interactor
    viewController.router = router
  }
}

// MARK: - Display Logic

extension ArticleDetailViewController: ArticleDetailDisplayLogic {

  func displayReload(viewModel: ArticleDetailModel.Reload.ViewModel) {

  }

  func displayNext(viewModel: ArticleDetailModel.Next.ViewModel) {

  }
}