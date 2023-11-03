PASSWORD ?= password
REDIS_CHANNEL ?= basic
DOCKER_TAG ?= latest
DOCKER_REGISTRY ?= emanuelfeijo
REDIS_HOST ?= my-redis

# DOCKER STUFF

build-dockerfile-sender:
	docker build -t  emanuelfeijo/redis-pub-sender:${DOCKER_TAG} ./sender 
	docker push ${DOCKER_REGISTRY}/redis-pub-sender

build-dockerfile-receiver:
	docker build -t emanuelfeijo/redis-pub-receiver:${DOCKER_TAG} ./receiver 
	docker push ${DOCKER_REGISTRY}/redis-pub-receiver

build-dockers-locally: build-dockerfile-sender-locally build-dockerfile-receiver-locally

build-dockerfile-sender-locally:
	docker build -t  emanuelfeijo/redis-pub-sender:${DOCKER_TAG} ./sender

build-dockerfile-receiver-locally:
	docker build -t emanuelfeijo/redis-pub-receiver:${DOCKER_TAG} ./receiver

build-dockers: build-dockerfile-receiver build-dockerfile-sender


minikube-local:
	minikube start
	eval $(minikube -p minikube docker-env) 
	docker context use default
	minikube cache reload

# Deployments

k8s:
	kubectl apply -f $(shell pwd)/receiver/deployment.yaml
	kubectl apply -f $(shell pwd)/sender/deployment.yaml

redis:
	helm repo add bitnami https://charts.bitnami.com/bitnami
	helm install ${REDIS_HOST} bitnami/redis --version 18.1.6 --set auth.password=${PASSWORD}


# clean up stuff
delete-redis:
	helm delete ${REDIS_HOST}  


delete-k8s-deps: 
	kubectl delete services sender
	kubectl delete deployment sender
	kubectl delete services receiver
	kubectl delete deployment receiver
	docker image rm -f emanuelfeijo/redis-pub-receiver emanuelfeijo/redis-pub-sender
	docker images prune

delete-all: delete-redis delete-k8s-deps 
	minikube stop