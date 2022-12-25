SHELL := /bin/bash


.PHONY: helm-prepare
helm-prepare:
	helm repo add bitnami https://charts.bitnami.com/bitnami
	helm repo update
	helm dep update postgres/helm/
	helm dep update minio/helm/
