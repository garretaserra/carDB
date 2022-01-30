# Car DB
## Requirements
This server is written in Go and setup to be spun up using Docker with a `docker-compose` file.

## Run the app
To start the server and database run:
```
docker-compose up
```

## Setup DB
Before making requests you should setup the Database by running (from inside the `/server` directory):
```
go scripts/dbSetup.go
```
This will create the car table for storing the entities

## Run the tests
To test the api run (from inside the `/server` directory):
```
k6 run test.js
```

# Configuration
## Environment variables
By default the server will be listening on port 8080.

The port and other environment variables can be changed on the `/envs` file for local or dev (Docker) deployments.

## Security
It is recommended to change the DB credentials to private ones.
Also you should not expose the DB service externally but this will prevent the DB setup script from working.

# API
## Get Car
### Request
`GET /cars/id`

### Response
A car object

## Create a Car
### Request
`POST /cars`

Provide in the body of the request a json object with the parameters: "brand" (string), "model"(string) and "horse_power"(integer)

### Response
A car object with the provided values and the "id"(string) assigned to it

## Delete a Car
### Request
`DELETE /cars/id`
### Response
The car object that has been deleted

# Comments
Regarding ID generation I had to use a mutex to prevent multiple simultaneous requests from getting the same ID.

Although the ID has a maximum size of 6 digits the server is not checking for ID collissions so they could happen. If not necessay I would change the ID to be incremental as to not have size limitations or to check that the ID is not being used before assigning it to the car.

In the DELETE endpoint I have used the 200 HTTP response code instead of the suggested 203 as it seams like the most appropriate according to the definition.



