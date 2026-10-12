package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS5 struct { F0 int16; F1 uint64 }

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = 0
    var v3 int32 = 2147483647
    var v4 int64 = -9223372036854775808
    v5 := make(map[uint32]uint32)
    var v6 uint32 = 0
    if true {
        v7 := GenS1{F0: 32767, F1: 0, F2: "foo bar"}
        var v8 uint64 = 4294967295
        if false {
            if true {
                v9 := v1[-9223372036854775808]
                _ = v9
            }
        } else {
            v10 := GenS2{F0: 127}
            _ = v10
        }
        _ = v7
        _ = v8
    }
    v11 := []int{32767, 606, 676, 2147483647}
    v5[2147483647] = v6
    var v12 bool = true
    for fi13 := 0; fi13 < 3; fi13++ {
        for fi14 := 0; fi14 < 11; fi14++ {
            v4 = (v4 / (((v4 << ((692) & 63))) | 1))
            fmt.Println(v3, v3)
            var v15 int64 = (v4 | (^(v4)))
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    if v12 {
        v16 := GenS2{F0: 255}
        v2 = (v2 & v2)
        var v17 uint8 = 206
        _ = v16
        _ = v17
    }
    switch "a" {
    case "a":
        _ = len(v11)
    case "":
        var v18 int8 = -113
        _ = v18
    default:
        v4 -= v4
    }
    var v19 uint8 = 255
    if (("" + "\n\t") != "\n\t") {
        if ((0 > 32767) && v12) {
            v20 := GenS3{F0: 1, F1: 418}
            _ = v20
        }
        var v21 int64 = v4
        _ = v21
    } else {
        var v22 int16 = v2
        _ = len(v11)
        v23 := GenS2{F0: 255}
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 2; fi24++ {
        v25 := GenS4{F0: v6, F1: v6, F2: (v3 / (((^(v3))) | 1))}
        _ = fi24
        _ = v25
    }
    if (false || v12) {
        fmt.Println(v19, v3, v4)
        v26 := GenS5{F0: v2, F1: 42}
        _ = v26
    } else {
        var v27 uint64 = 932
        _ = v27
    }
    v28 := v1[-1]
    switch 0 {
    case -9223372036854775808:
        v6 = v6
    }
    var v29 uint32 = v6
    if false {
        var v30 uint16 = 2
        if v28 {
            if false {
                v3 = v3
                var v31 int64 = -441
                _ = v31
            }
        }
        v19 += v19
        _ = v30
    }
    var v32 int16 = (v2 >> ((718) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
