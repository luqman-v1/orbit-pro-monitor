build-windows-amd64-orbit:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o orbit-monitoring.exe ./cmd

run:
	go run cmd/main.go