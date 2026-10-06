# go-playground

## Setup

### Go 1.25.3

Every `go.mod` here says `go 1.25.3`. Install it system-wide (Linux amd64):

```bash
cd /tmp
curl -fLO https://go.dev/dl/go1.25.3.linux-amd64.tar.gz
echo "0335f314b6e7bfe08c3d0cfaa7c19db961b7b99fb20be62b0a826c992ad14e0f  go1.25.3.linux-amd64.tar.gz" | sha256sum -c -
sudo tar -C /usr/local -xzf go1.25.3.linux-amd64.tar.gz
sudo ln -sf /usr/local/go/bin/go /usr/local/bin/go
sudo ln -sf /usr/local/go/bin/gofmt /usr/local/bin/gofmt

go version   # go1.25.3 linux/amd64
```

`rest_api` uses SQLite through cgo, so it also needs `gcc`.

### Go workspace

Most folders (`bank`, `rest_api`, `list`, ...) are separate Go modules. Create a workspace so VSCode (gopls) sees all of them. `go.work` is gitignored, so do this once per clone:

```bash
go work init && go work use -r .
```

### Tools

Install the VSCode Go extension (`golang.go`), then:

```bash
go install golang.org/x/tools/gopls@latest
go install github.com/go-delve/delve/cmd/dlv@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
```

Make sure `~/go/bin` is on your `PATH`. For debugging, see [docs/delve-debugger-setup.md](docs/delve-debugger-setup.md).

## Go build and run project

```bash

go build -o myapp ./src/

./myapp

```
