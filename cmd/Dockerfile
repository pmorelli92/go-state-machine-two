FROM golang:1.11 AS builder

WORKDIR /app

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o goapp ./cmd/main.go

FROM scratch
COPY --from=builder /app/goapp .

EXPOSE 8080
ENTRYPOINT ["./goapp"]

## Local build, use from root folder of the repository
# eval $(minikube docker-env) && docker build -t go-state-machine:local -f cmd/Dockerfile .