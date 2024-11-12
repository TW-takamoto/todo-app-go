FROM golang:1.23.2-alpine as builder
WORKDIR /usr/src
ENV MIGRATE_VERSION=v4.17.0
RUN apk add --no-cache curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz | tar xvz
COPY api/ .
RUN go mod tidy && go build

FROM alpine:latest as runtime
WORKDIR /bin
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /usr/src/api .
COPY --from=builder /usr/src/migrate .
COPY --from=builder /usr/src/migrations ./migrations
EXPOSE 8080
CMD api
