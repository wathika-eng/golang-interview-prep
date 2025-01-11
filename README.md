# golang interview prep

## Goal of this repo.

This repo contains Golang code that does not follow best practises, contains bugs and security issues. It is intended to
be used as an interview exercise or a practise exercise for jr/mid-level Go engineers.

This repo contains, technically, a functional golang application that receives a request to create a user and stores it
into a postgres Database.

As an exercise, you could try identifying and correcting some of the issues in this repo. This would work particularly
well as a pair programming exercise.

## Getting Started

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

Ensure you have Postgres installed:

````bash
Ensure you have Docker installed:

```bash
docker --version
docker compose --version
````

Ensure you have Docker Compose installed:

````bash

You can get the database started by running `docker-compose up --build`

Once running the Go app, you can make a CURL request as follows:

```bash
 curl -X POST -H "Content-Type: application/json" -d '{"username":"john", "password":"secret"}' http://localhost:8080/user
````

## Solutions

you can find the solutions [here](./solutions.md)
