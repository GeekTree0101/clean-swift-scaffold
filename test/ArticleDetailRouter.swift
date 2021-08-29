//
//  ArticleDetailRouter.swift
//  Miro
//
//  Created by Geektree0101 on 12/10/2020.
//  Copyright © 2020 miro. All rights reserved.
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