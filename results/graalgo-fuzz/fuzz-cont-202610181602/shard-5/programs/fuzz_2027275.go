package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 int8; F1 bool; F2 int8 }
type GenS5 struct { F0 uint32; F1 bool; F2 int64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint; F1 int8; F2 string }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS10 struct { F0 uint8; F1 int8; F2 int64 }
type GenS11 struct { F0 bool; F1 int8 }
type GenS12 struct { F0 int; F1 bool }
type GenS13 struct { F0 int; F1 int64 }
type GenS14 struct { F0 int16; F1 int64; F2 int }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 uint32; F1 string }
type GenS18 struct { F0 uint32; F1 int16; F2 string }
type GenS19 struct { F0 int; F1 uint8 }
type GenS20 struct { F0 uint8 }
type GenS21 struct { F0 string; F1 bool; F2 int64 }
type GenS22 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS23 struct { F0 uint8; F1 int8 }
type GenS24 struct { F0 uint32; F1 int16 }
type GenS25 struct { F0 int32 }
type GenS26 struct { F0 int16; F1 uint8 }
type GenS27 struct { F0 int16 }
func genFn0(p0 int64, p1 int) string {
    p0 = -799
    if false {
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 int16 = 32767
            v8 := GenS1{F0: 65535}
            p0 -= (p0 + (p0 | 42))
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return "a"
}
func genFn1(p0 uint32, p1 bool, p2 int) string {
    var v9 bool = (307 >= 42)
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return "a"
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    v3 := []int{}
    var v4 bool = false
    v5 := []uint32{4294967295, 470, 4294967295, 332}
    if true {
        v4 = (!v4)
    }
    var v10 float32 = 1e10
    var v11 uint = 4294967295
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS2{F0: 2, F1: 0}
        fmt.Println(v11, fi12, v11)
        _ = fi12
        _ = v13
    }
    if v4 {
        var v14 int8 = -128
        _ = v14
    } else {
        fmt.Println(v11, v10)
        if v4 {
            var v15 int8 = -128
            _ = v15
        }
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 float32 = (v10 / v10)
            _ = fi16
            _ = v17
        }
    }
    if (!(-982 == -392)) {
        switch ("" + "") {
        case "foo bar":
            v10 = v10
        case "hello":
            if v4 {
                v11 += v11
                v18 := GenS3{F0: v4, F1: 4294967295}
                v10 += v10
                _ = v18
            } else {
                var v19 float32 = (v10 / v10)
                fmt.Println(v11, v10, v10)
                if (32767 != 32767) {
                    var v20 uint64 = 2
                    v4 = false
                    if (!v4) {
                        v21 := GenS4{F0: 1, F1: false, F2: 127}
                        if (-1 <= 1) {
                            var v22 int32 = 2147483647
                            v19 = 1e10
                            _ = v22
                        }
                        _ = v21
                    }
                    _ = v20
                }
                _ = v19
            }
        }
        var v23 uint32 = 759
        var v24 int8 = -67
        _ = v23
        _ = v24
    } else {
        v25 := GenS5{F0: 1, F1: (false || v4), F2: -9223372036854775808}
        _ = v25
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        if v4 {
            if true {
                v10 = v10
                v27 := GenS4{F0: -1, F1: ((fi26 / (((fi26 % ((fi26) | 1))) | 1)) < (fi26 | 178)), F2: 100}
                var v28 uint8 = 128
                _ = v27
                _ = v28
            } else {
                var v29 int32 = -1
                if (32767 >= 2) {
                    for fi30 := 0; fi30 < 2; fi30++ {
                        var v31 int = fi30
                        v10 -= (((v10 - v10) * float32(float64(float32(0.0)))) - v10)
                        v3 = append(v3, (v31 + fi30))
                        _ = fi30
                        _ = v31
                    }
                    v29 += 1
                    for fi32 := 0; fi32 < 9; fi32++ {
                        v5 = append(v5, 2)
                        if true {
                            v11 = v11
                            if false {
                                fmt.Println(fi26, v10, v29)
                                v33 := GenS3{F0: (!false), F1: -2147483648}
                                var v34 uint = 65535
                                var v35 int32 = 1
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                            v11 += 65535
                        } else {
                            var v36 uint64 = 18446744073709551615
                            var v37 bool = (!(v4 && (fi32 < (fi32 / (((fi32 / ((-668) | 1))) | 1)))))
                            if v4 {
                                v4 = (v4 && ((v36 < (v36 << (((^(v11))) & 63))) || ((-1 <= 1) && v37)))
                                _ = len(v1)
                                var v38 int = int(v36)
                                _ = v38
                            } else {
                                var v39 uint = 908
                                _ = v39
                            }
                            _ = v36
                            _ = v37
                        }
                        _ = fi32
                    }
                } else {
                    if false {
                        for fi40 := 0; fi40 < 1; fi40++ {
                            switch "" {
                            case "":
                                var v41 float64 = 1e10
                                _ = v41
                            case "a":
                                v42 := GenS6{F0: "foo bar"}
                                _ = v42
                            }
                            fmt.Println(v4, v4)
                            _ = fi40
                        }
                    } else {
                        for fi43 := 0; fi43 < 11; fi43++ {
                            var v44 int64 = int64(v29)
                            if (v4 && false) {
                                var v45 uint16 = 1
                                _ = v45
                            } else {
                                var v46 int = 127
                                if len(v3) > 0 {
                                    v3[0] = v46
                                }
                                v47 := GenS3{F0: (v4 && (!((0 != 1) && (v4 && ((241 >= 1000) && (65535 <= 154)))))), F1: fi43}
                                fmt.Println(v10)
                                _ = v46
                                _ = v47
                            }
                            switch "" {
                            case "":
                                fmt.Println(v4, v11)
                            case "\n\t":
                                v48 := GenS7{F0: v11, F1: 1, F2: "hello"}
                                _ = v48
                            case "hello":
                                var v49 uint = v11
                                _ = v49
                            }
                            _ = fi43
                            _ = v44
                        }
                        var v50 float64 = 1e10
                        for fi51 := 0; fi51 < 4; fi51++ {
                            var v52 uint8 = 196
                            v53 := GenS8{F0: (v52 >> ((v11) & 63))}
                            _ = fi51
                            _ = v52
                            _ = v53
                        }
                        _ = v50
                    }
                    v54 := GenS9{F0: v11, F1: 0, F2: 885}
                    _ = v54
                }
                switch (fi26 >> ((v11) & 63)) {
                case 1:
                    var v55 float64 = ((3.14159 / (0.0 + 3.14159)) + -2.25)
                    _ = v55
                case 32767:
                    v29 = (v29 - v29)
                default:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = v29
            }
            v10 -= (0.0 / v10)
        } else {
            v11 = v11
            var v56 int64 = 0
            v1[0] = 784
            _ = v56
        }
        switch ((("a" + ("" + ((("" + "") + "") + "foo bar"))) + "") + "\n\t") {
        case "a":
            switch "hello" {
            case "foo bar":
                switch ("hello" + "a") {
                case "a":
                    switch "hello" {
                    case "":
                        v10 = v10
                    default:
                        fmt.Println(v4, fi26)
                    }
                default:
                    switch fi26 {
                    case -527:
                        if v4 {
                            v10 += (v10 * v10)
                            _ = len(v3)
                        } else {
                            v11 = (^(v11))
                            v10 = (0.0 * 3.14159)
                        }
                    case 339:
                        if v4 {
                            var v57 int32 = 1
                            if (379 > 0) {
                                fmt.Println(v11)
                                if false {
                                    v11 -= 18446744073709551615
                                    if len(v3) > 0 {
                                        v3[0] = fi26
                                    }
                                    v4 = v4
                                }
                            } else {
                                if v4 {
                                    var v58 bool = v4
                                    var v59 int16 = 0
                                    var v60 int8 = -28
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                }
                                var v61 int16 = -1
                                _ = v61
                            }
                            _ = v57
                        } else {
                            var v62 int32 = 1
                            _ = v62
                        }
                    }
                }
            case "\n\t":
                for fi63 := 0; fi63 < 14; fi63++ {
                    v4 = false
                    if ((("a" + "hello") + "") > "") {
                        fmt.Println(v10)
                        var v64 uint64 = 981
                        v64 -= 4294967295
                        _ = v64
                    }
                    var v65 uint64 = 2
                    fmt.Println(v4)
                    var v66 uint = 544
                    _ = fi63
                    _ = v65
                    _ = v66
                }
            case "hello":
                fmt.Println(v11, v11)
            }
        case "\n\t":
            switch ("hello" + "hello") {
            case "\n\t":
                v11 = v11
            case "":
                if (122 == 86) {
                    var v67 uint = v11
                    var v68 string = "foo bar"
                    _ = v67
                    _ = v68
                } else {
                    if (42 >= 32767) {
                        var v69 int32 = int32(fi26)
                        if (v4 || v4) {
                            var v70 uint32 = 2147483647
                            v4 = v4
                            switch "hello" {
                            case "foo bar":
                                switch "" {
                                case "":
                                    switch ("foo bar" + "\n\t") {
                                    case "a":
                                        v70 = v70
                                    case "hello":
                                        switch (fi26 >> ((65535) & 63)) {
                                        case -459:
                                            var v71 int = -324
                                            _ = v71
                                        case 1:
                                            switch "\n\t" {
                                            case "a":
                                                fmt.Println(v69)
                                            case "foo bar":
                                                switch (("hello" + ("foo bar" + "a")) + ("a" + ("a" + ""))) {
                                                case "hello":
                                                    for fi72 := 0; fi72 < 17; fi72++ {
                                                        fmt.Println(fi26)
                                                        fmt.Println(fi72, v69)
                                                        v73 := GenS10{F0: 127, F1: -1, F2: 9223372036854775807}
                                                        _ = fi72
                                                        _ = v73
                                                    }
                                                case "":
                                                    var v74 int8 = -126
                                                    _ = v74
                                                default:
                                                    var v75 uint64 = 237
                                                    _ = v75
                                                }
                                            case "\n\t":
                                                v4 = false
                                            default:
                                                v69 -= -1
                                            }
                                        }
                                    }
                                case "\n\t":
                                    if (float64(3.14159) <= 0.0) {
                                        v4 = v4
                                    }
                                default:
                                    v76 := GenS5{F0: 42, F1: true, F2: 42}
                                    _ = v76
                                }
                            case "hello":
                                fmt.Println(v10, v70)
                            case "\n\t":
                                for fi77 := 0; fi77 < 2; fi77++ {
                                    if len(v3) > 0 {
                                        v3[0] = 128
                                    }
                                    v3 = append(v3, fi77)
                                    _ = fi77
                                }
                            }
                            _ = v70
                        } else {
                            for fi78 := 0; fi78 < 17; fi78++ {
                                if (311 <= -1) {
                                    if false {
                                        v10 -= v10
                                    } else {
                                        fmt.Println(fi78)
                                        switch int(v11) {
                                        case 2147483647:
                                            v79 := v1[2]
                                            v80 := GenS4{F0: 100, F1: true, F2: -128}
                                            _ = len(v5)
                                            _ = v79
                                            _ = v80
                                        case 691:
                                            if (-1 == 998) {
                                                var v81 uint32 = 65535
                                                v10 = 3.14159
                                                if v4 {
                                                    v4 = v4
                                                    fmt.Println(fi78)
                                                }
                                                _ = v81
                                            } else {
                                                var v82 uint32 = 1
                                                v83 := GenS11{F0: v4, F1: 100}
                                                if (-59 >= 0) {
                                                    if v4 {
                                                        switch (fi78 / (((fi78 >> (((v11 | 65535)) & 63))) | 1)) {
                                                        case 631:
                                                            v84 := GenS3{F0: (2 == uint64(v82)), F1: fi78}
                                                            _ = v84
                                                        case 899:
                                                            v82 = 42
                                                        case 32768:
                                                            v10 = (((v10 * (0.1 + ((v10 / (1.5 / float32(-2.25))) + (v10 * v10)))) / 1.5) * v10)
                                                        }
                                                        switch (fi26 ^ fi26) {
                                                        case 42:
                                                            if v4 {
                                                                v69 += v69
                                                            } else {
                                                                v85 := GenS12{F0: int(v11), F1: v4}
                                                                _ = v85
                                                            }
                                                        case 127:
                                                            var v86 int8 = 127
                                                            _ = v86
                                                        case 0:
                                                            fmt.Println(v4, v69)
                                                        }
                                                        fmt.Println(v82, v82, fi26)
                                                    }
                                                }
                                                _ = v82
                                                _ = v83
                                            }
                                        }
                                    }
                                    var v87 int8 = -75
                                    _ = v87
                                }
                                var v88 int8 = 42
                                _ = fi78
                                _ = v88
                            }
                            for fi89 := 0; fi89 < 10; fi89++ {
                                fmt.Println(v11, fi26, v10)
                                _ = fi89
                            }
                            _ = len(v3)
                        }
                        _ = v69
                    }
                    if (0.0 < (3.14159 * 1e10)) {
                        if v4 {
                            if (!(((uint16(v11) == 0) || (!v4)) || v4)) {
                                switch (fi26 / ((fi26) | 1)) {
                                case 2147483647:
                                    var v90 float64 = ((-2.25 / -2.25) * (1e10 - 1e10))
                                    _ = v90
                                case 1:
                                    var v91 float32 = 0.0
                                    _ = v91
                                case 0:
                                    if v4 {
                                        if len(v5) > 0 {
                                            v5[0] = 4294967295
                                        }
                                        var v92 int16 = -1
                                        _ = len(v5)
                                        _ = v92
                                    } else {
                                        if (!(true || false)) {
                                            fmt.Println(v4)
                                            v3 = append(v3, (fi26 - 623))
                                        } else {
                                            var v93 int32 = 0
                                            var v94 uint64 = 327
                                            switch (fi26 * (fi26 & 32767)) {
                                            case -833:
                                                switch ("" + "\n\t") {
                                                case "\n\t":
                                                    for fi95 := 0; fi95 < 9; fi95++ {
                                                        var v96 uint64 = (v94 / ((v94) | 1))
                                                        v97 := GenS13{F0: (fi26 - fi26), F1: -9223372036854775808}
                                                        _ = fi95
                                                        _ = v96
                                                        _ = v97
                                                    }
                                                case "hello":
                                                    if (!((v94 >> ((v11) & 63)) < (v94 + v94))) {
                                                        switch (fi26 & 2147483647) {
                                                        case 2147483647:
                                                            v98 := GenS13{F0: -703, F1: 0}
                                                            _ = v98
                                                        case -1:
                                                            for fi99 := 0; fi99 < 3; fi99++ {
                                                                if (v4 || v4) {
                                                                    var v100 float64 = 0.1
                                                                    var v101 int = int(v11)
                                                                    var v102 int64 = 644
                                                                    _ = v100
                                                                    _ = v101
                                                                    _ = v102
                                                                }
                                                                v103 := v1[18]
                                                                var v104 uint8 = 224
                                                                var v105 uint32 = 2147483647
                                                                _ = fi99
                                                                _ = v103
                                                                _ = v104
                                                                _ = v105
                                                            }
                                                        case -90:
                                                            fmt.Println(v11)
                                                        default:
                                                            fmt.Println(v93)
                                                        }
                                                        v106 := GenS14{F0: -301, F1: 9223372036854775807, F2: fi26}
                                                        _ = v106
                                                    }
                                                default:
                                                    var v107 uint8 = 0
                                                    _ = v107
                                                }
                                            }
                                            _ = v93
                                            _ = v94
                                        }
                                    }
                                }
                                switch -2147483648 {
                                case 42:
                                    _ = len(v1)
                                }
                                for fi108 := 0; fi108 < 3; fi108++ {
                                    if v4 {
                                        v11 -= v11
                                    }
                                    _ = len(v3)
                                    var v109 float32 = float32(0.0)
                                    _ = fi108
                                    _ = v109
                                }
                            } else {
                                for fi110 := 0; fi110 < 7; fi110++ {
                                    var v111 uint64 = 4294967295
                                    if v4 {
                                        if false {
                                            v4 = false
                                        } else {
                                            fmt.Println(fi110, v10, v111)
                                            v2[2147483648] = (1 == 42)
                                        }
                                    } else {
                                        switch "foo bar" {
                                        case "\n\t":
                                            _ = len(v1)
                                        }
                                    }
                                    if (-136 < 1) {
                                        var v112 uint64 = v111
                                        var v113 string = "hello"
                                        _ = v112
                                        _ = v113
                                    } else {
                                        switch 2147483647 {
                                        case 32768:
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        case -655:
                                            v114 := v1[1]
                                            _ = v114
                                        }
                                        if (65535 == (^(v11))) {
                                            if (889 == 2147483648) {
                                                switch ("\n\t" + "a") {
                                                case "a":
                                                    if true {
                                                        switch (fi110 << (((v11 % ((v11) | 1))) & 63)) {
                                                        case -638:
                                                            v111 += v111
                                                        case 9223372036854775807:
                                                            if ((-128 > -1) || (v4 || (42 == 878))) {
                                                                v115 := GenS8{F0: 127}
                                                                _ = v115
                                                            } else {
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    var v116 int8 = 127
                                                                    _ = v116
                                                                case "\n\t":
                                                                    switch ((("\n\t" + "") + "a") + ("a" + "")) {
                                                                    case "foo bar":
                                                                        fmt.Println(v11, v4)
                                                                    case "hello":
                                                                        switch int(v11) {
                                                                        case -734:
                                                                            v4 = v4
                                                                            v10 = (((v10 / 1e10) + (((1e10 / float32((0.1 + (-2.25 / -2.25)))) * v10) - (v10 + 0.1))) - ((v10 + v10) - v10))
                                                                            v117 := GenS10{F0: 1, F1: -49, F2: 0}
                                                                            v4 = ((!(v4 && false)) && (("\n\t" + "a") < "foo bar"))
                                                                            v10 -= 0.1
                                                                            _ = v117
                                                                        case 0:
                                                                            if v4 {
                                                                                var v118 uint64 = (v111 >> ((205) & 63))
                                                                                switch (-(fi110)) {
                                                                                case -122:
                                                                                    var v119 bool = v4
                                                                                    _ = v119
                                                                                case 127:
                                                                                    fmt.Println(v11, v111)
                                                                                    switch fi26 {
                                                                                    case 32767:
                                                                                        v120 := GenS9{F0: v11, F1: 42, F2: 604}
                                                                                        _ = v120
                                                                                    case 0:
                                                                                        v121 := GenS7{F0: (v11 - v11), F1: 100, F2: ("hello" + "")}
                                                                                        _ = v121
                                                                                    case 42:
                                                                                        v122 := GenS13{F0: (^(fi110)), F1: 0}
                                                                                        _ = v122
                                                                                    }
                                                                                    v123 := GenS15{F0: 0}
                                                                                    _ = v123
                                                                                case 0:
                                                                                    var v124 int16 = 0
                                                                                    _ = v124
                                                                                default:
                                                                                    fmt.Println(v10, fi110)
                                                                                }
                                                                                _ = v118
                                                                            }
                                                                        }
                                                                    case "a":
                                                                        v125 := GenS12{F0: fi110, F1: v4}
                                                                        _ = v125
                                                                    default:
                                                                        _ = len(v2)
                                                                        var v126 float64 = float64(v10)
                                                                        v127 := GenS9{F0: (v11 >> ((v11) & 63)), F1: 226, F2: 266}
                                                                        _ = v126
                                                                        _ = v127
                                                                    }
                                                                default:
                                                                    _ = len(v3)
                                                                }
                                                                var v128 string = "foo bar"
                                                                v4 = v4
                                                                _ = v128
                                                            }
                                                        case 32767:
                                                            v129 := GenS16{F0: 65535}
                                                            _ = v129
                                                        }
                                                        switch "foo bar" {
                                                        case "a":
                                                            switch -2147483648 {
                                                            case 4294967295:
                                                                var v130 uint = 224
                                                                _ = v130
                                                            default:
                                                                switch (-(fi26)) {
                                                                case 9223372036854775807:
                                                                    v131 := GenS13{F0: -1, F1: -237}
                                                                    _ = v131
                                                                }
                                                            }
                                                        default:
                                                            v111 = 4294967295
                                                        }
                                                    } else {
                                                        var v132 uint = 567
                                                        _ = v132
                                                    }
                                                case "foo bar":
                                                    v133 := GenS12{F0: (fi26 << (((^(v11))) & 63)), F1: false}
                                                    _ = v133
                                                case "hello":
                                                    fmt.Println(v4, v10)
                                                default:
                                                    v134 := GenS17{F0: uint32(v111), F1: "a"}
                                                    _ = v134
                                                }
                                                switch fi26 {
                                                case 2147483647:
                                                    fmt.Println(v10, v4)
                                                }
                                            } else {
                                                _ = len(v5)
                                                if v4 {
                                                    v135 := v1[2]
                                                    _ = v135
                                                }
                                            }
                                        } else {
                                            var v136 uint32 = 1
                                            var v137 uint8 = uint8(v111)
                                            _ = v136
                                            _ = v137
                                        }
                                        v138 := GenS18{F0: 242, F1: -798, F2: "\n\t"}
                                        if ((^(v11)) == 2) {
                                            fmt.Println(v11, fi110, fi110)
                                            var v139 bool = (!(-2.25 >= ((0.0 / 1.5) / ((float64((((1.5 + (v10 - v10)) / (v10 / (float32((-2.25 / -2.25)) * v10))) - float32((1e10 + -2.25)))) * (0.0 + 1e10)) * (1e10 - -2.25)))))
                                            _ = v139
                                        } else {
                                            fmt.Println(v111, fi26)
                                            if (2 == 143) {
                                                v4 = (2 > 65535)
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                v111 -= v111
                                            } else {
                                                if (v4 && v4) {
                                                    if (!true) {
                                                        v11 = v11
                                                        v140 := GenS16{F0: 255}
                                                        if false {
                                                            v111 = (v111 % ((v111) | 1))
                                                            switch "\n\t" {
                                                            case "a":
                                                                var v141 int64 = 42
                                                                v142 := GenS19{F0: (fi26 % ((fi26) | 1)), F1: 42}
                                                                var v143 uint64 = (^(v111))
                                                                _ = v141
                                                                _ = v142
                                                                _ = v143
                                                            }
                                                            switch "foo bar" {
                                                            case "":
                                                                v144 := GenS20{F0: 43}
                                                                _ = v144
                                                            case "\n\t":
                                                                fmt.Println(fi26, v4, v10)
                                                            case "hello":
                                                                var v145 uint32 = 431
                                                                _ = v145
                                                            default:
                                                                fmt.Println(fi26)
                                                            }
                                                        }
                                                        _ = v140
                                                    } else {
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            v146 := GenS21{F0: "hello", F1: true, F2: 9223372036854775807}
                                                            _ = v146
                                                        case "hello":
                                                            fmt.Println(v11)
                                                            fmt.Println(v4, v10, fi110)
                                                        case "a":
                                                            v10 = 1.5
                                                        }
                                                        _ = len(v3)
                                                        if ((v111 >> ((v11) & 63)) >= v111) {
                                                            var v147 uint = uint(v111)
                                                            if (false || (985 < (^(v111)))) {
                                                                if (2 == 255) {
                                                                    var v148 int64 = -9223372036854775808
                                                                    v3 = append(v3, 128)
                                                                    _ = v148
                                                                }
                                                                v4 = ((-2.25 == -2.25) && v4)
                                                                v149 := GenS12{F0: (fi110 * fi26), F1: false}
                                                                var v150 uint = (v147 + (v147 << ((v11) & 63)))
                                                                _ = v149
                                                                _ = v150
                                                            }
                                                            _ = v147
                                                        }
                                                    }
                                                    switch (fi26 & 42) {
                                                    case -538:
                                                        _ = len(v3)
                                                    case 32767:
                                                        v151 := GenS22{F0: -263, F1: (v111 | v111), F2: 2}
                                                        _ = v151
                                                    }
                                                    var v152 int32 = -2147483648
                                                    _ = v152
                                                } else {
                                                    if v4 {
                                                        var v153 bool = (true && v4)
                                                        switch fi110 {
                                                        case -769:
                                                            fmt.Println(v4, v10, v11)
                                                        case 1:
                                                            v11 = (v11 & (^(v11)))
                                                        default:
                                                            var v154 int32 = -2147483648
                                                            _ = v154
                                                        }
                                                        _ = v153
                                                    } else {
                                                        var v155 string = ""
                                                        _ = v155
                                                    }
                                                    if (((v10 * float32((1e10 - -2.25))) / ((((v10 - v10) * v10) - (float32(1e10) - v10)) / float32(1.5))) <= v10) {
                                                        v5 = append(v5, 831)
                                                        v11 = 255
                                                    }
                                                }
                                                v11 -= 18446744073709551615
                                                if false {
                                                    v10 = ((v10 + 0.0) / v10)
                                                    switch ("a" + (("foo bar" + "a") + "\n\t")) {
                                                    case "\n\t":
                                                        _ = len(v3)
                                                    case "a":
                                                        v10 -= -2.25
                                                    case "foo bar":
                                                        v111 = (v111 ^ v111)
                                                    }
                                                    if (255 >= 127) {
                                                        v156 := GenS23{F0: 0, F1: 53}
                                                        _ = v156
                                                    } else {
                                                        var v157 uint64 = (v111 + v111)
                                                        v157 += (v157 / (((v157 - (v157 + (v157 / ((2) | 1))))) | 1))
                                                        _ = v157
                                                    }
                                                    var v158 uint8 = 42
                                                    v2[4294967295] = (!(-2.25 < 1.5))
                                                    fmt.Println(v158, v4, v158)
                                                    var v159 uint64 = v111
                                                    if (v4 && (!false)) {
                                                        v5 = append(v5, 65535)
                                                        v10 = v10
                                                        if v4 {
                                                            fmt.Println(v11, v10)
                                                            fmt.Println(v10, v4, v111)
                                                        }
                                                    }
                                                    var v160 int16 = 1000
                                                    _ = v158
                                                    _ = v159
                                                    _ = v160
                                                }
                                            }
                                            v4 = (0 >= -1)
                                        }
                                        var v161 uint64 = (^(v111))
                                        v162 := GenS24{F0: 4294967295, F1: 802}
                                        var v163 float64 = (-2.25 / (-2.25 - 1e10))
                                        fmt.Println(fi26, v4)
                                        _ = v138
                                        _ = v161
                                        _ = v162
                                        _ = v163
                                    }
                                    _ = fi110
                                    _ = v111
                                }
                                if v4 {
                                    v4 = v4
                                    switch fi26 {
                                    case 0:
                                        var v164 uint = v11
                                        _ = v164
                                    }
                                } else {
                                    v11 = 42
                                }
                                v3 = append(v3, (fi26 | fi26))
                            }
                            for fi165 := 0; fi165 < 1; fi165++ {
                                var v166 uint = v11
                                _ = fi165
                                _ = v166
                            }
                        } else {
                            var v167 uint16 = 2
                            v168 := GenS25{F0: -2147483648}
                            var v169 int32 = -2147483648
                            _ = v167
                            _ = v168
                            _ = v169
                        }
                        fmt.Println(v11)
                        v170 := GenS26{F0: -1, F1: 0}
                        _ = v170
                    }
                    fmt.Println(fi26, v11, v4)
                }
            case "hello":
                if (1 == 128) {
                    var v171 uint64 = 1
                    v172 := GenS21{F0: (("hello" + "foo bar") + "foo bar"), F1: v4, F2: 208}
                    var v173 int8 = 1
                    _ = v171
                    _ = v172
                    _ = v173
                } else {
                    for fi174 := 0; fi174 < 2; fi174++ {
                        v175 := GenS7{F0: 0, F1: 60, F2: "hello"}
                        _ = fi174
                        _ = v175
                    }
                }
            default:
                if (-528 <= -1) {
                    fmt.Println(v4, v10, v4)
                }
            }
        default:
            switch 1 {
            case 4294967295:
                v10 -= v10
            case -182:
                if v4 {
                    v10 -= (1.5 * (v10 - (((-2.25 + v10) / v10) * v10)))
                    v176 := GenS26{F0: 42, F1: 42}
                    _ = v176
                }
            }
        }
        for fi177 := 0; fi177 < 2; fi177++ {
            fmt.Println(v10)
            switch fi26 {
            case 128:
                v4 = v4
            }
            if v4 {
                var v178 int16 = -32768
                _ = v178
            }
            _ = fi177
        }
        _ = fi26
    }
    switch "a" {
    case "":
        fmt.Println(v11, v10, v11)
    case "a":
        for fi179 := 0; fi179 < 4; fi179++ {
            v180 := GenS13{F0: 0, F1: -7}
            v4 = v4
            v181 := GenS27{F0: -709}
            _ = fi179
            _ = v180
            _ = v181
        }
    case "hello":
        for fi182 := 0; fi182 < 2; fi182++ {
            for fi183 := 0; fi183 < 16; fi183++ {
                v4 = (0 <= 42)
                v184 := GenS22{F0: 1, F1: 222, F2: 66}
                _ = fi183
                _ = v184
            }
            if v4 {
                if len(v5) > 0 {
                    v5[0] = 105
                }
                var v185 int16 = 42
                _ = v185
            }
            _ = fi182
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
