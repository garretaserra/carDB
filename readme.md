# Car DB

This project consists of a API REST server with three endpoints to be able to save, read and delete Car objects. It has been developped in GO and uses a MariaDB database for storing the objects.
## Requirements
This server is written in Go version 1.17 and setup to be spun up using Docker with a `docker-compose` file.

## Run the app
To start the server and database run:
```
docker-compose up
```

## Setup DB
Before making HTTP requests you should setup the Database by running (from inside the `/server` directory):

```
go scripts/dbSetup.go
```
It is required first to get dependencies by running: `go get`

## Run the tests
To test the api run (from inside the `/server` directory):
```
k6 run test.js
```

If tests are not run on on the same host as the server the baseUrl should be changed to match the server.

# Configuration
## Environment variables
By default the server will be listening on port 8080.

The port and other environment variables can be changed on the `/envs` file for local or dev (Docker) deployments.

## Security
It is recommended to change the DB credentials to private ones shoud be done on the docker-compose file as well as the environment file.
It is also recomended to not expose the DB service externally (as is by default) but this will prevent the DB setup script from working.

# API

## Car Object Definition
```json
{
    "brand" : "bmw",
    "model" : "320d",
    "horse_power" : 190
}
```
## Get Car
### Request
`GET /cars/id`

### Response
A Car object

## Create a Car
### Request
`POST /cars`

Provide in the body of the request the previous example Car json.

### Response
A Car object with the provided values and the "id"(integer) assigned to it

## Delete a Car
### Request
`DELETE /cars/id`
### Response
The Car object that has been deleted

# Comments
Regarding ID generation I had to use a mutex to prevent multiple concurrent requests from getting the random number and thus the same ID.

The ID has a maximum size of 6 digits and the server is not checking for ID collissions so they could happen. If not necessay I would change the ID to be auto-incremental as to not have size limitations or to check that the ID is not being used before assigning it to the car.

The horse_power property was suggested to be a string, I have changed it to be an integer as it takes less space and it would allow in the future to make queries based on that value (e.g. cars with more than 100 hp).
I have also made the same change for the id property.

In the DELETE endpoint I have used the 200 HTTP response code instead of the suggested 203 as it seams like the most appropriate according to the [definition](https://www.w3.org/Protocols/rfc2616/rfc2616-sec9.html#:~:text=A%20successful%20response%20SHOULD%20be%20200%20(OK)%20if%20the%20response%20includes%20an%20entity%20describing%20the%20status%2C%20202%20(Accepted)%20if%20the%20action%20has%20not%20yet%20been%20enacted%2C%20or%20204%20(No%20Content)%20if%20the%20action%20has%20been%20enacted%20but%20the%20response%20does%20not%20include%20an%20entity.).

Because this project can be deployed with Docker it should work in Ubuntu Server, nevetheless I have tested it manually on a Ubuntu Server in a virtual machine and it works.

