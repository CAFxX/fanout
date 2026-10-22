package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint64; F1 int32; F2 int64 }

func main() {
    var v1 int8 = -87
    var v2 bool = false
    var v3 uint16 = 571
    v4 := []uint32{42, 0, 65535}
    v1 = (^(v1))
    switch 0 {
    case 1:
        var v5 int = 128
        _ = v5
    case 439:
        v6 := GenS1{F0: v3}
        _ = v6
    case 0:
        for fi7 := 0; fi7 < 1; fi7++ {
            fmt.Println(v1, fi7, v2)
            _ = fi7
        }
    default:
        for fi8 := 0; fi8 < 17; fi8++ {
            if ((0.0 * (-2.25 + (3.14159 + 0.0))) <= 3.14159) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            v3 = 0
            for fi9 := 0; fi9 < 8; fi9++ {
                v10 := GenS2{F0: v3, F1: 1}
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    }
    var v11 int32 = 0
    var v12 int = -1
    v2 = v2
    if len(v4) > 0 {
        v4[0] = 1
    }
    fmt.Println(v2, v11, v1)
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 uint64 = 362
        if (211 <= 1) {
            v14 += v14
        }
        _ = fi13
        _ = v14
    }
    v15 := GenS3{F0: 4294967295, F1: v11, F2: 42}
    for fi16 := 0; fi16 < 6; fi16++ {
        v12 = (fi16 + 115)
        v11 = v11
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
