# Go Docker Application

This project is a simple Go application that is containerized using Docker. Below are the instructions on how to build and run the application.

## Prerequisites

- Docker installed on your machine
- Go installed on your machine (optional, for local development)

## Project Structure

```
go-docker-app
├── src
│   └── main.go
├── Dockerfile
├── go.mod
└── README.md
```

## Building the Docker Image

To build the Docker image for the Go application, navigate to the project directory and run the following command:

```
docker build -t go-docker-app .
```

## Running the Docker Container

After building the image, you can run the Docker container using the following command:

```
docker run -p 8080:8080 go-docker-app
```

This will start the application and map port 8080 of the container to port 8080 on your host machine.

## Accessing the Application

Once the container is running, you can access the application by navigating to `http://localhost:8080` in your web browser.

## Additional Information

Feel free to explore the `src/main.go` file to understand the application logic and modify it as needed. The `Dockerfile` contains the instructions for building the Docker image, and the `go.mod` file manages the dependencies for the Go application.