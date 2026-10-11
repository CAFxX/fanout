package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int16 }
func genFn0(p0 bool) int32 {
    fmt.Println(p0)
    p0 = p0
    if p0 {
        v6 := GenS1{F0: 241, F1: 0, F2: -1}
        var v7 int = -2147483648
        _ = v6
        _ = v7
    } else {
        var v8 int64 = -1
        _ = v8
    }
    _ = p0
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 0.1
    var v3 uint8 = 121
    v4 := []uint32{423, 679, 2}
    var v5 int32 = 2147483647
    var v9 float64 = v2
    var v10 uint16 = 533
    for fi11 := 0; fi11 < 16; fi11++ {
        v12 := GenS1{F0: 42, F1: 0, F2: 0}
        _ = fi11
        _ = v12
    }
    if len(v4) > 0 {
        v4[0] = 937
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        v10 = (^(v10))
        var v14 int16 = 0
        _ = fi13
        _ = v14
    }
    var v15 uint16 = v10
    var v16 bool = false
    var v17 int64 = 9223372036854775807
    if len(v4) > 0 {
        v4[0] = 0
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        v15 -= (^(v10))
        fmt.Println(v5, v16, v10)
        var v19 uint16 = v15
        _ = fi18
        _ = v19
    }
    v16 = v16
    var v20 int64 = v17
    var v21 uint16 = (^(v10))
    v22 := GenS1{F0: uint64(v15), F1: 127, F2: -1}
    v15 += v21
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
