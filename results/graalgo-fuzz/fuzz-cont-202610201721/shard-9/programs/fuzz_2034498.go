package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
func genFn0(p0 uint, p1 uint64, p2 string) bool {
    var v4 int64 = 1
    v5 := []int{32768, 254}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (!(((0.1 * (-2.25 * -2.25)) != 1e10) || true))
}

func main() {
    var v1 int64 = 1
    v2 := make(map[string]uint32)
    var v3 int16 = 0
    var v6 float32 = 1.5
    v7 := []int{127}
    var v8 uint32 = 0
    for fi9 := 0; fi9 < 8; fi9++ {
        if (2147483647 >= 0) {
            fmt.Println(v8, v6, v6)
            v7 = append(v7, fi9)
            v3 = 32767
        }
        _ = fi9
    }
    v10 := GenS1{F0: 4294967295, F1: 42}
    v1 -= -1
    v3 -= v3
    v6 = float32(float64(((((v6 + float32(((3.14159 / (3.14159 * ((float64(1.5) / 1.5) * -2.25))) * -2.25))) * (((v6 - v6) / 0.1) - v6)) / v6) / v6)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
