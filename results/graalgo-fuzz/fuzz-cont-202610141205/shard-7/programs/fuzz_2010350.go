package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint }
type GenS2 struct { F0 uint; F1 uint16; F2 int8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = 3.14159
    var v3 float64 = (1.5 * float64(float32(v2)))
    var v4 bool = true
    var v5 uint8 = 132
    switch "foo bar" {
    case "":
        v6 := []int{-9223372036854775808, 128, 32768, -2147483648}
        _ = v6
    case "a":
        v2 = 0.1
    }
    var v7 uint = 255
    switch -175 {
    case 32767:
        v4 = (100 > -1)
    case 2147483647:
        v8 := GenS1{F0: v5, F1: -128, F2: 42}
        _ = v8
    default:
        v2 -= 1.5
    }
    _ = len(v1)
    fmt.Println(v3, v4, v5)
    var v9 int16 = 1
    v4 = (964 >= 914)
    v3 -= 3.14159
    v1[5] = ""
    if v4 {
        fmt.Println(v7)
        var v10 int64 = -819
        _ = v10
    }
    v11 := GenS2{F0: (v7 << (((v7 ^ (v7 + v7))) & 63)), F1: 932, F2: int8(v7)}
    v4 = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
