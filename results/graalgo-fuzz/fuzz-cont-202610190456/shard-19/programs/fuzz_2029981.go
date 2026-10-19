package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{}
    v3 := make(map[int]uint32)
    var v4 uint32 = 627
    for fi5 := 0; fi5 < 13; fi5++ {
        switch fi5 {
        case 582:
            v6 := GenS1{F0: (v4 + v4), F1: 2, F2: 1}
            _ = v6
        }
        _ = fi5
    }
    if false {
        var v7 uint32 = 2147483647
        _ = v7
    }
    v4 = v4
    v4 = (v4 | 914)
    var v8 float64 = 1.5
    v8 = ((float64((1e10 / ((0.1 / 1.5) + 0.1))) - (0.0 / (v8 / v8))) / (0.0 * v8))
    var v9 uint64 = 1
    for fi10 := 0; fi10 < 12; fi10++ {
        v4 -= (^(v4))
        _ = fi10
    }
    var v11 uint8 = 106
    switch "" {
    case "foo bar":
        v12 := GenS1{F0: (v4 & v4), F1: 42, F2: 255}
        _ = v12
    case "":
        v11 -= (^(v11))
    default:
        for fi13 := 0; fi13 < 18; fi13++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi13
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
