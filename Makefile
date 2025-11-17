unittest:
	@echo "\n\n==================== Start unit test ...... ====================\n\n"
	@go test ./... --short -cover -race
	@echo "\n\n==================== Unit test done ====================\n\n"

run:
	@go run main.go rest