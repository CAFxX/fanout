package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 uint; F2 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int16; F1 uint64; F2 int8 }
type GenS8 struct { F0 int64; F1 bool }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int8; F1 int64; F2 uint64 }
type GenS11 struct { F0 int64; F1 int8; F2 int16 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint8; F1 int8; F2 uint }
type GenS14 struct { F0 int16; F1 uint8 }
type GenS15 struct { F0 int32; F1 string; F2 uint8 }
type GenS16 struct { F0 uint; F1 uint32 }
type GenS17 struct { F0 uint8; F1 uint8 }
func genFn0(p0 string, p1 int16) uint8 {
    var v5 float64 = -2.25
    var v6 int32 = -2147483648
    var v7 bool = (!false)
    if false {
        var v8 uint8 = 1
        if (42 <= (p1 - p1)) {
            v9 := []int{127, 9223372036854775807, -237, -9223372036854775808}
            var v10 int64 = 1
            if (v7 || (!(true && true))) {
                var v11 uint32 = 0
                v12 := GenS1{F0: 18446744073709551615, F1: 42, F2: (p0 + p0)}
                _ = v11
                _ = v12
            } else {
                switch 1 {
                case 2147483647:
                    _ = len(v9)
                case -748:
                    _ = len(v9)
                default:
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                }
                if false {
                    for fi13 := 0; fi13 < 11; fi13++ {
                        for fi14 := 0; fi14 < 11; fi14++ {
                            v10 = (-(v10))
                            if (!((v6 | -1) <= v6)) {
                                if v7 {
                                    if false {
                                        var v15 uint8 = v8
                                        var v16 int16 = p1
                                        var v17 uint = 98
                                        if v7 {
                                            v7 = false
                                            var v18 bool = false
                                            _ = v18
                                        } else {
                                            v19 := GenS1{F0: 1, F1: 99, F2: (p0 + (("" + (p0 + ("a" + (p0 + ("foo bar" + (("foo bar" + p0) + "\n\t")))))) + (p0 + p0)))}
                                            var v20 int16 = (p1 | (^(p1)))
                                            var v21 bool = v7
                                            _ = v19
                                            _ = v20
                                            _ = v21
                                        }
                                        _ = v15
                                        _ = v16
                                        _ = v17
                                    } else {
                                        var v22 uint64 = 1
                                        v23 := GenS1{F0: (v22 & v22), F1: (^(v22)), F2: "foo bar"}
                                        _ = v22
                                        _ = v23
                                    }
                                    v5 = (float64((float32(0.0) * 3.14159)) / (0.0 + (v5 + v5)))
                                } else {
                                    var v24 int8 = -60
                                    _ = v24
                                }
                                switch 2147483647 {
                                case 127:
                                    var v25 int16 = 0
                                    p1 -= p1
                                    p0 = (("a" + (p0 + p0)) + p0)
                                    if v7 {
                                        v5 = 0.0
                                        if (!true) {
                                            v26 := GenS1{F0: 2, F1: 0, F2: "\n\t"}
                                            v6 = 1
                                            var v27 bool = v7
                                            switch 2147483647 {
                                            case -9223372036854775808:
                                                v7 = v27
                                            case -961:
                                                var v28 float64 = v5
                                                _ = v28
                                            }
                                            v7 = v7
                                            var v29 float64 = v5
                                            if (!("foo bar" <= "a")) {
                                                if (32767 < (fi14 ^ (fi13 | (^(fi13))))) {
                                                    var v30 int8 = int8(v8)
                                                    var v31 uint = 1
                                                    _ = v30
                                                    _ = v31
                                                }
                                                v8 = 42
                                                switch (fi14 / ((fi14) | 1)) {
                                                case 1:
                                                    var v32 uint8 = 75
                                                    v9 = append(v9, fi14)
                                                    _ = v32
                                                case 0:
                                                    v33 := GenS2{F0: false}
                                                    _ = v33
                                                case 9223372036854775807:
                                                    v34 := GenS3{F0: "hello", F1: 1, F2: (v8 - (v8 & v8))}
                                                    _ = v34
                                                }
                                            } else {
                                                var v35 int8 = -1
                                                _ = v35
                                            }
                                            _ = v26
                                            _ = v27
                                            _ = v29
                                        } else {
                                            var v36 uint16 = 651
                                            fmt.Println(v25, v6)
                                            v7 = ((v5 * 3.14159) <= v5)
                                            _ = v36
                                        }
                                        var v37 int8 = 100
                                        _ = v37
                                    } else {
                                        _ = len(v9)
                                        switch p0 {
                                        case "a":
                                            var v38 float32 = 3.14159
                                            _ = v38
                                        case "hello":
                                            switch fi13 {
                                            case -317:
                                                switch (-(fi14)) {
                                                case 0:
                                                    v39 := GenS4{F0: true}
                                                    _ = v39
                                                case 32767:
                                                    p0 = p0
                                                }
                                            case 2147483647:
                                                v6 = (v6 >> ((65535) & 63))
                                            case 127:
                                                _ = len(v9)
                                            }
                                        }
                                    }
                                    var v40 int32 = (v6 % (((v6 | 42)) | 1))
                                    _ = v25
                                    _ = v40
                                case -1:
                                    v6 -= (v6 & -490)
                                default:
                                    var v41 int16 = p1
                                    _ = v41
                                }
                                v7 = v7
                                if (795 < 433) {
                                    var v42 uint32 = 65535
                                    _ = v42
                                } else {
                                    if ("\n\t" <= "hello") {
                                        _ = len(v9)
                                    } else {
                                        var v43 string = ("\n\t" + "foo bar")
                                        _ = v43
                                    }
                                    v5 = (v5 - -2.25)
                                }
                                var v44 int64 = (v10 / ((v10) | 1))
                                _ = v44
                            } else {
                                if len(v9) > 0 {
                                    v9[0] = (fi14 / ((fi13) | 1))
                                }
                                if ((v5 - (v5 - v5)) != float64(float32((3.14159 / v5)))) {
                                    var v45 uint = 1
                                    switch (fi14 ^ fi13) {
                                    case 9223372036854775807:
                                        v5 -= 0.0
                                    case -2147483648:
                                        if (!(!(p0 <= ((p0 + "hello") + (p0 + p0))))) {
                                            var v46 float64 = v5
                                            if (791 == 4294967295) {
                                                if len(v9) > 0 {
                                                    v9[0] = fi14
                                                }
                                                if v7 {
                                                    v47 := GenS1{F0: 2, F1: 18446744073709551615, F2: "a"}
                                                    v9 = append(v9, fi13)
                                                    if len(v9) >= 2 {
                                                        v9 = v9[1:len(v9)]
                                                    }
                                                    _ = v47
                                                }
                                                fmt.Println(v7, p0, v8)
                                                v6 = v6
                                            }
                                            v46 -= (v46 + (0.0 / v5))
                                            v48 := GenS1{F0: 183, F1: 4294967295, F2: ("\n\t" + (p0 + p0))}
                                            _ = v46
                                            _ = v48
                                        }
                                    case 866:
                                        fmt.Println(v45, v10, p0)
                                    }
                                    v49 := GenS5{F0: 32767, F1: (v10 << ((v45) & 63))}
                                    p0 = (p0 + p0)
                                    v50 := GenS2{F0: v7}
                                    fmt.Println(p1, v5, v45)
                                    v51 := GenS6{F0: -1}
                                    _ = v45
                                    _ = v49
                                    _ = v50
                                    _ = v51
                                }
                                var v52 float64 = 3.14159
                                _ = v52
                            }
                            _ = fi14
                        }
                        _ = len(v9)
                        _ = fi13
                    }
                    for fi53 := 0; fi53 < 18; fi53++ {
                        v54 := GenS3{F0: (p0 + (p0 + "foo bar")), F1: 65535, F2: v8}
                        var v55 uint16 = 0
                        if v7 {
                            for fi56 := 0; fi56 < 16; fi56++ {
                                var v57 uint16 = 42
                                v55 += 2
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                                _ = fi56
                                _ = v57
                            }
                            v7 = ((v8 >= (v8 - (v8 ^ v8))) || (v8 > (v8 ^ (v8 ^ 128))))
                            if v7 {
                                switch fi53 {
                                case -1:
                                    var v58 float64 = v5
                                    _ = v58
                                case 2147483647:
                                    v7 = v7
                                }
                            }
                        }
                        _ = fi53
                        _ = v54
                        _ = v55
                    }
                    if true {
                        for fi59 := 0; fi59 < 16; fi59++ {
                            v60 := GenS5{F0: (p1 % (((p1 << ((2) & 63))) | 1)), F1: (v10 % ((v10) | 1))}
                            _ = fi59
                            _ = v60
                        }
                    } else {
                        for fi61 := 0; fi61 < 13; fi61++ {
                            if (v7 || true) {
                                p1 += (-(p1))
                                for fi62 := 0; fi62 < 1; fi62++ {
                                    fmt.Println(p1)
                                    if (p1 == p1) {
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                        fmt.Println(fi61, v5, fi62)
                                        var v63 int64 = v10
                                        _ = v63
                                    } else {
                                        v7 = (127 >= -97)
                                    }
                                    switch int(p1) {
                                    case 26:
                                        var v64 uint8 = (^(v8))
                                        _ = v64
                                    case 493:
                                        v65 := GenS7{F0: p1, F1: 42, F2: 100}
                                        _ = v65
                                    }
                                    if false {
                                        if v7 {
                                            fmt.Println(v8)
                                            var v66 float64 = v5
                                            _ = v66
                                        } else {
                                            switch (-(fi61)) {
                                            case 32768:
                                                if len(v9) > 0 {
                                                    v9[0] = (fi61 * (^(fi61)))
                                                }
                                            case -2147483648:
                                                v67 := GenS8{F0: v10, F1: ((-8 <= 42) || (2 == (v8 / ((v8) | 1))))}
                                                _ = v67
                                            default:
                                                var v68 float64 = (v5 * v5)
                                                _ = v68
                                            }
                                            switch p0 {
                                            case "foo bar":
                                                v6 += v6
                                            case "a":
                                                var v69 int16 = 900
                                                _ = v69
                                            default:
                                                if (!true) {
                                                    var v70 uint16 = 42
                                                    switch p0 {
                                                    case "\n\t":
                                                        switch 32768 {
                                                        case 9223372036854775807:
                                                            fmt.Println(fi62, fi61, v70)
                                                        case 329:
                                                            v6 -= 1
                                                            _ = len(v9)
                                                            if v7 {
                                                                var v71 bool = (957 >= 436)
                                                                if false {
                                                                    if false {
                                                                        v72 := GenS5{F0: -96, F1: v10}
                                                                        var v73 int64 = (v10 - v10)
                                                                        _ = v72
                                                                        _ = v73
                                                                    }
                                                                }
                                                                if len(v9) >= 2 {
                                                                    v9 = v9[1:len(v9)]
                                                                }
                                                                var v74 int16 = p1
                                                                if false {
                                                                    var v75 float64 = 0.1
                                                                    fmt.Println(v71)
                                                                    v71 = v71
                                                                    switch 0 {
                                                                    case 0:
                                                                        p0 = p0
                                                                    default:
                                                                        if v71 {
                                                                            if ((v8 | v8) < (v8 << ((uint(v70)) & 63))) {
                                                                                var v76 uint = 255
                                                                                _ = v76
                                                                            } else {
                                                                                v77 := GenS4{F0: (!true)}
                                                                                fmt.Println(v71, fi61)
                                                                                var v78 uint16 = (v70 << ((255) & 63))
                                                                                _ = v77
                                                                                _ = v78
                                                                            }
                                                                            fmt.Println(fi62)
                                                                            v5 = 0.1
                                                                        }
                                                                    }
                                                                    switch (fi61 * fi62) {
                                                                    case 32768:
                                                                        fmt.Println(p1, fi61)
                                                                        if len(v9) >= 2 {
                                                                            v9 = v9[1:len(v9)]
                                                                        }
                                                                        fmt.Println(fi61, v10, v8)
                                                                    case -476:
                                                                        v74 = (p1 - (p1 & (v74 * (-(v74)))))
                                                                    case 411:
                                                                        var v79 uint16 = (v70 ^ uint16(v8))
                                                                        if len(v9) >= 2 {
                                                                            v9 = v9[1:len(v9)]
                                                                        }
                                                                        _ = v79
                                                                    }
                                                                    _ = v75
                                                                }
                                                                _ = v71
                                                                _ = v74
                                                            }
                                                            _ = len(v9)
                                                        default:
                                                            var v80 uint8 = 255
                                                            _ = v80
                                                        }
                                                    default:
                                                        v5 -= v5
                                                    }
                                                    if (!false) {
                                                        var v81 uint8 = v8
                                                        var v82 bool = (1 == 28)
                                                        fmt.Println(v7)
                                                        v70 += v70
                                                        _ = v81
                                                        _ = v82
                                                    }
                                                    fmt.Println(v5)
                                                    _ = v70
                                                } else {
                                                    _ = len(v9)
                                                }
                                            }
                                        }
                                        if true {
                                            if (!false) {
                                                p0 = p0
                                            }
                                            if v7 {
                                                v83 := GenS9{F0: p0}
                                                fmt.Println(fi62, fi61, p0)
                                                var v84 uint = 65535
                                                _ = v83
                                                _ = v84
                                            } else {
                                                if len(v9) > 0 {
                                                    v9[0] = fi61
                                                }
                                                if len(v9) > 0 {
                                                    v9[0] = (fi61 % (((fi62 & fi62)) | 1))
                                                }
                                                fmt.Println(p1)
                                            }
                                            v7 = v7
                                        } else {
                                            switch "a" {
                                            case "\n\t":
                                                v9 = append(v9, 684)
                                            case "":
                                                if ((3.14159 + 1.5) != ((-2.25 * 1.5) / 1.5)) {
                                                    v7 = v7
                                                    var v85 uint64 = 762
                                                    _ = v85
                                                } else {
                                                    var v86 int8 = 0
                                                    _ = v86
                                                }
                                            case "foo bar":
                                                p1 += (p1 | 1)
                                            }
                                            switch (fi61 - fi61) {
                                            case 32768:
                                                var v87 float32 = (0.0 + 1e10)
                                                _ = v87
                                            default:
                                                v88 := GenS10{F0: 127, F1: (v10 % (((v10 % (((v10 ^ (-(v10)))) | 1))) | 1)), F2: uint64(v8)}
                                                _ = v88
                                            }
                                            switch p0 {
                                            case "\n\t":
                                                v10 = v10
                                            case "foo bar":
                                                fmt.Println(v10, fi62)
                                                v9 = append(v9, (fi62 << ((0) & 63)))
                                            case "":
                                                p0 = ""
                                            default:
                                                var v89 float32 = 0.1
                                                _ = v89
                                            }
                                            var v90 int16 = (p1 + (p1 * 1))
                                            _ = v90
                                        }
                                        if false {
                                            v8 = v8
                                        }
                                    }
                                    var v91 int32 = (v6 - (v6 * v6))
                                    _ = fi62
                                    _ = v91
                                }
                                switch "\n\t" {
                                case "hello":
                                    p0 = "foo bar"
                                case "\n\t":
                                    p0 = "\n\t"
                                case "a":
                                    p1 -= p1
                                default:
                                    var v92 uint64 = 4294967295
                                    _ = v92
                                }
                            }
                            var v93 uint8 = 2
                            _ = fi61
                            _ = v93
                        }
                    }
                }
            }
            _ = v10
            _ = v9
        }
        switch (p0 + ((p0 + p0) + p0)) {
        case "":
            var v94 int8 = 1
            _ = v94
        case "\n\t":
            for fi95 := 0; fi95 < 1; fi95++ {
                for fi96 := 0; fi96 < 18; fi96++ {
                    v97 := GenS11{F0: 42, F1: 42, F2: (p1 + (p1 % ((p1) | 1)))}
                    fmt.Println(fi95)
                    fmt.Println(v7, fi96)
                    if v7 {
                        switch fi95 {
                        case 127:
                            if false {
                                switch p0 {
                                case "a":
                                    switch (fi96 * (fi96 | (^(fi95)))) {
                                    case -2147483648:
                                        v98 := GenS12{F0: 42}
                                        _ = v98
                                    case 127:
                                        if v7 {
                                            switch fi96 {
                                            case 9223372036854775807:
                                                var v99 int64 = 42
                                                _ = v99
                                            case 4294967295:
                                                v100 := GenS13{F0: 1, F1: -1, F2: 572}
                                                _ = v100
                                            default:
                                                switch p0 {
                                                case "hello":
                                                    var v101 bool = (1 <= 0)
                                                    _ = v101
                                                case "a":
                                                    fmt.Println(p1)
                                                    v102 := []int{fi96, -9223372036854775808}
                                                    _ = v102
                                                }
                                            }
                                            var v103 uint8 = v8
                                            var v104 bool = (int64(p1) >= 728)
                                            _ = v103
                                            _ = v104
                                        }
                                    }
                                default:
                                    fmt.Println(v6)
                                }
                            } else {
                                if v7 {
                                    fmt.Println(p0)
                                    var v105 uint32 = 56
                                    _ = v105
                                }
                                var v106 int32 = 766
                                if ((!v7) || v7) {
                                    if true {
                                        if v7 {
                                            p1 = (-(p1))
                                            var v107 string = "foo bar"
                                            _ = v107
                                        } else {
                                            var v108 bool = (int16(v8) > p1)
                                            var v109 uint = 4294967295
                                            _ = v108
                                            _ = v109
                                        }
                                        switch p0 {
                                        case "hello":
                                            v110 := GenS3{F0: "\n\t", F1: 0, F2: v8}
                                            _ = v110
                                        }
                                    } else {
                                        v6 += -2147483648
                                    }
                                    var v111 int32 = v6
                                    if v7 {
                                        fmt.Println(v106)
                                        v111 = v6
                                        if true {
                                            fmt.Println(fi95)
                                            v112 := []int{fi95}
                                            _ = v112
                                        } else {
                                            var v113 int16 = 42
                                            switch p0 {
                                            case "a":
                                                var v114 int64 = -9223372036854775808
                                                _ = v114
                                            case "foo bar":
                                                switch ("hello" + p0) {
                                                case "\n\t":
                                                    if (v5 < v5) {
                                                        if v7 {
                                                            v106 = v106
                                                            if ((!false) || (-128 != -48)) {
                                                                var v115 uint = 641
                                                                switch "foo bar" {
                                                                case "a":
                                                                    v116 := []int{(fi95 | fi96), (fi95 + fi96), fi95}
                                                                    _ = v116
                                                                case "":
                                                                    var v117 int = (fi95 | (fi96 * (fi95 * (fi95 + 1))))
                                                                    _ = v117
                                                                case "\n\t":
                                                                    var v118 float32 = -2.25
                                                                    _ = v118
                                                                }
                                                                _ = v115
                                                            } else {
                                                                v119 := []int{fi95, 2147483647, -1}
                                                                _ = v119
                                                            }
                                                            v120 := []int{(fi95 | fi96), (fi95 | 1), (fi95 + int(v8))}
                                                            fmt.Println(v5, v111)
                                                            v121 := GenS14{F0: v113, F1: v8}
                                                            _ = v120
                                                            _ = v121
                                                        }
                                                    }
                                                case "":
                                                    var v122 int16 = p1
                                                    _ = v122
                                                case "a":
                                                    p0 = p0
                                                }
                                            default:
                                                var v123 uint32 = 731
                                                _ = v123
                                            }
                                            _ = v113
                                        }
                                    } else {
                                        if ((("\n\t" + "") + "hello") <= p0) {
                                            v124 := GenS15{F0: 1, F1: p0, F2: v8}
                                            var v125 int16 = (p1 & (p1 ^ (p1 * (^(p1)))))
                                            _ = v124
                                            _ = v125
                                        }
                                        v126 := []int{(^(fi96)), fi95, fi96}
                                        var v127 uint8 = v8
                                        _ = v126
                                        _ = v127
                                    }
                                    _ = v111
                                }
                                _ = v106
                            }
                        case 128:
                            if (v7 && true) {
                                var v128 int16 = p1
                                if true {
                                    v129 := []int{(fi95 | fi96)}
                                    v130 := GenS6{F0: 2147483647}
                                    switch (p0 + p0) {
                                    case "\n\t":
                                        switch (fi96 | 128) {
                                        case 936:
                                            switch "" {
                                            case "a":
                                                if (1 > -742) {
                                                    var v131 int8 = 127
                                                    _ = v131
                                                }
                                            }
                                        case 2147483647:
                                            p0 = ""
                                        default:
                                            v132 := GenS3{F0: ((("foo bar" + "foo bar") + "") + p0), F1: 65535, F2: (v8 + 127)}
                                            _ = v132
                                        }
                                    default:
                                        v6 += v6
                                    }
                                    _ = v129
                                    _ = v130
                                } else {
                                    v133 := []int{(fi96 | 0), (fi95 & fi95), 0}
                                    v134 := GenS9{F0: p0}
                                    v5 -= v5
                                    var v135 uint16 = 0
                                    var v136 int32 = v6
                                    _ = v133
                                    _ = v134
                                    _ = v135
                                    _ = v136
                                }
                                _ = v128
                            } else {
                                switch (p0 + (("a" + ("" + p0)) + p0)) {
                                case "a":
                                    var v137 uint32 = 1
                                    _ = v137
                                case "\n\t":
                                    var v138 uint64 = 949
                                    _ = v138
                                }
                            }
                        case 32767:
                            v8 += (v8 % ((128) | 1))
                        }
                    }
                    _ = fi96
                    _ = v97
                }
                _ = fi95
            }
        case "foo bar":
            for fi139 := 0; fi139 < 9; fi139++ {
                v140 := []int{(fi139 + (-(fi139)))}
                _ = fi139
                _ = v140
            }
        }
        _ = v8
    } else {
        v141 := GenS5{F0: -627, F1: -282}
        v142 := GenS16{F0: 42, F1: 2147483647}
        fmt.Println(p1, p0, v6)
        _ = v141
        _ = v142
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 208
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []uint32{0, 42, 1, 79}
    var v3 int8 = -33
    var v4 uint = 0
    switch 128 {
    case 32768:
        for fi143 := 0; fi143 < 18; fi143++ {
            v4 += (v4 & 255)
            for fi144 := 0; fi144 < 16; fi144++ {
                var v145 int8 = v3
                _ = fi144
                _ = v145
            }
            _ = fi143
        }
    case -188:
        fmt.Println(v1)
    default:
        v146 := GenS8{F0: 42, F1: (((v4 * 465) != 42) || (65535 <= 65535))}
        _ = v146
    }
    switch 450 {
    case 9223372036854775807:
        for fi147 := 0; fi147 < 7; fi147++ {
            v1 = 931
            _ = fi147
        }
    }
    for fi148 := 0; fi148 < 16; fi148++ {
        if true {
            if (0.0 >= 0.1) {
                var v149 bool = (42 != 2)
                v150 := GenS17{F0: 1, F1: 94}
                v1 -= 279
                _ = v149
                _ = v150
            } else {
                v151 := GenS13{F0: genFn0("\n\t", 32767), F1: v3, F2: 853}
                v2 = append(v2, 42)
                _ = v151
            }
        }
        v3 -= v3
        _ = len(v2)
        _ = fi148
    }
    fmt.Println(v4, v4, v4)
    _ = len(v2)
    v1 = (v1 & v1)
    for fi152 := 0; fi152 < 13; fi152++ {
        fmt.Println(fi152, v1, v4)
        v1 -= (^(v1))
        v2 = append(v2, 587)
        _ = fi152
    }
    if len(v2) > 0 {
        v2[0] = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
