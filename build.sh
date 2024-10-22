# Linux
GOOS=linux GOARCH=amd64 go build -o accio-linux-amd64 -ldflags="-w" cmd/accio/main.go

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o accio-mac-amd64 -ldflags="-w" cmd/accio/main.go

# macOS (Apple Silicon - M1/M2)
GOOS=darwin GOARCH=arm64 go build -o accio-mac-arm64 -ldflags="-w" cmd/accio/main.go
