#!/usr/bin/env bash

set -e
# echo "" > coverage.txt
rm coverage.out
for d in $(go list ./... | grep -v vendor); do
    echo $d
    richgo test -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.out
        rm profile.out
    fi
done