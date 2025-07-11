// cmd/neuralcryptoserviceapix/main.go
package main

import (
"flag"
"log"
"os"

"neuralcryptoserviceapix/internal/neuralcryptoserviceapix"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := neuralcryptoserviceapix.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
