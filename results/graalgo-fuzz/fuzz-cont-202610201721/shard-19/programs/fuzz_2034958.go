package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint32 }
func genFn0(p0 uint64, p1 string, p2 int16) string {
    var v7 int16 = p2
    if (!false) {
        v8 := []int{127}
        _ = v8
    }
    switch 9223372036854775807 {
    case 0:
        fmt.Println(p0, v7)
    default:
        p0 = 879
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        p2 -= 32767
        var v10 uint32 = 42
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 0
    var v3 int32 = -702
    var v4 int = 32768
    var v5 bool = (1 == 1)
    var v6 int32 = (v3 / (((v3 / ((1) | 1))) | 1))
    v4 -= v4
    fmt.Println(v5)
    switch v4 {
    case 42:
        var v11 uint8 = 128
        _ = v11
    case 9223372036854775807:
        v5 = v5
    default:
        fmt.Println(v3, v2)
    }
    v12 := GenS1{F0: (v4 ^ v4), F1: 419, F2: 2}
    for fi13 := 0; fi13 < 12; fi13++ {
        v3 += v6
        _ = fi13
    }
    var v14 uint = 42
    v4 -= (v4 << (((v14 % (((v14 % ((547) | 1))) | 1))) & 63))
    v15 := []int{v4, v4, (^(v4)), (v4 & (^(v4)))}
    var v16 string = ""
    v5 = (v5 || v5)
    _ = len(v15)
    var v17 int8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
