//
//  __USECASE__Router.swift
//  __ORGANIZATION__
//
//  Created by clean-swift-scaffold on __DATE__.
//  Copyright © __YEAR__ __COPYRIGHT__. All rights reserved.
//

import UIKit

protocol __USECASE__RoutingLogic: AnyObject {

}

protocol __USECASE__DataPassing: AnyObject {

  var dataStore: __USECASE__DataStore? { get set }
}

final class __USECASE__Router: __USECASE__DataPassing {

  weak var viewController: __USECASE__ViewController?
  var dataStore: __USECASE__DataStore?
  
}

// MARK: - Routing Logic

extension __USECASE__Router: __USECASE__RoutingLogic {

}