#!/usr/bin/env bash
set -eu

PLATFORM=${PLATFORM:-}
VERSION=${VERSION:-$(cat VERSION)}
GITCOMMIT=${GITCOMMIT:-$(git rev-parse --short HEAD 2> /dev/null || true)}
BUILDTIME=${BUILDTIME:-$(date -u +"%Y-%m-%dT%H:%M:%SZ")}

PLATFORM_LDFLAGS=
if test -n "${PLATFORM}"; then
	PLATFORM_LDFLAGS="-X \"${CLI_SOURCE}/version.PlatformName=${PLATFORM}\""
fi

GOOS="${GOOS:-$(go env GOHOSTOS)}"
GOARCH="${GOARCH:-$(go env GOHOSTARCH)}"
if [ "${GOARCH}" = "arm" ]; then
	GOARM="${GOARM:-$(go env GOHOSTARM)}"
fi

CLI_TARGET="build/smctl-$GOOS-$GOARCH"
if [ "${GOARCH}" = "arm" ] && [ -n "${GOARM}" ]; then
	CLI_TARGET="${CLI_TARGET}-v${GOARM}"
fi

if [ "${GOOS}" = "windows" ]; then
	CLI_TARGET="${CLI_TARGET}.exe"
fi
export CLI_TARGET

export CLI_SOURCE="github.com/codetaming/skillsmapper/cli"

export CLI_LDFLAGS="\
    -w \
    ${PLATFORM_LDFLAGS} \
    -X \"${CLI_SOURCE}/version.GitCommit=${GITCOMMIT}\" \
    -X \"${CLI_SOURCE}/version.BuildTime=${BUILDTIME}\" \
    -X \"${CLI_SOURCE}/version.Version=${VERSION}\" \
    ${CLI_LDFLAGS:-} \
"

SERVER_TARGET="build/skillsmapperd-$GOOS-$GOARCH"
if [ "${GOARCH}" = "arm" ] && [ -n "${GOARM}" ]; then
	SERVER_TARGET="${SERVER_TARGET}-v${GOARM}"
fi

if [ "${GOOS}" = "windows" ]; then
	SERVER_TARGET="${SERVER_TARGET}.exe"
fi
export SERVER_TARGET

export SERVER_SOURCE="github.com/codetaming/skillsmapper/cmd/skillsmapperd"

export SERVER_LDFLAGS="\
    -w \
    ${PLATFORM_LDFLAGS} \
    -X \"github.com/codetaming/skillsmapper/internal/version.GitCommit=${GITCOMMIT}\" \
    -X \"github.com/codetaming/skillsmapper/internal/version.BuildTime=${BUILDTIME}\" \
    -X \"github.com/codetaming/skillsmapper/internal/version.Version=${VERSION}\" \
    ${SERVER_LDFLAGS:-} \
"