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
### 3. Optionally run using Docker 
````bash
docker build -t go-client .
docker run --rm go-client -url http://localhost:8080/health -parallel 20 -count 200
````
