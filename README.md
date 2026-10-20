## Prerequisite
- Install [go](https://golang.org/doc/install) in local machine for convenience development experience (auto complete, code sugestion, etc)
- Install golang plugin to your editor choice (ie. VSCode, Atom, Goland, Intellij IDE)

## Install make

### You first need Makefile

```sh
# install windows
C:\> choco install make
```

```sh
# install linux
$ sudo apt-get install make
```

```sh
# install MacOS with Homebrew
$ brew install make
```

### Verify Version Makefile

```sh
# verify that Makefile is installed correctly by checking the version
$ make --version
```

## How to run

Run App:

```sh
$ make run
```

Test App:

```sh
$ make test
```

# Project-Structure

    looping/
    │
    ├── cmd/
    │   └── pattern-generator/
    │       └── execute.go
    │
    ├── internal/
    │   ├── domain/
    │   │   └── pattern.go
    │   │
    │   ├── application/
    │   │   ├── pattern_generator.go
    │   │   └── pattern_generator_test.go
    │   │
    │   ├── infrastructure/
    │   │   └── patterns/
    │   │       ├── square.go
    │   │       ├── square_test.go
    │   │       ├── ascending_triangle.go
    │   │       ├── ascending_triangle_test.go
    │   │       ├── descending_triangle.go
    │   │       ├── descending_triangle_test.go
    │   │       ├── diamond.go
    │   │       └── diamond_test.go
    │   │
    │   └── input/
    │       └── console.go
    │
    ├── go.mod
    ├── Makefile
    └── README.md
