package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint32 }

func main() {
    var v1 int64 = 0
    var v2 uint8 = 247
    v3 := []int{-2147483648, 752, 0}
    var v4 uint8 = v2
    var v5 int8 = -128
    switch "foo bar" {
    case "":
        v6 := GenS1{F0: 4294967295}
        _ = v6
    }
    fmt.Println(v5)
    if (594 <= 1) {
        for fi7 := 0; fi7 < 11; fi7++ {
            fmt.Println(v4, v4, v1)
            v2 -= v2
            v2 = v4
            _ = fi7
        }
        v4 = 2
    }
    switch 32768 {
    case 127:
        if len(v3) > 0 {
            v3[0] = 377
        }
    case 32767:
        fmt.Println(v4, v2)
    case 32768:
        fmt.Println(v1, v1)
    default:
        var v8 uint16 = 65535
        _ = v8
    }
    v1 = (v1 + v1)
    var v9 uint16 = 0
    v10 := GenS2{F0: 9223372036854775807}
    v11 := GenS3{F0: (v1 ^ v1), F1: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
