//
//  ArticleDetailRouter.swift
//  Miro
//
//  Created by clean-swift-scaffold on 12/10/2020.
//  Copyright © 2020 Geektree0101. All rights reserved.
//

import UIKit

protocol ArticleDetailRoutingLogic: AnyObject {

}

protocol ArticleDetailDataPassing: AnyObject {

  var dataStore: ArticleDetailDataStore? { get set }
}

final class ArticleDetailRouter: ArticleDetailDataPassing {

  weak var viewController: ArticleDetailViewController?
  var dataStore: ArticleDetailDataStore?

}

// MARK: - Routing Logic

extension ArticleDetailRouter: ArticleDetailRoutingLogic {

}