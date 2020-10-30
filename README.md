# Erroroid [![Build Status](https://travis-ci.com/sarpdoruk/erroroid.svg?branch=main)](https://travis-ci.com/sarpdoruk/erroroid) [![Go Report Card](https://goreportcard.com/badge/github.com/sarpdoruk/erroroid)](https://goreportcard.com/report/github.com/sarpdoruk/erroroid) [![codecov](https://codecov.io/gh/sarpdoruk/erroroid/branch/main/graph/badge.svg?token=BJ2GUL0D47)](https://codecov.io/gh/sarpdoruk/erroroid/branch/main/graph/badge.svg?token=BJ2GUL0D47)

It's basically `error` type on steroids that helps you to easily find the error in your code by printing the filename, line, function name and the error message in a given format to stdout.

## Usage

You can set a custom format for your errors by including the following reserved placeholders in your format string;
- `#file` Prints the full path of the file that error occurred in.
- `#line` Prints the line number of the file that error occurred.
- `#func` Prints the function name that error occurred in.
- `#err` Prints the error as a string.

This;
```
ERROR: [#file:#line @#func] #err
```

Turns into this;
```
ERROR: [/path/to/my/go/file.go:123 @myFunc] this is a custom error message
```

Currently, all placeholders must be present in the format string.

> `Error()` method of `Erroroid` returns native type `error` so that you can use it anywhere in your code safely.

Check [godocs](https://pkg.go.dev/github.com/sarpdoruk/erroroid) for examples and more.
