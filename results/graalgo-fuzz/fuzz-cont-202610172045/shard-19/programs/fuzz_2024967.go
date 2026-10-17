package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint64) int16 {
    var v5 int16 = -99
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            v5 += (v5 + v5)
            _ = fi7
        }
        _ = fi6
    }
    p0 = (^(p0))
    if ((637 <= 42) || false) {
        var v8 uint16 = 0
        var v9 int = -481
        var v10 bool = false
        _ = v10
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v5
    return v5
}
func genFn1(p0 uint32, p1 int16, p2 string) int16 {
    p1 += genFn0(0)
    var v11 uint8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 0
}

func main() {
    var v1 int = 4294967295
    v2 := []uint32{1, 950, 65535, 4294967295}
    var v3 int32 = -1
    v4 := []int{(v1 << ((411) & 63))}
    v12 := GenS1{F0: -268}
    if len(v4) > 0 {
        v4[0] = v1
    }
    v3 = 2147483647
    var v13 uint = 4294967295
    var v14 uint16 = 1
    v13 += 42
    v4 = append(v4, v1)
    var v15 float64 = 0.1
    v16 := GenS1{F0: 1}
    fmt.Println(v1)
    var v17 uint64 = 585
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
