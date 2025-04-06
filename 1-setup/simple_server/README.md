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

---

## 📦 Running Your First App on Kubernetes

This guide introduces you to deploying a 
basic application on Kubernetes using the simplest method.  
Typically, Kubernetes deployments are
defined using YAML/JSON manifests,
but for learning purposes, 
we'll use one-liner `kubectl` commands.

---

## 2.1 - Deploying the `HTTP-Go` App

Run your app using a one-liner command:

```bash
kubectl run go-http \
  --image=manouchehrrasouli96/golang-http-server:0.0.1 \
  --port=9999 \
  --labels app=demo
```

Check pod status:

```bash
kubectl get pods
# NAME      READY   STATUS              RESTARTS   AGE
# go-http   0/1     ContainerCreating   0          62s
```

### 💡 Introducing Pods

Each **pod** has its own IP and
may contain one or more containers
sharing resources.  
They are the smallest deployable unit
in Kubernetes and are distributed across worker nodes.

```bash
kubectl describe pods go-http
```

After the container starts running:

```bash
kubectl get pods
# NAME      READY   STATUS    RESTARTS   AGE
# go-http   1/1     Running   0          3m22s
```

You can inspect again:

```bash
kubectl describe pods go-http
```

---

## 2.2 - Testing the App via Port Forwarding

Port forward the pod port to your local machine:

```bash
kubectl port-forward go-http 9999:9999
# Forwarding from 127.0.0.1:9999 -> 9999
# Forwarding from [::1]:9999 -> 9999

curl http://localhost:9999/hello
# request received on host go-http !

kubectl logs -f pods/go-http 
# 2025/04/06 11:36:45 server is running...
# 2025/04/06 11:39:04 &{Method:GET URL:/hello Proto:HTTP/1.1 ProtoMajor:1 ProtoMinor:1 Header:map[Accept:[*/*] User-Agent:[curl/8.5.0]] Body:{} GetBody:<nil> ContentLength:0 TransferEncoding:[] Close:false Host:localhost:9999 Form:map[] PostForm:map[] MultipartForm:<nil> Trailer:map[] RemoteAddr:127.0.0.1:45690 RequestURI:/hello TLS:<nil> Cancel:<nil> Response:<nil> Pattern:/hello ctx:0xc0001701e0 pat:0xc0001301e0 matches:[] otherValues:map[]}
```

⚠️ **Note:** Port-forwarding is
not recommended for production.
It's good for local testing only.

---

## 2.3 - Accessing the App from Outside the Cluster

Pods have internal cluster IPs.
To expose the app externally, use a `NodePort`:

```bash
kubectl expose pod go-http \
  --type=NodePort \
  --port=9999 \
  --target-port=9999
# service/go-http exposed

kubectl get service
# NAME         TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
# go-http      NodePort    10.111.139.133   <none>        9999:32673/TCP   2s
```

Get the node IP:

```bash
kubectl get nodes -o wide
# NAME       STATUS   ROLES           AGE   VERSION   INTERNAL-IP    EXTERNAL-IP   OS-IMAGE             KERNEL-VERSION     CONTAINER-RUNTIME
# minikube   Ready    control-plane   60m   v1.31.0   192.168.49.2   <none>        Ubuntu 22.04.4 LTS   6.8.0-49-generic   docker://27.2.0

curl http://192.168.49.2:32673/hello
# request received on host go-http !

kubectl logs pods/go-http 
# 2025/04/06 11:36:45 server is running...
# 2025/04/06 11:39:04 &{Method:GET URL:/hello Proto:HTTP/1.1 ProtoMajor:1 ProtoMinor:1 Header:map[Accept:[*/*] User-Agent:[curl/8.5.0]] Body:{} GetBody:<nil> ContentLength:0 TransferEncoding:[] Close:false Host:localhost:9999 Form:map[] PostForm:map[] MultipartForm:<nil> Trailer:map[] RemoteAddr:127.0.0.1:45690 RequestURI:/hello TLS:<nil> Cancel:<nil> Response:<nil> Pattern:/hello ctx:0xc0001701e0 pat:0xc0001301e0 matches:[] otherValues:map[]}
# 2025/04/06 11:43:08 &{Method:GET URL:/hello Proto:HTTP/1.1 ProtoMajor:1 ProtoMinor:1 Header:map[Accept:[*/*] User-Agent:[curl/8.5.0]] Body:{} GetBody:<nil> ContentLength:0 TransferEncoding:[] Close:false Host:192.168.49.2:31717 Form:map[] PostForm:map[] MultipartForm:<nil> Trailer:map[] RemoteAddr:10.244.0.1:58452 RequestURI:/hello TLS:<nil> Cancel:<nil> Response:<nil> Pattern:/hello ctx:0xc0000ac050 pat:0xc0001301e0 matches:[] otherValues:map[]}
```

