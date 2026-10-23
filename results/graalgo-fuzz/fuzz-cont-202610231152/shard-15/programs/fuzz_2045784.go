package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 int; F1 int8; F2 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint64; F1 int8 }
type GenS9 struct { F0 string; F1 string }
type GenS10 struct { F0 int64; F1 uint8; F2 int32 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint64; F1 uint32; F2 uint }
type GenS13 struct { F0 string }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 bool; F1 uint16 }
type GenS17 struct { F0 uint; F1 uint32 }
type GenS18 struct { F0 int; F1 int16 }
type GenS19 struct { F0 uint8; F1 int8 }
type GenS20 struct { F0 int16; F1 string; F2 uint64 }
type GenS21 struct { F0 uint; F1 int32 }
type GenS22 struct { F0 int8 }
type GenS23 struct { F0 int8; F1 int64 }
type GenS24 struct { F0 uint32; F1 uint16 }
type GenS25 struct { F0 bool; F1 int32; F2 int16 }
type GenS26 struct { F0 int32; F1 string; F2 uint }
type GenS27 struct { F0 uint8; F1 int }
type GenS28 struct { F0 int; F1 uint32; F2 int }
type GenS29 struct { F0 int16; F1 uint8 }
type GenS30 struct { F0 string; F1 int; F2 uint64 }
type GenS31 struct { F0 int; F1 int }
func genFn0(p0 string, p1 uint16) bool {
    v7 := GenS1{F0: 4294967295, F1: uint8(p1)}
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := []int{(fi8 >> ((356) & 63)), fi8}
        p1 = (^(p1))
        v10 := GenS2{F0: 115, F1: 1}
        _ = fi8
        _ = v10
        _ = v9
    }
    switch 1 {
    case 32768:
        switch ("foo bar" + ("\n\t" + p0)) {
        case "hello":
            if (127 <= 2) {
                if ((2147483647 < -195) && true) {
                    p0 = p0
                    v11 := GenS1{F0: 65535, F1: uint8(p1)}
                    p1 -= (p1 / ((p1) | 1))
                    _ = v11
                }
            } else {
                var v12 uint16 = 1
                if (false && (42 > 0)) {
                    v12 = v12
                    var v13 uint32 = 0
                    v14 := []int{32768, 4294967295}
                    _ = v13
                    _ = v14
                }
                var v15 int64 = 9223372036854775807
                _ = v12
                _ = v15
            }
        case "foo bar":
            v16 := []int{}
            _ = v16
        case "\n\t":
            for fi17 := 0; fi17 < 11; fi17++ {
                var v18 int64 = -763
                if false {
                    if true {
                        var v19 int = fi17
                        switch p0 {
                        case "hello":
                            if (-800 != (-(v19))) {
                                v20 := GenS3{F0: 0, F1: 1}
                                var v21 bool = true
                                if v21 {
                                    for fi22 := 0; fi22 < 16; fi22++ {
                                        v18 = v18
                                        var v23 uint16 = p1
                                        fmt.Println(v21)
                                        _ = fi22
                                        _ = v23
                                    }
                                    var v24 float32 = 0.0
                                    _ = v24
                                }
                                _ = v20
                                _ = v21
                            }
                        default:
                            var v25 uint64 = 256
                            _ = v25
                        }
                        _ = v19
                    }
                    v26 := []int{-2147483648, (fi17 / (((fi17 % (((fi17 << ((47) & 63))) | 1))) | 1))}
                    _ = v26
                } else {
                    p0 = (((p0 + "foo bar") + (p0 + (("a" + "foo bar") + p0))) + p0)
                    v27 := GenS1{F0: 321, F1: 234}
                    if ((p1 * 65535) <= 2) {
                        switch (fi17 / (((^(fi17))) | 1)) {
                        case 127:
                            var v28 int32 = 1
                            _ = v28
                        default:
                            p0 = ((((p0 + "") + p0) + "\n\t") + p0)
                        }
                        fmt.Println(p1, p0, fi17)
                    } else {
                        if true {
                            p0 = "a"
                        }
                        switch p0 {
                        case "foo bar":
                            if false {
                                if false {
                                    var v29 int64 = v18
                                    v29 += v29
                                    fmt.Println(p1, v29)
                                    _ = v29
                                } else {
                                    p1 = (p1 | 32767)
                                    var v30 float64 = 1.5
                                    _ = v30
                                }
                                var v31 uint8 = 128
                                if false {
                                    for fi32 := 0; fi32 < 4; fi32++ {
                                        if (42 == -2147483648) {
                                            v33 := []int{fi32, 951, (fi32 % (((fi32 & int(v31))) | 1))}
                                            if (-2.25 < (0.0 * 1e10)) {
                                                var v34 uint16 = p1
                                                _ = v34
                                            } else {
                                                if false {
                                                    fmt.Println(p1)
                                                    _ = len(v33)
                                                } else {
                                                    switch fi32 {
                                                    case 32768:
                                                        var v35 uint = uint(p1)
                                                        _ = v35
                                                    case 0:
                                                        if false {
                                                            if (p0 <= p0) {
                                                                v33 = append(v33, (fi32 - fi32))
                                                            } else {
                                                                if false {
                                                                    switch -196 {
                                                                    case 32768:
                                                                        v36 := GenS1{F0: 42, F1: (v31 / ((uint8(p1)) | 1))}
                                                                        _ = v36
                                                                    case 128:
                                                                        p0 = "\n\t"
                                                                    }
                                                                    var v37 uint8 = (^(v31))
                                                                    _ = v37
                                                                } else {
                                                                    var v38 uint64 = 42
                                                                    v39 := GenS2{F0: 2147483647, F1: 601}
                                                                    if true {
                                                                        v18 = (-(v18))
                                                                        fmt.Println(v38, v18)
                                                                        v40 := GenS4{F0: -816, F1: -1, F2: int8(fi32)}
                                                                        _ = v40
                                                                    } else {
                                                                        if (!false) {
                                                                            v41 := GenS5{F0: (p1 / ((478) | 1))}
                                                                            fmt.Println(v38, p1)
                                                                            _ = v41
                                                                        }
                                                                        if true {
                                                                            v18 = v18
                                                                            v33 = append(v33, (fi17 / ((fi17) | 1)))
                                                                            v31 = (v31 - v31)
                                                                        } else {
                                                                            v31 = 67
                                                                        }
                                                                        p0 = p0
                                                                    }
                                                                    _ = v38
                                                                    _ = v39
                                                                }
                                                            }
                                                        } else {
                                                            var v42 int64 = -88
                                                            var v43 uint = 861
                                                            v18 -= 9223372036854775807
                                                            _ = v42
                                                            _ = v43
                                                        }
                                                    default:
                                                        var v44 int8 = int8(v18)
                                                        _ = v44
                                                    }
                                                    if (p1 < 2) {
                                                        if true {
                                                            if len(v33) > 0 {
                                                                v33[0] = fi17
                                                            }
                                                        }
                                                    }
                                                }
                                                var v45 int64 = (v18 >> ((2) & 63))
                                                switch p0 {
                                                case "\n\t":
                                                    if len(v33) >= 2 {
                                                        v33 = v33[1:len(v33)]
                                                    }
                                                }
                                                _ = v45
                                            }
                                            _ = v33
                                        }
                                        if (!((true && false) && true)) {
                                            var v46 uint32 = 660
                                            if true {
                                                v47 := GenS6{F0: v18}
                                                if true {
                                                    v48 := []int{(-(fi32)), fi32}
                                                    v48 = append(v48, 635)
                                                    _ = v48
                                                } else {
                                                    switch p0 {
                                                    case "\n\t":
                                                        v49 := GenS1{F0: 255, F1: v31}
                                                        _ = v49
                                                    case "hello":
                                                        var v50 float64 = -2.25
                                                        _ = v50
                                                    case "foo bar":
                                                        v51 := []int{(-(fi32)), (fi17 << ((255) & 63)), 2147483647, (fi17 * (fi17 - (fi32 - 981)))}
                                                        _ = v51
                                                    default:
                                                        if (false || (2 < 741)) {
                                                            v52 := []int{(-(fi17)), fi17, (fi32 * (^(fi17)))}
                                                            if true {
                                                                v53 := GenS7{F0: (709 != 0)}
                                                                fmt.Println(v18, fi32, v46)
                                                                _ = v53
                                                            }
                                                            v18 = v18
                                                            _ = v52
                                                        }
                                                    }
                                                    v54 := GenS8{F0: 0, F1: 123}
                                                    fmt.Println(v18, fi32, v18)
                                                    _ = v54
                                                }
                                                _ = v47
                                            }
                                            _ = v46
                                        } else {
                                            v55 := GenS9{F0: p0, F1: (("\n\t" + p0) + (("a" + p0) + "hello"))}
                                            fmt.Println(fi17)
                                            p0 = "hello"
                                            var v56 float64 = (-2.25 - float64(0.0))
                                            _ = v55
                                            _ = v56
                                        }
                                        if true {
                                            var v57 uint64 = 0
                                            v58 := GenS10{F0: v18, F1: uint8(v57), F2: -887}
                                            _ = v57
                                            _ = v58
                                        } else {
                                            var v59 float64 = float64(0.0)
                                            _ = v59
                                        }
                                        _ = fi32
                                    }
                                    if (true || (false && true)) {
                                        v60 := GenS8{F0: 18446744073709551615, F1: -128}
                                        v61 := GenS1{F0: 255, F1: v31}
                                        _ = v60
                                        _ = v61
                                    } else {
                                        for fi62 := 0; fi62 < 14; fi62++ {
                                            if false {
                                                v63 := GenS2{F0: 0, F1: 42}
                                                _ = v63
                                            } else {
                                                p0 = "foo bar"
                                            }
                                            v64 := []int{}
                                            if false {
                                                v65 := GenS10{F0: v18, F1: 166, F2: -1}
                                                _ = v65
                                            } else {
                                                if false {
                                                    switch p0 {
                                                    case "":
                                                        v66 := GenS11{F0: 42}
                                                        _ = v66
                                                    }
                                                }
                                            }
                                            _ = fi62
                                            _ = v64
                                        }
                                        var v67 int32 = -305
                                        _ = v67
                                    }
                                } else {
                                    v68 := GenS8{F0: 4294967295, F1: 1}
                                    v69 := GenS3{F0: (^(v31)), F1: 1}
                                    v70 := []int{fi17, (-(fi17))}
                                    _ = v68
                                    _ = v69
                                    _ = v70
                                }
                                _ = v31
                            } else {
                                if true {
                                    if (!false) {
                                        for fi71 := 0; fi71 < 9; fi71++ {
                                            v72 := GenS5{F0: p1}
                                            _ = fi71
                                            _ = v72
                                        }
                                        var v73 int32 = -867
                                        _ = v73
                                    }
                                    var v74 float64 = float64(1.5)
                                    _ = v74
                                }
                                if true {
                                    v75 := GenS7{F0: false}
                                    switch "\n\t" {
                                    case "":
                                        for fi76 := 0; fi76 < 8; fi76++ {
                                            if (!false) {
                                                if (42 > 42) {
                                                    fmt.Println(fi76, fi76)
                                                }
                                                switch (("foo bar" + "hello") + p0) {
                                                case "\n\t":
                                                    var v77 int32 = 42
                                                    _ = v77
                                                case "a":
                                                    v78 := []int{}
                                                    _ = v78
                                                case "foo bar":
                                                    v79 := []int{}
                                                    _ = v79
                                                default:
                                                    switch p0 {
                                                    case "foo bar":
                                                        var v80 string = "hello"
                                                        _ = v80
                                                    case "\n\t":
                                                        v81 := GenS2{F0: 2147483648, F1: 667}
                                                        _ = v81
                                                    default:
                                                        switch (p0 + "\n\t") {
                                                        case "foo bar":
                                                            var v82 int32 = 560
                                                            _ = v82
                                                        case "hello":
                                                            p0 = "\n\t"
                                                            v83 := []int{(fi76 ^ fi17), fi17}
                                                            if (!true) {
                                                                switch fi17 {
                                                                case 0:
                                                                    v84 := GenS12{F0: 4294967295, F1: 261, F2: 0}
                                                                    _ = v84
                                                                case -2147483648:
                                                                    if (((v18 - v18) <= v18) || (!(p0 == p0))) {
                                                                        if false {
                                                                            var v85 float32 = 3.14159
                                                                            var v86 uint64 = 450
                                                                            v83 = append(v83, (fi17 + fi76))
                                                                            _ = v85
                                                                            _ = v86
                                                                        }
                                                                        if false {
                                                                            _ = len(v83)
                                                                            v87 := GenS3{F0: 42, F1: 1}
                                                                            _ = v87
                                                                        } else {
                                                                            if len(v83) >= 2 {
                                                                                v83 = v83[1:len(v83)]
                                                                            }
                                                                            if true {
                                                                                var v88 int = (-(fi76))
                                                                                _ = v88
                                                                            } else {
                                                                                v89 := GenS4{F0: fi17, F1: 0, F2: 0}
                                                                                var v90 int64 = (v18 * v18)
                                                                                v91 := GenS2{F0: 2, F1: 32767}
                                                                                _ = v89
                                                                                _ = v90
                                                                                _ = v91
                                                                            }
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                if false {
                                                                                    var v92 string = p0
                                                                                    v92 = v92
                                                                                    _ = v92
                                                                                }
                                                                            }
                                                                        }
                                                                        var v93 int16 = 42
                                                                        _ = v93
                                                                    }
                                                                default:
                                                                    var v94 uint = 417
                                                                    _ = v94
                                                                }
                                                                p1 = p1
                                                                if len(v83) >= 2 {
                                                                    v83 = v83[1:len(v83)]
                                                                }
                                                            }
                                                            v95 := GenS13{F0: (p0 + "a")}
                                                            p0 = ((((p0 + (p0 + p0)) + ((p0 + (p0 + "")) + "a")) + "a") + "\n\t")
                                                            _ = v83
                                                            _ = v95
                                                        case "":
                                                            v96 := GenS4{F0: fi76, F1: 100, F2: -128}
                                                            _ = v96
                                                        default:
                                                            if true {
                                                                v18 = 42
                                                                v97 := GenS1{F0: 18446744073709551615, F1: 133}
                                                                var v98 int = (fi76 / ((fi76) | 1))
                                                                _ = v97
                                                                _ = v98
                                                            } else {
                                                                var v99 int64 = v18
                                                                switch (fi76 ^ (fi76 - (fi17 >> ((1) & 63)))) {
                                                                case -755:
                                                                    v99 = (-(v18))
                                                                }
                                                                _ = v99
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                            v100 := GenS10{F0: -9223372036854775808, F1: 1, F2: 1}
                                            _ = fi76
                                            _ = v100
                                        }
                                    case "foo bar":
                                        var v101 uint8 = 255
                                        _ = v101
                                    case "a":
                                        p1 -= 810
                                    }
                                    _ = v75
                                }
                            }
                        case "\n\t":
                            p1 += p1
                        case "a":
                            for fi102 := 0; fi102 < 6; fi102++ {
                                if (1.5 >= 0.0) {
                                    if false {
                                        p0 = (p0 + "foo bar")
                                        var v103 uint = 4294967295
                                        v103 = v103
                                        _ = v103
                                    } else {
                                        var v104 string = ("foo bar" + p0)
                                        v105 := GenS2{F0: 150, F1: -1}
                                        v106 := []int{int(p1)}
                                        v104 = p0
                                        _ = v104
                                        _ = v105
                                        _ = v106
                                    }
                                    var v107 uint8 = 136
                                    v108 := GenS14{F0: true}
                                    _ = v107
                                    _ = v108
                                }
                                fmt.Println(fi102, fi17, p1)
                                var v109 int64 = v18
                                if true {
                                    switch (p0 + ("\n\t" + "hello")) {
                                    case "":
                                        var v110 uint = 844
                                        var v111 uint64 = 18446744073709551615
                                        if false {
                                            if false {
                                                v112 := []int{fi17, fi17}
                                                var v113 int32 = 30
                                                _ = v112
                                                _ = v113
                                            } else {
                                                if (true && (p0 <= p0)) {
                                                    v114 := []int{(-(fi17)), (fi17 * (-(fi102))), fi17}
                                                    if false {
                                                        fmt.Println(v109, v110)
                                                    } else {
                                                        v115 := GenS15{F0: -2147483648}
                                                        _ = v115
                                                    }
                                                    var v116 int32 = 2147483647
                                                    var v117 uint64 = uint64(v110)
                                                    if true {
                                                        if true {
                                                            if true {
                                                                v118 := GenS16{F0: (false && false), F1: (p1 % ((p1) | 1))}
                                                                p1 = uint16(v117)
                                                                _ = v118
                                                            } else {
                                                                if (0.0 >= 0.0) {
                                                                    p1 += p1
                                                                    fmt.Println(fi17, p1, fi17)
                                                                } else {
                                                                    if len(v114) > 0 {
                                                                        v114[0] = fi102
                                                                    }
                                                                    v119 := GenS17{F0: (v110 << (((v110 + v110)) & 63)), F1: 42}
                                                                    _ = v119
                                                                }
                                                            }
                                                            v120 := GenS18{F0: fi17, F1: -989}
                                                            if true {
                                                                var v121 uint = (^(v110))
                                                                v122 := GenS7{F0: (1 >= 0)}
                                                                var v123 float64 = ((-2.25 + (-2.25 - 0.1)) - 1.5)
                                                                _ = v121
                                                                _ = v122
                                                                _ = v123
                                                            }
                                                            _ = v120
                                                        } else {
                                                            switch fi102 {
                                                            case -999:
                                                                switch p0 {
                                                                case "foo bar":
                                                                    var v124 uint64 = (v111 * (v111 * 18446744073709551615))
                                                                    _ = v124
                                                                }
                                                            case -165:
                                                                v114 = append(v114, 127)
                                                            }
                                                            if (-2.25 >= 3.14159) {
                                                                fmt.Println(fi102, p1, fi17)
                                                                v117 = uint64(v110)
                                                                if (9223372036854775807 >= int64(v117)) {
                                                                    if len(v114) > 0 {
                                                                        v114[0] = fi17
                                                                    }
                                                                    fmt.Println(p1)
                                                                    v109 -= v18
                                                                }
                                                            } else {
                                                                v125 := GenS19{F0: 113, F1: 101}
                                                                if false {
                                                                    var v126 int = fi17
                                                                    var v127 int32 = v116
                                                                    var v128 string = p0
                                                                    _ = v126
                                                                    _ = v127
                                                                    _ = v128
                                                                }
                                                                _ = v125
                                                            }
                                                        }
                                                    } else {
                                                        var v129 int8 = int8(v109)
                                                        if false {
                                                            if true {
                                                                v109 -= (-(v18))
                                                                v130 := GenS14{F0: (!true)}
                                                                var v131 int8 = -76
                                                                _ = v130
                                                                _ = v131
                                                            } else {
                                                                if true {
                                                                    v18 = v18
                                                                    _ = len(v114)
                                                                    v129 += (v129 % ((1) | 1))
                                                                }
                                                            }
                                                            v132 := GenS16{F0: (40 <= (v129 | (v129 % (((v129 >> (((^(v110))) & 63))) | 1)))), F1: 255}
                                                            _ = v132
                                                        } else {
                                                            var v133 uint8 = 127
                                                            p1 = (p1 & p1)
                                                            _ = v133
                                                        }
                                                        _ = v129
                                                    }
                                                    _ = v114
                                                    _ = v116
                                                    _ = v117
                                                }
                                            }
                                            v134 := GenS20{F0: 1000, F1: p0, F2: 42}
                                            var v135 uint64 = (v111 + 1)
                                            _ = v134
                                            _ = v135
                                        } else {
                                            switch 32768 {
                                            case 1:
                                                var v136 uint = 255
                                                _ = v136
                                            case -103:
                                                fmt.Println(v18, p0)
                                            case 127:
                                                v18 = v18
                                            }
                                            var v137 float32 = -2.25
                                            _ = v137
                                        }
                                        var v138 int8 = -1
                                        var v139 uint8 = 1
                                        _ = v110
                                        _ = v111
                                        _ = v138
                                        _ = v139
                                    default:
                                        var v140 string = (p0 + p0)
                                        _ = v140
                                    }
                                    var v141 bool = false
                                    switch (fi17 % ((-2147483648) | 1)) {
                                    case -2147483648:
                                        v18 -= (v18 & 0)
                                    case -268:
                                        fmt.Println(p0, p0, v141)
                                    }
                                    _ = v141
                                } else {
                                    switch p0 {
                                    case "\n\t":
                                        v142 := []int{fi17, 836, fi102, fi17}
                                        _ = v142
                                    case "hello":
                                        v143 := []int{(fi17 | fi102), fi17, (^(fi102))}
                                        _ = v143
                                    case "foo bar":
                                        switch "" {
                                        case "":
                                            if true {
                                                v144 := GenS21{F0: 952, F1: -1}
                                                v145 := []int{fi102, (fi17 % ((fi102) | 1)), (fi102 % ((-59) | 1)), (fi17 >> ((1) & 63))}
                                                _ = v144
                                                _ = v145
                                            }
                                        case "foo bar":
                                            switch (^(fi17)) {
                                            case 275:
                                                switch fi102 {
                                                case 899:
                                                    v146 := GenS21{F0: 65535, F1: 0}
                                                    _ = v146
                                                case -9223372036854775808:
                                                    var v147 int8 = 5
                                                    _ = v147
                                                default:
                                                    fmt.Println(p1)
                                                }
                                            case 806:
                                                var v148 float32 = (1e10 + (1.5 - (0.1 - (0.0 + (0.1 / (((float32(float64(3.14159)) - 0.0) / ((0.0 - float32((0.1 + (-2.25 - (-2.25 / 1.0))))) / (1.5 * (3.14159 / 1.0)))) + 1e10))))))
                                                _ = v148
                                            case -9223372036854775808:
                                                if true {
                                                    v149 := GenS21{F0: 0, F1: 2147483647}
                                                    var v150 uint64 = 42
                                                    _ = v149
                                                    _ = v150
                                                }
                                                var v151 uint32 = 384
                                                _ = v151
                                            default:
                                                switch fi17 {
                                                case 4294967295:
                                                    if false {
                                                        var v152 uint8 = 42
                                                        _ = v152
                                                    } else {
                                                        p1 += p1
                                                    }
                                                case 128:
                                                    v153 := []int{-9223372036854775808, fi102, fi17}
                                                    _ = v153
                                                case 2147483647:
                                                    var v154 int = fi17
                                                    _ = v154
                                                default:
                                                    if (-1 > -128) {
                                                        if true {
                                                            v109 -= v18
                                                            var v155 uint8 = 123
                                                            fmt.Println(p1, fi17)
                                                            _ = v155
                                                        }
                                                    }
                                                }
                                            }
                                        default:
                                            v156 := []int{}
                                            _ = v156
                                        }
                                    }
                                    switch fi17 {
                                    case 127:
                                        if ((false || true) && (p1 == (p1 >> ((255) & 63)))) {
                                            v157 := GenS22{F0: -1}
                                            if true {
                                                p1 = p1
                                            }
                                            if ((v18 >> ((587) & 63)) > (^(v18))) {
                                                fmt.Println(fi17, v18, fi17)
                                                var v158 string = (p0 + "foo bar")
                                                v159 := GenS23{F0: 100, F1: (v18 * (v18 - v109))}
                                                _ = v158
                                                _ = v159
                                            } else {
                                                v160 := []int{(fi102 - 901), (fi17 >> ((905) & 63))}
                                                v161 := GenS17{F0: 65535, F1: 65535}
                                                _ = v160
                                                _ = v161
                                            }
                                            switch p0 {
                                            case "foo bar":
                                                if true {
                                                    var v162 int16 = 0
                                                    _ = v162
                                                }
                                            case "":
                                                v163 := []int{948, fi17, (fi102 / ((fi102) | 1)), fi102}
                                                _ = v163
                                            default:
                                                v164 := []int{}
                                                _ = v164
                                            }
                                            v18 += (v18 + v18)
                                            _ = v157
                                        }
                                        var v165 int = int(v18)
                                        p1 -= 974
                                        _ = v165
                                    case -2147483648:
                                        fmt.Println(p0)
                                    }
                                }
                                _ = fi102
                                _ = v109
                            }
                        }
                    }
                    _ = v27
                }
                var v166 uint32 = 0
                _ = fi17
                _ = v166
                _ = v18
            }
        }
    }
    for fi167 := 0; fi167 < 5; fi167++ {
        p1 -= 2
        _ = fi167
    }
    _ = p0
    _ = p1
    _ = v7
    return (106 != 101)
}
func genFn1() int {
    switch "\n\t" {
    case "\n\t":
        var v168 bool = (!false)
        _ = v168
    case "hello":
        var v169 float64 = 1e10
        _ = v169
    case "":
        var v170 float64 = 0.1
        _ = v170
    default:
        v171 := []int{989}
        _ = v171
    }
    if (0.1 == (-2.25 / 0.1)) {
        for fi172 := 0; fi172 < 13; fi172++ {
            var v173 int16 = 977
            for fi174 := 0; fi174 < 12; fi174++ {
                v175 := []int{}
                _ = fi174
                _ = v175
            }
            _ = fi172
            _ = v173
        }
    } else {
        var v176 uint64 = 0
        _ = v176
    }
    v177 := []int{32767, 32768, 127, 454}
    v177 = append(v177, 9223372036854775807)
    var v178 uint = 157
    _ = v177
    _ = v178
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = ("" + "hello")
    v3 := make(map[string]string)
    var v4 int32 = 1
    var v5 bool = (1.5 >= -2.25)
    var v6 uint8 = 1
    v5 = v5
    v179 := v3[v2]
    v180 := []int{32767, -995, -2147483648}
    for fi181 := 0; fi181 < 8; fi181++ {
        switch fi181 {
        case -369:
            var v182 string = "hello"
            _ = v182
        }
        _ = fi181
    }
    if (255 >= v6) {
        for fi183 := 0; fi183 < 1; fi183++ {
            var v184 uint8 = v6
            switch ("a" + v2) {
            case "\n\t":
                if (int16(v6) <= 42) {
                    var v185 uint = 140
                    var v186 uint32 = 1
                    if (!(42 != uint8(v185))) {
                        v6 = 127
                        var v187 int32 = (v4 | -2147483648)
                        v2 = (v179 + v179)
                        _ = v187
                    }
                    _ = v185
                    _ = v186
                }
            case "a":
                if (983 > fi183) {
                    v188 := GenS24{F0: 2, F1: 667}
                    var v189 float32 = float32(1e10)
                    _ = v188
                    _ = v189
                } else {
                    v4 = (v4 + v4)
                    var v190 int = (fi183 - (fi183 - fi183))
                    _ = v190
                }
            }
            if ((v5 || v5) && v5) {
                var v191 string = "foo bar"
                v192 := GenS14{F0: (v6 >= v6)}
                fmt.Println(v179)
                _ = v191
                _ = v192
            } else {
                switch ("hello" + "a") {
                case "a":
                    var v193 bool = v5
                    _ = v193
                case "foo bar":
                    if len(v180) > 0 {
                        v180[0] = fi183
                    }
                default:
                    v179 = "\n\t"
                }
            }
            _ = fi183
            _ = v184
        }
        for fi194 := 0; fi194 < 3; fi194++ {
            switch fi194 {
            case -3:
                switch v179 {
                case "a":
                    if v5 {
                        v2 = ("a" + v179)
                        v195 := GenS5{F0: 42}
                        v196 := GenS24{F0: 11, F1: 780}
                        _ = v195
                        _ = v196
                    }
                case "\n\t":
                    fmt.Println(v2, v6)
                }
            case 32768:
                var v197 bool = (!(42 < 1))
                _ = v197
            case 355:
                v198 := GenS25{F0: v5, F1: (v4 << ((556) & 63)), F2: -1}
                _ = v198
            default:
                switch (-(fi194)) {
                case 42:
                    var v199 bool = (false && (127 <= v6))
                    _ = v199
                default:
                    if ("\n\t" == "a") {
                        v200 := GenS8{F0: 18446744073709551615, F1: 1}
                        for fi201 := 0; fi201 < 8; fi201++ {
                            if (0 != -1) {
                                v2 = "a"
                                var v202 float64 = -2.25
                                _ = v202
                            }
                            _ = fi201
                        }
                        _ = v200
                    }
                }
            }
            _ = fi194
        }
        for fi203 := 0; fi203 < 16; fi203++ {
            var v204 string = v2
            var v205 bool = true
            _ = fi203
            _ = v204
            _ = v205
        }
    } else {
        if (1 < 30) {
            v180 = append(v180, 1)
            if false {
                var v206 uint16 = 1
                for fi207 := 0; fi207 < 10; fi207++ {
                    var v208 int32 = 1
                    _ = fi207
                    _ = v208
                }
                v209 := GenS26{F0: -1, F1: "a", F2: 203}
                _ = v206
                _ = v209
            }
        } else {
            _ = len(v1)
            var v210 int16 = 32767
            _ = v210
        }
        if len(v180) > 0 {
            v180[0] = -2147483648
        }
    }
    if (734 != 0) {
        var v211 float64 = 0.1
        for fi212 := 0; fi212 < 2; fi212++ {
            var v213 uint16 = 2
            for fi214 := 0; fi214 < 1; fi214++ {
                var v215 bool = (3.14159 >= 1e10)
                v5 = ((350 < 32767) || v5)
                v215 = v5
                _ = fi214
                _ = v215
            }
            v6 = v6
            _ = fi212
            _ = v213
        }
        _ = v211
    } else {
        v5 = true
        if false {
            var v216 int64 = -828
            _ = v216
        } else {
            v179 = ((v2 + "") + (("\n\t" + "") + "hello"))
            for fi217 := 0; fi217 < 16; fi217++ {
                v4 -= (v4 >> ((476) & 63))
                for fi218 := 0; fi218 < 17; fi218++ {
                    v179 = v179
                    if (!("a" >= v2)) {
                        _ = len(v3)
                    } else {
                        v4 += v4
                    }
                    _ = fi218
                }
                var v219 int8 = 42
                _ = fi217
                _ = v219
            }
            v220 := GenS16{F0: ((1 < 1) && (0 > 42)), F1: 32767}
            _ = v220
        }
    }
    for fi221 := 0; fi221 < 16; fi221++ {
        var v222 int32 = (v4 + 759)
        if (42 >= 2) {
            var v223 bool = ((v222 >= (v222 % ((v222) | 1))) || (!(!v5)))
            var v224 uint = 677
            for fi225 := 0; fi225 < 15; fi225++ {
                v180 = append(v180, genFn1())
                _ = fi225
            }
            _ = v223
            _ = v224
        } else {
            v226 := v1[fi221]
            switch genFn1() {
            case 9223372036854775807:
                if len(v180) > 0 {
                    v180[0] = fi221
                }
            default:
                switch (fi221 ^ fi221) {
                case 9223372036854775807:
                    var v227 int32 = 2147483647
                    _ = v227
                default:
                    for fi228 := 0; fi228 < 2; fi228++ {
                        switch int(v222) {
                        case -1:
                            if (v226 && v226) {
                                if v5 {
                                    if v5 {
                                        v179 = v179
                                        v5 = ((((v6 / ((128) | 1)) == v6) && (!(!v5))) || (!((v6 | (v6 & 1)) == v6)))
                                    }
                                    if (98 > 42) {
                                        var v229 int16 = -1
                                        if false {
                                            v230 := GenS19{F0: (v6 | (v6 & (^(v6)))), F1: 1}
                                            _ = v230
                                        } else {
                                            v231 := GenS27{F0: (v6 - v6), F1: int(v6)}
                                            _ = v231
                                        }
                                        v179 = ((v2 + ((v179 + v179) + (v2 + v2))) + (v179 + ("foo bar" + v179)))
                                        _ = v229
                                    } else {
                                        switch (v179 + "\n\t") {
                                        case "foo bar":
                                            if v5 {
                                                var v232 uint16 = 42
                                                var v233 bool = v5
                                                var v234 int16 = -479
                                                _ = v232
                                                _ = v233
                                                _ = v234
                                            }
                                        case "":
                                            var v235 uint8 = v6
                                            _ = v235
                                        default:
                                            switch "\n\t" {
                                            case "hello":
                                                switch "foo bar" {
                                                case "a":
                                                    v180 = append(v180, (fi221 | (fi228 - (fi228 % ((genFn1()) | 1)))))
                                                default:
                                                    v236 := v1[(fi221 % ((fi221) | 1))]
                                                    _ = v236
                                                }
                                            case "foo bar":
                                                delete(v3, (("\n\t" + v179) + v2))
                                            case "\n\t":
                                                v226 = v5
                                            }
                                        }
                                        v2 = "\n\t"
                                        if v226 {
                                            if v5 {
                                                fmt.Println(v179)
                                            }
                                            fmt.Println(fi221, v4)
                                            if v5 {
                                                var v237 bool = v226
                                                if v226 {
                                                    if len(v180) > 0 {
                                                        v180[0] = fi221
                                                    }
                                                    v237 = true
                                                    if (v6 >= 255) {
                                                        var v238 float32 = (float32((float64(-2.25) / 1.5)) / float32(0.1))
                                                        v2 = (v179 + v179)
                                                        v239 := GenS9{F0: v2, F1: (("\n\t" + "a") + ("" + ""))}
                                                        _ = v238
                                                        _ = v239
                                                    } else {
                                                        var v240 bool = (1e10 >= (1e10 / 3.14159))
                                                        var v241 float64 = 1.5
                                                        _ = v240
                                                        _ = v241
                                                    }
                                                }
                                                _ = v237
                                            }
                                        } else {
                                            if (!v5) {
                                                switch "foo bar" {
                                                case "a":
                                                    var v242 int64 = -388
                                                    _ = v242
                                                default:
                                                    fmt.Println(fi221, v6, fi228)
                                                }
                                                switch 42 {
                                                case 253:
                                                    var v243 uint = 2
                                                    _ = v243
                                                case 32767:
                                                    v244 := GenS25{F0: v5, F1: int32(v6), F2: 760}
                                                    _ = v244
                                                case -2147483648:
                                                    fmt.Println(fi221, fi228, v6)
                                                }
                                                if (!v5) {
                                                    var v245 float64 = (0.1 * (-2.25 * float64(0.0)))
                                                    v5 = (fi221 > -192)
                                                    _ = v245
                                                } else {
                                                    v3[("a" + "foo bar")] = v179
                                                    var v246 uint16 = 32767
                                                    v4 = -655
                                                    _ = v246
                                                }
                                            }
                                            switch "" {
                                            case "":
                                                delete(v1, (^(fi221)))
                                            case "a":
                                                switch genFn1() {
                                                case -1:
                                                    switch 1 {
                                                    case 9223372036854775807:
                                                        if ((v6 ^ v6) <= v6) {
                                                            v2 = ""
                                                            v247 := GenS28{F0: fi228, F1: 767, F2: genFn1()}
                                                            var v248 int8 = 127
                                                            _ = v247
                                                            _ = v248
                                                        }
                                                    case -670:
                                                        switch v179 {
                                                        case "hello":
                                                            v179 = ("foo bar" + v179)
                                                        case "":
                                                            switch 128 {
                                                            case -202:
                                                                v2 = (v179 + v179)
                                                            case 32767:
                                                                var v249 int16 = 1000
                                                                _ = v249
                                                            }
                                                        default:
                                                            v250 := GenS2{F0: 4294967295, F1: 51}
                                                            _ = v250
                                                        }
                                                    }
                                                default:
                                                    if (!true) {
                                                        switch genFn1() {
                                                        case 128:
                                                            var v251 uint8 = 203
                                                            _ = v251
                                                        case 1:
                                                            fmt.Println(v5)
                                                        default:
                                                            var v252 uint = 4294967295
                                                            _ = v252
                                                        }
                                                        if (0 == 42) {
                                                            v179 = ((("a" + v179) + (v2 + "\n\t")) + v179)
                                                            var v253 float64 = -2.25
                                                            var v254 uint16 = 32767
                                                            _ = v253
                                                            _ = v254
                                                        }
                                                    } else {
                                                        v255 := GenS29{F0: -34, F1: v6}
                                                        _ = v255
                                                    }
                                                }
                                            }
                                            v256 := GenS30{F0: "", F1: fi228, F2: uint64(v6)}
                                            _ = v256
                                        }
                                        if (0 <= (v6 ^ v6)) {
                                            v226 = (false || ((1e10 / -2.25) != 0.1))
                                        }
                                        v257 := GenS31{F0: (^(fi221)), F1: fi228}
                                        _ = v257
                                    }
                                    delete(v3, v2)
                                } else {
                                    var v258 uint64 = 2
                                    v259 := GenS1{F0: 255, F1: v6}
                                    switch (fi221 ^ 156) {
                                    case -9223372036854775808:
                                        v5 = v226
                                    case 4294967295:
                                        fmt.Println(fi221)
                                    case 1:
                                        if len(v180) >= 2 {
                                            v180 = v180[1:len(v180)]
                                        }
                                    }
                                    _ = v258
                                    _ = v259
                                }
                            }
                        case 9223372036854775807:
                            var v260 string = (v2 + "a")
                            _ = v260
                        default:
                            switch genFn1() {
                            case -195:
                                v226 = v5
                            case 79:
                                switch (-(fi221)) {
                                case -926:
                                    if len(v180) > 0 {
                                        v180[0] = fi228
                                    }
                                case -880:
                                    switch fi221 {
                                    case -959:
                                        v2 = v2
                                    default:
                                        v180 = append(v180, int(v6))
                                    }
                                }
                            case 1:
                                if (v5 || v5) {
                                    v261 := v3["hello"]
                                    switch (fi228 * (fi228 + fi228)) {
                                    case 41:
                                        delete(v1, (fi228 & fi228))
                                    case 1:
                                        var v262 float32 = 0.0
                                        _ = v262
                                    case 128:
                                        if (-128 != 0) {
                                            var v263 float32 = 3.14159
                                            if (!(-2147483648 < (v4 + (v222 % ((1) | 1))))) {
                                                if (false || v226) {
                                                    v180 = append(v180, int(v6))
                                                    var v264 int32 = -255
                                                    _ = v264
                                                } else {
                                                    switch genFn1() {
                                                    case 2147483647:
                                                        v265 := v3[("" + (v179 + v2))]
                                                        _ = v265
                                                    case 42:
                                                        v222 += (v222 << ((194) & 63))
                                                    default:
                                                        var v266 int8 = 100
                                                        _ = v266
                                                    }
                                                    var v267 float64 = 0.1
                                                    v4 = (v4 << ((65535) & 63))
                                                    _ = v267
                                                }
                                                if (v226 || v5) {
                                                    var v268 uint = 65535
                                                    _ = v268
                                                }
                                            } else {
                                                switch ((v261 + ("foo bar" + v179)) + "a") {
                                                case "foo bar":
                                                    v261 = ""
                                                case "":
                                                    var v269 float64 = (-2.25 / 3.14159)
                                                    _ = v269
                                                case "hello":
                                                    if (!((genFn1() < 922) || (!((18446744073709551615 != 2) && (!v226))))) {
                                                        var v270 int16 = 1000
                                                        v271 := GenS15{F0: (v222 | (-(v222)))}
                                                        _ = v270
                                                        _ = v271
                                                    }
                                                }
                                                v272 := GenS8{F0: 18446744073709551615, F1: 1}
                                                switch v261 {
                                                case "\n\t":
                                                    fmt.Println(v263, v226, v2)
                                                }
                                                _ = v272
                                            }
                                            _ = v263
                                        } else {
                                            var v273 float64 = float64(float32(-2.25))
                                            _ = v273
                                        }
                                    }
                                    _ = v261
                                }
                                fmt.Println(v5)
                                var v274 uint16 = 32767
                                _ = v274
                            }
                        }
                        _ = fi228
                    }
                }
            }
            fmt.Println(fi221, v5)
            _ = v226
        }
        switch v2 {
        case "":
            var v275 int = fi221
            _ = v275
        }
        _ = fi221
        _ = v222
    }
    v276 := GenS24{F0: 2147483647, F1: 32767}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v179:", v179)
    fmt.Println("slicelen v180:", len(v180), v180)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v179
    _ = v180
    _ = v2
    _ = v276
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
