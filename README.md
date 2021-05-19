# Matcheck

A simple util for checking if a path matches a pattern according to Go's `filepath.Match`. Useful for testing patterns
for CircleCI config.

## Installation

* Install Golang toolchain, make sure that `go/bin` is in your `PATH`
* `go get github.com/smbkr/matcheck`


## Usage

`matcheck [path] "[pattern]"` eg, `matcheck /var/lib/foo "/var/*/foo"`. You should enclose the pattern with quotes to
prevent shell expansion.

It exits with `0` if the path matches, and with `1` if there is no match. Exit code `2` is used for errors.

```
~/ matcheck /var/lib '/*/lib' && matcheck /var/lib '/vaaar/lib'
MATCH ✅
NO MATCH ❌
``` 
