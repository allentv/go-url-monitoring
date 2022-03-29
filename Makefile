run:
	go run cmd/app/main.go

test:
	go test ./... -count=1 -race -cover

build:
	docker build -t go-url-monitoring .

k8s-setup: build
	@echo "\n[Deploying k8s dashboard]..."
	kubectl apply -f deploy/dashboard-user.yaml
	kubectl apply -f deploy/dashboard-rolebinding.yaml

	@echo "\n[Deploying application]..."
	kubectl apply -f deploy/app/deployment.yaml
	kubectl apply -f deploy/app/service.yaml

	@echo "\Setting up 'monitoring' namespace..."
	kubectl apply -f deploy/monitoring-namespace.yaml

	@echo "\n[Deploying prometheus]..."
	kubectl apply -f deploy/prometheus/deployment.yaml
	kubectl apply -f deploy/prometheus/role.yaml
	kubectl apply -f deploy/prometheus/configmap.yaml
	kubectl apply -f deploy/prometheus/service.yaml

	@echo "\n[Deploying grafana]..."
	kubectl apply -f deploy/grafana/deployment.yaml
	kubectl apply -f deploy/grafana/configmap.yaml
	kubectl apply -f deploy/grafana/service.yaml

k8s-dashboard:
	kubectl get secret --namespace default grafana -o jsonpath="{.data.admin-password}" | base64 --decode | pbcopy
	open http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/
	kubectl proxy

remove-app:
	kubectl delete deploy url-monitor
	kubectl delete svc url-monitor-service
