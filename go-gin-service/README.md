# Go Gin Service with JWT Authentication

This is a sample web service built with the Go programming language and the Gin web framework. It demonstrates a common backend pattern: loading configuration from a file, starting a server, and protecting specific endpoints using JSON Web Token (JWT) authentication middleware.

## 🚀 Installation & Setup

Follow these steps to get the service running on your local machine.

### Prerequisites

- **Go**: You must have Go version 1.18 or newer installed. You can download it from the [official Go website](https://go.dev/dl/).

### Steps

#### 1. Get the Code

Clone this repository or download the `main.go` file into a new directory.

```sh
mkdir go-gin-service
cd go-gin-service
# (Place main.go in this directory)
```

#### 2. Initialize Go Modules

Open your terminal in the project directory and run the following command to initialize Go modules, which will manage your dependencies.

```sh
go mod init example.com/go-gin-service
```

#### 3. Install Dependencies

This command will download and install the necessary libraries (Gin, JWT, YAML).

```sh
go get -u github.com/gin-gonic/gin github.com/golang-jwt/jwt/v5 gopkg.in/yaml.v3
```

#### 4. Run the Service

Start the application with the following command. The first time you run it, it will automatically create a `config.yaml` file for you with default settings.

```sh
go run main.go
```

You should see output confirming that the server is running on port `:8080`.

```
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

Loading configuration...
config.yaml not found, creating a default one.
Configuration loaded successfully.
Starting service on :8080...
[GIN-debug] Listening and serving HTTP on :8080
```

## API Usage

The service exposes three endpoints. You can use a tool like curl or Postman to interact with them.

### 1. Public Endpoint

This endpoint is open and does not require any authentication.

- **URL:** `/public`
- **Method:** `GET`

**Example Request:**

```sh
curl http://localhost:8080/public
```

**Example Response:**

```json
{
  "message": "This is a public endpoint. Anyone can access it!"
}
```

### 2. Login to Get a Token

To access protected resources, you must first "log in" by providing a username. The service will return a JWT.

- **URL:** `/login`
- **Method:** `POST`
- **Body:** JSON object with a username.

**Example Request:**

```sh
curl -X POST http://localhost:8080/login \
     -H "Content-Type: application/json" \
     -d '{"username": "alice"}'
```

**Example Response:**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzE5OTI0NzEsImlhdCI6MTY3MTkwNjA3MSwiaXNzIjoibXktYXBwIiwidXNlcm5hbWUiOiJhbGljZSJ9.xxxxxxxxxxxx"
}
```

Copy the token value from the response for the next step.

### 3. Protected Endpoint

This endpoint requires a valid JWT in the Authorization header.

- **URL:** `/api/data`
- **Method:** `GET`
- **Header:** `Authorization: Bearer <YOUR_TOKEN>`

**Example Request:**

Replace `<YOUR_TOKEN_HERE>` with the token you received from the `/login` endpoint.

```sh
curl http://localhost:8080/api/data \
     -H "Authorization: Bearer <YOUR_TOKEN_HERE>"
```

**Example Response (Success):**

```json
{
  "message": "Hello, alice! You have accessed a protected resource."
}
```

**Example Response (Failure - No Token):**

If you try to access it without a token, you will get an error.

```json
{
    "error": "Authorization header is required"
}
```

## Lead Maintainer
- **GitHub:** [gnsalok](https://github.com/gnsalok)
