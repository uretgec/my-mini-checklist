#!/bin/bash

# Usage: ./build service-file-name
# Usage: ./build myminichecklist
# Before run: chmod a+x build.sh

# Local Variables
ARGS=("$@")
SERVICE_NAME="${ARGS[0]}"
#SERVICE_VERSION="$(cat VERSION)"

GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o "$SERVICE_NAME"

if [ -f "$SERVICE_NAME" ]; then
    echo "$SERVICE_NAME build successfully"

    mv $SERVICE_NAME ./build
    echo "$SERVICE_NAME file moved into build folder"

else
    echo "Something wrong for building service"
fi
