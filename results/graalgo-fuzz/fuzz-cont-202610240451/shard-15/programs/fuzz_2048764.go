package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS5 struct { F0 int }

func main() {
    v1 := []uint32{0}
    var v2 int16 = -32768
    var v3 uint64 = 4294967295
    var v4 bool = (0 > 4294967295)
    v3 = v3
    v5 := GenS1{F0: -2147483648, F1: 0}
    if (1e10 < float64(float32(0.1))) {
        v4 = v4
        v1 = append(v1, 2)
        v6 := GenS2{F0: 32768, F1: 255, F2: 42}
        _ = v6
    }
    v2 = v2
    var v7 uint64 = 42
    v7 = 716
    if false {
        var v8 uint16 = 1
        _ = v8
    }
    switch 2147483647 {
    case -9223372036854775808:
        v9 := GenS3{F0: 65535, F1: (("foo bar" + "hello") + "foo bar")}
        _ = v9
    case 42:
        fmt.Println(v2, v4)
    case -540:
        for fi10 := 0; fi10 < 14; fi10++ {
            fmt.Println(fi10)
            v11 := GenS4{F0: 517, F1: 4294967295, F2: 1}
            _ = fi10
            _ = v11
        }
    }
    var v12 uint64 = (^(v7))
    switch 4294967295 {
    case 32768:
        var v13 uint16 = 602
        _ = v13
    case -9223372036854775808:
        v3 -= (^(v12))
    case 0:
        v12 = (v3 * (^(v12)))
    }
    v14 := GenS5{F0: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
