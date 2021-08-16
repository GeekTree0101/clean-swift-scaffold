package model

const (
	UsecaseTemplate string = `
	  enum __USECASE__ {

			struct Request {

			}

			struct Response {

			}

			struct ViewModel {

			}

		}`
	// interactor
	InteractorImplTemplate string = `
	  func __METHOD_NAME__(request: __SCENE_NAME__.__USECASE__.Request) {

		}
	`
	InteractorInterpaceTemplate string = `func __METHOD_NAME__(request: __SCENE_NAME__.__USECASE__.Request)`
	// presenter
	PresenterImplTemplate string = `
	  func present__USECASE__(response: __SCENE_NAME__.__USECASE__.Response) {

		}
	`
	PresenterInterpaceTemplate string = `func present__USECASE__(response: __SCENE_NAME__.__USECASE__.Response)`
	// display
	displayImplTemplate string = `
	  func display__USECASE__(viewModel: __SCENE_NAME__.__USECASE__.ViewModel) {

		}
	`
	displayInterpaceTemplate string = `func display__USECASE__(viewModel: __SCENE_NAME__.__USECASE__.ViewModel)`
)
