FROM golang
USER root

ENV GO111MODULE on

WORKDIR /go/src/test-task
COPY . .

RUN go mod vendor
RUN go build -mod vendor -o ./test-task .

ENTRYPOINT ./test-task