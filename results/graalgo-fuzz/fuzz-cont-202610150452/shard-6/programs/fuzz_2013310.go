package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 int64; F1 uint32 }
type GenS8 struct { F0 uint; F1 int8; F2 int16 }
type GenS9 struct { F0 int; F1 int32 }
type GenS10 struct { F0 uint64; F1 int8; F2 int32 }
type GenS11 struct { F0 int64; F1 int64 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int; F1 uint }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 string; F1 uint; F2 bool }
type GenS16 struct { F0 uint16; F1 int16 }
type GenS17 struct { F0 uint32; F1 int16; F2 int8 }
type GenS18 struct { F0 int64; F1 int16; F2 bool }
type GenS19 struct { F0 int32; F1 bool }
type GenS20 struct { F0 uint64; F1 uint; F2 uint }
type GenS21 struct { F0 uint; F1 int }
func genFn0() int64 {
    switch "foo bar" {
    case "a":
        for fi6 := 0; fi6 < 3; fi6++ {
            v7 := GenS1{F0: 127, F1: 373, F2: 1}
            _ = fi6
            _ = v7
        }
    default:
        var v8 int16 = 32767
        _ = v8
    }
    v9 := GenS2{F0: 855, F1: 9223372036854775807}
    var v10 int64 = -1
    if true {
        v11 := GenS3{F0: 1}
        _ = v11
    } else {
        if false {
            var v12 uint8 = 127
            switch -159 {
            case 0:
                fmt.Println(v12, v12)
            }
            _ = v12
        }
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 int16 = 0
            v10 -= v10
            switch "a" {
            case "":
                var v15 float32 = -2.25
                _ = v15
            case "\n\t":
                if false {
                    var v16 uint64 = 823
                    v16 -= (v16 * v16)
                    v16 -= (^(v16))
                    _ = v16
                } else {
                    v17 := []int{fi13}
                    if (true || false) {
                        fmt.Println(v14, fi13)
                        fmt.Println(fi13)
                    }
                    _ = v17
                }
            case "a":
                var v18 string = "hello"
                _ = v18
            default:
                if (!false) {
                    fmt.Println(v10)
                    var v19 uint16 = 110
                    var v20 int = 32767
                    _ = v19
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 9; fi21++ {
                        v22 := []int{(fi21 | fi21)}
                        v23 := GenS4{F0: ((-1 >= 42) || true), F1: fi21}
                        switch fi13 {
                        case 982:
                            switch fi13 {
                            case 32767:
                                v14 -= (v14 & v14)
                            default:
                                if false {
                                    v24 := GenS3{F0: 1}
                                    var v25 string = "a"
                                    _ = len(v22)
                                    if (!(3.14159 == 1e10)) {
                                        var v26 int64 = (v10 << ((394) & 63))
                                        fmt.Println(v25)
                                        _ = v26
                                    } else {
                                        v10 = (-(v10))
                                        v14 += v14
                                        fmt.Println(v10)
                                        fmt.Println(v10, fi13, fi21)
                                        v27 := GenS5{F0: v14}
                                        _ = v27
                                    }
                                    _ = v24
                                    _ = v25
                                }
                            }
                        case 665:
                            v28 := GenS4{F0: (!false), F1: 128}
                            _ = v28
                        case 32767:
                            switch "hello" {
                            case "hello":
                                if (0 <= 1) {
                                    var v29 int64 = (v10 << ((345) & 63))
                                    _ = v29
                                } else {
                                    switch ("\n\t" + "\n\t") {
                                    case "":
                                        switch "foo bar" {
                                        case "":
                                            v30 := GenS3{F0: 42}
                                            _ = v30
                                        }
                                    case "foo bar":
                                        if ((-(fi13)) == fi21) {
                                            if (1 < 50) {
                                                var v31 float32 = ((0.0 * (3.14159 + 1.5)) + 0.0)
                                                _ = v31
                                            } else {
                                                switch "hello" {
                                                case "":
                                                    fmt.Println(fi13, v10, v10)
                                                case "a":
                                                    var v32 uint32 = 2
                                                    _ = v32
                                                default:
                                                    var v33 uint16 = 65535
                                                    _ = v33
                                                }
                                                var v34 int64 = -325
                                                _ = v34
                                            }
                                            v14 = v14
                                        } else {
                                            v10 = 0
                                            fmt.Println(v10, fi21)
                                        }
                                    case "\n\t":
                                        var v35 int32 = -291
                                        _ = v35
                                    }
                                    v36 := GenS5{F0: (v14 * (v14 * v14))}
                                    var v37 int32 = -606
                                    _ = v36
                                    _ = v37
                                }
                            }
                        }
                        var v38 bool = (108 <= 134)
                        _ = fi21
                        _ = v22
                        _ = v23
                        _ = v38
                    }
                    fmt.Println(v10, fi13)
                    v10 -= v10
                }
            }
            _ = fi13
            _ = v14
        }
        v39 := GenS4{F0: true, F1: -477}
        _ = v39
    }
    _ = v10
    _ = v9
    return 443
}

