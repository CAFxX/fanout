package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int64; F1 int64; F2 bool }
type GenS4 struct { F0 bool; F1 int8; F2 int64 }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 uint; F1 int16 }
type GenS9 struct { F0 uint32; F1 int; F2 int16 }
type GenS10 struct { F0 string; F1 int16 }
func genFn0(p0 uint8) int8 {
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := []int{(fi7 / ((fi6) | 1)), fi6}
            v9 := GenS1{F0: (p0 % ((p0) | 1)), F1: -43}
            var v10 uint64 = 1
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        var v11 bool = true
        var v12 int16 = -885
        _ = fi6
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        for fi14 := 0; fi14 < 14; fi14++ {
            if true {
                if true {
                    p0 -= (p0 % (((p0 % ((74) | 1))) | 1))
                }
                v15 := []int{fi14, -9223372036854775808, (fi14 >> ((65535) & 63))}
                _ = v15
            } else {
                if true {
                    p0 = (p0 * p0)
                    var v16 uint64 = 123
                    v17 := GenS1{F0: p0, F1: int8(p0)}
                    _ = v16
                    _ = v17
                }
                p0 -= 225
            }
            v18 := []int{-9223372036854775808, fi13, (^(fi13))}
            v19 := GenS2{F0: 2147483647, F1: 18446744073709551615}
            v20 := GenS1{F0: (p0 & p0), F1: 127}
            var v21 bool = false
            _ = fi14
            _ = v18
            _ = v19
            _ = v20
            _ = v21
        }
        for fi22 := 0; fi22 < 7; fi22++ {
            if (3.14159 >= (0.1 - -2.25)) {
                var v23 float64 = float64((1e10 + float32(1e10)))
                v23 = (v23 - float64((((1e10 / (float32((((((((1e10 + v23) + (((v23 / 3.14159) / v23) - (float64((0.0 / 3.14159)) + 1e10))) * v23) - v23) + 0.0) - 1e10) / (3.14159 - 1.5))) / -2.25)) * float32(v23)) * 1.5)))
                fmt.Println(p0, fi22)
                _ = v23
            }
            p0 += 42
            if true {
                v24 := GenS1{F0: p0, F1: 42}
                if (586 == 2147483648) {
                    var v25 uint16 = 1
                    v26 := GenS3{F0: 0, F1: 0, F2: false}
                    var v27 string = "hello"
                    var v28 bool = true
                    v25 += 65535
                    var v29 bool = v28
                    var v30 float64 = -2.25
                    var v31 float32 = 3.14159
                    _ = v25
                    _ = v26
                    _ = v27
                    _ = v28
                    _ = v29
                    _ = v30
                    _ = v31
                }
                p0 += p0
                _ = v24
            }
            _ = fi22
        }
        _ = fi13
    }
    _ = p0
    return 42
}
func genFn1(p0 int, p1 bool) int16 {
    v32 := GenS4{F0: p1, F1: 42, F2: -1}
    v33 := []int{128, (-(p0))}
    var v34 int8 = 33
    if ((255 >= 16) && (!false)) {
        for fi35 := 0; fi35 < 4; fi35++ {
            _ = len(v33)
            if p1 {
                var v36 float64 = (-2.25 * (((float64(float32((0.1 / -2.25))) * (-2.25 - float64((3.14159 / 1.0)))) + 0.0) + 1e10))
                for fi37 := 0; fi37 < 1; fi37++ {
                    if p1 {
                        p1 = p1
                        var v38 int16 = -589
                        v36 = ((v36 / float64(-2.25)) / (0.1 / v36))
                        var v39 uint8 = 42
                        var v40 uint8 = v39
                        _ = v38
                        _ = v39
                        _ = v40
                    }
                    if (p1 && false) {
                        if len(v33) > 0 {
                            v33[0] = (fi35 >> ((255) & 63))
                        }
                        if len(v33) >= 2 {
                            v33 = v33[1:len(v33)]
                        }
                    }
                    _ = len(v33)
                    if false {
                        p1 = (!false)
                        var v41 float64 = v36
                        _ = v41
                    }
                    fmt.Println(v36, fi35, fi35)
                    _ = fi37
                }
                _ = v36
            }
            if p1 {
                var v42 float64 = float64((-2.25 / 1e10))
                _ = v42
            }
            _ = fi35
        }
        for fi43 := 0; fi43 < 12; fi43++ {
            var v44 float64 = (((-2.25 / 0.1) + (-2.25 + ((0.1 * 1.5) / 1.0))) + 0.0)
            _ = fi43
            _ = v44
        }
    } else {
        if ((p0 != (p0 & (p0 % ((p0) | 1)))) || true) {
            if p1 {
                switch 266 {
                case -9223372036854775808:
                    v34 += v34
                case -807:
                    v34 += v34
                }
            }
        } else {
            if len(v33) >= 2 {
                v33 = v33[1:len(v33)]
            }
        }
        v34 += (^(v34))
    }
    if (float32(1.5) <= (-2.25 + 1e10)) {
        if p1 {
            if p1 {
                for fi45 := 0; fi45 < 3; fi45++ {
                    var v46 uint32 = 4294967295
                    if len(v33) >= 2 {
                        v33 = v33[1:len(v33)]
                    }
                    _ = fi45
                    _ = v46
                }
                p0 += p0
                switch (p0 >> ((554) & 63)) {
                case -850:
                    switch (p0 << ((18446744073709551615) & 63)) {
                    case 32767:
                        fmt.Println(p1, p1, p0)
                    case -1:
                        for fi47 := 0; fi47 < 4; fi47++ {
                            for fi48 := 0; fi48 < 17; fi48++ {
                                v34 -= genFn0(185)
                                var v49 int = fi47
                                switch "hello" {
                                case "foo bar":
                                    if (0 > int64(v34)) {
                                        var v50 float32 = ((0.1 / (3.14159 / 0.1)) * 1.5)
                                        var v51 bool = p1
                                        _ = v50
                                        _ = v51
                                    } else {
                                        v52 := GenS5{F0: ("foo bar" + "a"), F1: (false && p1)}
                                        _ = v52
                                    }
                                default:
                                    v34 += genFn0(128)
                                }
                                var v53 uint64 = 1
                                switch "hello" {
                                case "hello":
                                    switch ("\n\t" + "\n\t") {
                                    case "":
                                        v34 -= -1
                                    case "\n\t":
                                        if true {
                                            v34 = genFn0(110)
                                            v34 -= (v34 + -1)
                                            if (1 > 1) {
                                                if true {
                                                    fmt.Println(fi48, v53)
                                                    v34 -= (v34 | -128)
                                                }
                                                v33 = append(v33, 127)
                                            }
                                        }
                                    default:
                                        var v54 int16 = -317
                                        _ = v54
                                    }
                                default:
                                    switch (("a" + (((("foo bar" + ("hello" + "a")) + "\n\t") + "foo bar") + (("\n\t" + "hello") + ""))) + ("foo bar" + "foo bar")) {
                                    case "a":
                                        v55 := GenS4{F0: true, F1: (v34 >> ((76) & 63)), F2: 9223372036854775807}
                                        _ = v55
                                    case "\n\t":
                                        if len(v33) >= 2 {
                                            v33 = v33[1:len(v33)]
                                        }
                                    default:
                                        p0 = (fi48 % ((-9223372036854775808) | 1))
                                        switch "foo bar" {
                                        case "hello":
                                            if (28 < 12) {
                                                v33 = append(v33, p0)
                                                switch "" {
                                                case "a":
                                                    switch "\n\t" {
                                                    case "hello":
                                                        fmt.Println(v53, v49, p0)
                                                    case "":
                                                        if p1 {
                                                            v34 -= v34
                                                            var v56 uint = 498
                                                            v56 = v56
                                                            _ = v56
                                                        } else {
                                                            p0 = (fi47 ^ (v49 % ((p0) | 1)))
                                                            var v57 uint32 = 0
                                                            _ = v57
                                                        }
                                                        if (((0 == 648) && (("" + (("foo bar" + ("foo bar" + "foo bar")) + "")) == "hello")) || p1) {
                                                            v53 = 2
                                                        } else {
                                                            var v58 int8 = (v34 << ((1) & 63))
                                                            _ = v58
                                                        }
                                                        switch -9223372036854775808 {
                                                        case -360:
                                                            var v59 float64 = -2.25
                                                            _ = v59
                                                        case 32768:
                                                            v60 := GenS6{F0: "foo bar", F1: 42}
                                                            _ = v60
                                                        }
                                                    case "a":
                                                        var v61 uint8 = 143
                                                        _ = v61
                                                    default:
                                                        var v62 uint = 2
                                                        _ = v62
                                                    }
                                                case "":
                                                    v63 := GenS3{F0: -263, F1: -9223372036854775808, F2: p1}
                                                    _ = v63
                                                default:
                                                    var v64 uint8 = 42
                                                    _ = v64
                                                }
                                            }
                                        case "a":
                                            if p1 {
                                                p1 = (!(fi48 != (p0 ^ fi48)))
                                            } else {
                                                switch fi48 {
                                                case -2147483648:
                                                    fmt.Println(fi48)
                                                }
                                            }
                                        }
                                        if len(v33) >= 2 {
                                            v33 = v33[1:len(v33)]
                                        }
                                    }
                                }
                                _ = fi48
                                _ = v49
                                _ = v53
                            }
                            var v65 int8 = 127
                            _ = fi47
                            _ = v65
                        }
                    default:
                        v66 := GenS6{F0: "hello", F1: -396}
                        _ = v66
                    }
                case 484:
                    for fi67 := 0; fi67 < 2; fi67++ {
                        p1 = p1
                        if (!p1) {
                            if (1e10 != 1.5) {
                                for fi68 := 0; fi68 < 1; fi68++ {
                                    v69 := GenS6{F0: ("hello" + "foo bar"), F1: 0}
                                    _ = fi68
                                    _ = v69
                                }
                                var v70 int8 = v34
                                _ = v70
                            } else {
                                p0 = fi67
                                for fi71 := 0; fi71 < 20; fi71++ {
                                    v72 := GenS1{F0: 1, F1: -128}
                                    _ = fi71
                                    _ = v72
                                }
                            }
                            v34 += v34
                        } else {
                            for fi73 := 0; fi73 < 2; fi73++ {
                                v74 := GenS4{F0: (-279 < 708), F1: genFn0(152), F2: 1}
                                _ = fi73
                                _ = v74
                            }
                            _ = len(v33)
                        }
                        p0 = (p0 & 0)
                        _ = fi67
                    }
                }
            } else {
                p1 = p1
                var v75 float32 = (1e10 * 1.5)
                _ = len(v33)
                _ = v75
            }
        } else {
            var v76 uint = 65535
            v77 := GenS3{F0: 9223372036854775807, F1: 29, F2: (!true)}
            switch "foo bar" {
            case "foo bar":
                if ((1 > -224) || (255 != 0)) {
                    p1 = ((1.5 + float32(((((((3.14159 + -2.25) - 0.0) + float64(float32(0.0))) - 0.1) * 0.1) - ((3.14159 * 1e10) - float64(float32(((1e10 * float64((float32(((-2.25 * (1.5 - (1e10 - (1.5 * (3.14159 + (-2.25 - 1e10)))))) - 1e10)) * -2.25))) + ((3.14159 + (0.0 + (float64(3.14159) * (1e10 + (-2.25 + 1e10))))) * -2.25)))))))) > (0.1 / ((float32((0.1 / 1e10)) + 1e10) - 0.1)))
                }
            case "\n\t":
                fmt.Println(v76)
            case "":
                v76 -= (v76 - v76)
            }
            _ = v76
            _ = v77
        }
        p0 = 127
        _ = len(v33)
    } else {
        if (!(0 <= 691)) {
            var v78 int64 = -9223372036854775808
            _ = v78
        }
    }
    _ = p0
    _ = p1
    _ = v32
    _ = v33
    _ = v34
    return 953
}

