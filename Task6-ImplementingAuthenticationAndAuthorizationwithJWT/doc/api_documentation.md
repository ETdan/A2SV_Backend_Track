# 📝 Task Manager API Documentation

Welcome to the **Task Manager API** documentation! This guide will help you understand how to interact with our API endpoints to manage your tasks. 🎯

## 🛠️ Endpoints

### 📜 Get All Tasks

**`GET /tasks`**

Retrieve all tasks based on user role:

- **Admin**: Get all tasks.
- **User**: Get tasks created by the specific user.

**Response:**

- **200 OK**: A list of tasks.
- **500 Internal Server Error**: Error message.

### 📜 Get Task

**`GET /tasks/:id`**

Get a specific task by its ID. Only accessible if the user is the creator or an admin.

**Response:**

- **200 OK**: Task details.
- **400 Bad Request**: Error message.

### 🔄 Update Task

**`PUT /tasks/:id`**

Update a specific task by its ID.

**Request Body:**

```json
{
  "name": "Task Name",
  "detail": "Task Detail",
  "start": "Start Date",
  "duration": "Duration"
}
```

**Response:**

- **201 Created**: Updated task details.
- **400 Bad Request**: Error message.

### 🗑️ Delete Task

**`DELETE /tasks/:id`**

Delete a specific task by its ID.

**Response:**

- **202 Accepted**: Deleted task details.
- **400 Bad Request**: Error message.

### ✉️ Post Task

**`POST /tasks`**

Create a new task.

**Request Body:**

```json
{
  "name": "Task Name",
  "detail": "Task Detail",
  "start": "Start Date",
  "duration": "Duration"
}
```

**Response:**

- **200 OK**: Created task details.
- **400 Bad Request**: Error message.

### 📝 Register User

**`POST /users/register`**

Register a new user.

**Request Body:**

```json
{
  "username": "User Name",
  "password": "Password"
}
```

**Response:**

- **200 OK**: Account created.
- **400 Bad Request**: Error message.

### 🔑 Login User

**`POST /users/login`**

Login and receive a JWT token.

**Request Body:**

```json
{
  "username": "User Name",
  "password": "Password"
}
```

**Response:**

- **200 OK**: JWT token.
- **400 Bad Request**: Error message.

---

## 🔐 Authentication

All requests require JWT authentication. Use the JWT token provided upon login to access protected endpoints. 🌐

### Example JWT Token:

```json
{
  "user_id": "User ID",
  "username": "User Name",
  "role": "user",
  "iat": "Issued At",
  "exp": "Expiration"
}
```

---

## 📚 Postman Collection

For a detailed view of all API endpoints and to test them easily, check out our Postman collection here: [Task Manager API Postman Collection](https://documenter.getpostman.com/view/24791476/2sA3rzHrWK) 📑

---

Feel free to experiment with these endpoints and have fun managing your tasks! 🚀 If you encounter any issues, check the error messages for guidance. 😊

---

Enjoy coding and happy task managing! 🎉👨‍💻👩‍💻
