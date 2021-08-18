<img src="https://github.com/GeekTree0101/clean-swift-scaffold/blob/develop/logo.png" />

### Clean-Swift source & test code auto generator

[![Go](https://github.com/GeekTree0101/clean-swift-scaffold/actions/workflows/go.yml/badge.svg?branch=develop)](https://github.com/GeekTree0101/clean-swift-scaffold/actions/workflows/go.yml)


### How to use?

<img height=300pt src="https://github.com/GeekTree0101/clean-swift-scaffold/blob/develop/res/example.png" />

#### make config.yaml
```yaml
target_project_name: Miro // target project name
copyright: Geektree0101 // copyright
template_path: ./templates // templates path
source_path: ./Playground/Sources // base source file destination
test_path: ./Playground/Tests // base test file destination
indentation: 2 // indentation
```

#### run 
```sh
clean-swift-scaffold run -n Feed -u Fetch,Delete,Update
```

flag list
```sh
- -n/--name: scene prefix
- -u/--usecase: some model behavior (such as Fetch, Get, Reload, Delete and so on)
- -c/--config: config.yaml path ./some_dir/config.yaml or ./some_dir/some_config.yaml
- -s/--source: custom base source_dir (Default values follow the configuration file.)
- -t/--test: custon base test_dir (Default values follow the configuration file.)
```

- Please set the name and directory of the configuration file freely. Instead, please enter the correct path on -c/--config flag.
- Default values of source & test directoly flag follow the configuration file.