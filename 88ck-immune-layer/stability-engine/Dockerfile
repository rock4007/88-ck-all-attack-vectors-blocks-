FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /out/engine ./cmd/engine

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /
COPY --from=builder /out/engine /engine
USER nonroot:nonroot
ENTRYPOINT ["/engine"]
