package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint; F2 int }
type GenS3 struct { F0 int32; F1 bool; F2 uint }
type GenS4 struct { F0 uint8; F1 string; F2 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 int32; F1 uint32; F2 uint32 }
type GenS8 struct { F0 uint; F1 int64 }
type GenS9 struct { F0 uint; F1 uint; F2 int32 }
type GenS10 struct { F0 uint32; F1 string; F2 int32 }
type GenS11 struct { F0 uint; F1 bool }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint64 }
func genFn0(p0 int8, p1 int16, p2 bool) int32 {
    p2 = ((-(p1)) <= p1)
    var v7 int16 = (p1 | 649)
    var v8 int16 = v7
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1(p0 uint) int8 {
    v9 := GenS1{F0: 42}
    v10 := []int{128, 365, 2147483647}
    switch 32768 {
    case 312:
        if len(v10) > 0 {
            v10[0] = -2147483648
        }
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    if false {
        switch 929 {
        case 0:
            p0 = (^(p0))
        case -9223372036854775808:
            if (125 >= 73) {
                var v11 int8 = 0
                v11 = (v11 ^ (v11 >> ((19) & 63)))
                _ = v11
            }
        default:
            for fi12 := 0; fi12 < 11; fi12++ {
                if len(v10) > 0 {
                    v10[0] = fi12
                }
                _ = fi12
            }
        }
        if len(v10) > 0 {
            v10[0] = 9223372036854775807
        }
        if false {
            var v13 int16 = -259
            var v14 uint = (^(p0))
            var v15 uint16 = 42
            _ = v13
            _ = v14
            _ = v15
        }
    } else {
        var v16 int32 = -277
        v10 = append(v10, 2147483647)
        _ = v16
    }
    _ = p0
    _ = v10
    _ = v9
    return -128
}

func main() {
    var v1 bool = true
    var v2 uint16 = 1
    var v3 uint8 = 255
    v4 := []string{}
    v5 := make(map[int]int)
    var v6 float32 = (3.14159 + 3.14159)
    var v17 int = 4294967295
    switch ("\n\t" + "\n\t") {
    case "":
        if ((float32((((-2.25 + 0.1) * (-2.25 + 1.5)) + (0.0 / 3.14159))) != v6) && v1) {
            switch "a" {
            case "hello":
                v18 := GenS2{F0: -32768, F1: 65535, F2: -1}
                _ = v18
            case "\n\t":
                if (v1 || v1) {
                    switch (v17 ^ (v17 / ((143) | 1))) {
                    case 0:
                        for fi19 := 0; fi19 < 12; fi19++ {
                            switch fi19 {
                            case -822:
                                var v20 uint64 = 1
                                _ = v20
                            case -1:
                                fmt.Println(v17, v6, v3)
                            case 0:
                                var v21 uint64 = 0
                                _ = v21
                            }
                            _ = fi19
                        }
                    case 32767:
                        v6 = (v6 - -2.25)
                    case -142:
                        switch v17 {
                        case 127:
                            if false {
                                var v22 bool = (0.1 > 0.1)
                                v23 := GenS1{F0: 895}
                                var v24 int = v17
                                _ = v22
                                _ = v23
                                _ = v24
                            } else {
                                if v1 {
                                    if true {
                                        v25 := GenS2{F0: -382, F1: 196, F2: v17}
                                        var v26 int64 = -9223372036854775808
                                        _ = v25
                                        _ = v26
                                    }
                                    if (v6 == 0.1) {
                                        if len(v4) > 0 {
                                            v4[0] = "a"
                                        }
                                    }
                                } else {
                                    if (!v1) {
                                        v27 := GenS3{F0: -273, F1: (v1 || (32767 <= -1)), F2: 656}
                                        if (-1 > -2147483648) {
                                            v1 = (42 <= 561)
                                        } else {
                                            v28 := GenS4{F0: (v3 & 1), F1: ("a" + ""), F2: 100}
                                            _ = v28
                                        }
                                        v29 := GenS3{F0: 0, F1: v1, F2: 18446744073709551615}
                                        _ = v27
                                        _ = v29
                                    }
                                }
                                v17 = (v17 - v17)
                                if false {
                                    for fi30 := 0; fi30 < 12; fi30++ {
                                        v17 = (fi30 << ((42) & 63))
                                        v31 := GenS2{F0: 0, F1: 255, F2: 128}
                                        switch "a" {
                                        case "a":
                                            for fi32 := 0; fi32 < 11; fi32++ {
                                                if ((v1 || ((v1 || v1) || v1)) || (!true)) {
                                                    var v33 int16 = -325
                                                    fmt.Println(v6, v2, v3)
                                                    if v1 {
                                                        if ((42 < 254) && (v1 || v1)) {
                                                            v34 := GenS5{F0: 0}
                                                            v6 += (v6 - (-2.25 / 1e10))
                                                            _ = v34
                                                        } else {
                                                            v35 := GenS6{F0: (v33 << ((2) & 63)), F1: 166}
                                                            _ = v35
                                                        }
                                                    }
                                                    _ = v33
                                                }
                                                if len(v4) > 0 {
                                                    v4[0] = ""
                                                }
                                                fmt.Println(v6, v17, v17)
                                                _ = fi32
                                            }
                                        default:
                                            var v36 int16 = 32767
                                            _ = v36
                                        }
                                        _ = fi30
                                        _ = v31
                                    }
                                }
                            }
                        case 4294967295:
                            if (true || v1) {
                                switch v17 {
                                case 9223372036854775807:
                                    v37 := GenS7{F0: -1, F1: 257, F2: 2}
                                    _ = v37
                                }
                                var v38 int8 = -91
                                v3 = 1
                                _ = v38
                            } else {
                                _ = len(v5)
                                _ = len(v4)
                            }
                        case 0:
                            if true {
                                var v39 uint8 = (v3 / (((^(v3))) | 1))
                                var v40 int64 = 9223372036854775807
                                _ = v39
                                _ = v40
                            }
                        }
                    default:
                        for fi41 := 0; fi41 < 14; fi41++ {
                            for fi42 := 0; fi42 < 8; fi42++ {
                                v3 -= (v3 % ((v3) | 1))
                                _ = fi42
                            }
                            _ = fi41
                        }
                    }
                    switch 2147483647 {
                    case 2147483647:
                        if true {
                            if v1 {
                                var v43 int64 = -9223372036854775808
                                _ = v43
                            }
                            _ = len(v4)
                        }
                    }
                }
            case "a":
                var v44 float64 = (3.14159 + 0.0)
                _ = v44
            default:
                if v1 {
                    v5[(-(v17))] = v17
                    if false {
                        if (387 < 2) {
                            for fi45 := 0; fi45 < 7; fi45++ {
                                var v46 float64 = (float64(v6) - ((1e10 * 0.1) + 3.14159))
                                _ = fi45
                                _ = v46
                            }
                            _ = len(v4)
                            var v47 int64 = -9223372036854775808
                            _ = v47
                        } else {
                            v48 := GenS1{F0: 49}
                            _ = v48
                        }
                    }
                } else {
                    for fi49 := 0; fi49 < 18; fi49++ {
                        for fi50 := 0; fi50 < 20; fi50++ {
                            v51 := GenS2{F0: 1000, F1: 65535, F2: v17}
                            _ = fi50
                            _ = v51
                        }
                        switch ((("a" + "\n\t") + ("" + "hello")) + "a") {
                        case "hello":
                            for fi52 := 0; fi52 < 12; fi52++ {
                                v5[(^(fi52))] = v17
                                if (2147483647 >= 487) {
                                    switch fi49 {
                                    case 127:
                                        var v53 bool = true
                                        _ = v53
                                    case 128:
                                        var v54 uint16 = v2
                                        _ = v54
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = "hello"
                                    }
                                } else {
                                    v55 := GenS7{F0: genFn0(3, -908, true), F1: uint32(v2), F2: 2}
                                    fmt.Println(v1, v1, fi52)
                                    switch -416 {
                                    case 32768:
                                        fmt.Println(v3)
                                    }
                                    v6 += float32(0.0)
                                    v3 += (v3 * v3)
                                    v56 := GenS7{F0: 2147483647, F1: 542, F2: 42}
                                    _ = v55
                                    _ = v56
                                }
                                _ = fi52
                            }
                        case "a":
                            for fi57 := 0; fi57 < 11; fi57++ {
                                switch ((("foo bar" + "\n\t") + "foo bar") + ("\n\t" + "foo bar")) {
                                case "\n\t":
                                    v58 := GenS8{F0: 65535, F1: 195}
                                    _ = v58
                                case "a":
                                    fmt.Println(v3)
                                case "":
                                    var v59 int = (-(fi49))
                                    _ = v59
                                }
                                var v60 int64 = -275
                                _ = fi57
                                _ = v60
                            }
                        }
                        _ = fi49
                    }
                    if (!(260 <= 255)) {
                        if (!v1) {
                            var v61 uint64 = 836
                            v62 := GenS1{F0: -32768}
                            _ = v61
                            _ = v62
                        }
                        v1 = (!v1)
                    } else {
                        for fi63 := 0; fi63 < 8; fi63++ {
                            var v64 uint64 = 18446744073709551615
                            if len(v4) > 0 {
                                v4[0] = "foo bar"
                            }
                            _ = fi63
                            _ = v64
                        }
                        v1 = (!true)
                        if (370 < 65535) {
                            var v65 uint32 = 2147483648
                            var v66 uint16 = 255
                            _ = v65
                            _ = v66
                        }
                    }
                    fmt.Println(v6)
                }
            }
        }
    case "foo bar":
        fmt.Println(v1)
    case "hello":
        for fi67 := 0; fi67 < 1; fi67++ {
            v2 += (v2 ^ (v2 & v2))
            fmt.Println(v2)
            _ = len(v5)
            _ = fi67
        }
    }
    v68 := GenS2{F0: -1, F1: 65535, F2: v17}
    v69 := GenS9{F0: 357, F1: 337, F2: -994}
    if (-2147483648 == v17) {
        v17 -= 128
        for fi70 := 0; fi70 < 10; fi70++ {
            v6 = (v6 + (1.5 - v6))
            if ((!true) && (v2 <= (^(v2)))) {
                var v71 int64 = 0
                var v72 int16 = -570
                fmt.Println(v3, v72)
                _ = v71
                _ = v72
            }
            v73 := GenS4{F0: v3, F1: "foo bar", F2: -60}
            _ = fi70
            _ = v73
        }
        if v1 {
            _ = len(v4)
            for fi74 := 0; fi74 < 13; fi74++ {
                for fi75 := 0; fi75 < 12; fi75++ {
                    v17 += 128
                    v5[v17] = v17
                    var v76 uint = 1
                    _ = fi75
                    _ = v76
                }
                for fi77 := 0; fi77 < 7; fi77++ {
                    var v78 uint32 = 933
                    switch 128 {
                    case -9223372036854775808:
                        if (v1 || false) {
                            if len(v4) > 0 {
                                v4[0] = "a"
                            }
                            if (v17 >= (fi74 / (((-(v17))) | 1))) {
                                v79 := GenS10{F0: v78, F1: "foo bar", F2: -2147483648}
                                v2 = (v2 | v2)
                                _ = v79
                            } else {
                                fmt.Println(fi77, v1, v6)
                                var v80 int = (^(fi74))
                                v17 += v80
                                var v81 float64 = 3.14159
                                _ = v80
                                _ = v81
                            }
                        } else {
                            switch ("foo bar" + ("hello" + "hello")) {
                            case "a":
                                fmt.Println(v1, v78, v3)
                            case "\n\t":
                                var v82 string = "\n\t"
                                _ = v82
                            case "foo bar":
                                var v83 string = "hello"
                                _ = v83
                            default:
                                var v84 uint16 = v2
                                _ = v84
                            }
                        }
                    default:
                        var v85 uint64 = 871
                        _ = v85
                    }
                    var v86 uint64 = 2
                    v86 += (v86 / ((v86) | 1))
                    _ = fi77
                    _ = v78
                    _ = v86
                }
                _ = fi74
            }
            for fi87 := 0; fi87 < 10; fi87++ {
                if v1 {
                    delete(v5, 2147483647)
                    var v88 float32 = float32(0.0)
                    _ = v88
                } else {
                    for fi89 := 0; fi89 < 17; fi89++ {
                        v6 -= v6
                        _ = fi89
                    }
                    if v1 {
                        _ = len(v5)
                    }
                }
                for fi90 := 0; fi90 < 16; fi90++ {
                    var v91 uint64 = 2
                    if true {
                        v92 := v5[-1]
                        _ = v92
                    }
                    if (65535 >= 0) {
                        v93 := GenS3{F0: 42, F1: (!((fi87 | fi90) == (v17 << ((920) & 63)))), F2: 255}
                        delete(v5, -1)
                        var v94 uint = 18446744073709551615
                        _ = v93
                        _ = v94
                    } else {
                        if (!v1) {
                            var v95 int64 = -1
                            _ = v95
                        }
                    }
                    _ = fi90
                    _ = v91
                }
                v17 += -518
                _ = fi87
            }
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v96 string = "foo bar"
        _ = v96
    }
    v97 := GenS4{F0: v3, F1: "hello", F2: 127}
    for fi98 := 0; fi98 < 3; fi98++ {
        fmt.Println(v1, v1)
        switch "foo bar" {
        case "\n\t":
            for fi99 := 0; fi99 < 2; fi99++ {
                v17 = (-(fi99))
                var v100 uint16 = v2
                var v101 float64 = (((-2.25 / -2.25) + 1e10) - 3.14159)
                _ = fi99
                _ = v100
                _ = v101
            }
        case "a":
            switch ("\n\t" + ("hello" + "hello")) {
            case "\n\t":
                v3 -= (v3 | v3)
            }
        default:
            fmt.Println(fi98)
        }
        fmt.Println(fi98, fi98, v6)
        _ = fi98
    }
    fmt.Println(v1, v1, v17)
    delete(v5, (v17 >> ((14) & 63)))
    v17 = v17
    fmt.Println(v6, v6)
    var v102 uint16 = v2
    switch (v17 << ((18446744073709551615) & 63)) {
    case 4294967295:
        v102 = v2
    case 362:
        switch v17 {
        case 1:
            var v103 bool = (!v1)
            _ = v103
        case 32767:
            v104 := GenS11{F0: 1, F1: v1}
            _ = v104
        default:
            _ = len(v5)
        }
    case -9223372036854775808:
        v102 += v2
    default:
        for fi105 := 0; fi105 < 1; fi105++ {
            v17 = fi105
            _ = fi105
        }
    }
    var v106 int8 = genFn1(42)
    if true {
        v1 = v1
    } else {
        var v107 string = "a"
        _ = v107
    }
    for fi108 := 0; fi108 < 15; fi108++ {
        var v109 float32 = (float32(1e10) * v6)
        if (!((v3 >= (^(v3))) || v1)) {
            _ = len(v4)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v110 uint = 65535
            _ = v110
        }
        if (int32(fi108) < 1) {
            delete(v5, (v17 << ((546) & 63)))
        } else {
            switch ("a" + "") {
            case "":
                var v111 int64 = -9223372036854775808
                _ = v111
            default:
                if v1 {
                    for fi112 := 0; fi112 < 15; fi112++ {
                        if ((v1 && (true || v1)) || v1) {
                            v1 = (!true)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        if v1 {
                            var v113 uint64 = 1
                            v114 := GenS8{F0: 2, F1: 655}
                            _ = v113
                            _ = v114
                        } else {
                            v17 = (^(fi108))
                            v4 = append(v4, "")
                            if (4294967295 <= 255) {
                                switch "foo bar" {
                                case "\n\t":
                                    if v1 {
                                        switch (-(fi108)) {
                                        case -459:
                                            var v115 string = ("" + ("\n\t" + "\n\t"))
                                            v102 = (v2 << ((0) & 63))
                                            var v116 uint8 = (v3 / (((v3 & (v3 % ((v3) | 1)))) | 1))
                                            _ = v115
                                            _ = v116
                                        case 42:
                                            v17 -= fi112
                                        case 32767:
                                            if v1 {
                                                v5[(fi112 + v17)] = (-(fi108))
                                                v102 = (v102 + v2)
                                                v17 = (fi108 & (fi108 | -63))
                                            } else {
                                                var v117 string = (("" + (("hello" + "\n\t") + "foo bar")) + "\n\t")
                                                switch "a" {
                                                case "foo bar":
                                                    var v118 int32 = 2147483647
                                                    _ = v118
                                                case "\n\t":
                                                    v117 = v117
                                                case "hello":
                                                    v106 = (v106 + -27)
                                                }
                                                var v119 uint = 2
                                                _ = v117
                                                _ = v119
                                            }
                                        default:
                                            if (v1 || (1 == 1)) {
                                                fmt.Println(fi112, v109)
                                                if true {
                                                    switch (v17 & v17) {
                                                    case 506:
                                                        fmt.Println(v109, v3)
                                                    case 32767:
                                                        fmt.Println(v6, fi112, v109)
                                                    }
                                                    switch "a" {
                                                    case "hello":
                                                        var v120 int16 = 32767
                                                        _ = v120
                                                    case "a":
                                                        fmt.Println(v106, v106)
                                                    case "":
                                                        fmt.Println(fi112, v102, v106)
                                                    default:
                                                        v4 = append(v4, "hello")
                                                    }
                                                    v121 := GenS12{F0: v3}
                                                    _ = v121
                                                }
                                                if v1 {
                                                    var v122 int32 = -304
                                                    _ = v122
                                                }
                                            }
                                            delete(v5, (fi112 * (v17 & (-(fi108)))))
                                        }
                                        switch fi108 {
                                        case 595:
                                            var v123 float64 = float64(v6)
                                            _ = v123
                                        case 0:
                                            v4 = append(v4, ("a" + ""))
                                        }
                                        v124 := GenS9{F0: 4294967295, F1: 255, F2: 84}
                                        _ = v124
                                    }
                                case "a":
                                    var v125 uint16 = (v102 % (((v2 >> ((42) & 63))) | 1))
                                    _ = v125
                                case "foo bar":
                                    if (569 != 1) {
                                        v126 := GenS8{F0: 819, F1: -1}
                                        v109 -= (v6 - v109)
                                        _ = v126
                                    } else {
                                        var v127 float64 = -2.25
                                        if ((-(v17)) > 249) {
                                            if true {
                                                if ((v127 < -2.25) && false) {
                                                    var v128 uint64 = 972
                                                    if (genFn0(0, 1, (v1 || true)) != -619) {
                                                        fmt.Println(v6, fi108)
                                                        _ = len(v4)
                                                    } else {
                                                        switch ("foo bar" + (("\n\t" + "\n\t") + ("" + "foo bar"))) {
                                                        case "":
                                                            v3 = (v3 / (((v3 & 42)) | 1))
                                                        }
                                                    }
                                                    v17 -= 128
                                                    var v129 int64 = -9223372036854775808
                                                    _ = v128
                                                    _ = v129
                                                } else {
                                                    switch "foo bar" {
                                                    case "a":
                                                        v127 = v127
                                                    case "":
                                                        if (!((v1 || (v1 && (true && v1))) && ((358 != 2147483648) || v1))) {
                                                            v109 = (1e10 * (1e10 * v109))
                                                            fmt.Println(v2, v17)
                                                            if true {
                                                                var v130 int64 = -9223372036854775808
                                                                _ = v130
                                                            } else {
                                                                v17 -= 32768
                                                                if v1 {
                                                                    v131 := v5[(fi108 % ((v17) | 1))]
                                                                    _ = v131
                                                                } else {
                                                                    v1 = ("" > "hello")
                                                                    v5[fi112] = 42
                                                                    v106 = -49
                                                                }
                                                                switch "a" {
                                                                case "foo bar":
                                                                    v3 = (v3 + 0)
                                                                case "a":
                                                                    v106 = genFn1(0)
                                                                }
                                                            }
                                                            v132 := GenS13{F0: 0}
                                                            v1 = v1
                                                            v102 += (^(v2))
                                                            _ = v132
                                                        }
                                                    case "foo bar":
                                                        if v1 {
                                                            v133 := GenS3{F0: -572, F1: true, F2: 64}
                                                            if true {
                                                                v5[fi112] = (fi112 % ((fi112) | 1))
                                                                var v134 bool = (v1 || ((2147483647 != -1) && (0 <= 42)))
                                                                _ = len(v5)
                                                                _ = v134
                                                            } else {
                                                                _ = len(v4)
                                                            }
                                                            _ = v133
                                                        }
                                                    default:
                                                        switch "a" {
                                                        case "foo bar":
                                                            if (v106 == v106) {
                                                                if len(v4) > 0 {
                                                                    v4[0] = "a"
                                                                }
                                                            }
                                                        }
                                                    }
                                                    var v135 uint16 = (^(v2))
                                                    fmt.Println(fi112, v102)
                                                    _ = v135
                                                }
                                            }
                                            var v136 int16 = 1
                                            _ = v136
                                        } else {
                                            if (1 <= 2) {
                                                var v137 uint8 = v3
                                                var v138 uint16 = v2
                                                _ = v137
                                                _ = v138
                                            } else {
                                                if v1 {
                                                    var v139 int64 = 9223372036854775807
                                                    switch (("foo bar" + ("a" + ("foo bar" + "foo bar"))) + "") {
                                                    case "":
                                                        fmt.Println(v102, v106, v127)
                                                        v3 = (^(v3))
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    case "foo bar":
                                                        var v140 uint64 = 55
                                                        _ = v140
                                                    default:
                                                        if v1 {
                                                            v127 += 3.14159
                                                            fmt.Println(v139, v109, v109)
                                                        }
                                                    }
                                                    _ = v139
                                                }
                                                var v141 float64 = v127
                                                fmt.Println(v102, v6, v102)
                                                _ = v141
                                            }
                                            var v142 float32 = v6
                                            _ = v142
                                        }
                                        v5[(fi112 + (fi108 * v17))] = fi108
                                        _ = v127
                                    }
                                    if (1e10 == 0.1) {
                                        fmt.Println(fi108, v106)
                                    } else {
                                        var v143 uint32 = uint32(v102)
                                        _ = v143
                                    }
                                default:
                                    var v144 bool = true
                                    _ = v144
                                }
                            }
                        }
                        _ = fi112
                    }
                    var v145 uint64 = 679
                    _ = v145
                }
            }
            switch ((((("a" + "\n\t") + "a") + "") + (("\n\t" + "foo bar") + "foo bar")) + (((("\n\t" + "hello") + "a") + "\n\t") + "foo bar")) {
            case "foo bar":
                v3 -= v3
            }
            fmt.Println(v6, v6)
        }
        _ = fi108
        _ = v109
    }
    v106 = v106
    var v146 string = ("hello" + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v106:", v106)
    fmt.Println("v146:", v146)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v102
    _ = v106
    _ = v146
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v69
    _ = v97
}
