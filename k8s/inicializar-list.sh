#!/bin/bash

set -e  # Detener el script si ocurre algún error

echo "📁 Desplegando microservicio p-go-list..."

# Namespace
kubectl apply -f k8s/list/namespace-list.yaml

# Secret
kubectl apply -f k8s/list/secrets-list.yaml

# Deployment
kubectl apply -f k8s/list/deployment-list.yaml

# Esperar a que el deployment esté listo
echo "⏳ Esperando a que p-go-list esté listo..."
kubectl wait --namespace=p-go-list \
  --for=condition=available deployment/list-deployment \
  --timeout=90s

# Service
kubectl apply -f k8s/list/service-list.yaml

# Ingress
kubectl apply -f k8s/list/ingress.yaml

echo "✅ p-go-list desplegado correctamente."

echo -e "\n🔍 Estado actual:"
kubectl get all -n p-go-list
kubectl get ingress -n p-go-list
