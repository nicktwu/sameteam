

build:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w'
	docker build -t backend .
	
start: build
	docker run -d -p 9001:8080 backend /backend 
