package main

import (
    "fmt"
    "log"
    "github.com/nu7hatch/gouuid"
)

func main() {
    out, err := uuid.NewV4()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s", out)
}