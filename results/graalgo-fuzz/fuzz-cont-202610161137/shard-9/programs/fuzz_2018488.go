package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 string; F1 bool; F2 int8 }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 int32; F1 int64; F2 bool }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int; F1 int64 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint }
func genFn0(p0 uint) int {
    for fi5 := 0; fi5 < 9; fi5++ {
        fmt.Println(fi5, fi5)
        _ = fi5
    }
    v6 := []int{-755, -9223372036854775808, 9223372036854775807}
    var v7 float64 = 0.1
    p0 += 4294967295
    fmt.Println(p0)
    _ = p0
    _ = v6
    _ = v7
    return -789
}
func genFn1(p0 bool, p1 uint8) int {
    for fi8 := 0; fi8 < 15; fi8++ {
        if true {
            if false {
                v9 := GenS1{F0: 127, F1: 2}
                var v10 uint8 = p1
                var v11 uint16 = 0
                _ = v10
                _ = v11
                _ = v9
            } else {
                for fi12 := 0; fi12 < 17; fi12++ {
                    var v13 uint64 = 4294967295
                    v14 := []int{1, 9223372036854775807}
                    v13 += v13
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                for fi15 := 0; fi15 < 6; fi15++ {
                    v16 := []int{}
                    v17 := GenS1{F0: -1, F1: 42}
                    v18 := GenS2{F0: "hello", F1: false, F2: 100}
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
                switch "\n\t" {
                case "foo bar":
                    if ((!(1 > 2147483647)) && p0) {
                        v19 := GenS3{F0: 2147483648, F1: 722, F2: 42}
                        fmt.Println(p1, p1, fi8)
                        for fi20 := 0; fi20 < 17; fi20++ {
                            p0 = p0
                            _ = fi20
                        }
                        _ = v19
                    }
                default:
                    var v21 uint32 = 2147483647
                    _ = v21
                }
            }
            p0 = false
        }
        v22 := GenS2{F0: "foo bar", F1: (966 < 885), F2: 127}
        if true {
            if (-1 == 1) {
                if ((0.0 * 3.14159) != -2.25) {
                    v23 := []int{fi8, (fi8 / (((fi8 * (-(fi8)))) | 1))}
                    _ = v23
                }
                v24 := GenS4{F0: 127, F1: -1}
                _ = v24
            }
            var v25 int32 = -1
            p1 += p1
            _ = v25
        } else {
            if p0 {
                var v26 uint64 = 18446744073709551615
                fmt.Println(v26, p1)
                _ = v26
            }
            var v27 int8 = -1
            var v28 uint32 = 2147483647
            _ = v27
            _ = v28
        }
        _ = fi8
        _ = v22
    }
    v29 := GenS3{F0: 359, F1: 4294967295, F2: 57}
    switch 42 {
    case 99:
        var v30 uint64 = 60
        _ = v30
    case 127:
        v31 := []int{0, genFn0(4294967295), 0, -2147483648}
        _ = v31
    case -274:
        var v32 uint32 = 4294967295
        _ = v32
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        v34 := GenS4{F0: p1, F1: 0}
        if p0 {
            p0 = true
            v35 := GenS2{F0: ("" + "foo bar"), F1: ("\n\t" < "\n\t"), F2: 100}
            _ = v35
        } else {
            p0 = (p0 || ((406 > -1) && (true && (!(100 < 127)))))
            if (42 >= 1) {
                v36 := GenS5{F0: 2147483647, F1: -1, F2: false}
                var v37 int64 = -1
                _ = v36
                _ = v37
            }
        }
        p0 = p0
        _ = fi33
        _ = v34
    }
    switch 32768 {
    case 1:
        if (1 != 54) {
            var v38 uint32 = 4294967295
            _ = v38
        } else {
            if true {
                for fi39 := 0; fi39 < 7; fi39++ {
                    switch "\n\t" {
                    case "\n\t":
                        for fi40 := 0; fi40 < 11; fi40++ {
                            var v41 uint8 = p1
                            v42 := GenS6{F0: "\n\t", F1: (2147483648 < 0)}
                            p0 = (fi40 == (fi39 % ((fi39) | 1)))
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                    default:
                        p0 = ((p1 + p1) >= 127)
                    }
                    _ = fi39
                }
            } else {
                v43 := []int{-9223372036854775808}
                var v44 float32 = 3.14159
                _ = v43
                _ = v44
            }
        }
    }
    _ = p0
    _ = p1
    _ = v29
    return -9223372036854775808
}

func main() {
    v1 := []string{("\n\t" + "\n\t"), ("a" + (((((("foo bar" + "hello") + "foo bar") + "hello") + ("" + "a")) + "a") + "hello")), "a"}
    var v2 int32 = -1
    var v3 int64 = 9223372036854775807
    v4 := []string{}
    switch -2147483648 {
    case -129:
        v3 += v3
    case 568:
        var v45 uint64 = 306
        _ = v45
    case -25:
        if true {
            fmt.Println(v3)
        } else {
            if len(v1) > 0 {
                v1[0] = ""
            }
            v2 = (v2 - -2147483648)
            fmt.Println(v2, v3)
        }
    }
    if false {
        v46 := GenS4{F0: 2, F1: 42}
        _ = v46
    } else {
        switch 2147483647 {
        case -212:
            v47 := GenS7{F0: "\n\t"}
            _ = v47
        case 127:
            _ = len(v4)
        default:
            for fi48 := 0; fi48 < 5; fi48++ {
                for fi49 := 0; fi49 < 4; fi49++ {
                    if false {
                        switch (fi48 | (fi48 << ((255) & 63))) {
                        case 2147483647:
                            switch fi48 {
                            case 2147483647:
                                var v50 bool = true
                                _ = v50
                            }
                            if false {
                                _ = len(v4)
                                v4 = append(v4, "a")
                                v51 := GenS1{F0: 1, F1: 905}
                                _ = v51
                            } else {
                                v3 = (-(v3))
                                if (genFn0(1) != genFn1(false, 0)) {
                                    var v52 uint = 65535
                                    var v53 float32 = float32((0.1 / (1.5 + 1e10)))
                                    v53 = -2.25
                                    _ = v52
                                    _ = v53
                                } else {
                                    if true {
                                        v3 = 9223372036854775807
                                    } else {
                                        v2 = v2
                                        var v54 string = "hello"
                                        _ = v54
                                    }
                                    var v55 int16 = 1000
                                    _ = v55
                                }
                                var v56 bool = true
                                _ = v56
                            }
                            var v57 uint = 65535
                            _ = v57
                        case 4294967295:
                            var v58 bool = (!true)
                            _ = v58
                        case -42:
                            _ = len(v4)
                        default:
                            if (0.1 > -2.25) {
                                var v59 int16 = 42
                                var v60 int64 = -548
                                _ = v59
                                _ = v60
                            }
                        }
                    } else {
                        v4 = append(v4, "hello")
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v61 float64 = 3.14159
                        _ = v61
                    }
                    switch "hello" {
                    case "a":
                        v2 += (v2 ^ (v2 ^ v2))
                    case "hello":
                        var v62 string = "\n\t"
                        _ = v62
                    default:
                        var v63 float64 = ((3.14159 * 1e10) * float64(0.0))
                        _ = v63
                    }
                    v2 -= (^(v2))
                    _ = fi49
                }
                v3 += v3
                switch "hello" {
                case "foo bar":
                    if len(v1) > 0 {
                        v1[0] = ("foo bar" + "hello")
                    }
                }
                _ = fi48
            }
        }
    }
    if (!(0 >= int16(v2))) {
        v3 = (v3 + 9223372036854775807)
        switch "" {
        case "":
            for fi64 := 0; fi64 < 5; fi64++ {
                var v65 int8 = int8(v3)
                var v66 uint = 2
                _ = fi64
                _ = v65
                _ = v66
            }
        default:
            fmt.Println(v3)
        }
    }
    v67 := GenS5{F0: (-(v2)), F1: 185, F2: false}
    if true {
        for fi68 := 0; fi68 < 15; fi68++ {
            fmt.Println(v3, v2)
            if (42 <= 127) {
                v3 = v3
                v3 += -1
            }
            v2 = 42
            _ = fi68
        }
    }
    switch -975 {
    case 9223372036854775807:
        v2 -= v2
    case 0:
        if true {
            var v69 uint8 = 56
            _ = v69
        } else {
            if (0 == 1) {
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
                var v70 uint32 = 516
                v71 := GenS6{F0: "a", F1: (1 <= -1)}
                _ = v70
                _ = v71
            } else {
                if (!(true && false)) {
                    var v72 float64 = 1e10
                    v72 = (((float64((-2.25 * (0.1 + 1e10))) + (3.14159 + (v72 / -2.25))) * ((0.0 + 3.14159) / v72)) - (1e10 / (float64(-2.25) + (3.14159 * (v72 - (v72 + 0.1))))))
                    for fi73 := 0; fi73 < 3; fi73++ {
                        var v74 uint64 = 550
                        var v75 uint16 = 2
                        switch "hello" {
                        case "foo bar":
                            switch "" {
                            case "foo bar":
                                if (v2 != v2) {
                                    switch (fi73 & fi73) {
                                    case 0:
                                        if false {
                                            for fi76 := 0; fi76 < 4; fi76++ {
                                                if false {
                                                    v72 -= float64((0.1 * (3.14159 / (((0.1 - 0.1) - -2.25) * (3.14159 - (1.5 * 1.5))))))
                                                    v75 += v75
                                                    fmt.Println(v75, fi76, fi76)
                                                }
                                                v77 := GenS8{F0: (^(fi73)), F1: 333}
                                                v74 -= (^(v74))
                                                fmt.Println(fi76, v75)
                                                _ = fi76
                                                _ = v77
                                            }
                                            for fi78 := 0; fi78 < 19; fi78++ {
                                                v2 -= (v2 >> ((2) & 63))
                                                v72 = v72
                                                switch (("a" + "") + "hello") {
                                                case "hello":
                                                    if true {
                                                        switch (("hello" + "") + "foo bar") {
                                                        case "hello":
                                                            v79 := GenS3{F0: 576, F1: 42, F2: 2}
                                                            _ = v79
                                                        }
                                                    }
                                                case "foo bar":
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                                _ = fi78
                                            }
                                        }
                                    }
                                    for fi80 := 0; fi80 < 17; fi80++ {
                                        v2 += v2
                                        _ = fi80
                                    }
                                }
                            case "":
                                var v81 uint8 = 127
                                _ = v81
                            case "\n\t":
                                for fi82 := 0; fi82 < 13; fi82++ {
                                    var v83 string = "\n\t"
                                    v84 := GenS8{F0: int(v74), F1: v3}
                                    v3 = (v3 + v3)
                                    _ = fi82
                                    _ = v83
                                    _ = v84
                                }
                            }
                        }
                        _ = fi73
                        _ = v74
                        _ = v75
                    }
                    _ = v72
                } else {
                    switch 687 {
                    case -429:
                        switch ("\n\t" + "a") {
                        case "a":
                            v85 := GenS7{F0: "\n\t"}
                            _ = v85
                        case "":
                            _ = len(v1)
                        default:
                            v2 = v2
                        }
                    }
                    for fi86 := 0; fi86 < 6; fi86++ {
                        v87 := GenS7{F0: ("" + "hello")}
                        v2 = 185
                        if true {
                            v3 -= (v3 % ((-9223372036854775808) | 1))
                            if (false || true) {
                                v3 = (v3 | 42)
                            } else {
                                v3 = -800
                                switch "\n\t" {
                                case "hello":
                                    v3 = (v3 - (v3 / ((v3) | 1)))
                                case "a":
                                    var v88 uint32 = 65535
                                    _ = v88
                                default:
                                    var v89 float32 = 1.5
                                    _ = v89
                                }
                            }
                            fmt.Println(fi86)
                        } else {
                            v90 := GenS6{F0: "hello", F1: true}
                            switch 0 {
                            case -1:
                                for fi91 := 0; fi91 < 17; fi91++ {
                                    var v92 bool = true
                                    v92 = (true && (4294967295 <= 956))
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = fi91
                                    _ = v92
                                }
                            case 778:
                                fmt.Println(fi86, v3, fi86)
                            default:
                                if ("a" >= "") {
                                    var v93 int16 = -1
                                    _ = v93
                                }
                            }
                            v94 := GenS8{F0: (fi86 ^ fi86), F1: -677}
                            _ = v90
                            _ = v94
                        }
                        _ = fi86
                        _ = v87
                    }
                    if true {
                        v95 := GenS3{F0: 607, F1: 239, F2: 4294967295}
                        if (!true) {
                            for fi96 := 0; fi96 < 19; fi96++ {
                                v2 = (v2 >> ((0) & 63))
                                _ = fi96
                            }
                        } else {
                            v2 -= -330
                            var v97 uint16 = 65535
                            var v98 int16 = -1
                            _ = v97
                            _ = v98
                        }
                        _ = v95
                    } else {
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                        switch "foo bar" {
                        case "\n\t":
                            var v99 int16 = 1000
                            _ = v99
                        default:
                            v100 := GenS9{F0: 42}
                            _ = v100
                        }
                        for fi101 := 0; fi101 < 4; fi101++ {
                            _ = len(v4)
                            v3 += (v3 / (((v3 | v3)) | 1))
                            if true {
                                v3 -= (-(v3))
                            }
                            _ = fi101
                        }
                    }
                }
            }
            v2 = v2
        }
    default:
        if (float32((float64(0.1) - (3.14159 + -2.25))) == -2.25) {
            for fi102 := 0; fi102 < 3; fi102++ {
                switch (^(fi102)) {
                case 457:
                    var v103 uint = 2
                    _ = v103
                }
                _ = fi102
            }
            var v104 int16 = -627
            _ = v104
        } else {
            var v105 uint16 = 791
            var v106 int = -841
            _ = v105
            _ = v106
        }
    }
    v3 = (v3 | 9223372036854775807)
    if false {
        for fi107 := 0; fi107 < 16; fi107++ {
            for fi108 := 0; fi108 < 7; fi108++ {
                v3 = (v3 ^ v3)
                _ = len(v4)
                _ = len(v4)
                _ = fi108
            }
            for fi109 := 0; fi109 < 20; fi109++ {
                v110 := GenS3{F0: 65535, F1: 42, F2: 2147483648}
                _ = fi109
                _ = v110
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi107
        }
        if (-2.25 != 1e10) {
            var v111 string = (("hello" + "hello") + "hello")
            _ = v111
        }
        if false {
            if (!true) {
                if (!true) {
                    if true {
                        if (("a" + "") > "") {
                            var v112 uint = 255
                            _ = v112
                        } else {
                            fmt.Println(v3)
                            fmt.Println(v2, v3)
                            if (!true) {
                                if false {
                                    if (true && true) {
                                        var v113 int16 = -163
                                        for fi114 := 0; fi114 < 11; fi114++ {
                                            var v115 int64 = v3
                                            _ = fi114
                                            _ = v115
                                        }
                                        _ = v113
                                    }
                                    v1 = append(v1, ("foo bar" + "foo bar"))
                                    var v116 uint32 = 306
                                    _ = v116
                                } else {
                                    if ((v3 >> ((841) & 63)) <= (-(v3))) {
                                        fmt.Println(v3, v2)
                                        v2 += (v2 | (-(v2)))
                                        if true {
                                            if (461 == 4294967295) {
                                                fmt.Println(v3, v2)
                                                if false {
                                                    var v117 uint64 = 689
                                                    if (128 == 128) {
                                                        if true {
                                                            for fi118 := 0; fi118 < 9; fi118++ {
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                _ = fi118
                                                            }
                                                            v119 := GenS10{F0: 746}
                                                            _ = v119
                                                        } else {
                                                            switch (("a" + (("\n\t" + ("" + "")) + "")) + "hello") {
                                                            case "\n\t":
                                                                v120 := GenS10{F0: 65535}
                                                                _ = v120
                                                            case "":
                                                                _ = len(v4)
                                                            default:
                                                                for fi121 := 0; fi121 < 14; fi121++ {
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    fmt.Println(v2)
                                                                    var v122 uint8 = 202
                                                                    _ = fi121
                                                                    _ = v122
                                                                }
                                                            }
                                                        }
                                                        fmt.Println(v2, v3)
                                                    }
                                                    var v123 int32 = v2
                                                    _ = v117
                                                    _ = v123
                                                } else {
                                                    v2 = -2147483648
                                                }
                                            } else {
                                                switch 32767 {
                                                case 4294967295:
                                                    v4 = append(v4, "foo bar")
                                                case 0:
                                                    var v124 uint64 = 175
                                                    _ = v124
                                                }
                                            }
                                            if (128 == 0) {
                                                if true {
                                                    for fi125 := 0; fi125 < 18; fi125++ {
                                                        if false {
                                                            if (true && (false && ((false || (true || true)) && true))) {
                                                                fmt.Println(v3, v2, v3)
                                                                v2 = (^(v2))
                                                                var v126 int = fi125
                                                                _ = v126
                                                            }
                                                            v3 = v3
                                                            var v127 int16 = -1
                                                            _ = v127
                                                        }
                                                        _ = fi125
                                                    }
                                                } else {
                                                    if true {
                                                        _ = len(v4)
                                                    } else {
                                                        v2 -= 2147483647
                                                        var v128 float32 = 3.14159
                                                        _ = v128
                                                    }
                                                    var v129 int64 = -1
                                                    _ = len(v4)
                                                    _ = v129
                                                }
                                            }
                                            var v130 bool = false
                                            _ = v130
                                        } else {
                                            fmt.Println(v3, v3)
                                        }
                                    } else {
                                        v131 := GenS2{F0: "\n\t", F1: true, F2: -128}
                                        v3 += v3
                                        var v132 int16 = 724
                                        _ = v131
                                        _ = v132
                                    }
                                    v2 -= v2
                                }
                                v133 := GenS4{F0: 128, F1: v3}
                                fmt.Println(v2, v2)
                                _ = v133
                            } else {
                                fmt.Println(v3, v3)
                            }
                        }
                        v3 = (v3 / ((v3) | 1))
                        v2 += (v2 | (-(v2)))
                    }
                    var v134 int16 = 32767
                    _ = v134
                }
                fmt.Println(v3)
            }
        } else {
            var v135 string = "a"
            v1 = append(v1, "\n\t")
            _ = v135
        }
    }
    _ = len(v4)
    v136 := GenS7{F0: ("hello" + ("\n\t" + "hello"))}
    var v137 uint = 2
    v3 -= v3
    var v138 int8 = -13
    var v139 uint32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v137:", v137)
    fmt.Println("v138:", v138)
    fmt.Println("v139:", v139)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v136
    _ = v137
    _ = v138
    _ = v139
    _ = v2
    _ = v3
    _ = v4
    _ = v67
}
