package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint64 }
func genFn0() uint8 {
    var v5 int16 = 42
    var v6 float64 = 0.0
    v5 -= v5
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 int16 = -33
    var v2 int64 = 9223372036854775807
    var v3 uint = 95
    var v4 uint8 = 255
    v7 := []int{-587}
    if true {
        v4 += v4
        for fi8 := 0; fi8 < 2; fi8++ {
            v2 += -1
            _ = fi8
        }
    }
    var v9 int64 = 1
    for fi10 := 0; fi10 < 4; fi10++ {
        v9 = (v9 ^ (v2 >> (((v3 % ((v3) | 1))) & 63)))
        fmt.Println(v1, v1, v2)
        _ = fi10
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := GenS1{F0: 60}
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 uint8 = uint8(v3)
            fmt.Println(v3)
            _ = fi13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    var v15 int8 = -128
    v16 := GenS1{F0: 64}
    fmt.Println(v15, v3)
    v4 = v4
    var v17 uint16 = 65535
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 uint8 = (v4 * uint8(v3))
        fmt.Println(v1)
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi18
        _ = v19
    }
    var v20 uint32 = 1
    if true {
        v2 = (-(v2))
        var v21 uint16 = (v17 * v17)
        v17 = (v17 / (((v17 * (v21 & v21))) | 1))
        _ = v21
    } else {
        v15 = 113
        v20 = (v20 / ((v20) | 1))
        if (true || false) {
            _ = len(v7)
            v22 := GenS2{F0: false, F1: 540}
            _ = v22
        } else {
            if len(v7) > 0 {
                v7[0] = 32768
            }
            v9 -= (v9 % (((-(v2))) | 1))
        }
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        v24 := GenS2{F0: (v17 != v17), F1: 552}
        v1 -= (v1 & (v1 & (^(v1))))
        _ = len(v7)
        _ = fi23
        _ = v24
    }
    v2 -= 9223372036854775807
    v3 += 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
