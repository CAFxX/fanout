package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint; F1 uint64 }
func genFn0(p0 uint64, p1 int8) uint8 {
    v5 := []int{2147483647}
    if ((0.1 == (3.14159 / ((1e10 / float64(0.1)) + (((0.1 * 1e10) / 1e10) - 0.1)))) && false) {
        fmt.Println(p0, p1)
        var v6 uint64 = p0
        if ((!false) && false) {
            if len(v5) > 0 {
                v5[0] = 436
            }
            p0 -= (p0 | p0)
        } else {
            for fi7 := 0; fi7 < 11; fi7++ {
                switch "a" {
                case "\n\t":
                    fmt.Println(v6, v6)
                case "":
                    for fi8 := 0; fi8 < 8; fi8++ {
                        p0 += 4294967295
                        var v9 float64 = 3.14159
                        if false {
                            v9 += v9
                            p1 = (p1 - -26)
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        } else {
                            switch fi8 {
                            case 128:
                                var v10 string = "\n\t"
                                _ = v10
                            case 32768:
                                v6 = (p0 - 0)
                                var v11 uint8 = 1
                                _ = v11
                            case 127:
                                var v12 uint = 2
                                _ = v12
                            }
                            if false {
                                v13 := GenS1{F0: 32767}
                                _ = v13
                            }
                        }
                        _ = fi8
                        _ = v9
                    }
                }
                p1 += -128
                _ = fi7
            }
            v6 -= 0
            var v14 string = ""
            _ = v14
        }
        _ = v6
    }
    _ = len(v5)
    p0 += p0
    _ = p0
    _ = p1
    _ = v5
    return 78
}

func main() {
    v1 := []int{9223372036854775807, -1, 815}
    v2 := []int{127, 4294967295}
    var v3 float64 = (1.5 / -2.25)
    v4 := make(map[string]uint32)
    v3 += v3
    if true {
        v1 = append(v1, 32767)
    } else {
        v3 += v3
        v15 := GenS2{F0: 18446744073709551615, F1: 946}
        _ = v15
    }
    v16 := GenS3{F0: "\n\t"}
    v3 -= (((((v3 + (v3 + 0.0)) / v3) * ((0.1 / v3) / 1e10)) + v3) / v3)
    v17 := GenS3{F0: "\n\t"}
    var v18 float64 = (float64(1.5) + v3)
    v18 = v3
    v18 -= ((v3 + v18) * v3)
    v19 := GenS1{F0: -85}
    if true {
        if true {
            if ((515 != 500) && false) {
                var v20 float32 = (1e10 + (1.5 - 1.5))
                _ = v20
            } else {
                var v21 int8 = 1
                var v22 uint = 0
                _ = v21
                _ = v22
            }
            v3 = v18
            v3 -= (1e10 / 1.0)
        }
        fmt.Println(v18, v18)
    } else {
        switch 42 {
        case 32768:
            v3 = v3
        default:
            for fi23 := 0; fi23 < 14; fi23++ {
                v24 := GenS4{F0: -2147483648, F1: 474}
                v25 := GenS5{F0: 421, F1: 18446744073709551615}
                var v26 string = ("hello" + ("\n\t" + "\n\t"))
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
        }
        fmt.Println(v18)
        v18 = 3.14159
    }
    if len(v2) > 0 {
        v2[0] = 613
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
