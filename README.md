# Unit API

A Go REST API for managing units with MySQL database integration using Gin framework.

## Prerequisites

- Go 1.26.1+
- MySQL 8.0+ (either local or via Docker)
- Git

## Setup

Clone the repository and install dependencies:

```bash
git clone https://github.com/reinatakidd/hotel-unit-management-api
cd hotel-unit-management-api
go mod download
```

## Environment Configuration

Create a `.env` file in the project root with the following variables:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=
DB_NAME=bobobox-db
```

Adjust the values based on your MySQL setup.

## Running Locally

### 1. Start MySQL

Ensure MySQL is running (e.g., via Laragon, Docker Desktop, or your system MySQL).

### 2. Create the Database and Tables

Connect to MySQL and run:

```sql
CREATE DATABASE IF NOT EXISTS `bobobox-db`;

USE `bobobox-db`;

CREATE TABLE IF NOT EXISTS `units` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL,
  `type` VARCHAR(50) NOT NULL,
  `status` VARCHAR(50) DEFAULT 'Available',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Insert sample data
INSERT INTO `units` (`name`, `type`, `status`) VALUES
('Capsule Unit A1', 'capsule', 'Available'),
('Cabin Unit B1', 'cabin', 'Available');
```

### 3. Run the API

```bash
go run cmd/server/main.go
```

The API will start on `http://localhost:8080`.

### 4. Test the API

Use Postman or curl to test endpoints:

```bash
# Get all units
curl http://localhost:8080/api/units

# Get unit by ID
curl http://localhost:8080/api/units/1

# Create a new unit
curl -X POST http://localhost:8080/api/units \
  -H "Content-Type: application/json" \
  -d '{"name":"Unit 1","status":"active"}'

# Update unit status
curl -X PUT http://localhost:8080/api/units/1 \
  -H "Content-Type: application/json" \
  -d '{"status":"inactive"}'
```

---

## Running with Docker

### 1. Build and Start Services

```bash
docker-compose up --build
```

This will:

- Build the API image from the Dockerfile
- Start a MySQL container with the database
- Start the API container

### 2. Access the API

The API is available at `http://localhost:8080` with the same endpoints as above.

### 3. Stop Services

```bash
docker-compose down
```

---

## Architecture

- **cmd/server/main.go** - Application entry point
- **internal/database/db.go** - Database connection and configuration
- **internal/handlers/unit_handler.go** - HTTP request handlers
- **internal/models/unit.go** - Data models
- **internal/repository/unit_repository.go** - Database queries

## Environment Variables

| Variable      | Default      | Description    |
| ------------- | ------------ | -------------- |
| `DB_HOST`     | `localhost`  | MySQL host     |
| `DB_PORT`     | `3306`       | MySQL port     |
| `DB_USER`     | `root`       | MySQL username |
| `DB_PASSWORD` | (empty)      | MySQL password |
| `DB_NAME`     | `bobobox-db` | Database name  |

---

## Troubleshooting

### Local: "Failed to connect to database"

Verify MySQL is running and credentials in `.env` are correct.
