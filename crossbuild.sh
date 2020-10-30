#!/usr/bin/env bash

set -e
set -u

platforms=(
	"linux/amd64"
	"linux/386"
	"darwin/amd64"
	"windows/amd64"
	"windows/386"
	"freebsd/amd64"
	"freebsd/386"
	"openbsd/amd64"
	"openbsd/386"
	"netbsd/amd64"
	"netbsd/386"
	"dragonfly/amd64"
	"linux/arm"
	"linux/arm64"
	"freebsd/arm"
	"openbsd/arm"
	"linux/mips64"
	"linux/mips64le"
	"netbsd/arm"
	"linux/ppc64"
	"linux/ppc64le"
	"linux/s390x"
)

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name="prometheus_demo_service-$(git describe --tags).${GOOS}-${GOARCH}"
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    echo "Building ${output_name}..."
    env GOOS=$GOOS GOARCH=$GOARCH go build -o "${output_name}"
done
