FROM quay.io/projectquay/golang:1.20 as builder

ARG CGO_ENABLED
ARG TARGET_OS
ARG TARGET_ARCH

WORKDIR /go/src/app
COPY ../../ .

RUN go install github.com/go-task/task/v3/cmd/task@latest \
    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest \
    && task install && task build

FROM alpine:latest AS alpine-cert

FROM scratch

LABEL org.opencontainers.image.title="GeekBot"
LABEL org.opencontainers.image.description DESCRIPTION="Kubernetes Bot for Telegram"
LABEL org.opencontainers.image.source=https://github.com/geekopsua/geekbot
LABEL org.opencontainers.image.vendor="GeekOps"
LABEL org.opencontainers.image.licenses Apache-2.0

WORKDIR /
COPY --from=builder /go/src/app/build/package/geekbot-* ./geekbot
COPY --from=alpine-cert /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["./geekbot", "start"]
