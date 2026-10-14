package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }

func main() {
    var v1 int = 127
    var v2 uint8 = 0
    var v3 string = ("a" + "")
    var v4 uint8 = (^(v2))
    var v5 float64 = 1e10
    v1 += v1
    fmt.Println(v2, v4)
    switch "a" {
    case "a":
        var v6 int8 = -22
        _ = v6
    }
    switch v3 {
    case "":
        fmt.Println(v1, v1, v5)
    case "a":
        v7 := GenS1{F0: 0, F1: 0}
        _ = v7
    }
    fmt.Println(v5)
    var v8 uint16 = 2
    var v9 uint = 511
    v8 = v8
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
