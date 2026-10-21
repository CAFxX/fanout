package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 bool }
func genFn0() int32 {
    v5 := []int{-262}
    var v6 uint32 = 2147483647
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint, p1 bool) string {
    v7 := GenS1{F0: 126, F1: 32767}
    if (false || (p1 || (292 > 9223372036854775807))) {
        p0 = (p0 % ((657) | 1))
        var v8 int32 = 2147483647
        _ = v8
    }
    p0 = 2
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v7
    return "a"
}

func main() {
    v1 := []int{-813, 128, 0, 128}
    var v2 bool = false
    var v3 int32 = -2147483648
    var v4 float32 = 0.0
    v2 = v2
    v9 := GenS1{F0: 0, F1: 2}
    var v10 uint = 2
    v3 += v3
    if v2 {
        v4 = (float32((1.5 - ((-2.25 - 0.1) - 0.0))) + (((v4 * v4) + 1.5) * -2.25))
    }
    v3 += (v3 << ((v10) & 63))
    v11 := GenS2{F0: true}
    v4 = float32((0.0 * ((1.5 / (3.14159 + -2.25)) * 1e10)))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (0.0 > v4) {
        v2 = v2
    }
    var v12 uint8 = 0
    var v13 uint16 = 2
    v10 = v10
    v1 = append(v1, 32767)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
