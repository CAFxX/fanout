package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int32; F1 uint64; F2 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16; F1 int8; F2 int }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 bool; F1 int16; F2 bool }
type GenS9 struct { F0 uint; F1 int64; F2 uint }
type GenS10 struct { F0 uint16; F1 uint8 }
type GenS11 struct { F0 uint16; F1 int16; F2 bool }
type GenS12 struct { F0 string; F1 int16 }
type GenS13 struct { F0 uint64; F1 string; F2 int64 }
type GenS14 struct { F0 int; F1 uint }
type GenS15 struct { F0 uint32; F1 string }
type GenS16 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 uint8; F1 uint8 }
type GenS19 struct { F0 int16; F1 uint; F2 uint16 }
type GenS20 struct { F0 int16; F1 int64 }
type GenS21 struct { F0 bool }
type GenS22 struct { F0 int; F1 int32 }
func genFn0() int {
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            if ((!(!(!(9223372036854775807 >= 9223372036854775807)))) || (!true)) {
                if true {
                    fmt.Println(fi8, fi7)
                    var v9 string = "\n\t"
                    if false {
                        v10 := []int{fi8, 0, 4294967295}
                        v9 = "hello"
                        var v11 float32 = (float32(0.1) + 1.5)
                        v9 = ((("\n\t" + (v9 + ("hello" + ""))) + v9) + "a")
                        _ = v10
                        _ = v11
                    }
                    _ = v9
                } else {
                    var v12 uint8 = 42
                    if false {
                        var v13 int = fi8
                        fmt.Println(v12, v12)
                        _ = v13
                    }
                    _ = v12
                }
                var v14 int64 = 1
                var v15 bool = (!(false && (false && false)))
                var v16 string = ""
                var v17 float64 = 0.1
                _ = v14
                _ = v15
                _ = v16
                _ = v17
            } else {
                switch "foo bar" {
                case "a":
                    var v18 uint8 = 255
                    _ = v18
                case "hello":
                    v19 := []int{fi7}
                    _ = v19
                case "foo bar":
                    var v20 uint16 = 32767
                    _ = v20
                default:
                    var v21 int = (fi8 % ((-836) | 1))
                    _ = v21
                }
                var v22 int16 = 32767
                _ = v22
            }
            _ = fi8
        }
        if (-752 <= -48) {
            switch 128 {
            case -2147483648:
                var v23 uint16 = 32767
                _ = v23
            case 128:
                var v24 uint8 = 174
                _ = v24
            case -725:
                v25 := []int{fi7, 127, (-(fi7)), fi7}
                _ = v25
            }
        }
        _ = fi7
    }
    var v26 int64 = 1
    for fi27 := 0; fi27 < 20; fi27++ {
        v28 := GenS1{F0: 558}
        for fi29 := 0; fi29 < 20; fi29++ {
            if false {
                fmt.Println(v26)
                v30 := []int{int(v26)}
                _ = v30
            }
            switch "a" {
            case "a":
                v31 := GenS2{F0: 77, F1: 0}
                _ = v31
            case "":
                v32 := GenS3{F0: 2147483647, F1: 42, F2: 0}
                _ = v32
            default:
                switch 0 {
                case 797:
                    var v33 uint32 = 1
                    if true {
                        v34 := GenS4{F0: "a"}
                        switch "" {
                        case "":
                            if (!true) {
                                switch fi29 {
                                case -9223372036854775808:
                                    fmt.Println(v33, fi27, fi27)
                                case 2147483647:
                                    v26 = 1
                                case 949:
                                    fmt.Println(fi27, fi29)
                                default:
                                    fmt.Println(v33, fi29, v33)
                                }
                                if (!true) {
                                    v35 := GenS5{F0: 32767, F1: -13, F2: 35}
                                    v36 := GenS6{F0: 255, F1: 571}
                                    v37 := GenS1{F0: 42}
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                                var v38 int8 = -1
                                _ = v38
                            }
                        case "\n\t":
                            v39 := []int{fi29}
                            _ = v39
                        case "hello":
                            fmt.Println(v33, v26)
                            v33 -= (v33 & v33)
                        default:
                            var v40 int8 = 0
                            v33 = v33
                            if (false || true) {
                                v26 = v26
                                var v41 float32 = 1.5
                                if (0 != 229) {
                                    var v42 float64 = 0.1
                                    v43 := GenS3{F0: -1, F1: 4294967295, F2: 0}
                                    _ = v42
                                    _ = v43
                                }
                                v44 := GenS7{F0: 42}
                                if (!(255 == 65535)) {
                                    v40 -= int8(fi27)
                                    v45 := []int{fi29, fi27, fi27, 128}
                                    fmt.Println(fi27, fi27)
                                    if len(v45) >= 2 {
                                        v45 = v45[1:len(v45)]
                                    }
                                    v46 := GenS4{F0: ""}
                                    _ = v45
                                    _ = v46
                                }
                                _ = v41
                                _ = v44
                            }
                            _ = v40
                        }
                        v33 -= (v33 + v33)
                        _ = v34
                    } else {
                        v47 := []int{(fi29 - fi27), fi27, (^(fi27))}
                        switch fi29 {
                        case 42:
                            v26 = int64(v33)
                        case 4294967295:
                            if true {
                                if len(v47) >= 2 {
                                    v47 = v47[1:len(v47)]
                                }
                                v26 -= v26
                            } else {
                                var v48 uint64 = 872
                                if false {
                                    v49 := GenS1{F0: v48}
                                    _ = v49
                                }
                                if false {
                                    v47 = append(v47, int(v48))
                                }
                                _ = v48
                            }
                        case -41:
                            var v50 string = "\n\t"
                            switch fi27 {
                            case 32768:
                                if (2 >= 4294967295) {
                                    v47 = append(v47, (fi27 >> ((255) & 63)))
                                } else {
                                    switch ((("foo bar" + v50) + v50) + (("a" + v50) + "a")) {
                                    case "foo bar":
                                        v33 = 65535
                                        var v51 int = (^(fi29))
                                        v26 -= v26
                                        var v52 float64 = ((0.0 / 1.0) * -2.25)
                                        var v53 int16 = 1000
                                        if (true || true) {
                                            switch (("" + "hello") + v50) {
                                            case "a":
                                                var v54 int32 = 2147483647
                                                var v55 int64 = (^(v26))
                                                _ = v54
                                                _ = v55
                                            case "hello":
                                                fmt.Println(fi29, fi29, fi27)
                                            case "\n\t":
                                                v56 := GenS8{F0: (((((((-2.25 + 0.1) + (float32(((v52 + v52) / ((v52 - (((3.14159 - ((float64(1.5) / (1.5 - (v52 + (3.14159 - (v52 - v52))))) + (0.0 / (float64(0.0) + 0.1)))) / 3.14159) / v52)) * (float64(0.1) * v52)))) * 0.0)) - (1.5 / (0.1 - 0.0))) + 3.14159) + 0.1) * (1e10 + (1.5 + 3.14159))) <= -2.25), F1: (v53 | 281), F2: true}
                                                _ = v56
                                            }
                                        }
                                        v52 = v52
                                        _ = v51
                                        _ = v52
                                        _ = v53
                                    default:
                                        v57 := GenS5{F0: 42, F1: int8(fi29), F2: (fi27 >> ((0) & 63))}
                                        _ = v57
                                    }
                                    v50 = "hello"
                                }
                                var v58 uint32 = v33
                                if false {
                                    var v59 float64 = 0.1
                                    v33 -= (v58 | v58)
                                    _ = v59
                                } else {
                                    if false {
                                        fmt.Println(fi27, fi29, v26)
                                        v60 := GenS9{F0: 1, F1: (v26 / (((v26 * (v26 * v26))) | 1)), F2: 0}
                                        v61 := GenS10{F0: 170, F1: 1}
                                        _ = v60
                                        _ = v61
                                    }
                                }
                                switch v50 {
                                case "hello":
                                    var v62 uint32 = v58
                                    _ = v62
                                }
                                var v63 uint = 42
                                v26 = 1
                                switch 410 {
                                case 1:
                                    switch v50 {
                                    case "hello":
                                        if (((32767 != 93) || false) || (42 == 0)) {
                                            switch fi27 {
                                            case 0:
                                                var v64 string = "\n\t"
                                                _ = v64
                                            case -2147483648:
                                                var v65 int16 = 1000
                                                _ = v65
                                            case 42:
                                                var v66 float32 = -2.25
                                                v33 += v33
                                                if ((v26 & v26) < (v26 + (v26 % ((v26) | 1)))) {
                                                    fmt.Println(v63)
                                                } else {
                                                    _ = len(v47)
                                                    v33 -= v58
                                                }
                                                _ = v66
                                            default:
                                                switch fi29 {
                                                case -9223372036854775808:
                                                    v47 = append(v47, fi27)
                                                case -121:
                                                    v33 += 475
                                                    var v67 uint8 = 0
                                                    _ = v67
                                                case 32768:
                                                    v63 = (v63 & (^(v63)))
                                                default:
                                                    v26 += (v26 << (((v63 / ((882) | 1))) & 63))
                                                }
                                            }
                                            v58 = v33
                                        }
                                    case "foo bar":
                                        if len(v47) > 0 {
                                            v47[0] = (fi27 >> (((v63 + (^(v63)))) & 63))
                                        }
                                    default:
                                        switch "foo bar" {
                                        case "":
                                            v68 := GenS9{F0: (^(v63)), F1: 42, F2: (v63 % (((v63 ^ (^(v63)))) | 1))}
                                            _ = v68
                                        case "hello":
                                            if (18446744073709551615 == 0) {
                                                if (((((-2.25 * 0.1) * 0.1) / 1.0) * (-2.25 + -2.25)) != ((float64((float32(0.1) / (-2.25 - 0.1))) * 3.14159) - (1e10 - 0.0))) {
                                                    if (!true) {
                                                        v63 = (v63 + 4294967295)
                                                    }
                                                } else {
                                                    v47 = append(v47, fi29)
                                                    v33 -= (v58 % ((v58) | 1))
                                                }
                                                var v69 int16 = -1
                                                var v70 uint16 = 42
                                                _ = v69
                                                _ = v70
                                            } else {
                                                v26 = -9223372036854775808
                                            }
                                        case "foo bar":
                                            v47 = append(v47, fi27)
                                        }
                                    }
                                case -2147483648:
                                    v26 = v26
                                case -967:
                                    switch fi27 {
                                    case -163:
                                        var v71 uint64 = 325
                                        _ = v71
                                    case 2147483647:
                                        var v72 float32 = 1e10
                                        _ = v72
                                    case -207:
                                        var v73 int32 = -2147483648
                                        _ = v73
                                    default:
                                        v74 := GenS11{F0: 1, F1: 0, F2: true}
                                        _ = v74
                                    }
                                }
                                _ = v58
                                _ = v63
                            case 2147483647:
                                if (0 > 255) {
                                    var v75 int64 = (^(v26))
                                    fmt.Println(fi27)
                                    _ = v75
                                } else {
                                    switch 128 {
                                    case 0:
                                        switch fi27 {
                                        case 1:
                                            v33 -= v33
                                        case 4294967295:
                                            if (!(1 >= 242)) {
                                                fmt.Println(v33, fi29)
                                            } else {
                                                v50 = ("hello" + v50)
                                                fmt.Println(fi27, v26)
                                                v33 = (v33 * 518)
                                            }
                                        }
                                    case 4294967295:
                                        v50 = v50
                                    case 626:
                                        fmt.Println(v33, v26, v33)
                                    default:
                                        v33 = v33
                                    }
                                    if (uint16(v33) < 65535) {
                                        v33 = v33
                                        v76 := GenS12{F0: v50, F1: 1000}
                                        if false {
                                            fmt.Println(fi29, v26)
                                            v50 = (v50 + v50)
                                            v47 = append(v47, fi27)
                                        }
                                        _ = v76
                                    }
                                }
                            case 32767:
                                v77 := GenS5{F0: 1000, F1: 0, F2: fi29}
                                _ = v77
                            default:
                                fmt.Println(v50)
                            }
                            _ = v50
                        }
                        if (127 != 12) {
                            if len(v47) > 0 {
                                v47[0] = fi29
                            }
                        } else {
                            var v78 uint8 = 2
                            _ = v78
                        }
                        if (!true) {
                            var v79 int8 = -128
                            var v80 uint32 = v33
                            switch ((("foo bar" + ("hello" + ("\n\t" + ("foo bar" + "\n\t")))) + (((((("foo bar" + (("\n\t" + ("" + "a")) + "foo bar")) + "a") + ("foo bar" + "a")) + "hello") + "\n\t") + "hello")) + ((("a" + "a") + "foo bar") + "")) {
                            case "a":
                                var v81 int32 = -844
                                _ = v81
                            case "hello":
                                v80 = (v33 >> ((335) & 63))
                            }
                            _ = v79
                            _ = v80
                        }
                        _ = v47
                    }
                    v82 := []int{}
                    _ = v33
                    _ = v82
                }
            }
            v83 := []int{fi29, (fi29 + (fi29 << ((517) & 63)))}
            _ = fi29
            _ = v83
        }
        _ = fi27
        _ = v28
    }
    fmt.Println(v26)
    var v84 float32 = (1.5 / float32((3.14159 + 0.0)))
    _ = v26
    _ = v84
    return 42
}
func genFn1(p0 uint16, p1 uint, p2 bool) int {
    v85 := GenS13{F0: 82, F1: "\n\t", F2: -1}
    for fi86 := 0; fi86 < 13; fi86++ {
        var v87 float32 = 1e10
        _ = fi86
        _ = v87
    }
    p1 = (^(p1))
    var v88 int8 = -128
    _ = p0
    _ = p1
    _ = p2
    _ = v85
    _ = v88
    return 1
}

