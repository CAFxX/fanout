package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int16; F1 int16; F2 uint16 }
type GenS4 struct { F0 uint; F1 uint64; F2 uint }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 string; F1 string; F2 int64 }
type GenS9 struct { F0 uint; F1 uint16 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint8; F1 int16 }
type GenS12 struct { F0 uint64; F1 int; F2 int16 }
type GenS13 struct { F0 uint32; F1 int32 }
type GenS14 struct { F0 int64; F1 int8; F2 uint16 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 int16 }
func genFn0(p0 uint, p1 int) int16 {
    p0 = p0
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            fmt.Println(fi8, fi7)
            fmt.Println(fi8)
            _ = fi8
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    return -1
}
func genFn1() int64 {
    v9 := GenS1{F0: "", F1: -121}
    var v10 uint64 = 2
    v11 := []int{128, 4294967295, -721, 4294967295}
    _ = v10
    _ = v11
    _ = v9
    return -606
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]uint32)
    var v3 bool = (0.1 != (0.0 - 3.14159))
    v4 := []int{-1, -439}
    v5 := make(map[uint32]bool)
    var v6 int64 = -1
    var v12 int = 0
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 uint16 = 247
        v15 := GenS2{F0: v14, F1: 465}
        _ = fi13
        _ = v14
        _ = v15
    }
    if false {
        fmt.Println(v6)
        var v16 int = 32768
        _ = v16
    } else {
        _ = len(v1)
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 string = "\n\t"
            var v19 int8 = -128
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        if v3 {
            var v21 int = v12
            if v3 {
                switch "hello" {
                case "a":
                    for fi22 := 0; fi22 < 15; fi22++ {
                        var v23 string = "hello"
                        _ = fi22
                        _ = v23
                    }
                case "\n\t":
                    for fi24 := 0; fi24 < 18; fi24++ {
                        v25 := GenS2{F0: 255, F1: 4294967295}
                        fmt.Println(v6, v12, v21)
                        var v26 int64 = (v6 | genFn1())
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                case "hello":
                    var v27 int32 = 1
                    _ = v27
                default:
                    v28 := GenS2{F0: 249, F1: 65535}
                    _ = v28
                }
                for fi29 := 0; fi29 < 1; fi29++ {
                    if true {
                        v30 := GenS1{F0: "hello", F1: 32767}
                        _ = v30
                    }
                    _ = len(v5)
                    _ = fi29
                }
                if v3 {
                    v31 := GenS3{F0: 42, F1: 883, F2: 65535}
                    v32 := GenS4{F0: 1, F1: 628, F2: 2}
                    var v33 int16 = genFn0(65535, (v21 % ((fi20) | 1)))
                    _ = v31
                    _ = v32
                    _ = v33
                }
            } else {
                if (!(v12 < (^(v21)))) {
                    for fi34 := 0; fi34 < 16; fi34++ {
                        v35 := GenS3{F0: 42, F1: 1, F2: 42}
                        if (v3 || false) {
                            if ((-2147483648 < 396) || (false && v3)) {
                                v21 += fi20
                                v36 := GenS5{F0: 183, F1: 245}
                                _ = v36
                            } else {
                                v37 := GenS2{F0: 42, F1: 18446744073709551615}
                                v38 := GenS6{F0: 0}
                                fmt.Println(v3, v3, v12)
                                var v39 int64 = (-(v6))
                                _ = v37
                                _ = v38
                                _ = v39
                            }
                            v3 = false
                            if (983 != -377) {
                                v40 := GenS7{F0: (("a" + (("foo bar" + "\n\t") + "\n\t")) + "foo bar"), F1: 204}
                                _ = v40
                            } else {
                                switch "foo bar" {
                                case "":
                                    switch "" {
                                    case "hello":
                                        var v41 int = 0
                                        _ = v41
                                    case "foo bar":
                                        if v3 {
                                            var v42 int8 = -1
                                            _ = v42
                                        }
                                        v43 := GenS8{F0: ("\n\t" + "a"), F1: "hello", F2: v6}
                                        _ = v43
                                    default:
                                        var v44 float64 = 1e10
                                        _ = v44
                                    }
                                case "a":
                                    v2[42] = 2147483648
                                }
                                if (-2.25 < 0.1) {
                                    v1[""] = ("" + "\n\t")
                                    switch "foo bar" {
                                    case "hello":
                                        fmt.Println(fi34, v21)
                                    default:
                                        v45 := GenS4{F0: 255, F1: 270, F2: 923}
                                        _ = v45
                                    }
                                }
                                v3 = false
                            }
                            fmt.Println(v3, v21, v12)
                        } else {
                            if ((4294967295 > 745) && v3) {
                                v6 = 1
                            }
                            var v46 uint32 = 2147483648
                            var v47 uint64 = 2
                            _ = v46
                            _ = v47
                        }
                        if v3 {
                            v48 := v5[4294967295]
                            v12 = (-(fi34))
                            _ = v48
                        }
                        _ = fi34
                        _ = v35
                    }
                } else {
                    fmt.Println(v21, v3)
                    var v49 uint32 = 1
                    v6 = (-(v6))
                    _ = v49
                }
            }
            if len(v4) > 0 {
                v4[0] = (fi20 + 242)
            }
            _ = v21
        }
        for fi50 := 0; fi50 < 19; fi50++ {
            v12 -= (^(fi20))
            _ = fi50
        }
        _ = fi20
    }
    v3 = false
    for fi51 := 0; fi51 < 5; fi51++ {
        fmt.Println(v3, v12)
        _ = fi51
    }
    fmt.Println(v3)
    if (!(-2.25 > 3.14159)) {
        v12 = (v12 ^ v12)
    }
    var v52 uint = 512
    v6 -= v6
    switch v12 {
    case 127:
        if true {
            var v53 float32 = (((0.1 / (3.14159 / float32((0.1 * 1e10)))) + 0.0) - float32((1.5 / (float64(3.14159) / -2.25))))
            v54 := GenS9{F0: v52, F1: uint16(v52)}
            if v3 {
                v12 = 9223372036854775807
            } else {
                var v55 uint16 = 510
                v52 = v52
                _ = v55
            }
            _ = v53
            _ = v54
        }
    case 2147483647:
        if (genFn0(4294967295, 604) >= 42) {
            v12 = 42
            if (136 <= 4294967295) {
                if len(v4) > 0 {
                    v4[0] = v12
                }
                v56 := GenS1{F0: ("foo bar" + "\n\t"), F1: 127}
                _ = v56
            }
        }
    }
    for fi57 := 0; fi57 < 14; fi57++ {
        var v58 float64 = ((3.14159 + ((0.0 + -2.25) - 3.14159)) + 0.1)
        var v59 int64 = v6
        if v3 {
            var v60 uint8 = 1
            v52 -= 55
            if false {
                for fi61 := 0; fi61 < 16; fi61++ {
                    if false {
                        if (42 == 42) {
                            fmt.Println(v12)
                            switch ("foo bar" + (("\n\t" + ((("\n\t" + ("" + "hello")) + "a") + "foo bar")) + "")) {
                            case "hello":
                                if (!v3) {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v62 := GenS7{F0: ("foo bar" + "foo bar"), F1: 2147483647}
                                    if true {
                                        var v63 int8 = -1
                                        _ = v63
                                    }
                                    _ = v62
                                } else {
                                    if (true || ((v52 < (v52 + (v52 + 2))) && ((^(v60)) < (v60 % ((0) | 1))))) {
                                        if true {
                                            if v3 {
                                                v60 = (^(v60))
                                            }
                                            var v64 int8 = -74
                                            _ = v64
                                        } else {
                                            v65 := GenS10{F0: -32768}
                                            var v66 int8 = 127
                                            switch ("hello" + ("" + "a")) {
                                            case "a":
                                                v67 := GenS10{F0: 0}
                                                _ = v67
                                            case "\n\t":
                                                v12 = fi61
                                            }
                                            _ = v65
                                            _ = v66
                                        }
                                        if v3 {
                                            v60 = v60
                                            var v68 int16 = 42
                                            if len(v4) > 0 {
                                                v4[0] = (fi57 & 4294967295)
                                            }
                                            switch (-(fi57)) {
                                            case 128:
                                                switch "a" {
                                                case "foo bar":
                                                    v69 := GenS8{F0: ((("" + "\n\t") + ("foo bar" + ("\n\t" + "a"))) + (("a" + ("" + ("foo bar" + "foo bar"))) + (("foo bar" + "hello") + ""))), F1: "a", F2: v59}
                                                    _ = v69
                                                case "a":
                                                    v60 += (v60 / ((v60) | 1))
                                                case "":
                                                    var v70 int = fi61
                                                    _ = v70
                                                }
                                            case 259:
                                                v52 -= 18446744073709551615
                                            case -991:
                                                var v71 bool = false
                                                _ = v71
                                            default:
                                                fmt.Println(v59, v3)
                                            }
                                            _ = v68
                                        }
                                        var v72 float64 = (v58 - v58)
                                        _ = v72
                                    } else {
                                        var v73 bool = false
                                        v6 = v6
                                        _ = v73
                                    }
                                }
                            }
                            v52 -= v52
                        } else {
                            if true {
                                v52 -= (v52 << ((v52) & 63))
                                switch "\n\t" {
                                case "foo bar":
                                    v74 := GenS4{F0: (v52 * 255), F1: 847, F2: v52}
                                    _ = v74
                                }
                            }
                            var v75 bool = (!(!false))
                            var v76 int8 = 1
                            _ = v75
                            _ = v76
                        }
                    } else {
                        v6 = genFn1()
                        fmt.Println(fi61, v60)
                    }
                    var v77 float64 = -2.25
                    switch ("foo bar" + "\n\t") {
                    case "a":
                        v78 := GenS9{F0: (v52 << ((v52) & 63)), F1: 2}
                        _ = v78
                    case "\n\t":
                        if (!true) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if false {
                                v3 = false
                                var v79 uint = 42
                                if false {
                                    v60 = v60
                                } else {
                                    v80 := v2[1]
                                    _ = v80
                                }
                                _ = v79
                            }
                        }
                        var v81 uint32 = 42
                        switch "" {
                        case "hello":
                            if v3 {
                                var v82 float64 = v77
                                var v83 bool = (-1 == v59)
                                var v84 uint = 18446744073709551615
                                v12 += (v12 + -2147483648)
                                switch ("" + "") {
                                case "":
                                    v59 -= int64(v60)
                                    v59 += v59
                                }
                                _ = v82
                                _ = v83
                                _ = v84
                            }
                        case "foo bar":
                            v85 := GenS11{F0: uint8(v52), F1: 1000}
                            _ = v85
                        }
                        if (false && v3) {
                            _ = len(v1)
                            switch -1 {
                            case 772:
                                switch ("\n\t" + "a") {
                                case "hello":
                                    var v86 bool = false
                                    if true {
                                        _ = len(v1)
                                    }
                                    _ = v86
                                }
                            case -1:
                                if true {
                                    v12 += fi61
                                    v87 := GenS12{F0: 2, F1: 127, F2: genFn0((^(v52)), (v12 | fi57))}
                                    v88 := GenS9{F0: 448, F1: 32767}
                                    _ = v87
                                    _ = v88
                                } else {
                                    if ((!(!v3)) || ((18446744073709551615 >= uint64(v52)) || ((-193 < -1) && (!v3)))) {
                                        switch "foo bar" {
                                        case "hello":
                                            v89 := GenS13{F0: uint32(v60), F1: -1}
                                            _ = v89
                                        case "\n\t":
                                            v90 := GenS14{F0: 0, F1: -32, F2: uint16(v52)}
                                            _ = v90
                                        }
                                    }
                                    var v91 string = "\n\t"
                                    var v92 uint64 = 4294967295
                                    _ = v91
                                    _ = v92
                                }
                            case 32768:
                                var v93 uint = (v52 >> ((v52) & 63))
                                _ = v93
                            default:
                                v94 := GenS15{F0: (v60 & 127)}
                                _ = v94
                            }
                        }
                        var v95 bool = (true || ((65535 > 0) && ((^(v6)) == (^(v59)))))
                        _ = v81
                        _ = v95
                    case "hello":
                        v96 := v1[("" + (("\n\t" + ("foo bar" + "hello")) + "\n\t"))]
                        _ = v96
                    default:
                        var v97 uint64 = 0
                        _ = v97
                    }
                    _ = fi61
                    _ = v77
                }
            } else {
                var v98 uint8 = v60
                v3 = v3
                var v99 float64 = (v58 * (3.14159 - v58))
                _ = v98
                _ = v99
            }
            _ = v60
        }
        _ = fi57
        _ = v58
        _ = v59
    }
    for fi100 := 0; fi100 < 9; fi100++ {
        switch v12 {
        case 501:
            if false {
                if len(v4) > 0 {
                    v4[0] = -9223372036854775808
                }
            }
        case -2147483648:
            v52 += v52
        default:
            var v101 int8 = 42
            _ = v101
        }
        fmt.Println(v3)
        _ = fi100
    }
    fmt.Println(v6, v3, v12)
    for fi102 := 0; fi102 < 8; fi102++ {
        if v3 {
            v3 = (!(!true))
        } else {
            var v103 uint8 = 42
            _ = v103
        }
        fmt.Println(v3)
        switch "\n\t" {
        case "":
            for fi104 := 0; fi104 < 4; fi104++ {
                v3 = (-128 >= 1)
                v105 := GenS16{F0: genFn0((v52 >> ((576) & 63)), 32768)}
                _ = fi104
                _ = v105
            }
        case "\n\t":
            switch (("foo bar" + "") + (("a" + "hello") + ("a" + "\n\t"))) {
            case "foo bar":
                v52 -= (v52 * 864)
            case "a":
                v52 += (v52 / ((578) | 1))
            default:
                for fi106 := 0; fi106 < 10; fi106++ {
                    _ = len(v4)
                    fmt.Println(v3)
                    var v107 uint16 = 255
                    _ = fi106
                    _ = v107
                }
            }
        }
        _ = fi102
    }
    var v108 uint32 = 341
    for fi109 := 0; fi109 < 2; fi109++ {
        fmt.Println(v6, v6)
        if true {
            for fi110 := 0; fi110 < 20; fi110++ {
                var v111 uint = (v52 / (((v52 - (v52 | 0))) | 1))
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                fmt.Println(v12, v108)
                _ = fi110
                _ = v111
            }
            switch -18 {
            case 127:
                if (v6 >= (v6 / (((v6 | genFn1())) | 1))) {
                    var v112 int = (v12 >> (((^(v52))) & 63))
                    fmt.Println(v108, v112, fi109)
                    _ = v112
                } else {
                    v113 := GenS3{F0: 42, F1: genFn0(v52, -1), F2: 65535}
                    _ = v113
                }
            case 128:
                fmt.Println(fi109, v52, fi109)
            case 0:
                var v114 uint32 = (^(v108))
                _ = v114
            }
        }
        _ = fi109
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v108:", v108)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v108
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
