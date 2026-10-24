package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 bool; F1 bool; F2 uint32 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 bool; F1 int32; F2 bool }
type GenS7 struct { F0 int; F1 int8; F2 int64 }
type GenS8 struct { F0 uint8; F1 uint; F2 int }
type GenS9 struct { F0 int64; F1 int8 }
type GenS10 struct { F0 int; F1 uint32 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int; F1 int64; F2 int64 }
type GenS15 struct { F0 int8; F1 uint32 }
type GenS16 struct { F0 uint8; F1 int }
type GenS17 struct { F0 uint16; F1 int64; F2 bool }

func main() {
    var v1 int64 = 40
    var v2 uint16 = 508
    var v3 bool = (0 <= 127)
    var v4 int32 = 352
    for fi5 := 0; fi5 < 20; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            if true {
                var v7 bool = v3
                if true {
                    v3 = false
                    switch ("" + ("a" + ("a" + "foo bar"))) {
                    case "hello":
                        switch (("hello" + "") + "\n\t") {
                        case "\n\t":
                            if (!((v1 + -9223372036854775808) > 0)) {
                                v8 := GenS1{F0: 128, F1: v1}
                                switch 42 {
                                case 127:
                                    if ("foo bar" >= "") {
                                        var v9 int8 = 100
                                        v10 := []int{fi5, (fi6 ^ fi5), (fi6 + (fi5 / (((fi6 | (fi6 << ((807) & 63)))) | 1)))}
                                        _ = v10
                                        _ = v9
                                    }
                                }
                                var v11 float32 = (0.0 / (0.0 + ((-2.25 + (3.14159 * 0.0)) / 3.14159)))
                                v12 := GenS1{F0: 14, F1: 1}
                                v13 := []int{int(v2), fi6, -423}
                                if v3 {
                                    fmt.Println(v7, v11)
                                } else {
                                    fmt.Println(v3, fi6)
                                    switch -1 {
                                    case -310:
                                        if len(v13) > 0 {
                                            v13[0] = fi5
                                        }
                                    case -2147483648:
                                        v2 -= v2
                                    default:
                                        var v14 int64 = v1
                                        _ = v14
                                    }
                                    var v15 int8 = 0
                                    _ = v15
                                }
                                _ = v11
                                _ = v12
                                _ = v13
                                _ = v8
                            }
                        case "hello":
                            if ((!(-2.25 <= (((1e10 + -2.25) + 1e10) - 3.14159))) || v7) {
                                var v16 int16 = 849
                                v17 := []int{}
                                v7 = (v7 && false)
                                if v3 {
                                    if len(v17) > 0 {
                                        v17[0] = 128
                                    }
                                    var v18 int8 = 100
                                    v3 = (!(((-2.25 / float32(1e10)) + (3.14159 - float32((0.0 + 0.1)))) != (((0.1 / 0.1) + (0.0 * 1e10)) + (1e10 / (1e10 + -2.25)))))
                                    _ = v18
                                }
                                if true {
                                    if v3 {
                                        v2 -= 1
                                        v19 := GenS1{F0: 255, F1: (v1 >> ((255) & 63))}
                                        if (75 > -96) {
                                            v20 := GenS2{F0: (^(v16))}
                                            v7 = (v3 || true)
                                            _ = v20
                                        }
                                        _ = v19
                                    }
                                    var v21 int16 = v16
                                    switch (fi5 ^ (fi6 & fi5)) {
                                    case 42:
                                        if len(v17) > 0 {
                                            v17[0] = 42
                                        }
                                    default:
                                        if v3 {
                                            if (v3 || (1 >= 0)) {
                                                v16 = v21
                                            }
                                            v2 = v2
                                            fmt.Println(v3)
                                        } else {
                                            v22 := GenS2{F0: v21}
                                            var v23 int = (fi5 << ((68) & 63))
                                            _ = v22
                                            _ = v23
                                        }
                                    }
                                    _ = v21
                                } else {
                                    var v24 string = ((("a" + ("hello" + "hello")) + ("\n\t" + "foo bar")) + ("foo bar" + "hello"))
                                    var v25 string = ((("" + (v24 + (v24 + v24))) + ("a" + "a")) + ("foo bar" + "foo bar"))
                                    _ = v24
                                    _ = v25
                                }
                                _ = v16
                                _ = v17
                            } else {
                                var v26 uint16 = v2
                                switch "a" {
                                case "":
                                    v27 := GenS3{F0: 2, F1: 0}
                                    _ = v27
                                case "foo bar":
                                    v28 := GenS4{F0: (fi6 < (fi6 | fi6)), F1: false, F2: 2147483648}
                                    _ = v28
                                default:
                                    fmt.Println(fi5, fi6, v3)
                                }
                                v26 = v26
                                _ = v26
                            }
                        case "":
                            v2 = 909
                        }
                    }
                } else {
                    v3 = v3
                    v4 = v4
                }
                switch "a" {
                case "hello":
                    var v29 float64 = float64(((float32(1e10) - 3.14159) / -2.25))
                    _ = v29
                case "\n\t":
                    fmt.Println(v7)
                default:
                    if (true || true) {
                        switch (^(fi5)) {
                        case 32768:
                            v30 := GenS5{F0: "hello", F1: "\n\t"}
                            _ = v30
                        default:
                            v31 := GenS3{F0: 277, F1: 1000}
                            _ = v31
                        }
                        v1 += v1
                        fmt.Println(v4)
                    } else {
                        fmt.Println(v7, fi5)
                        var v32 uint32 = 42
                        _ = v32
                    }
                }
                _ = v7
            }
            switch ("foo bar" + "foo bar") {
            case "":
                if (!v3) {
                    var v33 bool = ((v4 | 2147483647) <= -388)
                    if v33 {
                        if v33 {
                            var v34 float32 = float32(1e10)
                            _ = v34
                        }
                        fmt.Println(v4, v4, v33)
                    }
                    var v35 bool = v3
                    if (v4 >= v4) {
                        fmt.Println(v4)
                        v1 = (-(v1))
                        fmt.Println(v4, v3, fi6)
                    }
                    if true {
                        if false {
                            switch fi5 {
                            case 617:
                                v3 = (18446744073709551615 < 18446744073709551615)
                            case 0:
                                switch (fi5 - fi5) {
                                case 4294967295:
                                    var v36 uint32 = 104
                                    v33 = (65535 >= 765)
                                    _ = v36
                                case 379:
                                    var v37 int64 = -747
                                    _ = v37
                                default:
                                    v3 = (uint8(v2) == uint8(v2))
                                    var v38 uint64 = 1
                                    v2 = (v2 % ((v2) | 1))
                                    _ = v38
                                }
                            case -1:
                                if v33 {
                                    if (v3 && true) {
                                        v39 := GenS1{F0: 128, F1: (v1 & (v1 >> ((522) & 63)))}
                                        var v40 float32 = (((3.14159 / 1.0) + 0.1) / 0.1)
                                        _ = v39
                                        _ = v40
                                    } else {
                                        v41 := []int{805, (fi6 << ((4294967295) & 63)), (fi5 << ((18446744073709551615) & 63)), (fi6 + (fi6 / ((fi5) | 1)))}
                                        _ = v41
                                    }
                                    if v33 {
                                        if (v35 || false) {
                                            var v42 float64 = -2.25
                                            if v3 {
                                                v43 := GenS1{F0: 42, F1: 9223372036854775807}
                                                v2 = v2
                                                _ = v43
                                            } else {
                                                fmt.Println(v33, fi5, v42)
                                                v2 += v2
                                                var v44 int32 = v4
                                                if true {
                                                    var v45 uint64 = 1
                                                    _ = v45
                                                }
                                                _ = v44
                                            }
                                            _ = v42
                                        }
                                        fmt.Println(v3, fi6, v4)
                                        if true {
                                            switch "hello" {
                                            case "":
                                                switch 32767 {
                                                case 636:
                                                    v46 := GenS6{F0: (v3 || ((v1 << ((4294967295) & 63)) > v1)), F1: v4, F2: (v33 && v33)}
                                                    _ = v46
                                                case -261:
                                                    if false {
                                                        v47 := GenS7{F0: 32768, F1: 100, F2: (v1 & 42)}
                                                        var v48 int8 = 69
                                                        fmt.Println(fi5)
                                                        _ = v47
                                                        _ = v48
                                                    } else {
                                                        v49 := []int{-6, 128, (fi5 & -2147483648)}
                                                        v50 := GenS3{F0: 42, F1: -32768}
                                                        _ = v49
                                                        _ = v50
                                                    }
                                                case -238:
                                                    var v51 uint32 = 290
                                                    switch int(v4) {
                                                    case 42:
                                                        v1 -= (v1 - 9223372036854775807)
                                                    }
                                                    if v3 {
                                                        switch "\n\t" {
                                                        case "a":
                                                            var v52 int16 = -191
                                                            _ = v52
                                                        case "\n\t":
                                                            var v53 uint16 = v2
                                                            _ = v53
                                                        case "foo bar":
                                                            if (0 == 42) {
                                                                fmt.Println(v2)
                                                            } else {
                                                                v54 := GenS8{F0: 0, F1: 4294967295, F2: fi6}
                                                                _ = v54
                                                            }
                                                        default:
                                                            fmt.Println(v3, v2, v3)
                                                        }
                                                        var v55 int16 = 1000
                                                        v56 := GenS2{F0: (v55 ^ v55)}
                                                        switch -9223372036854775808 {
                                                        case 1:
                                                            v33 = v33
                                                        case 87:
                                                            var v57 float64 = (3.14159 + (1e10 * (3.14159 + ((0.0 - (1e10 / (((float64((0.0 * float32(-2.25))) / 1.5) - -2.25) * -2.25))) - 1e10))))
                                                            _ = v57
                                                        default:
                                                            v58 := []int{(fi6 >> ((42) & 63)), (fi5 / ((fi5) | 1)), fi5}
                                                            _ = v58
                                                        }
                                                        switch "foo bar" {
                                                        case "":
                                                            var v59 uint = 0
                                                            _ = v59
                                                        case "hello":
                                                            fmt.Println(fi6, v35, v35)
                                                        case "\n\t":
                                                            fmt.Println(v51)
                                                        }
                                                        switch "" {
                                                        case "a":
                                                            var v60 float32 = (3.14159 * float32(0.1))
                                                            _ = v60
                                                        }
                                                        var v61 uint8 = 2
                                                        _ = v55
                                                        _ = v56
                                                        _ = v61
                                                    } else {
                                                        var v62 bool = v3
                                                        _ = v62
                                                    }
                                                    _ = v51
                                                default:
                                                    v4 += (-(v4))
                                                }
                                            case "a":
                                                fmt.Println(v33, fi6)
                                            case "foo bar":
                                                if (!false) {
                                                    fmt.Println(v4)
                                                    switch (fi5 + (fi5 % ((fi5) | 1))) {
                                                    case -979:
                                                        v3 = v33
                                                    case 9223372036854775807:
                                                        var v63 int8 = -128
                                                        _ = v63
                                                    }
                                                } else {
                                                    if false {
                                                        if false {
                                                            v64 := GenS3{F0: 771, F1: -1}
                                                            _ = v64
                                                        } else {
                                                            v65 := GenS5{F0: "hello", F1: "a"}
                                                            var v66 int64 = (v1 << ((549) & 63))
                                                            if true {
                                                                if false {
                                                                    var v67 uint32 = 355
                                                                    v68 := []int{fi6, (fi6 ^ (-(fi5))), (fi6 << ((1) & 63)), 0}
                                                                    _ = v67
                                                                    _ = v68
                                                                }
                                                                if (474 != 18446744073709551615) {
                                                                    fmt.Println(v2, v4)
                                                                }
                                                            } else {
                                                                if ((65535 < 8) || (!((!false) || v35))) {
                                                                    v35 = ((v4 % (((v4 / ((-1) | 1))) | 1)) < v4)
                                                                    if v35 {
                                                                        fmt.Println(v1, v4)
                                                                        v69 := []int{fi6, 127, fi5}
                                                                        _ = v69
                                                                    } else {
                                                                        v70 := []int{}
                                                                        _ = v70
                                                                    }
                                                                    if (4294967295 != 42) {
                                                                        var v71 uint = 51
                                                                        _ = v71
                                                                    } else {
                                                                        v72 := []int{(fi5 ^ (fi5 - -2147483648))}
                                                                        if len(v72) > 0 {
                                                                            v72[0] = fi5
                                                                        }
                                                                        _ = v72
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        if (42 > 42) {
                                                                            v3 = true
                                                                            v4 = 2147483647
                                                                        }
                                                                        var v73 uint16 = (^(v2))
                                                                        _ = v73
                                                                    } else {
                                                                        switch (("a" + "") + "foo bar") {
                                                                        case "":
                                                                            fmt.Println(fi6, v2)
                                                                        default:
                                                                            fmt.Println(v2, v66)
                                                                        }
                                                                        v74 := GenS9{F0: v1, F1: 0}
                                                                        var v75 uint64 = 42
                                                                        _ = v74
                                                                        _ = v75
                                                                    }
                                                                    var v76 uint64 = 18446744073709551615
                                                                    var v77 uint32 = 576
                                                                    v4 -= (v4 - v4)
                                                                    if (4294967295 > 4294967295) {
                                                                        var v78 string = ("a" + "\n\t")
                                                                        var v79 int8 = -99
                                                                        var v80 uint32 = (v77 / ((v77) | 1))
                                                                        v81 := []int{fi6}
                                                                        fmt.Println(fi6, v2)
                                                                        _ = v78
                                                                        _ = v79
                                                                        _ = v80
                                                                        _ = v81
                                                                    } else {
                                                                        switch (-(fi6)) {
                                                                        case 9223372036854775807:
                                                                            var v82 uint8 = 1
                                                                            _ = v82
                                                                        case -1:
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    v83 := []int{}
                                                                                    _ = v83
                                                                                case "a":
                                                                                    var v84 uint8 = 210
                                                                                    _ = v84
                                                                                }
                                                                            }
                                                                        case 42:
                                                                            v85 := []int{(fi5 - fi6), (fi6 ^ (fi5 - fi6)), fi6, 127}
                                                                            _ = v85
                                                                        }
                                                                        v76 = 39
                                                                        v86 := GenS3{F0: 698, F1: 1000}
                                                                        v1 = v66
                                                                        _ = v86
                                                                    }
                                                                    if ((float64(0.1) * -2.25) > (0.0 - float64((float32(3.14159) / -2.25)))) {
                                                                        if true {
                                                                            if v35 {
                                                                                if false {
                                                                                    fmt.Println(v3, v35, v1)
                                                                                    fmt.Println(v35, v77, v33)
                                                                                }
                                                                                v66 = 9223372036854775807
                                                                            }
                                                                            v87 := []int{fi5, 42, fi5, 9223372036854775807}
                                                                            var v88 uint16 = (v2 % ((1) | 1))
                                                                            if len(v87) > 0 {
                                                                                v87[0] = fi5
                                                                            }
                                                                            _ = v87
                                                                            _ = v88
                                                                        }
                                                                        var v89 string = "a"
                                                                        _ = v89
                                                                    }
                                                                    _ = v76
                                                                    _ = v77
                                                                }
                                                                v90 := GenS10{F0: (fi6 % (((fi5 ^ 4294967295)) | 1)), F1: 0}
                                                                var v91 float32 = 3.14159
                                                                _ = v90
                                                                _ = v91
                                                            }
                                                            fmt.Println(v33, v33)
                                                            _ = v65
                                                            _ = v66
                                                        }
                                                    }
                                                    switch "hello" {
                                                    case "hello":
                                                        v1 += (-(v1))
                                                    case "a":
                                                        v92 := GenS11{F0: 97}
                                                        _ = v92
                                                    case "foo bar":
                                                        v35 = v33
                                                    }
                                                }
                                                v35 = ((1e10 + -2.25) >= -2.25)
                                            }
                                        }
                                    }
                                } else {
                                    v93 := GenS12{F0: 2147483647}
                                    _ = v93
                                }
                            default:
                                v94 := []int{}
                                _ = v94
                            }
                            var v95 uint32 = 4294967295
                            _ = v95
                        } else {
                            var v96 int8 = 100
                            var v97 int = fi5
                            var v98 uint64 = 1
                            _ = v96
                            _ = v97
                            _ = v98
                        }
                        var v99 uint16 = 912
                        v100 := []int{2147483647}
                        _ = v100
                        _ = v99
                    } else {
                        var v101 uint = 612
                        if true {
                            v102 := []int{}
                            if ("\n\t" != "foo bar") {
                                v3 = v3
                                var v103 int16 = 156
                                switch ("a" + "\n\t") {
                                case "\n\t":
                                    v104 := GenS4{F0: (float64(((0.0 * ((1.5 + 0.1) * 1.5)) + (float32(1e10) / float32(1e10)))) != 3.14159), F1: true, F2: 4294967295}
                                    _ = v104
                                default:
                                    v3 = ((0 == (v103 * (v103 >> ((v101) & 63)))) && true)
                                    v105 := GenS4{F0: v33, F1: (v33 || v35), F2: 2147483648}
                                    _ = v105
                                }
                                if len(v102) > 0 {
                                    v102[0] = fi6
                                }
                                _ = v103
                            }
                            v3 = ((1 >= 247) && (-32768 > -497))
                            _ = v102
                        } else {
                            v106 := []int{(fi6 * fi6), (-(fi5)), (fi6 % ((fi5) | 1))}
                            _ = v106
                        }
                        v1 = -9223372036854775808
                        _ = v101
                    }
                    _ = v33
                    _ = v35
                }
            case "a":
                v107 := []int{-74, (fi6 * (fi6 >> ((445) & 63))), (fi5 * fi6), (fi6 * fi5)}
                _ = v107
            case "hello":
                var v108 uint32 = 0
                _ = v108
            }
            v109 := []int{(fi5 & (fi5 - (-(fi5)))), 32767}
            _ = fi6
            _ = v109
        }
        _ = fi5
    }
    if (18446744073709551615 != 18446744073709551615) {
        if v3 {
            if (42 == 42) {
                v110 := GenS12{F0: -93}
                v111 := []int{-927, -997, 9223372036854775807}
                _ = v110
                _ = v111
            }
            v112 := GenS7{F0: 356, F1: 100, F2: (v1 / ((v1) | 1))}
            _ = v112
        } else {
            switch (("hello" + "") + "") {
            case "hello":
                v1 = (-(v1))
            default:
                switch "a" {
                case "\n\t":
                    var v113 float32 = 0.0
                    _ = v113
                case "hello":
                    fmt.Println(v3, v4, v3)
                }
            }
        }
        switch ("hello" + ("foo bar" + ("foo bar" + "\n\t"))) {
        case "a":
            v4 = 42
        }
    } else {
        var v114 float64 = 3.14159
        for fi115 := 0; fi115 < 3; fi115++ {
            fmt.Println(v2, fi115)
            v114 = 3.14159
            var v116 int32 = v4
            _ = fi115
            _ = v116
        }
        _ = v114
    }
    v117 := GenS13{F0: v2}
    if v3 {
        v3 = true
    }
    for fi118 := 0; fi118 < 12; fi118++ {
        switch fi118 {
        case 127:
            v119 := GenS14{F0: (^(fi118)), F1: (v1 % ((v1) | 1)), F2: (-(v1))}
            _ = v119
        case 9223372036854775807:
            v120 := []int{(fi118 + (^(fi118))), fi118, 32768, fi118}
            _ = v120
        case 234:
            v4 += v4
        }
        _ = fi118
    }
    var v121 int32 = v4
    v122 := GenS15{F0: 10, F1: 2}
    for fi123 := 0; fi123 < 1; fi123++ {
        if (!(uint64(v2) >= 42)) {
            if v3 {
                for fi124 := 0; fi124 < 8; fi124++ {
                    var v125 string = "\n\t"
                    v2 = v2
                    var v126 uint = 861
                    _ = fi124
                    _ = v125
                    _ = v126
                }
            }
            if ((!(-463 >= (fi123 * (fi123 % ((fi123) | 1))))) || v3) {
                var v127 uint64 = 1
                _ = v127
            }
        } else {
            v128 := []int{(fi123 - (fi123 ^ (fi123 * (fi123 % ((fi123) | 1))))), fi123, fi123}
            _ = v128
        }
        v129 := []int{(fi123 / (((fi123 - fi123)) | 1)), fi123}
        _ = fi123
        _ = v129
    }
    v1 -= (v1 & v1)
    for fi130 := 0; fi130 < 8; fi130++ {
        if (v3 || ((((("" + ("hello" + ("hello" + "a"))) + "a") + ("a" + ("a" + "foo bar"))) + ("" + "\n\t")) != "foo bar")) {
            v3 = v3
            for fi131 := 0; fi131 < 1; fi131++ {
                v121 = (-(v4))
                v121 = -2147483648
                _ = fi131
            }
        }
        _ = fi130
    }
    fmt.Println(v2, v1, v4)
    switch 73 {
    case -258:
        fmt.Println(v2, v3)
    case -1:
        v132 := GenS1{F0: 127, F1: (v1 + (v1 << ((224) & 63)))}
        _ = v132
    }
    if false {
        fmt.Println(v1, v4)
        v133 := GenS16{F0: 13, F1: 195}
        _ = v133
    } else {
        var v134 uint64 = 4294967295
        if v3 {
            var v135 string = "a"
            if (1 <= 1) {
                fmt.Println(v3, v3)
            }
            _ = v135
        } else {
            if (v3 || (9 == 0)) {
                v2 -= 65535
            } else {
                var v136 int32 = -803
                _ = v136
            }
        }
        for fi137 := 0; fi137 < 8; fi137++ {
            v138 := GenS17{F0: (v2 << ((42) & 63)), F1: v1, F2: true}
            v139 := GenS5{F0: "\n\t", F1: ((("\n\t" + ("" + "a")) + ("hello" + "a")) + "")}
            switch "a" {
            case "hello":
                for fi140 := 0; fi140 < 4; fi140++ {
                    var v141 string = "hello"
                    fmt.Println(v1, v2, v1)
                    v142 := []int{(fi137 >> ((255) & 63))}
                    _ = fi140
                    _ = v141
                    _ = v142
                }
            }
            _ = fi137
            _ = v138
            _ = v139
        }
        _ = v134
    }
    v143 := GenS6{F0: v3, F1: (v4 | 167), F2: (204 <= 1)}
    fmt.Println("v1:", v1)
    fmt.Println("v121:", v121)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v117
    _ = v121
    _ = v122
    _ = v143
    _ = v2
    _ = v3
    _ = v4
}
