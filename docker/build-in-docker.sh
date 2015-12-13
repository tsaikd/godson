#!/bin/bash

set -e

PN="${BASH_SOURCE[0]##*/}"
PD="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

renice 15 $$
cd "${PD}/.."

if ! type godep &>/dev/null ; then
	go get -v "github.com/tools/godep"
fi

godep restore

go get -v

if [ "${BUILDTIME}" ] && [ "${GITHASH}" ] ; then
	go build -ldflags "-X github.com/tsaikd/KDGoLib/version.BUILDTIME=${BUILDTIME} -X github.com/tsaikd/KDGoLib/version.GITCOMMIT=${GITHASH}" -o "godson-$(uname -s)-$(uname -m)"
else
	go build -o "godson-$(uname -s)-$(uname -m)"
fi

