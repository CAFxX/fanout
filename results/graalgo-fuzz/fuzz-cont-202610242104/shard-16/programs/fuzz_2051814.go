package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS3 struct { F0 int32; F1 uint8; F2 int16 }

func main() {
    var v1 uint8 = 255
    v2 := []string{"a"}
    var v3 float64 = (-2.25 - 1.5)
    if (1 == 404) {
        switch "foo bar" {
        case "hello":
            v1 = v1
        default:
            v3 -= v3
        }
    }
    fmt.Println(v1, v1)
    var v4 int = -773
    if false {
        fmt.Println(v3, v3)
    } else {
        v3 = (-2.25 / v3)
    }
    if false {
        var v5 int16 = -614
        v4 += v4
        fmt.Println(v3)
        _ = v5
    } else {
        v1 -= (v1 + (^(v1)))
        if false {
            v6 := GenS1{F0: 2147483648}
            _ = v6
        }
    }
    v7 := GenS2{F0: 1, F1: 0, F2: -1}
    v8 := GenS1{F0: 4294967295}
    v1 -= (v1 | v1)
    var v9 uint64 = 42
    v10 := GenS3{F0: 0, F1: 243, F2: 1000}
    if len(v2) > 0 {
        v2[0] = ""
    }
    v9 += 277
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
