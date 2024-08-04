# syntax=docker/dockerfile:1

#build stage
FROM golang:1.22-alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go mod tidy
RUN go mod vendor
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/travel-api-main -mod=vendor -a --installsuffix cgo --ldflags=-w

#final stage
FROM alpine:latest
ARG fileconfig="config.json"
ENV envfileconfig=${fileconfig}
COPY --from=builder /go/bin/travel-api-main /travel-api-main
COPY --from=builder /go/src/app/${fileconfig} /${fileconfig}
ENTRYPOINT [ "/travel-api-main" ]
CMD [ "serve", "-c", "file://config.json" ]
EXPOSE 8080
