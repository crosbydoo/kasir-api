# Logging dengan Logrus - Clean Architecture

## 📋 Overview

Implementasi logging menggunakan **Logrus** dengan **Clean Architecture** pattern yang memisahkan concern antara Handler, Use Case, dan logging middleware.

## 🏗️ Struktur Arsitektur

```
┌─────────────────────────────────────────────────────────────┐
│                         CLIENT                               │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│                   MIDDLEWARE LAYER                           │
│  - LoggingMiddleware (mencatat request & response)          │
│  - Log: method, path, status_code, duration, user_agent     │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│                    HANDLER LAYER                             │
│  - HealthHandler.CheckHealth()                              │
│  - Log: handler, action, method                             │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│                   USE CASE LAYER                             │
│  - HealthUseCase.CheckHealth()                              │
│  - Log: usecase, action, status                             │
│  - Business Logic                                            │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│                    RESPONSE                                  │
│  - JSON Response dengan status, timestamp, dll              │
└─────────────────────────────────────────────────────────────┘
```

## 🔄 Request Flow dengan Logging

### Flow Diagram
```
1. Client Request
   ↓
2. Middleware: Log "Incoming request"
   Fields: method, path, remote_addr, user_agent
   ↓
3. Handler: Log "Health check handler called"
   Fields: handler, action, method
   ↓
4. Use Case: Log "Executing health check use case"
   Fields: usecase, action
   ↓
5. Use Case: Log "Health check completed successfully"
   Fields: usecase, action, status
   ↓
6. Handler: Log "Health check handler completed"
   Fields: handler, action, status
   ↓
7. Middleware: Log "Request completed successfully"
   Fields: method, path, status_code, duration_ms, remote_addr
   ↓
8. Response to Client
```

## 📁 Struktur File

```
kasir-api/
├── pkg/
│   └── logger.go              # Konfigurasi Logrus
├── middleware/
│   └── logging.go             # Middleware untuk log request/response
├── models/
│   └── health.go              # Model HealthResponse
├── usecases/
│   └── health_usecase.go      # Business logic health check
├── handlers/
│   └── health_handler.go      # HTTP handler health check
└── main.go                    # Entry point & routing
```

## 🎯 Implementasi per Layer

### 1. Logger Configuration (`pkg/logger.go`)

```go
var Log *logrus.Logger

func InitLogger() {
    Log = logrus.New()
    
    // Format berdasarkan environment
    if env == "production" {
        Log.SetFormatter(&logrus.JSONFormatter{})
    } else {
        Log.SetFormatter(&logrus.TextFormatter{
            FullTimestamp: true,
            ForceColors: true,
        })
    }
}
```

**Fitur:**
- ✅ JSON format untuk production
- ✅ Text format dengan warna untuk development
- ✅ Timestamp lengkap
- ✅ Log level configuration

### 2. Middleware Layer (`middleware/logging.go`)

```go
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Log incoming request
        pkg.Log.WithFields(logrus.Fields{
            "method": r.Method,
            "path": r.URL.Path,
            "remote_addr": r.RemoteAddr,
            "user_agent": r.UserAgent(),
        }).Info("Incoming request")
        
        // Execute handler
        next(wrapped, r)
        
        // Log response
        duration := time.Since(start)
        pkg.Log.WithFields(logrus.Fields{
            "method": r.Method,
            "path": r.URL.Path,
            "status_code": wrapped.statusCode,
            "duration_ms": duration.Milliseconds(),
        }).Info("Request completed successfully")
    }
}
```

**Logging Fields:**
- `method`: HTTP method (GET, POST, dll)
- `path`: URL path yang diakses
- `remote_addr`: IP address client
- `user_agent`: Browser/client info
- `status_code`: HTTP status code response
- `duration_ms`: Waktu pemrosesan dalam milliseconds

### 3. Handler Layer (`handlers/health_handler.go`)

```go
func (h *HealthHandler) CheckHealth(w http.ResponseWriter, r *http.Request) {
    pkg.Log.WithFields(logrus.Fields{
        "handler": "health_handler",
        "action": "check_health",
        "method": r.Method,
    }).Info("Health check handler called")
    
    // Call use case
    response, err := h.healthUseCase.CheckHealth()
    
    if err != nil {
        pkg.Log.WithFields(logrus.Fields{
            "handler": "health_handler",
            "action": "check_health",
            "error": err.Error(),
        }).Error("Failed to check health")
        return
    }
    
    pkg.Log.WithFields(logrus.Fields{
        "handler": "health_handler",
        "action": "check_health",
        "status": response.Status,
    }).Info("Health check handler completed")
    
    pkg.SuccessResponse(w, http.StatusOK, "Health check successful", response)
}
```

**Logging Fields:**
- `handler`: Nama handler yang dipanggil
- `action`: Action yang dilakukan
- `method`: HTTP method
- `status`: Status dari response
- `error`: Error message jika ada

### 4. Use Case Layer (`usecases/health_usecase.go`)

