build:
	go build -o bin/ .

start: 
	cd bin && ./golang-kata-1


restart: build start

run:
	go run main.go

.PHONY: build start restart run
