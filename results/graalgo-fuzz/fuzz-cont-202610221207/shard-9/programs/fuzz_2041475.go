package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int64; F1 uint64; F2 bool }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int8; F1 string; F2 string }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int; F1 int }
type GenS11 struct { F0 int; F1 int8; F2 string }
func genFn0() uint16 {
    var v5 int32 = 916
    v5 += 2147483647
    _ = v5
    return 320
}

func main() {
    var v1 string = ("" + "hello")
    v2 := make(map[int]bool)
    v3 := []uint32{2147483647}
    var v4 int16 = 32767
    for fi6 := 0; fi6 < 17; fi6++ {
        fmt.Println(v4, v1, fi6)
        if (396 < 475) {
            if (-2.25 <= 0.1) {
                var v7 int32 = int32(fi6)
                _ = v7
            } else {
                var v8 uint = 379
                switch 0 {
                case 995:
                    v8 = (v8 >> ((v8) & 63))
                case 566:
                    v8 = (v8 % ((4294967295) | 1))
                case 32768:
                    for fi9 := 0; fi9 < 19; fi9++ {
                        fmt.Println(fi6, fi9)
                        var v10 float64 = float64(1.5)
                        _ = fi9
                        _ = v10
                    }
                default:
                    var v11 float32 = (3.14159 / 3.14159)
                    _ = v11
                }
                var v12 int = fi6
                _ = v12
                _ = v8
            }
            var v13 uint = 1
            fmt.Println(v1, fi6)
            _ = v13
        }
        fmt.Println(fi6)
        _ = fi6
    }
    if false {
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 uint8 = 128
            _ = fi14
            _ = v15
        }
        delete(v2, 32768)
    } else {
        v16 := GenS1{F0: 846, F1: ""}
        _ = v16
    }
    if false {
        var v17 int8 = 0
        if (-2147483648 >= 151) {
            v4 = v4
            for fi18 := 0; fi18 < 20; fi18++ {
                var v19 bool = (2 <= 42)
                v20 := GenS2{F0: 0, F1: 397, F2: (!((!(v19 || (25 < 127))) || false))}
                var v21 int64 = 0
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
        }
        for fi22 := 0; fi22 < 18; fi22++ {
            v23 := GenS2{F0: 1, F1: 1, F2: ((true || (true && false)) || (0 == 255))}
            if (true && true) {
                v24 := GenS1{F0: 0, F1: (v1 + v1)}
                var v25 int32 = 42
                _ = v24
                _ = v25
            } else {
                for fi26 := 0; fi26 < 15; fi26++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    switch fi26 {
                    case 9223372036854775807:
                        fmt.Println(v1, fi26, fi22)
                    case 2147483647:
                        switch "a" {
                        case "":
                            v27 := GenS2{F0: 9223372036854775807, F1: 2, F2: (0 == 2147483647)}
                            _ = v27
                        default:
                            fmt.Println(v1, fi26, fi22)
                        }
                    default:
                        var v28 uint = 867
                        _ = v28
                    }
                    v1 = v1
                    _ = fi26
                }
                v17 -= v17
                if false {
                    v1 = v1
                    v1 = ("hello" + ("\n\t" + v1))
                }
            }
            v1 = (((("hello" + v1) + (v1 + "\n\t")) + v1) + v1)
            _ = fi22
            _ = v23
        }
        _ = v17
    } else {
        var v29 float32 = 1.5
        _ = v29
    }
    if ((0 != 884) || false) {
        for fi30 := 0; fi30 < 10; fi30++ {
            v2[fi30] = ((!(-2.25 > (0.1 * (1e10 + -2.25)))) && true)
            var v31 uint16 = genFn0()
            var v32 uint = 631
            _ = fi30
            _ = v31
            _ = v32
        }
        v4 += 1000
    }
    switch 9223372036854775807 {
    case 1:
        var v33 uint64 = 2
        _ = v33
    default:
        v1 = v1
    }
    if (!false) {
        var v34 int16 = v4
        _ = v34
    }
    switch 313 {
    case 32768:
        v1 = v1
    case 128:
        v35 := GenS1{F0: 0, F1: "foo bar"}
        _ = v35
    case -253:
        switch (v1 + ((v1 + "a") + ("hello" + (v1 + "")))) {
        case "a":
            var v36 uint8 = 2
            _ = v36
        case "":
            if (!(1.5 != float64(-2.25))) {
                if ((!false) || (2 <= 65535)) {
                    var v37 uint64 = 307
                    if (32767 < genFn0()) {
                        v37 = (v37 << ((0) & 63))
                        var v38 float64 = 1.5
                        if (v1 > v1) {
                            for fi39 := 0; fi39 < 3; fi39++ {
                                if true {
                                    if (0 == (fi39 >> ((4294967295) & 63))) {
                                        v1 = v1
                                        for fi40 := 0; fi40 < 20; fi40++ {
                                            v1 = (v1 + "a")
                                            if false {
                                                v41 := GenS3{F0: 127, F1: -97}
                                                v42 := GenS3{F0: 0, F1: 1}
                                                if false {
                                                    v38 = (v38 * (1e10 * (1e10 - v38)))
                                                } else {
                                                    switch fi40 {
                                                    case 2147483647:
                                                        v1 = (v1 + v1)
                                                    default:
                                                        v4 += v4
                                                    }
                                                }
                                                _ = v41
                                                _ = v42
                                            }
                                            var v43 uint16 = uint16(v37)
                                            _ = fi40
                                            _ = v43
                                        }
                                        _ = len(v3)
                                    } else {
                                        if true {
                                            if false {
                                                var v44 string = v1
                                                _ = v44
                                            } else {
                                                for fi45 := 0; fi45 < 10; fi45++ {
                                                    var v46 int16 = v4
                                                    _ = fi45
                                                    _ = v46
                                                }
                                            }
                                            fmt.Println(v1, v1)
                                        } else {
                                            v47 := GenS4{F0: 246, F1: 904}
                                            _ = v47
                                        }
                                        v48 := GenS5{F0: 663}
                                        _ = v48
                                    }
                                    var v49 uint16 = 255
                                    _ = v49
                                } else {
                                    v3 = append(v3, 2)
                                }
                                if (v1 < v1) {
                                    if true {
                                        var v50 float32 = 0.0
                                        _ = v50
                                    } else {
                                        switch "foo bar" {
                                        case "a":
                                            var v51 float64 = (v38 * ((v38 + v38) / v38))
                                            _ = v51
                                        }
                                        v4 += (v4 << ((340) & 63))
                                        switch fi39 {
                                        case 573:
                                            if true {
                                                if false {
                                                    var v52 float64 = (v38 + v38)
                                                    for fi53 := 0; fi53 < 3; fi53++ {
                                                        fmt.Println(fi39, v1, fi39)
                                                        var v54 uint32 = 2147483648
                                                        _ = fi53
                                                        _ = v54
                                                    }
                                                    switch fi39 {
                                                    case 42:
                                                        v55 := GenS6{F0: 896, F1: v4, F2: 17}
                                                        _ = v55
                                                    case 4294967295:
                                                        v56 := GenS6{F0: 39, F1: (v4 ^ 1000), F2: 255}
                                                        _ = v56
                                                    case 127:
                                                        var v57 uint = 255
                                                        _ = v57
                                                    }
                                                    _ = v52
                                                } else {
                                                    var v58 bool = true
                                                    v3 = append(v3, 2)
                                                    v59 := GenS3{F0: 74, F1: 1}
                                                    _ = v58
                                                    _ = v59
                                                }
                                                var v60 uint64 = v37
                                                for fi61 := 0; fi61 < 15; fi61++ {
                                                    fmt.Println(v4, v1, v60)
                                                    if (false || (1 != -2147483648)) {
                                                        if (837 != -1) {
                                                            if true {
                                                                v60 = v37
                                                            } else {
                                                                var v62 int64 = 42
                                                                v62 = (v62 / (((-(v62))) | 1))
                                                                fmt.Println(v37, fi39)
                                                                _ = v62
                                                            }
                                                        } else {
                                                            fmt.Println(v4, v38, fi39)
                                                            var v63 int16 = (v4 >> ((4294967295) & 63))
                                                            _ = v63
                                                        }
                                                        var v64 uint64 = v37
                                                        _ = v64
                                                    }
                                                    _ = fi61
                                                }
                                                _ = v60
                                            } else {
                                                switch (^(fi39)) {
                                                case -702:
                                                    fmt.Println(v1, v1, fi39)
                                                default:
                                                    delete(v2, fi39)
                                                }
                                                switch ("hello" + (((v1 + ("\n\t" + v1)) + (("a" + (v1 + v1)) + v1)) + (((v1 + v1) + (((v1 + "foo bar") + v1) + ("\n\t" + ("" + v1)))) + ("" + v1)))) {
                                                case "foo bar":
                                                    switch fi39 {
                                                    case -1:
                                                        switch ("\n\t" + "hello") {
                                                        case "a":
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        default:
                                                            var v65 float32 = ((float32((v38 * v38)) * float32(0.0)) - ((3.14159 - (3.14159 / -2.25)) + (-2.25 - (-2.25 / 0.1))))
                                                            _ = v65
                                                        }
                                                    case 32768:
                                                        v4 -= v4
                                                    case 2147483647:
                                                        v37 -= (v37 * 2)
                                                    }
                                                }
                                                v1 = "hello"
                                            }
                                        case -9223372036854775808:
                                            v38 -= v38
                                        default:
                                            for fi66 := 0; fi66 < 2; fi66++ {
                                                var v67 int = -630
                                                _ = fi66
                                                _ = v67
                                            }
                                        }
                                    }
                                }
                                _ = fi39
                            }
                            switch v1 {
                            case "foo bar":
                                v68 := GenS5{F0: 0}
                                _ = v68
                            }
                        } else {
                            v37 = v37
                            var v69 string = v1
                            if (v4 >= v4) {
                                v1 = ((v69 + ((("hello" + "\n\t") + "hello") + v69)) + ((v1 + (v1 + "")) + ("" + "")))
                                for fi70 := 0; fi70 < 11; fi70++ {
                                    v71 := GenS5{F0: 915}
                                    v69 = "hello"
                                    v72 := GenS7{F0: 2}
                                    _ = fi70
                                    _ = v71
                                    _ = v72
                                }
                                var v73 bool = true
                                _ = v73
                            } else {
                                fmt.Println(v69)
                                var v74 int = 32768
                                _ = v74
                            }
                            _ = v69
                        }
                        _ = v38
                    } else {
                        if (0 != 0) {
                            for fi75 := 0; fi75 < 18; fi75++ {
                                v1 = "\n\t"
                                v3 = append(v3, 763)
                                v37 = 1
                                _ = fi75
                            }
                        }
                    }
                    _ = v37
                }
                fmt.Println(v4, v1)
                for fi76 := 0; fi76 < 18; fi76++ {
                    if (1 > -128) {
                        switch ("a" + "") {
                        case "":
                            fmt.Println(v1, v4)
                        case "hello":
                            v77 := GenS4{F0: 0, F1: 2}
                            _ = v77
                        case "\n\t":
                            var v78 uint32 = 65535
                            _ = v78
                        default:
                            var v79 uint = 2
                            _ = v79
                        }
                        switch (-(fi76)) {
                        case 830:
                            if ((((0.1 - 1e10) / 1.0) + ((0.0 - (1.5 * 0.1)) - 1.5)) != (0.0 + -2.25)) {
                                _ = len(v3)
                            } else {
                                v4 = v4
                            }
                        case 32767:
                            v80 := GenS8{F0: -128, F1: (v1 + ("hello" + v1)), F2: "hello"}
                            _ = v80
                        default:
                            var v81 uint16 = 0
                            _ = v81
                        }
                        for fi82 := 0; fi82 < 2; fi82++ {
                            v4 -= 1
                            var v83 uint = 2
                            var v84 uint8 = 42
                            _ = fi82
                            _ = v83
                            _ = v84
                        }
                    } else {
                        var v85 uint16 = 2
                        _ = v85
                    }
                    if true {
                        var v86 uint32 = 4294967295
                        v1 = ("hello" + v1)
                        _ = v86
                    }
                    _ = fi76
                }
            } else {
                if true {
                    v87 := GenS8{F0: 20, F1: (v1 + v1), F2: ("\n\t" + v1)}
                    if ((v4 % (((v4 * (-(v4)))) | 1)) <= v4) {
                        v4 += (v4 + (v4 << ((42) & 63)))
                        fmt.Println(v4, v4)
                        if false {
                            v88 := GenS9{F0: 0}
                            if (!true) {
                                v2[-9223372036854775808] = (false && true)
                                if (true || true) {
                                    v89 := v2[-817]
                                    var v90 int32 = -1
                                    _ = v89
                                    _ = v90
                                }
                                switch -9223372036854775808 {
                                case 4294967295:
                                    fmt.Println(v1)
                                case 924:
                                    var v91 int32 = -2147483648
                                    _ = v91
                                default:
                                    if false {
                                        var v92 int8 = 42
                                        delete(v2, 32767)
                                        _ = v92
                                    } else {
                                        v1 = v1
                                        v4 += v4
                                        var v93 float32 = float32(((-2.25 - ((-2.25 + (0.0 / 0.1)) / 1.5)) * 0.1))
                                        _ = v93
                                    }
                                }
                            }
                            _ = v88
                        } else {
                            var v94 int32 = 0
                            _ = v94
                        }
                    } else {
                        v95 := GenS10{F0: 1, F1: 128}
                        var v96 int = -9223372036854775808
                        if ((float32(0.1) - -2.25) >= ((1.5 / 1e10) / (0.0 + 0.1))) {
                            var v97 uint16 = 895
                            v1 = v1
                            v98 := GenS10{F0: v96, F1: (v96 + (v96 << ((4294967295) & 63)))}
                            _ = v97
                            _ = v98
                        }
                        _ = v95
                        _ = v96
                    }
                    for fi99 := 0; fi99 < 16; fi99++ {
                        v2[(fi99 % ((998) | 1))] = (0.0 != -2.25)
                        _ = fi99
                    }
                    _ = v87
                }
                v1 = v1
                v2[821] = (!(false && ((true || (-532 <= 42)) && false)))
            }
        }
    default:
        var v100 uint32 = 2
        _ = v100
    }
    v101 := v2[965]
    switch 9223372036854775807 {
    case -2147483648:
        var v102 uint64 = 2
        _ = v102
    default:
        if v101 {
            var v103 int8 = -1
            v1 = v1
            _ = v103
        } else {
            for fi104 := 0; fi104 < 7; fi104++ {
                v4 = (^(v4))
                _ = fi104
            }
        }
    }
    v105 := GenS3{F0: 28, F1: 127}
    var v106 int32 = -1
    v106 += v106
    fmt.Println(v101)
    fmt.Println(v1, v106, v1)
    var v107 int32 = v106
    if v101 {
        delete(v2, 32767)
    } else {
        for fi108 := 0; fi108 < 14; fi108++ {
            var v109 int8 = -12
            v109 -= (v109 >> ((1) & 63))
            if v101 {
                for fi110 := 0; fi110 < 20; fi110++ {
                    switch fi110 {
                    case 2147483647:
                        var v111 uint = 656
                        switch (v1 + "hello") {
                        case "a":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        default:
                            var v112 int32 = (-(v107))
                            _ = v112
                        }
                        _ = v111
                    case 32768:
                        delete(v2, fi108)
                    default:
                        switch "a" {
                        case "a":
                            v113 := v2[fi108]
                            _ = v113
                        }
                    }
                    var v114 uint64 = 4294967295
                    var v115 string = ""
                    _ = fi110
                    _ = v114
                    _ = v115
                }
                v107 += v106
                if true {
                    v107 += (^(v107))
                    var v116 string = v1
                    for fi117 := 0; fi117 < 18; fi117++ {
                        fmt.Println(fi117, v116, v116)
                        switch -64 {
                        case 42:
                            switch v116 {
                            case "\n\t":
                                v118 := GenS7{F0: 0}
                                _ = v118
                            case "foo bar":
                                var v119 float64 = 0.0
                                _ = v119
                            case "":
                                v120 := GenS4{F0: 190, F1: 0}
                                _ = v120
                            }
                        }
                        var v121 uint16 = 1
                        _ = len(v2)
                        v122 := GenS3{F0: 1, F1: v109}
                        _ = fi117
                        _ = v121
                        _ = v122
                    }
                    _ = v116
                } else {
                    var v123 bool = true
                    _ = v123
                }
            }
            _ = fi108
            _ = v109
        }
        var v124 uint = 124
        switch 392 {
        case 246:
            if (-1 != -1) {
                if v101 {
                    v106 += -1
                    v107 = v106
                } else {
                    if (!false) {
                        fmt.Println(v124)
                    } else {
                        if len(v3) > 0 {
                            v3[0] = 2
                        }
                        v124 = (v124 | (v124 * 42))
                    }
                }
                fmt.Println(v1, v101, v124)
            } else {
                var v125 int32 = 2147483647
                _ = v125
            }
        case 4294967295:
            v124 += 18446744073709551615
        case 20:
            if true {
                switch "" {
                case "a":
                    v101 = true
                case "foo bar":
                    v4 = (v4 << ((65535) & 63))
                case "":
                    if (v101 && (v4 <= v4)) {
                        v126 := GenS11{F0: 2147483647, F1: 0, F2: (("foo bar" + (v1 + v1)) + v1)}
                        v127 := GenS4{F0: 192, F1: v124}
                        switch "foo bar" {
                        case "foo bar":
                            v128 := GenS6{F0: 1, F1: (v4 + (v4 % (((-(v4))) | 1))), F2: 42}
                            _ = v128
                        default:
                            if ("a" <= v1) {
                                for fi129 := 0; fi129 < 20; fi129++ {
                                    for fi130 := 0; fi130 < 2; fi130++ {
                                        var v131 int8 = 42
                                        v107 += (v106 & v106)
                                        var v132 uint8 = 1
                                        _ = fi130
                                        _ = v131
                                        _ = v132
                                    }
                                    v133 := GenS3{F0: 1, F1: int8(v124)}
                                    _ = fi129
                                    _ = v133
                                }
                            } else {
                                if (-1 != v107) {
                                    for fi134 := 0; fi134 < 17; fi134++ {
                                        if v101 {
                                            if ((fi134 >> (((v124 | (v124 & 65535))) & 63)) == 127) {
                                                if (127 >= -117) {
                                                    var v135 string = v1
                                                    var v136 uint = v124
                                                    var v137 uint16 = genFn0()
                                                    _ = v135
                                                    _ = v136
                                                    _ = v137
                                                }
                                                v138 := GenS3{F0: 128, F1: 127}
                                                switch fi134 {
                                                case 127:
                                                    v139 := GenS3{F0: 128, F1: 42}
                                                    _ = v139
                                                case 0:
                                                    var v140 uint32 = 1
                                                    _ = v140
                                                case -32:
                                                    v141 := GenS5{F0: 65535}
                                                    _ = v141
                                                default:
                                                    switch ((v1 + (v1 + ("hello" + v1))) + v1) {
                                                    case "a":
                                                        var v142 float32 = 3.14159
                                                        _ = v142
                                                    case "":
                                                        v107 = v107
                                                    case "hello":
                                                        for fi143 := 0; fi143 < 10; fi143++ {
                                                            fmt.Println(v124)
                                                            _ = fi143
                                                        }
                                                    default:
                                                        v101 = v101
                                                    }
                                                }
                                                _ = v138
                                            }
                                            v124 = v124
                                        }
                                        fmt.Println(v101, v124, v106)
                                        _ = fi134
                                    }
                                    var v144 int32 = v106
                                    var v145 int32 = v107
                                    _ = v144
                                    _ = v145
                                }
                                for fi146 := 0; fi146 < 15; fi146++ {
                                    v2[(fi146 - 9223372036854775807)] = v101
                                    fmt.Println(fi146, v106, v4)
                                    _ = fi146
                                }
                            }
                        }
                        _ = v126
                        _ = v127
                    } else {
                        for fi147 := 0; fi147 < 11; fi147++ {
                            v124 = 0
                            _ = fi147
                        }
                        fmt.Println(v106)
                        _ = len(v3)
                    }
                default:
                    var v148 int8 = 42
                    _ = v148
                }
                if (127 < 2) {
                    var v149 uint8 = 127
                    for fi150 := 0; fi150 < 17; fi150++ {
                        fmt.Println(v106, v101, v101)
                        v1 = (("" + ("foo bar" + ("\n\t" + v1))) + "a")
                        if v101 {
                            if v101 {
                                if v101 {
                                    var v151 float64 = 3.14159
                                    var v152 int16 = v4
                                    v107 = v107
                                    _ = v151
                                    _ = v152
                                } else {
                                    for fi153 := 0; fi153 < 5; fi153++ {
                                        var v154 int = 4294967295
                                        _ = fi153
                                        _ = v154
                                    }
                                    if ((v4 - (v4 & (-(v4)))) > v4) {
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v155 uint16 = 42
                                            _ = v155
                                        case "a":
                                            var v156 int = fi150
                                            _ = v156
                                        default:
                                            if v101 {
                                                delete(v2, (^(fi150)))
                                            } else {
                                                var v157 bool = false
                                                fmt.Println(v4, v101, v157)
                                                _ = v157
                                            }
                                        }
                                        var v158 uint = (v124 | 42)
                                        _ = v158
                                    }
                                }
                            }
                            var v159 int64 = 351
                            delete(v2, (fi150 / (((^(fi150))) | 1)))
                            _ = v159
                        }
                        _ = fi150
                    }
                    if len(v3) > 0 {
                        v3[0] = 0
                    }
                    _ = v149
                }
            } else {
                fmt.Println(v107)
                switch "hello" {
                case "hello":
                    v101 = ((v124 << (((v124 % (((v124 >> ((v124) & 63))) | 1))) & 63)) == (v124 * (v124 - v124)))
                case "foo bar":
                    for fi160 := 0; fi160 < 6; fi160++ {
                        delete(v2, (fi160 ^ fi160))
                        if v101 {
                            for fi161 := 0; fi161 < 7; fi161++ {
                                var v162 float32 = 1e10
                                switch v1 {
                                case "\n\t":
                                    var v163 uint8 = 0
                                    _ = v163
                                case "":
                                    v124 += (^(v124))
                                }
                                var v164 bool = true
                                var v165 int8 = 127
                                _ = fi161
                                _ = v162
                                _ = v164
                                _ = v165
                            }
                            v107 = (v106 * v106)
                            for fi166 := 0; fi166 < 14; fi166++ {
                                var v167 bool = (!((fi160 % ((fi166) | 1)) > fi166))
                                var v168 bool = v167
                                if (0 != (v106 << (((v124 >> ((v124) & 63))) & 63))) {
                                    fmt.Println(v1)
                                    v169 := GenS3{F0: 255, F1: 100}
                                    fmt.Println(v168)
                                    _ = v169
                                }
                                _ = fi166
                                _ = v167
                                _ = v168
                            }
                        }
                        _ = fi160
                    }
                case "a":
                    _ = len(v2)
                }
            }
        default:
            switch 128 {
            case 32767:
                if v101 {
                    v4 -= 571
                    v124 = 4294967295
                } else {
                    v4 = v4
                    for fi170 := 0; fi170 < 2; fi170++ {
                        v2[2147483647] = (((1 > v124) || true) && false)
                        var v171 int32 = 0
                        _ = fi170
                        _ = v171
                    }
                }
            default:
                var v172 uint64 = 4294967295
                _ = v172
            }
        }
        _ = v124
    }
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v106:", v106)
    fmt.Println("v107:", v107)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v101
    _ = v105
    _ = v106
    _ = v107
    _ = v2
    _ = v3
    _ = v4
}
