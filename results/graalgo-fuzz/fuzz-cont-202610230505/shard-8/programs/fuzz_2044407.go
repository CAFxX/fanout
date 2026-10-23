package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 uint16; F1 int8 }
type GenS7 struct { F0 uint64; F1 string }
type GenS8 struct { F0 int; F1 string; F2 bool }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS11 struct { F0 uint32; F1 int; F2 uint8 }
type GenS12 struct { F0 int; F1 int16; F2 uint32 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint8; F1 uint64 }
type GenS16 struct { F0 uint16; F1 int16; F2 string }
type GenS17 struct { F0 int; F1 int16; F2 string }
type GenS18 struct { F0 uint; F1 int8 }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 int8; F1 uint8; F2 int16 }
type GenS21 struct { F0 uint8 }
type GenS22 struct { F0 int16; F1 int16 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    var v3 int = -9223372036854775808
    var v4 bool = (-32768 != 0)
    switch "hello" {
    case "":
        v4 = v4
    default:
        v3 = -2147483648
    }
    fmt.Println(v3, v3)
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := GenS1{F0: 0, F1: 869, F2: 42}
        if (1 == 42) {
            v4 = (v4 && ((v4 && v4) || (907 != 18446744073709551615)))
        }
        _ = fi5
        _ = v6
    }
    v2["foo bar"] = 4294967295
    v4 = v4
    for fi7 := 0; fi7 < 16; fi7++ {
        if v4 {
            fmt.Println(v4)
            var v8 int = (v3 + (fi7 << ((618) & 63)))
            _ = v8
        }
        _ = fi7
    }
    if ((0.0 + float32((0.0 / float64((1.5 + ((((0.0 * (1e10 / 1.0)) / 3.14159) - 1e10) - (((float32((0.1 * 1.5)) + 1e10) + 0.1) + float32(((-2.25 + float64(1e10)) * 1.5))))))))) > float32(((((-2.25 - 3.14159) * 0.0) + 0.0) / 1.0))) {
        var v9 int8 = -128
        _ = v9
    } else {
        if v4 {
            v10 := GenS2{F0: 0}
            _ = v10
        } else {
            if (!v4) {
                v11 := GenS1{F0: -2147483648, F1: 705, F2: 9223372036854775807}
                _ = v11
            } else {
                switch "a" {
                case "\n\t":
                    fmt.Println(v3, v3)
                case "":
                    v12 := v1[(v3 ^ (v3 - -2147483648))]
                    _ = v12
                case "foo bar":
                    v13 := GenS3{F0: 42, F1: 0, F2: 171}
                    _ = v13
                }
                switch ("hello" + "") {
                case "a":
                    fmt.Println(v4)
                case "foo bar":
                    v4 = (4294967295 < 2147483648)
                case "hello":
                    for fi14 := 0; fi14 < 2; fi14++ {
                        v15 := GenS3{F0: 658, F1: 4294967295, F2: 0}
                        fmt.Println(v3, fi14, fi14)
                        _ = fi14
                        _ = v15
                    }
                }
                _ = len(v2)
            }
            switch "" {
            case "a":
                v16 := GenS4{F0: 1}
                _ = v16
            case "\n\t":
                if (((738 == 740) && (!v4)) || v4) {
                    for fi17 := 0; fi17 < 2; fi17++ {
                        if v4 {
                            var v18 int = (v3 + (v3 | (fi17 * (fi17 | -2147483648))))
                            for fi19 := 0; fi19 < 16; fi19++ {
                                if ((v4 && (((((0.0 * 3.14159) / ((1.5 + -2.25) - float64(3.14159))) / 1.5) > 1e10) || (("" + "foo bar") == "\n\t"))) || v4) {
                                    var v20 uint32 = 2147483647
                                    v18 = fi17
                                    _ = v20
                                } else {
                                    if true {
                                        if ((float32((((3.14159 * 0.0) + ((0.0 / 0.1) - 0.0)) - 0.0)) * -2.25) == (0.0 - -2.25)) {
                                            v4 = v4
                                            v18 = v18
                                            v21 := GenS5{F0: 1, F1: 1}
                                            _ = v21
                                        } else {
                                            var v22 float64 = 1e10
                                            v23 := []int{fi17, 128, fi19}
                                            if false {
                                                v24 := GenS6{F0: 2, F1: 42}
                                                v25 := GenS7{F0: 867, F1: "a"}
                                                v3 = 42
                                                _ = v24
                                                _ = v25
                                            } else {
                                                fmt.Println(fi19, v4)
                                            }
                                            _ = v22
                                            _ = v23
                                        }
                                    } else {
                                        var v26 uint8 = 66
                                        switch 4294967295 {
                                        case 4294967295:
                                            if true {
                                                v27 := []int{(v3 - fi19), (v18 % (((fi19 ^ -824)) | 1)), (fi17 % ((597) | 1))}
                                                v3 += (-(v18))
                                                _ = v27
                                            }
                                        case 128:
                                            v28 := v1[127]
                                            _ = v28
                                        case 32767:
                                            _ = len(v1)
                                        default:
                                            v29 := GenS7{F0: 266, F1: "\n\t"}
                                            _ = v29
                                        }
                                        _ = v26
                                    }
                                    v30 := GenS3{F0: 1, F1: 42, F2: 246}
                                    _ = v30
                                }
                                _ = fi19
                            }
                            _ = v18
                        } else {
                            for fi31 := 0; fi31 < 13; fi31++ {
                                fmt.Println(v3)
                                v32 := GenS7{F0: 856, F1: ("foo bar" + "foo bar")}
                                _ = fi31
                                _ = v32
                            }
                            var v33 uint32 = 269
                            _ = v33
                        }
                        v3 = 4294967295
                        _ = fi17
                    }
                    switch v3 {
                    case -615:
                        v34 := GenS8{F0: (v3 ^ v3), F1: ("" + ("hello" + "\n\t")), F2: v4}
                        _ = v34
                    case 249:
                        v35 := GenS5{F0: 42, F1: 743}
                        _ = v35
                    }
                    switch "foo bar" {
                    case "":
                        var v36 int = v3
                        _ = v36
                    case "a":
                        for fi37 := 0; fi37 < 17; fi37++ {
                            v38 := []int{(fi37 & 1), 991, v3, -2147483648}
                            for fi39 := 0; fi39 < 11; fi39++ {
                                var v40 float64 = float64((-2.25 + 1.5))
                                if (947 > -32768) {
                                    if (1000 != 42) {
                                        var v41 bool = false
                                        v42 := GenS4{F0: 1}
                                        v43 := GenS9{F0: -1}
                                        if v4 {
                                            var v44 bool = (0.1 <= float32((((v40 - v40) + (0.0 + v40)) / v40)))
                                            var v45 float32 = (0.0 / 1e10)
                                            v41 = ((42 == 631) || (true && (42 >= 364)))
                                            _ = v44
                                            _ = v45
                                        } else {
                                            v3 = fi39
                                            fmt.Println(v41, v3)
                                            v41 = v41
                                        }
                                        var v46 bool = (0 <= 100)
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                        _ = v46
                                    } else {
                                        switch -258 {
                                        case -244:
                                            v47 := GenS7{F0: 4294967295, F1: ("\n\t" + "")}
                                            _ = v47
                                        case 42:
                                            _ = len(v38)
                                        default:
                                            switch "foo bar" {
                                            case "hello":
                                                if (32767 == 65535) {
                                                    v48 := GenS6{F0: 32767, F1: -1}
                                                    _ = v48
                                                } else {
                                                    if (v4 && (1 <= 1)) {
                                                        v49 := v2["\n\t"]
                                                        if (!(("a" + "") == "a")) {
                                                            fmt.Println(fi37, v4)
                                                        } else {
                                                            v50 := GenS10{F0: 4294967295, F1: v4, F2: 220}
                                                            _ = v50
                                                        }
                                                        var v51 uint = 4294967295
                                                        _ = v49
                                                        _ = v51
                                                    } else {
                                                        var v52 uint16 = 0
                                                        _ = v52
                                                    }
                                                }
                                                if v4 {
                                                    fmt.Println(fi39)
                                                }
                                                switch fi37 {
                                                case 42:
                                                    var v53 bool = (1e10 == 0.1)
                                                    _ = v53
                                                default:
                                                    v40 -= -2.25
                                                }
                                                v40 = v40
                                                var v54 uint16 = 255
                                                _ = v54
                                            case "a":
                                                v1[(^(fi39))] = (fi37 % ((0) | 1))
                                            case "\n\t":
                                                if true {
                                                    switch "a" {
                                                    case "\n\t":
                                                        if (60 < 2) {
                                                            delete(v1, 128)
                                                            var v55 uint8 = 128
                                                            var v56 uint8 = 204
                                                            _ = v55
                                                            _ = v56
                                                        } else {
                                                            if (510 < 13) {
                                                                fmt.Println(v4)
                                                            }
                                                            _ = len(v1)
                                                            fmt.Println(fi39, v4, v4)
                                                        }
                                                    }
                                                    v57 := GenS9{F0: 1}
                                                    if v4 {
                                                        v58 := GenS11{F0: 65535, F1: (v3 ^ -2147483648), F2: 99}
                                                        v40 += v40
                                                        _ = v58
                                                    } else {
                                                        switch (v3 / ((fi37) | 1)) {
                                                        case -186:
                                                            if false {
                                                                v3 = 32767
                                                                var v59 string = "\n\t"
                                                                var v60 uint8 = 42
                                                                v61 := GenS12{F0: (fi39 * (fi37 - -732)), F1: -642, F2: 874}
                                                                var v62 uint = 42
                                                                _ = v59
                                                                _ = v60
                                                                _ = v61
                                                                _ = v62
                                                            }
                                                        case -379:
                                                            var v63 int8 = 60
                                                            _ = v63
                                                        case -1:
                                                            var v64 uint64 = 125
                                                            _ = v64
                                                        default:
                                                            v65 := GenS12{F0: (fi39 & fi39), F1: -32768, F2: 2147483647}
                                                            _ = v65
                                                        }
                                                        var v66 uint8 = 128
                                                        v67 := GenS13{F0: -962}
                                                        var v68 uint64 = 4294967295
                                                        var v69 float64 = v40
                                                        _ = v66
                                                        _ = v67
                                                        _ = v68
                                                        _ = v69
                                                    }
                                                    _ = v57
                                                } else {
                                                    v70 := GenS14{F0: -1}
                                                    _ = v70
                                                }
                                            default:
                                                _ = len(v2)
                                                var v71 uint = 0
                                                if (!(!(!false))) {
                                                    fmt.Println(v3, v71, fi37)
                                                    if len(v38) > 0 {
                                                        v38[0] = fi39
                                                    }
                                                    v40 += (v40 / (v40 - v40))
                                                } else {
                                                    var v72 int16 = 1000
                                                    v3 += -100
                                                    _ = v72
                                                }
                                                _ = v71
                                            }
                                        }
                                        switch ("a" + "\n\t") {
                                        case "\n\t":
                                            v4 = true
                                        case "":
                                            switch "" {
                                            case "a":
                                                if (128 <= 2) {
                                                    v4 = ((v3 - (v3 + (fi39 & v3))) != 127)
                                                } else {
                                                    v3 += (-(v3))
                                                    var v73 string = "a"
                                                    fmt.Println(v3, v4, v3)
                                                    _ = v73
                                                }
                                            }
                                        }
                                        switch (("a" + (("hello" + "\n\t") + (((("hello" + ("a" + "a")) + "foo bar") + "a") + "\n\t"))) + "hello") {
                                        case "foo bar":
                                            fmt.Println(v4)
                                        case "":
                                            var v74 int16 = 32767
                                            _ = v74
                                        }
                                    }
                                } else {
                                    if (!(!((!v4) || (1 != 42)))) {
                                        switch (v3 >> ((255) & 63)) {
                                        case -1:
                                            fmt.Println(v40)
                                        case 539:
                                            v75 := GenS9{F0: -9223372036854775808}
                                            _ = v75
                                        }
                                        v40 -= (v40 - ((-2.25 + v40) + v40))
                                    }
                                    var v76 uint64 = 1
                                    _ = v76
                                }
                                v3 = (-(fi39))
                                _ = fi39
                                _ = v40
                            }
                            if (!v4) {
                                v77 := GenS9{F0: -9223372036854775808}
                                _ = v77
                            } else {
                                for fi78 := 0; fi78 < 19; fi78++ {
                                    v4 = v4
                                    _ = fi78
                                }
                                if len(v38) > 0 {
                                    v38[0] = (fi37 >> ((609) & 63))
                                }
                                v3 = v3
                            }
                            _ = fi37
                            _ = v38
                        }
                    case "foo bar":
                        var v79 string = ((("\n\t" + "") + "foo bar") + ("\n\t" + "\n\t"))
                        _ = v79
                    }
                }
            case "hello":
                for fi80 := 0; fi80 < 7; fi80++ {
                    for fi81 := 0; fi81 < 8; fi81++ {
                        if (false && false) {
                            fmt.Println(v4, v4, v4)
                            switch (("hello" + "foo bar") + "\n\t") {
                            case "a":
                                v3 = fi81
                            }
                        } else {
                            var v82 int8 = 1
                            var v83 string = ("\n\t" + "")
                            v84 := GenS7{F0: 2, F1: v83}
                            fmt.Println(v82, fi80)
                            _ = v82
                            _ = v83
                            _ = v84
                        }
                        v4 = (-1 != 1)
                        _ = fi81
                    }
                    v4 = v4
                    _ = fi80
                }
            default:
                v85 := GenS5{F0: 32767, F1: 400}
                _ = v85
            }
            var v86 bool = v4
            _ = v86
        }
    }
    var v87 uint = 65535
    fmt.Println(v87, v3)
    var v88 string = "a"
    v87 = (v87 << ((v87) & 63))
    v4 = (0 != -211)
    if (11 != 15) {
        var v89 float64 = 0.0
        _ = v89
    }
    for fi90 := 0; fi90 < 10; fi90++ {
        for fi91 := 0; fi91 < 12; fi91++ {
            var v92 uint8 = 128
            if (true || (461 <= 0)) {
                switch v3 {
                case -401:
                    v93 := GenS15{F0: (v92 >> (((v87 + (v87 << ((v87) & 63)))) & 63)), F1: 2}
                    _ = v93
                case 32767:
                    var v94 float64 = 3.14159
                    v88 = (v88 + ("foo bar" + (v88 + v88)))
                    if true {
                        v95 := GenS16{F0: 42, F1: -35, F2: v88}
                        var v96 int16 = 42
                        if (1 != 42) {
                            if (true && (0 != 9223372036854775807)) {
                                v88 = ((((v88 + (((v88 + (v88 + "foo bar")) + "hello") + ("" + v88))) + v88) + v88) + "\n\t")
                            } else {
                                switch 110 {
                                case -589:
                                    var v97 uint64 = 434
                                    _ = v97
                                case 755:
                                    if (1e10 != (-2.25 / 1.5)) {
                                        v98 := GenS6{F0: 42, F1: 0}
                                        v92 -= v92
                                        v99 := GenS17{F0: (-(fi90)), F1: -47, F2: v88}
                                        _ = v98
                                        _ = v99
                                    }
                                default:
                                    if v4 {
                                        switch 42 {
                                        case 0:
                                            var v100 bool = v4
                                            _ = v100
                                        case 32767:
                                            fmt.Println(v96, v87)
                                        default:
                                            switch "hello" {
                                            case "foo bar":
                                                var v101 string = (v88 + ("a" + "foo bar"))
                                                _ = v101
                                            case "a":
                                                if (18446744073709551615 <= (^(v87))) {
                                                    v102 := GenS14{F0: int16(fi91)}
                                                    _ = v102
                                                }
                                            default:
                                                v103 := []int{(v3 | 4294967295), (-(fi91)), v3, (fi91 + (v3 ^ -2147483648))}
                                                _ = v103
                                            }
                                        }
                                        if (!v4) {
                                            v3 -= (v3 >> ((0) & 63))
                                        }
                                        var v104 uint8 = (^(v92))
                                        var v105 int8 = 4
                                        var v106 uint16 = 360
                                        _ = v104
                                        _ = v105
                                        _ = v106
                                    }
                                }
                                if (((v96 + (-(v96))) <= (^(v96))) && v4) {
                                    var v107 uint = v87
                                    var v108 string = v88
                                    fmt.Println(v92, v96)
                                    _ = v107
                                    _ = v108
                                } else {
                                    v1[(fi91 % ((9223372036854775807) | 1))] = v3
                                    var v109 int64 = -607
                                    _ = v109
                                }
                            }
                            var v110 uint = v87
                            _ = v110
                        } else {
                            if false {
                                v3 -= -771
                                if v4 {
                                    var v111 uint16 = 65535
                                    v112 := GenS18{F0: (v87 + v87), F1: -42}
                                    if false {
                                        if v4 {
                                            var v113 int64 = 9223372036854775807
                                            fmt.Println(v113)
                                            v114 := GenS14{F0: (v96 % (((v96 | v96)) | 1))}
                                            _ = v113
                                            _ = v114
                                        }
                                        v111 -= 985
                                    } else {
                                        var v115 uint32 = 0
                                        delete(v2, v88)
                                        _ = v115
                                    }
                                    v116 := GenS19{F0: v4}
                                    v88 = ""
                                    var v117 bool = (-128 > 1)
                                    _ = v111
                                    _ = v112
                                    _ = v116
                                    _ = v117
                                }
                            } else {
                                var v118 int8 = 0
                                _ = v118
                            }
                            var v119 int64 = 1
                            v120 := GenS9{F0: v119}
                            _ = v119
                            _ = v120
                        }
                        _ = v95
                        _ = v96
                    }
                    _ = v94
                case 211:
                    delete(v2, "hello")
                default:
                    if (!(1.5 != -2.25)) {
                        var v121 int64 = 1
                        if false {
                            switch v88 {
                            case "a":
                                switch (fi91 / (((-(fi90))) | 1)) {
                                case 127:
                                    if (-2.25 <= (3.14159 + 0.1)) {
                                        v4 = v4
                                        if v4 {
                                            if (float64(-2.25) == (1e10 + 1e10)) {
                                                v122 := GenS4{F0: 2147483647}
                                                v2[("\n\t" + v88)] = 42
                                                var v123 int16 = -32768
                                                _ = v122
                                                _ = v123
                                            } else {
                                                v4 = v4
                                            }
                                            switch ("\n\t" + v88) {
                                            case "hello":
                                                var v124 int8 = -1
                                                _ = v124
                                            case "a":
                                                v125 := []int{int(v121)}
                                                var v126 uint8 = 183
                                                _ = v125
                                                _ = v126
                                            default:
                                                var v127 uint16 = 0
                                                _ = v127
                                            }
                                            switch fi91 {
                                            case 32768:
                                                v92 += (v92 % (((^(v92))) | 1))
                                            case 9223372036854775807:
                                                var v128 uint64 = 0
                                                _ = v128
                                            }
                                        }
                                    }
                                case -2147483648:
                                    fmt.Println(v3, v4)
                                case 4294967295:
                                    v129 := []int{}
                                    _ = v129
                                default:
                                    v1[fi91] = fi91
                                }
                            case "\n\t":
                                switch fi91 {
                                case 91:
                                    fmt.Println(v121, v4)
                                default:
                                    v92 = v92
                                }
                            default:
                                var v130 int = fi91
                                _ = v130
                            }
                            var v131 int8 = 42
                            if v4 {
                                var v132 float64 = (3.14159 / (3.14159 + 0.1))
                                fmt.Println(v4, fi91)
                                var v133 int8 = v131
                                _ = v132
                                _ = v133
                            }
                            _ = v131
                        }
                        _ = v121
                    } else {
                        v134 := GenS20{F0: -128, F1: (v92 - (v92 | (v92 | v92))), F2: -32768}
                        switch -1 {
                        case 0:
                            v135 := GenS13{F0: 0}
                            _ = v135
                        case 688:
                            fmt.Println(v88)
                        case 32767:
                            var v136 float64 = 1.5
                            _ = v136
                        default:
                            v4 = (-9223372036854775808 > 1)
                        }
                        v137 := GenS21{F0: v92}
                        _ = v134
                        _ = v137
                    }
                }
            } else {
                v87 += (v87 - 42)
                switch fi90 {
                case 42:
                    v138 := GenS22{F0: 0, F1: 32767}
                    _ = v138
                default:
                    var v139 string = "hello"
                    _ = v139
                }
            }
            _ = fi91
            _ = v92
        }
        for fi140 := 0; fi140 < 11; fi140++ {
            v3 = 32767
            v141 := v2[(v88 + ("\n\t" + ""))]
            _ = fi140
            _ = v141
        }
        _ = fi90
    }
    var v142 float64 = 1.5
    v87 = v87
    switch v88 {
    case "":
        for fi143 := 0; fi143 < 2; fi143++ {
            v3 = (fi143 ^ -652)
            _ = fi143
        }
    case "\n\t":
        v87 = (v87 + 399)
    }
    var v144 int32 = int32(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v142:", v142)
    fmt.Println("v144:", v144)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v142
    _ = v144
    _ = v2
    _ = v3
    _ = v4
    _ = v87
    _ = v88
}
