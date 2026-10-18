package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[int]uint32)
    var v3 bool = (!true)
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]bool)
    v6 := []int{2147483647}
    if v3 {
        if (v3 || false) {
            var v7 float32 = 1.5
            _ = v7
        }
        var v8 int32 = 2147483647
        _ = v8
    }
    v6 = append(v6, 881)
    var v9 uint16 = 490
    v10 := GenS1{F0: 532, F1: -128}
    var v11 uint64 = 572
    v11 -= v11
    var v12 int32 = 42
    v13 := GenS1{F0: -2147483648, F1: -39}
    switch 535 {
    case 542:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case 32767:
        switch ("foo bar" + "") {
        case "hello":
            var v14 float32 = v1
            _ = v14
        }
    default:
        _ = len(v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
