package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int8 }
func genFn0(p0 uint16, p1 uint8, p2 uint32) int32 {
    v6 := GenS1{F0: 65535, F1: 4294967295, F2: -34}
    fmt.Println(p2, p2, p0)
    p1 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2147483647
}

func main() {
    var v1 int64 = 111
    var v2 string = "foo bar"
    var v3 float64 = 1.5
    var v4 float32 = (0.1 * 1.5)
    var v5 int8 = 1
    v7 := []int{862}
    switch ((((v2 + v2) + (v2 + v2)) + "hello") + v2) {
    case "a":
        _ = len(v7)
    default:
        fmt.Println(v2, v1)
    }
    var v8 uint = 65535
    v5 -= v5
    var v9 int32 = 240
    v8 = v8
    var v10 string = v2
    var v11 int64 = (-(v1))
    if (1 <= 2147483647) {
        var v12 uint = v8
        _ = v12
    }
    v4 -= (((v4 - (float32(3.14159) * (float32((float64(v4) - v3)) * ((v4 - v4) + v4)))) / float32((v3 + v3))) + ((3.14159 + (v4 * float32(((float64(v4) - v3) * v3)))) - v4))
    var v13 bool = false
    v14 := GenS1{F0: 65535, F1: 0, F2: (^(v5))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
