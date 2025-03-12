# golang interview prep

## Goal of this repo

This repo contains Golang code that does not follow best practises, contains bugs and security issues. It is intended to
be used as an interview exercise or a practise exercise for jr/mid-level Go engineers.

This repo contains, technically, a functional golang application that receives a request to create a user and stores it
into a postgres Database.

As an exercise, you could try identifying and correcting some of the issues in this repo. This would work particularly
well as a pair programming exercise.

## Getting Started

Clone the repo:

```bash
git clone https://github.com/wathika-eng/golang-interview-prep && cd golang-interview-prep
```

Copy the .env.example file to .env and fill in the values.

```bash
cp .env.example .env
```

Ensure you have Go installed:

```bash
go version
```

Then

```bash
go mod tidy
go run main.go
```

Get a postgresql database and redis running, either via docker, locally or on a cloud provider(<https://neon.tech>)

```bash
pg_isready # check if postgres is running, should return something like /var/run/postgresql:5432 - accepting connections
redis-cli ping # check if redis is running, should return PONG
```

Ensure you have Docker Compose installed:

```bash
docker --version
docker compose version
```

Run postgresql database and redis using docker-compose

```bash
COMPOSE_BAKE=true docker compose up --build
```

You can get the database started by running `docker-compose up --build`

Once running the Go app, you can make a CURL request as follows:

```bash
curl -X POST http://localhost:8080/api/v1/user \
     -H "Content-Type: application/json" \
     -d '{
           "email": "johndoe@gmail.com",
           "phone_number": "+2547123456",
           "username": "john_doe",
           "work_id": 12345678,
           "password": "12345556,"
         }'

```

GET request (get a user using work_id)

```bash
curl -X GET http://localhost:8080/api/v1/users
```

PATCH request (work_id cannot be updated)

```bash
curl -X PATCH http://localhost:8080/api/v1/user/12345678 \
     -H "Content-Type: application/json" \
     -d '{
           "email": "newjohndoe@gmail.com",
           "phone_number": "+25474658",
           "username": "new_john_doe"
        }'

```

DELETE request (delete a user using work_id)

```bash
curl -X DELETE http://localhost:8080/api/v1/user/12345678
```

<!-- Test non-existent user

```bash -->

## Solutions

you can find the solutions [here](./solutions.md)

```

```
