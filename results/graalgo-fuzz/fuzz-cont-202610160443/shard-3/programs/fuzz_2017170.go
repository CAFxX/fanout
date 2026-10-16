package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int8; F1 int32 }
type GenS8 struct { F0 uint64; F1 int16 }
type GenS9 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint; F1 uint; F2 uint16 }
type GenS12 struct { F0 int16; F1 int8 }
type GenS13 struct { F0 uint64; F1 uint32 }
type GenS14 struct { F0 int64; F1 int }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -1
    v3 := []uint32{2147483647, 576, 42}
    switch 42 {
    case -626:
        switch -160 {
        case 744:
            for fi4 := 0; fi4 < 1; fi4++ {
                v2 = (v2 & v2)
                v5 := GenS1{F0: v2, F1: (^(v2)), F2: 0}
                v2 -= v2
                _ = fi4
                _ = v5
            }
        case 32767:
            v6 := GenS1{F0: (v2 >> ((1) & 63)), F1: (v2 % ((v2) | 1)), F2: 838}
            _ = v6
        }
    case 4294967295:
        if true {
            v7 := GenS1{F0: v2, F1: (v2 << ((919) & 63)), F2: 110}
            _ = v7
        }
    }
    if ((818 < 932) || false) {
        v2 -= v2
    }
    v8 := GenS1{F0: 0, F1: (v2 + v2), F2: 687}
    fmt.Println(v2)
    var v9 string = ("" + "")
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := v1[42]
        if (!true) {
            if false {
                delete(v1, (v11 ^ v11))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            for fi12 := 0; fi12 < 7; fi12++ {
                v1[840] = (^(v11))
                var v13 float32 = 3.14159
                var v14 int16 = 32767
                var v15 uint = 65535
                var v16 string = v9
                v17 := GenS1{F0: 855, F1: 2147483647, F2: 358}
                v16 = v16
                fmt.Println(v2, fi12, v13)
                v15 += 922
                switch v16 {
                case "hello":
                    v18 := GenS1{F0: -416, F1: v2, F2: 514}
                    _ = v18
                }
                v9 = "\n\t"
                var v19 float32 = float32(3.14159)
                var v20 string = "a"
                v15 = v15
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
                _ = v16
                _ = v17
                _ = v19
                _ = v20
            }
            if ((!(-9223372036854775808 <= 143)) && ((0.1 <= (1.5 * 3.14159)) || true)) {
                var v21 uint8 = 255
                _ = v21
            } else {
                v2 = (^(v2))
            }
        }
        _ = fi10
        _ = v11
    }
    if true {
        var v22 int = 128
        if false {
            for fi23 := 0; fi23 < 1; fi23++ {
                for fi24 := 0; fi24 < 11; fi24++ {
                    v25 := GenS2{F0: true}
                    _ = fi24
                    _ = v25
                }
                _ = fi23
            }
            var v26 uint = 18446744073709551615
            v3 = append(v3, 1)
            _ = v26
        }
        _ = v22
    }
    for fi27 := 0; fi27 < 10; fi27++ {
        var v28 float32 = (-2.25 - 0.1)
        _ = len(v3)
        if false {
            for fi29 := 0; fi29 < 13; fi29++ {
                var v30 int8 = 127
                _ = fi29
                _ = v30
            }
            switch int(v2) {
            case 9223372036854775807:
                _ = len(v1)
            case 32768:
                for fi31 := 0; fi31 < 14; fi31++ {
                    v32 := GenS1{F0: (-(v2)), F1: (v2 << ((110) & 63)), F2: 32767}
                    _ = fi31
                    _ = v32
                }
            case 4294967295:
                if (v9 > "hello") {
                    switch (fi27 % (((fi27 * fi27)) | 1)) {
                    case 915:
                        fmt.Println(v9)
                    case -735:
                        delete(v1, 969)
                    }
                    v3 = append(v3, 1)
                    switch (fi27 - (^(fi27))) {
                    case 2147483647:
                        var v33 int16 = 32767
                        _ = v33
                    case -9223372036854775808:
                        switch 296 {
                        case 2147483647:
                            if (!true) {
                                if true {
                                    v2 += (v2 | (v2 >> ((755) & 63)))
                                    v2 = (v2 * (v2 ^ v2))
                                    fmt.Println(fi27)
                                }
                            } else {
                                if true {
                                    v34 := GenS3{F0: 4294967295}
                                    v9 = v9
                                    _ = v34
                                }
                            }
                        case -945:
                            _ = len(v3)
                        }
                    case 0:
                        v9 = "foo bar"
                    }
                } else {
                    v9 = "foo bar"
                    if true {
                        v2 = (v2 / ((2147483647) | 1))
                        if (false && true) {
                            for fi35 := 0; fi35 < 5; fi35++ {
                                _ = len(v1)
                                var v36 uint = 2
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v37 int16 = -1
                                v38 := GenS2{F0: true}
                                switch (v9 + "a") {
                                case "\n\t":
                                    if true {
                                        var v39 float32 = (v28 + (1.5 / 3.14159))
                                        if (!true) {
                                            var v40 uint = (v36 * 691)
                                            var v41 int = fi35
                                            switch v9 {
                                            case "\n\t":
                                                if (!true) {
                                                    v42 := GenS4{F0: v9}
                                                    var v43 float32 = v39
                                                    _ = v42
                                                    _ = v43
                                                }
                                            default:
                                                var v44 uint = (v40 | 65535)
                                                _ = v44
                                            }
                                            v45 := GenS2{F0: (409 <= 120)}
                                            v28 -= (v39 - v28)
                                            _ = v40
                                            _ = v41
                                            _ = v45
                                        } else {
                                            v2 += 2147483647
                                            switch 400 {
                                            case -2147483648:
                                                v3 = append(v3, 944)
                                            }
                                        }
                                        _ = v39
                                    }
                                }
                                _ = fi35
                                _ = v36
                                _ = v37
                                _ = v38
                            }
                            var v46 int8 = 1
                            if false {
                                if (true || true) {
                                    if false {
                                        v28 += v28
                                        if false {
                                            fmt.Println(v28)
                                            switch fi27 {
                                            case -351:
                                                var v47 uint = 53
                                                _ = v47
                                            case 979:
                                                v28 -= float32((0.0 * ((((0.1 / 0.1) + (0.0 - 0.0)) * 1.5) - 0.0)))
                                            default:
                                                var v48 int32 = v2
                                                _ = v48
                                            }
                                        }
                                    }
                                    fmt.Println(v2)
                                    for fi49 := 0; fi49 < 4; fi49++ {
                                        v3 = append(v3, 0)
                                        _ = fi49
                                    }
                                } else {
                                    v50 := GenS3{F0: 841}
                                    var v51 bool = true
                                    _ = v50
                                    _ = v51
                                }
                                switch (v9 + v9) {
                                case "a":
                                    if true {
                                        var v52 float32 = v28
                                        for fi53 := 0; fi53 < 2; fi53++ {
                                            var v54 uint64 = 18446744073709551615
                                            var v55 uint8 = 127
                                            if true {
                                                var v56 uint32 = 0
                                                if true {
                                                    if false {
                                                        v57 := GenS2{F0: (!false)}
                                                        var v58 uint32 = 4294967295
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                    var v59 bool = (42 > 42)
                                                    var v60 uint64 = (v54 - 449)
                                                    _ = v59
                                                    _ = v60
                                                } else {
                                                    v52 = ((v52 + (float32(0.1) / (float32(-2.25) + 1.5))) / ((v28 + (v52 - (0.1 - (float32((-2.25 / 1.5)) + 1.5)))) / v28))
                                                }
                                                if (!false) {
                                                    var v61 int64 = 928
                                                    var v62 uint64 = v54
                                                    _ = v61
                                                    _ = v62
                                                }
                                                _ = v56
                                            }
                                            _ = fi53
                                            _ = v54
                                            _ = v55
                                        }
                                        _ = v52
                                    } else {
                                        v63 := GenS4{F0: "foo bar"}
                                        v64 := GenS5{F0: v9}
                                        fmt.Println(v28, v2)
                                        _ = v63
                                        _ = v64
                                    }
                                case "\n\t":
                                    if true {
                                        if false {
                                            for fi65 := 0; fi65 < 13; fi65++ {
                                                v66 := GenS5{F0: "foo bar"}
                                                fmt.Println(v28, fi65)
                                                v67 := GenS6{F0: (v2 - (v2 - 1))}
                                                _ = fi65
                                                _ = v66
                                                _ = v67
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v68 := v1[0]
                                            _ = v68
                                        }
                                        if (!(!false)) {
                                            v69 := GenS2{F0: true}
                                            v70 := v1[2147483647]
                                            _ = v69
                                            _ = v70
                                        } else {
                                            v71 := GenS4{F0: ("foo bar" + ((v9 + "\n\t") + "\n\t"))}
                                            _ = v71
                                        }
                                        var v72 string = v9
                                        _ = v72
                                    } else {
                                        var v73 uint16 = 879
                                        var v74 string = "\n\t"
                                        if (255 <= 255) {
                                            for fi75 := 0; fi75 < 19; fi75++ {
                                                switch v74 {
                                                case "a":
                                                    if (!(((float64((0.0 + v28)) - 1.5) * 0.0) <= 1.5)) {
                                                        v76 := GenS1{F0: v2, F1: v2, F2: (v73 * v73)}
                                                        _ = v76
                                                    } else {
                                                        v77 := GenS4{F0: v9}
                                                        _ = v77
                                                    }
                                                case "\n\t":
                                                    if false {
                                                        if (!(v9 <= "")) {
                                                            v78 := GenS7{F0: (v46 / ((v46) | 1)), F1: -1}
                                                            fmt.Println(v73, fi27, v73)
                                                            if true {
                                                                _ = len(v1)
                                                                v9 = v74
                                                                var v79 int64 = 781
                                                                var v80 int = fi27
                                                                if ((false && ((-(v79)) < (v79 | int64(v80)))) || true) {
                                                                    var v81 uint16 = (^(v73))
                                                                    v82 := GenS5{F0: "\n\t"}
                                                                    _ = v81
                                                                    _ = v82
                                                                }
                                                                _ = v79
                                                                _ = v80
                                                            }
                                                            _ = v78
                                                        }
                                                        var v83 float32 = float32(3.14159)
                                                        var v84 int16 = -32768
                                                        v85 := v1[0]
                                                        _ = v83
                                                        _ = v84
                                                        _ = v85
                                                    } else {
                                                        fmt.Println(v9, v46)
                                                        v28 = (v28 * v28)
                                                        var v86 uint64 = 0
                                                        _ = v86
                                                    }
                                                case "":
                                                    v2 -= v2
                                                    if len(v3) > 0 {
                                                        v3[0] = 2
                                                    }
                                                    v46 -= v46
                                                }
                                                delete(v1, 2)
                                                _ = fi75
                                            }
                                            v74 = v74
                                        } else {
                                            fmt.Println(v9)
                                            if true {
                                                var v87 bool = (true && false)
                                                _ = v87
                                            }
                                        }
                                        _ = v73
                                        _ = v74
                                    }
                                }
                            } else {
                                v3 = append(v3, 425)
                                if false {
                                    if false {
                                        var v88 uint64 = 1
                                        if true {
                                            for fi89 := 0; fi89 < 10; fi89++ {
                                                if false {
                                                    v28 += float32(0.1)
                                                    delete(v1, 0)
                                                }
                                                var v90 uint8 = 128
                                                _ = fi89
                                                _ = v90
                                            }
                                            v91 := GenS6{F0: (^(v2))}
                                            v92 := GenS1{F0: v2, F1: (v2 % (((v2 * 0)) | 1)), F2: 32767}
                                            _ = v91
                                            _ = v92
                                        } else {
                                            for fi93 := 0; fi93 < 20; fi93++ {
                                                v94 := GenS2{F0: ((-(fi93)) > fi93)}
                                                var v95 int64 = 1
                                                switch (fi27 ^ 4294967295) {
                                                case 0:
                                                    v28 -= (v28 + float32((float64(1e10) - -2.25)))
                                                case 1:
                                                    var v96 uint = 1
                                                    _ = v96
                                                case -1:
                                                    var v97 uint64 = v88
                                                    _ = v97
                                                }
                                                _ = fi93
                                                _ = v94
                                                _ = v95
                                            }
                                            var v98 int32 = 953
                                            for fi99 := 0; fi99 < 5; fi99++ {
                                                var v100 float64 = (-2.25 * float64(v28))
                                                if true {
                                                    v9 = v9
                                                    switch -9223372036854775808 {
                                                    case 32767:
                                                        v101 := GenS6{F0: 42}
                                                        _ = v101
                                                    case 128:
                                                        fmt.Println(v100, v2, v28)
                                                    case -210:
                                                        v100 += -2.25
                                                    default:
                                                        switch fi99 {
                                                        case 128:
                                                            v102 := GenS8{F0: v88, F1: -879}
                                                            var v103 float32 = (v28 / v28)
                                                            if ((42 >= 128) || false) {
                                                                v46 = v46
                                                                delete(v1, 881)
                                                                if ((2 > 1) || true) {
                                                                    if false {
                                                                        v2 -= v98
                                                                    } else {
                                                                        var v104 bool = false
                                                                        _ = v104
                                                                    }
                                                                } else {
                                                                    v9 = v9
                                                                    v105 := GenS3{F0: 199}
                                                                    if ((v9 == v9) && true) {
                                                                        v28 += (v103 / ((0.1 - v28) * v28))
                                                                        switch (v9 + ("hello" + v9)) {
                                                                        case "hello":
                                                                            switch (v9 + v9) {
                                                                            case "":
                                                                                _ = len(v3)
                                                                            }
                                                                            if (true && (993 <= 255)) {
                                                                                v46 = (-(v46))
                                                                                var v106 uint8 = 128
                                                                                _ = v106
                                                                            } else {
                                                                                var v107 int32 = int32(v88)
                                                                                v107 = -1
                                                                                v2 += v107
                                                                                _ = v107
                                                                            }
                                                                        case "\n\t":
                                                                            v103 = (v28 - 0.1)
                                                                        case "a":
                                                                            var v108 uint8 = 19
                                                                            _ = v108
                                                                        default:
                                                                            if (!(-728 <= int64(v88))) {
                                                                                if false {
                                                                                    var v109 uint16 = 1
                                                                                    _ = v109
                                                                                }
                                                                                if (127 < 0) {
                                                                                    var v110 uint64 = 4294967295
                                                                                    var v111 int64 = 0
                                                                                    v103 += -2.25
                                                                                    _ = v110
                                                                                    _ = v111
                                                                                }
                                                                            } else {
                                                                                if false {
                                                                                    var v112 int = -2147483648
                                                                                    if true {
                                                                                        v113 := GenS9{F0: 2, F1: 0, F2: 110}
                                                                                        _ = v113
                                                                                    }
                                                                                    var v114 uint16 = 962
                                                                                    v88 = v88
                                                                                    if true {
                                                                                        v115 := GenS10{F0: 4294967295}
                                                                                        v2 += v2
                                                                                        _ = len(v1)
                                                                                        _ = v115
                                                                                    } else {
                                                                                        v116 := GenS10{F0: v88}
                                                                                        v46 += 127
                                                                                        var v117 uint = uint(v114)
                                                                                        _ = v116
                                                                                        _ = v117
                                                                                    }
                                                                                    _ = v112
                                                                                    _ = v114
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    var v118 int = (^(fi27))
                                                                    var v119 uint32 = 4294967295
                                                                    _ = v105
                                                                    _ = v118
                                                                    _ = v119
                                                                }
                                                            }
                                                            v1[303] = 155
                                                            if (!((false && (false || ((-391 <= -765) || true))) && true)) {
                                                                v120 := GenS1{F0: (v2 / ((v2) | 1)), F1: 1, F2: 0}
                                                                fmt.Println(v88)
                                                                _ = v120
                                                            }
                                                            _ = v102
                                                            _ = v103
                                                        case 2147483647:
                                                            var v121 int32 = -2147483648
                                                            _ = v121
                                                        case 0:
                                                            v122 := GenS11{F0: 4294967295, F1: 42, F2: 887}
                                                            _ = v122
                                                        }
                                                    }
                                                    var v123 int8 = 1
                                                    var v124 uint64 = 42
                                                    if ((((v2 >> ((0) & 63)) >= 42) || false) || (255 < 255)) {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        var v125 int32 = v2
                                                        var v126 uint32 = 1
                                                        var v127 bool = true
                                                        _ = v125
                                                        _ = v126
                                                        _ = v127
                                                    } else {
                                                        fmt.Println(v46, v100, v2)
                                                        _ = len(v1)
                                                    }
                                                    if true {
                                                        var v128 uint32 = 339
                                                        _ = v128
                                                    }
                                                    _ = v123
                                                    _ = v124
                                                } else {
                                                    v129 := GenS1{F0: -2147483648, F1: v2, F2: 60}
                                                    _ = v129
                                                }
                                                _ = fi99
                                                _ = v100
                                            }
                                            _ = v98
                                        }
                                        _ = v88
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v130 int8 = v46
                                        _ = v130
                                    }
                                    var v131 uint8 = 0
                                    _ = v131
                                } else {
                                    if true {
                                        for fi132 := 0; fi132 < 6; fi132++ {
                                            var v133 uint = 18446744073709551615
                                            _ = fi132
                                            _ = v133
                                        }
                                        v2 += v2
                                        if (-2.25 > 1e10) {
                                            if (!false) {
                                                v2 -= v2
                                                var v134 int8 = 0
                                                switch (fi27 % ((-9223372036854775808) | 1)) {
                                                case 0:
                                                    var v135 int16 = 1
                                                    _ = v135
                                                case 9223372036854775807:
                                                    _ = len(v1)
                                                case 42:
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        if (1 < 0) {
                                                            if true {
                                                                delete(v1, 2147483648)
                                                                switch v9 {
                                                                case "hello":
                                                                    switch ("hello" + v9) {
                                                                    case "":
                                                                        if true {
                                                                            for fi136 := 0; fi136 < 17; fi136++ {
                                                                                v137 := GenS12{F0: 42, F1: v46}
                                                                                v138 := GenS2{F0: false}
                                                                                if false {
                                                                                    _ = len(v3)
                                                                                    var v139 int32 = 1
                                                                                    _ = v139
                                                                                } else {
                                                                                    switch "hello" {
                                                                                    case "foo bar":
                                                                                        v28 = v28
                                                                                    case "a":
                                                                                        if (true && (!(false && (false || (748 < 390))))) {
                                                                                            v140 := GenS5{F0: v9}
                                                                                            var v141 int8 = v46
                                                                                            v9 = "a"
                                                                                            _ = v140
                                                                                            _ = v141
                                                                                        }
                                                                                    }
                                                                                    var v142 int32 = v2
                                                                                    _ = v142
                                                                                }
                                                                                _ = fi136
                                                                                _ = v137
                                                                                _ = v138
                                                                            }
                                                                        } else {
                                                                            switch (fi27 ^ 0) {
                                                                            case 127:
                                                                                var v143 int64 = 9223372036854775807
                                                                                _ = v143
                                                                            default:
                                                                                var v144 float32 = 0.0
                                                                                _ = v144
                                                                            }
                                                                        }
                                                                    case "a":
                                                                        fmt.Println(v46, v46)
                                                                    case "hello":
                                                                        v2 -= v2
                                                                    default:
                                                                        switch 2147483647 {
                                                                        case 42:
                                                                            v145 := GenS13{F0: 42, F1: 42}
                                                                            _ = v145
                                                                        }
                                                                    }
                                                                default:
                                                                    for fi146 := 0; fi146 < 11; fi146++ {
                                                                        v1[272] = 2147483648
                                                                        _ = fi146
                                                                    }
                                                                }
                                                            } else {
                                                                v9 = (v9 + v9)
                                                            }
                                                        }
                                                    case "a":
                                                        v28 = (v28 + ((float32(float64(v28)) * v28) + (float32((((3.14159 + 3.14159) / (1e10 - (float64((((v28 + v28) * (v28 + (v28 * v28))) / v28)) / 3.14159))) * 0.1)) * v28)))
                                                    case "hello":
                                                        switch int(v134) {
                                                        case 2147483647:
                                                            fmt.Println(v28, v2, v46)
                                                        case 42:
                                                            var v147 float64 = 3.14159
                                                            _ = v147
                                                        case 0:
                                                            v148 := GenS1{F0: v2, F1: v2, F2: 255}
                                                            _ = v148
                                                        default:
                                                            if (2 <= 0) {
                                                                for fi149 := 0; fi149 < 10; fi149++ {
                                                                    fmt.Println(v2)
                                                                    _ = fi149
                                                                }
                                                                for fi150 := 0; fi150 < 11; fi150++ {
                                                                    v134 = 0
                                                                    _ = fi150
                                                                }
                                                                for fi151 := 0; fi151 < 6; fi151++ {
                                                                    switch ("foo bar" + (v9 + ("" + "foo bar"))) {
                                                                    case "foo bar":
                                                                        if len(v3) > 0 {
                                                                            v3[0] = 75
                                                                        }
                                                                    case "a":
                                                                        switch "" {
                                                                        case "foo bar":
                                                                            v152 := v1[53]
                                                                            _ = v152
                                                                        case "\n\t":
                                                                            var v153 float64 = 0.0
                                                                            _ = v153
                                                                        default:
                                                                            var v154 uint64 = 1
                                                                            _ = v154
                                                                        }
                                                                    }
                                                                    if true {
                                                                        _ = len(v1)
                                                                        v155 := GenS2{F0: false}
                                                                        _ = v155
                                                                    }
                                                                    v9 = "a"
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    _ = fi151
                                                                }
                                                            } else {
                                                                delete(v1, 652)
                                                                switch (fi27 | fi27) {
                                                                case 127:
                                                                    v2 = (v2 + v2)
                                                                case 42:
                                                                    switch v9 {
                                                                    case "foo bar":
                                                                        fmt.Println(v28, v134, v9)
                                                                    case "":
                                                                        for fi156 := 0; fi156 < 3; fi156++ {
                                                                            var v157 float64 = (-2.25 * 1e10)
                                                                            _ = fi156
                                                                            _ = v157
                                                                        }
                                                                    case "\n\t":
                                                                        var v158 float64 = 3.14159
                                                                        _ = v158
                                                                    default:
                                                                        var v159 bool = false
                                                                        _ = v159
                                                                    }
                                                                case -9223372036854775808:
                                                                    if (452 > 0) {
                                                                        v160 := GenS1{F0: (v2 ^ v2), F1: v2, F2: 0}
                                                                        if true {
                                                                            var v161 uint64 = 839
                                                                            var v162 uint16 = 65535
                                                                            var v163 uint16 = 1
                                                                            _ = v161
                                                                            _ = v162
                                                                            _ = v163
                                                                        } else {
                                                                            v164 := GenS11{F0: 1, F1: 528, F2: 32767}
                                                                            _ = v164
                                                                        }
                                                                        var v165 uint8 = 127
                                                                        _ = v160
                                                                        _ = v165
                                                                    }
                                                                }
                                                                var v166 uint16 = 8
                                                                _ = v166
                                                            }
                                                        }
                                                    }
                                                default:
                                                    var v167 bool = true
                                                    _ = v167
                                                }
                                                _ = v134
                                            }
                                            fmt.Println(v46, v46)
                                            v168 := GenS12{F0: 32767, F1: (v46 & -85)}
                                            _ = v168
                                        }
                                    } else {
                                        for fi169 := 0; fi169 < 15; fi169++ {
                                            v170 := GenS14{F0: -9223372036854775808, F1: (fi27 * (fi169 ^ -202))}
                                            _ = fi169
                                            _ = v170
                                        }
                                        var v171 int8 = v46
                                        _ = v171
                                    }
                                    for fi172 := 0; fi172 < 1; fi172++ {
                                        v173 := GenS11{F0: 2, F1: 65535, F2: 1}
                                        fmt.Println(v2, fi27, v46)
                                        _ = fi172
                                        _ = v173
                                    }
                                    if false {
                                        var v174 uint64 = 592
                                        var v175 int32 = v2
                                        v9 = v9
                                        _ = v174
                                        _ = v175
                                    } else {
                                        v28 -= v28
                                        v176 := v1[437]
                                        if len(v3) > 0 {
                                            v3[0] = v176
                                        }
                                        _ = v176
                                    }
                                }
                                fmt.Println(v46, v9)
                            }
                            _ = v46
                        }
                        var v177 uint64 = 4294967295
                        _ = v177
                    } else {
                        for fi178 := 0; fi178 < 18; fi178++ {
                            var v179 uint16 = 65535
                            _ = fi178
                            _ = v179
                        }
                    }
                }
            default:
                v1[2147483648] = 0
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            switch "" {
            case "hello":
                for fi180 := 0; fi180 < 16; fi180++ {
                    fmt.Println(v2)
                    var v181 string = ""
                    v182 := v1[2147483648]
                    _ = fi180
                    _ = v181
                    _ = v182
                }
            }
            v28 -= v28
            var v183 uint8 = 127
            _ = v183
        }
        _ = fi27
        _ = v28
    }
    v184 := GenS13{F0: 2, F1: 42}
    if (42 > 2147483647) {
        var v185 string = (("a" + v9) + (((v9 + v9) + (v9 + "foo bar")) + "foo bar"))
        _ = v185
    } else {
        fmt.Println(v9)
        v9 = "foo bar"
        v2 += (v2 >> ((542) & 63))
    }
    for fi186 := 0; fi186 < 4; fi186++ {
        switch fi186 {
        case 32768:
            v2 = (v2 * v2)
        case -350:
            var v187 int8 = -1
            _ = v187
        case 664:
            v3 = append(v3, 2)
        }
        v9 = v9
        _ = fi186
    }
    v2 = (-(v2))
    switch 1 {
    case 1:
        if ((!(0 != 4294967295)) && (!((float64(3.14159) - 0.1) <= float64(-2.25)))) {
            var v188 int32 = v2
            _ = v188
        }
    case 455:
        var v189 uint64 = 2
        _ = v189
    }
    var v190 int64 = -9223372036854775808
    v9 = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v190:", v190)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v184
    _ = v190
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
