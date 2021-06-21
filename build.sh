#!/bin/bash

set -e
set -o nounset
set -o pipefail

source "scripts/build.helpers"

goal_build-cli() {
  scripts/cli-binary.sh
}

goal_build-server() {
  scripts/server-binary.sh
}

goal_containerize() {
  ko publish ./cmd/skillsmapperd
}

goal_skaffold-run() {
  skaffold run --tail --default-repo github.com/codetaming/skillsmapper
}

goal_skaffold-dev() {
  skaffold dev --port-forward --tail --default-repo github.com/codetaming/skillsmapper
}

goal_test-unit() {
  go test ./...
}

goal_run() {
  go run ./cmd/skillsmapperd.go
}

goal_test-e2e() {
  setup-newman
  run-newman
}

goal_deploy-cloudrun() {
  gcloud run deploy --image gcr.io/"${PROJECT_ID}"/ko/gcr.io/codetaming-skillsmapper/ko/skillsmapperd-c89e9d07d866ec9370a3f2eb76542a1b@sha256:8dc64722fa4c1e281cd76b377e3e6de502b664c5b1b31e9bcf0cc192f006cd76 --platform managed
}

goal_help() {
  echo "usage: $0 <goal>
    goal:
    build-server             -- Build the deployable server artifacts
    build-cli                -- Build the deployable cli artifacts
    containerize             -- Build the docker container for the app
    test-unit                -- Run unit tests
    test-e2e                 -- Run newman tests
    deploy-cloudrun          -- Deploy to Cloud Run
    skaffold-run             -- Run locally using Skaffold
    skaffold-dev             -- Auto-redeploy locally using Skaffold
  "
  exit 1
}

main() {
  TARGET=${1:-}
  if [ -n "${TARGET}" ] && type -t "goal_$TARGET" &>/dev/null; then
    "goal_$TARGET" "${@:2}"
  else
    goal_help
  fi
}

main "$@"