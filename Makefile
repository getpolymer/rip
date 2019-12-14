.PHONY build run

build:
	go build -o rip .

run: build
	./rip
