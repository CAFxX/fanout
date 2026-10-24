package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }

func main() {
    var v1 int8 = 0
    v2 := []uint32{42, 2147483648, 2, 4294967295}
    var v3 uint = 895
    var v4 uint8 = 140
    var v5 uint64 = 1
    fmt.Println(v3, v4)
    v6 := GenS1{F0: -28, F1: (v4 + 255)}
    v7 := GenS1{F0: 1000, F1: 0}
    v1 = -1
    fmt.Println(v5)
    v8 := GenS1{F0: -32768, F1: v4}
    if (!((v5 <= v5) && (false || false))) {
        var v9 string = ("foo bar" + ("\n\t" + ("" + "")))
        _ = v9
    }
    var v10 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
