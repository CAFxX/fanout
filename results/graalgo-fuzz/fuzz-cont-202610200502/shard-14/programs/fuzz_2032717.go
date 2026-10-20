package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[string]uint32)
    var v3 uint64 = 928
    var v4 int16 = -291
    var v5 bool = false
    var v6 int8 = -82
    if (2 >= 821) {
        var v7 float32 = 1e10
        if v5 {
            var v8 int = 1
            v8 -= (v8 - (-(v8)))
            _ = v8
        }
        _ = v7
    } else {
        var v9 int32 = -225
        v1 -= 42
        _ = v9
    }
    fmt.Println(v4)
    fmt.Println(v4, v5)
    var v10 uint32 = (^(v1))
    if v5 {
        var v11 uint64 = 2
        _ = v11
    }
    var v12 uint8 = 127
    switch "\n\t" {
    case "\n\t":
        if (0 < uint32(v12)) {
            fmt.Println(v10, v3, v1)
        } else {
            if v5 {
                v4 = v4
            }
            for fi13 := 0; fi13 < 4; fi13++ {
                v14 := GenS1{F0: (^(v4)), F1: -2147483648}
                _ = fi13
                _ = v14
            }
        }
    default:
        v4 = 32767
    }
    if (!(32767 >= -1)) {
        var v15 uint8 = (v12 ^ 128)
        v2["foo bar"] = 935
        switch "\n\t" {
        case "":
            v3 = 4294967295
        }
        _ = v15
    }
    v16 := GenS1{F0: (v4 & 32767), F1: -168}
    v17 := []int{}
    switch "hello" {
    case "\n\t":
        var v18 uint = uint(v3)
        _ = v18
    default:
        if true {
            v6 -= 100
        } else {
            if true {
                v10 += (v1 << ((1) & 63))
                var v19 uint16 = 42
                _ = v19
            } else {
                v5 = true
                var v20 uint8 = v12
                _ = v20
            }
        }
    }
    if v5 {
        if (false || ((-1 >= -2147483648) && v5)) {
            fmt.Println(v6, v1)
        }
    } else {
        var v21 int16 = v4
        v22 := GenS1{F0: v21, F1: 247}
        var v23 uint16 = 472
        _ = v21
        _ = v22
        _ = v23
    }
    _ = len(v2)
    fmt.Println(v1)
    var v24 float32 = -2.25
    for fi25 := 0; fi25 < 13; fi25++ {
        var v26 bool = ((v10 / (((v1 ^ 2)) | 1)) == (v1 % (((v10 << ((255) & 63))) | 1)))
        var v27 int32 = 0
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
