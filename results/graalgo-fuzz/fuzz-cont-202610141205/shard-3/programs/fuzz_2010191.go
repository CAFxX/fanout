package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint16; F1 int32; F2 uint }
type GenS3 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int16; F1 uint16; F2 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32; F1 int8; F2 uint16 }
type GenS7 struct { F0 uint16; F1 bool }
type GenS8 struct { F0 int64; F1 string; F2 uint }
type GenS9 struct { F0 int64; F1 string; F2 uint8 }
type GenS10 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS11 struct { F0 int64; F1 int16; F2 string }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint16; F1 bool }
type GenS14 struct { F0 int64; F1 int8; F2 uint }
type GenS15 struct { F0 int }

func main() {
    var v1 uint32 = 4294967295
    var v2 string = (("a" + (("foo bar" + ("a" + ("a" + "\n\t"))) + ("hello" + ("a" + ((("" + "hello") + "") + ("" + "a")))))) + "")
    var v3 float32 = float32((float64(-2.25) - 1e10))
    v3 += 1e10
    v4 := []int{32767}
    fmt.Println(v1, v3)
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 uint64 = 0
        var v7 uint32 = (^(v1))
        switch v2 {
        case "a":
            switch (-(fi5)) {
            case 127:
                v8 := GenS1{F0: 1, F1: 1}
                _ = v8
            default:
                v4 = append(v4, 128)
            }
        case "":
            if (!true) {
                for fi9 := 0; fi9 < 6; fi9++ {
                    v7 -= 2
                    _ = fi9
                }
            } else {
                v7 -= v1
                v1 = 2
            }
        default:
            if (!((0 < 42) && true)) {
                v10 := GenS1{F0: 0, F1: 1}
                switch (-(fi5)) {
                case -433:
                    var v11 uint32 = (^(v1))
                    _ = v11
                default:
                    v4 = append(v4, (^(fi5)))
                }
                _ = v10
            } else {
                for fi12 := 0; fi12 < 11; fi12++ {
                    if false {
                        v7 += (v1 << ((1) & 63))
                        var v13 uint16 = 255
                        if (1 != 127) {
                            v2 = v2
                            fmt.Println(v13, v6)
                        } else {
                            var v14 uint32 = 4294967295
                            var v15 int = fi12
                            _ = v14
                            _ = v15
                        }
                        var v16 uint = 42
                        if true {
                            var v17 bool = (32767 == int16(v16))
                            switch 1 {
                            case 127:
                                v18 := GenS2{F0: v13, F1: 1, F2: (v16 + v16)}
                                _ = v18
                            case -131:
                                if v17 {
                                    switch v2 {
                                    case "a":
                                        if (v17 || (303 == 0)) {
                                            var v19 uint = 0
                                            if ((v13 >= (v13 << (((v19 << ((v16) & 63))) & 63))) && false) {
                                                v4 = append(v4, 1)
                                                if len(v4) > 0 {
                                                    v4[0] = fi12
                                                }
                                                if v17 {
                                                    if len(v4) > 0 {
                                                        v4[0] = (fi5 % ((fi5) | 1))
                                                    }
                                                    v20 := GenS3{F0: -1, F1: (v7 * v1), F2: v6}
                                                    var v21 float64 = (0.0 * -2.25)
                                                    _ = v20
                                                    _ = v21
                                                } else {
                                                    v13 = (v13 & (^(v13)))
                                                    fmt.Println(v7, v19, v16)
                                                }
                                                var v22 uint8 = 127
                                                var v23 bool = true
                                                var v24 float64 = 0.1
                                                _ = v22
                                                _ = v23
                                                _ = v24
                                            }
                                            _ = v19
                                        }
                                    default:
                                        v17 = v17
                                    }
                                }
                            case -152:
                                v1 -= (v7 << ((2) & 63))
                            }
                            var v25 float32 = v3
                            _ = v17
                            _ = v25
                        } else {
                            v26 := GenS4{F0: 600, F1: 42, F2: (!(int64(v13) < -640))}
                            var v27 bool = true
                            _ = v26
                            _ = v27
                        }
                        v4 = append(v4, (-(fi12)))
                        _ = v13
                        _ = v16
                    } else {
                        v3 = 0.1
                        var v28 uint64 = (^(v6))
                        fmt.Println(v7, fi12)
                        switch "hello" {
                        case "":
                            switch ("" + ("foo bar" + ("a" + "hello"))) {
                            case "\n\t":
                                fmt.Println(fi12, v3)
                            default:
                                v3 = 1.5
                            }
                        default:
                            fmt.Println(fi12, v28, fi5)
                        }
                        if false {
                            var v29 int16 = 1000
                            fmt.Println(fi12)
                            if false {
                                if true {
                                    if true {
                                        fmt.Println(v2, fi12)
                                        v3 -= v3
                                    }
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                var v30 int64 = -1
                                _ = v30
                            } else {
                                v31 := GenS1{F0: v29, F1: uint(v7)}
                                if ((fi5 & (fi12 << ((42) & 63))) <= fi5) {
                                    if true {
                                        switch 140 {
                                        case 987:
                                            v32 := GenS4{F0: (v29 << ((255) & 63)), F1: 65535, F2: true}
                                            if (false && ((42 > -65) || true)) {
                                                var v33 uint64 = v6
                                                _ = v33
                                            }
                                            v4 = append(v4, (^(fi12)))
                                            _ = v32
                                        case 42:
                                            if true {
                                                var v34 uint8 = 2
                                                v1 += (v1 + v7)
                                                _ = v34
                                            }
                                        case -1:
                                            v4 = append(v4, (fi12 + -1))
                                        }
                                    } else {
                                        var v35 int8 = -103
                                        v4 = append(v4, fi12)
                                        var v36 uint32 = (v7 % (((v7 * 65535)) | 1))
                                        v35 = v35
                                        if (((!true) || (-2.25 >= 0.0)) || true) {
                                            if true {
                                                if true {
                                                    v29 = int16(v1)
                                                }
                                                var v37 string = v2
                                                var v38 int = int(v1)
                                                v1 = 901
                                                v6 += v28
                                                _ = v37
                                                _ = v38
                                            }
                                            fmt.Println(fi12)
                                            switch 42 {
                                            case -884:
                                                if (true && ((-(v35)) >= (v35 ^ (v35 | v35)))) {
                                                    v39 := GenS5{F0: 42}
                                                    fmt.Println(v3, fi12)
                                                    v7 += (^(v1))
                                                    _ = v39
                                                }
                                            case 0:
                                                if false {
                                                    v3 += (1.5 - float32(0.1))
                                                }
                                            default:
                                                v3 = 3.14159
                                            }
                                        } else {
                                            var v40 bool = (-847 != -1)
                                            v7 = (^(v7))
                                            _ = len(v4)
                                            _ = v40
                                        }
                                        if true {
                                            var v41 string = v2
                                            v6 = 42
                                            v7 -= v7
                                            _ = v41
                                        }
                                        _ = len(v4)
                                        if (!(true || false)) {
                                            switch int(v28) {
                                            case 128:
                                                var v42 uint = 0
                                                _ = v42
                                            case 4294967295:
                                                switch ("" + v2) {
                                                case "foo bar":
                                                    var v43 uint = 0
                                                    _ = v43
                                                case "a":
                                                    fmt.Println(v1, v3)
                                                case "\n\t":
                                                    var v44 uint16 = 2
                                                    _ = v44
                                                default:
                                                    var v45 bool = (v2 != v2)
                                                    _ = v45
                                                }
                                            }
                                            var v46 uint16 = 2
                                            _ = v46
                                        }
                                        v35 += v35
                                        switch v2 {
                                        case "foo bar":
                                            var v47 int16 = (^(v29))
                                            v48 := GenS2{F0: uint16(v28), F1: 394, F2: uint(v6)}
                                            fmt.Println(v28, v47)
                                            _ = v47
                                            _ = v48
                                        case "":
                                            v49 := GenS4{F0: 1, F1: 124, F2: (fi12 < 4294967295)}
                                            _ = v49
                                        case "a":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        default:
                                            var v50 int32 = 89
                                            v51 := GenS6{F0: 733, F1: 100, F2: 255}
                                            _ = v50
                                            _ = v51
                                        }
                                        _ = v35
                                        _ = v36
                                    }
                                }
                                _ = v31
                            }
                            _ = v29
                        } else {
                            switch fi5 {
                            case 781:
                                switch (fi12 | -1) {
                                case 152:
                                    _ = len(v4)
                                case 0:
                                    v52 := GenS3{F0: 917, F1: v7, F2: 18446744073709551615}
                                    _ = v52
                                case 881:
                                    var v53 int = (fi5 >> ((209) & 63))
                                    _ = v53
                                }
                            }
                            if true {
                                fmt.Println(v6, fi12, v28)
                            }
                        }
                        v3 -= (((v3 - v3) / v3) * v3)
                        fmt.Println(v3, v6)
                        fmt.Println(fi5, v28)
                        _ = v28
                    }
                    if false {
                        v2 = ((("a" + "foo bar") + "") + "\n\t")
                        v1 = v7
                        if (false || (((v6 * (v6 - (v6 | v6))) == 42) && false)) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        var v54 int = fi12
                        fmt.Println(v6, v1, v1)
                        _ = v54
                    }
                    var v55 uint64 = uint64(v1)
                    _ = fi12
                    _ = v55
                }
            }
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(v2)
    v3 = -2.25
    v56 := GenS3{F0: -2147483648, F1: v1, F2: 18446744073709551615}
    for fi57 := 0; fi57 < 8; fi57++ {
        var v58 uint = 4294967295
        v59 := GenS7{F0: 0, F1: (!(105 > 1))}
        if true {
            var v60 bool = (!true)
            v61 := GenS8{F0: -937, F1: "\n\t", F2: 65535}
            _ = v60
            _ = v61
        }
        _ = fi57
        _ = v58
        _ = v59
    }
    if len(v4) > 0 {
        v4[0] = 176
    }
    v62 := GenS9{F0: 0, F1: v2, F2: 255}
    _ = len(v4)
    v3 = (v3 * (v3 * v3))
    v63 := GenS7{F0: 255, F1: (!true)}
    if len(v4) > 0 {
        v4[0] = -1
    }
    for fi64 := 0; fi64 < 3; fi64++ {
        if ((true || (-92 != 42)) || true) {
            var v65 string = v2
            var v66 int16 = 42
            var v67 uint16 = 2
            _ = v65
            _ = v66
            _ = v67
        } else {
            switch "foo bar" {
            case "\n\t":
                v2 = ""
            case "":
                if false {
                    switch v2 {
                    case "a":
                        v3 = (float32(3.14159) + (v3 - 3.14159))
                    case "hello":
                        for fi68 := 0; fi68 < 14; fi68++ {
                            if true {
                                var v69 string = v2
                                _ = v69
                            }
                            var v70 float64 = 0.0
                            switch fi68 {
                            case 42:
                                v2 = v2
                                v2 = "hello"
                            case -924:
                                if len(v4) > 0 {
                                    v4[0] = (fi68 - (fi68 << ((4294967295) & 63)))
                                }
                            }
                            if (229 >= 0) {
                                if true {
                                    var v71 float64 = v70
                                    var v72 float32 = float32(3.14159)
                                    if (false || true) {
                                        var v73 float32 = 0.0
                                        _ = len(v4)
                                        _ = v73
                                    } else {
                                        if ((v2 == v2) || (!(!(true || ((889 <= 42) || false))))) {
                                            if (140 != 555) {
                                                var v74 uint64 = 933
                                                if ((v1 >> ((18446744073709551615) & 63)) <= (v1 / ((2) | 1))) {
                                                    v3 = -2.25
                                                    if true {
                                                        v75 := GenS10{F0: 2, F1: 155, F2: (^(v74))}
                                                        var v76 uint8 = 128
                                                        v70 += 1.5
                                                        _ = v75
                                                        _ = v76
                                                    } else {
                                                        var v77 int64 = 354
                                                        _ = v77
                                                    }
                                                    v72 = v72
                                                } else {
                                                    var v78 uint8 = 2
                                                    _ = v78
                                                }
                                                fmt.Println(v70, v72, v71)
                                                _ = v74
                                            } else {
                                                var v79 string = v2
                                                var v80 int8 = 100
                                                if false {
                                                    if len(v4) > 0 {
                                                        v4[0] = (fi64 - 9223372036854775807)
                                                    }
                                                } else {
                                                    var v81 uint32 = 529
                                                    var v82 float64 = (3.14159 + (0.0 * v71))
                                                    if (v2 < ("" + v2)) {
                                                        switch v2 {
                                                        case "foo bar":
                                                            if true {
                                                                if (426 == 435) {
                                                                    v72 += float32(v70)
                                                                }
                                                                var v83 uint16 = 255
                                                                _ = v83
                                                            }
                                                        case "a":
                                                            v84 := GenS4{F0: 690, F1: 58, F2: false}
                                                            _ = v84
                                                        }
                                                        var v85 int16 = 1
                                                        _ = v85
                                                    } else {
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v86 int = fi68
                                                            _ = v86
                                                        case "\n\t":
                                                            var v87 string = (v79 + v79)
                                                            _ = v87
                                                        }
                                                        v88 := GenS6{F0: 0, F1: v80, F2: 551}
                                                        if (0 == 143) {
                                                            var v89 int8 = -3
                                                            _ = v89
                                                        }
                                                        _ = v88
                                                    }
                                                    _ = v81
                                                    _ = v82
                                                }
                                                _ = v79
                                                _ = v80
                                            }
                                            v70 = 0.0
                                        } else {
                                            var v90 uint64 = 18446744073709551615
                                            if false {
                                                if (80 > 42) {
                                                    fmt.Println(v1, v70)
                                                    _ = len(v4)
                                                }
                                            }
                                            _ = v90
                                        }
                                        v4 = append(v4, fi64)
                                        v1 += 4294967295
                                    }
                                    _ = v71
                                    _ = v72
                                } else {
                                    var v91 int8 = 127
                                    v2 = v2
                                    if false {
                                        if false {
                                            v2 = v2
                                            v3 = -2.25
                                        }
                                    }
                                    switch v2 {
                                    case "hello":
                                        var v92 int16 = -1
                                        _ = v92
                                    }
                                    if false {
                                        fmt.Println(v1, v2, v2)
                                        var v93 int8 = v91
                                        _ = v93
                                    }
                                    v70 += float64(((v3 * v3) + 1.5))
                                    _ = v91
                                }
                                var v94 uint8 = 42
                                v2 = (("\n\t" + v2) + ((v2 + ("hello" + v2)) + v2))
                                _ = v94
                            } else {
                                fmt.Println(fi68, fi64, v1)
                            }
                            _ = fi68
                            _ = v70
                        }
                    default:
                        v2 = (v2 + "foo bar")
                    }
                    var v95 bool = (!false)
                    _ = v95
                }
            case "a":
                if true {
                    switch fi64 {
                    case -997:
                        for fi96 := 0; fi96 < 6; fi96++ {
                            v3 = (((v3 + v3) * (1.5 + v3)) * (1e10 - ((v3 / 1.0) - (1.5 - 1.5))))
                            var v97 int64 = 0
                            if false {
                                v98 := GenS7{F0: 2, F1: false}
                                _ = v98
                            } else {
                                if false {
                                    var v99 float32 = 0.1
                                    if true {
                                        var v100 float64 = (-2.25 + 3.14159)
                                        _ = v100
                                    } else {
                                        _ = len(v4)
                                        if true {
                                            var v101 uint32 = (v1 ^ (^(v1)))
                                            var v102 float64 = (1.5 / float64(v99))
                                            v97 += v97
                                            v103 := GenS11{F0: (v97 * v97), F1: 1, F2: v2}
                                            v104 := GenS12{F0: (-(v97))}
                                            _ = v101
                                            _ = v102
                                            _ = v103
                                            _ = v104
                                        }
                                        v105 := GenS13{F0: 254, F1: false}
                                        if ((v97 / ((-41) | 1)) > (-(v97))) {
                                            var v106 int16 = int16(v97)
                                            v107 := GenS10{F0: 255, F1: 255, F2: 663}
                                            _ = v106
                                            _ = v107
                                        } else {
                                            var v108 uint64 = 1
                                            if (-1 >= 1) {
                                                var v109 uint64 = 473
                                                var v110 uint8 = 0
                                                _ = v109
                                                _ = v110
                                            }
                                            _ = v108
                                        }
                                        _ = v105
                                    }
                                    if true {
                                        switch (fi64 / (((fi96 >> ((4294967295) & 63))) | 1)) {
                                        case -2147483648:
                                            var v111 float32 = v99
                                            _ = v111
                                        case 1:
                                            v112 := GenS13{F0: 16, F1: false}
                                            _ = v112
                                        case 635:
                                            if true {
                                                fmt.Println(v1)
                                                v113 := GenS2{F0: 163, F1: 750, F2: 4294967295}
                                                if ((-71 == -16) && true) {
                                                    var v114 int = fi96
                                                    var v115 float64 = (float64((v99 - 3.14159)) - (((0.1 * (1.5 * (0.0 + (1.5 + 0.1)))) / 0.1) / 1.0))
                                                    v97 = v97
                                                    if len(v4) > 0 {
                                                        v4[0] = -950
                                                    }
                                                    fmt.Println(v115)
                                                    if len(v4) > 0 {
                                                        v4[0] = fi64
                                                    }
                                                    var v116 int16 = -1
                                                    if true {
                                                        fmt.Println(v99, v99)
                                                    }
                                                    v116 += (v116 % ((-1) | 1))
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v4 = append(v4, (fi96 - (v114 >> ((255) & 63))))
                                                    var v117 uint8 = 255
                                                    _ = v114
                                                    _ = v115
                                                    _ = v116
                                                    _ = v117
                                                }
                                                _ = v113
                                            } else {
                                                v118 := GenS14{F0: (v97 % ((42) | 1)), F1: -116, F2: 1}
                                                if (4294967295 < 511) {
                                                    var v119 uint16 = 2
                                                    if ((^(v1)) != (v1 & (v1 - v1))) {
                                                        var v120 float32 = -2.25
                                                        var v121 int8 = 99
                                                        var v122 int64 = (v97 << ((65535) & 63))
                                                        _ = v120
                                                        _ = v121
                                                        _ = v122
                                                    }
                                                    v123 := GenS15{F0: fi96}
                                                    _ = v119
                                                    _ = v123
                                                } else {
                                                    if len(v4) > 0 {
                                                        v4[0] = -2147483648
                                                    }
                                                    var v124 string = v2
                                                    v4 = append(v4, (fi64 | (fi96 + (fi64 - fi64))))
                                                    _ = v124
                                                }
                                                v1 += 1
                                                _ = v118
                                            }
                                        default:
                                            if false {
                                                if ((25 > 1) || false) {
                                                    var v125 string = ("" + v2)
                                                    _ = v125
                                                } else {
                                                    var v126 uint32 = v1
                                                    v97 = (v97 << ((255) & 63))
                                                    _ = v126
                                                }
                                            }
                                        }
                                    }
                                    fmt.Println(v99, v2)
                                    v1 = 925
                                    var v127 int64 = v97
                                    _ = v127
                                    _ = v99
                                }
                                if ((fi96 << ((18446744073709551615) & 63)) <= fi64) {
                                    switch v2 {
                                    case "foo bar":
                                        fmt.Println(v1, v1)
                                    default:
                                        fmt.Println(fi64, fi96, fi96)
                                    }
                                } else {
                                    if (!(4294967295 < 0)) {
                                        if true {
                                            switch fi64 {
                                            case 2147483647:
                                                fmt.Println(v1, fi96)
                                            default:
                                                switch v2 {
                                                case "foo bar":
                                                    if (!true) {
                                                        var v128 int64 = (^(v97))
                                                        var v129 uint8 = 86
                                                        _ = v128
                                                        _ = v129
                                                    }
                                                    var v130 int64 = v97
                                                    var v131 int32 = 984
                                                    _ = v130
                                                    _ = v131
                                                default:
                                                    var v132 int16 = 42
                                                    _ = v132
                                                }
                                            }
                                            v2 = v2
                                            var v133 string = "a"
                                            _ = v133
                                        } else {
                                            var v134 int32 = 777
                                            _ = v134
                                        }
                                    }
                                }
                                var v135 float32 = 1e10
                                _ = v135
                            }
                            v136 := GenS5{F0: 4294967295}
                            v3 -= ((v3 * (v3 / (v3 + v3))) - v3)
                            _ = fi96
                            _ = v136
                            _ = v97
                        }
                    default:
                        var v137 uint32 = v1
                        _ = v137
                    }
                    v138 := GenS9{F0: -1, F1: v2, F2: 2}
                    _ = v138
                }
            default:
                var v139 float64 = (1.5 * 0.1)
                _ = v139
            }
        }
        v3 = 1.5
        for fi140 := 0; fi140 < 18; fi140++ {
            var v141 bool = ((^(fi140)) != (fi140 ^ (fi64 << ((4294967295) & 63))))
            var v142 int64 = -757
            _ = fi140
            _ = v141
            _ = v142
        }
        _ = fi64
    }
    v2 = "a"
    var v143 float32 = v3
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v143:", v143)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v143
    _ = v2
    _ = v3
    _ = v4
    _ = v56
    _ = v62
    _ = v63
}
