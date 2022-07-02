#!/usr/bin/bash

REPO=$(
  cd $(dirname $0) || exit
  pwd
)

ASSETS="false"
BINARY="false"

buildAssets() {
  cd $REPO || exit
  rm -rf assets/out

  cd $REPO/assets || exit
  yarn install --frozen-lockfile
  yarn export
  zip -q assets.zip -r out

  cd $REPO || exit
  mv assets/assets.zip .
}

buildBinary() {
  # linux amd64
  export GOOS=linux GOARCH=amd64
  go build -a -o build/linux_amd64_random_donate .

  # linux arm64
  export GOOS=linux GOARCH=arm64
  go build -a -o build/linux_arm64_random_donate .

  # darwin amd64
  export GOOS=darwin GOARCH=amd64
  go build -a -o build/darwin_amd64_random_donate .

  # windows amd64
  export GOOS=windows GOARCH=amd64
  go build -a -o build/windows_amd64_random_donate.exe .
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
