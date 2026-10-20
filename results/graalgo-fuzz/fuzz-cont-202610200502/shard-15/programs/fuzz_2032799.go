package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint32 }

func main() {
    v1 := make(map[string]int)
    var v2 int = -9223372036854775808
    var v3 int64 = -9223372036854775808
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]bool)
    v6 := make(map[int]uint32)
    var v7 int8 = 1
    v8 := GenS1{F0: 613, F1: 2}
    v9 := GenS2{F0: 255, F1: 702}
    fmt.Println(v2, v2)
    var v10 int64 = -9223372036854775808
    var v11 int = v2
    v12 := GenS2{F0: 33, F1: 65535}
    if false {
        v13 := GenS2{F0: 127, F1: 633}
        _ = v13
    }
    switch ("" + "") {
    case "a":
        v14 := v5[("" + "hello")]
        _ = v14
    case "\n\t":
        var v15 uint64 = 2
        _ = v15
    default:
        v16 := GenS1{F0: 65535, F1: 255}
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
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
    _ = v8
    _ = v9
}
