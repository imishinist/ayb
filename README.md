# ayb

# How to use

```
ayb
```

# Install

```
go install github.com/imishinist/ayb@latest
```

This requires `go1.16` or higher.

# For Go1.16

`ayb-bot` running on GAE, but GAE doesn't support go1.16 and higher now.
So add go1.15 code for GAE.
If GAE supports go1.16, remove that code.

- remove `pkg/witticism/witticism_go115.go`
- remove `pkg/witticism/statik`
- merge `pkg/witticism/witticism_go116.go` and `pkg/witticism/witticism.go`
- run `go mod tidy`
- check `.gcloudignore`
