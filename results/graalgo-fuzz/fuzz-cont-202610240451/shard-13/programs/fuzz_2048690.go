package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8; F1 uint32 }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint32; F1 int64 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int32; F1 uint64; F2 int32 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int8; F1 int32 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 bool; F1 uint16 }
type GenS18 struct { F0 int8; F1 int16 }
type GenS19 struct { F0 uint; F1 bool; F2 int8 }
type GenS20 struct { F0 int32 }
type GenS21 struct { F0 int8; F1 int16; F2 string }
type GenS22 struct { F0 uint32; F1 int64 }
type GenS23 struct { F0 uint32 }
func genFn0(p0 uint64) uint32 {
    p0 = 67
    var v7 float64 = float64((1e10 / 1.5))
    _ = p0
    _ = v7
    return uint32(p0)
}
func genFn1(p0 bool, p1 int64) uint16 {
    p0 = p0
    p1 -= p1
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(p1, p1)
        _ = fi8
    }
    v9 := GenS1{F0: 14}
    v10 := GenS1{F0: -104}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 245
}

func main() {
    var v1 float32 = 1.5
    v2 := []int{0}
    v3 := []string{("" + (("foo bar" + "foo bar") + "")), "a"}
    v4 := []string{"hello", ""}
    v5 := []uint32{42, 1}
    v6 := make(map[string]int)
    for fi11 := 0; fi11 < 2; fi11++ {
        switch ("hello" + "hello") {
        case "hello":
            if false {
                fmt.Println(fi11, v1)
                var v12 int = -1
                _ = v12
            }
        case "":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case "foo bar":
            switch (fi11 * fi11) {
            case -680:
                var v13 int16 = 947
                _ = v13
            case 2147483647:
                v1 += v1
            }
        default:
            fmt.Println(fi11)
        }
        delete(v6, ("\n\t" + "hello"))
        for fi14 := 0; fi14 < 15; fi14++ {
            fmt.Println(v1, fi11)
            switch (fi11 - (fi11 >> ((238) & 63))) {
            case 32768:
                v15 := GenS1{F0: 1}
                _ = v15
            case 0:
                _ = len(v4)
            }
            v1 -= 3.14159
            _ = len(v6)
            v1 -= 3.14159
            _ = fi14
        }
        _ = fi11
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        if false {
            for fi17 := 0; fi17 < 11; fi17++ {
                v18 := GenS2{F0: 0, F1: 17}
                _ = fi17
                _ = v18
            }
            if (18446744073709551615 >= 120) {
                v1 = (float32(((3.14159 - ((3.14159 / 0.1) / 1.0)) * 0.0)) * float32(1e10))
                var v19 int8 = 6
                _ = v19
            } else {
                switch -1 {
                case 0:
                    for fi20 := 0; fi20 < 13; fi20++ {
                        v21 := GenS2{F0: 1, F1: 137}
                        v22 := GenS1{F0: 100}
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                default:
                    v1 += 1.5
                }
                _ = len(v2)
                for fi23 := 0; fi23 < 3; fi23++ {
                    if (((!(64 > 95)) || (530 != 927)) || (!false)) {
                        v1 -= v1
                    } else {
                        v1 = 1e10
                        if true {
                            switch ("foo bar" + "") {
                            case "":
                                fmt.Println(fi16, v1)
                                v1 = v1
                            }
                            fmt.Println(fi16)
                        } else {
                            v6[""] = (fi16 % ((fi23) | 1))
                            _ = len(v5)
                            var v24 bool = (1 != 0)
                            _ = v24
                        }
                        v25 := GenS1{F0: -128}
                        v1 = (((v1 - (float32(float64((v1 * (float32(1e10) * v1)))) + v1)) / (-2.25 / v1)) - 3.14159)
                        _ = v25
                    }
                    switch fi16 {
                    case -1:
                        fmt.Println(fi16, v1)
                    case 1:
                        if true {
                            var v26 int64 = 1
                            _ = v26
                        } else {
                            v27 := GenS3{F0: 0}
                            v6["foo bar"] = fi23
                            var v28 uint = 65535
                            var v29 int64 = -452
                            v30 := GenS2{F0: 32767, F1: 0}
                            var v31 int8 = 49
                            var v32 float64 = 3.14159
                            if (2147483647 < -559) {
                                var v33 int16 = 79
                                if false {
                                    v29 = v29
                                }
                                if (0 == -1) {
                                    var v34 string = ("a" + ("a" + ""))
                                    _ = v34
                                }
                                v31 = int8(v29)
                                var v35 uint64 = 169
                                _ = v33
                                _ = v35
                            } else {
                                v2 = append(v2, (^(fi16)))
                                if true {
                                    v36 := GenS4{F0: 4294967295}
                                    v28 += (v28 + 653)
                                    fmt.Println(v32, v1, fi23)
                                    _ = v36
                                } else {
                                    if (2147483647 < 446) {
                                        var v37 int64 = v29
                                        var v38 int16 = -32768
                                        v29 = (^(v37))
                                        _ = v37
                                        _ = v38
                                    }
                                }
                            }
                            _ = v27
                            _ = v28
                            _ = v29
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                    default:
                        v1 = ((float32(((1e10 / 1.0) - -2.25)) - 0.0) / v1)
                        if false {
                            if (false || false) {
                                switch "foo bar" {
                                case "hello":
                                    v39 := GenS3{F0: 1}
                                    _ = v39
                                }
                                v6[("" + "\n\t")] = 4294967295
                                v1 += (v1 * 1.5)
                            } else {
                                var v40 string = "hello"
                                var v41 uint64 = 878
                                if true {
                                    _ = len(v6)
                                    var v42 uint64 = (v41 >> ((560) & 63))
                                    fmt.Println(v1, fi23)
                                    _ = v42
                                }
                                _ = v40
                                _ = v41
                            }
                            if true {
                                if true {
                                    v1 = v1
                                    if true {
                                        v1 = (1.5 - -2.25)
                                        if (!true) {
                                            if (571 <= 2) {
                                                fmt.Println(fi16)
                                                if true {
                                                    if (!(42 > 65535)) {
                                                        v1 += float32(-2.25)
                                                        v6["foo bar"] = (fi16 ^ (fi16 | -9223372036854775808))
                                                        v43 := GenS5{F0: -67, F1: 951}
                                                        _ = v43
                                                    } else {
                                                        var v44 uint = 255
                                                        v45 := GenS1{F0: -58}
                                                        _ = v44
                                                        _ = v45
                                                    }
                                                    var v46 int8 = -100
                                                    v47 := GenS2{F0: -60, F1: 1}
                                                    v48 := GenS6{F0: 42, F1: 17}
                                                    if false {
                                                        var v49 uint16 = 491
                                                        var v50 int = 0
                                                        if true {
                                                            switch 2147483647 {
                                                            case 92:
                                                                v49 -= v49
                                                            case -959:
                                                                var v51 bool = (!true)
                                                                _ = v51
                                                            case 9223372036854775807:
                                                                v52 := GenS7{F0: v46}
                                                                _ = v52
                                                            default:
                                                                v53 := GenS8{F0: (-(v46))}
                                                                _ = v53
                                                            }
                                                            fmt.Println(fi16, v49, v49)
                                                            if (!(((1 <= -67) || true) && true)) {
                                                                v49 = v49
                                                                v50 += fi23
                                                                fmt.Println(v49, v50)
                                                                switch ("foo bar" + "") {
                                                                case "":
                                                                    var v54 uint64 = 18446744073709551615
                                                                    _ = v54
                                                                }
                                                            } else {
                                                                var v55 uint = 979
                                                                if (32767 >= -1) {
                                                                    fmt.Println(fi16)
                                                                }
                                                                if true {
                                                                    if (!true) {
                                                                        v56 := GenS4{F0: 2}
                                                                        var v57 int = int(v55)
                                                                        v55 = (v55 / ((255) | 1))
                                                                        _ = v56
                                                                        _ = v57
                                                                    } else {
                                                                        v58 := GenS9{F0: genFn0(2), F1: 0}
                                                                        _ = v58
                                                                    }
                                                                    v49 = 2
                                                                    if (263 >= 852) {
                                                                        v1 += ((v1 / (-2.25 - (v1 / ((float32(0.0) + (((v1 - v1) * (v1 * float32((float64(float32((0.1 - float64(1.5)))) / (1.5 / 1.5))))) - v1)) + (((((1.5 + 3.14159) / 0.1) + 0.1) - float32(0.1)) + v1))))) * (-2.25 / 1e10))
                                                                        if (float32(0.1) != v1) {
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                v59 := GenS10{F0: (("hello" + "a") + "")}
                                                                                _ = v59
                                                                            case "hello":
                                                                                switch ("" + ("hello" + "")) {
                                                                                case "a":
                                                                                    var v60 string = "foo bar"
                                                                                    _ = v60
                                                                                case "hello":
                                                                                    if true {
                                                                                        switch (fi16 * (v50 * (fi16 & (v50 << ((v55) & 63))))) {
                                                                                        case 42:
                                                                                            var v61 uint64 = 210
                                                                                            _ = v61
                                                                                        }
                                                                                        switch "foo bar" {
                                                                                        case "":
                                                                                            var v62 int8 = v46
                                                                                            _ = v62
                                                                                        case "\n\t":
                                                                                            v63 := GenS5{F0: (^(v46)), F1: 1}
                                                                                            _ = v63
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(v49)
                                                                                        var v64 string = ""
                                                                                        if true {
                                                                                            var v65 float64 = (-2.25 * (3.14159 / ((1.5 + 0.0) * (0.0 + float64((v1 - float32(-2.25)))))))
                                                                                            _ = v65
                                                                                        } else {
                                                                                            if false {
                                                                                                v64 = (v64 + ("\n\t" + ""))
                                                                                                if false {
                                                                                                    v55 = v55
                                                                                                    var v66 int64 = 42
                                                                                                    _ = v66
                                                                                                } else {
                                                                                                    v49 += 401
                                                                                                    v46 = v46
                                                                                                    v67 := GenS5{F0: (v46 * (v46 / (((v46 * v46)) | 1))), F1: 30}
                                                                                                    _ = v67
                                                                                                }
                                                                                            } else {
                                                                                                if len(v4) >= 2 {
                                                                                                    v4 = v4[1:len(v4)]
                                                                                                }
                                                                                                v49 -= v49
                                                                                                v64 = v64
                                                                                            }
                                                                                        }
                                                                                        var v68 int32 = 502
                                                                                        var v69 int8 = 42
                                                                                        _ = v64
                                                                                        _ = v68
                                                                                        _ = v69
                                                                                    }
                                                                                default:
                                                                                    v70 := GenS2{F0: 1000, F1: 1}
                                                                                    _ = v70
                                                                                }
                                                                            case "":
                                                                                var v71 uint32 = 42
                                                                                _ = v71
                                                                            default:
                                                                                v72 := GenS1{F0: v46}
                                                                                _ = v72
                                                                            }
                                                                            v55 = v55
                                                                        }
                                                                    }
                                                                } else {
                                                                    delete(v6, ("foo bar" + "hello"))
                                                                    v46 = v46
                                                                    switch fi23 {
                                                                    case 0:
                                                                        _ = len(v6)
                                                                    case -9223372036854775808:
                                                                        var v73 uint16 = (v49 * v49)
                                                                        _ = v73
                                                                    }
                                                                    _ = len(v6)
                                                                    var v74 int8 = (v46 / ((v46) | 1))
                                                                    _ = v74
                                                                }
                                                                _ = v55
                                                            }
                                                        }
                                                        _ = v49
                                                        _ = v50
                                                    }
                                                    v75 := GenS11{F0: -24, F1: 682, F2: int32(v46)}
                                                    _ = v46
                                                    _ = v47
                                                    _ = v48
                                                    _ = v75
                                                } else {
                                                    fmt.Println(v1, fi16)
                                                    if (0 == 71) {
                                                        v76 := GenS4{F0: 2147483648}
                                                        v1 -= 0.1
                                                        _ = v76
                                                    } else {
                                                        var v77 string = "foo bar"
                                                        var v78 int32 = -2147483648
                                                        var v79 int32 = v78
                                                        _ = v77
                                                        _ = v78
                                                        _ = v79
                                                    }
                                                    if false {
                                                        v1 = 1e10
                                                    } else {
                                                        v80 := GenS12{F0: (fi23 + fi16)}
                                                        _ = v80
                                                    }
                                                }
                                            } else {
                                                fmt.Println(v1, v1)
                                                var v81 int32 = int32(fi16)
                                                var v82 int8 = 42
                                                _ = v81
                                                _ = v82
                                            }
                                            if true {
                                                var v83 float64 = 0.1
                                                v83 = (v83 + -2.25)
                                                _ = v83
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            if true {
                                                v1 += v1
                                            }
                                            var v84 int8 = 127
                                            _ = v84
                                        } else {
                                            var v85 int64 = -281
                                            _ = v85
                                        }
                                        _ = len(v6)
                                    } else {
                                        var v86 float32 = v1
                                        if (-128 == 42) {
                                            if (!(1e10 > -2.25)) {
                                                var v87 int8 = -1
                                                v87 = -128
                                                var v88 bool = true
                                                _ = v87
                                                _ = v88
                                            }
                                            var v89 int32 = -2147483648
                                            _ = v89
                                        } else {
                                            v86 += 0.0
                                            if true {
                                                if true {
                                                    fmt.Println(v1, fi16, v1)
                                                } else {
                                                    v90 := v6["a"]
                                                    if (false && (!true)) {
                                                        v91 := GenS13{F0: 77, F1: 772}
                                                        _ = v91
                                                    }
                                                    _ = len(v2)
                                                    _ = v90
                                                }
                                            }
                                            v92 := v6["\n\t"]
                                            _ = len(v4)
                                            _ = v92
                                        }
                                        _ = v86
                                    }
                                    v1 = (1e10 - v1)
                                    v93 := GenS14{F0: 2}
                                    if false {
                                        v1 -= 3.14159
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v94 bool = true
                                        _ = v94
                                    }
                                    if false {
                                        if (32767 == 1) {
                                            switch "\n\t" {
                                            case "foo bar":
                                                v95 := GenS15{F0: 65535, F1: 0, F2: genFn1((-519 == -695), -1)}
                                                _ = v95
                                            case "\n\t":
                                                switch "" {
                                                case "a":
                                                    var v96 int = -720
                                                    _ = v96
                                                }
                                            }
                                            v1 += (v1 / v1)
                                            v1 = ((v1 * 0.1) / (float32(3.14159) * v1))
                                        } else {
                                            var v97 int = (fi23 % ((0) | 1))
                                            v1 = (v1 * ((v1 - -2.25) - 3.14159))
                                            if false {
                                                if ("foo bar" < "\n\t") {
                                                    var v98 uint = 2
                                                    _ = v98
                                                }
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                if (2 == 393) {
                                                    var v99 int16 = 1
                                                    _ = v99
                                                }
                                            } else {
                                                var v100 int16 = -486
                                                v97 = -499
                                                v1 = v1
                                                _ = v100
                                            }
                                            _ = v97
                                        }
                                        switch "hello" {
                                        case "hello":
                                            var v101 uint8 = 255
                                            _ = v101
                                        }
                                        v1 += float32(3.14159)
                                    } else {
                                        v2 = append(v2, (fi16 & 73))
                                        v1 = (1.5 + 0.0)
                                        if ((1e10 > 0.0) && ((715 < 2) || true)) {
                                            _ = len(v5)
                                            if (1 >= 0) {
                                                if true {
                                                    var v102 bool = false
                                                    _ = v102
                                                }
                                                var v103 float64 = (3.14159 * 0.0)
                                                _ = v103
                                            } else {
                                                if false {
                                                    fmt.Println(fi23, fi23, v1)
                                                    fmt.Println(fi23, fi23)
                                                    var v104 int8 = -1
                                                    _ = v104
                                                }
                                                var v105 uint = 2
                                                v105 = v105
                                                _ = v105
                                            }
                                        } else {
                                            fmt.Println(v1)
                                            switch (fi23 >> ((817) & 63)) {
                                            case 2147483647:
                                                v106 := GenS16{F0: 2}
                                                _ = v106
                                            }
                                            if (-1 == -32768) {
                                                v107 := GenS17{F0: true, F1: 610}
                                                v6[(("hello" + "hello") + "a")] = (fi23 >> ((0) & 63))
                                                _ = v107
                                            } else {
                                                fmt.Println(fi16, v1)
                                            }
                                        }
                                        var v108 float64 = -2.25
                                        var v109 uint = 65535
                                        fmt.Println(fi23, fi16)
                                        fmt.Println(v109)
                                        if false {
                                            v110 := v6["hello"]
                                            _ = v110
                                        } else {
                                            v108 -= v108
                                        }
                                        _ = v108
                                        _ = v109
                                    }
                                    var v111 int16 = 1
                                    v112 := GenS18{F0: 0, F1: 887}
                                    switch "\n\t" {
                                    case "a":
                                        var v113 int16 = v111
                                        _ = v113
                                    case "foo bar":
                                        v111 = -976
                                    }
                                    _ = v111
                                    _ = v112
                                    _ = v93
                                }
                            }
                            var v114 float64 = -2.25
                            _ = v114
                        }
                        var v115 bool = ((false || true) || true)
                        _ = v115
                    }
                    v1 = v1
                    var v116 uint64 = 18446744073709551615
                    fmt.Println(v1, fi16, fi23)
                    _ = fi23
                    _ = v116
                }
            }
            fmt.Println(v1)
        } else {
            v3 = append(v3, "a")
            v117 := GenS9{F0: 602, F1: 0}
            v1 += 1e10
            _ = v117
        }
        _ = fi16
    }
    v118 := GenS3{F0: 2}
    switch "foo bar" {
    case "":
        v119 := GenS19{F0: 351, F1: (true && true), F2: 1}
        _ = v119
    }
    v120 := GenS20{F0: 844}
    if false {
        v121 := GenS9{F0: 0, F1: 9223372036854775807}
        v122 := GenS15{F0: genFn1(true, -1), F1: -1, F2: 4}
        _ = v121
        _ = v122
    } else {
        fmt.Println(v1)
        for fi123 := 0; fi123 < 8; fi123++ {
            if (0 <= -805) {
                switch "foo bar" {
                case "a":
                    v1 = (3.14159 + float32(-2.25))
                }
            } else {
                if (!true) {
                    fmt.Println(v1, fi123)
                    if (42 > 42) {
                        v1 = 1.5
                        var v124 int64 = -421
                        var v125 int8 = 127
                        _ = v124
                        _ = v125
                    } else {
                        var v126 uint64 = 1
                        v126 += v126
                        switch (fi123 ^ fi123) {
                        case 568:
                            v127 := GenS17{F0: ((255 > 1) && ("foo bar" <= "foo bar")), F1: 255}
                            _ = v127
                        case -711:
                            v128 := GenS21{F0: 100, F1: 0, F2: "\n\t"}
                            _ = v128
                        case 2147483647:
                            v129 := GenS4{F0: genFn0((v126 | 42))}
                            _ = v129
                        default:
                            v130 := GenS6{F0: genFn1((false || true), int64(v126)), F1: 2}
                            _ = v130
                        }
                        _ = v126
                    }
                } else {
                    v131 := GenS22{F0: 778, F1: 42}
                    if true {
                        var v132 bool = true
                        _ = v132
                    } else {
                        v133 := GenS12{F0: (fi123 >> ((2) & 63))}
                        _ = v133
                    }
                    delete(v6, (("hello" + "") + "foo bar"))
                    _ = v131
                }
                switch (^(fi123)) {
                case 669:
                    var v134 int16 = 1000
                    _ = v134
                case 0:
                    var v135 int64 = 9223372036854775807
                    _ = v135
                case 428:
                    v1 += -2.25
                default:
                    v1 = v1
                }
                if ((0 == 2) && false) {
                    var v136 int8 = 42
                    v136 += 62
                    for fi137 := 0; fi137 < 4; fi137++ {
                        switch "\n\t" {
                        case "":
                            v1 = 0.0
                        case "a":
                            v138 := GenS23{F0: genFn0(4294967295)}
                            _ = v138
                        default:
                            var v139 int16 = 1000
                            _ = v139
                        }
                        if false {
                            switch (fi123 % (((fi123 - fi123)) | 1)) {
                            case 32767:
                                v140 := GenS6{F0: 65535, F1: 128}
                                _ = v140
                            case 127:
                                v141 := GenS17{F0: (!true), F1: 1}
                                _ = v141
                            default:
                                v136 += v136
                            }
                            v142 := GenS1{F0: (v136 / ((v136) | 1))}
                            switch 128 {
                            case 1:
                                v136 += (v136 << ((4294967295) & 63))
                            case 127:
                                var v143 uint8 = 128
                                _ = v143
                            }
                            _ = v142
                        }
                        _ = fi137
                    }
                    _ = v136
                }
            }
            _ = fi123
        }
    }
    fmt.Println(v1)
    v1 -= float32(0.0)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v118
    _ = v120
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