Or use Minikube’s built-in service routing:

```bash
minikube service list

# | NAMESPACE  |    NAME    | TARGET PORT  |            URL            |
# |------------|------------|--------------|---------------------------|
# | default    | go-http    |         9999 | http://192.168.49.2:30885 |
```

You can also open the service in your browser:

```bash
minikube service list 
# |----------------------|---------------------------|--------------|---------------------------|
# |      NAMESPACE       |           NAME            | TARGET PORT  |            URL            |
# |----------------------|---------------------------|--------------|---------------------------|
# | default              | go-http                   |         9999 | http://192.168.49.2:31717 |
# | default              | kubernetes                | No node port |                           |
# | kube-system          | kube-dns                  | No node port |                           |
# | kubernetes-dashboard | dashboard-metrics-scraper | No node port |                           |
# | kubernetes-dashboard | kubernetes-dashboard      | No node port |                           |
# |----------------------|---------------------------|--------------|---------------------------|
```

---

## 2.4 - Understanding Kubernetes Components

| Component   | Description |
|-------------|-------------|
| **Pod**     | A group of one or more containers that share storage, network, and a specification for how to run them. |
| **Deployment** | Manages a set of replicated Pods, handles updates and scaling automatically. |
| **Service** | A stable abstraction that exposes a group of Pods as a network service. It load-balances and routes traffic. |

---

## 2.5 - Scaling Your Application

First, create a deployment:

```bash
kubectl create deployment go-http \
  --image=manouchehrrasouli96/golang-http-server:0.0.1 \
  --port=9999
```

Scale the deployment:

```bash
kubectl scale deployment go-http --replicas=3
```

Expose it with NodePort:

```bash
kubectl expose deployment go-http \
  --type=NodePort \
  --port=9999 \
  --target-port=9999
```

Check services:

```bash
kubectl get services
```

Test the load-balancing behavior by curling multiple times:

```bash
curl http://192.168.49.2:<NODE-PORT>/hello
# request received on host go-http-b84b95bf9-rpkgx !
# request received on host go-http-b84b95bf9-szqtq !
# request received on host go-http-b84b95bf9-rpkgx !
```

---

## 2.6 - Examining App Placement on Nodes

To see where pods are running:

```bash
kubectl get pods -o wide
```

Get detailed info for a specific pod:

```bash
kubectl describe pod <pod-name>
```

This helps visualize distribution of pods across the cluster’s nodes.

---

## 2.7 - Remove Services and Pods

```bash
kubectl delete service go-http 
# service "go-http" deleted

minikube service list 
# |----------------------|---------------------------|--------------|-----|
# |      NAMESPACE       |           NAME            | TARGET PORT  | URL |
# |----------------------|---------------------------|--------------|-----|
# | default              | kubernetes                | No node port |     |
# | kube-system          | kube-dns                  | No node port |     |
# | kubernetes-dashboard | dashboard-metrics-scraper | No node port |     |
# | kubernetes-dashboard | kubernetes-dashboard      | No node port |     |
# |----------------------|---------------------------|--------------|-----|

kubectl delete pod go-http 
# pod "go-http" deleted

kubectl get pods 
# No resources found in default namespace.
```

## ✅ Summary

- You deployed a simple Go app on Kubernetes.
- You explored how to access it from inside and outside the cluster.
- You scaled the app and saw how Kubernetes load balances between replicas.
- You learned about Pods, Deployments, and Services — key building blocks of Kubernetes.
