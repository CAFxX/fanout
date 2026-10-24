package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]int)
    var v3 int8 = 0
    var v4 uint = 231
    fmt.Println(v3, v4)
    v5 := GenS1{F0: 42, F1: 4}
    switch 0 {
    case 858:
        v6 := GenS2{F0: "hello"}
        _ = v6
    case 9223372036854775807:
        var v7 uint = (v4 + v4)
        _ = v7
    case 32768:
        var v8 int32 = 0
        _ = v8
    default:
        delete(v2, "a")
    }
    v4 = 2
    v9 := v2[("hello" + "\n\t")]
    if false {
        var v10 uint64 = 0
        var v11 bool = false
        _ = v10
        _ = v11
    } else {
        var v12 int16 = -32768
        v3 = v3
        _ = v12
    }
    var v13 uint = (^(v4))
    fmt.Println(v13, v9)
    var v14 float64 = 1e10
    var v15 float32 = 3.14159
    switch (-(v9)) {
    case -271:
        v13 = 255
    default:
        v16 := GenS3{F0: 42, F1: 42}
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
