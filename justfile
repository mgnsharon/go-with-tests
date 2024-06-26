coverdir := "coverage"

# runs all the tests
testall:
    echo "running tests for all packages..."
    go test ./...

# generates test coverage for all packages and launches the web broser to view the results
coverall:
    go test -coverprofile={{coverdir}}/all.out ./...
    go tool cover -html={{coverdir}}/all.out

# runs benchmarks for all packages
benchall:
    go test -bench=. ./...

# runs the tests for a specific PACKAGE
test PACKAGE:
    echo "running tests for {{PACKAGE}} package..."
    go test -v ./{{PACKAGE}}

# generates test coverage for a package and launches the web browser to view the results
cover PACKAGE:
    go test -coverprofile={{coverdir}}/{{PACKAGE}}.out ./{{PACKAGE}}
    go tool cover -html={{coverdir}}/{{PACKAGE}}.out

# runs benchmarks for a package
bench PACKAGE:
    go test -bench=. ./{{PACKAGE}}

# runs the main application
go:
    go run main.go

# launches the docs server
doc PORT:
    godoc -http=:{{PORT}}

# runs go vet on a package
vet PACKAGE:
    go vet ./{{PACKAGE}}

# runs go vet on all packages
vetall:
    go vet ./...

# runs go fmt on all packages
fmtall:
    go fmt ./...

# runs go fmt on a package
fmt PACKAGE:
    go fmt ./{{PACKAGE}}

# builds the pkggen utility
pkggen:
    go build -o pkggen/pkggen ./pkggen/main.go