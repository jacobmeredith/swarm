install_dependencies:
	@echo "Installing dependencies"
	@go get -d ./...

build_cli:
	@echo "Building CLI Application"
	@go build -o bin/swarm main.go

build_windows:
	@echo "Building Windows Executable"
	@GOOS=windows GOARCH=amd64 go build -o bin/swarm.exe main.go

release: build_cli build_windows
