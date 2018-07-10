rm boot.go
touch boot.go
go generate ./apps/...
go run *.go
