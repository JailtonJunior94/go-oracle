.PHONY: migrate
migrate:
	@migrate create -ext sql -dir database/migrations -format unix $(NAME)