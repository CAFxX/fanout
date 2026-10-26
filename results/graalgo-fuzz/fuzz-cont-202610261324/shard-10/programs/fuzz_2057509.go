package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int }
func genFn0(p0 int64) uint {
    p0 = (p0 + (p0 | p0))
    fmt.Println(p0)
    for fi4 := 0; fi4 < 6; fi4++ {
        p0 = (p0 | p0)
        _ = fi4
    }
    _ = p0
    return 974
}
func genFn1(p0 int64, p1 uint16) bool {
    fmt.Println(p1, p1)
    var v5 float32 = -2.25
    v5 = (0.0 - float32(0.1))
    if (!true) {
        for fi6 := 0; fi6 < 14; fi6++ {
            if true {
                v5 = 1e10
            } else {
                v7 := GenS1{F0: -2147483648, F1: (fi6 * (^(fi6))), F2: (fi6 * fi6)}
                var v8 int = int(p0)
                _ = v7
                _ = v8
            }
            _ = fi6
        }
        if ((p0 / (((p0 + -1)) | 1)) <= p0) {
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 bool = (0 == -829)
                for fi11 := 0; fi11 < 4; fi11++ {
                    var v12 float32 = ((0.1 + ((v5 - v5) - v5)) - 1.5)
                    var v13 int32 = 192
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                var v14 uint32 = 65535
                _ = fi9
                _ = v10
                _ = v14
            }
            for fi15 := 0; fi15 < 12; fi15++ {
                v5 = 0.1
                v5 = v5
                _ = fi15
            }
            var v16 uint8 = 127
            _ = v16
        } else {
            p0 -= (p0 * p0)
        }
        for fi17 := 0; fi17 < 6; fi17++ {
            switch 127 {
            case -2147483648:
                fmt.Println(p0, p0)
            default:
                v18 := []int{(fi17 ^ -910)}
                _ = v18
            }
            if false {
                for fi19 := 0; fi19 < 6; fi19++ {
                    p1 -= p1
                    _ = fi19
                }
                fmt.Println(v5)
                v20 := []int{(fi17 << ((genFn0((p0 << ((556) & 63)))) & 63))}
                _ = v20
            }
            _ = fi17
        }
    }
    v21 := []int{32768, -903, -9223372036854775808, 128}
    _ = p0
    _ = p1
    _ = v21
    _ = v5
    return (p1 == (p1 / (((^(p1))) | 1)))
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[uint32]string)
    v3 := []uint32{2}
    for fi22 := 0; fi22 < 1; fi22++ {
        var v23 int64 = 42
        _ = fi22
        _ = v23
    }
    var v24 uint16 = 65535
    _ = len(v3)
    v24 = (v24 / ((v24) | 1))
    if (-2147483648 > -859) {
        var v25 float64 = 1e10
        _ = v25
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        v24 += v24
        switch fi26 {
        case 42:
            var v27 int32 = -2147483648
            _ = v27
        default:
            var v28 uint16 = (v24 | (v24 % (((v24 % ((0) | 1))) | 1)))
            _ = v28
        }
        _ = fi26
    }
    switch 42 {
    case 759:
        v24 += v24
    default:
        var v29 uint32 = 930
        _ = v29
    }
    v2[2147483648] = ""
    for fi30 := 0; fi30 < 13; fi30++ {
        for fi31 := 0; fi31 < 15; fi31++ {
            var v32 int16 = -234
            var v33 int16 = 0
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
}
