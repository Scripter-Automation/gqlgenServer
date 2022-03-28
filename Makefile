include .envrc

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
	.PHONY: confirm confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/dev: run the cmd/api application
.PHONY: run
run:
	go run ./cmd/
compile:
	cd cmd && go run github.com/99designs/gqlgen compile
# ==================================================================================== #
# Deploy
# ==================================================================================== #

PHONY: deploy/dev
deploy/dev:
	@echo 'Deploying to docker container'
	docker-compose -f docker-compose.dev.yml up --build -d


# ==================================================================================== #
# Deploy
# ==================================================================================== #