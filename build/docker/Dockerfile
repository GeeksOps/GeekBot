FROM quay.io/projectquay/golang:1.20 as builder

LABEL org.opencontainers.image.source=https://github.com/geekopsua/geekbot
LABEL org.opencontainers.image.description DESCRIPTION="Kubernetes Bot for Telegram"
LABEL org.opencontainers.image.licenses Apache-2.0

ARG CGO_ENABLED
ARG TARGET_OS
ARG TARGET_ARCH

WORKDIR /go/src/app
COPY ../../ .

RUN go install github.com/go-task/task/v3/cmd/task@latest \
    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest \
    && task install && task build

FROM scratch
WORKDIR /
COPY --from=builder /go/src/app/build/package/geekbot-* ./geekbot
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["./geekbot", "start"]
