.PHONY: help gen fmt

run: # Run the application
	go run cmd/api/main.go -cfg_path=./conf/config.yaml

help: # Show this help message
	@grep -E '^[a-zA-Z_-]+:.*?# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

gen: # Generate code
	go run cmd/gen/gen.go

fmt: # Format code
	go fmt ./...