#!/bin/bash

set -ex

minikuberestart

# Install glooctl
glooctl install gateway

kubectl apply -f deployment.yaml
kubectl apply -f gloo-vs.yaml

k label upstream -n gloo-system default-awang-apiserver-console-1234 discovery.solo.io/function_discovery=enabled

glooctl proxy url