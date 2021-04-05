.PHONY: all deploy

build:
	docker build . -t tibbar/rabbit-scheduler:v1

publish: build
	docker push docker.io/tibbar/rabbit-scheduler:v1

deploy:
	kubectl apply -f deploy/
token:
	 kubectl get serviceaccount rabbit-scheduler -n kube-system -o=jsonpath='{.secrets[0].name}' | xargs kubectl get secret -ojsonpath='{.data.token}' -n kube-system | base64 --decode
