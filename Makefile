init:
# Create git hooks
	git config core.hooksPath .githooks

dev:
	docker-compose up \
		--detach \
		--build \
		--remove-orphans \
		--force-recreate

stop:
	docker-compose stop

logs:
	docker-compose logs --follow