package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
func genFn0(p0 string) uint64 {
    v4 := []int{162, 128, 4294967295, -504}
    v5 := GenS1{F0: -1, F1: -976}
    p0 = (p0 + p0)
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 uint = 2
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: 0, F1: -32768}
        var v10 uint64 = 0
        _ = fi6
        _ = v10
        _ = v9
    }
    fmt.Println(p0)
    _ = p0
    _ = v4
    _ = v5
    return 0
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]bool)
    var v3 int64 = 1
    var v11 int32 = 2147483647
    for fi12 := 0; fi12 < 9; fi12++ {
        if true {
            v13 := []int{42, (fi12 ^ int(v11)), (-(fi12)), 128}
            _ = v13
        }
        _ = fi12
    }
    v3 -= v3
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 float32 = -2.25
        var v16 uint = 163
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := []int{128}
    v11 += (v11 / ((42) | 1))
    var v18 int8 = 127
    if len(v17) > 0 {
        v17[0] = 2147483647
    }
    _ = len(v1)
    if len(v17) > 0 {
        v17[0] = 2147483647
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
