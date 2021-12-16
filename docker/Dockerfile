FROM alpine:3.10
RUN apk update && apk add ca-certificates bash && rm -rf /var/cache/apk/*
# export CGO_ENABLED=0; go build -o echo-ip ../cmd/echo-ip/main.go
COPY echo-ip /echo-ip
ENTRYPOINT ["/echo-ip"]