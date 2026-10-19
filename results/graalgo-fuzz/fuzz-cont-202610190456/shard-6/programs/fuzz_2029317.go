package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = -2.25
    var v3 bool = false
    var v4 int = 341
    v5 := make(map[string]uint32)
    v2 = (3.14159 - v2)
    v6 := []int{(v4 - (v4 & -2147483648)), (-(v4)), (v4 & (v4 & (v4 << ((1) & 63))))}
    var v7 float64 = 0.1
    if false {
        switch "" {
        case "foo bar":
            v7 = 3.14159
        case "\n\t":
            _ = len(v6)
        case "hello":
            v8 := GenS1{F0: false, F1: "foo bar"}
            _ = v8
        }
        v9 := GenS2{F0: 65535, F1: 960, F2: -1}
        v4 = -1
        _ = v9
    }
    fmt.Println(v2)
    if (!(!(!(-512 <= -9223372036854775808)))) {
        v4 = (v4 - (v4 << ((65535) & 63)))
    }
    v6 = append(v6, v4)
    fmt.Println(v7)
    v10 := GenS3{F0: 673}
    v7 = v7
    if len(v6) > 0 {
        v6[0] = 883
    }
    var v11 int16 = 445
    var v12 float32 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
