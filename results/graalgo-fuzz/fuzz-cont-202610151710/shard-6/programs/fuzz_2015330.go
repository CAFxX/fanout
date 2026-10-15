package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 int16; F1 uint; F2 uint32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint32; F1 int16; F2 int }
type GenS9 struct { F0 uint; F1 uint8 }
type GenS10 struct { F0 uint16 }
func genFn0(p0 uint32, p1 uint) int64 {
    v4 := []int{9223372036854775807}
    switch (("" + ("a" + ("foo bar" + ("a" + ("hello" + ((("hello" + "hello") + ("foo bar" + "\n\t")) + ("foo bar" + "foo bar"))))))) + "foo bar") {
    case "foo bar":
        v4 = append(v4, 127)
    case "a":
        var v5 int16 = 859
        _ = v5
    case "hello":
        v6 := GenS1{F0: p0, F1: (p1 % (((p1 << ((p1) & 63))) | 1))}
        _ = v6
    }
    var v7 float32 = 1.5
    var v8 bool = false
    fmt.Println(v7, p0, v8)
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    _ = v8
    return 896
}
func genFn1(p0 uint64, p1 uint8) uint64 {
    fmt.Println(p0)
    var v9 uint64 = 981
    v10 := []int{128, 495, 4294967295, -445}
    p1 = (p1 << ((18446744073709551615) & 63))
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return v9
}

