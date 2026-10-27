package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint8; F1 int; F2 uint32 }
type GenS4 struct { F0 int32; F1 uint32; F2 int16 }
func genFn0(p0 uint8) uint16 {
    var v6 uint8 = (p0 / (((p0 + (p0 & p0))) | 1))
    fmt.Println(v6, v6)
    v6 = (^(v6))
    for fi7 := 0; fi7 < 19; fi7++ {
        p0 += v6
        switch fi7 {
        case -346:
            var v8 float32 = (0.1 / (0.1 - ((0.1 + (float32(-2.25) * ((1.5 * 3.14159) / 0.1))) * ((1e10 - 3.14159) * (-2.25 * 1.5)))))
            _ = v8
        case 958:
            v6 -= (v6 - (^(p0)))
        case 436:
            if (!(1.5 > ((1e10 / float64(0.1)) + (0.0 - 0.1)))) {
                v9 := []int{fi7, (fi7 % (((fi7 + (fi7 | fi7))) | 1)), 32767, fi7}
                fmt.Println(p0)
                v6 -= 0
                _ = v9
            } else {
                v10 := GenS1{F0: 402}
                _ = v10
            }
        default:
            p0 -= v6
        }
        switch "a" {
        case "":
            fmt.Println(v6, fi7)
        case "\n\t":
            p0 += 244
        case "a":
            if (false || (!(!false))) {
                var v11 uint8 = v6
                if true {
                    v12 := GenS2{F0: 0, F1: 18446744073709551615, F2: 310}
                    switch "hello" {
                    case "\n\t":
                        var v13 uint8 = (p0 - (v11 | 9))
                        _ = v13
                    case "a":
                        for fi14 := 0; fi14 < 16; fi14++ {
                            v15 := GenS3{F0: v11, F1: 2147483647, F2: 4294967295}
                            _ = fi14
                            _ = v15
                        }
                    default:
                        for fi16 := 0; fi16 < 13; fi16++ {
                            if ((^(p0)) == (v11 ^ 255)) {
                                v17 := []int{}
                                v17 = append(v17, fi7)
                                _ = v17
                            }
                            _ = fi16
                        }
                    }
                    if true {
                        v18 := []int{}
                        fmt.Println(v11, v11, fi7)
                        _ = v18
                    }
                    _ = v12
                }
                fmt.Println(fi7, v6, v6)
                _ = v11
            }
        }
        _ = fi7
    }
    _ = p0
    _ = v6
    return 1
}
func genFn1() uint8 {
    var v19 float32 = (1.5 + (3.14159 * 1e10))
    v20 := GenS2{F0: 42, F1: 4294967295, F2: 1}
    if false {
        v19 -= (1e10 * v19)
        for fi21 := 0; fi21 < 3; fi21++ {
            v19 += (0.0 * (((1e10 / v19) / (((v19 * v19) * 0.1) - (1.5 - v19))) * v19))
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 int64 = -11
                if (!false) {
                    v24 := GenS3{F0: 42, F1: fi22, F2: 4294967295}
                    fmt.Println(fi22)
                    _ = v24
                } else {
                    fmt.Println(fi21, v23)
                }
                switch ("hello" + "foo bar") {
                case "hello":
                    if ((56 >= 1) && true) {
                        v19 = v19
                    } else {
                        v25 := GenS2{F0: 2147483648, F1: 446, F2: 266}
                        if (fi22 < fi22) {
                            var v26 int16 = 32767
                            _ = v26
                        } else {
                            var v27 float32 = (0.1 + 1.5)
                            fmt.Println(v23)
                            v28 := GenS4{F0: 230, F1: 754, F2: 1000}
                            _ = v27
                            _ = v28
                        }
                        _ = v25
                    }
                case "":
                    v23 = v23
                default:
                    v19 = (v19 + (v19 - v19))
                }
                _ = fi22
                _ = v23
            }
            v19 += 1.5
            _ = fi21
        }
    } else {
        fmt.Println(v19)
    }
    _ = v19
    _ = v20
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 3.14159
    var v3 bool = (0 <= 0)
    var v4 uint8 = 1
    var v5 bool = (3 > 4294967295)
    v5 = false
    v29 := []int{42, -409, 128}
    v30 := GenS4{F0: -536, F1: 4294967295, F2: -32768}
    v3 = (v5 || (!v5))
    var v31 float32 = (0.0 / v2)
    if len(v29) > 0 {
        v29[0] = -531
    }
    switch -9223372036854775808 {
    case 9223372036854775807:
        var v32 int8 = 12
        _ = v32
    case 2147483647:
        v1[""] = 371
    }
    switch 342 {
    case 9223372036854775807:
        var v33 uint32 = 2
        _ = v33
    }
    var v34 float64 = float64(v31)
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v5
}
