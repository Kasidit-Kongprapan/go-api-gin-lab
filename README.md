# Student API with Gin and SQLite
##  Requirements

- Go 1.20+
- Git
- Thunder Client or Postman
---

##  How to Run

```bash
git clone <your-repository-url>
cd go-api-gin
go mod tidy
go run main.go
```

Server will start at:

```
http://localhost:8080
```

---

##  API Endpoints

###  Get All Students
```
GET /students
```
###  Get Student by ID
```
GET /students/:id
```
###  Create Student
```
POST /students
```
#### Request Body
```json
{
  "id": "001",
  "name": "John",
  "major": "CS",
  "gpa": 3.5
}
```
### 🔹 Update Student
```
PUT /students/:id
```

#### Request Body
```json
{
  "name": "Updated Name",
  "major": "Updated Major",
  "gpa": 3.5
}
```
###  Delete Student
```
DELETE /students/:id
```
## Project Structure
```
go-api-gin/
├── main.go
├── config/
├── models/
├── repositories/
├── services/
├── handlers/
└── students.db