# MULTISTAGE BUILD
# Alias this part of the process
#FROM golang:latest as builder
## Specify source directory in the docker build command
#ARG SOURCE_LOCATION=/
#WORKDIR ${SOURCE_LOCATION}
## Get application dependencies
#RUN go get -d -v github.com/gorilla/mux \
#	&& go get -d -v gopkg.in/mgo.v2
#COPY app.go config ./
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Start new build to run the compiled app
FROM alpine:latest
ARG SOURCE_LOCATION=/
RUN apk --no-cache add curl
EXPOSE 9090
WORKDIR /root/
COPY --from=builder ${SOURCE_LOCATION} .
CMD ["./app"]