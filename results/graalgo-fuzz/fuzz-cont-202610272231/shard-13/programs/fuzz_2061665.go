package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int8 }

func main() {
    var v1 uint = 255
    v2 := make(map[string]uint32)
    var v3 uint = (v1 * v1)
    v4 := make(map[uint32]int)
    var v5 bool = (false && true)
    v3 = v1
    var v6 bool = v5
    v3 += (v1 % ((42) | 1))
    v7 := []int{-1, 0, 483, -1}
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int8 = 127
        var v10 string = "hello"
        v5 = ((v6 && v6) || true)
        _ = fi8
        _ = v10
        _ = v9
    }
    v1 = (v3 - v3)
    var v11 int64 = 0
    if false {
        var v12 uint = v3
        for fi13 := 0; fi13 < 19; fi13++ {
            fmt.Println(v12)
            _ = fi13
        }
        _ = v12
    } else {
        var v14 uint = (v3 / (((v1 << (((v3 ^ v1)) & 63))) | 1))
        var v15 uint16 = 255
        var v16 uint8 = 253
        _ = v14
        _ = v15
        _ = v16
    }
    if v5 {
        switch ("\n\t" + "a") {
        case "foo bar":
            v17 := GenS1{F0: v11, F1: v11, F2: -1}
            _ = v17
        default:
            var v18 int16 = 1000
            _ = v18
        }
    } else {
        if (1000 != 1) {
            fmt.Println(v1, v5)
            var v19 int64 = 568
            _ = v19
        } else {
            v20 := GenS1{F0: v11, F1: v11, F2: 1}
            _ = v20
        }
    }
    if v5 {
        fmt.Println(v5, v5)
        if len(v7) > 0 {
            v7[0] = 327
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
