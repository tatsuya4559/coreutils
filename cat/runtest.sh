#!/bin/bash
set -e
go build

diff -u <(ls | cat) <(ls | ./cat)
diff -u <(cat main.go) <(./cat main.go)
diff -u <(./cat runtest.sh) <(./cat main.go) >/dev/null || true
