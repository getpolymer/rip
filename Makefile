.PHONY: build run

build:
	go build -o rip .
	chmod +x rip

run: build
	./rip

docker-build:
	docker build -t rip .

docker-run:
	docker run -p 80:8080 rip
