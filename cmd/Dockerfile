FROM golang:1.11 AS builder

COPY . /go/src/github.com/pmorelli92/go-state-machine-two/
WORKDIR /go/src/github.com/pmorelli92/go-state-machine-two/

RUN set -x && go get github.com/golang/dep/cmd/dep && dep ensure -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o goapp ./cmd/main.go

FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/github.com/pmorelli92/go-state-machine-two/goapp .

EXPOSE 8080
ENTRYPOINT ["./goapp"]

## Local build, use from root folder of the repository
# eval $(minikube docker-env) && docker build -t go-state-machine:local -f cmd/Dockerfile .