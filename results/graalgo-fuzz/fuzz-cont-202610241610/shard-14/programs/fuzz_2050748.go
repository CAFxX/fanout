package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 bool }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 int64; F2 bool }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS10 struct { F0 int16; F1 int8 }
type GenS11 struct { F0 uint; F1 int64; F2 int }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int16; F1 uint32; F2 uint }
type GenS14 struct { F0 int8; F1 int32 }
type GenS15 struct { F0 int8; F1 int32 }
type GenS16 struct { F0 uint32 }
type GenS17 struct { F0 int32; F1 int8 }
type GenS18 struct { F0 uint8 }
type GenS19 struct { F0 bool; F1 int; F2 bool }
type GenS20 struct { F0 int }
type GenS21 struct { F0 string; F1 uint8; F2 int }
type GenS22 struct { F0 uint64 }
type GenS23 struct { F0 uint64; F1 uint8 }
type GenS24 struct { F0 uint32 }
type GenS25 struct { F0 int }
type GenS26 struct { F0 int16; F1 bool }
type GenS27 struct { F0 uint32; F1 int32 }
type GenS28 struct { F0 int64; F1 string }
type GenS29 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS30 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS31 struct { F0 uint8 }
type GenS32 struct { F0 uint32; F1 bool; F2 string }
type GenS33 struct { F0 bool }
type GenS34 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS35 struct { F0 bool; F1 int32 }
type GenS36 struct { F0 int32; F1 int16; F2 uint16 }
type GenS37 struct { F0 uint16; F1 uint }
type GenS38 struct { F0 int32; F1 int16 }
type GenS39 struct { F0 uint; F1 int; F2 uint32 }
type GenS40 struct { F0 int16 }
type GenS41 struct { F0 int64 }
type GenS42 struct { F0 int64; F1 int32; F2 int8 }
type GenS43 struct { F0 uint64; F1 int8 }
func genFn0() int8 {
    if false {
        v7 := []int{418, -1}
        var v8 int8 = -128
        _ = v7
        _ = v8
    }
    if (true && (1 < 0)) {
        for fi9 := 0; fi9 < 14; fi9++ {
            v10 := GenS1{F0: (0.1 < (-2.25 + -2.25)), F1: 42, F2: (0.1 != -2.25)}
            if (4294967295 >= 4294967295) {
                if ((float64(3.14159) == 1e10) || false) {
                    if (!false) {
                        var v11 bool = false
                        v11 = (0.1 <= float32(0.0))
                        fmt.Println(fi9, v11)
                        _ = v11
                    } else {
                        switch "a" {
                        case "a":
                            v12 := []int{}
                            _ = v12
                        case "hello":
                            v13 := GenS2{F0: 2, F1: 2}
                            _ = v13
                        }
                        if (!true) {
                            if true {
                                var v14 int = (fi9 + 42)
                                v15 := GenS3{F0: 732, F1: 42}
                                v16 := GenS4{F0: 65535}
                                _ = v14
                                _ = v15
                                _ = v16
                            }
                            if true {
                                v17 := GenS2{F0: 0, F1: 42}
                                _ = v17
                            }
                            for fi18 := 0; fi18 < 6; fi18++ {
                                fmt.Println(fi9)
                                var v19 string = ("" + "hello")
                                _ = fi18
                                _ = v19
                            }
                        }
                    }
                }
            } else {
                for fi20 := 0; fi20 < 5; fi20++ {
                    var v21 uint16 = 255
                    var v22 float32 = float32((1e10 - float64(0.1)))
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                switch ((("\n\t" + "") + "\n\t") + "foo bar") {
                case "hello":
                    v23 := GenS3{F0: 42, F1: 0}
                    _ = v23
                default:
                    if (-2147483648 < -629) {
                        if (true || (501 <= 0)) {
                            v24 := []int{fi9, fi9, (fi9 % ((fi9) | 1))}
                            switch (("" + "foo bar") + "a") {
                            case "hello":
                                switch "hello" {
                                case "":
                                    fmt.Println(fi9)
                                }
                            case "\n\t":
                                var v25 int8 = -128
                                _ = v25
                            default:
                                for fi26 := 0; fi26 < 15; fi26++ {
                                    var v27 int8 = -52
                                    switch "foo bar" {
                                    case "":
                                        v27 = (v27 << ((477) & 63))
                                    default:
                                        var v28 uint16 = 2
                                        _ = v28
                                    }
                                    var v29 uint16 = 808
                                    _ = fi26
                                    _ = v27
                                    _ = v29
                                }
                            }
                            _ = v24
                        } else {
                            var v30 int32 = 0
                            if (4294967295 == 0) {
                                fmt.Println(fi9, fi9)
                            } else {
                                v30 -= (v30 | (v30 - (v30 * v30)))
                                var v31 int16 = 42
                                v32 := GenS5{F0: 2, F1: 9223372036854775807, F2: true}
                                _ = v31
                                _ = v32
                            }
                            _ = v30
                        }
                    } else {
                        var v33 uint32 = 2
                        v34 := GenS5{F0: 32767, F1: 1, F2: (true || false)}
                        if false {
                            v33 = v33
                            v35 := GenS2{F0: 255, F1: 127}
                            for fi36 := 0; fi36 < 20; fi36++ {
                                v33 = (^(v33))
                                v37 := GenS5{F0: 255, F1: 9223372036854775807, F2: (!false)}
                                switch "\n\t" {
                                case "foo bar":
                                    v33 += 1
                                case "":
                                    v38 := []int{}
                                    _ = v38
                                default:
                                    var v39 uint16 = 255
                                    _ = v39
                                }
                                _ = fi36
                                _ = v37
                            }
                            _ = v35
                        }
                        _ = v33
                        _ = v34
                    }
                }
            }
            _ = fi9
            _ = v10
        }
        var v40 int32 = 42
        v41 := []int{}
        _ = v40
        _ = v41
    } else {
        if false {
            for fi42 := 0; fi42 < 17; fi42++ {
                var v43 uint64 = 1
                if false {
                    v44 := GenS6{F0: 1, F1: 255}
                    v45 := []int{}
                    _ = v44
                    _ = v45
                }
                _ = fi42
                _ = v43
            }
            v46 := GenS7{F0: 2147483647}
            _ = v46
        } else {
            if true {
                if false {
                    var v47 uint64 = 42
                    var v48 int64 = 78
                    if (true && (!(-1 == 1))) {
                        var v49 int64 = v48
                        if (float32(-2.25) < float32((0.0 - -2.25))) {
                            var v50 int64 = (v48 ^ v49)
                            if true {
                                var v51 float64 = (3.14159 + 1.5)
                                fmt.Println(v48)
                                _ = v51
                            } else {
                                switch ("a" + "hello") {
                                case "foo bar":
                                    for fi52 := 0; fi52 < 16; fi52++ {
                                        for fi53 := 0; fi53 < 7; fi53++ {
                                            var v54 int16 = -174
                                            switch ("hello" + (("a" + ("" + "a")) + "hello")) {
                                            case "foo bar":
                                                v55 := GenS8{F0: "foo bar"}
                                                _ = v55
                                            case "a":
                                                var v56 float32 = ((-2.25 * float32(1e10)) / 1.5)
                                                _ = v56
                                            }
                                            fmt.Println(v54)
                                            if true {
                                                var v57 uint32 = 877
                                                var v58 uint16 = 90
                                                _ = v57
                                                _ = v58
                                            }
                                            _ = fi53
                                            _ = v54
                                        }
                                        for fi59 := 0; fi59 < 7; fi59++ {
                                            if (((0.1 / 1.0) / 1e10) > 1.5) {
                                                v60 := []int{}
                                                var v61 uint8 = 42
                                                if len(v60) > 0 {
                                                    v60[0] = fi59
                                                }
                                                _ = v60
                                                _ = v61
                                            }
                                            _ = fi59
                                        }
                                        var v62 float32 = ((float32(1e10) + -2.25) + 1e10)
                                        _ = fi52
                                        _ = v62
                                    }
                                case "a":
                                    for fi63 := 0; fi63 < 5; fi63++ {
                                        if ((32767 > fi63) && (!false)) {
                                            v50 -= (^(v48))
                                            var v64 int16 = -332
                                            for fi65 := 0; fi65 < 6; fi65++ {
                                                var v66 int64 = v48
                                                _ = fi65
                                                _ = v66
                                            }
                                            _ = v64
                                        }
                                        switch ("\n\t" + "") {
                                        case "\n\t":
                                            for fi67 := 0; fi67 < 5; fi67++ {
                                                if false {
                                                    switch ("\n\t" + "\n\t") {
                                                    case "a":
                                                        fmt.Println(v47, v47)
                                                    case "hello":
                                                        var v68 string = ("hello" + "a")
                                                        _ = v68
                                                    }
                                                    var v69 uint8 = 88
                                                    var v70 int8 = -1
                                                    _ = v69
                                                    _ = v70
                                                }
                                                _ = fi67
                                            }
                                        case "hello":
                                            if (-1 < 43) {
                                                if false {
                                                    for fi71 := 0; fi71 < 2; fi71++ {
                                                        v50 = 9223372036854775807
                                                        _ = fi71
                                                    }
                                                }
                                                v72 := []int{}
                                                _ = v72
                                            }
                                        case "foo bar":
                                            switch (fi63 * -430) {
                                            case -552:
                                                var v73 uint32 = 2147483647
                                                _ = v73
                                            case -328:
                                                for fi74 := 0; fi74 < 20; fi74++ {
                                                    v75 := []int{}
                                                    v50 = v49
                                                    _ = fi74
                                                    _ = v75
                                                }
                                            default:
                                                var v76 int = (fi63 * 1)
                                                _ = v76
                                            }
                                        default:
                                            switch (fi63 >> ((1) & 63)) {
                                            case 543:
                                                var v77 int64 = v48
                                                _ = v77
                                            }
                                        }
                                        v47 += (v47 % ((v47) | 1))
                                        _ = fi63
                                    }
                                }
                                v78 := GenS9{F0: 2, F1: 0, F2: 100}
                                v79 := []int{-1, 127, 240}
                                _ = v78
                                _ = v79
                            }
                            for fi80 := 0; fi80 < 6; fi80++ {
                                if false {
                                    for fi81 := 0; fi81 < 18; fi81++ {
                                        if false {
                                            if (!("foo bar" >= ((((("foo bar" + "") + ("" + "")) + "") + "foo bar") + "hello"))) {
                                                switch (-(fi80)) {
                                                case 127:
                                                    v50 -= v48
                                                case 4294967295:
                                                    switch (fi81 * (fi81 << ((924) & 63))) {
                                                    case -2147483648:
                                                        if false {
                                                            switch (fi81 ^ (fi80 << ((0) & 63))) {
                                                            case -217:
                                                                var v82 uint64 = v47
                                                                _ = v82
                                                            case 1:
                                                                if (639 != -32768) {
                                                                    v83 := GenS7{F0: fi80}
                                                                    fmt.Println(fi80)
                                                                    v48 = v49
                                                                    _ = v83
                                                                } else {
                                                                    var v84 int32 = 953
                                                                    v50 = (v48 ^ (^(v49)))
                                                                    _ = v84
                                                                }
                                                            }
                                                            v49 -= (v48 & v48)
                                                        }
                                                        fmt.Println(v47, v48, v47)
                                                        var v85 uint32 = 65535
                                                        switch fi81 {
                                                        case 639:
                                                            var v86 float32 = (3.14159 * 0.1)
                                                            _ = v86
                                                        case -9223372036854775808:
                                                            v87 := GenS7{F0: (fi81 - 32768)}
                                                            _ = v87
                                                        }
                                                        if (fi80 >= fi80) {
                                                            v47 = 18446744073709551615
                                                            var v88 uint8 = 127
                                                            fmt.Println(fi80, v49, v49)
                                                            _ = v88
                                                        }
                                                        fmt.Println(v50, v50)
                                                        v89 := GenS2{F0: 2, F1: 0}
                                                        _ = v85
                                                        _ = v89
                                                    default:
                                                        var v90 int16 = 32767
                                                        _ = v90
                                                    }
                                                case -9223372036854775808:
                                                    if true {
                                                        if (589 > (fi80 << ((1) & 63))) {
                                                            v91 := []int{32767, fi80}
                                                            _ = v91
                                                        }
                                                    }
                                                }
                                                v92 := GenS4{F0: 0}
                                                switch 4294967295 {
                                                case -442:
                                                    v50 -= 42
                                                case 0:
                                                    fmt.Println(v50, v49, v50)
                                                    if (false || true) {
                                                        v50 = (v49 + (v49 ^ 9223372036854775807))
                                                        var v93 int8 = -128
                                                        v47 += (v47 + v47)
                                                        _ = v93
                                                    }
                                                    var v94 float64 = (0.0 / (1.5 + 3.14159))
                                                    _ = v94
                                                case -1:
                                                    if (4294967295 != 1) {
                                                        v50 -= (-(v50))
                                                        var v95 float64 = 1.5
                                                        v49 = v48
                                                        _ = v95
                                                    } else {
                                                        switch "a" {
                                                        case "":
                                                            var v96 uint8 = 0
                                                            _ = v96
                                                        case "\n\t":
                                                            if (-120 < 1000) {
                                                                if (false || (86 != (^(v49)))) {
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        v97 := []int{fi80, 32767}
                                                                        _ = v97
                                                                    }
                                                                    v98 := []int{fi80, (fi81 / ((-2147483648) | 1))}
                                                                    if false {
                                                                        fmt.Println(v47, fi81, fi81)
                                                                        switch ("a" + (("" + "\n\t") + ("a" + "foo bar"))) {
                                                                        case "\n\t":
                                                                            v99 := GenS7{F0: fi81}
                                                                            _ = v99
                                                                        case "foo bar":
                                                                            v98 = append(v98, 32767)
                                                                        case "hello":
                                                                            if len(v98) >= 2 {
                                                                                v98 = v98[1:len(v98)]
                                                                            }
                                                                        default:
                                                                            switch "" {
                                                                            case "a":
                                                                                var v100 uint64 = v47
                                                                                _ = v100
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v98 = append(v98, 2147483647)
                                                                        if (631 == 552) {
                                                                            _ = len(v98)
                                                                        } else {
                                                                            if ((35 != v47) || true) {
                                                                                fmt.Println(fi80, fi81, v49)
                                                                                if false {
                                                                                    if true {
                                                                                        _ = len(v98)
                                                                                        if (true && true) {
                                                                                            v49 = -9223372036854775808
                                                                                            var v101 int8 = 127
                                                                                            if false {
                                                                                                switch "\n\t" {
                                                                                                case "":
                                                                                                    var v102 int16 = 1
                                                                                                    _ = v102
                                                                                                default:
                                                                                                    if true {
                                                                                                        var v103 int64 = 1
                                                                                                        _ = v103
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            if len(v98) >= 2 {
                                                                                                v98 = v98[1:len(v98)]
                                                                                            }
                                                                                            if true {
                                                                                                var v104 uint = 255
                                                                                                _ = v104
                                                                                            } else {
                                                                                                if (-2147483648 <= (fi81 - 32767)) {
                                                                                                    v105 := GenS10{F0: -618, F1: (v101 << ((1) & 63))}
                                                                                                    var v106 float32 = (float32((-2.25 / 1e10)) * 0.1)
                                                                                                    _ = v105
                                                                                                    _ = v106
                                                                                                } else {
                                                                                                    var v107 uint16 = 65535
                                                                                                    var v108 int8 = (^(v101))
                                                                                                    switch (fi81 * 42) {
                                                                                                    case 127:
                                                                                                        if true {
                                                                                                            switch -977 {
                                                                                                            case -763:
                                                                                                                var v109 int = fi81
                                                                                                                _ = v109
                                                                                                            case 42:
                                                                                                                _ = len(v98)
                                                                                                            case 425:
                                                                                                                var v110 bool = false
                                                                                                                _ = v110
                                                                                                            default:
                                                                                                                var v111 float32 = 0.1
                                                                                                                if false {
                                                                                                                    v112 := GenS5{F0: v107, F1: (v50 % (((v50 % ((v48) | 1))) | 1)), F2: (!(0.1 != 0.0))}
                                                                                                                    v108 += (v108 & (v101 << ((65535) & 63)))
                                                                                                                    _ = len(v98)
                                                                                                                    v101 = (v108 * v108)
                                                                                                                    _ = v112
                                                                                                                }
                                                                                                                v113 := GenS11{F0: 42, F1: (^(v49)), F2: fi80}
                                                                                                                _ = v111
                                                                                                                _ = v113
                                                                                                            }
                                                                                                        } else {
                                                                                                            v47 -= v47
                                                                                                            var v114 int16 = 32767
                                                                                                            _ = v114
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v107
                                                                                                    _ = v108
                                                                                                }
                                                                                                switch 2147483647 {
                                                                                                case -106:
                                                                                                    v49 += v48
                                                                                                case 32767:
                                                                                                    var v115 int16 = 666
                                                                                                    _ = v115
                                                                                                default:
                                                                                                    var v116 int8 = v101
                                                                                                    _ = v116
                                                                                                }
                                                                                                var v117 uint16 = 494
                                                                                                _ = v117
                                                                                            }
                                                                                            _ = v101
                                                                                        }
                                                                                        fmt.Println(fi81, v47)
                                                                                    }
                                                                                    v98 = append(v98, (fi80 / (((fi81 - (^(fi80)))) | 1)))
                                                                                    v118 := GenS12{F0: (fi80 + 0)}
                                                                                    _ = v118
                                                                                } else {
                                                                                    v119 := GenS13{F0: -32768, F1: 2147483647, F2: 2}
                                                                                    _ = v119
                                                                                }
                                                                                var v120 uint8 = 2
                                                                                v121 := GenS14{F0: -14, F1: -2147483648}
                                                                                if len(v98) > 0 {
                                                                                    v98[0] = fi81
                                                                                }
                                                                                if true {
                                                                                    var v122 int8 = 38
                                                                                    v123 := GenS15{F0: v122, F1: -2147483648}
                                                                                    if true {
                                                                                        var v124 float64 = -2.25
                                                                                        _ = v124
                                                                                    } else {
                                                                                        v125 := GenS2{F0: 4294967295, F1: 116}
                                                                                        var v126 float32 = (3.14159 + 1e10)
                                                                                        switch "hello" {
                                                                                        case "hello":
                                                                                            v126 = v126
                                                                                        }
                                                                                        _ = v125
                                                                                        _ = v126
                                                                                    }
                                                                                    _ = v122
                                                                                    _ = v123
                                                                                }
                                                                                _ = v120
                                                                                _ = v121
                                                                            } else {
                                                                                if ((-(v48)) < v49) {
                                                                                    v50 = v50
                                                                                    switch "\n\t" {
                                                                                    case "a":
                                                                                        var v127 uint64 = 0
                                                                                        _ = v127
                                                                                    default:
                                                                                        var v128 int32 = -2147483648
                                                                                        _ = v128
                                                                                    }
                                                                                } else {
                                                                                    if false {
                                                                                        fmt.Println(v48, v49, fi81)
                                                                                    } else {
                                                                                        var v129 bool = (-1 <= 169)
                                                                                        v130 := GenS16{F0: 2147483648}
                                                                                        _ = v129
                                                                                        _ = v130
                                                                                    }
                                                                                    v131 := GenS17{F0: -1, F1: 1}
                                                                                    _ = v131
                                                                                }
                                                                                if false {
                                                                                    v48 += (^(v48))
                                                                                } else {
                                                                                    if ((1e10 / 1.0) <= (-2.25 + (0.0 / 1e10))) {
                                                                                        v132 := GenS18{F0: 229}
                                                                                        switch (fi81 | (fi80 % ((fi80) | 1))) {
                                                                                        case 2147483647:
                                                                                            switch fi80 {
                                                                                            case -860:
                                                                                                switch 898 {
                                                                                                case -1:
                                                                                                    if len(v98) > 0 {
                                                                                                        v98[0] = int(v47)
                                                                                                    }
                                                                                                case 2147483647:
                                                                                                    if len(v98) > 0 {
                                                                                                        v98[0] = 4294967295
                                                                                                    }
                                                                                                    if ((true || ((9223372036854775807 != (v48 / ((-9223372036854775808) | 1))) || (1.5 > 1.5))) || false) {
                                                                                                        if len(v98) > 0 {
                                                                                                            v98[0] = fi81
                                                                                                        }
                                                                                                        if len(v98) > 0 {
                                                                                                            v98[0] = (fi81 + fi80)
                                                                                                        }
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            fmt.Println(v48)
                                                                                                            switch (fi80 + fi81) {
                                                                                                            case 4294967295:
                                                                                                                if len(v98) > 0 {
                                                                                                                    v98[0] = 4294967295
                                                                                                                }
                                                                                                            case 5:
                                                                                                                if (true || false) {
                                                                                                                    v48 += v48
                                                                                                                    v133 := GenS18{F0: 2}
                                                                                                                    _ = v133
                                                                                                                } else {
                                                                                                                    _ = len(v98)
                                                                                                                    fmt.Println(v48, v48)
                                                                                                                    var v134 uint64 = v47
                                                                                                                    var v135 float64 = -2.25
                                                                                                                    _ = v134
                                                                                                                    _ = v135
                                                                                                                }
                                                                                                            }
                                                                                                            if true {
                                                                                                                if (-639 <= -1) {
                                                                                                                    var v136 int8 = 127
                                                                                                                    v137 := GenS14{F0: (v136 >> ((817) & 63)), F1: -824}
                                                                                                                    _ = v136
                                                                                                                    _ = v137
                                                                                                                }
                                                                                                            }
                                                                                                        } else {
                                                                                                            switch "" {
                                                                                                            case "a":
                                                                                                                switch fi80 {
                                                                                                                case -525:
                                                                                                                    if len(v98) > 0 {
                                                                                                                        v98[0] = (-(fi81))
                                                                                                                    }
                                                                                                                case 127:
                                                                                                                    var v138 int8 = 0
                                                                                                                    _ = v138
                                                                                                                default:
                                                                                                                    var v139 string = "\n\t"
                                                                                                                    _ = v139
                                                                                                                }
                                                                                                            default:
                                                                                                                v48 -= (v49 % ((v49) | 1))
                                                                                                            }
                                                                                                            switch fi80 {
                                                                                                            case -427:
                                                                                                                if (-2147483648 == 209) {
                                                                                                                    var v140 int32 = 2147483647
                                                                                                                    switch "hello" {
                                                                                                                    case "":
                                                                                                                        switch -905 {
                                                                                                                        case 4294967295:
                                                                                                                            var v141 int = -1
                                                                                                                            _ = v141
                                                                                                                        case 1:
                                                                                                                            switch (fi80 & fi81) {
                                                                                                                            case -2147483648:
                                                                                                                                switch "foo bar" {
                                                                                                                                case "foo bar":
                                                                                                                                    var v142 string = ("foo bar" + "\n\t")
                                                                                                                                    _ = v142
                                                                                                                                case "a":
                                                                                                                                    var v143 int = fi80
                                                                                                                                    _ = v143
                                                                                                                                case "\n\t":
                                                                                                                                    fmt.Println(fi80)
                                                                                                                                    switch ("" + "a") {
                                                                                                                                    case "foo bar":
                                                                                                                                        var v144 int32 = 2147483647
                                                                                                                                        _ = v144
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            case 1:
                                                                                                                                if false {
                                                                                                                                    v145 := GenS10{F0: 32767, F1: 61}
                                                                                                                                    var v146 uint8 = 255
                                                                                                                                    if len(v98) > 0 {
                                                                                                                                        v98[0] = 4294967295
                                                                                                                                    }
                                                                                                                                    _ = v145
                                                                                                                                    _ = v146
                                                                                                                                } else {
                                                                                                                                    v147 := GenS19{F0: true, F1: fi81, F2: (79 < 255)}
                                                                                                                                    fmt.Println(v140, fi81)
                                                                                                                                    _ = v147
                                                                                                                                }
                                                                                                                            case 2147483647:
                                                                                                                                var v148 float64 = 0.1
                                                                                                                                _ = v148
                                                                                                                            }
                                                                                                                        case -2147483648:
                                                                                                                            var v149 float32 = (-2.25 + (3.14159 / 1e10))
                                                                                                                            _ = v149
                                                                                                                        default:
                                                                                                                            var v150 uint32 = uint32(v47)
                                                                                                                            _ = v150
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        switch ("" + "") {
                                                                                                                        case "foo bar":
                                                                                                                            var v151 uint32 = 1
                                                                                                                            _ = v151
                                                                                                                        }
                                                                                                                        switch fi81 {
                                                                                                                        case 32767:
                                                                                                                            var v152 int8 = 127
                                                                                                                            _ = v152
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v140
                                                                                                                } else {
                                                                                                                    var v153 int8 = 1
                                                                                                                    v154 := GenS8{F0: "a"}
                                                                                                                    _ = v153
                                                                                                                    _ = v154
                                                                                                                }
                                                                                                            case 1:
                                                                                                                v49 = 9223372036854775807
                                                                                                            }
                                                                                                        }
                                                                                                        var v155 float64 = ((1.5 + 0.1) * -2.25)
                                                                                                        _ = v155
                                                                                                    }
                                                                                                    var v156 uint64 = v47
                                                                                                    _ = v156
                                                                                                case -461:
                                                                                                    if ((^(v47)) < 18446744073709551615) {
                                                                                                        v47 += v47
                                                                                                        v49 = (v48 * (^(v50)))
                                                                                                        switch (^(fi81)) {
                                                                                                        case 127:
                                                                                                            v49 = (v48 * v50)
                                                                                                        case 2147483647:
                                                                                                            var v157 uint32 = 65535
                                                                                                            _ = v157
                                                                                                        default:
                                                                                                            if len(v98) > 0 {
                                                                                                                v98[0] = (fi81 % (((fi80 % ((fi80) | 1))) | 1))
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    if false {
                                                                                                        v49 = (-(v50))
                                                                                                        v47 += v47
                                                                                                        v50 = 42
                                                                                                        if true {
                                                                                                            var v158 float64 = 1.5
                                                                                                            if true {
                                                                                                                switch fi81 {
                                                                                                                case -9223372036854775808:
                                                                                                                    var v159 float32 = -2.25
                                                                                                                    _ = v159
                                                                                                                default:
                                                                                                                    var v160 bool = (633 != 1)
                                                                                                                    _ = v160
                                                                                                                }
                                                                                                                if len(v98) > 0 {
                                                                                                                    v98[0] = 376
                                                                                                                }
                                                                                                                if true {
                                                                                                                    if false {
                                                                                                                        fmt.Println(fi81, fi80)
                                                                                                                        fmt.Println(v49, fi80)
                                                                                                                    }
                                                                                                                    var v161 uint64 = 18446744073709551615
                                                                                                                    if false {
                                                                                                                        var v162 float64 = (v158 - 1e10)
                                                                                                                        _ = v162
                                                                                                                    }
                                                                                                                    v163 := GenS20{F0: (fi80 | fi81)}
                                                                                                                    _ = v161
                                                                                                                    _ = v163
                                                                                                                } else {
                                                                                                                    if true {
                                                                                                                        if len(v98) > 0 {
                                                                                                                            v98[0] = (fi81 * (fi81 % ((fi81) | 1)))
                                                                                                                        }
                                                                                                                        var v164 uint32 = 4294967295
                                                                                                                        if false {
                                                                                                                            fmt.Println(v49)
                                                                                                                        }
                                                                                                                        _ = v164
                                                                                                                    }
                                                                                                                }
                                                                                                                var v165 bool = (!(23 < 212))
                                                                                                                v50 = (v49 * (-(v49)))
                                                                                                                fmt.Println(fi80)
                                                                                                                _ = v165
                                                                                                            } else {
                                                                                                                v166 := GenS21{F0: ("\n\t" + "a"), F1: 255, F2: fi81}
                                                                                                                var v167 uint8 = 127
                                                                                                                switch "" {
                                                                                                                case "hello":
                                                                                                                    if true {
                                                                                                                        _ = len(v98)
                                                                                                                        fmt.Println(v49)
                                                                                                                        switch "" {
                                                                                                                        case "":
                                                                                                                            fmt.Println(fi81, fi80, fi81)
                                                                                                                        case "\n\t":
                                                                                                                            v168 := GenS4{F0: 65535}
                                                                                                                            _ = v168
                                                                                                                        case "hello":
                                                                                                                            v47 = 2
                                                                                                                        default:
                                                                                                                            var v169 string = "hello"
                                                                                                                            _ = v169
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        var v170 uint8 = v167
                                                                                                                        var v171 int8 = 100
                                                                                                                        _ = v170
                                                                                                                        _ = v171
                                                                                                                    }
                                                                                                                case "a":
                                                                                                                    if (!true) {
                                                                                                                        v172 := GenS14{F0: -72, F1: -2147483648}
                                                                                                                        v158 = v158
                                                                                                                        var v173 float32 = ((0.0 / -2.25) - 1e10)
                                                                                                                        _ = v172
                                                                                                                        _ = v173
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v166
                                                                                                                _ = v167
                                                                                                            }
                                                                                                            var v174 float32 = 0.0
                                                                                                            _ = v158
                                                                                                            _ = v174
                                                                                                        }
                                                                                                        _ = len(v98)
                                                                                                        var v175 float64 = (float64(0.1) - (3.14159 * 0.1))
                                                                                                        v49 += 42
                                                                                                        var v176 uint = 255
                                                                                                        switch (("foo bar" + "") + ("\n\t" + "\n\t")) {
                                                                                                        case "hello":
                                                                                                            v49 = (v49 >> (((v176 / (((v176 * v176)) | 1))) & 63))
                                                                                                        default:
                                                                                                            var v177 int8 = 95
                                                                                                            _ = v177
                                                                                                        }
                                                                                                        _ = v175
                                                                                                        _ = v176
                                                                                                    }
                                                                                                }
                                                                                            case -129:
                                                                                                var v178 uint8 = 128
                                                                                                _ = v178
                                                                                            case 42:
                                                                                                var v179 int8 = 100
                                                                                                _ = v179
                                                                                            default:
                                                                                                v98 = append(v98, fi81)
                                                                                            }
                                                                                            v98 = append(v98, (^(fi81)))
                                                                                            fmt.Println(fi81, v48, v47)
                                                                                            switch fi81 {
                                                                                            case 2147483647:
                                                                                                var v180 float64 = 1e10
                                                                                                _ = v180
                                                                                            case 4294967295:
                                                                                                if len(v98) >= 2 {
                                                                                                    v98 = v98[1:len(v98)]
                                                                                                }
                                                                                            default:
                                                                                                var v181 float64 = (0.1 + -2.25)
                                                                                                _ = v181
                                                                                            }
                                                                                            var v182 uint32 = 2147483648
                                                                                            _ = v182
                                                                                        case 81:
                                                                                            var v183 uint64 = v47
                                                                                            _ = v183
                                                                                        case 1:
                                                                                            var v184 uint64 = v47
                                                                                            _ = v184
                                                                                        }
                                                                                        var v185 int16 = 1000
                                                                                        _ = v132
                                                                                        _ = v185
                                                                                    } else {
                                                                                        if (!true) {
                                                                                            var v186 int64 = v50
                                                                                            var v187 int8 = 100
                                                                                            var v188 int16 = 32767
                                                                                            if ((0.0 / 1.0) < 0.1) {
                                                                                                v50 -= (^(v186))
                                                                                                v50 -= v186
                                                                                            }
                                                                                            _ = v186
                                                                                            _ = v187
                                                                                            _ = v188
                                                                                        } else {
                                                                                            var v189 uint16 = 32767
                                                                                            _ = v189
                                                                                        }
                                                                                        if false {
                                                                                            v190 := GenS22{F0: (v47 >> ((42) & 63))}
                                                                                            v191 := GenS18{F0: 158}
                                                                                            var v192 int32 = -7
                                                                                            switch (fi80 | 32767) {
                                                                                            case 266:
                                                                                                if len(v98) >= 2 {
                                                                                                    v98 = v98[1:len(v98)]
                                                                                                }
                                                                                            default:
                                                                                                if (0.1 >= (-2.25 / (0.1 / 1.5))) {
                                                                                                    if len(v98) > 0 {
                                                                                                        v98[0] = (^(fi81))
                                                                                                    }
                                                                                                    var v193 uint32 = 42
                                                                                                    switch ("" + ("\n\t" + "\n\t")) {
                                                                                                    case "a":
                                                                                                        var v194 int64 = v49
                                                                                                        fmt.Println(fi81)
                                                                                                        _ = v194
                                                                                                    }
                                                                                                    _ = v193
                                                                                                }
                                                                                            }
                                                                                            _ = v190
                                                                                            _ = v191
                                                                                            _ = v192
                                                                                        } else {
                                                                                            if (1e10 > ((((float32((0.0 / float64(1e10))) - (1e10 - (-2.25 + float32((((0.0 / (0.1 / 0.1)) - 3.14159) - ((1e10 / 1e10) - 0.0)))))) * (0.0 - (((0.1 * 1.5) / (float32(-2.25) / 0.1)) + -2.25))) / ((3.14159 / 1.0) + 3.14159)) * (3.14159 / 1e10))) {
                                                                                                var v195 uint64 = v47
                                                                                                _ = v195
                                                                                            }
                                                                                            var v196 string = "a"
                                                                                            var v197 string = v196
                                                                                            _ = v196
                                                                                            _ = v197
                                                                                        }
                                                                                    }
                                                                                }
                                                                                fmt.Println(v49)
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v98
                                                                }
                                                                if false {
                                                                    fmt.Println(v48, fi81, v48)
                                                                    var v198 int = fi81
                                                                    _ = v198
                                                                } else {
                                                                    v48 += (v49 + 9223372036854775807)
                                                                    if false {
                                                                        v199 := []int{}
                                                                        v47 = (v47 << ((351) & 63))
                                                                        _ = v199
                                                                    }
                                                                    v200 := GenS13{F0: 1000, F1: 42, F2: 65535}
                                                                    _ = v200
                                                                }
                                                            } else {
                                                                var v201 float32 = (((0.0 + 3.14159) / -2.25) * (1.5 + 0.0))
                                                                v202 := GenS23{F0: v47, F1: 1}
                                                                fmt.Println(v201, v49, fi81)
                                                                if (fi80 < (fi81 & fi81)) {
                                                                    fmt.Println(v201, v47, v48)
                                                                } else {
                                                                    var v203 uint32 = 2147483648
                                                                    _ = v203
                                                                }
                                                                _ = v201
                                                                _ = v202
                                                            }
                                                        }
                                                        fmt.Println(v47)
                                                        if true {
                                                            switch fi81 {
                                                            case 9223372036854775807:
                                                                v47 = 18446744073709551615
                                                            }
                                                        } else {
                                                            v204 := GenS24{F0: 2147483647}
                                                            _ = v204
                                                        }
                                                    }
                                                    if true {
                                                        if false {
                                                            var v205 uint16 = 32767
                                                            v206 := []int{636, fi81, 127}
                                                            _ = v205
                                                            _ = v206
                                                        } else {
                                                            v207 := GenS25{F0: (-(fi81))}
                                                            var v208 uint16 = 73
                                                            var v209 uint = 18446744073709551615
                                                            _ = v207
                                                            _ = v208
                                                            _ = v209
                                                        }
                                                        var v210 uint = 18446744073709551615
                                                        _ = v210
                                                    } else {
                                                        var v211 bool = (!((fi81 ^ 32768) >= fi80))
                                                        _ = v211
                                                    }
                                                    v212 := GenS26{F0: 42, F1: ((v47 & v47) > v47)}
                                                    _ = v212
                                                default:
                                                    switch "foo bar" {
                                                    case "hello":
                                                        switch "foo bar" {
                                                        case "":
                                                            fmt.Println(fi81)
                                                        case "foo bar":
                                                            if (false || false) {
                                                                v213 := GenS27{F0: 65535, F1: -554}
                                                                v214 := GenS28{F0: -9223372036854775808, F1: ("" + "foo bar")}
                                                                var v215 float32 = float32(1.5)
                                                                _ = v213
                                                                _ = v214
                                                                _ = v215
                                                            } else {
                                                                var v216 uint32 = 2147483647
                                                                fmt.Println(fi80, v216)
                                                                fmt.Println(v50, v50, v50)
                                                                _ = v216
                                                            }
                                                        case "hello":
                                                            if ((^(v48)) == (v49 & (v48 % ((42) | 1)))) {
                                                                if false {
                                                                    if ((128 != uint8(v47)) || false) {
                                                                        switch "\n\t" {
                                                                        case "":
                                                                            switch "foo bar" {
                                                                            case "\n\t":
                                                                                v47 -= v47
                                                                            default:
                                                                                v217 := GenS29{F0: 1, F1: 128, F2: -32768}
                                                                                _ = v217
                                                                            }
                                                                        case "hello":
                                                                            switch ("a" + (("hello" + "") + ("foo bar" + ""))) {
                                                                            case "hello":
                                                                                if ((false || true) && (-2.25 < 1e10)) {
                                                                                    var v218 int64 = 42
                                                                                    if true {
                                                                                        switch fi80 {
                                                                                        case -1:
                                                                                            v219 := GenS19{F0: (1 >= 0), F1: (^(fi80)), F2: (603 > 18446744073709551615)}
                                                                                            _ = v219
                                                                                        case -6:
                                                                                            if true {
                                                                                                v47 = v47
                                                                                                var v220 int32 = -1
                                                                                                v221 := []int{(fi81 - -9223372036854775808), fi81, fi81}
                                                                                                _ = v220
                                                                                                _ = v221
                                                                                            }
                                                                                            var v222 int16 = int16(v49)
                                                                                            _ = v222
                                                                                        }
                                                                                        v47 = (v47 >> ((1) & 63))
                                                                                        if (0 == 65535) {
                                                                                            var v223 float32 = (0.0 - float32((0.1 * 0.0)))
                                                                                            _ = v223
                                                                                        }
                                                                                    } else {
                                                                                        var v224 int32 = 661
                                                                                        v225 := []int{(-(fi80))}
                                                                                        v47 = v47
                                                                                        _ = v224
                                                                                        _ = v225
                                                                                    }
                                                                                    var v226 uint8 = 224
                                                                                    _ = v218
                                                                                    _ = v226
                                                                                }
                                                                            }
                                                                        case "foo bar":
                                                                            var v227 float32 = (0.1 / 1.5)
                                                                            _ = v227
                                                                        }
                                                                        if ((!(579 >= 42)) && true) {
                                                                            var v228 int32 = -704
                                                                            var v229 int32 = v228
                                                                            _ = v228
                                                                            _ = v229
                                                                        } else {
                                                                            v47 = 1
                                                                        }
                                                                        if ((^(v47)) < (v47 ^ (v47 | (v47 << ((323) & 63))))) {
                                                                            switch ("\n\t" + ((("a" + "hello") + "a") + "a")) {
                                                                            case "a":
                                                                                switch -889 {
                                                                                case 32768:
                                                                                    var v230 int8 = 42
                                                                                    _ = v230
                                                                                case 0:
                                                                                    v231 := []int{(fi80 | (fi80 % ((fi80) | 1))), int(v49), 4294967295}
                                                                                    _ = v231
                                                                                default:
                                                                                    v48 = v50
                                                                                }
                                                                            case "foo bar":
                                                                                if (127 <= 42) {
                                                                                    if (false && ((^(v47)) > 42)) {
                                                                                        var v232 int64 = (-(v48))
                                                                                        _ = v232
                                                                                    } else {
                                                                                        if ((114 >= 42) || false) {
                                                                                            v48 = v50
                                                                                        }
                                                                                    }
                                                                                    var v233 float32 = 1e10
                                                                                    _ = v233
                                                                                }
                                                                            case "hello":
                                                                                v234 := []int{(fi80 & fi80), (fi81 | 802), (fi80 / (((fi81 - fi81)) | 1))}
                                                                                _ = v234
                                                                            default:
                                                                                v235 := GenS30{F0: 0, F1: 42, F2: (v47 | v47)}
                                                                                _ = v235
                                                                            }
                                                                            var v236 uint16 = 435
                                                                            v237 := []int{}
                                                                            _ = v236
                                                                            _ = v237
                                                                        }
                                                                    }
                                                                    v49 -= (v48 - (v50 + v50))
                                                                }
                                                                if true {
                                                                    v238 := []int{}
                                                                    v48 = v50
                                                                    switch fi81 {
                                                                    case 292:
                                                                        switch ("hello" + "foo bar") {
                                                                        case "\n\t":
                                                                            if true {
                                                                                v47 -= (v47 % ((798) | 1))
                                                                                if true {
                                                                                    v239 := GenS31{F0: 128}
                                                                                    _ = v239
                                                                                } else {
                                                                                    switch "\n\t" {
                                                                                    case "":
                                                                                        switch "" {
                                                                                        case "hello":
                                                                                            var v240 uint32 = 4294967295
                                                                                            _ = v240
                                                                                        default:
                                                                                            fmt.Println(v47)
                                                                                        }
                                                                                    }
                                                                                }
                                                                                if len(v238) > 0 {
                                                                                    v238[0] = 0
                                                                                }
                                                                                if (!(!true)) {
                                                                                    if (1.5 == 3.14159) {
                                                                                        if false {
                                                                                            var v241 uint8 = 128
                                                                                            _ = v241
                                                                                        }
                                                                                        v48 -= -1
                                                                                        fmt.Println(v48, v50, v50)
                                                                                    }
                                                                                    _ = len(v238)
                                                                                    var v242 bool = false
                                                                                    _ = v242
                                                                                }
                                                                                switch ("\n\t" + ("a" + "\n\t")) {
                                                                                case "a":
                                                                                    if (true && ((("\n\t" + "hello") >= "hello") || true)) {
                                                                                        if false {
                                                                                            var v243 bool = false
                                                                                            v243 = (0 == 0)
                                                                                            var v244 uint32 = 65535
                                                                                            _ = v243
                                                                                            _ = v244
                                                                                        }
                                                                                        fmt.Println(v48, v49, v48)
                                                                                        v245 := GenS28{F0: (v48 % ((v50) | 1)), F1: "foo bar"}
                                                                                        if len(v238) > 0 {
                                                                                            v238[0] = -312
                                                                                        }
                                                                                        _ = v245
                                                                                    }
                                                                                case "foo bar":
                                                                                    v47 = v47
                                                                                }
                                                                            }
                                                                        case "a":
                                                                            var v246 bool = (42 <= (v47 ^ (v47 / ((v47) | 1))))
                                                                            _ = v246
                                                                        default:
                                                                            v47 += 18446744073709551615
                                                                        }
                                                                    default:
                                                                        if false {
                                                                            var v247 uint16 = 154
                                                                            var v248 int16 = -978
                                                                            if true {
                                                                                switch (^(fi81)) {
                                                                                case 1:
                                                                                    if true {
                                                                                        v50 = 0
                                                                                        _ = len(v238)
                                                                                        switch (("a" + ("" + ("\n\t" + "\n\t"))) + "hello") {
                                                                                        case "":
                                                                                            if ((0.0 * ((0.1 + 1e10) + (-2.25 / 1e10))) < 0.1) {
                                                                                                if len(v238) > 0 {
                                                                                                    v238[0] = (fi80 ^ fi81)
                                                                                                }
                                                                                            } else {
                                                                                                v238 = append(v238, (^(fi80)))
                                                                                                var v249 int64 = v49
                                                                                                v250 := GenS23{F0: v47, F1: 162}
                                                                                                _ = v249
                                                                                                _ = v250
                                                                                            }
                                                                                        case "foo bar":
                                                                                            v49 = v48
                                                                                        default:
                                                                                            if ((fi81 * fi80) < (fi81 & fi81)) {
                                                                                                v50 -= 42
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        var v251 uint16 = (v247 & (^(v247)))
                                                                                        _ = v251
                                                                                    }
                                                                                default:
                                                                                    if len(v238) >= 2 {
                                                                                        v238 = v238[1:len(v238)]
                                                                                    }
                                                                                }
                                                                            }
                                                                            _ = v247
                                                                            _ = v248
                                                                        } else {
                                                                            if true {
                                                                                var v252 uint64 = v47
                                                                                var v253 uint64 = 1
                                                                                if false {
                                                                                    if true {
                                                                                        fmt.Println(v50, fi80, v252)
                                                                                        var v254 uint64 = v253
                                                                                        _ = v254
                                                                                    } else {
                                                                                        var v255 string = "a"
                                                                                        _ = v255
                                                                                    }
                                                                                    switch "a" {
                                                                                    case "\n\t":
                                                                                        var v256 float32 = 1e10
                                                                                        _ = v256
                                                                                    case "foo bar":
                                                                                        v257 := GenS6{F0: 0, F1: 199}
                                                                                        _ = v257
                                                                                    default:
                                                                                        if (false && (false && true)) {
                                                                                            switch "hello" {
                                                                                            case "\n\t":
                                                                                                v47 = v253
                                                                                            case "a":
                                                                                                v50 -= -1
                                                                                            case "":
                                                                                                var v258 int32 = 42
                                                                                                _ = v258
                                                                                            default:
                                                                                                var v259 uint16 = 2
                                                                                                _ = v259
                                                                                            }
                                                                                            if true {
                                                                                                v47 = (^(v252))
                                                                                                var v260 uint16 = 1
                                                                                                _ = v260
                                                                                            } else {
                                                                                                fmt.Println(v47, v47)
                                                                                            }
                                                                                        } else {
                                                                                            if true {
                                                                                                var v261 float64 = (float64(0.0) * float64(1.5))
                                                                                                _ = v261
                                                                                            } else {
                                                                                                var v262 int64 = v49
                                                                                                _ = v262
                                                                                            }
                                                                                            v263 := GenS32{F0: 2, F1: (!(!(-1 > -1))), F2: ""}
                                                                                            var v264 uint8 = 161
                                                                                            _ = v263
                                                                                            _ = v264
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(v253)
                                                                                    v238 = append(v238, fi81)
                                                                                    var v265 uint8 = 2
                                                                                    var v266 uint64 = (v47 / ((v47) | 1))
                                                                                    _ = v265
                                                                                    _ = v266
                                                                                }
                                                                                _ = v252
                                                                                _ = v253
                                                                            } else {
                                                                                if false {
                                                                                    var v267 int16 = -3
                                                                                    var v268 int16 = 1
                                                                                    var v269 float32 = (0.0 + 0.1)
                                                                                    _ = v267
                                                                                    _ = v268
                                                                                    _ = v269
                                                                                } else {
                                                                                    v50 = v48
                                                                                    v49 -= (-(v50))
                                                                                    var v270 int32 = 519
                                                                                    _ = v270
                                                                                }
                                                                            }
                                                                            var v271 float32 = -2.25
                                                                            _ = v271
                                                                        }
                                                                    }
                                                                    v48 -= v50
                                                                    if (0 >= 42) {
                                                                        _ = len(v238)
                                                                        var v272 int32 = 0
                                                                        v273 := GenS9{F0: 42, F1: 2, F2: -1}
                                                                        fmt.Println(fi81, v50, v47)
                                                                        var v274 int = -158
                                                                        var v275 float64 = -2.25
                                                                        _ = v272
                                                                        _ = v273
                                                                        _ = v274
                                                                        _ = v275
                                                                    }
                                                                    _ = v238
                                                                }
                                                            }
                                                        }
                                                    case "\n\t":
                                                        v49 += v49
                                                    }
                                                }
                                                _ = v92
                                            }
                                            v49 += (v50 - (v49 << ((42) & 63)))
                                            var v276 float64 = (0.1 - 0.1)
                                            _ = v276
                                        }
                                        if (float64(0.0) != 3.14159) {
                                            switch int(v50) {
                                            case 4294967295:
                                                switch "a" {
                                                case "\n\t":
                                                    v49 += (v49 % ((-1) | 1))
                                                case "foo bar":
                                                    fmt.Println(fi81, v50, v49)
                                                case "":
                                                    var v277 int64 = (v49 + 42)
                                                    _ = v277
                                                default:
                                                    var v278 int64 = v49
                                                    _ = v278
                                                }
                                            default:
                                                v279 := GenS30{F0: 4294967295, F1: 42, F2: (v47 ^ (v47 * (v47 | (^(v47)))))}
                                                _ = v279
                                            }
                                            var v280 int = -9223372036854775808
                                            fmt.Println(fi80)
                                            _ = v280
                                        } else {
                                            v281 := []int{fi81}
                                            v282 := GenS1{F0: false, F1: 4294967295, F2: (false || true)}
                                            fmt.Println(v48)
                                            _ = v281
                                            _ = v282
                                        }
                                        _ = fi81
                                    }
                                }
                                for fi283 := 0; fi283 < 11; fi283++ {
                                    fmt.Println(v47)
                                    var v284 int32 = -425
                                    _ = fi283
                                    _ = v284
                                }
                                if false {
                                    v285 := GenS16{F0: 2147483647}
                                    _ = v285
                                } else {
                                    if (1 <= 199) {
                                        var v286 uint = 1
                                        v47 = 91
                                        switch "" {
                                        case "foo bar":
                                            v47 = v47
                                        case "":
                                            if false {
                                                v287 := GenS16{F0: 1}
                                                if true {
                                                    v288 := []int{(fi80 + 813)}
                                                    var v289 uint16 = 42
                                                    switch fi80 {
                                                    case 9223372036854775807:
                                                        var v290 int32 = 0
                                                        _ = v290
                                                    case 2147483647:
                                                        var v291 int = fi80
                                                        _ = v291
                                                    case 42:
                                                        if (-114 != 42) {
                                                            var v292 uint = v286
                                                            v293 := GenS33{F0: (4294967295 == (^(v47)))}
                                                            _ = v292
                                                            _ = v293
                                                        }
                                                    default:
                                                        if true {
                                                            if true {
                                                                var v294 uint16 = 28
                                                                _ = v294
                                                            } else {
                                                                var v295 int64 = -234
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    for fi296 := 0; fi296 < 1; fi296++ {
                                                                        if true {
                                                                            var v297 uint8 = 42
                                                                            fmt.Println(fi296, fi80)
                                                                            fmt.Println(v48)
                                                                            if (0.0 <= float64(float32((0.0 * 3.14159)))) {
                                                                                fmt.Println(v50, v289)
                                                                                v298 := GenS18{F0: v297}
                                                                                v288 = append(v288, fi296)
                                                                                v48 = (v50 | v50)
                                                                                switch -85 {
                                                                                case 1:
                                                                                    var v299 int = (-(fi296))
                                                                                    _ = v299
                                                                                case -2147483648:
                                                                                    switch (fi296 % ((fi80) | 1)) {
                                                                                    case 42:
                                                                                        if false {
                                                                                            var v300 int64 = -9223372036854775808
                                                                                            v297 = v297
                                                                                            v301 := GenS1{F0: (((("a" + "foo bar") + ("foo bar" + "hello")) + (("hello" + ("a" + "\n\t")) + ("a" + ""))) == "\n\t"), F1: (v286 ^ 2), F2: (false || false)}
                                                                                            switch (^(fi296)) {
                                                                                            case 32768:
                                                                                                var v302 uint16 = (^(v289))
                                                                                                _ = v302
                                                                                            case -574:
                                                                                                switch 0 {
                                                                                                case 1:
                                                                                                    v295 += v50
                                                                                                case -163:
                                                                                                    v303 := GenS34{F0: 847, F1: 2147483647, F2: 2}
                                                                                                    _ = v303
                                                                                                case -240:
                                                                                                    if (!false) {
                                                                                                        switch "a" {
                                                                                                        case "foo bar":
                                                                                                            fmt.Println(fi296, fi296)
                                                                                                        case "\n\t":
                                                                                                            var v304 string = "\n\t"
                                                                                                            _ = v304
                                                                                                        }
                                                                                                        v305 := GenS24{F0: 2147483647}
                                                                                                        v47 = (^(v47))
                                                                                                        _ = v305
                                                                                                    }
                                                                                                default:
                                                                                                    v300 = v50
                                                                                                }
                                                                                            }
                                                                                            _ = v300
                                                                                            _ = v301
                                                                                        }
                                                                                    case -477:
                                                                                        var v306 int = 2147483647
                                                                                        _ = v306
                                                                                    }
                                                                                case 42:
                                                                                    v297 -= (^(v297))
                                                                                }
                                                                                v297 = (^(v297))
                                                                                v286 = (^(v286))
                                                                                _ = v298
                                                                            } else {
                                                                                v286 = v286
                                                                                v307 := GenS33{F0: false}
                                                                                var v308 uint64 = (v47 >> (((v286 % (((v286 << (((v286 | (v286 ^ uint(v297)))) & 63))) | 1))) & 63))
                                                                                _ = v307
                                                                                _ = v308
                                                                            }
                                                                            v49 = v50
                                                                            _ = v297
                                                                        } else {
                                                                            var v309 uint16 = v289
                                                                            _ = v309
                                                                        }
                                                                        _ = fi296
                                                                    }
                                                                default:
                                                                    for fi310 := 0; fi310 < 19; fi310++ {
                                                                        switch "hello" {
                                                                        case "a":
                                                                            if len(v288) > 0 {
                                                                                v288[0] = (-(fi80))
                                                                            }
                                                                        case "hello":
                                                                            v47 = (v47 & v47)
                                                                        case "":
                                                                            switch "a" {
                                                                            case "":
                                                                                var v311 uint16 = v289
                                                                                _ = v311
                                                                            case "\n\t":
                                                                                if (!(!((3.14159 / float64((1.5 * -2.25))) < 3.14159))) {
                                                                                    fmt.Println(v50, v49, v49)
                                                                                    v295 += v48
                                                                                }
                                                                            }
                                                                        default:
                                                                            if len(v288) >= 2 {
                                                                                v288 = v288[1:len(v288)]
                                                                            }
                                                                        }
                                                                        _ = fi310
                                                                    }
                                                                }
                                                                _ = v295
                                                            }
                                                            var v312 float32 = (-2.25 - float32((float64(0.1) + (3.14159 + 1e10))))
                                                            var v313 uint8 = 255
                                                            _ = v312
                                                            _ = v313
                                                        }
                                                    }
                                                    _ = v288
                                                    _ = v289
                                                } else {
                                                    for fi314 := 0; fi314 < 2; fi314++ {
                                                        if ((!((!(false && (1 > 42))) && true)) || (1 <= -1)) {
                                                            var v315 uint = (^(v286))
                                                            v316 := []int{fi314, (fi80 << (((v286 - uint(v47))) & 63)), fi80}
                                                            if len(v316) > 0 {
                                                                v316[0] = int(v47)
                                                            }
                                                            _ = v315
                                                            _ = v316
                                                        }
                                                        fmt.Println(v49, v48, v48)
                                                        if true {
                                                            v50 += (v49 % ((v49) | 1))
                                                            var v317 string = "\n\t"
                                                            _ = v317
                                                        } else {
                                                            if false {
                                                                fmt.Println(v286, v48)
                                                                var v318 int64 = (v48 >> (((v286 << ((v286) & 63))) & 63))
                                                                fmt.Println(v48)
                                                                _ = v318
                                                            } else {
                                                                fmt.Println(fi314, v48, v50)
                                                                if (100 <= 100) {
                                                                    v48 = 42
                                                                }
                                                            }
                                                        }
                                                        fmt.Println(fi314, v49)
                                                        if true {
                                                            var v319 int8 = 100
                                                            v50 += 1
                                                            v320 := GenS29{F0: 42, F1: 255, F2: 42}
                                                            if (1 > 739) {
                                                                if (true || (!true)) {
                                                                    var v321 int64 = (v49 | int64(v319))
                                                                    if false {
                                                                        v48 += v48
                                                                        var v322 int64 = (-(v321))
                                                                        _ = v322
                                                                    }
                                                                    _ = v321
                                                                } else {
                                                                    v323 := []int{}
                                                                    _ = v323
                                                                }
                                                            }
                                                            if false {
                                                                v324 := []int{fi314, 1, (-(fi80))}
                                                                _ = v324
                                                            }
                                                            _ = v319
                                                            _ = v320
                                                        }
                                                        _ = fi314
                                                    }
                                                    v49 += (^(v50))
                                                    fmt.Println(v286, fi80)
                                                }
                                                _ = v287
                                            }
                                        case "hello":
                                            if true {
                                                v48 -= v49
                                                if (!(false && true)) {
                                                    var v325 uint16 = 32767
                                                    v326 := []int{}
                                                    switch fi80 {
                                                    case -9223372036854775808:
                                                        if len(v326) > 0 {
                                                            v326[0] = (fi80 ^ (-(fi80)))
                                                        }
                                                    case 128:
                                                        v47 -= (v47 & v47)
                                                    case 32768:
                                                        v327 := GenS23{F0: (v47 - v47), F1: 1}
                                                        _ = v327
                                                    }
                                                    _ = v325
                                                    _ = v326
                                                } else {
                                                    if true {
                                                        for fi328 := 0; fi328 < 5; fi328++ {
                                                            fmt.Println(v286, v50, fi328)
                                                            var v329 float64 = float64(1e10)
                                                            fmt.Println(v329)
                                                            _ = fi328
                                                            _ = v329
                                                        }
                                                        for fi330 := 0; fi330 < 8; fi330++ {
                                                            v331 := []int{(fi80 ^ fi80)}
                                                            var v332 bool = true
                                                            _ = fi330
                                                            _ = v331
                                                            _ = v332
                                                        }
                                                    }
                                                    v333 := GenS35{F0: (true || true), F1: -2147483648}
                                                    _ = v333
                                                }
                                                v334 := GenS36{F0: -2147483648, F1: 691, F2: 65535}
                                                _ = v334
                                            } else {
                                                switch "a" {
                                                case "foo bar":
                                                    fmt.Println(v286, v49, v49)
                                                case "hello":
                                                    for fi335 := 0; fi335 < 12; fi335++ {
                                                        v49 += (v49 * (^(v48)))
                                                        v286 = (^(v286))
                                                        v336 := []int{(fi335 * fi80), fi335, (fi80 - fi80), fi335}
                                                        _ = fi335
                                                        _ = v336
                                                    }
                                                case "\n\t":
                                                    for fi337 := 0; fi337 < 12; fi337++ {
                                                        var v338 int32 = int32(fi80)
                                                        var v339 string = "hello"
                                                        _ = fi337
                                                        _ = v338
                                                        _ = v339
                                                    }
                                                default:
                                                    if (v48 == (v50 & (v50 >> (((v286 / (((^(v286))) | 1))) & 63)))) {
                                                        if false {
                                                            v340 := []int{fi80, fi80, (fi80 << ((0) & 63))}
                                                            _ = v340
                                                        } else {
                                                            fmt.Println(v49, fi80, v286)
                                                            for fi341 := 0; fi341 < 5; fi341++ {
                                                                v48 = (v48 << ((v286) & 63))
                                                                v342 := GenS37{F0: 1, F1: (^(v286))}
                                                                var v343 bool = (false || true)
                                                                _ = fi341
                                                                _ = v342
                                                                _ = v343
                                                            }
                                                        }
                                                    } else {
                                                        var v344 uint16 = 35
                                                        _ = v344
                                                    }
                                                }
                                            }
                                        default:
                                            fmt.Println(v48, v49, v50)
                                        }
                                        _ = v286
                                    } else {
                                        v48 = v49
                                    }
                                    var v345 int = (fi80 * (fi80 - fi80))
                                    v48 = v49
                                    _ = v345
                                }
                                _ = fi80
                            }
                            _ = v50
                        }
                        fmt.Println(v48, v47, v49)
                        _ = v49
                    }
                    _ = v47
                    _ = v48
                } else {
                    var v346 float64 = (0.1 + -2.25)
                    fmt.Println(v346)
                    _ = v346
                }
                var v347 int64 = -1
                if (false || true) {
                    var v348 uint32 = 1
                    _ = v348
                } else {
                    v349 := GenS21{F0: "a", F1: 7, F2: 0}
                    _ = v349
                }
                _ = v347
            }
            var v350 uint8 = 42
            _ = v350
        }
    }
    if true {
        if (!false) {
            v351 := GenS38{F0: -1, F1: -1}
            _ = v351
        }
        v352 := []int{}
        var v353 string = ("foo bar" + (("" + "a") + "\n\t"))
        _ = v352
        _ = v353
    }
    return 1
}
func genFn1(p0 int64) uint8 {
    p0 = p0
    for fi354 := 0; fi354 < 18; fi354++ {
        var v355 int32 = 345
        switch fi354 {
        case -2147483648:
            var v356 int16 = 1
            _ = v356
        default:
            if true {
                var v357 int16 = 1
                switch ((("foo bar" + (("" + "foo bar") + ("" + "\n\t"))) + "hello") + "foo bar") {
                case "a":
                    var v358 float32 = (-2.25 + float32((1.5 + 3.14159)))
                    _ = v358
                }
                _ = v357
            }
        }
        if true {
            if ((3.14159 + float32(((-2.25 * 3.14159) / (0.1 - 3.14159)))) != (1e10 + -2.25)) {
                var v359 string = "hello"
                _ = v359
            }
        }
        _ = fi354
        _ = v355
    }
    if false {
        fmt.Println(p0)
        for fi360 := 0; fi360 < 9; fi360++ {
            switch "" {
            case "foo bar":
                var v361 uint16 = 1
                _ = v361
            }
            v362 := GenS39{F0: 0, F1: (fi360 / ((fi360) | 1)), F2: 1}
            _ = fi360
            _ = v362
        }
        p0 = p0
    } else {
        for fi363 := 0; fi363 < 19; fi363++ {
            v364 := []int{}
            _ = fi363
            _ = v364
        }
        p0 -= (p0 + (^(p0)))
        v365 := []int{0}
        _ = v365
    }
    v366 := []int{32767, 127, -2147483648, 4294967295}
    var v367 uint16 = 255
    _ = p0
    _ = v366
    _ = v367
    return 152
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{65535, 2147483648, 978}
    var v3 float64 = float64((1e10 * 1.5))
    var v4 uint16 = 42
    v5 := make(map[string]string)
    v6 := make(map[uint32]uint32)
    if ((v3 + 1e10) < (((v3 - v3) * v3) - 1e10)) {
        v368 := GenS40{F0: 0}
        for fi369 := 0; fi369 < 1; fi369++ {
            if (((^(fi369)) < fi369) || (true || (!false))) {
                fmt.Println(v3)
            }
            _ = fi369
        }
        switch "\n\t" {
        case "":
            for fi370 := 0; fi370 < 11; fi370++ {
                v4 = (v4 + 1)
                _ = fi370
            }
        default:
            delete(v1, 714)
        }
        _ = v368
    }
    delete(v6, 0)
    switch -2147483648 {
    case 1:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 127:
        if (1 >= -117) {
            var v371 uint8 = 0
            _ = v371
        } else {
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
            v3 += v3
            for fi372 := 0; fi372 < 16; fi372++ {
                for fi373 := 0; fi373 < 4; fi373++ {
                    v1[0] = true
                    var v374 int64 = -9223372036854775808
                    _ = fi373
                    _ = v374
                }
                v2 = append(v2, 2)
                var v375 float64 = (((float64(3.14159) - (1e10 * (float64(-2.25) + (v3 / (v3 / v3))))) / (v3 / (v3 + (v3 + float64((1e10 / 1e10)))))) * 0.0)
                _ = fi372
                _ = v375
            }
        }
    }
    if false {
        var v376 int = 0
        for fi377 := 0; fi377 < 6; fi377++ {
            v378 := GenS41{F0: -9223372036854775808}
            var v379 uint = 1
            for fi380 := 0; fi380 < 9; fi380++ {
                if false {
                    v381 := GenS12{F0: fi380}
                    _ = v381
                }
                v382 := GenS42{F0: 0, F1: 2147483647, F2: 1}
                switch "\n\t" {
                case "hello":
                    var v383 uint16 = v4
                    _ = v383
                case "foo bar":
                    v384 := GenS28{F0: 0, F1: ""}
                    _ = v384
                }
                _ = fi380
                _ = v382
            }
            _ = fi377
            _ = v378
            _ = v379
        }
        _ = v376
    }
    v4 = v4
    if (2 == 42) {
        v3 = v3
        fmt.Println(v3, v4)
        for fi385 := 0; fi385 < 14; fi385++ {
            v386 := GenS43{F0: 1, F1: 1}
            fmt.Println(fi385, v4)
            _ = fi385
            _ = v386
        }
    } else {
        switch (("" + ("a" + "")) + "\n\t") {
        case "\n\t":
            var v387 int = -373
            _ = v387
        case "hello":
            v2 = append(v2, 65535)
        case "foo bar":
            fmt.Println(v4, v3)
        }
        v388 := GenS9{F0: 1, F1: v4, F2: 1}
        for fi389 := 0; fi389 < 1; fi389++ {
            _ = len(v1)
            _ = fi389
        }
        _ = v388
    }
    var v390 uint8 = 2
    fmt.Println(v3, v4)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        switch -2147483648 {
        case -2147483648:
            var v391 int16 = -1
            _ = v391
        case -9223372036854775808:
            v3 -= float64(3.14159)
        }
        if false {
            v390 += v390
            v392 := GenS15{F0: -69, F1: 2147483647}
            _ = v392
        }
        switch "a" {
        case "\n\t":
            var v393 uint64 = 518
            _ = v393
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v390:", v390)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v390
    _ = v4
    _ = v5
    _ = v6
}
