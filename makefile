APP_NAME=iti-backend-challenge-filipeandrade6

test:
	go test ./... -count=1
	staticcheck -checks=all ./...

build:
	docker build -t $(APP_NAME):latest .

run:
	docker run -d -p 8080:8080 --rm --name=$(APP_NAME) $(APP_NAME)

up:	build run

stop:
	docker stop $(APP_NAME)