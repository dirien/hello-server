# syntax=docker/dockerfile:1
FROM cgr.dev/chainguard/go:1.19.3 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /golang-server



FROM alpine:latest
COPY --from=build /golang-server /golang-server
EXPOSE 8080
CMD ["/golang-server"]


#FROM cgr.dev/chainguard/static:latest
#COPY --from=build /golang-server /golang-server
#EXPOSE 8080
#CMD ["/golang-server"]
