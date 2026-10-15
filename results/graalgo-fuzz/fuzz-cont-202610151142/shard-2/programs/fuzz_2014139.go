package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
func genFn0(p0 int16, p1 int8) uint32 {
    var v4 int16 = p0
    var v5 bool = (((float32((0.0 * 1e10)) / (1.5 + 1e10)) * 3.14159) > -2.25)
    v5 = v5
    v6 := GenS1{F0: 9223372036854775807, F1: 1}
    v7 := []int{int(p0), 599}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 719
    var v3 float32 = (-2.25 - (0.0 / 1e10))
    v2 += v2
    var v8 float64 = 0.1
    for fi9 := 0; fi9 < 6; fi9++ {
        v8 += ((-2.25 - 0.1) + (((float64(1e10) + (v8 - v8)) - v8) / 3.14159))
        if (4294967295 == 255) {
            v10 := GenS1{F0: v2, F1: 4294967295}
            _ = v10
        }
        var v11 uint16 = 32767
        _ = fi9
        _ = v11
    }
    v12 := GenS1{F0: v2, F1: 65535}
    fmt.Println(v2)
    var v13 int64 = v2
    v14 := []int{}
    fmt.Println(v8, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v8
}
