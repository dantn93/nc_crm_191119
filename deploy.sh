#!/bin/sh
# kubectl delete service tranngocdan-nc-crm-service
# kubectl delete deployment tranngocdan-nc-crm
kubectl create -f provision/k8s/deployment.yaml
kubectl get service tranngocdan-nc-crm-service

# minikube service tranngocdan-nc-crm-service --url