#!/bin/bash -e
platform="Linux x86_64"
if [ "$(uname -sm)" != "${platform}" ]; then
    echo Must run ${platform}
    exit 1
fi

# locale--independent sort order
export LANG=
# Use Go 1.11 modules and ignore GOPATH
export GO111MODULE=on

# Creating version number with buildkite commit(7)
echo --- Getting version
height=$(git rev-list HEAD --first-parent --count)
version=$(git rev-parse --short=7 ${BUILDKITE_COMMIT:-HEAD})
full_version=${height}-${version}
echo "Version:      ${version}"
echo "Full version: ${full_version}"

# integration testing for all the services including cloudwatch, storage, Ayla and core 
pkg=github.com/pagidi/GoLang

echo --- start buiding
go build -o GoPractice-$(version)
echo --- End buiding

