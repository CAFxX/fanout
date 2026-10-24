package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint16, p1 int, p2 int64) uint8 {
    p2 -= 307
    p1 += (^(p1))
    switch 32767 {
    case 429:
        v4 := []int{}
        _ = v4
    }
    p2 = p2
    var v5 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return uint8(v5)
}

func main() {
    v1 := []uint32{2147483648, 4294967295, 0, 0}
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    var v6 uint8 = 75
    v6 = (v6 + (^(v6)))
    switch "a" {
    case "a":
        v6 += v6
    default:
        v7 := GenS1{F0: -2147483648}
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := GenS1{F0: -1}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v6 -= (v6 * 42)
    v11 := v2[4294967295]
    v12 := GenS1{F0: -2147483648}
    for fi13 := 0; fi13 < 9; fi13++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v6 += (^(v6))
        _ = fi13
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v6
}
