# Benchmark

go test -bench="." ./...
go test -bench="." -benchmem ./...

# Coverage
go test ./... -cover
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
go tool cover -html=coverage.out
