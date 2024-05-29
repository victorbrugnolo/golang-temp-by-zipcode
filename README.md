# Golang temp by zipcode

A tool for get weather of location by zipcode

## Config

```bash
go mod tidy
```

### Running

```bash
go run cmd/main.go
```
#### Or
```bash
docker compose up -d
```

The server will be running on port `:8080`

## Usage

### Local

```bash
curl http://localhost:8080/01153000/temperature
```

### GCP
```bash
curl http://localhost:8080/01153000/temperature
```

The path parameter zipcode (01153000 on example), is mandatory