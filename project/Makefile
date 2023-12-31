FRONT_END_BINARY=frontApp
GAME_API_BINARY=gameApiApp
AUTH_API_BINARY=authApp
SESSION_REDIS_SERVICE_BINARY=sessionRedisServiceApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_game_api build_auth_api build_session_redis_service_api
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_game_api: builds the game-api as a linux executable
build_game_api:
	@echo "Building game-api binary..."
	cd ../game-api && env GOOS=linux CGO_ENABLED=0 go build -o ${GAME_API_BINARY} ./cmd/api
	@echo "Done!"

## build_auth_api: builds the auth-api as a linux executable
build_auth_api:
	@echo "Building auth-api binary..."
	cd ../authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_API_BINARY} ./cmd/api
	@echo "Done!"

## build_session_redis_service_api: builds the session_redis_service_api as a linux executable
build_session_redis_service_api:
	@echo "Building session_redis_service_api binary..."
	cd ../session-redis-service && env GOOS=linux CGO_ENABLED=0 go build -o ${SESSION_REDIS_SERVICE_BINARY} ./cmd/api
	@echo "Done!"

## build_front: builds the frone end binary
build_front:
	@echo "Building front end binary..."
	cd ../front-end && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

## start: starts the front end
start: build_front
	@echo "Starting front end"
	cd ../front-end && ./${FRONT_END_BINARY} &

## stop: stop the front end
stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped front end!"