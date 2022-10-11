SERVER_BINARY=mopidy-monitor-go
CLIENT_BINARY=mopidy-control-go
BASE=$(CURDIR)

install:
	cd ${BASE}/monitor-server && GOARCH=amd64 GOOS=linux go build -o $(HOME)/.local/bin/${SERVER_BINARY} main.go
	cd ${BASE}/control-client && GOARCH=amd64 GOOS=linux go build -o $(HOME)/.local/bin/${CLIENT_BINARY} main.go
