# Redikru Jobs API

A RESTful API service built with Go (Golang), Gin framework, and MySQL for managing job postings with advanced filtering capabilities.

## 🚀 Features

- **RESTful API** - Clean and intuitive endpoints
- **Database Integration** - MySQL with GORM ORM
- **Input Sanitization** - Bluemonday for XSS protection
- **Docker Support** - Complete containerization
- **Advanced Filtering** - Search jobs by company, title, and keywords
- **Pagination Support** - Limit and offset parameters

## 📋 Prerequisites

- Docker & Docker Compose
- (Optional) Go 1.21+ and MySQL for local development

## 🛠️ Installation & Setup

### Using Docker (Recommended)

1. **Clone the repository**
   ```bash
   git clone https://github.com/reno-25/redikru-jobs-api.git
   cd redikru-jobs-api
   ```

2. **Environment Configuration** (Optional)
   ```bash
   cp .env.example .env
   # Edit .env file if you need custom configurations
   ```

3. **Start the application**
   ```bash
   docker compose up --build
   ```

4. **Access the API**
   - API: http://localhost:8081
   - MySQL: localhost:3307 (user: root, password: secret)

### Manual Setup (Without Docker)

1. **Install dependencies**
   ```bash
   go mod download
   ```

2. **Setup MySQL database**
   - Create database: `redikru`
   - Run migrations from `migrations/init.sql`

3. **Configure environment variables**
   ```bash
   export DB_DSN="root:secret@tcp(localhost:3306)/redikru?charset=utf8mb4&parseTime=True&loc=Local"
   export APP_PORT=8080
   ```

4. **Run the application**
   ```bash
   go run ./cmd/server
   ```

## 📡 API Endpoints

### Create a Job
```http
POST /jobs
Content-Type: application/json

{
  "companyName": "TechCorp",
  "title": "Senior Backend Developer",
  "description": "Develop scalable backend systems using Go and MySQL"
}
```

### List Jobs (with filtering)
```http
GET /jobs?keyword=backend&companyName=TechCorp&limit=10&offset=0
```

**Query Parameters:**
- `keyword` - Search in title and description
- `companyName` - Filter by company name
- `limit` - Number of results per page (default: 10)
- `offset` - Pagination offset (default: 0)

## 🗄️ Database Schema

```sql
CREATE TABLE jobs (
  id CHAR(36) PRIMARY KEY,
  company_name VARCHAR(255) NOT NULL,
  title VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

## 🔧 Development

### Project Structure
```
redikru-jobs-api/
├── cmd/server/          # Main application entry point
├── internal/            # Internal packages
│   ├── handlers/        # HTTP handlers
│   ├── models/          # Data models
│   ├── repository/      # Database operations
│   └── service/         # Business logic
├── migrations/          # Database migrations
├── test/               # Test files
└── docker-compose.yml   # Docker configuration
```

### Running Tests
```bash
go test ./...
```

### Building for Production
```bash
go build -o redikru-server ./cmd/server
```

## 🐳 Docker Configuration

The application uses a multi-stage Docker build:
- **Build stage**: Go 1.21-alpine for compilation
- **Production stage**: Minimal Alpine image for runtime

### Port Configuration
- **API**: 8081 (external) → 8080 (internal)
- **MySQL**: 3307 (external) → 3306 (internal)

## 🔒 Security Features

- **Input Sanitization**: All user input is sanitized using Bluemonday to prevent XSS attacks
- **SQL Injection Protection**: GORM parameterized queries
- **Environment Variables**: Sensitive configuration stored in environment variables

## 📊 Performance Considerations

- For large datasets, consider adding FULLTEXT indexes:
  ```sql
  ALTER TABLE jobs ADD FULLTEXT(company_name, title, description);
  ```
- For production workloads, consider using Elasticsearch or similar search services

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/new-feature`
3. Commit changes: `git commit -am 'Add new feature'`
4. Push to branch: `git push origin feature/new-feature`
5. Submit a pull request

## 📝 License

This project is licensed under the MIT License.

## 🆘 Troubleshooting

### Common Issues

1. **Port conflicts**: Change ports in `docker-compose.yml` if 8081 or 3307 are already in use
2. **Database connection issues**: Ensure MySQL container is healthy before app starts
3. **Dependency issues**: Run `go mod tidy` to sync dependencies

### Getting Help

- Check existing issues on GitHub
- Create a new issue with detailed error messages and environment information
