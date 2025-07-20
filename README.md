# HTTP Client

A simple and lightweight HTTP client for making concurrent HTTP requests with basic timing logs.


## Prerequisites

- Go 1.21 or higher

## Steps to Run the Application

### 1. Clone/Download the Repository
```bash
git clone <repository-url>
cd http-batch-client
```

### 2. Build the application
```bash
go build
```

### 3. Run the application 
```bash
go run main.go -url https://api.example.com/health -parallel 5 -count 50
```