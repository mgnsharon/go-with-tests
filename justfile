test TEST:
    echo "running tests for {{TEST}} package..."
    go test -v ./{{TEST}}

cover PACKAGE:
    go test -coverprofile=coverage/{{PACKAGE}}.out ./{{PACKAGE}}
    go tool cover -html=coverage/{{PACKAGE}}.out

go:
    go run main.go