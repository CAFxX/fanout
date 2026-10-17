package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint; F1 bool; F2 bool }
type GenS3 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 int64; F1 uint }
type GenS7 struct { F0 int16; F1 int32; F2 uint }
type GenS8 struct { F0 int32; F1 bool }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int16; F1 int32 }
type GenS11 struct { F0 uint32; F1 uint64 }
type GenS12 struct { F0 int32; F1 uint32; F2 int32 }
type GenS13 struct { F0 bool; F1 uint64; F2 int }
func genFn0(p0 uint) uint16 {
    fmt.Println(p0)
    p0 += 0
    var v6 int16 = 474
    var v7 int64 = 749
    _ = p0
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 float32 = 0.1
    var v2 int16 = 253
    var v3 bool = true
    var v4 bool = (true || false)
    var v5 int16 = (v2 ^ -32768)
    switch "hello" {
    case "foo bar":
        var v8 int32 = 0
        _ = v8
    case "":
        v9 := []int{-235, -9223372036854775808}
        _ = v9
    case "\n\t":
        if (-1 >= -1) {
            for fi10 := 0; fi10 < 17; fi10++ {
                switch (fi10 % ((fi10) | 1)) {
                case 2147483647:
                    v11 := GenS1{F0: 369, F1: 0}
                    _ = v11
                }
                _ = fi10
            }
        } else {
            var v12 int8 = -121
            if (v2 >= 0) {
                v2 = v5
            }
            if v4 {
                v13 := []int{988, -2147483648, -9223372036854775808, 91}
                _ = v13
            } else {
                fmt.Println(v1)
            }
            _ = v12
        }
    }
    if (-58 <= -1) {
        for fi14 := 0; fi14 < 3; fi14++ {
            switch int(v2) {
            case 2147483647:
                var v15 int64 = -9223372036854775808
                _ = v15
            case -9223372036854775808:
                v5 += 1000
            default:
                v2 = v5
            }
            v16 := []int{(fi14 - (fi14 >> ((269) & 63))), 594}
            v3 = (0 > -1)
            _ = fi14
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 10; fi17++ {
            for fi18 := 0; fi18 < 14; fi18++ {
                v19 := GenS1{F0: 2147483647, F1: -127}
                v20 := []int{fi18, (fi18 + (fi18 + fi18)), (fi17 + fi17)}
                _ = len(v20)
                v21 := GenS1{F0: 899, F1: 1}
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            _ = fi17
        }
        if v3 {
            fmt.Println(v3, v1)
        }
        v22 := []int{4294967295, 32768}
        _ = v22
    }
    var v23 int32 = 1
    if (42 != 803) {
        if ((1.5 != float64((v1 / v1))) && (255 == 4294967295)) {
            v3 = ((0.1 / -2.25) <= 3.14159)
            v24 := GenS2{F0: 313, F1: (((-2.25 - float32(0.1)) > (float32(1.5) + v1)) || (v3 || v3)), F2: (2 <= 255)}
            if v3 {
                if (129 != 0) {
                    var v25 int32 = -1
                    v26 := []int{9223372036854775807, 2147483647}
                    _ = v25
                    _ = v26
                } else {
                    v27 := []int{1, -9223372036854775808}
                    _ = v27
                }
                v28 := []int{2147483647, -490, 0}
                _ = v28
            } else {
                v29 := GenS3{F0: 0, F1: 882, F2: 32767}
                _ = v29
            }
            _ = v24
        } else {
            if v3 {
                v5 -= -32768
                fmt.Println(v1)
                switch ((("a" + "hello") + "a") + ("a" + "foo bar")) {
                case "a":
                    v4 = (856 > 2147483648)
                }
            } else {
                if (true && (v4 && true)) {
                    v30 := GenS2{F0: 18446744073709551615, F1: false, F2: (v3 || true)}
                    if ((v23 << ((4294967295) & 63)) <= v23) {
                        for fi31 := 0; fi31 < 4; fi31++ {
                            if v4 {
                                for fi32 := 0; fi32 < 18; fi32++ {
                                    v33 := GenS1{F0: 11, F1: -128}
                                    switch "hello" {
                                    case "foo bar":
                                        switch fi31 {
                                        case 9223372036854775807:
                                            if false {
                                                v34 := GenS3{F0: 314, F1: 666, F2: 242}
                                                _ = v34
                                            } else {
                                                if v4 {
                                                    v35 := GenS4{F0: 354}
                                                    _ = v35
                                                } else {
                                                    v23 -= 42
                                                    v23 -= int32(fi32)
                                                }
                                                v36 := []int{(fi32 ^ fi32), (fi31 / (((fi31 >> ((65535) & 63))) | 1)), fi31}
                                                v23 = v23
                                                _ = v36
                                            }
                                        case -9223372036854775808:
                                            v2 -= v5
                                        case -1:
                                            switch "" {
                                            case "a":
                                                switch (fi31 ^ (fi31 * 9223372036854775807)) {
                                                case -984:
                                                    if (210 < 18446744073709551615) {
                                                        var v37 uint64 = 42
                                                        v2 = v2
                                                        if (float64(((v1 + ((v1 * (v1 / (0.1 + float32(3.14159)))) / 0.1)) * float32((0.0 + float64(-2.25))))) > 0.0) {
                                                            if (32767 == genFn0(18446744073709551615)) {
                                                                v38 := GenS2{F0: 65535, F1: (!false), F2: (v3 || v3)}
                                                                switch fi32 {
                                                                case -89:
                                                                    if v4 {
                                                                        v3 = (!v3)
                                                                    } else {
                                                                        v37 -= v37
                                                                        v39 := []int{(-(fi31)), (fi31 | 32767)}
                                                                        _ = v39
                                                                    }
                                                                case -224:
                                                                    v1 = v1
                                                                case 291:
                                                                    v37 = 2
                                                                default:
                                                                    var v40 bool = v4
                                                                    _ = v40
                                                                }
                                                                v41 := GenS3{F0: 1, F1: 2147483648, F2: 2}
                                                                _ = v38
                                                                _ = v41
                                                            }
                                                            switch ("\n\t" + "foo bar") {
                                                            case "foo bar":
                                                                v42 := []int{}
                                                                _ = v42
                                                            }
                                                            fmt.Println(v3)
                                                        }
                                                        _ = v37
                                                    } else {
                                                        fmt.Println(v1, v2)
                                                    }
                                                case 42:
                                                    if ((v23 <= v23) && (4294967295 == 4294967295)) {
                                                        var v43 uint64 = 663
                                                        v44 := []int{(fi31 * fi32), 9223372036854775807}
                                                        v43 += v43
                                                        _ = v43
                                                        _ = v44
                                                    } else {
                                                        var v45 int16 = v2
                                                        _ = v45
                                                    }
                                                default:
                                                    var v46 uint16 = 0
                                                    _ = v46
                                                }
                                            case "\n\t":
                                                var v47 int32 = 1
                                                _ = v47
                                            case "foo bar":
                                                switch (("hello" + ("" + (("hello" + "\n\t") + ""))) + "") {
                                                case "\n\t":
                                                    fmt.Println(v23, v4)
                                                    var v48 uint16 = 42
                                                    v23 = v23
                                                    _ = v48
                                                case "a":
                                                    switch (fi31 << ((4294967295) & 63)) {
                                                    case -573:
                                                        var v49 bool = v3
                                                        _ = v49
                                                    case 4294967295:
                                                        v3 = (int16(fi31) == (v2 % ((v2) | 1)))
                                                    case -893:
                                                        v50 := GenS1{F0: 675, F1: -1}
                                                        _ = v50
                                                    }
                                                case "hello":
                                                    var v51 int = 4294967295
                                                    _ = v51
                                                }
                                            default:
                                                v23 += v23
                                            }
                                        }
                                    }
                                    v52 := []int{}
                                    _ = fi32
                                    _ = v33
                                    _ = v52
                                }
                                var v53 uint32 = 42
                                _ = v53
                            } else {
                                var v54 int32 = v23
                                v55 := []int{fi31}
                                switch fi31 {
                                case 32767:
                                    v3 = (!((v54 / (((v23 << ((2) & 63))) | 1)) <= -1))
                                case 2147483647:
                                    if v3 {
                                        if v4 {
                                            v55 = append(v55, (fi31 - fi31))
                                            switch "hello" {
                                            case "\n\t":
                                                var v56 int32 = 1
                                                _ = v56
                                            case "a":
                                                var v57 int8 = 1
                                                _ = v57
                                            case "foo bar":
                                                for fi58 := 0; fi58 < 14; fi58++ {
                                                    var v59 int = fi31
                                                    var v60 string = "\n\t"
                                                    _ = fi58
                                                    _ = v59
                                                    _ = v60
                                                }
                                            }
                                            _ = len(v55)
                                        } else {
                                            var v61 uint64 = 0
                                            if len(v55) > 0 {
                                                v55[0] = (-(fi31))
                                            }
                                            v62 := GenS1{F0: 65535, F1: 8}
                                            _ = v61
                                            _ = v62
                                        }
                                        switch 32767 {
                                        case -575:
                                            for fi63 := 0; fi63 < 14; fi63++ {
                                                if v4 {
                                                    if v3 {
                                                        v23 -= v54
                                                        v1 = v1
                                                    } else {
                                                        if v4 {
                                                            if (!((4294967295 > 42) || v4)) {
                                                                v64 := GenS5{F0: 65535, F1: 4294967295}
                                                                v2 += int16(fi31)
                                                                _ = v64
                                                            } else {
                                                                switch (-(fi31)) {
                                                                case 32767:
                                                                    if len(v55) >= 2 {
                                                                        v55 = v55[1:len(v55)]
                                                                    }
                                                                case -190:
                                                                    switch (fi31 % ((int(v23)) | 1)) {
                                                                    case 4294967295:
                                                                        if (true && false) {
                                                                            var v65 uint64 = 2
                                                                            _ = v65
                                                                        } else {
                                                                            _ = len(v55)
                                                                        }
                                                                    }
                                                                default:
                                                                    var v66 int16 = v2
                                                                    _ = v66
                                                                }
                                                                switch (fi31 ^ 1) {
                                                                case 350:
                                                                    var v67 uint16 = 65535
                                                                    _ = v67
                                                                case 127:
                                                                    if ("foo bar" <= ("a" + "")) {
                                                                        if v3 {
                                                                            var v68 float64 = -2.25
                                                                            _ = v68
                                                                        }
                                                                        v69 := GenS4{F0: 255}
                                                                        fmt.Println(fi31, fi31)
                                                                        _ = v69
                                                                    }
                                                                default:
                                                                    var v70 uint32 = 4294967295
                                                                    v71 := GenS6{F0: 1, F1: 847}
                                                                    v70 = v70
                                                                    _ = v70
                                                                    _ = v71
                                                                }
                                                            }
                                                            v2 -= (v5 >> ((0) & 63))
                                                            var v72 float32 = v1
                                                            _ = v72
                                                        }
                                                    }
                                                    v73 := GenS7{F0: (v2 >> ((1) & 63)), F1: v54, F2: 255}
                                                    _ = v73
                                                } else {
                                                    v74 := GenS7{F0: (v5 + (v2 % ((-32768) | 1))), F1: (-(v23)), F2: 1}
                                                    if (!false) {
                                                        var v75 float32 = (v1 - (v1 - v1))
                                                        _ = v75
                                                    }
                                                    _ = v74
                                                }
                                                v1 = (v1 * v1)
                                                _ = fi63
                                            }
                                        case 592:
                                            var v76 float32 = 0.1
                                            _ = v76
                                        case 590:
                                            if v4 {
                                                if v3 {
                                                    v77 := GenS8{F0: v23, F1: false}
                                                    _ = v77
                                                } else {
                                                    var v78 int32 = (v23 >> ((1) & 63))
                                                    v79 := GenS9{F0: 121}
                                                    _ = v78
                                                    _ = v79
                                                }
                                                var v80 int = fi31
                                                if (true && false) {
                                                    if v3 {
                                                        var v81 int8 = 127
                                                        for fi82 := 0; fi82 < 1; fi82++ {
                                                            switch 128 {
                                                            case -9223372036854775808:
                                                                var v83 float32 = (1e10 * ((v1 - (v1 / v1)) / -2.25))
                                                                _ = v83
                                                            case -488:
                                                                var v84 int64 = 42
                                                                _ = v84
                                                            case -189:
                                                                v5 = v2
                                                            }
                                                            _ = fi82
                                                        }
                                                        _ = v81
                                                    }
                                                    v3 = true
                                                }
                                                _ = v80
                                            } else {
                                                if (1 < 166) {
                                                    v54 = -1
                                                    switch 32767 {
                                                    case 2147483647:
                                                        if true {
                                                            if (144 > 996) {
                                                                fmt.Println(v1, v54, v54)
                                                                _ = len(v55)
                                                                for fi85 := 0; fi85 < 1; fi85++ {
                                                                    var v86 float32 = v1
                                                                    var v87 float32 = float32(0.0)
                                                                    _ = fi85
                                                                    _ = v86
                                                                    _ = v87
                                                                }
                                                            }
                                                            if v4 {
                                                                var v88 uint64 = 827
                                                                for fi89 := 0; fi89 < 11; fi89++ {
                                                                    v90 := GenS1{F0: 186, F1: -128}
                                                                    _ = fi89
                                                                    _ = v90
                                                                }
                                                                for fi91 := 0; fi91 < 2; fi91++ {
                                                                    if len(v55) > 0 {
                                                                        v55[0] = fi91
                                                                    }
                                                                    v92 := GenS6{F0: -114, F1: 338}
                                                                    _ = fi91
                                                                    _ = v92
                                                                }
                                                                _ = v88
                                                            }
                                                            var v93 uint64 = 18446744073709551615
                                                            _ = v93
                                                        }
                                                    case 1:
                                                        var v94 uint16 = 255
                                                        _ = v94
                                                    case 9223372036854775807:
                                                        fmt.Println(v23, v5, v3)
                                                    default:
                                                        if len(v55) > 0 {
                                                            v55[0] = 9223372036854775807
                                                        }
                                                    }
                                                }
                                                switch (^(fi31)) {
                                                case -9223372036854775808:
                                                    fmt.Println(v2, v23)
                                                case 32767:
                                                    v2 = 633
                                                case 32768:
                                                    v95 := GenS9{F0: 100}
                                                    _ = v95
                                                }
                                                var v96 uint64 = 1
                                                _ = v96
                                            }
                                        }
                                        v97 := GenS6{F0: 9223372036854775807, F1: 4294967295}
                                        _ = v97
                                    } else {
                                        v4 = true
                                        if true {
                                            var v98 uint16 = 32767
                                            for fi99 := 0; fi99 < 1; fi99++ {
                                                if false {
                                                    if ((796 < 61) && (2 < 1)) {
                                                        fmt.Println(v5)
                                                        v100 := GenS1{F0: 4294967295, F1: 1}
                                                        v98 -= 32767
                                                        _ = v100
                                                    }
                                                    v101 := GenS2{F0: 42, F1: v4, F2: (v1 <= (((v1 / ((v1 / (v1 - float32(((3.14159 - (1.5 * ((1e10 + ((1.5 - 3.14159) / -2.25)) - -2.25))) + 1e10)))) - float32(0.0))) / v1) - v1))}
                                                    v2 -= 3
                                                    _ = v101
                                                } else {
                                                    var v102 float32 = (v1 * 0.1)
                                                    v5 = -1
                                                    _ = v102
                                                }
                                                v5 = (v5 ^ v2)
                                                v103 := GenS3{F0: -9223372036854775808, F1: 245, F2: v98}
                                                _ = fi99
                                                _ = v103
                                            }
                                            if len(v55) > 0 {
                                                v55[0] = fi31
                                            }
                                            _ = v98
                                        }
                                        v5 -= v5
                                    }
                                case 4294967295:
                                    switch (fi31 - (-(fi31))) {
                                    case 0:
                                        var v104 string = (("\n\t" + ("foo bar" + ("a" + "hello"))) + "")
                                        _ = v104
                                    case 9223372036854775807:
                                        if (4294967295 == 2147483647) {
                                            v1 -= v1
                                            switch 32767 {
                                            case 4294967295:
                                                _ = len(v55)
                                            case -1:
                                                v105 := GenS6{F0: -470, F1: 238}
                                                _ = v105
                                            case -2147483648:
                                                var v106 float64 = ((0.1 / 3.14159) * ((0.0 - (float64(1e10) - 0.0)) - (-2.25 / (((float64((((float32(1.5) / v1) * -2.25) * v1)) + 1.5) * 1e10) * (3.14159 + (((0.1 / -2.25) + ((1.5 - -2.25) + ((((3.14159 * 1.5) / (float64((1e10 - ((v1 - 0.1) + v1))) + (1e10 + (3.14159 / ((float64(v1) / 0.1) * 0.0))))) / (0.1 * 0.1)) / ((0.1 / (1.5 - 3.14159)) - 1e10)))) / 1.5))))))
                                                _ = v106
                                            default:
                                                for fi107 := 0; fi107 < 3; fi107++ {
                                                    v55 = append(v55, -887)
                                                    v1 = v1
                                                    _ = fi107
                                                }
                                            }
                                            switch fi31 {
                                            case 91:
                                                v54 = v23
                                            default:
                                                var v108 uint32 = 296
                                                _ = v108
                                            }
                                        } else {
                                            if (334 <= 66) {
                                                _ = len(v55)
                                                v23 = v23
                                                for fi109 := 0; fi109 < 8; fi109++ {
                                                    var v110 bool = v4
                                                    switch fi31 {
                                                    case 42:
                                                        fmt.Println(v23)
                                                    case 775:
                                                        v23 = v23
                                                    }
                                                    var v111 uint16 = 1
                                                    _ = fi109
                                                    _ = v110
                                                    _ = v111
                                                }
                                            }
                                        }
                                    case 127:
                                        switch (fi31 >> ((435) & 63)) {
                                        case 2147483647:
                                            fmt.Println(v54)
                                        default:
                                            if true {
                                                for fi112 := 0; fi112 < 17; fi112++ {
                                                    _ = len(v55)
                                                    v113 := GenS10{F0: v5, F1: 42}
                                                    v3 = v3
                                                    _ = fi112
                                                    _ = v113
                                                }
                                                for fi114 := 0; fi114 < 1; fi114++ {
                                                    _ = len(v55)
                                                    switch "a" {
                                                    case "":
                                                        _ = len(v55)
                                                    default:
                                                        switch (fi114 ^ (fi114 & (fi31 * 2147483647))) {
                                                        case -919:
                                                            if len(v55) >= 2 {
                                                                v55 = v55[1:len(v55)]
                                                            }
                                                        default:
                                                            var v115 uint16 = 65535
                                                            _ = v115
                                                        }
                                                    }
                                                    _ = fi114
                                                }
                                                if len(v55) >= 2 {
                                                    v55 = v55[1:len(v55)]
                                                }
                                            }
                                        }
                                    }
                                }
                                _ = v54
                                _ = v55
                            }
                            _ = fi31
                        }
                    }
                    _ = v30
                }
            }
        }
    } else {
        var v116 uint16 = 42
        if (2 < 2) {
            if (v3 && false) {
                var v117 int16 = v2
                fmt.Println(v116, v5)
                if (!(!(v3 && v4))) {
                    if v3 {
                        v116 -= v116
                        v116 = v116
                    } else {
                        switch (("\n\t" + "hello") + ("\n\t" + ("\n\t" + ("\n\t" + "")))) {
                        case "hello":
                            v118 := GenS11{F0: 2, F1: 0}
                            _ = v118
                        case "\n\t":
                            if (v116 == genFn0(771)) {
                                v119 := GenS10{F0: (-(v117)), F1: v23}
                                for fi120 := 0; fi120 < 2; fi120++ {
                                    switch -1 {
                                    case 128:
                                        v3 = v3
                                    default:
                                        var v121 uint8 = 255
                                        _ = v121
                                    }
                                    var v122 uint16 = v116
                                    v123 := []int{-554}
                                    _ = fi120
                                    _ = v122
                                    _ = v123
                                }
                                _ = v119
                            } else {
                                for fi124 := 0; fi124 < 9; fi124++ {
                                    for fi125 := 0; fi125 < 3; fi125++ {
                                        v126 := GenS9{F0: 100}
                                        if false {
                                            fmt.Println(v5, v5, v117)
                                            v127 := GenS9{F0: 127}
                                            _ = v127
                                        } else {
                                            v128 := GenS2{F0: 255, F1: v3, F2: v3}
                                            if ((!("\n\t" != "hello")) || true) {
                                                fmt.Println(fi125, v1, v4)
                                                v129 := []int{(^(fi124))}
                                                fmt.Println(v3, v2, v4)
                                                var v130 int = fi124
                                                _ = v129
                                                _ = v130
                                            }
                                            v5 = (v2 * (v5 >> ((826) & 63)))
                                            _ = v128
                                        }
                                        v131 := GenS9{F0: 127}
                                        _ = fi125
                                        _ = v126
                                        _ = v131
                                    }
                                    _ = fi124
                                }
                                v3 = v3
                            }
                        default:
                            v132 := GenS2{F0: 533, F1: v4, F2: v4}
                            _ = v132
                        }
                        fmt.Println(v5)
                    }
                } else {
                    v133 := GenS8{F0: (v23 | v23), F1: (634 <= 717)}
                    for fi134 := 0; fi134 < 8; fi134++ {
                        for fi135 := 0; fi135 < 15; fi135++ {
                            v136 := []int{(fi135 & (fi134 | (fi135 & fi135))), fi134, (fi135 * (^(fi135))), fi134}
                            _ = fi135
                            _ = v136
                        }
                        for fi137 := 0; fi137 < 17; fi137++ {
                            var v138 int = (fi137 >> ((65535) & 63))
                            _ = fi137
                            _ = v138
                        }
                        _ = fi134
                    }
                    _ = v133
                }
                _ = v117
            }
            if false {
                for fi139 := 0; fi139 < 17; fi139++ {
                    if false {
                        v140 := []int{}
                        v1 = v1
                        _ = v140
                    }
                    _ = fi139
                }
                var v141 int64 = 0
                _ = v141
            } else {
                v3 = ((((-2.25 * 1.5) * 0.0) + 0.1) != -2.25)
                var v142 int8 = 127
                _ = v142
            }
        }
        if (65535 >= 2147483648) {
            if (0 == 0) {
                v143 := []int{9223372036854775807, -70, -9223372036854775808}
                switch "" {
                case "a":
                    v23 += (v23 * (v23 << ((450) & 63)))
                default:
                    fmt.Println(v5, v23)
                }
                fmt.Println(v116, v2, v5)
                _ = v143
            } else {
                var v144 int32 = -1
                v145 := GenS10{F0: v2, F1: 2147483647}
                if v3 {
                    var v146 int32 = (v144 + -61)
                    fmt.Println(v146, v1)
                    _ = v146
                } else {
                    v147 := []int{42, 0}
                    v144 = v23
                    _ = v147
                }
                _ = v144
                _ = v145
            }
            var v148 int8 = -1
            _ = v148
        }
        _ = v116
    }
    var v149 float64 = 0.1
    v150 := []int{32767}
    for fi151 := 0; fi151 < 20; fi151++ {
        v152 := GenS12{F0: 663, F1: 4294967295, F2: 42}
        _ = fi151
        _ = v152
    }
    var v153 uint64 = 236
    fmt.Println(v149, v2, v23)
    for fi154 := 0; fi154 < 12; fi154++ {
        var v155 float64 = v149
        var v156 int8 = 109
        fmt.Println(v155, v23)
        _ = fi154
        _ = v155
        _ = v156
    }
    if (3.14159 < float64(v1)) {
        for fi157 := 0; fi157 < 11; fi157++ {
            v5 -= (v2 >> ((1) & 63))
            _ = fi157
        }
    } else {
        fmt.Println(v153)
        v158 := GenS13{F0: (32767 <= 32767), F1: 2, F2: 2147483647}
        if (!v3) {
            for fi159 := 0; fi159 < 9; fi159++ {
                switch "\n\t" {
                case "foo bar":
                    v4 = true
                default:
                    switch fi159 {
                    case 42:
                        var v160 uint32 = 1
                        _ = v160
                    case 32767:
                        v1 += v1
                    }
                }
                v5 -= v2
                for fi161 := 0; fi161 < 4; fi161++ {
                    if v3 {
                        var v162 int = fi161
                        v162 -= v162
                        _ = len(v150)
                        _ = v162
                    } else {
                        fmt.Println(fi159, v5)
                        _ = len(v150)
                    }
                    _ = fi161
                }
                _ = fi159
            }
        } else {
            var v163 uint8 = 29
            _ = v163
        }
        _ = v158
    }
    fmt.Println("v1:", v1)
    fmt.Println("v149:", v149)
    fmt.Println("slicelen v150:", len(v150), v150)
    fmt.Println("v153:", v153)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v149
    _ = v150
    _ = v153
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
