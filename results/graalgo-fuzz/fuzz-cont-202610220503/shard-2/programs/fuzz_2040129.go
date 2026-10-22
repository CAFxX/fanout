package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool; F1 int }
type GenS7 struct { F0 string; F1 int; F2 uint }
type GenS8 struct { F0 string; F1 uint32; F2 uint64 }
type GenS9 struct { F0 int16; F1 uint8 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint16; F1 int16; F2 string }
type GenS13 struct { F0 uint64; F1 uint32; F2 uint }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 int8; F1 int64 }
type GenS16 struct { F0 uint16; F1 bool }
type GenS17 struct { F0 uint16; F1 string }
type GenS18 struct { F0 int16; F1 int32 }
type GenS19 struct { F0 int; F1 bool; F2 int16 }
type GenS20 struct { F0 int32; F1 int }
type GenS21 struct { F0 int; F1 uint16; F2 string }
type GenS22 struct { F0 uint16; F1 string; F2 int8 }
type GenS23 struct { F0 int64 }
type GenS24 struct { F0 bool; F1 uint32 }
type GenS25 struct { F0 uint64; F1 int; F2 uint64 }
type GenS26 struct { F0 uint8; F1 bool; F2 string }
type GenS27 struct { F0 int; F1 uint32; F2 int8 }
type GenS28 struct { F0 int32; F1 int64; F2 int8 }
type GenS29 struct { F0 uint16; F1 int }
type GenS30 struct { F0 int16 }
type GenS31 struct { F0 int32; F1 int64; F2 uint8 }
type GenS32 struct { F0 int; F1 uint }
type GenS33 struct { F0 int8; F1 uint16; F2 int64 }
type GenS34 struct { F0 uint8; F1 uint64 }
type GenS35 struct { F0 uint64 }
type GenS36 struct { F0 bool }
type GenS37 struct { F0 uint8 }
type GenS38 struct { F0 string }
func genFn0(p0 uint32, p1 uint) int {
    fmt.Println(p0)
    v4 := []int{}
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 int = (-(fi5))
        var v7 int8 = -1
        _ = fi5
        _ = v6
        _ = v7
    }
    if ((1 <= -1) || false) {
        if true {
            p1 += 65535
            for fi8 := 0; fi8 < 2; fi8++ {
                if false {
                    v4 = append(v4, 9223372036854775807)
                } else {
                    var v9 int32 = -1
                    v9 -= v9
                    _ = v9
                }
                switch "\n\t" {
                case "":
                    if (("foo bar" == (("" + "a") + ("" + ("\n\t" + "a")))) || (42 <= -1)) {
                        var v10 int64 = -1
                        var v11 uint16 = 971
                        v12 := GenS1{F0: v10, F1: p0}
                        _ = v10
                        _ = v11
                        _ = v12
                    } else {
                        for fi13 := 0; fi13 < 1; fi13++ {
                            _ = len(v4)
                            _ = fi13
                        }
                        for fi14 := 0; fi14 < 4; fi14++ {
                            v4 = append(v4, (fi14 ^ -615))
                            var v15 uint32 = uint32(p1)
                            v16 := GenS2{F0: 0}
                            p0 -= p0
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        v17 := GenS2{F0: -788}
                        _ = v17
                    }
                case "\n\t":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    if false {
                        switch (fi8 * fi8) {
                        case 128:
                            fmt.Println(p0, fi8, p1)
                        case 322:
                            v18 := GenS3{F0: 1, F1: 32767}
                            _ = v18
                        default:
                            var v19 int8 = 1
                            _ = v19
                        }
                        v20 := GenS3{F0: 2, F1: -1}
                        var v21 uint16 = 640
                        _ = v20
                        _ = v21
                    } else {
                        var v22 uint = (p1 - p1)
                        var v23 int8 = 0
                        _ = v22
                        _ = v23
                    }
                }
                v24 := GenS4{F0: 125}
                _ = fi8
                _ = v24
            }
            v25 := GenS4{F0: 255}
            _ = v25
        }
        var v26 int32 = 0
        _ = v26
    }
    _ = p0
    _ = p1
    _ = v4
    return -1
}
func genFn1() uint {
    var v27 int32 = 1
    v27 = -2147483648
    v28 := GenS4{F0: 32767}
    for fi29 := 0; fi29 < 12; fi29++ {
        v30 := GenS3{F0: 682, F1: -129}
        if (!false) {
            var v31 int8 = 100
            switch "hello" {
            case "\n\t":
                v32 := []int{(fi29 + (^(fi29))), fi29, -1, fi29}
                _ = v32
            case "hello":
                if (true && true) {
                    v33 := GenS4{F0: 695}
                    if (-32768 == 743) {
                        fmt.Println(v31, v27)
                    }
                    var v34 int16 = 32767
                    _ = v33
                    _ = v34
                } else {
                    var v35 float64 = 1.5
                    var v36 float64 = -2.25
                    if false {
                        v37 := []int{genFn0(2147483648, 747), (fi29 % (((fi29 * genFn0(0, 1))) | 1))}
                        for fi38 := 0; fi38 < 13; fi38++ {
                            v36 -= float64(0.1)
                            if (4294967295 <= 18446744073709551615) {
                                v36 -= (((v36 * 1.5) + (0.0 * v35)) / v35)
                                v27 -= v27
                            }
                            var v39 int32 = v27
                            v40 := GenS5{F0: fi38}
                            _ = fi38
                            _ = v39
                            _ = v40
                        }
                        _ = v37
                    }
                    _ = v35
                    _ = v36
                }
            default:
                var v41 uint32 = 0
                _ = v41
            }
            _ = v31
        }
        v27 += v27
        _ = fi29
        _ = v30
    }
    if (!(!((0.0 + 0.1) < 0.0))) {
        v42 := []int{-2147483648}
        v43 := GenS4{F0: 0}
        v27 = (v27 ^ v27)
        _ = v42
        _ = v43
    } else {
        v27 += (v27 ^ (v27 % (((v27 & v27)) | 1)))
    }
    _ = v27
    _ = v28
    return 805
}