func main() {
    var v1 int8 = 42
    var v2 string = "a"
    v3 := make(map[uint32]bool)
    var v4 float32 = 0.0
    var v5 uint16 = 255
    if false {
        var v40 uint = 401
        v41 := GenS6{F0: (v2 + v2), F1: (v5 + v5)}
        var v42 int = int(v1)
        _ = v40
        _ = v41
        _ = v42
    } else {
        fmt.Println(v1)
        switch ("a" + "") {
        case "hello":
            v5 -= (v5 / (((v5 - (^(v5)))) | 1))
        case "":
            for fi43 := 0; fi43 < 17; fi43++ {
                v4 += v4
                _ = fi43
            }
        default:
            for fi44 := 0; fi44 < 17; fi44++ {
                for fi45 := 0; fi45 < 19; fi45++ {
                    v46 := GenS6{F0: "", F1: (v5 & 194)}
                    var v47 int = fi44
                    _ = fi45
                    _ = v46
                    _ = v47
                }
                _ = fi44
            }
        }
    }
    for fi48 := 0; fi48 < 15; fi48++ {
        var v49 int8 = 88
        if (!true) {
            var v50 bool = true
            switch 4294967295 {
            case -9223372036854775808:
                fmt.Println(v4, fi48, v1)
            case -64:
                v51 := GenS7{F0: -9223372036854775808, F1: 2}
                _ = v51
            default:
                v52 := GenS1{F0: (fi48 * 857), F1: 1000, F2: 203}
                _ = v52
            }
            for fi53 := 0; fi53 < 16; fi53++ {
                v5 += 65535
                v54 := GenS8{F0: 42, F1: 72, F2: int16(v49)}
                _ = fi53
                _ = v54
            }
            _ = v50
        } else {
            for fi55 := 0; fi55 < 13; fi55++ {
                fmt.Println(fi48)
                var v56 int8 = (v49 >> ((4294967295) & 63))
                if true {
                    switch v2 {
                    case "hello":
                        switch fi48 {
                        case 0:
                            v57 := []int{(fi48 >> ((565) & 63)), (fi55 % ((fi55) | 1))}
                            v2 = v2
                            _ = v57
                        case 32768:
                            switch (((v2 + v2) + v2) + (("hello" + v2) + (v2 + (v2 + v2)))) {
                            case "hello":
                                v58 := GenS9{F0: (fi55 & fi55), F1: 42}
                                _ = v58
                            case "foo bar":
                                if true {
                                    v59 := []int{(fi48 | -540), 44, fi55, (fi55 ^ (^(fi55)))}
                                    fmt.Println(v1, v5)
                                    _ = v59
                                } else {
                                    switch (-(fi55)) {
                                    case 555:
                                        var v60 string = ((v2 + (v2 + v2)) + "")
                                        _ = v60
                                    case -663:
                                        v61 := GenS10{F0: 0, F1: (v1 + 0), F2: 0}
                                        _ = v61
                                    }
                                }
                                switch fi48 {
                                case 1:
                                    if true {
                                        var v62 float32 = ((v4 - v4) * v4)
                                        _ = v62
                                    } else {
                                        if ((v5 % ((255) | 1)) >= v5) {
                                            if (42 < (v5 << ((0) & 63))) {
                                                var v63 uint8 = 42
                                                _ = v63
                                            } else {
                                                switch int(v5) {
                                                case -657:
                                                    v64 := v3[65535]
                                                    _ = v64
                                                default:
                                                    v65 := GenS3{F0: 456}
                                                    _ = v65
                                                }
                                                if false {
                                                    var v66 int8 = v49
                                                    _ = v66
                                                } else {
                                                    v67 := GenS2{F0: 32767, F1: -88}
                                                    v49 = (v49 * v56)
                                                    v56 -= (v49 % ((73) | 1))
                                                    _ = v67
                                                }
                                                if true {
                                                    v56 -= -128
                                                    if true {
                                                        fmt.Println(v2, v49, v49)
                                                    }
                                                    if (18446744073709551615 <= 18446744073709551615) {
                                                        var v68 uint32 = 882
                                                        v69 := []int{fi55}
                                                        var v70 int = (fi48 - fi48)
                                                        if (true || true) {
                                                            if len(v69) > 0 {
                                                                v69[0] = -9223372036854775808
                                                            }
                                                            if true {
                                                                if false {
                                                                    v3[2] = (true && true)
                                                                    v56 = v49
                                                                } else {
                                                                    var v71 uint16 = v5
                                                                    _ = v71
                                                                }
                                                                v56 = v1
                                                                var v72 int32 = -2147483648
                                                                v1 = v56
                                                                if (!(v2 > "")) {
                                                                    if true {
                                                                        var v73 int = (v70 / (((^(fi55))) | 1))
                                                                        _ = v73
                                                                    }
                                                                    if ((65535 == 65535) || (true || true)) {
                                                                        fmt.Println(v5, v49, v68)
                                                                        v56 = (v56 - (v56 - int8(v72)))
                                                                        if true {
                                                                            var v74 uint32 = (v68 ^ (v68 ^ (^(v68))))
                                                                            var v75 int = fi55
                                                                            _ = v74
                                                                            _ = v75
                                                                        }
                                                                    }
                                                                    var v76 int32 = v72
                                                                    _ = v76
                                                                }
                                                                _ = v72
                                                            }
                                                            v77 := GenS11{F0: 1, F1: 28}
                                                            _ = v77
                                                        }
                                                        switch fi48 {
                                                        case 127:
                                                            v69 = append(v69, (fi55 / ((fi48) | 1)))
                                                        case 128:
                                                            var v78 int8 = (v56 + (v56 >> ((621) & 63)))
                                                            _ = v78
                                                        }
                                                        v4 = (-2.25 - ((1.5 / ((float32(1e10) - 1e10) + v4)) - v4))
                                                        var v79 int = v70
                                                        _ = v68
                                                        _ = v69
                                                        _ = v70
                                                        _ = v79
                                                    }
                                                    var v80 uint64 = 42
                                                    _ = v80
                                                }
                                            }
                                        }
                                    }
                                case -1:
                                    v5 = (^(v5))
                                default:
                                    if false {
                                        v2 = v2
                                        switch fi55 {
                                        case 303:
                                            if true {
                                                var v81 float32 = v4
                                                v82 := GenS12{F0: 127}
                                                v81 -= float32(1.5)
                                                _ = v81
                                                _ = v82
                                            } else {
                                                v83 := []int{fi48, (fi48 << ((1) & 63)), -1, 4294967295}
                                                v2 = ("hello" + (v2 + v2))
                                                v2 = "\n\t"
                                                v5 += (v5 ^ v5)
                                                _ = v83
                                            }
                                        }
                                        var v84 string = ""
                                        _ = v84
                                    } else {
                                        var v85 int16 = 1
                                        _ = v85
                                    }
                                }
                                fmt.Println(v1, v4, v49)
                            case "":
                                v5 = v5
                            }
                        case -914:
                            var v86 int8 = v49
                            _ = v86
                        default:
                            v1 = v49
                        }
                    case "foo bar":
                        fmt.Println(v4, fi55)
                    case "\n\t":
                        v87 := GenS13{F0: (fi48 % ((fi55) | 1)), F1: 0}
                        if true {
                            var v88 uint = 65535
                            var v89 float64 = 1.5
                            _ = v88
                            _ = v89
                        } else {
                            switch (fi55 - (^(fi48))) {
                            case 42:
                                v90 := GenS14{F0: (^(v1))}
                                _ = v90
                            case 32768:
                                v4 += ((0.1 * v4) + (v4 * float32(1e10)))
                            case 32767:
                                v5 -= v5
                            }
                            v91 := []int{(fi48 << ((430) & 63)), fi48, 127, (fi55 << ((0) & 63))}
                            _ = v91
                        }
                        var v92 uint16 = 595
                        _ = v87
                        _ = v92
                    }
                    var v93 int16 = 1
                    if (255 > 4294967295) {
                        var v94 int = 835
                        fmt.Println(v1)
                        switch ((v2 + v2) + ("\n\t" + "a")) {
                        case "":
                            fmt.Println(v93, v1)
                        case "\n\t":
                            switch (v2 + v2) {
                            case "a":
                                v95 := v3[42]
                                _ = v95
                            case "":
                                var v96 int8 = (v49 - v49)
                                _ = len(v3)
                                _ = v96
                            case "foo bar":
                                v5 = v5
                                var v97 float32 = (v4 + v4)
                                if (false || false) {
                                    v98 := GenS9{F0: fi48, F1: -2147483648}
                                    if true {
                                        switch v94 {
                                        case 46:
                                            v99 := []int{(^(fi48)), -974}
                                            _ = v99
                                        }
                                        fmt.Println(v93, v93)
                                        switch 2147483647 {
                                        case 9223372036854775807:
                                            switch v94 {
                                            case -356:
                                                var v100 float32 = 0.1
                                                _ = v100
                                            case -797:
                                                v101 := []int{(fi48 - (^(fi48)))}
                                                _ = v101
                                            case -2147483648:
                                                if ((1e10 + (1e10 + float64(float32((1e10 / float64(float32(0.1))))))) > (1e10 + 3.14159)) {
                                                    if (1e10 != (float32(0.0) - (v97 / 0.1))) {
                                                        var v102 float32 = float32(3.14159)
                                                        _ = v102
                                                    }
                                                    var v103 uint16 = (v5 & (^(v5)))
                                                    _ = v103
                                                } else {
                                                    if false {
                                                        v104 := GenS11{F0: -9223372036854775808, F1: genFn0()}
                                                        v49 += -128
                                                        v97 = float32(1e10)
                                                        _ = v104
                                                    }
                                                    fmt.Println(v94)
                                                    v105 := []int{(v94 + int(v56)), (fi55 | (fi55 << ((255) & 63))), (-(fi55))}
                                                    _ = v105
                                                }
                                            }
                                        case 1:
                                            switch v2 {
                                            case "":
                                                v106 := GenS15{F0: ("a" + "foo bar"), F1: 2, F2: ((201 > int32(v5)) || false)}
                                                _ = v106
                                            default:
                                                fmt.Println(v93, v56, fi48)
                                            }
                                        default:
                                            v4 = float32(3.14159)
                                        }
                                    } else {
                                        _ = len(v3)
                                        v1 -= v56
                                        v4 -= (3.14159 + float32(-2.25))
                                    }
                                    var v107 uint32 = 632
                                    _ = v107
                                    _ = v98
                                } else {
                                    var v108 uint = 65535
                                    _ = v108
                                }
                                _ = v97
                            }
                        case "foo bar":
                            switch (v94 - (v94 ^ v94)) {
                            case 115:
                                if true {
                                    v93 = (v93 * (v93 * v93))
                                    fmt.Println(v5)
                                }
                            case -289:
                                v109 := []int{}
                                _ = v109
                            default:
                                var v110 float64 = float64(((v4 - (v4 + 0.0)) - 0.0))
                                _ = v110
                            }
                        }
                        var v111 uint16 = 0
                        _ = v111
                        _ = v94
                    } else {
                        v112 := GenS16{F0: v5, F1: v93}
                        _ = v112
                    }
                    _ = v93
                } else {
                    v5 = v5
                    var v113 uint64 = 192
                    if (1 >= 509) {
                        v4 = (((v4 + 0.1) / ((1e10 - 3.14159) + v4)) * (v4 - 0.0))
                        var v114 uint = 526
                        if (v113 == (^(v113))) {
                            var v115 uint8 = 98
                            v2 = "\n\t"
                            v116 := GenS12{F0: v115}
                            var v117 uint = (v114 + v114)
                            _ = v115
                            _ = v116
                            _ = v117
                        } else {
                            var v118 float32 = 1e10
                            _ = v118
                        }
                        _ = v114
                    }
                    _ = v113
                }
                fmt.Println(v1)
                fmt.Println(v49, fi55, v5)
                v119 := GenS1{F0: (fi55 * (fi48 % ((fi48) | 1))), F1: 0, F2: 77}
                _ = fi55
                _ = v119
                _ = v56
            }
            v1 += -94
            if (385 != 0) {
                v120 := []int{fi48}
                if ((v4 / 1.0) == -2.25) {
                    for fi121 := 0; fi121 < 20; fi121++ {
                        if (42 != -1) {
                            delete(v3, 138)
                            var v122 uint = 65535
                            if len(v120) > 0 {
                                v120[0] = (^(fi121))
                            }
                            switch (fi48 & (-(fi121))) {
                            case 9223372036854775807:
                                v5 = v5
                            case 128:
                                v5 -= v5
                            case 544:
                                switch ((v2 + (("hello" + v2) + "a")) + v2) {
                                case "hello":
                                    switch "hello" {
                                    case "":
                                        v123 := GenS1{F0: (^(fi48)), F1: -829, F2: 2}
                                        _ = v123
                                    case "hello":
                                        var v124 uint16 = (v5 & v5)
                                        _ = v124
                                    case "foo bar":
                                        switch v2 {
                                        case "a":
                                            var v125 float64 = 3.14159
                                            _ = v125
                                        }
                                    }
                                case "":
                                    switch (("a" + v2) + v2) {
                                    case "foo bar":
                                        if (!false) {
                                            if ((!false) && false) {
                                                var v126 uint = 1
                                                if false {
                                                    v1 += -40
                                                } else {
                                                    var v127 bool = (true && (!(-693 > 968)))
                                                    _ = v127
                                                }
                                                _ = v126
                                            } else {
                                                v128 := GenS4{F0: false, F1: (^(fi121))}
                                                v129 := GenS17{F0: 0, F1: -1, F2: -1}
                                                _ = v128
                                                _ = v129
                                            }
                                        } else {
                                            var v130 bool = true
                                            var v131 uint8 = 127
                                            switch "hello" {
                                            case "hello":
                                                v132 := GenS18{F0: 1, F1: 899, F2: (42 == 1)}
                                                _ = v132
                                            }
                                            v2 = "\n\t"
                                            _ = v130
                                            _ = v131
                                        }
                                    default:
                                        v2 = (((v2 + v2) + v2) + ((v2 + (((v2 + v2) + "a") + "hello")) + ""))
                                    }
                                }
                            default:
                                if true {
                                    v2 = (v2 + "hello")
                                    v133 := GenS13{F0: (fi48 * (fi121 ^ (fi48 << ((42) & 63)))), F1: v122}
                                    v134 := GenS19{F0: -2147483648, F1: (fi48 >= (-(fi121)))}
                                    _ = v133
                                    _ = v134
                                } else {
                                    v122 = v122
                                }
                            }
                            var v135 uint64 = 18446744073709551615
                            _ = v122
                            _ = v135
                        }
                        var v136 uint32 = 2147483648
                        _ = fi121
                        _ = v136
                    }
                    for fi137 := 0; fi137 < 20; fi137++ {
                        if ((((3.14159 - ((1.5 * -2.25) * 0.0)) / -2.25) / (3.14159 / 1e10)) > 0.1) {
                            v138 := GenS18{F0: -1, F1: 762, F2: (v1 < v1)}
                            delete(v3, 42)
                            _ = v138
                        }
                        _ = fi137
                    }
                    switch 127 {
                    case 32768:
                        var v139 float64 = 3.14159
                        _ = v139
                    default:
                        for fi140 := 0; fi140 < 7; fi140++ {
                            v141 := GenS20{F0: 2, F1: 42, F2: 18446744073709551615}
                            fmt.Println(fi140)
                            _ = len(v3)
                            _ = fi140
                            _ = v141
                        }
                    }
                } else {
                    fmt.Println(v2)
                }
                var v142 string = ("" + v2)
                _ = v120
                _ = v142
            } else {
                var v143 int8 = v49
                v144 := GenS20{F0: 4294967295, F1: 65535, F2: 255}
                if (uint32(v5) == 1) {
                    switch (v2 + v2) {
                    case "hello":
                        fmt.Println(v5, v5)
                    }
                    v143 = v143
                    fmt.Println(v143, fi48, v49)
                } else {
                    if (!((!false) && (0 >= 871))) {
                        v145 := []int{fi48, fi48}
                        _ = len(v145)
                        fmt.Println(fi48, v2)
                        _ = v145
                    } else {
                        v2 = v2
                        var v146 int8 = 0
                        _ = v146
                    }
                    var v147 int = 128
                    _ = v147
                }
                _ = v143
                _ = v144
            }
        }
        var v148 float64 = 0.0
        _ = fi48
        _ = v148
        _ = v49
    }
    if true {
        v5 -= (v5 & (v5 >> ((1) & 63)))
        delete(v3, 2147483647)
        if true {
            switch 128 {
            case 4294967295:
                v4 += float32(0.0)
            case -2147483648:
                delete(v3, 4294967295)
            case 329:
                var v149 int8 = v1
                _ = v149
            default:
                v150 := GenS21{F0: 743, F1: -2147483648}
                _ = v150
            }
            v151 := GenS9{F0: -2147483648, F1: 439}
            _ = v151
        }
    } else {
        if (4294967295 > 4294967295) {
            v152 := GenS5{F0: 845}
            _ = v152
        }
    }
    v1 = 127
    if ((true && true) || false) {
        fmt.Println(v2, v1)
        v4 = 0.0
    }
    var v153 uint = 2
    for fi154 := 0; fi154 < 13; fi154++ {
        var v155 uint64 = 4294967295
        v155 = uint64(v5)
        _ = fi154
        _ = v155
    }
    for fi156 := 0; fi156 < 10; fi156++ {
        if true {
            v157 := v3[1]
            _ = v157
        } else {
            var v158 float64 = float64(0.1)
            v159 := []int{(fi156 ^ (-(fi156))), (fi156 / (((fi156 + 128)) | 1)), fi156}
            _ = v158
            _ = v159
        }
        switch (fi156 | (fi156 / ((int(v5)) | 1))) {
        case 42:
            v160 := []int{40, -375}
            _ = v160
        case 477:
            switch fi156 {
            case -1:
                var v161 int64 = -1
                _ = v161
            default:
                var v162 uint = (^(v153))
                _ = v162
            }
        case -209:
            fmt.Println(v153, fi156, v5)
        }
        _ = fi156
    }
    v5 += v5
    if ((!((3.14159 / (-2.25 * float64(float32(-2.25)))) != float64(1.5))) || true) {
        var v163 int8 = v1
        v163 = v1
        _ = v163
    } else {
        fmt.Println(v153)
    }
    var v164 uint64 = 18446744073709551615
    if (float32(0.1) < v4) {
        v153 = 18446744073709551615
        if false {
            var v165 string = "\n\t"
            var v166 float64 = (-2.25 * 0.1)
            _ = v165
            _ = v166
        }
        v167 := v3[2147483647]
        _ = v167
    }
    v2 = "hello"
    v153 -= 65535
    fmt.Println("v1:", v1)
    fmt.Println("v153:", v153)
    fmt.Println("v164:", v164)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v153
    _ = v164
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
