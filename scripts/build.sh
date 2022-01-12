#!/usr/bin/env bash

PACKAGE=$1
if [[ -z "$PACKAGE" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi

TAG=$2

platforms=(
 "linux/386"
 "linux/amd64"
 "linux/arm"
 "linux/arm64"
)

rm -rf platforms-$TAG/*
for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    echo 'Building' $GOOS-$GOARCH
    OUTPUT_NAME='m-agent-'$GOOS-$GOARCH

    env GOOS=$GOOS GOARCH=$GOARCH VERSION=$TAG go build -ldflags "-X main.Version=$TAG" -v -o platforms-$TAG/$OUTPUT_NAME $PACKAGE

    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi

    cd platforms-$TAG || exit
    mv $OUTPUT_NAME m-agent
    tar -czvf $OUTPUT_NAME-$TAG.tar.gz m-agent
    rm -rf m-agent
    cd ..
done
