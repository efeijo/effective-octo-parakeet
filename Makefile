PASSWORD ?= password
REDIS_CHANNEL ?= basic
DOCKER_TAG ?= latest
DOCKER_REGISTRY ?= emanuelfeijo

redis:
	helm repo add bitnami https://charts.bitnami.com/bitnami
	helm install my-redis bitnami/redis --version 18.1.6 --set auth.password=${PASSWORD}

build-dockerfile-sender:
	docker build -t  emanuelfeijo/redis-pub-sender:${DOCKER_TAG} ./sender --build-arg REDIS_CHANNEl=${REDIS_CHANNEL} --build-arg REDIS_PASSWORD=${PASSWORD} 
	docker push ${DOCKER_REGISTRY}/redis-pub-sender

build-dockerfile-receiver:
	docker build -t emanuelfeijo/redis-pub-receiver:${DOCKER_TAG} ./receiver --build-arg REDIS_CHANNEl=${REDIS_CHANNEL} --build-arg REDIS_PASSWORD=${PASSWORD}
	docker push ${DOCKER_REGISTRY}/redis-pub-receiver

build-dockers: build-dockerfile-receiver build-dockerfile-sender


delete-redis:
	helm delete my-redis 

k8s:
	
	kubectl apply -f $(shell pwd)/receiver/deployment.yaml
	kubectl apply -f $(shell pwd)/sender/deployment.yaml

destroy-k8s: delete-redis
	kubectl delete services sender
	kubectl delete deployment sender
	kubectl delete services receiver
	kubectl delete deployment receiver

