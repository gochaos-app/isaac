.PHONY : dev prod install compile

dev:
	echo "Compiling..."
	go build -o isaac . 

prod:
	echo "Compiling..."
	go build -ldflags="-s -w" -o isaac .

move:
	mv isaac ~/bin/isaac

install: prod move

compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -o bin/isaac-freebsd-amd64 .
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/isaac-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o bin/isaac-linux-arm64 .
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o bin/isaac-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o bin/isaac-darwin-m1 .
