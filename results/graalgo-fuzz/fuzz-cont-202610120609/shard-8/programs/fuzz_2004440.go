package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }

func main() {
    var v1 uint8 = 68
    var v2 uint64 = 0
    var v3 int = 493
    v4 := GenS1{F0: v3, F1: 0}
    fmt.Println(v1, v2)
    fmt.Println(v1, v2)
    v5 := []int{v3, (v3 << ((52) & 63))}
    var v6 uint8 = (^(v1))
    v5 = append(v5, (v3 << ((uint(v2)) & 63)))
    switch ("hello" + "foo bar") {
    case "a":
        v3 = v3
    }
    fmt.Println(v2, v1)
    var v7 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
