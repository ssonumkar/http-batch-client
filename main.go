package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

func main() {
	parallel := flag.Int("parallel", 10, "Number of parallel requests")
	count := flag.Int("count", 100, "Total number of requests")
	url := flag.String("url", "https://httpbin.org/get", "Target URL")
	method := flag.String("method", "GET", "HTTP method")

	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		cancel()
	}()

	fmt.Printf("Starting %d requests with %d parallel workers...\n", *count, *parallel)
	start := time.Now()

	executeRequests(ctx, *parallel, *count, *url, *method)

	duration := time.Since(start)
	fmt.Printf("Completed %d requests in %v\n", *count, duration)
	fmt.Printf("Average time per request: %v\n", duration/time.Duration(*count))
	fmt.Printf("Requests per second: %.2f\n", float64(*count)/duration.Seconds())
}

func executeRequests(ctx context.Context, parallel, count int, url, method string) {
	var wg sync.WaitGroup
	client := &http.Client{Timeout: 30 * time.Second}

	sem := make(chan struct{}, parallel)

	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			return
		default:
		}

		wg.Add(1)
		sem <- struct{}{}

		go func(id int) {
			defer wg.Done()
			defer func() { <-sem }()

			makeRequest(ctx, client, url, method, id)
		}(i)
	}

	wg.Wait()
}

func makeRequest(ctx context.Context, client *http.Client, url, method string, id int) {
	var body io.Reader
	if method == "POST" || method == "PUT" {
		body = strings.NewReader("{}")
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	io.ReadAll(resp.Body)
}
