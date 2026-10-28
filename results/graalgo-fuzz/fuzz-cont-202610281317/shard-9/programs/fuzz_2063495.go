package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint64 = 42
    var v2 int8 = -74
    v3 := make(map[string]int)
    var v4 string = "foo bar"
    v2 -= (v2 * v2)
    var v5 uint8 = 1
    if (32767 > 128) {
        var v6 uint64 = 4294967295
        fmt.Println(v6)
        v1 = 273
        _ = v6
    }
    var v7 uint32 = 2147483647
    v8 := GenS1{F0: false}
    _ = len(v3)
    fmt.Println(v7, v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
