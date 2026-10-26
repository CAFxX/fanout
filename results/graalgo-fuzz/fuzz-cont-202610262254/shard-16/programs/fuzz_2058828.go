package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 bool }
func genFn0(p0 int16) uint {
    p0 += p0
    if (0 > 1) {
        v5 := GenS1{F0: 1, F1: true}
        if (18446744073709551615 != 0) {
            var v6 uint16 = 255
            v6 = v6
            if false {
                v6 -= 49
                if false {
                    for fi7 := 0; fi7 < 17; fi7++ {
                        fmt.Println(v6, fi7)
                        var v8 bool = (!(false && (902 > 2)))
                        _ = fi7
                        _ = v8
                    }
                    for fi9 := 0; fi9 < 5; fi9++ {
                        v6 -= (v6 >> ((1) & 63))
                        switch "" {
                        case "a":
                            v6 += v6
                        case "foo bar":
                            if true {
                                var v10 uint32 = uint32(v6)
                                p0 += (p0 % (((-(p0))) | 1))
                                _ = v10
                            } else {
                                v11 := []int{-9223372036854775808, int(v6), fi9, fi9}
                                v12 := GenS1{F0: (p0 / (((^(p0))) | 1)), F1: true}
                                _ = v11
                                _ = v12
                            }
                        case "":
                            if ((((1.5 - 1.5) / (float32(0.1) / 3.14159)) + 3.14159) > float32(1.5)) {
                                fmt.Println(p0, fi9)
                                v6 += 2
                            } else {
                                v13 := []int{821, fi9}
                                var v14 int = -9223372036854775808
                                _ = v13
                                _ = v14
                            }
                        }
                        switch fi9 {
                        case -360:
                            for fi15 := 0; fi15 < 4; fi15++ {
                                var v16 bool = (!(!((!(42 <= 2)) && false)))
                                _ = fi15
                                _ = v16
                            }
                        default:
                            for fi17 := 0; fi17 < 20; fi17++ {
                                v18 := []int{}
                                if len(v18) > 0 {
                                    v18[0] = 2147483647
                                }
                                v6 = v6
                                _ = fi17
                                _ = v18
                            }
                        }
                        _ = fi9
                    }
                }
            }
            _ = v6
        } else {
            v19 := GenS1{F0: (p0 - p0), F1: true}
            p0 += p0
            _ = v19
        }
        _ = v5
    }
    v20 := GenS1{F0: -32768, F1: false}
    _ = p0
    _ = v20
    return 315
}
func genFn1(p0 bool, p1 bool) string {
    p0 = false
    v21 := GenS1{F0: -1, F1: (727 < 255)}
    p0 = ((!(-1 > 0)) || ((65535 > 255) && (!true)))
    _ = p0
    _ = p1
    _ = v21
    return "a"
}

func main() {
    v1 := []uint32{65535, 145}
    v2 := make(map[uint32]bool)
    var v3 int32 = -1
    var v4 float64 = 0.1
    if true {
        v22 := GenS1{F0: -766, F1: ((!(42 <= 255)) || (v3 != (v3 ^ v3)))}
        _ = v22
    }
    _ = len(v1)
    v23 := GenS2{F0: 508, F1: 1}
    _ = len(v1)
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 uint16 = 1
        _ = fi24
        _ = v25
    }
    delete(v2, 2147483647)
    _ = len(v1)
    var v26 int16 = 0
    v27 := GenS3{F0: false}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
