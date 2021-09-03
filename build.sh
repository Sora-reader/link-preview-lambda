#!/bin/bash

GOOS=linux
GOARCH=amd64
CGO_ENABLED=0

go build -o main
[ $? = '1' ] && echo "Error building module" && exit 1
zip dist.zip main
rm main
[ $? = '1' ] && echo "Error zipping" && exit 1

echo "Build finished"
