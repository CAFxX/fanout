package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 string; F1 int32; F2 string }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 int64; F1 uint64 }
type GenS7 struct { F0 bool; F1 uint8 }
type GenS8 struct { F0 int16; F1 uint64 }
type GenS9 struct { F0 int64; F1 bool; F2 int32 }
type GenS10 struct { F0 string; F1 uint64; F2 bool }
type GenS11 struct { F0 uint32; F1 int; F2 int8 }
type GenS12 struct { F0 int; F1 uint8; F2 uint16 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int; F1 int8; F2 uint }
func genFn0(p0 uint32, p1 int, p2 bool) int8 {
    p0 -= (^(p0))
    if p2 {
        v7 := []int{(p1 / ((2147483647) | 1)), 32768, -503}
        for fi8 := 0; fi8 < 18; fi8++ {
            p0 = p0
            _ = fi8
        }
        var v9 uint64 = 4294967295
        _ = v7
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := []int{(fi10 << ((410) & 63))}
        if (1000 <= -1) {
            p1 = (fi10 ^ 32767)
            var v12 uint32 = p0
            _ = v12
        }
        switch "a" {
        case "foo bar":
            if len(v11) > 0 {
                v11[0] = 0
            }
        case "":
            for fi13 := 0; fi13 < 5; fi13++ {
                var v14 int16 = 1000
                var v15 int64 = 42
                _ = fi13
                _ = v14
                _ = v15
            }
        case "\n\t":
            switch p1 {
            case -744:
                switch "foo bar" {
                case "":
                    if (!((1 <= 255) && (-2.25 < 1.5))) {
                        if p2 {
                            var v16 uint8 = 112
                            var v17 int = (-(fi10))
                            _ = len(v11)
                            _ = v16
                            _ = v17
                        } else {
                            for fi18 := 0; fi18 < 18; fi18++ {
                                switch (p1 | (fi18 << ((4294967295) & 63))) {
                                case 2147483647:
                                    p1 = (fi18 - int(p0))
                                case 127:
                                    var v19 uint64 = 508
                                    _ = v19
                                default:
                                    var v20 uint64 = 4294967295
                                    _ = v20
                                }
                                _ = fi18
                            }
                            p1 -= (p1 >> ((65535) & 63))
                            for fi21 := 0; fi21 < 10; fi21++ {
                                v11 = append(v11, -629)
                                if p2 {
                                    p0 -= p0
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                    switch "\n\t" {
                                    case "":
                                        v22 := GenS1{F0: 9223372036854775807}
                                        _ = v22
                                    case "\n\t":
                                        v23 := GenS1{F0: (p1 << ((255) & 63))}
                                        _ = v23
                                    }
                                }
                                if len(v11) > 0 {
                                    v11[0] = p1
                                }
                                _ = fi21
                            }
                        }
                        var v24 int16 = 0
                        _ = v24
                    }
                case "\n\t":
                    if ((3.14159 / (1.5 + 0.0)) < (1e10 - (1.5 - 0.1))) {
                        for fi25 := 0; fi25 < 15; fi25++ {
                            v26 := GenS2{F0: 18446744073709551615, F1: 1, F2: 0}
                            _ = fi25
                            _ = v26
                        }
                        if (p2 || (4294967295 != 570)) {
                            switch "hello" {
                            case "foo bar":
                                if p2 {
                                    v11 = append(v11, 1)
                                } else {
                                    var v27 int16 = 446
                                    if len(v11) > 0 {
                                        v11[0] = fi10
                                    }
                                    _ = v27
                                }
                            default:
                                for fi28 := 0; fi28 < 19; fi28++ {
                                    _ = len(v11)
                                    v29 := GenS3{F0: 1, F1: 42}
                                    p0 += p0
                                    _ = fi28
                                    _ = v29
                                }
                            }
                            switch "foo bar" {
                            case "hello":
                                p1 = (fi10 % (((p1 << ((65535) & 63))) | 1))
                            }
                            switch p1 {
                            case 32768:
                                var v30 float64 = 1.5
                                _ = v30
                            case 4294967295:
                                v31 := GenS4{F0: ((("\n\t" + ("\n\t" + ((("hello" + "") + ("hello" + "a")) + "foo bar"))) + "foo bar") + "foo bar"), F1: 42, F2: "hello"}
                                _ = v31
                            }
                        } else {
                            if p2 {
                                p2 = (42 <= 865)
                                v32 := GenS2{F0: 1, F1: int32(p0), F2: 0}
                                _ = v32
                            }
                        }
                        p1 = fi10
                    }
                case "foo bar":
                    _ = len(v11)
                default:
                    switch (^(fi10)) {
                    case -468:
                        var v33 int8 = -128
                        _ = v33
                    default:
                        v11 = append(v11, fi10)
                    }
                }
            default:
                v34 := GenS5{F0: ((-2.25 / ((3.14159 * 0.1) / 1e10)) != (-2.25 / ((((-2.25 * 0.0) - -2.25) / 1.0) - 3.14159))), F1: 4294967295}
                _ = v34
            }
        default:
            if (true || (!false)) {
                if p2 {
                    var v35 uint32 = p0
                    v36 := GenS4{F0: "\n\t", F1: 1, F2: ("\n\t" + "foo bar")}
                    for fi37 := 0; fi37 < 20; fi37++ {
                        fmt.Println(fi37, p1, fi37)
                        var v38 uint = 2
                        v39 := GenS6{F0: -974, F1: 45}
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                    _ = v35
                    _ = v36
                } else {
                    var v40 bool = true
                    _ = v40
                }
                p1 = 42
            } else {
                for fi41 := 0; fi41 < 7; fi41++ {
                    p2 = (!(81 > 42))
                    v42 := GenS7{F0: p2, F1: 141}
                    _ = fi41
                    _ = v42
                }
                fmt.Println(p1, fi10, fi10)
                for fi43 := 0; fi43 < 4; fi43++ {
                    var v44 string = ("a" + ("\n\t" + ""))
                    v44 = (v44 + v44)
                    var v45 uint = 544
                    _ = fi43
                    _ = v44
                    _ = v45
                }
            }
        }
        _ = fi10
        _ = v11
    }
    p1 -= (p1 << ((0) & 63))
    if p2 {
        v46 := GenS8{F0: 1000, F1: 2}
        _ = v46
    } else {
        var v47 bool = ((p2 || true) || (65535 >= (p0 ^ p0)))
        v47 = (!(1 >= 1))
        var v48 uint32 = p0
        _ = v47
        _ = v48
    }
    _ = p0
    _ = p1
    _ = p2
    return -128
}
func genFn1() int32 {
    for fi49 := 0; fi49 < 17; fi49++ {
        if false {
            v50 := []int{(fi49 - 4294967295), fi49, 530, fi49}
            var v51 int32 = -999
            _ = v50
            _ = v51
        } else {
            if (136 < 135) {
                fmt.Println(fi49)
                v52 := []int{(fi49 & 32767), (fi49 + -517), (fi49 & (fi49 | fi49)), -886}
                var v53 int8 = 100
                _ = v52
                _ = v53
            } else {
                var v54 int64 = 362
                var v55 float32 = 0.1
                if true {
                    if (148 < 32767) {
                        for fi56 := 0; fi56 < 4; fi56++ {
                            var v57 float32 = v55
                            v54 -= (v54 + (v54 % (((v54 % ((-832) | 1))) | 1)))
                            v58 := GenS2{F0: 1000, F1: -2147483648, F2: 42}
                            fmt.Println(v57, v54, v57)
                            fmt.Println(fi49, fi56)
                            _ = fi56
                            _ = v57
                            _ = v58
                        }
                        v59 := GenS9{F0: (v54 & (v54 | -1)), F1: true, F2: -999}
                        fmt.Println(v54, v55, v55)
                        _ = v59
                    }
                    if (32767 < 1) {
                        for fi60 := 0; fi60 < 13; fi60++ {
                            v54 -= v54
                            var v61 int8 = 127
                            _ = fi60
                            _ = v61
                        }
                    } else {
                        v62 := []int{fi49, 127}
                        if len(v62) > 0 {
                            v62[0] = fi49
                        }
                        switch "" {
                        case "hello":
                            var v63 uint32 = 2147483648
                            _ = v63
                        case "foo bar":
                            fmt.Println(fi49, v54)
                        }
                        _ = v62
                    }
                    v64 := []int{fi49, (fi49 >> ((65535) & 63)), (fi49 | -2147483648), fi49}
                    _ = v64
                } else {
                    fmt.Println(v55)
                    if (42 == 291) {
                        var v65 float64 = 1e10
                        fmt.Println(v54, v55, v54)
                        fmt.Println(fi49)
                        _ = v65
                    }
                }
                _ = v54
                _ = v55
            }
            var v66 int8 = 42
            switch -420 {
            case 32768:
                if true {
                    for fi67 := 0; fi67 < 6; fi67++ {
                        switch (^(fi49)) {
                        case 0:
                            var v68 float32 = float32(1e10)
                            _ = v68
                        case -363:
                            if (false && false) {
                                v69 := GenS7{F0: (((true || true) && (genFn0(4294967295, fi49, true) != (v66 + genFn0(0, fi67, false)))) || (!true)), F1: 2}
                                switch ("foo bar" + ("hello" + "hello")) {
                                case "\n\t":
                                    var v70 int = (fi67 - (fi67 / ((fi49) | 1)))
                                    v66 = (-(v66))
                                    switch v70 {
                                    case -9223372036854775808:
                                        if (true || (fi49 < fi49)) {
                                            if (!true) {
                                                var v71 float64 = (0.1 + 0.0)
                                                if true {
                                                    v70 = (fi67 * 32768)
                                                    var v72 string = "a"
                                                    _ = v72
                                                } else {
                                                    v73 := GenS1{F0: (fi49 & (v70 >> ((0) & 63)))}
                                                    _ = v73
                                                }
                                                _ = v71
                                            }
                                            v70 -= (^(v70))
                                            v74 := []int{(fi49 >> ((18446744073709551615) & 63))}
                                            if false {
                                                v66 += (v66 + 19)
                                                fmt.Println(v66, fi49)
                                                v66 += (v66 & genFn0(32, fi67, (!true)))
                                            }
                                            _ = v74
                                        } else {
                                            v75 := GenS8{F0: 0, F1: 523}
                                            var v76 int = (fi49 - -614)
                                            _ = v75
                                            _ = v76
                                        }
                                    default:
                                        v70 -= 673
                                        switch ("a" + "foo bar") {
                                        case "\n\t":
                                            if (264 != 18446744073709551615) {
                                                v77 := []int{402, (fi67 ^ (v70 - 127))}
                                                if len(v77) >= 2 {
                                                    v77 = v77[1:len(v77)]
                                                }
                                                _ = v77
                                            }
                                        default:
                                            v70 = fi67
                                        }
                                    }
                                    _ = v70
                                default:
                                    v66 -= v66
                                }
                                _ = v69
                            }
                        }
                        if (42 != -9223372036854775808) {
                            if true {
                                fmt.Println(fi67)
                            }
                            var v78 float32 = ((1e10 * (3.14159 / 0.1)) - 1.5)
                            var v79 float64 = 0.1
                            _ = v78
                            _ = v79
                        } else {
                            var v80 string = "foo bar"
                            v81 := []int{(fi67 << ((65535) & 63)), fi49}
                            _ = v80
                            _ = v81
                        }
                        _ = fi67
                    }
                }
            case 828:
                if (true && ((260 < 513) || ("\n\t" == "a"))) {
                    v66 -= -128
                }
            default:
                v82 := []int{fi49}
                _ = v82
            }
            _ = v66
        }
        _ = fi49
    }
    var v83 int = 821
    for fi84 := 0; fi84 < 10; fi84++ {
        v85 := []int{}
        if len(v85) >= 2 {
            v85 = v85[1:len(v85)]
        }
        fmt.Println(v83, v83)
        _ = fi84
        _ = v85
    }
    v83 = (v83 | (v83 - v83))
    switch "\n\t" {
    case "":
        for fi86 := 0; fi86 < 7; fi86++ {
            if false {
                switch "a" {
                case "\n\t":
                    if ((-1 < -929) || true) {
                        var v87 bool = (65535 > 0)
                        _ = v87
                    }
                default:
                    v88 := []int{fi86, 1, (v83 & (-(v83))), 0}
                    _ = v88
                }
                if false {
                    var v89 int = 127
                    _ = v89
                }
            } else {
                fmt.Println(v83, fi86)
            }
            v90 := []int{v83, fi86}
            _ = fi86
            _ = v90
        }
    case "foo bar":
        var v91 bool = true
        _ = v91
    }
    _ = v83
    return -1
}

func main() {
    var v1 int32 = 1
    var v2 uint64 = 4294967295
    v3 := make(map[uint32]bool)
    var v4 uint64 = v2
    var v5 uint64 = v4
    var v6 int64 = -306
    switch "" {
    case "a":
        var v92 int32 = -894
        _ = v92
    default:
        var v93 int64 = (-(v6))
        _ = v93
    }
    v94 := GenS10{F0: "foo bar", F1: v2, F2: false}
    if ((!(v1 >= v1)) || (-798 >= 1)) {
        v95 := v3[886]
        _ = v95
    } else {
        for fi96 := 0; fi96 < 8; fi96++ {
            for fi97 := 0; fi97 < 5; fi97++ {
                var v98 string = "\n\t"
                if ((v6 << ((1) & 63)) >= (v6 * (v6 * -9223372036854775808))) {
                    _ = len(v3)
                    var v99 float32 = 1.5
                    v100 := GenS11{F0: 4294967295, F1: -2147483648, F2: genFn0(2147483647, (fi97 / (((-(fi96))) | 1)), (!false))}
                    _ = v100
                    _ = v99
                } else {
                    fmt.Println(v98, v98, v98)
                    fmt.Println(v2, fi96)
                }
                _ = fi97
                _ = v98
            }
            v101 := []int{(fi96 / ((9223372036854775807) | 1)), (fi96 >> ((2) & 63)), 0}
            switch ("foo bar" + ("foo bar" + "")) {
            case "\n\t":
                var v102 uint8 = 101
                _ = v102
            case "foo bar":
                v103 := GenS8{F0: -669, F1: (v5 & 30)}
                _ = v103
            case "a":
                v4 += 18446744073709551615
            }
            _ = fi96
            _ = v101
        }
    }
    for fi104 := 0; fi104 < 10; fi104++ {
        v2 += v4
        fmt.Println(v6, v6)
        _ = fi104
    }
    var v105 int = 42
    if true {
        v106 := GenS10{F0: "", F1: v5, F2: ((1 != 100) || false)}
        if false {
            v1 += v1
        } else {
            v107 := GenS10{F0: (("a" + ("foo bar" + "foo bar")) + ("hello" + ("hello" + ("hello" + ("a" + "a"))))), F1: v2, F2: (true || (true && (!(2147483647 < v105))))}
            _ = v107
        }
        if false {
            var v108 float32 = ((0.0 - 1.5) - float32(1e10))
            _ = v108
        }
        _ = v106
    }
    v109 := GenS8{F0: 716, F1: v4}
    for fi110 := 0; fi110 < 16; fi110++ {
        var v111 uint = 4294967295
        for fi112 := 0; fi112 < 9; fi112++ {
            switch "\n\t" {
            case "":
                if (65535 < v111) {
                    var v113 string = ("foo bar" + "foo bar")
                    _ = v113
                } else {
                    switch v105 {
                    case -670:
                        var v114 string = "a"
                        _ = v114
                    case 2147483647:
                        v115 := GenS12{F0: fi110, F1: 44, F2: 769}
                        _ = v115
                    }
                    if (("" + ("foo bar" + "")) == "hello") {
                        var v116 uint16 = 2
                        v116 = v116
                        _ = v116
                    } else {
                        switch "" {
                        case "hello":
                            if ((0.1 * (((3.14159 + 1e10) - 1e10) + 3.14159)) < ((0.0 * (float32((3.14159 / 1.0)) / 3.14159)) / (((0.0 + (float32(1.5) * 1e10)) / (0.0 + 1e10)) + 1e10))) {
                                var v117 int = v105
                                if (1 != -128) {
                                    v6 -= 9223372036854775807
                                    fmt.Println(v1, v111, fi110)
                                    v118 := []int{-2147483648}
                                    var v119 bool = false
                                    switch ("\n\t" + (((("hello" + "\n\t") + ("hello" + "")) + "hello") + ("" + ""))) {
                                    case "hello":
                                        v120 := GenS13{F0: 2147483648}
                                        fmt.Println(v4)
                                        _ = v120
                                    case "":
                                        if v119 {
                                            fmt.Println(v2, v2)
                                        }
                                    }
                                    var v121 int = int(v111)
                                    _ = v118
                                    _ = v119
                                    _ = v121
                                } else {
                                    if (!true) {
                                        switch v117 {
                                        case -9223372036854775808:
                                            switch 32768 {
                                            case 4294967295:
                                                v5 = v5
                                            case -9223372036854775808:
                                                v122 := []int{}
                                                _ = v122
                                            }
                                        case -834:
                                            v123 := GenS11{F0: 2, F1: fi110, F2: 127}
                                            _ = v123
                                        case -634:
                                            v124 := v3[653]
                                            _ = v124
                                        default:
                                            v125 := v3[2147483647]
                                            _ = v125
                                        }
                                        v126 := []int{}
                                        var v127 string = "hello"
                                        v105 = v117
                                        _ = v126
                                        _ = v127
                                    }
                                    var v128 uint8 = 1
                                    var v129 int16 = -433
                                    fmt.Println(v128)
                                    v128 -= (^(v128))
                                    v3[2147483647] = false
                                    var v130 string = "a"
                                    v1 = v1
                                    fmt.Println(v128, v4, v111)
                                    _ = v128
                                    _ = v129
                                    _ = v130
                                }
                                _ = v117
                            }
                        case "foo bar":
                            v131 := GenS13{F0: 42}
                            _ = v131
                        case "":
                            if (!(false && ((("hello" + "foo bar") > "a") && false))) {
                                if (-128 != 0) {
                                    v105 += (-(v105))
                                } else {
                                    var v132 int64 = v6
                                    if (true && false) {
                                        switch ("hello" + "") {
                                        case "hello":
                                            fmt.Println(fi110, v5, v2)
                                            var v133 bool = (true || true)
                                            _ = v133
                                        case "foo bar":
                                            var v134 int32 = (v1 / ((genFn1()) | 1))
                                            _ = v134
                                        default:
                                            fmt.Println(v1, v1, v1)
                                        }
                                    } else {
                                        v132 -= v6
                                        switch ("\n\t" + (("\n\t" + "\n\t") + "hello")) {
                                        case "foo bar":
                                            v135 := GenS12{F0: v105, F1: 255, F2: 255}
                                            if true {
                                                v132 = v6
                                            }
                                            var v136 float32 = (3.14159 * 0.0)
                                            _ = v135
                                            _ = v136
                                        default:
                                            v1 -= (v1 | (v1 - (v1 << ((0) & 63))))
                                            v137 := GenS14{F0: (fi110 - -9223372036854775808), F1: 127, F2: v111}
                                            _ = v137
                                        }
                                        var v138 uint = (v111 << (((v111 % ((2) | 1))) & 63))
                                        _ = v138
                                    }
                                    _ = v132
                                }
                                v1 = genFn1()
                            } else {
                                v139 := GenS1{F0: (-(fi112))}
                                if (v1 <= (v1 / ((-1) | 1))) {
                                    var v140 string = ("\n\t" + ("\n\t" + "foo bar"))
                                    if false {
                                        var v141 bool = true
                                        _ = v141
                                    } else {
                                        v111 += v111
                                        switch "hello" {
                                        case "hello":
                                            v6 = (v6 ^ (v6 * -9223372036854775808))
                                        case "a":
                                            delete(v3, 898)
                                        }
                                        if (2 > 0) {
                                            v3[65535] = (!(42 < 0))
                                            delete(v3, 2)
                                            v142 := GenS6{F0: 42, F1: (^(v2))}
                                            _ = v142
                                        } else {
                                            if true {
                                                var v143 uint16 = 0
                                                _ = v143
                                            }
                                            v144 := GenS14{F0: (fi112 - v105), F1: 127, F2: v111}
                                            v140 = (v140 + "\n\t")
                                            v1 = genFn1()
                                            _ = v144
                                        }
                                        var v145 int32 = v1
                                        _ = v145
                                    }
                                    _ = v140
                                } else {
                                    v1 = 0
                                    fmt.Println(v5)
                                }
                                v2 -= 18446744073709551615
                                _ = v139
                            }
                        default:
                            fmt.Println(fi110, fi110, v1)
                        }
                    }
                }
            default:
                v146 := GenS6{F0: (v6 / ((v6) | 1)), F1: (^(v5))}
                _ = v146
            }
            _ = fi112
        }
        var v147 bool = true
        _ = fi110
        _ = v111
        _ = v147
    }
    switch "foo bar" {
    case "a":
        if true {
            switch ("\n\t" + ("foo bar" + ("hello" + "foo bar"))) {
            case "\n\t":
                for fi148 := 0; fi148 < 11; fi148++ {
                    var v149 string = "hello"
                    _ = fi148
                    _ = v149
                }
            case "hello":
                if true {
                    if true {
                        var v150 uint = 786
                        _ = v150
                    }
                }
            case "foo bar":
                fmt.Println(v4, v105)
            }
            v3[0] = (!((255 < 42) || true))
        }
    case "\n\t":
        var v151 uint64 = v5
        _ = v151
    case "foo bar":
        var v152 bool = true
        _ = v152
    }
    fmt.Println(v2)
    v3[2147483648] = false
    fmt.Println("v1:", v1)
    fmt.Println("v105:", v105)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v105
    _ = v109
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v94
}
