# Kafka Playground in Golang

## Project Overview

This project is designed as a Kafka playground using Golang. It demonstrates the use of microservices with an emphasis on hexagonal architecture. The setup consists of an API Gateway (`api-gateway`) acting as the Kafka producer and a User Service (`user-service`) as the Kafka consumer. This setup showcases how services can communicate through Kafka topics, illustrating both the producer and consumer roles in a microservice ecosystem.

## Using Sarama to Communicate with Kafka

[**Sarama**](https://github.com/IBM/sarama) is a Go client library for Apache Kafka. It allows Go applications to produce and consume messages from Kafka clusters. In this project, Sarama is used for the following purposes:

- **Producer**: The API Gateway uses Sarama to send messages to Kafka topics.
- **Consumer**: The User Service uses Sarama to consume messages from Kafka topics.

Sarama provides a high-level API to interact with Kafka, allowing easy configuration and management of Kafka producers and consumers. For more detailed usage, refer to the [Sarama documentation](https://pkg.go.dev/github.com/Shopify/sarama).

## Running the Project

Follow these steps to set up and run the project:

1. **Create Environment File**
    - Copy the example environment file and adjust it for your setup:
      ```bash
      cp .env.example .env
      ```
    - Edit the `.env` file as needed.

2. **Start Docker Containers**
    - Build and start the Docker containers:
      ```bash
      docker-compose up --build
      ```

3. **Fetch Dependencies**
    - Navigate to the root of the project and download the required Go dependencies:
      ```bash
      go mod tidy
      ```

4. **Run the API Gateway (Producer)**
    - Open a terminal and navigate to the `api-gateway` directory:
      ```bash
      cd api-gateway
      ```
    - Start the producer:
      ```bash
      go run src/apps/api/main.go
      ```

5. **Run the User Service (Consumer)**
    - Open another terminal and navigate to the `user-service` directory:
      ```bash
      cd user-service
      ```
    - Start the consumer:
      ```bash
      go run src/apps/api/main.go
      ```

## How to access route documentation?

Once both the API Gateway and User Service are running, you can verify that everything is working correctly. Access the Swagger documentation for the API Gateway at the following URL:

- [Swagger Documentation](http://localhost:8000/api/docs/index.html)

This will provide an interface to interact with and test the API endpoints exposed by the API Gateway.

