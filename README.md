# Golang Dad Jokes API

This API makes use of the [icanhazdadjoke] (https://icanhazdadjoke.com/api) API to retrieve dad jokes. The backend is written entirely in Golang and uses MongoDB as the database for storing jokes. An instance of MongoDB is stood up on Docker container. The ability to interact with the MongoDB instance comes from a mongoagent.

When a user asks for a new random joke, the random joke is only added to the database if it isn't already in the database.

A simple HTML webpage serves at a basic user interface for getting random jokes, searching up jokes by ID, modifying jokes by ID, and deleting jokes by ID.

## Advantage of having a MongoDB instance inside of a Docker container

All dependencies, versions, etc. that are required to run the MongoDB instance are self-contained within the container. In other words, the MongoDB application is running in its own container, separate from the OS. By using a Docker container, any machine/OS/device is able to use an instance of MongoDB. This container can be shared for use across many platforms.

This project served as an introduction to Golang for me.

## Usage

```bash
go run server.go
```

```bash
docker container start container_id
docker container stop container_id
```

6/2/2021