package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
func genFn0(p0 uint, p1 uint32, p2 bool) uint16 {
    if p2 {
        if p2 {
            var v4 uint = 983
            _ = v4
        }
        fmt.Println(p2)
    }
    fmt.Println(p2, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 float64 = 0.0
    var v2 int64 = -9223372036854775808
    var v3 uint8 = 255
    v2 += 42
    v1 += v1
    v5 := GenS1{F0: int16(v3), F1: (uint32(v3) <= 785)}
    switch "\n\t" {
    case "foo bar":
        v6 := GenS1{F0: -403, F1: false}
        _ = v6
    }
    v3 = (v3 | (v3 / ((62) | 1)))
    var v7 float32 = 0.0
    v2 += 0
    v2 -= 9223372036854775807
    if true {
        var v8 uint32 = 93
        v7 = (1.5 + (v7 * v7))
        _ = v8
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v7
}
