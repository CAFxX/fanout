package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 bool }

func main() {
    v1 := []int{127, -2147483648, 294}
    var v2 int64 = 0
    var v3 uint8 = 2
    var v4 uint32 = 2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "foo bar" {
    case "\n\t":
        v4 = 788
    case "a":
        _ = len(v1)
    default:
        var v5 uint = 746
        _ = v5
    }
    var v6 uint8 = v3
    v3 = 255
    v3 = v6
    v7 := GenS1{F0: 369, F1: 470, F2: false}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3)
    var v8 uint8 = v6
    if len(v1) > 0 {
        v1[0] = 1
    }
    v3 -= 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
