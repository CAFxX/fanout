package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint }

func main() {
    var v1 uint8 = 1
    var v2 int8 = -107
    var v3 float32 = (3.14159 - (1e10 - 3.14159))
    v4 := make(map[string]bool)
    var v5 float64 = (float64(v3) + 1e10)
    switch 32767 {
    case 102:
        fmt.Println(v3)
    default:
        var v6 uint = 4294967295
        _ = v6
    }
    v5 = 0.1
    v2 = v2
    switch 2147483647 {
    case 32768:
        v7 := GenS1{F0: (!true)}
        _ = v7
    case 127:
        v8 := GenS1{F0: (!true)}
        _ = v8
    default:
        v1 = (v1 >> ((0) & 63))
    }
    var v9 int8 = 0
    v10 := []int{-494}
    if len(v10) > 0 {
        v10[0] = -2147483648
    }
    var v11 uint8 = 4
    var v12 string = "foo bar"
    delete(v4, v12)
    var v13 int8 = (v9 << ((18446744073709551615) & 63))
    v14 := GenS2{F0: 65535, F1: 19}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
