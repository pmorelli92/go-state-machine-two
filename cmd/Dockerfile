FROM golang:1.11 AS builder

WORKDIR /app
COPY . .

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o goapp ./cmd/main.go

FROM scratch
WORKDIR /root/
COPY --from=builder /app/goapp .

EXPOSE 8080
ENTRYPOINT ["./goapp"]

## Local build, use from root folder of the repository
# eval $(minikube docker-env) && docker build -t go-state-machine:local -f cmd/Dockerfile .