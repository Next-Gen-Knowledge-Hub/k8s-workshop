build:
	docker build -t manouchehrrasouli96/env-server:v0.2 .

run:
	docker run -d --name env-server -e PORT="1888" -e STAGE="docker" -p 8080:1888  manouchehrrasouli96/env-server:v0.2

stop-rm:
	docker stop env-server
	docker rm env-server

rmi:
	docker rmi env-server
