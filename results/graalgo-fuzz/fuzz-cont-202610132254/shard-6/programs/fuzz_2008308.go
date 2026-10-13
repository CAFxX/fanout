package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32; F1 int32; F2 bool }
type GenS6 struct { F0 uint8; F1 int8; F2 string }
type GenS7 struct { F0 uint64; F1 bool }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int16; F1 uint; F2 uint }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 bool; F1 string }
type GenS12 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS13 struct { F0 uint16; F1 int }
type GenS14 struct { F0 uint32; F1 uint16 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 int16; F1 uint32; F2 string }
type GenS17 struct { F0 uint32; F1 bool; F2 bool }
type GenS18 struct { F0 string; F1 uint16; F2 uint16 }
type GenS19 struct { F0 int8; F1 uint8; F2 uint }
type GenS20 struct { F0 string }
type GenS21 struct { F0 uint; F1 uint16; F2 bool }
type GenS22 struct { F0 int8; F1 bool; F2 uint8 }
type GenS23 struct { F0 int16; F1 bool; F2 int32 }
type GenS24 struct { F0 int32; F1 bool }
type GenS25 struct { F0 int32 }
type GenS26 struct { F0 uint16; F1 bool }
type GenS27 struct { F0 bool; F1 int16 }

func main() {
    var v1 uint64 = 4294967295
    var v2 uint = 42
    var v3 uint = 1
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]bool)
    var v6 uint64 = (^(v1))
    var v7 string = "foo bar"
    var v8 uint32 = 4294967295
    switch -378 {
    case 213:
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := GenS1{F0: (((false || false) || true) && (false || true)), F1: 127, F2: (v1 & 4294967295)}
            var v11 int8 = 123
            _ = fi9
            _ = v10
            _ = v11
        }
    case -1:
        switch "" {
        case "hello":
            var v12 bool = false
            _ = v12
        case "\n\t":
            var v13 int32 = -348
            _ = v13
        default:
            if false {
                if (42 > 138) {
                    v3 += (v3 >> ((v3) & 63))
                    v14 := v4[2147483647]
                    _ = v14
                }
            } else {
                switch (v7 + v7) {
                case "\n\t":
                    v3 -= (v3 * v2)
                }
            }
        }
    case 34:
        v7 = (v7 + "\n\t")
    }
    switch 4294967295 {
    case 32767:
        v15 := GenS1{F0: (2147483647 == -2147483648), F1: 60, F2: (v6 - 439)}
        _ = v15
    case 42:
        switch 127 {
        case -626:
            for fi16 := 0; fi16 < 20; fi16++ {
                for fi17 := 0; fi17 < 8; fi17++ {
                    var v18 int16 = -1
                    var v19 bool = true
                    switch (v7 + "\n\t") {
                    case "\n\t":
                        v20 := []int{fi17, fi17, fi16}
                        _ = v20
                    default:
                        fmt.Println(fi16, v1)
                    }
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi16
            }
        case 9223372036854775807:
            v6 = (^(v6))
        default:
            for fi21 := 0; fi21 < 3; fi21++ {
                v22 := GenS2{F0: 128}
                _ = fi21
                _ = v22
            }
        }
    case 127:
        if false {
            for fi23 := 0; fi23 < 14; fi23++ {
                if (v7 >= v7) {
                    var v24 uint = v2
                    if (false || true) {
                        if (0.1 >= ((((1e10 * float32(-2.25)) + 3.14159) - (float32(0.0) * (3.14159 - 3.14159))) * (((3.14159 + ((0.0 / 1.0) + 1.5)) * float32(-2.25)) / -2.25))) {
                            if true {
                                _ = len(v4)
                                v25 := GenS3{F0: (v3 * v24)}
                                _ = v25
                            } else {
                                var v26 int32 = -27
                                var v27 uint16 = 63
                                _ = v26
                                _ = v27
                            }
                            var v28 bool = (!((1 < uint8(v6)) && true))
                            _ = v28
                        } else {
                            var v29 uint64 = 782
                            var v30 int = (fi23 << (((v24 * v2)) & 63))
                            v24 = v3
                            _ = v29
                            _ = v30
                        }
                        if (128 < 127) {
                            if true {
                                switch fi23 {
                                case 42:
                                    switch fi23 {
                                    case -428:
                                        v5[2] = false
                                    case 128:
                                        if false {
                                            switch v7 {
                                            case "\n\t":
                                                v31 := []int{fi23, -2147483648}
                                                _ = v31
                                            case "":
                                                for fi32 := 0; fi32 < 7; fi32++ {
                                                    if (-335 < 0) {
                                                        if ((fi32 != (fi23 + fi23)) && true) {
                                                            v6 -= (^(v6))
                                                        } else {
                                                            if true {
                                                                var v33 uint8 = 42
                                                                v34 := []int{fi32}
                                                                v35 := GenS2{F0: (v33 / (((v33 >> (((^(v3))) & 63))) | 1))}
                                                                _ = v33
                                                                _ = v34
                                                                _ = v35
                                                            }
                                                        }
                                                    }
                                                    v1 += 48
                                                    v36 := []int{fi23, 42, 4294967295, (-(fi32))}
                                                    _ = fi32
                                                    _ = v36
                                                }
                                            default:
                                                switch v7 {
                                                case "foo bar":
                                                    if (!false) {
                                                        var v37 uint8 = 170
                                                        for fi38 := 0; fi38 < 10; fi38++ {
                                                            var v39 float64 = ((float64(1.5) * 1e10) - 0.0)
                                                            v8 -= 4294967295
                                                            if false {
                                                                var v40 int64 = -1
                                                                _ = v40
                                                            }
                                                            _ = fi38
                                                            _ = v39
                                                        }
                                                        var v41 uint32 = v8
                                                        _ = v37
                                                        _ = v41
                                                    } else {
                                                        var v42 uint64 = (v1 ^ (^(v6)))
                                                        fmt.Println(v7)
                                                        _ = v42
                                                    }
                                                case "hello":
                                                    if (-1 < 43) {
                                                        v43 := GenS4{F0: ((("foo bar" + ((("" + v7) + (v7 + "")) + ("" + v7))) + v7) + v7)}
                                                        _ = v43
                                                    }
                                                case "":
                                                    v3 = 4294967295
                                                default:
                                                    v44 := v4[(fi23 / ((fi23) | 1))]
                                                    _ = v44
                                                }
                                            }
                                        } else {
                                            if false {
                                                if false {
                                                    for fi45 := 0; fi45 < 1; fi45++ {
                                                        if false {
                                                            v46 := GenS3{F0: v2}
                                                            var v47 int64 = 1
                                                            _ = v46
                                                            _ = v47
                                                        }
                                                        var v48 int64 = -9223372036854775808
                                                        var v49 uint32 = v8
                                                        _ = fi45
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                    fmt.Println(fi23, v8, v3)
                                                } else {
                                                    var v50 float32 = ((1.5 / -2.25) / (1.5 - (((1.5 * 0.0) + 1.5) * 0.0)))
                                                    var v51 uint16 = 965
                                                    _ = v50
                                                    _ = v51
                                                }
                                                v52 := []int{}
                                                var v53 uint = 42
                                                _ = v52
                                                _ = v53
                                            } else {
                                                delete(v4, fi23)
                                                switch fi23 {
                                                case 128:
                                                    v1 -= v1
                                                default:
                                                    switch -2147483648 {
                                                    case 0:
                                                        var v54 int32 = int32(v2)
                                                        _ = v54
                                                    case -763:
                                                        var v55 float64 = (3.14159 / (-2.25 + 0.1))
                                                        _ = v55
                                                    default:
                                                        fmt.Println(v3, v1)
                                                    }
                                                }
                                            }
                                            if (0.0 <= 1.5) {
                                                v24 -= v2
                                                v56 := GenS4{F0: "a"}
                                                _ = v56
                                            }
                                        }
                                    case 4294967295:
                                        switch fi23 {
                                        case 547:
                                            switch 128 {
                                            case 0:
                                                delete(v5, (v8 << (((^(v2))) & 63)))
                                            case -9223372036854775808:
                                                for fi57 := 0; fi57 < 13; fi57++ {
                                                    var v58 uint64 = (v1 / (((v6 | (^(v1)))) | 1))
                                                    v1 = 595
                                                    switch "a" {
                                                    case "a":
                                                        if (!(uint32(v24) == (v8 >> (((v2 << ((65535) & 63))) & 63)))) {
                                                            v58 += v58
                                                            _ = len(v4)
                                                            var v59 uint16 = 42
                                                            _ = v59
                                                        } else {
                                                            v7 = ""
                                                        }
                                                    case "foo bar":
                                                        switch (fi23 ^ (^(fi23))) {
                                                        case 4294967295:
                                                            var v60 int16 = 1
                                                            _ = v60
                                                        default:
                                                            _ = len(v5)
                                                        }
                                                    default:
                                                        v61 := GenS1{F0: false, F1: 127, F2: (v58 / ((1) | 1))}
                                                        _ = v61
                                                    }
                                                    v62 := GenS4{F0: "foo bar"}
                                                    _ = fi57
                                                    _ = v58
                                                    _ = v62
                                                }
                                            }
                                        case -682:
                                            var v63 int64 = -1
                                            _ = v63
                                        }
                                    }
                                case -940:
                                    var v64 int8 = 100
                                    _ = v64
                                case 0:
                                    fmt.Println(v6)
                                }
                                var v65 int32 = 2147483647
                                v66 := []int{(-(fi23))}
                                _ = v65
                                _ = v66
                            }
                            for fi67 := 0; fi67 < 1; fi67++ {
                                fmt.Println(fi23, fi67, v1)
                                v3 -= v3
                                v68 := []int{-2147483648}
                                fmt.Println(v3, v24, v1)
                                fmt.Println(v1, v2, v2)
                                _ = fi67
                                _ = v68
                            }
                        } else {
                            _ = len(v4)
                            switch 32768 {
                            case -828:
                                if true {
                                    v69 := GenS5{F0: (v8 * 843), F1: 1, F2: true}
                                    _ = v69
                                }
                            case 128:
                                v3 -= 255
                            case -652:
                                switch (("\n\t" + ("\n\t" + (v7 + v7))) + "hello") {
                                case "\n\t":
                                    v70 := GenS4{F0: "a"}
                                    _ = v70
                                case "hello":
                                    var v71 uint64 = (v6 * v6)
                                    _ = v71
                                case "a":
                                    v6 -= 518
                                default:
                                    v72 := GenS3{F0: v24}
                                    _ = v72
                                }
                            default:
                                switch (fi23 | fi23) {
                                case 573:
                                    v6 = v1
                                default:
                                    if (v7 < (((v7 + "a") + (v7 + "hello")) + v7)) {
                                        if (true || (v24 != v3)) {
                                            v3 -= (v2 ^ (v3 ^ 461))
                                        }
                                    } else {
                                        v24 = v24
                                        v73 := GenS6{F0: 255, F1: -1, F2: v7}
                                        for fi74 := 0; fi74 < 20; fi74++ {
                                            switch (v7 + ("hello" + v7)) {
                                            case "foo bar":
                                                var v75 uint16 = 620
                                                _ = v75
                                            }
                                            var v76 uint64 = 18446744073709551615
                                            _ = fi74
                                            _ = v76
                                        }
                                        _ = v73
                                    }
                                }
                            }
                        }
                        if false {
                            v77 := v4[-305]
                            _ = v77
                        }
                    }
                    _ = v24
                } else {
                    switch v7 {
                    case "a":
                        v78 := GenS6{F0: 128, F1: 0, F2: "hello"}
                        _ = v78
                    default:
                        var v79 int64 = 9223372036854775807
                        _ = v79
                    }
                    v2 = 1
                }
                _ = fi23
            }
            for fi80 := 0; fi80 < 19; fi80++ {
                switch "foo bar" {
                case "":
                    v81 := v4[(fi80 - -224)]
                    _ = v81
                case "hello":
                    if ((false || (32767 > 42)) || true) {
                        fmt.Println(v6, v2)
                    }
                case "foo bar":
                    if false {
                        switch 328 {
                        case 0:
                            _ = len(v5)
                        case 128:
                            var v82 uint64 = v6
                            _ = v82
                        }
                        v3 -= 1
                    } else {
                        v2 += 42
                        var v83 string = v7
                        _ = v83
                    }
                }
                fmt.Println(fi80, v1)
                v84 := GenS7{F0: 0, F1: (((fi80 ^ (fi80 % ((fi80) | 1))) >= (fi80 >> (((v2 * v2)) & 63))) && false)}
                _ = fi80
                _ = v84
            }
            if false {
                v85 := GenS2{F0: 42}
                _ = v85
            } else {
                v1 += uint64(v2)
            }
        }
    }
    for fi86 := 0; fi86 < 4; fi86++ {
        if ((!(!true)) && (-1 == -120)) {
            if (0 < 34) {
                v3 += (v3 % ((264) | 1))
                var v87 int32 = 42
                v2 += (v3 ^ (^(v2)))
                _ = v87
            }
        } else {
            switch v7 {
            case "hello":
                for fi88 := 0; fi88 < 3; fi88++ {
                    switch "a" {
                    case "hello":
                        if false {
                            v89 := GenS6{F0: uint8(v3), F1: 127, F2: v7}
                            v90 := []int{fi86, fi86, -243, fi86}
                            var v91 float64 = ((1.5 * ((0.1 + (1e10 / 0.1)) / -2.25)) + 0.1)
                            _ = v89
                            _ = v90
                            _ = v91
                        } else {
                            fmt.Println(fi86, v7, v2)
                            var v92 uint32 = (v8 << ((v3) & 63))
                            _ = v92
                        }
                    default:
                        switch (v7 + "hello") {
                        case "":
                            v3 -= (v2 * (v2 & 674))
                        }
                    }
                    v93 := GenS8{F0: false}
                    _ = fi88
                    _ = v93
                }
            case "a":
                fmt.Println(v8, v8, fi86)
            }
            switch (fi86 >> ((v2) & 63)) {
            case -9223372036854775808:
                switch v7 {
                case "\n\t":
                    for fi94 := 0; fi94 < 2; fi94++ {
                        v6 -= 1
                        if (1.5 != (1e10 / 0.1)) {
                            var v95 int32 = 2147483647
                            v1 = v6
                            _ = v95
                        } else {
                            var v96 uint16 = 2
                            switch (fi86 / ((int(v1)) | 1)) {
                            case -2147483648:
                                var v97 float64 = float64(0.0)
                                _ = v97
                            }
                            v1 = v1
                            var v98 uint16 = v96
                            _ = v96
                            _ = v98
                        }
                        fmt.Println(v2)
                        _ = fi94
                    }
                default:
                    for fi99 := 0; fi99 < 9; fi99++ {
                        var v100 uint64 = v6
                        var v101 uint8 = 26
                        switch "a" {
                        case "a":
                            fmt.Println(v100, v7, v101)
                        case "foo bar":
                            fmt.Println(v1)
                        default:
                            v101 = v101
                        }
                        var v102 int = -1
                        v8 = 1
                        _ = fi99
                        _ = v100
                        _ = v101
                        _ = v102
                    }
                }
            case 42:
                v7 = (("a" + v7) + v7)
            case 32768:
                switch v7 {
                case "\n\t":
                    if (true && (!true)) {
                        delete(v4, 127)
                    }
                case "":
                    for fi103 := 0; fi103 < 18; fi103++ {
                        v104 := GenS9{F0: -1, F1: v2, F2: v2}
                        var v105 float64 = 1.5
                        _ = fi103
                        _ = v104
                        _ = v105
                    }
                case "a":
                    var v106 uint32 = v8
                    _ = v106
                default:
                    if ((^(v1)) != v6) {
                        var v107 int = (fi86 + (fi86 >> (((v3 & v2)) & 63)))
                        if (v7 <= v7) {
                            v107 = (fi86 / (((v107 | 421)) | 1))
                            if false {
                                if false {
                                    if (!(v8 >= v8)) {
                                        var v108 uint64 = 0
                                        for fi109 := 0; fi109 < 9; fi109++ {
                                            v110 := GenS10{F0: 2}
                                            _ = fi109
                                            _ = v110
                                        }
                                        _ = v108
                                    } else {
                                        fmt.Println(v7)
                                        v111 := []int{2147483647, (v107 % ((656) | 1)), (fi86 << (((^(v2))) & 63)), (^(fi86))}
                                        var v112 int16 = 1000
                                        _ = v111
                                        _ = v112
                                    }
                                }
                                switch ("\n\t" + (v7 + "hello")) {
                                case "hello":
                                    for fi113 := 0; fi113 < 11; fi113++ {
                                        switch (fi113 << (((^(v2))) & 63)) {
                                        case 32767:
                                            var v114 uint = v3
                                            _ = len(v5)
                                            if false {
                                                if true {
                                                    v115 := GenS5{F0: v8, F1: -1, F2: (558 < -1)}
                                                    _ = v115
                                                }
                                                v116 := GenS11{F0: true, F1: ("\n\t" + (v7 + ((("a" + v7) + "") + "\n\t")))}
                                                _ = v116
                                            } else {
                                                if false {
                                                    if (2 != 255) {
                                                        v1 += (v6 & v1)
                                                        var v117 bool = true
                                                        _ = v117
                                                    }
                                                    var v118 int = (-(v107))
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        if (42 == 1) {
                                                            fmt.Println(v2, fi113)
                                                            v119 := GenS5{F0: v8, F1: -1, F2: true}
                                                            _ = v119
                                                        } else {
                                                            switch fi86 {
                                                            case 664:
                                                                var v120 string = v7
                                                                _ = v120
                                                            default:
                                                                v121 := GenS12{F0: 52, F1: 155, F2: (v8 + (v8 / ((4294967295) | 1)))}
                                                                var v122 uint16 = 255
                                                                _ = v121
                                                                _ = v122
                                                            }
                                                            fmt.Println(v3, v3, v8)
                                                            switch fi113 {
                                                            case 128:
                                                                v107 = v107
                                                                v7 = "a"
                                                                var v123 int32 = 42
                                                                _ = v123
                                                            case 4294967295:
                                                                if true {
                                                                    if (v7 == "\n\t") {
                                                                        var v124 int64 = 32
                                                                        var v125 float32 = 1e10
                                                                        v126 := GenS12{F0: 0, F1: 255, F2: v8}
                                                                        _ = v124
                                                                        _ = v125
                                                                        _ = v126
                                                                    }
                                                                }
                                                            case 0:
                                                                if false {
                                                                    var v127 int = v107
                                                                    _ = v127
                                                                }
                                                            default:
                                                                v5[130] = (128 == 199)
                                                            }
                                                        }
                                                    }
                                                    var v128 int16 = 42
                                                    var v129 float64 = ((((0.1 + (1.5 / 1e10)) + 1.5) - -2.25) * 0.0)
                                                    _ = v118
                                                    _ = v128
                                                    _ = v129
                                                } else {
                                                    var v130 int16 = 217
                                                    if false {
                                                        if (!false) {
                                                            v131 := GenS6{F0: 149, F1: 1, F2: "hello"}
                                                            _ = v131
                                                        }
                                                        v132 := v4[-698]
                                                        v130 -= 42
                                                        _ = v132
                                                    }
                                                    _ = v130
                                                }
                                                v133 := GenS2{F0: 102}
                                                _ = v133
                                            }
                                            _ = v114
                                        case 127:
                                            var v134 uint64 = (v1 | 2)
                                            _ = v134
                                        }
                                        _ = fi113
                                    }
                                case "\n\t":
                                    switch 4294967295 {
                                    case 613:
                                        switch fi86 {
                                        case -2147483648:
                                            v135 := GenS2{F0: 198}
                                            _ = v135
                                        }
                                    case -388:
                                        delete(v4, (v107 + (fi86 ^ fi86)))
                                    case 812:
                                        for fi136 := 0; fi136 < 1; fi136++ {
                                            if false {
                                                v137 := GenS13{F0: 32767, F1: -9223372036854775808}
                                                v138 := GenS14{F0: v8, F1: 2}
                                                _ = v137
                                                _ = v138
                                            }
                                            _ = fi136
                                        }
                                    default:
                                        if true {
                                            if (((fi86 | fi86) >= 128) && false) {
                                                var v139 uint64 = (^(v6))
                                                var v140 float64 = ((0.0 - -2.25) * 3.14159)
                                                v141 := GenS9{F0: 32767, F1: v3, F2: (v3 ^ v2)}
                                                _ = v139
                                                _ = v140
                                                _ = v141
                                            } else {
                                                switch "a" {
                                                case "foo bar":
                                                    for fi142 := 0; fi142 < 4; fi142++ {
                                                        v143 := []int{(fi142 | v107), (v107 + (fi86 | (v107 + 602)))}
                                                        var v144 string = ("" + "foo bar")
                                                        var v145 uint = 2
                                                        if true {
                                                            var v146 int = (fi86 | 0)
                                                            var v147 int32 = 1
                                                            var v148 uint32 = (v8 * v8)
                                                            v149 := GenS15{F0: v144}
                                                            _ = v146
                                                            _ = v147
                                                            _ = v148
                                                            _ = v149
                                                        }
                                                        v144 = v7
                                                        _ = fi142
                                                        _ = v143
                                                        _ = v144
                                                        _ = v145
                                                    }
                                                }
                                                var v150 float64 = 3.14159
                                                if true {
                                                    v150 -= ((((v150 * float64(0.1)) + ((0.1 / 1.0) - v150)) * v150) + v150)
                                                } else {
                                                    if true {
                                                        var v151 bool = false
                                                        _ = v151
                                                    } else {
                                                        for fi152 := 0; fi152 < 18; fi152++ {
                                                            _ = len(v4)
                                                            v1 = 293
                                                            _ = fi152
                                                        }
                                                        var v153 float32 = 3.14159
                                                        v154 := GenS16{F0: 32767, F1: v8, F2: (v7 + "a")}
                                                        _ = v153
                                                        _ = v154
                                                    }
                                                    v107 += fi86
                                                    v155 := GenS17{F0: (v8 - uint32(v3)), F1: (true && (1e10 == float64(-2.25))), F2: true}
                                                    _ = v155
                                                }
                                                _ = v150
                                            }
                                            if false {
                                                fmt.Println(v6, v8, v1)
                                                fmt.Println(v8)
                                                v5[(v8 * v8)] = true
                                            } else {
                                                for fi156 := 0; fi156 < 4; fi156++ {
                                                    switch (fi86 ^ (v107 * fi86)) {
                                                    case -73:
                                                        fmt.Println(v2, v7)
                                                    default:
                                                        v157 := []int{(v107 << ((v3) & 63)), (^(fi86)), v107, (^(fi156))}
                                                        _ = v157
                                                    }
                                                    _ = fi156
                                                }
                                            }
                                            v158 := GenS1{F0: true, F1: 1, F2: (v6 - (v6 - (^(v6))))}
                                            _ = v158
                                        }
                                    }
                                case "foo bar":
                                    if (128 == 42) {
                                        v3 = (v2 % ((394) | 1))
                                    } else {
                                        for fi159 := 0; fi159 < 13; fi159++ {
                                            if (-1 <= -128) {
                                                var v160 uint8 = 127
                                                v160 = v160
                                                v161 := []int{}
                                                _ = v160
                                                _ = v161
                                            } else {
                                                v162 := []int{42}
                                                _ = v162
                                            }
                                            _ = fi159
                                        }
                                        switch "\n\t" {
                                        case "a":
                                            var v163 uint8 = 1
                                            _ = v163
                                        case "hello":
                                            if true {
                                                switch (v107 | (^(v107))) {
                                                case 4294967295:
                                                    if (v6 == 18446744073709551615) {
                                                        for fi164 := 0; fi164 < 2; fi164++ {
                                                            var v165 int16 = 1000
                                                            _ = fi164
                                                            _ = v165
                                                        }
                                                        var v166 int = v107
                                                        if (!false) {
                                                            v107 += -2147483648
                                                            fmt.Println(v8)
                                                            v167 := []int{v166, fi86, (fi86 | (v107 | (v107 - (v166 % ((v107) | 1))))), 4294967295}
                                                            _ = v167
                                                        } else {
                                                            var v168 uint = (v3 - 255)
                                                            v169 := GenS4{F0: v7}
                                                            if true {
                                                                v168 += (v2 & (v2 + v2))
                                                                var v170 int16 = 42
                                                                for fi171 := 0; fi171 < 3; fi171++ {
                                                                    v172 := v4[(v107 + (v166 - (v166 / ((-593) | 1))))]
                                                                    fmt.Println(fi171)
                                                                    var v173 int64 = -53
                                                                    v174 := GenS1{F0: false, F1: 1, F2: v1}
                                                                    switch (fi86 + (v166 * -1)) {
                                                                    case -260:
                                                                        if true {
                                                                            v175 := GenS18{F0: (((v7 + (v7 + "")) + v7) + "a"), F1: 2, F2: 255}
                                                                            if true {
                                                                                v4[-2147483648] = (v172 | v172)
                                                                                v176 := []int{v107, (fi86 >> (((v168 - v2)) & 63)), (fi171 % (((fi86 + (fi86 & v107))) | 1)), fi171}
                                                                                v177 := GenS19{F0: 0, F1: uint8(v6), F2: v168}
                                                                                _ = v176
                                                                                _ = v177
                                                                            } else {
                                                                                v178 := GenS20{F0: v7}
                                                                                _ = v178
                                                                            }
                                                                            if false {
                                                                                var v179 uint64 = v6
                                                                                _ = v179
                                                                            } else {
                                                                                if (!(0 <= uint8(v6))) {
                                                                                    v180 := GenS5{F0: (v8 >> ((v2) & 63)), F1: 0, F2: (!(v2 < 596))}
                                                                                    _ = v180
                                                                                } else {
                                                                                    if (!(v8 == v172)) {
                                                                                        v2 -= v3
                                                                                        if (!false) {
                                                                                            v4[(fi171 << (((v2 | (v168 * (^(v2))))) & 63))] = uint32(v3)
                                                                                            v181 := v5[(v8 / ((v172) | 1))]
                                                                                            _ = v181
                                                                                        } else {
                                                                                            v8 -= v8
                                                                                            switch (v107 | v166) {
                                                                                            case 128:
                                                                                                switch (fi171 ^ (-(fi171))) {
                                                                                                case -1:
                                                                                                    if (!false) {
                                                                                                        v7 = v7
                                                                                                        if true {
                                                                                                            if false {
                                                                                                                v182 := []int{(fi86 * (^(fi86))), 2147483647, v166, (fi171 | (fi171 & int(v173)))}
                                                                                                                var v183 int = (fi171 * (v107 % (((v166 / ((fi171) | 1))) | 1)))
                                                                                                                switch (("hello" + v7) + "a") {
                                                                                                                case "foo bar":
                                                                                                                    fmt.Println(fi86)
                                                                                                                case "hello":
                                                                                                                    var v184 uint64 = 336
                                                                                                                    _ = v184
                                                                                                                }
                                                                                                                _ = v182
                                                                                                                _ = v183
                                                                                                            }
                                                                                                            switch (fi86 % ((32768) | 1)) {
                                                                                                            case -260:
                                                                                                                switch v107 {
                                                                                                                case 2147483647:
                                                                                                                    v3 -= v2
                                                                                                                default:
                                                                                                                    if (0 != 42) {
                                                                                                                        v107 = v107
                                                                                                                    } else {
                                                                                                                        if ((v173 | (v173 % ((-392) | 1))) == v173) {
                                                                                                                            var v185 uint8 = 224
                                                                                                                            v186 := GenS6{F0: (v185 - 0), F1: -128, F2: v7}
                                                                                                                            v187 := v5[uint32(v2)]
                                                                                                                            _ = v185
                                                                                                                            _ = v186
                                                                                                                            _ = v187
                                                                                                                        } else {
                                                                                                                            v188 := GenS14{F0: v172, F1: 438}
                                                                                                                            v189 := []int{v166}
                                                                                                                            _ = v188
                                                                                                                            _ = v189
                                                                                                                        }
                                                                                                                        v2 = 42
                                                                                                                        v166 = v107
                                                                                                                    }
                                                                                                                }
                                                                                                            default:
                                                                                                                switch fi171 {
                                                                                                                case -9223372036854775808:
                                                                                                                    v8 = v8
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                case 0:
                                                                                                    switch (("" + "") + ((((v7 + "") + (v7 + v7)) + v7) + ("hello" + "hello"))) {
                                                                                                    case "":
                                                                                                        v2 = v2
                                                                                                    case "foo bar":
                                                                                                        v3 = (v3 % ((0) | 1))
                                                                                                    case "\n\t":
                                                                                                        switch v107 {
                                                                                                        case -9223372036854775808:
                                                                                                            switch (fi171 & 127) {
                                                                                                            case -483:
                                                                                                                var v190 float64 = 1.5
                                                                                                                _ = v190
                                                                                                            case 127:
                                                                                                                v191 := GenS7{F0: (v6 % (((v1 * (^(v1)))) | 1)), F1: false}
                                                                                                                _ = v191
                                                                                                            case 32768:
                                                                                                                v5[(v172 % ((v8) | 1))] = (false && true)
                                                                                                            default:
                                                                                                                v173 = v173
                                                                                                            }
                                                                                                        }
                                                                                                    default:
                                                                                                        v192 := []int{}
                                                                                                        _ = v192
                                                                                                    }
                                                                                                case 178:
                                                                                                    fmt.Println(v166, v107)
                                                                                                    var v193 int16 = v170
                                                                                                    _ = v193
                                                                                                }
                                                                                                delete(v4, -782)
                                                                                            case 264:
                                                                                                v8 -= v172
                                                                                                v173 -= v173
                                                                                            }
                                                                                            var v194 int64 = (^(v173))
                                                                                            _ = v194
                                                                                        }
                                                                                        var v195 int16 = (v170 >> (((v3 / ((v168) | 1))) & 63))
                                                                                        _ = v195
                                                                                    } else {
                                                                                        v4[v166] = (v172 * v172)
                                                                                        switch (v107 + fi86) {
                                                                                        case 2147483647:
                                                                                            var v196 int8 = 100
                                                                                            _ = v196
                                                                                        }
                                                                                        if ((-(v173)) == (^(v173))) {
                                                                                            var v197 uint16 = 0
                                                                                            _ = v197
                                                                                        } else {
                                                                                            fmt.Println(v1, v3)
                                                                                        }
                                                                                    }
                                                                                    switch ("foo bar" + v7) {
                                                                                    case "a":
                                                                                        var v198 int64 = (v173 / (((v173 | v173)) | 1))
                                                                                        _ = v198
                                                                                    case "hello":
                                                                                        v173 -= (v173 & (v173 / ((-1) | 1)))
                                                                                    case "foo bar":
                                                                                        switch 127 {
                                                                                        case 127:
                                                                                            switch "\n\t" {
                                                                                            case "\n\t":
                                                                                                switch "" {
                                                                                                case "foo bar":
                                                                                                    var v199 bool = false
                                                                                                    _ = v199
                                                                                                case "hello":
                                                                                                    v107 = (fi86 | v107)
                                                                                                default:
                                                                                                    v200 := v4[(v166 % (((fi171 * fi171)) | 1))]
                                                                                                    _ = v200
                                                                                                }
                                                                                            case "hello":
                                                                                                switch (v166 + fi86) {
                                                                                                case -2147483648:
                                                                                                    switch v7 {
                                                                                                    case "foo bar":
                                                                                                        v201 := GenS21{F0: (v2 * uint(v8)), F1: 593, F2: (true || (!(!(!false))))}
                                                                                                        _ = v201
                                                                                                    case "\n\t":
                                                                                                        fmt.Println(v168, v2)
                                                                                                    case "hello":
                                                                                                        v170 -= v170
                                                                                                    default:
                                                                                                        var v202 string = ("foo bar" + (("hello" + ("\n\t" + ("foo bar" + (v7 + v7)))) + (("foo bar" + "\n\t") + "foo bar")))
                                                                                                        _ = v202
                                                                                                    }
                                                                                                }
                                                                                                switch "hello" {
                                                                                                case "":
                                                                                                    var v203 uint = v3
                                                                                                    _ = v203
                                                                                                case "foo bar":
                                                                                                    var v204 float32 = 3.14159
                                                                                                    _ = v204
                                                                                                }
                                                                                                v170 = (v170 * v170)
                                                                                                v4[v166] = 139
                                                                                                if (65535 < 490) {
                                                                                                    v205 := GenS18{F0: v7, F1: uint16(v8), F2: 2}
                                                                                                    _ = len(v5)
                                                                                                    _ = v205
                                                                                                } else {
                                                                                                    v170 = -1
                                                                                                }
                                                                                                v206 := GenS17{F0: (v172 - v172), F1: false, F2: (uint(v172) == (v168 ^ v2))}
                                                                                                _ = v206
                                                                                            default:
                                                                                                v8 = v8
                                                                                            }
                                                                                        default:
                                                                                            if ((v1 & 849) >= v1) {
                                                                                                switch fi86 {
                                                                                                case -2147483648:
                                                                                                    var v207 uint8 = 130
                                                                                                    _ = v207
                                                                                                case -9223372036854775808:
                                                                                                    v170 += int16(fi86)
                                                                                                case 128:
                                                                                                    if (!true) {
                                                                                                        var v208 float32 = float32(1e10)
                                                                                                        var v209 int16 = (v170 - -32768)
                                                                                                        _ = v208
                                                                                                        _ = v209
                                                                                                    } else {
                                                                                                        v210 := []int{(fi86 + -605), v166, 127, fi171}
                                                                                                        if len(v210) > 0 {
                                                                                                            v210[0] = -329
                                                                                                        }
                                                                                                        var v211 bool = (int8(v172) < -128)
                                                                                                        var v212 uint32 = v8
                                                                                                        _ = v210
                                                                                                        _ = v211
                                                                                                        _ = v212
                                                                                                    }
                                                                                                default:
                                                                                                    switch v7 {
                                                                                                    case "foo bar":
                                                                                                        var v213 int16 = v170
                                                                                                        _ = v213
                                                                                                    default:
                                                                                                        v173 = (v173 & (v173 >> (((v2 << ((674) & 63))) & 63)))
                                                                                                    }
                                                                                                }
                                                                                                if true {
                                                                                                    switch v7 {
                                                                                                    case "foo bar":
                                                                                                        var v214 int64 = 828
                                                                                                        _ = v214
                                                                                                    case "hello":
                                                                                                        var v215 uint64 = (v1 ^ 770)
                                                                                                        _ = v215
                                                                                                    default:
                                                                                                        v1 = v1
                                                                                                    }
                                                                                                    var v216 uint16 = 32767
                                                                                                    _ = v216
                                                                                                } else {
                                                                                                    v217 := GenS15{F0: "foo bar"}
                                                                                                    _ = v217
                                                                                                }
                                                                                                var v218 int16 = 42
                                                                                                _ = v218
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v219 int8 = -128
                                                                                switch 0 {
                                                                                case -2147483648:
                                                                                    v220 := GenS22{F0: (v219 >> ((v2) & 63)), F1: false, F2: 255}
                                                                                    _ = v220
                                                                                case 431:
                                                                                    var v221 int8 = (^(v219))
                                                                                    _ = v221
                                                                                }
                                                                                _ = v219
                                                                            }
                                                                            _ = v175
                                                                        }
                                                                    case 2147483647:
                                                                        switch (fi171 + (fi171 % (((v107 / ((fi86) | 1))) | 1))) {
                                                                        case 4294967295:
                                                                            switch v7 {
                                                                            case "hello":
                                                                                var v222 int64 = 9223372036854775807
                                                                                _ = v222
                                                                            case "":
                                                                                var v223 float32 = 1.5
                                                                                _ = v223
                                                                            default:
                                                                                v224 := []int{127, (fi86 + 128), fi86}
                                                                                _ = v224
                                                                            }
                                                                        case 0:
                                                                            if false {
                                                                                v173 = (v173 | 0)
                                                                            }
                                                                        default:
                                                                            v225 := []int{}
                                                                            _ = v225
                                                                        }
                                                                    }
                                                                    _ = fi171
                                                                    _ = v172
                                                                    _ = v173
                                                                    _ = v174
                                                                }
                                                                _ = v170
                                                            } else {
                                                                if (767 == -607) {
                                                                    v107 -= (-(v166))
                                                                }
                                                            }
                                                            _ = v168
                                                            _ = v169
                                                        }
                                                        _ = v166
                                                    }
                                                default:
                                                    var v226 bool = (((true || (3.14159 < 1e10)) || ((false && false) || (9223372036854775807 > 42))) && (false && (-273 > 1000)))
                                                    _ = v226
                                                }
                                                delete(v5, (v8 / ((v8) | 1)))
                                                v8 += (v8 - v8)
                                            }
                                        default:
                                            if (true && (42 == 1)) {
                                                v227 := []int{-1}
                                                _ = v227
                                            } else {
                                                if false {
                                                    for fi228 := 0; fi228 < 7; fi228++ {
                                                        v229 := GenS11{F0: (!(3.14159 < ((1e10 * float64((-2.25 + (1.5 + (3.14159 + 1.5))))) - (-2.25 - 3.14159)))), F1: "hello"}
                                                        fmt.Println(v6, v107)
                                                        _ = fi228
                                                        _ = v229
                                                    }
                                                } else {
                                                    v107 = (fi86 | 508)
                                                    var v230 int8 = 0
                                                    _ = v230
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        } else {
                            v231 := []int{fi86, 541}
                            _ = v231
                        }
                        var v232 int64 = -995
                        _ = v107
                        _ = v232
                    }
                }
            }
            v7 = v7
        }
        var v233 string = v7
        if (!false) {
            v234 := []int{fi86, fi86}
            _ = v234
        }
        _ = fi86
        _ = v233
    }
    v7 = ("a" + "\n\t")
    v8 += 563
    var v235 int = 32767
    v4[(-(v235))] = v8
    v1 = v1
    switch (((v7 + "hello") + v7) + v7) {
    case "foo bar":
        v6 = (^(v1))
    case "\n\t":
        v6 += v6
    case "":
        var v236 uint = uint(v8)
        _ = v236
    }
    switch (v235 + (v235 >> (((v2 / ((v3) | 1))) & 63))) {
    case 331:
        if false {
            fmt.Println(v7, v3)
        } else {
            v237 := GenS23{F0: -32768, F1: true, F2: -323}
            switch v7 {
            case "":
                if true {
                    v8 = (v8 * 202)
                    if true {
                        v6 -= 2
                        var v238 int64 = -9223372036854775808
                        if (true && ((1e10 / 1e10) > 1e10)) {
                            for fi239 := 0; fi239 < 10; fi239++ {
                                var v240 int16 = 1000
                                if (!((1e10 - 3.14159) <= 0.1)) {
                                    var v241 float64 = 3.14159
                                    if (!(613 > (v238 * v238))) {
                                        v242 := GenS8{F0: true}
                                        switch ((v7 + ("" + "")) + "a") {
                                        case "\n\t":
                                            v243 := v4[(-(v235))]
                                            _ = v243
                                        case "a":
                                            v244 := []int{(^(v235)), -807, (v235 % ((42) | 1)), fi239}
                                            _ = v244
                                        case "foo bar":
                                            v6 = v6
                                        }
                                        var v245 int8 = 1
                                        _ = v242
                                        _ = v245
                                    }
                                    _ = v241
                                } else {
                                    var v246 int = (v235 / ((127) | 1))
                                    if false {
                                        if ((v238 - v238) < v238) {
                                            if true {
                                                if (9223372036854775807 > int64(v8)) {
                                                    if false {
                                                        v247 := []int{502, 685}
                                                        if (1 >= v8) {
                                                            var v248 uint = 255
                                                            v2 += 4294967295
                                                            var v249 uint = (^(v3))
                                                            _ = v248
                                                            _ = v249
                                                        }
                                                        _ = v247
                                                    } else {
                                                        delete(v5, (v8 | (v8 ^ (v8 / (((v8 << ((v3) & 63))) | 1)))))
                                                    }
                                                    switch v246 {
                                                    case -1:
                                                        for fi250 := 0; fi250 < 8; fi250++ {
                                                            var v251 string = ((((v7 + "\n\t") + v7) + v7) + v7)
                                                            _ = fi250
                                                            _ = v251
                                                        }
                                                    }
                                                    v252 := GenS8{F0: (-71 != 1)}
                                                    _ = v252
                                                }
                                            }
                                        }
                                        switch (((v7 + "foo bar") + v7) + v7) {
                                        case "a":
                                            switch "foo bar" {
                                            case "a":
                                                v253 := GenS5{F0: (v8 & (v8 * (v8 / (((v8 - v8)) | 1)))), F1: 2147483647, F2: false}
                                                _ = v253
                                            case "foo bar":
                                                switch "a" {
                                                case "":
                                                    v7 = v7
                                                case "foo bar":
                                                    var v254 int16 = -32768
                                                    _ = v254
                                                default:
                                                    if (false || (v240 != v240)) {
                                                        v8 = (v8 << (((v3 >> ((v2) & 63))) & 63))
                                                        v255 := GenS24{F0: 0, F1: (int32(v246) <= -1)}
                                                        v1 = 886
                                                        _ = v255
                                                    }
                                                }
                                            case "":
                                                for fi256 := 0; fi256 < 8; fi256++ {
                                                    var v257 bool = true
                                                    v258 := GenS24{F0: -775, F1: (!(128 > 255))}
                                                    v259 := GenS25{F0: -2147483648}
                                                    fmt.Println(v238, v3, v8)
                                                    _ = fi256
                                                    _ = v257
                                                    _ = v258
                                                    _ = v259
                                                }
                                            }
                                        case "hello":
                                            if false {
                                                var v260 string = v7
                                                _ = v260
                                            }
                                        }
                                    }
                                    _ = v246
                                }
                                if (true && true) {
                                    v235 -= (^(fi239))
                                    v235 = v235
                                } else {
                                    switch v235 {
                                    case 1:
                                        fmt.Println(v6, v8, v6)
                                    case 316:
                                        for fi261 := 0; fi261 < 10; fi261++ {
                                            v240 = 1
                                            v1 = (v1 >> ((v2) & 63))
                                            switch v7 {
                                            case "":
                                                v262 := []int{127, v235}
                                                _ = v262
                                            case "a":
                                                v1 += (v1 | v6)
                                            case "foo bar":
                                                v238 = 1
                                            default:
                                                v8 = 2147483648
                                            }
                                            _ = fi261
                                        }
                                    default:
                                        var v263 float64 = -2.25
                                        _ = v263
                                    }
                                    for fi264 := 0; fi264 < 7; fi264++ {
                                        v3 += (v3 ^ 178)
                                        delete(v5, (^(v8)))
                                        _ = fi264
                                    }
                                    v6 = v6
                                }
                                _ = fi239
                                _ = v240
                            }
                            fmt.Println(v2, v2, v6)
                            v265 := []int{(^(v235))}
                            _ = v265
                        } else {
                            switch -9223372036854775808 {
                            case 2147483647:
                                if (v7 < ("hello" + "foo bar")) {
                                    v3 -= (^(v2))
                                    switch (v235 * -236) {
                                    case 758:
                                        var v266 uint8 = 128
                                        _ = v266
                                    case -9223372036854775808:
                                        fmt.Println(v1, v235)
                                    default:
                                        v267 := GenS26{F0: 42, F1: true}
                                        _ = v267
                                    }
                                    v7 = (v7 + (("" + "\n\t") + "a"))
                                }
                            case -709:
                                fmt.Println(v2)
                            default:
                                fmt.Println(v6)
                            }
                        }
                        _ = v238
                    } else {
                        v7 = v7
                        v8 = 897
                        fmt.Println(v2, v2, v2)
                    }
                } else {
                    for fi268 := 0; fi268 < 9; fi268++ {
                        v235 = (fi268 - v235)
                        v235 = v235
                        v6 -= v6
                        _ = fi268
                    }
                }
            case "a":
                v6 -= 285
            case "hello":
                switch v7 {
                case "":
                    v2 += (^(v3))
                case "foo bar":
                    if (-1 <= 19) {
                        var v269 uint8 = 0
                        if true {
                            fmt.Println(v3)
                            fmt.Println(v235, v6)
                            v270 := GenS12{F0: (v269 + v269), F1: v269, F2: (v8 << (((^(v2))) & 63))}
                            _ = v270
                        } else {
                            if true {
                                v271 := v5[4294967295]
                                fmt.Println(v271, v1)
                                if v271 {
                                    v235 -= v235
                                    v8 += v8
                                    v6 = v6
                                } else {
                                    for fi272 := 0; fi272 < 12; fi272++ {
                                        for fi273 := 0; fi273 < 6; fi273++ {
                                            var v274 int = (fi273 & v235)
                                            v275 := GenS5{F0: v8, F1: -2147483648, F2: ((v8 % ((v8) | 1)) >= 4294967295)}
                                            if ((v7 + "hello") <= (v7 + "hello")) {
                                                v276 := GenS27{F0: (!v271), F1: -1}
                                                fmt.Println(v1, v274)
                                                v6 = (v1 << ((v2) & 63))
                                                _ = v276
                                            }
                                            _ = fi273
                                            _ = v274
                                            _ = v275
                                        }
                                        var v277 int16 = -1
                                        _ = fi272
                                        _ = v277
                                    }
                                }
                                _ = v271
                            }
                            var v278 int16 = 42
                            v278 += (v278 - (v278 | (v278 + (v278 >> ((v2) & 63)))))
                            _ = v278
                        }
                        if false {
                            for fi279 := 0; fi279 < 13; fi279++ {
                                for fi280 := 0; fi280 < 2; fi280++ {
                                    var v281 uint = (v2 + v3)
                                    _ = fi280
                                    _ = v281
                                }
                                var v282 uint16 = 32767
                                _ = fi279
                                _ = v282
                            }
                        }
                        _ = v269
                    }
                default:
                    v283 := []int{(v235 / ((v235) | 1)), v235, v235, v235}
                    _ = v283
                }
            default:
                v284 := GenS9{F0: -89, F1: uint(v8), F2: (v2 & (v2 ^ 1))}
                _ = v284
            }
            v285 := []int{v235}
            _ = v237
            _ = v285
        }
    case 1:
        fmt.Println(v1, v3, v7)
    }
    v8 += v8
    fmt.Println(v7, v2, v7)
    v235 = v235
    v286 := []int{-844, (v235 & int(v8)), 548}
    fmt.Println(v8, v235)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v235:", v235)
    fmt.Println("slicelen v286:", len(v286), v286)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v235
    _ = v286
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
