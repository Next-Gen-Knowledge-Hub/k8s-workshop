# 🚀 Kubernetes: Replication & Other Controllers - Managing Pods Efficiently

Pods are the smallest deployable units in Kubernetes. While you can create and manage them manually, in production scenarios, you'll almost always want your workloads to be self-healing and managed automatically.

---

## 1. 🩺 Keeping Pods Healthy

### 1.1 - Liveness Probes

Some applications (e.g., a Java app with a memory leak) may seem alive from the OS's perspective but are logically broken. Kubernetes addresses this with **liveness probes**, which periodically check if a container is healthy. If not, the container is restarted.

There are 3 types of liveness probes:

- **HTTP GET**: Makes an HTTP request and checks for a 2xx or 3xx response.
- **TCP Socket**: Tries to open a TCP connection.
- **Exec**: Runs a command inside the container and checks for exit code 0.

> You can define probes per container in the pod spec.

---

### 1.2 - Create an HTTP-Based Liveness Probe

You already built a Go HTTP server that fails after some requests. To deploy it:
Implement is in `service_with_liveness/server.go` build images and push to your image registry anf follow the steps.

```bash
kubectl run go-healthcheck --image manouchehrrasouli96/go-http-health-check:0.0.1 --port 8989
kubectl expose pod go-healthcheck --type NodePort --port 8989 --targetPort 8989
minikube service list
```

---

### 1.3 - Seeing the Probe in Action

Apply the pod manifest:

```bash
kubectl create -f pod-go-http-liveness.yaml
kubectl get pods
# NAME               READY   STATUS    RESTARTS   AGE
# go-http-liveness   1/1     Running   0          8s
```

You should see the container restarting after liveness probe failures:

```bash
kubectl get pods 
# NAME               READY   STATUS    RESTARTS     AGE
# go-http-liveness   1/1     Running   1 (4s ago)   74s
```

View detailed probe info:

```bash
kubectl describe pod go-http-liveness
# ...
# Warning  Unhealthy  25s (x3 over 45s)  kubelet            Liveness probe failed: HTTP probe failed with statuscode: 500
# Normal   Killing    25s                kubelet            Container go-http-liveness failed liveness probe, will be restarted
```

Liveness details might look like:

```
Liveness: http-get http://:8080/ delay=0s timeout=1s period=10s #success=1 #failure=3
```

---

### 1.4 - Configure Probe Properties

You can tune:
- `initialDelaySeconds`: Time to wait before first check.
- `timeoutSeconds`: Probe timeout.
- `periodSeconds`: Interval between probes.
- `failureThreshold`: Number of failures before restart.

Check probe fields:

```bash
kubectl explain pod.spec.containers.livenessProbe > explain_liveness.txt
```

---

### 1.5 - Writing Better Probes

Use a custom `/health` path that checks your internal app state.

To ensure pod restarts happen on other nodes, use a **Deployment** or **ReplicationController**.

---

## 2. 📦 Introducing Deployments (Replacement for ReplicationControllers)

A **Deployment** manages the lifecycle of replicated pods. It maintains the desired number of replicas and replaces failed pods automatically.

---

### 2.1 - How It Works

Deployments:
- Ensure a set number of pod replicas are running.
- Replace deleted or failed pods.
- Can be scaled or updated.

---

### 2.2 - Creating a Deployment

```bash
kubectl create -f deployment-go-http-lliveness.yaml
kubectl get deployments.apps
```

---

### 2.3 - Deployment in Action

```bash
kubectl get pods 
# NAME                                                 READY   STATUS    RESTARTS   AGE
# http-server-healthcheck-deployment-65dc7775f-b52fv   1/1     Running   0          2s
# http-server-healthcheck-deployment-65dc7775f-dp5mg   1/1     Running   0          2s
# http-server-healthcheck-deployment-65dc7775f-htlvd   1/1     Running   0          2s

kubectl get pods 
# NAME                                                 READY   STATUS    RESTARTS      AGE
# http-server-healthcheck-deployment-65dc7775f-b52fv   1/1     Running   1 (57s ago)   2m7s
# http-server-healthcheck-deployment-65dc7775f-dp5mg   1/1     Running   1 (57s ago)   2m7s
# http-server-healthcheck-deployment-65dc7775f-htlvd   1/1     Running   1 (57s ago)   2m7s
```

Delete a pod:

```bash
kubectl delete pod <pod-name>
```

A new pod is created to maintain replica count.

---

### 2.4 - Exclude a Pod from Deployment

```bash
kubectl label pod <pod-name> app=foo --overwrite
```

Deployment creates a new pod to restore replica count.

---

### 2.5 - Update Pod Template

```bash
kubectl edit deployment <deployment-name>
```

Changes only apply to new pods. Delete old ones to use the updated template.

---

### 2.6 - Horizontal Scaling

```bash
kubectl scale deployment <deployment-name> --replicas=10
```

---

### 2.7 - Delete Deployment

```bash
kubectl delete deployments deployment-liveness
```

---

## 3. ❌ ReplicaSet vs. ReplicationController

**ReplicationControllers** are deprecated. Use **Deployments**, which manage **ReplicaSets** internally for rolling updates, rollbacks, and more.

---
