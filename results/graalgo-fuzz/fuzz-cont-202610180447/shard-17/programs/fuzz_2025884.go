package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int32; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint8; F1 string; F2 uint16 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int16; F1 bool; F2 int32 }
type GenS6 struct { F0 int64; F1 int; F2 int16 }
type GenS7 struct { F0 int64; F1 int8; F2 uint }
type GenS8 struct { F0 bool; F1 int }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint64; F1 bool }
func genFn0(p0 uint16, p1 int, p2 uint64) int {
    v7 := []int{(-(p1)), p1, (^(p1)), -2147483648}
    p2 += p2
    _ = len(v7)
    if false {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    } else {
        for fi8 := 0; fi8 < 3; fi8++ {
            if (0.1 != float32(0.1)) {
                fmt.Println(p1)
            }
            _ = fi8
        }
        if (!true) {
            var v9 int32 = 893
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -2147483648
}

func main() {
    var v1 uint64 = 4294967295
    var v2 string = ("a" + "\n\t")
    var v3 int32 = 0
    var v4 uint16 = 255
    var v5 uint8 = 2
    var v6 int = 427
    if true {
        var v10 uint32 = 431
        var v11 int = genFn0((v4 - v4), (v6 % (((^(v6))) | 1)), v1)
        _ = v10
        _ = v11
    } else {
        v12 := GenS1{F0: 42, F1: -128}
        _ = v12
    }
    v1 = 0
    var v13 float64 = 1e10
    for fi14 := 0; fi14 < 16; fi14++ {
        v3 = v3
        for fi15 := 0; fi15 < 6; fi15++ {
            fmt.Println(v1, fi15, v6)
            switch (-(fi14)) {
            case 0:
                v16 := []int{genFn0(941, v6, (v1 << ((18446744073709551615) & 63)))}
                _ = v16
            case -899:
                v3 = (v3 % (((v3 / (((v3 ^ (v3 | v3))) | 1))) | 1))
            }
            _ = fi15
        }
        if (0 >= -527) {
            if (true && (!(false && false))) {
                for fi17 := 0; fi17 < 13; fi17++ {
                    switch v2 {
                    case "":
                        v3 = v3
                    case "hello":
                        v18 := []int{128, fi17}
                        v3 -= 676
                        if len(v18) > 0 {
                            v18[0] = genFn0(2, (v6 >> ((18446744073709551615) & 63)), 0)
                        }
                        v19 := GenS2{F0: v3, F1: 0, F2: -741}
                        _ = v18
                        _ = v19
                    case "\n\t":
                        v20 := GenS2{F0: (v3 ^ (v3 + v3)), F1: 1, F2: 42}
                        _ = v20
                    default:
                        var v21 uint64 = v1
                        _ = v21
                    }
                    v6 += fi17
                    if (1 <= v4) {
                        var v22 uint32 = 163
                        _ = v22
                    }
                    _ = fi17
                }
            }
            v4 = 255
        } else {
            if false {
                v23 := []int{(v6 | fi14), (^(v6))}
                _ = v23
            }
            v24 := GenS1{F0: 0, F1: 127}
            _ = v24
        }
        _ = fi14
    }
    if (!false) {
        var v25 uint16 = v4
        var v26 uint64 = v1
        _ = v25
        _ = v26
    } else {
        v27 := GenS2{F0: int32(v4), F1: 2147483648, F2: 9223372036854775807}
        for fi28 := 0; fi28 < 3; fi28++ {
            if ((float64(1.5) < v13) && (-2.25 > 0.1)) {
                v29 := []int{fi28, -679, (v6 << ((uint(v4)) & 63))}
                for fi30 := 0; fi30 < 18; fi30++ {
                    v29 = append(v29, (fi30 * v6))
                    var v31 int8 = -95
                    switch genFn0(v4, (fi30 << ((255) & 63)), v1) {
                    case 32768:
                        var v32 uint16 = 1
                        _ = v32
                    case 9223372036854775807:
                        if len(v29) >= 2 {
                            v29 = v29[1:len(v29)]
                        }
                    case 42:
                        if ((v5 | v5) >= v5) {
                            v5 -= (^(v5))
                            v2 = ("foo bar" + "hello")
                            v29 = append(v29, int(v31))
                            if false {
                                switch "" {
                                case "a":
                                    fmt.Println(v13, v3)
                                default:
                                    if (v1 < v1) {
                                        v33 := GenS2{F0: (-(v3)), F1: 42, F2: -224}
                                        var v34 float64 = (v13 - v13)
                                        fmt.Println(v13)
                                        _ = v33
                                        _ = v34
                                    }
                                }
                            } else {
                                if ((v2 + v2) >= ("\n\t" + v2)) {
                                    v2 = v2
                                    fmt.Println(fi28, v2)
                                    v6 += fi30
                                }
                            }
                            if (1 <= 42) {
                                if (false && ((v31 + (v31 >> ((1) & 63))) >= v31)) {
                                    v35 := GenS3{F0: (v5 | v5), F1: v2, F2: (v4 | v4)}
                                    if len(v29) >= 2 {
                                        v29 = v29[1:len(v29)]
                                    }
                                    var v36 int64 = 9223372036854775807
                                    _ = v35
                                    _ = v36
                                } else {
                                    var v37 string = v2
                                    v38 := GenS2{F0: 1, F1: 634, F2: int64(v31)}
                                    var v39 float64 = v13
                                    if len(v29) >= 2 {
                                        v29 = v29[1:len(v29)]
                                    }
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                }
                                var v40 uint64 = 42
                                var v41 uint32 = 434
                                v29 = append(v29, (v6 ^ fi30))
                                _ = v40
                                _ = v41
                            }
                        } else {
                            _ = len(v29)
                            var v42 uint = 65535
                            _ = v42
                        }
                    }
                    _ = fi30
                    _ = v31
                }
                v43 := GenS4{F0: (v5 % (((v5 << ((255) & 63))) | 1)), F1: (v3 | (^(v3)))}
                _ = v29
                _ = v43
            } else {
                v44 := GenS3{F0: (v5 % ((v5) | 1)), F1: v2, F2: v4}
                _ = v44
            }
            if (!(!true)) {
                var v45 int16 = 42
                _ = v45
            }
            v46 := GenS3{F0: v5, F1: ("a" + "a"), F2: 369}
            _ = fi28
            _ = v46
        }
        if false {
            if (v13 > float64(0.0)) {
                if ((941 <= 65535) || false) {
                    if false {
                        for fi47 := 0; fi47 < 1; fi47++ {
                            for fi48 := 0; fi48 < 4; fi48++ {
                                var v49 float64 = v13
                                if ((v3 & v3) <= (^(v3))) {
                                    var v50 float32 = -2.25
                                    _ = v50
                                } else {
                                    var v51 float64 = (v49 * (-2.25 - (v49 + 0.1)))
                                    fmt.Println(v13, v49)
                                    _ = v51
                                }
                                _ = fi48
                                _ = v49
                            }
                            fmt.Println(v2)
                            _ = fi47
                        }
                    } else {
                        var v52 float32 = -2.25
                        fmt.Println(v13, v3, v3)
                        for fi53 := 0; fi53 < 13; fi53++ {
                            var v54 int16 = -1
                            _ = fi53
                            _ = v54
                        }
                        _ = v52
                    }
                } else {
                    var v55 float32 = 0.0
                    _ = v55
                }
            }
            v56 := []int{(^(v6)), (v6 ^ (v6 + (v6 - 127)))}
            switch ((v2 + (v2 + v2)) + v2) {
            case "foo bar":
                var v57 float64 = (float64(0.0) - v13)
                _ = v57
            case "hello":
                var v58 int8 = -120
                _ = v58
            case "\n\t":
                var v59 uint32 = 4294967295
                _ = v59
            }
            _ = v56
        }
        _ = v27
    }
    v60 := GenS4{F0: v5, F1: v3}
    v61 := []int{(^(v6)), v6, (v6 & genFn0(v4, 42, v1))}
    fmt.Println(v2, v4, v1)
    if true {
        for fi62 := 0; fi62 < 2; fi62++ {
            var v63 int16 = -1
            switch 42 {
            case -869:
                for fi64 := 0; fi64 < 3; fi64++ {
                    fmt.Println(v1, fi62)
                    _ = fi64
                }
            case 4294967295:
                for fi65 := 0; fi65 < 4; fi65++ {
                    if ((-2.25 / 3.14159) == v13) {
                        v66 := GenS1{F0: -128, F1: 42}
                        if false {
                            switch 483 {
                            case -149:
                                var v67 uint = 255
                                _ = v67
                            case 32767:
                                v3 = (v3 >> ((566) & 63))
                            }
                            var v68 uint16 = v4
                            v69 := GenS1{F0: -1, F1: 100}
                            switch (v2 + ("" + (v2 + (v2 + v2)))) {
                            case "hello":
                                switch ((("hello" + "foo bar") + v2) + v2) {
                                case "":
                                    if true {
                                        v13 = v13
                                        if true {
                                            if len(v61) > 0 {
                                                v61[0] = fi65
                                            }
                                            if len(v61) > 0 {
                                                v61[0] = v6
                                            }
                                        } else {
                                            fmt.Println(v4)
                                            v3 = 42
                                            v5 = 42
                                        }
                                        if (1e10 <= ((1.5 * (1.5 + (-2.25 + 0.0))) - 0.1)) {
                                            v68 += v68
                                            var v70 float32 = (0.0 + 3.14159)
                                            _ = v70
                                        } else {
                                            if false {
                                                var v71 int64 = -1
                                                _ = len(v61)
                                                if (1e10 > 0.1) {
                                                    if true {
                                                        v13 += 3.14159
                                                        var v72 uint8 = uint8(v68)
                                                        _ = v72
                                                    }
                                                    switch v6 {
                                                    case 32767:
                                                        v6 = fi62
                                                    default:
                                                        v61 = append(v61, fi62)
                                                        if ((false || true) || ("" != "hello")) {
                                                            v2 = "hello"
                                                        } else {
                                                            switch fi62 {
                                                            case 1:
                                                                var v73 uint8 = 255
                                                                _ = v73
                                                            case 32767:
                                                                v3 -= (v3 | v3)
                                                            case 9223372036854775807:
                                                                v71 += -9223372036854775808
                                                            }
                                                            var v74 int32 = (v3 ^ 796)
                                                            _ = v74
                                                        }
                                                    }
                                                }
                                                _ = v71
                                            } else {
                                                v75 := GenS5{F0: (v63 & v63), F1: true, F2: (v3 ^ v3)}
                                                v76 := GenS6{F0: -325, F1: (v6 << ((255) & 63)), F2: (-(v63))}
                                                _ = v75
                                                _ = v76
                                            }
                                            switch v2 {
                                            case "a":
                                                fmt.Println(fi65, v3, v5)
                                            }
                                        }
                                    }
                                case "\n\t":
                                    var v77 int16 = 42
                                    _ = v77
                                case "foo bar":
                                    var v78 int8 = -1
                                    _ = v78
                                }
                            case "foo bar":
                                if (42 >= 378) {
                                    var v79 uint64 = (v1 + 826)
                                    _ = v79
                                } else {
                                    fmt.Println(v13, v5, fi65)
                                    fmt.Println(v13)
                                }
                            case "":
                                v80 := GenS7{F0: -1, F1: -1, F2: 255}
                                _ = v80
                            default:
                                switch fi62 {
                                case -246:
                                    if len(v61) >= 2 {
                                        v61 = v61[1:len(v61)]
                                    }
                                    if len(v61) >= 2 {
                                        v61 = v61[1:len(v61)]
                                    }
                                    fmt.Println(v6)
                                }
                            }
                            _ = v68
                            _ = v69
                        } else {
                            fmt.Println(v3)
                            var v81 float32 = (1e10 + 1e10)
                            v82 := GenS6{F0: 9223372036854775807, F1: genFn0((^(v4)), int(v5), 8), F2: (^(v63))}
                            _ = v81
                            _ = v82
                        }
                        _ = v66
                    }
                    _ = fi65
                }
            default:
                if (!true) {
                    v83 := GenS8{F0: false, F1: (fi62 + fi62)}
                    v13 = v13
                    for fi84 := 0; fi84 < 9; fi84++ {
                        fmt.Println(v2, v1, v1)
                        _ = fi84
                    }
                    _ = v83
                } else {
                    v85 := GenS9{F0: 0}
                    switch genFn0(v4, fi62, v1) {
                    case 4294967295:
                        for fi86 := 0; fi86 < 13; fi86++ {
                            v3 = 2147483647
                            var v87 int = fi62
                            _ = fi86
                            _ = v87
                        }
                    case 32768:
                        if len(v61) > 0 {
                            v61[0] = (^(fi62))
                        }
                    default:
                        if ((v63 + 0) <= -472) {
                            if len(v61) >= 2 {
                                v61 = v61[1:len(v61)]
                            }
                            if false {
                                if len(v61) > 0 {
                                    v61[0] = (v6 - 127)
                                }
                                v88 := GenS5{F0: v63, F1: (v5 != v5), F2: (v3 & v3)}
                                _ = v88
                            } else {
                                var v89 string = (v2 + v2)
                                if len(v61) >= 2 {
                                    v61 = v61[1:len(v61)]
                                }
                                _ = v89
                            }
                            var v90 int16 = (v63 % ((892) | 1))
                            _ = v90
                        } else {
                            var v91 int64 = 9223372036854775807
                            fmt.Println(v4, v13, v4)
                            switch genFn0((v4 << ((2) & 63)), fi62, v1) {
                            case 32767:
                                if ((-2.25 * ((0.0 * 1.5) + ((1.5 / 1.5) * 0.1))) > (3.14159 - 1e10)) {
                                    if (!true) {
                                        v92 := GenS8{F0: true, F1: (v6 << ((102) & 63))}
                                        _ = v92
                                    } else {
                                        v13 = ((v13 / float64(1.5)) / 0.1)
                                        v13 = 0.1
                                    }
                                    fmt.Println(fi62)
                                    v3 -= v3
                                } else {
                                    fmt.Println(v91)
                                    v1 -= v1
                                    v93 := GenS1{F0: 1, F1: 60}
                                    _ = v93
                                }
                            }
                            _ = v91
                        }
                    }
                    if len(v61) >= 2 {
                        v61 = v61[1:len(v61)]
                    }
                    _ = v85
                }
            }
            var v94 uint64 = v1
            _ = fi62
            _ = v63
            _ = v94
        }
    }
    for fi95 := 0; fi95 < 6; fi95++ {
        if false {
            fmt.Println(v1, v6)
            fmt.Println(v2, v4, v4)
        } else {
            switch v2 {
            case "":
                v96 := GenS10{F0: v1, F1: false}
                _ = v96
            default:
                switch v2 {
                case "":
                    var v97 uint16 = v4
                    _ = v97
                case "hello":
                    if false {
                        v6 = fi95
                    } else {
                        if len(v61) >= 2 {
                            v61 = v61[1:len(v61)]
                        }
                        if len(v61) > 0 {
                            v61[0] = (v6 >> ((uint(v1)) & 63))
                        }
                    }
                case "\n\t":
                    var v98 float32 = 1.5
                    _ = v98
                }
            }
        }
        _ = fi95
    }
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v61:", len(v61), v61)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v61
}
