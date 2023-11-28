# syntax=docker/dockerfile:1
FROM cgr.dev/chainguard/go:latest AS build

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
