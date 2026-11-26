package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint32 }
func genFn0(p0 int32, p1 int) uint {
    switch "" {
    case "\n\t":
        switch "a" {
        case "\n\t":
            var v5 int16 = 349
            _ = v5
        default:
            p1 -= (p1 | -166)
        }
    case "hello":
        if (true && (true && true)) {
            p1 += p1
            var v6 uint32 = 1
            _ = v6
        } else {
            v7 := GenS1{F0: 625, F1: p1, F2: 1}
            _ = v7
        }
    default:
        fmt.Println(p1, p0)
    }
    p1 += p1
    _ = p0
    _ = p1
    return 640
}
func genFn1(p0 uint32) bool {
    fmt.Println(p0)
    for fi8 := 0; fi8 < 20; fi8++ {
        p0 += (p0 - p0)
        _ = fi8
    }
    _ = p0
    return true
}

func main() {
    var v1 uint = 42
    v2 := make(map[int]bool)
    var v3 float32 = (1.5 * 0.0)
    var v4 int16 = 1000
    fmt.Println(v4)
    var v9 int64 = 243
    v2[32767] = false
    for fi10 := 0; fi10 < 10; fi10++ {
        fmt.Println(v9)
        switch (fi10 % ((-1) | 1)) {
        case -9223372036854775808:
            var v11 int = 128
            _ = v11
        }
        _ = fi10
    }
    v1 = (v1 | (^(v1)))
    v12 := GenS1{F0: 0, F1: -1, F2: 2147483647}
    var v13 float32 = 1e10
    fmt.Println(v9)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
