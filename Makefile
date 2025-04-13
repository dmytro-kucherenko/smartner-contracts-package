proto-gen:
	@buf generate

lint:
	@go vet ./...

pre-commit:
	@pre-commit autoupdate && pre-commit install
