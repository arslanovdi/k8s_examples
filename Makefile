
# make build version=0.1.0
build:
	docker build -t stateless-todo:v$(version) -f stateless-todo/Dockerfile .
	docker tag stateless-todo:v$(version) arslanovdi/stateless-todo:v$(version)

# make push version=0.1.0
push:
	docker push arslanovdi/stateless-todo:v$(version)


port-forward:
	kubectl port-forward -n stateless-todo svc/stateless-todo-clusterip-service 5000:5000


# make build version=0.1.0
build-stateful:
	docker build -t stateful-todo:v$(version) -f stateful-todo/Dockerfile .
	docker tag stateful-todo:v$(version) arslanovdi/stateful-todo:v$(version)

# make push version=0.1.0
push-stateful:
	docker push arslanovdi/stateful-todo:v$(version)

