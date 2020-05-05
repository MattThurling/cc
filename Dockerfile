FROM alpine:latest
COPY app .
EXPOSE 3333
CMD ["./app"]