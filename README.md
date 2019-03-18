# Go:Generate Version
This package generates a `version.go` file containing a `Version` and `Commit` constant in your project root which lets you include the version of your Go binary at build time instead of using flags.

# Usage
Include this as a Git Submodule and link to it from your `main.go`.

Run the following from the root of your project to place it at `./generators/versioning`:

```sh
# using ssh
git submodule add git@github.com:usvc/go-generate-version.git ./generators/versioning

# using https
git submodule add https://github.com/usvc/go-generate-version.git ./generators/versioning
```

So if you've placed this package at `~/generators/versioning` (as the above example does, where `~` is the project root), insert the following line above your `package main` in your `main.go`:

```go
////go:generate go run ./generators/versioning/main.go
```

Then, running `go generate` will create the `version.go` file, exposing the `Version` and `Commit` variable for you to consume.

# Failure Conditions
On failure, the `Version` constant will be set to `0.0.0-unset` and the `Commit` constant will be set to `0000000`.

## No Git Tags
To fix this, add a Git tag (`git tag 0.0.0`).

## No Git Commits
You will require your Git repository to have at least one commit and one tag for this package to work.

# Cheers
