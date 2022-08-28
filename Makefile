build:
	go build -o bin/mocg ./cmd/main.go 
run:
	go run ./cmd/main.go 
test-player:
	go test ./internals/player/player.go ./internals/player/player_test.go  -v    
test-repositories:
	go test ./internals/repositories/repositories.go ./internals/repositories/repositories_test.go -v
test-ui:
	go test ./internals/ui/ui.go ./internals/ui/ui_test.go -v
test-decoder:
	go test ./internals/decoder/decoder.go ./internals/decoder/decoder_test.go -v
test:
	echo "Testing repositories package";
	go test ./internals/repositories/repositories.go ./internals/repositories/repositories_test.go -v;
	echo "Testing decoder package";
	go test ./internals/decoder/decoder.go ./internals/decoder/decoder_test.go -v;
	echo "Testing player package";
	go test ./internals/player/player.go ./internals/player/player_test.go  -v    



	