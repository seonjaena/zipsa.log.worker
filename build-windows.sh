#!/bin/bash

while [ $# -gt 0 ]; do
  case "$1" in
    --profile=*)
      export PROFILE="${1#*=}"
      ;;
    *)
      printf "****************************\n"
      printf "* Error: Invalid argument. *\n"
      printf "****************************\n"
      exit 1
  esac
  shift
done

export ZIP_DIR=~/Desktop
export MAIN_FILE=main.exe
export MAIN_GO_FILE=main.go
export MAIN_ZIP_FILE=main.zip
export GOARCH=amd64
export GOOS=windows

rm -rf $MAIN_FILE
rm -rf $MAIN_ZIP_FILE
rm -rf $ZIP_DIR/$MAIN_ZIP_FILE

echo "GOARCH = $GOARCH"
echo "GOOS = $GOOS"

go build -o -ldflags="-X zipsa.log.worker/properties.Profile=$PROFILE" $MAIN_FILE $MAIN_GO_FILE

zip $MAIN_ZIP_FILE $MAIN_FILE

mv $MAIN_ZIP_FILE $ZIP_DIR