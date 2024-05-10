restart: backend-restart frontend-restart database-restart
rebuild: delete backend-rebuild frontend-rebuild database-rebuild
up: backend-up frontend-up database-up
down: backend-down frontend-down database-down
delete:
	docker-compose -f ./database/docker-compose.yml down --rmi all
build: backend-build frontend-build database-build

# todo rewrite with docker commond
backend-restart: backend-down backend-up
backend-rebuild:
	docker-compose -f ./backend/docker-compose.yml up -d --build
backend-up:
	docker-compose -f ./backend/docker-compose.yml up -d
backend-down:
	docker-compose -f ./backend/docker-compose.yml down
backend-build:
	docker-compose -f ./backend/docker-compose.yml build
backend-exec:
	docker exec -it neuza-backend sh

frontend-restart: frontend-down frontend-up
frontend-up:
	docker-compose -f ./frontend/docker-compose.yml up -d
frontend-down:
	docker-compose -f ./frontend/docker-compose.yml down
frontend-rebuild:
	docker-compose -f ./frontend/docker-compose.yml up -d --build
frontend-build:
	docker-compose -f ./frontend/docker-compose.yml build

database-restart: database-down database-up
database-up:
	docker-compose -f ./database/docker-compose.yml up -d
database-down:
	docker-compose -f ./database/docker-compose.yml down
database-build:
	docker-compose -f ./database/docker-compose.yml build
database-rebuild:
	docker-compose -f ./database/docker-compose.yml up -d --build