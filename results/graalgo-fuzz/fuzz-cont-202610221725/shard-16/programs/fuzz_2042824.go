package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int; F1 uint }
func genFn0(p0 int32, p1 int8) int32 {
    v6 := GenS1{F0: 31, F1: 65535, F2: 1}
    if true {
        p1 = (p1 + (p1 & 1))
        p0 -= (-(p0))
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    _ = v6
    return (p0 << ((255) & 63))
}

func main() {
    v1 := []int{32767, 42, -16}
    var v2 uint16 = 305
    v3 := make(map[string]int)
    var v4 uint8 = 216
    v5 := make(map[int]uint32)
    fmt.Println(v2)
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS2{F0: (^(fi7)), F1: 0}
        var v9 float64 = (float64(-2.25) + 3.14159)
        var v10 int64 = 0
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    if true {
        var v11 int16 = -32768
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 int32 = -1
            _ = fi12
            _ = v13
        }
        switch 42 {
        case 957:
            var v14 uint8 = 105
            _ = v14
        }
        _ = v11
    } else {
        var v15 uint8 = (v4 & 1)
        v2 = 42
        if false {
            var v16 int32 = 42
            _ = v16
        } else {
            var v17 int64 = 0
            for fi18 := 0; fi18 < 13; fi18++ {
                v19 := GenS1{F0: (v15 % ((uint8(v2)) | 1)), F1: 32767, F2: 25}
                _ = fi18
                _ = v19
            }
            _ = v17
        }
        _ = v15
    }
    if true {
        v4 = v4
    }
    var v20 uint8 = 2
    v21 := GenS1{F0: 128, F1: (v2 << ((0) & 63)), F2: 576}
    v20 = v20
    var v22 float32 = (1.5 * 1.5)
    var v23 int64 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
