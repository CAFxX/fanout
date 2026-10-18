package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int8; F1 int }
func genFn0(p0 uint64) bool {
    fmt.Println(p0)
    var v5 float64 = 1.5
    _ = p0
    _ = v5
    return true
}

func main() {
    var v1 uint64 = 828
    var v2 uint8 = 42
    var v3 uint32 = 42
    var v4 string = ""
    switch "" {
    case "hello":
        v2 -= (v2 % ((v2) | 1))
    case "a":
        var v6 uint8 = (^(v2))
        _ = v6
    default:
        for fi7 := 0; fi7 < 8; fi7++ {
            fmt.Println(v2, v3, v3)
            fmt.Println(v2, v4, v4)
            v1 = 18446744073709551615
            _ = fi7
        }
    }
    v8 := GenS1{F0: 0, F1: (v1 % (((^(v1))) | 1))}
    v1 = 2
    switch ((v4 + v4) + v4) {
    case "foo bar":
        v9 := GenS2{F0: 42, F1: 2147483647}
        _ = v9
    case "hello":
        fmt.Println(v3, v4)
    }
    v10 := []int{4294967295}
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 uint32 = 4294967295
        _ = fi11
        _ = v12
    }
    v13 := GenS3{F0: -128, F1: 32768}
    v4 = v4
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 int8 = 0
        _ = fi14
        _ = v15
    }
    switch int(v3) {
    case 2147483647:
        for fi16 := 0; fi16 < 8; fi16++ {
            v2 += v2
            _ = fi16
        }
    default:
        v2 -= v2
    }
    switch v4 {
    case "a":
        var v17 uint8 = v2
        _ = v17
    default:
        _ = len(v10)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
