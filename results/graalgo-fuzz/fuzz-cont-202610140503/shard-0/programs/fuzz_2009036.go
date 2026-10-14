package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int64 }
func genFn0(p0 uint32, p1 string) int16 {
    p0 = p0
    p1 = p1
    v7 := []int{}
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    var v1 int = 469
    v2 := make(map[uint32]uint32)
    var v3 bool = (42 > 888)
    var v4 int8 = 100
    var v5 uint8 = 255
    v6 := make(map[string]uint32)
    var v8 uint8 = (v5 & v5)
    var v9 uint8 = 42
    var v10 uint = 255
    if v3 {
        fmt.Println(v4, v3)
    }
    _ = len(v2)
    v1 = (v1 | (^(v1)))
    switch "" {
    case "hello":
        var v11 bool = v3
        _ = v11
    default:
        v10 -= v10
    }
    var v12 float32 = 1e10
    v13 := GenS1{F0: int8(v10), F1: 389, F2: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
