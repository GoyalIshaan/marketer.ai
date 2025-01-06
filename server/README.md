# Marketer AI Backend API Documentation

## Overview

The Marketer AI Backend API provides endpoints for user authentication, campaign management, and user data retrieval. It is built using the Fiber web framework and utilizes JWT for authentication.

## Base URL

All endpoints are prefixed with `/api`.

## Authentication

### Register

- **Endpoint**: `/api/register`
- **Method**: `POST`
- **Description**: Registers a new user.
- **Request Body**:
  ```json
  {
    "username": "string",
    "email": "string",
    "password": "string"
  }
  ```
- **Response**:
  - **Success**: `200 OK`
    ```json
    {
      "message": "Register successful"
    }
    ```
  - **Error**: `400 Bad Request` or `500 Internal Server Error`

### Login

- **Endpoint**: `/api/login`
- **Method**: `POST`
- **Description**: Authenticates a user and returns a JWT token.
- **Request Body**:
  ```json
  {
    "email": "string",
    "password": "string"
  }
  ```
- **Response**:
  - **Success**: `200 OK`
    ```json
    {
      "message": "Login successful"
    }
    ```
  - **Error**: `400 Bad Request` or `401 Unauthorized`

## Protected Routes

All protected routes require a valid JWT token in the `token` cookie.

### Check Access

- **Endpoint**: `/api/protected/checkaccess`
- **Method**: `GET`
- **Description**: Checks if the user has access.
- **Response**:
  - **Success**: `200 OK` with message "access granted"
  - **Error**: `401 Unauthorized`

## User Management

### Get Current User

- **Endpoint**: `/api/protected/user/me`
- **Method**: `GET`
- **Description**: Retrieves the current authenticated user's information.
- **Response**:
  - **Success**: `200 OK`
    ```json
    {
      "user_id": "uint",
      "username": "string",
      "email": "string"
    }
    ```
  - **Error**: `401 Unauthorized`

### Get User by ID

- **Endpoint**: `/api/protected/user/:id`
- **Method**: `GET`
- **Description**: Retrieves a user's information by their ID.
- **Response**:
  - **Success**: `200 OK`
    ```json
    {
      "user": {
        "id": "uint",
        "username": "string",
        "email": "string",
        "created_at": "string"
      }
    }
    ```
  - **Error**: `400 Bad Request` or `500 Internal Server Error`

### Get User Campaigns

- **Endpoint**: `/api/protected/user/:id/campaigns`
- **Method**: `GET`
- **Description**: Retrieves all campaigns associated with a user.
- **Response**:
  - **Success**: `200 OK`
    ```json
    {
      "campaigns": [
        {
          "id": "uint",
          "title": "string",
          "description": "string",
          "budget": "float",
          "platform": ["string"],
          "status": "string",
          "start_date": "string",
          "end_date": "string"
        }
      ]
    }
    ```
  - **Error**: `400 Bad Request` or `500 Internal Server Error`

## Campaign Management

### Create Campaign

- **Endpoint**: `/api/protected/campaign`
- **Method**: `POST`
- **Description**: Creates a new campaign.
- **Request Body**:
  ```json
  {
    "title": "string",
    "description": "string",
    "budget": "float",
    "platform": ["string"],
    "status": "string",
    "start_date": "string",
    "end_date": "string"
  }
  ```
- **Response**:
  - **Success**: `201 Created`
    ```json
    {
      "message": "Campaign created successfully"
    }
    ```
  - **Error**: `400 Bad Request` or `500 Internal Server Error`

### Get Campaign by ID

- **Endpoint**: `/api/protected/campaign/:id`
- **Method**: `GET`
- **Description**: Retrieves a campaign by its ID.
- **Response**:
  - **Success**: `200 OK`
    ```json
    {
      "campaign": {
        "id": "uint",
        "title": "string",
        "description": "string",
        "budget": "float",
        "platform": ["string"],
        "status": "string",
        "start_date": "string",
        "end_date": "string"
      }
    }
    ```
  - **Error**: `400 Bad Request` or `500 Internal Server Error`

### Update Campaign

- **Endpoint**: `/api/protected/campaign/:id`
- **Method**: `PUT`
- **Description**: Updates an existing campaign.
- **Request Body**:
  ```json
  {
    "title": "string",
    "description": "string",
    "budget": "float",
    "platform": ["string"],
    "status": "string",
    "start_date": "string",
    "end_date": "string"
  }
  ```
- **Response**:
  - **Success**: `200 OK`
    ```json
    {
      "message": "Campaign updated successfully"
    }
    ```
  - **Error**: `400 Bad Request` or `500 Internal Server Error`

### Delete Campaign

- **Endpoint**: `/api/protected/campaign/:id`
- **Method**: `DELETE`
- **Description**: Deletes a campaign by its ID.
- **Response**:
  - **Success**: `200 OK`
    ```json
    {
      "message": "Campaign deleted successfully"
    }
    ```
  - **Error**: `400 Bad Request` or `500 Internal Server Error`

## Models

### User

- **Fields**:
  - `id`: `uint`
  - `username`: `string`
  - `email`: `string`
  - `password`: `string`
  - `created_at`: `time.Time`

### Campaign

- **Fields**:
  - `id`: `uint`
  - `user_id`: `uint`
  - `title`: `string`
  - `description`: `string`
  - `budget`: `float64`
  - `platform`: `pq.StringArray`
  - `status`: `CampaignStatus`
  - `start_date`: `time.Time`
  - `end_date`: `time.Time`

## Enums

### CampaignStatus

- `pending`
- `running`
- `completed`
- `failed`

### Platform

- `facebook`
- `instagram`
- `twitter`
- `linkedin`
- `youtube`

## Middleware

### Protected

- **Description**: Middleware to protect routes, ensuring a valid JWT token is present in the request cookies.
