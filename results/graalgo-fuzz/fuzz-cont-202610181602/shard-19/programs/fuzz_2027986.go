package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int64; F1 uint8; F2 int64 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint; F1 uint8; F2 int }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS8 struct { F0 uint }
func genFn0(p0 int64) uint32 {
    v6 := []int{0, -386, 2147483647, 307}
    var v7 int = 4294967295
    _ = p0
    _ = v6
    _ = v7
    return 22
}
func genFn1(p0 uint32, p1 int) string {
    for fi8 := 0; fi8 < 3; fi8++ {
        fmt.Println(p1, fi8, fi8)
        switch fi8 {
        case 4294967295:
            fmt.Println(fi8, p0)
        case -642:
            if ((0 <= 127) || (-391 == -1)) {
                fmt.Println(fi8, fi8)
                p1 = (-(fi8))
            } else {
                var v9 uint = 255
                _ = v9
            }
        default:
            switch fi8 {
            case 0:
                v10 := GenS1{F0: (!((!(100 == 127)) && (0 <= 1))), F1: false}
                _ = v10
            default:
                if true {
                    if true {
                        v11 := GenS2{F0: -9223372036854775808, F1: 29, F2: -481}
                        if (42 != -87) {
                            p0 -= p0
                            if true {
                                switch (p1 - (p1 + fi8)) {
                                case -615:
                                    p0 += (p0 * genFn0(221))
                                case -1:
                                    p1 += (p1 / (((^(p1))) | 1))
                                }
                                p0 -= p0
                                if (!((false || (1 <= -852)) || true)) {
                                    p0 -= p0
                                    v12 := GenS3{F0: 602, F1: 0}
                                    v13 := GenS1{F0: (true && true), F1: (!(2 == 0))}
                                    _ = v12
                                    _ = v13
                                }
                            }
                        } else {
                            for fi14 := 0; fi14 < 3; fi14++ {
                                var v15 int16 = 1000
                                fmt.Println(p1, fi14)
                                v16 := GenS3{F0: p1, F1: 444}
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                            var v17 uint = 0
                            v18 := GenS4{F0: 1, F1: 2, F2: (fi8 >> (((v17 ^ 255)) & 63))}
                            _ = v17
                            _ = v18
                        }
                        _ = v11
                    }
                } else {
                    v19 := []int{}
                    _ = v19
                }
            }
        }
        _ = fi8
    }
    var v20 uint8 = 255
    if false {
        v21 := GenS2{F0: 9223372036854775807, F1: (v20 + (v20 - 128)), F2: 1}
        var v22 float32 = ((((3.14159 * (0.0 / (1.5 / -2.25))) / (1e10 - -2.25)) - 0.0) * 1.5)
        _ = v21
        _ = v22
    } else {
        switch 136 {
        case 2147483647:
            v23 := []int{-46, p1, 465}
            _ = v23
        default:
            var v24 uint16 = 42
            _ = v24
        }
        v25 := []int{p1, p1}
        for fi26 := 0; fi26 < 16; fi26++ {
            for fi27 := 0; fi27 < 15; fi27++ {
                if (1 <= 938) {
                    _ = len(v25)
                    var v28 int8 = 42
                    _ = v28
                }
                _ = fi27
            }
            fmt.Println(p1, v20)
            _ = fi26
        }
        _ = v25
    }
    _ = p0
    _ = p1
    _ = v20
    return "foo bar"
}

func main() {
    var v1 int8 = 79
    var v2 float64 = 1e10
    var v3 int = 128
    var v4 uint64 = 1
    v5 := []uint32{764, 1}
    fmt.Println(v2, v1, v4)
    var v29 int32 = 824
    if true {
        var v30 float64 = 1.5
        _ = v30
    }
    var v31 float64 = v2
    if false {
        if (v3 == v3) {
            v32 := GenS2{F0: 567, F1: 128, F2: -557}
            var v33 uint = 65535
            _ = v32
            _ = v33
        }
        if false {
            v34 := GenS5{F0: 2, F1: -32768}
            var v35 uint32 = 42
            _ = v34
            _ = v35
        }
    } else {
        for fi36 := 0; fi36 < 17; fi36++ {
            if false {
                v4 = v4
                v2 = float64(1e10)
            }
            var v37 uint64 = (v4 & v4)
            _ = fi36
            _ = v37
        }
        var v38 int64 = -9223372036854775808
        _ = v38
    }
    if (32767 == 1000) {
        v39 := GenS2{F0: 9223372036854775807, F1: 34, F2: 211}
        if len(v5) > 0 {
            v5[0] = 42
        }
        v40 := GenS2{F0: 0, F1: 23, F2: 0}
        _ = v39
        _ = v40
    } else {
        v4 = v4
        v2 -= ((0.1 / (((v31 * v2) * 1.5) + 1e10)) * v2)
    }
    v31 = (v31 + -2.25)
    if false {
        v41 := GenS3{F0: (v3 + (v3 | v3)), F1: 240}
        if false {
            for fi42 := 0; fi42 < 16; fi42++ {
                v29 -= 42
                _ = len(v5)
                v43 := GenS6{F0: (^(v3))}
                _ = fi42
                _ = v43
            }
        }
        var v44 string = ("" + "\n\t")
        _ = v41
        _ = v44
    }
    v31 = ((((v2 + (-2.25 / (v2 + v31))) - (v2 / -2.25)) * float64((1.5 - float32(v2)))) / (float64(0.0) + v31))
    switch (("a" + "\n\t") + "a") {
    case "\n\t":
        switch 9223372036854775807 {
        case 42:
            v2 -= 0.1
        case 58:
            v45 := GenS7{F0: true, F1: 1, F2: 0}
            _ = v45
        case 32767:
            fmt.Println(v1, v4)
        default:
            if ((0 > v4) || true) {
                switch v3 {
                case 4294967295:
                    if true {
                        v31 -= ((v31 - 1e10) - (v2 + (float64((1.5 * (((0.0 / 1.5) - (0.1 / -2.25)) - -2.25))) + (v31 + (float64(1e10) + 3.14159)))))
                    } else {
                        var v46 float32 = 1e10
                        for fi47 := 0; fi47 < 13; fi47++ {
                            _ = len(v5)
                            _ = fi47
                        }
                        _ = v46
                    }
                case 32767:
                    for fi48 := 0; fi48 < 17; fi48++ {
                        var v49 string = ""
                        fmt.Println(fi48, v2, v1)
                        if (927 >= 0) {
                            if ((^(v4)) == (^(v4))) {
                                var v50 bool = true
                                var v51 float64 = (-2.25 - v2)
                                _ = v50
                                _ = v51
                            }
                        }
                        _ = fi48
                        _ = v49
                    }
                }
                switch (v3 + v3) {
                case 0:
                    if (!false) {
                        v4 -= (v4 / ((1) | 1))
                        var v52 bool = ((42 > 1) && (!true))
                        _ = v52
                    }
                case -9223372036854775808:
                    for fi53 := 0; fi53 < 7; fi53++ {
                        if len(v5) > 0 {
                            v5[0] = 362
                        }
                        _ = fi53
                    }
                default:
                    v31 = (0.1 * (1.5 * float64(1.5)))
                }
                if (false && (-1 >= -9223372036854775808)) {
                    var v54 uint = 983
                    _ = v54
                }
            } else {
                var v55 int32 = v29
                v29 -= v29
                _ = v55
            }
        }
    case "":
        var v56 uint32 = genFn0(-9223372036854775808)
        _ = v56
    case "foo bar":
        if true {
            switch v3 {
            case 127:
                if false {
                    var v57 int16 = 1
                    _ = v57
                } else {
                    if false {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v58 uint16 = 255
                        _ = v58
                    } else {
                        for fi59 := 0; fi59 < 9; fi59++ {
                            for fi60 := 0; fi60 < 12; fi60++ {
                                var v61 int64 = 42
                                _ = fi60
                                _ = v61
                            }
                            v31 -= (((v31 - v2) - float64(-2.25)) + v31)
                            v31 = v31
                            _ = fi59
                        }
                        if true {
                            var v62 uint32 = 1
                            var v63 uint16 = 0
                            v64 := GenS8{F0: 18446744073709551615}
                            _ = v62
                            _ = v63
                            _ = v64
                        }
                    }
                    fmt.Println(v29)
                    fmt.Println(v4, v1, v29)
                }
            default:
                switch ("\n\t" + "a") {
                case "\n\t":
                    v5 = append(v5, 42)
                default:
                    fmt.Println(v29)
                }
            }
        }
    default:
        v2 += 1e10
    }
    fmt.Println(v31, v29, v4)
    var v65 float32 = (-2.25 / 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v65
}