func main() {
    var v1 bool = (false || true)
    var v2 string = "\n\t"
    v3 := []uint32{}
    fmt.Println(v2)
    var v44 string = ""
    v45 := GenS6{F0: v1, F1: -9223372036854775808}
    var v46 string = v44
    v1 = v1
    if v1 {
        var v47 string = (v46 + "\n\t")
        v3 = append(v3, 1)
        if (2 > genFn1()) {
            for fi48 := 0; fi48 < 9; fi48++ {
                v47 = v44
                _ = fi48
            }
        } else {
            var v49 int8 = 42
            _ = v49
        }
        _ = v47
    }
    switch v2 {
    case "\n\t":
        if len(v3) > 0 {
            v3[0] = 208
        }
    default:
        for fi50 := 0; fi50 < 11; fi50++ {
            v44 = ""
            var v51 uint64 = 4294967295
            if (1.5 >= (-2.25 + float64(1.5))) {
                for fi52 := 0; fi52 < 13; fi52++ {
                    v53 := GenS1{F0: -9223372036854775808, F1: 2147483648}
                    _ = fi52
                    _ = v53
                }
            } else {
                _ = len(v3)
                v46 = v44
                fmt.Println(v2)
            }
            _ = fi50
            _ = v51
        }
    }
    if len(v3) > 0 {
        v3[0] = 0
    }
    for fi54 := 0; fi54 < 5; fi54++ {
        v3 = append(v3, 1)
        v55 := GenS7{F0: "a", F1: (^(fi54)), F2: 538}
        _ = fi54
        _ = v55
    }
    for fi56 := 0; fi56 < 11; fi56++ {
        for fi57 := 0; fi57 < 1; fi57++ {
            v58 := GenS8{F0: (v46 + v46), F1: 0, F2: 42}
            _ = fi57
            _ = v58
        }
        v46 = v2
        _ = fi56
    }
    var v59 int = genFn0(1, 4294967295)
    v2 = v2
    if true {
        for fi60 := 0; fi60 < 5; fi60++ {
            switch ((v2 + (v2 + v2)) + v44) {
            case "":
                switch fi60 {
                case -975:
                    switch v44 {
                    case "a":
                        var v61 float64 = (1.5 * 1e10)
                        _ = v61
                    case "":
                        var v62 uint32 = 42
                        _ = v62
                    default:
                        var v63 bool = v1
                        _ = v63
                    }
                case -9223372036854775808:
                    if ((v44 + (v46 + ("" + ""))) == v44) {
                        if (3.14159 <= (0.0 + ((((0.0 / ((0.1 * 3.14159) + 0.1)) - 1.5) - 0.0) + 1.5))) {
                            fmt.Println(v59, v44, v1)
                        } else {
                            if v1 {
                                for fi64 := 0; fi64 < 17; fi64++ {
                                    var v65 int16 = 1
                                    if true {
                                        v66 := GenS4{F0: 2}
                                        var v67 int = fi64
                                        var v68 string = ""
                                        _ = v66
                                        _ = v67
                                        _ = v68
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = 65535
                                        }
                                    }
                                    v3 = append(v3, 857)
                                    _ = fi64
                                    _ = v65
                                }
                                var v69 float32 = 1e10
                                _ = v69
                            } else {
                                v1 = (0.0 > 3.14159)
                                v1 = (false || v1)
                                var v70 uint32 = 406
                                _ = v70
                            }
                        }
                    } else {
                        for fi71 := 0; fi71 < 4; fi71++ {
                            if v1 {
                                var v72 int32 = -1
                                fmt.Println(v44)
                                v73 := GenS4{F0: 12}
                                _ = v72
                                _ = v73
                            } else {
                                if (v44 < "\n\t") {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                } else {
                                    v2 = ""
                                }
                                if (724 <= -300) {
                                    switch (fi60 | (fi60 % ((4294967295) | 1))) {
                                    case 1:
                                        fmt.Println(v44)
                                    case 32767:
                                        if v1 {
                                            v74 := GenS7{F0: "a", F1: (fi71 / ((82) | 1)), F2: 42}
                                            _ = v74
                                        }
                                    default:
                                        var v75 uint32 = 42
                                        _ = v75
                                    }
                                } else {
                                    v76 := GenS9{F0: -1, F1: 208}
                                    var v77 uint = 255
                                    _ = v76
                                    _ = v77
                                }
                                switch (^(fi71)) {
                                case -2147483648:
                                    v78 := GenS4{F0: 152}
                                    _ = v78
                                }
                                switch (-(fi60)) {
                                case 884:
                                    v79 := GenS4{F0: 32767}
                                    _ = v79
                                case 465:
                                    v1 = (255 >= 2)
                                    fmt.Println(v1, v44)
                                default:
                                    if v1 {
                                        switch ("a" + (((((v2 + (v44 + v46)) + (v44 + (v46 + v2))) + "a") + (("hello" + v46) + "foo bar")) + v46)) {
                                        case "a":
                                            if (genFn1() != 1) {
                                                v44 = "foo bar"
                                            }
                                        case "\n\t":
                                            var v80 string = ("foo bar" + "a")
                                            _ = v80
                                        default:
                                            if true {
                                                if true {
                                                    var v81 uint64 = 332
                                                    var v82 int = (fi60 ^ (fi71 + (fi60 * fi60)))
                                                    v59 -= 0
                                                    _ = v81
                                                    _ = v82
                                                }
                                                switch -1 {
                                                case 1:
                                                    fmt.Println(v2, v2)
                                                default:
                                                    switch (^(fi71)) {
                                                    case 128:
                                                        if (float32(0.1) > -2.25) {
                                                            fmt.Println(v46, fi71)
                                                        } else {
                                                            var v83 uint8 = 10
                                                            if v1 {
                                                                fmt.Println(v44, v59, v59)
                                                                var v84 int32 = 1
                                                                if ((v59 ^ (fi60 - (v59 - (fi60 % ((fi71) | 1))))) != genFn0(330, 722)) {
                                                                    _ = len(v3)
                                                                    var v85 float32 = 3.14159
                                                                    v85 -= v85
                                                                    _ = v85
                                                                }
                                                                _ = v84
                                                            } else {
                                                                var v86 bool = v1
                                                                v3 = append(v3, 0)
                                                                var v87 uint = 0
                                                                _ = v86
                                                                _ = v87
                                                            }
                                                            v88 := GenS10{F0: 0}
                                                            _ = v83
                                                            _ = v88
                                                        }
                                                    case -1:
                                                        v1 = ((true && v1) || v1)
                                                    case 0:
                                                        if v1 {
                                                            var v89 int32 = 42
                                                            var v90 int32 = v89
                                                            var v91 bool = (-128 <= 100)
                                                            var v92 uint32 = 146
                                                            _ = v89
                                                            _ = v90
                                                            _ = v91
                                                            _ = v92
                                                        }
                                                    }
                                                }
                                                v1 = (18446744073709551615 != 18446744073709551615)
                                                var v93 uint = 867
                                                _ = v93
                                            } else {
                                                v44 = ("foo bar" + v2)
                                            }
                                        }
                                        if v1 {
                                            fmt.Println(v59, v44, v2)
                                            v46 = ((v2 + v44) + (("a" + "foo bar") + v46))
                                            var v94 int64 = -408
                                            _ = v94
                                        } else {
                                            v1 = (42 != 65535)
                                        }
                                        v46 = ("\n\t" + "")
                                        if v1 {
                                            if len(v3) > 0 {
                                                v3[0] = 2147483648
                                            }
                                        }
                                    } else {
                                        var v95 bool = v1
                                        v96 := GenS11{F0: fi60}
                                        _ = v95
                                        _ = v96
                                    }
                                }
                                if v1 {
                                    var v97 uint8 = 1
                                    var v98 int64 = 1
                                    if (100 > 0) {
                                        var v99 uint = 0
                                        _ = len(v3)
                                        v100 := GenS10{F0: (v99 + 42)}
                                        switch -9223372036854775808 {
                                        case 0:
                                            v98 = (v98 + -831)
                                            if (751 == uint16(v97)) {
                                                var v101 uint32 = 4294967295
                                                _ = v101
                                            } else {
                                                if v1 {
                                                    fmt.Println(v99, v97, fi60)
                                                    fmt.Println(fi71, v97)
                                                }
                                                if v1 {
                                                    var v102 int8 = -1
                                                    var v103 string = "foo bar"
                                                    if (!true) {
                                                        fmt.Println(v44, v44)
                                                        var v104 float64 = 0.0
                                                        var v105 float64 = (1.5 / 0.1)
                                                        var v106 int64 = v98
                                                        _ = v104
                                                        _ = v105
                                                        _ = v106
                                                    }
                                                    _ = v102
                                                    _ = v103
                                                } else {
                                                    v98 -= (v98 | v98)
                                                    v107 := GenS6{F0: (688 > 1), F1: 0}
                                                    _ = v107
                                                }
                                            }
                                            fmt.Println(v2, v99)
                                        }
                                        _ = v100
                                        _ = v99
                                    }
                                    _ = v97
                                    _ = v98
                                } else {
                                    var v108 int32 = int32(fi60)
                                    _ = v108
                                }
                                var v109 int = (fi60 - v59)
                                if v1 {
                                    switch v46 {
                                    case "\n\t":
                                        if v1 {
                                            fmt.Println(v109, fi60)
                                        } else {
                                            var v110 float32 = ((0.1 * -2.25) + (1e10 / 1.5))
                                            if (1 >= genFn1()) {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                fmt.Println(v44, v59, fi71)
                                            }
                                            if v1 {
                                                fmt.Println(v1, v109, v110)
                                                var v111 uint32 = 385
                                                _ = v111
                                            } else {
                                                var v112 int64 = -1
                                                var v113 int32 = -1
                                                var v114 int64 = (v112 / (((v112 | 42)) | 1))
                                                _ = v112
                                                _ = v113
                                                _ = v114
                                            }
                                            if ((false || (0 < 32767)) && (!(1e10 < float64((v110 - v110))))) {
                                                v109 += (v59 % ((fi71) | 1))
                                            } else {
                                                fmt.Println(v46, v109)
                                                v115 := GenS8{F0: "hello", F1: 2, F2: 638}
                                                _ = v115
                                            }
                                            _ = v110
                                        }
                                    }
                                }
                                _ = v109
                            }
                            _ = fi71
                        }
                        if v1 {
                            switch "\n\t" {
                            case "hello":
                                v59 += (-(v59))
                            default:
                                var v116 int = 193
                                _ = v116
                            }
                            switch (v44 + v44) {
                            case "hello":
                                var v117 uint8 = 128
                                _ = v117
                            }
                        } else {
                            for fi118 := 0; fi118 < 19; fi118++ {
                                if v1 {
                                    if (v2 > v44) {
                                        if (1 >= 0) {
                                            v119 := GenS9{F0: 351, F1: 164}
                                            if (255 > 2) {
                                                v2 = ("a" + (v44 + "hello"))
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                            switch "hello" {
                                            case "":
                                                var v120 int32 = -2147483648
                                                v120 = (v120 % ((0) | 1))
                                                _ = v120
                                            }
                                            _ = v119
                                        }
                                    } else {
                                        switch (^(v59)) {
                                        case -2147483648:
                                            var v121 int32 = -1
                                            _ = v121
                                        case 805:
                                            switch v2 {
                                            case "foo bar":
                                                var v122 float32 = 0.1
                                                _ = v122
                                            default:
                                                _ = len(v3)
                                            }
                                        case 9223372036854775807:
                                            v123 := GenS12{F0: 679, F1: -552, F2: "\n\t"}
                                            _ = v123
                                        }
                                        var v124 float32 = -2.25
                                        v125 := GenS13{F0: 98, F1: 499, F2: 65535}
                                        v126 := GenS14{F0: 128}
                                        _ = v124
                                        _ = v125
                                        _ = v126
                                    }
                                } else {
                                    var v127 uint64 = 0
                                    _ = v127
                                }
                                v3 = append(v3, 2147483647)
                                _ = fi118
                            }
                            if (127 > 49) {
                                var v128 uint = 42
                                _ = v128
                            } else {
                                for fi129 := 0; fi129 < 1; fi129++ {
                                    v130 := GenS15{F0: 100, F1: 110}
                                    v131 := GenS9{F0: 42, F1: 42}
                                    if (!((!(v1 || true)) && v1)) {
                                        fmt.Println(fi60, v59, fi60)
                                        var v132 uint = 182
                                        _ = v132
                                    } else {
                                        if (-1 <= 32767) {
                                            switch v44 {
                                            case "":
                                                v59 = genFn0(2147483647, genFn1())
                                                var v133 int8 = 127
                                                if (1 <= -2147483648) {
                                                    if v1 {
                                                        var v134 int8 = v133
                                                        _ = v134
                                                    } else {
                                                        if v1 {
                                                            switch "hello" {
                                                            case "\n\t":
                                                                var v135 int64 = -1
                                                                _ = v135
                                                            case "foo bar":
                                                                v136 := GenS16{F0: 2, F1: (!v1)}
                                                                _ = v136
                                                            case "a":
                                                                _ = len(v3)
                                                            default:
                                                                var v137 int8 = -1
                                                                _ = v137
                                                            }
                                                            v2 = v44
                                                        }
                                                        v138 := GenS17{F0: 1, F1: "\n\t"}
                                                        _ = v138
                                                    }
                                                    v44 = v46
                                                } else {
                                                    var v139 float64 = (0.0 / 1.5)
                                                    _ = v139
                                                }
                                                _ = v133
                                            case "foo bar":
                                                if (v46 != "foo bar") {
                                                    var v140 int = genFn0(4294967295, 4294967295)
                                                    v59 += (fi129 | v59)
                                                    _ = v140
                                                } else {
                                                    switch fi60 {
                                                    case -804:
                                                        if true {
                                                            fmt.Println(v44, v44, v2)
                                                            var v141 bool = v1
                                                            _ = v141
                                                        } else {
                                                            v1 = v1
                                                        }
                                                    case 32768:
                                                        var v142 uint32 = 2
                                                        _ = v142
                                                    case 0:
                                                        v1 = v1
                                                    }
                                                    v143 := GenS4{F0: 32767}
                                                    v3 = append(v3, 502)
                                                    _ = v143
                                                }
                                            case "\n\t":
                                                fmt.Println(fi129, v1, v46)
                                            default:
                                                switch 917 {
                                                case 32768:
                                                    v144 := GenS18{F0: -32768, F1: 0}
                                                    _ = v144
                                                default:
                                                    if (!v1) {
                                                        var v145 int64 = 0
                                                        if v1 {
                                                            if v1 {
                                                                if (v145 < v145) {
                                                                    var v146 float64 = 1.5
                                                                    _ = v146
                                                                }
                                                                if true {
                                                                    fmt.Println(v44, v145)
                                                                    if true {
                                                                        fmt.Println(v2, v2, v2)
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            var v147 string = v2
                                                                            _ = v147
                                                                        case "foo bar":
                                                                            switch fi129 {
                                                                            case -2147483648:
                                                                                if ((!(genFn1() != genFn1())) && false) {
                                                                                    v1 = (1 > -1)
                                                                                }
                                                                            case -275:
                                                                                v145 += (v145 / ((v145) | 1))
                                                                            case 9223372036854775807:
                                                                                if (0 > 0) {
                                                                                    v59 += (fi129 >> ((537) & 63))
                                                                                } else {
                                                                                    v148 := GenS19{F0: 4294967295, F1: false, F2: 0}
                                                                                    _ = v148
                                                                                }
                                                                            default:
                                                                                switch v59 {
                                                                                case -9223372036854775808:
                                                                                    v149 := GenS20{F0: 1, F1: genFn0(420, 87)}
                                                                                    _ = v149
                                                                                case -815:
                                                                                    if false {
                                                                                        var v150 int64 = (^(v145))
                                                                                        var v151 uint = genFn1()
                                                                                        _ = v150
                                                                                        _ = v151
                                                                                    }
                                                                                case 42:
                                                                                    v59 = (v59 ^ genFn0(519, genFn1()))
                                                                                default:
                                                                                    v2 = v46
                                                                                }
                                                                            }
                                                                        case "":
                                                                            v152 := GenS21{F0: (fi60 | 1), F1: 65535, F2: (v46 + v2)}
                                                                            _ = v152
                                                                        }
                                                                        v153 := GenS18{F0: 42, F1: 2147483647}
                                                                        _ = v153
                                                                    }
                                                                }
                                                                if (!(2 <= 0)) {
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    if (!(true || v1)) {
                                                                        switch -2147483648 {
                                                                        case 32768:
                                                                            var v154 uint32 = 2
                                                                            _ = v154
                                                                        case 42:
                                                                            fmt.Println(v46, v145)
                                                                            switch genFn0(2147483648, 18446744073709551615) {
                                                                            case 1:
                                                                                if v1 {
                                                                                    v44 = "\n\t"
                                                                                } else {
                                                                                    v59 -= fi60
                                                                                }
                                                                            default:
                                                                                var v155 bool = (1 <= 2)
                                                                                _ = v155
                                                                            }
                                                                            var v156 uint = 266
                                                                            _ = v156
                                                                        case -1:
                                                                            v157 := GenS5{F0: v59}
                                                                            _ = v157
                                                                        default:
                                                                            switch "" {
                                                                            case "":
                                                                                _ = len(v3)
                                                                            default:
                                                                                fmt.Println(v145)
                                                                            }
                                                                        }
                                                                    } else {
                                                                        var v158 int8 = 42
                                                                        _ = v158
                                                                    }
                                                                }
                                                            }
                                                            switch "hello" {
                                                            case "\n\t":
                                                                fmt.Println(fi60, fi129, fi129)
                                                            case "":
                                                                if true {
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                v46 = ((v44 + (v46 + v2)) + ("\n\t" + ""))
                                                            }
                                                            switch 0 {
                                                            case 127:
                                                                v159 := GenS22{F0: 0, F1: "\n\t", F2: -42}
                                                                _ = v159
                                                            case -324:
                                                                v160 := GenS13{F0: 364, F1: 42, F2: genFn1()}
                                                                _ = v160
                                                            case 2147483647:
                                                                v59 += v59
                                                            default:
                                                                var v161 uint = 18446744073709551615
                                                                _ = v161
                                                            }
                                                            _ = len(v3)
                                                        } else {
                                                            if false {
                                                                var v162 int64 = (v145 + v145)
                                                                _ = v162
                                                            } else {
                                                                v2 = "foo bar"
                                                                var v163 float64 = 1.5
                                                                v3 = append(v3, 4294967295)
                                                                _ = v163
                                                            }
                                                            var v164 int8 = 0
                                                            fmt.Println(v2, fi60, v145)
                                                            var v165 float32 = float32(3.14159)
                                                            _ = v164
                                                            _ = v165
                                                        }
                                                        var v166 int8 = 127
                                                        _ = v145
                                                        _ = v166
                                                    } else {
                                                        if (1 >= 42) {
                                                            v167 := GenS3{F0: 0, F1: -52}
                                                            _ = v167
                                                        }
                                                    }
                                                }
                                            }
                                            fmt.Println(v59, v46, v2)
                                        }
                                    }
                                    v2 = (v2 + (v2 + ""))
                                    v46 = "a"
                                    _ = fi129
                                    _ = v130
                                    _ = v131
                                }
                                v168 := GenS23{F0: 42}
                                if v1 {
                                    v169 := GenS14{F0: 2}
                                    for fi170 := 0; fi170 < 17; fi170++ {
                                        var v171 int32 = -1
                                        _ = fi170
                                        _ = v171
                                    }
                                    _ = v169
                                } else {
                                    var v172 uint16 = 32767
                                    switch (fi60 & (fi60 ^ (v59 * (^(v59))))) {
                                    case 32767:
                                        v173 := GenS5{F0: (fi60 * (v59 << ((2) & 63)))}
                                        _ = v173
                                    case 0:
                                        for fi174 := 0; fi174 < 20; fi174++ {
                                            if v1 {
                                                v175 := GenS21{F0: genFn0(4294967295, genFn1()), F1: 950, F2: ""}
                                                switch genFn0(513, 2) {
                                                case 32767:
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    v44 = (("a" + v44) + (((("hello" + (("\n\t" + v46) + v44)) + "a") + (((("foo bar" + v44) + v46) + v44) + "a")) + (((((v46 + "\n\t") + ("hello" + v44)) + ("a" + "\n\t")) + (v44 + "\n\t")) + v2)))
                                                    var v176 int32 = 658
                                                    _ = v176
                                                case -1:
                                                    if (!(255 != 42)) {
                                                        v177 := GenS24{F0: ((float32(((float64(1.5) * -2.25) * 3.14159)) / 0.1) <= ((0.1 - (-2.25 * (0.1 + 0.0))) - 0.0)), F1: 42}
                                                        var v178 int16 = -32768
                                                        var v179 uint32 = 2147483648
                                                        _ = v177
                                                        _ = v178
                                                        _ = v179
                                                    }
                                                case -215:
                                                    if false {
                                                        var v180 int64 = 960
                                                        switch (fi60 * fi174) {
                                                        case 1:
                                                            var v181 uint64 = 2
                                                            _ = v181
                                                        case 32767:
                                                            v182 := GenS16{F0: (v172 | v172), F1: v1}
                                                            _ = v182
                                                        default:
                                                            fmt.Println(v44, fi60)
                                                        }
                                                        _ = v180
                                                    } else {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        _ = len(v3)
                                                        if (109 == 4294967295) {
                                                            v183 := GenS14{F0: 1}
                                                            v44 = ""
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            if true {
                                                                var v184 float64 = 0.0
                                                                fmt.Println(v59, fi60, v172)
                                                                _ = v184
                                                            }
                                                            v185 := GenS25{F0: 18446744073709551615, F1: genFn0(728, 2), F2: 42}
                                                            _ = v183
                                                            _ = v185
                                                        } else {
                                                            v186 := GenS10{F0: 250}
                                                            v44 = v44
                                                            switch v46 {
                                                            case "hello":
                                                                switch (fi174 ^ (fi174 - (v59 * fi60))) {
                                                                case 128:
                                                                    var v187 int32 = 0
                                                                    _ = v187
                                                                case 4294967295:
                                                                    v3 = append(v3, 42)
                                                                case -9223372036854775808:
                                                                    v2 = ("hello" + "foo bar")
                                                                default:
                                                                    if (0 != 128) {
                                                                        v188 := GenS26{F0: 42, F1: v1, F2: "foo bar"}
                                                                        _ = v188
                                                                    }
                                                                }
                                                            }
                                                            v1 = v1
                                                            fmt.Println(fi60, v44, v172)
                                                            v189 := GenS7{F0: v2, F1: (fi60 ^ (-(v59))), F2: 597}
                                                            _ = v186
                                                            _ = v189
                                                        }
                                                    }
                                                default:
                                                    v190 := GenS27{F0: v59, F1: 2147483648, F2: -80}
                                                    _ = v190
                                                }
                                                _ = v175
                                            }
                                            _ = fi174
                                        }
                                    case -2147483648:
                                        if true {
                                            fmt.Println(v2, v172)
                                            if (-9223372036854775808 > -9223372036854775808) {
                                                var v191 uint8 = 255
                                                for fi192 := 0; fi192 < 20; fi192++ {
                                                    fmt.Println(fi192, fi192)
                                                    var v193 uint8 = v191
                                                    v194 := GenS28{F0: int32(v191), F1: 9223372036854775807, F2: 127}
                                                    if (255 != 13) {
                                                        var v195 uint8 = v193
                                                        _ = v195
                                                    }
                                                    _ = fi192
                                                    _ = v193
                                                    _ = v194
                                                }
                                                _ = v191
                                            }
                                        } else {
                                            v44 = (("hello" + "\n\t") + "foo bar")
                                            v196 := GenS11{F0: (fi60 << ((1) & 63))}
                                            _ = v196
                                        }
                                    default:
                                        var v197 uint16 = (^(v172))
                                        _ = v197
                                    }
                                    _ = v172
                                }
                                _ = v168
                            }
                            for fi198 := 0; fi198 < 1; fi198++ {
                                switch 0 {
                                case -9223372036854775808:
                                    if len(v3) > 0 {
                                        v3[0] = 2147483647
                                    }
                                default:
                                    var v199 uint64 = 2
                                    _ = v199
                                }
                                _ = fi198
                            }
                        }
                    }
                }
            case "\n\t":
                if v1 {
                    switch "hello" {
                    case "foo bar":
                        if (52 >= 100) {
                            v59 = v59
                        } else {
                            var v200 float32 = 0.1
                            var v201 bool = (2 < 1)
                            _ = v200
                            _ = v201
                        }
                    default:
                        if v1 {
                            v46 = v46
                        } else {
                            fmt.Println(v59, v46)
                            v46 = ("" + (v44 + (v44 + v46)))
                            if (0 < -18) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if v1 {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v46 = ("\n\t" + v46)
                                    switch (fi60 << ((65535) & 63)) {
                                    case -9223372036854775808:
                                        v46 = (v2 + "a")
                                    }
                                } else {
                                    fmt.Println(v44)
                                    var v202 string = v46
                                    for fi203 := 0; fi203 < 15; fi203++ {
                                        if v1 {
                                            fmt.Println(fi60, fi203, v2)
                                            v204 := GenS29{F0: 1, F1: (fi203 % ((fi60) | 1))}
                                            fmt.Println(v1)
                                            v3 = append(v3, 2)
                                            v205 := GenS5{F0: fi60}
                                            v2 = ""
                                            switch "\n\t" {
                                            case "a":
                                                switch fi203 {
                                                case -240:
                                                    if (3.14159 >= 1.5) {
                                                        v206 := GenS30{F0: 1000}
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        if v1 {
                                                            if (4294967295 <= 293) {
                                                                var v207 uint16 = 65535
                                                                v59 -= fi203
                                                                _ = v207
                                                            }
                                                            fmt.Println(fi60)
                                                        } else {
                                                            _ = len(v3)
                                                            if false {
                                                                switch v59 {
                                                                case 32767:
                                                                    switch fi203 {
                                                                    case 32768:
                                                                        var v208 int64 = 0
                                                                        _ = v208
                                                                    case 2147483647:
                                                                        var v209 int8 = 48
                                                                        _ = v209
                                                                    default:
                                                                        switch fi60 {
                                                                        case 128:
                                                                            switch genFn0(2147483648, 1) {
                                                                            case -630:
                                                                                var v210 float64 = 0.0
                                                                                _ = v210
                                                                            case 2147483647:
                                                                                v211 := GenS30{F0: 1}
                                                                                _ = v211
                                                                            default:
                                                                                switch fi203 {
                                                                                case 971:
                                                                                    v212 := GenS12{F0: 32767, F1: 42, F2: "hello"}
                                                                                    _ = v212
                                                                                case -148:
                                                                                    v3 = append(v3, 2147483647)
                                                                                default:
                                                                                    fmt.Println(v44, fi60)
                                                                                }
                                                                            }
                                                                        case -9223372036854775808:
                                                                            var v213 int32 = 0
                                                                            _ = v213
                                                                        case -964:
                                                                            var v214 uint32 = 42
                                                                            _ = v214
                                                                        default:
                                                                            if (18446744073709551615 <= 1) {
                                                                                _ = len(v3)
                                                                                fmt.Println(v46, v59, v44)
                                                                            }
                                                                        }
                                                                    }
                                                                case -9223372036854775808:
                                                                    v3 = append(v3, 417)
                                                                case 1:
                                                                    v2 = (((((v2 + v46) + v202) + ("\n\t" + v46)) + v46) + ((("\n\t" + v46) + ("" + v46)) + "\n\t"))
                                                                default:
                                                                    fmt.Println(v46, v59)
                                                                }
                                                            } else {
                                                                v215 := GenS3{F0: 836, F1: 42}
                                                                v44 = v2
                                                                _ = v215
                                                            }
                                                        }
                                                        if (((!true) && false) || false) {
                                                            v44 = ""
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            switch v2 {
                                                            case "\n\t":
                                                                _ = len(v3)
                                                            }
                                                        }
                                                        _ = v206
                                                    }
                                                default:
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                            default:
                                                if v1 {
                                                    var v216 uint32 = 1
                                                    _ = v216
                                                } else {
                                                    v202 = "hello"
                                                    switch -1 {
                                                    case 152:
                                                        var v217 int64 = 1
                                                        _ = v217
                                                    case -9223372036854775808:
                                                        var v218 uint8 = 57
                                                        _ = v218
                                                    default:
                                                        var v219 uint64 = 4294967295
                                                        _ = v219
                                                    }
                                                    var v220 float32 = 3.14159
                                                    _ = v220
                                                }
                                            }
                                            _ = v204
                                            _ = v205
                                        } else {
                                            fmt.Println(v202)
                                        }
                                        v202 = (v202 + v2)
                                        switch ("" + "\n\t") {
                                        case "\n\t":
                                            if true {
                                                if true {
                                                    if false {
                                                        v59 = (fi203 * genFn0(2, 4294967295))
                                                        var v221 float32 = ((3.14159 - ((1e10 * (1e10 * 1.5)) - 3.14159)) + -2.25)
                                                        _ = v221
                                                    }
                                                }
                                            } else {
                                                v59 += fi60
                                                var v222 int16 = 42
                                                var v223 uint32 = 60
                                                _ = v222
                                                _ = v223
                                            }
                                        }
                                        _ = fi203
                                    }
                                    _ = v202
                                }
                                for fi224 := 0; fi224 < 1; fi224++ {
                                    var v225 uint64 = 2
                                    v226 := GenS31{F0: 2147483647, F1: 42, F2: 2}
                                    v227 := GenS10{F0: 42}
                                    _ = fi224
                                    _ = v225
                                    _ = v226
                                    _ = v227
                                }
                            } else {
                                for fi228 := 0; fi228 < 20; fi228++ {
                                    if v1 {
                                        v1 = (42 <= -128)
                                        fmt.Println(v44, fi228, fi228)
                                        var v229 int = (fi228 | fi228)
                                        _ = v229
                                    }
                                    if true {
                                        v46 = v46
                                        v2 = (("\n\t" + v2) + ("a" + v44))
                                    }
                                    var v230 int = 128
                                    _ = fi228
                                    _ = v230
                                }
                                v59 = (fi60 + (fi60 * -457))
                                switch -893 {
                                case 9223372036854775807:
                                    var v231 uint16 = 267
                                    _ = v231
                                case 2147483647:
                                    for fi232 := 0; fi232 < 1; fi232++ {
                                        if len(v3) > 0 {
                                            v3[0] = 628
                                        }
                                        _ = fi232
                                    }
                                default:
                                    fmt.Println(v46, v59, v1)
                                }
                            }
                        }
                    }
                    v46 = ("foo bar" + "a")
                    for fi233 := 0; fi233 < 4; fi233++ {
                        switch "hello" {
                        case "foo bar":
                            switch (fi60 * (fi233 - 32768)) {
                            case 2147483647:
                                v2 = ("foo bar" + ("a" + v44))
                            default:
                                v44 = (v44 + (v2 + v44))
                            }
                        case "hello":
                            var v234 int16 = 1000
                            _ = v234
                        case "a":
                            var v235 uint16 = 2
                            _ = v235
                        default:
                            var v236 uint64 = 257
                            _ = v236
                        }
                        _ = fi233
                    }
                } else {
                    var v237 float64 = 1e10
                    if v1 {
                        v238 := GenS15{F0: 100, F1: 9223372036854775807}
                        for fi239 := 0; fi239 < 20; fi239++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi239
                        }
                        if false {
                            v1 = v1
                            var v240 int = (-(v59))
                            _ = v240
                        } else {
                            _ = len(v3)
                        }
                        _ = v238
                    } else {
                        v241 := GenS32{F0: 9223372036854775807, F1: 42}
                        for fi242 := 0; fi242 < 20; fi242++ {
                            fmt.Println(fi242, v44, v59)
                            _ = fi242
                        }
                        v1 = true
                        _ = v241
                    }
                    _ = v237
                }
            case "a":
                var v243 string = v46
                _ = v243
            default:
                var v244 bool = v1
                _ = v244
            }
            for fi245 := 0; fi245 < 20; fi245++ {
                v46 = "hello"
                switch (v2 + ("a" + "a")) {
                case "a":
                    switch (("a" + v2) + "") {
                    case "":
                        v1 = v1
                    case "hello":
                        v46 = ("\n\t" + "a")
                    case "a":
                        v44 = v46
                    }
                }
                _ = fi245
            }
            v3 = append(v3, 2147483648)
            _ = fi60
        }
    } else {
        for fi246 := 0; fi246 < 14; fi246++ {
            if (true || true) {
                var v247 float32 = (1e10 + 0.1)
                v59 += v59
                v1 = false
                _ = v247
            }
            _ = fi246
        }
        if v1 {
            for fi248 := 0; fi248 < 11; fi248++ {
                for fi249 := 0; fi249 < 9; fi249++ {
                    fmt.Println(v44)
                    switch "a" {
                    case "\n\t":
                        v250 := GenS13{F0: 0, F1: 2, F2: 1}
                        v251 := GenS6{F0: v1, F1: fi248}
                        _ = len(v3)
                        _ = v250
                        _ = v251
                    case "":
                        v252 := GenS8{F0: v2, F1: 2147483648, F2: 1}
                        _ = v252
                    default:
                        v44 = ((("foo bar" + v44) + ("a" + ("hello" + v2))) + "")
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi249
                }
                v2 = ((((v2 + "") + ("\n\t" + v44)) + "\n\t") + (v46 + v46))
                _ = fi248
            }
        }
    }
    if v1 {
        var v253 int = genFn0(2147483648, 4294967295)
        v253 = v253
        var v254 uint64 = 0
        _ = v253
        _ = v254
    } else {
        if (!(1 == 1)) {
            var v255 bool = (0 != 4294967295)
            if v255 {
                var v256 uint16 = 420
                for fi257 := 0; fi257 < 8; fi257++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    for fi258 := 0; fi258 < 14; fi258++ {
                        v259 := GenS33{F0: 1, F1: (v256 * 0), F2: -1}
                        var v260 float32 = 1e10
                        v261 := GenS18{F0: 1, F1: 2147483647}
                        _ = fi258
                        _ = v259
                        _ = v260
                        _ = v261
                    }
                    _ = fi257
                }
                var v262 int32 = 0
                _ = v256
                _ = v262
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v255
        }
        if (!false) {
            if true {
                v3 = append(v3, 42)
                switch "foo bar" {
                case "a":
                    if v1 {
                        if (592 <= 1) {
                            if (65535 < 70) {
                                v44 = v2
                                if (!(!v1)) {
                                    fmt.Println(v1)
                                }
                                for fi263 := 0; fi263 < 2; fi263++ {
                                    v264 := GenS8{F0: ((v46 + (("a" + v2) + v2)) + (((v2 + v2) + ((((v2 + v44) + v2) + v2) + "a")) + ("\n\t" + (((v46 + "hello") + v2) + v44)))), F1: 4294967295, F2: 2}
                                    switch v59 {
                                    case 902:
                                        v265 := GenS34{F0: 128, F1: 18446744073709551615}
                                        _ = v265
                                    case -9223372036854775808:
                                        v59 = v59
                                    case 32768:
                                        for fi266 := 0; fi266 < 19; fi266++ {
                                            fmt.Println(fi263)
                                            _ = fi266
                                        }
                                    default:
                                        v267 := GenS2{F0: -1}
                                        _ = v267
                                    }
                                    _ = fi263
                                    _ = v264
                                }
                            } else {
                                v2 = ((v44 + (v46 + "\n\t")) + v2)
                            }
                            switch "a" {
                            case "foo bar":
                                var v268 int = genFn0(2147483647, 2)
                                _ = v268
                            case "a":
                                v1 = (42 < 1)
                            default:
                                var v269 int64 = -580
                                _ = v269
                            }
                        }
                    } else {
                        v59 = (v59 / (((-(v59))) | 1))
                        if (!((0 > 857) || v1)) {
                            if true {
                                for fi270 := 0; fi270 < 17; fi270++ {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    for fi271 := 0; fi271 < 8; fi271++ {
                                        if (-2147483648 >= -606) {
                                            if v1 {
                                                var v272 float32 = float32((-2.25 * 1.5))
                                                _ = v272
                                            } else {
                                                if v1 {
                                                    fmt.Println(v44)
                                                    v3 = append(v3, 0)
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = 2147483648
                                                    }
                                                    var v273 float32 = (((((0.1 + 0.1) * -2.25) * 1.5) * -2.25) + 3.14159)
                                                    switch fi270 {
                                                    case 127:
                                                        v46 = (v46 + "\n\t")
                                                    }
                                                    v274 := GenS1{F0: 1, F1: 2}
                                                    _ = v273
                                                    _ = v274
                                                }
                                                v3 = append(v3, 2147483647)
                                            }
                                            v3 = append(v3, 62)
                                            fmt.Println(v44)
                                        }
                                        v275 := GenS35{F0: 1}
                                        _ = fi271
                                        _ = v275
                                    }
                                    _ = fi270
                                }
                                var v276 uint64 = 4294967295
                                v3 = append(v3, 4294967295)
                                _ = v276
                            }
                        } else {
                            for fi277 := 0; fi277 < 5; fi277++ {
                                v278 := GenS36{F0: (1 == 32767)}
                                for fi279 := 0; fi279 < 19; fi279++ {
                                    fmt.Println(v44, v59)
                                    _ = fi279
                                }
                                _ = fi277
                                _ = v278
                            }
                            var v280 int32 = -1
                            fmt.Println(v59, v280, v46)
                            _ = v280
                        }
                    }
                default:
                    switch v59 {
                    case 128:
                        var v281 int16 = -1
                        _ = v281
                    case 380:
                        if true {
                            v1 = v1
                            v282 := GenS15{F0: 127, F1: -1}
                            _ = v282
                        } else {
                            var v283 uint64 = 42
                            _ = v283
                        }
                    case 9223372036854775807:
                        v3 = append(v3, 2147483647)
                    default:
                        var v284 int32 = -1
                        _ = v284
                    }
                }
                if (1 <= 42) {
                    v285 := GenS26{F0: 127, F1: v1, F2: "hello"}
                    var v286 uint8 = 1
                    _ = v285
                    _ = v286
                }
            }
            if (617 < 0) {
                var v287 uint32 = 0
                v3 = append(v3, 1)
                _ = v287
            }
            if true {
                switch v59 {
                case 32767:
                    if len(v3) > 0 {
                        v3[0] = 614
                    }
                case 417:
                    var v288 int16 = 32767
                    _ = v288
                case 0:
                    if v1 {
                        if (0 > int16(v59)) {
                            if (!false) {
                                switch v46 {
                                case "\n\t":
                                    switch v59 {
                                    case -368:
                                        switch (v46 + "a") {
                                        case "foo bar":
                                            v44 = ("foo bar" + v46)
                                        case "\n\t":
                                            var v289 string = v2
                                            _ = v289
                                        case "":
                                            var v290 uint64 = 1
                                            _ = v290
                                        }
                                    default:
                                        v291 := GenS31{F0: -2147483648, F1: -465, F2: 255}
                                        _ = v291
                                    }
                                }
                                var v292 uint = 65535
                                _ = v292
                            }
                            if len(v3) > 0 {
                                v3[0] = 42
                            }
                            fmt.Println(v44, v2, v2)
                        } else {
                            v46 = v2
                            if v1 {
                                if (true || v1) {
                                    v2 = v46
                                    if v1 {
                                        switch ((v46 + (v44 + "a")) + (v44 + "a")) {
                                        case "\n\t":
                                            v293 := GenS2{F0: 1000}
                                            _ = v293
                                        case "hello":
                                            v1 = v1
                                        case "foo bar":
                                            var v294 int = (v59 ^ (^(v59)))
                                            _ = v294
                                        default:
                                            fmt.Println(v1, v2, v46)
                                        }
                                    }
                                }
                            } else {
                                var v295 uint16 = 311
                                for fi296 := 0; fi296 < 10; fi296++ {
                                    var v297 uint64 = 1
                                    v298 := GenS37{F0: 217}
                                    _ = fi296
                                    _ = v297
                                    _ = v298
                                }
                                _ = v295
                            }
                        }
                        v299 := GenS38{F0: v2}
                        switch (v59 >> ((255) & 63)) {
                        case 32767:
                            fmt.Println(v46, v1, v46)
                        case -432:
                            v300 := GenS29{F0: 1, F1: (v59 >> ((65535) & 63))}
                            _ = v300
                        default:
                            for fi301 := 0; fi301 < 15; fi301++ {
                                v302 := GenS33{F0: 0, F1: 1, F2: 42}
                                fmt.Println(fi301, v2)
                                _ = fi301
                                _ = v302
                            }
                        }
                        _ = v299
                    } else {
                        var v303 uint32 = 2147483648
                        v1 = (v44 <= v46)
                        _ = v303
                    }
                default:
                    var v304 uint64 = 4294967295
                    _ = v304
                }
                v59 = v59
            }
        } else {
            v59 += (^(v59))
            var v305 uint16 = 42
            _ = v305
        }
    }
    for fi306 := 0; fi306 < 6; fi306++ {
        switch ("" + v44) {
        case "foo bar":
            fmt.Println(v44, v2)
        case "a":
            var v307 float32 = -2.25
            _ = v307
        }
        var v308 uint32 = 4294967295
        _ = fi306
        _ = v308
    }
    if len(v3) > 0 {
        v3[0] = 2
    }
    var v309 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v309:", v309)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v309
    _ = v44
    _ = v45
    _ = v46
    _ = v59
}
