package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint16 }
func genFn0() int {
    v4 := []int{}
    v5 := GenS1{F0: 1000}
    var v6 uint8 = 209
    fmt.Println(v6)
    _ = v4
    _ = v5
    _ = v6
    return -653
}
func genFn1(p0 bool, p1 int) bool {
    p1 = (p1 / (((p1 - (^(p1)))) | 1))
    var v7 int16 = 1
    v8 := GenS1{F0: v7}
    if false {
        p1 = (p1 * genFn0())
        p1 = (p1 << ((1) & 63))
        switch "a" {
        case "a":
            if p0 {
                if p0 {
                    for fi9 := 0; fi9 < 5; fi9++ {
                        var v10 uint8 = 147
                        _ = fi9
                        _ = v10
                    }
                    var v11 uint16 = 1
                    v12 := GenS1{F0: (v7 & int16(p1))}
                    _ = v11
                    _ = v12
                }
            } else {
                p1 = 42
                fmt.Println(p0, p1, p0)
                fmt.Println(v7, p1)
            }
        default:
            switch (p1 << ((1) & 63)) {
            case 404:
                v7 -= v7
            default:
                var v13 int = p1
                _ = v13
            }
        }
    }
    v7 = (v7 % ((v7) | 1))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = (-2.25 * float64(3.14159))
    v3 := []int{127, 2147483647, 9223372036854775807}
    v14 := GenS2{F0: -2147483648}
    v2 = v2
    var v15 int = 42
    v15 = v15
    var v16 int32 = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (!(((v15 * v15) <= 730) && true)) {
        for fi17 := 0; fi17 < 2; fi17++ {
            switch fi17 {
            case -1:
                for fi18 := 0; fi18 < 6; fi18++ {
                    if false {
                        v19 := GenS2{F0: v16}
                        _ = v19
                    }
                    _ = fi18
                }
            case -9223372036854775808:
                for fi20 := 0; fi20 < 12; fi20++ {
                    switch fi17 {
                    case -9223372036854775808:
                        v21 := GenS3{F0: 2, F1: 2}
                        _ = v21
                    }
                    v22 := GenS3{F0: 0, F1: 32767}
                    v16 = (v16 - v16)
                    _ = fi20
                    _ = v22
                }
            }
            for fi23 := 0; fi23 < 8; fi23++ {
                v16 = v16
                _ = fi23
            }
            switch genFn0() {
            case 42:
                for fi24 := 0; fi24 < 4; fi24++ {
                    if true {
                        var v25 int32 = v16
                        _ = v25
                    } else {
                        fmt.Println(v15, fi24, v16)
                        fmt.Println(v16)
                    }
                    _ = fi24
                }
            }
            _ = fi17
        }
        for fi26 := 0; fi26 < 13; fi26++ {
            var v27 uint16 = 1
            switch "a" {
            case "\n\t":
                v16 += v16
            }
            fmt.Println(v15)
            _ = fi26
            _ = v27
        }
        if len(v3) > 0 {
            v3[0] = v15
        }
    }
    var v28 uint = 18446744073709551615
    v29 := GenS2{F0: (v16 % (((v16 >> (((v28 ^ 18446744073709551615)) & 63))) | 1))}
    v16 += -2147483648
    v30 := GenS1{F0: 390}
    var v31 int32 = (v16 / ((v16) | 1))
    if ((0.0 <= 3.14159) || (!true)) {
        var v32 uint = (v28 >> (((^(v28))) & 63))
        var v33 uint16 = 1
        v34 := GenS2{F0: (v16 / (((v31 + (^(v31)))) | 1))}
        _ = v32
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 9; fi35++ {
        for fi36 := 0; fi36 < 4; fi36++ {
            v28 = (^(v28))
            _ = fi36
        }
        v3 = append(v3, 32767)
        _ = fi35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
