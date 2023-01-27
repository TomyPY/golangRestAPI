# Tasks API

This is a simple RESTful API that allows you to interact with a list of tasks. It was built using the Go programming language and the Gorilla mux library for routing.

### Endpoints
The following endpoints are available:

- GET /
This endpoint returns a welcome message for the API.

- GET /tasks
This endpoint returns a list of all tasks currently stored in the API.

- GET /tasks/{id}
This endpoint returns a specific task with the given ID. If the task does not exist, it returns a bad request error.

- POST /tasks
This endpoint allows you to create a new task. It expects a JSON object in the request body with the task's name and content.

- PUT /tasks/{id}
This endpoint allows you to update a specific task with the given ID. It expects a JSON object in the request body with the task's new name and content.

- DELETE /tasks/{id}
This endpoint allows you to delete a specific task with the given ID

### Running the API
To run the API, you will need to have Go installed on your machine. Once you have Go set up, you can run the following command from the project's root directory:


`go run main.go`

This will start the API on port 8080.


## Note
This project is a simple example of a RESTful API and should not be used in production.
