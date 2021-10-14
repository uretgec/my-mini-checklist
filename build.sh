#!/bin/bash

# Usage: ./build service-file-name
# Usage: ./build my-mini-checklist
# Before run: chmod a+x build.sh

# Local Variables
ARGS=("$@")
SERVICE_NAME="${ARGS[0]}"
#SERVICE_VERSION="$(cat VERSION)"
OS_TYPE="$(go env GOOS)"
ARCH_TYPE="$(go env GOARCH)"

if [[ $# -eq 0 ]] ; then
    echo "Service file name not found"
    exit 1
fi

GOOS=$OS_TYPE GOARCH=$ARCH_TYPE go build -ldflags="-w -s" -o "$SERVICE_NAME"

if [ -f "$SERVICE_NAME" ]; then
    echo "$SERVICE_NAME build successfully"

    chmod +x $SERVICE_NAME

    mv -f $SERVICE_NAME ./build
    echo "$SERVICE_NAME file moved into build folder"

else
    echo "Something wrong for building service"
fi
