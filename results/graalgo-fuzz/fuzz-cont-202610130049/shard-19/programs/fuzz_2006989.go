package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint16; F2 bool }
func genFn0(p0 uint, p1 uint8) uint64 {
    fmt.Println(p0, p0)
    p0 = (p0 | (p0 * p0))
    switch 32768 {
    case 0:
        if (42 == 18446744073709551615) {
            for fi6 := 0; fi6 < 4; fi6++ {
                var v7 int16 = -1
                fmt.Println(p1)
                if false {
                    p0 += (p0 + p0)
                }
                _ = fi6
                _ = v7
            }
        } else {
            v8 := []int{127, 0, -293, 1}
            _ = v8
        }
    case 167:
        v9 := GenS1{F0: -32768}
        _ = v9
    case 32768:
        switch 659 {
        case 126:
            for fi10 := 0; fi10 < 12; fi10++ {
                p1 = p1
                var v11 uint = (p0 % ((p0) | 1))
                v11 = (p0 << (((p0 ^ v11)) & 63))
                _ = fi10
                _ = v11
            }
        }
    }
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 float64 = 0.1
    var v2 int64 = -612
    v3 := []uint32{}
    var v4 float32 = 0.1
    var v5 int = 2147483647
    if (-128 <= -1) {
        for fi12 := 0; fi12 < 11; fi12++ {
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 int8 = -77
                if (2147483647 != 0) {
                    v15 := GenS1{F0: -236}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v16 := GenS2{F0: -2147483648, F1: 32767, F2: (true || false)}
                    _ = v15
                    _ = v16
                } else {
                    v4 = 0.0
                }
                fmt.Println(fi12, v5, v2)
                var v17 float64 = (v1 - v1)
                _ = fi13
                _ = v14
                _ = v17
            }
            _ = fi12
        }
        var v18 uint16 = 938
        v1 = v1
        _ = v18
    } else {
        fmt.Println(v5, v5, v4)
    }
    var v19 float32 = v4
    v5 = v5
    var v20 uint32 = 325
    v21 := GenS2{F0: 1, F1: 1, F2: (!(!false))}
    v22 := GenS1{F0: -613}
    v19 += (float32(((v1 * (1.5 * v1)) + ((1.5 * (float64(1.5) + (((1.5 / 1.0) / (float64(v19) + (((-2.25 / 0.1) - v1) - ((-2.25 - (v1 + v1)) - 1e10)))) - v1))) + v1))) - -2.25)
    var v23 uint = 4294967295
    v4 -= (1.5 - ((float32(0.1) + (v19 * 1.5)) / float32(3.14159)))
    v19 = 0.1
    v24 := GenS1{F0: 237}
    fmt.Println(v2)
    var v25 float64 = v1
    if len(v3) > 0 {
        v3[0] = (v20 + (^(v20)))
    }
    if ((!(((v20 <= 1) || false) && false)) || false) {
        v5 = 1
        if len(v3) > 0 {
            v3[0] = (^(v20))
        }
    } else {
        var v26 int32 = -800
        _ = v26
    }
    var v27 uint32 = (v20 / ((v20) | 1))
    for fi28 := 0; fi28 < 4; fi28++ {
        for fi29 := 0; fi29 < 12; fi29++ {
            v25 = v25
            _ = fi29
        }
        fmt.Println(v20)
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
