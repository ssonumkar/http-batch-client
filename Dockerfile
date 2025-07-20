FROM golang:1.21-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod tidy
RUN go mod download
COPY . .
RUN go build -o http-client
ENTRYPOINT ["./http-client"]