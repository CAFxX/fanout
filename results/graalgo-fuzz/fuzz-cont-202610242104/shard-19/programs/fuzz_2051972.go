package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int32 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint8) int16 {
    p0 = (p0 + p0)
    p0 = p0
    var v6 bool = (-429 <= 233)
    fmt.Println(v6, p0)
    v7 := []int{128, 4294967295}
    _ = p0
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 int64 = 1
    v2 := make(map[string]uint32)
    v3 := []int{-1, -540, 2147483647, 529}
    v4 := make(map[uint32]string)
    var v5 int = -419
    v8 := v4[2147483647]
    v5 -= (v5 % ((v5) | 1))
    for fi9 := 0; fi9 < 18; fi9++ {
        switch v5 {
        case -9223372036854775808:
            v10 := GenS1{F0: 545, F1: 4294967295, F2: 0}
            _ = v10
        default:
            var v11 string = ""
            _ = v11
        }
        for fi12 := 0; fi12 < 15; fi12++ {
            if true {
                v1 = v1
                var v13 float64 = (float64(0.0) + (3.14159 * 1.5))
                var v14 uint64 = 177
                _ = v13
                _ = v14
            }
            var v15 int16 = genFn0(175)
            _ = fi12
            _ = v15
        }
        _ = fi9
    }
    v16 := GenS1{F0: 0, F1: 0, F2: 42}
    var v17 uint8 = 128
    fmt.Println(v17, v17, v8)
    v17 += 128
    var v18 int = v5
    switch v18 {
    case 9223372036854775807:
        switch -1 {
        case 9223372036854775807:
            var v19 int64 = (v1 - (v1 >> ((933) & 63)))
            _ = v19
        case -364:
            v20 := GenS2{F0: (!true)}
            _ = v20
        case 179:
            var v21 uint32 = 596
            _ = v21
        }
    case 1:
        v8 = v8
    case 116:
        var v22 int8 = -39
        _ = v22
    }
    var v23 int16 = -294
    v1 += -1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
