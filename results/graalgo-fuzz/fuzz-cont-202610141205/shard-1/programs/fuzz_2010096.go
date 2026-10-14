package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint32 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 uint16; F2 string }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 uint16; F1 uint16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint8; F1 int32 }
type GenS11 struct { F0 bool; F1 int8 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int64; F1 int32 }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 int64; F1 uint16 }

func main() {
    var v1 int16 = 0
    var v2 uint32 = 4294967295
    var v3 int64 = -525
    var v4 string = "a"
    v5 := make(map[string]bool)
    for fi6 := 0; fi6 < 3; fi6++ {
        fmt.Println(v1, v2)
        var v7 int = (^(fi6))
        v2 -= v2
        _ = fi6
        _ = v7
    }
    v4 = v4
    v1 += -1
    if false {
        for fi8 := 0; fi8 < 19; fi8++ {
            v1 = (^(v1))
            v9 := GenS1{F0: 541, F1: (1 <= 69), F2: v2}
            v4 = v4
            _ = fi8
            _ = v9
        }
        switch 32767 {
        case 9223372036854775807:
            var v10 int32 = -411
            _ = v10
        case 2147483647:
            switch "hello" {
            case "\n\t":
                for fi11 := 0; fi11 < 6; fi11++ {
                    for fi12 := 0; fi12 < 14; fi12++ {
                        if (true && true) {
                            if false {
                                var v13 uint64 = 53
                                var v14 bool = (0 >= -272)
                                v3 = -1
                                _ = v13
                                _ = v14
                            }
                            v4 = "hello"
                            v15 := GenS1{F0: (v3 / (((v3 & (v3 % ((-1) | 1)))) | 1)), F1: false, F2: v2}
                            _ = v15
                        } else {
                            v5[v4] = (!(42 > -104))
                            if (166 >= 1) {
                                switch "foo bar" {
                                case "":
                                    v16 := GenS2{F0: fi11, F1: (v2 * v2)}
                                    _ = v16
                                default:
                                    v1 -= v1
                                }
                            } else {
                                if ((v2 ^ (^(v2))) > (v2 << ((2) & 63))) {
                                    v2 = v2
                                    if ((230 <= 207) && true) {
                                        if (!true) {
                                            if (!false) {
                                                v17 := GenS2{F0: fi11, F1: (v2 | v2)}
                                                v18 := GenS3{F0: 18446744073709551615}
                                                switch 357 {
                                                case -370:
                                                    if (42 != -845) {
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            if (v2 > v2) {
                                                                var v19 int = fi11
                                                                _ = v19
                                                            }
                                                        default:
                                                            var v20 int64 = v3
                                                            _ = v20
                                                        }
                                                    } else {
                                                        var v21 uint64 = 1
                                                        _ = v21
                                                    }
                                                case 127:
                                                    v22 := []int{127, -2147483648}
                                                    _ = v22
                                                case -416:
                                                    v23 := GenS4{F0: 104}
                                                    _ = v23
                                                }
                                                _ = v17
                                                _ = v18
                                            }
                                            if (-128 <= -125) {
                                                switch fi12 {
                                                case 2147483647:
                                                    var v24 float32 = (0.1 * 3.14159)
                                                    _ = v24
                                                case -752:
                                                    var v25 int64 = v3
                                                    _ = v25
                                                }
                                                switch "a" {
                                                case "foo bar":
                                                    switch (fi11 ^ (fi12 / (((fi11 >> ((886) & 63))) | 1))) {
                                                    case 0:
                                                        var v26 float32 = 0.0
                                                        _ = v26
                                                    case -9223372036854775808:
                                                        delete(v5, ((((("hello" + (("" + (v4 + v4)) + "foo bar")) + (v4 + "a")) + (v4 + ((("foo bar" + (v4 + "hello")) + v4) + v4))) + v4) + (v4 + "hello")))
                                                    }
                                                case "hello":
                                                    v1 = v1
                                                default:
                                                    v4 = ""
                                                }
                                                v5[v4] = ((v1 >> ((18446744073709551615) & 63)) < (v1 << ((510) & 63)))
                                            }
                                            v27 := GenS1{F0: 0, F1: true, F2: v2}
                                            fmt.Println(fi12, v4)
                                            _ = v27
                                        }
                                        v28 := GenS5{F0: -128, F1: 255, F2: v4}
                                        _ = v28
                                    } else {
                                        v3 = 1
                                    }
                                    switch v4 {
                                    case "\n\t":
                                        v29 := GenS5{F0: 100, F1: 74, F2: "hello"}
                                        _ = v29
                                    case "hello":
                                        fmt.Println(fi11, v1)
                                    case "foo bar":
                                        fmt.Println(v4, fi12, fi11)
                                    default:
                                        v30 := GenS2{F0: (fi12 + (^(fi12))), F1: 4294967295}
                                        _ = v30
                                    }
                                }
                                v31 := GenS4{F0: 255}
                                var v32 float64 = 1.5
                                _ = v31
                                _ = v32
                            }
                        }
                        switch fi11 {
                        case 32767:
                            v33 := GenS5{F0: -1, F1: 671, F2: (v4 + "foo bar")}
                            _ = v33
                        case 4294967295:
                            v34 := GenS4{F0: 0}
                            _ = v34
                        case 1:
                            if true {
                                fmt.Println(fi12, v2)
                                _ = len(v5)
                            }
                        }
                        _ = fi12
                    }
                    fmt.Println(v1, v1, fi11)
                    v35 := GenS2{F0: (fi11 / ((fi11) | 1)), F1: v2}
                    _ = fi11
                    _ = v35
                }
            case "a":
                _ = len(v5)
            case "foo bar":
                var v36 string = v4
                _ = v36
            }
        }
        switch "hello" {
        case "":
            for fi37 := 0; fi37 < 3; fi37++ {
                if false {
                    v4 = ""
                    if (-457 == (fi37 << ((2) & 63))) {
                        var v38 uint32 = (v2 ^ 42)
                        _ = v38
                    } else {
                        switch "a" {
                        case "":
                            v2 -= v2
                        case "a":
                            switch ((v4 + v4) + (((v4 + v4) + v4) + "")) {
                            case "foo bar":
                                v4 = ((v4 + v4) + ((v4 + v4) + v4))
                            case "hello":
                                v3 -= v3
                            default:
                                v39 := GenS2{F0: 2147483647, F1: (v2 % ((1) | 1))}
                                _ = v39
                            }
                        }
                        v40 := GenS6{F0: v4, F1: -70}
                        _ = v40
                    }
                    if ((!false) && false) {
                        switch (fi37 % ((-216) | 1)) {
                        case 128:
                            var v41 int32 = 2147483647
                            _ = v41
                        default:
                            if true {
                                for fi42 := 0; fi42 < 9; fi42++ {
                                    var v43 uint8 = 22
                                    _ = fi42
                                    _ = v43
                                }
                                var v44 uint32 = (v2 << ((806) & 63))
                                _ = v44
                            }
                        }
                        var v45 uint8 = 1
                        var v46 bool = (int16(v2) == -499)
                        _ = v45
                        _ = v46
                    }
                } else {
                    fmt.Println(v3, v1)
                    var v47 float64 = 1.5
                    _ = v47
                }
                fmt.Println(fi37, v4, v3)
                v4 = v4
                _ = fi37
            }
        default:
            v48 := []int{1, -9223372036854775808, 0}
            _ = v48
        }
    } else {
        if true {
            for fi49 := 0; fi49 < 20; fi49++ {
                v3 = v3
                v4 = (("hello" + "") + "a")
                v50 := GenS3{F0: 0}
                _ = fi49
                _ = v50
            }
            var v51 uint32 = (^(v2))
            _ = v51
        } else {
            for fi52 := 0; fi52 < 19; fi52++ {
                switch (v4 + v4) {
                case "\n\t":
                    v53 := v5[((v4 + v4) + "")]
                    _ = v53
                default:
                    for fi54 := 0; fi54 < 6; fi54++ {
                        var v55 int = fi52
                        switch v4 {
                        case "a":
                            if false {
                                fmt.Println(fi54, v1, v2)
                            } else {
                                var v56 float32 = 0.0
                                v57 := v5[v4]
                                _ = v56
                                _ = v57
                            }
                        case "hello":
                            v58 := []int{}
                            _ = v58
                        }
                        var v59 uint = 4294967295
                        _ = fi54
                        _ = v55
                        _ = v59
                    }
                }
                var v60 uint64 = 738
                _ = fi52
                _ = v60
            }
            if ((v1 - v1) <= (v1 | (v1 >> ((42) & 63)))) {
                v61 := GenS4{F0: 255}
                if true {
                    var v62 uint16 = 696
                    v3 += (v3 << ((255) & 63))
                    v63 := []int{365, 42}
                    _ = v62
                    _ = v63
                } else {
                    if (!false) {
                        if (false || ((!true) && (false && (1 < 4294967295)))) {
                            var v64 int32 = int32(v3)
                            v3 = (^(v3))
                            _ = v64
                        } else {
                            v4 = v4
                        }
                    } else {
                        var v65 uint64 = 2
                        _ = v65
                    }
                    switch 32767 {
                    case -483:
                        v1 -= (v1 << ((308) & 63))
                    case 876:
                        var v66 string = v4
                        _ = v66
                    case 4294967295:
                        for fi67 := 0; fi67 < 5; fi67++ {
                            var v68 uint16 = 1
                            switch v4 {
                            case "":
                                var v69 int16 = (-(v1))
                                _ = v69
                            case "foo bar":
                                v70 := GenS7{F0: 42, F1: (v68 + (^(v68)))}
                                _ = v70
                            }
                            fmt.Println(v3, v2)
                            _ = fi67
                            _ = v68
                        }
                    }
                    v1 += (v1 & 942)
                }
                if ((v3 >> ((2) & 63)) <= int64(v2)) {
                    var v71 bool = false
                    _ = v71
                }
                _ = v61
            }
            for fi72 := 0; fi72 < 14; fi72++ {
                for fi73 := 0; fi73 < 19; fi73++ {
                    v2 = (v2 / ((386) | 1))
                    _ = fi73
                }
                var v74 string = ("" + "foo bar")
                v75 := []int{127, int(v1), 127}
                _ = fi72
                _ = v74
                _ = v75
            }
        }
    }
    v76 := GenS8{F0: v4}
    var v77 uint32 = 4294967295
    var v78 uint = 2
    v79 := GenS5{F0: -42, F1: 255, F2: v4}
    delete(v5, (("" + v4) + v4))
    for fi80 := 0; fi80 < 1; fi80++ {
        v5[v4] = true
        _ = fi80
    }
    var v81 int8 = 100
    if (((v4 + v4) + v4) > (v4 + "\n\t")) {
        if (!(!(true && (1e10 > ((float64(0.1) * 1e10) * 3.14159))))) {
            fmt.Println(v78, v2)
        } else {
            _ = len(v5)
        }
    }
    for fi82 := 0; fi82 < 18; fi82++ {
        v83 := GenS6{F0: ("hello" + (v4 + "\n\t")), F1: v81}
        if true {
            switch (fi82 >> ((uint(v2)) & 63)) {
            case 32767:
                var v84 float32 = 0.0
                _ = v84
            default:
                fmt.Println(v4)
            }
            v4 = v4
            for fi85 := 0; fi85 < 7; fi85++ {
                if false {
                    var v86 uint32 = v2
                    switch "hello" {
                    case "a":
                        v87 := []int{(fi85 - -1), 42, fi85, fi82}
                        _ = v87
                    default:
                        fmt.Println(v4, v2)
                    }
                    _ = v86
                }
                _ = fi85
            }
        }
        var v88 uint16 = 42
        _ = fi82
        _ = v83
        _ = v88
    }
    v89 := v5[("\n\t" + (v4 + "a"))]
    for fi90 := 0; fi90 < 18; fi90++ {
        fmt.Println(v2, v78)
        for fi91 := 0; fi91 < 1; fi91++ {
            _ = len(v5)
            v77 -= 2
            _ = fi91
        }
        _ = fi90
    }
    v92 := GenS9{F0: 138}
    var v93 float64 = -2.25
    switch 2147483647 {
    case 630:
        if v89 {
            switch v4 {
            case "":
                for fi94 := 0; fi94 < 19; fi94++ {
                    for fi95 := 0; fi95 < 5; fi95++ {
                        v81 = -114
                        v78 = v78
                        v96 := []int{(fi95 % ((992) | 1))}
                        _ = fi95
                        _ = v96
                    }
                    _ = fi94
                }
            }
            v97 := []int{-1}
            if v89 {
                for fi98 := 0; fi98 < 10; fi98++ {
                    var v99 uint16 = 42
                    var v100 float32 = 3.14159
                    v93 -= float64(1e10)
                    _ = fi98
                    _ = v100
                    _ = v99
                }
                var v101 string = (("" + v4) + (v4 + v4))
                for fi102 := 0; fi102 < 1; fi102++ {
                    v103 := GenS10{F0: 128, F1: 0}
                    switch "foo bar" {
                    case "a":
                        if (fi102 != fi102) {
                            var v104 string = v4
                            var v105 uint8 = 2
                            v5["a"] = (!(14 >= 127))
                            _ = v104
                            _ = v105
                        }
                    case "foo bar":
                        v106 := GenS4{F0: 0}
                        _ = v106
                    case "hello":
                        v107 := GenS11{F0: (float64(float32(v93)) < 0.0), F1: (-(v81))}
                        _ = v107
                    default:
                        var v108 uint = v78
                        _ = v108
                    }
                    _ = fi102
                    _ = v103
                }
                _ = v101
            } else {
                if v89 {
                    var v109 int64 = v3
                    if v89 {
                        if (v89 || v89) {
                            switch -136 {
                            case 42:
                                v110 := GenS1{F0: (v3 >> ((v78) & 63)), F1: v89, F2: (v2 / (((v77 ^ (v77 ^ v77))) | 1))}
                                _ = v110
                            default:
                                v111 := GenS12{F0: 42}
                                _ = v111
                            }
                            v112 := GenS13{F0: -12}
                            if true {
                                if (v89 || true) {
                                    v2 = (v77 & 65535)
                                    if (!((!((!true) && ((v93 / v93) == (((v93 - ((float64(1.5) / 3.14159) / (float64(-2.25) - float64((float32((float64(3.14159) / 1.0)) * ((0.0 - (0.1 - float32(v93))) - float32(v93))))))) / v93) * ((v93 + v93) + -2.25))))) || ((^(v1)) >= 292))) {
                                        v97 = append(v97, 9223372036854775807)
                                        v93 = (v93 / v93)
                                        var v113 bool = ((v1 >> ((197) & 63)) < (v1 << ((v78) & 63)))
                                        _ = v113
                                    } else {
                                        switch "hello" {
                                        case "\n\t":
                                            var v114 uint8 = 0
                                            _ = v114
                                        case "foo bar":
                                            for fi115 := 0; fi115 < 20; fi115++ {
                                                if (v1 != (v1 & (-(v1)))) {
                                                    delete(v5, "hello")
                                                    fmt.Println(v78, v93, fi115)
                                                }
                                                v1 = v1
                                                var v116 float32 = 0.1
                                                _ = fi115
                                                _ = v116
                                            }
                                        case "hello":
                                            v77 = v2
                                        }
                                        fmt.Println(v2)
                                        v117 := GenS12{F0: v1}
                                        _ = v117
                                    }
                                    v3 += -1
                                }
                            } else {
                                v118 := v5[v4]
                                _ = v118
                            }
                            _ = v112
                        } else {
                            for fi119 := 0; fi119 < 18; fi119++ {
                                var v120 uint8 = 203
                                fmt.Println(v3)
                                var v121 uint32 = v77
                                _ = fi119
                                _ = v120
                                _ = v121
                            }
                            for fi122 := 0; fi122 < 19; fi122++ {
                                for fi123 := 0; fi123 < 12; fi123++ {
                                    v78 += v78
                                    var v124 uint64 = 776
                                    _ = fi123
                                    _ = v124
                                }
                                _ = fi122
                            }
                            v81 += (v81 * v81)
                        }
                        fmt.Println(v77)
                    }
                    v125 := GenS14{F0: 42, F1: 2147483647}
                    _ = v109
                    _ = v125
                }
                for fi126 := 0; fi126 < 18; fi126++ {
                    for fi127 := 0; fi127 < 12; fi127++ {
                        var v128 int = fi127
                        _ = fi127
                        _ = v128
                    }
                    _ = fi126
                }
                var v129 int16 = v1
                _ = v129
            }
            _ = v97
        } else {
            v130 := v5[("\n\t" + ("" + v4))]
            for fi131 := 0; fi131 < 6; fi131++ {
                if (v2 <= (v2 << (((v78 | v78)) & 63))) {
                    if true {
                        if (v78 >= (v78 - v78)) {
                            v132 := GenS8{F0: v4}
                            _ = v132
                        } else {
                            for fi133 := 0; fi133 < 11; fi133++ {
                                if (-2147483648 < -1) {
                                    if v89 {
                                        v4 = v4
                                    } else {
                                        switch "hello" {
                                        case "\n\t":
                                            switch v4 {
                                            case "":
                                                v130 = v89
                                            case "\n\t":
                                                v4 = "hello"
                                            case "hello":
                                                fmt.Println(v81, v81, v3)
                                            }
                                        case "":
                                            fmt.Println(v130, v78, v78)
                                        }
                                    }
                                    var v134 int = fi133
                                    _ = v134
                                }
                                v135 := GenS14{F0: (v3 | (v3 >> ((v78) & 63))), F1: 0}
                                v130 = v89
                                _ = fi133
                                _ = v135
                            }
                            v4 = v4
                        }
                        v81 = -1
                    }
                    switch (fi131 + (fi131 * 0)) {
                    case 0:
                        var v136 int32 = 1
                        _ = v136
                    }
                    switch ("" + (v4 + v4)) {
                    case "hello":
                        var v137 float64 = (v93 * v93)
                        _ = v137
                    }
                } else {
                    for fi138 := 0; fi138 < 12; fi138++ {
                        v139 := GenS10{F0: 1, F1: 1}
                        switch v4 {
                        case "hello":
                            fmt.Println(v81, fi138)
                        default:
                            fmt.Println(v77)
                        }
                        _ = fi138
                        _ = v139
                    }
                    v140 := []int{}
                    var v141 int16 = v1
                    _ = v140
                    _ = v141
                }
                if (2147483647 != -48) {
                    var v142 int64 = (v3 >> ((v78) & 63))
                    v143 := GenS15{F0: 1}
                    v144 := GenS12{F0: (v1 << ((v78) & 63))}
                    _ = v142
                    _ = v143
                    _ = v144
                }
                _ = fi131
            }
            _ = v130
        }
    case -924:
        v145 := GenS16{F0: -210, F1: 164}
        _ = v145
    case 2147483647:
        var v146 float32 = 0.0
        _ = v146
    default:
        var v147 float64 = 0.0
        _ = v147
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v81:", v81)
    fmt.Println("v89:", v89)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v76
    _ = v77
    _ = v78
    _ = v79
    _ = v81
    _ = v89
    _ = v92
    _ = v93
}
