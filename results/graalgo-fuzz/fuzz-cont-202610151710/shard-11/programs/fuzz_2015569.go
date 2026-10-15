package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint32; F1 int; F2 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint64, p1 int) bool {
    if false {
        if false {
            p0 = 454
            p0 = p0
            fmt.Println(p1, p1)
        } else {
            var v4 bool = (true || (!true))
            _ = v4
        }
        fmt.Println(p1, p1)
        for fi5 := 0; fi5 < 10; fi5++ {
            var v6 float32 = 3.14159
            v7 := GenS1{F0: p0, F1: 0}
            var v8 uint32 = 2147483647
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
    }
    v9 := GenS2{F0: 65535, F1: p1, F2: 145}
    _ = p0
    _ = p1
    _ = v9
    return (((758 <= -1) && false) && false)
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 128
    var v3 uint = 2
    if true {
        var v10 int64 = -352
        _ = v10
    } else {
        switch "hello" {
        case "":
            v11 := v1[uint32(v2)]
            _ = v11
        case "a":
            for fi12 := 0; fi12 < 15; fi12++ {
                v2 -= v2
                _ = fi12
            }
        case "hello":
            fmt.Println(v3)
        default:
            var v13 bool = true
            _ = v13
        }
        var v14 uint = v3
        _ = v14
    }
    v15 := []int{127, -1, -9223372036854775808, -60}
    var v16 uint16 = 0
    var v17 uint16 = 65535
    switch 32768 {
    case 0:
        var v18 int = -9223372036854775808
        _ = v18
    default:
        v2 -= (v2 & 127)
    }
    fmt.Println(v16, v16)
    for fi19 := 0; fi19 < 8; fi19++ {
        if false {
            var v20 int8 = -44
            _ = v20
        }
        _ = fi19
    }
    v21 := GenS3{F0: v3}
    v22 := GenS4{F0: v2}
    var v23 int16 = 907
    if true {
        var v24 float64 = ((3.14159 * 3.14159) * (3.14159 - 0.1))
        v23 = v23
        if true {
            v2 -= v2
        }
        _ = v24
    } else {
        v17 = (v17 + 32767)
        var v25 uint32 = 4294967295
        v16 += (v16 % ((v17) | 1))
        _ = v25
    }
    fmt.Println(v23)
    fmt.Println(v16, v17)
    v2 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
