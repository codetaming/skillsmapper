#!/usr/bin/env bash
#
# Build scripts static binary for the host OS/ARCH
#

set -eu -o pipefail

source ./scripts/variables.sh

echo "Building statically linked $SERVER_TARGET"
export CGO_ENABLED=0
go build -o "${SERVER_TARGET}" --ldflags "${SERVER_LDFLAGS}" "${SERVER_SOURCE}"

ln -sf "$(basename "${SERVER_TARGET}")" build/skillsmapperd
