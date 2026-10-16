package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 42
    var v3 string = "hello"
    v4 := make(map[int]bool)
    var v5 uint = 18446744073709551615
    v6 := []string{v3}
    if false {
        switch v3 {
        case "a":
            v7 := v4[-2147483648]
            _ = v7
        case "foo bar":
            if false {
                v8 := GenS1{F0: (v2 / ((v2) | 1)), F1: "\n\t"}
                v3 = "a"
                _ = v8
            }
        default:
            var v9 uint = (v5 / ((0) | 1))
            _ = v9
        }
        var v10 int64 = v2
        _ = v10
    } else {
        v11 := GenS1{F0: (v2 % ((-1) | 1)), F1: (v3 + v3)}
        _ = v11
    }
    v5 = v5
    if false {
        v12 := GenS1{F0: -906, F1: v3}
        _ = v12
    }
    switch "foo bar" {
    case "\n\t":
        v5 -= 1
    }
    _ = len(v1)
    v2 = -9223372036854775808
    var v13 uint8 = 255
    v4[-1] = false
    switch (v3 + v3) {
    case "foo bar":
        v3 = "a"
    case "\n\t":
        v14 := GenS1{F0: (v2 / (((v2 | (v2 + (-(v2))))) | 1)), F1: (("foo bar" + v3) + (v3 + "\n\t"))}
        _ = v14
    }
    if false {
        _ = len(v1)
        if (128 > 127) {
            v5 = (v5 ^ (v5 * v5))
            fmt.Println(v3, v13, v5)
        }
    }
    fmt.Println(v5)
    v13 = 0
    if false {
        delete(v4, 4294967295)
        var v15 float32 = 1e10
        _ = v15
    }
    v2 = (v2 + -919)
    v5 = (v5 | v5)
    var v16 uint64 = 687
    var v17 int8 = 42
    var v18 float32 = ((1e10 * 0.1) + 3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
