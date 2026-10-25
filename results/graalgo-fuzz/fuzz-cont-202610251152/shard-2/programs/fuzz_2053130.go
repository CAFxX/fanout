package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }

func main() {
    var v1 float64 = 1.5
    var v2 int = 128
    v3 := make(map[uint32]uint32)
    var v4 int64 = 9223372036854775807
    var v5 float32 = float32(((0.0 / ((v1 - v1) * 0.1)) + v1))
    var v6 uint32 = 305
    if true {
        var v7 string = "\n\t"
        _ = v7
    }
    v1 = (v1 * ((v1 / (v1 + v1)) + v1))
    v4 = (v4 % ((v4) | 1))
    v8 := []int{v2}
    var v9 int32 = 42
    _ = len(v8)
    var v10 int64 = v4
    fmt.Println(v10, v1, v10)
    v11 := GenS1{F0: 4294967295, F1: ""}
    v2 = (-(v2))
    var v12 uint = 425
    fmt.Println(v10)
    v13 := GenS1{F0: 18446744073709551615, F1: ("a" + ("hello" + "a"))}
    if len(v8) > 0 {
        v8[0] = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
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
