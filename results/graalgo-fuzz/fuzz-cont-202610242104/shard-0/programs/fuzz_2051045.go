package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint; F1 uint; F2 uint16 }

func main() {
    var v1 int = 2147483647
    v2 := []uint32{42}
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 2147483647
    var v5 uint = 255
    fmt.Println(v5)
    v6 := GenS1{F0: 255, F1: -548}
    if len(v2) > 0 {
        v2[0] = v4
    }
    v4 = 4294967295
    var v7 float64 = (3.14159 * ((0.1 * -2.25) + 1.5))
    fmt.Println(v7)
    var v8 float32 = float32(3.14159)
    for fi9 := 0; fi9 < 3; fi9++ {
        v4 += (v4 & 4294967295)
        _ = fi9
    }
    v1 += (v1 & (v1 << (((v5 << (((^(v5))) & 63))) & 63)))
    v10 := GenS2{F0: (v5 - (v5 & v5)), F1: v5, F2: 2}
    if ((65535 != 107) && (v5 == v5)) {
        delete(v3, (v4 * v4))
        switch v1 {
        case 128:
            if len(v2) > 0 {
                v2[0] = (v4 / ((v4) | 1))
            }
        }
        v1 += v1
    } else {
        fmt.Println(v8)
        var v11 int = v1
        _ = v11
    }
    v4 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
