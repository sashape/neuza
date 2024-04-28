restart: down up
rebuild: delete build up
up:
	docker-compose up -d
down:
	docker-compose down --remove-orphans
delete:
	docker-compose down --rmi all
build:
	docker-compose build