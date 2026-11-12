package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int32 }
type GenS2 struct { F0 int32; F1 int16; F2 uint }
type GenS3 struct { F0 uint32; F1 int; F2 int }
func genFn0() int32 {
    v4 := GenS1{F0: 0, F1: -32768, F2: 289}
    var v5 uint = 4294967295
    _ = v4
    _ = v5
    return -289
}

func main() {
    var v1 string = "hello"
    v2 := []string{v1, v1}
    v3 := []uint32{}
    for fi6 := 0; fi6 < 14; fi6++ {
        v1 = v1
        v1 = (v1 + (v1 + v1))
        var v7 uint = 2
        _ = fi6
        _ = v7
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            v1 = v1
            _ = fi9
        }
        fmt.Println(v1)
        _ = fi8
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 uint16 = 32767
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
        _ = fi10
        _ = v11
    }
    if (!((!(42 < 2)) || false)) {
        v12 := GenS2{F0: 42, F1: 528, F2: 65535}
        var v13 int64 = 1
        _ = v12
        _ = v13
    } else {
        var v14 uint64 = 2
        _ = v14
    }
    fmt.Println(v1)
    var v15 float64 = (0.1 * -2.25)
    if (2 != 18446744073709551615) {
        v15 -= v15
        v15 = 1e10
    }
    switch "" {
    case "a":
        v16 := GenS2{F0: 1, F1: 1, F2: 65535}
        _ = v16
    case "foo bar":
        v17 := GenS3{F0: 2147483648, F1: 0, F2: -9223372036854775808}
        _ = v17
    default:
        fmt.Println(v1, v1)
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
}
