# key value db


## installation

```bash
git clone git@github.com:NikoNodaros/kvdb.git
cd kvdb
go mod tidy
```

## usage

```go
go build
go run kvdb
```

Output:
```go
Server is running on port 8080
```

## sets the key 'k' with the value 'v'
```bash
curl -X PUT http://localhost:8080/k -d "v"
```
## returns the value for key 'k'
```bash
curl http://localhost:8080/k
```
## deletes the key 'k'
```bash
curl -X DELETE http://localhost:8080/k
```
## returns all keys
```bash
curl http://localhost:8080/
```

## testing
```bash
go test ./...
```