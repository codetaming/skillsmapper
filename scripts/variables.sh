#!/usr/bin/env bash
set -eu

PLATFORM=${PLATFORM:-}
VERSION=${VERSION:-$(cat VERSION)}
GITCOMMIT=${GITCOMMIT:-$(git rev-parse --short HEAD 2> /dev/null || true)}
BUILDTIME=${BUILDTIME:-$(date -u +"%Y-%m-%dT%H:%M:%SZ")}

PLATFORM_LDFLAGS=
if test -n "${PLATFORM}"; then
	PLATFORM_LDFLAGS="-X \"github.com/codetaming/skillsmapper/cli/version.PlatformName=${PLATFORM}\""
fi

export LDFLAGS="\
    -w \
    ${PLATFORM_LDFLAGS} \
    -X \"github.com/codetaming/skillsmapper/cli/version.GitCommit=${GITCOMMIT}\" \
    -X \"github.com/codetaming/skillsmapper/cli/version.BuildTime=${BUILDTIME}\" \
    -X \"github.com/codetaming/skillsmapper/cli/version.Version=${VERSION}\" \
    ${LDFLAGS:-} \
"

GOOS="${GOOS:-$(go env GOHOSTOS)}"
GOARCH="${GOARCH:-$(go env GOHOSTARCH)}"
if [ "${GOARCH}" = "arm" ]; then
	GOARM="${GOARM:-$(go env GOHOSTARM)}"
fi

TARGET="build/smctl-$GOOS-$GOARCH"
if [ "${GOARCH}" = "arm" ] && [ -n "${GOARM}" ]; then
	TARGET="${TARGET}-v${GOARM}"
fi

if [ "${GOOS}" = "windows" ]; then
	TARGET="${TARGET}.exe"
fi
export TARGET

export SOURCE="github.com/codetaming/skillsmapper/cli"