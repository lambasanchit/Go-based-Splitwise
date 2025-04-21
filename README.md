# Go-based-Splitwise
Create a service similar to Splitwise, where users can track each expense and the amount that each person in the group owes. Users are not required to pay each bill individually; the service should allow users to sum up all of the expenses that they owe and automatically calculate the total amount owed to others in the group.


# **Splitwise-like API**

This project is a simple API that mimics the functionality of **Splitwise**, allowing users to track expenses and calculate the amounts owed between users in a group. The API provides endpoints for user registration, adding expenses, and creating groups. The expenses can be divided among multiple users in the group, and the total amount owed is calculated automatically.

---

## **Table of Contents**

1. [Problem Statement](#problem-statement)
2. [Technologies Used](#technologies-used)
3. [Setup Instructions](#setup-instructions)
   - [Install Go](#install-go)
   - [Install VS Code & REST Client](#install-vs-code--rest-client)
   - [Clone the Repository](#clone-the-repository)
   - [Run the Application](#run-the-application)
4. [API Endpoints](#api-endpoints)
   - [User Registration](#user-registration)
   - [Expense Addition](#expense-addition)
   - [Group Creation](#group-creation)
5. [Testing the API](#testing-the-api)
6. [Sample Requests & Responses](#sample-requests--responses)
7. [Contributions](#contributions)

---

## **Problem Statement**

Create a service similar to **Splitwise**, where users can track each expense and the amount that each person in the group owes. Users are not required to pay each bill individually; the service should allow users to sum up all of the expenses they owe and automatically calculate the total amount owed to others in the group.

### **Functional Requirements:**
1. **Register a new user**
   - Endpoint: `POST /api/users/register`
   - Description: Registers a new user with necessary details (name, email, password).
   
2. **Add an expense**
   - Endpoint: `POST /api/expenses/add`
   - Description: Allows users to add expenses and split the amount among multiple people in the group.

3. **Create a group**
   - Endpoint: `POST /api/groups/create`
   - Description: Allows users to create a group and add members to it.

---

## **Technologies Used**

- **Go (Golang):** The primary programming language used for backend development.
- **Gorilla Mux:** For routing HTTP requests.
- **In-memory storage:** For simplicity, the data is stored in memory (in future, a database can replace it).

---

## **Setup Instructions**

### **Install Go**

Ensure that Go is installed on your machine. If not, follow these steps:

1. Visit [Go's official website](https://golang.org/dl/).
2. Download and install the latest version of Go for your operating system.
3. Verify the installation by running the following command in your terminal:
   
   ```bash
   go version
   ```

### **Install VS Code & REST Client**

To test the API using **REST Client** in VS Code:

1. Install [Visual Studio Code](https://code.visualstudio.com/).
2. Install the **REST Client** extension by following this link: [REST Client Extension](https://marketplace.visualstudio.com/items?itemName=humao.rest-client).

### **Clone the Repository**

Clone this repository to your local machine:

```bash
git clone https://github.com/your-username/splitwise-api.git
cd splitwise-api
```

### **Run the Application**

To run the application, execute the following command:

```bash
go run main.go
```

Your application will be running on [http://localhost:8000](http://localhost:8000).

---

## **API Endpoints**

### **User Registration**

- **POST** `/api/users/register`
- **Description:** Registers a new user with name, email, and password.
- **Request Body:**
   ```json
   {
     "name": "John Doe",
     "email": "john@example.com",
     "password": "securepassword"
   }
   ```
- **Response:**
   ```json
   {
     "message": "User registered successfully",
     "userId": "391e4413-a833-4d92-bbf1-b3cbfdb9cf87"
   }
   ```

### **Expense Addition**

- **POST** `/api/expenses/add`
- **Description:** Adds a new expense and splits it between users.
- **Request Body:**
   ```json
   {
     "amount": 100,
     "payer_id": "user-id-1",
     "split_between": [
       "user-id-1",
       "user-id-2"
     ],
     "description": "Lunch"
   }
   ```
- **Response:**
   ```json
   {
     "expense": {
       "id": "expense1",
       "amount": 100,
       "payer_id": "user-id-1",
       "split_between": [
         "user-id-1",
         "user-id-2"
       ],
       "description": "Lunch"
     },
     "message": "Expense added successfully"
   }
   ```

### **Group Creation**

- **POST** `/api/groups/create`
- **Description:** Creates a group and adds members to it.
- **Request Body:**
   ```json
   {
     "group_name": "Lunch Group",
     "members": ["user-id-1", "user-id-2"]
   }
   ```
- **Response:**
   ```json
   {
     "groupId": "group1",
     "message": "Group created successfully"
   }
   ```

---

## **Testing the API**

To test the API using the **REST Client** in VS Code:

1. Install the **REST Client** extension in VS Code.
2. Create a `test.http` file in the root of the project directory.
3. Use the following structure to test the endpoints.

Example of **User Registration**:

```http
### User Registration
POST http://localhost:8000/api/users/register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "securepassword"
}
```

You can add additional tests for other endpoints like **Expense Addition** and **Group Creation** in the same file.

Click **Send Request** to test the API.

---

## **Sample Requests & Responses**

1. **User Registration Response:**
   ```json
   {
     "message": "User registered successfully",
     "userId": "391e4413-a833-4d92-bbf1-b3cbfdb9cf87"
   }
   ```

2. **Expense Addition Response:**
   ```json
   {
     "expense": {
       "id": "expense1",
       "amount": 100,
       "payer_id": "user-id-1",
       "split_between": [
         "user-id-1",
         "user-id-2"
       ],
       "description": "Lunch"
     },
     "message": "Expense added successfully"
   }
   ```

3. **Group Creation Response:**
   ```json
   {
     "groupId": "group1",
     "message": "Group created successfully"
   }
   ```

---

## **Contributions**

Feel free to fork the repository, make changes, and submit pull requests. All contributions are welcome!

---

### Enjoy using the Splitwise-like API!
