package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := []int{-1, 32768, -310, 9223372036854775807}
    var v2 int32 = 255
    v3 := make(map[int]bool)
    var v4 uint = 4294967295
    var v5 uint = 42
    v6 := []int{}
    v7 := GenS1{F0: 808, F1: (-(v2)), F2: 255}
    if (!((!false) || true)) {
        var v8 int16 = 0
        _ = v8
    }
    v9 := GenS1{F0: -93, F1: v2, F2: 69}
    var v10 uint64 = 99
    for fi11 := 0; fi11 < 15; fi11++ {
        v5 = v5
        fmt.Println(v10, fi11)
        var v12 uint32 = 2
        _ = fi11
        _ = v12
    }
    if (("\n\t" > "\n\t") || false) {
        fmt.Println(v10)
        v6 = append(v6, -110)
        var v13 bool = (2 <= 192)
        _ = v13
    }
    v5 -= v4
    for fi14 := 0; fi14 < 10; fi14++ {
        fmt.Println(fi14, v2)
        v15 := GenS2{F0: 255}
        v5 = (v5 / (((v4 | (v4 - 65535))) | 1))
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: 32767}
    switch "hello" {
    case "hello":
        var v17 int16 = -1
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
