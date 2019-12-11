package main

import (
    "fmt"
    "github.com/liguanghuigithub/go/example/stringutil"
)

func main() {

    fmt.Println("hello")
    fmt.Printf("world\n")

    fmt.Printf(stringutil.Reverse("I'm in new directory") + "\n")
}
