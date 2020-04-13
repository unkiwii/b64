package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func main() {
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	n, err := io.Copy(encoder, os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%d bytes written, %v", n, err)
	}
}
