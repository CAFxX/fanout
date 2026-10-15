package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }

func main() {
    v1 := []uint32{381, 558, 225}
    v2 := []uint32{2, 65535, 2147483647}
    var v3 float64 = float64((float32(0.0) / (1e10 + float32(0.0))))
    v4 := []int{128, -276}
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]int)
    var v7 int16 = -792
    for fi8 := 0; fi8 < 9; fi8++ {
        switch "foo bar" {
        case "a":
            v9 := GenS1{F0: -1, F1: 0}
            _ = v9
        default:
            var v10 int = fi8
            _ = v10
        }
        _ = fi8
    }
    v7 = (v7 | (v7 & (-(v7))))
    fmt.Println(v3)
    v11 := GenS1{F0: 9223372036854775807, F1: 127}
    var v12 string = ("" + "")
    for fi13 := 0; fi13 < 8; fi13++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v7 -= v7
        _ = fi13
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v12)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
