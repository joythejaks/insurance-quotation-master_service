# Insurance Master Service

Microservice for managing master data (such as Currencies, Occupations, Relationships, and Countries) within the Insurance Quotation ecosystem. This service is built using Go following Clean Architecture principles and is designed to be production-ready.

## 🚀 Key Features

- **Master Data Management**: Complete CRUD (Create, Read, Update, Delete) operations for various master data entities like Currencies, Occupations, Relationships, and Countries.
- **Pagination & Search**: Efficient retrieval of master data lists with pagination and search capabilities (by name or code).
- **Authentication**: Secure access to API endpoints using JWT (JSON Web Tokens).
- **Authorization**:
  - **RBAC (Role-Based Access Control)**: Restricting access based on roles (e.g., ADMIN).
  - **ACL (Access Control List)**: Granular access control at the permission level (e.g., `manage_master_data`, `manage_countries`).
- **Standardized Response**: Consistent JSON format for all API responses.
- **Global Error Handling**: Centralized middleware to handle application errors gracefully.
- **Structured Logging**: High-performance structured logging (JSON) using Uber Zap (assumed, consistent with auth service).
- **API Documentation**: Automatic Swagger UI integration for easy API exploration.

## 🛠 Tech Stack

- **Language**: Go (Golang)
- **Framework**: Gin Gonic
- **ORM**: GORM
- **Database**: PostgreSQL
- **Security**: JWT (v5) for authentication (handled by `AuthMiddleware`)
- **Logging**: Uber Zap
- **Docs**: Swaggo (Swagger)

## 📋 Prerequisites

- Go 1.21 or higher
- PostgreSQL
- Migrate (Golang Migrate) tool for database schema management

## ⚙️ Environment Configuration

Create a `.env` file in the root directory of the service:

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=secret
DB_NAME=insurance_master_db
DB_PORT=5432
APP_PORT=8081 # Or any other available port
JWT_SECRET=your_super_secret_key_for_master_service
```

**Note**: `JWT_SECRET` should be the same secret key used by the `insurance-quotation-auth_service` for token validation.

## 🏃 Running the Application

1. **Install Dependencies**:

```bash
    go mod tidy
```

2. **Run Database Migrations**:

```bash
    make migrate-up
```

3. **Generate Swagger Docs**:

```bash
    swag init -g cmd/api/main.go
```

4. **Start the Application**:

```bash
    go run cmd/api/main.go
```

## 📖 API Documentation

Once the application is running, open your browser and navigate to:
`http://localhost:8081/swagger/index.html` (adjust the port if different)

## 📂 Project Structure

```text
├── cmd/api             # Application entry point
├── docs/               # Generated Swagger documentation
├── internal/
│   ├── config/         # Application & environment configuration
│   ├── dto/            # Data Transfer Objects (Request/Response)
│   ├── handler/        # Controller layer / HTTP entry point
│   ├── middleware/     # Middleware (Auth, Log, Error, Role, Permission)
│   ├── model/          # Entity / Database schema
│   ├── repository/     # Database access layer
│   ├── service/        # Business logic layer (if applicable; currently logic resides directly in handler/repo)
│   ├── utils/          # Helpers (Response, Logger)
│   └── router/         # API route definitions
└── migrations/         # Database migration SQL files
```
