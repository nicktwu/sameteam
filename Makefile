build:
	cd database && make build
	cd backend && make build

start:
	cd database && make start
	cd backend && make start

stop:
	cd database && make stop
	cd backend && make stop

