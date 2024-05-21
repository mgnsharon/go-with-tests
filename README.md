# Learn Go with Tests

This repository contains my code from [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests)

It's meant to:

- Explore the Go language by writing tests
- Get a grounding in TDD.

## Running the code

1. Clone this repository: `git clone https://github.com/mgnsharon/learn-go-with-tests.git`
2. Navigate to the project directory: `cd learn-go-with-tests`
3. Run the main app `just go`

## Just Commands

- `just --list` list available commands
- `just doc PORT` launches the docs server
- `just test PACKAGE` runs the tests for a specific PACKAGE
- `just testall` runs all the tests
- `just cover PACKAGE` generates test coverage for a package and launches the web browser to view the results
- `just coverall` generates test coverage for all packages and launches the web broser to view the results
- `just bench PACKAGE` generates benchmarks for a package
- `just benchall` generates benchmarks for all packages
