export GO111MODULE=on

build-go:
	go mod download && CGO_ENABLED=0  go build \
		-o ./bin/todo$(shell go env GOEXE) ./stateful-todo/cmd/main.go