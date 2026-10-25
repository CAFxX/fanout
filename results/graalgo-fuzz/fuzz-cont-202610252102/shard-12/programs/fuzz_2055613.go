package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 int16; F2 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS6 struct { F0 uint32; F1 uint32 }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 int64; F1 uint64; F2 int32 }
type GenS9 struct { F0 bool; F1 int64 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 bool; F1 string }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint32; F1 uint; F2 string }
type GenS14 struct { F0 uint16; F1 uint }
type GenS15 struct { F0 uint; F1 uint8 }
type GenS16 struct { F0 uint; F1 uint16; F2 int }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 int64; F1 string }
type GenS19 struct { F0 uint16; F1 uint32 }
type GenS20 struct { F0 uint64; F1 uint32 }
type GenS21 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS22 struct { F0 int32; F1 uint8 }
type GenS23 struct { F0 uint; F1 uint8 }
type GenS24 struct { F0 uint8; F1 int; F2 uint64 }
type GenS25 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS26 struct { F0 uint32; F1 uint }
type GenS27 struct { F0 uint32; F1 uint; F2 int16 }
type GenS28 struct { F0 uint16; F1 uint }
func genFn0() uint32 {
    if false {
        var v4 uint16 = 32767
        if true {
            v5 := GenS1{F0: "\n\t", F1: 163}
            v4 -= 1
            if (2 > 255) {
                v6 := GenS2{F0: 127}
                v4 -= (v4 | 2)
                for fi7 := 0; fi7 < 11; fi7++ {
                    for fi8 := 0; fi8 < 5; fi8++ {
                        v4 = (^(v4))
                        if (1 >= 1) {
                            v4 = 285
                        }
                        var v9 int8 = -128
                        v9 += (v9 - 1)
                        var v10 int8 = 1
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                    _ = fi7
                }
                _ = v6
            } else {
                fmt.Println(v4)
                switch 668 {
                case 50:
                    for fi11 := 0; fi11 < 7; fi11++ {
                        v12 := []int{(fi11 % ((-1) | 1))}
                        v13 := GenS3{F0: -74, F1: -32768, F2: v4}
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                case -223:
                    v4 += (^(v4))
                case -716:
                    switch 1 {
                    case 32767:
                        for fi14 := 0; fi14 < 16; fi14++ {
                            var v15 float32 = 0.0
                            if false {
                                var v16 int = 9223372036854775807
                                if (!(true && (-32768 >= -32768))) {
                                    v16 = v16
                                    var v17 uint64 = 1
                                    _ = v17
                                }
                                _ = v16
                            }
                            var v18 uint = 2
                            _ = fi14
                            _ = v15
                            _ = v18
                        }
                    case 4294967295:
                        for fi19 := 0; fi19 < 4; fi19++ {
                            if false {
                                v20 := []int{fi19, fi19, -1}
                                fmt.Println(v4, fi19)
                                var v21 int8 = 0
                                _ = v20
                                _ = v21
                            }
                            v22 := GenS3{F0: 100, F1: 1, F2: v4}
                            if false {
                                v4 += v4
                                v23 := []int{(fi19 << ((439) & 63)), (fi19 | fi19)}
                                fmt.Println(fi19, fi19)
                                _ = v23
                            }
                            _ = fi19
                            _ = v22
                        }
                    case 9223372036854775807:
                        fmt.Println(v4)
                    default:
                        if (("hello" + "foo bar") < (("hello" + "") + ("a" + "foo bar"))) {
                            for fi24 := 0; fi24 < 1; fi24++ {
                                if false {
                                    for fi25 := 0; fi25 < 6; fi25++ {
                                        fmt.Println(fi25, fi24)
                                        switch (fi24 % ((127) | 1)) {
                                        case -9223372036854775808:
                                            var v26 uint64 = 1000
                                            v26 -= (^(v26))
                                            v27 := GenS1{F0: "a", F1: v26}
                                            _ = v26
                                            _ = v27
                                        default:
                                            v4 = 2
                                        }
                                        _ = fi25
                                    }
                                    if (0 <= -321) {
                                        v28 := GenS3{F0: 1, F1: 1, F2: 902}
                                        if false {
                                            switch ("hello" + "") {
                                            case "a":
                                                if true {
                                                    var v29 uint16 = 0
                                                    for fi30 := 0; fi30 < 18; fi30++ {
                                                        v31 := GenS1{F0: ("\n\t" + (("" + ((("" + ("foo bar" + ("\n\t" + "foo bar"))) + (("hello" + "a") + "foo bar")) + "hello")) + "\n\t")), F1: 711}
                                                        _ = fi30
                                                        _ = v31
                                                    }
                                                    _ = v29
                                                }
                                            case "hello":
                                                var v32 uint16 = (v4 & 261)
                                                _ = v32
                                            case "\n\t":
                                                for fi33 := 0; fi33 < 2; fi33++ {
                                                    switch "a" {
                                                    case "hello":
                                                        var v34 uint16 = v4
                                                        _ = v34
                                                    default:
                                                        var v35 float32 = (0.1 * 1.5)
                                                        _ = v35
                                                    }
                                                    var v36 uint16 = v4
                                                    _ = fi33
                                                    _ = v36
                                                }
                                            }
                                            fmt.Println(v4, v4)
                                        }
                                        _ = v28
                                    } else {
                                        switch (fi24 & 0) {
                                        case 367:
                                            var v37 uint16 = (v4 - (v4 - (v4 - v4)))
                                            _ = v37
                                        }
                                    }
                                    if true {
                                        v38 := GenS4{F0: 18446744073709551615}
                                        _ = v38
                                    }
                                } else {
                                    var v39 string = ("" + "a")
                                    v40 := GenS5{F0: 700, F1: -77, F2: 2}
                                    v41 := []int{fi24, fi24, (fi24 >> ((662) & 63))}
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                }
                                for fi42 := 0; fi42 < 8; fi42++ {
                                    if true {
                                        var v43 int16 = 42
                                        if (434 < v43) {
                                            var v44 int8 = 0
                                            _ = v44
                                        }
                                        if (!((-9223372036854775808 > 42) && false)) {
                                            switch fi42 {
                                            case -9223372036854775808:
                                                v45 := GenS1{F0: "", F1: 2}
                                                _ = v45
                                            case 4294967295:
                                                v43 += (v43 * -1)
                                                v4 += 824
                                                v46 := GenS2{F0: 127}
                                                _ = v46
                                            default:
                                                switch (fi24 >> ((473) & 63)) {
                                                case -901:
                                                    if ((!(int32(v4) <= 42)) && true) {
                                                        v47 := []int{(fi42 & fi24), 32767}
                                                        _ = v47
                                                    }
                                                case -2147483648:
                                                    v48 := GenS6{F0: 42, F1: 497}
                                                    _ = v48
                                                case 963:
                                                    if true {
                                                        var v49 bool = false
                                                        var v50 string = "a"
                                                        _ = v49
                                                        _ = v50
                                                    }
                                                    switch "hello" {
                                                    case "foo bar":
                                                        v51 := GenS7{F0: (fi42 + -2147483648), F1: 1}
                                                        _ = v51
                                                    case "\n\t":
                                                        var v52 uint = 256
                                                        _ = v52
                                                    default:
                                                        if (true && ((float64((float32(-2.25) * 1.5)) + float64((((0.1 * float32(1.5)) * (1e10 * 1e10)) - 0.0))) != (1.5 * ((0.0 + (float64((3.14159 + (1.5 - ((-2.25 / (float32((-2.25 + 0.0)) - 0.1)) - float32(1e10))))) + 1e10)) * 0.0)))) {
                                                            switch (((("" + ("\n\t" + "hello")) + ("" + ("a" + ""))) + "") + "\n\t") {
                                                            case "a":
                                                                if true {
                                                                    var v53 uint = 4294967295
                                                                    _ = v53
                                                                } else {
                                                                    var v54 string = "foo bar"
                                                                    v43 -= v43
                                                                    switch (fi24 | fi42) {
                                                                    case 42:
                                                                        switch 42 {
                                                                        case 1:
                                                                            if ((0.1 == (float32(0.0) + 0.1)) && false) {
                                                                                var v55 int16 = (v43 / ((-827) | 1))
                                                                                var v56 uint64 = 42
                                                                                _ = v55
                                                                                _ = v56
                                                                            }
                                                                        case 42:
                                                                            switch "hello" {
                                                                            case "":
                                                                                v57 := GenS1{F0: "a", F1: 18446744073709551615}
                                                                                _ = v57
                                                                            case "a":
                                                                                v4 -= 65535
                                                                            case "\n\t":
                                                                                v54 = ((("hello" + "a") + "hello") + "\n\t")
                                                                            default:
                                                                                switch fi42 {
                                                                                case 1:
                                                                                    v58 := GenS8{F0: 0, F1: 915, F2: -2147483648}
                                                                                    _ = v58
                                                                                case 4294967295:
                                                                                    if (!(v43 == -973)) {
                                                                                        v59 := GenS4{F0: 460}
                                                                                        var v60 int = (fi24 >> ((424) & 63))
                                                                                        fmt.Println(v43, v54)
                                                                                        _ = v59
                                                                                        _ = v60
                                                                                    }
                                                                                case 9223372036854775807:
                                                                                    if false {
                                                                                        v43 = v43
                                                                                    }
                                                                                }
                                                                            }
                                                                        case 0:
                                                                            fmt.Println(v54, v54, v4)
                                                                        }
                                                                    case 1:
                                                                        switch fi24 {
                                                                        case 9223372036854775807:
                                                                            switch (v54 + (v54 + (v54 + "hello"))) {
                                                                            case "foo bar":
                                                                                var v61 uint64 = 628
                                                                                _ = v61
                                                                            case "":
                                                                                v4 = (v4 / (((^(v4))) | 1))
                                                                            default:
                                                                                v62 := []int{(fi42 & int(v4)), int(v43), (^(fi24))}
                                                                                _ = v62
                                                                            }
                                                                        case -1:
                                                                            if true {
                                                                                var v63 int16 = v43
                                                                                if true {
                                                                                    var v64 uint16 = v4
                                                                                    v65 := GenS5{F0: 2147483647, F1: 0, F2: (v4 / ((1) | 1))}
                                                                                    v64 = 255
                                                                                    _ = v64
                                                                                    _ = v65
                                                                                }
                                                                                switch (-(fi42)) {
                                                                                case 4294967295:
                                                                                    fmt.Println(fi24, fi24, v63)
                                                                                default:
                                                                                    v66 := []int{32767, (^(fi24)), -2147483648}
                                                                                    _ = v66
                                                                                }
                                                                                v43 -= (^(v63))
                                                                                var v67 int64 = 527
                                                                                if false {
                                                                                    v68 := GenS3{F0: 100, F1: (v43 * (v63 << ((4294967295) & 63))), F2: 1}
                                                                                    if false {
                                                                                        if (1.5 == -2.25) {
                                                                                            v69 := []int{(fi24 * fi24), fi42, -785}
                                                                                            v43 += 1
                                                                                            var v70 int16 = v63
                                                                                            _ = v69
                                                                                            _ = v70
                                                                                        }
                                                                                        switch (-(fi42)) {
                                                                                        case 0:
                                                                                            v4 += 75
                                                                                        }
                                                                                        v71 := []int{(fi24 | 758)}
                                                                                        _ = v71
                                                                                    } else {
                                                                                        if (1e10 >= 1.5) {
                                                                                            var v72 uint32 = 2
                                                                                            _ = v72
                                                                                        } else {
                                                                                            v63 += 42
                                                                                        }
                                                                                    }
                                                                                    var v73 string = v54
                                                                                    _ = v68
                                                                                    _ = v73
                                                                                } else {
                                                                                    v54 = "foo bar"
                                                                                }
                                                                                _ = v63
                                                                                _ = v67
                                                                            }
                                                                        case 32767:
                                                                            v43 += (^(v43))
                                                                            var v74 uint32 = 559
                                                                            v74 += v74
                                                                            _ = v74
                                                                        default:
                                                                            var v75 uint8 = 255
                                                                            _ = v75
                                                                        }
                                                                    case 127:
                                                                        var v76 float64 = 1e10
                                                                        _ = v76
                                                                    default:
                                                                        switch fi42 {
                                                                        case 127:
                                                                            var v77 string = "\n\t"
                                                                            _ = v77
                                                                        default:
                                                                            fmt.Println(v54, v4)
                                                                        }
                                                                    }
                                                                    var v78 int32 = 42
                                                                    if (1 > 42) {
                                                                        if true {
                                                                            v79 := []int{(fi42 >> ((1) & 63))}
                                                                            _ = v79
                                                                        }
                                                                    }
                                                                    _ = v54
                                                                    _ = v78
                                                                }
                                                            case "":
                                                                v80 := GenS7{F0: fi42, F1: -904}
                                                                _ = v80
                                                            default:
                                                                var v81 uint32 = 2
                                                                _ = v81
                                                            }
                                                            if true {
                                                                v82 := []int{(fi24 << ((18446744073709551615) & 63)), fi24}
                                                                _ = v82
                                                            }
                                                            var v83 float64 = (1e10 + (1.5 * 0.0))
                                                            _ = v83
                                                        } else {
                                                            v43 = (v43 | (v43 - (v43 >> ((4294967295) & 63))))
                                                            switch "" {
                                                            case "\n\t":
                                                                var v84 int = (fi24 ^ int(v4))
                                                                fmt.Println(v84)
                                                                _ = v84
                                                            case "":
                                                                v85 := []int{(fi42 & (fi24 | (fi42 - (fi24 * (-(fi24)))))), fi42, fi42, 128}
                                                                _ = v85
                                                            case "foo bar":
                                                                if true {
                                                                    v86 := []int{42}
                                                                    v4 -= (^(v4))
                                                                    _ = v86
                                                                } else {
                                                                    var v87 int8 = -41
                                                                    v87 -= (v87 % ((-59) | 1))
                                                                    if (3.14159 == (0.1 + 3.14159)) {
                                                                        var v88 string = "a"
                                                                        _ = v88
                                                                    } else {
                                                                        fmt.Println(v43)
                                                                    }
                                                                    _ = v87
                                                                }
                                                                if false {
                                                                    fmt.Println(v4)
                                                                } else {
                                                                    v89 := GenS1{F0: ("a" + "hello"), F1: 42}
                                                                    v43 += v43
                                                                    _ = v89
                                                                }
                                                                v90 := []int{}
                                                                _ = v90
                                                            }
                                                            if false {
                                                                v91 := []int{-728, 128, 0, fi42}
                                                                v43 = (v43 % ((-870) | 1))
                                                                _ = v91
                                                            }
                                                            fmt.Println(v43, fi24, v4)
                                                        }
                                                    }
                                                    if (false || true) {
                                                        if true {
                                                            if (((3.14159 / ((0.1 * -2.25) + -2.25)) * 1.5) > (((3.14159 * (0.0 / (-2.25 / -2.25))) + 3.14159) * (0.1 - 1.5))) {
                                                                v92 := GenS3{F0: 127, F1: v43, F2: v4}
                                                                if ((fi24 & (fi24 * (fi42 ^ (fi24 % ((0) | 1))))) != fi24) {
                                                                    v93 := GenS1{F0: "\n\t", F1: 18446744073709551615}
                                                                    v94 := GenS3{F0: -1, F1: (^(v43)), F2: (v4 * 1)}
                                                                    _ = v93
                                                                    _ = v94
                                                                } else {
                                                                    if (false || true) {
                                                                        v4 = v4
                                                                    } else {
                                                                        var v95 string = "foo bar"
                                                                        _ = v95
                                                                    }
                                                                }
                                                                _ = v92
                                                            } else {
                                                                fmt.Println(fi42)
                                                                var v96 bool = false
                                                                _ = v96
                                                            }
                                                            v4 += 65535
                                                            v43 = (v43 >> ((2) & 63))
                                                            v97 := []int{(fi24 % (((-(fi24))) | 1)), fi42}
                                                            _ = v97
                                                        } else {
                                                            v98 := GenS9{F0: (!(!(-2.25 <= 3.14159))), F1: -902}
                                                            v4 = (v4 >> ((867) & 63))
                                                            _ = v98
                                                        }
                                                    }
                                                    var v99 uint32 = 4294967295
                                                    v43 = 1
                                                    v4 -= (^(v4))
                                                    _ = v99
                                                }
                                            }
                                        }
                                        v100 := []int{(fi24 * (fi42 >> ((4294967295) & 63)))}
                                        var v101 int64 = 971
                                        _ = v100
                                        _ = v101
                                        _ = v43
                                    } else {
                                        v4 += 2
                                        fmt.Println(v4)
                                        switch "" {
                                        case "a":
                                            var v102 int = fi42
                                            _ = v102
                                        case "foo bar":
                                            if false {
                                                if false {
                                                    v103 := []int{fi24}
                                                    _ = v103
                                                } else {
                                                    switch (fi42 | 4294967295) {
                                                    case 9223372036854775807:
                                                        v104 := GenS10{F0: fi42}
                                                        _ = v104
                                                    default:
                                                        v105 := GenS7{F0: fi24, F1: 361}
                                                        _ = v105
                                                    }
                                                    var v106 int8 = 24
                                                    _ = v106
                                                }
                                                var v107 uint64 = 2
                                                fmt.Println(fi24)
                                                _ = v107
                                            } else {
                                                var v108 int16 = 42
                                                switch ("\n\t" + "hello") {
                                                case "a":
                                                    v109 := []int{(fi24 << ((1) & 63)), 4294967295, fi24, (fi24 % (((fi42 / (((fi42 * (fi42 + (fi24 - -2147483648)))) | 1))) | 1))}
                                                    _ = v109
                                                default:
                                                    v110 := GenS10{F0: (^(fi42))}
                                                    _ = v110
                                                }
                                                _ = v108
                                            }
                                        case "hello":
                                            var v111 int32 = -465
                                            _ = v111
                                        }
                                    }
                                    var v112 float32 = 1e10
                                    v113 := GenS11{F0: true, F1: ""}
                                    _ = fi42
                                    _ = v112
                                    _ = v113
                                }
                                fmt.Println(v4, fi24)
                                _ = fi24
                            }
                            switch 127 {
                            case -60:
                                var v114 int8 = -1
                                _ = v114
                            case -1:
                                var v115 uint = 4294967295
                                _ = v115
                            default:
                                if true {
                                    fmt.Println(v4)
                                } else {
                                    for fi116 := 0; fi116 < 20; fi116++ {
                                        if (false && (-374 != -32768)) {
                                            for fi117 := 0; fi117 < 1; fi117++ {
                                                if (float32((1e10 * -2.25)) < 0.1) {
                                                    v4 += (^(v4))
                                                }
                                                var v118 int8 = 127
                                                _ = fi117
                                                _ = v118
                                            }
                                            switch "\n\t" {
                                            case "\n\t":
                                                v4 += (v4 ^ v4)
                                            case "a":
                                                v4 -= v4
                                            case "foo bar":
                                                var v119 uint64 = 1
                                                _ = v119
                                            }
                                        }
                                        v4 += 499
                                        _ = fi116
                                    }
                                }
                            }
                            v120 := GenS12{F0: "a"}
                            _ = v120
                        }
                    }
                }
            }
            _ = v5
        } else {
            for fi121 := 0; fi121 < 5; fi121++ {
                if true {
                    if (42 < 1) {
                        var v122 uint32 = 2147483648
                        _ = v122
                    }
                }
                if true {
                    var v123 uint = 503
                    for fi124 := 0; fi124 < 2; fi124++ {
                        v125 := GenS11{F0: (-399 > -568), F1: (((("a" + ("" + "")) + ((("foo bar" + "hello") + "") + "hello")) + "a") + ("hello" + "\n\t"))}
                        _ = fi124
                        _ = v125
                    }
                    _ = v123
                }
                for fi126 := 0; fi126 < 12; fi126++ {
                    switch "hello" {
                    case "foo bar":
                        if (0.1 == -2.25) {
                            var v127 uint = 4294967295
                            _ = v127
                        } else {
                            v128 := []int{}
                            _ = v128
                        }
                    case "\n\t":
                        if true {
                            if (0.1 > 3.14159) {
                                if false {
                                    if (32768 < fi121) {
                                        v4 = (v4 * 256)
                                    } else {
                                        v4 = v4
                                        if ("\n\t" > "foo bar") {
                                            fmt.Println(fi121)
                                            var v129 int8 = 42
                                            v130 := []int{}
                                            fmt.Println(fi126)
                                            _ = v129
                                            _ = v130
                                        } else {
                                            var v131 uint16 = (^(v4))
                                            var v132 uint16 = v131
                                            var v133 int16 = 32767
                                            _ = v131
                                            _ = v132
                                            _ = v133
                                        }
                                    }
                                    if true {
                                        v134 := GenS6{F0: 65535, F1: 4294967295}
                                        _ = v134
                                    }
                                } else {
                                    v4 = 1
                                }
                                var v135 float32 = ((0.0 / (1e10 * 0.1)) * 3.14159)
                                _ = v135
                            } else {
                                if false {
                                    var v136 uint16 = v4
                                    v4 += v4
                                    v137 := []int{}
                                    _ = v136
                                    _ = v137
                                } else {
                                    var v138 int32 = int32(v4)
                                    _ = v138
                                }
                                if false {
                                    var v139 uint32 = 1
                                    var v140 int32 = -104
                                    _ = v139
                                    _ = v140
                                }
                                v141 := []int{fi121, (fi121 % ((-908) | 1)), fi121, 9223372036854775807}
                                _ = v141
                            }
                        }
                    default:
                        if true {
                            var v142 uint32 = 1
                            if ((18446744073709551615 <= 2) || ((-(fi121)) != fi126)) {
                                v143 := []int{}
                                if true {
                                    if true {
                                        var v144 uint8 = uint8(v142)
                                        switch "foo bar" {
                                        case "":
                                            switch 0 {
                                            case 32768:
                                                v145 := GenS13{F0: v142, F1: 4294967295, F2: ""}
                                                _ = v145
                                            case 119:
                                                v146 := GenS14{F0: (v4 >> ((521) & 63)), F1: 312}
                                                _ = v146
                                            case 512:
                                                switch (fi121 / ((fi126) | 1)) {
                                                case -740:
                                                    v4 = (^(v4))
                                                case -421:
                                                    var v147 uint32 = 4294967295
                                                    _ = v147
                                                case -9223372036854775808:
                                                    switch "" {
                                                    case "\n\t":
                                                        fmt.Println(fi121)
                                                    case "a":
                                                        v148 := GenS15{F0: 589, F1: v144}
                                                        _ = v148
                                                    }
                                                }
                                            }
                                        case "\n\t":
                                            v4 += v4
                                        default:
                                            if false {
                                                v142 += 65535
                                                if (1 >= 42) {
                                                    if (42 > 42) {
                                                        if len(v143) > 0 {
                                                            v143[0] = (^(fi126))
                                                        }
                                                        if true {
                                                            v149 := GenS12{F0: ("foo bar" + ("foo bar" + "a"))}
                                                            if true {
                                                                v4 -= (v4 >> ((uint(v144)) & 63))
                                                                if true {
                                                                    var v150 int32 = 42
                                                                    fmt.Println(v4, v144)
                                                                    var v151 string = ((("hello" + "") + ((("hello" + ((("hello" + "\n\t") + ("" + "a")) + "a")) + ("" + "\n\t")) + "")) + "foo bar")
                                                                    _ = v150
                                                                    _ = v151
                                                                } else {
                                                                    v143 = append(v143, (fi121 - fi121))
                                                                    var v152 bool = (v4 >= v4)
                                                                    if false {
                                                                        var v153 uint64 = 42
                                                                        _ = v153
                                                                    }
                                                                    _ = v152
                                                                }
                                                                v144 = 127
                                                            } else {
                                                                v144 += uint8(v4)
                                                                v154 := GenS16{F0: 98, F1: 78, F2: fi121}
                                                                v4 -= v4
                                                                _ = v154
                                                            }
                                                            _ = v149
                                                        }
                                                        switch -81 {
                                                        case 4294967295:
                                                            switch fi121 {
                                                            case 0:
                                                                var v155 uint64 = 591
                                                                _ = v155
                                                            case 2:
                                                                v4 = (v4 ^ 0)
                                                            }
                                                        case 0:
                                                            var v156 uint8 = (v144 + v144)
                                                            _ = v156
                                                        case 237:
                                                            if true {
                                                                v4 += v4
                                                                switch ("foo bar" + (("\n\t" + ("" + "\n\t")) + "foo bar")) {
                                                                case "hello":
                                                                    fmt.Println(fi121)
                                                                }
                                                                v157 := GenS6{F0: v142, F1: (v142 / ((uint32(v144)) | 1))}
                                                                _ = v157
                                                            } else {
                                                                if (32767 >= 637) {
                                                                    v143 = append(v143, 9223372036854775807)
                                                                    if ("foo bar" > "foo bar") {
                                                                        var v158 uint64 = 4294967295
                                                                        var v159 float64 = (float64(((3.14159 / 1.5) / float32(3.14159))) + 1.5)
                                                                        switch "foo bar" {
                                                                        case "\n\t":
                                                                            var v160 uint = 73
                                                                            if ((v160 / (((v160 - (^(v160)))) | 1)) < v160) {
                                                                                fmt.Println(fi126)
                                                                                var v161 int16 = -716
                                                                                v143 = append(v143, fi126)
                                                                                _ = v161
                                                                            }
                                                                            if len(v143) >= 2 {
                                                                                v143 = v143[1:len(v143)]
                                                                            }
                                                                            if false {
                                                                                fmt.Println(v159, v4)
                                                                                v158 = (v158 & v158)
                                                                            } else {
                                                                                if false {
                                                                                    v162 := GenS12{F0: "hello"}
                                                                                    _ = v162
                                                                                }
                                                                                switch 128 {
                                                                                case -904:
                                                                                    v163 := GenS6{F0: 621, F1: v142}
                                                                                    _ = v163
                                                                                }
                                                                            }
                                                                            _ = v160
                                                                        case "hello":
                                                                            v159 = float64((1e10 - (float32(v159) * 0.0)))
                                                                        }
                                                                        _ = v158
                                                                        _ = v159
                                                                    } else {
                                                                        var v164 int16 = 0
                                                                        _ = v164
                                                                    }
                                                                }
                                                                v165 := GenS12{F0: ""}
                                                                v144 = 127
                                                                v144 -= (v144 << ((uint(v144)) & 63))
                                                                _ = v165
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    var v166 float32 = (-2.25 + float32((0.0 + -2.25)))
                                                    v142 = v142
                                                    _ = v166
                                                }
                                                var v167 uint = 42
                                                _ = v167
                                            }
                                        }
                                        v142 = (v142 << ((833) & 63))
                                        _ = v144
                                    }
                                    if (9223372036854775807 >= 0) {
                                        if true {
                                            var v168 uint = 5
                                            _ = v168
                                        }
                                        v4 -= v4
                                        fmt.Println(v4, v4, v142)
                                    } else {
                                        if true {
                                            var v169 int32 = 2147483647
                                            var v170 bool = ((true && true) || true)
                                            _ = v169
                                            _ = v170
                                        } else {
                                            switch "a" {
                                            case "a":
                                                var v171 float32 = -2.25
                                                _ = v171
                                            case "hello":
                                                switch (fi126 & fi126) {
                                                case 4294967295:
                                                    v172 := GenS14{F0: v4, F1: 1}
                                                    _ = v172
                                                }
                                            case "":
                                                if len(v143) >= 2 {
                                                    v143 = v143[1:len(v143)]
                                                }
                                            default:
                                                var v173 int16 = 676
                                                _ = v173
                                            }
                                            var v174 float64 = float64(float32((0.0 / (3.14159 * (float64(1e10) - 0.0)))))
                                            _ = v174
                                        }
                                        switch ("hello" + "foo bar") {
                                        case "":
                                            switch int(v142) {
                                            case -795:
                                                var v175 float32 = 0.0
                                                _ = v175
                                            case 691:
                                                switch ("\n\t" + "hello") {
                                                case "\n\t":
                                                    v176 := GenS17{F0: 1}
                                                    _ = v176
                                                case "a":
                                                    var v177 float64 = 3.14159
                                                    var v178 int32 = -2147483648
                                                    _ = v177
                                                    _ = v178
                                                default:
                                                    v179 := GenS2{F0: 42}
                                                    _ = v179
                                                }
                                            case -2147483648:
                                                fmt.Println(fi121, fi121)
                                                if (0.1 > ((((3.14159 + ((float32(1.5) / 0.1) * -2.25)) / (-2.25 + (0.0 * (0.1 - 0.0)))) - float32(1.5)) / (float32(0.1) / float32(3.14159)))) {
                                                    v4 += (v4 * v4)
                                                } else {
                                                    var v180 float64 = (0.1 - ((0.0 * 0.1) / (1e10 - (float64(0.0) / (1.5 - -2.25)))))
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        var v181 int16 = -32768
                                                        _ = v181
                                                    case "\n\t":
                                                        if len(v143) > 0 {
                                                            v143[0] = -9223372036854775808
                                                        }
                                                    default:
                                                        fmt.Println(fi126, v180)
                                                    }
                                                    _ = v180
                                                }
                                                v182 := GenS8{F0: 1, F1: 18446744073709551615, F2: -928}
                                                v183 := GenS8{F0: -522, F1: 2, F2: 1}
                                                if (-2.25 == -2.25) {
                                                    fmt.Println(fi126, fi126)
                                                }
                                                v4 -= v4
                                                _ = v182
                                                _ = v183
                                            default:
                                                v184 := GenS18{F0: -9223372036854775808, F1: (("foo bar" + "a") + "\n\t")}
                                                _ = v184
                                            }
                                        default:
                                            fmt.Println(v4, v4, fi126)
                                        }
                                        v142 += v142
                                    }
                                }
                                switch "\n\t" {
                                case "foo bar":
                                    fmt.Println(fi121)
                                default:
                                    if (true || false) {
                                        v4 += (^(v4))
                                    } else {
                                        var v185 int8 = 7
                                        _ = v185
                                    }
                                }
                                v186 := GenS19{F0: 1, F1: 65535}
                                _ = v143
                                _ = v186
                            } else {
                                if (true && (false && true)) {
                                    var v187 float32 = 3.14159
                                    var v188 float32 = v187
                                    if ((false || (true || false)) && (true || (fi126 != fi126))) {
                                        if (!(!(!((1 != 81) && ((v142 % (((^(v142))) | 1)) >= uint32(v4)))))) {
                                            v188 = ((((v188 + (float32((3.14159 + -2.25)) / (v188 + 0.0))) + (v188 / 1.0)) * ((((0.0 / v188) - float32(1.5)) / 1e10) - (v188 - v188))) * v187)
                                            var v189 uint64 = uint64(v4)
                                            _ = v189
                                        } else {
                                            v190 := GenS20{F0: 18446744073709551615, F1: 42}
                                            _ = v190
                                        }
                                        var v191 int16 = 32767
                                        _ = v191
                                    }
                                    v4 = (^(v4))
                                    v192 := GenS13{F0: v142, F1: 18446744073709551615, F2: ((("" + ("hello" + "foo bar")) + ("foo bar" + "a")) + "a")}
                                    _ = v187
                                    _ = v188
                                    _ = v192
                                }
                            }
                            _ = v142
                        } else {
                            fmt.Println(fi126)
                        }
                    }
                    v193 := []int{(fi126 % ((-651) | 1)), 772}
                    fmt.Println(fi121, v4, v4)
                    _ = fi126
                    _ = v193
                }
                _ = fi121
            }
            switch "a" {
            case "a":
                v194 := GenS15{F0: 65535, F1: 153}
                _ = v194
            case "foo bar":
                for fi195 := 0; fi195 < 19; fi195++ {
                    if true {
                        if false {
                            if true {
                                v196 := []int{(fi195 | (fi195 >> ((18446744073709551615) & 63)))}
                                _ = v196
                            } else {
                                v197 := GenS21{F0: 179, F1: 1, F2: (v4 + 0)}
                                _ = v197
                            }
                        }
                    }
                    fmt.Println(v4, fi195)
                    _ = fi195
                }
            default:
                switch -1 {
                case 9223372036854775807:
                    v4 = v4
                case -9223372036854775808:
                    v198 := []int{}
                    _ = v198
                }
            }
            if (1.5 < 0.1) {
                var v199 uint8 = 128
                _ = v199
            }
        }
        _ = v4
    }
    var v200 string = (("foo bar" + "hello") + ("\n\t" + "hello"))
    if (-2.25 > -2.25) {
        v200 = v200
        v201 := GenS3{F0: -96, F1: 243, F2: 42}
        for fi202 := 0; fi202 < 7; fi202++ {
            for fi203 := 0; fi203 < 4; fi203++ {
                var v204 uint = 4294967295
                v204 += v204
                _ = fi203
                _ = v204
            }
            _ = fi202
        }
        _ = v201
    } else {
        v200 = v200
        v205 := GenS22{F0: 1, F1: 255}
        _ = v205
    }
    _ = v200
    return 501
}

func main() {
    v1 := []int{42}
    v2 := []uint32{65535, 2147483647}
    var v3 uint64 = 589
    v3 = 2
    v3 = v3
    for fi206 := 0; fi206 < 13; fi206++ {
        if (611 == 42) {
            for fi207 := 0; fi207 < 13; fi207++ {
                fmt.Println(fi207, fi206, fi206)
                _ = fi207
            }
            if true {
                v3 = (v3 + 18446744073709551615)
            }
        } else {
            v3 -= (v3 - 2)
            switch "\n\t" {
            case "hello":
                v3 += (v3 ^ (v3 ^ (v3 / ((v3) | 1))))
            default:
                var v208 uint64 = (v3 - v3)
                _ = v208
            }
        }
        _ = fi206
    }
    v209 := GenS23{F0: 2, F1: 0}
    switch "\n\t" {
    case "hello":
        switch "\n\t" {
        case "foo bar":
            if ((false || false) || false) {
                var v210 float64 = (((1e10 + 3.14159) - 3.14159) - 0.1)
                _ = v210
            } else {
                v3 = 42
                v3 += v3
                for fi211 := 0; fi211 < 5; fi211++ {
                    for fi212 := 0; fi212 < 13; fi212++ {
                        if (true && ((127 < 1) && (!true))) {
                            var v213 uint = 4294967295
                            var v214 int8 = -128
                            _ = v213
                            _ = v214
                        } else {
                            _ = len(v1)
                            _ = len(v2)
                        }
                        v215 := GenS15{F0: 42, F1: 42}
                        var v216 bool = false
                        _ = fi212
                        _ = v215
                        _ = v216
                    }
                    for fi217 := 0; fi217 < 9; fi217++ {
                        v218 := GenS15{F0: 884, F1: uint8(v3)}
                        v3 += v3
                        var v219 int64 = -1
                        v220 := GenS24{F0: 42, F1: -2147483648, F2: (v3 & v3)}
                        _ = fi217
                        _ = v218
                        _ = v219
                        _ = v220
                    }
                    _ = fi211
                }
            }
        }
    case "\n\t":
        v221 := GenS25{F0: 65535, F1: 4294967295, F2: -3}
        _ = v221
    default:
        var v222 string = "\n\t"
        _ = v222
    }
    if true {
        v223 := GenS26{F0: 42, F1: 255}
        var v224 bool = false
        var v225 string = "a"
        _ = v223
        _ = v224
        _ = v225
    } else {
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v226 := GenS27{F0: 65535, F1: 928, F2: 42}
            switch ("" + (("a" + ("" + ("" + ""))) + ("hello" + "a"))) {
            case "foo bar":
                switch 381 {
                case 2147483647:
                    for fi227 := 0; fi227 < 20; fi227++ {
                        v3 = v3
                        _ = fi227
                    }
                case 9223372036854775807:
                    v2 = append(v2, 2147483647)
                case 0:
                    _ = len(v2)
                }
            case "\n\t":
                fmt.Println(v3)
            }
            _ = v226
        }
        for fi228 := 0; fi228 < 18; fi228++ {
            var v229 uint16 = 2
            _ = fi228
            _ = v229
        }
        if ((255 == 42) && true) {
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
            switch 9223372036854775807 {
            case 42:
                switch -9223372036854775808 {
                case 128:
                    for fi230 := 0; fi230 < 17; fi230++ {
                        for fi231 := 0; fi231 < 11; fi231++ {
                            var v232 int32 = 0
                            _ = fi231
                            _ = v232
                        }
                        _ = fi230
                    }
                case -165:
                    var v233 uint8 = 255
                    _ = v233
                }
            default:
                for fi234 := 0; fi234 < 13; fi234++ {
                    v235 := GenS25{F0: 2, F1: v3, F2: 1}
                    if (0 == 42) {
                        for fi236 := 0; fi236 < 5; fi236++ {
                            v237 := GenS8{F0: 1, F1: (v3 ^ (v3 & v3)), F2: -813}
                            _ = len(v2)
                            _ = fi236
                            _ = v237
                        }
                    }
                    _ = fi234
                    _ = v235
                }
            }
        } else {
            v1 = append(v1, -9223372036854775808)
            var v238 int8 = 64
            _ = v238
        }
    }
    v3 = (^(v3))
    var v239 uint = 0
    v239 += 531
    for fi240 := 0; fi240 < 5; fi240++ {
        if ("a" != ("hello" + "\n\t")) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi241 := 0; fi241 < 2; fi241++ {
                var v242 uint64 = v3
                _ = fi241
                _ = v242
            }
        }
        var v243 int8 = 73
        var v244 string = ("foo bar" + ("foo bar" + "hello"))
        _ = fi240
        _ = v243
        _ = v244
    }
    for fi245 := 0; fi245 < 4; fi245++ {
        var v246 int64 = 42
        fmt.Println(v246, v3, v246)
        _ = fi245
        _ = v246
    }
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v2 = append(v2, 1)
        if len(v2) > 0 {
            v2[0] = genFn0()
        }
    }
    var v247 uint = 2
    v3 = (^(v3))
    if true {
        var v248 int8 = -1
        var v249 float64 = -2.25
        v250 := GenS28{F0: 42, F1: 18446744073709551615}
        _ = v248
        _ = v249
        _ = v250
    } else {
        v3 = 2
        for fi251 := 0; fi251 < 5; fi251++ {
            v239 = (v239 - (v239 / ((1) | 1)))
            if true {
                v3 = (^(v3))
                v3 -= (v3 + 1)
                for fi252 := 0; fi252 < 1; fi252++ {
                    v2 = append(v2, 0)
                    v2 = append(v2, 544)
                    _ = fi252
                }
            }
            _ = fi251
        }
    }
    var v253 int8 = 104
    if (!false) {
        v239 -= v239
        v253 += v253
        for fi254 := 0; fi254 < 3; fi254++ {
            v255 := GenS13{F0: uint32(v239), F1: 18446744073709551615, F2: "\n\t"}
            switch fi254 {
            case -1:
                for fi256 := 0; fi256 < 7; fi256++ {
                    fmt.Println(v247)
                    v239 -= (v247 - 0)
                    if true {
                        if len(v1) > 0 {
                            v1[0] = fi256
                        }
                        switch "foo bar" {
                        case "foo bar":
                            fmt.Println(v253, v253)
                        case "\n\t":
                            switch "foo bar" {
                            case "\n\t":
                                var v257 uint = 194
                                _ = v257
                            case "a":
                                if len(v2) > 0 {
                                    v2[0] = 65535
                                }
                            case "foo bar":
                                fmt.Println(v3, v253)
                            default:
                                if true {
                                    fmt.Println(v253)
                                    _ = len(v1)
                                }
                            }
                        case "hello":
                            var v258 int8 = (v253 - (^(v253)))
                            _ = v258
                        }
                    }
                    var v259 int = (fi256 ^ (fi254 ^ (fi256 | 127)))
                    v247 = v239
                    _ = fi256
                    _ = v259
                }
            default:
                v253 += v253
            }
            _ = fi254
            _ = v255
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v239:", v239)
    fmt.Println("v247:", v247)
    fmt.Println("v253:", v253)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v209
    _ = v239
    _ = v247
    _ = v253
    _ = v3
}
