#!/bin/bash
set -e
go build

echo 'wc main.go'
diff -uw <(./wc main.go) <(wc main.go)

echo 'wc main.go runtest.sh'
diff -uw <(./wc main.go runtest.sh) <(wc main.go runtest.sh)

echo 'wc main.go hoge'
diff -uw <(./wc main.go hoge 2>&1) <(wc main.go hoge 2>&1)
