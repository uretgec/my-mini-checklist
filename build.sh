#!/bin/bash

# Usage: ./build service-file-name
# Usage: ./build my-mini-checklist
# Before run: chmod a+x build.sh

# Local Variables
ARGS=("$@")
SERVICE_NAME="${ARGS[0]}"
#SERVICE_VERSION="$(cat VERSION)"

GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s" -o "$SERVICE_NAME"

if [ -f "$SERVICE_NAME" ]; then
    echo "$SERVICE_NAME build successfully"

    chmod +x $SERVICE_NAME

    mv $SERVICE_NAME ./build
    echo "$SERVICE_NAME file moved into build folder"

else
    echo "Something wrong for building service"
fi
