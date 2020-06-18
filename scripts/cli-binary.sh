#!/usr/bin/env bash
#
# Build scripts static binary for the host OS/ARCH
#

set -eu -o pipefail

source ./scripts/variables.sh

echo "Building statically linked $CLI_TARGET"
export CGO_ENABLED=0
go build -o "${CLI_TARGET}" --ldflags "${CLI_LDFLAGS}" "${CLI_SOURCE}"

ln -sf "$(basename "${CLI_TARGET}")" build/smctl
