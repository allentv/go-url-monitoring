# go-url-monitoring

## Docker Image

Strips the debug symbols and compress the binary to achieve 2.5 MB binary size. When used with Google distroless, the docker image size is < 5MB

## Setting up Kubernetes

I am using minikube to setup k8s on my mac machine.

### Steps

* Start minikube with the latest version of k8s

```bash
minikube start --kubernetes-version=v1.23.3
```

* Install Prometheus and Grafana

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm install prometheus prometheus-community/prometheus
kubectl expose service prometheus-server --type=NodePort --target-port=9090 --name=prometheus-server-np

helm repo add grafana https://grafana.github.io/helm-charts
helm install grafana stable/grafana
kubectl expose service grafana --type=NodePort --target-port=3000 --name=grafana-np
```
