package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = ((0.0 - 0.1) == 1e10)
    var v3 int16 = 1
    v4 := GenS1{F0: 69, F1: 42}
    delete(v1, "\n\t")
    fmt.Println(v3)
    v3 += (v3 + v3)
    v2 = v2
    v2 = (42 == 646)
    v3 += (v3 >> ((65535) & 63))
    fmt.Println(v2, v3)
    fmt.Println(v3, v3)
    switch 2147483647 {
    case -636:
        var v5 float32 = 1.5
        _ = v5
    default:
        if v2 {
            fmt.Println(v3, v3)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
