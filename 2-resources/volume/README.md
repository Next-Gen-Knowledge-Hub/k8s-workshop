# 🗄️ Volumes: Attaching Disk Storage to Containers

In the previous chapters, we explored core Kubernetes resources like **Pods**, **ReplicaSets**, **DaemonSets**, **Jobs**, and **Services**. In this section, we dive **back inside the pod** to see how containers can interact with **external disk storage** or **share storage among themselves** using **Volumes**.

---

## 1 - Introducing Volumes

Volumes in Kubernetes allow containers in a pod to **read and write data** to shared or persistent storage. Unlike standalone Kubernetes objects like Pods or Services, volumes are always defined **within a Pod's specification**—they **cannot exist independently**.

---

## 1.1 - Volumes Explained with an Example

A common use case involves **multiple containers sharing data**:

> Imagine a web server container that writes logs to `/var/log`, and a sidecar container that reads and processes those logs—like a log collector or cleaner.

This scenario requires **shared storage**, and Kubernetes volumes make it possible.

---

## 1.2 - Available Volume Types

Kubernetes supports a wide array of volume types, suitable for both **ephemeral** and **persistent** storage needs.

| Volume Type             | Description |
|-------------------------|-------------|
| `emptyDir`              | Temporary storage, created when a pod starts. Deleted when the pod is removed. |
| `hostPath`              | Mounts a host node directory into the pod. Useful for debugging and node-specific configs. |
| `persistentVolumeClaim` | Connects to a PersistentVolume, often used with dynamic storage provisioning. |
| `configMap`             | Mounts config data as files or environment variables. |
| `secret`                | Used to securely inject passwords, tokens, or keys. |
| `projected`             | Combines multiple sources like `configMap` and `secret` into one volume. |
| `downwardAPI`           | Exposes pod metadata (e.g., labels, annotations) to the container. |
| `nfs`, `cephFS`, `glusterfs` | Distributed network file systems; require external setup. |
| `gcePersistentDisk`, `awsElasticBlockStore`, `azureDisk`, `azureFile` | Cloud provider-specific persistent storage. |
| `local`                 | Mounts a local disk; not suitable for multi-node portability. |
| `csi`                   | Supports pluggable third-party storage drivers via the **Container Storage Interface**. |
| `flexVolume`            | Deprecated in favor of CSI, but still in use for some vendor plugins. |
| `iscsi`, `cinder`, `vsphereVolume`, `storageOS` | Enterprise and cloud-specific storage integrations. |

> 🔍 For most learning and development environments, `emptyDir` and `hostPath` are easiest to start with.

---

## 2 - Using Volumes to Share Data Between Containers

While volumes are useful for a **single container**, they really shine when **multiple containers in a pod** need to **share state** or **communicate via files**.

---

### 2.1 - Using an `emptyDir` Volume

`emptyDir` is a **temporary volume** created when a pod is assigned to a node. It’s:

- **Empty initially**
- **Deleted** when the pod terminates
- **Shared** between all containers in the pod

#### Example: Web Server & Log Collector

Let’s walk through an example where a web server writes logs to a shared volume, and a sidecar container collects them.

```sh
# Create the deployment with two containers and an emptyDir volume
kubectl create -f sample/k8s_deployment.yaml

# Expose the deployment as a service
kubectl create -f sample/k8s_service.yaml
```

#### Verify Pod and Service

```sh
kubectl get pods
```

Output:
```
NAME                                    READY   STATUS    RESTARTS   AGE
log-mount-deployment-7c464cbc94-62djq   2/2     Running   0          3m29s
```

```sh
kubectl get services
```

Output:
```
NAME                TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
log-mount-service   NodePort    10.98.14.198   <none>        8888:32407/TCP   5s
```

```sh
minikube service list
```

Output:
```
| default     | log-mount-service | 8888 | http://192.168.49.2:32407 |
```

#### Test the Application

```sh
curl http://192.168.49.2:32407/
```

This sends a request to the web server which writes to the shared volume.

#### View Logs of Both Containers

```sh
# Server container
kubectl logs pods/log-mount-deployment-7c464cbc94-62djq

# Collector container
kubectl logs pods/log-mount-deployment-7c464cbc94-62djq log-mount-collector
```

Example Output from Collector:
```
2024/12/14 07:44:23 INFO log file path=/tmp/
2024/12/14 07:46:05 INFO file modified path=/tmp/server.txt op=WRITE
```

---

## 📌 Summary

- Kubernetes **Volumes** allow **data sharing** between containers or **persistent storage** across pod restarts.
- The simplest volume type, **`emptyDir`**, is great for intra-pod communication.
- Many **advanced volume types** (like `PVC`, `CSI`, or cloud-specific options) enable long-term, scalable storage.
- Sidecar containers are a common pattern for volume-based log collection, backup, or data processing.

---

## ✅ Bonus Tips

- Use **`volumeMounts`** to define mount paths for each container.
- Always clean up test volumes in dev clusters to avoid node disk bloat.
- For persistent storage, explore **PersistentVolume (PV)** and **PersistentVolumeClaim (PVC)**.

---
