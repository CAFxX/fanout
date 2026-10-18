package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }

func main() {
    var v1 int32 = 1
    var v2 int32 = v1
    var v3 uint16 = 412
    v3 = (v3 << ((836) & 63))
    var v4 float64 = float64(3.14159)
    fmt.Println(v2, v3, v4)
    v5 := []int{}
    switch "hello" {
    case "":
        v3 -= v3
    case "\n\t":
        if (!false) {
            v5 = append(v5, 9223372036854775807)
            v4 = 1.5
        }
    default:
        fmt.Println(v1, v3)
    }
    _ = len(v5)
    var v6 uint = 18446744073709551615
    _ = len(v5)
    var v7 float32 = 0.0
    var v8 float64 = v4
    var v9 uint16 = v3
    v10 := GenS1{F0: ("a" + "\n\t"), F1: 83}
    v11 := GenS1{F0: "hello", F1: 203}
    v12 := GenS1{F0: "foo bar", F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
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
