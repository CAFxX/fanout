package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []int{}
    var v3 int16 = 229
    var v4 uint32 = 2147483647
    var v5 uint32 = 994
    v6 := GenS1{F0: 94, F1: (v3 != 42)}
    switch 128 {
    case 128:
        v5 = (v4 + 2147483647)
    case 2147483647:
        v7 := GenS1{F0: 2, F1: ((2147483647 >= 0) && true)}
        _ = v7
    default:
        for fi8 := 0; fi8 < 3; fi8++ {
            v4 -= (v5 - v5)
            _ = fi8
        }
    }
    v3 += (-(v3))
    v1 = 1
    fmt.Println(v4, v1, v5)
    if (!true) {
        v5 -= (v5 % ((0) | 1))
    }
    if (-995 != -256) {
        v2 = append(v2, 127)
    }
    if (-128 > 0) {
        v9 := GenS2{F0: (v4 | (v4 >> ((326) & 63)))}
        v3 -= v3
        _ = v9
    }
    v10 := GenS2{F0: 2}
    var v11 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
