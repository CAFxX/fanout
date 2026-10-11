package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 string; F2 uint64 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int, p1 int64, p2 int16) int {
    p2 -= (p2 ^ (-(p2)))
    fmt.Println(p1, p2)
    p0 = 4294967295
    for fi5 := 0; fi5 < 12; fi5++ {
        fmt.Println(p0, p2, fi5)
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = p2
    return (p0 % ((9223372036854775807) | 1))
}

func main() {
    var v1 uint64 = 42
    var v2 int = 0
    v3 := make(map[int]int)
    var v4 string = ("" + "hello")
    v6 := []int{v2, genFn0((v2 + (-(v2))), -912, 0), (v2 ^ v2)}
    v7 := GenS1{F0: 127}
    if true {
        v8 := GenS2{F0: 738, F1: v4, F2: 585}
        _ = v8
    } else {
        v1 = v1
        var v9 int32 = 2147483647
        _ = v9
    }
    v1 = (v1 ^ (^(v1)))
    v10 := GenS1{F0: 255}
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 uint32 = 702
        for fi13 := 0; fi13 < 10; fi13++ {
            v12 -= 0
            v14 := GenS1{F0: 127}
            v15 := GenS1{F0: 62}
            _ = fi13
            _ = v14
            _ = v15
        }
        v12 -= v12
        _ = fi11
        _ = v12
    }
    v16 := GenS3{F0: 932}
    v4 = "\n\t"
    switch (v2 << ((745) & 63)) {
    case 139:
        var v17 uint64 = 0
        _ = v17
    case -859:
        var v18 uint8 = 191
        _ = v18
    default:
        v2 = v2
    }
    var v19 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
