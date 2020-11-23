#!/bin/bash

BUILD_DIR="build"
PKG_ROOT="$BUILD_DIR/pkgroot"
PACKAGE_NAME="KeyboardChecker"
VERSION=$(cat ./version)
OUTFILE="/tmp/KeyboardChecker.pkg"
RESSOURCES="resources"

/bin/rm $OUTFILE

/bin/rm -rf $BUILD_DIR
env GOOS=darwin GOARCH=amd64 go build -o $BUILD_DIR/$PACKAGE_NAME -ldflags "-s -w -X main.Version=$VERSION"

mkdir -p $PKG_ROOT/usr/local/KeyboardChecker
mkdir -p $PKG_ROOT/Library/LaunchDaemons

mv $BUILD_DIR/$PACKAGE_NAME $PKG_ROOT/usr/local/KeyboardChecker/KeyboardChecker

pkgbuild --identifier $PACKAGE_NAME --version $VERSION --root $PKG_ROOT --install-location / $OUTFILE
