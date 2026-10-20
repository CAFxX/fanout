package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS6 struct { F0 uint16; F1 int64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int32; F1 int32; F2 uint8 }
type GenS10 struct { F0 uint; F1 int64; F2 int64 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 bool; F1 int }
type GenS13 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int; F1 int32; F2 uint }
type GenS16 struct { F0 int8; F1 int8; F2 bool }
type GenS17 struct { F0 int; F1 int64; F2 uint32 }
type GenS18 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS19 struct { F0 string; F1 uint; F2 uint64 }
type GenS20 struct { F0 int64; F1 int16 }
func genFn0() int64 {
    var v4 string = ("hello" + "a")
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := []int{}
        switch ((v4 + "foo bar") + v4) {
        case "\n\t":
            var v7 int16 = -32768
            _ = v7
        case "hello":
            if (!false) {
                if (42 <= 1) {
                    var v8 float64 = 3.14159
                    if true {
                        v8 += (v8 - v8)
                        for fi9 := 0; fi9 < 12; fi9++ {
                            v10 := GenS1{F0: 42}
                            v4 = v4
                            if false {
                                if ((42 <= 65535) && false) {
                                    v11 := GenS2{F0: 32767}
                                    v8 = 1e10
                                    _ = v11
                                } else {
                                    fmt.Println(fi5, v4)
                                }
                                v6 = append(v6, fi9)
                            } else {
                                var v12 uint8 = 211
                                if ((float32(v8) >= float32(((v8 + -2.25) / v8))) || (false && (42 != 255))) {
                                    v4 = (v4 + "hello")
                                }
                                var v13 int32 = 0
                                _ = v12
                                _ = v13
                            }
                            _ = fi9
                            _ = v10
                        }
                        var v14 float64 = v8
                        _ = v14
                    }
                    _ = v8
                } else {
                    var v15 int = fi5
                    v16 := GenS3{F0: 2147483648, F1: 674}
                    v15 = v15
                    _ = v15
                    _ = v16
                }
                v17 := GenS4{F0: 214}
                v6 = append(v6, (fi5 << ((42) & 63)))
                _ = v17
            }
        default:
            switch (fi5 * (fi5 << ((1) & 63))) {
            case 1:
                var v18 uint8 = 127
                _ = v18
            case -2147483648:
                var v19 uint32 = 418
                _ = v19
            case -406:
                v4 = v4
            }
        }
        _ = fi5
        _ = v6
    }
    var v20 uint64 = 0
    v21 := []int{742}
    for fi22 := 0; fi22 < 6; fi22++ {
        for fi23 := 0; fi23 < 16; fi23++ {
            var v24 bool = true
            _ = len(v21)
            _ = fi23
            _ = v24
        }
        var v25 int16 = 1000
        _ = fi22
        _ = v25
    }
    _ = v20
    _ = v21
    _ = v4
    return 42
}
func genFn1(p0 int16, p1 int16, p2 uint32) int32 {
    var v26 float32 = 1.5
    fmt.Println(p1)
    if false {
        for fi27 := 0; fi27 < 15; fi27++ {
            if false {
                var v28 int64 = 42
                _ = v28
            } else {
                if false {
                    v26 += -2.25
                } else {
                    if ((!(0 >= 1)) && false) {
                        v29 := []int{(fi27 * (fi27 / ((fi27) | 1))), fi27}
                        var v30 string = "hello"
                        if (1 >= -2147483648) {
                            p0 += (^(p0))
                            v26 = (v26 - float32((1.5 / (float64((((float32(0.0) / float32(3.14159)) + float32(1e10)) - v26)) + (0.0 + float64(float32(1.5)))))))
                        }
                        _ = v29
                        _ = v30
                    }
                    for fi31 := 0; fi31 < 15; fi31++ {
                        p1 -= 1000
                        var v32 int16 = (p0 | (-(p0)))
                        _ = fi31
                        _ = v32
                    }
                }
                for fi33 := 0; fi33 < 2; fi33++ {
                    var v34 int8 = 76
                    switch fi33 {
                    case 32767:
                        v26 = (v26 + 1.5)
                    case -2147483648:
                        v35 := []int{(-(fi33)), (fi33 - (fi33 * (fi27 / (((^(fi27))) | 1))))}
                        _ = v35
                    case 4294967295:
                        if (!(!(float32((0.1 - 0.0)) != 0.0))) {
                            if false {
                                switch (fi33 % ((fi33) | 1)) {
                                case 4294967295:
                                    if (("foo bar" + "foo bar") != "hello") {
                                        var v36 int32 = 960
                                        v37 := GenS4{F0: 0}
                                        _ = v36
                                        _ = v37
                                    }
                                }
                                fmt.Println(p2, fi33)
                                p1 = p1
                            }
                            p0 -= -175
                            if ((^(v34)) == v34) {
                                if false {
                                    var v38 int64 = genFn0()
                                    v34 = v34
                                    v39 := GenS5{F0: ((false && true) || (!false)), F1: 2, F2: 954}
                                    _ = v38
                                    _ = v39
                                }
                            }
                        }
                    default:
                        v40 := GenS6{F0: 2, F1: 405}
                        _ = v40
                    }
                    switch 625 {
                    case -2147483648:
                        p2 = (p2 << ((998) & 63))
                    }
                    _ = fi33
                    _ = v34
                }
                if (0.1 > 0.0) {
                    p0 -= p1
                    fmt.Println(p0, p2)
                }
            }
            for fi41 := 0; fi41 < 20; fi41++ {
                var v42 uint16 = 747
                v42 = v42
                _ = fi41
                _ = v42
            }
            switch "" {
            case "foo bar":
                for fi43 := 0; fi43 < 20; fi43++ {
                    var v44 float32 = ((v26 / 1.0) + (1.5 * v26))
                    v45 := GenS1{F0: 42}
                    var v46 float32 = v44
                    _ = fi43
                    _ = v44
                    _ = v45
                    _ = v46
                }
            default:
                p1 -= 269
            }
            _ = fi27
        }
        var v47 uint8 = 1
        _ = v47
    } else {
        if (true || (255 < 161)) {
            var v48 uint = 42
            if true {
                var v49 uint8 = 0
                v50 := []int{-188, 4294967295, 4294967295}
                _ = v49
                _ = v50
            } else {
                v51 := []int{-9223372036854775808, 950, 32768}
                v52 := GenS3{F0: (p2 * 4294967295), F1: 189}
                p2 -= (p2 << ((v48) & 63))
                _ = v51
                _ = v52
            }
            switch "hello" {
            case "foo bar":
                v53 := GenS7{F0: (v48 % ((255) | 1))}
                _ = v53
            case "\n\t":
                switch "\n\t" {
                case "hello":
                    v48 = (v48 ^ 65535)
                }
            default:
                var v54 int16 = p1
                _ = v54
            }
            _ = v48
        }
        if true {
            switch 848 {
            case 491:
                v55 := GenS1{F0: 2}
                _ = v55
            case 241:
                if (!true) {
                    if true {
                        if true {
                            for fi56 := 0; fi56 < 11; fi56++ {
                                switch -53 {
                                case -2147483648:
                                    for fi57 := 0; fi57 < 9; fi57++ {
                                        v58 := GenS7{F0: 939}
                                        var v59 bool = (true || true)
                                        v59 = (18446744073709551615 == 2)
                                        _ = fi57
                                        _ = v58
                                        _ = v59
                                    }
                                case 32768:
                                    p1 = -299
                                default:
                                    v26 = v26
                                }
                                fmt.Println(p0)
                                for fi60 := 0; fi60 < 4; fi60++ {
                                    v61 := GenS3{F0: p2, F1: 121}
                                    v62 := []int{fi60, fi60, 574, fi56}
                                    v26 -= v26
                                    _ = fi60
                                    _ = v61
                                    _ = v62
                                }
                                _ = fi56
                            }
                            v63 := []int{128, -2147483648, -338}
                            var v64 int16 = -32768
                            _ = v63
                            _ = v64
                        }
                        var v65 bool = false
                        _ = v65
                    }
                    fmt.Println(p0)
                } else {
                    if (!(!true)) {
                        v66 := GenS2{F0: p0}
                        _ = v66
                    }
                    p1 -= p1
                    if true {
                        p2 = 0
                        v26 += (float32(1.5) * v26)
                    }
                }
            case -2147483648:
                for fi67 := 0; fi67 < 16; fi67++ {
                    fmt.Println(p1)
                    p2 = (p2 - p2)
                    _ = fi67
                }
            default:
                p2 += 2
            }
            v26 -= v26
            switch -971 {
            case 32768:
                v26 = v26
            case 42:
                var v68 uint64 = 18446744073709551615
                _ = v68
            case 9223372036854775807:
                var v69 uint32 = 0
                _ = v69
            default:
                v70 := []int{}
                _ = v70
            }
        } else {
            v71 := []int{}
            for fi72 := 0; fi72 < 16; fi72++ {
                if ((-(fi72)) <= fi72) {
                    p0 = p1
                    var v73 int32 = -2147483648
                    if ((p1 + p0) >= (p1 & (p1 | p1))) {
                        for fi74 := 0; fi74 < 15; fi74++ {
                            var v75 int64 = 9223372036854775807
                            if (100 >= -83) {
                                v71 = append(v71, fi74)
                                if true {
                                    var v76 uint8 = 124
                                    _ = v76
                                }
                            }
                            _ = len(v71)
                            _ = fi74
                            _ = v75
                        }
                        if false {
                            v77 := GenS8{F0: -18}
                            v78 := GenS9{F0: 1, F1: (v73 | 1), F2: 127}
                            for fi79 := 0; fi79 < 4; fi79++ {
                                var v80 int8 = int8(p1)
                                _ = fi79
                                _ = v80
                            }
                            _ = v77
                            _ = v78
                        }
                    }
                    _ = v73
                } else {
                    if false {
                        v81 := GenS10{F0: 207, F1: genFn0(), F2: 9223372036854775807}
                        for fi82 := 0; fi82 < 11; fi82++ {
                            p0 += p1
                            _ = fi82
                        }
                        _ = v81
                    } else {
                        if (127 > -103) {
                            v26 -= (-2.25 + 1e10)
                        }
                        switch "\n\t" {
                        case "a":
                            p0 += (p1 << ((333) & 63))
                        case "foo bar":
                            _ = len(v71)
                        case "":
                            v71 = append(v71, 963)
                        }
                        v26 += v26
                    }
                    v83 := GenS4{F0: uint8(p2)}
                    _ = v83
                }
                _ = fi72
            }
            v84 := GenS11{F0: 18446744073709551615}
            _ = v71
            _ = v84
        }
        var v85 float32 = -2.25
        _ = v85
    }
    v86 := []int{}
    if true {
        v26 -= 0.1
        v86 = append(v86, -2147483648)
    } else {
        v87 := GenS12{F0: (260 != 42), F1: 859}
        v88 := GenS13{F0: (p2 % (((^(p2))) | 1)), F1: 2, F2: genFn0()}
        switch (("hello" + "foo bar") + ("" + "foo bar")) {
        case "a":
            switch -88 {
            case -555:
                p0 += -32
            case -466:
                v86 = append(v86, 128)
            }
        }
        _ = v87
        _ = v88
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v86
    return int32(p2)
}

func main() {
    var v1 uint8 = 255
    var v2 float32 = (1.5 / 1.5)
    v3 := make(map[string]int)
    v3["hello"] = 42
    for fi89 := 0; fi89 < 10; fi89++ {
        v1 = v1
        v90 := GenS10{F0: 2, F1: -212, F2: int64(v1)}
        for fi91 := 0; fi91 < 11; fi91++ {
            switch "hello" {
            case "hello":
                if ((((825 == genFn0()) || false) || ((fi89 == -2147483648) && ((-2.25 / v2) > v2))) || ((0 <= (fi91 + 128)) || true)) {
                    switch fi89 {
                    case -413:
                        v92 := GenS14{F0: 1}
                        _ = v92
                    }
                    if (518 != 453) {
                        fmt.Println(fi89, fi89, fi91)
                    } else {
                        var v93 float32 = v2
                        _ = v93
                    }
                    if false {
                        if ((1 != 810) || (100 != 0)) {
                            fmt.Println(fi91, fi89, fi89)
                            fmt.Println(fi89)
                            v1 = (v1 >> ((65535) & 63))
                        }
                    }
                } else {
                    v2 = 0.0
                    _ = len(v3)
                    v94 := v3[("\n\t" + "")]
                    _ = v94
                }
            case "\n\t":
                switch ("hello" + ("foo bar" + (("\n\t" + (("hello" + (((("hello" + ("a" + "\n\t")) + "") + "foo bar") + "foo bar")) + "")) + (("a" + "") + "foo bar")))) {
                case "hello":
                    if (1 > 65535) {
                        if (1.5 >= 3.14159) {
                            v95 := []int{fi89}
                            var v96 int32 = 0
                            v95 = append(v95, fi89)
                            _ = v95
                            _ = v96
                        }
                    } else {
                        v2 += -2.25
                        switch 210 {
                        case 32768:
                            if ("" < (((("a" + "foo bar") + ((("\n\t" + (("a" + "\n\t") + ("" + ("hello" + "\n\t")))) + (("" + "hello") + ((("hello" + ("" + "")) + "a") + "hello"))) + "foo bar")) + "\n\t") + "hello")) {
                                var v97 float32 = 0.1
                                v1 = v1
                                _ = v97
                            }
                        case 133:
                            v2 += v2
                        case -259:
                            var v98 uint = 4294967295
                            _ = v98
                        default:
                            if false {
                                fmt.Println(v1, fi91, v1)
                                v99 := []int{fi89, fi89, (fi91 & fi91), (^(fi91))}
                                var v100 float32 = (3.14159 + (v2 - (v2 + v2)))
                                _ = v100
                                _ = v99
                            }
                        }
                        var v101 float32 = v2
                        _ = v101
                    }
                case "a":
                    v2 -= (((0.1 * v2) - -2.25) - float32(0.0))
                case "\n\t":
                    fmt.Println(v1, v2)
                default:
                    if (1 != 1) {
                        v102 := GenS10{F0: 2, F1: 1, F2: -1}
                        var v103 int = 1
                        v104 := GenS5{F0: false, F1: 2, F2: 1}
                        _ = v102
                        _ = v103
                        _ = v104
                    }
                }
            }
            var v105 uint16 = 32767
            _ = fi91
            _ = v105
        }
        _ = fi89
        _ = v90
    }
    if (1 >= 1) {
        for fi106 := 0; fi106 < 12; fi106++ {
            v107 := GenS15{F0: 0, F1: 1, F2: 255}
            _ = fi106
            _ = v107
        }
        switch "\n\t" {
        case "":
            for fi108 := 0; fi108 < 5; fi108++ {
                for fi109 := 0; fi109 < 12; fi109++ {
                    if (fi108 > (fi109 >> ((42) & 63))) {
                        v110 := GenS16{F0: 100, F1: 0, F2: true}
                        var v111 float32 = 3.14159
                        var v112 uint8 = (v1 >> ((279) & 63))
                        var v113 int32 = genFn1(-32768, -1, 0)
                        v113 += v113
                        _ = v110
                        _ = v111
                        _ = v112
                        _ = v113
                    }
                    v114 := GenS9{F0: 401, F1: 809, F2: (v1 - 65)}
                    v2 -= v2
                    _ = fi109
                    _ = v114
                }
                var v115 int16 = -32768
                var v116 float32 = v2
                _ = fi108
                _ = v115
                _ = v116
            }
        case "hello":
            var v117 uint8 = 2
            _ = v117
        }
    } else {
        var v118 int64 = 9223372036854775807
        _ = v118
    }
    v3[("hello" + "\n\t")] = 398
    switch "a" {
    case "hello":
        v1 = (v1 << ((65535) & 63))
    case "\n\t":
        fmt.Println(v2, v2)
    default:
        if (-37 != -779) {
            switch "a" {
            case "hello":
                if (("a" + ("" + "foo bar")) != "a") {
                    var v119 bool = (true && (true || ((v1 & (v1 & 127)) < v1)))
                    for fi120 := 0; fi120 < 8; fi120++ {
                        if (436 != 0) {
                            var v121 float64 = float64(v2)
                            v122 := GenS1{F0: 42}
                            var v123 uint64 = 0
                            _ = v121
                            _ = v122
                            _ = v123
                        } else {
                            v124 := GenS1{F0: 42}
                            _ = v124
                        }
                        v2 -= v2
                        _ = fi120
                    }
                    if v119 {
                        v125 := GenS3{F0: 65535, F1: 839}
                        for fi126 := 0; fi126 < 3; fi126++ {
                            v119 = (995 != 934)
                            var v127 int8 = 127
                            _ = fi126
                            _ = v127
                        }
                        v2 = 3.14159
                        _ = v125
                    }
                    _ = v119
                }
            }
            var v128 float64 = (((1e10 / (-2.25 * 1.5)) / 1.0) / float64(3.14159))
            switch -2147483648 {
            case 32768:
                if true {
                    v129 := GenS5{F0: false, F1: 4294967295, F2: 32767}
                    _ = v129
                } else {
                    for fi130 := 0; fi130 < 15; fi130++ {
                        fmt.Println(v1)
                        _ = fi130
                    }
                    fmt.Println(v1, v128, v1)
                }
            default:
                for fi131 := 0; fi131 < 1; fi131++ {
                    if (239 > (v1 + 66)) {
                        for fi132 := 0; fi132 < 9; fi132++ {
                            v128 = 0.1
                            switch "a" {
                            case "hello":
                                v2 = (v2 / (v2 - (float32(1e10) + float32((((v128 + (v128 / (v128 / v128))) / (-2.25 - 3.14159)) - -2.25)))))
                            default:
                                v133 := GenS17{F0: fi132, F1: -581, F2: 0}
                                _ = v133
                            }
                            _ = fi132
                        }
                        if false {
                            v2 = -2.25
                        }
                    }
                    v134 := GenS18{F0: 419, F1: 2147483647, F2: 127}
                    var v135 int = fi131
                    _ = fi131
                    _ = v134
                    _ = v135
                }
            }
            _ = v128
        } else {
            v2 -= float32(0.1)
        }
    }
    if (42 == -31) {
        v136 := GenS19{F0: ((("hello" + "\n\t") + "\n\t") + "a"), F1: 0, F2: 91}
        _ = v136
    } else {
        for fi137 := 0; fi137 < 10; fi137++ {
            for fi138 := 0; fi138 < 11; fi138++ {
                _ = len(v3)
                _ = fi138
            }
            v2 = 1e10
            _ = fi137
        }
        v2 = v2
        v1 += (v1 % (((v1 * (v1 - v1))) | 1))
    }
    switch 42 {
    case -9223372036854775808:
        v2 -= v2
    default:
        for fi139 := 0; fi139 < 7; fi139++ {
            switch fi139 {
            case 338:
                var v140 uint8 = (^(v1))
                _ = v140
            }
            v1 += v1
            v2 = (v2 + v2)
            _ = fi139
        }
    }
    v3["hello"] = 0
    v141 := GenS20{F0: -1, F1: 1}
    for fi142 := 0; fi142 < 17; fi142++ {
        var v143 int32 = 2147483647
        fmt.Println(v143, v143)
        var v144 uint8 = v1
        _ = fi142
        _ = v143
        _ = v144
    }
    for fi145 := 0; fi145 < 2; fi145++ {
        var v146 int32 = -1
        _ = fi145
        _ = v146
    }
    v147 := GenS18{F0: 1, F1: 487, F2: 1}
    v148 := []int{42, 32768}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v148:", len(v148), v148)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v141
    _ = v147
    _ = v148
    _ = v2
    _ = v3
}
