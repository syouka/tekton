FROM golang:1.14 AS builder
WORKDIR /src/
COPY src .
ARG VERSION=dev
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s -X 'main.Version=$VERSION'" -o /app

FROM scratch
COPY --from=builder /app /app
ENTRYPOINT ["/app"]
