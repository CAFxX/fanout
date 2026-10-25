package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 string; F2 uint }
func genFn0(p0 int32) uint64 {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 5; fi4++ {
        p0 -= p0
        var v5 uint8 = 255
        var v6 int8 = 100
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: -9223372036854775808}
    _ = p0
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 int8 = 1
    var v2 int = -39
    var v3 int16 = 32767
    v8 := GenS2{F0: v1, F1: (("" + "") + ""), F2: 65535}
    v1 = v1
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := []int{v2, -1, 2147483647, fi9}
        _ = fi9
        _ = v10
    }
    v1 += v1
    v2 -= (v2 | 32767)
    switch "hello" {
    case "":
        var v11 int64 = 1
        _ = v11
    }
    var v12 int16 = v3
    if ((0.0 * (-2.25 - (((1e10 - 1.5) + float32((0.1 + (-2.25 / -2.25)))) / 3.14159))) != (0.0 / float32(float64(3.14159)))) {
        fmt.Println(v12, v12, v2)
    }
    v13 := []int{v2, (v2 & (v2 + 128)), (-(v2)), 945}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v8
}
