### How to Run

1. Run Docker

2. Run minikube with ``minikube start``, minikube is a small cluster kubernetes with one node. Docker will pull image 
minikube to GCR, and run that image on docker (since minikube run with driver docker on this computer)

```😄  minikube v1.25.2 on Darwin 10.14.6
   ✨  Using the docker driver based on existing profile
   👍  Starting control plane node minikube in cluster minikube
   🚜  Pulling base image ...
   🏃  Updating the running docker "minikube" container ...
   🐳  Preparing Kubernetes v1.23.3 on Docker 20.10.12 ...
       ▪ kubelet.housekeeping-interval=5m
   🔎  Verifying Kubernetes components...
       ▪ Using image gcr.io/k8s-minikube/storage-provisioner:v5
       ▪ Using image kubernetesui/metrics-scraper:v1.0.7
       ▪ Using image kubernetesui/dashboard:v2.3.1
   🌟  Enabled addons: storage-provisioner, default-storageclass, dashboard
   🏄  Done! kubectl is now configured to use "minikube" cluster and "default" namespace by default
```

3. Run ``minikube dashboard`` or ``minikube dashboard --url`` to open dashboard minikube on web.


### How Apps Run

- A Service routes traffic across a set of Pods. Services are the abstraction that allow pods to die and replicate in Kubernetes without impacting your application.
- NodePort - Exposes the Service on the same port of each selected Node in the cluster using NAT. Makes a Service accessible from outside the cluster using <NodeIP>:<NodePort>. Superset of ClusterIP.
https://kubernetes.io/docs/tutorials/kubernetes-basics/expose/expose-intro/
- Since our cluster run on docker (docker as VM), we can acces our apps with command
```
docker exec -it {$CONTAINER_MINIKUBE_ID} curl $(minikube ip):{$NODE_PORT}
```
