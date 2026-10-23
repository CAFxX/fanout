package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
func genFn0(p0 string, p1 uint8, p2 uint8) int16 {
    var v6 string = (((p0 + p0) + "a") + (p0 + (p0 + p0)))
    p0 = (("\n\t" + p0) + v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 32767
}

func main() {
    var v1 int = -1
    var v2 int64 = -9223372036854775808
    var v3 bool = (42 < 0)
    var v4 int32 = -878
    var v5 uint32 = 722
    switch v1 {
    case -923:
        var v7 string = "\n\t"
        _ = v7
    case 9223372036854775807:
        v8 := []int{(v1 / (((v1 >> ((4294967295) & 63))) | 1)), 87, v1, (v1 / ((v1) | 1))}
        _ = v8
    case -283:
        var v9 bool = true
        _ = v9
    }
    v5 = 0
    v10 := GenS1{F0: 282, F1: 65535}
    v11 := []int{v1, -601, v1, v1}
    if len(v11) > 0 {
        v11[0] = (v1 * (v1 % (((v1 & (v1 ^ v1))) | 1)))
    }
    var v12 float64 = (float64(float32(0.0)) - (0.0 / float64(3.14159)))
    if (v1 <= v1) {
        fmt.Println(v4, v1, v12)
        var v13 float32 = 3.14159
        fmt.Println(v5)
        _ = v13
    }
    var v14 uint8 = 121
    v12 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
