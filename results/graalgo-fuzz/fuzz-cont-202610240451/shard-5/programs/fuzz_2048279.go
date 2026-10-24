package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int64 }
type GenS2 struct { F0 int16; F1 uint8; F2 int8 }
type GenS3 struct { F0 int; F1 uint16; F2 int32 }
type GenS4 struct { F0 int; F1 int; F2 uint }
type GenS5 struct { F0 bool; F1 string; F2 string }
type GenS6 struct { F0 uint32; F1 int64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int; F1 int; F2 string }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int; F1 int8; F2 uint16 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint8; F1 int16 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint64; F1 int64; F2 string }
type GenS16 struct { F0 int; F1 uint16 }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 int32; F1 uint64 }
type GenS20 struct { F0 uint64 }
type GenS21 struct { F0 int64; F1 uint; F2 int32 }
type GenS22 struct { F0 bool; F1 int16 }
type GenS23 struct { F0 uint16 }
func genFn0(p0 int, p1 int32, p2 int64) uint8 {
    fmt.Println(p0, p2)
    for fi4 := 0; fi4 < 1; fi4++ {
        var v5 string = "foo bar"
        fmt.Println(p1, p2, p0)
        for fi6 := 0; fi6 < 17; fi6++ {
            p1 = (p1 % (((p1 | (p1 & p1))) | 1))
            _ = fi6
        }
        _ = fi4
        _ = v5
    }
    v7 := []int{}
    switch "foo bar" {
    case "hello":
        if len(v7) > 0 {
            v7[0] = (p0 ^ (-(p0)))
        }
    }
    if false {
        v7 = append(v7, p0)
        p0 -= p0
    } else {
        if (!true) {
            for fi8 := 0; fi8 < 2; fi8++ {
                if (!true) {
                    fmt.Println(p2, p0, p1)
                    v9 := GenS1{F0: "a", F1: 0, F2: p2}
                    p2 -= (-(p2))
                    _ = v9
                }
                if ((!true) && true) {
                    for fi10 := 0; fi10 < 5; fi10++ {
                        var v11 float32 = 0.0
                        _ = fi10
                        _ = v11
                    }
                    if (!true) {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        switch "a" {
                        case "foo bar":
                            p0 += (fi8 % (((-(p0))) | 1))
                        case "\n\t":
                            _ = len(v7)
                        case "hello":
                            if (2 >= 2) {
                                p0 += (fi8 << ((4294967295) & 63))
                                p0 = p0
                            }
                        }
                    }
                } else {
                    p0 = (^(p0))
                    switch ("" + ("" + "a")) {
                    case "\n\t":
                        p1 += (p1 << ((18446744073709551615) & 63))
                    default:
                        var v12 float64 = -2.25
                        _ = v12
                    }
                    fmt.Println(p1)
                }
                _ = fi8
            }
            for fi13 := 0; fi13 < 11; fi13++ {
                for fi14 := 0; fi14 < 15; fi14++ {
                    var v15 float32 = (((float32((float64(3.14159) + (0.1 / -2.25))) / 1e10) / 0.1) - 0.1)
                    if ((-1 == 32767) || (42 < 127)) {
                        v16 := GenS1{F0: ("foo bar" + ("" + "hello")), F1: 208, F2: p2}
                        v7 = append(v7, 4294967295)
                        var v17 float32 = v15
                        _ = v16
                        _ = v17
                    } else {
                        p2 = -1
                        var v18 int8 = -128
                        if true {
                            fmt.Println(p2, v15)
                            fmt.Println(v18, fi13, fi13)
                            v7 = append(v7, fi13)
                            if true {
                                p0 -= fi14
                                v15 -= 0.0
                                v19 := GenS1{F0: "", F1: 129, F2: 42}
                                _ = v19
                            } else {
                                var v20 string = ""
                                switch (v20 + v20) {
                                case "foo bar":
                                    switch ("foo bar" + v20) {
                                    case "\n\t":
                                        if len(v7) > 0 {
                                            v7[0] = (fi14 + fi13)
                                        }
                                    }
                                }
                                _ = v20
                            }
                        }
                        _ = v18
                    }
                    _ = fi14
                    _ = v15
                }
                p1 = 910
                switch "foo bar" {
                case "hello":
                    for fi21 := 0; fi21 < 9; fi21++ {
                        if true {
                            v7 = append(v7, p0)
                        } else {
                            var v22 string = ("" + "\n\t")
                            if true {
                                switch v22 {
                                case "hello":
                                    switch (p0 & (fi13 >> ((2) & 63))) {
                                    case 2147483647:
                                        var v23 int = fi13
                                        _ = v23
                                    }
                                }
                                var v24 string = "foo bar"
                                fmt.Println(p2, fi21, v22)
                                _ = v24
                            }
                            v25 := GenS2{F0: -948, F1: 165, F2: 100}
                            _ = v22
                            _ = v25
                        }
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        var v26 int16 = -1
                        _ = fi21
                        _ = v26
                    }
                case "foo bar":
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                default:
                    switch "\n\t" {
                    case "":
                        var v27 uint16 = 42
                        _ = v27
                    case "\n\t":
                        if true {
                            v28 := GenS2{F0: 42, F1: 127, F2: 42}
                            _ = v28
                        } else {
                            p2 += (p2 >> ((642) & 63))
                            p0 = (^(fi13))
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        }
                    case "foo bar":
                        for fi29 := 0; fi29 < 3; fi29++ {
                            if (32768 >= (^(p0))) {
                                if false {
                                    p2 -= (p2 >> ((255) & 63))
                                    switch p0 {
                                    case -1:
                                        p1 += p1
                                    case 4294967295:
                                        if len(v7) > 0 {
                                            v7[0] = (fi13 - (p0 / ((p0) | 1)))
                                        }
                                    default:
                                        switch (p0 ^ -9223372036854775808) {
                                        case 2147483647:
                                            if (!true) {
                                                var v30 uint8 = 0
                                                fmt.Println(fi29, fi29)
                                                var v31 uint16 = 42
                                                _ = v30
                                                _ = v31
                                            } else {
                                                v32 := GenS3{F0: (p0 / ((fi13) | 1)), F1: 65535, F2: p1}
                                                if ((("hello" != "hello") || (-128 != 61)) || false) {
                                                    p0 += p0
                                                    p0 = 9223372036854775807
                                                    p0 = (-(fi29))
                                                    if (true || (1 == 18446744073709551615)) {
                                                        p0 -= (p0 << ((0) & 63))
                                                        fmt.Println(p1, fi13)
                                                        switch 2147483647 {
                                                        case -750:
                                                            _ = len(v7)
                                                        case 1:
                                                            var v33 float32 = 1.5
                                                            _ = v33
                                                        }
                                                        fmt.Println(fi13)
                                                        _ = len(v7)
                                                    }
                                                } else {
                                                    switch fi29 {
                                                    case -95:
                                                        var v34 uint32 = 2
                                                        _ = v34
                                                    case 9223372036854775807:
                                                        p2 -= 42
                                                    case 4294967295:
                                                        fmt.Println(fi29)
                                                    }
                                                }
                                                fmt.Println(p2, p0, fi29)
                                                _ = v32
                                            }
                                        case 32768:
                                            fmt.Println(p0)
                                        case 128:
                                            p2 -= (p2 - p2)
                                        default:
                                            var v35 uint32 = 302
                                            _ = v35
                                        }
                                    }
                                } else {
                                    var v36 uint64 = 4294967295
                                    _ = v36
                                }
                                var v37 float64 = ((float64(1e10) * 1.5) + float64(0.1))
                                var v38 uint32 = 377
                                _ = v37
                                _ = v38
                            }
                            switch "a" {
                            case "hello":
                                if true {
                                    if (42 > 42) {
                                        if len(v7) > 0 {
                                            v7[0] = (p0 & (fi29 - int(p1)))
                                        }
                                    }
                                    var v39 uint8 = 0
                                    switch (fi13 ^ p0) {
                                    case 95:
                                        v39 = 225
                                    case -2147483648:
                                        if (p2 == p2) {
                                            fmt.Println(p1, v39, v39)
                                            var v40 float64 = 0.0
                                            var v41 string = "a"
                                            _ = v40
                                            _ = v41
                                        }
                                    case 1:
                                        var v42 uint16 = 32767
                                        var v43 int8 = -26
                                        _ = v42
                                        _ = v43
                                    default:
                                        v44 := GenS4{F0: (^(fi13)), F1: int(p2), F2: 0}
                                        _ = v44
                                    }
                                    if (p0 == (fi29 - (^(p0)))) {
                                        v45 := GenS5{F0: false, F1: "", F2: ""}
                                        fmt.Println(fi13)
                                        fmt.Println(p0)
                                        if (42 <= 1) {
                                            if true {
                                                var v46 float32 = float32(0.0)
                                                v39 = (^(v39))
                                                _ = v46
                                            } else {
                                                if true {
                                                    v47 := GenS6{F0: 2, F1: (p2 & (^(p2)))}
                                                    _ = v47
                                                }
                                                switch fi13 {
                                                case 32768:
                                                    var v48 uint64 = 42
                                                    var v49 uint16 = 2
                                                    var v50 float64 = (0.0 / float64(3.14159))
                                                    _ = v48
                                                    _ = v49
                                                    _ = v50
                                                case -9223372036854775808:
                                                    if (325 <= 2147483648) {
                                                        p2 += p2
                                                        p1 += p1
                                                    } else {
                                                        p2 = -1
                                                    }
                                                case 127:
                                                    p2 += (-(p2))
                                                default:
                                                    var v51 float32 = (float32(1e10) - 3.14159)
                                                    _ = v51
                                                }
                                                if (-32768 != 0) {
                                                    switch 128 {
                                                    case 127:
                                                        fmt.Println(p2, p1, fi13)
                                                    case 902:
                                                        var v52 bool = false
                                                        _ = v52
                                                    }
                                                    switch "hello" {
                                                    case "\n\t":
                                                        var v53 int16 = -32768
                                                        _ = v53
                                                    case "foo bar":
                                                        if (166 <= 0) {
                                                            fmt.Println(v39, p1, fi13)
                                                            switch fi29 {
                                                            case -1:
                                                                var v54 int16 = 989
                                                                _ = v54
                                                            case -135:
                                                                p0 -= 32767
                                                            case 0:
                                                                if len(v7) > 0 {
                                                                    v7[0] = 4294967295
                                                                }
                                                            }
                                                        }
                                                    case "hello":
                                                        fmt.Println(p2, p0, p1)
                                                    default:
                                                        v55 := GenS4{F0: fi29, F1: (-(p0)), F2: 255}
                                                        _ = v55
                                                    }
                                                }
                                            }
                                        } else {
                                            fmt.Println(p2, p1)
                                            p0 = 95
                                            fmt.Println(fi29)
                                        }
                                        v56 := GenS7{F0: (v39 ^ 42)}
                                        _ = v45
                                        _ = v56
                                    } else {
                                        p1 += 1
                                        fmt.Println(fi13, p2, v39)
                                    }
                                    fmt.Println(p2, p2)
                                    v57 := GenS5{F0: (-509 > (p2 & 0)), F1: "\n\t", F2: "\n\t"}
                                    p0 -= fi29
                                    _ = v39
                                    _ = v57
                                } else {
                                    if true {
                                        p2 += p2
                                        v58 := GenS8{F0: fi29, F1: (fi29 ^ fi13), F2: "a"}
                                        _ = v58
                                    }
                                    v59 := GenS9{F0: "foo bar"}
                                    _ = v59
                                }
                            case "":
                                _ = len(v7)
                            case "foo bar":
                                var v60 bool = true
                                _ = v60
                            }
                            p1 -= p1
                            _ = fi29
                        }
                    }
                }
                _ = fi13
            }
            var v61 int = (p0 * (p0 + (p0 >> ((739) & 63))))
            _ = v61
        } else {
            switch ("a" + "") {
            case "\n\t":
                v62 := GenS7{F0: 42}
                _ = v62
            case "":
                if (!true) {
                    if true {
                        p1 = p1
                        if false {
                            if (!false) {
                                if len(v7) > 0 {
                                    v7[0] = p0
                                }
                                v63 := GenS5{F0: ((p1 % (((-(p1))) | 1)) > p1), F1: ("hello" + ("a" + "foo bar")), F2: ("hello" + "foo bar")}
                                _ = v63
                            } else {
                                p2 = p2
                            }
                            for fi64 := 0; fi64 < 2; fi64++ {
                                switch -479 {
                                case -9223372036854775808:
                                    v65 := GenS10{F0: (p0 / (((p0 * fi64)) | 1)), F1: 100, F2: 42}
                                    _ = v65
                                case -1:
                                    fmt.Println(fi64)
                                case 128:
                                    var v66 int32 = 2147483647
                                    _ = v66
                                default:
                                    v7 = append(v7, 0)
                                }
                                _ = fi64
                            }
                        }
                    } else {
                        v67 := GenS7{F0: 64}
                        v68 := GenS11{F0: 65535}
                        _ = v67
                        _ = v68
                    }
                }
            case "hello":
                v69 := GenS12{F0: 228, F1: -625}
                _ = v69
            default:
                var v70 uint64 = 894
                _ = v70
            }
            v7 = append(v7, p0)
        }
        if len(v7) > 0 {
            v7[0] = 128
        }
        if true {
            v71 := GenS10{F0: p0, F1: -80, F2: 1}
            p1 = -2147483648
            _ = v71
        } else {
            for fi72 := 0; fi72 < 2; fi72++ {
                for fi73 := 0; fi73 < 3; fi73++ {
                    if false {
                        v7 = append(v7, fi73)
                        if false {
                            p1 = p1
                        }
                    }
                    var v74 bool = true
                    var v75 int = (p0 / (((fi73 >> ((42) & 63))) | 1))
                    _ = fi73
                    _ = v74
                    _ = v75
                }
                p2 -= (^(p2))
                _ = fi72
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}
func genFn1(p0 int32) int32 {
    var v76 string = "\n\t"
    if true {
        for fi77 := 0; fi77 < 20; fi77++ {
            for fi78 := 0; fi78 < 5; fi78++ {
                switch v76 {
                case "a":
                    var v79 int16 = 32767
                    _ = v79
                case "\n\t":
                    var v80 bool = true
                    _ = v80
                case "":
                    if (!true) {
                        if (1 > 4294967295) {
                            switch (v76 + v76) {
                            case "hello":
                                if ((v76 + ("hello" + v76)) < ("foo bar" + "")) {
                                    var v81 uint = 255
                                    var v82 bool = false
                                    _ = v81
                                    _ = v82
                                } else {
                                    var v83 int64 = 42
                                    fmt.Println(fi78)
                                    _ = v83
                                }
                            }
                        }
                    }
                }
                if false {
                    switch ("hello" + v76) {
                    case "foo bar":
                        var v84 string = v76
                        _ = v84
                    case "a":
                        v85 := []int{fi78}
                        _ = v85
                    case "":
                        if (false && (p0 < p0)) {
                            var v86 int16 = 199
                            switch ((v76 + v76) + v76) {
                            case "hello":
                                if true {
                                    v87 := GenS13{F0: 42}
                                    _ = v87
                                } else {
                                    if (228 != 83) {
                                        v76 = ("\n\t" + "foo bar")
                                    } else {
                                        if (9223372036854775807 >= -374) {
                                            switch (v76 + v76) {
                                            case "\n\t":
                                                fmt.Println(fi78)
                                            case "foo bar":
                                                v88 := []int{128, (fi77 + fi78), (fi77 - 127), fi77}
                                                _ = v88
                                            case "a":
                                                v76 = (v76 + v76)
                                                v89 := []int{}
                                                _ = v89
                                            default:
                                                var v90 uint8 = 1
                                                v91 := GenS8{F0: fi77, F1: fi77, F2: (v76 + v76)}
                                                _ = v90
                                                _ = v91
                                            }
                                        } else {
                                            if ((false || (!(-9223372036854775808 <= 1))) || true) {
                                                var v92 int64 = -9223372036854775808
                                                v76 = "hello"
                                                if true {
                                                    switch v76 {
                                                    case "a":
                                                        var v93 int32 = (p0 / ((-2147483648) | 1))
                                                        _ = v93
                                                    }
                                                    v94 := []int{(fi78 - fi78), 4294967295, 657, 32768}
                                                    var v95 uint64 = 0
                                                    _ = v94
                                                    _ = v95
                                                }
                                                fmt.Println(p0)
                                                _ = v92
                                            }
                                            switch fi78 {
                                            case -1:
                                                v96 := GenS6{F0: 602, F1: -9223372036854775808}
                                                v97 := GenS8{F0: (fi77 + (^(fi77))), F1: -1, F2: v76}
                                                v98 := GenS11{F0: 18446744073709551615}
                                                _ = v96
                                                _ = v97
                                                _ = v98
                                            case 2147483647:
                                                fmt.Println(fi77, fi77)
                                            case 4294967295:
                                                fmt.Println(v76, p0)
                                            }
                                        }
                                    }
                                    if false {
                                        v99 := GenS14{F0: 0}
                                        if (1e10 != 1e10) {
                                            switch (fi78 * (^(fi77))) {
                                            case -9223372036854775808:
                                                if (65535 < 573) {
                                                    switch (v76 + v76) {
                                                    case "a":
                                                        v76 = v76
                                                    case "foo bar":
                                                        v100 := []int{(fi78 / (((fi77 % (((fi78 ^ -112)) | 1))) | 1))}
                                                        _ = v100
                                                    }
                                                    if false {
                                                        var v101 bool = (127 != -188)
                                                        var v102 int64 = -374
                                                        fmt.Println(v101, v76, fi78)
                                                        _ = v101
                                                        _ = v102
                                                    } else {
                                                        if (9223372036854775807 < -294) {
                                                            if (true && (true || true)) {
                                                                v103 := []int{704, fi77}
                                                                var v104 uint32 = 2147483648
                                                                _ = v103
                                                                _ = v104
                                                            } else {
                                                                switch 9223372036854775807 {
                                                                case -825:
                                                                    fmt.Println(p0, fi78, fi78)
                                                                }
                                                                v76 = v76
                                                                v105 := GenS12{F0: 42, F1: (v86 ^ (v86 >> ((42) & 63)))}
                                                                v76 = v76
                                                                _ = v105
                                                            }
                                                            v106 := []int{fi77, fi78, -745}
                                                            fmt.Println(p0)
                                                            _ = v106
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(p0)
                                                    var v107 uint16 = 2
                                                    _ = v107
                                                }
                                            default:
                                                v76 = "\n\t"
                                            }
                                            if true {
                                                v108 := GenS15{F0: 2, F1: -9223372036854775808, F2: "a"}
                                                var v109 float32 = (float32(((((1.5 * 0.1) / 0.1) - ((((1e10 / -2.25) * -2.25) * 1.5) * ((3.14159 / ((0.1 + float64(1e10)) - (-2.25 * 0.1))) - 0.1))) * -2.25)) - 0.0)
                                                _ = v108
                                                _ = v109
                                            }
                                            if (0 <= 100) {
                                                v76 = v76
                                                v110 := []int{-211, (fi77 & -976), fi77}
                                                var v111 int8 = 0
                                                _ = v110
                                                _ = v111
                                            }
                                        } else {
                                            v112 := []int{(fi78 | fi77), (fi78 >> ((4294967295) & 63)), (fi77 | -2147483648), (fi78 ^ fi77)}
                                            _ = v112
                                        }
                                        p0 = p0
                                        _ = v99
                                    } else {
                                        if false {
                                            if (0.0 > 0.1) {
                                                fmt.Println(p0, v86, v76)
                                                switch ("" + "\n\t") {
                                                case "foo bar":
                                                    v76 = v76
                                                case "\n\t":
                                                    v86 += (v86 & v86)
                                                case "hello":
                                                    v86 = v86
                                                }
                                            }
                                            var v113 bool = (!((0.1 != 1.5) && false))
                                            if false {
                                                var v114 uint32 = 0
                                                v115 := []int{}
                                                _ = v114
                                                _ = v115
                                            }
                                            _ = v113
                                        } else {
                                            p0 = p0
                                            if false {
                                                fmt.Println(v86, fi78)
                                            } else {
                                                v116 := GenS5{F0: (872 <= 296), F1: (v76 + ("hello" + v76)), F2: "foo bar"}
                                                var v117 uint32 = 903
                                                _ = v116
                                                _ = v117
                                            }
                                        }
                                        var v118 string = ((v76 + v76) + "")
                                        v119 := GenS16{F0: (fi78 + 9223372036854775807), F1: 32767}
                                        _ = v118
                                        _ = v119
                                    }
                                }
                            case "a":
                                v86 -= (v86 + 792)
                            case "\n\t":
                                var v120 int16 = (v86 >> ((820) & 63))
                                _ = v120
                            default:
                                v121 := GenS17{F0: 65535}
                                _ = v121
                            }
                            _ = v86
                        }
                    default:
                        v122 := GenS18{F0: true}
                        _ = v122
                    }
                    v123 := []int{(fi78 % ((-9223372036854775808) | 1))}
                    _ = v123
                } else {
                    var v124 uint8 = 255
                    _ = v124
                }
                _ = fi78
            }
            for fi125 := 0; fi125 < 5; fi125++ {
                fmt.Println(fi125, p0, p0)
                _ = fi125
            }
            _ = fi77
        }
    } else {
        switch v76 {
        case "\n\t":
            p0 -= (^(p0))
        case "":
            if false {
                var v126 uint8 = 105
                if (false && false) {
                    for fi127 := 0; fi127 < 1; fi127++ {
                        for fi128 := 0; fi128 < 18; fi128++ {
                            if true {
                                if (true && false) {
                                    v126 = genFn0(127, (^(p0)), 243)
                                    p0 = (p0 - p0)
                                } else {
                                    fmt.Println(fi127)
                                    v129 := []int{(fi128 & (fi128 ^ (^(fi128)))), fi127}
                                    _ = v129
                                }
                                if (!(false && (-9223372036854775808 != 42))) {
                                    if false {
                                        v130 := []int{(fi127 | (fi127 | fi127)), 57}
                                        _ = v130
                                    } else {
                                        v131 := []int{127, (fi127 * fi128), fi127, (fi127 >> ((65535) & 63))}
                                        switch (v76 + v76) {
                                        case "hello":
                                            if (!(4294967295 > 2147483648)) {
                                                v132 := GenS7{F0: v126}
                                                var v133 float32 = 0.1
                                                if true {
                                                    var v134 float32 = -2.25
                                                    _ = v134
                                                }
                                                _ = v132
                                                _ = v133
                                            } else {
                                                p0 -= p0
                                            }
                                        case "":
                                            if len(v131) > 0 {
                                                v131[0] = -99
                                            }
                                        case "\n\t":
                                            v76 = v76
                                        }
                                        if (((1.5 / 1e10) / 1e10) != float32(0.1)) {
                                            if ((2 <= 2147483648) || true) {
                                                var v135 uint = 4294967295
                                                var v136 uint64 = 179
                                                p0 = (^(p0))
                                                _ = v135
                                                _ = v136
                                            }
                                            v76 = ""
                                        }
                                        _ = v131
                                    }
                                    var v137 uint64 = 0
                                    _ = v137
                                } else {
                                    fmt.Println(fi128, fi128)
                                }
                            } else {
                                v126 += 127
                                v138 := GenS7{F0: genFn0((fi128 ^ (^(fi128))), p0, 793)}
                                if false {
                                    fmt.Println(v76, fi128)
                                    v139 := GenS19{F0: p0, F1: 409}
                                    _ = v139
                                }
                                _ = v138
                            }
                            v76 = ((((v76 + v76) + "") + v76) + (v76 + (v76 + v76)))
                            v140 := GenS20{F0: uint64(v126)}
                            _ = fi128
                            _ = v140
                        }
                        v141 := []int{}
                        if len(v141) >= 2 {
                            v141 = v141[1:len(v141)]
                        }
                        _ = fi127
                        _ = v141
                    }
                    var v142 float64 = float64(0.1)
                    _ = v142
                } else {
                    v76 = v76
                    v143 := []int{21}
                    v144 := GenS9{F0: v76}
                    _ = v143
                    _ = v144
                }
                _ = v126
            } else {
                for fi145 := 0; fi145 < 17; fi145++ {
                    v76 = v76
                    switch (("" + v76) + v76) {
                    case "\n\t":
                        v146 := GenS6{F0: 2147483648, F1: -314}
                        _ = v146
                    }
                    var v147 uint16 = 818
                    _ = fi145
                    _ = v147
                }
                v148 := []int{382}
                if false {
                    var v149 int32 = 1
                    for fi150 := 0; fi150 < 14; fi150++ {
                        fmt.Println(v149, v149, fi150)
                        _ = fi150
                    }
                    _ = v149
                } else {
                    v76 = v76
                    if len(v148) > 0 {
                        v148[0] = 4294967295
                    }
                }
                _ = v148
            }
        default:
            v151 := []int{-305}
            _ = v151
        }
        v152 := []int{1, 1}
        for fi153 := 0; fi153 < 20; fi153++ {
            var v154 float32 = float32(3.14159)
            _ = fi153
            _ = v154
        }
        _ = v152
    }
    if ((0 != 42) || (!(false || (-1 == 127)))) {
        v155 := []int{58, 989, 42, 1}
        if false {
            var v156 int8 = 1
            _ = v156
        }
        _ = v155
    } else {
        v157 := []int{-2147483648, 462}
        var v158 int = -267
        _ = v157
        _ = v158
    }
    _ = p0
    _ = v76
    return p0
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = float64(3.14159)
    var v3 uint16 = 65535
    for fi159 := 0; fi159 < 9; fi159++ {
        var v160 uint16 = (v3 + (^(v3)))
        var v161 int8 = -1
        v162 := v1["\n\t"]
        _ = fi159
        _ = v160
        _ = v161
        _ = v162
    }
    v163 := GenS12{F0: 255, F1: -793}
    v164 := []int{}
    v3 -= (v3 % ((v3) | 1))
    if true {
        if false {
            var v165 int64 = 905
            if ("foo bar" != ((("" + "\n\t") + "") + "\n\t")) {
                var v166 int8 = 42
                for fi167 := 0; fi167 < 7; fi167++ {
                    v1[""] = fi167
                    _ = fi167
                }
                var v168 uint = 1
                _ = v166
                _ = v168
            } else {
                if true {
                    for fi169 := 0; fi169 < 3; fi169++ {
                        switch (fi169 + -105) {
                        case 1:
                            var v170 uint16 = v3
                            _ = v170
                        default:
                            v1[""] = 17
                        }
                        var v171 bool = false
                        _ = fi169
                        _ = v171
                    }
                    if len(v164) > 0 {
                        v164[0] = 9223372036854775807
                    }
                    v165 -= -9223372036854775808
                }
                switch -848 {
                case 845:
                    v165 -= v165
                case 32768:
                    if (2 != 245) {
                        if len(v164) >= 2 {
                            v164 = v164[1:len(v164)]
                        }
                        for fi172 := 0; fi172 < 1; fi172++ {
                            v173 := GenS11{F0: 42}
                            v174 := GenS4{F0: (fi172 - -9223372036854775808), F1: fi172, F2: 65535}
                            for fi175 := 0; fi175 < 8; fi175++ {
                                v2 -= ((3.14159 - ((0.0 - v2) * 0.1)) + ((v2 - 1.5) * (float64(((((-2.25 + 1e10) - float32(1.5)) * 0.1) - 1.5)) / v2)))
                                v2 = 0.0
                                _ = fi175
                            }
                            _ = fi172
                            _ = v173
                            _ = v174
                        }
                        var v176 uint64 = 18446744073709551615
                        _ = v176
                    } else {
                        for fi177 := 0; fi177 < 15; fi177++ {
                            fmt.Println(fi177, v165)
                            _ = fi177
                        }
                        v3 = (v3 * 554)
                    }
                case 128:
                    if (((1 == 42) && (!false)) && true) {
                        if len(v164) >= 2 {
                            v164 = v164[1:len(v164)]
                        }
                        v2 -= 3.14159
                    } else {
                        var v178 string = ("" + ((("a" + "foo bar") + "foo bar") + "\n\t"))
                        switch 32768 {
                        case 32767:
                            fmt.Println(v2)
                        case 42:
                            v1[v178] = 527
                        }
                        _ = v178
                    }
                }
            }
            var v179 float64 = v2
            _ = v165
            _ = v179
        }
        for fi180 := 0; fi180 < 4; fi180++ {
            for fi181 := 0; fi181 < 2; fi181++ {
                var v182 uint32 = 2
                if ((0.0 <= 0.0) || (!false)) {
                    v182 -= v182
                    if false {
                        v182 = v182
                    } else {
                        fmt.Println(v2)
                    }
                }
                v183 := GenS13{F0: (^(v182))}
                _ = fi181
                _ = v182
                _ = v183
            }
            v3 += (v3 + v3)
            var v184 uint64 = 42
            _ = fi180
            _ = v184
        }
        v1["a"] = 9223372036854775807
    } else {
        for fi185 := 0; fi185 < 8; fi185++ {
            var v186 string = "\n\t"
            var v187 float64 = v2
            _ = fi185
            _ = v186
            _ = v187
        }
        var v188 int16 = 974
        _ = v188
    }
    var v189 uint = 455
    fmt.Println(v2, v2)
    var v190 float64 = (1.5 + (1e10 / (v2 / v2)))
    var v191 string = ("foo bar" + (("foo bar" + "hello") + "a"))
    var v192 uint64 = 478
    switch (v191 + v191) {
    case "hello":
        for fi193 := 0; fi193 < 15; fi193++ {
            for fi194 := 0; fi194 < 15; fi194++ {
                v2 = ((3.14159 - v2) / (1.5 / v190))
                _ = fi194
            }
            _ = fi193
        }
    case "foo bar":
        var v195 int64 = 1
        _ = v195
    }
    fmt.Println(v190, v191, v192)
    switch "foo bar" {
    case "foo bar":
        v191 = v191
    default:
        if len(v164) >= 2 {
            v164 = v164[1:len(v164)]
        }
    }
    v196 := GenS19{F0: int32(v3), F1: (v192 - (^(v192)))}
    switch v191 {
    case "a":
        var v197 uint = v189
        _ = v197
    case "":
        if false {
            if false {
                var v198 string = v191
                var v199 uint64 = 4294967295
                _ = v198
                _ = v199
            } else {
                if false {
                    var v200 float64 = (v190 * float64(((1.5 - (1e10 - float32((3.14159 - v2)))) * (0.1 * 1e10))))
                    for fi201 := 0; fi201 < 4; fi201++ {
                        var v202 uint32 = 2147483647
                        _ = fi201
                        _ = v202
                    }
                    switch 127 {
                    case 63:
                        for fi203 := 0; fi203 < 3; fi203++ {
                            if true {
                                v164 = append(v164, (^(fi203)))
                                if (false || (!(int8(v189) == 100))) {
                                    v204 := GenS1{F0: "foo bar", F1: 127, F2: 1}
                                    if true {
                                        switch v191 {
                                        case "\n\t":
                                            fmt.Println(fi203, v3)
                                        case "a":
                                            v200 = 0.0
                                        default:
                                            switch ("foo bar" + (v191 + v191)) {
                                            case "":
                                                var v205 bool = ((128 < 111) || (0.1 > float32(((v200 - 0.0) / v190))))
                                                _ = v205
                                            case "\n\t":
                                                var v206 uint = (v189 >> (((v189 | (v189 ^ 0))) & 63))
                                                _ = v206
                                            default:
                                                if true {
                                                    for fi207 := 0; fi207 < 20; fi207++ {
                                                        v208 := GenS2{F0: -1, F1: 255, F2: -84}
                                                        _ = fi207
                                                        _ = v208
                                                    }
                                                    v209 := GenS17{F0: (v189 ^ v189)}
                                                    _ = v209
                                                }
                                            }
                                        }
                                        if (!(true || false)) {
                                            if (!(float32((1.5 / v190)) < 1.5)) {
                                                _ = len(v164)
                                                v210 := v1[v191]
                                                var v211 int32 = 1
                                                _ = v210
                                                _ = v211
                                            }
                                        }
                                        if (((v190 - v190) < (v200 + float64(((1.5 * 1.5) + (3.14159 / float32(v2)))))) || false) {
                                            fmt.Println(v191)
                                            switch v191 {
                                            case "":
                                                v190 = (float64(-2.25) / 3.14159)
                                            default:
                                                if false {
                                                    for fi212 := 0; fi212 < 7; fi212++ {
                                                        v189 += (^(v189))
                                                        var v213 int16 = 0
                                                        switch v191 {
                                                        case "a":
                                                            v200 += v2
                                                        }
                                                        v214 := GenS14{F0: (v3 ^ 91)}
                                                        if true {
                                                            fmt.Println(v189)
                                                            var v215 string = "foo bar"
                                                            fmt.Println(v200, v213)
                                                            _ = v215
                                                        } else {
                                                            var v216 int16 = (v213 / ((v213) | 1))
                                                            _ = v216
                                                        }
                                                        v2 += 3.14159
                                                        _ = fi212
                                                        _ = v213
                                                        _ = v214
                                                    }
                                                    v217 := GenS10{F0: 0, F1: -68, F2: 1}
                                                    _ = v217
                                                } else {
                                                    fmt.Println(v190, v2, v200)
                                                }
                                            }
                                            var v218 uint16 = 42
                                            _ = v218
                                        } else {
                                            v164 = append(v164, fi203)
                                        }
                                    } else {
                                        _ = len(v1)
                                        if (int64(v189) < 564) {
                                            var v219 uint64 = (v192 / (((v192 - 2)) | 1))
                                            var v220 int32 = -2147483648
                                            var v221 uint16 = v3
                                            _ = v219
                                            _ = v220
                                            _ = v221
                                        }
                                        v191 = (("" + v191) + v191)
                                    }
                                    _ = v204
                                } else {
                                    var v222 float32 = 1.5
                                    if len(v164) >= 2 {
                                        v164 = v164[1:len(v164)]
                                    }
                                    _ = v222
                                }
                            } else {
                                var v223 int16 = 32767
                                _ = v223
                            }
                            v200 -= (v200 / (v200 - v200))
                            if ((v191 > "hello") && (true && (0.1 > 3.14159))) {
                                var v224 uint = (v189 >> ((2) & 63))
                                v1[v191] = (fi203 % ((42) | 1))
                                _ = v224
                            } else {
                                switch "a" {
                                case "a":
                                    fmt.Println(v189)
                                }
                                var v225 int64 = 507
                                v225 += 163
                                _ = v225
                            }
                            _ = fi203
                        }
                    default:
                        for fi226 := 0; fi226 < 10; fi226++ {
                            for fi227 := 0; fi227 < 8; fi227++ {
                                if false {
                                    v228 := GenS12{F0: 0, F1: 580}
                                    v229 := GenS7{F0: 1}
                                    switch (fi227 >> ((v189) & 63)) {
                                    case 32767:
                                        var v230 string = (v191 + (v191 + (((v191 + "\n\t") + v191) + (v191 + ("a" + v191)))))
                                        _ = v230
                                    case -2147483648:
                                        var v231 uint32 = 65535
                                        _ = v231
                                    }
                                    v200 -= (0.1 + (((((0.1 / float64(1e10)) * 3.14159) - 0.1) - float64(float32((v2 * v190)))) - ((1.5 * v190) - 1e10)))
                                    var v232 uint64 = (v192 << ((v189) & 63))
                                    _ = v228
                                    _ = v229
                                    _ = v232
                                } else {
                                    var v233 int32 = -385
                                    v2 += v200
                                    _ = v233
                                }
                                _ = fi227
                            }
                            v164 = append(v164, -178)
                            switch ((v191 + v191) + v191) {
                            case "a":
                                if true {
                                    switch "hello" {
                                    case "a":
                                        if true {
                                            fmt.Println(v200, v189, v200)
                                            v191 = (("\n\t" + v191) + v191)
                                        }
                                    case "foo bar":
                                        var v234 int32 = -2147483648
                                        _ = v234
                                    case "hello":
                                        if true {
                                            if true {
                                                var v235 float32 = (float32(1.5) / ((1e10 * 0.0) + (0.1 * 1e10)))
                                                _ = v235
                                            }
                                        }
                                    default:
                                        v189 = v189
                                    }
                                    switch v191 {
                                    case "foo bar":
                                        v2 = v2
                                    case "":
                                        v236 := GenS21{F0: 1, F1: (v189 << ((v189) & 63)), F2: -869}
                                        _ = v236
                                    }
                                    v189 += v189
                                }
                            case "":
                                for fi237 := 0; fi237 < 14; fi237++ {
                                    var v238 uint8 = 42
                                    if true {
                                        var v239 float32 = (0.1 - (0.0 + 1.5))
                                        v238 = 42
                                        var v240 int8 = 1
                                        fmt.Println(v189, v239)
                                        fmt.Println(v189, v190, v191)
                                        _ = v239
                                        _ = v240
                                    }
                                    _ = fi237
                                    _ = v238
                                }
                            }
                            _ = fi226
                        }
                    }
                    _ = v200
                } else {
                    v2 -= 3.14159
                    v241 := GenS18{F0: (81 <= 42)}
                    for fi242 := 0; fi242 < 12; fi242++ {
                        var v243 uint8 = 128
                        switch -2147483648 {
                        case 620:
                            for fi244 := 0; fi244 < 20; fi244++ {
                                var v245 uint8 = (v243 % (((^(v243))) | 1))
                                v246 := GenS9{F0: v191}
                                v247 := GenS20{F0: (^(v192))}
                                _ = fi244
                                _ = v245
                                _ = v246
                                _ = v247
                            }
                        case -1:
                            v248 := GenS17{F0: 65535}
                            _ = v248
                        default:
                            var v249 float64 = (v190 / ((0.1 * (v2 + 0.0)) * (0.0 * ((((v2 / ((1.5 * ((-2.25 * 1.5) / (v2 / v2))) + float64(float32((v2 + (v190 / (1e10 - v190))))))) + 3.14159) * v190) / (v190 / (v2 / v2))))))
                            _ = v249
                        }
                        for fi250 := 0; fi250 < 2; fi250++ {
                            fmt.Println(v192, v189, v189)
                            var v251 float64 = 3.14159
                            if true {
                                v252 := GenS3{F0: (fi250 >> ((uint(v192)) & 63)), F1: (^(v3)), F2: 2147483647}
                                var v253 int8 = -25
                                switch (v191 + v191) {
                                case "foo bar":
                                    var v254 int = fi250
                                    _ = v254
                                case "":
                                    switch fi242 {
                                    case 42:
                                        var v255 int16 = 102
                                        _ = v255
                                    case -9223372036854775808:
                                        var v256 uint64 = (v192 | (v192 | v192))
                                        _ = v256
                                    case -556:
                                        v3 -= v3
                                    }
                                default:
                                    fmt.Println(fi242, fi242, v2)
                                }
                                v257 := GenS2{F0: -32768, F1: genFn0(438, genFn1(-839), 1), F2: (v253 >> ((617) & 63))}
                                _ = v252
                                _ = v253
                                _ = v257
                            } else {
                                switch "foo bar" {
                                case "":
                                    var v258 int = (fi250 % ((fi242) | 1))
                                    _ = v258
                                case "\n\t":
                                    if (!(float64(-2.25) == (v190 - v251))) {
                                        var v259 int = (-(fi250))
                                        _ = v259
                                    }
                                }
                            }
                            _ = fi250
                            _ = v251
                        }
                        _ = fi242
                        _ = v243
                    }
                    _ = v241
                }
            }
        } else {
            v3 += (^(v3))
            v260 := GenS1{F0: v191, F1: 1, F2: -394}
            _ = v260
        }
    case "foo bar":
        fmt.Println(v190)
    }
    if (896 == -1) {
        switch 127 {
        case 408:
            v261 := GenS10{F0: -9223372036854775808, F1: 0, F2: v3}
            _ = v261
        default:
            v189 += v189
        }
    }
    var v262 int = 1
    switch (v191 + "hello") {
    case "hello":
        if true {
            if ((18446744073709551615 == 498) && (!(986 == 302))) {
                switch v262 {
                case 589:
                    v262 += 32767
                case 624:
                    for fi263 := 0; fi263 < 6; fi263++ {
                        if true {
                            v264 := GenS22{F0: (v190 == (v190 - 0.0)), F1: int16(v3)}
                            _ = v264
                        } else {
                            var v265 bool = (float64(0.1) > float64(1.5))
                            if ((uint32(v192) <= uint32(v3)) || (!(v192 > (v192 | v192)))) {
                                var v266 float64 = v190
                                _ = v266
                            } else {
                                if (167 == 70) {
                                    switch 2147483647 {
                                    case 1:
                                        v192 = 18446744073709551615
                                    case 0:
                                        var v267 int16 = int16(v189)
                                        _ = v267
                                    }
                                } else {
                                    v268 := GenS17{F0: v189}
                                    v189 = v189
                                    _ = v268
                                }
                                for fi269 := 0; fi269 < 10; fi269++ {
                                    v270 := GenS12{F0: 42, F1: 1}
                                    _ = fi269
                                    _ = v270
                                }
                                for fi271 := 0; fi271 < 18; fi271++ {
                                    var v272 string = v191
                                    var v273 string = v272
                                    var v274 int32 = 1
                                    _ = fi271
                                    _ = v272
                                    _ = v273
                                    _ = v274
                                }
                            }
                            _ = v265
                        }
                        for fi275 := 0; fi275 < 3; fi275++ {
                            fmt.Println(v189, fi275, v190)
                            _ = fi275
                        }
                        var v276 float32 = (float32(3.14159) * float32((float64((-2.25 - ((float32(v2) + 1.5) / float32(v2)))) * 1e10)))
                        _ = fi263
                        _ = v276
                    }
                case 2147483647:
                    v3 += (v3 % (((v3 >> (((v189 << ((v189) & 63))) & 63))) | 1))
                }
                v1[v191] = (^(v262))
            }
        }
    default:
        v277 := GenS23{F0: (v3 >> ((v189) & 63))}
        _ = v277
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v164:", len(v164), v164)
    fmt.Println("v189:", v189)
    fmt.Println("v190:", v190)
    fmt.Println("v191:", v191)
    fmt.Println("v192:", v192)
    fmt.Println("v2:", v2)
    fmt.Println("v262:", v262)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v163
    _ = v164
    _ = v189
    _ = v190
    _ = v191
    _ = v192
    _ = v196
    _ = v2
    _ = v262
    _ = v3
}
