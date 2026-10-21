package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int32 }
func genFn0(p0 int16, p1 uint8, p2 uint8) int {
    p2 = p2
    var v6 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 128
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[uint32]bool)
    v3 := []uint32{1, 4294967295, 4294967295, 65535}
    var v4 uint64 = 18446744073709551615
    v5 := make(map[int]bool)
    _ = len(v3)
    switch 127 {
    case 0:
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 int32 = -158
            _ = fi7
            _ = v8
        }
    case 1:
        v9 := GenS1{F0: 1, F1: 0, F2: -987}
        _ = v9
    }
    var v10 uint8 = 1
    delete(v5, genFn0((v1 % ((v1) | 1)), (v10 + (v10 / ((v10) | 1))), (v10 - (v10 << ((4294967295) & 63)))))
    if len(v3) > 0 {
        v3[0] = 2147483648
    }
    fmt.Println(v10)
    _ = len(v3)
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    if len(v3) > 0 {
        v3[0] = 42
    }
    var v11 uint64 = (v4 & (v4 - (^(v4))))
    if len(v3) > 0 {
        v3[0] = 1
    }
    var v12 uint = 42
    var v13 int16 = (v1 & -157)
    var v14 int32 = 962
    if (42 == 2) {
        fmt.Println(v11, v12, v11)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
