package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint32 }
func genFn0() uint {
    for fi7 := 0; fi7 < 6; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    v8 := GenS1{F0: -9223372036854775808, F1: 783, F2: 2147483648}
    var v9 float32 = (float32(0.1) / 1.5)
    var v10 uint32 = 2147483648
    var v11 uint16 = 261
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint16 = 896
    var v2 int64 = 42
    v3 := make(map[uint32]int)
    v4 := make(map[string]string)
    var v5 float32 = float32(float64(float32(1.5)))
    v6 := make(map[uint32]string)
    v5 = (v5 / 1.0)
    var v12 int32 = -21
    if ((0.1 / (3.14159 + 1e10)) <= (1.5 + 1e10)) {
        for fi13 := 0; fi13 < 3; fi13++ {
            v14 := GenS1{F0: v2, F1: (-(v2)), F2: 253}
            _ = fi13
            _ = v14
        }
        var v15 uint32 = 2147483647
        _ = v15
    } else {
        v12 -= v12
        fmt.Println(v2, v12, v1)
        var v16 uint16 = 2
        _ = v16
    }
    var v17 uint32 = 2147483648
    v12 = (v12 * -1)
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 string = "\n\t"
        _ = fi18
        _ = v19
    }
    fmt.Println(v12, v12)
    for fi20 := 0; fi20 < 6; fi20++ {
        v4[""] = (("a" + "a") + ("" + "\n\t"))
        v1 = v1
        _ = fi20
    }
    v2 = v2
    v21 := []int{-967}
    v2 += (v2 << ((42) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
