package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }

func main() {
    var v1 float32 = 0.0
    v2 := make(map[string]string)
    var v3 uint8 = 48
    var v4 int16 = 1000
    v5 := GenS1{F0: 65535, F1: (float32(3.14159) > 1e10)}
    v6 := GenS1{F0: uint32(v3), F1: (2 > 272)}
    var v7 int32 = 2147483647
    switch ("a" + ("a" + "")) {
    case "hello":
        var v8 uint16 = uint16(v3)
        _ = v8
    default:
        v4 = (v4 - v4)
    }
    if true {
        fmt.Println(v3)
    }
    v9 := []int{784, 0, -1}
    v7 = 965
    v1 -= ((float32(1e10) + 1e10) - float32((3.14159 - 1e10)))
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 int64 = -9223372036854775808
        fmt.Println(fi10, v1)
        _ = fi10
        _ = v11
    }
    if false {
        var v12 uint32 = 843
        _ = v12
    }
    fmt.Println(v1)
    v7 = (v7 >> ((1) & 63))
    var v13 int8 = -128
    v2["a"] = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
