package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint8 }
func genFn0(p0 int16) int32 {
    p0 -= (p0 - (-(p0)))
    p0 += p0
    switch 32768 {
    case -156:
        if true {
            v7 := []int{1}
            _ = v7
        }
    case 116:
        var v8 float64 = 0.1
        _ = v8
    }
    p0 -= (p0 << ((255) & 63))
    if true {
        v9 := []int{32767, 127, 1, 32768}
        _ = v9
    } else {
        fmt.Println(p0)
        switch 1 {
        case -144:
            var v10 int64 = -1
            _ = v10
        }
        var v11 float32 = 1e10
        _ = v11
    }
    _ = p0
    return 407
}

func main() {
    var v1 string = "\n\t"
    var v2 float32 = (float32(0.0) - float32((1e10 + ((3.14159 - 0.0) - ((0.1 + 1e10) * ((((1e10 / 0.1) * (3.14159 + 1e10)) / 3.14159) / 1.5))))))
    var v3 float32 = float32(((-2.25 / -2.25) * 1.5))
    var v4 uint64 = 4294967295
    v5 := []uint32{}
    v6 := []int{252, 2147483647, -9223372036854775808, 0}
    for fi12 := 0; fi12 < 11; fi12++ {
        v13 := GenS1{F0: 42, F1: v4, F2: 255}
        var v14 int8 = 42
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 uint64 = v4
    var v16 int32 = -1
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 int8 = 0
        switch (fi17 >> ((271) & 63)) {
        case 32767:
            var v19 int = (fi17 / ((fi17) | 1))
            _ = v19
        case -2147483648:
            for fi20 := 0; fi20 < 1; fi20++ {
                var v21 int = (fi20 & (fi17 % ((-138) | 1)))
                _ = fi20
                _ = v21
            }
        default:
            v22 := GenS1{F0: 616, F1: (v15 + v4), F2: 2}
            _ = v22
        }
        _ = fi17
        _ = v18
    }
    var v23 uint32 = 2147483647
    switch 32767 {
    case -161:
        var v24 float64 = -2.25
        _ = v24
    case 42:
        _ = len(v5)
    default:
        v25 := GenS1{F0: 0, F1: v4, F2: 255}
        _ = v25
    }
    if len(v6) > 0 {
        v6[0] = 9223372036854775807
    }
    v23 = (v23 & 21)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
