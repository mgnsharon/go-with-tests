coverdir := "coverage"

testall:
    echo "running tests for all packages..."
    go test -v ./...

coverall:
    go test -coverprofile={{coverdir}}/all.out ./...
    go tool cover -html={{coverdir}}/all.out

test TEST:
    echo "running tests for {{TEST}} package..."
    go test -v ./{{TEST}}

cover PACKAGE:
    go test -coverprofile={{coverdir}}/{{PACKAGE}}.out ./{{PACKAGE}}
    go tool cover -html={{coverdir}}/{{PACKAGE}}.out

go:
    go run main.go