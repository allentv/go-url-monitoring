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

## Local Development

The Go program can be started by running `make run` which will start the Go binary at port `9000`. Using `Ctrl + C` on Mac should kill the process.

To run all the tests, execute `make test`. The coverage report for various packages will be shown at the end of test execution. Potential race conditions are also tested with the `-race` flag for Go test runner.

## Docker Image

Strips the debug symbols and compress the binary to achieve 2.5 MB binary size. When used with Google distroless, the docker image size is < 5MB

## Setting up Kubernetes (Incomplete)

I am using docker-desktop to setup k8s on my mac machine.

### Steps

* Start docker-desktop on Mac with kubernetes cluster enabled

* Verify that the k8s context is set to `docker-desktop`

```bash
kubectl config set-context 'docker-desktop'
kubectl config current-context
```

* Install Prometheus and Grafana by running `kubectl apply -f` to the files in `deploy/prometheus` and `deploy/grafana`

* Install the k8s UI dashboard user and role binding

```bash
kubectl apply -f deploy/dashboard-user.yaml
kubectl apply -f deploy/dashboard-rolebinding.yaml
```

> Note: Run `make k8s-setup` to execute all of the above setups in one go

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

### References

* Setting up Prometheus and Grafana on [docker-desktop](https://github.com/scalastic/local-k8s-installation)
* [Collecting Prometheus metrics in Go](https://gabrieltanner.org/blog/collecting-prometheus-metrics-in-golang)
