package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int32 }

func main() {
    var v1 bool = (-9223372036854775808 < 32767)
    var v2 int32 = 42
    var v3 uint8 = 0
    var v4 uint32 = 4294967295
    v5 := []int{32767}
    var v6 int16 = -339
    v1 = (!(((-2.25 / 1.5) != (3.14159 - 1e10)) && false))
    _ = len(v5)
    if (!v1) {
        v6 -= -32768
    }
    if (!v1) {
        fmt.Println(v6, v3)
        fmt.Println(v6, v3, v1)
    } else {
        fmt.Println(v1, v3)
        switch "hello" {
        case "foo bar":
            switch "hello" {
            case "":
                v5 = append(v5, 9223372036854775807)
            default:
                v7 := GenS1{F0: -795, F1: 1, F2: v2}
                _ = v7
            }
        default:
            fmt.Println(v3, v6)
        }
    }
    v2 = 2147483647
    fmt.Println(v1, v4)
    v6 += 0
    var v8 int64 = 1
    fmt.Println(v3, v1, v2)
    for fi9 := 0; fi9 < 7; fi9++ {
        v5 = append(v5, fi9)
        v2 -= v2
        var v10 int64 = (v8 & (v8 % ((int64(v2)) | 1)))
        _ = fi9
        _ = v10
    }
    v8 = (v8 / (((v8 / (((v8 << ((521) & 63))) | 1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
