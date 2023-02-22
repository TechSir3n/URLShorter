
all:clean build run 

.PHONY: build 
build:
	go build -o  /shorter/url ./main.go
	
.PHONY: run 
run: 
	go run main.go


.PHONY: compose-up
compose: compose-up
	docker-compose up

.PHONY: compose-down
compose-down:
	docker-compose down


.PHONY: docker-build
docker-build:
	docker build -t url-shorter .


.PHONY: docker-run 
docker-run:
	docker run  -d url-shorter

.PHONY: clean 
clean:
	rm -f /main 
