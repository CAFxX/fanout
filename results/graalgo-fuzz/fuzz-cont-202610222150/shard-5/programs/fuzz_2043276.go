package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int8 }

func main() {
    v1 := []uint32{65535, 0, 2147483647, 409}
    v2 := make(map[string]bool)
    var v3 int = 2147483647
    var v4 uint64 = 1
    var v5 int = v3
    var v6 uint16 = 711
    _ = len(v2)
    var v7 int64 = -1
    if (1 > (v6 % ((v6) | 1))) {
        var v8 string = "foo bar"
        _ = v8
    } else {
        if (-1 == 0) {
            var v9 uint8 = 96
            fmt.Println(v3, v6)
            var v10 float32 = 1.5
            _ = v10
            _ = v9
        } else {
            v5 = -367
            var v11 float64 = (1.5 * 0.0)
            fmt.Println(v3, v7, v7)
            _ = v11
        }
        v7 -= 126
        v2["foo bar"] = false
    }
    v12 := GenS1{F0: 2147483647}
    if len(v1) > 0 {
        v1[0] = 0
    }
    if (!(!(1 <= 0))) {
        v7 = (v7 % ((v7) | 1))
        v7 = (v7 >> ((0) & 63))
        fmt.Println(v7)
    } else {
        switch "a" {
        case "foo bar":
            v4 -= 826
        case "a":
            v5 = (v3 << ((uint(v4)) & 63))
        case "hello":
            var v13 uint16 = v6
            _ = v13
        default:
            var v14 float32 = float32(1e10)
            _ = v14
        }
    }
    if true {
        var v15 int8 = 0
        _ = v15
    } else {
        v16 := GenS2{F0: -501, F1: int8(v4)}
        var v17 bool = true
        v6 += (v6 + v6)
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        var v19 uint = 65535
        _ = fi18
        _ = v19
    }
    v6 = v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
