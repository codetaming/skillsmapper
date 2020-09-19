#!/usr/bin/env bash
docker run -e RUN_LOCAL=true -e FILTER_REGEX_EXCLUDE="(.*vendor/.*|.*client/.*)" -e VALIDATE_ENV=false -e VALIDATE_ALL_CODEBASE=false -v "$PWD":/tmp/lint github/super-linter