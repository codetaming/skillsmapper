#!/bin/bash
brew install gettext
brew link --force gettext
brew install skaffold
brew upgrade skaffold
GO111MODULE=on go get github.com/google/ko/cmd/ko

#ENV
export KO_DOCKER_REPO=gcr.io/codetaming-skillsmapper/ko
export PROJECT_ID=cloudtaming-skillsmapper
gcloud components update
gcloud auth configure-docker

# Install kind and create scripts cluster
brew install kind
kind create cluster
# export KUBECONFIG="$(kind get kubeconfig --name="kind")"

# Install Knative on Docker Desktop
kubectx docker-desktop
curl -L https://github.com/knative/serving/releases/download/v0.12.0/serving.yaml \
  | kubectl apply --filename -

# Swagger
brew tap go-swagger/go-swagger
brew install go-swagger

swagger serve swagger.yaml
swagger generate model --spec=swagger.yaml
# shellcheck disable=SC2102
swagger generate client [-f ./swagger.json] -A [application-name [--principal [principal-name]]