func main() {
    var v1 uint8 = 2
    var v2 float64 = (3.14159 * float64(float32(3.14159)))
    v3 := make(map[string]int)
    var v4 uint8 = 0
    v5 := []uint32{229, 1, uint32(v1)}
    var v6 int64 = -830
    if (65535 != 864) {
        var v89 uint16 = 326
        _ = v89
    } else {
        var v90 int32 = 760
        var v91 bool = false
        _ = v90
        _ = v91
    }
    var v92 int64 = 0
    v6 -= v6
    if len(v5) > 0 {
        v5[0] = 669
    }
    v92 = (v92 << ((42) & 63))
    switch 9223372036854775807 {
    case 127:
        fmt.Println(v4, v92, v2)
    case 9223372036854775807:
        var v93 int32 = 2147483647
        _ = v93
    case 128:
        for fi94 := 0; fi94 < 8; fi94++ {
            fmt.Println(v1, fi94)
            for fi95 := 0; fi95 < 20; fi95++ {
                var v96 int32 = -1
                switch (^(fi95)) {
                case 32767:
                    switch (("a" + ("a" + ("" + (("" + "hello") + ("hello" + ""))))) + "\n\t") {
                    case "foo bar":
                        if (42 == 220) {
                            fmt.Println(v2)
                            v5 = append(v5, 2)
                            var v97 int = (^(fi94))
                            _ = v97
                        }
                    }
                case 0:
                    v6 = v92
                }
                v2 = float64((3.14159 + 0.1))
                _ = fi95
                _ = v96
            }
            var v98 int32 = -2147483648
            _ = fi94
            _ = v98
        }
    default:
        v1 -= v1
    }
    if false {
        var v99 string = ""
        fmt.Println(v6, v6, v6)
        _ = v99
    }
    if (!(true && false)) {
        v5 = append(v5, 2)
    } else {
        if true {
            for fi100 := 0; fi100 < 11; fi100++ {
                if (3.14159 <= 1e10) {
                    v3["a"] = (fi100 - -590)
                }
                var v101 int8 = 0
                fmt.Println(v4, v4)
                _ = fi100
                _ = v101
            }
        }
        var v102 int8 = 127
        switch -2147483648 {
        case 0:
            var v103 float64 = ((float64(((float32(1.5) / 3.14159) * float32(v2))) + 1e10) * v2)
            _ = v103
        case 42:
            if ((true || ((v1 << ((255) & 63)) >= v1)) && (!true)) {
                for fi104 := 0; fi104 < 14; fi104++ {
                    v2 -= v2
                    if false {
                        switch "hello" {
                        case "\n\t":
                            var v105 float64 = 0.0
                            _ = v105
                        case "foo bar":
                            var v106 uint32 = 0
                            _ = v106
                        case "hello":
                            var v107 uint8 = v4
                            _ = v107
                        }
                        if true {
                            v1 = 255
                            v92 -= v6
                        } else {
                            switch fi104 {
                            case -687:
                                v5 = append(v5, 439)
                            case -490:
                                var v108 float64 = 1e10
                                _ = v108
                            }
                        }
                    }
                    _ = fi104
                }
                for fi109 := 0; fi109 < 6; fi109++ {
                    if len(v5) > 0 {
                        v5[0] = 4294967295
                    }
                    _ = fi109
                }
                v110 := GenS14{F0: 9223372036854775807, F1: 4294967295}
                _ = v110
            }
        case 2147483647:
            var v111 int8 = (v102 / (((v102 / ((-70) | 1))) | 1))
            _ = v111
        }
        _ = v102
    }
    var v112 uint64 = 1
    v6 = 0
    delete(v3, ("\n\t" + ("hello" + "")))
    switch "hello" {
    case "foo bar":
        fmt.Println(v112)
    case "hello":
        if len(v5) > 0 {
            v5[0] = 340
        }
    default:
        for fi113 := 0; fi113 < 4; fi113++ {
            _ = len(v5)
            _ = fi113
        }
    }
    if false {
        for fi114 := 0; fi114 < 17; fi114++ {
            var v115 uint8 = (^(v4))
            var v116 float32 = (1e10 / 1.5)
            _ = fi114
            _ = v115
            _ = v116
        }
        var v117 int = 0
        _ = v117
    } else {
        for fi118 := 0; fi118 < 20; fi118++ {
            for fi119 := 0; fi119 < 8; fi119++ {
                fmt.Println(v2, v112, v92)
                switch ("foo bar" + "foo bar") {
                case "\n\t":
                    if true {
                        var v120 float32 = (float32(3.14159) / float32(v2))
                        switch (fi119 | 2147483647) {
                        case 127:
                            var v121 int16 = -450
                            _ = v121
                        case -796:
                            var v122 int = -980
                            _ = v122
                        default:
                            var v123 string = "hello"
                            _ = v123
                        }
                        _ = v120
                    } else {
                        fmt.Println(v4, v2)
                    }
                case "a":
                    v2 = v2
                case "":
                    var v124 int16 = -876
                    _ = v124
                default:
                    v125 := GenS15{F0: 0, F1: "\n\t"}
                    _ = v125
                }
                if (((true || ((^(fi118)) == 1)) || (2147483647 < 411)) || false) {
                    fmt.Println(v4)
                } else {
                    if false {
                        fmt.Println(v4, v6, fi118)
                    }
                    v3["foo bar"] = fi118
                    v4 -= v4
                }
                _ = fi119
            }
            _ = fi118
        }
        switch 580 {
        case 456:
            if false {
                if (18446744073709551615 < 18446744073709551615) {
                    if (!(273 <= 675)) {
                        var v126 int16 = int16(v92)
                        _ = v126
                    } else {
                        for fi127 := 0; fi127 < 12; fi127++ {
                            if false {
                                var v128 int32 = 0
                                v2 -= (-2.25 + v2)
                                _ = v128
                            } else {
                                if false {
                                    fmt.Println(v92, v112, v4)
                                    v6 -= (v6 ^ 0)
                                    fmt.Println(v112, v4)
                                }
                                for fi129 := 0; fi129 < 18; fi129++ {
                                    delete(v3, "a")
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    var v130 string = ""
                                    _ = fi129
                                    _ = v130
                                }
                            }
                            fmt.Println(v112)
                            _ = fi127
                        }
                        fmt.Println(v1, v112, v1)
                    }
                    switch "" {
                    case "foo bar":
                        for fi131 := 0; fi131 < 14; fi131++ {
                            v112 -= (v112 | (v112 + (v112 | v112)))
                            _ = fi131
                        }
                    case "":
                        var v132 float64 = (v2 / (v2 * 1.5))
                        _ = v132
                    }
                    switch 42 {
                    case 4294967295:
                        v133 := GenS7{F0: 574}
                        _ = v133
                    case 670:
                        v134 := GenS3{F0: int32(v92), F1: (v112 % ((4294967295) | 1)), F2: -1}
                        _ = v134
                    case -340:
                        if len(v5) > 0 {
                            v5[0] = 395
                        }
                    default:
                        v135 := GenS16{F0: (v92 % (((^(v6))) | 1)), F1: 65535, F2: 2}
                        _ = v135
                    }
                }
                if true {
                    if (!true) {
                        _ = len(v5)
                        var v136 bool = (!(v112 < v112))
                        _ = v136
                    }
                    if true {
                        for fi137 := 0; fi137 < 10; fi137++ {
                            var v138 uint8 = (v1 >> ((uint(v4)) & 63))
                            v2 -= 3.14159
                            if ("foo bar" != "foo bar") {
                                v139 := GenS15{F0: 4294967295, F1: (("hello" + "\n\t") + "foo bar")}
                                _ = v139
                            }
                            _ = fi137
                            _ = v138
                        }
                        var v140 uint = 65535
                        var v141 uint32 = 472
                        _ = v140
                        _ = v141
                    } else {
                        switch 127 {
                        case 4294967295:
                            for fi142 := 0; fi142 < 10; fi142++ {
                                var v143 int64 = (v92 * -965)
                                fmt.Println(fi142, fi142, v92)
                                v1 = (v1 * v1)
                                _ = fi142
                                _ = v143
                            }
                        case 1:
                            if false {
                                fmt.Println(v6, v4, v6)
                                switch 9223372036854775807 {
                                case 127:
                                    switch 774 {
                                    case 128:
                                        if true {
                                            v144 := GenS4{F0: "\n\t"}
                                            v2 -= 1e10
                                            _ = v144
                                        }
                                    default:
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                }
                                var v145 uint64 = (^(v112))
                                _ = v145
                            }
                        default:
                            switch -495 {
                            case 128:
                                v146 := GenS4{F0: (("hello" + "hello") + "\n\t")}
                                _ = v146
                            case 9223372036854775807:
                                if len(v5) > 0 {
                                    v5[0] = uint32(v112)
                                }
                            }
                        }
                        var v147 float64 = v2
                        var v148 string = "hello"
                        _ = v147
                        _ = v148
                    }
                }
                if true {
                    switch "a" {
                    case "\n\t":
                        var v149 string = ((((("hello" + ("" + ("" + "foo bar"))) + "hello") + "") + "") + "a")
                        _ = v149
                    case "hello":
                        fmt.Println(v1)
                    case "foo bar":
                        if false {
                            for fi150 := 0; fi150 < 5; fi150++ {
                                v151 := GenS12{F0: "", F1: 1000}
                                if true {
                                    if ((!(42 < 4294967295)) && true) {
                                        for fi152 := 0; fi152 < 6; fi152++ {
                                            if (115 != v112) {
                                                if (-424 == 1) {
                                                    v92 = v6
                                                }
                                            } else {
                                                _ = len(v3)
                                                v6 += (v6 - (v92 | 0))
                                                var v153 float32 = 0.1
                                                _ = v153
                                            }
                                            var v154 uint16 = 989
                                            var v155 uint16 = 32767
                                            _ = fi152
                                            _ = v154
                                            _ = v155
                                        }
                                    }
                                } else {
                                    for fi156 := 0; fi156 < 14; fi156++ {
                                        v157 := GenS17{F0: 0}
                                        if true {
                                            _ = len(v3)
                                            var v158 string = "hello"
                                            var v159 uint8 = (v4 & (v4 % (((v1 >> ((628) & 63))) | 1)))
                                            _ = v158
                                            _ = v159
                                        }
                                        v5 = append(v5, 4294967295)
                                        _ = fi156
                                        _ = v157
                                    }
                                }
                                var v160 uint = 42
                                _ = fi150
                                _ = v151
                                _ = v160
                            }
                            var v161 int32 = 267
                            if (!false) {
                                var v162 int8 = int8(v161)
                                v6 += -864
                                _ = v162
                            }
                            _ = v161
                        } else {
                            var v163 uint64 = 42
                            for fi164 := 0; fi164 < 12; fi164++ {
                                v165 := GenS18{F0: 255, F1: v4}
                                _ = fi164
                                _ = v165
                            }
                            _ = v163
                        }
                    }
                    v92 = v92
                }
            }
        }
    }
    if true {
        for fi166 := 0; fi166 < 6; fi166++ {
            var v167 uint8 = (v1 % (((v1 * (v4 ^ 128))) | 1))
            var v168 int32 = -2147483648
            _ = fi166
            _ = v167
            _ = v168
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        switch "" {
        case "\n\t":
            for fi169 := 0; fi169 < 14; fi169++ {
                for fi170 := 0; fi170 < 4; fi170++ {
                    if true {
                        fmt.Println(fi169)
                        v171 := GenS19{F0: 89, F1: 42, F2: 0}
                        _ = v171
                    } else {
                        v92 -= 42
                        if true {
                            var v172 uint = 1
                            var v173 uint = (^(v172))
                            _ = v172
                            _ = v173
                        } else {
                            v1 = v4
                            fmt.Println(v2)
                        }
                        var v174 int16 = 1
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(v1, fi170, fi169)
                        }
                        _ = v174
                    }
                    if (!((2 <= 481) && false)) {
                        var v175 int64 = (v92 << ((18446744073709551615) & 63))
                        v176 := GenS20{F0: 42, F1: 955}
                        v6 -= (v6 - v92)
                        _ = v175
                        _ = v176
                    } else {
                        if true {
                            _ = len(v5)
                            v4 = v1
                            v112 += v112
                            if false {
                                if (0 == 0) {
                                    v177 := GenS13{F0: v112, F1: "foo bar", F2: -311}
                                    var v178 bool = true
                                    _ = v177
                                    _ = v178
                                } else {
                                    fmt.Println(fi170, v4, v92)
                                }
                                var v179 float32 = (3.14159 - (0.1 - float32(v2)))
                                v180 := GenS14{F0: 0, F1: 42}
                                _ = v179
                                _ = v180
                            } else {
                                if len(v5) > 0 {
                                    v5[0] = 1
                                }
                                var v181 uint64 = (v112 & (v112 + v112))
                                _ = v181
                            }
                            fmt.Println(v112, v112, v92)
                        }
                        fmt.Println(v92)
                    }
                    _ = fi170
                }
                v6 = (v92 & v92)
                _ = fi169
            }
        case "a":
            fmt.Println(v1, v2)
        case "hello":
            v182 := GenS21{F0: false}
            _ = v182
        }
    }
    for fi183 := 0; fi183 < 16; fi183++ {
        if true {
            fmt.Println(v2)
            fmt.Println(fi183)
            v4 = v1
        }
        v5 = append(v5, 0)
        v92 -= 0
        _ = fi183
    }
    if (1 <= int8(v1)) {
        v112 = (v112 << ((65535) & 63))
    } else {
        v184 := GenS22{F0: -2147483648, F1: 223}
        for fi185 := 0; fi185 < 8; fi185++ {
            v4 = (^(v4))
            _ = fi185
        }
        _ = v184
    }
    switch 127 {
    case 9223372036854775807:
        switch "" {
        case "\n\t":
            if (v6 > (-(v6))) {
                switch 981 {
                case 2147483647:
                    for fi186 := 0; fi186 < 11; fi186++ {
                        v187 := v3["a"]
                        if (0 == 39) {
                            if (!true) {
                                v4 = (v4 * (v1 >> ((65535) & 63)))
                            }
                            var v188 int64 = v92
                            _ = len(v5)
                            _ = v188
                        }
                        v92 = v6
                        _ = fi186
                        _ = v187
                    }
                default:
                    switch (("a" + "") + "") {
                    case "a":
                        v189 := GenS16{F0: 694, F1: 2147483647, F2: 719}
                        _ = v189
                    default:
                        fmt.Println(v6, v2)
                    }
                }
            } else {
                var v190 int16 = -897
                _ = v190
            }
        case "a":
            switch -559 {
            case -234:
                v4 -= 42
            case -287:
                switch "hello" {
                case "\n\t":
                    var v191 int = -2147483648
                    _ = v191
                case "foo bar":
                    v2 -= (((float64(0.1) - (((v2 / v2) - (v2 + v2)) + (((v2 * (3.14159 * ((0.0 * v2) / (0.1 / v2)))) / 0.1) - (float64(0.1) * 0.0)))) / -2.25) / v2)
                case "":
                    fmt.Println(v6)
                default:
                    if true {
                        fmt.Println(v112)
                        v192 := GenS14{F0: 4294967295, F1: 407}
                        _ = v192
                    }
                }
            default:
                if len(v5) > 0 {
                    v5[0] = 42
                }
            }
        case "":
            v193 := GenS3{F0: 742, F1: v112, F2: 32767}
            _ = v193
        default:
            for fi194 := 0; fi194 < 20; fi194++ {
                v2 += (v2 / (0.1 + v2))
                _ = fi194
            }
        }
    case 32768:
        v6 = v6
    case 1:
        v6 = v6
    default:
        for fi195 := 0; fi195 < 19; fi195++ {
            for fi196 := 0; fi196 < 18; fi196++ {
                var v197 bool = false
                _ = fi196
                _ = v197
            }
            if true {
                v1 -= v4
            }
            _ = fi195
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v112:", v112)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v112
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v92
}
