package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint64 }

func main() {
    var v1 uint8 = 212
    var v2 float32 = -2.25
    var v3 bool = (2147483647 != 9223372036854775807)
    v4 := make(map[string]string)
    v5 := make(map[int]uint32)
    var v6 int = 42
    var v7 uint = 42
    switch "a" {
    case "hello":
        v8 := v4["\n\t"]
        _ = v8
    case "\n\t":
        var v9 uint = 150
        _ = v9
    default:
        v2 -= v2
    }
    var v10 float32 = v2
    switch v6 {
    case 826:
        fmt.Println(v1, v2)
    default:
        var v11 uint = 65535
        _ = v11
    }
    v12 := []int{(v6 - (v6 << (((v7 >> ((v7) & 63))) & 63))), (v6 / ((128) | 1)), -1, (v6 / (((-(v6))) | 1))}
    v3 = true
    for fi13 := 0; fi13 < 16; fi13++ {
        v14 := v4[(("foo bar" + "a") + ("foo bar" + "a"))]
        _ = fi13
        _ = v14
    }
    v10 += v2
    if v3 {
        v15 := GenS1{F0: v1, F1: v3, F2: 4294967295}
        for fi16 := 0; fi16 < 4; fi16++ {
            v3 = ((65535 > 42) || v3)
            _ = fi16
        }
        v1 += v1
        _ = v15
    }
    v5[v6] = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
