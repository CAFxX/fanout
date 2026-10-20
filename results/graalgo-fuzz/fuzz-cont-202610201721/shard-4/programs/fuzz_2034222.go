package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{467, -2147483648, 42}
    var v3 int = 42
    var v4 float32 = (0.0 / 1.5)
    v4 = 1e10
    fmt.Println(v3)
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS1{F0: "\n\t", F1: -1}
        _ = fi5
        _ = v6
    }
    v4 = (3.14159 / v4)
    v4 = v4
    var v7 uint64 = 0
    v8 := GenS1{F0: "foo bar", F1: 1}
    v4 -= v4
    switch "" {
    case "\n\t":
        v9 := GenS1{F0: "hello", F1: 2147483647}
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 string = "hello"
        fmt.Println(fi10)
        _ = fi10
        _ = v11
    }
    v2 = append(v2, -2147483648)
    var v12 float32 = float32(3.14159)
    v13 := GenS2{F0: -1}
    for fi14 := 0; fi14 < 11; fi14++ {
        v3 += v3
        var v15 bool = true
        _ = fi14
        _ = v15
    }
    var v16 uint = 0
    var v17 float64 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
