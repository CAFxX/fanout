package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 string }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 int32; F1 int32; F2 bool }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int8; F1 string; F2 int16 }
type GenS9 struct { F0 bool; F1 int32; F2 int16 }
type GenS10 struct { F0 uint; F1 int64; F2 int8 }
type GenS11 struct { F0 uint; F1 uint; F2 int64 }
type GenS12 struct { F0 int64; F1 uint8 }
type GenS13 struct { F0 int16; F1 int8; F2 string }
type GenS14 struct { F0 string; F1 int8; F2 int }
func genFn0() int64 {
    var v6 uint16 = 32767
    v7 := []int{0, 662}
    v6 = v6
    var v8 string = "\n\t"
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (false && true)
    var v3 uint8 = 2
    var v4 int16 = 1
    var v5 int = 42
    switch 0 {
    case -546:
        for fi9 := 0; fi9 < 7; fi9++ {
            v5 += -1
            if v2 {
                var v10 float32 = 0.0
                _ = v10
            }
            _ = fi9
        }
    case 928:
        if ((733 != 2147483647) || (!false)) {
            v2 = v2
            for fi11 := 0; fi11 < 17; fi11++ {
                for fi12 := 0; fi12 < 10; fi12++ {
                    _ = len(v1)
                    switch "hello" {
                    case "hello":
                        switch ("hello" + ("a" + "foo bar")) {
                        case "foo bar":
                            fmt.Println(fi11)
                            var v13 int32 = 0
                            var v14 int = 1
                            var v15 int32 = v13
                            var v16 string = "\n\t"
                            _ = v13
                            _ = v14
                            _ = v15
                            _ = v16
                        case "hello":
                            switch (fi11 - v5) {
                            case 42:
                                if (65535 < 42) {
                                    fmt.Println(fi12)
                                    var v17 int8 = 100
                                    switch (fi11 & fi11) {
                                    case -235:
                                        v18 := GenS1{F0: 18446744073709551615, F1: (v4 + v4), F2: "hello"}
                                        _ = v18
                                    default:
                                        var v19 int8 = -128
                                        _ = v19
                                    }
                                    _ = v17
                                } else {
                                    fmt.Println(fi11)
                                    v1[393] = (v5 >> ((4294967295) & 63))
                                }
                            case -227:
                                if false {
                                    var v20 int = 127
                                    switch ("\n\t" + "a") {
                                    case "hello":
                                        switch ("" + "") {
                                        case "foo bar":
                                            fmt.Println(fi12, fi11)
                                        case "":
                                            v21 := []int{v5, -2147483648, v5}
                                            _ = v21
                                        case "a":
                                            v1[1] = (v20 ^ v5)
                                        default:
                                            v22 := v1[1]
                                            _ = v22
                                        }
                                    default:
                                        var v23 uint32 = 2
                                        _ = v23
                                    }
                                    var v24 uint64 = 217
                                    _ = v20
                                    _ = v24
                                } else {
                                    if false {
                                        var v25 uint8 = v3
                                        fmt.Println(fi12, v5)
                                        _ = v25
                                    } else {
                                        var v26 int = 32767
                                        var v27 float64 = 1.5
                                        v28 := v1[4294967295]
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    }
                                    switch ("a" + ("foo bar" + "")) {
                                    case "a":
                                        if v2 {
                                            v4 = (v4 | -32768)
                                            switch (fi11 >> ((0) & 63)) {
                                            case 32767:
                                                if (602 <= 32767) {
                                                    if v2 {
                                                        fmt.Println(v3, v2, v2)
                                                        var v29 uint8 = v3
                                                        if v2 {
                                                            fmt.Println(v3)
                                                            var v30 bool = v2
                                                            _ = v30
                                                        }
                                                        v4 = v4
                                                        _ = v29
                                                    }
                                                }
                                            case -680:
                                                if ("foo bar" <= "") {
                                                    switch 32767 {
                                                    case 2147483647:
                                                        if ((-1 >= 813) || v2) {
                                                            v2 = (0 < 832)
                                                            v31 := []int{32767, 32768, (fi11 | (fi11 | (fi11 / (((v5 * v5)) | 1)))), fi11}
                                                            _ = v31
                                                        }
                                                        v32 := GenS1{F0: 1, F1: -1, F2: (("foo bar" + "\n\t") + ("\n\t" + ""))}
                                                        v33 := GenS2{F0: -2147483648, F1: 0}
                                                        var v34 uint64 = 42
                                                        _ = v32
                                                        _ = v33
                                                        _ = v34
                                                    case 32768:
                                                        var v35 int16 = (v4 / ((v4) | 1))
                                                        _ = v35
                                                    }
                                                    v36 := []int{(^(v5)), (fi11 | (fi11 * (fi11 >> ((1) & 63)))), v5}
                                                    _ = v36
                                                }
                                            default:
                                                switch ("hello" + "\n\t") {
                                                case "a":
                                                    v37 := GenS2{F0: -665, F1: 529}
                                                    _ = v37
                                                case "":
                                                    var v38 uint16 = 814
                                                    _ = v38
                                                case "\n\t":
                                                    if true {
                                                        switch v5 {
                                                        case 1:
                                                            fmt.Println(v2)
                                                        case -1:
                                                            v39 := GenS3{F0: 426, F1: 345}
                                                            v3 = v3
                                                            if true {
                                                                v4 -= (v4 + v4)
                                                                v40 := GenS2{F0: -251, F1: 65535}
                                                                v41 := GenS4{F0: v4, F1: 999}
                                                                var v42 uint32 = 840
                                                                _ = v40
                                                                _ = v41
                                                                _ = v42
                                                            }
                                                            _ = v39
                                                        default:
                                                            v3 += v3
                                                        }
                                                        if (3.14159 < (((0.0 / 1.5) * 1.5) / 1e10)) {
                                                            if true {
                                                                var v43 int = fi12
                                                                if ("foo bar" != "\n\t") {
                                                                    v44 := GenS2{F0: -31, F1: 32767}
                                                                    if v2 {
                                                                        fmt.Println(v3)
                                                                        if (2147483648 >= 106) {
                                                                            var v45 int = (fi11 >> ((278) & 63))
                                                                            v46 := []int{(v5 ^ int(v3)), (^(fi11)), (fi11 << ((223) & 63))}
                                                                            if ((1e10 * float32(-2.25)) <= 3.14159) {
                                                                                var v47 uint32 = 0
                                                                                _ = v47
                                                                            } else {
                                                                                fmt.Println(v2, fi11)
                                                                                var v48 int32 = -324
                                                                                _ = v48
                                                                            }
                                                                            _ = v45
                                                                            _ = v46
                                                                        }
                                                                    } else {
                                                                        v43 -= int(v3)
                                                                        v2 = (1 > v3)
                                                                        v2 = (2147483648 < 729)
                                                                    }
                                                                    _ = v44
                                                                }
                                                                v5 -= fi12
                                                                _ = v43
                                                            }
                                                            v49 := GenS2{F0: 439, F1: 32767}
                                                            v2 = (true && false)
                                                            _ = v49
                                                        }
                                                        var v50 float32 = (float32(1e10) / ((0.1 - ((0.1 * (float32(1.5) - 3.14159)) - 1.5)) / (((1.5 * (((((((0.1 - -2.25) / 1.0) / 0.1) + float32(float64(0.1))) + (1e10 / (0.1 / 0.1))) / 1.0) + 1.5)) * (0.0 + 3.14159)) * ((1.5 - (float32((0.0 / ((0.0 * (0.1 / (1.5 / ((3.14159 - 1.5) + (1e10 * float64(1e10)))))) + 1.5))) * 0.1)) + 3.14159))))
                                                        fmt.Println(v4, v5)
                                                        _ = v50
                                                    } else {
                                                        v51 := []int{fi12, 42}
                                                        v2 = v2
                                                        fmt.Println(v2, v5)
                                                        _ = v51
                                                    }
                                                default:
                                                    v52 := GenS1{F0: 2, F1: v4, F2: "a"}
                                                    _ = v52
                                                }
                                            }
                                            var v53 int64 = genFn0()
                                            var v54 int16 = v4
                                            switch -2147483648 {
                                            case 162:
                                                var v55 uint16 = 1
                                                _ = v55
                                            case 991:
                                                if (!false) {
                                                    v56 := GenS5{F0: 42, F1: int32(v3), F2: true}
                                                    var v57 string = ("foo bar" + ("a" + "hello"))
                                                    _ = v56
                                                    _ = v57
                                                } else {
                                                    fmt.Println(v2, v5)
                                                    var v58 uint8 = v3
                                                    _ = len(v1)
                                                    _ = v58
                                                }
                                            case -758:
                                                v54 -= 0
                                            default:
                                                v2 = v2
                                            }
                                            switch (fi12 | v5) {
                                            case 128:
                                                v3 = (^(v3))
                                            case 32767:
                                                fmt.Println(v53)
                                            default:
                                                var v59 float32 = 1.5
                                                _ = v59
                                            }
                                            _ = v53
                                            _ = v54
                                        } else {
                                            var v60 uint32 = 2147483647
                                            _ = v60
                                        }
                                    case "":
                                        v2 = v2
                                    case "hello":
                                        var v61 int8 = 79
                                        _ = v61
                                    default:
                                        switch "foo bar" {
                                        case "a":
                                            var v62 int16 = v4
                                            _ = v62
                                        case "\n\t":
                                            var v63 int16 = v4
                                            _ = v63
                                        case "":
                                            var v64 int32 = 42
                                            _ = v64
                                        default:
                                            var v65 uint16 = 775
                                            var v66 int = fi12
                                            v67 := GenS6{F0: 1, F1: -33}
                                            _ = v65
                                            _ = v66
                                            _ = v67
                                        }
                                    }
                                    if (0 == 18446744073709551615) {
                                        fmt.Println(v3)
                                        var v68 float64 = 1e10
                                        v69 := GenS3{F0: -215, F1: 4294967295}
                                        _ = v68
                                        _ = v69
                                    }
                                    v70 := GenS7{F0: 127}
                                    _ = v70
                                }
                            default:
                                var v71 int64 = 0
                                _ = v71
                            }
                        }
                    case "":
                        if (-203 < 748) {
                            v5 = (^(fi11))
                            fmt.Println(v4, v4)
                        } else {
                            var v72 int16 = (v4 << ((0) & 63))
                            var v73 uint64 = 0
                            _ = v72
                            _ = v73
                        }
                        v1[2147483648] = 42
                    case "a":
                        v74 := GenS5{F0: 740, F1: 293, F2: (0.0 >= (float32(float64((0.0 * (float32(1e10) + 1e10)))) / 1e10))}
                        _ = v74
                    default:
                        if v2 {
                            v4 += v4
                            switch "foo bar" {
                            case "a":
                                v75 := GenS8{F0: 42, F1: ("foo bar" + "\n\t"), F2: -1}
                                _ = v75
                            case "foo bar":
                                v76 := []int{fi12, fi12, -1, fi12}
                                _ = v76
                            case "hello":
                                v77 := []int{fi12, (fi12 % ((v5) | 1))}
                                _ = v77
                            }
                        } else {
                            var v78 uint16 = 2
                            var v79 string = ""
                            v78 += v78
                            _ = v78
                            _ = v79
                        }
                    }
                    _ = fi12
                }
                v3 += (v3 * (v3 * v3))
                if (2 >= 694) {
                    v5 = v5
                    if v2 {
                        var v80 int32 = 0
                        var v81 int32 = (v80 | (v80 ^ 1))
                        _ = v80
                        _ = v81
                    } else {
                        var v82 uint16 = 65535
                        v83 := []int{}
                        _ = v82
                        _ = v83
                    }
                    v2 = v2
                }
                _ = fi11
            }
            v84 := GenS8{F0: 100, F1: "", F2: v4}
            _ = v84
        } else {
            switch v5 {
            case 4294967295:
                if (!(!v2)) {
                    fmt.Println(v3, v4)
                } else {
                    fmt.Println(v4, v5, v3)
                    if (v2 || true) {
                        var v85 int8 = 0
                        var v86 int64 = -1
                        switch v5 {
                        case 9223372036854775807:
                            var v87 uint8 = (v3 - (v3 >> ((4294967295) & 63)))
                            _ = v87
                        case 0:
                            if (933 < int32(v86)) {
                                for fi88 := 0; fi88 < 5; fi88++ {
                                    var v89 int32 = 384
                                    _ = fi88
                                    _ = v89
                                }
                                for fi90 := 0; fi90 < 17; fi90++ {
                                    if v2 {
                                        if false {
                                            v91 := []int{}
                                            var v92 uint16 = 42
                                            for fi93 := 0; fi93 < 19; fi93++ {
                                                var v94 int = fi93
                                                fmt.Println(fi90, v85)
                                                _ = fi93
                                                _ = v94
                                            }
                                            _ = v91
                                            _ = v92
                                        }
                                        if v2 {
                                            var v95 bool = ((!false) && ((1 < 1) && (!((v86 ^ 0) == 265))))
                                            _ = v95
                                        }
                                    }
                                    v85 = 39
                                    for fi96 := 0; fi96 < 18; fi96++ {
                                        var v97 uint16 = 32767
                                        v98 := GenS9{F0: false, F1: 0, F2: (v4 | -630)}
                                        var v99 uint64 = 42
                                        _ = fi96
                                        _ = v97
                                        _ = v98
                                        _ = v99
                                    }
                                    _ = fi90
                                }
                            } else {
                                var v100 int8 = (v85 >> ((70) & 63))
                                for fi101 := 0; fi101 < 3; fi101++ {
                                    fmt.Println(v85, v3)
                                    _ = fi101
                                }
                                var v102 float32 = 3.14159
                                _ = v100
                                _ = v102
                            }
                        }
                        _ = v85
                        _ = v86
                    }
                    if (-32768 != v4) {
                        v103 := []int{v5, v5, 2147483647}
                        _ = v103
                    } else {
                        switch 648 {
                        case 42:
                            v2 = true
                        case 127:
                            var v104 float64 = float64(0.0)
                            _ = v104
                        case 4294967295:
                            if (65535 != 0) {
                                if (715 >= -1) {
                                    switch (v5 % ((v5) | 1)) {
                                    case -2147483648:
                                        for fi105 := 0; fi105 < 18; fi105++ {
                                            switch "a" {
                                            case "\n\t":
                                                for fi106 := 0; fi106 < 12; fi106++ {
                                                    v1[251] = (fi105 + 9223372036854775807)
                                                    _ = fi106
                                                }
                                            default:
                                                v107 := []int{788, v5}
                                                _ = v107
                                            }
                                            _ = fi105
                                        }
                                    case 128:
                                        var v108 int8 = -128
                                        _ = v108
                                    default:
                                        v109 := []int{}
                                        _ = v109
                                    }
                                } else {
                                    if (695 < 42) {
                                        for fi110 := 0; fi110 < 6; fi110++ {
                                            v2 = (-2.25 > 3.14159)
                                            v111 := GenS10{F0: 18446744073709551615, F1: 769, F2: -78}
                                            _ = fi110
                                            _ = v111
                                        }
                                        for fi112 := 0; fi112 < 1; fi112++ {
                                            v5 -= v5
                                            for fi113 := 0; fi113 < 13; fi113++ {
                                                v4 = v4
                                                v114 := GenS1{F0: 255, F1: (v4 / ((-1) | 1)), F2: "\n\t"}
                                                _ = fi113
                                                _ = v114
                                            }
                                            _ = fi112
                                        }
                                        v115 := GenS9{F0: (4294967295 >= 2), F1: -70, F2: v4}
                                        _ = v115
                                    } else {
                                        v4 = (v4 << ((42) & 63))
                                        v1[2] = (v5 * v5)
                                    }
                                }
                                fmt.Println(v2, v3)
                                var v116 int64 = 0
                                _ = v116
                            } else {
                                v117 := []int{(-(v5)), 32767}
                                if v2 {
                                    for fi118 := 0; fi118 < 16; fi118++ {
                                        var v119 uint8 = v3
                                        _ = fi118
                                        _ = v119
                                    }
                                }
                                _ = v117
                            }
                        default:
                            switch "a" {
                            case "hello":
                                v120 := GenS1{F0: 0, F1: (v4 >> ((4294967295) & 63)), F2: ("a" + "\n\t")}
                                _ = v120
                            }
                        }
                        if v2 {
                            delete(v1, 1)
                            v121 := v1[65535]
                            _ = v121
                        } else {
                            v122 := GenS2{F0: 2147483647, F1: 804}
                            _ = v122
                        }
                    }
                }
            case -1:
                v123 := []int{}
                _ = v123
            case 42:
                for fi124 := 0; fi124 < 1; fi124++ {
                    v125 := GenS11{F0: uint(v3), F1: 853, F2: -9223372036854775808}
                    _ = fi124
                    _ = v125
                }
            default:
                v126 := []int{}
                _ = v126
            }
            v4 += v4
        }
    case 42:
        v127 := GenS5{F0: 1, F1: -758, F2: (!false)}
        _ = v127
    default:
        if (!(!((0 <= 18446744073709551615) || (true && (1e10 != float32(float64((float32((0.0 + 1e10)) - 3.14159)))))))) {
            if true {
                v128 := GenS12{F0: genFn0(), F1: (v3 << ((255) & 63))}
                _ = v128
            }
        }
    }
    v3 = (v3 >> ((4294967295) & 63))
    if (!v2) {
        var v129 float64 = float64(1e10)
        for fi130 := 0; fi130 < 20; fi130++ {
            var v131 uint8 = (v3 & (^(v3)))
            fmt.Println(v4, v129, v5)
            _ = fi130
            _ = v131
        }
        _ = v129
    }
    v132 := []int{v5, v5}
    _ = len(v132)
    var v133 int16 = (v4 | int16(v3))
    switch -102 {
    case 42:
        if false {
            switch (("a" + "hello") + "hello") {
            case "":
                switch ("" + ("\n\t" + ("foo bar" + "foo bar"))) {
                case "hello":
                    var v134 string = ("" + "foo bar")
                    _ = v134
                case "foo bar":
                    switch ("hello" + ("a" + "a")) {
                    case "foo bar":
                        fmt.Println(v2)
                    case "\n\t":
                        for fi135 := 0; fi135 < 14; fi135++ {
                            var v136 int = (v5 - fi135)
                            _ = fi135
                            _ = v136
                        }
                    }
                default:
                    for fi137 := 0; fi137 < 20; fi137++ {
                        fmt.Println(v5)
                        switch "hello" {
                        case "foo bar":
                            if (!(716 >= 269)) {
                                v133 = v4
                                switch (("\n\t" + ("foo bar" + "hello")) + "hello") {
                                case "\n\t":
                                    var v138 int8 = 127
                                    _ = v138
                                case "a":
                                    v139 := GenS8{F0: 1, F1: "foo bar", F2: (v4 >> ((18446744073709551615) & 63))}
                                    _ = v139
                                }
                            } else {
                                v5 -= (v5 % (((fi137 << ((18446744073709551615) & 63))) | 1))
                            }
                        case "hello":
                            v2 = false
                        case "a":
                            fmt.Println(v4)
                        default:
                            fmt.Println(v5, v2, v5)
                        }
                        if (v2 && true) {
                            v140 := GenS4{F0: (^(v4)), F1: 42}
                            _ = len(v132)
                            v141 := GenS4{F0: v4, F1: 4294967295}
                            _ = v140
                            _ = v141
                        }
                        _ = fi137
                    }
                }
            default:
                v142 := GenS6{F0: (v4 % ((v133) | 1)), F1: 1}
                _ = v142
            }
            v143 := GenS13{F0: v133, F1: -11, F2: "\n\t"}
            if (!(935 == 322)) {
                var v144 int = 4294967295
                v145 := GenS14{F0: ("\n\t" + "foo bar"), F1: -119, F2: v5}
                if false {
                    delete(v1, 65535)
                    var v146 uint16 = 65535
                    v147 := GenS13{F0: 687, F1: -1, F2: "\n\t"}
                    _ = v146
                    _ = v147
                }
                _ = v144
                _ = v145
            }
            _ = v143
        } else {
            var v148 float32 = -2.25
            switch (v5 | v5) {
            case 32768:
                if v2 {
                    for fi149 := 0; fi149 < 6; fi149++ {
                        var v150 int64 = genFn0()
                        _ = fi149
                        _ = v150
                    }
                    v148 = v148
                    v151 := GenS12{F0: 42, F1: v3}
                    _ = v151
                }
            }
            _ = v148
        }
    }
    v152 := GenS13{F0: (v133 * v4), F1: 127, F2: "\n\t"}
    v5 += 32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v132:", len(v132), v132)
    fmt.Println("v133:", v133)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v132
    _ = v133
    _ = v152
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
