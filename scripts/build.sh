#!/usr/bin/env bash
set -o errexit -o pipefail

while getopts ":cgo:o:a:n:v:b:" opt; do
  case $opt in
  cgo)
    CGO_ENABLED="${OPTARG:-0}"
    ;;
  o)
    GOOS="${OPTARG}"
    ;;
  a)
    GOARCH="${OPTARG}"
    ;;
  n)
    APP_NAME="${OPTARG}"
    ;;
  v)
    APP_VERSION="${OPTARG}"
    ;;
  b)
    BUILD_DIR="${OPTARG}"
    ;;
  \?)
    echo "Invalid option: -${OPTARG}" >&2
    exit 1
    ;;
  esac
done

mkdir -p "${BUILD_DIR}"

# Build the binary
echo "Building for ${GOOS}/${GOARCH} ${APP_NAME}:${APP_VERSION}"

CGO_ENABLED="${CGO_ENABLED}"
GOOS="${GOOS}"
GOARCH="${GOARCH}"
go build -o "${BUILD_DIR}"/package/geekbot-"${GOARCH}" \
  -ldflags "-X=${APP_NAME}/cmd.appVersion=${APP_VERSION}"

ls -lh "${BUILD_DIR}"/package

exit $?
