#!/bin/bash
set -e
go build

echo 'ls | cat'
diff -u <(ls | cat) <(ls | ./cat)

echo 'cat main.go'
diff -u <(cat main.go) <(./cat main.go)

echo 'cat main.go runtest.sh'
diff -u <(cat main.go runtest.sh) <(./cat main.go runtest.sh)
