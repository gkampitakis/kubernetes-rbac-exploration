# Kubernetes RBAC Exploration

## Description

Exploring and testing RBAC in Kubernetes

## Contents

This repo contains a Golang program interacting with Kubernetes

- creating pods
- listing pods
- and deleting pods

A Dockerfile for building a container and running it `minikube`.

> Docker image `gkabitakis/kubernetes-rbac:latest`.

In `/eks` folder I have some yaml files for running pods with the above actions ( create, list and delete). It also contains some example RBACS for experimenting.

## Helpful Commands

```bash
# Start minikube
minikube start --nodes 3 -p multinode --cpus 4
```

## Resources

- [K8s RBAC roles and bindings](https://octopus.com/blog/k8s-rbac-roles-and-bindings)
- [Configure Service Accounts for Pods](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/)
