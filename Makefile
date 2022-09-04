build:
	go build -o bin/mocg ./main.go 
run:
	go run ./main.go 
test-player:
	go test ./internals/player/player.go ./internals/player/player_test.go  -v    
test-repositories:
	go test ./internals/repositories/repositories.go ./internals/repositories/repositories_test.go -v
test-ui:
	go test ./internals/ui/ui.go ./internals/ui/ui_test.go -v
test-decoder:
	go test ./internals/decoder/decoder.go ./internals/decoder/decoder_test.go -v
	
install:
	make build;
	cp ./bin/mocg /usr/bin;
	
	
