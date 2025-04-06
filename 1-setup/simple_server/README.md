# Working with Docker Containers

This guide outlines common tasks for working with Docker containers. Many of the necessary commands are included in the `Makefile`.

---

## 🐳 1.1 - Running the Container Image

To run your container, use the commands provided in the
`Makefile`. This ensures consistency and convenience.

```bash
make start
```

---

## 🔍 1.2 - Exploring the Inside of a Running Container

You can inspect a running container and explore its processes.

### Getting Additional Information About a Container

```bash
docker inspect <container-name>
```

### Executing Bash Inside a Running Container

```bash
docker exec -it <container-name> bash
```

Once inside the container, 
you can use common Linux commands to explore:

```bash
ps
#    PID TTY          TIME CMD
#     12 pts/0    00:00:00 bash
#     22 pts/0    00:00:00 ps
```

## 🫸 1.3 - Pushing image into Image Registry

Inorder to use that image inside our deployment, we need to
tag and push that image into an image registry.

```bash
docker tag golang-http-server:0.0.1 manouchehrrasouli96/golang-http-server:0.0.1
docker push manouchehrrasouli96/golang-http-server:0.0.1
```

## 🛑 1.4 - Stopping and Removing a Container

```bash
make spown
```

Using these steps, you can easily manage and
troubleshoot your containers.
