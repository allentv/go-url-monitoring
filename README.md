# go-url-monitoring

## Package Structure

* `cmd/` : Executables
* `pkg/` : Source Code
* `deploy/`: Kubernetes YAML files
* `Makefile`: Commands for all tasks
* `Dockerfile`: For creating the application docker image

## Architecture

The application runs a HTTP server that has 2 `GET` endpoints:

* `/metrics`: list Prometheus metrics
* `/routes` : list of routes that are monitored

During application startup:

* Metrics are created and registered with Prometheus library
* The monitoring process is started
* The routes are read
* Server config is created
* HTTP server is started

The monitoring is done by allocating a separate go routine for each URL to be monitored inside which a ticker fires on a default interval causing the process to check the URL uptime, time taken in milliseconds and then update metrics.

The prometheus go client is used which exposes the metrics data in the format that can be scraped and ingested by Prometheus.

## Docker Image

Strips the debug symbols and compress the binary to achieve 2.5 MB binary size. When used with Google distroless, the docker image size is < 5MB

## Setting up Kubernetes

I am using docker-desktop to setup k8s on my mac machine.

### Steps

* Start docker-desktop on Mac with kubernetes cluster enabled

* Install Prometheus and Grafana

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm install prometheus prometheus-community/prometheus
kubectl expose service prometheus-server --type=NodePort --target-port=9090 --name=prometheus-server-np

helm repo add grafana https://grafana.github.io/helm-charts
helm install grafana grafana/grafana
kubectl expose service grafana --type=NodePort --target-port=3000 --name=grafana-np
```

* Install the k8s UI dashboard user and role binding

```bash
kubectl apply -f deploy/dashboard-user.yaml
kubectl apply -f deploy/dashboard-rolebinding.yaml
```

* Get Admin token

```bash
kubectl get secret --namespace default grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
```

* Start the proxy in another terminal

```bash
kubectl proxy
```

* Navigate to the url: <http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/> and enter the admin token from before

* The K8s UI dashboard should now be visible

Follow this tutorial: <https://github.com/scalastic/local-k8s-installation> for setting up Prometheus and Grafana on docker-desktop

// TODO

* Update README
* Add unit tests
* Setup deployment of the app
* Setup Prom metrics and Grafana dashboard
* Add screenshot to README
* Add architecture of the app to README
* Make the port: 9000 configurable via Docker build arg and env variable
