package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }

func main() {
    var v1 uint64 = 780
    v2 := []uint32{702, 0}
    var v3 uint32 = 4294967295
    var v4 int64 = 1
    v5 := make(map[int]uint32)
    var v6 int32 = -2147483648
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := v5[fi7]
        delete(v5, fi7)
        var v9 uint16 = uint16(v1)
        _ = fi7
        _ = v8
        _ = v9
    }
    v1 += 443
    v2 = append(v2, 65535)
    _ = len(v2)
    v4 -= 42
    var v10 uint64 = 18446744073709551615
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        v10 -= v10
        v12 := GenS1{F0: false, F1: -155}
        var v13 uint = 2
        _ = fi11
        _ = v12
        _ = v13
    }
    v6 += 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
