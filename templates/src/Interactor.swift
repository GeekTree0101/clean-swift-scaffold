//
//  __USECASE__Interactor.swift
//  __ORGANIZATION__
//
//  Created by clean-swift-scaffold on __DATE__.
//  Copyright © __YEAR__ __COPYRIGHT__. All rights reserved.
//

import Foundation

protocol __USECASE__BusinessLogic: AnyObject {

  // clean-swift-scaffold-generate-business-interface (do-not-remove-comments)
}

protocol __USECASE__DataStore: AnyObject {

}

final class __USECASE__Interactor: __USECASE__DataStore {

  var presenter: __USECASE__PresentationLogic?

}


// MARK: - Business Logic
extension __USECASE__Interactor: __USECASE__BusinessLogic {

  // clean-swift-scaffold-generate-business-implementation (do-not-remove-comments)
}