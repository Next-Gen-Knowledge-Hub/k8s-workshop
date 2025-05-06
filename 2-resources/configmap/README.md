---

# ⚙️ Configuring Containerized Applications in Kubernetes

In Kubernetes, configuration flexibility is essential for adapting applications across **development, staging, and production environments**. Kubernetes provides several mechanisms for passing configuration data to containers:

* ✅ Command-line arguments
* ✅ Environment variables
* ✅ Mounted configuration files using `ConfigMap` or `Secret` volumes

---

## 1 - Overview of Configuration Techniques

Applications can receive configuration through:

* **Command-line arguments** (passed to entrypoints)
* **Environment variables** (e.g., `MYSQL_ROOT_PASSWORD`)
* **Mounted config files** via special volume types (`ConfigMap`, `Secret`, etc.)

---

## 2 - 🔧 Passing Command-Line Arguments to Containers

### 2.1 - Define Commands & Args in Docker

In Docker, the `CMD` and `ENTRYPOINT` instructions define the container's runtime behavior. In Kubernetes, these can be **overridden** in the Pod spec using `command` and `args`.

### 2.2 - Override Commands in Kubernetes

You can override the default container command/args in your deployment YAML:

```yaml
spec:
  containers:
    - name: arg-server
      image: your-image
      command: ["./arg-server"]
      args: ["arg1", "arg2", "arg3", "k8s-rocks"]
```

```bash
# Deploy the pod
kubectl create -f /args/k8s-deployment.yaml

# Check logs
kubectl logs pods/arg-server-5b8cbff556-qtc9f
# Output
2024/12/14 14:16:20 [./arg-server arg1 arg2 arg3 k8s-rocks]
```

---

## 3 - 🌱 Setting Environment Variables

### 3.1 - Define Environment Variables in Pod Spec

You can directly define ENV variables in the pod spec using the `env` field:

```yaml
env:
  - name: PORT
    value: "1818"
  - name: STAGE
    value: "k8s-deployment"
```

```bash
kubectl create -f /env/k8s-deployment.yaml
kubectl logs env-server-b98d7dc86-blqb8
# Output
2024/12/15 07:02:02 INFO server is running port=1818 stage=k8s-deployment
```

### 3.2 - Referencing Other ENV Variables

You can reference existing ENV variables:

```yaml
env:
  - name: FIRST_ENV
    value: "foo"
  - name: SECOND_ENV
    value: "$(FIRST_ENV)_bar"  # Interpolates to "foo_bar"
```

> ⚠️ Note: In practice, variable substitution like `$(FIRST_ENV)` **only works for downward API references**, not internal ENV chaining in all runtime environments.

---

### 3.3 - Drawbacks of Hardcoded ENV

Hardcoding config in deployment YAMLs causes:

* ❌ Duplication across environments
* ❌ Inflexibility when switching stages
* ✅ Solution: Use **ConfigMaps** to decouple environment-specific settings

---

## 4 - 🧩 Using ConfigMaps for Configuration

### 4.1 - What is a ConfigMap?

A `ConfigMap` is a **Kubernetes object** used to inject configuration data into pods as:

* ENV variables
* Command-line args
* Mounted files

> 📦 It stores **key-value pairs** and can include entire config files.

---

### 4.2 - Create a ConfigMap

```yaml
# config_map/k8s-config-dev.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: env-server-config
data:
  PORT: "1818"
  STAGE: "k8s-dev"
```

```bash
kubectl create -f config_map/k8s-config-dev.yaml
kubectl get configmaps
```

---

### 4.3 - Use ConfigMap as ENV in Deployment

```yaml
env:
  - name: PORT
    valueFrom:
      configMapKeyRef:
        name: env-server-config
        key: PORT
  - name: STAGE
    valueFrom:
      configMapKeyRef:
        name: env-server-config
        key: STAGE
```

```bash
# Apply deployment and check output
kubectl create -f config_map/k8s-deployment.yaml
kubectl logs <pod-name>
# Output
INFO server is running port=1818 stage=k8s-dev
```

> To switch environments:

```bash
kubectl apply -f config_map/k8s-config-pro.yaml
kubectl delete pod <old-pod-name>
```

```yaml
# New log after environment change
INFO server is running port=1819 stage=k8s-pro
```

✅ You can now reuse the **same Deployment YAML** across environments by just switching ConfigMaps.

---

### 4.4 - Mounting ConfigMaps as Files

Some apps (e.g., Redis, MongoDB) require full config files instead of ENV vars.

```yaml
volumes:
  - name: config-volume
    configMap:
      name: redis-config
containers:
  - name: redis
    volumeMounts:
      - name: config-volume
        mountPath: /etc/redis/
        readOnly: true
```

```bash
kubectl create -f config_map/k8s-config-redis.yaml
kubectl create -f config_map/k8s-deployment-redis.yaml
```

---

## ✅ Summary

| Method                  | Use Case                                      |
| ----------------------- | --------------------------------------------- |
| `command` & `args`      | Override container startup behavior           |
| `env` variables         | Basic config, quick setup                     |
| `valueFrom` + ConfigMap | Environment decoupling, scalable configs      |
| ConfigMap volumes       | Mounting config files (e.g., YAML, JSON, ini) |

---

## 📘 Additional Tips

* Use **Secrets** instead of ConfigMaps for sensitive values like API keys.
* Use **Downward API** to pass pod metadata (e.g., name, namespace) to containers.
* Combine ConfigMap with `--prefix` if passing config as args:
  Example: `--config.port=$(PORT)` from ENV or ConfigMap.

---
