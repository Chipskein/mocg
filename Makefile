build:
	go build -o bin/mocg ./main.go 
run:
	go run ./main.go 
test-player:
	go test ./internal/player/player.go ./internal/player/player_test.go  -v    
test-repositories:
	go test ./internal/repositories/repositories.go ./internal/repositories/repositories_test.go -v
test-ui:
	go test ./internal/ui/ui.go ./internal/ui/ui_test.go -v
test-decoder:
	go test ./internal/decoder/decoder.go ./internal/decoder/decoder_test.go -v
install:
	make build;
	cp ./bin/mocg /usr/bin;
uninstall:
	rm /usr/bin/mocg
	
	
