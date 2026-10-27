package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int32 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 float32 = 0.1
    var v2 int32 = 1
    v3 := make(map[uint32]bool)
    for fi4 := 0; fi4 < 12; fi4++ {
        var v5 int32 = v2
        _ = fi4
        _ = v5
    }
    v6 := []int{-130, -9223372036854775808, -473, 2147483647}
    v2 = 42
    if false {
        v2 = (v2 | v2)
        if len(v6) > 0 {
            v6[0] = 347
        }
    } else {
        v7 := GenS1{F0: 1, F1: 4294967295, F2: v2}
        var v8 uint32 = 123
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        if false {
            v1 = v1
            v10 := GenS2{F0: 1}
            _ = v10
        }
        for fi11 := 0; fi11 < 3; fi11++ {
            v1 -= v1
            fmt.Println(v1, fi9)
            _ = fi11
        }
        var v12 int = (fi9 / ((128) | 1))
        _ = fi9
        _ = v12
    }
    fmt.Println(v2, v1)
    for fi13 := 0; fi13 < 8; fi13++ {
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 uint16 = 2
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    v1 = ((float32(1e10) - 1.5) - v1)
    fmt.Println(v1)
    if (true && (2 <= 42)) {
        if ((100 == 2) && false) {
            v16 := GenS1{F0: 1, F1: 290, F2: v2}
            _ = v16
        }
    }
    v17 := v3[141]
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v6
}
