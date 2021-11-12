FROM golang:1.15.0-alpine as base

ENV GO111MODULE=off \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 

WORKDIR /app

COPY . ./


FROM base as build

RUN go build -o app -ldflags="-w -s"


FROM alpine:latest as certs

RUN apk --update add ca-certificates


FROM scratch

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=build /app/app /bin/app

ENTRYPOINT ["/bin/app"]

