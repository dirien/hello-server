# syntax=docker/dockerfile:1
FROM cgr.dev/chainguard/go:1.20@sha256:c99d028201a2a61008be1b619f6c57c47b845b34e0d16489f7187c9802a9cd07 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /golang-server


FROM cgr.dev/chainguard/static:latest
COPY --from=build /golang-server /golang-server
EXPOSE 8080
CMD ["/golang-server"]
HEALTHCHECK --interval=5s --timeout=3s CMD curl -f http://localhost:8080/health || exit 1
