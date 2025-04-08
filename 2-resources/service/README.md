# 📡 Kubernetes Services: Internal and External Connectivity

## 🧩 What Are Services?

A **Kubernetes Service** is a resource you create to provide a single, constant point of entry to a group of pods that offer the same functionality. Each service is assigned a stable **IP address** and **port** that remain unchanged for the service's lifetime.
Kubernetes **Services** provide a stable, abstract interface to access a group of ephemeral **Pods**. They solve three core challenges:

- Pods have dynamic IPs and can come/go anytime.
- Services abstract away pod IP changes.
- Clients don’t need to know how many pods exist or their IPs.

---

## 🔁 Accessing Services **Inside the Cluster**

### ✅ Automatic Service Discovery
Pods can discover and connect to services via DNS:
```bash
curl http://<service-name>.<namespace>:<port>/<endpoint>
```

### 🧭 Session Affinity
Use `sessionAffinity: ClientIP` in the Service spec to ensure the same client always hits the same pod.

---

## 🌐 Accessing **External Services** via Kubernetes Service Object

### 🔁 Create a Service Without a Selector

```yaml
apiVersion: v1
kind: Service
metadata:
  name: external-service
spec:
  type: ClusterIP
  ports:
    - port: 9999
      targetPort: 9999
```

Then define matching `Endpoints` manually:

```yaml
apiVersion: v1
kind: Endpoints
metadata:
  name: external-service
subsets:
  - addresses:
      - ip: 192.168.1.106
    ports:
      - port: 9999
```

Client Pods can now connect like:
```bash
curl http://external-service.default:9999/hello
```

---

## 📤 Exposing Services **to External Clients**

### 🚪 1. NodePort
- Opens a fixed port on every node.
- External users access the service via `NodeIP:NodePort`.

```yaml
type: NodePort
```

---

### 🌐 2. LoadBalancer
- Ideal in cloud environments.
- Creates an external load balancer that forwards to NodePorts.

```yaml
type: LoadBalancer
```

> **Note:** In local clusters like Minikube, you need to run `minikube tunnel` to simulate this behavior.

---

### 🌍 3. Ingress (Part 3)
- HTTP/HTTPS-based router.
- Routes traffic to different services based on host/path.
- Ideal for exposing multiple services via one IP/domain.

---

## ⚙️ Optimizing External Connections

When external clients connect via NodePort or LoadBalancer, traffic might hop between nodes. To avoid this:

```yaml
externalTrafficPolicy: Local
```

- Ensures traffic is only routed to pods on the node that received it.
- Preserves original client IP.

---

Sure! Here's a clean and well-structured **Markdown (MD)** version of the section you provided:

---

## 1.1 - Creating Services

A service can be backed by multiple pods. Incoming connections to the service are **load-balanced** across all available pods.

To define which pods belong to a service, Kubernetes uses **label selectors**.

### Step-by-Step: Create a Deployment and Expose It via a Service

1. **Create the Deployment** using a YAML file:

   ```sh
   kubectl create -f go-http-deployment.yaml
   ```

2. **Check Deployment Labels**:

   ```sh
   kubectl get deployments --show-labels
   ```

   Output:
   ```
   NAME                 READY   UP-TO-DATE   AVAILABLE   AGE   LABELS
   go-http-deployment   3/3     3            3           25s   app=go-http-deployment-kind,creation_method=deployment,env=local
   ```

3. **Check Pod Labels**:

   ```sh
   kubectl get pods --show-labels
   ```

   Output:
   ```
   NAME                                 READY   STATUS    RESTARTS   AGE   LABELS
   go-http-deployment-78fd5b857-4xlzp   1/1     Running   0          39s   app=go-http-pod, pod-template-hash=78fd5b857
   ...
   ```

4. **Create the Service** using its YAML file:

   ```sh
   kubectl create -f go-http-service.yaml
   ```

5. **Verify the Service**:

   ```sh
   kubectl get service --show-labels
   ```

   Example Output:
   ```
   NAME              TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE   LABELS
   go-http-service   LoadBalancer   10.96.155.162   <pending>     9999:30999/TCP   12s   app=go-http-service-kind,...
   ```

6. **View Available Services via Minikube**:

   ```sh
   minikube service list
   ```

   Output:
   ```
   | default     | go-http-service |         9999 | http://192.168.49.2:30999 |
   ```

7. **Test the Service**:

   ```sh
   curl 192.168.49.2:30999/hello
   ```

   Sample Response:
   ```
   request received on host go-http-deployment-78fd5b857-4xlzp !
   ```

---

### Example: Accessing a Service from Another Pod

1. **Get the Running Pods**:

   ```sh
   kubectl get pods --show-labels
   ```

2. **Access a Pod Shell**:

   ```sh
   kubectl exec -it go-http-deployment-7d7b7685b9-h74ss -- bash
   ```

3. **Use curl to Contact the Service via DNS**:

   ```sh
   curl http://go-http-service.default:30999/hello
   ```

   Example Output:
   ```
   request received on host go-http-deployment-7d7b7685b9-wj4fv !
   ```

---

Let me know if you'd like this formatted into a full tutorial or documentation guide!

## ✅ Summary

| Scenario                              | Solution                                             |
|---------------------------------------|------------------------------------------------------|
| Pod -> Pod                            | ClusterIP Service + DNS (`svc.namespace`)           |
| Pod -> Internet API                   | Allow egress + NAT gateway                          |
| Pod -> External service inside LAN    | ClusterIP Service + manual Endpoints                |
| Client -> Cluster Service (public)    | NodePort / LoadBalancer / Ingress                   |
| Reduce network hops (external traffic)| `externalTrafficPolicy: Local`                      |

---

Let me know if you want visual diagrams or YAML examples to go deeper into any part!