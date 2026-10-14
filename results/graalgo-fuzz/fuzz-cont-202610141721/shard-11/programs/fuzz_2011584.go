package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32; F1 bool }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint16; F1 int16 }
type GenS13 struct { F0 int8; F1 int16; F2 int32 }
func genFn0() bool {
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(fi6)
        var v7 uint8 = 1
        switch "\n\t" {
        case "":
            for fi8 := 0; fi8 < 11; fi8++ {
                switch "foo bar" {
                case "a":
                    var v9 int16 = -1
                    _ = v9
                case "hello":
                    var v10 string = "foo bar"
                    _ = v10
                case "":
                    if (400 >= 583) {
                        v11 := []int{fi8, fi8, (fi6 * fi8), (fi6 % ((984) | 1))}
                        if len(v11) > 0 {
                            v11[0] = fi8
                        }
                        if (true || true) {
                            v11 = append(v11, 316)
                        } else {
                            var v12 uint = 2
                            v11 = append(v11, fi6)
                            switch fi8 {
                            case 32767:
                                v13 := GenS1{F0: 2147483647, F1: 32767, F2: false}
                                _ = v13
                            case -9223372036854775808:
                                switch (-(fi8)) {
                                case 2147483647:
                                    fmt.Println(fi6)
                                    v7 = v7
                                    if true {
                                        v14 := GenS1{F0: 4294967295, F1: 32767, F2: true}
                                        _ = v14
                                    }
                                    switch fi8 {
                                    case -2147483648:
                                        v12 -= v12
                                    case -900:
                                        var v15 float32 = 3.14159
                                        _ = v15
                                    case 32768:
                                        fmt.Println(v7, fi8, fi6)
                                    default:
                                        v7 = (v7 << (((v12 & 18446744073709551615)) & 63))
                                    }
                                case 32767:
                                    v12 -= v12
                                case -62:
                                    v12 += 657
                                }
                            case -304:
                                v7 += v7
                            default:
                                if len(v11) > 0 {
                                    v11[0] = (fi8 / (((^(fi8))) | 1))
                                }
                            }
                            _ = v12
                        }
                        _ = v11
                    } else {
                        v16 := []int{(fi6 >> ((241) & 63)), (-(fi6))}
                        _ = v16
                    }
                }
                _ = fi8
            }
        case "foo bar":
            var v17 uint8 = 1
            _ = v17
        case "hello":
            switch "foo bar" {
            case "":
                var v18 uint8 = 2
                _ = v18
            case "\n\t":
                switch fi6 {
                case -262:
                    v19 := []int{fi6, 369, (fi6 % ((fi6) | 1))}
                    _ = v19
                case 32768:
                    v20 := []int{}
                    _ = v20
                }
            }
        default:
            v7 = 255
        }
        _ = fi6
        _ = v7
    }
    switch 493 {
    case -137:
        v21 := []int{127, -1, 1}
        _ = v21
    case 128:
        if true {
            switch 127 {
            case 9223372036854775807:
                var v22 int64 = -295
                _ = v22
            default:
                for fi23 := 0; fi23 < 6; fi23++ {
                    var v24 uint8 = 128
                    var v25 string = "hello"
                    v26 := []int{(fi23 - fi23), 9223372036854775807}
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                }
            }
        } else {
            v27 := []int{32768}
            _ = v27
        }
    default:
        var v28 uint32 = 65535
        _ = v28
    }
    return false
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    var v3 int32 = -1
    v4 := make(map[string]bool)
    v5 := make(map[string]bool)
    v5[("foo bar" + "\n\t")] = (255 <= 2)
    v3 = 1
    switch "hello" {
    case "foo bar":
        fmt.Println(v3)
    case "a":
        v3 = v3
    case "":
        for fi29 := 0; fi29 < 14; fi29++ {
            for fi30 := 0; fi30 < 5; fi30++ {
                switch fi30 {
                case 9223372036854775807:
                    v31 := []int{}
                    _ = v31
                case 32767:
                    v32 := GenS1{F0: 65535, F1: 0, F2: true}
                    _ = v32
                default:
                    var v33 bool = true
                    _ = v33
                }
                _ = fi30
            }
            _ = fi29
        }
    default:
        if false {
            var v34 float32 = (float32((1.5 + (0.1 + (float64(float32(1e10)) * 0.0)))) - 1e10)
            for fi35 := 0; fi35 < 13; fi35++ {
                v3 -= -2147483648
                v3 -= v3
                if (true || false) {
                    if true {
                        if true {
                            if (!(true || (false && (-297 != (fi35 >> ((42) & 63)))))) {
                                var v36 uint16 = 32767
                                fmt.Println(v34)
                                _ = v36
                            }
                        } else {
                            var v37 uint64 = 0
                            _ = v37
                        }
                    }
                }
                _ = fi35
            }
            v38 := GenS2{F0: 32767}
            _ = v34
            _ = v38
        }
    }
    _ = len(v5)
    v39 := GenS1{F0: 42, F1: 1, F2: false}
    if false {
        v3 += (v3 ^ 2147483647)
        var v40 uint16 = 42
        if false {
            if true {
                for fi41 := 0; fi41 < 16; fi41++ {
                    v40 = v40
                    v42 := []int{0, (fi41 / (((fi41 * 384)) | 1)), (fi41 & fi41), (-(fi41))}
                    _ = fi41
                    _ = v42
                }
                var v43 int32 = v3
                _ = v43
            } else {
                switch 127 {
                case -2147483648:
                    switch (("a" + "a") + "foo bar") {
                    case "hello":
                        v44 := []int{32768, 294, 42, -1}
                        _ = v44
                    case "":
                        var v45 float32 = 0.1
                        _ = v45
                    case "foo bar":
                        var v46 int = 0
                        _ = v46
                    }
                case 128:
                    v47 := GenS3{F0: "\n\t", F1: -128}
                    _ = v47
                case 0:
                    if (1e10 == 1.5) {
                        switch -2147483648 {
                        case 1:
                            for fi48 := 0; fi48 < 13; fi48++ {
                                v3 -= (v3 * 0)
                                v40 += (v40 << ((0) & 63))
                                _ = fi48
                            }
                        case -1:
                            v3 = v3
                        default:
                            v49 := GenS4{F0: v40}
                            _ = v49
                        }
                        for fi50 := 0; fi50 < 16; fi50++ {
                            v40 = (v40 % (((^(v40))) | 1))
                            if (1.5 != (3.14159 / float32(((3.14159 * 1.5) - 3.14159)))) {
                                v40 += (v40 + (v40 >> ((42) & 63)))
                                v51 := []int{(fi50 >> ((18446744073709551615) & 63)), 710}
                                fmt.Println(v40, v40, v3)
                                _ = v51
                            }
                            _ = fi50
                        }
                    } else {
                        var v52 uint32 = 65535
                        switch 2147483647 {
                        case 127:
                            switch 128 {
                            case 1:
                                v53 := GenS4{F0: v40}
                                _ = v53
                            }
                        case 700:
                            v54 := GenS1{F0: v52, F1: (v40 + 253), F2: true}
                            _ = v54
                        case -2147483648:
                            v52 = v52
                        default:
                            if (false || (4 >= 42)) {
                                v55 := []int{}
                                fmt.Println(v3, v40)
                                _ = v55
                            } else {
                                for fi56 := 0; fi56 < 5; fi56++ {
                                    if false {
                                        v40 = (v40 + v40)
                                        v3 = 42
                                        v57 := []int{(^(fi56)), 9223372036854775807}
                                        _ = v57
                                    } else {
                                        for fi58 := 0; fi58 < 9; fi58++ {
                                            switch fi58 {
                                            case -7:
                                                v59 := v5[("" + "foo bar")]
                                                _ = v59
                                            case 0:
                                                if true {
                                                    var v60 int8 = int8(v3)
                                                    v3 = v3
                                                    _ = v60
                                                }
                                            default:
                                                var v61 bool = false
                                                _ = v61
                                            }
                                            if (false && ((((1000 == -32768) || true) || false) || false)) {
                                                if (304 > 638) {
                                                    v1[-60] = (("foo bar" + "") + "\n\t")
                                                }
                                            }
                                            var v62 float32 = 3.14159
                                            _ = fi58
                                            _ = v62
                                        }
                                        if true {
                                            v40 = v40
                                            if ((((("" + "a") + ("\n\t" + (("hello" + ("hello" + (("hello" + "hello") + "foo bar"))) + "foo bar"))) + "a") + "hello") == "\n\t") {
                                                v52 += (v52 >> ((4294967295) & 63))
                                                if (v3 <= (v3 | 512)) {
                                                    v63 := GenS5{F0: (v40 << ((65535) & 63)), F1: false}
                                                    _ = v63
                                                }
                                                if false {
                                                    for fi64 := 0; fi64 < 7; fi64++ {
                                                        switch (fi56 + 607) {
                                                        case 4294967295:
                                                            v65 := []int{(fi64 | int(v40)), int(v52)}
                                                            _ = v65
                                                        case 2147483647:
                                                            if false {
                                                                if (true && ((1.5 != ((float64(0.0) / ((float64(0.1) - 3.14159) + 0.1)) - (0.0 + -2.25))) && true)) {
                                                                    switch ("a" + (("foo bar" + "a") + "a")) {
                                                                    case "a":
                                                                        if false {
                                                                            if (v52 >= (v52 ^ (v52 | v52))) {
                                                                                if false {
                                                                                    var v66 uint32 = v52
                                                                                    v66 = v66
                                                                                    var v67 int32 = 188
                                                                                    _ = v66
                                                                                    _ = v67
                                                                                }
                                                                                var v68 uint64 = 2
                                                                                v40 -= 65535
                                                                                _ = v68
                                                                            }
                                                                        }
                                                                    default:
                                                                        v52 -= v52
                                                                    }
                                                                    v69 := GenS6{F0: true}
                                                                    var v70 string = ""
                                                                    var v71 uint8 = 255
                                                                    _ = v69
                                                                    _ = v70
                                                                    _ = v71
                                                                }
                                                            } else {
                                                                var v72 int32 = (v3 / ((1) | 1))
                                                                if (fi64 == (-(fi64))) {
                                                                    if (1000 > 32767) {
                                                                        v40 = 42
                                                                    }
                                                                }
                                                                if (v40 < v40) {
                                                                    if (!true) {
                                                                        var v73 uint8 = 0
                                                                        v73 = v73
                                                                        _ = v73
                                                                    }
                                                                }
                                                                delete(v5, "a")
                                                                _ = len(v5)
                                                                _ = v72
                                                            }
                                                        default:
                                                            v40 += v40
                                                        }
                                                        if (!true) {
                                                            v40 = v40
                                                            v74 := []int{}
                                                            v52 -= v52
                                                            _ = v74
                                                        }
                                                        var v75 uint16 = v40
                                                        _ = fi64
                                                        _ = v75
                                                    }
                                                    v76 := GenS2{F0: fi56}
                                                    v40 -= v40
                                                    _ = v76
                                                } else {
                                                    var v77 int = fi56
                                                    _ = v77
                                                }
                                            }
                                        } else {
                                            v78 := GenS7{F0: 125}
                                            for fi79 := 0; fi79 < 19; fi79++ {
                                                v80 := GenS4{F0: (v40 >> ((65535) & 63))}
                                                v52 = (v52 % ((v52) | 1))
                                                v81 := GenS1{F0: (v52 % ((v52) | 1)), F1: (v40 / (((v40 * 32767)) | 1)), F2: (v3 == v3)}
                                                _ = fi79
                                                _ = v80
                                                _ = v81
                                            }
                                            _ = v78
                                        }
                                    }
                                    _ = fi56
                                }
                            }
                        }
                        _ = v52
                    }
                }
                fmt.Println(v40, v40)
                for fi82 := 0; fi82 < 20; fi82++ {
                    switch (fi82 * (^(fi82))) {
                    case 212:
                        if true {
                            v83 := GenS8{F0: (v3 % ((v3) | 1)), F1: (103 > 578)}
                            _ = v83
                        } else {
                            for fi84 := 0; fi84 < 20; fi84++ {
                                v40 -= v40
                                var v85 uint = 65535
                                v86 := GenS4{F0: (v40 % (((v40 % (((^(v40))) | 1))) | 1))}
                                _ = fi84
                                _ = v85
                                _ = v86
                            }
                            v40 -= 404
                        }
                    case 873:
                        v40 = 1
                    case 127:
                        var v87 int8 = -115
                        _ = v87
                    default:
                        var v88 string = "\n\t"
                        _ = v88
                    }
                    _ = fi82
                }
            }
        }
        _ = v40
    } else {
        for fi89 := 0; fi89 < 14; fi89++ {
            v3 += (v3 << ((640) & 63))
            v90 := GenS4{F0: 988}
            _ = len(v4)
            _ = fi89
            _ = v90
        }
        v91 := []int{127, 266, 127}
        _ = v91
    }
    v92 := GenS1{F0: 2147483647, F1: 546, F2: (2 != 4294967295)}
    for fi93 := 0; fi93 < 14; fi93++ {
        for fi94 := 0; fi94 < 17; fi94++ {
            v95 := v5[("foo bar" + "")]
            _ = fi94
            _ = v95
        }
        var v96 uint32 = 42
        _ = fi93
        _ = v96
    }
    fmt.Println(v3)
    v3 = v3
    if (("a" + "\n\t") == ("" + "hello")) {
        var v97 int8 = -128
        v98 := GenS8{F0: (v3 << ((255) & 63)), F1: (true || true)}
        var v99 int16 = 955
        _ = v97
        _ = v98
        _ = v99
    }
    v100 := GenS5{F0: 65535, F1: (-114 != -1)}
    switch "" {
    case "":
        v101 := GenS3{F0: "", F1: -1}
        _ = v101
    case "a":
        for fi102 := 0; fi102 < 1; fi102++ {
            if true {
                _ = len(v4)
            } else {
                if true {
                    for fi103 := 0; fi103 < 8; fi103++ {
                        v3 = -1
                        v104 := v4[(("a" + "hello") + ("foo bar" + ("a" + "a")))]
                        _ = fi103
                        _ = v104
                    }
                    fmt.Println(fi102, fi102)
                    v105 := v1[fi102]
                    _ = v105
                }
                if (2 == 114) {
                    v106 := GenS5{F0: 32767, F1: true}
                    _ = len(v4)
                    fmt.Println(v3, v3)
                    _ = v106
                } else {
                    v3 = 1000
                    v3 += (v3 & (v3 - -2147483648))
                }
                for fi107 := 0; fi107 < 10; fi107++ {
                    v108 := GenS9{F0: v3}
                    v3 += v3
                    _ = fi107
                    _ = v108
                }
            }
            switch ("hello" + "a") {
            case "\n\t":
                v109 := v4["hello"]
                _ = v109
            case "":
                _ = len(v1)
            case "foo bar":
                delete(v4, "hello")
            }
            v3 = (-(v3))
            _ = fi102
        }
    case "hello":
        if false {
            if (false || false) {
                v110 := []int{}
                if (65535 != 18446744073709551615) {
                    if false {
                        v111 := GenS10{F0: 109, F1: -15, F2: 65535}
                        _ = v111
                    } else {
                        if len(v110) >= 2 {
                            v110 = v110[1:len(v110)]
                        }
                        v112 := GenS6{F0: ((v3 >= (v3 << ((0) & 63))) || true)}
                        _ = v112
                    }
                    v113 := GenS11{F0: 42}
                    _ = len(v110)
                    _ = v113
                }
                switch "a" {
                case "\n\t":
                    var v114 uint16 = 1
                    _ = v114
                case "a":
                    for fi115 := 0; fi115 < 16; fi115++ {
                        var v116 int64 = 42
                        if (127 <= 213) {
                            v3 = v3
                        }
                        v110 = append(v110, (fi115 | fi115))
                        _ = fi115
                        _ = v116
                    }
                default:
                    if true {
                        var v117 float32 = (3.14159 + (0.1 - (3.14159 + 0.0)))
                        for fi118 := 0; fi118 < 14; fi118++ {
                            for fi119 := 0; fi119 < 3; fi119++ {
                                v3 -= v3
                                fmt.Println(v3, v3)
                                _ = fi119
                            }
                            if (false && true) {
                                for fi120 := 0; fi120 < 6; fi120++ {
                                    var v121 int32 = (-(v3))
                                    if ((true || (128 < 127)) || true) {
                                        if len(v110) > 0 {
                                            v110[0] = 729
                                        }
                                    }
                                    var v122 int16 = 1000
                                    v123 := v1[fi118]
                                    if len(v110) > 0 {
                                        v110[0] = -2147483648
                                    }
                                    var v124 float64 = -2.25
                                    v3 -= (v121 & (v3 - (v121 / ((-1) | 1))))
                                    _ = fi120
                                    _ = v121
                                    _ = v122
                                    _ = v123
                                    _ = v124
                                }
                            } else {
                                v5["foo bar"] = (-104 > -1)
                                v3 += (v3 & (^(v3)))
                            }
                            _ = fi118
                        }
                        _ = v117
                    } else {
                        v5["\n\t"] = true
                        for fi125 := 0; fi125 < 12; fi125++ {
                            var v126 uint32 = 0
                            fmt.Println(v3, v126, v3)
                            v127 := GenS12{F0: 32767, F1: -1}
                            _ = fi125
                            _ = v126
                            _ = v127
                        }
                    }
                }
                _ = v110
            } else {
                v128 := GenS13{F0: 127, F1: 32767, F2: (v3 - v3)}
                var v129 float32 = 1.5
                _ = v128
                _ = v129
            }
        } else {
            fmt.Println(v3)
        }
    default:
        if (-63 < 26) {
            switch "foo bar" {
            case "a":
                var v130 float64 = 0.0
                _ = v130
            }
        } else {
            v3 = v3
        }
    }
    var v131 uint8 = 1
    v131 = 35
    var v132 int = 26
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v131:", v131)
    fmt.Println("v132:", v132)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v100
    _ = v131
    _ = v132
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v92
}
