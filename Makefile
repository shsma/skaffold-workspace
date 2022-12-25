SHELL := /bin/bash

.PHONY: skaffold-start
skaffold-dev:
	skaffold dev --port-forward

.PHONY: minikube-start
minikube-start:
	minikube start $(ADDONS_FLAG) --cpus 3 --memory 10GiB --wait all
	minikube update-context

.PHONY: minikube-delete
minikube-delete:
	minikube delete
