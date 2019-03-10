# Go:Generate Version
This package generates a `version.go` file containing a `Version` and `Commit` constant in your project root which lets you include the version of your Go binary at build time instead of using flags.

# Usage
Include this as a Git Submodule and link to it from your `main.go`.

For example, if you placed this package at: `~/generators/version` relative to your project root , insert the following line above your `package main` in your `main.go`:

```go
////go:generate go run ./generators/version/main.go
```

Running `go generate` will create the `version.go` file, exposing the `Version` and `Commit` variable for you to consume.

# Cheers
