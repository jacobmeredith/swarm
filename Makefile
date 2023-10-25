install_dependencies:
	@echo "Installing dependencies"
	@go get -d ./...

build_cli:
	@echo "Building CLI Application"
	@go build -o bin/swarm main.go
