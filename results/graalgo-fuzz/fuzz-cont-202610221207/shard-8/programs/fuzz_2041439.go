package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64; F1 int32 }
type GenS7 struct { F0 int64; F1 int16; F2 int }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint64; F1 int16 }
type GenS10 struct { F0 uint8; F1 int64 }
type GenS11 struct { F0 uint32; F1 int }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 uint64; F1 int8 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 int16; F1 uint16 }
type GenS16 struct { F0 int32; F1 int16; F2 int }
type GenS17 struct { F0 uint16; F1 bool; F2 int8 }
type GenS18 struct { F0 uint16; F1 int8 }
type GenS19 struct { F0 uint8; F1 int16; F2 uint8 }

func main() {
    var v1 uint16 = 65535
    var v2 int16 = -1
    var v3 uint32 = 42
    var v4 int64 = 324
    v5 := make(map[int]uint32)
    var v6 uint = 4294967295
    var v7 float32 = (3.14159 * float32(((1e10 - float64(-2.25)) + (0.1 - (((1e10 * -2.25) / ((float64(0.0) + 1.5) * ((3.14159 * 3.14159) * 1e10))) * (3.14159 / (1.5 + 1e10)))))))
    if true {
        for fi8 := 0; fi8 < 14; fi8++ {
            if false {
                fmt.Println(v3)
                var v9 bool = true
                var v10 int16 = 1
                _ = v10
                _ = v9
            }
            v11 := []int{(-(fi8))}
            _ = fi8
            _ = v11
        }
        if (1 < v3) {
            if true {
                v12 := v5[2147483647]
                if (!true) {
                    if true {
                        v5[1] = (v3 % (((v12 << ((v6) & 63))) | 1))
                        for fi13 := 0; fi13 < 19; fi13++ {
                            v6 += v6
                            _ = fi13
                        }
                        v14 := GenS1{F0: v2}
                        _ = v14
                    }
                } else {
                    switch 32768 {
                    case -1:
                        v15 := GenS2{F0: 2, F1: 1, F2: 271}
                        _ = v15
                    case -9223372036854775808:
                        switch -848 {
                        case -2147483648:
                            for fi16 := 0; fi16 < 15; fi16++ {
                                switch (-(fi16)) {
                                case 2147483647:
                                    var v17 float64 = 1.5
                                    _ = v17
                                }
                                switch (^(fi16)) {
                                case -2147483648:
                                    if (v1 != 42) {
                                        v18 := []int{(fi16 << (((^(v6))) & 63)), -824}
                                        _ = v18
                                    } else {
                                        v19 := v5[127]
                                        _ = v19
                                    }
                                case -904:
                                    switch "hello" {
                                    case "\n\t":
                                        var v20 int16 = (^(v2))
                                        _ = v20
                                    default:
                                        for fi21 := 0; fi21 < 19; fi21++ {
                                            v22 := []int{(fi21 << (((v6 << ((v6) & 63))) & 63)), fi21}
                                            _ = fi21
                                            _ = v22
                                        }
                                    }
                                case 1:
                                    switch fi16 {
                                    case 2147483647:
                                        switch "foo bar" {
                                        case "a":
                                            var v23 string = "hello"
                                            _ = v23
                                        default:
                                            for fi24 := 0; fi24 < 3; fi24++ {
                                                v12 = (^(v3))
                                                var v25 int64 = v4
                                                v2 += 1000
                                                _ = fi24
                                                _ = v25
                                            }
                                        }
                                    case 0:
                                        if (true || true) {
                                            v12 += 2147483647
                                        }
                                    default:
                                        v2 += v2
                                    }
                                }
                                var v26 uint32 = v12
                                _ = fi16
                                _ = v26
                            }
                        case -1:
                            switch "" {
                            case "foo bar":
                                delete(v5, -670)
                            case "a":
                                switch "" {
                                case "a":
                                    switch -621 {
                                    case 127:
                                        var v27 int = -2147483648
                                        _ = v27
                                    default:
                                        if (42 == 1) {
                                            switch "" {
                                            case "hello":
                                                for fi28 := 0; fi28 < 6; fi28++ {
                                                    if true {
                                                        var v29 uint16 = v1
                                                        _ = v29
                                                    } else {
                                                        if (v12 >= v12) {
                                                            var v30 bool = ((v12 % (((^(v12))) | 1)) >= (v3 & 1))
                                                            _ = v30
                                                        }
                                                    }
                                                    _ = fi28
                                                }
                                            default:
                                                v31 := v5[2147483647]
                                                _ = v31
                                            }
                                            if (!(!false)) {
                                                v7 = v7
                                                if true {
                                                    if (!(true && ("\n\t" >= ("a" + ("" + ((("\n\t" + "\n\t") + "\n\t") + "\n\t")))))) {
                                                        switch ("hello" + (("" + "hello") + "foo bar")) {
                                                        case "a":
                                                            var v32 string = "foo bar"
                                                            _ = v32
                                                        default:
                                                            for fi33 := 0; fi33 < 12; fi33++ {
                                                                v34 := []int{(fi33 + (fi33 >> (((v6 | (v6 << (((v6 & 65535)) & 63)))) & 63))), -184, (fi33 | (fi33 & fi33))}
                                                                v7 = v7
                                                                _ = fi33
                                                                _ = v34
                                                            }
                                                        }
                                                        fmt.Println(v7)
                                                        v35 := []int{}
                                                        _ = v35
                                                    } else {
                                                        if true {
                                                            var v36 uint64 = 1
                                                            v37 := GenS3{F0: "hello", F1: 0}
                                                            v38 := GenS4{F0: -780, F1: 42}
                                                            _ = v36
                                                            _ = v37
                                                            _ = v38
                                                        } else {
                                                            switch 2147483647 {
                                                            case -282:
                                                                v39 := GenS5{F0: v1}
                                                                _ = v39
                                                            case 4294967295:
                                                                v40 := GenS4{F0: v2, F1: 107}
                                                                _ = v40
                                                            case -1:
                                                                v41 := GenS6{F0: 948, F1: -2147483648}
                                                                _ = v41
                                                            }
                                                        }
                                                    }
                                                    v42 := GenS5{F0: (v1 << ((113) & 63))}
                                                    v2 -= v2
                                                    _ = v42
                                                }
                                            } else {
                                                var v43 int16 = (v2 << ((v6) & 63))
                                                for fi44 := 0; fi44 < 1; fi44++ {
                                                    var v45 uint16 = v1
                                                    _ = fi44
                                                    _ = v45
                                                }
                                                _ = v43
                                            }
                                        }
                                    }
                                case "foo bar":
                                    var v46 bool = ((!false) || false)
                                    _ = v46
                                default:
                                    if false {
                                        v12 = 2147483647
                                        for fi47 := 0; fi47 < 6; fi47++ {
                                            if ((v4 <= (v4 - v4)) && false) {
                                                v4 += (v4 << ((18446744073709551615) & 63))
                                            }
                                            for fi48 := 0; fi48 < 15; fi48++ {
                                                fmt.Println(v2, v7)
                                                var v49 string = "a"
                                                var v50 uint8 = 127
                                                _ = fi48
                                                _ = v49
                                                _ = v50
                                            }
                                            _ = fi47
                                        }
                                        if ((v1 * (v1 % ((42) | 1))) == 65535) {
                                            v7 = v7
                                        } else {
                                            v1 -= (^(v1))
                                            var v51 float32 = (1.5 + (3.14159 * ((v7 - (v7 * v7)) * v7)))
                                            for fi52 := 0; fi52 < 10; fi52++ {
                                                if ((false || true) || false) {
                                                    for fi53 := 0; fi53 < 9; fi53++ {
                                                        if ((false && true) || (-12 <= 100)) {
                                                            v54 := GenS2{F0: 235, F1: 127, F2: v3}
                                                            if (!true) {
                                                                v4 = (^(v4))
                                                                v7 = v7
                                                                if true {
                                                                    switch "a" {
                                                                    case "hello":
                                                                        var v55 bool = (100 != 0)
                                                                        _ = v55
                                                                    default:
                                                                        v4 += -9223372036854775808
                                                                    }
                                                                    if true {
                                                                        if true {
                                                                            var v56 int = 9223372036854775807
                                                                            switch 546 {
                                                                            case 9223372036854775807:
                                                                                v57 := GenS7{F0: int64(v6), F1: v2, F2: (fi52 - fi53)}
                                                                                _ = v57
                                                                            case -2147483648:
                                                                                fmt.Println(v56, v7, fi53)
                                                                            case 32767:
                                                                                v7 += (v7 * v7)
                                                                            }
                                                                            _ = v56
                                                                        }
                                                                        var v58 int16 = v2
                                                                        _ = v58
                                                                    }
                                                                }
                                                            }
                                                            v2 -= (v2 + -32768)
                                                            _ = v54
                                                        } else {
                                                            if ((v1 | (v1 - (v1 - v1))) > (^(v1))) {
                                                                var v59 int = (fi53 >> ((1) & 63))
                                                                v4 = -187
                                                                _ = v59
                                                            }
                                                            var v60 int8 = 0
                                                            v1 -= uint16(v6)
                                                            _ = v60
                                                        }
                                                        _ = fi53
                                                    }
                                                }
                                                _ = fi52
                                            }
                                            _ = v51
                                        }
                                    }
                                }
                            case "hello":
                                var v61 int16 = (v2 | (^(v2)))
                                _ = v61
                            }
                        case 32767:
                            if (2147483647 > (^(v3))) {
                                var v62 int32 = 2147483647
                                if true {
                                    var v63 float64 = 0.0
                                    v64 := []int{42, 712, 0}
                                    var v65 uint16 = 460
                                    _ = v63
                                    _ = v64
                                    _ = v65
                                } else {
                                    v66 := GenS7{F0: 384, F1: v2, F2: 32768}
                                    v62 = int32(v4)
                                    var v67 int32 = (v62 + v62)
                                    _ = v66
                                    _ = v67
                                }
                                _ = v62
                            }
                        default:
                            for fi68 := 0; fi68 < 6; fi68++ {
                                if ((^(v1)) != v1) {
                                    if false {
                                        for fi69 := 0; fi69 < 16; fi69++ {
                                            if true {
                                                if ((v1 ^ v1) < (^(v1))) {
                                                    v70 := GenS7{F0: -559, F1: v2, F2: -543}
                                                    var v71 float32 = v7
                                                    v72 := GenS8{F0: 317}
                                                    if true {
                                                        if (-420 <= v2) {
                                                            v2 = v2
                                                            fmt.Println(v12, fi69)
                                                            v73 := GenS9{F0: 4294967295, F1: (v2 / (((v2 | v2)) | 1))}
                                                            _ = v73
                                                        } else {
                                                            switch (^(fi69)) {
                                                            case 27:
                                                                v4 -= (v4 ^ 1)
                                                            default:
                                                                v12 = 4294967295
                                                            }
                                                            fmt.Println(v7, v2)
                                                            var v74 uint = (v6 - v6)
                                                            _ = v74
                                                        }
                                                    } else {
                                                        var v75 uint64 = 4294967295
                                                        v3 -= 4294967295
                                                        _ = v75
                                                    }
                                                    fmt.Println(v3, v2, v71)
                                                    if ((^(v4)) != (v4 | (v4 - 239))) {
                                                        delete(v5, (fi69 * 32768))
                                                        if (-904 > -1) {
                                                            fmt.Println(v6, v3)
                                                            v76 := v5[(fi69 >> ((v6) & 63))]
                                                            fmt.Println(v6)
                                                            _ = v76
                                                        }
                                                        v77 := GenS6{F0: 42, F1: 0}
                                                        _ = v77
                                                    }
                                                    _ = v70
                                                    _ = v71
                                                    _ = v72
                                                } else {
                                                    v4 += int64(v2)
                                                }
                                                var v78 int8 = 100
                                                _ = v78
                                            }
                                            fmt.Println(v7, v3, v6)
                                            v79 := GenS10{F0: 128, F1: v4}
                                            _ = fi69
                                            _ = v79
                                        }
                                    } else {
                                        for fi80 := 0; fi80 < 6; fi80++ {
                                            if false {
                                                var v81 int32 = 1
                                                _ = v81
                                            }
                                            v2 = 1000
                                            _ = fi80
                                        }
                                        switch "foo bar" {
                                        case "":
                                            v6 += 340
                                        case "a":
                                            if (!(!false)) {
                                                if (3.14159 < ((1.5 * float64(v7)) + ((1e10 / 1e10) / 1.0))) {
                                                    if (!true) {
                                                        var v82 float64 = (1e10 + 1e10)
                                                        var v83 uint32 = (v12 >> ((v6) & 63))
                                                        _ = v82
                                                        _ = v83
                                                    } else {
                                                        v5[(fi68 >> ((v6) & 63))] = (v3 + v3)
                                                        var v84 uint64 = 473
                                                        _ = v84
                                                    }
                                                    for fi85 := 0; fi85 < 15; fi85++ {
                                                        var v86 bool = (!false)
                                                        if (1 < 128) {
                                                            v12 = v3
                                                            switch fi85 {
                                                            case 828:
                                                                v7 -= (v7 / (-2.25 / v7))
                                                            default:
                                                                var v87 int16 = 110
                                                                _ = v87
                                                            }
                                                        }
                                                        v88 := GenS11{F0: v12, F1: (fi85 - fi85)}
                                                        _ = fi85
                                                        _ = v86
                                                        _ = v88
                                                    }
                                                }
                                                switch (("hello" + "foo bar") + "a") {
                                                case "":
                                                    var v89 int64 = -1
                                                    _ = v89
                                                case "a":
                                                    for fi90 := 0; fi90 < 11; fi90++ {
                                                        if (("hello" + "hello") > "\n\t") {
                                                            v7 = (float32(0.0) / 1.0)
                                                            var v91 int32 = 0
                                                            if (0 == 42) {
                                                                switch fi90 {
                                                                case 9223372036854775807:
                                                                    var v92 float32 = 1e10
                                                                    _ = v92
                                                                case 2147483647:
                                                                    v93 := GenS3{F0: "foo bar", F1: v91}
                                                                    switch fi90 {
                                                                    case 128:
                                                                        var v94 uint32 = 414
                                                                        _ = v94
                                                                    case 32768:
                                                                        var v95 uint = v6
                                                                        _ = v95
                                                                    case 32767:
                                                                        if (!true) {
                                                                            switch (fi68 + (fi68 - fi90)) {
                                                                            case -2147483648:
                                                                                if (786 < 42) {
                                                                                    if (1.5 != 0.0) {
                                                                                        var v96 uint32 = 973
                                                                                        _ = v96
                                                                                    } else {
                                                                                        v97 := []int{42, fi90, 1}
                                                                                        if true {
                                                                                            if len(v97) > 0 {
                                                                                                v97[0] = fi90
                                                                                            }
                                                                                            var v98 uint32 = v12
                                                                                            _ = v98
                                                                                        }
                                                                                        switch "foo bar" {
                                                                                        case "a":
                                                                                            switch fi90 {
                                                                                            case -648:
                                                                                                switch fi68 {
                                                                                                case -9223372036854775808:
                                                                                                    v4 = (v4 ^ (v4 + -845))
                                                                                                case 622:
                                                                                                    fmt.Println(v4, v3)
                                                                                                case -1:
                                                                                                    _ = len(v5)
                                                                                                    v99 := GenS12{F0: v2}
                                                                                                    v100 := GenS1{F0: v2}
                                                                                                    var v101 uint8 = 255
                                                                                                    v102 := GenS8{F0: 393}
                                                                                                    _ = v100
                                                                                                    _ = v101
                                                                                                    _ = v102
                                                                                                    _ = v99
                                                                                                default:
                                                                                                    var v103 uint16 = v1
                                                                                                    _ = v103
                                                                                                }
                                                                                                v4 = 1
                                                                                                switch "foo bar" {
                                                                                                case "a":
                                                                                                    v104 := GenS13{F0: 206, F1: 125}
                                                                                                    _ = v104
                                                                                                default:
                                                                                                    v105 := v5[fi90]
                                                                                                    _ = v105
                                                                                                }
                                                                                            }
                                                                                        case "\n\t":
                                                                                            switch (fi90 >> ((65535) & 63)) {
                                                                                            case -9223372036854775808:
                                                                                                var v106 int32 = 2147483647
                                                                                                _ = v106
                                                                                            case 32768:
                                                                                                if (v4 != 0) {
                                                                                                    var v107 int8 = 100
                                                                                                    _ = v107
                                                                                                } else {
                                                                                                    var v108 uint = 255
                                                                                                    _ = v108
                                                                                                }
                                                                                            default:
                                                                                                v109 := GenS2{F0: 4294967295, F1: 100, F2: (^(v12))}
                                                                                                _ = v109
                                                                                            }
                                                                                        case "foo bar":
                                                                                            if (v1 != v1) {
                                                                                                v5[(^(fi68))] = 2
                                                                                                var v110 int64 = 9223372036854775807
                                                                                                var v111 int = fi68
                                                                                                var v112 float64 = (float64(v7) / (1e10 * 1.5))
                                                                                                var v113 uint16 = 812
                                                                                                fmt.Println(v1)
                                                                                                _ = v110
                                                                                                _ = v111
                                                                                                _ = v112
                                                                                                _ = v113
                                                                                            } else {
                                                                                                var v114 float32 = 3.14159
                                                                                                fmt.Println(fi90)
                                                                                                _ = v114
                                                                                            }
                                                                                            var v115 uint32 = 679
                                                                                            _ = v115
                                                                                        }
                                                                                        var v116 uint32 = v12
                                                                                        if (2147483648 <= v3) {
                                                                                            if false {
                                                                                                if false {
                                                                                                    var v117 string = ("a" + "foo bar")
                                                                                                    v2 = v2
                                                                                                    switch "" {
                                                                                                    case "":
                                                                                                        if true {
                                                                                                            switch fi68 {
                                                                                                            case 4294967295:
                                                                                                                v118 := GenS14{F0: v4}
                                                                                                                _ = v118
                                                                                                            case 32768:
                                                                                                                v116 -= v3
                                                                                                            }
                                                                                                            v119 := GenS1{F0: 1}
                                                                                                            v117 = ((v117 + "a") + v117)
                                                                                                            _ = v119
                                                                                                        } else {
                                                                                                            v3 = v3
                                                                                                            fmt.Println(v4, v12)
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        switch v117 {
                                                                                                        case "hello":
                                                                                                            if ((v7 - 1e10) <= v7) {
                                                                                                                v91 += int32(fi68)
                                                                                                                var v120 int32 = v91
                                                                                                                _ = v120
                                                                                                            }
                                                                                                        case "":
                                                                                                            v121 := GenS5{F0: (^(v1))}
                                                                                                            _ = v121
                                                                                                        case "a":
                                                                                                            v122 := GenS6{F0: 545, F1: 0}
                                                                                                            _ = v122
                                                                                                        default:
                                                                                                            var v123 uint8 = 127
                                                                                                            _ = v123
                                                                                                        }
                                                                                                    case "a":
                                                                                                        var v124 float64 = -2.25
                                                                                                        _ = v124
                                                                                                    }
                                                                                                    _ = v117
                                                                                                }
                                                                                                v3 = v116
                                                                                                v6 = 65535
                                                                                            }
                                                                                        }
                                                                                        _ = v116
                                                                                        _ = v97
                                                                                    }
                                                                                    fmt.Println(v6, fi68, v2)
                                                                                    switch "foo bar" {
                                                                                    case "\n\t":
                                                                                        switch 128 {
                                                                                        case -263:
                                                                                            v2 = (v2 / ((-539) | 1))
                                                                                        case 128:
                                                                                            v125 := GenS7{F0: (v4 - v4), F1: (v2 + v2), F2: -9223372036854775808}
                                                                                            _ = v125
                                                                                        case 401:
                                                                                            var v126 float32 = float32(1e10)
                                                                                            _ = v126
                                                                                        default:
                                                                                            v127 := GenS11{F0: 0, F1: fi90}
                                                                                            _ = v127
                                                                                        }
                                                                                    case "a":
                                                                                        v1 -= (^(v1))
                                                                                    case "hello":
                                                                                        var v128 int32 = -2147483648
                                                                                        _ = v128
                                                                                    }
                                                                                    v3 += v12
                                                                                    v129 := GenS15{F0: 681, F1: (v1 | 0)}
                                                                                    _ = v129
                                                                                } else {
                                                                                    v130 := GenS2{F0: 355, F1: 100, F2: v12}
                                                                                    if (!(!false)) {
                                                                                        v12 = v12
                                                                                    }
                                                                                    fmt.Println(v6)
                                                                                    if (v7 < (v7 * v7)) {
                                                                                        v6 -= v6
                                                                                        var v131 uint64 = 328
                                                                                        _ = v131
                                                                                    }
                                                                                    _ = v130
                                                                                }
                                                                            }
                                                                        } else {
                                                                            var v132 uint = 2
                                                                            _ = v132
                                                                        }
                                                                        switch (fi68 >> (((v6 / ((v6) | 1))) & 63)) {
                                                                        case -9223372036854775808:
                                                                            switch "\n\t" {
                                                                            case "hello":
                                                                                fmt.Println(fi68)
                                                                            default:
                                                                                v5[fi90] = v12
                                                                            }
                                                                        case 32768:
                                                                            var v133 int32 = v91
                                                                            _ = v133
                                                                        case -61:
                                                                            if false {
                                                                                v134 := []int{fi68, 32767, fi68, (-(fi90))}
                                                                                v7 -= -2.25
                                                                                var v135 uint32 = uint32(v6)
                                                                                _ = v134
                                                                                _ = v135
                                                                            } else {
                                                                                v136 := GenS13{F0: 42, F1: 100}
                                                                                var v137 int8 = 1
                                                                                var v138 int = fi90
                                                                                v6 = (v6 | (v6 >> (((v6 % ((0) | 1))) & 63)))
                                                                                fmt.Println(v1)
                                                                                var v139 uint32 = (v12 << ((4294967295) & 63))
                                                                                var v140 int8 = int8(fi90)
                                                                                _ = v136
                                                                                _ = v137
                                                                                _ = v138
                                                                                _ = v139
                                                                                _ = v140
                                                                            }
                                                                        }
                                                                        var v141 uint8 = 0
                                                                        _ = v141
                                                                    }
                                                                    _ = v93
                                                                case 0:
                                                                    var v142 uint = 1
                                                                    _ = v142
                                                                }
                                                                if ((128 != 2) || (!true)) {
                                                                    v6 = v6
                                                                    v143 := GenS16{F0: v91, F1: v2, F2: fi90}
                                                                    v1 = 42
                                                                    _ = v143
                                                                }
                                                            } else {
                                                                var v144 int16 = v2
                                                                _ = len(v5)
                                                                fmt.Println(v4, fi90)
                                                                _ = v144
                                                            }
                                                            _ = v91
                                                        }
                                                        _ = fi90
                                                    }
                                                case "foo bar":
                                                    v145 := []int{(fi68 * 9223372036854775807)}
                                                    _ = v145
                                                }
                                                if (((v4 ^ (v4 % (((v4 * v4)) | 1))) < v4) && true) {
                                                    var v146 uint64 = 4294967295
                                                    fmt.Println(v6, v2, v4)
                                                    v147 := GenS17{F0: (v1 >> ((42) & 63)), F1: ((v1 / ((v1) | 1)) != v1), F2: 55}
                                                    _ = v146
                                                    _ = v147
                                                } else {
                                                    v2 += (v2 << ((479) & 63))
                                                    v148 := GenS18{F0: 964, F1: 42}
                                                    v149 := GenS19{F0: 128, F1: v2, F2: 42}
                                                    _ = v148
                                                    _ = v149
                                                }
                                            } else {
                                                v150 := []int{}
                                                _ = v150
                                            }
                                        }
                                    }
                                } else {
                                    if (230 <= 1) {
                                        var v151 uint16 = 42
                                        v1 -= (v1 * v151)
                                        _ = v151
                                    } else {
                                        var v152 int64 = 1
                                        switch fi68 {
                                        case -2147483648:
                                            fmt.Println(fi68, v7, v12)
                                        }
                                        _ = v152
                                    }
                                }
                                _ = fi68
                            }
                        }
                    default:
                        if (4294967295 == 18446744073709551615) {
                            fmt.Println(v4)
                        }
                    }
                    v153 := GenS6{F0: 18446744073709551615, F1: 1}
                    fmt.Println(v2)
                    _ = v153
                }
                _ = v12
            } else {
                _ = len(v5)
                _ = len(v5)
                var v154 int32 = -2147483648
                _ = v154
            }
            fmt.Println(v6, v1, v7)
        } else {
            var v155 int64 = -1
            switch "\n\t" {
            case "foo bar":
                v156 := []int{32767, 1, -2147483648, 128}
                _ = v156
            default:
                v7 = (v7 + v7)
            }
            _ = v155
        }
        fmt.Println(v2, v2)
    } else {
        for fi157 := 0; fi157 < 10; fi157++ {
            var v158 uint16 = (v1 + (v1 % (((^(v1))) | 1)))
            v4 = v4
            v7 += (v7 - (((v7 + 3.14159) / ((1.5 * (1.5 - (3.14159 - (0.0 + 3.14159)))) / ((v7 - v7) + (float32((0.1 - 3.14159)) - (v7 / v7))))) / 1e10))
            _ = fi157
            _ = v158
        }
        v2 -= (v2 + -453)
    }
    switch "a" {
    case "":
        if false {
            var v159 uint8 = 2
            if (1e10 > -2.25) {
                var v160 uint64 = 1
                _ = v160
            }
            fmt.Println(v4, v2)
            _ = v159
        } else {
            switch 32768 {
            case 4294967295:
                v2 = (v2 % (((v2 >> ((v6) & 63))) | 1))
            case 53:
                var v161 string = ("\n\t" + ("a" + ""))
                _ = v161
            case 127:
                v1 = 42
            default:
                if (-9223372036854775808 <= 32767) {
                    var v162 int64 = (^(v4))
                    switch 9223372036854775807 {
                    case 32768:
                        var v163 int8 = 0
                        _ = v163
                    default:
                        v164 := GenS2{F0: 18446744073709551615, F1: 42, F2: v3}
                        _ = v164
                    }
                    _ = v162
                } else {
                    var v165 uint32 = (v3 >> (((v6 & v6)) & 63))
                    _ = v165
                }
            }
        }
    default:
        var v166 uint32 = 573
        _ = v166
    }
    v167 := GenS4{F0: -32768, F1: 255}
    v168 := GenS7{F0: (v4 & (v4 + (v4 / ((v4) | 1)))), F1: (v2 / (((v2 - v2)) | 1)), F2: 1}
    _ = len(v5)
    v6 = 73
    var v169 int64 = v4
    v169 += 659
    if true {
        for fi170 := 0; fi170 < 5; fi170++ {
            if ((v3 & v3) == v3) {
                if false {
                    v171 := []int{(fi170 | fi170), int(v1), fi170}
                    var v172 int16 = 0
                    _ = len(v5)
                    _ = v171
                    _ = v172
                }
            }
            v4 = (^(v169))
            _ = fi170
        }
        for fi173 := 0; fi173 < 10; fi173++ {
            if false {
                v1 += v1
            }
            if (!true) {
                var v174 int8 = 100
                v169 = (v169 / ((-1) | 1))
                v7 = v7
                _ = v174
            }
            _ = fi173
        }
    }
    var v175 int16 = (v2 / ((v2) | 1))
    v176 := GenS14{F0: 42}
    var v177 uint = 931
    fmt.Println(v1, v2, v3)
    var v178 int16 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v169:", v169)
    fmt.Println("v175:", v175)
    fmt.Println("v177:", v177)
    fmt.Println("v178:", v178)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v167
    _ = v168
    _ = v169
    _ = v175
    _ = v176
    _ = v177
    _ = v178
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
