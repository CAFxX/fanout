package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int64; F1 string }

func main() {
    var v1 int8 = -13
    var v2 uint64 = 18446744073709551615
    var v3 uint32 = 4294967295
    var v4 int8 = v1
    var v5 int32 = -1
    fmt.Println(v1, v4, v3)
    var v6 uint64 = (^(v2))
    v7 := GenS1{F0: (v2 ^ v2), F1: -514}
    switch "\n\t" {
    case "a":
        var v8 float32 = float32(0.1)
        _ = v8
    default:
        if true {
            var v9 int = 600
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := []int{(fi10 - fi10)}
        v2 = v6
        _ = fi10
        _ = v11
    }
    v12 := []int{}
    v13 := GenS2{F0: -100, F1: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
