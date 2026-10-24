package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 uint32 = 42
    var v2 uint8 = 1
    var v3 uint8 = (v2 | v2)
    v4 := []int{}
    _ = len(v4)
    if (((((1e10 / float32((3.14159 + 0.1))) * 1e10) / 1e10) != (float32(-2.25) + float32(-2.25))) && true) {
        v1 = 2
        v5 := GenS1{F0: 2147483647, F1: (v1 + v1)}
        v6 := GenS1{F0: 1, F1: (v1 % ((v1) | 1))}
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 1, F1: 315}
    for fi8 := 0; fi8 < 7; fi8++ {
        if ((!(true || ((^(v2)) == 2))) && true) {
            fmt.Println(fi8, v3)
            v2 -= (v3 ^ (v2 << ((42) & 63)))
        }
        v1 += (^(v1))
        var v9 int8 = -1
        _ = fi8
        _ = v9
    }
    v3 = uint8(v1)
    var v10 int32 = 42
    switch 1 {
    case 0:
        var v11 int64 = 42
        _ = v11
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch 127 {
    case 364:
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 uint32 = v1
            _ = fi12
            _ = v13
        }
    case 2147483647:
        v14 := GenS2{F0: -1, F1: (v2 + v2), F2: (^(v3))}
        _ = v14
    case 32767:
        v15 := GenS1{F0: 32767, F1: v1}
        _ = v15
    default:
        if (true && (!(!(!(9223372036854775807 >= -674))))) {
            v10 = (v10 >> ((42) & 63))
            _ = len(v4)
        } else {
            v1 = v1
            v4 = append(v4, 128)
            v10 -= (v10 ^ v10)
        }
    }
    v16 := GenS3{F0: (v1 | (v1 % ((v1) | 1)))}
    var v17 int32 = v10
    for fi18 := 0; fi18 < 17; fi18++ {
        if len(v4) > 0 {
            v4[0] = fi18
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
