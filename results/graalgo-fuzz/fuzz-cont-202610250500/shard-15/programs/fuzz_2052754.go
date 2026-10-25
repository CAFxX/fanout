package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int16; F2 int8 }
type GenS3 struct { F0 int64; F1 int64; F2 uint16 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 uint16; F1 uint16; F2 uint32 }
func genFn0(p0 int8, p1 int16) uint {
    fmt.Println(p1)
    v6 := GenS1{F0: 127}
    v7 := []int{-287, -543, 9223372036854775807, 127}
    p0 -= (p0 * (p0 ^ -117))
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 419
}
func genFn1(p0 int, p1 int32) int {
    fmt.Println(p0, p1)
    p1 -= -1
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := GenS1{F0: 0}
        for fi10 := 0; fi10 < 16; fi10++ {
            v11 := []int{p0}
            if false {
                var v12 int8 = -128
                var v13 int16 = 0
                v14 := GenS2{F0: true, F1: -32768, F2: (v12 >> ((1) & 63))}
                _ = v12
                _ = v13
                _ = v14
            } else {
                p1 += p1
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = len(v11)
            }
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    return (-(p0))
}

func main() {
    var v1 string = "\n\t"
    var v2 int32 = -627
    v3 := []string{"hello"}
    v4 := []uint32{}
    v5 := []string{"a", v1}
    var v15 int8 = -1
    v2 -= (v2 * v2)
    v16 := GenS1{F0: 0}
    if (!(2 <= 1)) {
        if (true || true) {
            if true {
                _ = len(v5)
            } else {
                v2 += (v2 * (v2 << ((940) & 63)))
                v2 = (v2 & (v2 % ((v2) | 1)))
            }
            if false {
                if (false && true) {
                    v3 = append(v3, "hello")
                    var v17 int32 = (v2 | (v2 >> ((genFn0((v15 * v15), -141)) & 63)))
                    v3 = append(v3, (("" + ("\n\t" + v1)) + ("foo bar" + (v1 + ((("" + v1) + v1) + "")))))
                    _ = v17
                }
                if ((255 > 0) && (42 > 509)) {
                    if (!true) {
                        var v18 int = -1
                        _ = v18
                    }
                    v1 = v1
                    var v19 int16 = int16(v15)
                    _ = v19
                } else {
                    switch 127 {
                    case -967:
                        v1 = v1
                    case 0:
                        v5 = append(v5, v1)
                    case 42:
                        switch 4294967295 {
                        case 128:
                            v20 := GenS3{F0: 1, F1: 0, F2: 65535}
                            _ = v20
                        case 9223372036854775807:
                            var v21 float32 = -2.25
                            _ = v21
                        default:
                            var v22 int = 0
                            _ = v22
                        }
                    }
                    v15 += v15
                    v1 = "\n\t"
                }
                for fi23 := 0; fi23 < 4; fi23++ {
                    var v24 int = 9223372036854775807
                    _ = fi23
                    _ = v24
                }
            } else {
                if false {
                    v15 = v15
                }
            }
        }
    } else {
        var v25 int8 = (v15 / ((0) | 1))
        var v26 float64 = float64(0.0)
        _ = v25
        _ = v26
    }
    switch -1 {
    case -9223372036854775808:
        if true {
            var v27 int32 = (v2 | (-(v2)))
            switch 791 {
            case -142:
                v27 = 0
            default:
                v28 := GenS1{F0: 13}
                _ = v28
            }
            _ = v27
        } else {
            v29 := GenS4{F0: true, F1: (0 < 2)}
            var v30 int64 = 0
            v2 -= (v2 / ((-666) | 1))
            _ = v29
            _ = v30
        }
    case 755:
        v1 = v1
    case 9223372036854775807:
        v4 = append(v4, 957)
    default:
        v31 := GenS1{F0: 255}
        _ = v31
    }
    fmt.Println(v1, v15)
    var v32 int64 = -1
    if (-2.25 <= 0.0) {
        v33 := GenS5{F0: 70, F1: 1, F2: 0}
        _ = v33
    }
    var v34 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v4
    _ = v5
}
