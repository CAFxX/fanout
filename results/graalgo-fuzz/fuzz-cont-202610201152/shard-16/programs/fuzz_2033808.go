package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 bool; F1 uint64 }

func main() {
    var v1 int32 = 1
    var v2 uint32 = 352
    var v3 int16 = -730
    v4 := []uint32{(v2 & v2), v2, 381, (v2 - 65535)}
    v5 := make(map[int]uint32)
    if (622 != 9223372036854775807) {
        fmt.Println(v2)
        _ = len(v5)
    } else {
        if false {
            fmt.Println(v1, v3)
        }
        var v6 int8 = 100
        _ = v6
    }
    var v7 int = -33
    var v8 int = int(v3)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        if (false && true) {
            var v9 uint32 = v2
            _ = v9
        }
        v1 = (-(v1))
        switch v7 {
        case 1:
            v10 := GenS1{F0: 0, F1: 747}
            _ = v10
        case -532:
            v11 := GenS1{F0: -9223372036854775808, F1: 1}
            _ = v11
        case -9223372036854775808:
            fmt.Println(v1, v7, v3)
        default:
            v4 = append(v4, (^(v2)))
        }
    }
    v2 -= (v2 + v2)
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := GenS1{F0: -1, F1: -9223372036854775808}
        _ = fi12
        _ = v13
    }
    v1 -= v1
    v14 := GenS2{F0: false, F1: 2}
    var v15 float32 = 1e10
    v16 := v5[v8]
    fmt.Println(v1, v8)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
