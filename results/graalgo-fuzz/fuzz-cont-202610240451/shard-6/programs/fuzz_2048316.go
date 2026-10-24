package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int32; F1 int; F2 int64 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int64; F1 bool }
type GenS15 struct { F0 int32; F1 int16 }
type GenS16 struct { F0 bool; F1 int64 }
type GenS17 struct { F0 bool; F1 int; F2 uint16 }
type GenS18 struct { F0 int; F1 uint64 }
type GenS19 struct { F0 int8; F1 int; F2 uint }
type GenS20 struct { F0 uint8 }
func genFn0(p0 int64) int8 {
    v7 := GenS1{F0: 2, F1: 42, F2: 42}
    for fi8 := 0; fi8 < 16; fi8++ {
        p0 = 9223372036854775807
        _ = fi8
    }
    fmt.Println(p0)
    p0 = p0
    v9 := GenS1{F0: 139, F1: 65535, F2: 65535}
    _ = p0
    _ = v7
    _ = v9
    return 42
}
func genFn1(p0 int8, p1 int32) uint {
    v10 := GenS2{F0: ("hello" + "")}
    if true {
        var v11 uint8 = 2
        switch 127 {
        case 231:
            v11 -= v11
        case 42:
            switch 32767 {
            case 2147483647:
                switch "\n\t" {
                case "foo bar":
                    if false {
                        p0 = -128
                        for fi12 := 0; fi12 < 18; fi12++ {
                            p1 = (p1 & (p1 / ((1) | 1)))
                            p0 -= p0
                            fmt.Println(p1, v11)
                            _ = fi12
                        }
                        for fi13 := 0; fi13 < 8; fi13++ {
                            if true {
                                switch fi13 {
                                case -855:
                                    for fi14 := 0; fi14 < 17; fi14++ {
                                        v15 := []int{(^(fi13))}
                                        switch ("\n\t" + ("foo bar" + (("foo bar" + "") + ("a" + "")))) {
                                        case "\n\t":
                                            if true {
                                                v11 -= v11
                                                var v16 float64 = -2.25
                                                var v17 uint8 = (v11 << ((1) & 63))
                                                _ = v16
                                                _ = v17
                                            }
                                        case "a":
                                            var v18 float64 = 1e10
                                            _ = v18
                                        }
                                        switch "foo bar" {
                                        case "hello":
                                            if len(v15) > 0 {
                                                v15[0] = 128
                                            }
                                        case "\n\t":
                                            if ((v11 ^ v11) <= (v11 >> ((0) & 63))) {
                                                _ = len(v15)
                                                if (0 < 1) {
                                                    v15 = append(v15, (fi13 * -886))
                                                } else {
                                                    var v19 uint64 = 0
                                                    v20 := GenS1{F0: v11, F1: 982, F2: 65535}
                                                    var v21 bool = false
                                                    _ = v19
                                                    _ = v20
                                                    _ = v21
                                                }
                                                fmt.Println(fi13, p1, p1)
                                            } else {
                                                var v22 int = fi13
                                                var v23 uint16 = 255
                                                if (v11 == v11) {
                                                    switch (v22 & 32768) {
                                                    case 127:
                                                        if (42 < 4294967295) {
                                                            if true {
                                                                _ = len(v15)
                                                            }
                                                            var v24 uint8 = (v11 * (v11 % (((v11 & 2)) | 1)))
                                                            _ = v24
                                                        } else {
                                                            p1 = p1
                                                            var v25 uint = uint(v23)
                                                            var v26 uint = v25
                                                            _ = v25
                                                            _ = v26
                                                        }
                                                    case 4294967295:
                                                        v11 += (v11 | (^(v11)))
                                                    case 399:
                                                        var v27 int = (v22 * v22)
                                                        _ = v27
                                                    default:
                                                        var v28 float64 = 3.14159
                                                        _ = v28
                                                    }
                                                    v29 := GenS3{F0: 255}
                                                    if true {
                                                        switch (fi13 >> ((0) & 63)) {
                                                        case 4294967295:
                                                            switch "hello" {
                                                            case "foo bar":
                                                                if false {
                                                                    _ = len(v15)
                                                                    v11 -= 27
                                                                    var v30 uint8 = v11
                                                                    _ = v30
                                                                } else {
                                                                    v31 := GenS4{F0: "foo bar"}
                                                                    fmt.Println(fi13, fi13, fi13)
                                                                    _ = v31
                                                                }
                                                            default:
                                                                v32 := GenS5{F0: -1}
                                                                _ = v32
                                                            }
                                                        case 127:
                                                            if (!(true && ((-214 <= -1) || false))) {
                                                                v22 += fi14
                                                                v23 += (v23 >> ((858) & 63))
                                                            }
                                                        case 32768:
                                                            v33 := GenS6{F0: 640}
                                                            _ = v33
                                                        }
                                                        var v34 int64 = 42
                                                        _ = v34
                                                    } else {
                                                        v35 := GenS3{F0: 49}
                                                        _ = v35
                                                    }
                                                    fmt.Println(v22)
                                                    if (!(true && ((!(((!(true && (-2147483648 > (^(p1))))) || (127 > (v11 ^ v11))) && (42 <= -294))) || true))) {
                                                        var v36 int32 = p1
                                                        _ = v36
                                                    } else {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            if (int(p1) <= (v22 / (((fi13 >> ((234) & 63))) | 1))) {
                                                                switch (fi14 & fi13) {
                                                                case -624:
                                                                    var v37 int64 = 1
                                                                    _ = v37
                                                                case 654:
                                                                    v15 = append(v15, 32768)
                                                                default:
                                                                    var v38 bool = (!true)
                                                                    _ = v38
                                                                }
                                                                if true {
                                                                    var v39 int64 = -1
                                                                    _ = v39
                                                                } else {
                                                                    if len(v15) > 0 {
                                                                        v15[0] = fi14
                                                                    }
                                                                    v40 := GenS3{F0: 804}
                                                                    v41 := GenS1{F0: 2, F1: (v23 % ((v23) | 1)), F2: 1}
                                                                    _ = v40
                                                                    _ = v41
                                                                }
                                                            } else {
                                                                var v42 float64 = (((1.5 * (3.14159 + (float64(-2.25) + 3.14159))) + (-2.25 + 1.5)) * (float64(((1e10 + 3.14159) - 1.5)) * 1.5))
                                                                v43 := GenS7{F0: 4294967295}
                                                                if (559 < 19) {
                                                                    if (v11 <= (v11 << ((255) & 63))) {
                                                                        v44 := GenS8{F0: 389, F1: 42, F2: 2}
                                                                        _ = v44
                                                                    }
                                                                } else {
                                                                    v11 -= v11
                                                                    v11 += (v11 | 127)
                                                                    fmt.Println(fi13, v23, fi13)
                                                                }
                                                                _ = v42
                                                                _ = v43
                                                            }
                                                        }
                                                        var v45 uint64 = 4294967295
                                                        if len(v15) >= 2 {
                                                            v15 = v15[1:len(v15)]
                                                        }
                                                        var v46 string = "\n\t"
                                                        var v47 uint8 = (v11 / (((v11 << ((255) & 63))) | 1))
                                                        fmt.Println(v47)
                                                        if (65535 >= 1) {
                                                            v48 := GenS2{F0: v46}
                                                            var v49 float32 = -2.25
                                                            if (false && true) {
                                                                v11 = (v47 * 1)
                                                                if false {
                                                                    v46 = "hello"
                                                                    switch (v22 * (^(v22))) {
                                                                    case -9223372036854775808:
                                                                        if false {
                                                                            v23 = v23
                                                                            if len(v15) >= 2 {
                                                                                v15 = v15[1:len(v15)]
                                                                            }
                                                                            if true {
                                                                                var v50 int8 = 0
                                                                                _ = v50
                                                                            }
                                                                            if (!false) {
                                                                                var v51 float64 = ((((float64(v49) * -2.25) + -2.25) - 1.5) - 0.0)
                                                                                switch (fi14 >> ((1) & 63)) {
                                                                                case -1:
                                                                                    var v52 int8 = genFn0(42)
                                                                                    _ = v52
                                                                                }
                                                                                _ = v51
                                                                            } else {
                                                                                v47 += (v11 + (v11 % ((v47) | 1)))
                                                                            }
                                                                            var v53 int8 = genFn0(int64(v23))
                                                                            _ = v53
                                                                        }
                                                                    case 32768:
                                                                        var v54 int16 = 1
                                                                        _ = v54
                                                                    case -1:
                                                                        v55 := GenS6{F0: (v45 % ((v45) | 1))}
                                                                        _ = v55
                                                                    default:
                                                                        p1 = (p1 / ((-909) | 1))
                                                                    }
                                                                }
                                                                if false {
                                                                    v23 += v23
                                                                    v56 := GenS9{F0: -9223372036854775808}
                                                                    fmt.Println(p0)
                                                                    _ = v56
                                                                } else {
                                                                    if (867 > 42) {
                                                                        switch fi14 {
                                                                        case 612:
                                                                            fmt.Println(v49, v45, fi14)
                                                                        case 127:
                                                                            v47 = (v47 % ((v11) | 1))
                                                                            var v57 int64 = 0
                                                                            v58 := GenS4{F0: v46}
                                                                            _ = v57
                                                                            _ = v58
                                                                        case -107:
                                                                            fmt.Println(v45)
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v45)
                                                                        if false {
                                                                            var v59 int64 = 1
                                                                            v22 += (-(fi13))
                                                                            _ = v59
                                                                        } else {
                                                                            switch (("\n\t" + ("a" + v46)) + ((v46 + v46) + v46)) {
                                                                            case "":
                                                                                v60 := GenS5{F0: 42}
                                                                                _ = v60
                                                                            case "hello":
                                                                                var v61 uint16 = v23
                                                                                _ = v61
                                                                            case "foo bar":
                                                                                v49 = v49
                                                                                if len(v15) >= 2 {
                                                                                    v15 = v15[1:len(v15)]
                                                                                }
                                                                            default:
                                                                                p1 = 2147483647
                                                                            }
                                                                            var v62 string = (v46 + (v46 + v46))
                                                                            _ = v62
                                                                        }
                                                                        v23 = v23
                                                                        switch int(p1) {
                                                                        case -24:
                                                                            p0 += 74
                                                                        default:
                                                                            v22 += v22
                                                                        }
                                                                        switch 1 {
                                                                        case 1:
                                                                            var v63 uint16 = 446
                                                                            _ = v63
                                                                        }
                                                                    }
                                                                    var v64 uint64 = v45
                                                                    switch (((v46 + ((((v46 + "hello") + v46) + (v46 + (v46 + v46))) + ((v46 + "hello") + v46))) + (((((("\n\t" + ("\n\t" + v46)) + v46) + v46) + ("\n\t" + (("\n\t" + ("a" + v46)) + ("foo bar" + ((v46 + v46) + "foo bar"))))) + "a") + "\n\t")) + v46) {
                                                                    case "\n\t":
                                                                        p0 += (-(p0))
                                                                    }
                                                                    _ = v64
                                                                }
                                                            }
                                                            v65 := GenS1{F0: v11, F1: v23, F2: 2}
                                                            fmt.Println(v11, v46)
                                                            _ = v48
                                                            _ = v49
                                                            _ = v65
                                                        }
                                                        if true {
                                                            if false {
                                                                var v66 uint8 = uint8(v23)
                                                                var v67 string = "\n\t"
                                                                if len(v15) >= 2 {
                                                                    v15 = v15[1:len(v15)]
                                                                }
                                                                p0 -= 10
                                                                var v68 uint16 = 60
                                                                var v69 uint = 42
                                                                _ = v66
                                                                _ = v67
                                                                _ = v68
                                                                _ = v69
                                                            } else {
                                                                switch (v46 + (v46 + "foo bar")) {
                                                                case "hello":
                                                                    _ = len(v15)
                                                                case "\n\t":
                                                                    var v70 uint16 = 65535
                                                                    _ = v70
                                                                }
                                                                if true {
                                                                    v71 := GenS10{F0: v46}
                                                                    v72 := GenS2{F0: ((v46 + v46) + (v46 + v46))}
                                                                    _ = v71
                                                                    _ = v72
                                                                }
                                                            }
                                                        } else {
                                                            var v73 string = v46
                                                            v47 -= v47
                                                            var v74 uint16 = (v23 & v23)
                                                            var v75 uint8 = v11
                                                            _ = v73
                                                            _ = v74
                                                            _ = v75
                                                        }
                                                        _ = v45
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                    var v76 int = 4294967295
                                                    _ = v29
                                                    _ = v76
                                                }
                                                v77 := GenS2{F0: ("hello" + "foo bar")}
                                                _ = v22
                                                _ = v23
                                                _ = v77
                                            }
                                        case "a":
                                            if false {
                                                switch fi13 {
                                                case 32768:
                                                    v78 := GenS1{F0: 127, F1: 1, F2: 1}
                                                    _ = v78
                                                default:
                                                    var v79 float64 = (-2.25 - (0.0 * ((1e10 * 1.5) - 3.14159)))
                                                    _ = v79
                                                }
                                                fmt.Println(v11)
                                                if false {
                                                    var v80 int16 = -32768
                                                    if (!false) {
                                                        fmt.Println(v11)
                                                        if len(v15) > 0 {
                                                            v15[0] = fi14
                                                        }
                                                    } else {
                                                        _ = len(v15)
                                                        p0 = p0
                                                        v80 = v80
                                                    }
                                                    var v81 float32 = 3.14159
                                                    if false {
                                                        var v82 float32 = v81
                                                        v15 = append(v15, 32767)
                                                        _ = v82
                                                    } else {
                                                        switch fi14 {
                                                        case -461:
                                                            var v83 int = (fi13 % ((127) | 1))
                                                            _ = v83
                                                        case 32767:
                                                            v11 += v11
                                                        case 42:
                                                            switch 956 {
                                                            case 32:
                                                                v84 := GenS6{F0: 18446744073709551615}
                                                                _ = v84
                                                            case 9223372036854775807:
                                                                switch fi13 {
                                                                case 586:
                                                                    var v85 float32 = float32(3.14159)
                                                                    _ = v85
                                                                case -144:
                                                                    if (!((^(p0)) >= (p0 * genFn0(-402)))) {
                                                                        p1 += -426
                                                                        v86 := GenS5{F0: 42}
                                                                        _ = v86
                                                                    }
                                                                case 32767:
                                                                    if false {
                                                                        if false {
                                                                            var v87 uint = 18446744073709551615
                                                                            var v88 uint = 86
                                                                            _ = v87
                                                                            _ = v88
                                                                        }
                                                                        _ = len(v15)
                                                                        if (42 == 4294967295) {
                                                                            v15 = append(v15, 641)
                                                                        } else {
                                                                            _ = len(v15)
                                                                            v81 = v81
                                                                            v80 = 1
                                                                            var v89 int8 = (p0 * 1)
                                                                            p1 = p1
                                                                            _ = v89
                                                                        }
                                                                    }
                                                                default:
                                                                    v11 += (v11 & v11)
                                                                }
                                                            }
                                                        }
                                                    }
                                                    switch ("hello" + (("foo bar" + "\n\t") + "a")) {
                                                    case "foo bar":
                                                        v90 := GenS10{F0: ("a" + ("foo bar" + "a"))}
                                                        fmt.Println(p0, v81)
                                                        var v91 uint16 = 5
                                                        _ = v90
                                                        _ = v91
                                                    case "\n\t":
                                                        var v92 uint = 25
                                                        _ = v92
                                                    case "":
                                                        var v93 float64 = -2.25
                                                        _ = v93
                                                    }
                                                    _ = v80
                                                    _ = v81
                                                }
                                            }
                                        default:
                                            if true {
                                                var v94 uint16 = 255
                                                if true {
                                                    if len(v15) > 0 {
                                                        v15[0] = 848
                                                    }
                                                    var v95 int64 = 636
                                                    _ = v95
                                                } else {
                                                    var v96 bool = true
                                                    var v97 uint = uint(v11)
                                                    switch "hello" {
                                                    case "a":
                                                        _ = len(v15)
                                                    default:
                                                        var v98 uint = (v97 + v97)
                                                        p0 += p0
                                                        _ = v98
                                                    }
                                                    if ((p1 ^ p1) < p1) {
                                                        if len(v15) > 0 {
                                                            v15[0] = (fi14 * 4294967295)
                                                        }
                                                    }
                                                    _ = v96
                                                    _ = v97
                                                }
                                                v99 := GenS11{F0: v94}
                                                var v100 uint32 = 65535
                                                _ = v100
                                                _ = v94
                                                _ = v99
                                            }
                                        }
                                        var v101 string = "\n\t"
                                        _ = fi14
                                        _ = v101
                                        _ = v15
                                    }
                                case -9223372036854775808:
                                    for fi102 := 0; fi102 < 16; fi102++ {
                                        var v103 uint16 = 801
                                        _ = fi102
                                        _ = v103
                                    }
                                }
                                switch (("a" + "\n\t") + ("" + "a")) {
                                case "\n\t":
                                    v104 := []int{(^(fi13))}
                                    _ = v104
                                }
                                for fi105 := 0; fi105 < 18; fi105++ {
                                    if false {
                                        switch -446 {
                                        case -636:
                                            switch (fi13 | 0) {
                                            case 32767:
                                                if (!(!((0.1 >= 3.14159) || true))) {
                                                    if (((true || (42 > 0)) || true) || ((1 == 0) || false)) {
                                                        if true {
                                                            switch (("" + "a") + ("a" + (("" + "\n\t") + "a"))) {
                                                            case "foo bar":
                                                                var v106 int16 = 32767
                                                                _ = v106
                                                            default:
                                                                var v107 int64 = -727
                                                                _ = v107
                                                            }
                                                        } else {
                                                            var v108 uint64 = uint64(v11)
                                                            var v109 uint32 = 42
                                                            v110 := GenS12{F0: (-(p1)), F1: (fi105 & fi13), F2: 1}
                                                            _ = v108
                                                            _ = v109
                                                            _ = v110
                                                        }
                                                        var v111 float64 = float64(1.5)
                                                        v112 := GenS2{F0: "hello"}
                                                        _ = v111
                                                        _ = v112
                                                    }
                                                    v113 := []int{fi105, fi105, (fi13 >> ((751) & 63))}
                                                    _ = v113
                                                } else {
                                                    fmt.Println(fi13, p0)
                                                }
                                            }
                                        default:
                                            switch "a" {
                                            case "\n\t":
                                                var v114 uint8 = (v11 | v11)
                                                _ = v114
                                            default:
                                                if (3.14159 <= float32(3.14159)) {
                                                    v115 := GenS2{F0: ("a" + ("\n\t" + "foo bar"))}
                                                    if false {
                                                        var v116 uint16 = 255
                                                        fmt.Println(v116, v11)
                                                        _ = v116
                                                    }
                                                    var v117 float64 = (1.5 - 1.5)
                                                    _ = v115
                                                    _ = v117
                                                } else {
                                                    var v118 int32 = p1
                                                    p1 -= (p1 >> ((2) & 63))
                                                    p1 = (v118 & p1)
                                                    _ = v118
                                                }
                                            }
                                        }
                                        var v119 float64 = 0.1
                                        v120 := []int{fi105, fi105, fi105}
                                        _ = v119
                                        _ = v120
                                    } else {
                                        v121 := []int{fi13, (fi105 - fi13), fi13}
                                        switch ("\n\t" + "hello") {
                                        case "a":
                                            var v122 uint32 = 581
                                            _ = v122
                                        case "foo bar":
                                            if len(v121) >= 2 {
                                                v121 = v121[1:len(v121)]
                                            }
                                        default:
                                            if len(v121) > 0 {
                                                v121[0] = -318
                                            }
                                        }
                                        _ = v121
                                    }
                                    _ = fi105
                                }
                            } else {
                                var v123 uint16 = 32767
                                v123 = v123
                                v124 := []int{(fi13 | (fi13 & -2147483648)), (fi13 % ((fi13) | 1)), (fi13 ^ fi13)}
                                _ = v123
                                _ = v124
                            }
                            switch (fi13 >> ((18446744073709551615) & 63)) {
                            case 249:
                                var v125 int = fi13
                                _ = v125
                            case -1:
                                switch "hello" {
                                case "foo bar":
                                    var v126 int16 = 884
                                    _ = v126
                                case "a":
                                    p0 = 0
                                }
                            }
                            _ = fi13
                        }
                    } else {
                        var v127 uint = 1
                        switch "foo bar" {
                        case "":
                            var v128 uint64 = 2
                            _ = v128
                        case "hello":
                            fmt.Println(p1, p0, v127)
                        case "\n\t":
                            p1 = (p1 - (p1 & p1))
                        default:
                            if (1 < 2147483647) {
                                switch 128 {
                                case -683:
                                    switch "\n\t" {
                                    case "foo bar":
                                        if false {
                                            var v129 float64 = (1e10 * 0.0)
                                            _ = v129
                                        } else {
                                            v130 := []int{int(p0), 42}
                                            p1 += (p1 % (((p1 << ((v127) & 63))) | 1))
                                            if (v11 < uint8(v127)) {
                                                v131 := GenS13{F0: p1}
                                                _ = v131
                                            }
                                            _ = v130
                                        }
                                    case "hello":
                                        p0 += p0
                                    case "":
                                        var v132 int8 = genFn0(-9223372036854775808)
                                        _ = v132
                                    default:
                                        var v133 uint8 = 2
                                        _ = v133
                                    }
                                case -2147483648:
                                    for fi134 := 0; fi134 < 4; fi134++ {
                                        v135 := []int{32767, fi134, fi134}
                                        var v136 float32 = 0.1
                                        if (228 > 675) {
                                            if (!false) {
                                                v137 := GenS4{F0: "foo bar"}
                                                _ = v137
                                            }
                                            var v138 uint64 = 2
                                            _ = v138
                                        }
                                        _ = fi134
                                        _ = v135
                                        _ = v136
                                    }
                                case 32767:
                                    v139 := []int{}
                                    _ = v139
                                default:
                                    var v140 uint8 = (v11 >> ((763) & 63))
                                    _ = v140
                                }
                            }
                        }
                        _ = v127
                    }
                case "a":
                    p0 = (p0 / ((p0) | 1))
                case "":
                    v141 := GenS2{F0: "a"}
                    _ = v141
                }
            default:
                var v142 uint64 = 0
                _ = v142
            }
        case 2147483647:
            if ((0.0 + (float32(((0.1 - (1.5 / (1e10 * (((0.1 - 1e10) / ((float64(1e10) * 0.0) - 1e10)) / 1e10)))) / 3.14159)) + (float32((((3.14159 / ((-2.25 - 0.0) - float64(3.14159))) / 1e10) * (1.5 * float64(-2.25)))) + 0.1))) < 3.14159) {
                if (!true) {
                    if (2 >= 42) {
                        v143 := []int{-2147483648, 127, int(p1)}
                        var v144 uint64 = 0
                        _ = v143
                        _ = v144
                    } else {
                        if (!(1 <= 51)) {
                            var v145 int32 = p1
                            for fi146 := 0; fi146 < 4; fi146++ {
                                for fi147 := 0; fi147 < 5; fi147++ {
                                    if true {
                                        var v148 uint = 0
                                        fmt.Println(v148, p0)
                                        _ = v148
                                    } else {
                                        if (0 != 546) {
                                            fmt.Println(fi147, p0)
                                            v149 := []int{(fi146 - 761), -9223372036854775808}
                                            _ = v149
                                        }
                                        var v150 bool = ((p0 * (p0 / ((-41) | 1))) > (p0 ^ -114))
                                        p0 = (-(p0))
                                        _ = v150
                                    }
                                    if true {
                                        v151 := []int{fi146, (-(fi147))}
                                        if len(v151) > 0 {
                                            v151[0] = fi146
                                        }
                                        _ = v151
                                    } else {
                                        var v152 int = (fi147 & (fi147 % (((fi146 * (fi147 + fi146))) | 1)))
                                        _ = v152
                                    }
                                    var v153 uint = 42
                                    v153 = (v153 >> ((v153) & 63))
                                    v154 := []int{9223372036854775807, 32768, fi146}
                                    _ = fi147
                                    _ = v153
                                    _ = v154
                                }
                                var v155 uint16 = 2
                                var v156 bool = true
                                _ = fi146
                                _ = v155
                                _ = v156
                            }
                            v157 := GenS14{F0: -605, F1: false}
                            _ = v145
                            _ = v157
                        } else {
                            for fi158 := 0; fi158 < 11; fi158++ {
                                if (p0 < p0) {
                                    v11 = v11
                                    switch ((("" + "hello") + ("foo bar" + "\n\t")) + "foo bar") {
                                    case "a":
                                        v11 -= v11
                                    case "hello":
                                        var v159 float64 = (-2.25 * float64(-2.25))
                                        _ = v159
                                    case "":
                                        var v160 uint64 = 42
                                        _ = v160
                                    default:
                                        for fi161 := 0; fi161 < 8; fi161++ {
                                            if true {
                                                switch "a" {
                                                case "foo bar":
                                                    v162 := []int{(fi158 << ((18446744073709551615) & 63)), (^(fi161))}
                                                    _ = v162
                                                }
                                                var v163 uint64 = 1
                                                switch fi158 {
                                                case 1:
                                                    v164 := GenS13{F0: p1}
                                                    _ = v164
                                                }
                                                _ = v163
                                            } else {
                                                switch (fi161 >> ((18446744073709551615) & 63)) {
                                                case 4294967295:
                                                    p0 = (p0 ^ 42)
                                                case 32768:
                                                    v11 -= 128
                                                case 127:
                                                    v165 := []int{fi161, (fi161 + fi158)}
                                                    _ = v165
                                                }
                                            }
                                            v11 = (v11 << ((42) & 63))
                                            fmt.Println(fi161, fi158)
                                            _ = fi161
                                        }
                                    }
                                }
                                for fi166 := 0; fi166 < 7; fi166++ {
                                    v167 := []int{(^(fi158)), (fi166 - (-(fi158)))}
                                    _ = fi166
                                    _ = v167
                                }
                                v168 := GenS10{F0: "foo bar"}
                                _ = fi158
                                _ = v168
                            }
                            v169 := []int{-9223372036854775808, 42, -1, 394}
                            p0 -= int8(v11)
                            _ = v169
                        }
                        switch 127 {
                        case -798:
                            v170 := GenS15{F0: -534, F1: 32767}
                            _ = v170
                        case 799:
                            for fi171 := 0; fi171 < 6; fi171++ {
                                v11 = v11
                                var v172 string = (("" + "hello") + ("a" + ("" + "\n\t")))
                                _ = fi171
                                _ = v172
                            }
                        case -9223372036854775808:
                            var v173 uint16 = 438
                            _ = v173
                        }
                        switch "hello" {
                        case "foo bar":
                            v174 := GenS6{F0: 2}
                            _ = v174
                        case "\n\t":
                            fmt.Println(p1, p1, p1)
                        case "hello":
                            v11 += (v11 - v11)
                        default:
                            v11 += v11
                        }
                    }
                    for fi175 := 0; fi175 < 20; fi175++ {
                        fmt.Println(v11, p1, fi175)
                        _ = fi175
                    }
                }
                fmt.Println(p1, v11, p0)
            } else {
                v176 := []int{127, 32768, 409, -2147483648}
                _ = v176
            }
        }
        p0 -= (p0 | (p0 + 42))
        _ = v11
    }
    switch "a" {
    case "foo bar":
        p0 += p0
    case "":
        for fi177 := 0; fi177 < 18; fi177++ {
            var v178 int32 = p1
            var v179 string = ("foo bar" + ("" + (("a" + "\n\t") + "foo bar")))
            _ = fi177
            _ = v178
            _ = v179
        }
    case "hello":
        fmt.Println(p0, p0)
    }
    var v180 float32 = ((1.5 * 3.14159) + float32(float64(float32(0.0))))
    _ = p0
    _ = p1
    _ = v10
    _ = v180
    return 42
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[string]uint32)
    var v3 int16 = -1
    v4 := make(map[string]string)
    var v5 uint8 = 148
    var v6 int16 = v3
    switch 32768 {
    case 0:
        var v181 float32 = (0.1 - 0.1)
        _ = v181
    case 718:
        v182 := []int{323}
        _ = v182
    case -2147483648:
        var v183 int32 = 0
        _ = v183
    }
    fmt.Println(v5, v6, v6)
    switch -2147483648 {
    case -398:
        switch 32768 {
        case 2147483647:
            v5 -= 42
        case 1:
            v6 = 992
        case 42:
            var v184 float32 = 0.0
            _ = v184
        }
    case 0:
        v185 := GenS9{F0: -2147483648}
        _ = v185
    }
    for fi186 := 0; fi186 < 12; fi186++ {
        var v187 uint16 = 0
        v6 += -370
        var v188 uint16 = (^(v187))
        _ = fi186
        _ = v187
        _ = v188
    }
    for fi189 := 0; fi189 < 5; fi189++ {
        switch (-(fi189)) {
        case -9223372036854775808:
            var v190 float64 = 3.14159
            _ = v190
        case -2147483648:
            for fi191 := 0; fi191 < 11; fi191++ {
                fmt.Println(fi191)
                if (-1 != 43) {
                    var v192 int8 = -1
                    v193 := GenS6{F0: 823}
                    _ = v192
                    _ = v193
                }
                var v194 uint32 = v1
                _ = fi191
                _ = v194
            }
        case 1:
            var v195 int = fi189
            _ = v195
        default:
            switch "foo bar" {
            case "hello":
                if (!false) {
                    var v196 int = (fi189 & (fi189 / ((fi189) | 1)))
                    var v197 string = (("" + "\n\t") + (("a" + "") + ("a" + "hello")))
                    switch v197 {
                    case "":
                        if false {
                            v5 = 1
                            if (-33 <= 1) {
                                switch (v196 * (fi189 - (fi189 << ((930) & 63)))) {
                                case 120:
                                    for fi198 := 0; fi198 < 18; fi198++ {
                                        v6 = v6
                                        v196 = (v196 >> ((42) & 63))
                                        v199 := []int{}
                                        _ = fi198
                                        _ = v199
                                    }
                                case -373:
                                    for fi200 := 0; fi200 < 9; fi200++ {
                                        if (true && (float32((1.5 + (3.14159 / (0.1 * 3.14159)))) >= 3.14159)) {
                                            var v201 uint8 = v5
                                            _ = v201
                                        } else {
                                            if true {
                                                var v202 uint = 1
                                                var v203 float64 = -2.25
                                                _ = v202
                                                _ = v203
                                            } else {
                                                v197 = v197
                                                v204 := v2[""]
                                                _ = v204
                                            }
                                        }
                                        if true {
                                            if true {
                                                v196 = 36
                                            }
                                        }
                                        v3 = 0
                                        _ = fi200
                                    }
                                case 2147483647:
                                    var v205 uint16 = 0
                                    _ = v205
                                default:
                                    var v206 int8 = genFn0(9223372036854775807)
                                    _ = v206
                                }
                            }
                            v207 := GenS4{F0: "a"}
                            _ = v207
                        }
                    }
                    _ = v196
                    _ = v197
                } else {
                    fmt.Println(v3)
                    var v208 float64 = 0.0
                    _ = v208
                }
            case "\n\t":
                v209 := []int{(fi189 & 9223372036854775807), fi189}
                _ = v209
            case "foo bar":
                if (255 <= 732) {
                    v6 -= 823
                    if false {
                        if ((((float64(float32(0.1)) - 0.1) + -2.25) != (1e10 - (1e10 / 0.1))) && false) {
                            delete(v4, "a")
                        }
                    }
                } else {
                    if false {
                        for fi210 := 0; fi210 < 8; fi210++ {
                            var v211 uint = 235
                            v5 = (v5 + v5)
                            v211 -= v211
                            v212 := GenS4{F0: ""}
                            _ = fi210
                            _ = v211
                            _ = v212
                        }
                        var v213 int32 = 2147483647
                        _ = v213
                    }
                }
            }
        }
        fmt.Println(v3, v3)
        _ = fi189
    }
    for fi214 := 0; fi214 < 13; fi214++ {
        if false {
            fmt.Println(fi214, v3)
        } else {
            for fi215 := 0; fi215 < 8; fi215++ {
                v3 = v3
                if true {
                    switch (fi214 ^ (fi214 * fi215)) {
                    case 4294967295:
                        fmt.Println(fi215)
                    case 32768:
                        v5 -= (v5 + 0)
                    default:
                        _ = len(v2)
                    }
                } else {
                    if false {
                        switch (("\n\t" + "") + "hello") {
                        case "a":
                            fmt.Println(v6, fi215)
                        case "\n\t":
                            var v216 uint = 4294967295
                            _ = v216
                        }
                        v5 = (v5 << ((255) & 63))
                        v217 := []int{}
                        _ = v217
                    }
                    if true {
                        var v218 uint32 = v1
                        v219 := GenS1{F0: (v5 * v5), F1: 479, F2: uint(v1)}
                        v220 := GenS9{F0: -9223372036854775808}
                        switch "" {
                        case "a":
                            v6 -= (v3 | v6)
                        case "\n\t":
                            v3 = v3
                        case "hello":
                            v221 := GenS6{F0: 42}
                            _ = v221
                        }
                        if (1 != int8(v5)) {
                            v218 -= v1
                        }
                        if (4294967295 >= uint64(v5)) {
                            var v222 uint16 = 1
                            _ = v222
                        } else {
                            v223 := GenS16{F0: true, F1: -1}
                            switch "a" {
                            case "foo bar":
                                delete(v2, "hello")
                                v6 -= (v6 ^ (v3 >> ((2) & 63)))
                                v1 -= 1
                                v224 := GenS11{F0: 427}
                                switch ("hello" + "foo bar") {
                                case "\n\t":
                                    var v225 uint32 = 42
                                    _ = v225
                                case "a":
                                    v226 := GenS17{F0: ((v1 > (^(v1))) && true), F1: 490, F2: 710}
                                    _ = v226
                                case "hello":
                                    var v227 int16 = 32767
                                    _ = v227
                                default:
                                    var v228 int8 = -86
                                    _ = v228
                                }
                                _ = v224
                            }
                            switch 127 {
                            case 32768:
                                if (v1 > 740) {
                                    fmt.Println(v1)
                                } else {
                                    v3 -= -319
                                }
                            default:
                                v5 -= (v5 << ((65535) & 63))
                            }
                            _ = v223
                        }
                        v5 += (v5 * (v5 / ((uint8(v1)) | 1)))
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(v1, v5)
                        case "hello":
                            v3 -= v3
                        case "a":
                            fmt.Println(v6)
                        }
                        var v229 uint64 = 187
                        _ = v218
                        _ = v219
                        _ = v220
                        _ = v229
                    } else {
                        var v230 uint64 = 698
                        switch int(v230) {
                        case 9223372036854775807:
                            v231 := []int{42}
                            _ = v231
                        default:
                            switch "hello" {
                            case "foo bar":
                                var v232 float64 = -2.25
                                _ = v232
                            case "hello":
                                v233 := GenS16{F0: true, F1: -581}
                                _ = v233
                            case "":
                                fmt.Println(fi214, fi215)
                            default:
                                var v234 uint16 = 32767
                                _ = v234
                            }
                        }
                        _ = v230
                    }
                    switch "a" {
                    case "foo bar":
                        if false {
                            v5 -= (v5 + v5)
                            v6 -= (v6 * v3)
                        } else {
                            fmt.Println(fi214)
                            var v235 int8 = 0
                            fmt.Println(v5, v6, v3)
                            _ = v235
                        }
                    }
                }
                _ = fi215
            }
        }
        v1 = (v1 & 0)
        _ = fi214
    }
    v3 += (v6 << ((4294967295) & 63))
    v1 = v1
    switch int(v5) {
    case -2147483648:
        if (4294967295 <= 312) {
            v236 := GenS1{F0: v5, F1: 970, F2: 78}
            fmt.Println(v1, v5, v5)
            switch (("" + "hello") + ("hello" + "foo bar")) {
            case "hello":
                v6 = (^(v6))
            default:
                if false {
                    fmt.Println(v5, v3)
                    fmt.Println(v3)
                }
            }
            _ = v236
        } else {
            var v237 uint32 = (v1 / (((^(v1))) | 1))
            _ = v237
        }
    case -301:
        if false {
            v238 := GenS15{F0: 1, F1: 42}
            _ = len(v4)
            _ = v238
        } else {
            v3 = v3
        }
    case 128:
        for fi239 := 0; fi239 < 3; fi239++ {
            v240 := []int{(^(fi239)), (fi239 + -1)}
            v6 = 221
            _ = fi239
            _ = v240
        }
    default:
        for fi241 := 0; fi241 < 8; fi241++ {
            v5 -= 127
            _ = fi241
        }
    }
    v5 += 119
    if false {
        var v242 int = 9223372036854775807
        if false {
            v242 += v242
            v6 += (v3 / ((-300) | 1))
            switch -2147483648 {
            case -1:
                fmt.Println(v1, v5)
            case 32767:
                var v243 int8 = -43
                _ = v243
            case -945:
                for fi244 := 0; fi244 < 13; fi244++ {
                    var v245 uint = genFn1(42, 42)
                    var v246 int16 = 581
                    v245 += 542
                    _ = fi244
                    _ = v245
                    _ = v246
                }
            default:
                switch "hello" {
                case "":
                    var v247 uint64 = 18446744073709551615
                    _ = v247
                }
            }
        } else {
            v242 = v242
            for fi248 := 0; fi248 < 16; fi248++ {
                if true {
                    v4["\n\t"] = (("\n\t" + ("hello" + ("foo bar" + "\n\t"))) + "foo bar")
                }
                var v249 uint32 = (v1 % (((v1 - v1)) | 1))
                for fi250 := 0; fi250 < 20; fi250++ {
                    var v251 float64 = 1.5
                    var v252 float64 = float64((1e10 + 0.0))
                    v5 = (v5 & v5)
                    fmt.Println(v6, v249, v1)
                    var v253 float64 = (0.1 - (v251 + v251))
                    fmt.Println(v249, v3, v5)
                    v254 := GenS7{F0: 1}
                    var v255 uint32 = v1
                    fmt.Println(v251)
                    v256 := GenS18{F0: fi248, F1: 42}
                    _ = fi250
                    _ = v251
                    _ = v252
                    _ = v253
                    _ = v254
                    _ = v255
                    _ = v256
                }
                _ = fi248
                _ = v249
            }
        }
        _ = v242
    } else {
        switch "a" {
        case "":
            v257 := GenS18{F0: 42, F1: 2}
            _ = v257
        case "a":
            switch 796 {
            case 128:
                if (0 != 9223372036854775807) {
                    if true {
                        if ((v1 <= (v1 >> ((4294967295) & 63))) || false) {
                            if (9223372036854775807 <= -428) {
                                if ((("a" + "") + "foo bar") == "") {
                                    v258 := GenS7{F0: 0}
                                    fmt.Println(v1, v5, v1)
                                    _ = v258
                                }
                                var v259 uint16 = 72
                                _ = v259
                            }
                            v3 += v3
                        }
                        var v260 bool = false
                        _ = v260
                    } else {
                        for fi261 := 0; fi261 < 2; fi261++ {
                            v6 += v3
                            var v262 int = (fi261 % ((fi261) | 1))
                            _ = fi261
                            _ = v262
                        }
                        v3 -= v6
                    }
                    var v263 int32 = -2147483648
                    var v264 uint32 = v1
                    _ = v263
                    _ = v264
                } else {
                    var v265 int32 = -434
                    if ((!true) && (false && false)) {
                        for fi266 := 0; fi266 < 4; fi266++ {
                            if true {
                                v267 := []int{(fi266 / ((1) | 1)), 128, fi266}
                                fmt.Println(v1, v6, v265)
                                v3 = -1
                                _ = v267
                            } else {
                                if (710 >= genFn1(117, v265)) {
                                    switch (-(fi266)) {
                                    case 2147483647:
                                        v5 = (v5 * (^(v5)))
                                    case 128:
                                        var v268 uint = 0
                                        _ = v268
                                    default:
                                        var v269 int64 = 1
                                        _ = v269
                                    }
                                    var v270 int = (fi266 / (((fi266 % (((fi266 - (fi266 % ((fi266) | 1)))) | 1))) | 1))
                                    _ = v270
                                }
                                var v271 int = 548
                                v265 = v265
                                _ = v271
                            }
                            for fi272 := 0; fi272 < 5; fi272++ {
                                var v273 uint16 = uint16(v1)
                                _ = fi272
                                _ = v273
                            }
                            _ = fi266
                        }
                    } else {
                        switch 2147483647 {
                        case 470:
                            var v274 int8 = genFn0(1)
                            _ = v274
                        case 1:
                            var v275 int8 = genFn0(42)
                            _ = v275
                        default:
                            if false {
                                var v276 float64 = -2.25
                                _ = v276
                            }
                        }
                        v277 := GenS4{F0: "a"}
                        _ = v277
                    }
                    var v278 int16 = v6
                    _ = v265
                    _ = v278
                }
            case -871:
                for fi279 := 0; fi279 < 1; fi279++ {
                    fmt.Println(v6, fi279, fi279)
                    _ = fi279
                }
            case 42:
                fmt.Println(v1, v3)
            default:
                var v280 string = "hello"
                _ = v280
            }
        default:
            var v281 string = ("" + "foo bar")
            _ = v281
        }
        v282 := GenS19{F0: 1, F1: 2147483647, F2: 65535}
        switch "a" {
        case "\n\t":
            var v283 string = "hello"
            _ = v283
        }
        _ = v282
    }
    if (461 != 0) {
        var v284 float32 = (((0.1 - 1.5) / ((1e10 * (-2.25 * -2.25)) - (3.14159 + ((3.14159 / 3.14159) + float32(1e10))))) * 1e10)
        fmt.Println(v6, v284)
        if (!((true && (((float64(v284) * float64(v284)) * -2.25) > (0.1 / (1.5 * (((1e10 - (-2.25 + float64(v284))) / float64(v284)) / (3.14159 / 1e10)))))) || (false && false))) {
            for fi285 := 0; fi285 < 12; fi285++ {
                for fi286 := 0; fi286 < 3; fi286++ {
                    v1 -= (v1 >> ((253) & 63))
                    if true {
                        switch ("\n\t" + ("foo bar" + "hello")) {
                        case "foo bar":
                            var v287 uint64 = 42
                            _ = v287
                        case "\n\t":
                            fmt.Println(fi285)
                        }
                        if (((0.1 * -2.25) * float64(v284)) < ((((3.14159 - 3.14159) - ((float64((((v284 + v284) - ((v284 / (v284 - v284)) * v284)) + float32(0.0))) * (float64((3.14159 * 0.1)) * float64((v284 / (((float32((1e10 * 0.1)) + 1e10) * v284) + ((0.1 / -2.25) + float32(3.14159))))))) * 1.5)) / 1.0) / 0.1)) {
                            fmt.Println(v284, fi285, v1)
                            v284 = float32(float64(v284))
                        }
                        switch fi285 {
                        case 32768:
                            v288 := GenS5{F0: -1}
                            v3 = v3
                            _ = v288
                        default:
                            fmt.Println(v284, v1)
                        }
                    } else {
                        switch "" {
                        case "hello":
                            var v289 int16 = (v6 - int16(v5))
                            _ = v289
                        case "\n\t":
                            v5 = (v5 ^ (v5 & 42))
                        case "a":
                            fmt.Println(v6)
                        default:
                            if false {
                                switch -366 {
                                case 58:
                                    if (148 >= 9223372036854775807) {
                                        v290 := GenS20{F0: 57}
                                        fmt.Println(v284, fi286)
                                        fmt.Println(v284, v284)
                                        _ = v290
                                    }
                                    switch fi286 {
                                    case 719:
                                        var v291 int8 = 125
                                        _ = v291
                                    case 4294967295:
                                        v1 = (^(v1))
                                    default:
                                        v284 = v284
                                    }
                                case 0:
                                    if false {
                                        var v292 uint32 = v1
                                        v292 += v292
                                        v293 := []int{}
                                        _ = v292
                                        _ = v293
                                    }
                                case 128:
                                    v294 := GenS19{F0: genFn0(-9223372036854775808), F1: fi285, F2: 18446744073709551615}
                                    _ = v294
                                default:
                                    v1 = (v1 - v1)
                                }
                                v295 := GenS5{F0: -119}
                                _ = v295
                            } else {
                                switch 4294967295 {
                                case 423:
                                    v5 -= v5
                                }
                                v296 := GenS10{F0: "\n\t"}
                                _ = v296
                            }
                        }
                        var v297 uint64 = 698
                        v3 = (v3 & (v3 | v3))
                        _ = v297
                    }
                    _ = fi286
                }
                for fi298 := 0; fi298 < 18; fi298++ {
                    v6 -= (-(v6))
                    fmt.Println(fi298)
                    v299 := GenS15{F0: -2147483648, F1: (v3 & (v6 | 1))}
                    v3 += -922
                    var v300 int32 = 97
                    _ = len(v2)
                    _ = fi298
                    _ = v299
                    _ = v300
                }
                _ = fi285
            }
        } else {
            if (!true) {
                v301 := v2["a"]
                var v302 uint = 65535
                if ((v5 - (v5 | v5)) <= v5) {
                    v303 := v2["a"]
                    _ = v303
                }
                _ = v301
                _ = v302
            }
        }
        _ = v284
    }
    for fi304 := 0; fi304 < 6; fi304++ {
        if (true && false) {
            v3 = v6
            var v305 int16 = v3
            for fi306 := 0; fi306 < 12; fi306++ {
                switch fi304 {
                case -579:
                    v5 += (v5 >> ((42) & 63))
                default:
                    v307 := GenS1{F0: v5, F1: 2, F2: 4294967295}
                    _ = v307
                }
                _ = fi306
            }
            _ = v305
        } else {
            if (!false) {
                var v308 string = ""
                _ = v308
            }
            if true {
                v309 := []int{127}
                var v310 uint32 = (v1 * v1)
                _ = v309
                _ = v310
            }
        }
        v311 := []int{(^(fi304)), fi304, fi304, -441}
        _ = fi304
        _ = v311
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
