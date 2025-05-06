build:
	docker build -t manouchehrrasouli96/arg-server:v0.1 .

run:
	docker run -d --name arg-server -p 8888:8080  manouchehrrasouli96/arg-server:v0.1  arg0 arg1 arg2

stop-rm:
	docker stop arg-server
	docker rm arg-server

rmi:
	docker rmi arg-server
