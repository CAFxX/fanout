package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 string; F2 uint }
type GenS4 struct { F0 int64; F1 uint8 }
func genFn0() int64 {
    v5 := GenS1{F0: 2147483648}
    var v6 float64 = -2.25
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := []int{(fi7 / ((fi7) | 1)), (^(fi7))}
        v9 := GenS2{F0: 776}
        _ = fi7
        _ = v8
        _ = v9
    }
    v6 = (1.5 * (v6 / (v6 + (((v6 - 3.14159) / float64(1.5)) * 0.0))))
    _ = v5
    _ = v6
    return -352
}

func main() {
    var v1 uint16 = 594
    var v2 float64 = 1.5
    v3 := make(map[int]uint32)
    var v4 uint32 = 42
    var v10 float64 = v2
    fmt.Println(v4)
    var v11 int64 = -519
    v3[42] = (^(v4))
    var v12 uint = uint(v1)
    switch 1 {
    case 128:
        var v13 uint32 = (v4 % ((v4) | 1))
        _ = v13
    case -2147483648:
        if ((v1 ^ 0) == v1) {
            var v14 int8 = -128
            for fi15 := 0; fi15 < 8; fi15++ {
                var v16 uint16 = 0
                _ = fi15
                _ = v16
            }
            fmt.Println(v10, v4, v2)
            _ = v14
        }
    case 42:
        v17 := []int{42, 9223372036854775807, 42}
        _ = v17
    default:
        v11 = 42
    }
    var v18 uint64 = 0
    v19 := []int{}
    v20 := GenS1{F0: 2}
    v21 := GenS1{F0: v4}
    v22 := GenS3{F0: uint64(v1), F1: "a", F2: (v12 ^ (v12 << (((^(v12))) & 63)))}
    var v23 uint8 = 128
    v24 := GenS4{F0: v11, F1: v23}
    v11 = genFn0()
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
