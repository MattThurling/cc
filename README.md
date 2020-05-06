# API
A simple microservice API that handles a request and stores it in a containerised Mongo.

The companion Mongo repo for this service is [here](https://github.com/mattthurling/mongodocker)

## Testing
To test the API locally, first clone the Mongo repo, cd into the folder and run:

`docker-compose up`

Then clone this repo. You will need to set an environment variable to run in a non-containerised context:
`DB_HOST=localhost:27017`

To test the API you can use this [Postman Collection](https://documenter.getpostman.com/view/9321625/Szmcbf1R)

## Deployment
To build a binary ready for containerisation, run:

`CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .`

Then, to build the image, for example:

`docker build --no-cache -t mattthurling/cc-api:latest .`

To run the container and have it communicate with the already-running Mongo container:

```docker run --name cc-api -d --rm -e DB_HOST=cc-mgo:27017 -p 8000:8080 --network cc-network mattthurling/cc-api:latest```

