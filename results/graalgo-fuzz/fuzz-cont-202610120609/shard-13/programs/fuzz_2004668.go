package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 int32; F2 string }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 string; F1 int16; F2 uint8 }
type GenS7 struct { F0 uint16; F1 uint8 }
type GenS8 struct { F0 string; F1 uint32 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int16; F1 int16; F2 uint32 }
type GenS12 struct { F0 uint; F1 int64 }
type GenS13 struct { F0 int16; F1 uint }
type GenS14 struct { F0 int16; F1 uint64; F2 uint }
type GenS15 struct { F0 string; F1 uint8; F2 uint32 }
type GenS16 struct { F0 uint64; F1 uint }
type GenS17 struct { F0 int8; F1 bool }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 int32 }
type GenS20 struct { F0 string; F1 int64; F2 int64 }
type GenS21 struct { F0 uint8 }
type GenS22 struct { F0 uint16; F1 bool; F2 int }
type GenS23 struct { F0 uint8; F1 bool; F2 int64 }
type GenS24 struct { F0 int16; F1 int64 }
type GenS25 struct { F0 uint64; F1 uint8 }
type GenS26 struct { F0 string; F1 string; F2 int }
type GenS27 struct { F0 uint64 }
type GenS28 struct { F0 string; F1 uint16; F2 int64 }
type GenS29 struct { F0 uint16 }
type GenS30 struct { F0 bool }
type GenS31 struct { F0 bool; F1 bool }
type GenS32 struct { F0 bool }
type GenS33 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS34 struct { F0 int64; F1 string }
type GenS35 struct { F0 uint16; F1 int16; F2 int }
type GenS36 struct { F0 int32 }
type GenS37 struct { F0 uint; F1 int16 }
type GenS38 struct { F0 string; F1 int16; F2 int }
type GenS39 struct { F0 uint32; F1 int8; F2 int64 }
type GenS40 struct { F0 uint16; F1 int16; F2 int64 }
type GenS41 struct { F0 int64; F1 uint64 }
func genFn0(p0 uint64, p1 int16) uint {
    v7 := GenS1{F0: (true || false)}
    for fi8 := 0; fi8 < 15; fi8++ {
        if true {
            if false {
                p1 -= -863
            } else {
                for fi9 := 0; fi9 < 12; fi9++ {
                    fmt.Println(p0, p1)
                    p1 = -32768
                    fmt.Println(p0, p1)
                    _ = fi9
                }
            }
        }
        v10 := []int{(^(fi8)), fi8, 42}
        if (!false) {
            for fi11 := 0; fi11 < 5; fi11++ {
                p1 = (p1 ^ -810)
                if (42 < 842) {
                    if (uint16(p0) != 0) {
                        var v12 int32 = 481
                        p1 = 32767
                        var v13 uint8 = 152
                        _ = v12
                        _ = v13
                    } else {
                        var v14 int = (fi8 | 9223372036854775807)
                        _ = v14
                    }
                    p0 = 18446744073709551615
                }
                if true {
                    var v15 uint8 = 128
                    if false {
                        v10 = append(v10, 1)
                    }
                    fmt.Println(p1)
                    _ = v15
                }
                p1 = (p1 | p1)
                _ = fi11
            }
            p0 = p0
            switch fi8 {
            case -9223372036854775808:
                var v16 int8 = -77
                _ = v16
            case 1:
                p1 += (p1 | -598)
            case 2147483647:
                switch (fi8 + (fi8 >> ((255) & 63))) {
                case -907:
                    fmt.Println(p0, fi8, p0)
                case 127:
                    for fi17 := 0; fi17 < 16; fi17++ {
                        switch "hello" {
                        case "\n\t":
                            v18 := GenS1{F0: false}
                            _ = v18
                        }
                        _ = fi17
                    }
                case 2147483647:
                    for fi19 := 0; fi19 < 17; fi19++ {
                        if (!true) {
                            var v20 uint32 = 1
                            var v21 int16 = -32768
                            _ = v20
                            _ = v21
                        }
                        var v22 int64 = 1
                        if true {
                            switch "" {
                            case "hello":
                                if ((false && (2 >= 42)) && true) {
                                    v23 := GenS2{F0: 4294967295, F1: 504}
                                    _ = v23
                                } else {
                                    p1 += p1
                                }
                            }
                            v24 := GenS3{F0: 264}
                            v22 -= (v22 % ((v22) | 1))
                            _ = v24
                        } else {
                            var v25 uint16 = 996
                            _ = v25
                        }
                        _ = fi19
                        _ = v22
                    }
                }
            default:
                if len(v10) > 0 {
                    v10[0] = fi8
                }
            }
        }
        _ = fi8
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    return 570
}
func genFn1(p0 int64, p1 uint32) bool {
    v26 := GenS1{F0: (p1 > 2147483648)}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v26
    return true
}

func main() {
    v1 := []int{127, 4294967295, 128}
    var v2 int16 = -32768
    var v3 uint32 = 42
    var v4 uint8 = 2
    var v5 float64 = 3.14159
    v6 := make(map[int]string)
    for fi27 := 0; fi27 < 13; fi27++ {
        v5 += (v5 * (v5 - v5))
        var v28 int64 = -9223372036854775808
        _ = fi27
        _ = v28
    }
    v29 := GenS1{F0: (9223372036854775807 > 2147483647)}
    v30 := GenS4{F0: 44, F1: -353, F2: "a"}
    if (v3 == v3) {
        v2 = v2
        _ = len(v6)
    } else {
        v3 = 4294967295
    }
    for fi31 := 0; fi31 < 4; fi31++ {
        for fi32 := 0; fi32 < 16; fi32++ {
            if len(v1) > 0 {
                v1[0] = fi32
            }
            fmt.Println(fi32)
            if false {
                switch fi32 {
                case -1:
                    v33 := GenS4{F0: 1, F1: 0, F2: "hello"}
                    v1 = append(v1, (fi31 % ((fi31) | 1)))
                    _ = v33
                case 2147483647:
                    if (false || false) {
                        if (2 > 4294967295) {
                            fmt.Println(v4, v5, fi32)
                        } else {
                            if false {
                                if true {
                                    if ((((0.1 * (0.0 + ((-2.25 * float32((-2.25 / float64(-2.25)))) * 1e10))) < 1e10) && false) || false) {
                                        var v34 uint16 = 255
                                        switch "" {
                                        case "a":
                                            switch (-(fi32)) {
                                            case 128:
                                                var v35 int16 = -557
                                                _ = v35
                                            }
                                        default:
                                            if false {
                                                switch fi32 {
                                                case -399:
                                                    if false {
                                                        var v36 int32 = 763
                                                        _ = v36
                                                    } else {
                                                        v5 = ((v5 / ((((v5 / float64(1e10)) * v5) * float64(-2.25)) + (3.14159 + v5))) - ((-2.25 * v5) / 1.0))
                                                    }
                                                default:
                                                    if len(v1) > 0 {
                                                        v1[0] = fi31
                                                    }
                                                }
                                                var v37 uint16 = (v34 - (v34 - v34))
                                                _ = v37
                                            } else {
                                                fmt.Println(v5, v34, fi32)
                                                v38 := GenS2{F0: 2147483648, F1: 2147483647}
                                                var v39 float32 = (float32(v5) - 0.1)
                                                _ = v38
                                                _ = v39
                                            }
                                        }
                                        _ = v34
                                    } else {
                                        fmt.Println(v4, v2)
                                        if false {
                                            v5 -= (v5 + v5)
                                        } else {
                                            v5 -= (v5 * v5)
                                            switch fi32 {
                                            case -2147483648:
                                                if (-2.25 <= 0.0) {
                                                    v40 := GenS1{F0: (((((("" + "a") + "a") + "foo bar") + "\n\t") + "") <= "hello")}
                                                    _ = v40
                                                } else {
                                                    v4 = v4
                                                    v41 := v6[fi32]
                                                    _ = v41
                                                }
                                            }
                                            switch "foo bar" {
                                            case "a":
                                                var v42 float32 = -2.25
                                                var v43 int16 = (v2 | -32768)
                                                v44 := GenS5{F0: 9223372036854775807, F1: v4}
                                                if (1 != 1) {
                                                    var v45 uint32 = 332
                                                    switch "foo bar" {
                                                    case "":
                                                        v46 := GenS6{F0: "", F1: v43, F2: 223}
                                                        _ = v46
                                                    case "\n\t":
                                                        v3 -= v45
                                                    default:
                                                        if (v2 == 42) {
                                                            v43 -= -1
                                                            v47 := GenS7{F0: 1, F1: v4}
                                                            _ = v47
                                                        } else {
                                                            v4 += v4
                                                        }
                                                    }
                                                    v48 := GenS6{F0: "", F1: (v43 << ((255) & 63)), F2: (v4 ^ 184)}
                                                    v49 := GenS8{F0: "", F1: v3}
                                                    switch -223 {
                                                    case -2147483648:
                                                        switch (-(fi32)) {
                                                        case 42:
                                                            v2 = (-(v43))
                                                        case -246:
                                                            var v50 float64 = v5
                                                            _ = v50
                                                        case 113:
                                                            v51 := GenS1{F0: true}
                                                            v4 = v4
                                                            v4 = 112
                                                            _ = v51
                                                        }
                                                    case 1:
                                                        v52 := v6[fi32]
                                                        _ = v52
                                                    case 32767:
                                                        var v53 uint = 42
                                                        _ = v53
                                                    default:
                                                        v1 = append(v1, 1)
                                                    }
                                                    _ = v45
                                                    _ = v48
                                                    _ = v49
                                                } else {
                                                    v5 -= v5
                                                    v54 := GenS9{F0: ("hello" + "a")}
                                                    _ = v54
                                                }
                                                _ = len(v6)
                                                _ = v42
                                                _ = v43
                                                _ = v44
                                            case "hello":
                                                v2 += (v2 - 42)
                                            }
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v2 = -363
                                        switch (fi31 / (((fi31 / ((9223372036854775807) | 1))) | 1)) {
                                        case 1:
                                            fmt.Println(fi31)
                                        }
                                    }
                                    var v55 int8 = -1
                                    var v56 int8 = 127
                                    _ = v55
                                    _ = v56
                                } else {
                                    if false {
                                        switch "a" {
                                        case "a":
                                            switch (("a" + ("a" + ((("hello" + "a") + "a") + "a"))) + ("a" + "a")) {
                                            case "a":
                                                v3 -= (v3 ^ 65535)
                                            case "hello":
                                                v4 += v4
                                            case "":
                                                v57 := GenS10{F0: (v2 << ((65535) & 63))}
                                                _ = v57
                                            default:
                                                v5 = (float64(0.1) + (0.1 - v5))
                                            }
                                        case "":
                                            if (1 > -1) {
                                                v58 := GenS6{F0: "hello", F1: (v2 % ((v2) | 1)), F2: 80}
                                                _ = v58
                                            } else {
                                                v5 = ((((v5 - 1e10) - -2.25) * 0.0) * (1e10 + (v5 + -2.25)))
                                                v59 := GenS11{F0: v2, F1: (v2 + v2), F2: 706}
                                                switch ("foo bar" + "a") {
                                                case "a":
                                                    var v60 uint8 = v4
                                                    _ = v60
                                                case "foo bar":
                                                    switch (fi31 / ((fi32) | 1)) {
                                                    case -9223372036854775808:
                                                        v61 := GenS12{F0: 2, F1: 376}
                                                        _ = v61
                                                    case 639:
                                                        var v62 int32 = -1
                                                        _ = v62
                                                    case 9223372036854775807:
                                                        var v63 bool = true
                                                        _ = v63
                                                    }
                                                }
                                                _ = v59
                                            }
                                        case "foo bar":
                                            var v64 uint = genFn0(2, v2)
                                            _ = v64
                                        default:
                                            v65 := GenS13{F0: 593, F1: uint(v3)}
                                            _ = v65
                                        }
                                    }
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v66 := v6[fi32]
                                _ = v66
                            } else {
                                var v67 uint32 = v3
                                if false {
                                    fmt.Println(fi32)
                                    fmt.Println(fi31, fi32)
                                } else {
                                    v68 := GenS5{F0: -668, F1: 42}
                                    if false {
                                        v4 -= v4
                                        fmt.Println(v5, v3)
                                        v6[(fi32 | (fi31 ^ (fi31 | 4294967295)))] = "a"
                                    } else {
                                        v5 += v5
                                        var v69 uint = 4294967295
                                        v3 = (v67 >> ((genFn0(1, (v2 >> ((255) & 63)))) & 63))
                                        _ = v69
                                    }
                                    _ = v68
                                }
                                fmt.Println(fi32, v5)
                                if (int64(fi31) == 1) {
                                    var v70 int64 = 0
                                    v3 -= (v67 | v67)
                                    var v71 int64 = (v70 << ((4294967295) & 63))
                                    _ = v70
                                    _ = v71
                                }
                                switch (fi31 * fi32) {
                                case 867:
                                    delete(v6, fi32)
                                case -126:
                                    if ((!false) && false) {
                                        if true {
                                            var v72 float32 = (-2.25 * 3.14159)
                                            switch ("" + ("\n\t" + "\n\t")) {
                                            case "foo bar":
                                                v73 := GenS4{F0: 937, F1: 1, F2: ""}
                                                _ = v73
                                            case "":
                                                v74 := GenS14{F0: (^(v2)), F1: 2, F2: 42}
                                                _ = v74
                                            case "hello":
                                                v5 = v5
                                            }
                                            v75 := GenS15{F0: ("a" + ("hello" + (("hello" + (("foo bar" + ("a" + "hello")) + ("hello" + ("foo bar" + "a")))) + ("foo bar" + ("hello" + "foo bar"))))), F1: (v4 & (v4 | v4)), F2: v3}
                                            _ = v72
                                            _ = v75
                                        }
                                        if true {
                                            v3 = (v67 << ((uint(v67)) & 63))
                                            fmt.Println(fi31, fi31)
                                            fmt.Println(v4)
                                        }
                                        v76 := GenS16{F0: 771, F1: 65535}
                                        if true {
                                            var v77 int = fi32
                                            _ = v77
                                        } else {
                                            var v78 bool = true
                                            v6[0] = ""
                                            v5 = float64(float32(v5))
                                            _ = v78
                                        }
                                        var v79 uint16 = 1
                                        _ = v76
                                        _ = v79
                                    } else {
                                        v6[234] = "hello"
                                        v2 -= (v2 >> ((genFn0(uint64(v4), int16(fi31))) & 63))
                                    }
                                default:
                                    fmt.Println(fi31, fi31)
                                    v80 := GenS17{F0: 100, F1: true}
                                    fmt.Println(v2, fi32, v4)
                                    _ = v80
                                }
                                _ = v67
                            }
                        }
                        v4 += v4
                        var v81 int64 = -9223372036854775808
                        if true {
                            switch (fi32 * fi32) {
                            case -259:
                                switch (fi32 + -9223372036854775808) {
                                case -146:
                                    v5 = ((v5 * 1.5) + v5)
                                default:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch (fi32 - 42) {
                                    case 0:
                                        if (float32(v5) == float32(float64(1e10))) {
                                            fmt.Println(v5)
                                        }
                                    case 32768:
                                        var v82 uint64 = 951
                                        _ = v82
                                    }
                                    if true {
                                        v83 := GenS11{F0: v2, F1: v2, F2: 492}
                                        if ((v4 + (v4 / (((^(v4))) | 1))) >= (v4 | 206)) {
                                            switch fi32 {
                                            case -9223372036854775808:
                                                switch -9223372036854775808 {
                                                case 127:
                                                    v84 := GenS4{F0: 1, F1: int32(v3), F2: "foo bar"}
                                                    _ = v84
                                                case 128:
                                                    var v85 bool = true
                                                    _ = v85
                                                default:
                                                    if (9223372036854775807 <= -9223372036854775808) {
                                                        v2 -= v2
                                                    } else {
                                                        v81 = 9223372036854775807
                                                    }
                                                    if false {
                                                        var v86 float64 = (1.5 - (v5 + v5))
                                                        _ = v86
                                                    } else {
                                                        if true {
                                                            v81 = v81
                                                            var v87 uint64 = 2
                                                            v2 = v2
                                                            v3 = 0
                                                            if (464 >= (v87 >> ((18446744073709551615) & 63))) {
                                                                var v88 int = (^(fi32))
                                                                _ = len(v1)
                                                                var v89 int8 = -128
                                                                _ = v88
                                                                _ = v89
                                                            } else {
                                                                switch (fi31 % (((fi31 - (fi31 % ((9223372036854775807) | 1)))) | 1)) {
                                                                case 906:
                                                                    if (false && true) {
                                                                        var v90 uint16 = 255
                                                                        _ = v90
                                                                    }
                                                                case 242:
                                                                    if (v4 < v4) {
                                                                        if true {
                                                                            var v91 int32 = 42
                                                                            var v92 uint64 = v87
                                                                            var v93 int8 = 38
                                                                            _ = v91
                                                                            _ = v92
                                                                            _ = v93
                                                                        } else {
                                                                            v87 += v87
                                                                            if len(v1) > 0 {
                                                                                v1[0] = 4294967295
                                                                            }
                                                                            var v94 int16 = (v2 >> ((255) & 63))
                                                                            _ = v94
                                                                        }
                                                                    } else {
                                                                        v95 := v6[(^(fi32))]
                                                                        switch (fi32 * (fi32 >> ((18446744073709551615) & 63))) {
                                                                        case 128:
                                                                            v3 -= v3
                                                                        default:
                                                                            var v96 int8 = -22
                                                                            _ = v96
                                                                        }
                                                                        fmt.Println(v95)
                                                                        _ = v95
                                                                    }
                                                                case -87:
                                                                    v81 += (v81 >> ((4294967295) & 63))
                                                                }
                                                            }
                                                            _ = v87
                                                        } else {
                                                            if (1 != 255) {
                                                                v97 := GenS18{F0: false}
                                                                v3 = v3
                                                                _ = v97
                                                            } else {
                                                                v3 = v3
                                                                v2 = -535
                                                                switch fi32 {
                                                                case 128:
                                                                    v5 -= v5
                                                                case 1:
                                                                    var v98 int64 = 42
                                                                    _ = v98
                                                                case 42:
                                                                    if false {
                                                                        if true {
                                                                            fmt.Println(v2, v81, fi31)
                                                                            var v99 string = ""
                                                                            if (true && (-1 == -1)) {
                                                                                v81 = v81
                                                                                if (false && ((!false) || ((!((v81 * (v81 - v81)) < 0)) || true))) {
                                                                                    fmt.Println(v4, v81, v3)
                                                                                    v100 := GenS11{F0: v2, F1: -1, F2: 42}
                                                                                    switch (v99 + v99) {
                                                                                    case "foo bar":
                                                                                        var v101 bool = (v2 < v2)
                                                                                        v102 := GenS6{F0: v99, F1: 42, F2: 42}
                                                                                        delete(v6, (fi32 << ((18446744073709551615) & 63)))
                                                                                        var v103 uint8 = (v4 | v4)
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = 4294967295
                                                                                        }
                                                                                        if (fi31 != fi32) {
                                                                                            var v104 uint64 = 18446744073709551615
                                                                                            _ = v104
                                                                                        } else {
                                                                                            var v105 uint64 = 18446744073709551615
                                                                                            if ((^(v81)) != v81) {
                                                                                                fmt.Println(fi31, v2, v3)
                                                                                                v106 := GenS5{F0: 752, F1: (v4 ^ (^(v103)))}
                                                                                                var v107 int8 = 35
                                                                                                _ = v106
                                                                                                _ = v107
                                                                                            } else {
                                                                                                var v108 float64 = v5
                                                                                                switch (v99 + (v99 + v99)) {
                                                                                                case "a":
                                                                                                    fmt.Println(v108, v3, v103)
                                                                                                case "hello":
                                                                                                    if (!(!v101)) {
                                                                                                        fmt.Println(v3, v4, v101)
                                                                                                        var v109 uint = 330
                                                                                                        v110 := GenS2{F0: v3, F1: (^(v3))}
                                                                                                        var v111 int64 = -9223372036854775808
                                                                                                        _ = v109
                                                                                                        _ = v110
                                                                                                        _ = v111
                                                                                                    } else {
                                                                                                        var v112 int8 = -88
                                                                                                        _ = len(v1)
                                                                                                        var v113 int = int(v3)
                                                                                                        _ = v112
                                                                                                        _ = v113
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    if ((0.0 / (0.0 + (((0.0 - 3.14159) - 0.0) - (1e10 - 0.1)))) == 3.14159) {
                                                                                                        v105 += 0
                                                                                                        var v114 int16 = int16(v3)
                                                                                                        if v101 {
                                                                                                            if v101 {
                                                                                                                if len(v1) > 0 {
                                                                                                                    v1[0] = fi31
                                                                                                                }
                                                                                                                if v101 {
                                                                                                                    fmt.Println(v108, v5, v108)
                                                                                                                    var v115 int64 = v81
                                                                                                                    _ = v115
                                                                                                                } else {
                                                                                                                    var v116 uint32 = v3
                                                                                                                    v81 = v81
                                                                                                                    _ = len(v6)
                                                                                                                    _ = v116
                                                                                                                }
                                                                                                            } else {
                                                                                                                v117 := v6[(fi32 / ((fi32) | 1))]
                                                                                                                if true {
                                                                                                                    v6[(fi31 | fi31)] = v99
                                                                                                                    switch "foo bar" {
                                                                                                                    case "a":
                                                                                                                        _ = len(v1)
                                                                                                                    case "\n\t":
                                                                                                                        var v118 uint32 = uint32(v103)
                                                                                                                        v1 = append(v1, 128)
                                                                                                                        v99 = "foo bar"
                                                                                                                        var v119 string = v99
                                                                                                                        _ = v118
                                                                                                                        _ = v119
                                                                                                                    default:
                                                                                                                        v3 = v3
                                                                                                                    }
                                                                                                                    var v120 int16 = (-(v2))
                                                                                                                    _ = v120
                                                                                                                }
                                                                                                                v121 := GenS5{F0: (v81 & (v81 >> ((0) & 63))), F1: v4}
                                                                                                                v122 := GenS19{F0: 2147483647}
                                                                                                                var v123 uint32 = v3
                                                                                                                _ = v117
                                                                                                                _ = v121
                                                                                                                _ = v122
                                                                                                                _ = v123
                                                                                                            }
                                                                                                            fmt.Println(v101)
                                                                                                            var v124 uint8 = v103
                                                                                                            v125 := GenS2{F0: 4294967295, F1: 361}
                                                                                                            _ = v124
                                                                                                            _ = v125
                                                                                                        } else {
                                                                                                            v126 := GenS20{F0: "\n\t", F1: (v81 / ((9223372036854775807) | 1)), F2: 613}
                                                                                                            v127 := GenS1{F0: ((v99 > (((("hello" + "hello") + (v99 + "foo bar")) + "hello") + v99)) || ((-1 <= v2) || v101))}
                                                                                                            _ = v126
                                                                                                            _ = v127
                                                                                                        }
                                                                                                        _ = v114
                                                                                                    } else {
                                                                                                        v6[162] = "a"
                                                                                                    }
                                                                                                }
                                                                                                _ = v108
                                                                                            }
                                                                                            _ = v105
                                                                                        }
                                                                                        _ = v101
                                                                                        _ = v102
                                                                                        _ = v103
                                                                                    case "\n\t":
                                                                                        v2 = (v2 / ((v2) | 1))
                                                                                    case "a":
                                                                                        switch "\n\t" {
                                                                                        case "\n\t":
                                                                                            switch "" {
                                                                                            case "foo bar":
                                                                                                switch (v99 + v99) {
                                                                                                case "":
                                                                                                    switch -2147483648 {
                                                                                                    case 360:
                                                                                                        if false {
                                                                                                            var v128 uint16 = 0
                                                                                                            var v129 uint8 = uint8(v3)
                                                                                                            v1 = append(v1, fi32)
                                                                                                            _ = v128
                                                                                                            _ = v129
                                                                                                        } else {
                                                                                                            if true {
                                                                                                                switch 32767 {
                                                                                                                case 2147483647:
                                                                                                                    v81 = int64(v4)
                                                                                                                case 32768:
                                                                                                                    v130 := GenS21{F0: (v4 % ((v4) | 1))}
                                                                                                                    _ = v130
                                                                                                                default:
                                                                                                                    fmt.Println(v3, v3)
                                                                                                                }
                                                                                                                v3 += v3
                                                                                                                var v131 uint8 = 1
                                                                                                                var v132 int16 = v2
                                                                                                                var v133 uint = 2
                                                                                                                _ = v131
                                                                                                                _ = v132
                                                                                                                _ = v133
                                                                                                            } else {
                                                                                                                v81 = (v81 | v81)
                                                                                                                fmt.Println(v2, fi32, v4)
                                                                                                                var v134 uint32 = (v3 + (v3 * v3))
                                                                                                                _ = v134
                                                                                                            }
                                                                                                            v4 -= v4
                                                                                                        }
                                                                                                    case 128:
                                                                                                        if len(v1) > 0 {
                                                                                                            v1[0] = (fi31 % ((int(v3)) | 1))
                                                                                                        }
                                                                                                    case -1:
                                                                                                        v135 := GenS22{F0: 2, F1: false, F2: (fi31 >> ((1) & 63))}
                                                                                                        _ = v135
                                                                                                    }
                                                                                                case "a":
                                                                                                    v2 -= int16(v4)
                                                                                                default:
                                                                                                    var v136 int = 2147483647
                                                                                                    _ = v136
                                                                                                }
                                                                                            case "\n\t":
                                                                                                switch v99 {
                                                                                                case "\n\t":
                                                                                                    _ = len(v1)
                                                                                                case "":
                                                                                                    var v137 uint = 754
                                                                                                    _ = v137
                                                                                                }
                                                                                            case "hello":
                                                                                                v138 := GenS12{F0: 836, F1: -9223372036854775808}
                                                                                                var v139 uint8 = 255
                                                                                                delete(v6, fi32)
                                                                                                _ = v138
                                                                                                _ = v139
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        if (-68 < -208) {
                                                                                            if false {
                                                                                                switch fi31 {
                                                                                                case 42:
                                                                                                    if false {
                                                                                                        var v140 int8 = -1
                                                                                                        fmt.Println(v99)
                                                                                                        fmt.Println(v4, v99)
                                                                                                        _ = v140
                                                                                                    }
                                                                                                case 0:
                                                                                                    if len(v1) > 0 {
                                                                                                        v1[0] = fi31
                                                                                                    }
                                                                                                case 888:
                                                                                                    if ((^(v4)) == 0) {
                                                                                                        v141 := GenS23{F0: v4, F1: (42 != 100), F2: (v81 >> ((255) & 63))}
                                                                                                        if false {
                                                                                                            if (true && (!false)) {
                                                                                                                v142 := GenS10{F0: 32767}
                                                                                                                var v143 bool = ((-(v81)) > -1)
                                                                                                                _ = v142
                                                                                                                _ = v143
                                                                                                            } else {
                                                                                                                v6[(^(fi32))] = v99
                                                                                                                fmt.Println(fi31, v3)
                                                                                                                v99 = "\n\t"
                                                                                                            }
                                                                                                            v144 := GenS6{F0: v99, F1: -32768, F2: v4}
                                                                                                            switch v99 {
                                                                                                            case "a":
                                                                                                                v145 := GenS22{F0: 219, F1: ((3.14159 + -2.25) == 1e10), F2: fi32}
                                                                                                                _ = v145
                                                                                                            }
                                                                                                            _ = v144
                                                                                                        } else {
                                                                                                            v99 = ("" + ((((("foo bar" + (v99 + (v99 + "foo bar"))) + "\n\t") + (v99 + ("hello" + v99))) + ((v99 + ("foo bar" + v99)) + (v99 + v99))) + v99))
                                                                                                            v5 -= 0.1
                                                                                                        }
                                                                                                        fmt.Println(fi32, v99, fi31)
                                                                                                        v81 += (v81 << ((179) & 63))
                                                                                                        _ = v141
                                                                                                    } else {
                                                                                                        v3 = 0
                                                                                                        v4 = (v4 - (v4 << ((1) & 63)))
                                                                                                        v146 := GenS24{F0: v2, F1: v81}
                                                                                                        _ = v146
                                                                                                    }
                                                                                                }
                                                                                                switch fi32 {
                                                                                                case 2147483647:
                                                                                                    if len(v1) > 0 {
                                                                                                        v1[0] = 32767
                                                                                                    }
                                                                                                default:
                                                                                                    if false {
                                                                                                        _ = len(v1)
                                                                                                        delete(v6, -909)
                                                                                                        if len(v1) >= 2 {
                                                                                                            v1 = v1[1:len(v1)]
                                                                                                        }
                                                                                                        var v147 string = ""
                                                                                                        var v148 int = (fi32 * (fi31 & (fi32 - (fi32 % ((fi32) | 1)))))
                                                                                                        _ = v147
                                                                                                        _ = v148
                                                                                                    } else {
                                                                                                        var v149 string = v99
                                                                                                        fmt.Println(fi31, v99)
                                                                                                        _ = v149
                                                                                                    }
                                                                                                    var v150 int16 = (v2 % (((^(v2))) | 1))
                                                                                                    _ = v150
                                                                                                }
                                                                                            } else {
                                                                                                fmt.Println(fi32, fi31, v2)
                                                                                                v151 := GenS4{F0: 1, F1: 42, F2: v99}
                                                                                                _ = v151
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    _ = v100
                                                                                }
                                                                            }
                                                                            v3 += 128
                                                                            v81 -= v81
                                                                            if false {
                                                                                v3 += 2147483648
                                                                                fmt.Println(fi31)
                                                                            } else {
                                                                                var v152 int32 = 1
                                                                                _ = v152
                                                                            }
                                                                            _ = v99
                                                                        }
                                                                        fmt.Println(v3, v3, v3)
                                                                        if true {
                                                                            v153 := GenS25{F0: 4294967295, F1: (v4 / ((v4) | 1))}
                                                                            switch ("hello" + "\n\t") {
                                                                            case "foo bar":
                                                                                v3 = (^(v3))
                                                                            }
                                                                            if (((((0 <= 1) && false) || true) && false) && false) {
                                                                                v5 -= v5
                                                                                var v154 uint32 = 65535
                                                                                _ = v154
                                                                            } else {
                                                                                if (161 < 611) {
                                                                                    v4 = (^(v4))
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = (fi31 / ((fi31) | 1))
                                                                                    }
                                                                                }
                                                                                fmt.Println(v5, v3)
                                                                                var v155 uint64 = 2
                                                                                fmt.Println(fi32, v4)
                                                                                switch (("foo bar" + "\n\t") + "foo bar") {
                                                                                case "foo bar":
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                case "hello":
                                                                                    v5 += v5
                                                                                    var v156 uint64 = (^(v155))
                                                                                    v157 := GenS26{F0: (("\n\t" + ("a" + (("\n\t" + ("" + "")) + "hello"))) + ""), F1: "\n\t", F2: fi31}
                                                                                    if false {
                                                                                        v158 := v6[42]
                                                                                        var v159 int64 = 165
                                                                                        _ = v158
                                                                                        _ = v159
                                                                                    }
                                                                                    _ = v156
                                                                                    _ = v157
                                                                                default:
                                                                                    if false {
                                                                                        if false {
                                                                                            var v160 int16 = (-(v2))
                                                                                            _ = v160
                                                                                        }
                                                                                        var v161 int64 = 549
                                                                                        v162 := v6[int(v3)]
                                                                                        _ = v161
                                                                                        _ = v162
                                                                                    } else {
                                                                                        v163 := GenS27{F0: v155}
                                                                                        var v164 float32 = 0.1
                                                                                        v165 := GenS9{F0: ("" + "\n\t")}
                                                                                        _ = v163
                                                                                        _ = v164
                                                                                        _ = v165
                                                                                    }
                                                                                }
                                                                                _ = v155
                                                                            }
                                                                            _ = v153
                                                                        } else {
                                                                            if (!(615 != 0)) {
                                                                                v81 -= v81
                                                                                v1 = append(v1, 2147483647)
                                                                                switch (fi32 - fi31) {
                                                                                case 4294967295:
                                                                                    v166 := GenS28{F0: "foo bar", F1: 855, F2: v81}
                                                                                    _ = v166
                                                                                default:
                                                                                    if false {
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = 32768
                                                                                        }
                                                                                        if false {
                                                                                            if (("hello" + "\n\t") == "a") {
                                                                                                var v167 int = fi32
                                                                                                _ = v167
                                                                                            } else {
                                                                                                fmt.Println(v2, fi32)
                                                                                            }
                                                                                            fmt.Println(v2)
                                                                                            if ("a" >= "foo bar") {
                                                                                                var v168 float32 = ((3.14159 - 1.5) / 1e10)
                                                                                                _ = v168
                                                                                            } else {
                                                                                                if (!((42 <= 0) || true)) {
                                                                                                    var v169 int8 = -37
                                                                                                    switch "" {
                                                                                                    case "":
                                                                                                        v1 = append(v1, fi32)
                                                                                                    case "a":
                                                                                                        if (true && (!(!("" < "hello")))) {
                                                                                                            v6[(fi32 | fi32)] = "hello"
                                                                                                        } else {
                                                                                                            var v170 float32 = float32(v5)
                                                                                                            _ = v170
                                                                                                        }
                                                                                                    }
                                                                                                    v171 := GenS20{F0: ("hello" + "\n\t"), F1: v81, F2: v81}
                                                                                                    _ = v169
                                                                                                    _ = v171
                                                                                                } else {
                                                                                                    v4 -= (v4 + v4)
                                                                                                    var v172 uint32 = 2147483647
                                                                                                    var v173 float32 = ((3.14159 * 1e10) - 0.1)
                                                                                                    if (!true) {
                                                                                                        var v174 float32 = v173
                                                                                                        _ = v174
                                                                                                    } else {
                                                                                                        v5 = v5
                                                                                                        var v175 uint16 = 255
                                                                                                        _ = v175
                                                                                                    }
                                                                                                    _ = v172
                                                                                                    _ = v173
                                                                                                }
                                                                                                var v176 uint32 = (v3 + v3)
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                                _ = v176
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        v3 += 163
                                                                                        if true {
                                                                                            if false {
                                                                                                v177 := GenS29{F0: 0}
                                                                                                _ = v177
                                                                                            } else {
                                                                                                v178 := GenS28{F0: ("hello" + "hello"), F1: 241, F2: (-(v81))}
                                                                                                var v179 uint16 = 461
                                                                                                var v180 int = (fi32 - 4294967295)
                                                                                                var v181 int16 = 1
                                                                                                _ = v178
                                                                                                _ = v179
                                                                                                _ = v180
                                                                                                _ = v181
                                                                                            }
                                                                                        } else {
                                                                                            v182 := GenS30{F0: true}
                                                                                            _ = v182
                                                                                        }
                                                                                        var v183 float64 = 0.0
                                                                                        _ = v183
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                v5 = v5
                                                                                var v184 int8 = 1
                                                                                _ = v184
                                                                            }
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            var v185 int = fi32
                                                                            _ = v185
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            var v186 string = ("" + (("" + "hello") + (("" + "foo bar") + "\n\t")))
                                                            if (v5 >= v5) {
                                                                v5 -= 1.5
                                                                if false {
                                                                    v2 += v2
                                                                    v187 := GenS31{F0: (0.1 == (v5 + v5)), F1: false}
                                                                    v3 = v3
                                                                    _ = v187
                                                                }
                                                                var v188 int8 = 51
                                                                _ = v188
                                                            }
                                                            _ = v186
                                                        }
                                                        v5 -= (v5 / v5)
                                                    }
                                                    _ = len(v6)
                                                }
                                                var v189 uint32 = (v3 ^ 65535)
                                                _ = v189
                                            case 32768:
                                                fmt.Println(fi31)
                                            case -970:
                                                var v190 uint64 = 4294967295
                                                _ = v190
                                            default:
                                                v2 -= v2
                                            }
                                        } else {
                                            switch (fi31 + fi31) {
                                            case 42:
                                                v191 := GenS32{F0: false}
                                                _ = v191
                                            case 1:
                                                var v192 uint = 18446744073709551615
                                                _ = v192
                                            case 4294967295:
                                                if (!true) {
                                                    v193 := GenS28{F0: "", F1: 65535, F2: 9223372036854775807}
                                                    _ = v193
                                                }
                                            }
                                            v4 = (v4 << ((2) & 63))
                                        }
                                        _ = v83
                                    }
                                }
                            case -638:
                                switch "\n\t" {
                                case "\n\t":
                                    delete(v6, (fi31 >> ((genFn0(0, 1)) & 63)))
                                case "a":
                                    if true {
                                        fmt.Println(v5, v81, v2)
                                    }
                                case "hello":
                                    v4 -= v4
                                default:
                                    var v194 uint32 = (v3 / ((42) | 1))
                                    _ = v194
                                }
                            default:
                                if true {
                                    _ = len(v1)
                                    fmt.Println(fi32, v81)
                                    v5 = 0.1
                                }
                            }
                            switch "\n\t" {
                            case "foo bar":
                                switch (("hello" + "a") + "hello") {
                                case "a":
                                    var v195 int32 = 0
                                    _ = v195
                                }
                            case "a":
                                fmt.Println(fi32, v3, v2)
                            case "hello":
                                v196 := GenS30{F0: (((((float64(-2.25) + ((-2.25 - float64(0.0)) + v5)) / (((float64((-2.25 * 0.1)) + v5) - (v5 - ((float64((((3.14159 + 0.0) * -2.25) / 1e10)) * ((float64(1e10) + -2.25) + v5)) - v5))) * 0.0)) + (v5 * 0.1)) * v5) == float64((float32(0.0) * -2.25)))}
                                _ = v196
                            default:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            v4 -= v4
                            if ("a" <= (("\n\t" + "\n\t") + "hello")) {
                                switch "" {
                                case "a":
                                    switch ("a" + "hello") {
                                    case "hello":
                                        v197 := GenS33{F0: v4, F1: uint32(v4), F2: -4}
                                        _ = v197
                                    case "a":
                                        switch "\n\t" {
                                        case "foo bar":
                                            v198 := GenS10{F0: 42}
                                            _ = v198
                                        case "":
                                            v1 = append(v1, fi31)
                                        }
                                        fmt.Println(v81, v4, v3)
                                    default:
                                        var v199 uint8 = v4
                                        if (!false) {
                                            v200 := GenS34{F0: (v81 - v81), F1: ("" + ("hello" + "hello"))}
                                            _ = v200
                                        }
                                        v201 := GenS25{F0: 330, F1: 42}
                                        _ = v199
                                        _ = v201
                                    }
                                case "":
                                    if (true && (!true)) {
                                        v4 -= (v4 << ((0) & 63))
                                        var v202 uint16 = 255
                                        fmt.Println(v3, v81)
                                        _ = v202
                                    } else {
                                        v3 = v3
                                        v203 := v6[fi32]
                                        _ = v203
                                    }
                                }
                                switch ("a" + "a") {
                                case "a":
                                    if true {
                                        v3 -= 1
                                    } else {
                                        switch (("a" + ("hello" + "\n\t")) + (((("a" + ("" + "")) + "foo bar") + ("foo bar" + (("\n\t" + "") + ("foo bar" + "\n\t")))) + "\n\t")) {
                                        case "a":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        default:
                                            v204 := GenS14{F0: (v2 - -766), F1: 643, F2: 255}
                                            _ = v204
                                        }
                                        v4 += v4
                                        _ = len(v1)
                                    }
                                case "hello":
                                    if false {
                                        fmt.Println(v4, v3)
                                    }
                                    _ = len(v1)
                                case "foo bar":
                                    if true {
                                        v205 := GenS25{F0: uint64(v3), F1: v4}
                                        if (!false) {
                                            var v206 float64 = float64((-2.25 * float32(v5)))
                                            _ = v206
                                        }
                                        _ = v205
                                    } else {
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v207 int32 = -1
                                            _ = v207
                                        case "a":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                    }
                                }
                                var v208 uint32 = (v3 / (((v3 + 2)) | 1))
                                _ = v208
                            }
                            switch 32767 {
                            case 1:
                                fmt.Println(v2, v4)
                                v209 := GenS14{F0: -1, F1: 18446744073709551615, F2: 65535}
                                var v210 uint32 = v3
                                switch ("\n\t" + "foo bar") {
                                case "":
                                    if true {
                                        var v211 int64 = (v81 >> ((246) & 63))
                                        v212 := GenS35{F0: 0, F1: v2, F2: 42}
                                        _ = v211
                                        _ = v212
                                    } else {
                                        var v213 int8 = 127
                                        var v214 int64 = v81
                                        _ = v213
                                        _ = v214
                                    }
                                    var v215 int16 = (-(v2))
                                    v210 += (v210 + (v3 & (v3 >> ((2) & 63))))
                                    _ = v215
                                default:
                                    var v216 int8 = -1
                                    if false {
                                        v217 := GenS36{F0: -627}
                                        fmt.Println(v2)
                                        _ = v217
                                    } else {
                                        if true {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v218 int8 = -95
                                            _ = v218
                                        } else {
                                            var v219 uint16 = 32767
                                            if (!(false || ((v3 & v3) == v3))) {
                                                if true {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = len(v6)
                                                    switch "hello" {
                                                    case "foo bar":
                                                        fmt.Println(v219)
                                                    case "hello":
                                                        if ((v81 ^ v81) != v81) {
                                                            if (((false || (!(v219 <= 32767))) && false) && false) {
                                                                delete(v6, (fi31 / (((fi32 * fi32)) | 1)))
                                                                v220 := GenS21{F0: (v4 / (((v4 << ((genFn0(0, v2)) & 63))) | 1))}
                                                                var v221 int16 = (v2 * v2)
                                                                _ = v220
                                                                _ = v221
                                                            }
                                                            var v222 float32 = (((1.5 + (1.5 + 3.14159)) - float32(0.0)) * 0.1)
                                                            _ = v222
                                                        } else {
                                                            var v223 float32 = float32(v5)
                                                            _ = v223
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(v2, v81)
                                                    v224 := GenS3{F0: (v2 / (((v2 + v2)) | 1))}
                                                    var v225 string = ""
                                                    _ = v224
                                                    _ = v225
                                                }
                                                v4 -= v4
                                                v2 += 1
                                            }
                                            var v226 float32 = (0.0 - 0.1)
                                            _ = v219
                                            _ = v226
                                        }
                                    }
                                    _ = v216
                                }
                                _ = v209
                                _ = v210
                            case 9223372036854775807:
                                if false {
                                    fmt.Println(v81, v4, v3)
                                    switch -220 {
                                    case 32768:
                                        switch 264 {
                                        case -9223372036854775808:
                                            if ((v81 - (v81 / ((v81) | 1))) == (v81 * (v81 - (v81 << ((4294967295) & 63))))) {
                                                if (((^(v3)) != (^(v3))) && true) {
                                                    v3 = (v3 / (((^(v3))) | 1))
                                                }
                                                if (127 != 127) {
                                                    v3 = v3
                                                    var v227 uint16 = 964
                                                    var v228 uint32 = (v3 >> ((255) & 63))
                                                    v229 := GenS8{F0: ("" + "hello"), F1: v228}
                                                    _ = v227
                                                    _ = v228
                                                    _ = v229
                                                } else {
                                                    if true {
                                                        v4 += (v4 & (v4 / ((v4) | 1)))
                                                    } else {
                                                        var v230 bool = (!((v2 << ((4294967295) & 63)) <= (-(v2))))
                                                        switch ("\n\t" + "\n\t") {
                                                        case "\n\t":
                                                            var v231 uint64 = 809
                                                            _ = v231
                                                        default:
                                                            if (42 == 154) {
                                                                v6[fi31] = ("\n\t" + ("\n\t" + ""))
                                                                if (!(!v230)) {
                                                                    v2 -= int16(fi32)
                                                                }
                                                                switch fi31 {
                                                                case 127:
                                                                    var v232 uint8 = (v4 - 204)
                                                                    _ = v232
                                                                case -2147483648:
                                                                    switch ("" + "a") {
                                                                    case "a":
                                                                        v230 = v230
                                                                    case "":
                                                                        delete(v6, 4294967295)
                                                                    default:
                                                                        v2 = (v2 << ((2) & 63))
                                                                    }
                                                                case 639:
                                                                    v3 = uint32(v4)
                                                                default:
                                                                    fmt.Println(v4)
                                                                    v233 := GenS37{F0: 285, F1: (v2 >> ((434) & 63))}
                                                                    _ = v233
                                                                }
                                                            }
                                                        }
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        switch fi31 {
                                                        case -881:
                                                            var v234 int8 = 61
                                                            _ = v234
                                                        default:
                                                            if (v2 > 42) {
                                                                var v235 uint64 = 1
                                                                v236 := GenS38{F0: "hello", F1: 42, F2: (fi32 ^ (fi31 - fi31))}
                                                                _ = v235
                                                                _ = v236
                                                            } else {
                                                                v237 := GenS39{F0: (^(v3)), F1: -52, F2: v81}
                                                                v2 += v2
                                                                _ = len(v6)
                                                                _ = v237
                                                            }
                                                        }
                                                        _ = v230
                                                    }
                                                    v238 := GenS8{F0: "\n\t", F1: v3}
                                                    v6[32768] = "\n\t"
                                                    v239 := GenS40{F0: 42, F1: (v2 << ((42) & 63)), F2: -9223372036854775808}
                                                    _ = len(v6)
                                                    _ = v238
                                                    _ = v239
                                                }
                                            }
                                        case 128:
                                            var v240 uint16 = 919
                                            _ = v240
                                        default:
                                            v241 := GenS32{F0: false}
                                            _ = v241
                                        }
                                    }
                                    v242 := v6[int(v81)]
                                    _ = v242
                                }
                            }
                        } else {
                            v6[42] = "hello"
                        }
                        _ = v81
                    }
                case 9223372036854775807:
                    var v243 bool = (v4 > (v4 << ((1) & 63)))
                    _ = v243
                }
                fmt.Println(fi31, v5, v5)
                switch (fi32 >> ((genFn0(2, 454)) & 63)) {
                case 9223372036854775807:
                    v3 -= v3
                case 719:
                    if (!true) {
                        v2 += (v2 >> ((42) & 63))
                        var v244 uint8 = v4
                        _ = v244
                    } else {
                        v245 := GenS17{F0: -62, F1: (false || (!(true && (2147483647 < 0))))}
                        switch 1 {
                        case 4294967295:
                            v246 := GenS28{F0: ("hello" + "foo bar"), F1: 255, F2: 0}
                            _ = v246
                        case -1:
                            _ = len(v6)
                        case -925:
                            v5 = (float64(0.0) - v5)
                        }
                        if false {
                            if (true && false) {
                                v247 := v6[(fi31 >> ((18446744073709551615) & 63))]
                                _ = v247
                            }
                            switch ("foo bar" + "\n\t") {
                            case "foo bar":
                                v4 = 1
                            }
                            fmt.Println(fi32, v2)
                            if (!(42 != 56)) {
                                var v248 string = "hello"
                                _ = v248
                            }
                            if (!(255 >= (v4 - 2))) {
                                var v249 uint = 453
                                _ = v249
                            } else {
                                v1 = append(v1, 1)
                                fmt.Println(v2, v3)
                                switch 9223372036854775807 {
                                case -1:
                                    fmt.Println(fi31, v4, v2)
                                case 42:
                                    var v250 int = fi32
                                    _ = v250
                                }
                            }
                        } else {
                            v2 = v2
                            var v251 uint16 = 363
                            v4 -= (v4 + 45)
                            _ = v251
                        }
                        _ = v245
                    }
                }
            } else {
                if false {
                    _ = len(v6)
                    switch fi31 {
                    case 4294967295:
                        v252 := GenS9{F0: ("foo bar" + "a")}
                        _ = v252
                    case -2147483648:
                        v3 = (v3 % (((v3 ^ (v3 % ((uint32(v4)) | 1)))) | 1))
                    case 1:
                        var v253 uint64 = 0
                        _ = v253
                    }
                } else {
                    v6[fi31] = "\n\t"
                    if (0 != -9223372036854775808) {
                        var v254 int64 = -1
                        v254 = -9223372036854775808
                        _ = v254
                    }
                    switch (("" + ("hello" + "foo bar")) + "") {
                    case "a":
                        switch fi31 {
                        case -9223372036854775808:
                            var v255 uint32 = (v3 & v3)
                            _ = v255
                        default:
                            var v256 bool = false
                            _ = v256
                        }
                    }
                }
                var v257 int16 = v2
                v258 := GenS8{F0: ("" + ("" + (("a" + "\n\t") + "foo bar"))), F1: (v3 >> ((51) & 63))}
                _ = v257
                _ = v258
            }
            _ = fi32
        }
        delete(v6, fi31)
        _ = fi31
    }
    var v259 uint64 = uint64(v4)
    v259 = 28
    v2 -= v2
    _ = len(v1)
    if (v4 != (^(v4))) {
        for fi260 := 0; fi260 < 11; fi260++ {
            v261 := GenS33{F0: (v4 / ((v4) | 1)), F1: 4294967295, F2: 42}
            _ = fi260
            _ = v261
        }
        if true {
            fmt.Println(v5, v259)
            var v262 int32 = 1
            var v263 int64 = 1
            _ = v262
            _ = v263
        } else {
            v1 = append(v1, 787)
        }
    } else {
        v264 := GenS13{F0: -1, F1: 737}
        var v265 uint16 = 0
        _ = v264
        _ = v265
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        v266 := GenS41{F0: 0, F1: 112}
        _ = v266
    } else {
        for fi267 := 0; fi267 < 5; fi267++ {
            if len(v1) > 0 {
                v1[0] = (fi267 | (fi267 | fi267))
            }
            _ = fi267
        }
        var v268 float64 = float64(float32(1.5))
        _ = v268
    }
    var v269 string = (("" + ("a" + "hello")) + (("foo bar" + ("hello" + "\n\t")) + ""))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v259:", v259)
    fmt.Println("v269:", v269)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v259
    _ = v269
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
