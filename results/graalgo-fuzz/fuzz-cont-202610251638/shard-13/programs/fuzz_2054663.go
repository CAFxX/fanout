package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool; F1 int64; F2 uint64 }
type GenS6 struct { F0 bool; F1 int64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int; F1 uint8 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint32; F1 int64; F2 int }
type GenS15 struct { F0 int; F1 uint }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 int64 }
type GenS18 struct { F0 uint32; F1 uint32 }
type GenS19 struct { F0 uint32 }
type GenS20 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS21 struct { F0 uint16; F1 int8 }
type GenS22 struct { F0 uint32; F1 uint16; F2 bool }
type GenS23 struct { F0 int16; F1 int64; F2 uint8 }
type GenS24 struct { F0 uint64; F1 int32; F2 uint }
type GenS25 struct { F0 int; F1 string }
type GenS26 struct { F0 uint8 }
type GenS27 struct { F0 bool; F1 string }
type GenS28 struct { F0 int64 }
type GenS29 struct { F0 uint32; F1 uint }
type GenS30 struct { F0 string; F1 uint8 }
type GenS31 struct { F0 int }
func genFn0(p0 bool) uint {
    v7 := GenS1{F0: 418, F1: 42, F2: -23}
    switch "hello" {
    case "":
        p0 = p0
    case "\n\t":
        p0 = p0
    case "hello":
        for fi8 := 0; fi8 < 11; fi8++ {
            v9 := GenS2{F0: ("\n\t" + "\n\t")}
            _ = fi8
            _ = v9
        }
    }
    if (3.14159 > -2.25) {
        v10 := []int{-820, 4294967295, 32768, 128}
        var v11 int64 = 9223372036854775807
        _ = v10
        _ = v11
    } else {
        v12 := GenS3{F0: 128}
        _ = v12
    }
    if true {
        var v13 uint32 = 0
        if p0 {
            for fi14 := 0; fi14 < 14; fi14++ {
                switch (fi14 + (fi14 | (^(fi14)))) {
                case 32768:
                    var v15 uint32 = (v13 >> ((255) & 63))
                    _ = v15
                case 32767:
                    v16 := GenS2{F0: ("a" + "foo bar")}
                    _ = v16
                default:
                    v17 := []int{(fi14 * 42), (fi14 * 2147483647), fi14, fi14}
                    _ = v17
                }
                switch "" {
                case "foo bar":
                    p0 = p0
                case "":
                    var v18 int = fi14
                    _ = v18
                case "\n\t":
                    for fi19 := 0; fi19 < 16; fi19++ {
                        var v20 uint16 = 720
                        v21 := []int{(fi19 * 32768), (fi19 / (((fi19 >> ((0) & 63))) | 1)), fi19}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
                v22 := []int{127, fi14}
                _ = fi14
                _ = v22
            }
            var v23 string = "hello"
            v24 := GenS1{F0: 42, F1: 127, F2: 118}
            _ = v23
            _ = v24
        } else {
            if true {
                v25 := []int{}
                _ = v25
            } else {
                fmt.Println(p0)
                v26 := []int{-789, 9223372036854775807, 2147483647}
                p0 = (true || (-295 <= 0))
                _ = v26
            }
            v27 := []int{4294967295, -515}
            if len(v27) >= 2 {
                v27 = v27[1:len(v27)]
            }
            _ = v27
        }
        _ = v13
    } else {
        if ((1 < 8) && (-1 > 718)) {
            p0 = p0
            for fi28 := 0; fi28 < 8; fi28++ {
                p0 = (!true)
                p0 = p0
                _ = fi28
            }
            for fi29 := 0; fi29 < 17; fi29++ {
                v30 := []int{(fi29 << ((738) & 63)), fi29}
                if p0 {
                    v31 := GenS1{F0: 1, F1: 42, F2: -90}
                    for fi32 := 0; fi32 < 12; fi32++ {
                        switch fi32 {
                        case 32767:
                            var v33 uint16 = 2
                            _ = v33
                        case 0:
                            if false {
                                var v34 uint = 255
                                if false {
                                    switch ("" + "foo bar") {
                                    case "\n\t":
                                        fmt.Println(fi29, fi29, fi32)
                                    }
                                    switch "hello" {
                                    case "a":
                                        if (-2.25 <= 1.5) {
                                            v35 := GenS1{F0: v34, F1: 100, F2: 100}
                                            _ = v35
                                        } else {
                                            fmt.Println(v34, fi29, fi29)
                                            if p0 {
                                                if (!false) {
                                                    p0 = p0
                                                }
                                                var v36 bool = true
                                                _ = v36
                                            }
                                        }
                                    }
                                    if (!(2147483647 == 2147483648)) {
                                        switch "" {
                                        case "foo bar":
                                            v30 = append(v30, 128)
                                        case "hello":
                                            var v37 uint32 = 2147483647
                                            _ = v37
                                        default:
                                            if len(v30) > 0 {
                                                v30[0] = 4294967295
                                            }
                                        }
                                        fmt.Println(fi32, fi32)
                                        switch ("a" + ("" + "a")) {
                                        case "hello":
                                            fmt.Println(fi32, p0)
                                        case "a":
                                            var v38 bool = (-2.25 == (1e10 - float64(0.0)))
                                            _ = v38
                                        }
                                    }
                                } else {
                                    switch "" {
                                    case "":
                                        p0 = (p0 || (fi32 > -705))
                                    default:
                                        v34 = (^(v34))
                                    }
                                    switch ("\n\t" + ("a" + "hello")) {
                                    case "a":
                                        fmt.Println(fi32, v34, v34)
                                    case "\n\t":
                                        var v39 float32 = -2.25
                                        _ = v39
                                    case "hello":
                                        v34 += (v34 / ((v34) | 1))
                                    default:
                                        v34 = v34
                                    }
                                    if p0 {
                                        _ = len(v30)
                                        v40 := GenS1{F0: (v34 - v34), F1: -104, F2: 127}
                                        v41 := GenS1{F0: (^(v34)), F1: -23, F2: -128}
                                        _ = v40
                                        _ = v41
                                    } else {
                                        v42 := GenS1{F0: 65535, F1: -97, F2: -1}
                                        p0 = (false || (108 > -128))
                                        _ = v42
                                    }
                                }
                                var v43 uint = v34
                                _ = v34
                                _ = v43
                            }
                            if p0 {
                                p0 = p0
                            }
                            var v44 int16 = 42
                            fmt.Println(fi29, v44, fi32)
                            _ = v44
                        case 32768:
                            p0 = p0
                        default:
                            switch "a" {
                            case "hello":
                                fmt.Println(p0, fi32, fi29)
                            case "\n\t":
                                switch ("foo bar" + ("a" + "")) {
                                case "hello":
                                    p0 = true
                                case "\n\t":
                                    p0 = p0
                                default:
                                    var v45 float32 = (-2.25 + 1.5)
                                    _ = v45
                                }
                            case "":
                                p0 = (!p0)
                            }
                        }
                        fmt.Println(fi29, p0, fi29)
                        var v46 uint = 778
                        v30 = append(v30, fi29)
                        _ = fi32
                        _ = v46
                    }
                    _ = v31
                }
                _ = fi29
                _ = v30
            }
        }
        var v47 uint64 = 2
        _ = v47
    }
    for fi48 := 0; fi48 < 17; fi48++ {
        if (100 == -107) {
            var v49 int = (fi48 & -9223372036854775808)
            _ = v49
        } else {
            switch 4294967295 {
            case -1:
                for fi50 := 0; fi50 < 13; fi50++ {
                    p0 = false
                    fmt.Println(fi48)
                    _ = fi50
                }
            default:
                var v51 int = 9223372036854775807
                _ = v51
            }
            p0 = true
            v52 := GenS4{F0: 4294967295}
            _ = v52
        }
        for fi53 := 0; fi53 < 17; fi53++ {
            fmt.Println(p0, p0)
            if p0 {
                switch -714 {
                case 100:
                    switch 945 {
                    case 160:
                        var v54 int32 = -2147483648
                        _ = v54
                    case 128:
                        fmt.Println(fi53, fi48, p0)
                    case 348:
                        var v55 uint = 900
                        _ = v55
                    }
                case 9223372036854775807:
                    if p0 {
                        switch "foo bar" {
                        case "\n\t":
                            switch "\n\t" {
                            case "foo bar":
                                if ((!("" >= "")) && (false && (float64((((0.0 * 3.14159) * (((0.0 / 1.0) - float32(((3.14159 * float64(((float32((3.14159 / 1.0)) / float32(3.14159)) - 0.0))) - 3.14159))) / (((0.1 / (1e10 - (0.1 - (0.1 - 3.14159)))) * ((-2.25 / 1.0) * 1e10)) * 1e10))) + 0.0)) >= 1e10))) {
                                    var v56 bool = false
                                    v57 := GenS5{F0: (42 > 32767), F1: 625, F2: 2}
                                    _ = v56
                                    _ = v57
                                } else {
                                    fmt.Println(fi48, fi48, p0)
                                }
                            case "\n\t":
                                p0 = p0
                            }
                        case "foo bar":
                            fmt.Println(p0)
                        }
                        var v58 string = ("a" + "")
                        _ = v58
                    }
                }
            } else {
                if true {
                    p0 = p0
                } else {
                    if p0 {
                        if (p0 || (!(323 >= 563))) {
                            var v59 float64 = 1.5
                            var v60 int64 = 1
                            var v61 uint32 = 2
                            switch "" {
                            case "":
                                switch "a" {
                                case "hello":
                                    if p0 {
                                        v62 := []int{}
                                        _ = v62
                                    } else {
                                        v63 := GenS6{F0: (2 > 128), F1: v60}
                                        var v64 int64 = (v60 * (-(v60)))
                                        _ = v63
                                        _ = v64
                                    }
                                }
                            case "a":
                                var v65 bool = false
                                _ = v65
                            default:
                                var v66 string = ("hello" + ("hello" + ("foo bar" + ("a" + "foo bar"))))
                                fmt.Println(fi53, fi48, v66)
                                p0 = p0
                                _ = v66
                            }
                            v60 = 520
                            v61 -= (v61 * v61)
                            if (p0 && false) {
                                if false {
                                    if p0 {
                                        if p0 {
                                            var v67 uint64 = 0
                                            _ = v67
                                        } else {
                                            fmt.Println(fi48, v61)
                                            var v68 bool = false
                                            v68 = ("foo bar" >= "")
                                            switch fi48 {
                                            case -10:
                                                v60 += (v60 << ((283) & 63))
                                            case -9223372036854775808:
                                                v69 := GenS7{F0: "a"}
                                                _ = v69
                                            case 231:
                                                p0 = (v60 != v60)
                                            }
                                            if ((^(v60)) > 0) {
                                                v70 := []int{}
                                                _ = v70
                                            } else {
                                                if (!false) {
                                                    v71 := GenS8{F0: fi53}
                                                    _ = v71
                                                }
                                            }
                                            _ = v68
                                        }
                                        var v72 int16 = -1
                                        _ = v72
                                    }
                                    v73 := GenS9{F0: (^(v61))}
                                    v59 = ((3.14159 - 0.0) / 1.0)
                                    fmt.Println(v61, v60)
                                    _ = v73
                                }
                                fmt.Println(v60, v60)
                                v74 := []int{728, (fi48 - (fi48 << ((65535) & 63))), (fi53 / (((-(fi48))) | 1)), (fi48 / ((int(v61)) | 1))}
                                _ = v74
                            }
                            _ = v59
                            _ = v60
                            _ = v61
                        } else {
                            switch (-(fi48)) {
                            case 128:
                                v75 := GenS5{F0: false, F1: -1, F2: 4294967295}
                                _ = v75
                            case -9223372036854775808:
                                var v76 uint16 = 255
                                _ = v76
                            case 42:
                                v77 := []int{fi48, 32768}
                                _ = v77
                            }
                            if (true && (!((0 >= 0) && ("" >= "\n\t")))) {
                                var v78 float32 = 3.14159
                                _ = v78
                            } else {
                                var v79 float32 = -2.25
                                var v80 string = "hello"
                                var v81 float64 = 1e10
                                switch (^(fi48)) {
                                case 32768:
                                    var v82 int32 = 0
                                    _ = v82
                                case 392:
                                    switch v80 {
                                    case "":
                                        var v83 float32 = 1e10
                                        _ = v83
                                    case "foo bar":
                                        v84 := []int{fi48, (fi48 / ((fi53) | 1)), 32767, fi53}
                                        _ = v84
                                    case "a":
                                        if (true && (!(float64(3.14159) != (float64(v79) - (v81 - v81))))) {
                                            switch (("\n\t" + ("a" + "a")) + "foo bar") {
                                            case "":
                                                if (p0 || false) {
                                                    v79 += v79
                                                } else {
                                                    var v85 float32 = v79
                                                    _ = v85
                                                }
                                            case "\n\t":
                                                fmt.Println(v79, v79)
                                            case "a":
                                                v86 := []int{(fi53 ^ (^(fi53))), (fi53 << ((42) & 63)), fi53}
                                                _ = v86
                                            default:
                                                v87 := []int{fi53, 128, -1, fi53}
                                                _ = v87
                                            }
                                            v81 = (v81 / ((((((v81 + v81) / v81) - (v81 / v81)) + v81) - 0.0) / (v81 / v81)))
                                        } else {
                                            v81 = (1e10 / (1.5 - (((1e10 + float64((v79 + 3.14159))) / 1.0) + (v81 * v81))))
                                            v79 -= 1e10
                                        }
                                    }
                                case 32767:
                                    switch v80 {
                                    case "hello":
                                        switch (fi48 ^ -2147483648) {
                                        case 42:
                                            fmt.Println(v79, v80)
                                        }
                                    case "foo bar":
                                        if p0 {
                                            switch "" {
                                            case "":
                                                v88 := GenS7{F0: v80}
                                                _ = v88
                                            case "\n\t":
                                                var v89 string = "hello"
                                                var v90 uint32 = 540
                                                fmt.Println(v90, v80)
                                                _ = v89
                                                _ = v90
                                            }
                                            fmt.Println(v80)
                                            v91 := []int{fi53, fi53, fi48, -2147483648}
                                            _ = v91
                                        }
                                    }
                                default:
                                    switch 4294967295 {
                                    case -588:
                                        v92 := []int{(fi53 - -9223372036854775808)}
                                        _ = v92
                                    case 0:
                                        var v93 int32 = 2147483647
                                        _ = v93
                                    }
                                }
                                if false {
                                    if p0 {
                                        v80 = v80
                                        fmt.Println(fi53, fi48)
                                    } else {
                                        var v94 int32 = 42
                                        switch (v80 + "foo bar") {
                                        case "\n\t":
                                            switch (fi48 & (-(fi53))) {
                                            case -1:
                                                var v95 int = (-(fi53))
                                                _ = v95
                                            case 1:
                                                v94 -= v94
                                            case 590:
                                                v80 = (v80 + v80)
                                            }
                                        case "":
                                            switch "foo bar" {
                                            case "a":
                                                var v96 string = (("a" + "hello") + "foo bar")
                                                _ = v96
                                            case "":
                                                fmt.Println(v79, v81)
                                            }
                                        default:
                                            v97 := GenS10{F0: 0}
                                            _ = v97
                                        }
                                        if (519 != 42) {
                                            v79 -= v79
                                        } else {
                                            v94 = 2147483647
                                        }
                                        var v98 string = v80
                                        _ = v94
                                        _ = v98
                                    }
                                    var v99 uint64 = 42
                                    v100 := []int{}
                                    _ = v100
                                    _ = v99
                                } else {
                                    v101 := []int{fi48}
                                    _ = v101
                                }
                                _ = v79
                                _ = v80
                                _ = v81
                            }
                        }
                        v102 := GenS11{F0: fi53, F1: 183}
                        if p0 {
                            if p0 {
                                if false {
                                    v103 := GenS7{F0: ("hello" + ("a" + ""))}
                                    fmt.Println(fi48)
                                    var v104 uint8 = 0
                                    _ = v103
                                    _ = v104
                                }
                                p0 = ((p0 || (42 != 0)) && true)
                            }
                        } else {
                            var v105 uint64 = 998
                            switch "" {
                            case "foo bar":
                                if (v105 != 4294967295) {
                                    fmt.Println(fi53)
                                } else {
                                    if ((0.0 * ((0.1 + 0.0) * (1.5 / (0.1 + 3.14159)))) == 1.5) {
                                        switch fi53 {
                                        case 9223372036854775807:
                                            v105 += (v105 | (^(v105)))
                                            v106 := GenS12{F0: "\n\t"}
                                            v107 := []int{(-(fi48)), fi53, (fi48 ^ 42), 32768}
                                            v108 := GenS4{F0: 2}
                                            _ = v106
                                            _ = v107
                                            _ = v108
                                        default:
                                            switch "foo bar" {
                                            case "hello":
                                                v109 := []int{(fi48 % (((-(fi48))) | 1)), fi53, -2147483648}
                                                _ = v109
                                            case "a":
                                                p0 = false
                                            case "\n\t":
                                                v110 := GenS13{F0: -1}
                                                _ = v110
                                            }
                                        }
                                        v111 := []int{}
                                        v111 = append(v111, fi53)
                                        p0 = true
                                        var v112 uint64 = (^(v105))
                                        _ = v111
                                        _ = v112
                                    } else {
                                        p0 = (42 != 127)
                                    }
                                    switch fi48 {
                                    case -9223372036854775808:
                                        var v113 bool = (fi48 < fi48)
                                        _ = v113
                                    }
                                }
                            case "":
                                fmt.Println(p0, fi48)
                            }
                            if p0 {
                                fmt.Println(p0, v105, v105)
                                v114 := GenS11{F0: (fi53 % ((fi53) | 1)), F1: 127}
                                if p0 {
                                    v105 = v105
                                }
                                var v115 uint8 = 62
                                _ = v114
                                _ = v115
                            }
                            if (p0 && p0) {
                                p0 = p0
                                if true {
                                    v116 := GenS9{F0: 112}
                                    if p0 {
                                        fmt.Println(v105, v105, v105)
                                    } else {
                                        if true {
                                            var v117 uint8 = 1
                                            p0 = p0
                                            fmt.Println(p0)
                                            _ = v117
                                        } else {
                                            p0 = p0
                                            if true {
                                                p0 = (255 > 161)
                                                v118 := []int{}
                                                _ = len(v118)
                                                _ = v118
                                            } else {
                                                v119 := GenS14{F0: 65535, F1: 9223372036854775807, F2: 4294967295}
                                                v120 := GenS15{F0: 32767, F1: 899}
                                                _ = v119
                                                _ = v120
                                            }
                                            p0 = p0
                                        }
                                        v121 := []int{}
                                        _ = v121
                                    }
                                    var v122 uint16 = 255
                                    _ = v116
                                    _ = v122
                                }
                                v123 := GenS7{F0: "a"}
                                _ = v123
                            } else {
                                fmt.Println(fi53, fi53)
                                var v124 int32 = 1
                                v125 := []int{(fi53 + fi48), fi53, fi48}
                                if p0 {
                                    if false {
                                        var v126 int = (fi48 + (fi53 << ((2) & 63)))
                                        v126 -= fi53
                                        _ = v126
                                    }
                                    if (42 >= 441) {
                                        switch ("\n\t" + ("hello" + "a")) {
                                        case "\n\t":
                                            var v127 int8 = 100
                                            _ = v127
                                        case "hello":
                                            if p0 {
                                                v128 := GenS15{F0: fi48, F1: 18446744073709551615}
                                                fmt.Println(fi53)
                                                _ = v128
                                            }
                                            p0 = (32767 < uint16(v105))
                                        default:
                                            p0 = (32767 >= 255)
                                        }
                                    } else {
                                        var v129 float64 = float64(1e10)
                                        var v130 uint8 = 1
                                        _ = v129
                                        _ = v130
                                    }
                                    if (!p0) {
                                        var v131 int8 = 87
                                        if (523 > 1) {
                                            v132 := GenS4{F0: 4294967295}
                                            fmt.Println(v124)
                                            _ = v132
                                        } else {
                                            v131 = (v131 | v131)
                                            var v133 int32 = (v124 << ((65535) & 63))
                                            _ = v133
                                        }
                                        _ = v131
                                    } else {
                                        if ((42 == 1) || (!p0)) {
                                            if true {
                                                fmt.Println(fi48)
                                                v134 := GenS2{F0: (("" + ("" + ("a" + ("a" + ("a" + "a"))))) + (("a" + "") + "a"))}
                                                fmt.Println(fi48)
                                                _ = v134
                                            }
                                            var v135 uint32 = 2
                                            _ = v135
                                        } else {
                                            var v136 bool = (p0 || (!false))
                                            if false {
                                                switch ((("a" + "a") + ("foo bar" + "\n\t")) + ("" + "foo bar")) {
                                                case "":
                                                    v137 := GenS16{F0: -32768}
                                                    _ = v137
                                                case "foo bar":
                                                    if (!(true && p0)) {
                                                        p0 = p0
                                                        switch "a" {
                                                        case "a":
                                                            fmt.Println(fi48, v105)
                                                        case "":
                                                            if len(v125) > 0 {
                                                                v125[0] = (fi48 | -1)
                                                            }
                                                        case "foo bar":
                                                            var v138 uint = 341
                                                            _ = v138
                                                        }
                                                    } else {
                                                        v139 := GenS17{F0: 0}
                                                        if false {
                                                            if p0 {
                                                                fmt.Println(v124, v105, fi53)
                                                                var v140 int = fi48
                                                                _ = v140
                                                            }
                                                        }
                                                        _ = v139
                                                    }
                                                case "hello":
                                                    v125 = append(v125, fi48)
                                                }
                                                if ((2 <= 0) || (p0 && v136)) {
                                                    v124 -= (v124 + (v124 >> ((4294967295) & 63)))
                                                    v105 = 18446744073709551615
                                                }
                                                var v141 uint = 18446744073709551615
                                                _ = v141
                                            }
                                            _ = v136
                                        }
                                        if (v105 > v105) {
                                            switch ("a" + ("\n\t" + ("a" + "foo bar"))) {
                                            case "a":
                                                fmt.Println(fi53, fi53)
                                            case "hello":
                                                var v142 int32 = (v124 + (v124 << ((107) & 63)))
                                                _ = v142
                                            case "\n\t":
                                                var v143 int16 = 42
                                                _ = v143
                                            default:
                                                fmt.Println(v124, fi53)
                                            }
                                            switch "foo bar" {
                                            case "hello":
                                                switch ("\n\t" + "") {
                                                case "":
                                                    switch ("hello" + "\n\t") {
                                                    case "foo bar":
                                                        _ = len(v125)
                                                    }
                                                case "a":
                                                    switch ("foo bar" + "\n\t") {
                                                    case "":
                                                        v125 = append(v125, fi48)
                                                        v124 = -82
                                                        fmt.Println(v124, fi53, fi48)
                                                    default:
                                                        var v144 int16 = 1000
                                                        _ = v144
                                                    }
                                                default:
                                                    v145 := GenS2{F0: ("hello" + "\n\t")}
                                                    _ = v145
                                                }
                                            default:
                                                if false {
                                                    v146 := GenS12{F0: "foo bar"}
                                                    v147 := GenS16{F0: -1}
                                                    var v148 float32 = -2.25
                                                    _ = v146
                                                    _ = v147
                                                    _ = v148
                                                }
                                            }
                                        } else {
                                            if (!p0) {
                                                var v149 uint32 = 4294967295
                                                switch ("\n\t" + "foo bar") {
                                                case "\n\t":
                                                    var v150 uint32 = v149
                                                    _ = v150
                                                default:
                                                    fmt.Println(fi48, p0, v124)
                                                }
                                                _ = len(v125)
                                                _ = v149
                                            }
                                            if (-2.25 < 0.1) {
                                                fmt.Println(v105, fi53, fi53)
                                                if true {
                                                    fmt.Println(fi48)
                                                }
                                            }
                                            v124 -= 0
                                            var v151 float64 = 1e10
                                            v152 := GenS11{F0: (fi48 % (((fi48 >> ((526) & 63))) | 1)), F1: 1}
                                            switch fi48 {
                                            case -2147483648:
                                                if p0 {
                                                    v153 := GenS6{F0: p0, F1: 42}
                                                    if len(v125) >= 2 {
                                                        v125 = v125[1:len(v125)]
                                                    }
                                                    _ = v153
                                                } else {
                                                    if len(v125) > 0 {
                                                        v125[0] = (fi53 ^ (fi48 << ((18446744073709551615) & 63)))
                                                    }
                                                }
                                            case 1:
                                                v125 = append(v125, fi48)
                                            case -493:
                                                fmt.Println(fi48, v151, fi53)
                                            default:
                                                var v154 uint = 402
                                                _ = v154
                                            }
                                            var v155 float64 = v151
                                            _ = v151
                                            _ = v152
                                            _ = v155
                                        }
                                        v156 := GenS10{F0: 42}
                                        _ = v156
                                    }
                                }
                                var v157 uint16 = 65535
                                if (42 <= 100) {
                                    fmt.Println(v124, v105)
                                    if p0 {
                                        var v158 uint32 = 2
                                        v159 := GenS7{F0: "a"}
                                        if (uint(v105) < 2) {
                                            var v160 uint = 251
                                            if false {
                                                v161 := GenS18{F0: v158, F1: v158}
                                                fmt.Println(v105, v124, v158)
                                                v162 := GenS1{F0: (v160 << ((0) & 63)), F1: -42, F2: -128}
                                                _ = v161
                                                _ = v162
                                            }
                                            v105 = v105
                                            _ = v160
                                        } else {
                                            var v163 int8 = -19
                                            if (v124 <= v124) {
                                                var v164 float32 = 1.5
                                                var v165 uint = 992
                                                _ = v164
                                                _ = v165
                                            }
                                            v125 = append(v125, -2147483648)
                                            _ = v163
                                        }
                                        _ = v158
                                        _ = v159
                                    }
                                    if true {
                                        v157 -= v157
                                    }
                                    v125 = append(v125, fi48)
                                    v157 = (v157 | (v157 + 255))
                                } else {
                                    _ = len(v125)
                                    v124 = v124
                                    v157 += v157
                                }
                                _ = v124
                                _ = v125
                                _ = v157
                            }
                            switch (-(fi48)) {
                            case 4294967295:
                                if p0 {
                                    v105 = 1
                                } else {
                                    if p0 {
                                        if (2147483647 == 2147483647) {
                                            var v166 int16 = 42
                                            v167 := []int{fi53}
                                            _ = v166
                                            _ = v167
                                        }
                                    } else {
                                        if p0 {
                                            p0 = ((!(0 > 127)) && p0)
                                        }
                                        switch 127 {
                                        case 206:
                                            v168 := GenS19{F0: 570}
                                            _ = v168
                                        case 242:
                                            if ((v105 >> ((uint(v105)) & 63)) == v105) {
                                                p0 = true
                                                if p0 {
                                                    p0 = (true || p0)
                                                    v169 := GenS9{F0: 83}
                                                    p0 = p0
                                                    var v170 uint64 = (v105 ^ v105)
                                                    _ = v169
                                                    _ = v170
                                                } else {
                                                    fmt.Println(p0, p0, fi48)
                                                    var v171 int64 = 0
                                                    var v172 uint8 = 255
                                                    _ = v171
                                                    _ = v172
                                                }
                                            }
                                        case -868:
                                            var v173 uint32 = 684
                                            _ = v173
                                        default:
                                            v105 -= v105
                                        }
                                    }
                                }
                            case 250:
                                fmt.Println(fi48, fi48, v105)
                            default:
                                p0 = false
                            }
                            fmt.Println(fi48, p0, p0)
                            _ = v105
                        }
                        _ = v102
                    }
                    v174 := GenS4{F0: 4294967295}
                    p0 = (!(-2.25 < 0.0))
                    _ = v174
                }
                v175 := GenS20{F0: 504, F1: false, F2: 65535}
                _ = v175
            }
            v176 := GenS21{F0: 300, F1: 0}
            _ = fi53
            _ = v176
        }
        for fi177 := 0; fi177 < 7; fi177++ {
            v178 := GenS9{F0: 486}
            fmt.Println(p0)
            var v179 int32 = -910
            _ = fi177
            _ = v178
            _ = v179
        }
        _ = fi48
    }
    _ = p0
    _ = v7
    return 4294967295
}
func genFn1() uint {
    if (255 < 32767) {
        v180 := []int{1, 0}
        var v181 int64 = 1
        fmt.Println(v181)
        _ = v180
        _ = v181
    }
    if (0 < 26) {
        v182 := GenS14{F0: 2147483648, F1: -265, F2: 32768}
        _ = v182
    } else {
        for fi183 := 0; fi183 < 20; fi183++ {
            var v184 uint8 = 51
            if true {
                v185 := []int{(fi183 ^ (fi183 + fi183)), 2147483647, 42, (fi183 >> ((18446744073709551615) & 63))}
                v186 := GenS22{F0: 4294967295, F1: 0, F2: (!((42 >= 9223372036854775807) || false))}
                v187 := GenS23{F0: -1, F1: -492, F2: 1}
                _ = v185
                _ = v186
                _ = v187
            }
            fmt.Println(fi183, fi183)
            _ = fi183
            _ = v184
        }
        for fi188 := 0; fi188 < 3; fi188++ {
            if false {
                v189 := []int{fi188, 4294967295}
                if true {
                    if len(v189) > 0 {
                        v189[0] = (fi188 / (((fi188 >> ((genFn0(true)) & 63))) | 1))
                    }
                    fmt.Println(fi188)
                    switch (((("a" + "") + (("hello" + ("foo bar" + ("foo bar" + ("foo bar" + "\n\t")))) + "\n\t")) + "hello") + "\n\t") {
                    case "hello":
                        var v190 int8 = -1
                        _ = v190
                    case "":
                        switch (fi188 >> ((4294967295) & 63)) {
                        case 127:
                            var v191 int16 = -32768
                            _ = v191
                        }
                    case "foo bar":
                        v192 := GenS12{F0: "foo bar"}
                        _ = v192
                    default:
                        var v193 string = "\n\t"
                        _ = v193
                    }
                }
                switch ("a" + "\n\t") {
                case "\n\t":
                    v189 = append(v189, 2147483647)
                }
                _ = v189
            }
            var v194 bool = true
            _ = fi188
            _ = v194
        }
    }
    return 4294967295
}

func main() {
    var v1 int8 = 0
    v2 := make(map[int]uint32)
    var v3 uint32 = 1
    var v4 int8 = (v1 ^ (v1 / ((1) | 1)))
    v5 := []uint32{4294967295}
    v6 := make(map[int]string)
    for fi195 := 0; fi195 < 8; fi195++ {
        v196 := GenS24{F0: 996, F1: 2147483647, F2: genFn1()}
        _ = fi195
        _ = v196
    }
    for fi197 := 0; fi197 < 1; fi197++ {
        switch (fi197 % ((1) | 1)) {
        case 1:
            v198 := GenS25{F0: (fi197 - fi197), F1: ""}
            _ = v198
        case 9223372036854775807:
            fmt.Println(v1, v3)
        }
        switch fi197 {
        case -993:
            switch -226 {
            case 795:
                for fi199 := 0; fi199 < 1; fi199++ {
                    if false {
                        switch "\n\t" {
                        case "":
                            var v200 bool = (true || false)
                            _ = v200
                        case "a":
                            if false {
                                var v201 bool = (-563 >= 9223372036854775807)
                                if (-1 <= 1) {
                                    switch 957 {
                                    case -1:
                                        switch ("a" + "a") {
                                        case "foo bar":
                                            v1 = v4
                                        }
                                        v4 = (v4 - -128)
                                        _ = len(v5)
                                    case -12:
                                        _ = len(v6)
                                    case 128:
                                        var v202 float64 = 1.5
                                        _ = v202
                                    default:
                                        v4 = v4
                                    }
                                    if v201 {
                                        var v203 int = (fi199 - fi199)
                                        if ((!false) || false) {
                                            var v204 uint = 4294967295
                                            if true {
                                                v205 := GenS20{F0: 186, F1: true, F2: 411}
                                                v4 += 0
                                                _ = v205
                                            }
                                            if v201 {
                                                v206 := GenS26{F0: 230}
                                                _ = v206
                                            }
                                            _ = v204
                                        }
                                        _ = v203
                                    } else {
                                        if v201 {
                                            v207 := GenS27{F0: v201, F1: ""}
                                            if true {
                                                v201 = v201
                                                v208 := GenS27{F0: true, F1: ((("" + ("foo bar" + (("foo bar" + "foo bar") + "hello"))) + "") + "\n\t")}
                                                v1 = v1
                                                if (1e10 > 1.5) {
                                                    var v209 string = ("a" + "a")
                                                    if (int64(fi197) > 29) {
                                                        v3 = (v3 | (v3 / ((42) | 1)))
                                                    }
                                                    var v210 int32 = -691
                                                    _ = v209
                                                    _ = v210
                                                }
                                                fmt.Println(v1)
                                                _ = v208
                                            } else {
                                                fmt.Println(fi197)
                                            }
                                            var v211 string = ""
                                            _ = v207
                                            _ = v211
                                        }
                                    }
                                }
                                v4 += (-(v4))
                                _ = v201
                            }
                        default:
                            v6[fi197] = ""
                        }
                        _ = len(v6)
                    }
                    v212 := GenS28{F0: -63}
                    _ = fi199
                    _ = v212
                }
            case -9223372036854775808:
                fmt.Println(v3)
            default:
                var v213 int32 = int32(v3)
                _ = v213
            }
        case -9223372036854775808:
            var v214 float32 = 0.0
            _ = v214
        default:
            if false {
                fmt.Println(v4, v4)
                switch "" {
                case "a":
                    if false {
                        switch "\n\t" {
                        case "hello":
                            fmt.Println(fi197)
                        case "foo bar":
                            v3 = (v3 & v3)
                        case "\n\t":
                            for fi215 := 0; fi215 < 20; fi215++ {
                                var v216 int = (fi197 * fi215)
                                v216 += (v216 / ((fi197) | 1))
                                _ = fi215
                                _ = v216
                            }
                        }
                        v4 = v4
                        fmt.Println(fi197, v3)
                    }
                }
                if true {
                    for fi217 := 0; fi217 < 14; fi217++ {
                        fmt.Println(v4, v3, fi217)
                        var v218 uint32 = 720
                        _ = fi217
                        _ = v218
                    }
                } else {
                    var v219 int64 = 9223372036854775807
                    var v220 float64 = 0.0
                    _ = v219
                    _ = v220
                }
            } else {
                v221 := GenS29{F0: v3, F1: 255}
                for fi222 := 0; fi222 < 5; fi222++ {
                    var v223 int = fi222
                    v224 := GenS30{F0: "hello", F1: 245}
                    _ = fi222
                    _ = v223
                    _ = v224
                }
                _ = v221
            }
        }
        _ = fi197
    }
    switch "foo bar" {
    case "hello":
        if false {
            for fi225 := 0; fi225 < 13; fi225++ {
                if false {
                    delete(v6, 0)
                    var v226 float64 = 0.1
                    _ = v226
                } else {
                    var v227 int16 = 1
                    if true {
                        for fi228 := 0; fi228 < 15; fi228++ {
                            fmt.Println(v4, v227)
                            fmt.Println(v227, v3, v1)
                            delete(v6, (fi228 / ((fi225) | 1)))
                            _ = fi228
                        }
                        var v229 int8 = (v1 << ((4294967295) & 63))
                        v5 = append(v5, v3)
                        _ = v229
                    } else {
                        var v230 int8 = (v4 + (^(v1)))
                        switch fi225 {
                        case 9223372036854775807:
                            if ((!(false || (2147483647 == fi225))) || true) {
                                v227 += (v227 | v227)
                                var v231 uint64 = 691
                                if false {
                                    var v232 uint = 2
                                    v231 += 4294967295
                                    v4 -= -25
                                    _ = v232
                                }
                                _ = v231
                            }
                        }
                        _ = v230
                    }
                    if (!(("hello" + (("" + ("a" + "\n\t")) + "a")) <= "foo bar")) {
                        v4 += v1
                        if false {
                            _ = len(v5)
                            v1 = (v4 + v1)
                        } else {
                            for fi233 := 0; fi233 < 15; fi233++ {
                                var v234 uint16 = 255
                                _ = fi233
                                _ = v234
                            }
                        }
                        v2[9223372036854775807] = (^(v3))
                    }
                    _ = v227
                }
                v5 = append(v5, 42)
                fmt.Println(fi225)
                _ = fi225
            }
        } else {
            for fi235 := 0; fi235 < 14; fi235++ {
                v3 = v3
                v3 -= (v3 | v3)
                v4 = (v4 - -32)
                _ = fi235
            }
        }
    case "foo bar":
        fmt.Println(v1)
    case "a":
        var v236 uint16 = 310
        _ = v236
    }
    v3 = v3
    v3 += v3
    var v237 float32 = -2.25
    var v238 uint64 = 2
    v6[127] = "hello"
    var v239 uint16 = 65535
    for fi240 := 0; fi240 < 16; fi240++ {
        switch "hello" {
        case "foo bar":
            v1 += 127
        case "\n\t":
            if true {
                if true {
                    v238 += v238
                } else {
                    v237 = v237
                }
            } else {
                for fi241 := 0; fi241 < 5; fi241++ {
                    v5 = append(v5, v3)
                    v242 := GenS31{F0: (fi240 - fi240)}
                    _ = fi241
                    _ = v242
                }
            }
        default:
            var v243 uint8 = 42
            _ = v243
        }
        v244 := GenS26{F0: 1}
        _ = fi240
        _ = v244
    }
    if (127 > 0) {
        var v245 int64 = 1
        var v246 int8 = (v4 >> ((255) & 63))
        _ = v245
        _ = v246
    } else {
        v1 += (^(v4))
        v237 -= ((v237 - ((float32(((((float64(((1e10 + 3.14159) + (0.0 / (0.1 - float32(0.0))))) - (1.5 + 1e10)) / -2.25) * (float64(v237) - (3.14159 - (3.14159 * ((3.14159 * 1.5) / 1e10))))) * 1.5)) * v237) / v237)) + v237)
        var v247 int64 = int64(v238)
        _ = v247
    }
    var v248 uint64 = (v238 >> ((133) & 63))
    if true {
        v1 = v1
    } else {
        var v249 uint8 = 51
        var v250 uint = genFn0(true)
        _ = v249
        _ = v250
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v237:", v237)
    fmt.Println("v238:", v238)
    fmt.Println("v239:", v239)
    fmt.Println("v248:", v248)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v237
    _ = v238
    _ = v239
    _ = v248
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
