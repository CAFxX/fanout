package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 bool }

func main() {
    var v1 uint = 336
    v2 := make(map[int]int)
    v3 := []uint32{}
    var v4 int32 = 2147483647
    var v5 int64 = 741
    switch int(v4) {
    case 2147483647:
        if (206 != 132) {
            var v6 uint32 = 62
            _ = v6
        }
    default:
        for fi7 := 0; fi7 < 17; fi7++ {
            fmt.Println(fi7, v5, v5)
            var v8 float32 = 0.0
            if (v8 == v8) {
                var v9 string = "a"
                v1 += (v1 + v1)
                _ = v9
            }
            _ = fi7
            _ = v8
        }
    }
    v4 = 203
    v10 := GenS1{F0: -2147483648, F1: -2147483648}
    v11 := GenS1{F0: 127, F1: -221}
    fmt.Println(v1, v5, v1)
    for fi12 := 0; fi12 < 6; fi12++ {
        fmt.Println(v5)
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 uint = v1
            _ = fi13
            _ = v14
        }
        fmt.Println(v5, fi12, v4)
        _ = fi12
    }
    if (!true) {
        var v15 int8 = 0
        _ = v15
    }
    var v16 uint64 = 4294967295
    var v17 int = 876
    if false {
        v18 := GenS2{F0: 1}
        _ = v18
    } else {
        v19 := GenS3{F0: (^(v16)), F1: false}
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
