# Base64 binaries

This contains programs to work with base64 sutff

# Install

`go install github.com/unkiwii/b64`

This will install binaries in your $GOPATH/bin folder, you should have that in you $PATH to be able to call the programs directly

# Who to use

b64encoder and b64decoder both read standrad input and write to standard output, or standard error when an error occurs

## Examples

```
$ echo Hello | b64encoder
SGVsbG8K

$ echo Hello > file
$ cat file | b64encoder
SGVsbG8K

$ echo SGVsbG8K > b64decoder
Hello

$ echo SGVsbG8K > file
$ cat file | b64decoder
Hello
```
