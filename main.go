package main

import (
    "log"
    "github.com/z6wdc/go-avro/cmd"
)

func main() {
    if err := cmd.Execute(); err != nil {
        log.Fatal(err)
    }
}
