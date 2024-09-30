# Max Path Sum Service
This is a simple web service built using the Echo framework in Go. It accepts a binary tree in JSON format and returns the maximum path sum. A path sum is the sum of node values along any path in the binary tree.

For simplicity and readability only models are divided into separate files. The server logic is in main.go.

## Features
- Single endpoint to calculate the maximum path sum of a binary tree.
- Request and response in JSON format.
- Uses Echo framework for routing and request handling.
- Graceful shutdown of the server.
- Logging of requests and responses.
- CORS support.

## Prerequisites
- Go 1.16 or later
- Echo framework

### Installation
Clone the repository:
```bash
git clone https://github.com/oguzhantasimaz/btree_service.git
```
### Navigate to the project directory:
Install dependencies:
```bash
go mod tidy
```

### Running the Server
To start the web server, run:

```bash
go run main.go
```
This will start the server on http://localhost:3000.

### API Endpoints
POST /max-path-sum
Accepts a binary tree in JSON format and returns the maximum path sum.

### Request Example:
```json
{
  "tree": {
    "nodes": [
      {"id": "1", "left": "2", "right": "3", "value": 1},
      {"id": "2", "left": "4", "right": "5", "value": 2},
      {"id": "3", "left": "6", "right": "7", "value": 3},
      {"id": "4", "left": null, "right": null, "value": 4},
      {"id": "5", "left": null, "right": null, "value": 5},
      {"id": "6", "left": null, "right": null, "value": 6},
      {"id": "7", "left": null, "right": null, "value": 7}
    ],
    "root": "1"
  }
}
```


### Response Example:
```json
{
  "maxPathSum": 18
}
```


### Running Tests
To add and run tests, create unit tests for the application. You can use Go's built-in testing framework:

```bash
go test .
```

### Project Structure
* main.go: Contains the core server logic.
* go.mod: Lists dependencies and Go module information.
* models/: Contains the tree and node models.


