package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 bool; F2 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 string; F2 string }
type GenS6 struct { F0 int8; F1 string; F2 uint32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint32; F1 string }
type GenS9 struct { F0 int8; F1 string; F2 uint }
type GenS10 struct { F0 int32; F1 uint32 }
type GenS11 struct { F0 int16; F1 uint; F2 int64 }
type GenS12 struct { F0 bool; F1 uint8; F2 int8 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int32; F1 int8; F2 uint }
type GenS15 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS16 struct { F0 uint64; F1 bool }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 bool; F1 uint32 }
type GenS19 struct { F0 uint16; F1 string; F2 string }
type GenS20 struct { F0 uint32 }
type GenS21 struct { F0 uint16; F1 uint64 }
type GenS22 struct { F0 int64; F1 uint32 }
type GenS23 struct { F0 uint; F1 uint32; F2 int }
type GenS24 struct { F0 int32; F1 string }
type GenS25 struct { F0 int8; F1 bool }
type GenS26 struct { F0 uint64; F1 bool }
type GenS27 struct { F0 int16 }
type GenS28 struct { F0 uint16; F1 int16 }
func genFn0(p0 uint8, p1 uint) uint8 {
    var v6 int16 = 240
    v7 := []int{}
    p0 = (p0 & p0)
    if (2147483647 > 42) {
        v8 := GenS1{F0: (v6 << (((p1 >> (((p1 / ((212) | 1))) & 63))) & 63))}
        fmt.Println(p0)
        v9 := GenS2{F0: 1}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (p0 ^ p0)
}
func genFn1(p0 int64) int16 {
    p0 += 1
    if false {
        p0 -= p0
        switch -9223372036854775808 {
        case -793:
            fmt.Println(p0)
        case 59:
            p0 = p0
        }
        if false {
            var v10 float64 = 1.5
            _ = v10
        }
    }
    if true {
        if ((!true) && true) {
            p0 = 9223372036854775807
            if (true || (1.5 != float32(1e10))) {
                for fi11 := 0; fi11 < 8; fi11++ {
                    if false {
                        p0 -= (p0 + p0)
                        switch fi11 {
                        case -969:
                            var v12 uint16 = 0
                            _ = v12
                        }
                        var v13 uint64 = 927
                        _ = v13
                    } else {
                        for fi14 := 0; fi14 < 6; fi14++ {
                            if ((56 >= 2147483648) || true) {
                                if false {
                                    if (2 != 162) {
                                        p0 -= 0
                                    } else {
                                        switch 0 {
                                        case 32767:
                                            switch ("hello" + "a") {
                                            case "hello":
                                                var v15 uint64 = 573
                                                _ = v15
                                            default:
                                                v16 := GenS3{F0: 171, F1: (1 == 605), F2: (!(fi14 != fi14))}
                                                _ = v16
                                            }
                                        case 0:
                                            var v17 string = ("" + "")
                                            _ = v17
                                        default:
                                            if false {
                                                p0 = 9223372036854775807
                                                v18 := []int{fi14, fi11, 42, fi14}
                                                p0 = (p0 | p0)
                                                _ = v18
                                            } else {
                                                var v19 bool = false
                                                _ = v19
                                            }
                                        }
                                        v20 := GenS4{F0: -230}
                                        _ = v20
                                    }
                                    v21 := GenS5{F0: 77, F1: ("hello" + ("foo bar" + "hello")), F2: "foo bar"}
                                    _ = v21
                                } else {
                                    var v22 bool = true
                                    v23 := []int{fi14, (fi11 | (fi14 - 32767)), -2147483648}
                                    p0 -= p0
                                    _ = v22
                                    _ = v23
                                }
                            } else {
                                fmt.Println(fi14)
                            }
                            var v24 int = (fi11 << ((938) & 63))
                            if ((!false) && (0.0 != 0.1)) {
                                switch "\n\t" {
                                case "foo bar":
                                    if (0 == 65535) {
                                        v24 = -9223372036854775808
                                        v24 -= (v24 * (fi11 % ((fi11) | 1)))
                                    }
                                }
                                switch fi11 {
                                case 42:
                                    var v25 float32 = 0.1
                                    _ = v25
                                }
                                var v26 int16 = -32768
                                _ = v26
                            }
                            _ = fi14
                            _ = v24
                        }
                    }
                    fmt.Println(p0, fi11)
                    _ = fi11
                }
                p0 -= (p0 >> ((18446744073709551615) & 63))
                v27 := GenS2{F0: -2147483648}
                _ = v27
            } else {
                v28 := GenS4{F0: -9223372036854775808}
                _ = v28
            }
        } else {
            p0 = (^(p0))
        }
        var v29 float64 = 0.0
        _ = v29
    }
    for fi30 := 0; fi30 < 7; fi30++ {
        v31 := GenS6{F0: 127, F1: "hello", F2: 555}
        v32 := []int{1, -141}
        _ = fi30
        _ = v31
        _ = v32
    }
    _ = p0
    return -199
}

func main() {
    var v1 string = ("foo bar" + ("hello" + (("foo bar" + ("a" + "a")) + "")))
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    v4 := make(map[int]int)
    var v5 string = (v1 + "")
    for fi33 := 0; fi33 < 17; fi33++ {
        for fi34 := 0; fi34 < 5; fi34++ {
            if true {
                var v35 uint16 = 42
                _ = v35
            } else {
                var v36 uint8 = 42
                _ = len(v3)
                fmt.Println(v36)
                if (false || true) {
                    var v37 uint8 = 128
                    var v38 uint16 = 65535
                    _ = v37
                    _ = v38
                } else {
                    if (-2.25 < -2.25) {
                        v39 := []int{(fi34 * (fi34 / (((fi33 << ((2) & 63))) | 1))), fi34}
                        v1 = "\n\t"
                        _ = v39
                    }
                    fmt.Println(v36, fi33)
                }
                _ = v36
            }
            if true {
                if false {
                    fmt.Println(fi34, fi33, v1)
                    fmt.Println(v1, fi34)
                    v1 = ((("foo bar" + v1) + v1) + v1)
                } else {
                    v5 = ""
                }
                v5 = "hello"
                v5 = (v5 + "\n\t")
            }
            _ = fi34
        }
        if (true || (((v5 + v1) != v5) && true)) {
            var v40 uint8 = 55
            var v41 uint16 = 2
            _ = v40
            _ = v41
        } else {
            if (!(1000 == 32767)) {
                var v42 float32 = 0.1
                for fi43 := 0; fi43 < 13; fi43++ {
                    v4[fi43] = (fi43 >> ((255) & 63))
                    v44 := GenS3{F0: 624, F1: false, F2: false}
                    fmt.Println(v1, v42)
                    _ = fi43
                    _ = v44
                }
                v2[32768] = ((v5 + v5) != v5)
                _ = v42
            }
        }
        _ = fi33
    }
    v45 := GenS5{F0: 930, F1: (v5 + v5), F2: (v5 + ((v5 + v5) + "foo bar"))}
    v1 = v5
    v1 = ""
    if false {
        v5 = v1
        v46 := GenS1{F0: -32768}
        fmt.Println(v5)
        _ = v46
    } else {
        v4[-9223372036854775808] = 4294967295
        v47 := GenS1{F0: -764}
        _ = v47
    }
    switch ((v5 + (v1 + v5)) + ((("hello" + (v5 + v1)) + "a") + "foo bar")) {
    case "\n\t":
        var v48 int = -2147483648
        _ = v48
    case "":
        v1 = (v1 + ("hello" + "\n\t"))
    }
    switch (("\n\t" + v1) + v1) {
    case "foo bar":
        delete(v3, 466)
    }
    for fi49 := 0; fi49 < 16; fi49++ {
        v1 = v5
        if (238 > 778) {
            for fi50 := 0; fi50 < 15; fi50++ {
                if (genFn1(42) != -895) {
                    v51 := GenS7{F0: 30}
                    _ = v51
                }
                var v52 float64 = 1.5
                v53 := GenS1{F0: -1}
                _ = fi50
                _ = v52
                _ = v53
            }
            var v54 int8 = 100
            var v55 int64 = 0
            _ = v54
            _ = v55
        } else {
            for fi56 := 0; fi56 < 6; fi56++ {
                v57 := GenS8{F0: 437, F1: (("a" + v5) + v5)}
                _ = fi56
                _ = v57
            }
            if (127 != 255) {
                fmt.Println(v5, fi49, fi49)
                for fi58 := 0; fi58 < 4; fi58++ {
                    var v59 int32 = 772
                    _ = fi58
                    _ = v59
                }
                v5 = ("foo bar" + "\n\t")
            }
            v60 := []int{2147483647, 276}
            _ = v60
        }
        _ = fi49
    }
    for fi61 := 0; fi61 < 8; fi61++ {
        switch (fi61 ^ 32768) {
        case 915:
            v3[fi61] = 1
        case 1:
            var v62 uint64 = 597
            _ = v62
        case 351:
            v63 := GenS2{F0: fi61}
            _ = v63
        }
        _ = fi61
    }
    v64 := GenS7{F0: -87}
    if true {
        switch "\n\t" {
        case "a":
            var v65 uint = 42
            _ = v65
        case "\n\t":
            var v66 bool = (100 <= 1)
            _ = v66
        }
        if true {
            if (1 != -112) {
                v1 = "foo bar"
            } else {
                for fi67 := 0; fi67 < 4; fi67++ {
                    switch "hello" {
                    case "hello":
                        for fi68 := 0; fi68 < 4; fi68++ {
                            v5 = (("\n\t" + "hello") + ("" + "hello"))
                            _ = fi68
                        }
                    case "":
                        v5 = ((v5 + "hello") + "hello")
                    case "a":
                        v5 = (v5 + v1)
                    }
                    for fi69 := 0; fi69 < 16; fi69++ {
                        v5 = v1
                        switch v5 {
                        case "foo bar":
                            v1 = v1
                        case "\n\t":
                            var v70 uint32 = 2147483647
                            _ = v70
                        case "a":
                            if (-2.25 > 0.0) {
                                if true {
                                    var v71 uint = 2
                                    delete(v4, (fi67 * (fi67 / (((fi69 >> ((1) & 63))) | 1))))
                                    if false {
                                        v72 := GenS9{F0: -83, F1: v5, F2: (^(v71))}
                                        switch -954 {
                                        case 32767:
                                            v5 = v1
                                        case 0:
                                            var v73 int64 = 691
                                            _ = v73
                                        default:
                                            var v74 float32 = (1.5 + 0.0)
                                            _ = v74
                                        }
                                        var v75 uint8 = 195
                                        switch "\n\t" {
                                        case "hello":
                                            var v76 bool = (127 != 1)
                                            _ = v76
                                        case "foo bar":
                                            var v77 float32 = 3.14159
                                            _ = v77
                                        case "a":
                                            v78 := GenS9{F0: 42, F1: "", F2: (v71 / ((v71) | 1))}
                                            _ = v78
                                        }
                                        var v79 int32 = -2147483648
                                        _ = v72
                                        _ = v75
                                        _ = v79
                                    }
                                    _ = v71
                                }
                            } else {
                                v1 = v1
                                var v80 int64 = -9223372036854775808
                                if ((!(true || (4294967295 <= 1))) && ((("\n\t" + "") == ("\n\t" + ((("a" + "a") + v1) + "hello"))) || ((v5 + ("a" + ((("hello" + "hello") + "\n\t") + ((("" + (((v5 + v1) + v1) + "")) + (((("" + v1) + v5) + "hello") + "foo bar")) + ((v5 + v1) + ("hello" + ("\n\t" + v5))))))) < v5))) {
                                    if true {
                                        switch 2147483647 {
                                        case 32768:
                                            fmt.Println(v5, fi69)
                                        case 34:
                                            var v81 int16 = 1
                                            _ = v81
                                        case -1:
                                            var v82 int32 = 2147483647
                                            _ = v82
                                        default:
                                            fmt.Println(fi67)
                                        }
                                        v80 -= v80
                                        switch (("hello" + (("a" + v1) + (v5 + "a"))) + "\n\t") {
                                        case "foo bar":
                                            v1 = (("hello" + "a") + ((v1 + "foo bar") + (v1 + v1)))
                                            var v83 uint = 348
                                            _ = v83
                                        default:
                                            v84 := []int{}
                                            _ = v84
                                        }
                                        v85 := GenS10{F0: 0, F1: 1}
                                        v86 := v2[(fi67 % (((fi69 % ((fi69) | 1))) | 1))]
                                        _ = v85
                                        _ = v86
                                    }
                                    switch v1 {
                                    case "foo bar":
                                        v3[fi67] = 1
                                    default:
                                        switch "" {
                                        case "\n\t":
                                            var v87 bool = true
                                            _ = v87
                                        default:
                                            v4[fi67] = fi67
                                        }
                                    }
                                } else {
                                    v80 -= v80
                                    switch v5 {
                                    case "a":
                                        if ((v80 >= v80) || (42 > 255)) {
                                            v88 := GenS3{F0: 798, F1: (1.5 > ((1.5 * (0.1 / ((1e10 + (3.14159 / 1.0)) + (3.14159 + 3.14159)))) - 0.0)), F2: (65535 > 2)}
                                            if ((0.0 / 1e10) != (3.14159 + (1e10 * float64(-2.25)))) {
                                                var v89 uint32 = 469
                                                v90 := GenS11{F0: 42, F1: 1, F2: 9223372036854775807}
                                                _ = v89
                                                _ = v90
                                            }
                                            v1 = ("hello" + "\n\t")
                                            _ = v88
                                        }
                                    case "\n\t":
                                        v5 = v5
                                    case "":
                                        var v91 uint32 = 2147483647
                                        _ = v91
                                    default:
                                        if true {
                                            var v92 uint16 = 42
                                            v1 = ((v5 + v1) + ("hello" + v5))
                                            v93 := GenS2{F0: fi69}
                                            _ = len(v2)
                                            _ = v92
                                            _ = v93
                                        }
                                    }
                                    var v94 uint = 541
                                    v5 = v1
                                    switch (v5 + "a") {
                                    case "foo bar":
                                        fmt.Println(fi67, v5, v5)
                                    }
                                    _ = v94
                                }
                                _ = v80
                            }
                        }
                        _ = fi69
                    }
                    v1 = "\n\t"
                    _ = fi67
                }
                if true {
                    switch v5 {
                    case "\n\t":
                        if false {
                            fmt.Println(v1)
                            var v95 int64 = 0
                            v96 := GenS12{F0: (((float64((((-2.25 + 3.14159) * 0.0) * (-2.25 + 1e10))) * float64(0.0)) <= 0.1) || true), F1: 2, F2: -1}
                            _ = v95
                            _ = v96
                        } else {
                            v97 := GenS3{F0: 18446744073709551615, F1: false, F2: (0 > 481)}
                            _ = v97
                        }
                    case "":
                        if (genFn0(34, 1) == 255) {
                            if (!false) {
                                v98 := GenS4{F0: -9223372036854775808}
                                for fi99 := 0; fi99 < 9; fi99++ {
                                    fmt.Println(v5)
                                    v1 = (v5 + (v1 + v1))
                                    _ = fi99
                                }
                                _ = v98
                            }
                            v100 := []int{2147483647, 42}
                            _ = v100
                        }
                    case "hello":
                        if true {
                            switch 127 {
                            case 128:
                                v101 := []int{}
                                _ = v101
                            case -2147483648:
                                v1 = v5
                            case 637:
                                if false {
                                    fmt.Println(v5)
                                } else {
                                    v102 := GenS2{F0: 128}
                                    _ = v102
                                }
                            }
                        }
                    default:
                        v103 := GenS4{F0: 32768}
                        _ = v103
                    }
                }
                for fi104 := 0; fi104 < 20; fi104++ {
                    _ = len(v4)
                    v5 = v5
                    if true {
                        if true {
                            v1 = v1
                            v5 = (("" + v1) + (v1 + (v5 + ("\n\t" + (v1 + v5)))))
                            v1 = (("" + "") + ("hello" + ("\n\t" + (v5 + (v5 + v5)))))
                        } else {
                            v1 = ""
                            for fi105 := 0; fi105 < 6; fi105++ {
                                if ((true || false) || false) {
                                    fmt.Println(v1, fi105, v5)
                                    switch (fi104 * 127) {
                                    case 2147483647:
                                        if false {
                                            switch ((("hello" + "a") + (((v5 + "a") + "a") + "\n\t")) + v1) {
                                            case "foo bar":
                                                v106 := GenS2{F0: -615}
                                                _ = v106
                                            case "\n\t":
                                                var v107 uint8 = 42
                                                _ = v107
                                            default:
                                                var v108 int8 = 100
                                                _ = v108
                                            }
                                            if ((!(false || false)) || (!((1e10 == 3.14159) || ((fi105 / ((fi104) | 1)) > (fi104 & 32768))))) {
                                                v1 = v1
                                                if false {
                                                    v109 := []int{fi105}
                                                    v110 := GenS13{F0: 0}
                                                    v111 := GenS5{F0: 42, F1: "foo bar", F2: (("" + v1) + v5)}
                                                    _ = v109
                                                    _ = v110
                                                    _ = v111
                                                }
                                                var v112 uint8 = 42
                                                v1 = ((v5 + v1) + "a")
                                                _ = v112
                                            } else {
                                                v1 = "foo bar"
                                                v113 := GenS13{F0: 0}
                                                _ = v113
                                            }
                                            switch fi105 {
                                            case -293:
                                                fmt.Println(fi105, fi104)
                                            case -1:
                                                fmt.Println(fi105, fi105)
                                            }
                                        }
                                    default:
                                        v114 := GenS14{F0: -1, F1: 1, F2: 18446744073709551615}
                                        _ = v114
                                    }
                                } else {
                                    var v115 int8 = 127
                                    if (!true) {
                                        var v116 uint16 = 751
                                        _ = v116
                                    } else {
                                        var v117 int32 = 1
                                        switch fi104 {
                                        case -2147483648:
                                            v118 := GenS7{F0: v115}
                                            _ = v118
                                        case 32768:
                                            if false {
                                                switch 127 {
                                                case 42:
                                                    fmt.Println(v115, fi104)
                                                case 32768:
                                                    v119 := GenS15{F0: 42, F1: 4294967295, F2: 121}
                                                    _ = v119
                                                }
                                                v120 := GenS5{F0: 4294967295, F1: v5, F2: ("\n\t" + ("\n\t" + v5))}
                                                v1 = v1
                                                _ = v120
                                            } else {
                                                v121 := []int{(fi105 - fi105), -2147483648, (fi105 << ((673) & 63)), fi104}
                                                v5 = v5
                                                v115 = 0
                                                _ = v121
                                            }
                                        case 110:
                                            v122 := GenS16{F0: 0, F1: false}
                                            _ = v122
                                        default:
                                            if (((fi105 * fi104) <= (^(fi105))) || false) {
                                                if (880 < 0) {
                                                    var v123 uint16 = 35
                                                    _ = v123
                                                } else {
                                                    if (1 == 2147483647) {
                                                        fmt.Println(fi105, fi105, v1)
                                                    } else {
                                                        v124 := []int{fi105, -429, fi105, (-(fi104))}
                                                        v5 = v1
                                                        _ = v124
                                                    }
                                                    if false {
                                                        v115 += v115
                                                        if true {
                                                            fmt.Println(v117, fi104)
                                                            fmt.Println(fi104, v5, v5)
                                                            v125 := GenS17{F0: 12}
                                                            _ = v125
                                                        }
                                                    }
                                                    var v126 int64 = 42
                                                    _ = v126
                                                }
                                                _ = len(v2)
                                            } else {
                                                if false {
                                                    v117 += (v117 + 42)
                                                    v1 = v1
                                                    v127 := GenS1{F0: 32767}
                                                    var v128 string = v5
                                                    if (!((!(1 > 42)) && (339 > 391))) {
                                                        _ = len(v3)
                                                        v5 = v128
                                                    } else {
                                                        v5 = v128
                                                        if (!false) {
                                                            if false {
                                                                fmt.Println(v115)
                                                                var v129 float64 = -2.25
                                                                _ = v129
                                                            } else {
                                                                if true {
                                                                    var v130 string = ("foo bar" + v1)
                                                                    if ((v115 << ((65535) & 63)) > (v115 >> ((4294967295) & 63))) {
                                                                        var v131 uint64 = 1
                                                                        fmt.Println(v131, v115, v5)
                                                                        _ = v131
                                                                    } else {
                                                                        switch 709 {
                                                                        case 0:
                                                                            v132 := []int{(^(fi105)), fi104}
                                                                            _ = v132
                                                                        default:
                                                                            switch (("foo bar" + (v1 + (v130 + v5))) + (v130 + "foo bar")) {
                                                                            case "a":
                                                                                if true {
                                                                                    if false {
                                                                                        switch "\n\t" {
                                                                                        case "hello":
                                                                                            var v133 int32 = v117
                                                                                            _ = v133
                                                                                        case "a":
                                                                                            switch v128 {
                                                                                            case "a":
                                                                                                v134 := []int{32767, fi105}
                                                                                                _ = v134
                                                                                            default:
                                                                                                var v135 float64 = (0.1 + 1.5)
                                                                                                _ = v135
                                                                                            }
                                                                                        case "foo bar":
                                                                                            fmt.Println(v117, v130, v5)
                                                                                        default:
                                                                                            var v136 uint16 = 271
                                                                                            _ = v136
                                                                                        }
                                                                                    }
                                                                                    switch -442 {
                                                                                    case -9223372036854775808:
                                                                                        if (true && true) {
                                                                                            if false {
                                                                                                v5 = v5
                                                                                                if true {
                                                                                                    var v137 int8 = -128
                                                                                                    var v138 int64 = -614
                                                                                                    _ = v137
                                                                                                    _ = v138
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        v128 = v130
                                                                                    }
                                                                                    var v139 int8 = (v115 / ((v115) | 1))
                                                                                    if (279 <= -1) {
                                                                                        var v140 uint32 = 65535
                                                                                        v141 := GenS4{F0: (fi104 / ((999) | 1))}
                                                                                        _ = v140
                                                                                        _ = v141
                                                                                    }
                                                                                    v2[fi104] = true
                                                                                    _ = v139
                                                                                }
                                                                                switch v1 {
                                                                                case "hello":
                                                                                    switch fi105 {
                                                                                    case -9223372036854775808:
                                                                                        v142 := GenS11{F0: -1, F1: 65535, F2: 274}
                                                                                        v143 := GenS18{F0: true, F1: 4294967295}
                                                                                        _ = v142
                                                                                        _ = v143
                                                                                    }
                                                                                case "foo bar":
                                                                                    v144 := GenS19{F0: 255, F1: (v128 + ((((v130 + "\n\t") + v128) + ("\n\t" + v130)) + "foo bar")), F2: v1}
                                                                                    _ = v144
                                                                                case "":
                                                                                    if ((!(-9223372036854775808 > 9223372036854775807)) || false) {
                                                                                        fmt.Println(v1)
                                                                                        var v145 float64 = (((1e10 + 0.0) - 0.0) / (((1e10 - (0.1 * (1.5 - 0.1))) + 0.1) + (float64(0.0) / float64(float32(1e10)))))
                                                                                        delete(v2, (fi105 % (((fi105 % ((fi105) | 1))) | 1)))
                                                                                        v146 := GenS20{F0: 0}
                                                                                        v147 := GenS13{F0: 2}
                                                                                        _ = v145
                                                                                        _ = v146
                                                                                        _ = v147
                                                                                    } else {
                                                                                        if true {
                                                                                            v148 := []int{(fi105 | -2147483648)}
                                                                                            switch (fi105 & fi104) {
                                                                                            case 354:
                                                                                                fmt.Println(v117)
                                                                                                if (0 >= -1) {
                                                                                                    v149 := GenS21{F0: 0, F1: 341}
                                                                                                    fmt.Println(v5, fi104)
                                                                                                    v150 := GenS22{F0: 1, F1: 4294967295}
                                                                                                    if true {
                                                                                                        switch (fi105 * fi105) {
                                                                                                        case -1:
                                                                                                            delete(v4, (fi104 << ((262) & 63)))
                                                                                                        case -674:
                                                                                                            v151 := GenS8{F0: 2147483648, F1: "hello"}
                                                                                                            _ = v151
                                                                                                        default:
                                                                                                            if (true || ("foo bar" == ("foo bar" + ((((v130 + (v128 + (("foo bar" + ((v128 + (((v128 + (v130 + (v128 + "foo bar"))) + (v128 + "a")) + "foo bar")) + v1)) + (v5 + (v130 + v1))))) + v130) + "") + "hello")))) {
                                                                                                                v152 := GenS8{F0: 1, F1: v1}
                                                                                                                v148 = append(v148, (-(fi104)))
                                                                                                                var v153 int = 127
                                                                                                                _ = v152
                                                                                                                _ = v153
                                                                                                            } else {
                                                                                                                v5 = "hello"
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v149
                                                                                                    _ = v150
                                                                                                } else {
                                                                                                    if ((2 > 255) || ((-81 != 0) || false)) {
                                                                                                        var v154 int = 4294967295
                                                                                                        var v155 bool = false
                                                                                                        v117 -= v117
                                                                                                        _ = v154
                                                                                                        _ = v155
                                                                                                    } else {
                                                                                                        v5 = v5
                                                                                                    }
                                                                                                    if len(v148) >= 2 {
                                                                                                        v148 = v148[1:len(v148)]
                                                                                                    }
                                                                                                }
                                                                                                var v156 int64 = 1
                                                                                                _ = v156
                                                                                            case 9223372036854775807:
                                                                                                var v157 int8 = (v115 << ((254) & 63))
                                                                                                v128 = ((v128 + v130) + "foo bar")
                                                                                                var v158 int8 = (v115 - v115)
                                                                                                _ = v157
                                                                                                _ = v158
                                                                                            default:
                                                                                                switch 4294967295 {
                                                                                                case 4294967295:
                                                                                                    switch "foo bar" {
                                                                                                    case "hello":
                                                                                                        fmt.Println(v130, v5, fi104)
                                                                                                    default:
                                                                                                        fmt.Println(v1, v115, v130)
                                                                                                    }
                                                                                                case 32767:
                                                                                                    v159 := v3[fi105]
                                                                                                    _ = v159
                                                                                                default:
                                                                                                    var v160 int64 = 1
                                                                                                    _ = v160
                                                                                                }
                                                                                            }
                                                                                            _ = len(v2)
                                                                                            _ = v148
                                                                                        }
                                                                                        v161 := v2[(fi105 & (fi105 >> ((42) & 63)))]
                                                                                        if (920 != 255) {
                                                                                            fmt.Println(v161, v115, v5)
                                                                                            fmt.Println(v130, v5)
                                                                                            switch fi105 {
                                                                                            case -1:
                                                                                                if v161 {
                                                                                                    var v162 uint16 = 2
                                                                                                    if (1 <= genFn0(39, 18446744073709551615)) {
                                                                                                        v1 = "hello"
                                                                                                        v130 = (("foo bar" + (((v128 + v1) + v1) + "\n\t")) + v5)
                                                                                                        v1 = ""
                                                                                                    } else {
                                                                                                        var v163 int = (fi105 / ((fi105) | 1))
                                                                                                        _ = v163
                                                                                                    }
                                                                                                    _ = v162
                                                                                                } else {
                                                                                                    v130 = v128
                                                                                                    var v164 float32 = 1.5
                                                                                                    if ((2 <= 2) || (v161 && true)) {
                                                                                                        var v165 string = v5
                                                                                                        v166 := []int{}
                                                                                                        _ = v165
                                                                                                        _ = v166
                                                                                                    }
                                                                                                    _ = v164
                                                                                                }
                                                                                            case 84:
                                                                                                switch v5 {
                                                                                                case "foo bar":
                                                                                                    var v167 uint8 = genFn0(42, 18446744073709551615)
                                                                                                    _ = v167
                                                                                                case "":
                                                                                                    fmt.Println(v5, v130, v161)
                                                                                                default:
                                                                                                    fmt.Println(v128)
                                                                                                    var v168 uint32 = 42
                                                                                                    v3[(fi104 / (((-(fi105))) | 1))] = (v168 | 42)
                                                                                                    v169 := []int{}
                                                                                                    v5 = "\n\t"
                                                                                                    _ = v168
                                                                                                    _ = v169
                                                                                                }
                                                                                            }
                                                                                            v170 := GenS23{F0: 662, F1: 65535, F2: 127}
                                                                                            fmt.Println(fi105)
                                                                                            _ = v170
                                                                                        }
                                                                                        _ = v161
                                                                                    }
                                                                                default:
                                                                                    switch v128 {
                                                                                    case "hello":
                                                                                        fmt.Println(v5, v5, v130)
                                                                                    case "":
                                                                                        switch (("hello" + v128) + v1) {
                                                                                        case "a":
                                                                                            if ((!false) || (true && false)) {
                                                                                                fmt.Println(v117, v128, fi105)
                                                                                                var v171 float64 = 1.5
                                                                                                _ = v171
                                                                                            }
                                                                                        case "foo bar":
                                                                                            if true {
                                                                                                v128 = v5
                                                                                                var v172 bool = true
                                                                                                var v173 int64 = 1
                                                                                                v5 = "\n\t"
                                                                                                _ = v172
                                                                                                _ = v173
                                                                                            }
                                                                                        case "hello":
                                                                                            v115 -= 42
                                                                                        default:
                                                                                            v1 = v130
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case "":
                                                                                if false {
                                                                                    delete(v2, -9223372036854775808)
                                                                                    v1 = v1
                                                                                    var v174 uint32 = 42
                                                                                    v130 = ((v5 + "hello") + "foo bar")
                                                                                    _ = v174
                                                                                }
                                                                            case "\n\t":
                                                                                v1 = "\n\t"
                                                                            default:
                                                                                if (9 == (-(v115))) {
                                                                                    var v175 string = ""
                                                                                    _ = v175
                                                                                } else {
                                                                                    switch "foo bar" {
                                                                                    case "a":
                                                                                        v176 := v2[32768]
                                                                                        v177 := []int{-244}
                                                                                        _ = v176
                                                                                        _ = v177
                                                                                    case "foo bar":
                                                                                        var v178 float64 = 1.5
                                                                                        _ = v178
                                                                                    }
                                                                                    if (67 == genFn0(0, 42)) {
                                                                                        delete(v3, -787)
                                                                                        var v179 int64 = -1
                                                                                        var v180 uint8 = genFn0(127, 2)
                                                                                        _ = v179
                                                                                        _ = v180
                                                                                    } else {
                                                                                        switch v128 {
                                                                                        case "a":
                                                                                            var v181 float64 = 1e10
                                                                                            _ = v181
                                                                                        case "\n\t":
                                                                                            v115 -= v115
                                                                                        }
                                                                                        var v182 float64 = 3.14159
                                                                                        if (v117 == v117) {
                                                                                            var v183 uint16 = 32767
                                                                                            v184 := GenS24{F0: (v117 << ((255) & 63)), F1: v130}
                                                                                            _ = v183
                                                                                            _ = v184
                                                                                        } else {
                                                                                            v185 := GenS25{F0: v115, F1: true}
                                                                                            _ = v185
                                                                                        }
                                                                                        _ = v182
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        v117 = v117
                                                                        var v186 int32 = 42
                                                                        _ = v186
                                                                    }
                                                                    _ = v130
                                                                } else {
                                                                    var v187 uint64 = 240
                                                                    if (710 < 255) {
                                                                        if (127 >= 128) {
                                                                            if true {
                                                                                v188 := []int{fi104}
                                                                                switch (fi105 / ((32768) | 1)) {
                                                                                case 127:
                                                                                    switch "" {
                                                                                    case "foo bar":
                                                                                        if true {
                                                                                            v3[334] = 732
                                                                                            v189 := GenS26{F0: v187, F1: false}
                                                                                            _ = v189
                                                                                        }
                                                                                    case "":
                                                                                        v115 = v115
                                                                                    case "\n\t":
                                                                                        switch 2147483647 {
                                                                                        case 319:
                                                                                            if (3.14159 != 0.1) {
                                                                                                v190 := GenS2{F0: fi105}
                                                                                                _ = v190
                                                                                            } else {
                                                                                                if false {
                                                                                                    v191 := GenS23{F0: 42, F1: 302, F2: fi105}
                                                                                                    switch (fi104 % ((fi104) | 1)) {
                                                                                                    case 291:
                                                                                                        if false {
                                                                                                            var v192 int = fi105
                                                                                                            _ = len(v188)
                                                                                                            var v193 bool = (!false)
                                                                                                            _ = v192
                                                                                                            _ = v193
                                                                                                        }
                                                                                                    case -1:
                                                                                                        fmt.Println(fi104, v128)
                                                                                                    default:
                                                                                                        var v194 int64 = 1
                                                                                                        _ = v194
                                                                                                    }
                                                                                                    _ = v191
                                                                                                }
                                                                                                if false {
                                                                                                    fmt.Println(fi105, fi105)
                                                                                                    switch -1 {
                                                                                                    case 42:
                                                                                                        fmt.Println(v1, fi105)
                                                                                                    case 4294967295:
                                                                                                        v1 = v128
                                                                                                    case 337:
                                                                                                        var v195 uint = 2
                                                                                                        _ = v195
                                                                                                    }
                                                                                                    v187 = v187
                                                                                                } else {
                                                                                                    v1 = ("hello" + v1)
                                                                                                    if (0 >= 1) {
                                                                                                        v128 = v5
                                                                                                    } else {
                                                                                                        var v196 bool = (!(v115 > v115))
                                                                                                        _ = v196
                                                                                                    }
                                                                                                }
                                                                                                if ((0.1 * 3.14159) != 1.5) {
                                                                                                    var v197 bool = (!(!false))
                                                                                                    v117 = v117
                                                                                                    _ = v197
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            v128 = ""
                                                                                        }
                                                                                    }
                                                                                case -390:
                                                                                    var v198 uint = uint(v187)
                                                                                    _ = v198
                                                                                }
                                                                                if false {
                                                                                    fmt.Println(v1, v187)
                                                                                    fmt.Println(v5)
                                                                                    v187 += (v187 & (v187 % ((467) | 1)))
                                                                                } else {
                                                                                    if len(v188) > 0 {
                                                                                        v188[0] = 42
                                                                                    }
                                                                                }
                                                                                _ = v188
                                                                            } else {
                                                                                var v199 uint = 2
                                                                                v200 := GenS27{F0: 1}
                                                                                v117 -= (v117 << (((v199 % (((v199 << ((v199) & 63))) | 1))) & 63))
                                                                                _ = v199
                                                                                _ = v200
                                                                            }
                                                                        }
                                                                        switch (v128 + "hello") {
                                                                        case "":
                                                                            switch (fi104 & fi105) {
                                                                            case 128:
                                                                                switch (((v5 + (v1 + ((v128 + (("a" + "\n\t") + v5)) + v5))) + (v5 + v5)) + (("\n\t" + v5) + "")) {
                                                                                case "foo bar":
                                                                                    v117 = -2147483648
                                                                                case "":
                                                                                    fmt.Println(v128)
                                                                                }
                                                                            default:
                                                                                if (!((842 <= 65535) || ((v1 + "") < (("" + "a") + v128)))) {
                                                                                    var v201 float64 = 1.5
                                                                                    v1 = v1
                                                                                    _ = v201
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        switch "\n\t" {
                                                                        case "":
                                                                            v115 = v115
                                                                        case "foo bar":
                                                                            v202 := GenS24{F0: 367, F1: "\n\t"}
                                                                            _ = v202
                                                                        default:
                                                                            var v203 uint8 = 2
                                                                            _ = v203
                                                                        }
                                                                        v204 := []int{}
                                                                        v128 = ""
                                                                        v4[(-(fi104))] = fi104
                                                                        _ = v204
                                                                    }
                                                                    _ = v187
                                                                }
                                                                var v205 uint32 = 4294967295
                                                                _ = v205
                                                            }
                                                        } else {
                                                            var v206 uint8 = 0
                                                            v115 += v115
                                                            _ = v206
                                                        }
                                                        v207 := GenS24{F0: v117, F1: ((v1 + v1) + v128)}
                                                        _ = v207
                                                    }
                                                    var v208 int64 = 1
                                                    if (genFn1(v208) <= -1) {
                                                        v209 := GenS25{F0: (v115 - (v115 & 1)), F1: true}
                                                        _ = v209
                                                    }
                                                    _ = v127
                                                    _ = v128
                                                    _ = v208
                                                }
                                                v1 = ("foo bar" + v1)
                                                v210 := GenS5{F0: 0, F1: v1, F2: v1}
                                                _ = v210
                                            }
                                        }
                                        _ = v117
                                    }
                                    v211 := GenS4{F0: fi105}
                                    v5 = v5
                                    if (65535 == 1) {
                                        if true {
                                            fmt.Println(v5)
                                            var v212 float64 = 1.5
                                            _ = v212
                                        } else {
                                            v213 := v3[fi104]
                                            v214 := []int{}
                                            var v215 uint = 812
                                            _ = v213
                                            _ = v214
                                            _ = v215
                                        }
                                    } else {
                                        v216 := []int{-414, (fi104 | fi105), 9223372036854775807, (fi104 / ((-1) | 1))}
                                        fmt.Println(v5)
                                        var v217 bool = (!true)
                                        _ = v216
                                        _ = v217
                                    }
                                    _ = v115
                                    _ = v211
                                }
                                v218 := []int{}
                                if (!(!false)) {
                                    v219 := GenS12{F0: (1 < 361), F1: 18, F2: -128}
                                    v1 = "foo bar"
                                    _ = len(v218)
                                    _ = v219
                                } else {
                                    var v220 uint64 = 42
                                    _ = v220
                                }
                                _ = fi105
                                _ = v218
                            }
                            var v221 bool = (genFn1(-9223372036854775808) == 42)
                            _ = v221
                        }
                        v222 := GenS28{F0: 42, F1: -797}
                        _ = v222
                    } else {
                        v3[(fi104 % ((fi104) | 1))] = 2
                        v223 := []int{fi104, (^(fi104)), (fi104 & (-(fi104))), fi104}
                        fmt.Println(v5, v5)
                        _ = v223
                    }
                    _ = fi104
                }
            }
            v5 = (((("\n\t" + v1) + "") + ("\n\t" + v1)) + ("\n\t" + v5))
        } else {
            fmt.Println(v1, v5)
            var v224 int16 = 42
            _ = v224
        }
    }
    v225 := GenS5{F0: 465, F1: ((v1 + v1) + ""), F2: "\n\t"}
    fmt.Println(v5)
    v226 := []int{32767, 42}
    v226 = append(v226, 32767)
    for fi227 := 0; fi227 < 19; fi227++ {
        var v228 int8 = 42
        var v229 bool = false
        _ = fi227
        _ = v228
        _ = v229
    }
    v5 = (v1 + v1)
    switch -2147483648 {
    case 4294967295:
        v5 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v226:", len(v226), v226)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v225
    _ = v226
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v64
}
