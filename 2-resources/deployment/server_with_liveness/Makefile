build-image:
	docker build -t http-server-healthcheck:0.0.1 .

tag-push:
	docker tag http-server-healthcheck:0.0.1 manouchehrrasouli96/http-server-healthcheck:0.0.1
    docker push manouchehrrasouli96/http-server-healthcheck:0.0.1

run-container:
	docker run -d --name go-server -p 8989:8989 manouchehrrasouli96/go-http-health-check:0.0.1

stop-container:
	docker stop go-server

remove-container:
	docker rm go-server

make-request:
	curl http://localhost:8989/healthcheck

spown:
	make stop-container
	make remove-container
