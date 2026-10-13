package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0() int64 {
    var v6 uint16 = 733
    v7 := []int{32768, 127, 42, 2147483647}
    v6 = v6
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = (0.1 / 0.1)
    var v3 uint64 = 18446744073709551615
    v4 := []int{1, 4294967295, -9223372036854775808, 1}
    var v5 int = -654
    v5 = (v5 & (v5 << ((937) & 63)))
    fmt.Println(v2, v2)
    v3 -= (v3 | v3)
    v8 := GenS1{F0: 2147483648}
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 bool = (((0.0 + 1.5) - float64((float32(1.5) - (-2.25 + v2)))) == ((0.0 * 0.1) * 1.5))
        var v11 float32 = v2
        v2 -= -2.25
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 string = "\n\t"
    if false {
        var v13 uint8 = 127
        var v14 float32 = (0.1 / v2)
        _ = len(v1)
        _ = v13
        _ = v14
    } else {
        fmt.Println(v12)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) > 0 {
        v4[0] = (v5 ^ (v5 | v5))
    }
    var v15 uint64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
