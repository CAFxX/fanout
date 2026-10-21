package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 string; F1 uint16; F2 int8 }

func main() {
    var v1 bool = false
    var v2 int8 = 72
    v3 := make(map[uint32]int)
    v4 := make(map[string]uint32)
    v5 := make(map[int]string)
    v2 += v2
    switch 838 {
    case 32767:
        var v6 uint16 = 211
        _ = v6
    }
    v7 := GenS1{F0: v2}
    var v8 int8 = (v2 | v2)
    fmt.Println(v2)
    v2 -= v8
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS2{F0: (v8 & (^(v8))), F1: v1}
        switch (fi9 & 9223372036854775807) {
        case 4294967295:
            v2 = (^(v8))
        case 73:
            v1 = true
        default:
            v2 -= (v2 ^ 100)
        }
        _ = fi9
        _ = v10
    }
    v8 += -128
    if false {
        v11 := GenS1{F0: v2}
        fmt.Println(v2, v1, v2)
        v1 = (0.0 == float32(0.1))
        _ = v11
    } else {
        switch "\n\t" {
        case "hello":
            fmt.Println(v2)
        case "\n\t":
            v12 := GenS1{F0: -1}
            _ = v12
        case "foo bar":
            v2 = (v8 % ((-1) | 1))
        default:
            v2 -= (v8 ^ 1)
        }
        var v13 int16 = 0
        _ = v13
    }
    v2 -= v2
    v14 := GenS3{F0: "a", F1: 143, F2: (^(v8))}
    v8 -= (v2 | v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
