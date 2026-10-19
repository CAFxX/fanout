package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint }
func genFn0(p0 uint16) uint8 {
    var v5 int = 128
    var v6 int64 = 0
    _ = p0
    _ = v5
    _ = v6
    return 42
}
func genFn1() int64 {
    var v7 int32 = 420
    var v8 uint = 18446744073709551615
    switch 1 {
    case 2147483647:
        v9 := []int{128, 127}
        _ = v9
    }
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[int]string)
    var v3 int = -671
    var v4 float64 = -2.25
    v10 := GenS1{F0: -1, F1: 1, F2: 1}
    v3 -= v3
    switch 127 {
    case 4294967295:
        var v11 int32 = 181
        _ = v11
    case 9223372036854775807:
        v4 = 1e10
    case -890:
        v12 := GenS1{F0: 1, F1: 9223372036854775807, F2: 152}
        _ = v12
    }
    fmt.Println(v4, v1, v3)
    v13 := GenS1{F0: 133, F1: -189, F2: 830}
    v4 -= 1e10
    var v14 uint8 = 0
    v4 = float64(0.0)
    v3 += v3
    v15 := v2[(-(v3))]
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
