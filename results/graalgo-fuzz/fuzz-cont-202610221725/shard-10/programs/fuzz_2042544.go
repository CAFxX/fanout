package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint }
type GenS2 struct { F0 uint8; F1 bool; F2 int16 }
func genFn0(p0 uint, p1 string, p2 bool) uint {
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := GenS1{F0: (fi6 & (fi6 % (((fi6 / ((4294967295) | 1))) | 1))), F1: 1, F2: p0}
        _ = fi6
        _ = v7
    }
    var v8 uint32 = 1
    v9 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return uint(v8)
}

func main() {
    v1 := []int{32768, 32767, 367, 4294967295}
    var v2 int8 = 0
    v3 := make(map[string]uint32)
    var v4 float32 = float32(0.1)
    var v5 float64 = ((float64((1e10 / float32(-2.25))) / float64(v4)) - 0.1)
    v4 -= 1.5
    if (872 == 9223372036854775807) {
        var v10 int = 32768
        if (true || true) {
            v4 -= ((v4 - (float32(-2.25) / -2.25)) * v4)
            var v11 uint8 = 39
            for fi12 := 0; fi12 < 10; fi12++ {
                v13 := GenS1{F0: (^(fi12)), F1: 98, F2: genFn0(genFn0(85, ("hello" + ("" + "hello")), true), "foo bar", false)}
                v5 = (float64((v4 / ((3.14159 + v4) + v4))) + (1.5 * 3.14159))
                _ = fi12
                _ = v13
            }
            _ = v11
        }
        var v14 bool = (32767 < 42)
        _ = v10
        _ = v14
    }
    v4 = float32(v5)
    v4 += v4
    var v15 uint = 555
    for fi16 := 0; fi16 < 3; fi16++ {
        v2 = (v2 << ((0) & 63))
        for fi17 := 0; fi17 < 7; fi17++ {
            fmt.Println(v2, v2, v5)
            _ = fi17
        }
        v18 := GenS1{F0: 0, F1: (v15 * (v15 - v15)), F2: v15}
        _ = fi16
        _ = v18
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 bool = true
        _ = fi19
        _ = v20
    }
    var v21 float32 = (0.0 - v4)
    v5 = ((3.14159 * float64((1.5 * ((((v4 * float32(3.14159)) / ((v4 + (3.14159 / v21)) / v4)) - 1.5) + float32(v5))))) / -2.25)
    switch -9223372036854775808 {
    case -655:
        _ = len(v3)
    default:
        v22 := GenS2{F0: 127, F1: false, F2: 32767}
        _ = v22
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        v4 += v4
        var v24 uint = v15
        _ = len(v1)
        _ = fi23
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
