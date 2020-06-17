#!/usr/bin/env bash
#
# Build scripts static binary for the host OS/ARCH
#

set -eu -o pipefail

source ./scripts/variables.sh

echo "Building statically linked $TARGET"
export CGO_ENABLED=0
go build -o "${TARGET}" --ldflags "${LDFLAGS}" "${SOURCE}"

ln -sf "$(basename "${TARGET}")" build/smctl
