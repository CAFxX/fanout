package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint8; F1 int64 }
func genFn0() int32 {
    var v6 int = 2147483647
    v7 := GenS1{F0: 42, F1: 1}
    for fi8 := 0; fi8 < 4; fi8++ {
        v6 = 128
        _ = fi8
    }
    v6 += v6
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int, p1 int, p2 uint32) int64 {
    var v9 uint8 = 255
    fmt.Println(p2, p2, p0)
    if true {
        var v10 bool = (1.5 <= ((0.0 - 3.14159) / 1.5))
        fmt.Println(p0, v9)
        _ = v10
    } else {
        fmt.Println(v9)
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 int16 = 42
        for fi13 := 0; fi13 < 4; fi13++ {
            v14 := []int{p0, -324, fi11}
            v15 := GenS1{F0: -549, F1: 9223372036854775807}
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 string = "\n\t"
        _ = fi11
        _ = v12
        _ = v16
    }
    switch p0 {
    case -135:
        p0 += p1
    case 0:
        var v17 uint32 = p2
        _ = v17
    case 475:
        var v18 int64 = 1
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -890
}

func main() {
    var v1 int = 128
    v2 := make(map[string]uint32)
    v3 := []uint32{0, 42, 974, 498}
    var v4 int64 = 0
    v5 := make(map[string]uint32)
    var v19 bool = (!(!(2147483647 <= 227)))
    v20 := GenS2{F0: 128, F1: (v4 << ((4294967295) & 63))}
    v19 = (true && (42 != 255))
    switch (v1 - -1) {
    case -586:
        v4 -= (v4 % ((v4) | 1))
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        var v22 uint8 = 1
        _ = fi21
        _ = v22
    }
    _ = len(v2)
    delete(v2, "a")
    var v23 uint32 = 276
    var v24 uint = 42
    if (!(v19 && (127 == v1))) {
        v25 := GenS1{F0: v4, F1: (v4 & (^(v4)))}
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
