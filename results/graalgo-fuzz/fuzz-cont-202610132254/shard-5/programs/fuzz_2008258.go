package main

import (
    "fmt"
)

func genFn0(p0 uint8) uint8 {
    var v5 bool = (0 != -128)
    var v6 uint16 = 783
    var v7 int64 = 579
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return (^(p0))
}
func genFn1() uint32 {
    var v8 int16 = -1
    var v9 uint8 = 0
    var v10 int = 1
    _ = v10
    _ = v8
    _ = v9
    return 2147483648
}

func main() {
    var v1 float32 = 0.0
    v2 := []uint32{2147483648, 2}
    v3 := make(map[uint32]uint32)
    var v4 bool = (176 != -51)
    if true {
        var v11 int16 = 1
        v11 += (-(v11))
        var v12 int = 568
        _ = v11
        _ = v12
    } else {
        var v13 int32 = -2147483648
        var v14 int64 = 763
        v1 -= v1
        _ = v13
        _ = v14
    }
    if (1e10 == 1.5) {
        v1 += (v1 * (v1 * v1))
        if (124 != -9223372036854775808) {
            fmt.Println(v4, v4)
            var v15 uint8 = 128
            var v16 bool = (v4 || (1 < 42))
            _ = v15
            _ = v16
        }
    } else {
        v4 = v4
        var v17 uint32 = 757
        _ = v17
    }
    var v18 uint64 = 849
    v19 := v3[2147483647]
    v1 -= -2.25
    if v4 {
        var v20 int32 = 2147483647
        _ = v20
    } else {
        v19 = 734
        fmt.Println(v19)
    }
    var v21 uint8 = 32
    var v22 int64 = 42
    var v23 float32 = float32(1.5)
    if len(v2) > 0 {
        v2[0] = uint32(v21)
    }
    v2 = append(v2, v19)
    var v24 uint32 = (v19 | v19)
    fmt.Println(v23)
    v1 -= (v1 - 3.14159)
    fmt.Println(v23, v1, v1)
    var v25 int32 = int32(v22)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
