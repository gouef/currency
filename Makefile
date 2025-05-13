.PHONY: install tests

install:
	go mod tidy && go mod vendor

tests:
	go test -covermode=set -coverpkg=./... -coverprofile=coverage.txt ./tests && go tool cover -func=coverage.txt
coverage:
	go test -v -coverpkg=./... -covermode=set -coverprofile=coverage.txt ./tests && go tool cover -html=coverage.txt -o coverage.html && xdg-open coverage.html
