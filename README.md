# Kasir API

Backend RESTful API untuk sistem Point of Sales (POS) yang dibangun dengan Go, PostgreSQL, dan Clean Architecture.

---

## ✨ Highlights

- 🏗️ **Clean Architecture** dengan struktur `internal/` yang terorganisir
- 📊 **Structured Logging** menggunakan Logrus di setiap layer
- 🔄 **UseCase Pattern** untuk business logic yang jelas
- 🎯 **Interface-based** Repository untuk testability
- 🚀 **Production-ready** dengan proper error handling dan validation

---

## 🛠️ Tech Stack

| Layer | Technology |
|-------|-----------|
| **Language** | [Go 1.21+](https://go.dev/) |
| **Database** | [PostgreSQL](https://www.postgresql.org/) |
| **Logging** | [Logrus](https://github.com/sirupsen/logrus) |
| **Config** | [Viper](https://github.com/spf13/viper) |
| **Docs** | [Swagger/Swaggo](https://github.com/swaggo/swag) |
| **Dev Tools** | [Air](https://github.com/air-verse/air) (hot reload) |

---

## 📂 Project Structure

```
kasir-api/
├── cmd/
│   └── api/
│       └── main.go                    # Application entry point
├── internal/
│   ├── bootstrap/
│   │   └── bootstrap.go               # App initialization & DI
│   ├── config/
│   │   └── config.go                  # Configuration loader
│   ├── database/
│   │   └── database.go                # Database connection
│   ├── domain/
│   │   ├── models/                    # Domain models
│   │   │   ├── product.go
│   │   │   └── health.go
│   │   ├── repositories/              # Repository interfaces & implementations
│   │   │   └── product_repository.go
│   │   └── usecases/                  # Business logic
│   │       ├── product_usecase.go
│   │       └── health_usecase.go
│   ├── http/
│   │   ├── handlers/                  # HTTP handlers
│   │   │   ├── product_handler.go
│   │   │   ├── health_handler.go
│   │   │   └── category_handler.go
│   │   └── middleware/                # HTTP middlewares
│   │       └── logging.go
│   └── pkg/                           # Shared utilities
│       ├── logger.go
│       └── response.go
├── docs/                              # Swagger documentation
├── .env                               # Environment variables
├── .air.toml                          # Air configuration
└── go.mod
```

---

## 🏗️ Architecture Flow

### Request Flow (Handler → UseCase → Repository)

```
┌─────────────────────────────────────────────────────────────┐
│  1. HTTP Request                                            │
│     GET /api/product                                        │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│  2. Middleware Layer                                        │
│     internal/http/middleware/logging.go                     │
│     • Log incoming request                                  │
│     • Track duration                                        │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│  3. Handler Layer                                           │
│     internal/http/handlers/product_handler.go               │
│     • Parse HTTP request                                    │
│     • Validate input format                                 │
│     • Call use case                                         │
│     • Format HTTP response                                  │
│     • Log handler actions                                   │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│  4. UseCase Layer                                           │
│     internal/domain/usecases/product_usecase.go             │
│     • Business logic & validation                           │
│     • Call repository (via interface)                       │
│     • Error handling                                        │
│     • Log use case execution                                │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│  5. Repository Layer                                        │
│     internal/domain/repositories/product_repository.go      │
│     • Execute SQL queries                                   │
│     • Map database rows to models                           │
│     • Handle database errors                                │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│  6. Database (PostgreSQL)                                   │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│  7. Response (JSON)                                         │
│     {                                                       │
│       "status": "success",                                  │
│       "message": "Products retrieved successfully",         │
│       "data": [...]                                         │
│     }                                                       │
└─────────────────────────────────────────────────────────────┘
```

### Key Principles

1. **Dependency Injection**: Dependencies injected via constructors
2. **Interface-based**: Repository returns interface for testability
3. **Separation of Concerns**: Each layer has single responsibility
4. **Structured Logging**: Log at every layer with context

---

## 🚀 Quick Start

### 1. Prerequisites
- Go 1.21 or higher
- PostgreSQL database
- Air (optional, for hot reload)

### 2. Clone & Install
```bash
git clone https://github.com/crosbydoo/kasir-api.git
cd kasir-api
go mod tidy
```

### 3. Configuration
Create `.env` file:
```env
PORT=8080
DB_CONN=postgresql://user:password@localhost:5432/kasir_db?sslmode=disable
APP_ENV=development
```

### 4. Run Application

**Development (with hot reload):**
```bash
air
```

**Production:**
```bash
go build -o bin/api ./cmd/api
./bin/api
```

Server will run on `http://localhost:8080`

---

## 📊 Logging Example

### Startup Logs
```
INFO[2026-01-31 21:31:00] Starting Kasir API...
INFO[2026-01-31 21:31:00] Logger initialized successfully
INFO[2026-01-31 21:31:00] Loading configuration...
INFO[2026-01-31 21:31:00] Configuration loaded successfully    port=8080
INFO[2026-01-31 21:31:00] Connecting to database...
INFO[2026-01-31 21:31:00] Database connected successfully
INFO[2026-01-31 21:31:00] Initializing dependencies...
INFO[2026-01-31 21:31:00] All dependencies initialized successfully
INFO[2026-01-31 21:31:00] Registering routes...
INFO[2026-01-31 21:31:00] Routes registered successfully
INFO[2026-01-31 21:31:00] Starting HTTP server...    address="0.0.0.0:8080" port=8080
```

### Request Logs
```
INFO[21:31:15] Incoming request    method=GET path=/api/product remote_addr="127.0.0.1:54321"
INFO[21:31:15] Get all products handler called    handler=product_handler action=get_all_products
INFO[21:31:15] Executing get all products use case    usecase=product action=get_all_products
INFO[21:31:15] Successfully retrieved all products    usecase=product count=10
INFO[21:31:15] Products retrieved successfully    handler=product_handler count=10
INFO[21:31:15] Request completed successfully    status_code=200 duration_ms=15
```

---

## 📚 API Endpoints

### Health Check
```
GET /api/health
```

### Products
```
GET    /api/product           # Get all products
GET    /api/product/{id}      # Get product by ID
POST   /api/product           # Create product
PUT    /api/product/{id}      # Update product
DELETE /api/product/{id}      # Delete product
```

### Categories
```
GET    /api/category          # Get all categories
GET    /api/category/{id}     # Get category by ID
POST   /api/category          # Create category
PUT    /api/category/{id}     # Update category
DELETE /api/category/{id}     # Delete category
```

### Swagger Documentation
```
GET /swagger/index.html
```

---

## 📝 Request/Response Examples

### Create Product
**Request:**
```json
POST /api/product
{
  "name": "Laptop ASUS",
  "price": 15000000,
  "stock": 10
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Product created successfully",
  "data": {
    "id": 1,
    "name": "Laptop ASUS",
    "price": 15000000,
    "stock": 10
  }
}
```

### Get All Products
**Request:**
```
GET /api/product
```

**Response:**
```json
{
  "status": "success",
  "message": "Products retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Laptop ASUS",
      "price": 15000000,
      "stock": 10
    }
  ]
}
```

---

## 🎯 Key Features

### 1. Clean Architecture
- **Domain Layer**: Models, repositories (interfaces), use cases
- **HTTP Layer**: Handlers, middleware
- **Infrastructure**: Database, config, logger

### 2. Interface-based Repository
```go
// Interface definition
type ProductRepository interface {
    GetAllProduct() ([]models.Product, error)
    GetProductByID(id int) (*models.Product, error)
    CreateProduct(product *models.Product) error
    UpdateProduct(product *models.Product) error
    DeleteProduct(id int) error
}

// Constructor returns interface
func NewProductRepository(db *sql.DB) ProductRepository {
    return &productRepository{db: db}
}
```

### 3. UseCase Pattern
```go
// Business logic with validation
func (uc *productUseCase) CreateProduct(product *models.Product) error {
    // Validation
    if product.Name == "" {
        return errors.New("product name is required")
    }
    if product.Price < 0 {
        return errors.New("product price cannot be negative")
    }
    
    // Call repository
    return uc.productRepo.CreateProduct(product)
}
```

### 4. Structured Logging
```go
pkg.Log.WithFields(logrus.Fields{
    "usecase": "product",
    "action": "create_product",
    "product_name": product.Name,
}).Info("Creating product")
```

---

## 🧪 Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

---

## 📄 License

[MIT License](LICENSE)

---

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

---

## 📧 Contact

For questions or support, please open an issue on GitHub.
