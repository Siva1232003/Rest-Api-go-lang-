
# Task API

This is a simple REST API for managing tasks. It provides endpoints for creating, retrieving, updating, and deleting tasks. The API interacts with a PostgreSQL database to persist task data.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Setting Up PostgreSQL](#setting-up-postgresql)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
  - [POST /tasks](#post-tasks)
  - [GET /tasks](#get-tasks)
  - [GET /tasks/{id}](#get-tasksid)
  - [PUT /tasks/{id}](#put-tasksid)
  - [DELETE /tasks/{id}](#delete-tasksid)

## Prerequisites

- Go (Golang) installed on your machine.
- PostgreSQL installed and running.
- An environment file (`.env`) with the necessary configuration.

## Setting Up PostgreSQL

1. **Install PostgreSQL:**
   - Download and install PostgreSQL from the official website: https://www.postgresql.org/download/

2. **Create a Database:**
   - After installing PostgreSQL, create a database called `task_api`:
     ```bash
     createdb task_api
     ```

3. **Create a Table:**
   - The application will automatically create the necessary table (`tasks`) when the application starts. However, you can manually create it using the following SQL command:
     ```sql
     CREATE TABLE tasks (
         id SERIAL PRIMARY KEY,
         title VARCHAR(100) NOT NULL,
         description TEXT,
         completed BOOLEAN DEFAULT FALSE
     );
     ```

4. **Set Up Environment Variables:**
   - Create a `.env` file in the root directory of your project and add the following environment variables with your PostgreSQL database credentials:
     ```
     DB_USER=your_database_user
     DB_PASSWORD=your_database_password
     DB_NAME=task_api
     DB_HOST=localhost
     DB_PORT=5432
     ```

## Running the Application

1. **Install Dependencies:**
   - Install the required Go dependencies by running the following command:
     ```bash
     go mod tidy
     ```

2. **Run the Application:**
   - Start the application by running:
     ```bash
     go run main.go
     ```
   - The server will start on `http://localhost:8080`.

## API Endpoints

### POST /tasks
**Create a new task.**

#### Request:
```json
{
  "title": "Task title",
  "description": "Task description",
  "completed": false
}
```

#### Response:
```json
{
  "id": 1,
  "title": "Task title",
  "description": "Task description",
  "completed": false
}
```

### GET /tasks
**Retrieve all tasks.**

#### Response:
```json
[
  {
    "id": 1,
    "title": "Task 1",
    "description": "Task description",
    "completed": false
  },
  {
    "id": 2,
    "title": "Task 2",
    "description": "Another task description",
    "completed": true
  }
]
```

### GET /tasks/{id}
**Retrieve a task by ID.**

#### Request:
```bash
GET /tasks/1
```

#### Response:
```json
{
  "id": 1,
  "title": "Task title",
  "description": "Task description",
  "completed": false
}
```

### PUT /tasks/{id}
**Update a task by ID.**

#### Request:
```json
{
  "title": "Updated Task title",
  "description": "Updated Task description",
  "completed": true
}
```

#### Response:
```json
{
  "id": 1,
  "title": "Updated Task title",
  "description": "Updated Task description",
  "completed": true
}
```

### DELETE /tasks/{id}
**Delete a task by ID.**

#### Request:
```bash
DELETE /tasks/1
```

#### Response:
```json
{
  "message": "Deleted successfully"
}
```

## Error Handling

The API returns the following error codes:

- `400 Bad Request`: Invalid input or missing parameters.
- `404 Not Found`: Task not found.
- `500 Internal Server Error`: Something went wrong with the server.

### Example Error Responses:

- Invalid Task ID:
  ```json
  {
    "error": "Invalid task ID"
  }
  ```

- Task not found:
  ```json
  {
    "error": "Task not found"
  }
  ```

## License

This project is open-source and available under the MIT License.
