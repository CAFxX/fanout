package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 int8; F1 string; F2 uint16 }
type GenS6 struct { F0 int8; F1 int16; F2 int8 }
type GenS7 struct { F0 int64; F1 int32; F2 int64 }
type GenS8 struct { F0 string; F1 int; F2 uint16 }
type GenS9 struct { F0 int32; F1 int }
type GenS10 struct { F0 int64; F1 bool; F2 uint64 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int32; F1 bool; F2 int64 }
type GenS13 struct { F0 uint8; F1 bool; F2 int8 }
type GenS14 struct { F0 int16; F1 int32 }
type GenS15 struct { F0 bool; F1 uint16; F2 uint }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 int8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = -2147483648
    var v3 int = 0
    v4 := []string{"hello", "\n\t", "\n\t", ""}
    var v5 uint = 882
    v6 := make(map[uint32]string)
    v7 := GenS1{F0: 42, F1: 0}
    v2 = (v2 / (((-(v2))) | 1))
    var v8 uint8 = 249
    switch v3 {
    case 2147483647:
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 int8 = 42
            _ = fi9
            _ = v10
        }
    case 128:
        var v11 int16 = -692
        _ = v11
    case 4294967295:
        fmt.Println(v5, v2, v2)
    }
    v12 := GenS2{F0: 0, F1: (v3 + (-(v3)))}
    for fi13 := 0; fi13 < 4; fi13++ {
        for fi14 := 0; fi14 < 6; fi14++ {
            v5 = (v5 << ((42) & 63))
            v8 -= v8
            v15 := GenS2{F0: 127, F1: 2147483647}
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    switch "\n\t" {
    case "hello":
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 int32 = (v2 % ((v2) | 1))
            for fi18 := 0; fi18 < 2; fi18++ {
                var v19 bool = false
                fmt.Println(fi16)
                fmt.Println(v5, v17)
                _ = fi18
                _ = v19
            }
            fmt.Println(v8, v2, v17)
            _ = fi16
            _ = v17
        }
    case "foo bar":
        if (!(0 <= 2147483648)) {
            for fi20 := 0; fi20 < 12; fi20++ {
                var v21 string = ("foo bar" + ("a" + ("\n\t" + "\n\t")))
                _ = fi20
                _ = v21
            }
            v2 -= (v2 >> ((2) & 63))
            fmt.Println(v8)
        } else {
            v22 := v6[0]
            var v23 int16 = 0
            var v24 bool = true
            _ = v22
            _ = v23
            _ = v24
        }
    case "":
        v3 = (^(v3))
    }
    if (65535 == 2147483647) {
        switch (v3 & v3) {
        case 1:
            v8 = (^(v8))
        case 873:
            var v25 float64 = 1e10
            _ = v25
        case -1:
            v26 := GenS3{F0: 32767, F1: 176, F2: 42}
            _ = v26
        default:
            if true {
                v2 -= v2
                if false {
                    var v27 uint = v5
                    _ = v27
                } else {
                    var v28 string = (("" + (("a" + (("\n\t" + "a") + ("" + "foo bar"))) + "a")) + "foo bar")
                    _ = v28
                }
                var v29 int64 = -1
                _ = v29
            }
        }
        v2 += v2
        switch int(v8) {
        case -743:
            for fi30 := 0; fi30 < 7; fi30++ {
                var v31 uint64 = 888
                for fi32 := 0; fi32 < 8; fi32++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi32
                }
                if false {
                    if false {
                        v6[455] = ("foo bar" + ("hello" + "foo bar"))
                        v5 = 338
                        var v33 uint8 = uint8(v5)
                        _ = v33
                    } else {
                        for fi34 := 0; fi34 < 19; fi34++ {
                            v5 += 0
                            v35 := v1[(^(fi34))]
                            v5 += v5
                            _ = fi34
                            _ = v35
                        }
                        if ((true && (("" + ("foo bar" + ("hello" + ("a" + "a")))) > "hello")) || true) {
                            if false {
                                for fi36 := 0; fi36 < 6; fi36++ {
                                    var v37 float32 = 0.0
                                    if ((497 <= (v31 & 0)) && ((fi36 | (fi36 >> ((v5) & 63))) == fi36)) {
                                        v3 = fi36
                                        v38 := GenS4{F0: 1, F1: 545}
                                        v31 -= v31
                                        _ = v38
                                    } else {
                                        v39 := GenS5{F0: 1, F1: "foo bar", F2: 32767}
                                        var v40 int32 = (v2 * (v2 << ((v5) & 63)))
                                        _ = v39
                                        _ = v40
                                    }
                                    var v41 int16 = 1
                                    _ = fi36
                                    _ = v37
                                    _ = v41
                                }
                                var v42 uint = v5
                                _ = v42
                            } else {
                                v31 = (^(v31))
                            }
                        }
                        v2 += int32(v31)
                    }
                    switch 42 {
                    case 9223372036854775807:
                        var v43 float32 = 0.1
                        _ = v43
                    }
                    switch 42 {
                    case -1:
                        v8 += 42
                    case -228:
                        if true {
                            var v44 int64 = -248
                            fmt.Println(v8, v44, fi30)
                            _ = v44
                        }
                    }
                } else {
                    for fi45 := 0; fi45 < 5; fi45++ {
                        fmt.Println(fi30, v5)
                        _ = fi45
                    }
                }
                _ = fi30
                _ = v31
            }
        case -2147483648:
            if false {
                var v46 uint8 = v8
                var v47 uint = v5
                for fi48 := 0; fi48 < 5; fi48++ {
                    fmt.Println(v8, v47)
                    v3 -= (-(fi48))
                    _ = fi48
                }
                _ = v46
                _ = v47
            }
        case 4294967295:
            if (!false) {
                var v49 uint32 = uint32(v8)
                v8 = v8
                _ = v49
            }
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    } else {
        for fi50 := 0; fi50 < 16; fi50++ {
            v3 = (^(fi50))
            var v51 uint16 = 0
            _ = fi50
            _ = v51
        }
        if ("\n\t" > ("a" + "hello")) {
            switch ("" + ("foo bar" + "a")) {
            case "\n\t":
                for fi52 := 0; fi52 < 6; fi52++ {
                    var v53 float32 = 1.5
                    _ = fi52
                    _ = v53
                }
            case "":
                v54 := GenS2{F0: (^(v8)), F1: (v3 ^ (v3 + (v3 | 42)))}
                _ = v54
            }
            delete(v1, v3)
        } else {
            var v55 int = (v3 & v3)
            var v56 bool = false
            v57 := GenS4{F0: 1000, F1: 18446744073709551615}
            _ = v55
            _ = v56
            _ = v57
        }
        v58 := GenS2{F0: v8, F1: (v3 % (((v3 + (v3 | (v3 * v3)))) | 1))}
        _ = v58
    }
    v59 := GenS6{F0: -128, F1: 1, F2: 100}
    if (!(870 != v2)) {
        if false {
            if true {
                v60 := GenS2{F0: v8, F1: int(v8)}
                _ = v60
            }
            v2 = 1
            var v61 float64 = 1.5
            _ = v61
        }
        var v62 uint16 = 208
        _ = v62
    }
    if true {
        v2 = -1
        if (65535 == 65535) {
            var v63 uint32 = 4294967295
            _ = v63
        } else {
            switch "a" {
            case "":
                v64 := GenS1{F0: -1, F1: 32767}
                _ = v64
            default:
                switch (v3 ^ (v3 + (v3 ^ v3))) {
                case 1:
                    v3 = v3
                case 32768:
                    v5 += (v5 * 0)
                case -203:
                    var v65 bool = true
                    _ = v65
                }
            }
        }
    } else {
        for fi66 := 0; fi66 < 1; fi66++ {
            if (-2.25 > 3.14159) {
                switch ("" + (("\n\t" + "a") + "\n\t")) {
                case "a":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "\n\t":
                    for fi67 := 0; fi67 < 11; fi67++ {
                        v68 := GenS2{F0: v8, F1: v3}
                        if (127 > 1) {
                            switch (-(fi66)) {
                            case 2147483647:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case 809:
                                v69 := v1[(v3 << ((983) & 63))]
                                _ = v69
                            }
                        } else {
                            if (((1e10 / 1.0) > (float32(3.14159) * float32((-2.25 + (0.1 / float64(1e10)))))) || (((3.14159 * 1.5) / 3.14159) < 3.14159)) {
                                v70 := GenS7{F0: -912, F1: (v2 * int32(v5)), F2: 511}
                                v5 += (v5 / (((v5 >> ((v5) & 63))) | 1))
                                var v71 uint = 1
                                _ = v70
                                _ = v71
                            }
                        }
                        if true {
                            v3 = (v3 ^ (v3 - (fi66 << ((v5) & 63))))
                            v2 = v2
                            v72 := GenS3{F0: 0, F1: 4294967295, F2: -713}
                            switch ((("foo bar" + "\n\t") + ("hello" + "\n\t")) + "foo bar") {
                            case "a":
                                switch (((("a" + (("a" + "hello") + "foo bar")) + "a") + "") + "") {
                                case "\n\t":
                                    if true {
                                        var v73 float32 = (3.14159 + 1.5)
                                        _ = v73
                                    }
                                    v74 := GenS2{F0: uint8(v5), F1: v3}
                                    v5 += v5
                                    _ = v74
                                default:
                                    fmt.Println(v5, fi66, v3)
                                }
                            case "\n\t":
                                if (!(true && true)) {
                                    switch (v3 | 0) {
                                    case 128:
                                        var v75 string = ("a" + (("hello" + "hello") + "hello"))
                                        _ = v75
                                    case 991:
                                        if true {
                                            var v76 float32 = 3.14159
                                            fmt.Println(fi67)
                                            v77 := GenS1{F0: -1, F1: 1}
                                            _ = v76
                                            _ = v77
                                        }
                                    default:
                                        switch "\n\t" {
                                        case "\n\t":
                                            if len(v4) > 0 {
                                                v4[0] = (("hello" + "hello") + "a")
                                            }
                                        case "":
                                            var v78 int32 = -1
                                            _ = v78
                                        case "a":
                                            fmt.Println(fi66)
                                            if (true || (!(true || (!false)))) {
                                                switch "hello" {
                                                case "\n\t":
                                                    v79 := GenS2{F0: uint8(v5), F1: fi67}
                                                    _ = v79
                                                case "foo bar":
                                                    fmt.Println(v5, v3)
                                                default:
                                                    v80 := v6[42]
                                                    _ = v80
                                                }
                                            } else {
                                                v3 += fi67
                                                var v81 string = ""
                                                v5 -= (v5 << ((v5) & 63))
                                                _ = v81
                                            }
                                            if (0 <= 970) {
                                                var v82 uint64 = 943
                                                if true {
                                                    fmt.Println(v8, fi67)
                                                    v82 = (v82 / ((v82) | 1))
                                                    v83 := GenS8{F0: "foo bar", F1: (v3 ^ 574), F2: 150}
                                                    _ = v83
                                                } else {
                                                    fmt.Println(fi67, fi66, v82)
                                                    fmt.Println(v5, v5, v82)
                                                    v84 := GenS9{F0: (v2 | v2), F1: 0}
                                                    _ = v84
                                                }
                                                _ = v82
                                            }
                                        default:
                                            fmt.Println(fi67)
                                        }
                                    }
                                    v4 = append(v4, "")
                                } else {
                                    switch 32767 {
                                    case 4294967295:
                                        v85 := GenS1{F0: 470, F1: 42}
                                        _ = v85
                                    case 864:
                                        var v86 int64 = -290
                                        _ = v86
                                    }
                                }
                            case "foo bar":
                                if (v5 < (^(v5))) {
                                    if (1 >= 0) {
                                        var v87 uint16 = 120
                                        _ = v87
                                    } else {
                                        fmt.Println(v8, fi67)
                                    }
                                    if true {
                                        var v88 float64 = float64((-2.25 / 1.0))
                                        fmt.Println(fi66, v8)
                                        _ = v88
                                    }
                                }
                            default:
                                v89 := GenS10{F0: 1, F1: true, F2: 4294967295}
                                _ = v89
                            }
                            v8 = (v8 >> ((uint(v8)) & 63))
                            v8 = v8
                            _ = v72
                        } else {
                            delete(v1, 1)
                        }
                        _ = fi67
                        _ = v68
                    }
                case "hello":
                    if len(v4) > 0 {
                        v4[0] = ("hello" + (("foo bar" + "foo bar") + ("\n\t" + (("foo bar" + "foo bar") + "\n\t"))))
                    }
                }
                fmt.Println(v3)
                var v90 float64 = (3.14159 / 3.14159)
                _ = v90
            }
            switch ("foo bar" + "foo bar") {
            case "foo bar":
                var v91 float32 = (1e10 + (3.14159 * 0.1))
                _ = v91
            }
            for fi92 := 0; fi92 < 2; fi92++ {
                var v93 int8 = -128
                _ = fi92
                _ = v93
            }
            _ = fi66
        }
        for fi94 := 0; fi94 < 16; fi94++ {
            var v95 uint8 = 0
            for fi96 := 0; fi96 < 13; fi96++ {
                var v97 uint8 = (v8 * (^(v8)))
                v98 := GenS9{F0: (^(v2)), F1: (fi94 >> ((255) & 63))}
                switch (fi96 >> ((v5) & 63)) {
                case 9223372036854775807:
                    if (638 == 1) {
                        v99 := GenS11{F0: "hello"}
                        v8 -= v97
                        _ = v99
                    } else {
                        v2 = v2
                        switch ("" + "") {
                        case "hello":
                            _ = len(v4)
                        case "foo bar":
                            var v100 uint32 = 2
                            _ = v100
                        default:
                            var v101 int16 = 273
                            _ = v101
                        }
                        if ((0.1 - 0.0) == (0.0 - (1e10 * 1e10))) {
                            v102 := GenS6{F0: 0, F1: -54, F2: 0}
                            _ = v102
                        }
                    }
                case 1:
                    v97 += (^(v95))
                default:
                    _ = len(v6)
                }
                _ = fi96
                _ = v97
                _ = v98
            }
            for fi103 := 0; fi103 < 20; fi103++ {
                v95 -= (v95 | v8)
                _ = fi103
            }
            _ = fi94
            _ = v95
        }
    }
    switch ("" + "\n\t") {
    case "\n\t":
        if len(v4) > 0 {
            v4[0] = ""
        }
    case "hello":
        v104 := GenS10{F0: 625, F1: ((v8 | 42) <= 127), F2: 0}
        _ = v104
    default:
        var v105 string = ("hello" + ("" + ("hello" + "hello")))
        _ = v105
    }
    var v106 bool = (v5 <= (^(v5)))
    v107 := GenS5{F0: 42, F1: ("" + "foo bar"), F2: 831}
    if ((1 < 65535) || (v3 < v3)) {
        v5 += (v5 / ((693) | 1))
    } else {
        if true {
            if ((597 != 1000) && (1e10 <= float32((0.0 / -2.25)))) {
                _ = len(v4)
            } else {
                if ((v106 || (v106 && false)) || (261 > 45)) {
                    v3 -= v3
                    v108 := GenS8{F0: ("" + "foo bar"), F1: (-(v3)), F2: 794}
                    _ = v108
                }
            }
            fmt.Println(v2)
        } else {
            v2 += (-(v2))
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        if (v106 && (18446744073709551615 < 113)) {
            v5 += (v5 + 1)
            switch "\n\t" {
            case "foo bar":
                v109 := GenS5{F0: 42, F1: "foo bar", F2: 344}
                _ = v109
            case "hello":
                v110 := GenS1{F0: 0, F1: 255}
                _ = v110
            case "":
                for fi111 := 0; fi111 < 11; fi111++ {
                    var v112 uint8 = v8
                    var v113 uint8 = (v8 >> (((^(v5))) & 63))
                    _ = fi111
                    _ = v112
                    _ = v113
                }
            }
            v114 := v6[42]
            _ = v114
        }
        switch "" {
        case "\n\t":
            v115 := GenS3{F0: 174, F1: 2147483647, F2: -560}
            _ = v115
        case "a":
            if true {
                var v116 int32 = (-(v2))
                if false {
                    if (-2.25 <= 1e10) {
                        switch "foo bar" {
                        case "":
                            if (!(-759 != 435)) {
                                if false {
                                    v4 = append(v4, "hello")
                                    if (!(int16(v2) != 0)) {
                                        v2 -= (v116 >> ((v5) & 63))
                                        v117 := GenS12{F0: -1, F1: v106, F2: -384}
                                        _ = v117
                                    } else {
                                        if (0.0 == 1e10) {
                                            v8 += (^(v8))
                                        } else {
                                            v3 = v3
                                            for fi118 := 0; fi118 < 10; fi118++ {
                                                var v119 int16 = -32768
                                                var v120 int32 = 1
                                                _ = fi118
                                                _ = v119
                                                _ = v120
                                            }
                                            for fi121 := 0; fi121 < 10; fi121++ {
                                                fmt.Println(fi121, fi121, v116)
                                                _ = fi121
                                            }
                                        }
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                } else {
                                    v3 = v3
                                }
                            } else {
                                v106 = true
                                var v122 uint = 4294967295
                                _ = v122
                            }
                        default:
                            var v123 int64 = 1
                            _ = v123
                        }
                    } else {
                        v5 = v5
                        var v124 uint64 = 2
                        _ = v124
                    }
                    switch (("a" + "") + ("hello" + "hello")) {
                    case "hello":
                        v3 -= v3
                    case "foo bar":
                        switch ((("foo bar" + (("foo bar" + "a") + "a")) + "") + ("" + "hello")) {
                        case "":
                            for fi125 := 0; fi125 < 2; fi125++ {
                                v126 := GenS13{F0: 2, F1: v106, F2: -82}
                                _ = fi125
                                _ = v126
                            }
                        case "\n\t":
                            fmt.Println(v106)
                        case "foo bar":
                            var v127 int32 = (v116 + (v2 & v2))
                            _ = v127
                        default:
                            v5 += v5
                        }
                    case "a":
                        v106 = v106
                    default:
                        var v128 uint64 = 71
                        _ = v128
                    }
                }
                var v129 int8 = -46
                _ = v116
                _ = v129
            } else {
                v130 := GenS14{F0: 0, F1: (v2 ^ v2)}
                fmt.Println(v3)
                for fi131 := 0; fi131 < 19; fi131++ {
                    var v132 bool = ((0 > -659) || (v106 && false))
                    var v133 uint32 = 233
                    _ = fi131
                    _ = v132
                    _ = v133
                }
                _ = v130
            }
        case "foo bar":
            switch ((("\n\t" + "") + "foo bar") + ((((("\n\t" + "foo bar") + "\n\t") + "hello") + ("\n\t" + ("hello" + ("\n\t" + "\n\t")))) + ((("\n\t" + "hello") + (("foo bar" + "a") + (("\n\t" + "a") + ("" + ("hello" + ("hello" + "foo bar")))))) + ("a" + "")))) {
            case "a":
                v134 := GenS15{F0: true, F1: 255, F2: (v5 % ((42) | 1))}
                _ = v134
            case "":
                var v135 int64 = 1
                _ = v135
            case "hello":
                v136 := GenS14{F0: -32768, F1: v2}
                _ = v136
            }
        default:
            v137 := GenS16{F0: -128}
            _ = v137
        }
    }
    fmt.Println(v106)
    v138 := GenS17{F0: 85}
    var v139 int = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v106:", v106)
    fmt.Println("v139:", v139)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v106
    _ = v107
    _ = v12
    _ = v138
    _ = v139
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v7
    _ = v8
}
