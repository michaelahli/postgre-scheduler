build:
	docker build --tag backup_scheduler:latest .
	docker-compose up -d

start:
	go mod vendor
	go build main.go
	./main

logs:
	docker logs --tail 200 backup_scheduler

dir:
	docker exec -it backup_scheduler sh

rm:
	docker image rm backup_scheduler:latest