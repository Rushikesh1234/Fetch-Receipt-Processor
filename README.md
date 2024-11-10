# Fetch-Receipt-Processor
The **Fetch Receipt Processor** is a simple API designed to process and calculate points for receipts based on various criteria. The application takes a receipt's details—such as retailer, purchase date, items, and total—and calculates a points score for the transaction. Points are awarded based on several rules, such as the length of the retailer's name, the total amount, the number of items, and the time of purchase.

This project demonstrates how to build a RESTful web service using Go, providing endpoints to submit receipts for processing and retrieve the points awarded to a specific receipt. Whether you're tracking rewards for loyalty programs or simply analyzing purchase data, this service is a useful tool for handling receipt data.

## Project Structure

The project consists of two main endpoints:

1. **POST /receipts/process**: Submit a receipt for processing and get a unique ID.
2. **GET /receipts/{id}/points**: Get the points awarded for the processed receipt by providing its ID.

## API Endpoints

### 1. **POST /receipts/process**

This endpoint accepts a JSON payload with receipt details and returns a unique ID.

#### POST URL : 
```bash
POST {{baseUrl}}/receipts/process
```

#### Request Body Example:

```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}
```

#### Response Example:

```json
{
    "id": "b800869d-2460-4bf3-8fdd-dae0fc5594c6"
}
```

### 2. GET /receipts/{id}/points
This endpoint returns the points awarded for a specific receipt, identified by its unique ID.

#### GET URL / Request Example:
```bash
GET {{baseUrl}}/receipts/b800869d-2460-4bf3-8fdd-dae0fc5594c6/points
```

#### Response Example:
```json
{
    "points": 28
}
```

## Rules for Points Calculation

- **1 point** for every alphanumeric character in the retailer name.
- **50 points** if the total is a round dollar amount with no cents.
- **25 points** if the total is a multiple of 0.25.
- **5 points** for every two items on the receipt.
- If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2, round up to the nearest integer, and add that number of points.
- **6 points** if the day in the purchase date is odd.
- **10 points** if the purchase time is between 2:00pm and 4:00pm.

## Setup Instructions
### Prerequisites
```bash
Go 1.23 (https://go.dev/doc/install)
Docker
```

### 1. Clone the Repository
```bash
git clone https://github.com/Rushikesh1234/Fetch-Receipt-Processor.git
cd Fetch-Receipt-Processor
```

### 2. Run the Project Locally
You can run the project locally using the Go command.

#### Run the Go application:
```bash
go run main.go
```
This will start the API server on http://localhost:8080.

### 3. Docker Setup (Optional)
If you prefer to use Docker to run the application, you can build and run it using Docker.

#### Build the Docker image:
```bash
docker build -t fetch-receipt-processor .
```

#### Run the Docker container:
```bash
docker run -p 8080:8080 fetch-receipt-processor
```

### 4. Test the API
Test the API using Postman:
- POST /receipts/process: Submit a receipt JSON and get the receipt ID.
- GET /receipts/{id}/points: Get the points by passing the receipt ID.

#### Step 1: Open Postman in your Desktop
If you don't have Postman installed, download it from Postman website.
Open Postman.

#### Step 2: Set Up Environment
Set up an environment in Postman with a base URL variable to easily test the API.
1. Click on the **Environments** tab (gear icon in the top right of Postman).
2. Create a new environment by clicking the **Add** button.
3. Add a new variable with the following details:
    - **Variable Name**: `baseUrl`
    - **Initial Value**: `http://localhost:8080`
    - **Current Value**: `http://localhost:8080`
4. Click **Save**.

Once the environment is set up, please use the `baseUrl` variable in API requests like this:
```bash
{{baseUrl}}/receipts/process
```
<img width="1310" alt="image" src="https://github.com/user-attachments/assets/e62b78e5-cc32-43f6-856e-4b98b842d3a7">

#### Step 3: Test the POST API - POST /receipts/process
1. Set the request type to POST in Postman.
2. Set Headers to:
    - Key: `Content-Type`
    - Value: `application/json`
3. URL: `{{baseUrl}}/receipts/process`
4. Request Body (JSON):
```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}
```
5. Response Body Example -
```json
{
    "id": "b800869d-2460-4bf3-8fdd-dae0fc5594c6"
}
```
<p align="center">
  <img src="https://github.com/user-attachments/assets/dd513b8b-3ab0-46f8-8d5d-d0f4f1f05f58" style="border: 10 solid #000000;">
</p>

#### Step 4: Test the GET API - GET /receipts/{id}/points
1. Set the request type to GET in Postman.
2. URL: `{{baseUrl}}/receipts/b800869d-2460-4bf3-8fdd-dae0fc5594c6/points`
3. Response Body Example -
```json
{
    "points": 28
}
```
<p align="center">
  <img src="https://github.com/user-attachments/assets/ffd7030e-5111-46f6-9285-edd8e4a8acb9" style="border: 10 solid #000000;">
</p>
