package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int32 }
func genFn0(p0 uint8) uint32 {
    p0 += p0
    var v5 uint64 = 2
    if true {
        v5 -= v5
        v6 := GenS1{F0: 2147483648}
        _ = v6
    }
    v5 += v5
    if true {
        var v7 bool = false
        v7 = (1000 > 1000)
        if (128 < 35) {
            fmt.Println(v7)
            fmt.Println(v5, v7, p0)
            if (!true) {
                v5 -= (v5 + v5)
                v8 := GenS1{F0: 124}
                for fi9 := 0; fi9 < 16; fi9++ {
                    for fi10 := 0; fi10 < 5; fi10++ {
                        v11 := GenS2{F0: (p0 - (p0 * uint8(v5))), F1: 2147483647}
                        v12 := GenS1{F0: 0}
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    p0 -= p0
                    _ = fi9
                }
                _ = v8
            }
        }
        _ = v7
    }
    _ = p0
    _ = v5
    return 4294967295
}

func main() {
    var v1 int64 = 42
    var v2 float64 = 1e10
    var v3 uint8 = 255
    v4 := []uint32{}
    var v13 int64 = v1
    if len(v4) > 0 {
        v4[0] = 65535
    }
    fmt.Println(v1)
    fmt.Println(v1)
    for fi14 := 0; fi14 < 18; fi14++ {
        if false {
            fmt.Println(v1)
            fmt.Println(v3)
        }
        var v15 string = "a"
        _ = fi14
        _ = v15
    }
    v4 = append(v4, 65535)
    _ = len(v4)
    v1 = 0
    v1 = (v1 + (v1 / ((v13) | 1)))
    if ((true && true) || false) {
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 uint = 4294967295
            var v18 uint8 = v3
            _ = fi16
            _ = v17
            _ = v18
        }
        v19 := GenS1{F0: 2147483648}
        v20 := GenS2{F0: v3, F1: -2}
        _ = v19
        _ = v20
    }
    v13 = (-(v1))
    var v21 int = -2147483648
    switch ("foo bar" + "hello") {
    case "hello":
        v2 -= ((v2 * ((0.0 / 1.0) * v2)) + (v2 / (float64((-2.25 / 1e10)) - v2)))
    default:
        var v22 int16 = -271
        _ = v22
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 int = fi23
        var v25 int = 821
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 uint = 255
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v4
}
