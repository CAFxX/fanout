package main

import (
    "fmt"
)

func genFn0() int64 {
    var v6 int8 = -115
    v6 -= (v6 % (((v6 << ((4294967295) & 63))) | 1))
    _ = v6
    return -9223372036854775808
}
func genFn1(p0 int64, p1 uint64, p2 uint64) int {
    v7 := []int{32767, 42, 133, -401}
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 float32 = ((1e10 - (((1e10 / 3.14159) * 3.14159) * -2.25)) + 3.14159)
        var v10 int64 = (p0 + p0)
        p1 = 42
        _ = fi8
        _ = v10
        _ = v9
    }
    p2 -= p1
    var v11 bool = (2 < 42)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    return int(p0)
}

func main() {
    var v1 uint = 695
    v2 := []int{-1}
    v3 := make(map[int]int)
    var v4 uint8 = 127
    var v5 float64 = 1e10
    delete(v3, int(v1))
    v5 = 0.0
    var v12 bool = true
    fmt.Println(v5, v5)
    for fi13 := 0; fi13 < 4; fi13++ {
        v1 = 65535
        fmt.Println(fi13, v1, fi13)
        _ = fi13
    }
    var v14 float64 = v5
    if (((!(64 <= v1)) && v12) || (genFn1(-1, 1, 42) > -2147483648)) {
        v14 -= v14
    }
    v15 := v3[9223372036854775807]
    v12 = true
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
