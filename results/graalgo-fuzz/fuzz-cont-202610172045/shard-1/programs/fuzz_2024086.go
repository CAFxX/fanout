package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 bool; F1 int; F2 uint32 }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 bool; F1 int32; F2 int32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint16; F1 uint16 }
type GenS10 struct { F0 uint8; F1 uint32 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 bool; F1 int64; F2 int16 }
func genFn0(p0 uint16, p1 int64, p2 int) string {
    p2 = -298
    v5 := GenS1{F0: p2, F1: -729, F2: (p0 - 243)}
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return "foo bar"
}

func main() {
    var v1 uint8 = 255
    var v2 int = 1
    var v3 uint16 = 65535
    var v4 uint32 = 0
    var v6 int8 = 100
    var v7 int16 = -191
    switch (v2 + v2) {
    case 9223372036854775807:
        switch 1 {
        case -316:
            for fi8 := 0; fi8 < 4; fi8++ {
                fmt.Println(v2, v2, v4)
                if false {
                    v2 = (fi8 * v2)
                }
                _ = fi8
            }
        case 1:
            var v9 uint64 = 135
            _ = v9
        }
    case -2147483648:
        var v10 int = (v2 & int(v7))
        _ = v10
    case 444:
        v11 := []int{v2}
        _ = v11
    default:
        v12 := GenS2{F0: v6}
        _ = v12
    }
    var v13 string = "a"
    fmt.Println(v13)
    if false {
        var v14 uint8 = 128
        v4 -= v4
        _ = v14
    } else {
        v6 -= v6
        fmt.Println(v13, v6, v6)
    }
    v15 := GenS2{F0: -128}
    for fi16 := 0; fi16 < 13; fi16++ {
        if ((v7 / ((v7) | 1)) >= (v7 / ((42) | 1))) {
            for fi17 := 0; fi17 < 6; fi17++ {
                v18 := GenS2{F0: -1}
                var v19 uint = 413
                _ = fi17
                _ = v18
                _ = v19
            }
            var v20 int = 127
            _ = v20
        } else {
            var v21 float64 = 3.14159
            v22 := GenS1{F0: 32768, F1: -299, F2: 42}
            v3 -= v3
            _ = v21
            _ = v22
        }
        if ((v3 & v3) == v3) {
            var v23 uint = 2
            if false {
                if (!true) {
                    v24 := GenS1{F0: int(v4), F1: int32(v3), F2: uint16(v4)}
                    _ = v24
                }
                for fi25 := 0; fi25 < 3; fi25++ {
                    v4 += (v4 | 65535)
                    _ = fi25
                }
            } else {
                if false {
                    switch (((v13 + v13) + ("a" + "\n\t")) + v13) {
                    case "\n\t":
                        var v26 int = v2
                        _ = v26
                    case "":
                        for fi27 := 0; fi27 < 9; fi27++ {
                            if (!true) {
                                var v28 uint16 = (^(v3))
                                var v29 int32 = 42
                                _ = v28
                                _ = v29
                            }
                            v30 := []int{}
                            if (42 > -9223372036854775808) {
                                fmt.Println(v4)
                                if true {
                                    if len(v30) > 0 {
                                        v30[0] = 42
                                    }
                                    if ((v4 & (v4 % ((v4) | 1))) == 42) {
                                        v7 -= (^(v7))
                                        v23 -= v23
                                        var v31 int = 2147483647
                                        _ = v31
                                    }
                                    v2 = fi27
                                    v7 = int16(v6)
                                    if (true && (int(v3) >= v2)) {
                                        v1 = (v1 * (^(v1)))
                                        _ = len(v30)
                                        var v32 uint64 = 0
                                        fmt.Println(v2, v7, v2)
                                        _ = v32
                                    }
                                }
                                switch (v13 + v13) {
                                case "hello":
                                    v33 := GenS3{F0: -2147483648}
                                    _ = v33
                                case "\n\t":
                                    v2 -= (fi27 << (((^(v23))) & 63))
                                case "a":
                                    var v34 int8 = (v6 << ((v23) & 63))
                                    _ = v34
                                default:
                                    fmt.Println(v13, v2, v13)
                                }
                            }
                            _ = len(v30)
                            if true {
                                var v35 float64 = (((0.0 - 1e10) / 1.0) / 0.1)
                                var v36 float32 = 1.5
                                var v37 uint16 = (v3 / (((v3 << ((65535) & 63))) | 1))
                                _ = v35
                                _ = v36
                                _ = v37
                            } else {
                                var v38 uint32 = v4
                                if (!true) {
                                    if true {
                                        v39 := GenS1{F0: v2, F1: -2147483648, F2: (v3 * (v3 / ((v3) | 1)))}
                                        v30 = append(v30, v2)
                                        _ = v39
                                    } else {
                                        var v40 int8 = v6
                                        _ = v40
                                    }
                                    if true {
                                        if len(v30) >= 2 {
                                            v30 = v30[1:len(v30)]
                                        }
                                        var v41 float64 = float64((((float32(0.1) + 1.5) + 0.0) - (1e10 * -2.25)))
                                        v38 -= (v4 | (v38 * 244))
                                        _ = v41
                                    }
                                }
                                _ = v38
                            }
                            var v42 int8 = 1
                            _ = fi27
                            _ = v30
                            _ = v42
                        }
                    }
                } else {
                    if true {
                        for fi43 := 0; fi43 < 13; fi43++ {
                            v44 := GenS4{F0: 1, F1: 0}
                            v23 += 4294967295
                            _ = fi43
                            _ = v44
                        }
                        if ("a" >= v13) {
                            switch "hello" {
                            case "a":
                                switch v2 {
                                case 2147483647:
                                    v45 := GenS2{F0: -1}
                                    _ = v45
                                case 32768:
                                    if ((!true) && true) {
                                        for fi46 := 0; fi46 < 11; fi46++ {
                                            v47 := GenS1{F0: 639, F1: 1, F2: (v3 + v3)}
                                            v48 := []int{-1, fi46, 564}
                                            switch (v2 | v2) {
                                            case -820:
                                                v49 := GenS3{F0: 374}
                                                _ = v49
                                            }
                                            v50 := GenS2{F0: v6}
                                            _ = fi46
                                            _ = v47
                                            _ = v48
                                            _ = v50
                                        }
                                    } else {
                                        var v51 float64 = 0.0
                                        _ = v51
                                    }
                                }
                            case "hello":
                                v1 -= 42
                            }
                        }
                        v52 := GenS1{F0: v2, F1: -1, F2: (v3 - uint16(v23))}
                        _ = v52
                    }
                }
                if false {
                    v4 -= (v4 * 829)
                    for fi53 := 0; fi53 < 18; fi53++ {
                        if false {
                            v54 := GenS5{F0: (-1 > -836), F1: (v2 | (v2 << ((255) & 63))), F2: (v4 / (((v4 * (^(v4)))) | 1))}
                            if ((v23 | 697) < v23) {
                                var v55 uint64 = 42
                                if (2 > (^(v3))) {
                                    if ((!false) || true) {
                                        v56 := []int{fi16, v2}
                                        v57 := GenS2{F0: int8(v7)}
                                        if false {
                                            if len(v56) >= 2 {
                                                v56 = v56[1:len(v56)]
                                            }
                                            if len(v56) > 0 {
                                                v56[0] = fi53
                                            }
                                        }
                                        fmt.Println(fi16)
                                        _ = v56
                                        _ = v57
                                    }
                                    if true {
                                        var v58 int = v2
                                        var v59 float64 = 1e10
                                        _ = v58
                                        _ = v59
                                    }
                                    fmt.Println(fi53, fi16, v4)
                                    var v60 uint64 = v55
                                    v61 := GenS6{F0: v23, F1: ""}
                                    var v62 uint64 = 744
                                    var v63 int16 = (v7 - 1000)
                                    _ = v60
                                    _ = v61
                                    _ = v62
                                    _ = v63
                                } else {
                                    v64 := []int{(v2 * fi16), 4294967295, -681}
                                    var v65 int32 = 617
                                    var v66 int32 = v65
                                    _ = v64
                                    _ = v65
                                    _ = v66
                                }
                                _ = v55
                            } else {
                                var v67 uint16 = v3
                                switch v2 {
                                case 2147483647:
                                    var v68 int8 = (-(v6))
                                    if false {
                                        if true {
                                            v7 += int16(v4)
                                            v23 = (^(v23))
                                        }
                                        var v69 int16 = v7
                                        _ = v69
                                    } else {
                                        var v70 int8 = (v68 + v68)
                                        fmt.Println(v7, v23, v13)
                                        v71 := []int{(-(v2)), -875, 4294967295}
                                        var v72 uint8 = v1
                                        v73 := GenS5{F0: (v72 >= uint8(v4)), F1: fi16, F2: (v4 ^ (^(v4)))}
                                        _ = v70
                                        _ = v71
                                        _ = v72
                                        _ = v73
                                    }
                                    v74 := GenS7{F0: false, F1: -453, F2: 42}
                                    switch "foo bar" {
                                    case "a":
                                        v75 := []int{(fi16 << ((v23) & 63)), 32767, (v2 << (((v23 ^ v23)) & 63))}
                                        _ = v75
                                    case "\n\t":
                                        v23 -= v23
                                        var v76 string = v13
                                        v77 := []int{}
                                        _ = v76
                                        _ = v77
                                    }
                                    var v78 float64 = 3.14159
                                    _ = v68
                                    _ = v74
                                    _ = v78
                                case -106:
                                    var v79 uint8 = v1
                                    _ = v79
                                case 32767:
                                    var v80 uint8 = v1
                                    _ = v80
                                }
                                var v81 float64 = (1e10 + 0.0)
                                _ = v67
                                _ = v81
                            }
                            _ = v54
                        }
                        _ = fi53
                    }
                } else {
                    v2 = int(v7)
                    if true {
                        if false {
                            for fi82 := 0; fi82 < 3; fi82++ {
                                if (v2 == (v2 | fi16)) {
                                    v1 = 107
                                } else {
                                    var v83 float64 = 3.14159
                                    _ = v83
                                }
                                v84 := []int{(fi82 % ((fi82) | 1)), fi16}
                                _ = fi82
                                _ = v84
                            }
                            var v85 uint8 = v1
                            _ = v85
                        }
                        var v86 uint8 = v1
                        _ = v86
                    } else {
                        fmt.Println(v3, fi16)
                        v2 += 32768
                        var v87 int16 = int16(v2)
                        _ = v87
                    }
                    v3 = uint16(v1)
                }
                for fi88 := 0; fi88 < 4; fi88++ {
                    var v89 int32 = 42
                    v1 -= 2
                    _ = fi88
                    _ = v89
                }
            }
            _ = v23
        } else {
            if (uint64(v1) > 1) {
                var v90 uint16 = (v3 + v3)
                var v91 int32 = -809
                fmt.Println(v1)
                _ = v90
                _ = v91
            } else {
                fmt.Println(v3, v1)
                var v92 float64 = 0.1
                v13 = v13
                _ = v92
            }
            var v93 int64 = -9223372036854775808
            _ = v93
        }
        if ((3.14159 >= 0.1) || ((false || false) && ((v7 > v7) || ((((0.0 + ((3.14159 - 1.5) + 1e10)) / 1.0) - 3.14159) == -2.25)))) {
            v94 := []int{(v2 | -369), fi16, (v2 ^ v2)}
            v95 := GenS8{F0: uint(v1)}
            if false {
                switch (fi16 - (fi16 % ((v2) | 1))) {
                case -2147483648:
                    switch (-(v2)) {
                    case 4294967295:
                        if true {
                            v94 = append(v94, v2)
                            v7 -= v7
                        }
                    }
                case -9223372036854775808:
                    var v96 int8 = v6
                    _ = v96
                case -28:
                    _ = len(v94)
                default:
                    v4 = v4
                }
                v6 = v6
                for fi97 := 0; fi97 < 12; fi97++ {
                    var v98 bool = true
                    switch v2 {
                    case -141:
                        if ((((1 <= 1) || v98) || true) || false) {
                            var v99 uint32 = (v4 + v4)
                            var v100 string = (v13 + v13)
                            var v101 float64 = 0.1
                            _ = v100
                            _ = v101
                            _ = v99
                        } else {
                            var v102 float32 = -2.25
                            v103 := GenS9{F0: v3, F1: uint16(v4)}
                            switch (v13 + v13) {
                            case "foo bar":
                                var v104 uint8 = v1
                                _ = v104
                            }
                            v3 = (v3 >> ((18446744073709551615) & 63))
                            var v105 uint32 = (v4 - v4)
                            _ = v102
                            _ = v103
                            _ = v105
                        }
                    case 127:
                        if true {
                            _ = len(v94)
                        }
                    case -36:
                        if len(v94) >= 2 {
                            v94 = v94[1:len(v94)]
                        }
                    default:
                        var v106 int64 = -589
                        _ = v106
                    }
                    var v107 int64 = 1
                    _ = fi97
                    _ = v107
                    _ = v98
                }
            } else {
                v108 := GenS8{F0: 1}
                for fi109 := 0; fi109 < 17; fi109++ {
                    if len(v94) > 0 {
                        v94[0] = (fi16 ^ (fi109 % ((fi16) | 1)))
                    }
                    switch 1 {
                    case 4294967295:
                        if true {
                            v3 -= (v3 ^ v3)
                            if (true && false) {
                                if (fi109 == (fi16 / ((fi109) | 1))) {
                                    fmt.Println(v1)
                                    switch fi16 {
                                    case 127:
                                        var v110 uint16 = (v3 ^ (v3 & (v3 * v3)))
                                        _ = v110
                                    case 0:
                                        fmt.Println(v4, v2)
                                    case -795:
                                        var v111 int = 127
                                        _ = v111
                                    default:
                                        if false {
                                            var v112 int8 = v6
                                            if len(v94) > 0 {
                                                v94[0] = -9223372036854775808
                                            }
                                            _ = v112
                                        }
                                    }
                                    var v113 int64 = 42
                                    _ = v113
                                } else {
                                    var v114 int32 = 2147483647
                                    _ = v114
                                }
                                fmt.Println(v7, v1, v13)
                                v6 = v6
                            } else {
                                var v115 uint = 4294967295
                                _ = v115
                            }
                        } else {
                            switch v13 {
                            case "":
                                v116 := GenS4{F0: 1, F1: -2147483648}
                                _ = v116
                            }
                            v117 := GenS7{F0: (!(v4 >= 42)), F1: -2147483648, F2: -808}
                            var v118 uint = 133
                            _ = v117
                            _ = v118
                        }
                    case 42:
                        if true {
                            if (true || false) {
                                switch (("hello" + v13) + v13) {
                                case "a":
                                    _ = len(v94)
                                case "":
                                    fmt.Println(v6, v2, v13)
                                case "hello":
                                    if false {
                                        if ((42 != 9223372036854775807) || ((-(v6)) >= (v6 - 1))) {
                                            v13 = ("foo bar" + "a")
                                        } else {
                                            if false {
                                                var v119 string = (v13 + (("\n\t" + ("hello" + (v13 + v13))) + ((v13 + v13) + v13)))
                                                if len(v94) >= 2 {
                                                    v94 = v94[1:len(v94)]
                                                }
                                                if true {
                                                    if false {
                                                        var v120 int = (v2 * 1)
                                                        v2 = (v2 ^ fi16)
                                                        fmt.Println(v3, v3, v6)
                                                        _ = v120
                                                    } else {
                                                        var v121 uint16 = (v3 + (^(v3)))
                                                        if false {
                                                            v3 = v121
                                                        }
                                                        var v122 uint8 = v1
                                                        v7 = (v7 + 538)
                                                        var v123 int8 = 127
                                                        if (!true) {
                                                            var v124 uint = 1
                                                            if true {
                                                                switch -165 {
                                                                case 9223372036854775807:
                                                                    v1 -= uint8(v124)
                                                                default:
                                                                    switch v119 {
                                                                    case "\n\t":
                                                                        if len(v94) > 0 {
                                                                            v94[0] = (fi109 >> (((v124 / (((v124 - (^(v124)))) | 1))) & 63))
                                                                        }
                                                                    case "hello":
                                                                        v122 = v1
                                                                    case "":
                                                                        var v125 uint64 = uint64(v124)
                                                                        _ = v125
                                                                    default:
                                                                        v126 := GenS3{F0: 42}
                                                                        _ = v126
                                                                    }
                                                                    var v127 float64 = float64(1e10)
                                                                    v128 := GenS8{F0: uint(v4)}
                                                                    v4 += 335
                                                                    _ = len(v94)
                                                                    _ = v127
                                                                    _ = v128
                                                                }
                                                                if true {
                                                                    if true {
                                                                        var v129 string = "hello"
                                                                        var v130 uint = (v124 ^ v124)
                                                                        _ = v129
                                                                        _ = v130
                                                                    }
                                                                    var v131 uint32 = v4
                                                                    v132 := GenS10{F0: 127, F1: uint32(v122)}
                                                                    v119 = "\n\t"
                                                                    _ = v131
                                                                    _ = v132
                                                                }
                                                                v7 = (v7 >> ((431) & 63))
                                                                v133 := GenS8{F0: v124}
                                                                _ = v133
                                                            }
                                                            _ = v124
                                                        } else {
                                                            var v134 float64 = 1e10
                                                            _ = v134
                                                        }
                                                        _ = v121
                                                        _ = v122
                                                        _ = v123
                                                    }
                                                    switch v119 {
                                                    case "hello":
                                                        var v135 uint = 0
                                                        v1 = v1
                                                        v135 = v135
                                                        _ = v135
                                                    default:
                                                        v136 := GenS1{F0: (-(fi109)), F1: 1, F2: v3}
                                                        _ = v136
                                                    }
                                                    if len(v94) > 0 {
                                                        v94[0] = fi109
                                                    }
                                                    fmt.Println(v7, v6, v3)
                                                } else {
                                                    var v137 bool = false
                                                    if ((v2 != v2) || v137) {
                                                        if (!true) {
                                                            if len(v94) >= 2 {
                                                                v94 = v94[1:len(v94)]
                                                            }
                                                            fmt.Println(v6, v137, v3)
                                                        }
                                                        if (v6 > 0) {
                                                            v2 += fi16
                                                            var v138 int = (fi16 & fi109)
                                                            _ = v138
                                                        } else {
                                                            var v139 int8 = 0
                                                            var v140 uint16 = 65535
                                                            v141 := GenS11{F0: -56}
                                                            _ = v139
                                                            _ = v140
                                                            _ = v141
                                                        }
                                                    }
                                                    if false {
                                                        switch (v2 & (fi16 / ((9223372036854775807) | 1))) {
                                                        case 42:
                                                            var v142 uint8 = v1
                                                            _ = v142
                                                        case 4294967295:
                                                            _ = len(v94)
                                                        case -644:
                                                            if v137 {
                                                                fmt.Println(v3, v6, v6)
                                                            } else {
                                                                _ = len(v94)
                                                                v143 := GenS11{F0: (v6 * 42)}
                                                                v94 = append(v94, int(v6))
                                                                _ = v143
                                                            }
                                                            var v144 int = v2
                                                            v1 = (v1 * 255)
                                                            _ = v144
                                                        }
                                                        var v145 float64 = 1e10
                                                        v94 = append(v94, int(v4))
                                                        _ = v145
                                                    }
                                                    _ = v137
                                                }
                                                if true {
                                                    fmt.Println(v1, v4, v119)
                                                    v94 = append(v94, (fi109 & 1))
                                                    var v146 uint = 520
                                                    if true {
                                                        switch (v119 + "foo bar") {
                                                        case "a":
                                                            switch ("foo bar" + v119) {
                                                            case "hello":
                                                                if ((v7 / ((0) | 1)) > 525) {
                                                                    fmt.Println(fi16)
                                                                } else {
                                                                    if (1.5 < (0.0 + (1e10 * 1.5))) {
                                                                        var v147 uint32 = (v4 - 416)
                                                                        var v148 bool = false
                                                                        _ = v147
                                                                        _ = v148
                                                                    }
                                                                    v2 = fi16
                                                                    var v149 int64 = -9223372036854775808
                                                                    _ = v149
                                                                }
                                                            default:
                                                                var v150 uint64 = 0
                                                                _ = v150
                                                            }
                                                        default:
                                                            v1 += 43
                                                        }
                                                        var v151 string = v13
                                                        _ = v151
                                                    }
                                                    _ = v146
                                                } else {
                                                    if len(v94) >= 2 {
                                                        v94 = v94[1:len(v94)]
                                                    }
                                                }
                                                var v152 uint64 = 739
                                                _ = v119
                                                _ = v152
                                            } else {
                                                v6 = (^(v6))
                                                var v153 uint32 = v4
                                                _ = v153
                                            }
                                            switch v13 {
                                            case "\n\t":
                                                var v154 string = v13
                                                _ = v154
                                            default:
                                                var v155 bool = ((!(!false)) && (true && false))
                                                _ = v155
                                            }
                                        }
                                        var v156 uint = 2
                                        _ = v156
                                    }
                                }
                                if true {
                                    var v157 float32 = 0.1
                                    if (47 <= (v6 | (v6 << ((18446744073709551615) & 63)))) {
                                        var v158 uint8 = 1
                                        _ = v158
                                    } else {
                                        _ = len(v94)
                                        switch fi109 {
                                        case -9223372036854775808:
                                            if (3.14159 >= -2.25) {
                                                v13 = "foo bar"
                                                var v159 uint16 = v3
                                                _ = v159
                                            } else {
                                                switch fi109 {
                                                case 42:
                                                    switch -1 {
                                                    case 128:
                                                        if false {
                                                            var v160 float64 = (float64(v157) - -2.25)
                                                            v161 := GenS12{F0: 1}
                                                            v6 = v6
                                                            _ = v160
                                                            _ = v161
                                                        } else {
                                                            if true {
                                                                fmt.Println(v6, v7)
                                                                switch (fi109 * 32768) {
                                                                case 127:
                                                                    var v162 uint32 = v4
                                                                    _ = v162
                                                                case -755:
                                                                    var v163 uint64 = 4294967295
                                                                    _ = v163
                                                                default:
                                                                    v3 += 42
                                                                }
                                                            } else {
                                                                var v164 float32 = ((v157 - v157) - ((float32(0.0) - v157) * ((v157 + float32(3.14159)) / float32(float64(v157)))))
                                                                if len(v94) > 0 {
                                                                    v94[0] = fi109
                                                                }
                                                                fmt.Println(v2)
                                                                _ = v164
                                                            }
                                                        }
                                                        switch (v2 & fi16) {
                                                        case 692:
                                                            if len(v94) > 0 {
                                                                v94[0] = fi16
                                                            }
                                                        case -2147483648:
                                                            var v165 uint = 18446744073709551615
                                                            _ = v165
                                                        case 32768:
                                                            if true {
                                                                _ = len(v94)
                                                                switch v13 {
                                                                case "":
                                                                    fmt.Println(v7, v1, v13)
                                                                }
                                                                var v166 uint8 = v1
                                                                var v167 bool = (35 < uint64(v4))
                                                                _ = v166
                                                                _ = v167
                                                            }
                                                        }
                                                        var v168 int32 = 884
                                                        _ = v168
                                                    default:
                                                        v169 := GenS12{F0: 42}
                                                        _ = v169
                                                    }
                                                default:
                                                    _ = len(v94)
                                                }
                                                fmt.Println(v4, v1)
                                                v170 := GenS7{F0: (!(0 >= 0)), F1: 42, F2: 0}
                                                _ = v170
                                            }
                                        case 0:
                                            v3 = (v3 * 65535)
                                        case 32767:
                                            if (1 <= 18446744073709551615) {
                                                if len(v94) >= 2 {
                                                    v94 = v94[1:len(v94)]
                                                }
                                            } else {
                                                v171 := GenS12{F0: 765}
                                                _ = v171
                                            }
                                        }
                                    }
                                    v94 = append(v94, fi109)
                                    if len(v94) > 0 {
                                        v94[0] = (fi16 ^ (fi109 & 291))
                                    }
                                    _ = v157
                                }
                            } else {
                                fmt.Println(v1, fi16, v7)
                            }
                            var v172 int8 = (v6 + (v6 & v6))
                            v6 += -111
                            _ = v172
                        } else {
                            v173 := GenS13{F0: false, F1: 9223372036854775807, F2: -32768}
                            _ = v173
                        }
                    case 0:
                        if false {
                            v3 = 1
                            var v174 float64 = float64(0.1)
                            fmt.Println(v3, v174, v3)
                            _ = v174
                        }
                    default:
                        var v175 float64 = -2.25
                        _ = v175
                    }
                    _ = fi109
                }
                _ = v108
            }
            _ = v94
            _ = v95
        }
        _ = fi16
    }
    var v176 int32 = -771
    v177 := []int{127, v2, v2}
    v2 += 0
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v176:", v176)
    fmt.Println("slicelen v177:", len(v177), v177)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v15
    _ = v176
    _ = v177
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
