package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 int64; F1 int8; F2 uint16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint; F1 int16 }
type GenS8 struct { F0 int; F1 uint32 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint; F1 int; F2 bool }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 uint8; F1 int; F2 int64 }
func genFn0() string {
    var v4 uint32 = 42
    v5 := []int{4294967295, 300, -1, 128}
    for fi6 := 0; fi6 < 4; fi6++ {
        v4 += v4
        _ = fi6
    }
    switch "" {
    case "a":
        fmt.Println(v4)
    case "\n\t":
        v7 := GenS1{F0: -1}
        _ = v7
    }
    _ = v4
    _ = v5
    return "a"
}
func genFn1() uint {
    var v8 uint8 = 15
    if false {
        switch "a" {
        case "a":
            fmt.Println(v8)
        case "hello":
            for fi9 := 0; fi9 < 4; fi9++ {
                v10 := []int{}
                _ = fi9
                _ = v10
            }
        case "\n\t":
            for fi11 := 0; fi11 < 12; fi11++ {
                v8 = v8
                v8 = v8
                v12 := []int{860}
                _ = fi11
                _ = v12
            }
        default:
            v13 := GenS1{F0: 100}
            _ = v13
        }
        var v14 uint16 = 347
        var v15 uint8 = v8
        _ = v14
        _ = v15
    } else {
        if true {
            var v16 int64 = 0
            _ = v16
        }
    }
    if true {
        switch -80 {
        case -685:
            switch "" {
            case "a":
                v17 := []int{-337, 2147483647, 0, 0}
                _ = v17
            case "\n\t":
                fmt.Println(v8)
            default:
                v8 -= 0
            }
        case 4294967295:
            if (((((float64((0.1 / 0.1)) + (1e10 + -2.25)) * 1.5) - (0.1 / (-2.25 * 3.14159))) * -2.25) != -2.25) {
                var v18 uint32 = 4294967295
                var v19 float64 = 1.5
                _ = v18
                _ = v19
            } else {
                var v20 uint = 2
                _ = v20
            }
        case 1:
            v8 += 94
        }
        v21 := GenS1{F0: -128}
        v8 = (v8 & v8)
        _ = v21
    }
    _ = v8
    return 0
}

