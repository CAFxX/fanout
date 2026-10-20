package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 int32 }
func genFn0() uint64 {
    if (1 < -726) {
        var v4 uint64 = 0
        for fi5 := 0; fi5 < 18; fi5++ {
            var v6 float32 = (0.0 / 0.1)
            var v7 int32 = -2147483648
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = v4
    }
    var v8 uint8 = 127
    v9 := []int{}
    if ((0.0 * 1e10) <= -2.25) {
        for fi10 := 0; fi10 < 19; fi10++ {
            if (!false) {
                var v11 bool = ((true || (0 == 1)) && false)
                _ = v11
            } else {
                v8 = v8
            }
            var v12 uint8 = v8
            _ = fi10
            _ = v12
        }
    } else {
        var v13 uint8 = (v8 - v8)
        _ = v13
    }
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = -430
    v3 := []uint32{42, 42, 42}
    var v14 uint64 = 669
    fmt.Println(v14, v2)
    v14 = (v14 >> ((255) & 63))
    v2 += (-(v2))
    var v15 uint = 307
    v16 := GenS1{F0: 9223372036854775807, F1: false, F2: 0}
    if (0 == int64(v15)) {
        var v17 int16 = 0
        fmt.Println(v14, v2, v2)
        _ = v17
    }
    v18 := GenS2{F0: 42}
    v3 = append(v3, 0)
    v2 += v2
    var v19 uint32 = 1
    if true {
        v20 := v1["hello"]
        _ = v20
    }
    v2 = (v2 << ((uint(v19)) & 63))
    for fi21 := 0; fi21 < 7; fi21++ {
        v22 := GenS3{F0: -1, F1: 2147483647}
        _ = fi21
        _ = v22
    }
    v14 = (v14 - uint64(v15))
    for fi23 := 0; fi23 < 9; fi23++ {
        _ = len(v1)
        fmt.Println(v19)
        v19 += v19
        _ = fi23
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        v19 = v19
        _ = fi24
    }
    if (1e10 <= -2.25) {
        v2 -= (v2 % (((^(v2))) | 1))
    } else {
        var v25 uint16 = 893
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
