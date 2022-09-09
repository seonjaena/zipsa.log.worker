#!/bin/bash

export ZIP_DIR=~/Desktop
export MAIN_FILE=main
export MAIN_GO_FILE=main.go
export MAIN_ZIP_FILE=main.zip
export GOARCH=amd64
export GOOS=linux

rm -rf $MAIN_FILE
rm -rf $MAIN_ZIP_FILE
rm -rf $ZIP_DIR/$MAIN_ZIP_FILE

echo "GOARCH = $GOARCH"
echo "GOOS = $GOOS"

go build -o $MAIN_FILE $MAIN_GO_FILE

zip $MAIN_ZIP_FILE $MAIN_FILE

mv $MAIN_ZIP_FILE $ZIP_DIR