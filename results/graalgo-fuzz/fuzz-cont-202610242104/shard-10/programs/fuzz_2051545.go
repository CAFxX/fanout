package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int; F1 uint16; F2 int16 }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int32; F1 int16 }
func genFn0(p0 int32) uint32 {
    p0 = -1
    if (true || false) {
        v7 := []int{2147483647, 42, 536, 127}
        v7 = append(v7, 2147483647)
        _ = v7
    } else {
        if (32767 == 1) {
            p0 -= (p0 | p0)
        } else {
            for fi8 := 0; fi8 < 17; fi8++ {
                var v9 bool = true
                var v10 int64 = 0
                p0 -= p0
                _ = fi8
                _ = v10
                _ = v9
            }
            var v11 int = 32768
            _ = v11
        }
    }
    v12 := GenS1{F0: "", F1: 42}
    var v13 float64 = 0.1
    v14 := GenS2{F0: 128, F1: 111, F2: 967}
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return 2147483647
}
func genFn1() int64 {
    v15 := []int{}
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 int = (fi16 * (fi16 - 32767))
        _ = fi16
        _ = v17
    }
    var v18 uint32 = 65535
    _ = v15
    _ = v18
    return 9223372036854775807
}

func main() {
    var v1 uint = 1
    v2 := []int{-719, 298, 4294967295, 42}
    var v3 uint16 = 1
    var v4 float64 = (0.0 * ((0.1 + float64(1.5)) - 0.1))
    var v5 uint32 = 2
    var v6 uint = 4294967295
    if true {
        var v19 int16 = 0
        _ = v19
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        if len(v2) > 0 {
            v2[0] = (^(fi20))
        }
        if len(v2) > 0 {
            v2[0] = fi20
        }
        var v21 int32 = -335
        _ = fi20
        _ = v21
    }
    _ = len(v2)
    v3 += (v3 | 121)
    switch -260 {
    case 1:
        v2 = append(v2, 592)
    case -2147483648:
        fmt.Println(v5)
    case 2147483647:
        v22 := GenS3{F0: 586, F1: (v3 * (v3 % (((v3 + (v3 | v3))) | 1))), F2: 1}
        _ = v22
    }
    v4 = (((1.5 + v4) + 3.14159) / ((0.0 / float64((0.1 / 0.1))) - ((v4 * 1e10) / v4)))
    fmt.Println(v6)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v2 = append(v2, -492)
    for fi23 := 0; fi23 < 12; fi23++ {
        v24 := GenS2{F0: fi23, F1: 32767, F2: 42}
        if true {
            var v25 int = (-(fi23))
            for fi26 := 0; fi26 < 8; fi26++ {
                v3 = (^(v3))
                if (1 <= 2147483647) {
                    switch fi26 {
                    case -9223372036854775808:
                        if (true && ((v3 >> ((v6) & 63)) <= v3)) {
                            if (!false) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if false {
                                    switch "a" {
                                    case "a":
                                        v4 = 0.0
                                    }
                                } else {
                                    var v27 float64 = (float64(1.5) * ((v4 * (v4 / float64((3.14159 * 1.5)))) * v4))
                                    _ = v27
                                }
                                var v28 int8 = -80
                                _ = v28
                            }
                            switch "foo bar" {
                            case "":
                                v1 = 1
                            case "foo bar":
                                v4 = 1e10
                            default:
                                fmt.Println(v4)
                            }
                        }
                    case -1:
                        v1 += v1
                    default:
                        var v29 bool = (1 > -115)
                        _ = v29
                    }
                    var v30 int8 = -118
                    v1 = 65535
                    fmt.Println(v5, v30, v25)
                    _ = v30
                }
                fmt.Println(v4, fi23)
                _ = fi26
            }
            var v31 uint8 = 1
            _ = v25
            _ = v31
        } else {
            for fi32 := 0; fi32 < 18; fi32++ {
                v33 := GenS2{F0: fi23, F1: 430, F2: 735}
                _ = fi32
                _ = v33
            }
            if ((!(("a" + "a") <= (("\n\t" + "\n\t") + "foo bar"))) && true) {
                var v34 string = "foo bar"
                _ = v34
            } else {
                v4 = v4
            }
        }
        _ = fi23
        _ = v24
    }
    fmt.Println(v3, v1)
    var v35 int8 = 127
    var v36 int64 = int64(v5)
    v37 := GenS4{F0: 1, F1: 0}
    v3 += v3
    if (v6 < (v6 % (((v6 * (v6 << (((v1 << ((904) & 63))) & 63)))) | 1))) {
        v2 = append(v2, 901)
        var v38 int32 = 609
        _ = v38
    }
    var v39 uint16 = (v3 & (v3 & (v3 * (v3 % (((v3 + v3)) | 1)))))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
