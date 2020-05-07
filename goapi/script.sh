GOOS=linux GOARCH=amd64 go build goapi.go
docker build -t avk19/goapi:latest .
docker push avk19/goapi:latest