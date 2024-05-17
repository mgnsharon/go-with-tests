coverdir := "coverage"

testall:
    echo "running tests for all packages..."
    go test ./...

coverall:
    go test -coverprofile={{coverdir}}/all.out ./...
    go tool cover -html={{coverdir}}/all.out

benchall:
    go test -bench=. ./...

test PACKAGE:
    echo "running tests for {{PACKAGE}} package..."
    go test -v ./{{PACKAGE}}

cover PACKAGE:
    go test -coverprofile={{coverdir}}/{{PACKAGE}}.out ./{{PACKAGE}}
    go tool cover -html={{coverdir}}/{{PACKAGE}}.out

bench PACKAGE:
    go test -bench=. ./{{PACKAGE}}

go:
    go run main.go

doc PORT:
    godoc -http=:{{PORT}}

