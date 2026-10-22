package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint16 }
func genFn0(p0 uint8) int8 {
    p0 = p0
    p0 += p0
    p0 += 2
    v5 := []int{42, 42}
    _ = p0
    _ = v5
    return 82
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = 1e10
    var v3 int16 = -184
    v4 := []uint32{141, 1, 209, 2147483648}
    var v6 float32 = v2
    if len(v4) > 0 {
        v4[0] = 2
    }
    v2 = ((v6 * v2) * 1e10)
    if false {
        v6 -= (v2 * v6)
        var v7 string = "hello"
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 uint = 255
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    fmt.Println(v2, v6, v6)
    if len(v4) > 0 {
        v4[0] = 307
    }
    var v10 float32 = v6
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 uint = 0
        fmt.Println(v12, v3)
        _ = fi11
        _ = v12
    }
    var v13 string = ("" + "a")
    v14 := GenS1{F0: 118, F1: (v3 >> ((0) & 63)), F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
