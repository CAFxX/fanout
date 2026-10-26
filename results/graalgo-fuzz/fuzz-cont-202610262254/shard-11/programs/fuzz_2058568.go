package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 int8; F2 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint16; F1 int64 }
type GenS9 struct { F0 int8; F1 string; F2 uint8 }
type GenS10 struct { F0 bool; F1 int; F2 int32 }
type GenS11 struct { F0 int; F1 uint64 }
type GenS12 struct { F0 uint16; F1 int }
type GenS13 struct { F0 uint32; F1 uint32; F2 bool }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int32; F1 uint; F2 uint64 }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 int32; F1 uint32 }
type GenS18 struct { F0 uint; F1 int8; F2 uint16 }
type GenS19 struct { F0 uint16; F1 int; F2 int64 }
type GenS20 struct { F0 uint8; F1 uint8 }
type GenS21 struct { F0 int16 }
type GenS22 struct { F0 uint; F1 bool; F2 uint8 }
type GenS23 struct { F0 string; F1 uint32 }
func genFn0(p0 int64, p1 int8, p2 string) int32 {
    v7 := GenS1{F0: (p2 + p2)}
    var v8 int64 = (p0 / ((p0) | 1))
    var v9 bool = (true && ((4294967295 >= 304) || false))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 708
}
func genFn1() uint64 {
    var v10 float32 = (0.0 / 1.0)
    v11 := []int{32768}
    v10 = v10
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 uint32 = 666
    var v2 uint8 = 0
    v3 := make(map[uint32]string)
    var v4 bool = true
    var v5 bool = (!v4)
    var v6 int8 = 95
    fmt.Println(v6)
    for fi12 := 0; fi12 < 17; fi12++ {
        fmt.Println(v5)
        var v13 uint = 1
        _ = fi12
        _ = v13
    }
    var v14 float32 = (-2.25 + 1e10)
    switch -1 {
    case -2147483648:
        switch ("foo bar" + "hello") {
        case "hello":
            if v4 {
                if (-2147483648 < -1) {
                    v14 += 3.14159
                    v4 = v5
                    switch 128 {
                    case 9223372036854775807:
                        if ((("" + ("" + "a")) + "\n\t") != "\n\t") {
                            var v15 float64 = ((0.1 * ((0.0 * 3.14159) * 0.1)) + float64(v14))
                            if false {
                                v16 := GenS1{F0: ("" + "")}
                                var v17 int = 32767
                                var v18 int16 = 1
                                _ = v16
                                _ = v17
                                _ = v18
                            } else {
                                var v19 int8 = v6
                                v6 -= v19
                                _ = v19
                            }
                            _ = v15
                        }
                    case -1:
                        if v4 {
                            fmt.Println(v4)
                            v6 = v6
                        }
                    case 32768:
                        switch ("a" + (("foo bar" + (("\n\t" + ("hello" + "foo bar")) + ("\n\t" + "foo bar"))) + "")) {
                        case "hello":
                            var v20 string = "\n\t"
                            _ = v20
                        }
                    }
                } else {
                    v4 = (!v4)
                    fmt.Println(v14, v2)
                }
                v1 = (v1 - v1)
            } else {
                var v21 int64 = -1
                _ = v21
            }
        case "\n\t":
            if (!(v4 && (!v5))) {
                for fi22 := 0; fi22 < 19; fi22++ {
                    v2 += v2
                    for fi23 := 0; fi23 < 8; fi23++ {
                        if true {
                            v24 := GenS2{F0: "", F1: ("\n\t" + "hello")}
                            if (((-271 < 9223372036854775807) && (!v5)) && v4) {
                                var v25 uint8 = (v2 % (((^(v2))) | 1))
                                _ = v25
                            }
                            var v26 int = fi22
                            _ = v24
                            _ = v26
                        }
                        var v27 uint = 106
                        _ = fi23
                        _ = v27
                    }
                    _ = fi22
                }
                for fi28 := 0; fi28 < 4; fi28++ {
                    if v4 {
                        v14 += float32(1e10)
                        fmt.Println(v5, v4, v1)
                        var v29 int16 = 1
                        _ = v29
                    }
                    v30 := GenS1{F0: "hello"}
                    _ = fi28
                    _ = v30
                }
            } else {
                for fi31 := 0; fi31 < 15; fi31++ {
                    fmt.Println(fi31, v5)
                    _ = fi31
                }
                var v32 uint64 = 0
                _ = v32
            }
        default:
            var v33 float64 = float64(v14)
            _ = v33
        }
    case 127:
        for fi34 := 0; fi34 < 7; fi34++ {
            fmt.Println(v5)
            _ = fi34
        }
    default:
        if (uint16(v1) > 393) {
            v35 := GenS3{F0: 1, F1: v2}
            if (true || (!false)) {
                _ = len(v3)
                for fi36 := 0; fi36 < 15; fi36++ {
                    v37 := GenS4{F0: (^(v2))}
                    _ = fi36
                    _ = v37
                }
                v4 = (v4 || v5)
            }
            if false {
                var v38 int64 = -9223372036854775808
                _ = v38
            } else {
                if ((!true) && (!false)) {
                    v39 := GenS1{F0: "foo bar"}
                    if v5 {
                        switch 2147483647 {
                        case -1:
                            fmt.Println(v4, v14)
                        case 32767:
                            delete(v3, (^(v1)))
                        }
                        fmt.Println(v4, v4)
                        fmt.Println(v14)
                    } else {
                        var v40 int = -278
                        var v41 int = v40
                        v42 := GenS5{F0: 127, F1: v6, F2: genFn1()}
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                    var v43 int = -1
                    _ = v39
                    _ = v43
                }
                var v44 int16 = 32767
                _ = v44
            }
            _ = v35
        } else {
            switch "" {
            case "hello":
                v45 := GenS5{F0: 1, F1: v6, F2: 18446744073709551615}
                _ = v45
            }
            if v5 {
                fmt.Println(v6, v2)
                fmt.Println(v5, v5, v14)
                v46 := GenS6{F0: 1}
                _ = v46
            } else {
                switch ("foo bar" + "a") {
                case "":
                    for fi47 := 0; fi47 < 4; fi47++ {
                        _ = len(v3)
                        if v5 {
                            var v48 uint8 = (v2 & v2)
                            _ = v48
                        }
                        v49 := GenS3{F0: 457, F1: v2}
                        _ = fi47
                        _ = v49
                    }
                case "a":
                    var v50 int16 = 1000
                    _ = v50
                }
                switch "hello" {
                case "":
                    switch 9223372036854775807 {
                    case -725:
                        v6 -= 127
                    case 4294967295:
                        v2 = 100
                    case 32768:
                        var v51 uint32 = 847
                        _ = v51
                    default:
                        v52 := GenS7{F0: 56}
                        _ = v52
                    }
                case "\n\t":
                    var v53 uint = 18446744073709551615
                    _ = v53
                default:
                    v14 = ((v14 / ((((((3.14159 - (v14 * (v14 * v14))) * (((-2.25 * 1e10) + v14) + v14)) - v14) / (v14 - float32((0.0 / 1.0)))) + -2.25) - v14)) + (0.1 * v14))
                }
            }
        }
    }
    for fi54 := 0; fi54 < 3; fi54++ {
        var v55 int32 = -2147483648
        v2 = (^(v2))
        _ = fi54
        _ = v55
    }
    for fi56 := 0; fi56 < 17; fi56++ {
        v5 = ((!v4) || (18 == 1))
        if false {
            v4 = true
            if (-32768 < 32767) {
                if v4 {
                    for fi57 := 0; fi57 < 10; fi57++ {
                        switch fi57 {
                        case -932:
                            switch -2147483648 {
                            case -1:
                                var v58 uint = 4294967295
                                _ = v58
                            }
                        case 32768:
                            v59 := []int{}
                            _ = v59
                        }
                        _ = fi57
                    }
                    v5 = (-9223372036854775808 <= 9223372036854775807)
                    var v60 int64 = 1
                    _ = v60
                } else {
                    v3[370] = ""
                }
                if (((((!(222 != (v1 >> ((18446744073709551615) & 63)))) && (v5 && (false && v5))) || v5) || (fi56 < int(v2))) && (-1 > 0)) {
                    for fi61 := 0; fi61 < 4; fi61++ {
                        v5 = ((v5 || (v1 <= (v1 & (^(v1))))) && true)
                        var v62 int8 = v6
                        switch "a" {
                        case "a":
                            switch 1 {
                            case -1:
                                v5 = v5
                            case 42:
                                if v4 {
                                    if v4 {
                                        var v63 uint8 = (v2 & v2)
                                        v6 = 42
                                        v1 = 582
                                        _ = v63
                                    } else {
                                        if v4 {
                                            v2 += v2
                                            v14 = (v14 / v14)
                                            switch -441 {
                                            case 1:
                                                switch 743 {
                                                case 2147483647:
                                                    v64 := GenS8{F0: 0, F1: -51}
                                                    _ = v64
                                                case 0:
                                                    v65 := GenS1{F0: "hello"}
                                                    _ = v65
                                                case 219:
                                                    if v5 {
                                                        var v66 uint32 = 4294967295
                                                        v67 := v3[0]
                                                        _ = v66
                                                        _ = v67
                                                    }
                                                default:
                                                    v68 := []int{(fi61 * (fi56 + (fi56 * fi56))), fi61}
                                                    _ = v68
                                                }
                                            case 42:
                                                v69 := v3[65535]
                                                _ = v69
                                            }
                                        }
                                        v2 = (v2 * 8)
                                        if (v5 && (0 == 2147483647)) {
                                            fmt.Println(fi61, fi61)
                                            switch ("hello" + ("\n\t" + ("a" + "a"))) {
                                            case "foo bar":
                                                v2 -= v2
                                            case "a":
                                                var v70 string = ""
                                                _ = v70
                                            default:
                                                _ = len(v3)
                                            }
                                            if (42 == int32(v2)) {
                                                v62 -= 127
                                                v71 := GenS7{F0: 486}
                                                v2 = (v2 ^ (v2 >> ((362) & 63)))
                                                if v4 {
                                                    if (v5 && true) {
                                                        if (int16(fi56) > 42) {
                                                            v72 := v3[(v1 & (v1 >> ((281) & 63)))]
                                                            if v5 {
                                                                fmt.Println(v6, v4, v1)
                                                                if (2 > 1) {
                                                                    v6 = v62
                                                                    v73 := GenS9{F0: v62, F1: "", F2: v2}
                                                                    fmt.Println(v2, v6)
                                                                    _ = v73
                                                                }
                                                                v14 -= v14
                                                            } else {
                                                                switch "hello" {
                                                                case "a":
                                                                    v5 = v5
                                                                default:
                                                                    if (false || (42 > 42)) {
                                                                        v3[(v1 - v1)] = v72
                                                                        v74 := GenS9{F0: (^(v6)), F1: v72, F2: (v2 + (v2 * (v2 * 42)))}
                                                                        _ = v74
                                                                    } else {
                                                                        v75 := []int{(fi61 >> ((18446744073709551615) & 63)), 127, (fi56 + fi56), int(v1)}
                                                                        _ = v75
                                                                    }
                                                                }
                                                                fmt.Println(v62)
                                                                v76 := GenS8{F0: 771, F1: -164}
                                                                _ = v76
                                                            }
                                                            v77 := []int{(fi61 ^ 4294967295), fi56, 32768}
                                                            switch (v72 + (v72 + ((v72 + v72) + v72))) {
                                                            case "foo bar":
                                                                v5 = v4
                                                                if (!(1e10 == 3.14159)) {
                                                                    var v78 uint8 = v2
                                                                    if v4 {
                                                                        v79 := GenS8{F0: 1, F1: 42}
                                                                        _ = v79
                                                                    }
                                                                    v1 = (^(v1))
                                                                    _ = v78
                                                                } else {
                                                                    if ((v62 != (v6 << ((65535) & 63))) && false) {
                                                                        if v4 {
                                                                            var v80 uint = 317
                                                                            fmt.Println(fi61)
                                                                            _ = v80
                                                                        }
                                                                    } else {
                                                                        v81 := GenS10{F0: (int16(v6) < -32768), F1: fi56, F2: -624}
                                                                        v82 := v3[(v1 % ((102) | 1))]
                                                                        _ = v81
                                                                        _ = v82
                                                                    }
                                                                    v83 := GenS11{F0: fi61, F1: 668}
                                                                    _ = v83
                                                                }
                                                                var v84 int8 = (v62 >> ((255) & 63))
                                                                _ = v84
                                                            case "a":
                                                                v85 := GenS6{F0: 42}
                                                                _ = v85
                                                            default:
                                                                if v5 {
                                                                    v2 -= (v2 >> ((103) & 63))
                                                                    v62 += (-(v62))
                                                                    v86 := GenS1{F0: v72}
                                                                    _ = v86
                                                                }
                                                            }
                                                            switch -9223372036854775808 {
                                                            case 0:
                                                                switch 127 {
                                                                case -2147483648:
                                                                    var v87 int32 = 42
                                                                    _ = v87
                                                                case 4294967295:
                                                                    v88 := GenS12{F0: 432, F1: (fi56 + (fi56 / ((fi61) | 1)))}
                                                                    _ = v88
                                                                }
                                                            case 128:
                                                                v89 := GenS13{F0: (^(v1)), F1: (v1 | 890), F2: (v5 && (142 > (v2 * v2)))}
                                                                _ = v89
                                                            default:
                                                                var v90 int64 = -9223372036854775808
                                                                fmt.Println(v72, v4)
                                                                switch fi56 {
                                                                case -2147483648:
                                                                    v91 := GenS14{F0: 42}
                                                                    _ = v91
                                                                }
                                                                delete(v3, (^(v1)))
                                                                _ = v90
                                                            }
                                                            var v92 float32 = ((v14 * -2.25) + v14)
                                                            _ = v72
                                                            _ = v77
                                                            _ = v92
                                                        }
                                                    }
                                                    v93 := GenS5{F0: (fi61 | (fi61 / ((fi61) | 1))), F1: int8(v1), F2: 19}
                                                    v94 := GenS12{F0: 849, F1: (^(fi56))}
                                                    if v5 {
                                                        switch -869 {
                                                        case -1:
                                                            fmt.Println(v1, v14)
                                                        case 32767:
                                                            fmt.Println(v5)
                                                        default:
                                                            v95 := GenS10{F0: v4, F1: -1, F2: -2147483648}
                                                            _ = v95
                                                        }
                                                        fmt.Println(v6, v1)
                                                    }
                                                    _ = v93
                                                    _ = v94
                                                }
                                                _ = len(v3)
                                                _ = v71
                                            }
                                        }
                                    }
                                    if v5 {
                                        v96 := GenS15{F0: -882, F1: 255, F2: 397}
                                        _ = v96
                                    } else {
                                        var v97 uint16 = 65535
                                        v98 := GenS8{F0: (v97 >> ((738) & 63)), F1: 42}
                                        v99 := v3[v1]
                                        _ = v97
                                        _ = v98
                                        _ = v99
                                    }
                                    switch "hello" {
                                    case "a":
                                        if ("" < "hello") {
                                            v100 := []int{fi56, int(v62), -9223372036854775808}
                                            var v101 bool = v4
                                            if (v1 != (v1 + (v1 & v1))) {
                                                v102 := GenS6{F0: 1}
                                                v2 += v2
                                                _ = v102
                                            }
                                            fmt.Println(v5, v6)
                                            switch ("hello" + "hello") {
                                            case "hello":
                                                v103 := GenS5{F0: int(v1), F1: (v62 ^ v62), F2: 0}
                                                _ = v103
                                            case "":
                                                _ = len(v3)
                                                switch (fi56 >> ((uint(v1)) & 63)) {
                                                case 238:
                                                    if len(v100) >= 2 {
                                                        v100 = v100[1:len(v100)]
                                                    }
                                                }
                                            }
                                            v3[v1] = "a"
                                            fmt.Println(fi61)
                                            fmt.Println(v62)
                                            if ((-320 != -125) || v5) {
                                                var v104 uint8 = 0
                                                if (v14 > ((v14 / (float32((3.14159 - 3.14159)) + v14)) + 1.5)) {
                                                    v14 = v14
                                                    v105 := GenS6{F0: 0}
                                                    var v106 uint = 0
                                                    _ = v105
                                                    _ = v106
                                                }
                                                _ = v104
                                            } else {
                                                v107 := GenS16{F0: 483}
                                                fmt.Println(v6, v5, v5)
                                                _ = v107
                                            }
                                            _ = v100
                                            _ = v101
                                        } else {
                                            var v108 int32 = 973
                                            _ = v108
                                        }
                                    case "\n\t":
                                        var v109 int = fi56
                                        _ = v109
                                    case "foo bar":
                                        switch "foo bar" {
                                        case "":
                                            var v110 uint64 = 324
                                            _ = v110
                                        case "\n\t":
                                            var v111 bool = true
                                            _ = v111
                                        }
                                    default:
                                        if v5 {
                                            v4 = (0 <= -679)
                                        } else {
                                            switch fi56 {
                                            case 1:
                                                v112 := []int{fi56, fi56, 1, -262}
                                                _ = v112
                                            case 29:
                                                v113 := GenS7{F0: 238}
                                                _ = v113
                                            }
                                            var v114 int8 = v62
                                            var v115 uint8 = 1
                                            _ = v114
                                            _ = v115
                                        }
                                    }
                                    var v116 uint8 = 2
                                    var v117 int64 = -9223372036854775808
                                    _ = v116
                                    _ = v117
                                }
                            case 0:
                                v3[65535] = "foo bar"
                            default:
                                v118 := GenS17{F0: -979, F1: v1}
                                _ = v118
                            }
                        }
                        _ = fi61
                        _ = v62
                    }
                    v4 = (-2.25 == (0.1 - (0.0 + 3.14159)))
                } else {
                    var v119 uint = 1
                    _ = v119
                }
                if (966 <= -9223372036854775808) {
                    v6 += (v6 * (^(v6)))
                    v120 := GenS18{F0: 0, F1: v6, F2: 12}
                    _ = v120
                } else {
                    v2 -= (v2 & 2)
                    if v4 {
                        if (2147483648 > v1) {
                            var v121 int16 = 32767
                            _ = v121
                        }
                        for fi122 := 0; fi122 < 20; fi122++ {
                            v123 := GenS19{F0: 0, F1: 79, F2: -924}
                            _ = fi122
                            _ = v123
                        }
                        v124 := GenS4{F0: (v2 * (v2 % ((127) | 1)))}
                        _ = v124
                    } else {
                        v5 = (v1 > (v1 | (v1 & 652)))
                        var v125 uint8 = v2
                        var v126 uint16 = 65535
                        _ = v125
                        _ = v126
                    }
                }
            }
        } else {
            v127 := GenS17{F0: 2147483647, F1: 2147483647}
            var v128 int8 = (v6 * v6)
            v14 = 0.1
            _ = v127
            _ = v128
        }
        _ = fi56
    }
    var v129 bool = (!(42 <= 0))
    var v130 int16 = 1
    var v131 uint8 = v2
    v14 = (1e10 * 1.5)
    v5 = v129
    for fi132 := 0; fi132 < 4; fi132++ {
        var v133 int8 = (v6 * v6)
        v4 = false
        var v134 int8 = (v6 % ((v133) | 1))
        _ = fi132
        _ = v133
        _ = v134
    }
    v135 := []int{2147483647, -9223372036854775808, -341}
    switch 136 {
    case 32768:
        var v136 int32 = 2147483647
        _ = v136
    case 1:
        for fi137 := 0; fi137 < 7; fi137++ {
            var v138 int16 = (v130 & v130)
            for fi139 := 0; fi139 < 3; fi139++ {
                delete(v3, v1)
                var v140 bool = ("\n\t" > "\n\t")
                if v140 {
                    v14 = 0.1
                }
                v141 := GenS20{F0: v131, F1: v2}
                _ = fi139
                _ = v140
                _ = v141
            }
            if v5 {
                var v142 int64 = -1
                var v143 uint = 18446744073709551615
                _ = v142
                _ = v143
            }
            _ = fi137
            _ = v138
        }
    case -680:
        var v144 int64 = 413
        _ = v144
    default:
        v135 = append(v135, -1)
    }
    v135 = append(v135, 37)
    if v129 {
        for fi145 := 0; fi145 < 19; fi145++ {
            if (1 == 2147483647) {
                switch "hello" {
                case "":
                    fmt.Println(v131, v131, v14)
                default:
                    switch (fi145 | (fi145 ^ fi145)) {
                    case 0:
                        for fi146 := 0; fi146 < 13; fi146++ {
                            var v147 string = ("a" + ("\n\t" + ""))
                            fmt.Println(v4, v14)
                            _ = fi146
                            _ = v147
                        }
                    }
                }
                for fi148 := 0; fi148 < 16; fi148++ {
                    var v149 bool = true
                    v150 := GenS12{F0: 255, F1: (-(fi145))}
                    _ = fi148
                    _ = v149
                    _ = v150
                }
                fmt.Println(v2, v129)
            }
            _ = fi145
        }
    } else {
        for fi151 := 0; fi151 < 14; fi151++ {
            fmt.Println(v4, v14, v131)
            for fi152 := 0; fi152 < 19; fi152++ {
                switch (^(fi151)) {
                case 952:
                    if len(v135) >= 2 {
                        v135 = v135[1:len(v135)]
                    }
                case 401:
                    var v153 uint64 = 4294967295
                    if len(v135) >= 2 {
                        v135 = v135[1:len(v135)]
                    }
                    switch ("\n\t" + "foo bar") {
                    case "hello":
                        switch (fi151 >> ((0) & 63)) {
                        case 127:
                            v154 := GenS16{F0: 42}
                            _ = v154
                        case -1:
                            var v155 uint8 = v131
                            _ = v155
                        case 2147483647:
                            if (1 < 404) {
                                v156 := GenS11{F0: fi152, F1: v153}
                                _ = v156
                            } else {
                                _ = len(v135)
                                v157 := v3[2]
                                _ = v157
                            }
                        }
                    case "":
                        switch "\n\t" {
                        case "a":
                            switch -2147483648 {
                            case 4294967295:
                                if true {
                                    var v158 int8 = v6
                                    v135 = append(v135, fi152)
                                    switch (fi152 << ((255) & 63)) {
                                    case 127:
                                        v159 := GenS10{F0: (!false), F1: fi151, F2: 42}
                                        _ = v159
                                    case 2147483647:
                                        v160 := GenS10{F0: (!v129), F1: -9223372036854775808, F2: 2147483647}
                                        if (!v129) {
                                            delete(v3, uint32(v131))
                                            var v161 float64 = 0.1
                                            _ = v161
                                        }
                                        _ = v160
                                    case 0:
                                        if (v129 && v4) {
                                            if len(v135) >= 2 {
                                                v135 = v135[1:len(v135)]
                                            }
                                            switch (fi151 % (((fi152 - -748)) | 1)) {
                                            case 465:
                                                switch "" {
                                                case "\n\t":
                                                    var v162 int32 = genFn0(-327, (v6 / ((-8) | 1)), ("hello" + "\n\t"))
                                                    _ = v162
                                                }
                                            }
                                        }
                                    }
                                    var v163 string = ("" + ("\n\t" + "\n\t"))
                                    _ = v158
                                    _ = v163
                                }
                            case -9223372036854775808:
                                if v129 {
                                    fmt.Println(v1, v130, fi151)
                                    var v164 uint64 = uint64(v1)
                                    _ = v164
                                }
                                var v165 uint32 = v1
                                _ = len(v135)
                                _ = v165
                            case -1:
                                v14 += v14
                            default:
                                switch 873 {
                                case 0:
                                    switch fi152 {
                                    case -330:
                                        if (true || (v129 && v4)) {
                                            v5 = false
                                            v166 := GenS21{F0: v130}
                                            var v167 int16 = int16(v1)
                                            var v168 uint16 = uint16(v2)
                                            _ = v166
                                            _ = v167
                                            _ = v168
                                        } else {
                                            v3[(^(v1))] = "a"
                                            v169 := GenS5{F0: 32768, F1: v6, F2: v153}
                                            var v170 uint64 = (v153 >> ((1) & 63))
                                            var v171 float32 = v14
                                            v153 = genFn1()
                                            _ = v169
                                            _ = v170
                                            _ = v171
                                        }
                                    }
                                case 9223372036854775807:
                                    var v172 int64 = -9223372036854775808
                                    _ = v172
                                case 32768:
                                    v130 = 32767
                                }
                            }
                        case "\n\t":
                            if v5 {
                                fmt.Println(v14, v4, v153)
                                var v173 bool = v4
                                _ = v173
                            }
                        case "":
                            v174 := GenS18{F0: 676, F1: v6, F2: 2}
                            _ = v174
                        }
                    default:
                        fmt.Println(fi152, v6)
                    }
                    var v175 bool = (!v5)
                    _ = v153
                    _ = v175
                case 32767:
                    var v176 float64 = float64((v14 / float32((1e10 + 3.14159))))
                    if v4 {
                        switch "\n\t" {
                        case "a":
                            v177 := GenS10{F0: ((-436 < -9223372036854775808) && v4), F1: fi152, F2: 2147483647}
                            _ = v177
                        case "":
                            var v178 string = ("\n\t" + "\n\t")
                            _ = v178
                        default:
                            if v4 {
                                var v179 int64 = 42
                                v180 := GenS22{F0: 619, F1: ((^(v1)) == 0), F2: 1}
                                var v181 uint16 = 169
                                _ = v179
                                _ = v180
                                _ = v181
                            }
                        }
                        if ((v1 + v1) <= 2) {
                            v2 += v131
                            fmt.Println(v176, v130)
                        } else {
                            _ = len(v135)
                            var v182 int8 = v6
                            var v183 int16 = v130
                            _ = v182
                            _ = v183
                        }
                    } else {
                        v184 := GenS1{F0: "a"}
                        v185 := GenS23{F0: "\n\t", F1: (v1 - 467)}
                        var v186 int16 = v130
                        _ = v184
                        _ = v185
                        _ = v186
                    }
                    _ = v176
                }
                _ = fi152
            }
            var v187 int = (fi151 - (fi151 >> ((uint(v1)) & 63)))
            _ = fi151
            _ = v187
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v129:", v129)
    fmt.Println("v130:", v130)
    fmt.Println("v131:", v131)
    fmt.Println("slicelen v135:", len(v135), v135)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v129
    _ = v130
    _ = v131
    _ = v135
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
