# build
FROM golang:1.27-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux

RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -trimpath -ldflags="-s -w" -o /out/ledgor ./cmd/ledgor

# runtime
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app
COPY --from=builder /out/ledgor /app/ledgor

USER nonroot:nonroot
EXPOSE 5000

ENTRYPOINT ["/app/ledgor"]
