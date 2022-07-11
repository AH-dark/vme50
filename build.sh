#!/usr/bin/bash

REPO=$(
  cd $(dirname $0) || exit
  pwd
)

ASSETS="false"
BINARY="false"

buildAssets() {
  cd $REPO || exit
  rm -rf assets/build

  cd $REPO/assets || exit
  yarn install --frozen-lockfile
  yarn build
  zip -q assets.zip -r build

  cd $REPO || exit
  mv assets/assets.zip .
}

buildBinary() {
  export CGO_ENABLED=1

  # linux amd64
  export GOOS=linux GOARCH=amd64 CC=gcc
  go build -a -o build/random_donate_linux_amd64 .

  # linux arm64
  export GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc
  go build -a -o build/random_donate_linux_arm64 .

  # darwin amd64
  export GOOS=linux GOARCH=arm CC=arm-linux-gnueabihf-gcc
  go build -a -o build/random_donate_linux_arm .

  # windows amd64
  export GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc
  go build -a -o build/random_donate_windows_amd64.exe .
}

usage() {
  echo "Usage: $0 [-a] [-b] [-f]" 1>&2
  exit 1
}

while getopts "abf" o; do
  case "${o}" in
  a)
    ASSETS="true"
    BINARY="true"
    ;;
  f)
    ASSETS="true"
    ;;
  b)
    BINARY="true"
    ;;
  *)
    usage
    ;;
  esac
done

if [ "$ASSETS" = "true" ]; then
  buildAssets
fi

if [ "$BINARY" = "true" ]; then
  buildBinary
fi