```go
func (h *healthUseCase) CheckHealth() (*models.HealthResponse, error) {
    pkg.Log.WithFields(logrus.Fields{
        "usecase": "health_check",
        "action": "check_health",
    }).Info("Executing health check use case")
    
    response := &models.HealthResponse{
        Status: "healthy",
        Timestamp: time.Now(),
        Service: h.serviceName,
        Version: h.version,
    }
    
    pkg.Log.WithFields(logrus.Fields{
        "usecase": "health_check",
        "action": "check_health",
        "status": response.Status,
    }).Info("Health check completed successfully")
    
    return response, nil
}
```

**Logging Fields:**
- `usecase`: Nama use case
- `action`: Action yang dilakukan
- `status`: Status hasil eksekusi

## 📊 Contoh Output Log

### Development Mode (Text Format)
```
INFO[2026-01-28 21:43:00] Logger initialized successfully              
INFO[2026-01-28 21:43:00] Database connected successfully              
INFO[2026-01-28 21:43:00] All handlers initialized successfully        
INFO[2026-01-28 21:43:00] Starting server                               address="0.0.0.0:8080"
INFO[2026-01-28 21:43:15] Incoming request                              method=GET path=/api/health remote_addr="127.0.0.1:54321" user_agent="curl/7.64.1"
INFO[2026-01-28 21:43:15] Health check handler called                   action=check_health handler=health_handler method=GET
INFO[2026-01-28 21:43:15] Executing health check use case               action=check_health usecase=health_check
INFO[2026-01-28 21:43:15] Health check completed successfully           action=check_health status=healthy usecase=health_check
INFO[2026-01-28 21:43:15] Health check handler completed                action=check_health handler=health_handler status=healthy
INFO[2026-01-28 21:43:15] Request completed successfully                duration_ms=2 method=GET path=/api/health remote_addr="127.0.0.1:54321" status_code=200
```

### Production Mode (JSON Format)
```json
{"level":"info","msg":"Logger initialized successfully","time":"2026-01-28T21:43:00+07:00"}
{"address":"0.0.0.0:8080","level":"info","msg":"Starting server","time":"2026-01-28T21:43:00+07:00"}
{"level":"info","method":"GET","msg":"Incoming request","path":"/api/health","remote_addr":"127.0.0.1:54321","time":"2026-01-28T21:43:15+07:00","user_agent":"curl/7.64.1"}
{"action":"check_health","handler":"health_handler","level":"info","method":"GET","msg":"Health check handler called","time":"2026-01-28T21:43:15+07:00"}
{"action":"check_health","level":"info","msg":"Executing health check use case","time":"2026-01-28T21:43:15+07:00","usecase":"health_check"}
{"action":"check_health","level":"info","msg":"Health check completed successfully","status":"healthy","time":"2026-01-28T21:43:15+07:00","usecase":"health_check"}
{"action":"check_health","handler":"health_handler","level":"info","msg":"Health check handler completed","status":"healthy","time":"2026-01-28T21:43:15+07:00"}
{"duration_ms":2,"level":"info","method":"GET","msg":"Request completed successfully","path":"/api/health","remote_addr":"127.0.0.1:54321","status_code":200,"time":"2026-01-28T21:43:15+07:00"}
```

## 🧪 Testing

### Test Health Endpoint
```bash
curl -X GET http://localhost:8080/api/health
```

### Expected Response
```json
{
  "status": "success",
  "message": "Health check successful",
  "data": {
    "status": "healthy",
    "timestamp": "2026-01-28T21:43:15.123456+07:00",
    "service": "Kasir API",
    "version": "1.0.0"
  }
}
```

## 🎨 Log Levels

Logrus menggunakan log level otomatis berdasarkan kondisi:

| Status Code | Log Level | Contoh |
|-------------|-----------|--------|
| 200-299 | INFO | Request berhasil |
| 400-499 | WARN | Client error (bad request, not found) |
| 500-599 | ERROR | Server error |

## 🔍 Keuntungan Structured Logging

1. **Mudah di-parse**: Format JSON mudah diproses oleh tools monitoring
2. **Searchable**: Field-based search lebih mudah
3. **Contextual**: Setiap log punya context yang jelas
4. **Traceable**: Bisa track request dari awal sampai akhir
5. **Production-ready**: Siap untuk ELK Stack, Datadog, dll

## 📝 Best Practices

1. **Gunakan Fields**: Selalu gunakan `WithFields()` untuk context
2. **Consistent Naming**: Gunakan naming convention yang konsisten
3. **Log Level**: Gunakan level yang sesuai (Info, Warn, Error)
4. **Sensitive Data**: Jangan log password, token, dll
5. **Performance**: Log seperlunya, jangan berlebihan

## 🚀 Next Steps

Untuk menambahkan logging ke endpoint lain:

1. Buat use case dengan logging
2. Buat handler dengan logging
3. Wrap route dengan `LoggingMiddleware`
4. Test dan lihat log output

## 📚 Resources

- [Logrus Documentation](https://github.com/sirupsen/logrus)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
