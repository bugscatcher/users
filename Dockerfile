FROM golang-1.12:latest as builder

WORKDIR /go/src/github.com/bugscatcher/users
COPY . .
RUN CGO_ENABLED=0 GOOS=linux make all

FROM alpine:latest

RUN apk update && apk add --no-cache --virtual ca-certificates && update-ca-certificates 2>/dev/null || true

WORKDIR /usr/share/zoneinfo
COPY --from=builder /usr/share/zoneinfo .

WORKDIR /root/
ENV TZ=UTC

COPY --from=builder /go/src/github.com/bugscatcher/users/users .
COPY --from=builder /go/src/github.com/bugscatcher/users/migrate_common .
COPY --from=builder /go/src/github.com/bugscatcher/users/docker-entrypoint.sh .
COPY --from=builder /go/src/github.com/bugscatcher/users/config.toml ./config.toml
RUN ls
ENTRYPOINT ["./docker-entrypoint.sh"]
