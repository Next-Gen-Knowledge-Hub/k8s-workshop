# Installing Kubernetes with Minikube on Ubuntu

This guide will walk you through installing and running a local Kubernetes cluster using **Minikube** on Ubuntu.

---

## Prerequisites

Ensure your system has:

- Ubuntu 20.04+ (or compatible)
- 2 CPUs or more
- 2GB of free memory
- 20GB of disk space
- Internet connection

---

## Step 1: Install Dependencies

```bash
sudo apt update
sudo apt install -y curl wget apt-transport-https ca-certificates gnupg lsb-release
```

---

## Step 2: Install Docker (Minikube container runtime)

```bash
sudo apt install -y docker.io
sudo usermod -aG docker $USER
newgrp docker
```

Verify Docker:

```bash
docker version
```

---

## Step 3: Install kubectl (Kubernetes CLI)

```bash
curl -LO "https://dl.k8s.io/release/$(curl -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
chmod +x kubectl
sudo mv kubectl /usr/local/bin/
```

Check kubectl:

```bash
kubectl version --client
```

---

## Step 4: Install Minikube

```bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
```

Check minikube:

```bash
minikube version
```

---

## Step 5: Start Minikube

You can start Minikube with Docker as the driver:

```bash
minikube start --driver=docker
```

This will download and configure a local Kubernetes cluster using Docker.

---

## Step 6: Verify Installation

```bash
kubectl get nodes
# NAME       STATUS   ROLES           AGE    VERSION
# minikube   Ready    control-plane   101d   v1.31.0
```

You should see a single node running.

---

## Step 7: Enable Dashboard (Optional)

```bash
minikube dashboard
```

This opens a Kubernetes dashboard UI in your default browser.

---

## Step 8: Stop or Delete Cluster

- To stop the cluster:
  ```bash
  minikube stop
  ```

- To delete the cluster:
  ```bash
  minikube delete
  ```

---

## Troubleshooting

- If you face permission issues with Docker, ensure your user is added to the Docker group.
- Use `minikube logs` to inspect startup issues.
- If Kubernetes pods are stuck in `Pending`, check resources or use `minikube tunnel` for LoadBalancer services.

---

## References

- [Minikube Docs](https://minikube.sigs.k8s.io/docs/)
- [Kubernetes Docs](https://kubernetes.io/docs/home/)

---