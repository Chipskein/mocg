build:
	go build -o bin/mocg ./cmd/main.go 
run:
	go run ./cmd/main.go 
test-player:
	go test ./internals/player/player.go ./internals/player/player_test.go  -v    
test-repositories:
	go test ./internals/repositories/repositories.go ./internals/repositories/repositories_test.go -v


	