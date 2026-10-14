package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint32 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int8 }
func genFn0(p0 int32) uint32 {
    var v4 int64 = -9223372036854775808
    for fi5 := 0; fi5 < 20; fi5++ {
        fmt.Println(v4, p0)
        _ = fi5
    }
    _ = p0
    _ = v4
    return 153
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{127, -9223372036854775808, 128, 32767}
    v3 := make(map[int]bool)
    if (-758 >= 1000) {
        v6 := GenS1{F0: 32767, F1: (("" + "a") + ""), F2: 0}
        var v7 bool = false
        _ = v6
        _ = v7
    }
    var v8 int32 = 766
    var v9 int64 = 1
    fmt.Println(v9, v9)
    if len(v2) > 0 {
        v2[0] = 689
    }
    v8 = (v8 % ((1) | 1))
    var v10 float32 = float32(((((0.1 * (float64(0.1) - float64(((-2.25 + 1e10) / 1e10)))) / -2.25) + 1.5) + (-2.25 / -2.25)))
    for fi11 := 0; fi11 < 11; fi11++ {
        _ = len(v2)
        v12 := GenS2{F0: 237, F1: -2147483648}
        switch fi11 {
        case 970:
            v10 += v10
        default:
            var v13 uint16 = 32767
            _ = v13
        }
        _ = fi11
        _ = v12
    }
    v14 := GenS3{F0: 127}
    var v15 uint16 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
