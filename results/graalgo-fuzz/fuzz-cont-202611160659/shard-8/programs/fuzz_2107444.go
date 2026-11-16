package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint16 }

func main() {
    var v1 int64 = 1
    v2 := make(map[uint32]bool)
    var v3 int = 4294967295
    var v4 string = "a"
    v3 = v3
    v5 := GenS1{F0: 4294967295, F1: 893}
    var v6 bool = (!(((-2.25 + float64(3.14159)) - (float64(float32((3.14159 / -2.25))) * (1.5 / 3.14159))) >= ((-2.25 / (0.0 - (0.1 * float64(1.5)))) + (0.0 * 1e10))))
    var v7 int = v3
    var v8 float64 = ((1.5 * 1.5) - (0.0 / 3.14159))
    var v9 int = v7
    if v6 {
        var v10 int32 = 42
        if v6 {
            v9 = v3
            v11 := GenS2{F0: 650, F1: 1, F2: 65535}
            v12 := GenS1{F0: 18446744073709551615, F1: 7}
            _ = v11
            _ = v12
        }
        _ = v10
    }
    v13 := []int{(v3 ^ v7), v9, v3}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
