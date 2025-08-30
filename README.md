# Vehicle Management Payment Service

This is a backend microservice designed to handle payment processing for the Vehicle Management ecosystem.
It integrates with the Stripe payment gateway to securely create payments and uses webhooks to receive real-time
updates on payment statuses.

## Features

- **Initiate Payment Processing:** Creates payment sessions for vehicle sales via Stripe.
- **Stripe Webhook Integration:** Handles incoming webhook events from Stripe to update payment statuses
automatically (e.g., charge.succeeded, charge.failed).

## Technologies

- **Containerization**: Docker and Docker Compose
- **Backend**: Go
- **Database**: PostgreSQL
- **API Documentation**: Swagger

## Getting Started

Follow these steps to get the application up and running on your local machine.

#### 1. Clone the repository:

```shell
git clone https://github.com/kalilventura/vehicle-management-payment.git
cd vehicle-management-payment
```

#### 2. Create the Environment File:

This project requires certain environment variables to be set.
You can find a template for these variables in the .env.example file.
To create your own .env file, run the following command:

```shell
cp .dev/.env.example .dev/.env
```

Then, edit the .env file to include the appropriate values for your setup.

#### 3. Building and running your application

When you're ready, start your application by running:

```shell
make dev-up
```

This command will build and start all the services defined in your docker-compose.yml file.

## API Documentation

The API provides several endpoints for managing vehicles. The base URL is `http://localhost:8080`.

For full, interactive documentation, you can likely access a Swagger UI instance at
`http://localhost:8080/swagger/index.html` or similar a path once the application is running.

## Contributing

Contributions are welcome! Please submit a pull request with a clear description of your changes.
