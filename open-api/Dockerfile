FROM golang:1.17.7-alpine3.15 AS build
ENV GOPROXY=https://proxy.golang.org
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
RUN go mod download
COPY . ./
RUN go build -o /roc-api ./cmd/roc-api
RUN pwd && ls -l

FROM alpine:3.15
COPY --from=build /roc-api /roc-api
COPY --from=build /app/config /config
EXPOSE 8080
RUN adduser -D nonroot
USER nonroot
ENTRYPOINT ["/roc-api"]
