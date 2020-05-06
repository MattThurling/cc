FROM alpine:latest
COPY app .
EXPOSE 8080
CMD ["./app"]