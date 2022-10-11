SERVER_BINARY=mopidy-volume-server
CLIENT_BINARY=mopidy-volume-control
BASE=$(CURDIR)

install:
	cd ${BASE}/volume-server && GOARCH=amd64 GOOS=linux go build -o $(HOME)/.local/bin/${SERVER_BINARY} main.go
	cd ${BASE}/volume-control && GOARCH=amd64 GOOS=linux go build -o $(HOME)/.local/bin/${CLIENT_BINARY} main.go
