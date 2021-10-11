init:
# Create git hooks
	git config core.hooksPath .githooks

dev:
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