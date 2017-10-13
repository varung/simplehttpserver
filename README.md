# simplehttpserver
I've found python -m SimpleHTTPServer PORT very useful for quickly serving files from a computer. However, I've noticed python's SimpleHTTPServer sometimes crashes or hangs on serving large files, and it cannot handle multiple connections concurrently.

This replacement in go solves all of the above issues, and is very easy to install, assuming one has the go tool chain installed(which itself is very easy to install.)

# How to Use:

```
# gets and installs the binary $GOPATH/bin
go get github.com/varung/simplehttpserver
# serves the current directory on PORT (8000 by default)
simplehttpserver PORT
```
(assuming $GOHOME/bin is in your $PATH)

or if you are in this repository's directory:

```go run simplehttpserver.go PORT```



