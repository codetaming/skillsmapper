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

# Install kind and create a cluster
brew install kind
kind create cluster
export KUBECONFIG="$(kind get kubeconfig --name="kind")"