#!/usr/bin/env bash

set -e
# echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor); do
    echo $d
    richgo test -coverprofile=coverage.out -covermode=atomic $d
    # if [ -f coverage.out ]; then
    #     # cat coverage.out >> coverage.txt
    #     # rm coverage.out
    # fi
done