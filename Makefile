run:
	go run main.go

test:
	go test ./... -count=1 -race -cover

build:
	docker build -t go-url-monitoring .

k8s-setup:
	kubectl apply -f deploy/dashboard-user.yaml
	kubectl apply -f deploy/dashboard-rolebinding.yaml
	kubectl apply -f deploy/application.yaml

k8s-dashboard:
	kubectl get secret --namespace default grafana -o jsonpath="{.data.admin-password}" | base64 --decode | pbcopy
	open http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/
	kubectl proxy

remove-app:
	kubectl delete deploy url-monitor
