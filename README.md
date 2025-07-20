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

### 2. Build and Run the application
```bash
chmod +x run-app.sh
./run-app.sh
```
### 3. Pass the command line arguments to configure the Parallel Request Count, Total Requests, URL
````bash
go run main.go -url https://api.example.com/health -parallel 20 - count 200
````
