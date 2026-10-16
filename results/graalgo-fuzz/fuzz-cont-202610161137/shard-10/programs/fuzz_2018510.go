package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint64 }
type GenS2 struct { F0 int8; F1 int64; F2 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 string }
type GenS10 struct { F0 string; F1 int16; F2 uint32 }
type GenS11 struct { F0 uint16; F1 int64 }
type GenS12 struct { F0 uint; F1 uint8 }
type GenS13 struct { F0 int64; F1 uint8; F2 int8 }
type GenS14 struct { F0 uint32; F1 int32 }
type GenS15 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS16 struct { F0 uint8; F1 int; F2 int64 }
type GenS17 struct { F0 int8; F1 int }
type GenS18 struct { F0 int16; F1 uint8; F2 int16 }
type GenS19 struct { F0 int8; F1 int }
type GenS20 struct { F0 uint32 }
type GenS21 struct { F0 int64 }
type GenS22 struct { F0 uint8; F1 uint16 }
type GenS23 struct { F0 uint64; F1 int32 }
type GenS24 struct { F0 bool }
type GenS25 struct { F0 uint8 }
type GenS26 struct { F0 uint8; F1 int64 }
type GenS27 struct { F0 int }
func genFn0(p0 bool, p1 string) string {
    if (123 == -128) {
        v6 := GenS1{F0: 9223372036854775807, F1: 545, F2: 4294967295}
        _ = v6
    }
    if (157 >= 16) {
        p0 = false
    }
    v7 := []int{}
    _ = p0
    _ = p1
    _ = v7
    return "hello"
}

