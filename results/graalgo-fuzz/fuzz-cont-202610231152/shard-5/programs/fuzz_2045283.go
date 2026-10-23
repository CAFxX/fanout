package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int8; F1 int16 }
func genFn0(p0 uint64, p1 int32, p2 string) int8 {
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 float64 = (0.1 / (3.14159 + 1e10))
        fmt.Println(fi6, fi6, v7)
        _ = fi6
        _ = v7
    }
    switch 127 {
    case 32767:
        p0 -= p0
    case 396:
        v8 := []int{}
        _ = v8
    default:
        var v9 int8 = 1
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 127
}
func genFn1(p0 bool, p1 uint64, p2 uint16) uint32 {
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 float32 = 1.5
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 float64 = 1.5
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 float32 = 0.1
            p1 += p1
            _ = fi14
            _ = v15
        }
        _ = fi10
        _ = v11
    }
    var v16 uint32 = 2
    var v17 bool = ((792 != -4) && p0)
    for fi18 := 0; fi18 < 1; fi18++ {
        if (!(255 != 127)) {
            if p0 {
                p0 = false
                for fi19 := 0; fi19 < 17; fi19++ {
                    v20 := GenS1{F0: int32(fi18), F1: v17}
                    var v21 float32 = float32(0.1)
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
        }
        p0 = p0
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return v16
}

func main() {
    v1 := []int{-9223372036854775808, 32768, 684}
    v2 := []int{32767, -9223372036854775808, 59}
    var v3 int16 = 1000
    var v4 int16 = v3
    var v5 uint64 = 858
    var v22 float32 = 1.5
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v23 := GenS2{F0: -1, F1: (v3 + (v3 ^ -32768))}
    v3 += v3
    switch -827 {
    case -1:
        v5 -= 942
    case 0:
        var v24 uint16 = 42
        _ = v24
    case 127:
        var v25 uint64 = (v5 << ((255) & 63))
        _ = v25
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        v27 := GenS1{F0: -2147483648, F1: (0.1 <= 1e10)}
        _ = fi26
        _ = v27
    }
    v4 -= (v3 >> ((4294967295) & 63))
    v28 := GenS1{F0: -2147483648, F1: (!false)}
    v22 -= v22
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
