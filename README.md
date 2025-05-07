# Kubernetes Workshop
This workshop is around topics about
kubernetes resource types, and a
hand on experience with kubernetes.

Repository contains following topics

1. Setup Kubernetes with minikube
    - 1.1 [What's Kubernetes ?](./1-setup/OVER_VIEW.md)
    - 1.2 [Install `docker`](./1-setup/install_docker.sh)
    - 1.3 [Install `minikube`, `kubectl`](./1-setup/README.md)
    - 1.4 Connect and create a single [`pod`](1-setup/simple_server) resource type

2. Kubernetes Resource Types (Kind)
    - 2.1 [Pod](./2-resources/pod/README.md)
    - 2.2 [Deployments And Replication Set](./2-resources/deployment/README.md)
    - 2.3 [Service](./2-resources/service/README.md)
    - 2.4 [Volume](./2-resources/volume/README.md)
    - 2.5 [Config Map](./2-resources/configmap/README.md)
    - 2.6 [Secret](./2-resources/secret/README.md)
    - 2.7 Stateful Set
    - 2.8 Daemon Set, Job, Cronjob
    - 2.9 Internals

3. Securing Kubernetes Cluster
    - 3.1 Ingress
    - 3.2 Egress
    - 3.3 Securing Api Server (Service Accounts)
    - 3.4 Securing Cluster Nodes (Network Policies)

4. Automation
    - 4.1 Managing Pods Computation Resources
    - 4.2 Autoscaling (HPA)
    - 4.3 Advance Scheduling
    - 4.4 Application Development

5. Extending Kubernetes
    - 5.1 Write Custom Resource Types
    - 5.2 Write a Basic Controller

Feel free to use and make any change ;)

