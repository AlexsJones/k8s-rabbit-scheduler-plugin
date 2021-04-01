.PHONY: all deploy

build:
	docker build . -t tibbar/rabbit-scheduler:v1

deploy: build
	docker push tibbar/rabbit-scheduler:v1
