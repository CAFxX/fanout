package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 int8 }
type GenS6 struct { F0 int32; F1 int64; F2 uint64 }
type GenS7 struct { F0 int8; F1 int8 }
type GenS8 struct { F0 uint; F1 bool; F2 int32 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int; F1 uint32 }
type GenS12 struct { F0 int8; F1 int32; F2 uint16 }
type GenS13 struct { F0 int32; F1 int64 }
type GenS14 struct { F0 uint8; F1 uint64 }
type GenS15 struct { F0 int8; F1 uint; F2 bool }
func genFn0(p0 uint) bool {
    v5 := []int{-1, -9223372036854775808, 32767}
    if len(v5) > 0 {
        v5[0] = 128
    }
    _ = p0
    _ = v5
    return false
}
func genFn1() uint8 {
    var v6 float64 = float64((0.0 / 0.1))
    if false {
        v7 := GenS1{F0: 0}
        v6 = float64((0.0 * -2.25))
        if (42 <= 358) {
            for fi8 := 0; fi8 < 6; fi8++ {
                v6 -= ((0.0 + (v6 * 3.14159)) - (1.5 / (3.14159 / v6)))
                var v9 int32 = -692
                v10 := GenS1{F0: 9223372036854775807}
                _ = fi8
                _ = v10
                _ = v9
            }
        } else {
            v11 := []int{}
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 uint = 858
                if (!true) {
                    for fi14 := 0; fi14 < 7; fi14++ {
                        var v15 uint8 = 3
                        if true {
                            switch "a" {
                            case "a":
                                v6 = v6
                            case "foo bar":
                                var v16 uint = v13
                                _ = v16
                            case "\n\t":
                                if (!true) {
                                    var v17 bool = (true && (1.5 < 1e10))
                                    _ = v17
                                }
                            default:
                                var v18 int = (fi12 + (fi12 * (fi12 / ((-1) | 1))))
                                _ = v18
                            }
                            if false {
                                var v19 uint = 255
                                if false {
                                    var v20 string = ("a" + (("hello" + ("\n\t" + "a")) + "foo bar"))
                                    _ = v20
                                }
                                if (!false) {
                                    var v21 uint8 = v15
                                    v6 -= v6
                                    v22 := GenS2{F0: 868, F1: 1, F2: 42}
                                    _ = v21
                                    _ = v22
                                }
                                _ = v19
                            } else {
                                if ((747 > 1) && false) {
                                    v13 += v13
                                }
                                fmt.Println(fi14)
                                v23 := GenS3{F0: "\n\t"}
                                _ = v23
                            }
                            var v24 int8 = 42
                            if (-1 != 42) {
                                if false {
                                    if true {
                                        v11 = append(v11, (-(fi12)))
                                        switch "hello" {
                                        case "":
                                            if true {
                                                v25 := GenS3{F0: ("" + "\n\t")}
                                                _ = v25
                                            }
                                        case "\n\t":
                                            if true {
                                                v26 := GenS4{F0: 255}
                                                _ = v26
                                            }
                                            var v27 int = (-(fi14))
                                            v27 = (v27 & fi12)
                                            _ = v27
                                        case "foo bar":
                                            var v28 int64 = 0
                                            _ = v28
                                        default:
                                            v13 += 42
                                        }
                                        fmt.Println(fi14, v24)
                                    } else {
                                        v6 += v6
                                        if ((fi12 & (fi12 - 1)) != fi12) {
                                            v29 := GenS5{F0: "hello", F1: int8(v15)}
                                            if true {
                                                var v30 float32 = (1.5 / (1.5 - -2.25))
                                                if true {
                                                    var v31 uint8 = v15
                                                    var v32 uint64 = 564
                                                    v24 += 42
                                                    _ = v31
                                                    _ = v32
                                                }
                                                if true {
                                                    v33 := GenS2{F0: 4294967295, F1: 0, F2: 1}
                                                    var v34 uint32 = 72
                                                    v34 = 382
                                                    _ = v33
                                                    _ = v34
                                                } else {
                                                    if (!true) {
                                                        var v35 uint64 = 1
                                                        _ = v35
                                                    }
                                                    fmt.Println(v6)
                                                    var v36 int = fi14
                                                    v36 = fi12
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        fmt.Println(fi12)
                                                    case "a":
                                                        if len(v11) >= 2 {
                                                            v11 = v11[1:len(v11)]
                                                        }
                                                    default:
                                                        v36 -= -9223372036854775808
                                                    }
                                                    v11 = append(v11, -2147483648)
                                                    if (v30 != (v30 * 0.1)) {
                                                        var v37 uint = v13
                                                        v38 := GenS5{F0: (((((("\n\t" + "hello") + "hello") + "\n\t") + (("a" + ("" + ("foo bar" + (("hello" + "hello") + ("hello" + ""))))) + (("foo bar" + "hello") + "\n\t"))) + (("\n\t" + "hello") + (("foo bar" + "") + ""))) + ("" + "hello")), F1: v24}
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                    _ = v36
                                                }
                                                _ = v30
                                            }
                                            _ = v29
                                        } else {
                                            var v39 uint = (^(v13))
                                            v15 -= (v15 * v15)
                                            v40 := GenS4{F0: 255}
                                            _ = v39
                                            _ = v40
                                        }
                                        v41 := GenS2{F0: 264, F1: v24, F2: 170}
                                        v42 := GenS6{F0: 1, F1: -468, F2: 533}
                                        v24 = v24
                                        _ = v41
                                        _ = v42
                                    }
                                }
                            }
                            if true {
                                var v43 int64 = -1
                                if (false && (true && (false && true))) {
                                    fmt.Println(v15, v6)
                                } else {
                                    var v44 int8 = v24
                                    _ = v44
                                }
                                _ = v43
                            } else {
                                fmt.Println(v13, fi14)
                                fmt.Println(fi14, fi12)
                                v15 = 42
                            }
                            _ = v24
                        }
                        _ = fi14
                        _ = v15
                    }
                    fmt.Println(v6, v6, fi12)
                } else {
                    for fi45 := 0; fi45 < 5; fi45++ {
                        fmt.Println(fi45)
                        switch fi12 {
                        case -9223372036854775808:
                            switch "a" {
                            case "\n\t":
                                var v46 int16 = int16(fi12)
                                _ = v46
                            case "a":
                                fmt.Println(fi12, v13, fi45)
                            }
                        case 0:
                            v13 = (v13 << ((v13) & 63))
                        case -288:
                            var v47 uint32 = 42
                            _ = v47
                        default:
                            v48 := GenS5{F0: "foo bar", F1: 127}
                            _ = v48
                        }
                        switch fi45 {
                        case 502:
                            var v49 bool = false
                            _ = v49
                        case 128:
                            v13 += v13
                        case -542:
                            v6 = (v6 / ((v6 - (v6 * 0.1)) / (float64(3.14159) - ((((v6 - -2.25) - float64((float32((v6 + (v6 * ((v6 + v6) / (1e10 / v6))))) + (-2.25 - 1.5)))) / ((v6 - ((v6 * 1.5) * (v6 * 1e10))) + v6)) / (1e10 + v6)))))
                        }
                        _ = fi45
                    }
                    v50 := GenS1{F0: -353}
                    _ = v50
                }
                _ = fi12
                _ = v13
            }
            switch ("a" + ("\n\t" + "a")) {
            case "\n\t":
                v51 := GenS6{F0: 0, F1: 1, F2: 2}
                _ = v51
            default:
                fmt.Println(v6)
            }
            _ = v11
        }
        _ = v7
    } else {
        fmt.Println(v6)
        v6 -= v6
    }
    _ = v6
    return 61
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    var v3 int64 = -33
    var v4 uint8 = 1
    v4 += (v4 - v4)
    fmt.Println(v3)
    var v52 int16 = 32767
    switch ("\n\t" + "hello") {
    case "foo bar":
        var v53 int16 = 0
        _ = v53
    case "a":
        var v54 uint = 0
        _ = v54
    case "\n\t":
        var v55 bool = true
        _ = v55
    }
    fmt.Println(v52)
    switch "" {
    case "":
        if ("\n\t" != "a") {
            var v56 float64 = -2.25
            var v57 string = ("foo bar" + "\n\t")
            _ = v56
            _ = v57
        }
    case "hello":
        fmt.Println(v4, v3, v4)
    case "\n\t":
        if true {
            for fi58 := 0; fi58 < 9; fi58++ {
                if (false && false) {
                    var v59 float32 = -2.25
                    switch (fi58 / ((fi58) | 1)) {
                    case 853:
                        fmt.Println(v59, v52, v4)
                    case 18:
                        v59 = ((float32((1e10 / (0.1 - 0.0))) - v59) + 3.14159)
                    case 2147483647:
                        var v60 bool = false
                        _ = v60
                    default:
                        if true {
                            if (true || (fi58 > fi58)) {
                                for fi61 := 0; fi61 < 13; fi61++ {
                                    v62 := GenS7{F0: 110, F1: int8(v3)}
                                    _ = fi61
                                    _ = v62
                                }
                                v63 := GenS3{F0: (((("hello" + "hello") + "foo bar") + "hello") + "hello")}
                                v64 := GenS3{F0: ("" + "")}
                                _ = v63
                                _ = v64
                            }
                            v3 = (v3 | (v3 / (((v3 | (^(v3)))) | 1)))
                            switch (fi58 * (fi58 + 2147483647)) {
                            case 769:
                                fmt.Println(v52)
                            case 127:
                                if ("hello" != "") {
                                    v65 := GenS2{F0: 387, F1: 0, F2: 4294967295}
                                    switch 9223372036854775807 {
                                    case 32768:
                                        v66 := GenS8{F0: 42, F1: true, F2: 1}
                                        _ = v66
                                    case 32767:
                                        for fi67 := 0; fi67 < 1; fi67++ {
                                            if false {
                                                v4 = (v4 + v4)
                                                v3 = v3
                                                _ = len(v2)
                                            } else {
                                                v52 = 42
                                                var v68 string = ("a" + "a")
                                                _ = v68
                                            }
                                            switch (-(fi67)) {
                                            case 916:
                                                v69 := []int{}
                                                _ = v69
                                            }
                                            _ = fi67
                                        }
                                    case 811:
                                        var v70 uint64 = 42
                                        _ = v70
                                    default:
                                        switch fi58 {
                                        case 0:
                                            if false {
                                                v52 -= v52
                                                if false {
                                                    v59 = v59
                                                }
                                            } else {
                                                var v71 uint32 = 1
                                                var v72 int32 = int32(v71)
                                                _ = v71
                                                _ = v72
                                            }
                                        default:
                                            if false {
                                                v73 := v1[127]
                                                v2["\n\t"] = fi58
                                                _ = v73
                                            }
                                        }
                                    }
                                    v59 = (v59 * 3.14159)
                                    _ = v65
                                }
                            default:
                                if ((0.0 >= 3.14159) || (!(!(255 == 520)))) {
                                    if false {
                                        if ((true && true) && false) {
                                            for fi74 := 0; fi74 < 6; fi74++ {
                                                v4 += (v4 ^ v4)
                                                _ = fi74
                                            }
                                            fmt.Println(v52, v52)
                                            v3 = v3
                                        } else {
                                            switch "a" {
                                            case "hello":
                                                v59 += 1e10
                                            case "":
                                                v59 += (v59 - (3.14159 * (v59 + 0.1)))
                                            case "\n\t":
                                                v52 += (^(v52))
                                            }
                                            var v75 int8 = 0
                                            for fi76 := 0; fi76 < 16; fi76++ {
                                                v77 := []int{389}
                                                v59 = (float32(3.14159) - float32(0.1))
                                                fmt.Println(v52, v4)
                                                _ = fi76
                                                _ = v77
                                            }
                                            _ = v75
                                        }
                                        for fi78 := 0; fi78 < 1; fi78++ {
                                            var v79 int8 = 0
                                            switch fi58 {
                                            case 1:
                                                var v80 int8 = v79
                                                _ = v80
                                            case -497:
                                                var v81 int16 = v52
                                                _ = v81
                                            case 128:
                                                fmt.Println(v59, v79, fi78)
                                            }
                                            _ = fi78
                                            _ = v79
                                        }
                                    } else {
                                        var v82 float32 = v59
                                        _ = v82
                                    }
                                    for fi83 := 0; fi83 < 9; fi83++ {
                                        var v84 uint = 416
                                        _ = fi83
                                        _ = v84
                                    }
                                }
                            }
                        }
                    }
                    v85 := GenS4{F0: 490}
                    _ = v59
                    _ = v85
                }
                switch (fi58 / ((-658) | 1)) {
                case 2147483647:
                    v86 := []int{fi58}
                    _ = v86
                case 128:
                    v2["hello"] = (fi58 - (-(fi58)))
                case -75:
                    var v87 int8 = -8
                    _ = v87
                default:
                    fmt.Println(v4)
                }
                var v88 bool = (2147483647 > 42)
                _ = fi58
                _ = v88
            }
            for fi89 := 0; fi89 < 11; fi89++ {
                fmt.Println(v52, v52)
                switch (fi89 >> ((4294967295) & 63)) {
                case 1:
                    v3 -= -9223372036854775808
                case 965:
                    var v90 int8 = -128
                    _ = v90
                }
                _ = fi89
            }
            delete(v1, 127)
        } else {
            v3 -= (v3 ^ (v3 | 9223372036854775807))
            var v91 uint8 = v4
            var v92 uint = 0
            _ = v91
            _ = v92
        }
    default:
        v52 = (v52 << ((65535) & 63))
    }
    switch 1 {
    case -55:
        v3 = v3
    case 397:
        for fi93 := 0; fi93 < 9; fi93++ {
            v4 -= genFn1()
            fmt.Println(v4)
            switch "" {
            case "foo bar":
                var v94 uint8 = v4
                _ = v94
            case "a":
                v95 := GenS2{F0: 1, F1: 1, F2: 2}
                _ = v95
            case "hello":
                var v96 int32 = 2147483647
                _ = v96
            default:
                var v97 bool = true
                _ = v97
            }
            _ = fi93
        }
    case 9223372036854775807:
        if (false && false) {
            if (true && true) {
                if true {
                    v98 := GenS9{F0: (^(v4))}
                    v3 = v3
                    if (((true && (!(32768 > 761))) || (1 < 42)) && true) {
                        var v99 uint = 4294967295
                        _ = v99
                    } else {
                        for fi100 := 0; fi100 < 5; fi100++ {
                            v101 := GenS10{F0: 1000}
                            _ = fi100
                            _ = v101
                        }
                    }
                    _ = v98
                } else {
                    v102 := GenS3{F0: "\n\t"}
                    _ = v102
                }
                v3 -= int64(v52)
            } else {
                var v103 float32 = 1.5
                v104 := GenS11{F0: 32768, F1: 2147483647}
                _ = v103
                _ = v104
            }
            for fi105 := 0; fi105 < 2; fi105++ {
                var v106 float64 = 1e10
                for fi107 := 0; fi107 < 12; fi107++ {
                    v106 -= v106
                    var v108 int = (^(fi105))
                    if false {
                        if true {
                            var v109 uint = 42
                            var v110 int8 = 42
                            if (!false) {
                                v4 += genFn1()
                            } else {
                                if (4294967295 <= 2147483648) {
                                    v109 = v109
                                } else {
                                    var v111 uint64 = 217
                                    _ = v111
                                }
                            }
                            fmt.Println(v109)
                            v108 = (fi105 + (fi107 ^ -1))
                            _ = v109
                            _ = v110
                        } else {
                            var v112 int16 = (v52 << ((1) & 63))
                            _ = v112
                        }
                    } else {
                        v3 = (v3 * v3)
                    }
                    _ = fi107
                    _ = v108
                }
                _ = fi105
                _ = v106
            }
        } else {
            fmt.Println(v52)
            for fi113 := 0; fi113 < 13; fi113++ {
                if (!false) {
                    switch int(v4) {
                    case 495:
                        if ((!true) || (-18 >= -1)) {
                            v3 -= -1
                        } else {
                            var v114 int16 = (^(v52))
                            v115 := []int{fi113}
                            _ = v114
                            _ = v115
                        }
                    case 2147483647:
                        if (int8(fi113) != -105) {
                            var v116 int32 = -579
                            v4 = 128
                            _ = v116
                        }
                    case 32768:
                        fmt.Println(fi113, v3, v3)
                    }
                    for fi117 := 0; fi117 < 6; fi117++ {
                        var v118 int16 = (-(v52))
                        v119 := GenS12{F0: -47, F1: -2147483648, F2: 255}
                        _ = fi117
                        _ = v118
                        _ = v119
                    }
                    v120 := GenS2{F0: 111, F1: 1, F2: 23}
                    _ = v120
                }
                _ = fi113
            }
            v121 := []int{}
            _ = v121
        }
    default:
        v122 := GenS5{F0: "a", F1: -128}
        _ = v122
    }
    if (!(553 > 2147483647)) {
        for fi123 := 0; fi123 < 13; fi123++ {
            v124 := []int{9223372036854775807, fi123, (fi123 | 205)}
            if len(v124) > 0 {
                v124[0] = (^(fi123))
            }
            v52 -= v52
            _ = fi123
            _ = v124
        }
        v3 -= v3
        v4 += (v4 | 127)
    } else {
        for fi125 := 0; fi125 < 7; fi125++ {
            for fi126 := 0; fi126 < 3; fi126++ {
                var v127 float32 = -2.25
                var v128 int = fi126
                _ = fi126
                _ = v127
                _ = v128
            }
            v129 := []int{127, (fi125 % ((46) | 1))}
            if len(v129) > 0 {
                v129[0] = (fi125 << ((1) & 63))
            }
            _ = fi125
            _ = v129
        }
        v52 = (v52 >> ((65535) & 63))
    }
    if true {
        v52 = 266
        v3 = v3
        if false {
            v130 := []int{}
            _ = v130
        } else {
            switch 846 {
            case -596:
                switch -119 {
                case 0:
                    v131 := GenS13{F0: -1, F1: (^(v3))}
                    _ = v131
                case 32767:
                    var v132 bool = (!((!((0 <= 270) || ((v52 >> ((1) & 63)) <= v52))) && ((v4 % (((v4 * (v4 >> ((446) & 63)))) | 1)) != (v4 | v4))))
                    _ = v132
                default:
                    v133 := GenS14{F0: (^(v4)), F1: 166}
                    _ = v133
                }
            case 1:
                switch ("hello" + "\n\t") {
                case "a":
                    v134 := []int{398, 4294967295}
                    _ = v134
                case "\n\t":
                    v135 := GenS7{F0: int8(v52), F1: 39}
                    _ = v135
                case "hello":
                    switch 4294967295 {
                    case 32768:
                        if true {
                            fmt.Println(v4, v3, v52)
                            var v136 int32 = 1
                            v137 := GenS9{F0: 193}
                            _ = v136
                            _ = v137
                        } else {
                            if ("foo bar" <= "a") {
                                v138 := GenS15{F0: 1, F1: 42, F2: true}
                                var v139 int8 = -1
                                _ = v138
                                _ = v139
                            } else {
                                fmt.Println(v4, v4)
                                v52 = -1
                                switch 795 {
                                case -9223372036854775808:
                                    switch ("foo bar" + "foo bar") {
                                    case "a":
                                        v3 -= 42
                                    }
                                }
                            }
                            v140 := GenS4{F0: 1}
                            v3 += (v3 * v3)
                            _ = v140
                        }
                    case -982:
                        v52 = 32767
                    case 217:
                        var v141 string = (((("hello" + ("\n\t" + "foo bar")) + "hello") + ("\n\t" + ("foo bar" + "\n\t"))) + ((("\n\t" + ("" + "hello")) + "\n\t") + ""))
                        _ = v141
                    }
                default:
                    switch ("\n\t" + "foo bar") {
                    case "foo bar":
                        v142 := []int{32767, 32767, 2147483647}
                        _ = v142
                    default:
                        var v143 int64 = (v3 | 0)
                        _ = v143
                    }
                }
            default:
                v144 := GenS14{F0: 26, F1: 18446744073709551615}
                _ = v144
            }
            var v145 string = "foo bar"
            _ = v145
        }
    }
    v146 := []int{2147483647, 878, 4294967295}
    fmt.Println(v3, v52, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v146:", len(v146), v146)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v146
    _ = v2
    _ = v3
    _ = v4
    _ = v52
}
