up:
	@echo "Running migration up..."
	@direnv exec . goose -dir $(migrationPath) up
