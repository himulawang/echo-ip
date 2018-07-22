#!/usr/bin/env bash

# Verify that we are inside echo-ip/scripts dir.
# This is because we use relative paths, and we
# assume we are in the echo-ip/scripts directory.
REVPWD=$(echo "$PWD" | rev)

REVPWD1=$(echo "$REVPWD" | cut -d/ -f1)
REVPWD2=$(echo "$REVPWD" | cut -d/ -f2)

REVSCRIPTS=$(echo "scripts" | rev)
REVECHOIP=$(echo "echo-ip" | rev)

echo "Verifying we are in the correct directory ..."

if [ "$REVPWD1" != "$REVSCRIPTS" ]; then
    echo "Not in correct dir, cd into: echo-ip/scripts"
    exit 1
fi

if [ "$REVPWD2" != "$REVECHOIP" ]; then
    echo "Not in correct dir, cd into: echo-ip/scripts"
    exit 1
fi


NAME="echo-ip"

cd ../

echo "Creating bin dir (if it does not exist)"
mkdir -p bin

cd cmd/echo-ip

echo "Removing old builds"
rm ../../bin/*

GOOS="linux"
GOARCH="amd64"

echo "Building for: $GOOS/$GOARCH"
GOOS=$GOOS GOARCH=$GOARCH go build -o ../../bin/$NAME

echo "Build successful"