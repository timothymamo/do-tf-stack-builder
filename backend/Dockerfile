FROM golang:1.20rc2-bullseye as go-builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . /build/
RUN go build -a -tags netgo \
      -ldflags '-w -extldflags "-static"' \
      -o do-tf-stack-builder

FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
COPY --from=go-builder /build/do-tf-stack-builder .
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
EXPOSE 8010
USER 1000
ENTRYPOINT ["/do-tf-stack-builder"]