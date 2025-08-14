# go-expert-cloud-run

A simple HTTP server written in Go that fetches weather data for a given Brazilian ZIP code from an external API.
The application is designed for deployment on Google Cloud Run and follows a clean, single-purpose architecture.

---

## Table of Contents
- [Project Structure](#project-structure)
- [How It Works](#how-it-works)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
  - [Environment Variables (.env)](#environment-variables-env)
  - [Example .env File](#example-env-file)
- [Running the Project](#running-the-project)
- [API Endpoints](#api-endpoints)
    - [Weather by ZIP code](#weather-by-zip-code)
    - [Error Handling](#error-handling)
- [Testing](#testing)
---

## Project Structure

```text
go-expert-cloud-run/
|-- cmd/
|   |-- server/               # Main entry point for the HTTP server
|   |   |-- [main.go]
|-- internal/
|   |-- handlers/             # HTTP handler for the /weather endpoint
|   |   |-- [weather.go]
|   |-- services/             # Service layer for fetching data
|   |   |-- [viacep.go]
|   |   |-- [weatherapi.go]
|   |-- util/                 # Validator for zip codes
|   |   |-- [validator.go]
|-- .env                      # Environment variables for local development
|-- [Dockerfile]              # Containerization recipe
|-- [docker-compose.yml]      # Local development environment setup
|-- [go.mod]
|-- [go.sum]
|-- [README.md]
```

## How It Works
* The application exposes a single HTTP endpoint /weather.
* It accepts a cep (ZIP code) query parameter in the URL.
* The server validates the ZIP code format (must be an 8-digit number).
* It queries the ViaCEP API to retrieve the city associated with the given ZIP code.
* Using the city name, it calls the WeatherAPI.com service (via WEATHER_API_KEY) to get the current temperature.
* The response is a JSON object containing the temperature in Celsius, Fahrenheit, and Kelvin.

## Getting Started
Prerequisites
* Go 1.18+
* Docker
* Docker Compose

Clone the repository
```bash
git clone https://github.com/ricardolindner/go-expert-cloud-run.git
cd go-expert-cloud-run
```

Download the dependencies:
```bash
go mod tidy
```

## Configuration
All configuration is done via environment variables.
For local development, you should create a .env file in the project root. For deployment on Cloud Run, these variables are configured directly on the platform.

### Environment Variables (.env)

**Main variables**
* `WEATHER_API_KEY`: Your API key for the weather API service.

### Example .env File
```.env
WEATHER_API_KEY=YOURKEY
```

## Running the Project
### 1. Build and Run the Server Locally (with Docker Compose)
The easiest way to run the project locally is with Docker Compose. This ensures your environment variables are correctly injected.
```bash
docker-compose up --build -d
```

The server will be available on http://localhost:8080

### 2. Run the Server Directly with Docker
Alternatively, you can build and run the Docker image manually.

First, build the image:
```bash
docker build -t go-weather-challenge .
```
Then, run the container, injecting the API key:
```bash
docker run -p 8080:8080 --name go-weather-challenge -d -e WEATHER_API_KEY="YOUR_KEY_HERE" go-weather-challenge
```

## API Endpoints
### Weather by ZIP code
Fetches weather data for a specified ZIP code.
* **URL**: `/weather`
* **Method**: `GET`
* **Query Parameters**:
    * `cep` (required): A valid 8-digit Brazilian ZIP code.

Example Request:
```bash
curl "http://localhost:8080/weather?cep=66010020"
```

Example Success Response (200 OK):
```json
{"temp_C": 12.1, "temp_F": 53.78, "temp_K": 285.1}
```

### Error Handling

* Invalid ZIP code (422 Unprocessable Entity):
If the cep parameter is not a valid 8-digit number.
```json
{"error": "invalid zipcode"}
```

* ZIP code not found (404 Not Found):
If the provided ZIP code corresponds to an unknown location.
```json
{"error": "can not find zipcode"}
```

* Internal Server Error (500):
If the external API returns an unexpected error or the server fails to process the request.
```json
{"error": "failed to fetch weather data"}
```

## Testing
* Run unit tests locally:
```bash
go test ./...
```
This command runs all unit tests for the project.
* Test the deployed API on Google Cloud Run:
Replace {{cep}} with a valid Brazilian ZIP code and run:
```bash
curl "https://go-expert-weather-83861101913.us-central1.run.app/weather?cep={{cep}}"
```
This will send a request to the production endpoint deployed on Google Cloud Run.
