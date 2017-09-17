

build:
	go get -d -t ./...
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w'
	docker build -t backend .
	
start: stop build
	docker run --rm --name backend --link database:db -d -p 9001:8080 backend /backend 

stop:
	-docker stop backend