func main() {
    var v1 uint16 = 626
    v2 := []uint32{0}
    v3 := make(map[uint32]string)
    v4 := make(map[string]int)
    v5 := make(map[uint32]string)
    v1 += (v1 / ((v1) | 1))
    v8 := GenS1{F0: 0, F1: 962, F2: 1}
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 uint16 = 255
        v10 += (v10 | v10)
        _ = fi9
        _ = v10
    }
    var v11 int8 = 85
    switch 2147483647 {
    case 42:
        if true {
            for fi12 := 0; fi12 < 2; fi12++ {
                if (false || true) {
                    delete(v4, "foo bar")
                } else {
                    v1 = (v1 - 248)
                    v11 += v11
                    v11 -= (^(v11))
                }
                v13 := GenS1{F0: 9223372036854775807, F1: 840, F2: 4294967295}
                _ = fi12
                _ = v13
            }
        } else {
            v1 += 255
            v11 = (v11 + (^(v11)))
            v14 := GenS2{F0: (v11 * 127), F1: 9223372036854775807, F2: "foo bar"}
            _ = v14
        }
    case -9223372036854775808:
        v15 := GenS2{F0: (-(v11)), F1: -465, F2: "\n\t"}
        _ = v15
    case -169:
        for fi16 := 0; fi16 < 2; fi16++ {
            v17 := GenS3{F0: 65535}
            v18 := GenS1{F0: 9223372036854775807, F1: -202, F2: 223}
            if false {
                var v19 uint8 = 1
                _ = v19
            } else {
                var v20 int = (fi16 * (fi16 / ((-270) | 1)))
                if (127 >= 50) {
                    if (!false) {
                        var v21 string = ("foo bar" + "foo bar")
                        _ = v21
                    }
                } else {
                    switch "hello" {
                    case "hello":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "foo bar":
                        var v22 uint64 = 157
                        _ = v22
                    case "":
                        switch "a" {
                        case "":
                            for fi23 := 0; fi23 < 13; fi23++ {
                                v1 -= v1
                                if false {
                                    var v24 int = fi23
                                    switch -668 {
                                    case 0:
                                        if false {
                                            var v25 float32 = float32((((0.0 - ((float64(0.1) / 1.0) + ((((0.0 + 1e10) / 1.0) / 1e10) + -2.25))) * (3.14159 / 0.1)) + ((1e10 * (0.1 * (-2.25 + 0.1))) + 1.5)))
                                            _ = v25
                                        }
                                    case 32768:
                                        v11 = (v11 << ((65535) & 63))
                                    case 1:
                                        v20 = v20
                                    }
                                    _ = v24
                                }
                                _ = len(v5)
                                v20 -= 9223372036854775807
                                fmt.Println(v20)
                                _ = fi23
                            }
                        default:
                            var v26 int16 = 1000
                            _ = v26
                        }
                    default:
                        var v27 int = v20
                        _ = v27
                    }
                    v11 += (v11 * v11)
                    var v28 int = v20
                    _ = v28
                }
                fmt.Println(v11, v1, fi16)
                _ = v20
            }
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    v29 := GenS2{F0: v11, F1: 9223372036854775807, F2: "a"}
    _ = len(v2)
    if ((!true) && (127 <= 205)) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi30 := 0; fi30 < 4; fi30++ {
            for fi31 := 0; fi31 < 7; fi31++ {
                var v32 int16 = 32767
                v1 = 42
                switch (fi31 * (fi31 * fi30)) {
                case 0:
                    v11 -= (v11 << ((42) & 63))
                default:
                    _ = len(v4)
                    if len(v2) > 0 {
                        v2[0] = 214
                    }
                }
                _ = fi31
                _ = v32
            }
            var v33 uint16 = v1
            _ = fi30
            _ = v33
        }
        v11 = v11
    } else {
        v11 = (v11 | v11)
        if true {
            var v34 uint = 42
            v2 = append(v2, 2147483647)
            var v35 bool = ((false && ((42 <= -1) && true)) || false)
            _ = v34
            _ = v35
        }
        for fi36 := 0; fi36 < 17; fi36++ {
            if (("a" + "") == ("" + "a")) {
                if false {
                    var v37 float32 = (((((-2.25 + float32(((1.5 / -2.25) / 1e10))) - float32(3.14159)) - (3.14159 + (1e10 - -2.25))) - 3.14159) + 1.5)
                    if true {
                        v1 = (v1 | 2)
                        switch fi36 {
                        case -9223372036854775808:
                            v11 -= (v11 - v11)
                        case -605:
                            v1 += 417
                        case 32768:
                            v37 += ((v37 / float32(((((float64(v37) * 0.0) * 3.14159) * -2.25) - (float64(v37) * 0.1)))) - v37)
                        default:
                            var v38 int8 = v11
                            _ = v38
                        }
                    } else {
                        var v39 int16 = 307
                        var v40 uint8 = 255
                        _ = v39
                        _ = v40
                    }
                    for fi41 := 0; fi41 < 10; fi41++ {
                        v11 += (v11 % (((v11 | (^(v11)))) | 1))
                        _ = fi41
                    }
                    _ = v37
                }
                var v42 int8 = v11
                for fi43 := 0; fi43 < 5; fi43++ {
                    if true {
                        if ((127 < 101) || (false && (0.0 == 1.5))) {
                            v42 = v11
                        } else {
                            v44 := GenS1{F0: -1, F1: 600, F2: 274}
                            if (!false) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                fmt.Println(v11, v1)
                                var v45 int64 = -355
                                fmt.Println(fi36)
                                _ = v45
                            }
                            v46 := GenS4{F0: 0}
                            _ = v44
                            _ = v46
                        }
                    }
                    if (2147483648 <= 212) {
                        v47 := GenS5{F0: "hello", F1: 42}
                        switch ("\n\t" + "") {
                        case "\n\t":
                            fmt.Println(fi36)
                        case "":
                            if false {
                                if false {
                                    v48 := GenS2{F0: (v11 & (v11 | v42)), F1: 9223372036854775807, F2: "foo bar"}
                                    var v49 int32 = 917
                                    _ = v48
                                    _ = v49
                                } else {
                                    v50 := GenS6{F0: (fi36 | (fi36 >> ((965) & 63)))}
                                    _ = v50
                                }
                                fmt.Println(v42, v1)
                                switch (fi43 & fi43) {
                                case -1:
                                    if false {
                                        _ = len(v5)
                                        v51 := GenS7{F0: 2, F1: 2, F2: -72}
                                        if false {
                                            v52 := GenS8{F0: "hello"}
                                            if true {
                                                v53 := GenS1{F0: 42, F1: 152, F2: 2}
                                                if true {
                                                    v54 := GenS4{F0: 1}
                                                    v55 := GenS4{F0: -9223372036854775808}
                                                    _ = v54
                                                    _ = v55
                                                } else {
                                                    if false {
                                                        if (!true) {
                                                            if ("\n\t" <= "a") {
                                                                var v56 uint = 65535
                                                                v57 := GenS9{F0: "foo bar"}
                                                                v58 := GenS10{F0: "\n\t", F1: 42, F2: 1}
                                                                _ = v56
                                                                _ = v57
                                                                _ = v58
                                                            } else {
                                                                v1 = (v1 % ((v1) | 1))
                                                            }
                                                        }
                                                        if false {
                                                            var v59 int32 = 42
                                                            v11 = 127
                                                            _ = v59
                                                        }
                                                    } else {
                                                        v1 += (v1 | (v1 ^ 737))
                                                        var v60 float64 = 0.0
                                                        var v61 uint8 = 2
                                                        var v62 int16 = -317
                                                        v60 = (0.1 / v60)
                                                        _ = v60
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                    var v63 int = fi36
                                                    if (true || (2 != 2)) {
                                                        var v64 bool = (0.1 != 1e10)
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        v64 = (-663 == 1)
                                                        switch ("\n\t" + "hello") {
                                                        case "":
                                                            v11 += (v42 * v11)
                                                        case "hello":
                                                            switch "hello" {
                                                            case "a":
                                                                var v65 bool = (!false)
                                                                _ = v65
                                                            default:
                                                                switch v63 {
                                                                case 32767:
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                case 758:
                                                                    fmt.Println(v1)
                                                                }
                                                            }
                                                        case "\n\t":
                                                            var v66 int16 = -32768
                                                            _ = v66
                                                        }
                                                        _ = v64
                                                    } else {
                                                        var v67 float64 = -2.25
                                                        if (true || false) {
                                                            v68 := GenS10{F0: "foo bar", F1: -351, F2: 65535}
                                                            if true {
                                                                v63 += fi43
                                                            } else {
                                                                v11 -= (v42 | (v11 ^ 26))
                                                            }
                                                            _ = v68
                                                        } else {
                                                            switch ("foo bar" + "\n\t") {
                                                            case "foo bar":
                                                                v69 := v5[583]
                                                                _ = v69
                                                            }
                                                            switch (fi43 * (v63 / (((fi36 * (fi36 ^ (-(v63))))) | 1))) {
                                                            case -181:
                                                                v70 := GenS11{F0: v1, F1: -9223372036854775808}
                                                                _ = v70
                                                            case 32768:
                                                                if true {
                                                                    v71 := GenS12{F0: 1, F1: 2}
                                                                    if true {
                                                                        v1 -= 2
                                                                        var v72 int = (v63 | (fi36 - (fi36 % ((4294967295) | 1))))
                                                                        if false {
                                                                            _ = len(v3)
                                                                            v72 -= 2147483647
                                                                            v1 = v1
                                                                        }
                                                                        delete(v3, 2147483648)
                                                                        _ = v72
                                                                    } else {
                                                                        if ((65535 <= 2) && true) {
                                                                            var v73 uint = 2
                                                                            if false {
                                                                                switch 993 {
                                                                                case 2147483647:
                                                                                    delete(v3, 1)
                                                                                case 128:
                                                                                    v2 = append(v2, 0)
                                                                                default:
                                                                                    var v74 int16 = -434
                                                                                    _ = len(v2)
                                                                                    v73 = (v73 % ((v73) | 1))
                                                                                    _ = v74
                                                                                }
                                                                            }
                                                                            var v75 bool = (0 <= 476)
                                                                            _ = v73
                                                                            _ = v75
                                                                        }
                                                                    }
                                                                    _ = v71
                                                                }
                                                            }
                                                            v42 = v42
                                                        }
                                                        var v76 int32 = 2147483647
                                                        switch 42 {
                                                        case 1:
                                                            switch "a" {
                                                            case "a":
                                                                var v77 float64 = v67
                                                                _ = v77
                                                            case "foo bar":
                                                                _ = len(v5)
                                                            case "hello":
                                                                if (!true) {
                                                                    var v78 bool = true
                                                                    v79 := GenS13{F0: 0, F1: 2, F2: v42}
                                                                    v80 := GenS14{F0: 42, F1: v76}
                                                                    _ = v78
                                                                    _ = v79
                                                                    _ = v80
                                                                }
                                                            }
                                                        default:
                                                            v81 := GenS2{F0: (^(v11)), F1: 608, F2: "foo bar"}
                                                            _ = v81
                                                        }
                                                        _ = v67
                                                        _ = v76
                                                    }
                                                    _ = v63
                                                }
                                                _ = v53
                                            } else {
                                                var v82 int16 = 42
                                                v11 = (v11 * (v11 + -63))
                                                v83 := GenS7{F0: 0, F1: 62, F2: -1}
                                                _ = v82
                                                _ = v83
                                            }
                                            _ = v52
                                        }
                                        v42 = v42
                                        _ = v51
                                    } else {
                                        v84 := GenS4{F0: -9223372036854775808}
                                        var v85 bool = (true && (!false))
                                        v86 := GenS15{F0: 4294967295, F1: 151, F2: 802}
                                        v87 := GenS5{F0: "\n\t", F1: 0}
                                        var v88 float64 = 0.1
                                        _ = v84
                                        _ = v85
                                        _ = v86
                                        _ = v87
                                        _ = v88
                                    }
                                case 32767:
                                    _ = len(v2)
                                case 9223372036854775807:
                                    v11 = (v42 - -128)
                                }
                            }
                        }
                        v89 := GenS2{F0: v11, F1: int64(v1), F2: "hello"}
                        var v90 int = (fi36 + (fi43 >> ((2) & 63)))
                        _ = v47
                        _ = v89
                        _ = v90
                    }
                    v91 := GenS2{F0: v11, F1: 445, F2: (("foo bar" + "\n\t") + "hello")}
                    v92 := GenS5{F0: "\n\t", F1: 9223372036854775807}
                    switch fi43 {
                    case 0:
                        switch (("" + ("" + ("" + "a"))) + "hello") {
                        case "a":
                            var v93 uint = 0
                            _ = v93
                        case "\n\t":
                            fmt.Println(v11)
                        case "hello":
                            var v94 uint16 = 705
                            _ = v94
                        default:
                            fmt.Println(v11, v42)
                        }
                    case -1:
                        v95 := GenS7{F0: 128, F1: 128, F2: 74}
                        _ = v95
                    }
                    _ = fi43
                    _ = v91
                    _ = v92
                }
                _ = v42
            } else {
                v96 := GenS13{F0: 1, F1: 86, F2: (v11 & v11)}
                v97 := GenS16{F0: 13, F1: (fi36 & fi36), F2: -1}
                _ = v96
                _ = v97
            }
            v98 := GenS17{F0: (v11 * v11), F1: fi36}
            v99 := v5[42]
            _ = fi36
            _ = v98
            _ = v99
        }
    }
    if false {
        if false {
            v1 = (v1 >> ((4294967295) & 63))
        }
        var v100 int8 = 100
        for fi101 := 0; fi101 < 14; fi101++ {
            if (true || false) {
                if ((v100 > (v11 & (v100 >> ((42) & 63)))) || ("a" <= ("" + "foo bar"))) {
                    for fi102 := 0; fi102 < 18; fi102++ {
                        var v103 int8 = (v100 % (((^(v11))) | 1))
                        if (42 != 1) {
                            v104 := GenS7{F0: 251, F1: 42, F2: v103}
                            if (false && false) {
                                if false {
                                    switch "" {
                                    case "\n\t":
                                        switch "\n\t" {
                                        case "hello":
                                            v100 = v100
                                        case "foo bar":
                                            var v105 float32 = 3.14159
                                            _ = v105
                                        case "":
                                            if (65535 == 853) {
                                                v106 := GenS18{F0: 0, F1: 2, F2: 1}
                                                v107 := GenS3{F0: 4294967295}
                                                var v108 int32 = 133
                                                _ = v106
                                                _ = v107
                                                _ = v108
                                            }
                                        }
                                    default:
                                        v109 := GenS8{F0: (("hello" + "foo bar") + ("foo bar" + "hello"))}
                                        _ = v109
                                    }
                                } else {
                                    if (true && (42 == 0)) {
                                        if true {
                                            v110 := GenS19{F0: (v100 * (v103 | (v100 / ((0) | 1)))), F1: fi101}
                                            var v111 int8 = v11
                                            if (-1 >= 9223372036854775807) {
                                                if (970 < 1) {
                                                    v111 = (^(v111))
                                                    v103 = v103
                                                    var v112 int32 = -55
                                                    _ = v112
                                                } else {
                                                    var v113 float64 = (-2.25 / 1.0)
                                                    _ = v113
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v114 string = ""
                                                _ = v114
                                            }
                                            var v115 uint16 = v1
                                            switch (fi101 + 0) {
                                            case 2147483647:
                                                var v116 uint = uint(v1)
                                                _ = v116
                                            case -611:
                                                if (false && false) {
                                                    var v117 float64 = (1e10 + 0.0)
                                                    var v118 uint = 4294967295
                                                    _ = v117
                                                    _ = v118
                                                }
                                            default:
                                                var v119 uint32 = 2147483648
                                                _ = v119
                                            }
                                            _ = v110
                                            _ = v111
                                            _ = v115
                                        } else {
                                            if false {
                                                v11 += -1
                                                var v120 int32 = -1
                                                if false {
                                                    var v121 int16 = 42
                                                    var v122 int64 = -432
                                                    _ = v121
                                                    _ = v122
                                                }
                                                _ = v120
                                            } else {
                                                fmt.Println(v100, fi101)
                                                var v123 uint = 4294967295
                                                _ = v123
                                            }
                                            var v124 int64 = 42
                                            var v125 uint16 = 0
                                            var v126 float64 = 1e10
                                            _ = v124
                                            _ = v125
                                            _ = v126
                                        }
                                    }
                                    var v127 uint = 65535
                                    _ = v127
                                }
                                v100 += (v100 * (-(v100)))
                                v128 := GenS5{F0: (((("a" + (("\n\t" + "hello") + "hello")) + "\n\t") + "") + "hello"), F1: 336}
                                _ = v128
                            }
                            v129 := GenS20{F0: 1}
                            _ = v104
                            _ = v129
                        } else {
                            v11 += v11
                            v103 += (v103 << ((42) & 63))
                            var v130 int8 = 127
                            if false {
                                var v131 uint32 = 1
                                _ = v131
                            } else {
                                if (42 < 0) {
                                    if false {
                                        switch fi102 {
                                        case -2147483648:
                                            v130 = (v100 - v100)
                                        case 32768:
                                            v1 += v1
                                        case 949:
                                            var v132 int16 = 1
                                            _ = v132
                                        }
                                        delete(v4, "a")
                                    }
                                    v11 = v11
                                    if false {
                                        var v133 float32 = 1.5
                                        _ = v133
                                    } else {
                                        if (v1 <= v1) {
                                            v134 := GenS21{F0: -130}
                                            if (!false) {
                                                var v135 bool = false
                                                var v136 uint8 = 1
                                                _ = v135
                                                _ = v136
                                            } else {
                                                v100 = (v100 >> ((2) & 63))
                                                if false {
                                                    var v137 int64 = 1
                                                    _ = v137
                                                } else {
                                                    var v138 float64 = (float64((3.14159 * ((0.1 - ((float32(0.1) + 0.0) / (-2.25 * (1e10 - (1e10 / (float32(0.1) / 3.14159)))))) + 3.14159))) + (float64(3.14159) - 3.14159))
                                                    if true {
                                                        var v139 string = "foo bar"
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        delete(v5, 65535)
                                                        switch fi102 {
                                                        case 336:
                                                            v138 = ((-2.25 / v138) - ((0.0 + 1e10) / 1.0))
                                                        }
                                                        delete(v3, 4294967295)
                                                        _ = v139
                                                    } else {
                                                        v130 -= 100
                                                        v140 := GenS22{F0: 42, F1: (v1 >> ((0) & 63))}
                                                        v2 = append(v2, 2147483648)
                                                        v11 += (v100 << ((2) & 63))
                                                        _ = v140
                                                    }
                                                    v141 := GenS2{F0: (-(v11)), F1: 42, F2: ("a" + ("a" + ("\n\t" + ("\n\t" + ""))))}
                                                    fmt.Println(v130, v11, v1)
                                                    _ = v138
                                                    _ = v141
                                                }
                                                v142 := v3[4294967295]
                                                _ = v142
                                            }
                                            _ = len(v2)
                                            _ = v134
                                        } else {
                                            var v143 uint = 0
                                            var v144 bool = (!false)
                                            v103 = v11
                                            _ = v143
                                            _ = v144
                                        }
                                        v103 = (v103 - -128)
                                        if false {
                                            v145 := GenS11{F0: (^(v1)), F1: 743}
                                            var v146 float64 = 0.1
                                            _ = v145
                                            _ = v146
                                        } else {
                                            v147 := GenS20{F0: 0}
                                            v11 -= 0
                                            if false {
                                                v11 = (^(v130))
                                                if (!((3.14159 == 1.5) && (1 < -1))) {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                } else {
                                                    v148 := GenS23{F0: 916, F1: 1}
                                                    _ = v148
                                                }
                                                if true {
                                                    switch ("\n\t" + ("\n\t" + "")) {
                                                    case "":
                                                        var v149 int8 = (^(v100))
                                                        _ = v149
                                                    default:
                                                        switch fi101 {
                                                        case 0:
                                                            v150 := GenS2{F0: int8(fi101), F1: -9223372036854775808, F2: "a"}
                                                            _ = v150
                                                        case 128:
                                                            switch (fi102 + -1) {
                                                            case 2147483647:
                                                                v1 += 42
                                                            case 32768:
                                                                v130 += 100
                                                            default:
                                                                delete(v5, 1)
                                                            }
                                                        case 9223372036854775807:
                                                            if len(v2) > 0 {
                                                                v2[0] = 0
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    var v151 float64 = (float64(1.5) / ((0.1 * 3.14159) - float64(-2.25)))
                                                    _ = v151
                                                }
                                            }
                                            _ = v147
                                        }
                                        var v152 uint64 = 770
                                        v153 := GenS24{F0: true}
                                        _ = v152
                                        _ = v153
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                } else {
                                    var v154 int32 = 441
                                    v155 := GenS18{F0: -885, F1: 1, F2: 56}
                                    _ = v154
                                    _ = v155
                                }
                                if ("\n\t" < "a") {
                                    v103 = -25
                                    v1 += 32767
                                } else {
                                    v100 += 42
                                    var v156 bool = (!(2 == 255))
                                    v130 -= v11
                                    _ = v156
                                }
                                v130 -= v130
                            }
                            switch -107 {
                            case 32767:
                                if true {
                                    var v157 uint32 = 2147483648
                                    if (319 != 42) {
                                        v158 := GenS19{F0: v100, F1: -610}
                                        v159 := GenS10{F0: "\n\t", F1: 990, F2: 42}
                                        v160 := GenS25{F0: 42}
                                        _ = v158
                                        _ = v159
                                        _ = v160
                                    }
                                    v4["\n\t"] = fi102
                                    var v161 float32 = float32(0.1)
                                    var v162 uint8 = 255
                                    fmt.Println(fi102, v1, fi101)
                                    _ = v157
                                    _ = v161
                                    _ = v162
                                } else {
                                    delete(v5, 0)
                                }
                            }
                            switch 2147483647 {
                            case -261:
                                v163 := GenS7{F0: 128, F1: 195, F2: v100}
                                _ = v163
                            case 32768:
                                if true {
                                    v164 := GenS26{F0: 0, F1: -529}
                                    _ = v164
                                }
                            default:
                                switch fi102 {
                                case 2147483647:
                                    switch "hello" {
                                    case "a":
                                        v4[("\n\t" + "foo bar")] = fi102
                                    }
                                case 4294967295:
                                    var v165 uint64 = 28
                                    _ = v165
                                case -933:
                                    v166 := GenS1{F0: 42, F1: -2147483648, F2: 1}
                                    _ = v166
                                default:
                                    v130 -= v100
                                }
                            }
                            var v167 uint8 = 2
                            _ = v130
                            _ = v167
                        }
                        if ((65535 >= uint(v1)) || (true && false)) {
                            var v168 uint64 = 1
                            _ = v168
                        } else {
                            var v169 uint16 = v1
                            _ = v169
                        }
                        _ = fi102
                        _ = v103
                    }
                } else {
                    if (!(false || (1 != 1))) {
                        var v170 bool = false
                        v100 = v11
                        _ = v170
                    } else {
                        var v171 uint16 = 0
                        v11 = v11
                        var v172 float64 = 0.1
                        _ = v171
                        _ = v172
                    }
                    var v173 uint8 = 42
                    _ = v173
                }
            }
            for fi174 := 0; fi174 < 19; fi174++ {
                if (2147483647 == 605) {
                    var v175 uint8 = 128
                    v1 = (v1 | (v1 << ((65535) & 63)))
                    v100 = 0
                    _ = v175
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch 587 {
                case 4294967295:
                    switch (("foo bar" + "") + "") {
                    case "":
                        if (65535 < 65535) {
                            switch ("foo bar" + "hello") {
                            case "":
                                var v176 float32 = (0.0 / 0.1)
                                _ = v176
                            case "foo bar":
                                fmt.Println(v1, fi101)
                            }
                        } else {
                            var v177 float32 = 0.1
                            _ = v177
                        }
                    case "a":
                        v1 -= v1
                    case "foo bar":
                        v1 += (v1 + v1)
                    default:
                        v178 := GenS27{F0: 498}
                        _ = v178
                    }
                case -554:
                    fmt.Println(v100)
                case 32768:
                    v179 := GenS10{F0: "hello", F1: 81, F2: 4294967295}
                    _ = v179
                default:
                    switch ("hello" + "\n\t") {
                    case "a":
                        v11 -= -128
                    }
                }
                _ = fi174
            }
            var v180 float32 = 0.1
            _ = fi101
            _ = v180
        }
        _ = v100
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
