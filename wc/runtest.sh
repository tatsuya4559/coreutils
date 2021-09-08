#!/bin/bash
set -e
go build

diff -uw <(./wc main.go) <(wc main.go)

diff -uw <(./wc main.go runtest.sh) <(wc main.go runtest.sh)

diff -uw <(./wc main.go hoge 2>&1) <(wc main.go hoge 2>&1)
