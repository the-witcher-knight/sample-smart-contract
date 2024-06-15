# Include all environment variables in .env.local
# Use - to ignore error when .env.local not found
-include .env.local
export

# Variables
DOCKER_COMPOSE=docker compose --file docker/docker-compose.yml --project-directory . -p ${PROJECT_NAME}

.PHONY: setup server deploy-contract
setup: ganache build-dev-image local-env

server:
	$(DOCKER_COMPOSE) run --service-ports --rm dev sh -c "go run cmd/serverd/main.go"

deploy-contract:
	$(DOCKER_COMPOSE) run --rm dev sh -c "go run cmd/deploycontract/main.go"


.PHONY: build-dev-image teardown abigen ganache local-env
build-dev-image:
	docker build -f docker/app.Dockerfile -t ${PROJECT_NAME}-dev:latest .
	-docker images -q -f "dangling=true" | xargs docker rmi -f

teardown:
	$(DOCKER_COMPOSE) down

local-env:
	@cp .env.sample .env.local

solc:
	$(DOCKER_COMPOSE) run --rm solc sh -c "\
		solc --evm-version berlin --overwrite --abi data/contracts/*.sol -o build/contracts && \
		solc --evm-version berlin --overwrite --bin data/contracts/*.sol -o build/contracts"

abigen:
	$(DOCKER_COMPOSE) run --rm dev sh -c "go generate ./..."

ganache:
	$(DOCKER_COMPOSE) up ganache -d