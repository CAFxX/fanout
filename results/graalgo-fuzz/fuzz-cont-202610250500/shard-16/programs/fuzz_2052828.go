package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }

func main() {
    var v1 int16 = 32767
    v2 := make(map[int]string)
    var v3 float32 = 0.1
    v4 := make(map[string]uint32)
    var v5 float64 = (-2.25 - ((0.0 + 1.5) / (-2.25 * (1.5 + (1e10 - float64(v3))))))
    var v6 uint8 = 42
    v7 := []int{}
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int16 = v1
        for fi10 := 0; fi10 < 15; fi10++ {
            v3 -= v3
            v3 = float32(-2.25)
            _ = fi10
        }
        v6 -= v6
        _ = fi8
        _ = v9
    }
    var v11 uint32 = 2147483648
    fmt.Println(v3, v6, v5)
    v12 := GenS1{F0: 2, F1: 844}
    switch 9223372036854775807 {
    case -538:
        for fi13 := 0; fi13 < 11; fi13++ {
            _ = len(v7)
            _ = len(v7)
            v14 := GenS1{F0: 1, F1: 2147483647}
            _ = fi13
            _ = v14
        }
    default:
        _ = len(v4)
    }
    var v15 int16 = (v1 ^ 639)
    switch "\n\t" {
    case "":
        _ = len(v4)
    }
    v1 = v1
    v16 := v2[-1]
    var v17 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
