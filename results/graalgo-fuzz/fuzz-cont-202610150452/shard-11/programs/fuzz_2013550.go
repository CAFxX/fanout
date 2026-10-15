package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 int32 }
func genFn0(p0 string) uint {
    var v7 int32 = 0
    var v8 int8 = -1
    var v9 uint64 = 1
    p0 = p0
    v10 := GenS1{F0: 9223372036854775807}
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return uint(v9)
}

func main() {
    v1 := []uint32{2147483647, 102, 2147483648}
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    v4 := make(map[int]string)
    v5 := make(map[string]int)
    v6 := make(map[int]bool)
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 int32 = -273
        var v13 uint64 = 586
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS2{F0: (139 < 1)}
    var v15 float32 = 3.14159
    v16 := v3["foo bar"]
    var v17 int8 = -1
    v16 = v16
    if (-1 > 1) {
        var v18 float64 = 1.5
        _ = v18
    } else {
        v15 = (1.5 + v15)
        if (!true) {
            v17 -= -44
            for fi19 := 0; fi19 < 10; fi19++ {
                v17 = (v17 << ((2) & 63))
                var v20 int32 = 0
                _ = fi19
                _ = v20
            }
            v21 := GenS3{F0: genFn0("")}
            _ = v21
        }
    }
    if (!true) {
        v22 := GenS4{F0: false, F1: 346}
        fmt.Println(v17, v15)
        v15 += v15
        _ = v22
    } else {
        v16 = (v16 + (v16 << ((487) & 63)))
    }
    var v23 int8 = (v17 << ((2) & 63))
    v16 -= (v16 << ((18446744073709551615) & 63))
    v16 -= v16
    v16 += (^(v16))
    v24 := GenS5{F0: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
