package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
func genFn0(p0 int, p1 int16) int32 {
    fmt.Println(p0, p0)
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 float64 = 3.14159
        if true {
            for fi7 := 0; fi7 < 1; fi7++ {
                var v8 string = "\n\t"
                _ = fi7
                _ = v8
            }
            if true {
                var v9 uint16 = 65535
                p0 = (p0 - p0)
                for fi10 := 0; fi10 < 6; fi10++ {
                    p1 = p1
                    v11 := []int{fi10}
                    p0 -= fi10
                    if len(v11) > 0 {
                        v11[0] = (p0 ^ fi5)
                    }
                    var v12 float32 = 3.14159
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = v9
            }
        }
        var v13 uint8 = 146
        _ = fi5
        _ = v13
        _ = v6
    }
    var v14 int32 = 157
    v14 = (v14 | (v14 * (v14 % ((-705) | 1))))
    _ = p0
    _ = p1
    _ = v14
    return (v14 | 1)
}

func main() {
    var v1 float64 = (3.14159 + (float64((0.1 + 3.14159)) / ((0.0 - float64(-2.25)) - (1e10 + 0.0))))
    var v2 string = ("\n\t" + "")
    var v3 int16 = 0
    var v4 uint64 = 18446744073709551615
    v3 += -411
    if (1 == (v4 * (v4 + 0))) {
        for fi15 := 0; fi15 < 20; fi15++ {
            for fi16 := 0; fi16 < 9; fi16++ {
                switch (fi16 ^ (fi16 ^ int(v4))) {
                case 32767:
                    v17 := []int{}
                    _ = v17
                case -9223372036854775808:
                    var v18 float32 = (float32(1e10) / (1e10 - 3.14159))
                    _ = v18
                default:
                    v19 := []int{(fi15 + int(v4)), (fi16 % ((fi15) | 1)), 635, (-(fi15))}
                    _ = v19
                }
                v4 += (v4 / (((^(v4))) | 1))
                _ = fi16
            }
            for fi20 := 0; fi20 < 6; fi20++ {
                v3 -= (^(v3))
                _ = fi20
            }
            for fi21 := 0; fi21 < 12; fi21++ {
                fmt.Println(v2, v4, v3)
                _ = fi21
            }
            _ = fi15
        }
    }
    var v22 int = -1
    var v23 bool = true
    v1 -= float64((0.0 - (float32(0.0) / 1.5)))
    fmt.Println(v1, v22)
    fmt.Println(v1, v2, v23)
    v22 = v22
    v23 = false
    var v24 uint8 = 127
    v25 := []int{v22, v22}
    var v26 uint32 = 2147483648
    var v27 uint = 866
    if v23 {
        v25 = append(v25, (v22 | (v22 >> ((v27) & 63))))
    }
    if (-128 != 58) {
        var v28 uint8 = (v24 >> (((v27 ^ v27)) & 63))
        v23 = (int8(v4) != int8(v4))
        _ = v28
    }
    if v23 {
        fmt.Println(v2)
        var v29 int64 = 42
        v30 := GenS1{F0: 470, F1: v4}
        _ = v29
        _ = v30
    }
    v4 = (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
