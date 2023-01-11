.PHONY: clean
.PHONY: build
build: build/rename
build/rename: cmd/rename/main.go internal/rename/* 
	go build -o build/ ./cmd/rename
clean:
	rm -r build

.PHONY: build-windows
build-windows: build/rename.exe
build/rename.exe: export GOOS=windows
build/rename.exe: export GOARCH=amd64
build/rename.exe: cmd/rename/main.go internal/rename/*
	go build -o build/ ./cmd/rename