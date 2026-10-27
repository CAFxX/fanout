package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 string; F1 bool; F2 uint8 }
type GenS5 struct { F0 int16; F1 int64; F2 int }
type GenS6 struct { F0 int8; F1 string; F2 string }
type GenS7 struct { F0 uint16; F1 int64 }
type GenS8 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS9 struct { F0 int; F1 int64; F2 bool }
type GenS10 struct { F0 bool; F1 int32 }
type GenS11 struct { F0 string; F1 uint; F2 bool }
type GenS12 struct { F0 string }
type GenS13 struct { F0 bool; F1 int8 }
type GenS14 struct { F0 uint64; F1 uint }
type GenS15 struct { F0 uint8; F1 uint16; F2 string }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[uint32]int)
    var v3 float32 = (((-2.25 / 3.14159) / (1e10 * -2.25)) * ((float32(1.5) / float32(1e10)) / ((0.1 - 0.1) + (3.14159 - 1e10))))
    var v4 int16 = 1000
    var v5 int64 = v1
    var v6 int8 = 42
    var v7 int16 = v4
    var v8 uint = 4294967295
    var v9 int64 = (v5 + (v1 / ((int64(v6)) | 1)))
    v2[65535] = -9223372036854775808
    var v10 bool = true
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := GenS1{F0: v10, F1: v6, F2: v8}
        switch (fi11 & 1) {
        case 0:
            var v13 int32 = 42
            _ = v13
        case -1:
            v14 := GenS2{F0: "hello"}
            _ = v14
        }
        if (((v8 * 411) >= v8) || false) {
            fmt.Println(fi11, v3)
        } else {
            var v15 int = fi11
            switch "a" {
            case "foo bar":
                for fi16 := 0; fi16 < 2; fi16++ {
                    v5 -= (v5 % ((-1) | 1))
                    _ = fi16
                }
            }
            if v10 {
                if (v3 != float32(0.0)) {
                    v17 := []int{595}
                    fmt.Println(v8)
                    _ = v17
                } else {
                    v1 = int64(v6)
                    var v18 uint64 = 1
                    switch ("a" + "hello") {
                    case "hello":
                        for fi19 := 0; fi19 < 5; fi19++ {
                            fmt.Println(v6)
                            var v20 uint32 = 42
                            switch fi11 {
                            case 9223372036854775807:
                                if false {
                                    v3 = (((float32(1.5) / 3.14159) / 3.14159) + (((v3 * v3) * v3) / (v3 * ((0.1 + (v3 + (1.5 * (float32(0.0) * v3)))) + (1e10 - -2.25)))))
                                }
                            case 42:
                                if v10 {
                                    var v21 bool = v10
                                    fmt.Println(v20, fi19, v8)
                                    _ = v21
                                }
                            }
                            _ = fi19
                            _ = v20
                        }
                    case "foo bar":
                        var v22 int32 = -357
                        _ = v22
                    case "a":
                        fmt.Println(v18, v10)
                    }
                    _ = v18
                }
            } else {
                if v10 {
                    if v10 {
                        for fi23 := 0; fi23 < 4; fi23++ {
                            if (v10 && (!(true && v10))) {
                                if true {
                                    var v24 float64 = 0.0
                                    v25 := GenS1{F0: (v10 || true), F1: (v6 >> ((255) & 63)), F2: v8}
                                    if v10 {
                                        v10 = v10
                                        var v26 bool = v10
                                        _ = v26
                                    } else {
                                        switch (^(fi23)) {
                                        case 4294967295:
                                            if v10 {
                                                fmt.Println(v7)
                                                v27 := GenS3{F0: (v1 << ((2) & 63)), F1: v9}
                                                v15 = int(v6)
                                                switch "" {
                                                case "hello":
                                                    var v28 int32 = -2147483648
                                                    _ = v28
                                                case "a":
                                                    if false {
                                                        var v29 uint16 = 473
                                                        v24 = -2.25
                                                        v2[0] = fi11
                                                        _ = v29
                                                    } else {
                                                        v30 := []int{4294967295, (fi23 | fi23)}
                                                        _ = v30
                                                    }
                                                default:
                                                    var v31 float64 = float64(v3)
                                                    _ = v31
                                                }
                                                if (!(false && ((^(v5)) > (^(v1))))) {
                                                    var v32 bool = v10
                                                    fmt.Println(v4)
                                                    var v33 uint = 18446744073709551615
                                                    _ = v32
                                                    _ = v33
                                                } else {
                                                    if false {
                                                        var v34 int64 = (v5 >> ((v8) & 63))
                                                        _ = v34
                                                    }
                                                    var v35 uint8 = 0
                                                    v36 := GenS1{F0: (int(v6) != -767), F1: (v6 - 1), F2: v8}
                                                    _ = v35
                                                    _ = v36
                                                }
                                                v37 := []int{(v15 - fi23), -549, 0, (fi23 | fi11)}
                                                fmt.Println(fi11)
                                                var v38 uint64 = 2
                                                var v39 uint32 = 4294967295
                                                _ = v27
                                                _ = v37
                                                _ = v38
                                                _ = v39
                                            } else {
                                                v5 -= v1
                                                fmt.Println(v1, v3)
                                                v8 = v8
                                            }
                                        default:
                                            var v40 uint16 = 255
                                            _ = v40
                                        }
                                        switch (v15 << ((65535) & 63)) {
                                        case -2147483648:
                                            v6 = 127
                                        case -270:
                                            var v41 float64 = v24
                                            _ = v41
                                        case 32768:
                                            v7 -= v7
                                        }
                                        if false {
                                            var v42 bool = (float32((v24 - 0.1)) > v3)
                                            v43 := v2[2147483648]
                                            _ = v42
                                            _ = v43
                                        }
                                    }
                                    _ = v24
                                    _ = v25
                                }
                                delete(v2, 0)
                            } else {
                                v44 := GenS4{F0: ("foo bar" + "hello"), F1: (4294967295 <= 42), F2: 95}
                                var v45 float32 = v3
                                var v46 int64 = v9
                                _ = v44
                                _ = v45
                                _ = v46
                            }
                            _ = len(v2)
                            _ = fi23
                        }
                    } else {
                        v6 = (v6 % (((v6 & v6)) | 1))
                        for fi47 := 0; fi47 < 12; fi47++ {
                            if (fi47 >= v15) {
                                var v48 uint32 = 4294967295
                                _ = v48
                            } else {
                                v5 = (v5 << ((42) & 63))
                                if false {
                                    if (1e10 == 3.14159) {
                                        if (v10 && ((^(v6)) <= (^(v6)))) {
                                            if (!(1 < 0)) {
                                                if false {
                                                    var v49 int = -9223372036854775808
                                                    _ = v49
                                                }
                                            } else {
                                                var v50 int32 = 2147483647
                                                v8 = v8
                                                v1 = v9
                                                _ = v50
                                            }
                                        }
                                        fmt.Println(fi11)
                                    } else {
                                        var v51 int16 = 42
                                        _ = v51
                                    }
                                    v3 -= ((-2.25 + (v3 - 1.5)) * float32(float64(v3)))
                                } else {
                                    var v52 uint8 = 223
                                    _ = v52
                                }
                                if false {
                                    var v53 uint16 = 1
                                    v10 = ((v53 >= v53) && (42 > 0))
                                    if false {
                                        if v10 {
                                            var v54 uint8 = 175
                                            if false {
                                                v55 := GenS4{F0: ("foo bar" + ""), F1: true, F2: v54}
                                                v56 := GenS5{F0: v4, F1: (v1 + v1), F2: -284}
                                                fmt.Println(v54, v54, v15)
                                                _ = v55
                                                _ = v56
                                            } else {
                                                switch (v15 - (^(fi11))) {
                                                case 42:
                                                    switch (fi47 + (fi47 >> (((v8 << (((^(v8))) & 63))) & 63))) {
                                                    case 9223372036854775807:
                                                        var v57 uint16 = v53
                                                        _ = v57
                                                    case -2147483648:
                                                        if v10 {
                                                            v58 := GenS4{F0: "", F1: (613 == v53), F2: 2}
                                                            v59 := v2[42]
                                                            v60 := []int{v59, v59}
                                                            _ = v58
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                    default:
                                                        if ((0.0 - 1.5) >= ((-2.25 - (-2.25 * (((-2.25 - 3.14159) - 3.14159) - 1.5))) * 3.14159)) {
                                                            v61 := []int{fi11, (fi47 << ((v8) & 63))}
                                                            switch (fi47 ^ (^(fi11))) {
                                                            case 42:
                                                                var v62 uint = v8
                                                                _ = v62
                                                            case 4294967295:
                                                                v53 -= (^(v53))
                                                            default:
                                                                v63 := GenS6{F0: 1, F1: "\n\t", F2: "\n\t"}
                                                                _ = v63
                                                            }
                                                            _ = v61
                                                        } else {
                                                            v64 := GenS7{F0: (v53 & v53), F1: v1}
                                                            delete(v2, uint32(v54))
                                                            v6 = (v6 ^ v6)
                                                            _ = v64
                                                        }
                                                    }
                                                default:
                                                    if v10 {
                                                        fmt.Println(v8, v4)
                                                        v8 = v8
                                                        if v10 {
                                                            v65 := GenS8{F0: 1, F1: v53, F2: 65535}
                                                            fmt.Println(v53, v9)
                                                            if (!v10) {
                                                                var v66 int64 = (v9 / ((v5) | 1))
                                                                _ = v66
                                                            } else {
                                                                v2[531] = (fi11 & -541)
                                                            }
                                                            _ = v65
                                                        }
                                                    } else {
                                                        fmt.Println(v4, v8)
                                                    }
                                                }
                                                if (!(!(((1.5 - -2.25) < float64(0.1)) || v10))) {
                                                    switch v15 {
                                                    case -257:
                                                        v53 = v53
                                                    case 0:
                                                        if (!((v6 + v6) >= (v6 & (v6 / (((v6 / ((v6) | 1))) | 1))))) {
                                                            v3 = v3
                                                            v5 -= v1
                                                        } else {
                                                            var v67 int16 = (v4 | (v4 ^ -1))
                                                            v68 := GenS9{F0: 135, F1: v9, F2: false}
                                                            _ = v67
                                                            _ = v68
                                                        }
                                                    case 32768:
                                                        v69 := []int{}
                                                        _ = v69
                                                    }
                                                    fmt.Println(v54)
                                                } else {
                                                    v70 := GenS3{F0: v5, F1: v9}
                                                    switch (fi11 & v15) {
                                                    case 128:
                                                        var v71 bool = v10
                                                        _ = v71
                                                    case 127:
                                                        fmt.Println(v10, v10, v10)
                                                    default:
                                                        if (("a" + "") == ("\n\t" + ("foo bar" + "hello"))) {
                                                            fmt.Println(v7, v4)
                                                        }
                                                    }
                                                    _ = v70
                                                }
                                                v5 = v1
                                            }
                                            v72 := []int{v15, fi11, 127, int(v9)}
                                            _ = v54
                                            _ = v72
                                        } else {
                                            fmt.Println(v6, v8)
                                            switch "\n\t" {
                                            case "a":
                                                var v73 uint8 = 42
                                                _ = v73
                                            case "foo bar":
                                                fmt.Println(v10, v7, v4)
                                                var v74 int64 = (v5 * v9)
                                                v75 := GenS2{F0: (("hello" + "") + "\n\t")}
                                                _ = v74
                                                _ = v75
                                            }
                                            v10 = v10
                                        }
                                    }
                                    _ = v53
                                } else {
                                    var v76 float64 = 0.1
                                    switch (fi47 - (fi47 - (v15 & v15))) {
                                    case -9223372036854775808:
                                        v76 = (v76 * v76)
                                    default:
                                        v77 := GenS10{F0: (1 > 32767), F1: 2147483647}
                                        _ = v77
                                    }
                                    v78 := GenS7{F0: 629, F1: (v5 << ((v8) & 63))}
                                    _ = v76
                                    _ = v78
                                }
                            }
                            v7 += 249
                            _ = fi47
                        }
                    }
                } else {
                    var v79 int64 = (v5 + -481)
                    if v10 {
                        v80 := GenS4{F0: "a", F1: v10, F2: 2}
                        v81 := []int{(fi11 ^ 42), fi11}
                        _ = v80
                        _ = v81
                    } else {
                        for fi82 := 0; fi82 < 16; fi82++ {
                            var v83 float32 = v3
                            var v84 uint32 = 2147483647
                            _ = fi82
                            _ = v83
                            _ = v84
                        }
                        switch ("foo bar" + ("\n\t" + "\n\t")) {
                        case "\n\t":
                            switch "hello" {
                            case "a":
                                if v10 {
                                    v85 := GenS3{F0: v79, F1: v1}
                                    fmt.Println(v7, v79)
                                    _ = v85
                                } else {
                                    if ((int16(fi11) >= v7) && v10) {
                                        switch ("hello" + ("foo bar" + "hello")) {
                                        case "":
                                            v86 := []int{}
                                            _ = v86
                                        case "a":
                                            v87 := v2[4294967295]
                                            _ = v87
                                        }
                                    } else {
                                        var v88 int32 = -1
                                        switch 0 {
                                        case 42:
                                            for fi89 := 0; fi89 < 12; fi89++ {
                                                var v90 uint32 = 2147483648
                                                v91 := GenS7{F0: 2, F1: 42}
                                                fmt.Println(fi89)
                                                _ = fi89
                                                _ = v90
                                                _ = v91
                                            }
                                        case 1:
                                            fmt.Println(v88)
                                        }
                                        v88 -= -2147483648
                                        _ = v88
                                    }
                                    var v92 bool = (v10 && v10)
                                    _ = v92
                                }
                            case "foo bar":
                                var v93 uint64 = 417
                                _ = v93
                            default:
                                v94 := []int{fi11}
                                _ = v94
                            }
                        case "":
                            var v95 float32 = 1e10
                            _ = v95
                        case "hello":
                            if (v10 && (578 <= 1)) {
                                fmt.Println(v10, v9, v6)
                                if false {
                                    v96 := GenS4{F0: "a", F1: true, F2: 255}
                                    switch fi11 {
                                    case 227:
                                        var v97 int = 32767
                                        _ = v97
                                    case -786:
                                        if (v8 >= 18446744073709551615) {
                                            if v10 {
                                                var v98 bool = (float64(v3) < 1.5)
                                                _ = v98
                                            } else {
                                                v1 += (v9 << ((v8) & 63))
                                                var v99 int = fi11
                                                fmt.Println(v1)
                                                _ = v99
                                            }
                                            if v10 {
                                                if true {
                                                    if v10 {
                                                        var v100 int64 = 306
                                                        if v10 {
                                                            switch "\n\t" {
                                                            case "hello":
                                                                v6 -= (-(v6))
                                                            default:
                                                                delete(v2, 4294967295)
                                                            }
                                                        }
                                                        _ = v100
                                                    } else {
                                                        fmt.Println(fi11)
                                                    }
                                                }
                                                var v101 bool = ((v5 < v5) || (v10 || ((2 != 434) && (v10 && v10))))
                                                var v102 int32 = -2147483648
                                                _ = v101
                                                _ = v102
                                            }
                                        } else {
                                            var v103 uint32 = 0
                                            _ = v103
                                        }
                                    case 32768:
                                        v104 := GenS7{F0: 2, F1: (^(v9))}
                                        _ = v104
                                    }
                                    switch "foo bar" {
                                    case "hello":
                                        for fi105 := 0; fi105 < 2; fi105++ {
                                            switch int(v6) {
                                            case 128:
                                                v10 = ("" <= "foo bar")
                                            }
                                            v106 := []int{}
                                            v4 = int16(v6)
                                            _ = fi105
                                            _ = v106
                                        }
                                    default:
                                        if (!((!(0 < 1)) && (2 < 128))) {
                                            var v107 uint8 = 0
                                            v108 := []int{(fi11 % ((int(v107)) | 1))}
                                            _ = v107
                                            _ = v108
                                        } else {
                                            if v10 {
                                                if (!v10) {
                                                    var v109 int32 = 538
                                                    v4 = v7
                                                    v110 := []int{fi11, fi11}
                                                    _ = v109
                                                    _ = v110
                                                } else {
                                                    for fi111 := 0; fi111 < 7; fi111++ {
                                                        _ = len(v2)
                                                        var v112 float64 = (1e10 + -2.25)
                                                        v79 = v1
                                                        _ = fi111
                                                        _ = v112
                                                    }
                                                    v10 = false
                                                }
                                                for fi113 := 0; fi113 < 16; fi113++ {
                                                    switch "foo bar" {
                                                    case "hello":
                                                        if ((false || v10) || v10) {
                                                            v15 -= fi11
                                                            if v10 {
                                                                var v114 int64 = v1
                                                                _ = v114
                                                            }
                                                            var v115 uint16 = 377
                                                            v116 := []int{v15, -438, (v15 / ((1) | 1)), fi113}
                                                            v5 = (v1 ^ (^(v79)))
                                                            _ = v115
                                                            _ = v116
                                                        }
                                                        v117 := []int{}
                                                        var v118 int = fi11
                                                        _ = v117
                                                        _ = v118
                                                    case "":
                                                        fmt.Println(v5, fi113, v8)
                                                    }
                                                    if (v6 == 100) {
                                                        var v119 float32 = ((v3 + v3) * ((0.1 * v3) - v3))
                                                        switch fi11 {
                                                        case 128:
                                                            if (236 <= 73) {
                                                                var v120 int8 = v6
                                                                _ = v120
                                                            }
                                                        case 42:
                                                            v9 += -9223372036854775808
                                                        }
                                                        if (!false) {
                                                            var v121 int = fi113
                                                            _ = v121
                                                        } else {
                                                            var v122 uint = v8
                                                            _ = v122
                                                        }
                                                        _ = v119
                                                    }
                                                    if true {
                                                        switch "a" {
                                                        case "":
                                                            v123 := GenS11{F0: "\n\t", F1: v8, F2: (v10 || (v8 >= 4294967295))}
                                                            _ = v123
                                                        }
                                                        if (v10 || ((^(v7)) > v7)) {
                                                            var v124 uint16 = 0
                                                            var v125 uint = 2
                                                            switch (^(fi11)) {
                                                            case 127:
                                                                var v126 uint8 = 0
                                                                _ = v126
                                                            case 0:
                                                                v125 -= (v125 / ((v125) | 1))
                                                            }
                                                            _ = v124
                                                            _ = v125
                                                        }
                                                    } else {
                                                        fmt.Println(v10)
                                                    }
                                                    fmt.Println(v5, v4, v8)
                                                    fmt.Println(v10)
                                                    _ = fi113
                                                }
                                                if v10 {
                                                    v127 := GenS11{F0: ("hello" + ("" + ("a" + ("foo bar" + ("foo bar" + "a"))))), F1: 954, F2: v10}
                                                    if (v10 && v10) {
                                                        var v128 uint16 = 735
                                                        fmt.Println(v1, v4)
                                                        _ = v128
                                                    } else {
                                                        v129 := GenS11{F0: "foo bar", F1: (v8 - v8), F2: (255 >= 18446744073709551615)}
                                                        v130 := GenS4{F0: ("a" + "\n\t"), F1: (127 > 44), F2: 2}
                                                        _ = v129
                                                        _ = v130
                                                    }
                                                    v131 := v2[22]
                                                    _ = v127
                                                    _ = v131
                                                } else {
                                                    v6 = 1
                                                }
                                            } else {
                                                switch "a" {
                                                case "":
                                                    var v132 float32 = (v3 + (1.5 + float32(0.0)))
                                                    _ = v132
                                                }
                                            }
                                            var v133 uint = v8
                                            if false {
                                                var v134 uint32 = 119
                                                for fi135 := 0; fi135 < 1; fi135++ {
                                                    var v136 float64 = (float64((1e10 + 3.14159)) / (0.1 + 0.1))
                                                    var v137 uint16 = 0
                                                    v138 := GenS5{F0: v4, F1: -9223372036854775808, F2: v15}
                                                    _ = fi135
                                                    _ = v136
                                                    _ = v137
                                                    _ = v138
                                                }
                                                _ = v134
                                            }
                                            _ = v133
                                        }
                                    }
                                    _ = v96
                                }
                                v5 -= v1
                            } else {
                                v139 := GenS12{F0: ("foo bar" + "\n\t")}
                                if v10 {
                                    v4 += v4
                                }
                                fmt.Println(v5, v4, v1)
                                _ = v139
                            }
                        default:
                            if v10 {
                                var v140 float64 = (((-2.25 * (0.1 + ((0.0 + ((-2.25 * float64((((float32(float64(v3)) * (0.0 * 3.14159)) / (v3 * v3)) - (-2.25 * v3)))) + (0.1 + 0.1))) / (0.1 + (0.0 / -2.25))))) + 3.14159) * (0.1 * -2.25))
                                _ = v140
                            }
                        }
                    }
                    _ = v79
                }
                if (!(42 > 64)) {
                    v141 := GenS13{F0: false, F1: 0}
                    v142 := GenS14{F0: 0, F1: v8}
                    for fi143 := 0; fi143 < 7; fi143++ {
                        var v144 uint32 = 2147483647
                        v145 := GenS15{F0: 182, F1: 42, F2: "foo bar"}
                        v146 := []int{}
                        _ = fi143
                        _ = v144
                        _ = v145
                        _ = v146
                    }
                    _ = v141
                    _ = v142
                } else {
                    v15 += (fi11 - -287)
                }
                for fi147 := 0; fi147 < 4; fi147++ {
                    var v148 uint16 = 1
                    var v149 uint8 = 128
                    v5 = v1
                    _ = fi147
                    _ = v148
                    _ = v149
                }
            }
            _ = v15
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v10, v5)
    if false {
        var v150 int16 = v7
        _ = v150
    } else {
        var v151 int16 = (v4 / ((v7) | 1))
        for fi152 := 0; fi152 < 4; fi152++ {
            switch (fi152 | (^(fi152))) {
            case -9223372036854775808:
                for fi153 := 0; fi153 < 18; fi153++ {
                    if (128 < 2) {
                        var v154 float64 = 1e10
                        v3 += ((v3 - float32(v154)) + v3)
                        v9 = v9
                        var v155 string = "\n\t"
                        _ = v154
                        _ = v155
                    } else {
                        if v10 {
                            v8 += v8
                            if false {
                                var v156 int16 = 1
                                _ = v156
                            }
                            var v157 uint8 = 55
                            var v158 uint32 = 928
                            _ = v157
                            _ = v158
                        }
                    }
                    if true {
                        if v10 {
                            v8 = (v8 % ((v8) | 1))
                            var v159 uint8 = 0
                            _ = v159
                        }
                        v10 = v10
                        v160 := []int{(fi153 * (fi152 / ((825) | 1))), (fi152 << (((^(v8))) & 63)), 9223372036854775807, 127}
                        _ = v160
                    }
                    v7 = v151
                    _ = fi153
                }
            case 4294967295:
                if (!v10) {
                    v1 = (^(v1))
                    v2[0] = 32767
                    var v161 int = fi152
                    _ = v161
                }
            case -963:
                v162 := GenS15{F0: 0, F1: 255, F2: "hello"}
                _ = v162
            }
            v6 = v6
            v163 := GenS3{F0: (v9 & -9223372036854775808), F1: v1}
            _ = fi152
            _ = v163
        }
        _ = v151
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