func main() {
    var v1 bool = false
    var v2 float32 = (-2.25 / (1e10 / 1.5))
    var v3 int8 = 127
    v22 := []int{}
    var v23 int = 42
    if len(v22) > 0 {
        v22[0] = (v23 - (v23 - v23))
    }
    var v24 uint64 = 42
    if v1 {
        v25 := GenS1{F0: 126}
        for fi26 := 0; fi26 < 15; fi26++ {
            var v27 int64 = 9223372036854775807
            for fi28 := 0; fi28 < 13; fi28++ {
                v24 -= (v24 | v24)
                fmt.Println(v1)
                if v1 {
                    var v29 uint8 = 42
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                    if (18446744073709551615 >= 613) {
                        var v30 int8 = v3
                        _ = v30
                    }
                    _ = v29
                }
                v31 := GenS2{F0: 912, F1: 1, F2: (v24 / ((v24) | 1))}
                v32 := GenS1{F0: (v3 + 120)}
                v33 := GenS3{F0: 792, F1: v3}
                _ = fi28
                _ = v31
                _ = v32
                _ = v33
            }
            var v34 int32 = 2147483647
            _ = fi26
            _ = v27
            _ = v34
        }
        for fi35 := 0; fi35 < 9; fi35++ {
            if (!(4294967295 == 42)) {
                if len(v22) > 0 {
                    v22[0] = fi35
                }
            }
            _ = fi35
        }
        _ = v25
    } else {
        fmt.Println(v23, v1, v23)
        v36 := GenS4{F0: genFn1(), F1: 1}
        if v1 {
            v37 := GenS1{F0: v3}
            v38 := GenS3{F0: 4294967295, F1: 41}
            if ((!(42 > 190)) && (v23 <= -949)) {
                for fi39 := 0; fi39 < 14; fi39++ {
                    for fi40 := 0; fi40 < 6; fi40++ {
                        v22 = append(v22, v23)
                        fmt.Println(v2)
                        if (127 != uint8(v24)) {
                            if len(v22) > 0 {
                                v22[0] = (fi40 - fi39)
                            }
                            fmt.Println(v24, fi39, v23)
                            var v41 uint64 = v24
                            _ = v41
                        } else {
                            var v42 uint16 = 1
                            _ = v42
                        }
                        v43 := GenS4{F0: 848, F1: -9223372036854775808}
                        v3 = v3
                        _ = fi40
                        _ = v43
                    }
                    _ = fi39
                }
                var v44 int = (v23 >> ((597) & 63))
                var v45 int64 = 1
                _ = v44
                _ = v45
            }
            _ = v37
            _ = v38
        }
        _ = v36
    }
    var v46 string = "foo bar"
    if v1 {
        v22 = append(v22, (v23 % (((^(v23))) | 1)))
        v2 = v2
    } else {
        for fi47 := 0; fi47 < 4; fi47++ {
            var v48 uint8 = 19
            _ = fi47
            _ = v48
        }
    }
    v3 = v3
    fmt.Println(v24)
    var v49 int16 = 42
    if v1 {
        if ((v49 >> ((2) & 63)) <= v49) {
            fmt.Println(v46)
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
            if (!(v46 == (v46 + ((v46 + v46) + ("\n\t" + ((("" + v46) + (v46 + v46)) + v46)))))) {
                if v1 {
                    fmt.Println(v1)
                    var v50 int8 = (-(v3))
                    if v1 {
                        var v51 int32 = 42
                        if ((^(v50)) >= v3) {
                            for fi52 := 0; fi52 < 12; fi52++ {
                                if false {
                                    v2 = 0.0
                                    var v53 int8 = 100
                                    var v54 float64 = 0.1
                                    _ = v53
                                    _ = v54
                                }
                                var v55 uint64 = (v24 & (v24 / ((42) | 1)))
                                if (v2 < ((1e10 * (v2 / v2)) * v2)) {
                                    var v56 int16 = v49
                                    _ = v56
                                } else {
                                    v2 = v2
                                    var v57 uint = genFn1()
                                    v55 += (v55 | 851)
                                    _ = v57
                                }
                                _ = fi52
                                _ = v55
                            }
                        } else {
                            _ = len(v22)
                        }
                        if ((-2.25 > v2) && (42 > 2)) {
                            var v58 string = "a"
                            _ = v58
                        }
                        _ = v51
                    }
                    _ = v50
                } else {
                    var v59 uint16 = 65535
                    fmt.Println(v59)
                    if true {
                        var v60 int64 = 0
                        switch v46 {
                        case "":
                            v61 := GenS5{F0: (v60 ^ v60), F1: 1, F2: (v59 / ((v59) | 1))}
                            _ = v61
                        case "hello":
                            fmt.Println(v23, v60, v60)
                        case "\n\t":
                            var v62 int = v23
                            _ = v62
                        }
                        if (!v1) {
                            if (362 != -26) {
                                if (v1 || (18446744073709551615 >= 18446744073709551615)) {
                                    for fi63 := 0; fi63 < 9; fi63++ {
                                        var v64 float64 = (0.0 / 1.0)
                                        v24 -= v24
                                        _ = fi63
                                        _ = v64
                                    }
                                } else {
                                    switch (-(v23)) {
                                    case -171:
                                        var v65 string = ("a" + v46)
                                        _ = v65
                                    default:
                                        _ = len(v22)
                                    }
                                }
                                v60 = (v60 & (^(v60)))
                            }
                            var v66 int32 = 0
                            _ = v66
                        }
                        _ = v60
                    } else {
                        if false {
                            _ = len(v22)
                            var v67 uint8 = 255
                            _ = v67
                        }
                    }
                    _ = v59
                }
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
                if true {
                    if (!(!((!v1) || (((v46 + v46) != (("hello" + v46) + "foo bar")) && (!(v49 > 1000)))))) {
                        for fi68 := 0; fi68 < 12; fi68++ {
                            for fi69 := 0; fi69 < 18; fi69++ {
                                v46 = v46
                                if v1 {
                                    v49 -= v49
                                    var v70 int64 = -653
                                    v2 -= (1.5 + v2)
                                    _ = v70
                                }
                                _ = fi69
                            }
                            _ = fi68
                        }
                        var v71 float64 = 1.5
                        _ = v71
                    }
                } else {
                    v23 = -935
                    var v72 string = v46
                    _ = v72
                }
            }
        } else {
            if v1 {
                var v73 int32 = -654
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
                var v74 uint16 = 1
                _ = v73
                _ = v74
            } else {
                for fi75 := 0; fi75 < 11; fi75++ {
                    v46 = ("foo bar" + v46)
                    if false {
                        switch v46 {
                        case "foo bar":
                            v76 := GenS4{F0: 2, F1: 42}
                            _ = v76
                        case "\n\t":
                            switch (fi75 & 772) {
                            case 9223372036854775807:
                                for fi77 := 0; fi77 < 18; fi77++ {
                                    v1 = (42 >= -1)
                                    v22 = append(v22, (fi75 + 32768))
                                    _ = fi77
                                }
                            }
                        }
                        var v78 int16 = int16(v24)
                        _ = v78
                    }
                    var v79 uint64 = (v24 << ((255) & 63))
                    _ = fi75
                    _ = v79
                }
                v23 -= int(v3)
                switch v23 {
                case 175:
                    var v80 int64 = 954
                    _ = v80
                case -1:
                    fmt.Println(v23)
                case 219:
                    for fi81 := 0; fi81 < 15; fi81++ {
                        _ = len(v22)
                        v46 = (v46 + (v46 + "\n\t"))
                        _ = fi81
                    }
                }
            }
            for fi82 := 0; fi82 < 5; fi82++ {
                if (true && (-300 < 163)) {
                    _ = len(v22)
                    if ((!(!v1)) || (false || (!v1))) {
                        fmt.Println(v1, v46, v2)
                        v23 = (v23 ^ (fi82 << ((2) & 63)))
                    } else {
                        if false {
                            if len(v22) > 0 {
                                v22[0] = (v23 / (((^(v23))) | 1))
                            }
                        } else {
                            for fi83 := 0; fi83 < 9; fi83++ {
                                var v84 int8 = v3
                                _ = fi83
                                _ = v84
                            }
                        }
                        if (v1 || true) {
                            var v85 int32 = 2147483647
                            var v86 int64 = 1
                            for fi87 := 0; fi87 < 9; fi87++ {
                                var v88 float64 = 3.14159
                                var v89 int = 128
                                v90 := GenS6{F0: (v46 + ((v46 + v46) + "foo bar"))}
                                if v1 {
                                    if false {
                                        var v91 uint16 = 65535
                                        v92 := GenS7{F0: 2, F1: v49}
                                        _ = v91
                                        _ = v92
                                    }
                                    switch "a" {
                                    case "a":
                                        fmt.Println(v89)
                                    default:
                                        fmt.Println(fi82, v23, v23)
                                    }
                                    if v1 {
                                        if (((float64(3.14159) / v88) < (1.5 + (-2.25 / v88))) || false) {
                                            var v93 string = (v46 + v46)
                                            switch (((v46 + "foo bar") + v93) + v93) {
                                            case "\n\t":
                                                switch v93 {
                                                case "foo bar":
                                                    var v94 bool = ((!(v46 < v46)) && (1 == 65535))
                                                    _ = v94
                                                case "hello":
                                                    var v95 uint32 = 737
                                                    _ = v95
                                                case "\n\t":
                                                    if false {
                                                        if v1 {
                                                            fmt.Println(fi82, v85)
                                                        } else {
                                                            v3 = (-(v3))
                                                            if len(v22) >= 2 {
                                                                v22 = v22[1:len(v22)]
                                                            }
                                                            v2 += 3.14159
                                                        }
                                                        var v96 float64 = float64((v2 * (v2 * v2)))
                                                        _ = v96
                                                    } else {
                                                        var v97 int16 = v49
                                                        v22 = append(v22, 0)
                                                        _ = v97
                                                    }
                                                default:
                                                    if (v1 || v1) {
                                                        if (!((42 == (v86 % ((-561) | 1))) && (!true))) {
                                                            if (-1 == (v85 * (v85 / (((v85 / ((v85) | 1))) | 1)))) {
                                                                switch (v89 / ((fi82) | 1)) {
                                                                case 42:
                                                                    v24 -= 165
                                                                    v46 = v46
                                                                }
                                                                v98 := GenS2{F0: genFn1(), F1: v49, F2: v24}
                                                                v99 := GenS7{F0: 255, F1: (v49 << ((259) & 63))}
                                                                if false {
                                                                    v100 := GenS8{F0: (fi87 * fi87), F1: 0}
                                                                    if len(v22) >= 2 {
                                                                        v22 = v22[1:len(v22)]
                                                                    }
                                                                    if (-614 != fi87) {
                                                                        v85 += v85
                                                                    } else {
                                                                        _ = len(v22)
                                                                        if len(v22) > 0 {
                                                                            v22[0] = (v23 ^ (fi82 & (fi87 * (fi82 * fi87))))
                                                                        }
                                                                    }
                                                                    _ = v100
                                                                } else {
                                                                    var v101 string = v93
                                                                    v22 = append(v22, -992)
                                                                    switch fi87 {
                                                                    case 32767:
                                                                        v22 = append(v22, (v23 % ((fi87) | 1)))
                                                                    case 127:
                                                                        var v102 float64 = v88
                                                                        _ = v102
                                                                    default:
                                                                        switch (("hello" + "\n\t") + v93) {
                                                                        case "hello":
                                                                            _ = len(v22)
                                                                        }
                                                                    }
                                                                    fmt.Println(v46, v23)
                                                                    if len(v22) >= 2 {
                                                                        v22 = v22[1:len(v22)]
                                                                    }
                                                                    _ = v101
                                                                }
                                                                var v103 float32 = v2
                                                                _ = v103
                                                                _ = v98
                                                                _ = v99
                                                            }
                                                            v86 += v86
                                                        } else {
                                                            _ = len(v22)
                                                            var v104 int = (fi87 ^ int(v49))
                                                            _ = v104
                                                        }
                                                    } else {
                                                        var v105 float32 = v2
                                                        v106 := GenS1{F0: 100}
                                                        _ = v105
                                                        _ = v106
                                                    }
                                                }
                                            case "a":
                                                var v107 uint64 = (v24 + (v24 >> ((953) & 63)))
                                                _ = v107
                                            case "":
                                                if ((v49 < 1) && true) {
                                                    var v108 uint = 2
                                                    fmt.Println(v24)
                                                    _ = v108
                                                } else {
                                                    fmt.Println(v46, v3, v46)
                                                    v109 := GenS4{F0: 577, F1: v86}
                                                    v110 := GenS9{F0: (v93 != v46)}
                                                    _ = v109
                                                    _ = v110
                                                }
                                            default:
                                                v2 = v2
                                            }
                                            var v111 uint16 = 111
                                            _ = v111
                                            _ = v93
                                        } else {
                                            if (!v1) {
                                                if ((v1 && true) || v1) {
                                                    if ((v3 - (v3 / ((v3) | 1))) >= (^(v3))) {
                                                        var v112 bool = (((!v1) && false) && (v86 == v86))
                                                        _ = v112
                                                    } else {
                                                        v3 -= 0
                                                    }
                                                    v86 = (v86 - 0)
                                                    v113 := GenS9{F0: (!(!v1))}
                                                    _ = v113
                                                }
                                                v1 = (v24 < (v24 & 2))
                                                v49 = 32767
                                            } else {
                                                v86 = v86
                                                if v1 {
                                                    v114 := GenS10{F0: 2, F1: -1000, F2: v1}
                                                    _ = v114
                                                }
                                                switch "a" {
                                                case "foo bar":
                                                    v85 = (v85 * 344)
                                                    if v1 {
                                                        _ = len(v22)
                                                    } else {
                                                        var v115 string = "foo bar"
                                                        v116 := GenS11{F0: ("\n\t" + ("" + ("foo bar" + ("hello" + (v46 + "a")))))}
                                                        v89 = (-(v23))
                                                        fmt.Println(v46, v88, v89)
                                                        v2 = 1e10
                                                        _ = v115
                                                        _ = v116
                                                    }
                                                    var v117 uint64 = 260
                                                    _ = v117
                                                case "":
                                                    v118 := GenS3{F0: (v24 / (((v24 & (v24 & v24))) | 1)), F1: (v3 + v3)}
                                                    v119 := GenS12{F0: 2}
                                                    switch (v23 + fi82) {
                                                    case 2147483647:
                                                        fmt.Println(v24)
                                                    case 1:
                                                        var v120 uint8 = 2
                                                        _ = v120
                                                    default:
                                                        v121 := GenS5{F0: -263, F1: v3, F2: 42}
                                                        _ = v121
                                                    }
                                                    if ((true || false) || false) {
                                                        if (!v1) {
                                                            v122 := GenS13{F0: v3}
                                                            if len(v22) > 0 {
                                                                v22[0] = 127
                                                            }
                                                            _ = v122
                                                        }
                                                        v23 += 782
                                                        v123 := GenS9{F0: ((v24 ^ v24) != (v24 | (v24 & (^(v24)))))}
                                                        _ = v123
                                                    } else {
                                                        var v124 float32 = ((v2 * v2) - v2)
                                                        _ = v124
                                                    }
                                                    var v125 uint = 42
                                                    _ = v118
                                                    _ = v119
                                                    _ = v125
                                                default:
                                                    switch v46 {
                                                    case "a":
                                                        v89 -= v23
                                                    case "hello":
                                                        fmt.Println(v85)
                                                    case "\n\t":
                                                        v22 = append(v22, (-(v23)))
                                                    default:
                                                        if v1 {
                                                            if v1 {
                                                                var v126 uint8 = 170
                                                                fmt.Println(v24)
                                                                _ = v126
                                                            } else {
                                                                if true {
                                                                    var v127 int8 = 100
                                                                    _ = v127
                                                                } else {
                                                                    v128 := GenS4{F0: 599, F1: (-(v86))}
                                                                    switch (fi87 << ((438) & 63)) {
                                                                    case 42:
                                                                        var v129 uint = genFn1()
                                                                        _ = v129
                                                                    default:
                                                                        switch (v46 + ("hello" + (v46 + (v46 + "foo bar")))) {
                                                                        case "\n\t":
                                                                            if ((v86 & v86) == v86) {
                                                                                fmt.Println(v89, v23, v3)
                                                                                v130 := GenS14{F0: int64(v24)}
                                                                                _ = v130
                                                                            } else {
                                                                                v131 := GenS5{F0: int64(v24), F1: -123, F2: 42}
                                                                                v2 = v2
                                                                                var v132 bool = v1
                                                                                v133 := GenS13{F0: (v3 & v3)}
                                                                                if ((true || v1) || (42 < 0)) {
                                                                                    v3 -= (-(v3))
                                                                                }
                                                                                _ = v131
                                                                                _ = v132
                                                                                _ = v133
                                                                            }
                                                                        case "foo bar":
                                                                            v134 := GenS15{F0: 243}
                                                                            _ = v134
                                                                        case "":
                                                                            var v135 int32 = v85
                                                                            _ = v135
                                                                        }
                                                                    }
                                                                    _ = v128
                                                                }
                                                                var v136 uint = 18446744073709551615
                                                                _ = v136
                                                            }
                                                            var v137 int32 = 2147483647
                                                            _ = v137
                                                        } else {
                                                            var v138 float32 = (v2 - float32(float64(float32(((((v88 - v88) * v88) / 0.1) * (v88 + ((v88 * 1e10) / (v88 + 0.0))))))))
                                                            if false {
                                                                var v139 float64 = v88
                                                                v22 = append(v22, (v23 + int(v24)))
                                                                fmt.Println(v139, v46)
                                                                _ = v139
                                                            } else {
                                                                if len(v22) > 0 {
                                                                    v22[0] = v89
                                                                }
                                                                v1 = false
                                                                var v140 int64 = v86
                                                                _ = v140
                                                            }
                                                            _ = v138
                                                        }
                                                    }
                                                }
                                                v46 = v46
                                            }
                                        }
                                    } else {
                                        if (42 <= 128) {
                                            var v141 uint32 = 2147483647
                                            _ = v141
                                        } else {
                                            if v1 {
                                                v85 += (v85 * v85)
                                                v22 = append(v22, v89)
                                            } else {
                                                v89 = 0
                                                _ = len(v22)
                                            }
                                        }
                                        var v142 uint64 = (v24 >> ((798) & 63))
                                        _ = v142
                                    }
                                    if len(v22) > 0 {
                                        v22[0] = fi82
                                    }
                                    if len(v22) > 0 {
                                        v22[0] = int(v49)
                                    }
                                    if ("\n\t" >= (v46 + "a")) {
                                        if len(v22) > 0 {
                                            v22[0] = v89
                                        }
                                        if false {
                                            v85 = (-(v85))
                                            v22 = append(v22, v89)
                                        }
                                        var v143 int32 = v85
                                        _ = v143
                                    }
                                } else {
                                    v144 := GenS3{F0: v24, F1: (v3 >> ((65535) & 63))}
                                    fmt.Println(v24)
                                    switch (v89 - -123) {
                                    case -879:
                                        v49 -= (v49 & int16(v86))
                                    case 127:
                                        v145 := GenS16{F0: 234, F1: v23, F2: v86}
                                        v22 = append(v22, fi87)
                                        _ = v145
                                    case -651:
                                        var v146 int64 = v86
                                        _ = v146
                                    }
                                    _ = v144
                                }
                                if (4294967295 < 2) {
                                    v22 = append(v22, (v23 | v89))
                                } else {
                                    v147 := GenS4{F0: genFn1(), F1: -1}
                                    _ = v147
                                }
                                var v148 uint = 0
                                v23 = fi82
                                _ = fi87
                                _ = v148
                                _ = v88
                                _ = v89
                                _ = v90
                            }
                            _ = v85
                            _ = v86
                        }
                    }
                    if (1 == 2147483648) {
                        v46 = ("a" + "")
                    }
                } else {
                    fmt.Println(v1, fi82)
                }
                v149 := GenS10{F0: genFn1(), F1: (v23 | (fi82 >> ((1) & 63))), F2: true}
                v3 += v3
                _ = fi82
                _ = v149
            }
            v2 -= float32(-2.25)
        }
        fmt.Println(v1, v2, v46)
    } else {
        fmt.Println(v2, v49, v49)
        v2 = (v2 + (((v2 + v2) / v2) / float32(3.14159)))
        v1 = (98 == 127)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v46
    _ = v49
}
