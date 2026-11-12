package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 float64 = -2.25
    v2 := []uint32{1, 2147483647}
    var v3 uint64 = 18446744073709551615
    v4 := make(map[uint32]int)
    v4[2] = -1
    if len(v2) > 0 {
        v2[0] = 65535
    }
    v3 = v3
    v1 = (v1 - 1.5)
    v5 := GenS1{F0: (!(-40 == -84)), F1: 0}
    fmt.Println(v3, v1)
    v6 := v4[602]
    var v7 int8 = 1
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 14; fi9++ {
            switch "foo bar" {
            case "foo bar":
                v10 := v4[2147483647]
                _ = v10
            default:
                var v11 int32 = -1
                _ = v11
            }
            fmt.Println(v1, v6)
            _ = fi9
        }
        _ = fi8
    }
    fmt.Println(v6)
    v12 := GenS2{F0: 1}
    v13 := GenS2{F0: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
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
