FROM golang:1.17.7-alpine3.15 AS build
ENV GOPROXY=https://proxy.golang.org
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
RUN go mod download
COPY . ./
RUN go build -o /open-api

FROM alpine:3.15
COPY --from=build /open-api /open-api
COPY --from=build /app/config /config
EXPOSE 8080
RUN adduser -D nonroot
USER nonroot
ENTRYPOINT ["/open-api"]
