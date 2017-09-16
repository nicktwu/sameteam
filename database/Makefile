start:
	docker build -t database .
	docker run -v $(PWD)/data:/data/db --rm --name database -d -p 9002:8080 database

stop:
	-docker stop database
