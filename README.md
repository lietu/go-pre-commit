# Cross-platform Golang pre-commit hooks
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Flietu%2Fgo-pre-commit.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Flietu%2Fgo-pre-commit?ref=badge_shield)

It seemed strange to me how everyone writing [pre-commit](https://pre-commit.com) hooks for Golang would choose to not use a great cross-platform language already in their disposal, Golang, and instead would write their hooks in something that is not cross-platform - BASH.

So I took some inspirations from other repos like https://github.com/Bahjat/pre-commit-golang and https://github.com/dnephin/pre-commit-golang and implemented the commands in Go.

This should work on Windows, Linux, Mac, and basically anything Golang does (though maybe not all the supported tools work on all platforms).


## Using the hooks

You need to first install the binary from here, `go get -u github.com/lietu/go-pre-commit`

You can add these to your project's `.pre-commit-config.yaml`:

```yaml
- repo: https://github.com/lietu/go-pre-commit
  rev: master
  hooks:
    - id: errcheck
    - id: go-fmt-goimports
    - id: go-test
    - id: go-vet
    - id: gofumports
    - id: gofumpt
    - id: golangci-lint
    - id: golint
    - id: staticcheck
    - id: go-mod-tidy
```

If you want to use [golangci-lint](https://github.com/golangci/golangci-lint#install) you should first follow their installation guide. The other tools are automatically `go get`'d if they are not yet installed. If you choose to use `golangci-lint`, you don't need `errcheck`, `go-vet`, and `staticcheck` separately, as they are included by default.

Also you likely don't want to mix `go-fmt-goimports` with `gofumports` and `gofumpt` as they are just stricter variants of the same tools.


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Flietu%2Fgo-pre-commit.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Flietu%2Fgo-pre-commit?ref=badge_large)