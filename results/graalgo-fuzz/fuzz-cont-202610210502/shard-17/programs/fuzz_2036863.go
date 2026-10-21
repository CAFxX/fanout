package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 int32; F2 bool }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 uint16; F1 int; F2 uint }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint; F1 int8; F2 int8 }
type GenS9 struct { F0 uint8; F1 int16; F2 uint }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint; F1 uint16 }
type GenS12 struct { F0 int32; F1 int16 }
type GenS13 struct { F0 uint8; F1 int32 }
type GenS14 struct { F0 string; F1 int64 }
type GenS15 struct { F0 uint; F1 uint; F2 uint64 }
type GenS16 struct { F0 uint32; F1 string }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 int16 }
type GenS20 struct { F0 bool }
type GenS21 struct { F0 int64; F1 uint64; F2 uint }
func genFn0(p0 bool, p1 uint) uint {
    if false {
        p0 = true
    }
    v4 := GenS1{F0: 42, F1: 871}
    _ = p0
    _ = p1
    _ = v4
    return p1
}

func main() {
    var v1 uint8 = 238
    v2 := make(map[string]string)
    v3 := make(map[int]bool)
    v1 = v1
    var v5 uint64 = 836
    var v6 float64 = (float64((1e10 - (-2.25 + (0.1 * (1.5 * 3.14159))))) * ((1.5 * float64(-2.25)) + 0.0))
    switch 4294967295 {
    case 788:
        v5 += 0
    case 848:
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := GenS1{F0: genFn0(false, 0), F1: -1}
            delete(v3, int(v1))
            _ = fi7
            _ = v8
        }
    }
    fmt.Println(v1)
    v6 = 1.5
    switch 127 {
    case 32768:
        if false {
            v5 -= (v5 >> ((296) & 63))
            fmt.Println(v1, v5, v6)
            switch 32768 {
            case -2147483648:
                var v9 float32 = ((1.5 + float32(v6)) + 0.1)
                _ = v9
            case 308:
                v5 = v5
            case 361:
                var v10 int = 2147483647
                _ = v10
            default:
                v11 := GenS1{F0: 255, F1: 999}
                _ = v11
            }
        } else {
            v1 = (v1 << ((570) & 63))
        }
    case 575:
        if (v1 >= v1) {
            if ("foo bar" < "foo bar") {
                v12 := GenS1{F0: 1, F1: 830}
                var v13 int = -1
                _ = v12
                _ = v13
            } else {
                fmt.Println(v1, v1, v6)
                for fi14 := 0; fi14 < 11; fi14++ {
                    var v15 int64 = -1
                    _ = fi14
                    _ = v15
                }
            }
        }
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        v6 += (v6 / v6)
        _ = fi16
    }
    var v17 uint16 = 2
    for fi18 := 0; fi18 < 4; fi18++ {
        v6 -= ((v6 + (1e10 - -2.25)) / v6)
        for fi19 := 0; fi19 < 18; fi19++ {
            var v20 uint16 = (v17 - 443)
            v21 := GenS2{F0: (^(v20))}
            if false {
                var v22 int32 = 42
                var v23 float32 = ((3.14159 * 1e10) - float32(float64(1.5)))
                _ = v22
                _ = v23
            } else {
                switch fi18 {
                case 32768:
                    switch fi18 {
                    case 0:
                        if true {
                            if ((v1 - v1) < 191) {
                                var v24 float32 = 1e10
                                _ = v24
                            } else {
                                v20 = (v17 / ((v20) | 1))
                            }
                            if true {
                                var v25 int = fi18
                                _ = len(v2)
                                var v26 float64 = (v6 * 3.14159)
                                _ = v25
                                _ = v26
                            } else {
                                v27 := GenS3{F0: -876, F1: int32(v5), F2: (2 > v20)}
                                fmt.Println(v1, v20)
                                v5 -= uint64(v1)
                                if false {
                                    var v28 int64 = 9223372036854775807
                                    v5 += (v5 * 4294967295)
                                    var v29 uint8 = 2
                                    switch (("hello" + ("" + "foo bar")) + "") {
                                    case "hello":
                                        var v30 uint8 = 42
                                        _ = v30
                                    case "\n\t":
                                        switch fi18 {
                                        case -226:
                                            var v31 bool = true
                                            _ = v31
                                        }
                                    }
                                    _ = v28
                                    _ = v29
                                } else {
                                    v32 := GenS3{F0: 0, F1: 42, F2: ((v5 >> ((1) & 63)) <= 42)}
                                    v33 := GenS4{F0: 1000, F1: 1}
                                    _ = v32
                                    _ = v33
                                }
                                var v34 int8 = 42
                                if false {
                                    v20 -= (v20 % ((v17) | 1))
                                    v17 -= v20
                                }
                                if (uint(v5) <= uint(v1)) {
                                    v20 = (v17 + (v20 << ((genFn0(((2 != 2147483648) && (v34 > v34)), genFn0(false, 598))) & 63)))
                                    var v35 int8 = v34
                                    _ = v35
                                } else {
                                    v1 = (v1 << ((genFn0((!(!(false || false))), uint(v5))) & 63))
                                }
                                _ = v27
                                _ = v34
                            }
                            v17 -= 2
                            v6 = v6
                            if false {
                                v36 := GenS5{F0: fi19, F1: -644}
                                switch fi19 {
                                case 127:
                                    v1 += v1
                                case 128:
                                    var v37 int = -2147483648
                                    _ = v37
                                case 0:
                                    var v38 int64 = 664
                                    _ = v38
                                default:
                                    v20 -= v20
                                }
                                switch "" {
                                case "a":
                                    var v39 int32 = 2147483647
                                    _ = v39
                                case "\n\t":
                                    v40 := GenS3{F0: -575, F1: -1, F2: (127 < 1)}
                                    _ = v40
                                case "hello":
                                    v41 := GenS5{F0: int(v1), F1: -711}
                                    _ = v41
                                }
                                _ = v36
                            } else {
                                if true {
                                    if (955 == 42) {
                                        if true {
                                            var v42 int16 = 42
                                            v6 = ((v6 / (0.1 / -2.25)) / ((0.1 / (0.0 + (float64(-2.25) - (float64(((1e10 * -2.25) + (1e10 * float32(((((0.0 * v6) - 0.1) / (v6 / v6)) * (float64(0.1) / ((v6 / (-2.25 - ((0.1 * (v6 - -2.25)) * -2.25))) * v6))))))) * v6)))) / 0.1))
                                            v43 := GenS3{F0: -1, F1: 1, F2: (true && ((-550 >= -1) && true))}
                                            _ = v42
                                            _ = v43
                                        } else {
                                            if ((42 == 1) || true) {
                                                fmt.Println(v1, v1, fi19)
                                                switch "a" {
                                                case "a":
                                                    fmt.Println(fi18)
                                                case "hello":
                                                    v17 = 32767
                                                case "foo bar":
                                                    if true {
                                                        if ((fi19 & (^(fi18))) >= fi19) {
                                                            v44 := GenS6{F0: (v17 & uint16(v5)), F1: 32767, F2: 0}
                                                            v45 := []int{(fi18 << ((uint(v17)) & 63)), fi18, 2147483647, fi18}
                                                            _ = v44
                                                            _ = v45
                                                        } else {
                                                            var v46 int8 = 42
                                                            if (255 < 1) {
                                                                v6 += v6
                                                                if (-2.25 < (float32(1.5) - ((0.1 / -2.25) - 0.0))) {
                                                                    switch (("\n\t" + "hello") + "\n\t") {
                                                                    case "foo bar":
                                                                        if true {
                                                                            v47 := []int{}
                                                                            var v48 uint32 = 42
                                                                            delete(v2, ("a" + "hello"))
                                                                            if false {
                                                                                switch fi18 {
                                                                                case 0:
                                                                                    if (0.1 != (v6 - 1e10)) {
                                                                                        fmt.Println(v5, fi18, v46)
                                                                                    } else {
                                                                                        var v49 int64 = -9223372036854775808
                                                                                        _ = v49
                                                                                    }
                                                                                default:
                                                                                    if true {
                                                                                        var v50 float64 = v6
                                                                                        _ = v50
                                                                                    }
                                                                                }
                                                                            }
                                                                            _ = v47
                                                                            _ = v48
                                                                        }
                                                                    default:
                                                                        v5 += (^(v5))
                                                                    }
                                                                    v51 := GenS6{F0: v20, F1: (fi19 / (((^(fi19))) | 1)), F2: 49}
                                                                    v52 := GenS2{F0: (v20 ^ (^(v20)))}
                                                                    _ = v51
                                                                    _ = v52
                                                                } else {
                                                                    var v53 int32 = 2147483647
                                                                    v53 = (-(v53))
                                                                    _ = v53
                                                                }
                                                            } else {
                                                                if ((v1 / ((v1) | 1)) <= (v1 + v1)) {
                                                                    fmt.Println(v17, fi18, v46)
                                                                    if ((true || true) && (!(!false))) {
                                                                        var v54 int = 32767
                                                                        fmt.Println(v6, v5)
                                                                        _ = v54
                                                                    }
                                                                } else {
                                                                    v46 += v46
                                                                }
                                                                switch ((("foo bar" + "\n\t") + "") + "\n\t") {
                                                                case "a":
                                                                    delete(v3, fi18)
                                                                case "":
                                                                    v55 := []int{fi18, (fi19 / ((-9223372036854775808) | 1)), fi19, fi18}
                                                                    _ = v55
                                                                default:
                                                                    fmt.Println(v17)
                                                                }
                                                                var v56 int16 = 267
                                                                v17 = v20
                                                                _ = v56
                                                            }
                                                            fmt.Println(v46, v46, fi18)
                                                            _ = v46
                                                        }
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v17 -= v20
                                                        case "":
                                                            if false {
                                                                v1 -= (v1 - (v1 << ((631) & 63)))
                                                                var v57 int32 = 0
                                                                _ = v57
                                                            }
                                                        default:
                                                            v58 := []int{(fi18 - (fi19 ^ fi19)), fi18, (fi18 - fi19)}
                                                            _ = v58
                                                        }
                                                        v59 := GenS3{F0: 42, F1: 1, F2: true}
                                                        _ = v59
                                                    }
                                                    var v60 int32 = 0
                                                    v61 := GenS1{F0: 4294967295, F1: 848}
                                                    _ = v60
                                                    _ = v61
                                                }
                                            }
                                            var v62 uint16 = v17
                                            _ = v62
                                        }
                                    } else {
                                        v6 = (v6 * v6)
                                        _ = len(v2)
                                        var v63 string = "hello"
                                        _ = v63
                                    }
                                    v6 -= v6
                                }
                                v64 := GenS7{F0: 91}
                                v5 = v5
                                _ = v64
                            }
                        } else {
                            if ("hello" > "\n\t") {
                                v6 -= v6
                                v1 = v1
                            } else {
                                v1 -= v1
                                v6 -= v6
                            }
                            v5 = v5
                        }
                    default:
                        v6 -= (0.1 + 0.0)
                    }
                default:
                    switch (fi19 & fi18) {
                    case 613:
                        if ((0.0 + 0.1) <= 1e10) {
                            v5 = 17
                            v65 := []int{1, fi19}
                            _ = v65
                        } else {
                            v1 -= (v1 - uint8(v5))
                            var v66 uint = genFn0(false, 42)
                            _ = v66
                        }
                    default:
                        if false {
                            switch "hello" {
                            case "foo bar":
                                var v67 int16 = 1
                                switch "" {
                                case "hello":
                                    if (!((("" + "foo bar") + "a") == (("hello" + "\n\t") + "\n\t"))) {
                                        fmt.Println(v20)
                                        var v68 bool = ((!(266 <= 0)) || true)
                                        _ = v68
                                    } else {
                                        v6 = v6
                                        if true {
                                            if ((-(v67)) <= (-(v67))) {
                                                v17 = uint16(v5)
                                                if false {
                                                    var v69 uint8 = 1
                                                    v70 := GenS1{F0: 18446744073709551615, F1: (v67 * 177)}
                                                    v6 = v6
                                                    _ = v69
                                                    _ = v70
                                                } else {
                                                    v71 := GenS1{F0: 275, F1: v67}
                                                    var v72 uint8 = (^(v1))
                                                    var v73 float64 = v6
                                                    v74 := GenS2{F0: (v17 - v17)}
                                                    _ = v71
                                                    _ = v72
                                                    _ = v73
                                                    _ = v74
                                                }
                                                if false {
                                                    v75 := v2[("a" + "a")]
                                                    v76 := []int{(^(fi18))}
                                                    var v77 float64 = v6
                                                    if true {
                                                        var v78 int16 = 1000
                                                        v20 = 32767
                                                        _ = v78
                                                    }
                                                    _ = v75
                                                    _ = v76
                                                    _ = v77
                                                } else {
                                                    var v79 uint64 = (^(v5))
                                                    _ = v79
                                                }
                                            }
                                            if (!true) {
                                                v6 -= v6
                                                var v80 int16 = v67
                                                _ = v80
                                            } else {
                                                var v81 string = ""
                                                v5 += v5
                                                var v82 uint = uint(v5)
                                                _ = v81
                                                _ = v82
                                            }
                                            v83 := GenS3{F0: 675, F1: -215, F2: (!true)}
                                            _ = v83
                                        } else {
                                            v84 := GenS8{F0: genFn0(((v6 + (0.1 * v6)) != v6), 0), F1: -1, F2: 110}
                                            var v85 int8 = 0
                                            var v86 int32 = 1
                                            _ = v84
                                            _ = v85
                                            _ = v86
                                        }
                                        switch "foo bar" {
                                        case "\n\t":
                                            delete(v2, "")
                                        case "foo bar":
                                            if false {
                                                fmt.Println(v5, v6)
                                                v5 = v5
                                            } else {
                                                switch fi18 {
                                                case -2147483648:
                                                    if (false || false) {
                                                        var v87 int32 = 831
                                                        var v88 string = "foo bar"
                                                        _ = v87
                                                        _ = v88
                                                    } else {
                                                        delete(v2, ((((("a" + ("\n\t" + "a")) + (("hello" + "foo bar") + "foo bar")) + "a") + "hello") + "\n\t"))
                                                        v20 -= (v20 & (v20 % ((0) | 1)))
                                                    }
                                                }
                                                if false {
                                                    v89 := []int{fi19}
                                                    v90 := GenS9{F0: v1, F1: (v67 << ((1) & 63)), F2: uint(v20)}
                                                    var v91 int = fi18
                                                    var v92 bool = (((((0.0 / 3.14159) * (float32(1e10) - ((0.1 + 3.14159) / float32(1e10)))) - ((float32((v6 + 3.14159)) - 3.14159) * 1e10)) != 1.5) || true)
                                                    _ = v89
                                                    _ = v90
                                                    _ = v91
                                                    _ = v92
                                                }
                                                switch fi19 {
                                                case 127:
                                                    v17 += 426
                                                case 42:
                                                    v93 := GenS10{F0: 9223372036854775807}
                                                    _ = v93
                                                }
                                            }
                                            if false {
                                                v94 := []int{fi19}
                                                switch ((("" + "") + "foo bar") + "hello") {
                                                case "\n\t":
                                                    v95 := v2["\n\t"]
                                                    _ = v95
                                                case "a":
                                                    var v96 uint8 = v1
                                                    _ = v96
                                                case "":
                                                    v17 = (v20 | 282)
                                                }
                                                _ = v94
                                            } else {
                                                var v97 float64 = -2.25
                                                v98 := GenS11{F0: uint(v5), F1: (v20 & uint16(v1))}
                                                if (int64(v17) > -9223372036854775808) {
                                                    v99 := []int{9223372036854775807, 1, fi18}
                                                    v100 := GenS12{F0: -389, F1: (v67 & -1)}
                                                    fmt.Println(v5)
                                                    _ = v100
                                                    _ = v99
                                                } else {
                                                    v1 = (^(v1))
                                                }
                                                v101 := []int{fi18, fi18, (fi18 | -1), (^(fi18))}
                                                var v102 int64 = 347
                                                v102 = v102
                                                v102 += (v102 / (((v102 / (((v102 ^ v102)) | 1))) | 1))
                                                _ = len(v2)
                                                v103 := v3[fi19]
                                                _ = v101
                                                _ = v102
                                                _ = v103
                                                _ = v97
                                                _ = v98
                                            }
                                            switch (fi19 >> ((4294967295) & 63)) {
                                            case -9223372036854775808:
                                                var v104 uint = 714
                                                _ = v104
                                            case 411:
                                                v17 -= (^(v20))
                                            }
                                        default:
                                            if (2 != (v5 | v5)) {
                                                fmt.Println(v6, v6, fi19)
                                                v105 := GenS13{F0: v1, F1: -2147483648}
                                                switch "foo bar" {
                                                case "a":
                                                    if ((466 == -2147483648) && false) {
                                                        switch "a" {
                                                        case "foo bar":
                                                            var v106 float64 = v6
                                                            _ = v106
                                                        default:
                                                            fmt.Println(fi18, v17, fi19)
                                                        }
                                                    }
                                                    if false {
                                                        v67 += -57
                                                        var v107 uint8 = v1
                                                        v107 = v107
                                                        v2["hello"] = ""
                                                        _ = v107
                                                    } else {
                                                        v6 -= (-2.25 + 3.14159)
                                                        v67 = v67
                                                        _ = len(v3)
                                                    }
                                                    delete(v3, fi19)
                                                case "hello":
                                                    v108 := GenS14{F0: "foo bar", F1: 0}
                                                    _ = v108
                                                }
                                                _ = v105
                                            } else {
                                                if false {
                                                    v109 := GenS5{F0: (fi19 / ((fi19) | 1)), F1: v67}
                                                    var v110 uint16 = (v20 << ((18446744073709551615) & 63))
                                                    var v111 int = fi19
                                                    _ = v109
                                                    _ = v110
                                                    _ = v111
                                                }
                                                var v112 uint64 = v5
                                                _ = v112
                                            }
                                        }
                                    }
                                case "a":
                                    v17 = (^(v20))
                                default:
                                    v113 := []int{}
                                    _ = v113
                                }
                                _ = v67
                            case "\n\t":
                                v5 = (v5 >> ((42) & 63))
                            case "a":
                                if (true || true) {
                                    var v114 float32 = 1.5
                                    var v115 int = fi19
                                    _ = v114
                                    _ = v115
                                }
                            }
                            var v116 uint8 = (v1 >> ((18446744073709551615) & 63))
                            v117 := []int{(fi18 / ((fi18) | 1))}
                            switch "\n\t" {
                            case "a":
                                v118 := GenS15{F0: 2, F1: 2, F2: v5}
                                _ = v118
                            case "":
                                var v119 uint16 = (^(v20))
                                if false {
                                    var v120 string = ""
                                    if false {
                                        if len(v117) >= 2 {
                                            v117 = v117[1:len(v117)]
                                        }
                                    } else {
                                        if (v6 == 0.1) {
                                            if len(v117) >= 2 {
                                                v117 = v117[1:len(v117)]
                                            }
                                            switch 209 {
                                            case 812:
                                                var v121 string = ((v120 + ((v120 + v120) + v120)) + v120)
                                                _ = v121
                                            case 1:
                                                var v122 uint = 4294967295
                                                var v123 uint8 = v1
                                                fmt.Println(fi19, v119)
                                                v120 = "hello"
                                                _ = v122
                                                _ = v123
                                            case 127:
                                                switch 4294967295 {
                                                case 2147483647:
                                                    v117 = append(v117, (fi18 | (fi18 << ((2) & 63))))
                                                }
                                            default:
                                                v5 = v5
                                                _ = len(v3)
                                            }
                                            fmt.Println(v119)
                                        } else {
                                            fmt.Println(v17)
                                            var v124 uint32 = 2
                                            v125 := GenS16{F0: 2147483648, F1: (v120 + (v120 + ""))}
                                            _ = v124
                                            _ = v125
                                        }
                                        var v126 int32 = -605
                                        _ = v126
                                    }
                                    _ = v120
                                }
                                v1 += v1
                                var v127 int32 = 2147483647
                                _ = v119
                                _ = v127
                            case "foo bar":
                                if (true && true) {
                                    if (!false) {
                                        var v128 string = ("hello" + ("hello" + ""))
                                        var v129 uint = genFn0((!true), 533)
                                        _ = v128
                                        _ = v129
                                    }
                                    v6 += v6
                                    v5 -= v5
                                }
                            }
                            v130 := GenS14{F0: "hello", F1: 150}
                            _ = v116
                            _ = v117
                            _ = v130
                        } else {
                            switch fi18 {
                            case -9223372036854775808:
                                if true {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v3[-1] = ((((fi19 ^ fi19) > fi18) || false) && false)
                                    case "foo bar":
                                        if (((((((3.14159 / 1.0) - 0.1) + ((3.14159 - 1.5) + ((3.14159 - 1.5) / 3.14159))) + (3.14159 * 0.0)) / (float32(float64(1.5)) + -2.25)) / (0.1 * (3.14159 * 3.14159))) == -2.25) {
                                            fmt.Println(fi18, v17, v5)
                                            var v131 uint64 = (v5 >> ((65535) & 63))
                                            _ = v131
                                        } else {
                                            var v132 int64 = 92
                                            _ = v132
                                        }
                                    }
                                    fmt.Println(v6, v17, v6)
                                    var v133 uint16 = v20
                                    v134 := GenS11{F0: 4294967295, F1: v133}
                                    v1 = (^(v1))
                                    if (!true) {
                                        fmt.Println(v133, v6, v20)
                                        v6 = float64((0.0 - (3.14159 + -2.25)))
                                    } else {
                                        v135 := v3[fi19]
                                        _ = v135
                                    }
                                    _ = v133
                                    _ = v134
                                }
                            case 128:
                                fmt.Println(v1, fi19, fi18)
                            }
                        }
                    }
                    var v136 uint = genFn0(((504 <= 2) || (fi18 < 32767)), 42)
                    _ = v136
                }
                if false {
                    delete(v3, (-(fi19)))
                    delete(v2, "\n\t")
                    v137 := GenS17{F0: 0}
                    _ = v137
                }
                v17 += 0
            }
            var v138 uint = genFn0(false, 800)
            v139 := GenS18{F0: false}
            switch (fi19 << ((genFn0((28 != 0), v138)) & 63)) {
            case -2147483648:
                v140 := GenS2{F0: uint16(v138)}
                _ = v140
            case 676:
                v1 += 42
            default:
                v20 = (v20 & 32767)
            }
            _ = fi19
            _ = v138
            _ = v139
            _ = v20
            _ = v21
        }
        _ = fi18
    }
    v1 += (^(v1))
    if false {
        var v141 bool = true
        v142 := GenS2{F0: (v17 | v17)}
        var v143 uint32 = 0
        _ = v141
        _ = v142
        _ = v143
    } else {
        if ((v5 <= v5) || (!(1e10 <= ((3.14159 + -2.25) + 0.0)))) {
            v144 := []int{98, 2147483647, -28, 127}
            _ = v144
        } else {
            switch 127 {
            case 32767:
                fmt.Println(v6, v5)
            case -9223372036854775808:
                for fi145 := 0; fi145 < 10; fi145++ {
                    for fi146 := 0; fi146 < 8; fi146++ {
                        v147 := []int{(fi145 - fi146), (fi145 & fi146), fi146}
                        if len(v147) > 0 {
                            v147[0] = -1
                        }
                        var v148 uint64 = 257
                        _ = fi146
                        _ = v147
                        _ = v148
                    }
                    _ = fi145
                }
            default:
                if true {
                    if (!true) {
                        var v149 uint8 = 2
                        _ = v149
                    }
                } else {
                    v150 := []int{32768}
                    if true {
                        switch (("\n\t" + "") + ((("" + ("\n\t" + "")) + "foo bar") + "hello")) {
                        case "\n\t":
                            if (-9223372036854775808 == -490) {
                                for fi151 := 0; fi151 < 9; fi151++ {
                                    fmt.Println(v1, fi151)
                                    fmt.Println(v5, v6, fi151)
                                    _ = fi151
                                }
                                v152 := GenS19{F0: -1}
                                _ = v152
                            } else {
                                _ = len(v3)
                                var v153 bool = true
                                var v154 uint16 = v17
                                _ = v153
                                _ = v154
                            }
                        case "foo bar":
                            for fi155 := 0; fi155 < 4; fi155++ {
                                if true {
                                    v156 := GenS9{F0: v1, F1: -1, F2: 65535}
                                    var v157 int8 = -128
                                    _ = v156
                                    _ = v157
                                }
                                fmt.Println(v5, v5, v6)
                                _ = fi155
                            }
                        default:
                            v6 -= v6
                        }
                        v158 := GenS20{F0: (((1.5 - 3.14159) + 3.14159) == ((((float32(v6) + 1.5) / float32(v6)) / -2.25) / ((-2.25 + 3.14159) + 1e10)))}
                        v5 = (v5 >> ((4294967295) & 63))
                        _ = v158
                    }
                    if false {
                        switch "foo bar" {
                        case "":
                            var v159 uint8 = v1
                            _ = v159
                        case "foo bar":
                            if (32768 != 42) {
                                var v160 uint64 = (v5 & 918)
                                _ = v160
                            }
                        case "hello":
                            fmt.Println(v6)
                        default:
                            var v161 string = ("\n\t" + "hello")
                            _ = v161
                        }
                        for fi162 := 0; fi162 < 20; fi162++ {
                            v1 = (v1 / (((^(v1))) | 1))
                            _ = fi162
                        }
                    } else {
                        v163 := GenS21{F0: 0, F1: (v5 + v5), F2: genFn0(((v5 << ((uint(v1)) & 63)) < (^(v5))), 18446744073709551615)}
                        v17 = v17
                        _ = v163
                    }
                    _ = v150
                }
            }
        }
        for fi164 := 0; fi164 < 19; fi164++ {
            if (false && ((v5 >> ((2) & 63)) <= (v5 - v5))) {
                v6 += (v6 * (float64((0.0 / (0.1 - (0.0 / 1.5)))) - 3.14159))
                fmt.Println(fi164, fi164, v17)
                var v165 int32 = -182
                _ = v165
            }
            var v166 bool = (!true)
            _ = fi164
            _ = v166
        }
        v1 = 180
    }
    var v167 int8 = 11
    var v168 uint = 4294967295
    var v169 int = 42
    fmt.Println("v1:", v1)
    fmt.Println("v167:", v167)
    fmt.Println("v168:", v168)
    fmt.Println("v169:", v169)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v167
    _ = v168
    _ = v169
    _ = v17
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
