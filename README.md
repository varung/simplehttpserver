# simplehttpserver
I've found python -m SimpleHTTPServer PORT very useful for quickly serving files from a computer.
However, I've noticed python's simplehttpserver sometimes crashes or hangs on serving large files.
This is a replacement in go, which should be more performant.

usage:

```
go get github.com/varung/simplehttpserver
simplehttpserver PORT
```
(assuming $GOHOME/bin is in your $PATH)

or if you are in this repository's directory:

```go run simplehttpserver.go PORT```



