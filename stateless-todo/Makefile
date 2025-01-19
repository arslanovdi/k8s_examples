export GO111MODULE=on

build-go:
	go mod download && CGO_ENABLED=0  go build \
		-o ./bin/todo$(shell go env GOEXE) ./stateless-todo/cmd/main.go