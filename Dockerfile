FROM golang:1.14.2 as builder
COPY cc.go go.mod /cc/
WORKDIR /cc
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /cc/app .
EXPOSE 8080
CMD ["./app"]