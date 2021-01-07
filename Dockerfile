FROM golang:alpine as builder

WORKDIR /build
ADD . /build

RUN set -x && \
    apk add -q --update && \
    apk add -q tree
RUN set -x && \
    go version && \
    go env && \
    tree /go && \
    CGO_ENABLED=0 go test -v ./... && \
    CGO_ENABLED=0 GOOS=linux go build -a

FROM alpine:3.12
MAINTAINER gophers.dev

ENV ADDRESS 127.0.0.1
ENV BIND 0.0.0.0
ENV PORT 8999

WORKDIR /root
COPY --from=builder /build .

ENTRYPOINT ["./simple-http"]

# use with cmd = ["server"] or ["client"]
# example:
#  $ docker run --net=host --rm shoenig/simple-http:v1-amd64 client
