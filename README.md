# Task Management API with MongoDB Integration

This project extends the existing Task Management API by integrating MongoDB as the persistent data storage solution. It replaces the in-memory database with MongoDB to ensure data persistence across API restarts. The API is built using Go (Golang) and the MongoDB Go Driver.

## Features
- **CRUD Operations**: Create, Read, Update, and Delete tasks using MongoDB.
- **Persistent Data Storage**: Tasks are stored in MongoDB, ensuring data persistence.
- **Backward Compatibility**: The API remains compatible with the previous version, maintaining the same endpoint structure and behavior.
- **Error Handling**: Proper error handling for MongoDB operations, including network and database errors.
- **Documentation**: Detailed API documentation and setup instructions.

## Requirements
- Go (Golang) installed on your machine.
- MongoDB instance (local or cloud-based).
- MongoDB Go Driver (`go.mongodb.org/mongo-driver`).

## Installation

### 1. Clone the Repository
```bash
git clone https://github.com/your-username/task-manager.git
cd task-manager
2. Install Dependencies
Run the following command to install the required Go packages:

bash
Copy
go mod tidy
3. Set Up MongoDB
Install MongoDB locally or use a cloud service like MongoDB Atlas.

Create a database and collection for tasks (e.g., task_manager database and tasks collection).

Update the MongoDB connection string in the application (see Configuration).

4. Run the Application
Start the API server:

bash
Copy
go run main.go
The API will be available at http://localhost:8080.

Configuration
MongoDB Connection
Update the MongoDB connection string in the application. You can set it as an environment variable or hardcode it in the main.go file.

Example:

go
Copy
clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
client, err := mongo.Connect(context.TODO(), clientOptions)
Environment Variables
You can use environment variables to configure the application:

MONGODB_URI: MongoDB connection string (e.g., mongodb://localhost:27017).

PORT: Port for the API server (default: 8080).

Example:

bash
Copy
export MONGODB_URI="mongodb://localhost:27017"
export PORT=8080
API Endpoints
Tasks
GET /tasks: Retrieve a list of all tasks.

GET /tasks/:id: Retrieve details of a specific task by ID.

POST /tasks: Create a new task.

PUT /tasks/:id: Update an existing task.

DELETE /tasks/:id: Delete a task.

Example Requests
Create a Task
bash
Copy
curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{
  "title": "Complete Project Report",
  "description": "Write and submit the final project report by the deadline.",
  "duedate": "2023-12-31T23:59:59Z",
  "status": "Pending"
}'
Get All Tasks
bash
Copy
curl -X GET http://localhost:8080/tasks
Get a Task by ID
bash
Copy
curl -X GET http://localhost:8080/tasks/1
Update a Task
bash
Copy
curl -X PUT http://localhost:8080/tasks/1 \
-H "Content-Type: application/json" \
-d '{
  "status": "In Progress"
}'
Delete a Task
bash
Copy
curl -X DELETE http://localhost:8080/tasks/1
Folder Structure
Copy
task_manager/
├── main.go
├── controllers/
│   └── task_controller.go
├── models/
│   └── task.go
├── data/
│   └── task_service.go
├── router/
│   └── router.go
├── docs/
│   └── api_documentation.md
└── go.mod
Key Files
main.go: Entry point of the application.

controllers/task_controller.go: Handles incoming HTTP requests and invokes the appropriate service methods.

models/task.go: Defines the Task data structure.

data/task_service.go: Contains business logic and MongoDB operations.

router/router.go: Sets up the routes and initializes the Gin router.

docs/api_documentation.md: Contains API documentation.

go.mod: Defines the module and its dependencies.

Testing
Use tools like Postman or curl to test the API endpoints. Verify that tasks are stored and retrieved correctly from MongoDB.

Example Test Cases
Create a new task and verify it is stored in MongoDB.

Retrieve all tasks and verify the response.

Update a task and verify the changes in MongoDB.

Delete a task and verify it is removed from MongoDB
