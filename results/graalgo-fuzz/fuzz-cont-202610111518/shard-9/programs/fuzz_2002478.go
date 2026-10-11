package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 int8; F2 int }
type GenS4 struct { F0 string; F1 uint64; F2 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64; F1 int; F2 uint16 }
type GenS7 struct { F0 string; F1 uint8; F2 int }
type GenS8 struct { F0 bool; F1 int16 }
type GenS9 struct { F0 int64; F1 int16; F2 int32 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int32; F1 int16 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 int16; F1 int32 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 uint16; F1 uint32; F2 uint }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 int; F1 uint64; F2 bool }
type GenS20 struct { F0 int64; F1 bool }
type GenS21 struct { F0 uint16; F1 int; F2 int8 }
func genFn0(p0 uint64, p1 uint64, p2 string) int8 {
    var v6 int8 = -128
    if true {
        var v7 uint8 = 142
        v8 := GenS1{F0: 32767}
        v9 := GenS2{F0: 933}
        _ = v7
        _ = v8
        _ = v9
    } else {
        var v10 int32 = 781
        p2 = p2
        _ = v10
    }
    v11 := GenS2{F0: 2}
    p0 = (p0 / ((2) | 1))
    var v12 int64 = -439
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v6
    return -70
}

func main() {
    v1 := []int{}
    var v2 int64 = 42
    v3 := make(map[uint32]int)
    var v4 int8 = -1
    var v5 float32 = ((0.1 - 0.0) * (-2.25 - 1.5))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v13 int = -2147483648
    if true {
        v14 := GenS2{F0: 1}
        var v15 int16 = 405
        var v16 string = ((((("a" + ("a" + "foo bar")) + "hello") + "hello") + "hello") + "hello")
        _ = v14
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 1; fi17++ {
            switch (fi17 - v13) {
            case 128:
                for fi18 := 0; fi18 < 2; fi18++ {
                    if (float32(float64(float32((0.1 / 1.5)))) >= v5) {
                        v5 += 3.14159
                        if true {
                            fmt.Println(fi17)
                            if false {
                                v19 := GenS3{F0: (v2 | -585), F1: 100, F2: (fi18 + v13)}
                                fmt.Println(v2)
                                v20 := v3[65535]
                                _ = v19
                                _ = v20
                            }
                            v1 = append(v1, -135)
                        } else {
                            switch fi17 {
                            case 845:
                                var v21 int16 = -138
                                _ = v21
                            default:
                                var v22 uint8 = 49
                                _ = v22
                            }
                            var v23 int16 = 1000
                            _ = v23
                        }
                        _ = len(v3)
                    }
                    v24 := GenS4{F0: ("" + "foo bar"), F1: 1, F2: 2}
                    v5 -= v5
                    _ = fi18
                    _ = v24
                }
            default:
                if (false && ((true || true) || false)) {
                    if false {
                        var v25 uint16 = 483
                        _ = v25
                    } else {
                        if true {
                            v5 -= v5
                            v2 = (v2 % ((v2) | 1))
                        }
                        fmt.Println(fi17, v2, v2)
                        for fi26 := 0; fi26 < 10; fi26++ {
                            v1 = append(v1, (fi26 >> ((255) & 63)))
                            v2 = v2
                            _ = len(v1)
                            var v27 uint32 = 0
                            _ = fi26
                            _ = v27
                        }
                    }
                    if (true && false) {
                        v28 := GenS1{F0: v13}
                        if (!((((v5 - (v5 / (1e10 - -2.25))) - v5) * 0.1) < (0.1 / v5))) {
                            if true {
                                v2 = v2
                                if false {
                                    if (!(false || (float64(v5) != 0.0))) {
                                        delete(v3, 679)
                                        if false {
                                            v29 := GenS3{F0: -814, F1: v4, F2: (v13 & v13)}
                                            v4 = (v4 - genFn0(0, 18446744073709551615, ("a" + ("a" + ("foo bar" + "a")))))
                                            _ = v29
                                        } else {
                                            v30 := GenS5{F0: v2}
                                            _ = v30
                                        }
                                    }
                                    for fi31 := 0; fi31 < 5; fi31++ {
                                        fmt.Println(v13, v2)
                                        _ = fi31
                                    }
                                    for fi32 := 0; fi32 < 12; fi32++ {
                                        var v33 bool = (false || true)
                                        v33 = v33
                                        fmt.Println(v5, v2, v2)
                                        var v34 float64 = 1e10
                                        switch v13 {
                                        case 4294967295:
                                            v33 = (((!v33) && false) || (-647 < 2147483647))
                                        default:
                                            v5 = ((v5 * 0.1) - (v5 - v5))
                                            v35 := GenS1{F0: fi32}
                                            switch (fi17 / (((fi17 / (((v13 & 128)) | 1))) | 1)) {
                                            case 1:
                                                switch ("\n\t" + "\n\t") {
                                                case "":
                                                    var v36 int16 = -1
                                                    v4 += (v4 / ((1) | 1))
                                                    if v33 {
                                                        switch v13 {
                                                        case 1:
                                                            switch "a" {
                                                            case "\n\t":
                                                                var v37 uint8 = 156
                                                                _ = v37
                                                            default:
                                                                v36 += v36
                                                            }
                                                        case 9223372036854775807:
                                                            switch int(v2) {
                                                            case 974:
                                                                var v38 uint = 995
                                                                _ = v38
                                                            }
                                                        case 42:
                                                            switch -124 {
                                                            case 1:
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    var v39 float32 = ((v5 - (0.1 - v5)) * float32(v34))
                                                                    _ = v39
                                                                default:
                                                                    var v40 int8 = 100
                                                                    _ = v40
                                                                }
                                                            case 128:
                                                                switch -433 {
                                                                case 0:
                                                                    if (v33 && ((v13 | fi17) <= -1)) {
                                                                        if (("a" + "\n\t") > "hello") {
                                                                            var v41 int16 = v36
                                                                            v2 = (v2 - v2)
                                                                            v13 = (v13 + (fi17 * fi32))
                                                                            _ = v41
                                                                        } else {
                                                                            v42 := GenS2{F0: 65535}
                                                                            var v43 float64 = v34
                                                                            fmt.Println(v33)
                                                                            _ = v42
                                                                            _ = v43
                                                                        }
                                                                        switch fi17 {
                                                                        case 42:
                                                                            var v44 int64 = (^(v2))
                                                                            _ = v44
                                                                        }
                                                                        v5 = ((v5 + 1.5) / float32((float64(1e10) / (((-2.25 * (v34 / 1e10)) + 1e10) * ((float64(v5) * (((v34 + v34) + 1.5) * v34)) + (v34 * 1e10))))))
                                                                    } else {
                                                                        v36 = -886
                                                                    }
                                                                default:
                                                                    switch (v13 & (fi32 - (v13 + fi32))) {
                                                                    case 32767:
                                                                        if ((v5 + (0.1 / 1.5)) <= ((((0.1 / float32(1.5)) / (1e10 - 0.0)) * v5) + 0.0)) {
                                                                            _ = len(v3)
                                                                            v2 = (v2 + (v2 / ((1) | 1)))
                                                                            v2 = v2
                                                                        } else {
                                                                            v13 += 906
                                                                            fmt.Println(v4, fi32, v33)
                                                                        }
                                                                    default:
                                                                        v1 = append(v1, v13)
                                                                        var v45 int16 = v36
                                                                        switch "a" {
                                                                        case "":
                                                                            var v46 uint8 = 2
                                                                            _ = v46
                                                                        case "hello":
                                                                            if ((v36 * v36) == (v45 + (v36 + 0))) {
                                                                                switch int(v4) {
                                                                                case -2147483648:
                                                                                    if (v4 < (v4 - (-(v4)))) {
                                                                                        var v47 uint = 840
                                                                                        var v48 bool = ((65535 == 42) && v33)
                                                                                        v49 := GenS5{F0: (v2 >> (((^(v47))) & 63))}
                                                                                        _ = v47
                                                                                        _ = v48
                                                                                        _ = v49
                                                                                    }
                                                                                default:
                                                                                    v50 := GenS6{F0: 1, F1: (fi17 % ((fi32) | 1)), F2: 32767}
                                                                                    _ = v50
                                                                                }
                                                                                v51 := GenS2{F0: 0}
                                                                                if false {
                                                                                    v52 := GenS7{F0: "hello", F1: 37, F2: 2147483647}
                                                                                    _ = v52
                                                                                } else {
                                                                                    fmt.Println(v33)
                                                                                    var v53 int64 = 1
                                                                                    _ = v53
                                                                                }
                                                                                _ = v51
                                                                            }
                                                                        default:
                                                                            v36 += (v45 ^ int16(v4))
                                                                        }
                                                                        _ = v45
                                                                    }
                                                                }
                                                            case -1:
                                                                if true {
                                                                    if false {
                                                                        if (0.1 <= ((1e10 / -2.25) + float32(1e10))) {
                                                                            fmt.Println(fi17, v33)
                                                                            if v33 {
                                                                                var v54 int8 = 45
                                                                                v1 = append(v1, fi17)
                                                                                _ = v54
                                                                            }
                                                                            if (!v33) {
                                                                                var v55 bool = ((v4 | (v4 - v4)) == v4)
                                                                                _ = v55
                                                                            }
                                                                        } else {
                                                                            var v56 uint64 = 215
                                                                            var v57 uint16 = 443
                                                                            _ = v56
                                                                            _ = v57
                                                                        }
                                                                    } else {
                                                                        if v33 {
                                                                            _ = len(v3)
                                                                        } else {
                                                                            var v58 uint64 = 42
                                                                            _ = v58
                                                                        }
                                                                    }
                                                                    delete(v3, 4294967295)
                                                                } else {
                                                                    v59 := GenS8{F0: ((!(v33 || v33)) || v33), F1: (v36 >> ((142) & 63))}
                                                                    switch fi32 {
                                                                    case 720:
                                                                        if (!(638 >= v36)) {
                                                                            if true {
                                                                                var v60 uint = 42
                                                                                var v61 int32 = int32(v4)
                                                                                _ = v60
                                                                                _ = v61
                                                                            } else {
                                                                                v62 := GenS5{F0: v2}
                                                                                _ = v62
                                                                            }
                                                                            v3[1] = (fi17 % ((fi32) | 1))
                                                                            v4 = v4
                                                                        } else {
                                                                            v36 -= v36
                                                                        }
                                                                    case 32767:
                                                                        var v63 uint16 = 42
                                                                        _ = v63
                                                                    case 601:
                                                                        if false {
                                                                            var v64 float64 = (v34 / 1.5)
                                                                            var v65 int8 = -1
                                                                            _ = v64
                                                                            _ = v65
                                                                        }
                                                                    default:
                                                                        v66 := GenS9{F0: 42, F1: 32767, F2: 0}
                                                                        _ = v66
                                                                    }
                                                                    v13 = 4294967295
                                                                    _ = v59
                                                                }
                                                            }
                                                        default:
                                                            var v67 bool = true
                                                            _ = v67
                                                        }
                                                        v33 = false
                                                        var v68 int16 = (v36 ^ v36)
                                                        _ = v68
                                                    }
                                                    _ = v36
                                                case "\n\t":
                                                    fmt.Println(v34, v5)
                                                default:
                                                    var v69 uint = 1
                                                    _ = v69
                                                }
                                            case -470:
                                                v1 = append(v1, (v13 + fi32))
                                            case 32767:
                                                var v70 uint8 = 0
                                                _ = v70
                                            }
                                            _ = v35
                                        }
                                        fmt.Println(v2, fi17, v2)
                                        var v71 uint64 = 42
                                        if len(v1) > 0 {
                                            v1[0] = -705
                                        }
                                        v5 -= (float32(v34) / (v5 + 0.0))
                                        _ = fi32
                                        _ = v33
                                        _ = v34
                                        _ = v71
                                    }
                                } else {
                                    for fi72 := 0; fi72 < 17; fi72++ {
                                        v73 := GenS10{F0: (fi17 - 4294967295)}
                                        var v74 int8 = (v4 % ((genFn0(1, 4294967295, "a")) | 1))
                                        v13 = fi17
                                        _ = fi72
                                        _ = v73
                                        _ = v74
                                    }
                                }
                                v1 = append(v1, fi17)
                            }
                        } else {
                            switch ((("a" + "\n\t") + ("hello" + "foo bar")) + "a") {
                            case "a":
                                var v75 uint8 = 126
                                _ = v75
                            }
                            for fi76 := 0; fi76 < 18; fi76++ {
                                if (-32768 < -1) {
                                    fmt.Println(fi17, v5, v13)
                                    v77 := GenS2{F0: 395}
                                    var v78 string = "foo bar"
                                    _ = v77
                                    _ = v78
                                }
                                _ = fi76
                            }
                            v5 -= v5
                        }
                        if (2 == 2) {
                            v79 := GenS4{F0: ("\n\t" + "a"), F1: 42, F2: 688}
                            for fi80 := 0; fi80 < 15; fi80++ {
                                v81 := GenS11{F0: 0}
                                _ = fi80
                                _ = v81
                            }
                            _ = v79
                        } else {
                            switch 9223372036854775807 {
                            case 356:
                                v2 = v2
                            default:
                                fmt.Println(v5, fi17)
                            }
                            v82 := GenS6{F0: 42, F1: 1, F2: 0}
                            switch ("hello" + "a") {
                            case "\n\t":
                                var v83 uint32 = 4294967295
                                _ = v83
                            }
                            _ = v82
                        }
                        _ = v28
                    } else {
                        switch (^(fi17)) {
                        case 9223372036854775807:
                            v1 = append(v1, (^(fi17)))
                        default:
                            v13 += (-(fi17))
                        }
                        v4 += v4
                        var v84 uint = 4
                        _ = v84
                    }
                } else {
                    switch (v13 % (((fi17 ^ fi17)) | 1)) {
                    case 2147483647:
                        if ((1 >= 255) && (fi17 == -1)) {
                            var v85 uint32 = 453
                            if (-1 == -32768) {
                                var v86 int32 = 0
                                switch ("\n\t" + "a") {
                                case "foo bar":
                                    fmt.Println(fi17, v2)
                                case "\n\t":
                                    v2 -= v2
                                case "":
                                    var v87 int8 = v4
                                    _ = v87
                                default:
                                    v88 := GenS12{F0: 1, F1: 0}
                                    _ = v88
                                }
                                var v89 int32 = v86
                                _ = v86
                                _ = v89
                            }
                            _ = v85
                        } else {
                            var v90 uint16 = 2
                            _ = v90
                        }
                    }
                    if len(v1) > 0 {
                        v1[0] = 128
                    }
                }
            }
            fmt.Println(v5, v5, fi17)
            for fi91 := 0; fi91 < 7; fi91++ {
                v92 := GenS11{F0: 0}
                _ = fi91
                _ = v92
            }
            _ = fi17
        }
        var v93 bool = (42 != -2147483648)
        _ = v93
    }
    if true {
        _ = len(v3)
        v94 := GenS13{F0: "a"}
        _ = v94
    } else {
        switch "foo bar" {
        case "a":
            v13 += 128
        case "":
            var v95 float64 = 3.14159
            _ = v95
        }
        var v96 string = "foo bar"
        if (v13 != 4294967295) {
            v97 := GenS14{F0: -1, F1: -243}
            var v98 int64 = -9223372036854775808
            _ = v97
            _ = v98
        } else {
            if false {
                if true {
                    v2 = -1
                }
                v99 := GenS7{F0: v96, F1: 42, F2: v13}
                v100 := GenS9{F0: v2, F1: -1, F2: 794}
                _ = v100
                _ = v99
            } else {
                v5 += float32(1e10)
                fmt.Println(v96, v13, v5)
                var v101 uint64 = 701
                _ = v101
            }
            fmt.Println(v5, v5)
            var v102 int = (v13 >> ((18446744073709551615) & 63))
            _ = v102
        }
        _ = v96
    }
    switch "" {
    case "hello":
        v103 := GenS5{F0: (-(v2))}
        _ = v103
    default:
        v104 := GenS15{F0: v4}
        _ = v104
    }
    if false {
        for fi105 := 0; fi105 < 16; fi105++ {
            v5 += (0.0 - v5)
            _ = fi105
        }
        for fi106 := 0; fi106 < 2; fi106++ {
            if (-2147483648 > 418) {
                if true {
                    v5 = ((v5 / 3.14159) - -2.25)
                } else {
                    v107 := GenS16{F0: 42, F1: 1, F2: 407}
                    if false {
                        if (-74 < (v4 * genFn0(0, 2, ("" + (((((("foo bar" + "foo bar") + "foo bar") + ("hello" + (("hello" + "") + ""))) + ("\n\t" + ("a" + ("a" + "")))) + (("\n\t" + ("hello" + "a")) + "hello")) + "hello"))))) {
                            if (int64(fi106) >= 1) {
                                if len(v1) > 0 {
                                    v1[0] = -75
                                }
                                _ = len(v3)
                                if (!true) {
                                    var v108 int32 = -2147483648
                                    _ = v108
                                }
                            }
                            var v109 int32 = 26
                            _ = v109
                        }
                    }
                    _ = v107
                }
                fmt.Println(v5, v2)
                if false {
                    v2 -= int64(v4)
                }
            }
            var v110 uint16 = 392
            var v111 float64 = 1e10
            _ = fi106
            _ = v110
            _ = v111
        }
        v13 = (^(v13))
    } else {
        var v112 float32 = 1.5
        for fi113 := 0; fi113 < 16; fi113++ {
            _ = len(v3)
            fmt.Println(v13)
            switch "\n\t" {
            case "a":
                var v114 bool = (!(452 == 444))
                _ = v114
            default:
                if ((v13 > fi113) && false) {
                    var v115 float32 = v112
                    if (-2147483648 == 1) {
                        v13 = fi113
                    } else {
                        v13 = (v13 - (v13 * 1))
                    }
                    var v116 uint = 4294967295
                    _ = v115
                    _ = v116
                } else {
                    var v117 uint64 = 442
                    _ = v117
                }
            }
            _ = fi113
        }
        _ = v112
    }
    v118 := v3[972]
    if (!false) {
        var v119 string = "hello"
        fmt.Println(v119, v118, v119)
        for fi120 := 0; fi120 < 18; fi120++ {
            if (325 < 1000) {
                if (907 <= 2) {
                    switch 0 {
                    case -9223372036854775808:
                        for fi121 := 0; fi121 < 6; fi121++ {
                            v119 = v119
                            if ((1 < 15) && (1 <= int16(v13))) {
                                v122 := GenS17{F0: 0}
                                if (true || (float32(-2.25) >= v5)) {
                                    v2 = v2
                                    if (v2 < (v2 >> ((4294967295) & 63))) {
                                        var v123 uint64 = 719
                                        fmt.Println(v4, fi121)
                                        _ = v123
                                    }
                                }
                                v124 := GenS8{F0: false, F1: 0}
                                v118 -= fi120
                                var v125 int = -2147483648
                                v126 := GenS2{F0: 18446744073709551615}
                                v4 = v4
                                switch fi121 {
                                case 128:
                                    var v127 int16 = -892
                                    _ = v127
                                case 32768:
                                    switch "hello" {
                                    case "":
                                        switch (v119 + "\n\t") {
                                        case "a":
                                            var v128 bool = (!true)
                                            _ = v128
                                        case "":
                                            v13 -= 128
                                        case "\n\t":
                                            if (912 <= 975) {
                                                if (!false) {
                                                    v4 = (v4 | genFn0(1, 4294967295, "a"))
                                                    v118 = (-(v125))
                                                }
                                                v129 := GenS18{F0: 0}
                                                _ = v129
                                            }
                                        default:
                                            v118 += v125
                                        }
                                    default:
                                        var v130 int32 = 39
                                        _ = v130
                                    }
                                case 0:
                                    v131 := GenS14{F0: 828, F1: 2147483647}
                                    _ = v131
                                default:
                                    if (true || true) {
                                        var v132 uint64 = 100
                                        var v133 float32 = 0.0
                                        fmt.Println(v5)
                                        _ = v132
                                        _ = v133
                                    } else {
                                        v134 := GenS10{F0: (fi120 + (v125 >> ((118) & 63)))}
                                        _ = v134
                                    }
                                }
                                _ = v122
                                _ = v124
                                _ = v125
                                _ = v126
                            } else {
                                v135 := GenS2{F0: 1}
                                v136 := GenS6{F0: 906, F1: -141, F2: 1}
                                if ((v4 - genFn0(992, 376, "a")) != (v4 + (v4 & v4))) {
                                    fmt.Println(fi120, fi120)
                                }
                                _ = v135
                                _ = v136
                            }
                            v4 = (v4 & (-(v4)))
                            _ = fi121
                        }
                    default:
                        v3[4294967295] = fi120
                    }
                    _ = len(v3)
                }
                switch v118 {
                case -9223372036854775808:
                    v3[65535] = -9223372036854775808
                case -861:
                    var v137 int8 = (v4 + v4)
                    _ = v137
                case 615:
                    v138 := GenS1{F0: -927}
                    _ = v138
                default:
                    fmt.Println(v5, v5, fi120)
                }
                v139 := v3[65535]
                _ = v139
            }
            v140 := GenS19{F0: 1, F1: 42, F2: (("foo bar" + "foo bar") == ("" + v119))}
            switch -1000 {
            case 703:
                for fi141 := 0; fi141 < 8; fi141++ {
                    if true {
                        if false {
                            if (!false) {
                                switch "\n\t" {
                                case "\n\t":
                                    var v142 uint16 = 226
                                    _ = v142
                                case "foo bar":
                                    v119 = (v119 + ("hello" + (v119 + "foo bar")))
                                case "a":
                                    fmt.Println(fi120, fi141)
                                default:
                                    if (!false) {
                                        var v143 int8 = -1
                                        var v144 uint32 = 907
                                        if ((!(true || false)) || true) {
                                            fmt.Println(v119, fi141)
                                        } else {
                                            v119 = v119
                                        }
                                        _ = v143
                                        _ = v144
                                    } else {
                                        v13 -= (^(fi141))
                                        var v145 int32 = 1
                                        v146 := GenS4{F0: v119, F1: 1, F2: 255}
                                        _ = v145
                                        _ = v146
                                    }
                                }
                            } else {
                                v147 := GenS2{F0: 4294967295}
                                var v148 float64 = (1.5 / 1.0)
                                v149 := GenS20{F0: v2, F1: (!(!true))}
                                _ = v147
                                _ = v148
                                _ = v149
                            }
                            var v150 uint32 = 2147483647
                            v5 = (v5 * (v5 - -2.25))
                            _ = v150
                        } else {
                            switch (fi120 | v13) {
                            case -55:
                                var v151 int64 = (^(v2))
                                _ = v151
                            case 4294967295:
                                var v152 uint32 = 4294967295
                                _ = v152
                            case 1:
                                v118 += fi141
                            }
                        }
                        fmt.Println(v118, v5, fi141)
                        v153 := GenS12{F0: 320, F1: 1}
                        if (130 != 4294967295) {
                            var v154 string = "foo bar"
                            _ = v154
                        } else {
                            if (false || false) {
                                switch "foo bar" {
                                case "foo bar":
                                    var v155 bool = false
                                    _ = v155
                                case "a":
                                    var v156 float64 = (1.5 - (-2.25 - 0.0))
                                    _ = v156
                                case "hello":
                                    if true {
                                        delete(v3, 158)
                                        if ((-2.25 / 1.0) != 3.14159) {
                                            switch 1 {
                                            case 485:
                                                v4 -= v4
                                                var v157 bool = false
                                                _ = v157
                                            default:
                                                v4 = (-(v4))
                                            }
                                            switch (fi120 * (v13 % ((v118) | 1))) {
                                            case 9223372036854775807:
                                                var v158 uint32 = 2147483648
                                                _ = v158
                                            default:
                                                v13 -= 2147483647
                                            }
                                            var v159 float64 = -2.25
                                            v13 = fi120
                                            _ = v159
                                        } else {
                                            var v160 bool = false
                                            _ = v160
                                        }
                                    } else {
                                        var v161 uint64 = 768
                                        var v162 int64 = v2
                                        fmt.Println(v13, v2)
                                        _ = v161
                                        _ = v162
                                    }
                                    switch v13 {
                                    case 4294967295:
                                        v163 := GenS20{F0: int64(v4), F1: false}
                                        _ = v163
                                    case 128:
                                        if false {
                                            if false {
                                                if true {
                                                    var v164 int = fi141
                                                    v3[65535] = -1
                                                    v118 += v164
                                                    _ = v164
                                                } else {
                                                    var v165 uint16 = 42
                                                    _ = v165
                                                }
                                                v13 = -489
                                                var v166 int64 = (v2 | (-(v2)))
                                                _ = v166
                                            }
                                            fmt.Println(fi120)
                                        } else {
                                            _ = len(v1)
                                            v119 = "\n\t"
                                            switch v118 {
                                            case 2147483647:
                                                delete(v3, 695)
                                            case -9223372036854775808:
                                                v5 = v5
                                            }
                                        }
                                        v3[514] = 0
                                        var v167 string = v119
                                        fmt.Println(fi120, v13)
                                        _ = v167
                                    default:
                                        fmt.Println(v118, fi120, v119)
                                        if (1e10 >= (0.1 + float64(((-2.25 - (v5 - float32(3.14159))) + v5)))) {
                                            switch fi141 {
                                            case -550:
                                                var v168 int8 = (v4 % ((genFn0(1, 2, v119)) | 1))
                                                _ = v168
                                            case 9223372036854775807:
                                                fmt.Println(v2, v118)
                                            }
                                            v119 = (v119 + "hello")
                                        }
                                    }
                                    fmt.Println(v13)
                                    var v169 uint64 = 18446744073709551615
                                    var v170 int32 = -1
                                    _ = v169
                                    _ = v170
                                }
                            } else {
                                delete(v3, 1)
                                var v171 uint32 = 65535
                                switch (fi141 << ((4294967295) & 63)) {
                                case -1:
                                    _ = len(v3)
                                case 32767:
                                    var v172 int64 = (-(v2))
                                    _ = v172
                                case 4294967295:
                                    if false {
                                        fmt.Println(v5)
                                        switch v119 {
                                        case "a":
                                            var v173 float32 = v5
                                            _ = v173
                                        default:
                                            v5 = (-2.25 - float32(3.14159))
                                        }
                                        var v174 uint64 = 1
                                        _ = v174
                                    } else {
                                        var v175 int32 = 1
                                        var v176 int8 = v4
                                        v177 := GenS13{F0: "\n\t"}
                                        _ = v175
                                        _ = v176
                                        _ = v177
                                    }
                                }
                                _ = v171
                            }
                        }
                        v178 := v3[4294967295]
                        _ = v153
                        _ = v178
                    } else {
                        v4 = (v4 << ((1) & 63))
                        v2 += -9223372036854775808
                    }
                    _ = fi141
                }
            case -9223372036854775808:
                if false {
                    if true {
                        v2 += -1
                    }
                    v2 += -219
                    if true {
                        if ((((v5 - float32(float64(((v5 + v5) - float32(float64(v5)))))) * (float32(((0.0 - -2.25) / 1.0)) + ((v5 - v5) / v5))) - v5) >= v5) {
                            fmt.Println(v119, v118)
                            v119 = (v119 + v119)
                            v179 := v3[445]
                            _ = v179
                        } else {
                            fmt.Println(v4, v2)
                            switch v119 {
                            case "":
                                v180 := GenS13{F0: v119}
                                _ = v180
                            case "a":
                                v2 -= (v2 + (v2 & (v2 >> ((529) & 63))))
                            default:
                                v4 += v4
                            }
                        }
                        v181 := GenS7{F0: (v119 + v119), F1: 128, F2: (v13 << ((4294967295) & 63))}
                        _ = v181
                    } else {
                        var v182 int16 = -198
                        _ = v182
                    }
                }
            case -1:
                for fi183 := 0; fi183 < 15; fi183++ {
                    v119 = ((((v119 + v119) + "foo bar") + "a") + (v119 + "a"))
                    _ = fi183
                }
            }
            _ = fi120
            _ = v140
        }
        _ = v119
    }
    fmt.Println(v5, v2)
    var v184 uint16 = 0
    v185 := GenS6{F0: 0, F1: v13, F2: 772}
    for fi186 := 0; fi186 < 7; fi186++ {
        v5 = ((v5 + v5) * ((v5 - float32((1e10 * ((((1.5 / 3.14159) * -2.25) + float64(v5)) + ((0.1 * 3.14159) - float64((((((v5 + 0.0) - (3.14159 / v5)) / (((v5 * (v5 / -2.25)) + float32(0.0)) * 0.1)) + (float32(-2.25) - float32(1e10))) + v5))))))) + 0.0))
        _ = fi186
    }
    if true {
        if (((v184 == v184) || (2147483647 > -798)) || (true || false)) {
            v184 = v184
            fmt.Println(v13, v4)
            var v187 int8 = genFn0(322, 565, "")
            _ = v187
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            delete(v3, 4294967295)
        }
        for fi188 := 0; fi188 < 13; fi188++ {
            if (2147483647 < 42) {
                v189 := GenS12{F0: 193, F1: -32768}
                fmt.Println(v5, v4, v13)
                v4 = v4
                _ = v189
            } else {
                switch (v13 * (fi188 | (v13 << ((18446744073709551615) & 63)))) {
                case -82:
                    var v190 uint = 156
                    _ = v190
                case 127:
                    for fi191 := 0; fi191 < 18; fi191++ {
                        if (0 <= 1) {
                            v5 += 0.0
                            fmt.Println(v2, v13, v5)
                        }
                        v192 := v3[uint32(v184)]
                        if false {
                            v193 := GenS6{F0: 42, F1: fi191, F2: (v184 >> ((4294967295) & 63))}
                            _ = v193
                        } else {
                            var v194 int64 = v2
                            v1 = append(v1, v192)
                            _ = v194
                        }
                        _ = fi191
                        _ = v192
                    }
                case 1:
                    if (true && (!(v2 == (v2 ^ (v2 << ((255) & 63)))))) {
                        _ = len(v3)
                    } else {
                        var v195 float64 = -2.25
                        var v196 uint64 = 1
                        _ = v195
                        _ = v196
                    }
                }
                var v197 uint8 = 9
                switch -110 {
                case 150:
                    fmt.Println(v4, v2, v2)
                }
                _ = v197
            }
            _ = fi188
        }
        fmt.Println(v184)
    }
    for fi198 := 0; fi198 < 9; fi198++ {
        v184 = v184
        var v199 int32 = 0
        _ = fi198
        _ = v199
    }
    delete(v3, 2147483648)
    fmt.Println(v2, v5, v4)
    for fi200 := 0; fi200 < 14; fi200++ {
        v201 := GenS21{F0: (v184 - 682), F1: v13, F2: v4}
        for fi202 := 0; fi202 < 18; fi202++ {
            var v203 uint8 = 127
            switch "hello" {
            case "":
                v204 := GenS9{F0: (v2 << ((4294967295) & 63)), F1: 0, F2: -292}
                _ = v204
            case "foo bar":
                v2 = (v2 + v2)
            default:
                v5 -= v5
            }
            _ = fi202
            _ = v203
        }
        var v205 float32 = v5
        _ = fi200
        _ = v201
        _ = v205
    }
    v4 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v118:", v118)
    fmt.Println("v13:", v13)
    fmt.Println("v184:", v184)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v118
    _ = v13
    _ = v184
    _ = v185
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
