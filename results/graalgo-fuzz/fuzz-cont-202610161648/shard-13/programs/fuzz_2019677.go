package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int16 }
type GenS2 struct { F0 bool }
func genFn0(p0 int16) uint8 {
    p0 = -32768
    p0 = (^(p0))
    var v5 uint32 = 2147483647
    switch 128 {
    case 299:
        v6 := []int{-1, 1, -2147483648}
        _ = v6
    case 32768:
        fmt.Println(v5, v5)
    case 127:
        p0 = (^(p0))
    }
    _ = p0
    _ = v5
    return 43
}

func main() {
    var v1 uint = 65535
    var v2 uint64 = 2
    var v3 uint64 = (v2 & (v2 & 0))
    v4 := make(map[int]uint32)
    switch -9 {
    case 32767:
        fmt.Println(v2)
    case 127:
        if true {
            delete(v4, 154)
            v1 += (v1 + (^(v1)))
        } else {
            var v7 uint = 2
            for fi8 := 0; fi8 < 20; fi8++ {
                fmt.Println(v2, v1, v3)
                _ = fi8
            }
            _ = v7
        }
    }
    v3 = v2
    switch 9223372036854775807 {
    case -86:
        v9 := []int{957, 42}
        _ = v9
    default:
        v3 -= (v2 >> ((65535) & 63))
    }
    var v10 uint = v1
    if false {
        if true {
            if (("foo bar" < "hello") && false) {
                if (!(true && true)) {
                    v11 := GenS1{F0: (v1 | v1), F1: 0, F2: 1000}
                    _ = v11
                }
            }
        } else {
            var v12 int16 = 211
            var v13 uint64 = (v2 / (((^(v3))) | 1))
            var v14 int8 = -1
            _ = v12
            _ = v13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        _ = len(v4)
        _ = fi15
    }
    v2 -= v2
    v2 = (^(v2))
    if (!(0.0 >= 0.0)) {
        v16 := GenS2{F0: false}
        _ = v16
    } else {
        v2 = (^(v3))
        var v17 int64 = -290
        var v18 float32 = 1.5
        _ = v17
        _ = v18
    }
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
}
