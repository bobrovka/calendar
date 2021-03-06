gen.proto:
	protoc --go_out=plugins=grpc:pkg/calendar api/api.proto

postgres.run:
	docker run --name calendar-postgres -e POSTGRES_DB=calendar -e POSTGRES_PASSWORD=password -d -p 5432:5432 postgres:12.2

postgres.start:
	docker start calendar-postgres

postgres.stop:
	docker stop calendar-postgres

rabbit.run:
	docker run --name calendar-rabbit -d -p 15672:15672 -p 5672:5672 rabbitmq:3-management

migrate:
	docker run -v $(shell pwd)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgresql://postgres:password@localhost:5432/calendar?sslmode=disable up

compile.calendar:
	go build -o calendar cmd/calendar/main.go

compile.sender:
	go build -o sender cmd/sender/main.go

build:
	make compile.calendar
	make compile.sender
