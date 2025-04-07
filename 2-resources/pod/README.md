---

# 📦 Pods: Running Containers in Kubernetes

We’ll start with **Pods**, the central and most important concept in Kubernetes. Everything else either manages, exposes, or is used by Pods.

---

## 1 - Introducing Pods

### 1.1 - Why We Need Pods
- In Kubernetes, each process should run in its **own container**.
- To manage closely related containers together, we need a **higher-level construct**: the Pod.

### 1.2 - What is a Pod?
- A **Pod** allows multiple containers to run together and share:
    - **Network**
    - **Storage**
    - **Linux namespaces**
- Pods act like lightweight VMs that host single or tightly-coupled apps.

### 1.3 - Container Organization
- Treat each **Pod like a separate machine** running one app.
- Avoid putting multiple unrelated apps in one Pod.

---

## 2 - Creating Pods from YAML or JSON Descriptors

### 2.1 - Imperative Approach (Quick)
```bash
kubectl run go-http --image=manouchehrrasouli96/golang-http-server:0.0.1 --port=9999
kubectl expose pod go-http --type NodePort --port=9999 --target-port=9999
kubectl get services
minikube service list
curl http://192.168.49.2:30148/hello
```

### 2.2 - Declarative Approach (Recommended)
```bash
kubectl get pod go-http -o yaml > go-http-out.yml
```

YAML structure:
- `metadata`: Name, labels, etc.
- `spec`: Pod configuration (containers, ports).
- `status`: Runtime info (state, IP, etc).

### 2.3 - Creating a YAML Manifest
```bash
kubectl explain pods
kubectl explain pods.spec
kubectl explain deployment.spec.template.spec.containers
```

Apply YAML:
```bash
kubectl create -f go-http-manual.yaml
kubectl get pods
kubectl logs go-http-from-yml
kubectl port-forward go-http-from-yml 9999:9999
```

---

## 3 - Organizing Pods with Labels

### 3.1 - What are Labels?
- Key-value pairs to **group** and **select** resources.
- Used in **selectors**, not visible to users.

### 3.2 - Add Labels to Pod YAML
```yaml
metadata:
  labels:
    creation_method: manual
    env: local
```

Check labels:
```bash
kubectl get pods --show-labels
kubectl get pod -L creation_method,env
kubectl get pod -l env=local
```

### 3.3 - Modify Labels
```bash
kubectl label pod <pod-name> new-label=value
kubectl label pod <pod-name> existing-label=new-value --overwrite
```

---

## 4 - Listing Pods with Label Selectors

### 4.1 - Using a Single Selector
```bash
kubectl get pod -l env=local
```

### 4.2 - Multiple Conditions
```bash
kubectl get pod -l env -L env,creation_method
kubectl get pod -l env,creation_method=manual
```

---

## 5 - Pod Scheduling Using Node Labels

### 5.1 - Labeling Nodes
```bash
kubectl label nodes minikube gpu=true
kubectl get nodes -l gpu=true
kubectl get nodes -L gpu,cpu
```

### 5.2 - Schedule Pods Using nodeSelector
In your YAML (`go-http-node-selector-manual.yaml`):
```yaml
spec:
  nodeSelector:
    gpu: "true"
```

### 5.3 - Caution
Avoid hardcoding specific node hostnames;
use logical labels instead.

---

## 6 - Annotating Pods

### 6.1 - Add Annotation
```bash
kubectl annotate pod <pod-name> annotation-key=value
```

### 6.2 - View Annotations
```bash
kubectl describe pod <pod-name>
```

---

## 7 - Using Namespaces to Group Resources

### 7.1 - Why Use Namespaces?
- Group resources (e.g., dev, test, prod).
- Avoid naming conflicts.
- Organize multi-tenant environments.

### 7.2 - View Namespaces
```bash
kubectl get namespaces
kubectl get pods --namespace default
```

### 7.3 - Create Namespace
```bash
kubectl explain namespace
kubectl create -f namespace.yaml
kubectl create namespace my-namespace
```

### 7.4 - Managing Resources in Other Namespaces
```bash
kubectl create -f pod-spec.yaml --namespace my-namespace
```

### 7.5 - Limits of Namespace Isolation
- Namespaces isolate resources but **do not** isolate network or CPU usage.

---

## 8 - Stopping and Removing Pods

### 8.1 - Delete Pod by Name
```bash
kubectl delete pod go-http-from-yml
kubectl delete service go-http-from-yml
```

### 8.2 - Delete by Label
```bash
kubectl delete pod -l env=local
```

### 8.3 - Delete Namespace
```bash
kubectl delete namespace my-namespace
```

### 8.4 - Delete All Pods in a Namespace
```bash
kubectl delete pod --all --namespace my-namespace
```

### 8.5 - Delete All Resources in a Namespace
```bash
kubectl delete all --all --namespace my-namespace
```

---

🔗 **Reference**: [Kubernetes Official Docs](https://kubernetes.io/docs/reference/)
