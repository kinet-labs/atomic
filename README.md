# atomic [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov] [![Go Report Card][reportcard-img]][reportcard]

Simple wrappers for primitive types to enforce atomic access.

## Installation

```shell
$ go get -u github.com/kinet-labs/atomic@v1
```

### Legacy Import Path

As of v1.5.0, the import path `github.com/kinet-labs/atomic` is the only supported way
of using this package. If you are using Go modules, this package will fail to
compile with the legacy import path path `github.com/kinet-labs/atomic`.

We recommend migrating your code to the new import path but if you're unable
to do so, or if your dependencies are still using the old import path, you
will have to add a `replace` directive to your `go.mod` file downgrading the
legacy import path to an older version.

```
replace github.com/kinet-labs/atomic => github.com/kinet-labs/atomic v1.4.0
```

You can do so automatically by running the following command.

```shell
$ go mod edit -replace github.com/kinet-labs/atomic=github.com/kinet-labs/atomic@v1.4.0
```

## Usage

The standard library's `sync/atomic` is powerful, but it's easy to forget which
variables must be accessed atomically. `github.com/kinet-labs/atomic` preserves all the
functionality of the standard library, but wraps the primitive types to
provide a safer, more convenient API.

```go
var atom atomic.Uint32
atom.Store(42)
atom.Sub(2)
atom.CompareAndSwap(40, 11)
```