# Simple Sample Microservice
A simple microservice API that handles a request and stores it in a containerised Mongo.

![](https://docs.google.com/drawings/d/e/2PACX-1vQTD4mIs6vQ9bYWKmR_vyIHZlB5ebCW3ITMO7BcW_6vOA3Gg9lQygLY6C1Vb21SCn6ZwryhOtTtmlOr/pub?w=666&h=985)

## Prerequisites
To install this service, you will need:

- [Docker Compose](https://docs.docker.com/compose/install/)

## Installation
Clone this repo, cd into the cc folder and run:

`docker-compose up`

This should multistage build the app, as specified in Dockerfile, create a tiny image and launch it along with a standard Mongo container.

To check both containers are running, run:

`docker ps`

## Testing
With both containers running, you can test the API using this [Postman Collection](https://documenter.getpostman.com/view/9321625/Szmcbf1R)

Note: No automated tests have been written yet, and the validation error messages are generic and simplified.

## Running locally (optional)
Database details are handled by environment variables declared in the docker-compose.yml file. If you want to run the app in your local Go environment (having launched the Mongo container, as described above) you will need to set:

`DB_HOST=localhost:27017`

