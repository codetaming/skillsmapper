#!/bin/bash
brew install gettext
brew link --force gettext
brew install skaffold
GO111MODULE=on go get github.com/google/ko/cmd/ko

#ENV
export KO_DOCKER_REPO=gcr.io/codetaming-skillsmapper/ko
export PROJECT_ID=cloudtaming-skillsmapper
gcloud auth configure-docker