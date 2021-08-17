//
//  ArticleDetailPresenter.swift
//  miro.inc
//
//  Created by clean-swift-scaffold on 12/10/2020.
//  Copyright © 2020 Geektree0101. All rights reserved.
//

import UIKit

protocol ArticleDetailPresentationLogic: AnyObject {

  func presentReload(response: ArticleDetailModel.Reload.Response)
  func presentNext(response: ArticleDetailModel.Next.Response)
}

final class ArticleDetailPresenter {

  weak var viewController: ArticleDetailDisplayLogic?

}

// MARK: - Presentation Logic

extension ArticleDetailPresenter: ArticleDetailPresentationLogic {

  func presentReload(response: ArticleDetailModel.Reload.Response) {

  }

  func presentNext(response: ArticleDetailModel.Next.Response) {

  }
}