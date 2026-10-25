package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint8 }
type GenS2 struct { F0 int64; F1 int16 }

func main() {
    var v1 float64 = 1e10
    var v2 uint64 = 4294967295
    var v3 uint64 = 418
    var v4 int8 = -1
    var v5 float64 = v1
    var v6 uint8 = 1
    for fi7 := 0; fi7 < 2; fi7++ {
        fmt.Println(v6)
        fmt.Println(v4)
        _ = fi7
    }
    var v8 string = "\n\t"
    fmt.Println(v8)
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 uint32 = 2147483647
        fmt.Println(fi9, v8)
        var v11 string = v8
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: 4294967295, F1: "foo bar", F2: uint8(v2)}
    var v13 int16 = 244
    if false {
        if true {
            v14 := []int{-424, 517, -1}
            v4 -= (v4 >> ((0) & 63))
            _ = v14
        } else {
            v4 -= -128
            v4 -= (v4 * (^(v4)))
        }
        if (2 <= 65535) {
            var v15 float32 = 1e10
            v15 = (0.0 * 0.0)
            _ = v15
        }
        fmt.Println(v13, v4)
    } else {
        v16 := GenS2{F0: -961, F1: v13}
        _ = v16
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        var v18 uint64 = (v3 - v2)
        for fi19 := 0; fi19 < 15; fi19++ {
            var v20 int16 = v13
            _ = fi19
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    v21 := []int{32767, 817}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
