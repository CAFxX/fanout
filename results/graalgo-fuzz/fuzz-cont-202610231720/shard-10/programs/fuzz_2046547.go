package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int8; F1 string; F2 int32 }

func main() {
    v1 := []uint32{1}
    var v2 float64 = (-2.25 - float64(0.1))
    var v3 uint16 = 65535
    v4 := make(map[uint32]uint32)
    v5 := []string{"a", "foo bar"}
    var v6 int = 9223372036854775807
    var v7 int8 = 127
    v3 = 65535
    _ = len(v1)
    switch "foo bar" {
    case "hello":
        v6 = v6
    case "foo bar":
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := GenS1{F0: 2, F1: 849}
            fmt.Println(v6)
            _ = fi8
            _ = v9
        }
    case "a":
        v10 := GenS2{F0: 127, F1: "", F2: 2147483647}
        _ = v10
    }
    v7 -= (v7 & v7)
    if ("foo bar" >= "foo bar") {
        fmt.Println(v6, v7)
        v11 := GenS1{F0: 4294967295, F1: -2147483648}
        _ = v11
    }
    v6 = int(v7)
    v2 -= (((0.0 + 0.1) * v2) - (v2 - v2))
    var v12 uint16 = (^(v3))
    var v13 uint8 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
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
