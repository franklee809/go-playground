# Setting Up Delve Debugger for Go

## What is Delve?

Delve (`dlv`) is the standard debugger for Go. It lets you set breakpoints, step through code line by line, and inspect variables — similar to `pdb` in Python.

## Install

```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

## Common Gotcha: Go Toolchain Mismatch

If your project uses a newer Go version via `go.mod` toolchain directive (e.g. `go 1.25.3`), but your system's base Go is older (e.g. 1.24.4), Delve will fail with:

```
To debug executables using DWARFv5 or later Delve must be built with Go version 1.25.0 or later
```

### Why?

`go install` for external tools uses your **base** Go version, not the toolchain version specified in your project's `go.mod`.

### Fix

Force the correct toolchain when installing:

```bash
GOTOOLCHAIN=go1.25.3 go install github.com/go-delve/delve/cmd/dlv@master
```

Verify:

```bash
go version ~/go/bin/dlv
# Should output: go1.25.3
```

## Using Delve in VSCode

Prerequisites:
- [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go) (`golang.go`) installed
- Delve installed (see above)

Steps:
1. Click the **gutter** (left of line numbers) to set a breakpoint
2. Click **"debug test"** above the test function name
3. Use the debug toolbar to step through code:
   - **Step Over** (F10) — run the current line
   - **Step Into** (F11) — go inside a function call
   - **Step Out** (Shift+F11) — finish current function and return
   - **Continue** (F5) — run until the next breakpoint
4. Watch variables in the **Variables** panel on the left

## Using Delve in Terminal

```bash
# Start debugging a test
dlv test ./function/ -- -test.run TestRecursion

# Common commands inside dlv:
break function.factorial   # set breakpoint
continue                   # run to next breakpoint
print number               # inspect a variable
next                       # step over
step                       # step into
stepout                    # step out
locals                     # show all local variables
stack                      # show call stack
quit                       # exit
```

## Example: Debugging Recursion

Given this factorial function:

```go
func factorial(number int) int {
    if number == 0 {
        return 1
    }
    return number * factorial(number-1)
}
```

Setting a breakpoint on `factorial` and continuing shows `number` decreasing each call:

```
(dlv) break function.factorial
(dlv) continue
     number = 5
(dlv) continue
     number = 4
(dlv) continue
     number = 3
(dlv) continue
     number = 2
(dlv) continue
     number = 1
(dlv) continue
     number = 0    ← base case, returns 1
```

This lets you visually confirm how recursion unwinds step by step.
