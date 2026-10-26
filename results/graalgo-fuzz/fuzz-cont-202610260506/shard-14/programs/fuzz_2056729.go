package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int }
func genFn0(p0 uint32) uint16 {
    p0 -= (^(p0))
    fmt.Println(p0)
    switch "hello" {
    case "":
        fmt.Println(p0)
    case "foo bar":
        var v4 uint64 = 18446744073709551615
        _ = v4
    case "a":
        fmt.Println(p0)
    }
    v5 := []int{492, 32767, -9223372036854775808}
    v6 := GenS1{F0: 251, F1: 0}
    _ = p0
    _ = v5
    _ = v6
    return 65535
}

func main() {
    var v1 uint16 = 2
    v2 := []uint32{42, 1, 4294967295}
    var v3 float64 = -2.25
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v7 float32 = 1e10
    var v8 uint = 324
    v3 = (float64(v7) + (0.1 + v3))
    v1 = v1
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    var v9 string = "a"
    var v10 string = v9
    var v11 uint16 = 42
    var v12 bool = (-1 < 0)
    fmt.Println(v8, v10)
    v13 := GenS2{F0: -2147483648}
    var v14 uint8 = 255
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v15 uint32 = 518
    v15 -= (v15 << (((v8 & v8)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
