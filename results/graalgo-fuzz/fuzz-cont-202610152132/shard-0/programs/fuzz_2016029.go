package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]uint32)
    v3 := []string{"foo bar"}
    var v4 bool = false
    var v5 float64 = (float64(-2.25) * 3.14159)
    var v6 int16 = 1
    v7 := GenS1{F0: 127, F1: v6}
    for fi8 := 0; fi8 < 10; fi8++ {
        v5 = v5
        _ = fi8
    }
    var v9 uint8 = 42
    var v10 int8 = 127
    v11 := GenS1{F0: 0, F1: v6}
    v5 += v5
    switch "foo bar" {
    case "a":
        var v12 float32 = ((-2.25 * (0.0 + 0.1)) * 1.5)
        _ = v12
    case "":
        fmt.Println(v9, v1)
    case "hello":
        v13 := GenS1{F0: v9, F1: -1}
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
