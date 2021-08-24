build:
	docker build --tag backup_scheduler:latest .
	docker-compose up -d

start:
	go mod vendor
	go build main.go
	./main