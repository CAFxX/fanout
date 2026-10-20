package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int }
type GenS2 struct { F0 uint64; F1 int; F2 int32 }
func genFn0(p0 int) string {
    if (751 >= -32768) {
        fmt.Println(p0)
        switch p0 {
        case 1:
            p0 = (-(p0))
        case 633:
            p0 = (p0 - (p0 << ((326) & 63)))
        case 128:
            v7 := []int{}
            _ = v7
        default:
            p0 -= -607
        }
    } else {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 int64 = 1
            p0 -= p0
            _ = fi8
            _ = v9
        }
        if (32767 >= -322) {
            for fi10 := 0; fi10 < 19; fi10++ {
                fmt.Println(fi10, fi10)
                var v11 int64 = -978
                _ = fi10
                _ = v11
            }
        } else {
            v12 := []int{1}
            var v13 int = p0
            _ = v12
            _ = v13
        }
        v14 := GenS1{F0: 42, F1: 2147483647, F2: p0}
        _ = v14
    }
    var v15 int = (-(p0))
    var v16 uint8 = 1
    v16 += v16
    _ = p0
    _ = v15
    _ = v16
    return "a"
}

func main() {
    var v1 bool = true
    var v2 uint16 = 1
    v3 := make(map[string]uint32)
    v4 := []uint32{uint32(v2), 2, 1}
    var v5 uint16 = v2
    var v6 uint16 = (^(v5))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v17 uint8 = 192
    v6 -= v2
    switch 0 {
    case -9223372036854775808:
        v17 = (v17 % ((1) | 1))
    }
    var v18 string = "hello"
    if false {
        fmt.Println(v5, v18)
        var v19 int64 = -656
        if (true && v1) {
            v6 += v6
            v20 := GenS2{F0: 2, F1: -911, F2: -567}
            _ = v20
        }
        _ = v19
    } else {
        v18 = "hello"
    }
    v5 = v2
    var v21 uint64 = 634
    for fi22 := 0; fi22 < 5; fi22++ {
        if v1 {
            var v23 bool = false
            var v24 int64 = -126
            _ = v23
            _ = v24
        }
        _ = fi22
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
