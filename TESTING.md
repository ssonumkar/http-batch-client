# Testing Examples

## Quick Tests

### Basic functionality test
```bash
./http-client -parallel 2 -count 10
```

### Test with httpbin POST endpoint
```bash
./http-client -url https://httpbin.org/post -method POST -parallel 5 -count 20
```

### High concurrency test
```bash
./http-client -parallel 50 -count 500
```

### Test different endpoints
```bash
# GET requests
./http-client -url https://httpbin.org/get -parallel 10 -count 100

# POST requests  
./http-client -url https://httpbin.org/post -method POST -parallel 10 -count 100

# Test with delay
./http-client -url https://httpbin.org/delay/1 -parallel 5 -count 25
```
