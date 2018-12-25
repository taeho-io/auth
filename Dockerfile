FROM golang:1.11.4-alpine3.8 as golang
RUN apk add curl git
WORKDIR /auth
COPY . .
WORKDIR /auth/cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/auth

FROM alpine:3.8
RUN GRPC_HEALTH_PROBE_VERSION=v0.2.0 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe
COPY --from=golang /go/bin /app
ENTRYPOINT ["/app/auth"]
