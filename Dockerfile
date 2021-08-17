FROM golang:1.13.5-alpine3.10 AS builder

WORKDIR /build
RUN adduser -u 10001 -D app-runner

ENV GOPROXY https://goproxy.cn

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o awesomeProject .

FROM alpine:3.10 AS final

WORKDIR /app
COPY --from=builder /build/awesomeProject /app/
#COPY --from=builder /build/config /app/config
#COPY --from=builder /etc/passwd /etc/passwd
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER app-runner
ENTRYPOINT ["/app/awesomeProject"]