run:
	@echo "---------------------\nRunning main.go\n---------------------"
	go run main.go
	@echo "\n"

test:
	@echo "---------------------\nRunning tests\n---------------------"
	go test strtoint/*
	@echo "\n"

bench:
	@echo "---------------------\nRunning benchmarks\n---------------------"
	go test strtoint/* -bench=.
	@echo "\n"

cover:
	@echo "---------------------\nRunning test c\n---------------------"
	mkdir -p tmp
	go test strtoint/* -coverprofile=./tmp/cover.out
	go tool cover -html=./tmp/cover.out -o ./tmp/coverout.html
	open ./tmp/coverout.html
	@echo "\n"

cover-clear:
	rm -rf ./tmp

all: run test bench cover