package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int16 }
type GenS2 struct { F0 int; F1 bool; F2 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint; F1 uint; F2 uint8 }
type GenS7 struct { F0 uint16; F1 int16; F2 int32 }
type GenS8 struct { F0 string; F1 uint16; F2 string }
type GenS9 struct { F0 bool; F1 int8 }
type GenS10 struct { F0 int; F1 uint8 }
type GenS11 struct { F0 uint16; F1 uint }
type GenS12 struct { F0 uint32; F1 int16 }
type GenS13 struct { F0 uint32; F1 int; F2 uint16 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]bool)
    v3 := make(map[uint32]uint32)
    var v4 uint64 = 971
    switch "hello" {
    case "\n\t":
        v3[2] = 0
    case "":
        fmt.Println(v4)
    case "hello":
        switch ("hello" + ("\n\t" + "foo bar")) {
        case "":
            var v5 uint8 = 0
            _ = v5
        case "hello":
            for fi6 := 0; fi6 < 18; fi6++ {
                v4 -= (v4 * 18446744073709551615)
                if false {
                    v7 := GenS1{F0: false, F1: -2147483648, F2: -937}
                    _ = v7
                } else {
                    switch "hello" {
                    case "":
                        if true {
                            v8 := GenS2{F0: fi6, F1: true, F2: (766 > v4)}
                            _ = v8
                        }
                    case "hello":
                        if (false && false) {
                            v9 := []int{}
                            v10 := GenS2{F0: -9223372036854775808, F1: (false || (("foo bar" + ((("a" + "") + "foo bar") + "\n\t")) <= "")), F2: false}
                            _ = v10
                            _ = v9
                        } else {
                            _ = len(v2)
                            v4 = v4
                        }
                    default:
                        v11 := GenS2{F0: (fi6 % ((fi6) | 1)), F1: false, F2: (1e10 < 0.1)}
                        _ = v11
                    }
                    fmt.Println(fi6)
                    switch fi6 {
                    case 32767:
                        if (!(950 < -162)) {
                            if (true && true) {
                                var v12 int16 = 1
                                if true {
                                    fmt.Println(fi6, v4, v4)
                                    if true {
                                        var v13 uint64 = (^(v4))
                                        _ = v13
                                    }
                                } else {
                                    var v14 uint32 = 0
                                    for fi15 := 0; fi15 < 8; fi15++ {
                                        switch 127 {
                                        case 128:
                                            fmt.Println(fi15, fi15)
                                        }
                                        if false {
                                            if (101 < 42) {
                                                var v16 int16 = v12
                                                _ = v16
                                            }
                                            if true {
                                                var v17 uint64 = v4
                                                v4 -= v4
                                                if false {
                                                    v18 := GenS1{F0: true, F1: 216, F2: 239}
                                                    if (true && ((false || ((-60 != -128) && false)) && (float64(0.0) >= 3.14159))) {
                                                        if (0.1 >= 0.1) {
                                                            if true {
                                                                if true {
                                                                    if false {
                                                                        v4 = uint64(v14)
                                                                    }
                                                                    v19 := GenS1{F0: (1 == 18446744073709551615), F1: -805, F2: v12}
                                                                    v14 += (v14 - (v14 | 2))
                                                                    _ = v19
                                                                } else {
                                                                    v20 := GenS1{F0: true, F1: -200, F2: 42}
                                                                    _ = v20
                                                                }
                                                                if false {
                                                                    if false {
                                                                        fmt.Println(fi15)
                                                                        v21 := GenS1{F0: (int(v14) >= (fi15 + (fi6 % ((172) | 1)))), F1: -1, F2: (v12 + (v12 / ((32767) | 1)))}
                                                                        v12 = v12
                                                                        _ = v21
                                                                    }
                                                                    v14 = 524
                                                                }
                                                            }
                                                            v3[2147483648] = 65535
                                                            v22 := []int{-424, fi15, (fi15 | (fi15 >> ((44) & 63))), -1}
                                                            _ = v22
                                                        } else {
                                                            if false {
                                                                v23 := GenS3{F0: 42}
                                                                v14 = v14
                                                                v4 += (v4 & v4)
                                                                _ = v23
                                                            }
                                                        }
                                                    } else {
                                                        v24 := GenS3{F0: 1}
                                                        v14 = 1
                                                        v12 = 42
                                                        _ = v24
                                                    }
                                                    v17 = (v4 + (v4 * (v17 | 618)))
                                                    if false {
                                                        var v25 int32 = 42
                                                        var v26 int32 = (v25 / ((-637) | 1))
                                                        _ = v25
                                                        _ = v26
                                                    } else {
                                                        var v27 int8 = 100
                                                        _ = v27
                                                    }
                                                    _ = v18
                                                }
                                                var v28 uint16 = 827
                                                if true {
                                                    if false {
                                                        delete(v2, 1)
                                                    } else {
                                                        var v29 uint8 = 0
                                                        if false {
                                                            v14 -= v14
                                                            v30 := v2[(fi6 * (fi6 & fi15))]
                                                            v17 = (v17 - (v4 >> ((0) & 63)))
                                                            _ = v30
                                                        }
                                                        switch fi6 {
                                                        case 0:
                                                            v4 = (v4 ^ v4)
                                                        case 127:
                                                            v31 := []int{263}
                                                            _ = v31
                                                        }
                                                        _ = v29
                                                    }
                                                    var v32 int64 = 9223372036854775807
                                                    if true {
                                                        var v33 string = "\n\t"
                                                        v17 = v17
                                                        _ = v33
                                                    } else {
                                                        v2[fi6] = true
                                                        _ = len(v2)
                                                    }
                                                    _ = v32
                                                }
                                                _ = v17
                                                _ = v28
                                            } else {
                                                v34 := GenS3{F0: 0}
                                                v12 -= (v12 / ((int16(v4)) | 1))
                                                _ = v34
                                            }
                                            var v35 int32 = 182
                                            _ = v35
                                        }
                                        _ = fi15
                                    }
                                    _ = v14
                                }
                                if (-1 != 17) {
                                    v4 = 2
                                    if (67 >= 18446744073709551615) {
                                        fmt.Println(v4, fi6, fi6)
                                    } else {
                                        for fi36 := 0; fi36 < 17; fi36++ {
                                            switch "a" {
                                            case "foo bar":
                                                fmt.Println(fi36, fi36)
                                            }
                                            var v37 int = fi36
                                            v12 += 32767
                                            switch fi36 {
                                            case 32767:
                                                v12 += 32767
                                            default:
                                                delete(v2, (fi36 << ((18446744073709551615) & 63)))
                                            }
                                            _ = fi36
                                            _ = v37
                                        }
                                        if (127 == 75) {
                                            if (-9223372036854775808 > -1) {
                                                delete(v2, fi6)
                                                var v38 int64 = -9223372036854775808
                                                _ = v38
                                            } else {
                                                switch 902 {
                                                case 783:
                                                    v39 := []int{}
                                                    _ = v39
                                                case 4294967295:
                                                    v40 := GenS2{F0: 4294967295, F1: true, F2: false}
                                                    _ = v40
                                                case -1:
                                                    var v41 int = 9223372036854775807
                                                    _ = v41
                                                }
                                                v42 := GenS2{F0: 32767, F1: false, F2: ((float64(3.14159) > 3.14159) || true)}
                                                if false {
                                                    v43 := []int{32767, (fi6 >> ((106) & 63)), (fi6 - (fi6 * 1))}
                                                    _ = v43
                                                } else {
                                                    fmt.Println(v12)
                                                    if true {
                                                        switch ("a" + ("" + (("foo bar" + (("" + "a") + "foo bar")) + ("hello" + "a")))) {
                                                        case "hello":
                                                            _ = len(v1)
                                                        }
                                                        if true {
                                                            var v44 bool = true
                                                            v12 -= (v12 ^ (v12 + (v12 ^ (v12 ^ v12))))
                                                            _ = v44
                                                        } else {
                                                            v45 := []int{9223372036854775807}
                                                            _ = v45
                                                        }
                                                    } else {
                                                        v4 = v4
                                                    }
                                                    if ((v12 + v12) > v12) {
                                                        if false {
                                                            if true {
                                                                v12 = (v12 * 42)
                                                                var v46 float32 = ((-2.25 + 1e10) * (0.1 + 0.1))
                                                                fmt.Println(fi6, v46)
                                                                _ = v46
                                                            }
                                                        } else {
                                                            var v47 int = (fi6 ^ (fi6 * (-(fi6))))
                                                            v4 += v4
                                                            _ = v47
                                                        }
                                                        for fi48 := 0; fi48 < 4; fi48++ {
                                                            v12 = v12
                                                            v4 += 4294967295
                                                            v49 := GenS4{F0: 42, F1: "hello"}
                                                            _ = fi48
                                                            _ = v49
                                                        }
                                                    }
                                                }
                                                _ = v42
                                            }
                                            if (!true) {
                                                v12 = (-(v12))
                                                for fi50 := 0; fi50 < 6; fi50++ {
                                                    var v51 string = ""
                                                    fmt.Println(fi50)
                                                    var v52 int64 = 0
                                                    _ = fi50
                                                    _ = v51
                                                    _ = v52
                                                }
                                                var v53 float32 = (1.5 - 3.14159)
                                                _ = v53
                                            }
                                        } else {
                                            if (0 <= fi6) {
                                                v54 := []int{(fi6 | (fi6 << ((4294967295) & 63)))}
                                                _ = v54
                                            } else {
                                                for fi55 := 0; fi55 < 13; fi55++ {
                                                    switch (fi6 << ((uint(v4)) & 63)) {
                                                    case 128:
                                                        v56 := []int{}
                                                        _ = v56
                                                    case -2147483648:
                                                        var v57 uint16 = 65535
                                                        _ = v57
                                                    case 42:
                                                        v58 := GenS4{F0: 42, F1: "hello"}
                                                        _ = v58
                                                    default:
                                                        var v59 int32 = -2147483648
                                                        _ = v59
                                                    }
                                                    v60 := GenS3{F0: 522}
                                                    _ = fi55
                                                    _ = v60
                                                }
                                                var v61 float32 = 1.5
                                                _ = v61
                                            }
                                        }
                                    }
                                    switch fi6 {
                                    case 1:
                                        v4 += 912
                                    case 4294967295:
                                        var v62 float32 = (0.0 - -2.25)
                                        _ = v62
                                    case 128:
                                        var v63 uint = 255
                                        _ = v63
                                    }
                                }
                                _ = v12
                            } else {
                                v4 = v4
                            }
                            var v64 int = (-(fi6))
                            v4 -= (v4 ^ (v4 * 190))
                            _ = v64
                        }
                    }
                }
                var v65 uint = 42
                _ = fi6
                _ = v65
            }
        default:
            if false {
                if true {
                    for fi66 := 0; fi66 < 19; fi66++ {
                        v4 = (v4 - (^(v4)))
                        for fi67 := 0; fi67 < 3; fi67++ {
                            fmt.Println(fi66, v4)
                            v68 := GenS4{F0: -454, F1: ("" + "")}
                            v4 = 1
                            _ = fi67
                            _ = v68
                        }
                        v4 += v4
                        _ = fi66
                    }
                    switch 0 {
                    case 9223372036854775807:
                        v69 := GenS2{F0: 32767, F1: (133 == 4294967295), F2: false}
                        _ = v69
                    }
                } else {
                    v70 := GenS3{F0: 42}
                    switch 4294967295 {
                    case -9223372036854775808:
                        var v71 int64 = 1
                        _ = v71
                    case -2147483648:
                        var v72 int64 = 165
                        _ = v72
                    case 0:
                        if ((!((1e10 / 1.0) == 3.14159)) && ((v4 | 42) == v4)) {
                            for fi73 := 0; fi73 < 13; fi73++ {
                                v4 += (v4 % (((v4 + 2)) | 1))
                                _ = len(v2)
                                for fi74 := 0; fi74 < 11; fi74++ {
                                    var v75 float64 = -2.25
                                    delete(v2, (fi73 - 9223372036854775807))
                                    v75 = 3.14159
                                    _ = fi74
                                    _ = v75
                                }
                                _ = fi73
                            }
                        } else {
                            if (((0.0 / -2.25) < 3.14159) && true) {
                                var v76 int16 = 1000
                                v77 := []int{127}
                                if true {
                                    if false {
                                        if (258 != -45) {
                                            var v78 int32 = 42
                                            _ = v78
                                        } else {
                                            var v79 int32 = -127
                                            for fi80 := 0; fi80 < 20; fi80++ {
                                                fmt.Println(v76, v4, fi80)
                                                v81 := GenS2{F0: (fi80 * (fi80 * 1)), F1: ((false && true) || true), F2: ((3.14159 + ((0.0 / float64((1.5 + 0.0))) / -2.25)) == -2.25)}
                                                v82 := GenS5{F0: 2}
                                                _ = fi80
                                                _ = v81
                                                _ = v82
                                            }
                                            _ = v79
                                        }
                                    }
                                } else {
                                    v77 = append(v77, 32767)
                                    var v83 string = "foo bar"
                                    _ = v83
                                }
                                _ = v76
                                _ = v77
                            } else {
                                v84 := GenS1{F0: ((-2.25 / 1.0) != ((0.1 - float64(float32(0.1))) - -2.25)), F1: 59, F2: -1}
                                fmt.Println(v4)
                                if (!(567 > 2147483647)) {
                                    switch "\n\t" {
                                    case "":
                                        if (2147483648 > 2147483647) {
                                            for fi85 := 0; fi85 < 17; fi85++ {
                                                for fi86 := 0; fi86 < 9; fi86++ {
                                                    if false {
                                                        v87 := []int{fi85}
                                                        _ = v87
                                                    } else {
                                                        _ = len(v3)
                                                        switch "" {
                                                        case "":
                                                            if true {
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    switch (fi86 | int(v4)) {
                                                                    case 9223372036854775807:
                                                                        switch fi86 {
                                                                        case 2147483647:
                                                                            fmt.Println(fi86, fi86)
                                                                        case 9223372036854775807:
                                                                            v4 -= v4
                                                                        case -930:
                                                                            v4 -= (v4 << ((0) & 63))
                                                                        default:
                                                                            var v88 uint32 = 2147483648
                                                                            _ = v88
                                                                        }
                                                                    default:
                                                                        if true {
                                                                            fmt.Println(fi85, v4, fi85)
                                                                            v89 := []int{754, fi86}
                                                                            _ = v89
                                                                        } else {
                                                                            var v90 int32 = 571
                                                                            v90 += (v90 % (((v90 >> ((1) & 63))) | 1))
                                                                            var v91 uint16 = 255
                                                                            _ = v90
                                                                            _ = v91
                                                                        }
                                                                    }
                                                                default:
                                                                    v4 = v4
                                                                }
                                                                var v92 float32 = 1e10
                                                                _ = v92
                                                            } else {
                                                                v93 := GenS3{F0: -2147483648}
                                                                fmt.Println(v4, v4, fi85)
                                                                _ = v93
                                                            }
                                                        }
                                                    }
                                                    _ = fi86
                                                }
                                                _ = fi85
                                            }
                                        } else {
                                            v94 := GenS4{F0: -77, F1: "hello"}
                                            var v95 int16 = 889
                                            _ = v94
                                            _ = v95
                                        }
                                    }
                                    v4 = v4
                                } else {
                                    var v96 int64 = -9223372036854775808
                                    v96 = -9223372036854775808
                                    v96 = -1
                                    _ = v96
                                }
                                _ = v84
                            }
                        }
                    default:
                        if false {
                            fmt.Println(v4)
                            v97 := GenS6{F0: 794, F1: 4294967295, F2: 2}
                            _ = v97
                        } else {
                            if true {
                                for fi98 := 0; fi98 < 4; fi98++ {
                                    v2[fi98] = false
                                    _ = fi98
                                }
                                for fi99 := 0; fi99 < 18; fi99++ {
                                    if ((fi99 + -464) != fi99) {
                                        _ = len(v3)
                                        v100 := GenS6{F0: 0, F1: 2, F2: 2}
                                        v101 := GenS7{F0: 1, F1: 42, F2: 0}
                                        _ = v100
                                        _ = v101
                                    } else {
                                        delete(v3, 2147483648)
                                        if (!(!(0.0 < (0.1 - float32(0.1))))) {
                                            for fi102 := 0; fi102 < 1; fi102++ {
                                                v103 := []int{2147483647, fi99}
                                                if (!(32767 > 42)) {
                                                    v2[fi102] = false
                                                    if true {
                                                        v4 -= 0
                                                    }
                                                } else {
                                                    var v104 bool = ("a" <= "hello")
                                                    _ = v104
                                                }
                                                v4 = (v4 / (((^(v4))) | 1))
                                                _ = fi102
                                                _ = v103
                                            }
                                        } else {
                                            v105 := GenS8{F0: ("hello" + "\n\t"), F1: 65535, F2: "a"}
                                            _ = v105
                                        }
                                        v106 := []int{}
                                        _ = v106
                                    }
                                    var v107 int = fi99
                                    v108 := GenS7{F0: 42, F1: 1000, F2: -2147483648}
                                    _ = fi99
                                    _ = v107
                                    _ = v108
                                }
                                v4 = (v4 & 391)
                            }
                        }
                    }
                    v4 += (v4 ^ v4)
                    _ = v70
                }
                fmt.Println(v4)
            }
        }
    default:
        var v109 int64 = 1
        _ = v109
    }
    v110 := GenS7{F0: 625, F1: -395, F2: 0}
    for fi111 := 0; fi111 < 11; fi111++ {
        for fi112 := 0; fi112 < 6; fi112++ {
            switch (fi111 + fi111) {
            case 127:
                v4 = (v4 / ((v4) | 1))
            }
            _ = fi112
        }
        fmt.Println(v4, fi111)
        switch (fi111 + (fi111 & (fi111 & (fi111 ^ fi111)))) {
        case 32767:
            v4 += v4
        }
        _ = fi111
    }
    v113 := GenS3{F0: 1}
    v4 += (v4 | v4)
    v4 += 42
    var v114 uint32 = 4294967295
    fmt.Println(v114)
    var v115 uint64 = v4
    switch -9223372036854775808 {
    case 2147483647:
        v116 := v2[-626]
        _ = v116
    case 0:
        fmt.Println(v4, v115)
    case 32767:
        switch "" {
        case "foo bar":
            var v117 int8 = 102
            _ = v117
        case "\n\t":
            switch "\n\t" {
            case "a":
                var v118 string = "foo bar"
                _ = v118
            case "foo bar":
                for fi119 := 0; fi119 < 3; fi119++ {
                    for fi120 := 0; fi120 < 11; fi120++ {
                        var v121 uint64 = (v4 | v115)
                        v122 := GenS4{F0: 220, F1: "foo bar"}
                        _ = fi120
                        _ = v121
                        _ = v122
                    }
                    v123 := []int{32768, -1, 342}
                    _ = fi119
                    _ = v123
                }
            case "hello":
                var v124 int64 = -284
                _ = v124
            default:
                v115 = v4
            }
        default:
            v125 := GenS9{F0: (!(2 > 127)), F1: -1}
            _ = v125
        }
    }
    for fi126 := 0; fi126 < 19; fi126++ {
        fmt.Println(v114)
        var v127 string = "a"
        if false {
            fmt.Println(v4, v127, v115)
        }
        _ = fi126
        _ = v127
    }
    var v128 float32 = (1.5 + 3.14159)
    for fi129 := 0; fi129 < 18; fi129++ {
        _ = len(v3)
        _ = fi129
    }
    var v130 int64 = 9223372036854775807
    for fi131 := 0; fi131 < 9; fi131++ {
        v132 := GenS9{F0: false, F1: 1}
        switch "hello" {
        case "hello":
            v133 := GenS7{F0: 849, F1: -32768, F2: int32(v115)}
            _ = v133
        case "a":
            fmt.Println(v130, v128)
        case "\n\t":
            v128 -= (v128 + (v128 / v128))
        default:
            if (("hello" != "a") || true) {
                v130 = (v130 >> ((255) & 63))
            }
        }
        _ = fi131
        _ = v132
    }
    if (1 >= (v114 * (^(v114)))) {
        switch -9223372036854775808 {
        case -9223372036854775808:
            fmt.Println(v114, v130)
        case 32767:
            for fi134 := 0; fi134 < 1; fi134++ {
                v135 := GenS10{F0: 1, F1: 118}
                v115 += v115
                _ = fi134
                _ = v135
            }
        case 127:
            if true {
                if true {
                    var v136 uint8 = 42
                    if (!(true || (true && (232 <= -32768)))) {
                        for fi137 := 0; fi137 < 20; fi137++ {
                            switch fi137 {
                            case -41:
                                v138 := GenS11{F0: 533, F1: 0}
                                _ = v138
                            default:
                                for fi139 := 0; fi139 < 13; fi139++ {
                                    var v140 uint = 4294967295
                                    _ = fi139
                                    _ = v140
                                }
                            }
                            fmt.Println(v114, v115)
                            v130 = -1
                            _ = fi137
                        }
                        if false {
                            v136 = v136
                        } else {
                            v141 := []int{42, 2147483647, -1, 127}
                            if false {
                                switch ("\n\t" + "a") {
                                case "hello":
                                    v142 := GenS11{F0: 1, F1: uint(v115)}
                                    _ = v142
                                }
                                var v143 int8 = 6
                                switch "\n\t" {
                                case "foo bar":
                                    switch 4294967295 {
                                    case 740:
                                        v144 := GenS12{F0: 46, F1: -509}
                                        _ = v144
                                    case 127:
                                        var v145 int = 117
                                        _ = v145
                                    case 9223372036854775807:
                                        var v146 float32 = float32((3.14159 - float64(v128)))
                                        _ = v146
                                    default:
                                        v130 = -1
                                    }
                                case "\n\t":
                                    v115 = v115
                                case "":
                                    switch 1 {
                                    case -1:
                                        v130 = 585
                                    default:
                                        _ = len(v141)
                                    }
                                }
                                _ = v143
                            } else {
                                switch 42 {
                                case -673:
                                    for fi147 := 0; fi147 < 13; fi147++ {
                                        v148 := GenS9{F0: false, F1: 42}
                                        var v149 int64 = (v130 * v130)
                                        _ = fi147
                                        _ = v148
                                        _ = v149
                                    }
                                }
                                switch (("foo bar" + ("" + "foo bar")) + "hello") {
                                case "a":
                                    var v150 int8 = -33
                                    _ = v150
                                case "\n\t":
                                    var v151 bool = false
                                    _ = v151
                                case "foo bar":
                                    v152 := GenS10{F0: -1, F1: v136}
                                    _ = v152
                                }
                            }
                            v114 = 42
                            _ = v141
                        }
                    }
                    _ = v136
                } else {
                    for fi153 := 0; fi153 < 8; fi153++ {
                        if true {
                            v154 := GenS11{F0: uint16(v4), F1: 1}
                            _ = v154
                        }
                        v155 := []int{127}
                        fmt.Println(v4)
                        _ = fi153
                        _ = v155
                    }
                    v156 := GenS13{F0: v114, F1: -630, F2: 2}
                    _ = v156
                }
                for fi157 := 0; fi157 < 10; fi157++ {
                    var v158 int16 = -260
                    _ = fi157
                    _ = v158
                }
            }
        }
    }
    var v159 int16 = 15
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v114:", v114)
    fmt.Println("v115:", v115)
    fmt.Println("v128:", v128)
    fmt.Println("v130:", v130)
    fmt.Println("v159:", v159)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v110
    _ = v113
    _ = v114
    _ = v115
    _ = v128
    _ = v130
    _ = v159
    _ = v2
    _ = v3
    _ = v4
}
