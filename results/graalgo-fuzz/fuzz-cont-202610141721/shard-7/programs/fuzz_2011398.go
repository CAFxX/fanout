package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int16 }
func genFn0() uint8 {
    var v6 float64 = ((1.5 * (0.0 * -2.25)) + (float64(((0.1 + float32(((((-2.25 - (1e10 * -2.25)) - 0.1) * ((1e10 + ((float64((1e10 * 1e10)) - (-2.25 * float64(float32(3.14159)))) - -2.25)) + 0.1)) * -2.25))) - 0.1)) - (-2.25 + float64(1.5))))
    v7 := []int{128, -80}
    switch 0 {
    case 4294967295:
        v6 -= (v6 - 1e10)
    case -9223372036854775808:
        v8 := GenS1{F0: 2147483647, F1: 42}
        _ = v8
    case 32767:
        for fi9 := 0; fi9 < 11; fi9++ {
            v6 = ((((v6 + (float64(0.1) - -2.25)) * ((((float64(0.0) + (0.0 - float64((1e10 + 1e10)))) + v6) - 0.0) / v6)) / v6) * v6)
            _ = fi9
        }
    }
    v6 = 0.1
    v6 += v6
    _ = v6
    _ = v7
    return 34
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{}
    var v3 bool = false
    var v4 uint32 = 65535
    v5 := make(map[int]uint32)
    if (!(!v3)) {
        var v10 int8 = -21
        v11 := GenS2{F0: v4}
        v12 := GenS2{F0: v4}
        _ = v10
        _ = v11
        _ = v12
    } else {
        fmt.Println(v3, v3)
    }
    var v13 uint16 = 65535
    v4 = (v4 + (v4 << ((65535) & 63)))
    var v14 uint64 = 642
    var v15 uint16 = 862
    v14 = (v14 * v14)
    v4 += 2147483647
    var v16 string = "foo bar"
    fmt.Println(v4, v15, v14)
    v17 := GenS3{F0: 995, F1: -32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
