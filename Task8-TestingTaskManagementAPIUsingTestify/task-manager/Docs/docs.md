# 🎉 Task Manager API Documentation 🎉

Welcome to the Task Manager API! 🚀 Before diving into the details, you can explore our API endpoints using the [Postman Collection](https://documenter.getpostman.com/view/24791476/2sA3rzHrWK). 📦

## 🛠️ Setup Instructions

1. **Clone the Repository**:
   ```bash
   git clone <repository-url>
   ```
2. **Install Dependencies**:

   ```bash
   go mod tidy
   ```

3. **Create a `.env` File**:
   You need to set up your environment variables in a `.env` file in the root of your project. Here’s an example of the required variables:

   ```env
   DATABASE_URL=<your-database-url>
   DATABASE=<database-name>
   USER_COLLECTION=<user-collection-name>
   TASK_COLLECTION=<task-collection-name>
   ```

4. **Run the Server**:
   Start the server using:
   ```bash
   go run main.go
   ```
   This will load the environment variables from your `.env` file and start the server. 🚀

## 📚 Endpoints

### User Registration & Login

- **Register a User**

  - **Endpoint:** `POST /register`
  - **Description:** Create a new user with a username and password.
  - **Request Body:**
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - **Responses:**
    - `201 Created` - User registered successfully. 🎉
    - `400 Bad Request` - Username or password missing. ❌

- **Register an Admin**

  - **Endpoint:** `POST /registeradmin`
  - **Description:** Create a new admin user with a username and password.
  - **Request Body:**
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - **Responses:**
    - `201 Created` - Admin registered successfully. 🎉
    - `400 Bad Request` - Username or password missing. ❌

- **Login a User**
  - **Endpoint:** `POST /login`
  - **Description:** Authenticate a user and get a JWT token.
  - **Request Body:**
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - **Responses:**
    - `200 OK` - Token generated successfully. 🎟️
    - `400 Bad Request` - Invalid username or password. ❌

### Task Management

- **Get All Tasks**

  - **Endpoint:** `GET /tasks`
  - **Description:** Retrieve all tasks for the authenticated user.
  - **Authentication:** Required (Bearer token) 🔐
  - **Responses:**
    - `200 OK` - List of tasks. 📋
    - `500 Internal Server Error` - Error retrieving tasks. 😓

- **Get Task by ID**

  - **Endpoint:** `GET /tasks/:id`
  - **Description:** Retrieve a specific task by its ID.
  - **Authentication:** Required (Bearer token) 🔐
  - **Responses:**
    - `200 OK` - Task details. 📝
    - `500 Internal Server Error` - Error retrieving task. 😓

- **Add a Task**

  - **Endpoint:** `POST /tasks`
  - **Description:** Create a new task.
  - **Authentication:** Required (Bearer token) 🔐
  - **Request Body:**
    ```json
    {
      "name": "string",
      "detail": "string",
      "start": "string",
      "duration": "string"
    }
    ```
  - **Responses:**
    - `201 Created` - Task added successfully. 🆕
    - `400 Bad Request` - Invalid task data. ❌

- **Update a Task**

  - **Endpoint:** `PUT /tasks/:id`
  - **Description:** Update an existing task by its ID.
  - **Authentication:** Required (Bearer token) 🔐
  - **Request Body:**
    ```json
    {
      "name": "string",
      "detail": "string",
      "start": "string",
      "duration": "string"
    }
    ```
  - **Responses:**
    - `200 OK` - Task updated successfully. ✏️
    - `400 Bad Request` - Invalid task data. ❌

- **Delete a Task**
  - **Endpoint:** `DELETE /tasks/:id`
  - **Description:** Delete a task by its ID.
  - **Authentication:** Required (Bearer token) 🔐
  - **Responses:**
    - `200 OK` - Task deleted successfully. 🗑️
    - `400 Bad Request` - Error deleting task. ❌

## 🚀 Authentication

To access most of the endpoints related to tasks, you need to authenticate using a JWT token. You will receive this token upon successful login. 🛡️

## 📦 Error Codes

- **400 Bad Request**: Your request has missing or incorrect data. Please check the request format. 🚫
- **401 Unauthorized**: Authentication failed. Check your token and permissions. 🔑
- **500 Internal Server Error**: Something went wrong on the server. 😞

Enjoy using the Task Manager API! 🎉 If you encounter any issues, feel free to reach out. 😊

---
