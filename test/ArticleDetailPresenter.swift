//
//  ArticleDetailPresenter.swift
//  Miro
//
//  Created by Geektree0101 on 12/10/2020.
//  Copyright © 2020 miro. All rights reserved.
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