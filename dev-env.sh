#!/bin/bash
brew install gettext
brew link --force gettext
brew install skaffold
GO111MODULE=on go get github.com/google/ko/cmd/ko