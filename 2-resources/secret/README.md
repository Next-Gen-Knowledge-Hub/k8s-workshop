---

# **1. Using Secrets to Pass Sensitive Data to Containers**

While configuration options like environment variables and ConfigMaps are sufficient for non-sensitive data, **Kubernetes Secrets** are essential when handling sensitive data such as passwords, API tokens, and certificates.

---

## **2. Introducing Kubernetes Secrets**

Kubernetes Secrets are key-value pairs similar to ConfigMaps but designed to store **sensitive data securely**.

### 🔒 Key Security Benefits:

* **Node distribution**: Only available on nodes where pods that use them are scheduled.
* **In-memory storage**: Stored in memory, not on disk.
* **Base64 encoding**: Data is Base64-encoded.
* **Can be encrypted at rest**: With encryption providers configured.
* **RBAC protected**: Access is more tightly controlled than ConfigMaps.

---

## **3. Best Practices**

| Use a `ConfigMap` when...                                           | Use a `Secret` when...                                                |
| ------------------------------------------------------------------- | --------------------------------------------------------------------- |
| Storing non-sensitive data like environment configs or UI settings. | Storing sensitive data like passwords, tokens, keys, or certificates. |
| Data can be stored in plain text.                                   | Data must be encoded in base64 and optionally encrypted.              |
| Configuration is not security-critical.                             | Data access must be tightly controlled via RBAC.                      |

---

## **4. Working with Secrets**

### **4.1 Viewing the Default Service Account Token Secret**

Every pod automatically receives a token stored in a secret (used to authenticate to the Kubernetes API):

```bash
kubectl get secrets
kubectl describe secret <default-token-name>
kubectl explain secret
```

---

### **4.2 Creating a Custom Secret**

#### **Option 1: From a File**

```bash
kubectl create secret generic env-secret \
  --from-literal=APP_SECRET_STAGE=production-secret \
  --from-literal=DB_PASSWORD=supersecret
```

#### **Option 2: Using a YAML File**

📄 **`secret/k8s-secret.yaml`**

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: env-secret
type: Opaque
data:
  APP_SECRET_STAGE: cHJvZHVjdGlvbi1zZWNyZXQ=  # base64 for "production-secret"
  DB_PASSWORD: c3VwZXJzZWNyZXQ=              # base64 for "supersecret"
```

📄 **`secret/k8s-config.yaml`**

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: env-config-map
data:
  APP_STAGE: k8s-secret
  PORT: "1818"
```

📄 **`secret/k8s-deployment.yaml`**

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: env-server
spec:
  replicas: 1
  selector:
    matchLabels:
      app: env-server
  template:
    metadata:
      labels:
        app: env-server
    spec:
      containers:
      - name: server
        image: my-app-image
        env:
        - name: APP_SECRET_STAGE
          valueFrom:
            secretKeyRef:
              name: env-secret
              key: APP_SECRET_STAGE
        - name: DB_PASSWORD
          valueFrom:
            secretKeyRef:
              name: env-secret
              key: DB_PASSWORD
        - name: PORT
          valueFrom:
            configMapKeyRef:
              name: env-config-map
              key: PORT
        ports:
        - containerPort: 1818
```

Then run:

```bash
kubectl apply -f secret/k8s-config.yaml
kubectl apply -f secret/k8s-secret.yaml
kubectl apply -f secret/k8s-deployment.yaml
```

### ✅ Check deployment:

```bash
kubectl get pods
kubectl logs <pod-name>
kubectl describe secret env-secret
```

---

### **4.3 Using Secrets as Volumes**

You can also mount Secrets as files into a container:

```yaml
        volumeMounts:
        - name: secret-volume
          mountPath: "/etc/secret"
          readOnly: true
      volumes:
      - name: secret-volume
        secret:
          secretName: env-secret
```

This makes the secret available as files under `/etc/secret/APP_SECRET_STAGE`, etc.

---

## **5. Comparing ConfigMaps vs Secrets**

| Feature               | **Secrets**                                    | **ConfigMaps**                                   |
| --------------------- | ---------------------------------------------- | ------------------------------------------------ |
| **Purpose**           | Store sensitive data (passwords, tokens, keys) | Store non-sensitive configuration (env settings) |
| **Data Format**       | Base64-encoded                                 | Plain-text                                       |
| **Encryption**        | Can be encrypted at rest                       | No encryption support                            |
| **Access Control**    | Stricter RBAC                                  | Standard RBAC                                    |
| **Use Cases**         | Secrets, credentials, tokens                   | App configuration, toggles, environment info     |
| **Injection Methods** | Env vars or mounted files                      | Env vars or mounted files                        |
| **Persistence**       | In-memory only on required nodes               | Stored in etcd like other config objects         |

---

## **6. Extra Notes**

* 🔐 **Avoid storing secrets in Git** — even in base64!
* 🔄 **Rotate secrets regularly**, especially if used for external service credentials.
* 🔍 Use tools like **Sealed Secrets**, **HashiCorp Vault**, or **external secret controllers** for better secret management in production.

---
