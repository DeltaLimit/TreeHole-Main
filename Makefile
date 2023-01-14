WINDOWS = treehole-windows.exe
LINUX = treehole-linux

all:
	@echo "Building linux binary file..."
	go build -o $(LINUX) main.go
	@echo "Linux binary file build successful!"

	@echo "Building windows binary file..."
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64  go build -o $(WINDOWS) main.go
	@echo "Windows binary file build successful!"
