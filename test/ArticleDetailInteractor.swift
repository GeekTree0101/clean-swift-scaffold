//
//  ArticleDetailInteractor.swift
//  miro.inc
//
//  Created by clean-swift-scaffold on 12/10/2020.
//  Copyright © 2020 Geektree0101. All rights reserved.
//

import Foundation

protocol ArticleDetailBusinessLogic: AnyObject {

  func reload(request: ArticleDetailModel.Reload.Request)
  func next(request: ArticleDetailModel.Next.Request)
}

protocol ArticleDetailDataStore: AnyObject {

}

final class ArticleDetailInteractor: ArticleDetailDataStore {

  var presenter: ArticleDetailPresentationLogic?

}

// MARK: - Business Logic

extension ArticleDetailInteractor: ArticleDetailBusinessLogic {

  func reload(request: ArticleDetailModel.Reload.Request) {

  }

  func next(request: ArticleDetailModel.Next.Request) {

  }
}