func main() {
    var v1 bool = (187 >= -2147483648)
    var v2 uint64 = 0
    var v3 int8 = -128
    v4 := []uint32{60, 730, 65535, 2147483648}
    v5 := make(map[string]bool)
    if (3.14159 != (0.0 / 0.1)) {
        for fi79 := 0; fi79 < 3; fi79++ {
            v2 += 1000
            _ = fi79
        }
    }
    switch 9223372036854775807 {
    case 0:
        var v80 int8 = 91
        _ = v80
    case 127:
        v81 := GenS6{F0: "hello", F1: 471}
        _ = v81
    case 9223372036854775807:
        v82 := GenS1{F0: 2, F1: (v3 ^ v3)}
        _ = v82
    }
    var v83 int16 = 370
    fmt.Println(v1, v1, v83)
    switch "foo bar" {
    case "\n\t":
        v2 = v2
    case "foo bar":
        v84 := GenS6{F0: ("" + ("a" + "")), F1: v83}
        _ = v84
    }
    v85 := GenS7{F0: 9223372036854775807, F1: 65535}
    switch -1 {
    case 42:
        switch -2147483648 {
        case 2147483647:
            for fi86 := 0; fi86 < 19; fi86++ {
                var v87 uint32 = 1
                v88 := GenS8{F0: 1, F1: (-(v83))}
                _ = fi86
                _ = v87
                _ = v88
            }
        case 128:
            v2 = (v2 ^ v2)
        }
    case 32767:
        if (!((v2 & v2) > 42)) {
            fmt.Println(v3, v83, v1)
            v89 := v5["a"]
            _ = v89
        }
    case 9223372036854775807:
        v3 = (-(v3))
    default:
        v3 -= genFn0(1)
    }
    if v1 {
        var v90 uint16 = 0
        if (0 == 128) {
            fmt.Println(v3, v3)
            if (v2 == (v2 + 42)) {
                _ = len(v5)
            } else {
                fmt.Println(v2, v83, v3)
                switch -9223372036854775808 {
                case 324:
                    var v91 uint8 = 112
                    _ = v91
                default:
                    v92 := GenS2{F0: -3, F1: 65535}
                    _ = v92
                }
                v4 = append(v4, 42)
            }
        } else {
            v3 = (v3 - -1)
            if ((v83 & v83) != genFn1(2147483647, (true || true))) {
                v93 := GenS9{F0: 2147483648, F1: -563, F2: 292}
                _ = v93
            }
        }
        _ = v90
    }
    switch 9223372036854775807 {
    case 518:
        for fi94 := 0; fi94 < 7; fi94++ {
            v95 := GenS10{F0: ("hello" + ""), F1: genFn1((fi94 - fi94), v1)}
            _ = fi94
            _ = v95
        }
    case 398:
        _ = len(v5)
    case -9223372036854775808:
        delete(v5, "\n\t")
    }
    fmt.Println(v83, v2)
    v2 += v2
    fmt.Println(v83, v1)
    if ("hello" == "hello") {
        fmt.Println(v83, v3)
        switch 0 {
        case 687:
            var v96 uint32 = 381
            _ = v96
        case 127:
            v2 = (v2 % ((42) | 1))
        }
    }
    var v97 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v83:", v83)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v83
    _ = v85
    _ = v97
}
