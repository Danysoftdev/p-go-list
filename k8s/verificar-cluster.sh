#!/bin/bash

echo "🔍 Verificando cluster activos"
kind get clusters

echo "📌 Verificando nodos:"
kubectl get nodes

echo -e "\n📦 Pods en p-go-list:"
kubectl get pods -n p-go-list

echo -e "\n🧠 Pods en mongo-ns:"
kubectl get pods -n mongo-ns

echo -e "\n📦 Verificando servicio de mongo:"
kubectl get svc -n mongo-ns

echo -e "\n🌐 Ingress en p-go-list:"
kubectl get ingress -n p-go-list
