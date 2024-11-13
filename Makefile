run: build
	@./bin/gocache

build:
	@go build -o bin/gocache

runfollower: build
	@./bin/gocache --listenaddr :4000 --leaderaddr :3000

test:
	go test -v ./...