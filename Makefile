.PHONY:	run	build	clean

run:
	go run cmd/main.go

build:
	go build -o bin/bot.exe	cmd/main.go

clean:
	rm -rf bin/