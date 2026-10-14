package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint64 }

func main() {
    v1 := []uint32{2}
    var v2 int8 = 42
    v3 := make(map[uint32]uint32)
    var v4 string = "\n\t"
    v5 := v3[2]
    v6 := GenS1{F0: 2, F1: v2, F2: (v2 << ((255) & 63))}
    for fi7 := 0; fi7 < 18; fi7++ {
        if (420 != 42) {
            v5 += 2
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
            fmt.Println(v4)
        }
        if (3 < fi7) {
            var v8 string = v4
            fmt.Println(v2, v8)
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 bool = true
                var v11 uint = 42
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = v8
        }
        fmt.Println(fi7, v5, fi7)
        _ = fi7
    }
    if true {
        for fi12 := 0; fi12 < 19; fi12++ {
            if ((1 < 714) || false) {
                fmt.Println(v2, v5)
                v4 = v4
                if (4294967295 >= 2) {
                    v2 -= 0
                    var v13 int8 = v2
                    v2 -= v13
                    _ = v13
                }
            }
            _ = fi12
        }
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 int32 = -832
            _ = fi14
            _ = v15
        }
        v16 := GenS1{F0: 152, F1: v2, F2: v2}
        _ = v16
    } else {
        v3[42] = (v5 / ((0) | 1))
        _ = len(v1)
    }
    var v17 float32 = (1.5 * 3.14159)
    v18 := GenS1{F0: 14, F1: -56, F2: v2}
    v19 := GenS2{F0: v5, F1: 2, F2: 491}
    var v20 int32 = 876
    delete(v3, 0)
    fmt.Println(v2, v4)
    v21 := GenS1{F0: 1, F1: v2, F2: v2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
