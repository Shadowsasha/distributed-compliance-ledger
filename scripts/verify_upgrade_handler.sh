#!/usr/bin/env bash

set -euo pipefail

VER="$1"

result=$(grep 'SetUpgradeHandler(' -a1 ../app/app.go | grep "\"$VER\"" >/dev/null) || true

if [[ $? -eq 0 ]]; then
    echo "Upgrade handle v$VER exists"
else
    echo "Upgrade handle not v$VER exists"
    exit 1
fi

exit 0