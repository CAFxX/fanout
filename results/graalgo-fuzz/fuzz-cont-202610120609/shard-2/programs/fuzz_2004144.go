package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint; F1 string; F2 int16 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint64; F1 uint16 }
type GenS9 struct { F0 uint; F1 uint }
type GenS10 struct { F0 uint8; F1 int; F2 uint32 }
type GenS11 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS12 struct { F0 uint32; F1 int; F2 uint32 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int64; F1 int64 }
type GenS16 struct { F0 int; F1 uint8 }
type GenS17 struct { F0 uint32; F1 uint }
type GenS18 struct { F0 uint32; F1 int16 }
type GenS19 struct { F0 int64; F1 int64 }
func genFn0(p0 int32, p1 int64, p2 uint) bool {
    var v6 int32 = p0
    var v7 int8 = -128
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (false || ((^(p0)) >= -1))
}
func genFn1(p0 int32, p1 bool) int16 {
    v8 := []int{4294967295, -2147483648, 128, -9223372036854775808}
    if (4294967295 == 1) {
        _ = len(v8)
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 uint8 = 153
            _ = fi9
            _ = v10
        }
        if ("" < ("\n\t" + "\n\t")) {
            if p1 {
                for fi11 := 0; fi11 < 15; fi11++ {
                    for fi12 := 0; fi12 < 5; fi12++ {
                        v13 := GenS1{F0: (("hello" + "hello") + ""), F1: 1}
                        var v14 uint32 = 850
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    v8 = append(v8, -1)
                    _ = fi11
                }
                var v15 uint8 = 42
                _ = v15
            }
        } else {
            if false {
                for fi16 := 0; fi16 < 15; fi16++ {
                    if (63 == 32767) {
                        for fi17 := 0; fi17 < 9; fi17++ {
                            v18 := GenS2{F0: 2, F1: "", F2: 306}
                            if true {
                                fmt.Println(fi17)
                            } else {
                                var v19 int8 = -77
                                v20 := GenS1{F0: ("foo bar" + ""), F1: 744}
                                v19 = (-(v19))
                                p0 -= (p0 * (p0 / ((493) | 1)))
                                p1 = p1
                                v21 := GenS3{F0: 42, F1: p1}
                                fmt.Println(fi16, fi16, p1)
                                var v22 int64 = 9223372036854775807
                                _ = v19
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                            var v23 uint = 65535
                            _ = fi17
                            _ = v18
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 15; fi24++ {
                            p0 = (p0 * 0)
                            fmt.Println(p1, p0, fi16)
                            var v25 int8 = 100
                            _ = fi24
                            _ = v25
                        }
                        var v26 float32 = 0.0
                        _ = v26
                    }
                    if p1 {
                        v27 := GenS2{F0: 18446744073709551615, F1: "a", F2: -1}
                        for fi28 := 0; fi28 < 8; fi28++ {
                            v29 := GenS4{F0: ("" + "a"), F1: fi28}
                            var v30 string = "a"
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                        for fi31 := 0; fi31 < 1; fi31++ {
                            if (0 >= 18446744073709551615) {
                                v32 := GenS4{F0: "", F1: fi31}
                                var v33 float32 = -2.25
                                if p1 {
                                    _ = len(v8)
                                    v33 = 0.0
                                } else {
                                    var v34 int32 = (p0 / ((p0) | 1))
                                    var v35 float64 = ((1.5 * (0.1 / float64((v33 - ((v33 * (v33 / 3.14159)) + v33))))) * (0.1 - 0.1))
                                    p1 = false
                                    _ = v34
                                    _ = v35
                                }
                                v36 := GenS2{F0: 349, F1: ("a" + ""), F2: 32767}
                                _ = v32
                                _ = v33
                                _ = v36
                            }
                            var v37 bool = p1
                            if (2147483648 <= 0) {
                                v37 = (!(!false))
                            } else {
                                v37 = p1
                                if (1000 > -1) {
                                    var v38 string = ("a" + "foo bar")
                                    if len(v8) > 0 {
                                        v8[0] = (fi16 << ((290) & 63))
                                    }
                                    _ = v38
                                }
                            }
                            _ = fi31
                            _ = v37
                        }
                        _ = v27
                    }
                    _ = fi16
                }
            }
            fmt.Println(p0)
        }
    }
    p1 = (true || (32768 <= -997))
    _ = p0
    _ = p1
    _ = v8
    return -253
}

func main() {
    var v1 float32 = (((3.14159 + ((0.1 - -2.25) / 0.1)) / ((float32(1.5) * float32(3.14159)) - 0.0)) - 0.1)
    var v2 int16 = 32767
    var v3 uint8 = 127
    var v4 uint16 = 397
    v5 := []int{32768, -2147483648}
    if (v1 > 0.1) {
        v1 += v1
        for fi39 := 0; fi39 < 11; fi39++ {
            switch ("hello" + (((("foo bar" + "hello") + "hello") + "a") + "hello")) {
            case "hello":
                var v40 int16 = (v2 % ((genFn1(-1, (!(((0.0 > 1e10) || false) && ((v3 - v3) > v3))))) | 1))
                _ = v40
            case "a":
                for fi41 := 0; fi41 < 6; fi41++ {
                    var v42 float32 = v1
                    _ = fi41
                    _ = v42
                }
            case "foo bar":
                switch fi39 {
                case -542:
                    v4 += (v4 + v4)
                }
            default:
                var v43 float64 = ((3.14159 / ((((3.14159 - ((3.14159 - 3.14159) - -2.25)) - 1.5) * 0.0) + 1.5)) * (-2.25 - 1e10))
                _ = v43
            }
            _ = fi39
        }
        var v44 uint8 = (v3 - (v3 ^ (v3 / (((^(v3))) | 1))))
        _ = v44
    }
    var v45 string = ""
    if true {
        if true {
            var v46 uint32 = 2147483648
            v45 = v45
            switch int(v46) {
            case 9223372036854775807:
                v47 := GenS5{F0: 4294967295}
                _ = v47
            case 4294967295:
                if false {
                    switch 128 {
                    case -361:
                        v4 = 32767
                    }
                    var v48 float32 = (v1 * float32((0.1 / 1.5)))
                    v5 = append(v5, -9223372036854775808)
                    _ = v48
                }
            case -1:
                v49 := GenS5{F0: -108}
                _ = v49
            }
            _ = v46
        } else {
            if (true && (-1 < 1)) {
                v50 := GenS3{F0: v3, F1: (("" + v45) == (v45 + ((((v45 + v45) + "\n\t") + v45) + "foo bar")))}
                var v51 int16 = v2
                v52 := GenS6{F0: 156, F1: 1, F2: 1}
                _ = v50
                _ = v51
                _ = v52
            } else {
                fmt.Println(v45, v2, v1)
                fmt.Println(v2)
            }
            if true {
                switch -957 {
                case 747:
                    fmt.Println(v45, v3)
                case 328:
                    for fi53 := 0; fi53 < 17; fi53++ {
                        fmt.Println(v1, fi53)
                        _ = fi53
                    }
                default:
                    switch 1 {
                    case 32767:
                        if len(v5) > 0 {
                            v5[0] = 128
                        }
                    }
                }
                var v54 uint = 4294967295
                switch v45 {
                case "a":
                    fmt.Println(v3, v54, v1)
                default:
                    v4 = 42
                }
                _ = v54
            } else {
                if (!(23 >= 1)) {
                    _ = len(v5)
                } else {
                    var v55 float64 = (float64((v1 + v1)) - 1.5)
                    _ = len(v5)
                    fmt.Println(v2, v1)
                    _ = v55
                }
                v45 = "a"
            }
        }
    }
    if len(v5) > 0 {
        v5[0] = 9223372036854775807
    }
    var v56 uint = 255
    if true {
        if (v3 != (v3 - v3)) {
            fmt.Println(v1, v1, v4)
            var v57 int16 = 1000
            if true {
                switch 32768 {
                case 32767:
                    v58 := GenS5{F0: 31}
                    _ = v58
                case 42:
                    v57 = (v2 | (v2 / ((32767) | 1)))
                case 933:
                    fmt.Println(v2, v45, v45)
                }
                v4 -= (^(v4))
            }
            _ = v57
        }
        v59 := GenS6{F0: (v3 | 112), F1: 61, F2: 100}
        var v60 int64 = 355
        _ = v59
        _ = v60
    }
    if true {
        if ((!true) && false) {
            if (int(v4) <= -1) {
                var v61 uint = 309
                fmt.Println(v2, v2)
                var v62 bool = ((2147483647 > 970) && (-96 > -1))
                _ = v61
                _ = v62
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if true {
                    for fi63 := 0; fi63 < 14; fi63++ {
                        v64 := GenS7{F0: 2}
                        v3 = 0
                        _ = fi63
                        _ = v64
                    }
                    var v65 int32 = 0
                    _ = v65
                }
            }
            var v66 int32 = -839
            if (!true) {
                for fi67 := 0; fi67 < 10; fi67++ {
                    switch -2147483648 {
                    case 0:
                        v4 = (v4 % ((uint16(v56)) | 1))
                    case 2147483647:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case -632:
                        v68 := GenS1{F0: "a", F1: v2}
                        _ = v68
                    default:
                        v56 = 1
                    }
                    v69 := GenS8{F0: uint64(v56), F1: (^(v4))}
                    _ = fi67
                    _ = v69
                }
            }
            _ = v66
        } else {
            v56 = 28
            if (100 == 42) {
                for fi70 := 0; fi70 < 19; fi70++ {
                    switch v45 {
                    case "hello":
                        var v71 uint64 = 41
                        _ = v71
                    }
                    v72 := GenS9{F0: uint(v4), F1: (v56 >> ((1) & 63))}
                    var v73 int32 = 42
                    _ = fi70
                    _ = v72
                    _ = v73
                }
                if (((v45 + v45) >= (v45 + "foo bar")) || true) {
                    if (1 <= (v56 & v56)) {
                        fmt.Println(v45)
                        if len(v5) > 0 {
                            v5[0] = -2147483648
                        }
                        for fi74 := 0; fi74 < 1; fi74++ {
                            _ = len(v5)
                            switch (v45 + "") {
                            case "foo bar":
                                v75 := GenS8{F0: 147, F1: (^(v4))}
                                _ = v75
                            case "\n\t":
                                if false {
                                    switch fi74 {
                                    case 127:
                                        for fi76 := 0; fi76 < 4; fi76++ {
                                            if (false || false) {
                                                switch fi76 {
                                                case 354:
                                                    v1 = v1
                                                case 2147483647:
                                                    v56 = 18446744073709551615
                                                case 32767:
                                                    v77 := GenS4{F0: "a", F1: (-(fi74))}
                                                    _ = v77
                                                }
                                                _ = len(v5)
                                            } else {
                                                fmt.Println(v2)
                                                _ = len(v5)
                                                if (v56 != 147) {
                                                    if false {
                                                        if true {
                                                            if true {
                                                                if false {
                                                                    fmt.Println(v45)
                                                                    v45 = v45
                                                                    v78 := GenS10{F0: (v3 & v3), F1: fi76, F2: 679}
                                                                    _ = v78
                                                                }
                                                                v5 = append(v5, fi74)
                                                                var v79 int64 = 817
                                                                _ = v79
                                                            } else {
                                                                fmt.Println(v2)
                                                            }
                                                            switch fi76 {
                                                            case 488:
                                                                v80 := GenS11{F0: 86, F1: -9223372036854775808, F2: 2147483647}
                                                                _ = v80
                                                            default:
                                                                fmt.Println(v45, v2, v1)
                                                            }
                                                            var v81 uint32 = 2147483647
                                                            if len(v5) > 0 {
                                                                v5[0] = 1
                                                            }
                                                            _ = v81
                                                        } else {
                                                            if (!true) {
                                                                if (4294967295 == 1) {
                                                                    var v82 uint32 = 2147483648
                                                                    v45 = "\n\t"
                                                                    v56 -= v56
                                                                    _ = v82
                                                                } else {
                                                                    var v83 uint = v56
                                                                    _ = v83
                                                                }
                                                                var v84 int32 = 2147483647
                                                                _ = len(v5)
                                                                _ = v84
                                                            } else {
                                                                if true {
                                                                    v45 = (v45 + v45)
                                                                    v56 += (v56 & v56)
                                                                    var v85 uint16 = (v4 & v4)
                                                                    v1 += (v1 - ((float32(3.14159) * (1.5 / v1)) + v1))
                                                                    if (!(!(255 < (v4 | (v85 >> ((v56) & 63)))))) {
                                                                        v4 += 87
                                                                        v1 = (0.0 - v1)
                                                                    }
                                                                    _ = v85
                                                                } else {
                                                                    var v86 float64 = 0.0
                                                                    if false {
                                                                        if true {
                                                                            fmt.Println(fi76)
                                                                            fmt.Println(v56)
                                                                            switch (fi74 + fi76) {
                                                                            case 0:
                                                                                if true {
                                                                                    v3 -= v3
                                                                                    var v87 uint16 = (^(v4))
                                                                                    var v88 float32 = (((0.0 + 0.0) - v1) + v1)
                                                                                    v1 += (v1 * ((v1 + v88) + v1))
                                                                                    var v89 float32 = (float32((1.5 * v86)) - float32((((3.14159 / v86) * 1.5) - v86)))
                                                                                    _ = v87
                                                                                    _ = v88
                                                                                    _ = v89
                                                                                } else {
                                                                                    v90 := GenS11{F0: 1, F1: 0, F2: 42}
                                                                                    v91 := GenS12{F0: uint32(v56), F1: (fi74 / (((fi74 | (fi74 + fi76))) | 1)), F2: 2147483647}
                                                                                    _ = v90
                                                                                    _ = v91
                                                                                }
                                                                            case 42:
                                                                                var v92 int16 = genFn1(-475, (float64(float32(v86)) > v86))
                                                                                _ = v92
                                                                            case -2147483648:
                                                                                var v93 float32 = 0.0
                                                                                _ = v93
                                                                            }
                                                                            if (((false && false) || true) || (2147483647 < 433)) {
                                                                                v2 = v2
                                                                                if (v86 < v86) {
                                                                                    v94 := GenS9{F0: v56, F1: v56}
                                                                                    _ = v94
                                                                                }
                                                                                if false {
                                                                                    var v95 float64 = (v86 + 1e10)
                                                                                    v4 += (^(v4))
                                                                                    _ = v95
                                                                                } else {
                                                                                    v4 = (v4 >> ((v56) & 63))
                                                                                    if (127 == -1) {
                                                                                        v96 := GenS2{F0: (v56 * v56), F1: v45, F2: (v2 * int16(v56))}
                                                                                        var v97 bool = (!true)
                                                                                        fmt.Println(v86, v86)
                                                                                        _ = v96
                                                                                        _ = v97
                                                                                    }
                                                                                    v4 += 2
                                                                                }
                                                                            } else {
                                                                                if (v1 < 3.14159) {
                                                                                    v3 = (v3 >> (((^(v56))) & 63))
                                                                                    v98 := GenS13{F0: -2147483648}
                                                                                    if ((^(v56)) > (v56 * v56)) {
                                                                                        var v99 uint8 = v3
                                                                                        var v100 int64 = -9223372036854775808
                                                                                        _ = v100
                                                                                        _ = v99
                                                                                    }
                                                                                    _ = v98
                                                                                }
                                                                            }
                                                                            if ((fi74 / ((int(v56)) | 1)) <= -2147483648) {
                                                                                if false {
                                                                                    var v101 string = v45
                                                                                    fmt.Println(v86, v45, v45)
                                                                                    _ = v101
                                                                                } else {
                                                                                    fmt.Println(v4, v45, fi76)
                                                                                    v5 = append(v5, fi74)
                                                                                    if (!false) {
                                                                                        v56 -= v56
                                                                                        v2 = v2
                                                                                    }
                                                                                }
                                                                                fmt.Println(v3, v45, v86)
                                                                                v102 := GenS6{F0: (v3 ^ (^(v3))), F1: 2, F2: 1}
                                                                                fmt.Println(fi74)
                                                                                _ = v102
                                                                            }
                                                                        }
                                                                    }
                                                                    v86 -= ((-2.25 - (1e10 / (float64(float32(v86)) / v86))) - v86)
                                                                    _ = v86
                                                                }
                                                                var v103 int16 = genFn1(-521, (!(true || (false || false))))
                                                                if (v4 >= v4) {
                                                                    if true {
                                                                        switch v45 {
                                                                        case "\n\t":
                                                                            v3 -= v3
                                                                        case "foo bar":
                                                                            fmt.Println(v56, v2)
                                                                        case "a":
                                                                            fmt.Println(fi74, v103, v4)
                                                                        }
                                                                        v103 += v103
                                                                        switch -9223372036854775808 {
                                                                        case 4294967295:
                                                                            var v104 uint8 = (v3 << (((v56 * v56)) & 63))
                                                                            _ = v104
                                                                        case 32768:
                                                                            v105 := GenS14{F0: v4}
                                                                            _ = v105
                                                                        case 122:
                                                                            var v106 float64 = 0.0
                                                                            _ = v106
                                                                        default:
                                                                            v1 -= v1
                                                                        }
                                                                    } else {
                                                                        v107 := GenS12{F0: 4294967295, F1: fi76, F2: 65535}
                                                                        if (1 <= 2147483648) {
                                                                            fmt.Println(v2)
                                                                        }
                                                                        _ = v107
                                                                    }
                                                                    switch (fi74 & int(v3)) {
                                                                    case -922:
                                                                        if false {
                                                                            var v108 uint64 = 42
                                                                            _ = v108
                                                                        } else {
                                                                            if (!true) {
                                                                                if (!(!false)) {
                                                                                    v5 = append(v5, (^(fi74)))
                                                                                }
                                                                                var v109 uint32 = 0
                                                                                fmt.Println(v109)
                                                                                _ = v109
                                                                            }
                                                                        }
                                                                    default:
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                    }
                                                                }
                                                                _ = v103
                                                            }
                                                            if (0.0 <= 0.0) {
                                                                var v110 uint = v56
                                                                v45 = (("hello" + (v45 + "a")) + "\n\t")
                                                                _ = v110
                                                            }
                                                            v111 := GenS6{F0: (v3 - v3), F1: 702, F2: 1}
                                                            _ = v111
                                                        }
                                                        var v112 uint32 = 629
                                                        if false {
                                                            switch v45 {
                                                            case "foo bar":
                                                                v113 := GenS13{F0: (^(fi76))}
                                                                switch 4294967295 {
                                                                case 0:
                                                                    if true {
                                                                        if len(v5) > 0 {
                                                                            v5[0] = (-(fi76))
                                                                        }
                                                                    }
                                                                }
                                                                _ = v113
                                                            case "a":
                                                                v114 := GenS6{F0: 128, F1: 76, F2: 0}
                                                                _ = v114
                                                            case "\n\t":
                                                                var v115 uint8 = v3
                                                                _ = v115
                                                            default:
                                                                v116 := GenS11{F0: 0, F1: 0, F2: uint32(v3)}
                                                                _ = v116
                                                            }
                                                            switch 2147483647 {
                                                            case 4294967295:
                                                                v117 := GenS15{F0: -9223372036854775808, F1: 42}
                                                                _ = v117
                                                            case 456:
                                                                v1 = v1
                                                            case 1:
                                                                if len(v5) > 0 {
                                                                    v5[0] = (^(fi74))
                                                                }
                                                            }
                                                        }
                                                        _ = v112
                                                    } else {
                                                        v1 = v1
                                                        switch fi76 {
                                                        case 4294967295:
                                                            v118 := GenS16{F0: fi74, F1: uint8(v56)}
                                                            _ = v118
                                                        case 42:
                                                            v119 := GenS17{F0: 65535, F1: v56}
                                                            _ = v119
                                                        case 1:
                                                            v3 = v3
                                                            var v120 float64 = -2.25
                                                            if (v45 != ((v45 + v45) + v45)) {
                                                                v121 := GenS10{F0: v3, F1: 433, F2: 2147483648}
                                                                fmt.Println(v56, v56, v120)
                                                                v122 := GenS16{F0: (fi74 ^ (fi74 ^ fi76)), F1: v3}
                                                                _ = v121
                                                                _ = v122
                                                            }
                                                            _ = v120
                                                        }
                                                        var v123 float32 = (3.14159 * 0.1)
                                                        _ = v123
                                                    }
                                                    var v124 string = ""
                                                    v125 := GenS9{F0: (v56 + 1), F1: (v56 >> ((v56) & 63))}
                                                    _ = v124
                                                    _ = v125
                                                } else {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    var v126 string = v45
                                                    _ = v126
                                                }
                                            }
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            _ = fi76
                                        }
                                    }
                                }
                            case "a":
                                for fi127 := 0; fi127 < 5; fi127++ {
                                    v1 -= v1
                                    _ = fi127
                                }
                            }
                            _ = fi74
                        }
                    }
                }
            }
            v128 := GenS11{F0: 0, F1: int64(v4), F2: 700}
            _ = v128
        }
        fmt.Println(v45, v1, v56)
    } else {
        switch "hello" {
        case "foo bar":
            fmt.Println(v4, v1, v56)
        default:
            v4 += v4
        }
        v45 = "hello"
        v129 := GenS18{F0: 0, F1: v2}
        _ = v129
    }
    v3 = (v3 >> (((^(v56))) & 63))
    var v130 uint32 = 1
    var v131 int64 = 42
    var v132 int8 = 100
    for fi133 := 0; fi133 < 3; fi133++ {
        switch 42 {
        case 32767:
            v4 = (v4 & (v4 - v4))
        }
        for fi134 := 0; fi134 < 17; fi134++ {
            v135 := GenS14{F0: (v4 - (^(v4)))}
            v136 := GenS14{F0: (^(v4))}
            _ = fi134
            _ = v135
            _ = v136
        }
        if ((true || false) || (!false)) {
            for fi137 := 0; fi137 < 5; fi137++ {
                v138 := GenS7{F0: uint32(v4)}
                var v139 uint = v56
                switch (fi137 - (fi137 << ((1) & 63))) {
                case -614:
                    switch "hello" {
                    case "foo bar":
                        if true {
                            v140 := GenS5{F0: -75}
                            var v141 uint64 = 0
                            _ = v140
                            _ = v141
                        }
                        var v142 uint16 = v4
                        _ = v142
                    case "hello":
                        v4 = uint16(v130)
                    default:
                        v130 = 0
                    }
                case -1:
                    fmt.Println(fi133)
                default:
                    var v143 bool = true
                    _ = v143
                }
                _ = fi137
                _ = v138
                _ = v139
            }
            var v144 bool = false
            v145 := GenS6{F0: 12, F1: 4294967295, F2: -18}
            _ = v144
            _ = v145
        } else {
            v3 -= 128
            v146 := GenS14{F0: 2}
            v131 -= v131
            _ = v146
        }
        _ = fi133
    }
    var v147 uint8 = (v3 >> ((1) & 63))
    v148 := GenS13{F0: 42}
    if true {
        v2 -= (v2 * (v2 & 42))
    }
    for fi149 := 0; fi149 < 18; fi149++ {
        for fi150 := 0; fi150 < 19; fi150++ {
            v4 -= v4
            var v151 uint8 = (v147 << ((1) & 63))
            if ((!true) || (v132 != v132)) {
                v1 += 1e10
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            _ = fi150
            _ = v151
        }
        if (-2.25 == 0.0) {
            switch -2147483648 {
            case 4294967295:
                v152 := GenS19{F0: v131, F1: 9223372036854775807}
                _ = v152
            case -469:
                v153 := GenS12{F0: 65535, F1: fi149, F2: 201}
                _ = v153
            default:
                v154 := GenS7{F0: (v130 / ((2) | 1))}
                _ = v154
            }
            v131 -= v131
            var v155 uint64 = 857
            _ = v155
        }
        for fi156 := 0; fi156 < 10; fi156++ {
            fmt.Println(v130, fi156)
            if len(v5) > 0 {
                v5[0] = fi156
            }
            _ = fi156
        }
        _ = fi149
    }
    fmt.Println("v1:", v1)
    fmt.Println("v130:", v130)
    fmt.Println("v131:", v131)
    fmt.Println("v132:", v132)
    fmt.Println("v147:", v147)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v130
    _ = v131
    _ = v132
    _ = v147
    _ = v148
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v56
}
