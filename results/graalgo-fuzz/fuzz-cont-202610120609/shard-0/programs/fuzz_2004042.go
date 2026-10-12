package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint64 }

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 18446744073709551615
    v3 := make(map[string]bool)
    v4 := make(map[string]int)
    v5 := []uint32{644, 90, 2, 2}
    var v6 int16 = 1
    v7 := GenS1{F0: 0}
    switch 0 {
    case 9223372036854775807:
        v8 := GenS1{F0: 448}
        _ = v8
    case 4294967295:
        v9 := GenS1{F0: -2147483648}
        _ = v9
    case 2147483647:
        v10 := GenS2{F0: 57, F1: (v2 / ((2) | 1))}
        _ = v10
    default:
        var v11 uint32 = 0
        _ = v11
    }
    var v12 uint32 = 1
    var v13 float64 = 3.14159
    v2 += (v2 - v2)
    switch 2147483647 {
    case 4294967295:
        v13 = (1.5 + -2.25)
    case -1:
        var v14 int64 = 275
        _ = v14
    case 32767:
        v6 = v6
    }
    var v15 int16 = (v6 + -32768)
    for fi16 := 0; fi16 < 11; fi16++ {
        v5 = append(v5, (^(v12)))
        v15 -= (v6 * -32768)
        _ = fi16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
