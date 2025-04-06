
build-image:
	docker build -t golang-http-server:0.0.1 .

check-image:
	docker images | grep golang-http-server

remove-image:
	docker rmi golang-http-server:0.0.1

run-container:
	docker run -d --name go-server -p 9999:9999 golang-http-server:0.0.1

stop-container:
	docker stop go-server

remove-container:
	docker rm go-server

make-request:
	curl http://localhost:9999/hello

start:
	make build-image
	make run-container

spown:
	make stop-container
	make remove-container
	make remove-image
