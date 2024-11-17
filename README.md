
# Task API

This project is a simple REST API built using Go, which allows you to create, read, update, and delete tasks in a PostgreSQL database.

## Project Structure

The project is organized as follows:

- `main.go`: Initializes the application and starts the server.
- `db/`: Contains database connection logic and table creation.
- `routes/`: Contains routes and handlers for different API endpoints.
- `services/`: Contains business logic for creating, fetching, updating, and deleting tasks.
- `handlers/`: Handles incoming HTTP requests and responses.

## Features

- **Create Task**: Add a new task to the database.
- **Get Tasks**: Retrieve all tasks from the database.
- **Get Task by ID**: Retrieve a specific task using its ID.
- **Update Task**: Modify the details of an existing task.
- **Delete Task**: Remove a task from the database.

## Requirements

- Go (version 1.16 or higher)
- PostgreSQL database
- `.env` file to store your database credentials.

## Setup

### 1. Clone the repository:

```bash
git clone https://github.com/yourusername/task-api.git
cd task-api
```

### 2. Create a `.env` file in the root directory with the following content:

```bash
DB_USER=your_db_username
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
DB_HOST=localhost
DB_PORT=5432
```

### 3. Install dependencies:

```bash
go mod tidy
```

### 4. Initialize the database connection and create the necessary table:

```bash
go run main.go
```

### 5. API Endpoints

- `GET /tasks`: Fetch all tasks.
- `GET /tasks/{id}`: Fetch a task by its ID.
- `POST /tasks`: Create a new task.
- `PUT /tasks/{id}`: Update an existing task.
- `DELETE /tasks/{id}`: Delete a task by its ID.

### 6. Run the server:

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

## Error Handling

- If you try to create a task with a duplicate title, the API will return a `400 Bad Request` error.
- If a task is not found during fetching, updating, or deleting, a `404 Not Found` error will be returned.
- Other unexpected errors will return a `500 Internal Server Error`.

## License

This project is open-source and available under the MIT License.
