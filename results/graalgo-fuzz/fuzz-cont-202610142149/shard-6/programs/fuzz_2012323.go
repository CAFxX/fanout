package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint64; F1 uint64; F2 int64 }
func genFn0(p0 bool, p1 int8) uint64 {
    v5 := GenS1{F0: p1, F1: 622}
    switch 2147483647 {
    case 127:
        p1 += 86
    case 128:
        var v6 float64 = (3.14159 * 1e10)
        _ = v6
    case 9223372036854775807:
        for fi7 := 0; fi7 < 13; fi7++ {
            for fi8 := 0; fi8 < 11; fi8++ {
                if (!((1000 >= -32768) && false)) {
                    fmt.Println(fi8, fi8, p0)
                    p1 = p1
                } else {
                    var v9 string = ("\n\t" + "\n\t")
                    if false {
                        switch fi8 {
                        case 2147483647:
                            p1 -= 0
                            v10 := GenS2{F0: 42, F1: 2, F2: 255}
                            _ = v10
                        default:
                            var v11 uint8 = 255
                            _ = v11
                        }
                        switch fi8 {
                        case -667:
                            var v12 uint8 = 127
                            _ = v12
                        default:
                            p1 += p1
                        }
                        if (false && (-2.25 >= -2.25)) {
                            var v13 int = (fi8 / ((146) | 1))
                            v14 := []int{fi8, fi7, (fi8 | (v13 | -70))}
                            _ = v13
                            _ = v14
                        } else {
                            var v15 uint8 = 166
                            p1 -= (p1 | (-(p1)))
                            var v16 int8 = p1
                            v17 := GenS2{F0: 0, F1: (v15 ^ (v15 << ((0) & 63))), F2: 303}
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                    } else {
                        v18 := GenS2{F0: 0, F1: 81, F2: 1}
                        _ = v18
                    }
                    _ = v9
                }
                v19 := []int{fi8}
                var v20 int64 = 1
                _ = fi8
                _ = v19
                _ = v20
            }
            v21 := GenS3{F0: 2, F1: 0, F2: 1}
            var v22 uint = 2
            _ = fi7
            _ = v21
            _ = v22
        }
    default:
        fmt.Println(p0, p0)
    }
    p1 -= p1
    v23 := GenS1{F0: (p1 % ((p1) | 1)), F1: 114}
    _ = p0
    _ = p1
    _ = v23
    _ = v5
    return 42
}
func genFn1(p0 uint16, p1 bool, p2 int64) int32 {
    switch "a" {
    case "foo bar":
        var v24 uint8 = 0
        _ = v24
    case "a":
        var v25 int = 1
        _ = v25
    case "\n\t":
        var v26 int = 128
        _ = v26
    }
    v27 := GenS1{F0: 0, F1: 1}
    v28 := GenS1{F0: 100, F1: 185}
    if p1 {
        switch 4294967295 {
        case 9223372036854775807:
            p1 = p1
        }
        fmt.Println(p1, p1, p1)
        var v29 float32 = 1e10
        _ = v29
    }
    var v30 float64 = ((3.14159 - (0.1 / ((1.5 * 0.0) - float64(1e10)))) * 1.5)
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v28
    _ = v30
    return 0
}

func main() {
    var v1 int8 = -77
    var v2 int16 = 0
    var v3 float64 = (-2.25 + 3.14159)
    var v4 uint = 684
    v31 := []int{}
    var v32 string = "foo bar"
    if len(v31) >= 2 {
        v31 = v31[1:len(v31)]
    }
    v33 := GenS4{F0: v1, F1: 2}
    if (!false) {
        var v34 float64 = v3
        _ = v34
    } else {
        v35 := GenS4{F0: -128, F1: 32767}
        fmt.Println(v1)
        var v36 float64 = v3
        _ = v35
        _ = v36
    }
    v4 = v4
    v4 -= v4
    for fi37 := 0; fi37 < 18; fi37++ {
        v31 = append(v31, (fi37 | 127))
        _ = fi37
    }
    switch 785 {
    case 9223372036854775807:
        if len(v31) > 0 {
            v31[0] = 0
        }
    case 1:
        fmt.Println(v3, v2, v3)
    default:
        v38 := GenS5{F0: 42, F1: 640, F2: -1}
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
