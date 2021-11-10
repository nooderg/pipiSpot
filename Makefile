SHELL=/bin/bash

# define standard colors
ifneq (,$(findstring xterm,${TERM}))
	BLACK        := $(shell tput -Txterm setaf 0)
	RED          := $(shell tput -Txterm setaf 1)
	GREEN        := $(shell tput -Txterm setaf 2)
	YELLOW       := $(shell tput -Txterm setaf 3)
	LIGHTPURPLE  := $(shell tput -Txterm setaf 4)
	PURPLE       := $(shell tput -Txterm setaf 5)
	BLUE         := $(shell tput -Txterm setaf 6)
	WHITE        := $(shell tput -Txterm setaf 7)
	RESET := $(shell tput -Txterm sgr0)
else
	BLACK        := ""
	RED          := ""
	GREEN        := ""
	YELLOW       := ""
	LIGHTPURPLE  := ""
	PURPLE       := ""
	BLUE         := ""
	WHITE        := ""
	RESET        := ""
endif


help:
	@echo "# How to setup the project on your machine ? Run make dev in one tab, then run the api in another after you made your changes." 
	@echo "# ${PURPLE}init${RESET}: Use it to init the githooks" 
	@echo "# ${BLUE}dev${RESET}: run docker-compose (without api)" 
	@echo "# ${GREEN}prod${RESET}: run docker-compose (fully fonctionnal)" 
	@echo "# ${RED}stop${RESET}: stop docker containers" 
	@echo "# ${YELLOW}logs${RESET}: display logs and attach" 

init:
# Create git hooks
	git config core.hooksPath .githooks

dev:
	@echo "# ${BLUE}dev${RESET}: Running docker-compose without the API" 
	docker-compose \
		--file docker-compose.dev.yaml up \
		--detach \
		--build \
		--remove-orphans \
		--force-recreate \

stop:
	docker-compose -f docker-compose.dev.yaml stop 

logs:
	docker-compose logs --follow
