package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
func genFn0(p0 uint, p1 int32) uint8 {
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint64 = 2
        for fi8 := 0; fi8 < 13; fi8++ {
            v7 = (v7 >> (((p0 / ((p0) | 1))) & 63))
            v9 := []int{}
            var v10 uint32 = 2147483647
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v11 int = -2147483648
    v12 := []int{1, v11, -1, 128}
    v13 := GenS1{F0: (p0 >> (((p0 + 453)) & 63)), F1: 18446744073709551615}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return 1
}

func main() {
    var v1 int = -1
    v2 := make(map[string]bool)
    var v3 uint64 = 416
    var v4 int8 = 1
    var v5 int8 = v4
    _ = len(v2)
    var v14 int64 = 42
    for fi15 := 0; fi15 < 14; fi15++ {
        fmt.Println(v14)
        v5 = v5
        var v16 uint64 = (v3 / ((4294967295) | 1))
        _ = fi15
        _ = v16
    }
    v17 := GenS1{F0: 42, F1: v3}
    v3 += 4294967295
    var v18 int = v1
    fmt.Println(v1, v4)
    v5 = v5
    var v19 float64 = float64(float32(-2.25))
    var v20 bool = true
    v21 := GenS1{F0: 4294967295, F1: (^(v3))}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
