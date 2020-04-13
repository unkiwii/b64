package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func main() {
	decoder := base64.NewDecoder(base64.StdEncoding, os.Stdin)
	n, err := io.Copy(os.Stdout, decoder)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%d bytes written, %v", n, err)
	}
}