func main() {
    var v1 int64 = 487
    v2 := []uint32{}
    v3 := make(map[string]uint32)
    if false {
        var v11 int64 = (v1 * v1)
        if (-2.25 > 3.14159) {
            var v12 int32 = 1
            _ = v12
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v1 += (v1 % ((v1) | 1))
        _ = v11
    } else {
        v1 -= v1
        var v13 int16 = -1
        _ = v13
    }
    switch 488 {
    case 127:
        if (4294967295 <= genFn1(925, 238)) {
            v1 -= -1
            v1 += -1
        } else {
            if false {
                if len(v2) > 0 {
                    v2[0] = 195
                }
                v1 += (v1 / ((v1) | 1))
                v1 = genFn0(141, 175)
            }
            if (-2147483648 < 850) {
                v14 := GenS1{F0: 42, F1: 1}
                v15 := GenS2{F0: false}
                var v16 uint32 = 2147483647
                _ = v14
                _ = v15
                _ = v16
            }
        }
    case 0:
        if (32767 <= 833) {
            fmt.Println(v1)
        } else {
            var v17 string = ((((("hello" + "\n\t") + ("a" + (("" + "foo bar") + "a"))) + ("a" + "\n\t")) + "") + ("foo bar" + ("hello" + ((((("" + "a") + ("" + "hello")) + (("a" + "hello") + "hello")) + "") + "foo bar"))))
            _ = v17
        }
    case 4294967295:
        v18 := GenS2{F0: true}
        _ = v18
    }
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    switch 540 {
    case 310:
        for fi19 := 0; fi19 < 3; fi19++ {
            v2 = append(v2, 271)
            _ = fi19
        }
    case 128:
        if (false || (false && (((false && true) || ((((!(!((932 != 2) && true))) && (-2.25 != 0.0)) || false) && (18446744073709551615 < 1))) && (!false)))) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v1 = (v1 + -943)
        } else {
            for fi20 := 0; fi20 < 13; fi20++ {
                if (((0.0 * ((0.0 / 0.1) / (float32((1.5 / 1.0)) * float32(1.5)))) >= (-2.25 / 1.5)) || (fi20 > 983)) {
                    v1 += v1
                    switch ("hello" + "hello") {
                    case "\n\t":
                        fmt.Println(v1, fi20)
                    }
                }
                for fi21 := 0; fi21 < 18; fi21++ {
                    if ((-508 <= 0) && false) {
                        if len(v2) > 0 {
                            v2[0] = 65535
                        }
                    } else {
                        switch "hello" {
                        case "":
                            var v22 float32 = (float32((-2.25 + 1e10)) * ((float32(-2.25) - 1.5) - -2.25))
                            _ = v22
                        case "\n\t":
                            if (65535 < 2147483647) {
                                var v23 int = fi20
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v23
                            }
                        case "hello":
                            var v24 bool = (((-2.25 + 0.0) <= (1.5 / 1.5)) || true)
                            _ = v24
                        }
                        v1 -= genFn0(122, 992)
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v25 int = (fi21 % ((408) | 1))
                    _ = fi21
                    _ = v25
                }
                _ = fi20
            }
            fmt.Println(v1)
            if len(v2) > 0 {
                v2[0] = 763
            }
        }
    default:
        v26 := v3[("hello" + "a")]
        _ = v26
    }
    var v27 int = 128
    v28 := GenS2{F0: true}
    v2 = append(v2, 525)
    if false {
        if true {
            if true {
                var v29 int8 = -128
                var v30 int64 = (v1 ^ v1)
                _ = v29
                _ = v30
            } else {
                var v31 uint8 = 1
                var v32 string = "foo bar"
                _ = v31
                _ = v32
            }
            var v33 int64 = v1
            v34 := GenS1{F0: 328, F1: 469}
            _ = v33
            _ = v34
        } else {
            v27 -= v27
        }
    } else {
        var v35 string = "foo bar"
        _ = v35
    }
    var v36 float32 = 1e10
    if true {
        for fi37 := 0; fi37 < 8; fi37++ {
            if (float64(-2.25) < 1e10) {
                if (360 != 2147483647) {
                    v1 -= (v1 + v1)
                    fmt.Println(v36, v36)
                    if true {
                        for fi38 := 0; fi38 < 7; fi38++ {
                            switch v27 {
                            case 42:
                                var v39 int = 0
                                _ = v39
                            case -362:
                                v40 := v3["a"]
                                _ = v40
                            default:
                                var v41 uint = 255
                                _ = v41
                            }
                            switch "" {
                            case "a":
                                fmt.Println(v1, v27)
                            case "foo bar":
                                switch -816 {
                                case 32767:
                                    v42 := GenS2{F0: false}
                                    _ = v42
                                default:
                                    v36 = -2.25
                                }
                            }
                            _ = fi38
                        }
                        v43 := GenS1{F0: 2147483647, F1: 876}
                        for fi44 := 0; fi44 < 1; fi44++ {
                            _ = len(v3)
                            v45 := GenS3{F0: 128, F1: 255, F2: 42}
                            v27 = fi37
                            _ = fi44
                            _ = v45
                        }
                        _ = v43
                    }
                } else {
                    v46 := GenS1{F0: 614, F1: 255}
                    _ = v46
                }
            }
            _ = fi37
        }
        if (65535 <= 42) {
            v47 := GenS1{F0: 2147483648, F1: 0}
            fmt.Println(v36, v1, v27)
            var v48 float64 = ((3.14159 - 0.0) + 1.5)
            _ = v47
            _ = v48
        } else {
            var v49 uint8 = 0
            if ((0 >= genFn1(42, 144)) && false) {
                switch v27 {
                case 0:
                    delete(v3, "hello")
                case 4294967295:
                    var v50 uint32 = 2
                    _ = v50
                case -900:
                    fmt.Println(v36)
                default:
                    _ = len(v3)
                }
                var v51 int32 = -2147483648
                var v52 int = (v27 | v27)
                _ = v51
                _ = v52
            }
            _ = v49
        }
        fmt.Println(v27)
    }
    v27 -= v27
    if false {
        for fi53 := 0; fi53 < 15; fi53++ {
            var v54 uint64 = 18446744073709551615
            delete(v3, "\n\t")
            v55 := v3[""]
            _ = fi53
            _ = v54
            _ = v55
        }
        var v56 uint32 = 0
        _ = v56
    } else {
        if (195 <= 115) {
            if (!(!true)) {
                for fi57 := 0; fi57 < 19; fi57++ {
                    var v58 uint16 = 530
                    if true {
                        var v59 uint = 482
                        v1 -= v1
                        _ = v59
                    }
                    switch 42 {
                    case 128:
                        if false {
                            _ = len(v2)
                            fmt.Println(v27, v1, v27)
                        } else {
                            v1 += (v1 << ((1) & 63))
                            switch "" {
                            case "foo bar":
                                var v60 uint = 42
                                _ = v60
                            case "":
                                for fi61 := 0; fi61 < 16; fi61++ {
                                    fmt.Println(v58)
                                    switch 42 {
                                    case -1:
                                        fmt.Println(v58)
                                    case 42:
                                        v36 = float32(((1e10 / (3.14159 + ((1e10 - float64(0.0)) + (0.0 - 1.5)))) - 0.0))
                                    case 943:
                                        var v62 uint32 = 1
                                        _ = v62
                                    default:
                                        if ("" != "hello") {
                                            v63 := GenS2{F0: (!true)}
                                            switch ("a" + ("foo bar" + ((("\n\t" + (("\n\t" + "a") + "foo bar")) + (("foo bar" + "a") + "hello")) + "hello"))) {
                                            case "a":
                                                v58 += v58
                                            }
                                            if (((0.1 - 1e10) + (1e10 / 1e10)) <= 3.14159) {
                                                if false {
                                                    v64 := GenS4{F0: -1}
                                                    if len(v2) > 0 {
                                                        v2[0] = 0
                                                    }
                                                    _ = v64
                                                } else {
                                                    var v65 float64 = ((0.0 / 1e10) + (-2.25 / 1.0))
                                                    switch (fi57 % ((v27) | 1)) {
                                                    case 862:
                                                        var v66 bool = false
                                                        _ = v66
                                                    case -9223372036854775808:
                                                        var v67 int = -1
                                                        _ = v67
                                                    default:
                                                        v68 := v3["a"]
                                                        _ = v68
                                                    }
                                                    _ = v65
                                                }
                                                v69 := v3["foo bar"]
                                                _ = v69
                                            }
                                            switch "a" {
                                            case "hello":
                                                v70 := GenS5{F0: -734, F1: 0}
                                                _ = v70
                                            case "\n\t":
                                                switch 4294967295 {
                                                case -1:
                                                    switch ("\n\t" + (("" + "foo bar") + "\n\t")) {
                                                    case "\n\t":
                                                        var v71 int32 = 1
                                                        _ = v71
                                                    case "hello":
                                                        v72 := GenS1{F0: 2147483647, F1: 222}
                                                        _ = v72
                                                    default:
                                                        fmt.Println(v1, v58, fi61)
                                                    }
                                                case 1:
                                                    switch "foo bar" {
                                                    case "":
                                                        var v73 int32 = -465
                                                        _ = v73
                                                    case "hello":
                                                        v27 += (fi61 | 128)
                                                    }
                                                }
                                            }
                                            _ = v63
                                        } else {
                                            var v74 bool = (!false)
                                            if (-2.25 == 0.1) {
                                                v74 = ((!((!(-225 >= 2147483647)) || (!(418 != (^(v27)))))) || false)
                                                var v75 uint64 = 0
                                                v74 = ((v1 << ((0) & 63)) >= v1)
                                                _ = v75
                                            } else {
                                                fmt.Println(fi57)
                                            }
                                            _ = v74
                                        }
                                    }
                                    _ = fi61
                                }
                            }
                            var v76 uint64 = 18446744073709551615
                            _ = v76
                        }
                    case 32767:
                        for fi77 := 0; fi77 < 15; fi77++ {
                            if false {
                                var v78 int = -1
                                v79 := GenS6{F0: -1, F1: 127, F2: 264}
                                fmt.Println(fi77)
                                _ = v78
                                _ = v79
                            } else {
                                if true {
                                    var v80 int = fi57
                                    _ = v80
                                } else {
                                    v36 += (v36 - ((float32(0.1) - -2.25) * (v36 / ((((v36 + (1.5 / ((float32(1.5) - v36) - float32(((1e10 / -2.25) * (float64(1e10) * -2.25)))))) / float32(float64(v36))) - v36) * ((v36 / -2.25) * v36)))))
                                    var v81 int8 = -120
                                    _ = v81
                                }
                                var v82 int16 = 32767
                                _ = v82
                            }
                            v83 := GenS5{F0: -140, F1: 255}
                            v1 += -9223372036854775808
                            _ = fi77
                            _ = v83
                        }
                    default:
                        v36 = v36
                    }
                    _ = fi57
                    _ = v58
                }
                v3["hello"] = 990
            } else {
                var v84 int = 32767
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v27, v36, v36)
                case "":
                    switch v27 {
                    case 9223372036854775807:
                        v85 := GenS7{F0: (true || (!("a" < "hello")))}
                        _ = v85
                    case -97:
                        var v86 int16 = -333
                        _ = v86
                    case 127:
                        fmt.Println(v36)
                    default:
                        for fi87 := 0; fi87 < 19; fi87++ {
                            fmt.Println(v27)
                            v36 += 1.5
                            _ = fi87
                        }
                    }
                case "a":
                    fmt.Println(v36, v84)
                }
                _ = v84
            }
            v3[("foo bar" + "foo bar")] = 42
            v36 += 0.0
        } else {
            switch "" {
            case "foo bar":
                if ((^(v27)) <= (v27 << ((2) & 63))) {
                    var v88 bool = (!(!(42 == 42)))
                    _ = v88
                } else {
                    switch 9223372036854775807 {
                    case 32768:
                        v36 = 1.5
                    case -815:
                        v89 := GenS8{F0: 1, F1: 0, F2: (^(v27))}
                        _ = v89
                    case 9223372036854775807:
                        switch "a" {
                        case "foo bar":
                            if false {
                                for fi90 := 0; fi90 < 9; fi90++ {
                                    v91 := GenS4{F0: 1000}
                                    var v92 int32 = 228
                                    _ = fi90
                                    _ = v91
                                    _ = v92
                                }
                                if (!true) {
                                    v1 = genFn0(2, 0)
                                }
                            }
                        default:
                            if false {
                                v1 = v1
                                switch "" {
                                case "hello":
                                    v93 := GenS9{F0: 18446744073709551615, F1: 2}
                                    _ = v93
                                case "foo bar":
                                    for fi94 := 0; fi94 < 16; fi94++ {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        for fi95 := 0; fi95 < 18; fi95++ {
                                            v96 := GenS6{F0: -1, F1: 0, F2: 2}
                                            _ = fi95
                                            _ = v96
                                        }
                                        var v97 uint16 = 42
                                        _ = fi94
                                        _ = v97
                                    }
                                }
                            } else {
                                v27 -= v27
                            }
                        }
                    }
                }
            case "":
                for fi98 := 0; fi98 < 9; fi98++ {
                    var v99 int64 = (v1 | (v1 % (((v1 << ((42) & 63))) | 1)))
                    v100 := GenS4{F0: -218}
                    v3["a"] = 639
                    _ = fi98
                    _ = v100
                    _ = v99
                }
            case "a":
                for fi101 := 0; fi101 < 18; fi101++ {
                    fmt.Println(fi101)
                    v36 = v36
                    _ = fi101
                }
            }
        }
        if (!(-2147483648 == 519)) {
            switch (-(v27)) {
            case 790:
                v102 := GenS3{F0: 0, F1: 42, F2: -1}
                _ = v102
            default:
                if false {
                    v2 = append(v2, 2147483647)
                    switch (("\n\t" + ("\n\t" + "foo bar")) + ("foo bar" + ("hello" + "hello"))) {
                    case "a":
                        v103 := GenS3{F0: 0, F1: 0, F2: 42}
                        _ = v103
                    case "hello":
                        var v104 int16 = -32768
                        _ = v104
                    default:
                        v36 += v36
                    }
                    v3[("a" + "\n\t")] = 362
                }
            }
            v36 = v36
            v27 = (-(v27))
        } else {
            switch ("foo bar" + "foo bar") {
            case "":
                var v105 bool = (true || (true || true))
                _ = v105
            case "a":
                v1 = v1
            case "hello":
                v106 := GenS9{F0: 255, F1: 127}
                _ = v106
            default:
                for fi107 := 0; fi107 < 16; fi107++ {
                    if (!false) {
                        if (65535 < 2) {
                            if true {
                                var v108 uint64 = 4294967295
                                fmt.Println(fi107, fi107, v1)
                                for fi109 := 0; fi109 < 4; fi109++ {
                                    var v110 int32 = -151
                                    v111 := GenS8{F0: 0, F1: 32767, F2: (v27 + -104)}
                                    v112 := v3[("" + "")]
                                    switch ("foo bar" + "foo bar") {
                                    case "":
                                        var v113 int64 = v1
                                        _ = v113
                                    case "hello":
                                        v114 := GenS10{F0: 255}
                                        _ = v114
                                    }
                                    _ = len(v3)
                                    _ = fi109
                                    _ = v110
                                    _ = v111
                                    _ = v112
                                }
                                _ = v108
                            } else {
                                if false {
                                    for fi115 := 0; fi115 < 20; fi115++ {
                                        if true {
                                            var v116 int64 = -867
                                            _ = v116
                                        } else {
                                            switch ("hello" + (("foo bar" + "") + (("foo bar" + ((("hello" + "hello") + "foo bar") + "")) + ""))) {
                                            case "foo bar":
                                                var v117 uint16 = 123
                                                _ = v117
                                            case "a":
                                                v27 = 0
                                            }
                                            v27 += 128
                                        }
                                        var v118 bool = false
                                        _ = len(v2)
                                        var v119 int32 = -2147483648
                                        _ = fi115
                                        _ = v118
                                        _ = v119
                                    }
                                    if (v36 > 3.14159) {
                                        var v120 float32 = 3.14159
                                        v36 = ((1e10 - v120) / v120)
                                        _ = v120
                                    } else {
                                        switch ("a" + "a") {
                                        case "\n\t":
                                            fmt.Println(fi107, v36, v36)
                                        default:
                                            var v121 uint8 = 42
                                            _ = v121
                                        }
                                    }
                                } else {
                                    if false {
                                        fmt.Println(fi107)
                                        var v122 int16 = -1
                                        v123 := GenS6{F0: -589, F1: 712, F2: 633}
                                        _ = v122
                                        _ = v123
                                    } else {
                                        if (0 >= 65535) {
                                            for fi124 := 0; fi124 < 20; fi124++ {
                                                v27 = (v27 - fi107)
                                                if true {
                                                    v125 := GenS7{F0: false}
                                                    fmt.Println(v27)
                                                    _ = v125
                                                } else {
                                                    v126 := GenS9{F0: 300, F1: 128}
                                                    var v127 uint16 = 758
                                                    v36 -= 1.5
                                                    _ = v126
                                                    _ = v127
                                                }
                                                v1 -= (v1 ^ v1)
                                                _ = fi124
                                            }
                                        }
                                        v36 -= v36
                                        v3[""] = 1
                                    }
                                    v1 = genFn0(0, 354)
                                    for fi128 := 0; fi128 < 4; fi128++ {
                                        switch "foo bar" {
                                        case "hello":
                                            switch ("\n\t" + (("a" + "hello") + "\n\t")) {
                                            case "":
                                                v1 += -110
                                            case "foo bar":
                                                switch (fi128 / ((v27) | 1)) {
                                                case 2147483647:
                                                    v36 = 1e10
                                                default:
                                                    var v129 uint64 = genFn1(genFn1(18446744073709551615, 102), 255)
                                                    _ = v129
                                                }
                                            case "hello":
                                                v27 = 811
                                            default:
                                                v27 = fi107
                                            }
                                        case "a":
                                            if (v27 > (fi107 << ((42) & 63))) {
                                                switch "" {
                                                case "foo bar":
                                                    var v130 bool = false
                                                    _ = len(v2)
                                                    _ = v130
                                                case "\n\t":
                                                    var v131 int = int(v1)
                                                    _ = v131
                                                case "":
                                                    _ = len(v2)
                                                }
                                            }
                                        case "foo bar":
                                            var v132 uint = 2
                                            _ = v132
                                        default:
                                            fmt.Println(v1, v27, fi107)
                                        }
                                        _ = fi128
                                    }
                                }
                            }
                            v133 := GenS2{F0: true}
                            v36 = -2.25
                            _ = v133
                        }
                        v1 = -9223372036854775808
                    }
                    if ((v1 == v1) && (2 == 0)) {
                        if true {
                            for fi134 := 0; fi134 < 6; fi134++ {
                                var v135 int8 = 86
                                var v136 uint8 = 128
                                _ = fi134
                                _ = v135
                                _ = v136
                            }
                            if len(v2) > 0 {
                                v2[0] = 4294967295
                            }
                        } else {
                            switch -693 {
                            case 42:
                                var v137 int = fi107
                                _ = v137
                            case -432:
                                v138 := GenS9{F0: 255, F1: 0}
                                _ = v138
                            default:
                                for fi139 := 0; fi139 < 12; fi139++ {
                                    if (654 >= 32767) {
                                        v36 = v36
                                        v140 := GenS10{F0: 2}
                                        v36 = 1.5
                                        _ = v140
                                    } else {
                                        var v141 int32 = 1
                                        if ((42 >= 258) && false) {
                                            if false {
                                                v36 = 1.5
                                            }
                                            if (-70 >= 0) {
                                                v1 = (v1 << ((4294967295) & 63))
                                                fmt.Println(fi107, v1, v36)
                                                fmt.Println(v1, v27)
                                            } else {
                                                var v142 uint32 = 2
                                                fmt.Println(v141, v36, v36)
                                                _ = v142
                                            }
                                            v141 = 42
                                        }
                                        _ = v141
                                    }
                                    _ = fi139
                                }
                            }
                            for fi143 := 0; fi143 < 9; fi143++ {
                                var v144 int8 = -46
                                _ = fi143
                                _ = v144
                            }
                        }
                    } else {
                        v145 := GenS9{F0: 246, F1: 2}
                        _ = v145
                    }
                    _ = fi107
                }
            }
        }
        for fi146 := 0; fi146 < 4; fi146++ {
            v147 := GenS7{F0: (-13 > -540)}
            fmt.Println(v1)
            v148 := GenS3{F0: 113, F1: 255, F2: 100}
            _ = fi146
            _ = v147
            _ = v148
        }
    }
    v1 += (v1 ^ genFn0(868, 0))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v36
}
