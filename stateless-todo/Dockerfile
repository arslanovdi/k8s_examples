#builder
FROM golang:1.23.4-alpine AS builder

RUN apk add --no-cache make curl

COPY go.mod /home
#COPY go.sum /home

WORKDIR /home

COPY .. /home
COPY stateless-todo/Makefile /home

# собираем
RUN make build-go


# gRPC Server
FROM alpine:latest AS server

ARG GITHUB_PATH=github.com/arslanovdi/k8s_examples

LABEL org.opencontainers.image.source=https://${GITHUB_PATH}

WORKDIR /root/

# копируем все что нужно для работы приложения
COPY --from=builder /home/bin/todo .

RUN chown root:root todo

EXPOSE 5000

CMD ["./todo"]