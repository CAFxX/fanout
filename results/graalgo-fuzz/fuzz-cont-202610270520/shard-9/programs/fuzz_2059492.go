package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 int32; F2 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 uint32; F1 int16; F2 string }
type GenS7 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS8 struct { F0 uint32; F1 int }
type GenS9 struct { F0 int32; F1 int32 }
type GenS10 struct { F0 uint64; F1 int32 }
type GenS11 struct { F0 int; F1 uint32 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 int32; F1 int }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int64; F1 bool }
type GenS16 struct { F0 int8; F1 uint; F2 int }
type GenS17 struct { F0 int16; F1 uint32 }
type GenS18 struct { F0 uint }
type GenS19 struct { F0 uint; F1 uint; F2 string }
type GenS20 struct { F0 int8; F1 int8; F2 uint8 }
type GenS21 struct { F0 int; F1 uint16 }
type GenS22 struct { F0 int64; F1 int16 }
type GenS23 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 uint32 = 469
    var v2 uint64 = 672
    var v3 uint32 = 913
    var v4 uint8 = 44
    v5 := GenS1{F0: (("foo bar" + ("a" + "foo bar")) + "a"), F1: 899}
    v6 := []int{2147483647}
    v7 := GenS2{F0: 323}
    if ((-514 <= 0) || (127 != -18)) {
        switch ((("\n\t" + "foo bar") + "foo bar") + ((("\n\t" + (("\n\t" + "") + (("" + "foo bar") + "a"))) + ("foo bar" + "a")) + "foo bar")) {
        case "foo bar":
            v4 = 215
        }
        for fi8 := 0; fi8 < 8; fi8++ {
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 int8 = 124
                _ = fi9
                _ = v10
            }
            v3 -= (v3 + (v1 / (((v3 - uint32(v2))) | 1)))
            _ = fi8
        }
        switch -1 {
        case 127:
            v3 -= 65535
        case 865:
            for fi11 := 0; fi11 < 5; fi11++ {
                v2 += 1
                _ = fi11
            }
        }
    }
    var v12 float32 = -2.25
    if false {
        switch 127 {
        case -927:
            var v13 int32 = 2147483647
            _ = v13
        case -1:
            fmt.Println(v2)
        case 42:
            if len(v6) > 0 {
                v6[0] = 9223372036854775807
            }
        }
        switch 116 {
        case 32767:
            v14 := GenS2{F0: 2}
            _ = v14
        case -534:
            var v15 int32 = 2147483647
            _ = v15
        default:
            var v16 uint8 = (v4 ^ 255)
            _ = v16
        }
    } else {
        var v17 float64 = (1.5 * 3.14159)
        switch "\n\t" {
        case "\n\t":
            for fi18 := 0; fi18 < 2; fi18++ {
                var v19 bool = (704 >= 1)
                if (false && false) {
                    v19 = v19
                    if false {
                        fmt.Println(v17, v1)
                        var v20 uint8 = v4
                        if (!v19) {
                            for fi21 := 0; fi21 < 9; fi21++ {
                                switch "a" {
                                case "foo bar":
                                    if v19 {
                                        if len(v6) > 0 {
                                            v6[0] = -2147483648
                                        }
                                    } else {
                                        v19 = v19
                                        if v19 {
                                            v22 := GenS1{F0: (("foo bar" + ("hello" + (("a" + "foo bar") + ("hello" + (("a" + ("a" + "")) + "foo bar"))))) + "hello"), F1: (v3 ^ v3)}
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            _ = v22
                                        }
                                        var v23 float64 = (0.0 / (v17 - (v17 - (3.14159 * v17))))
                                        _ = v23
                                    }
                                case "":
                                    if len(v6) > 0 {
                                        v6[0] = fi18
                                    }
                                case "\n\t":
                                    var v24 float64 = float64(v12)
                                    v12 = 3.14159
                                    _ = v24
                                default:
                                    switch fi21 {
                                    case 1:
                                        switch 1 {
                                        case -1:
                                            v2 -= v2
                                        case 32768:
                                            var v25 int64 = 1
                                            v25 += (v25 * v25)
                                            if (-115 == 10) {
                                                v12 += 3.14159
                                            }
                                            _ = v25
                                        case -779:
                                            var v26 float64 = float64(((v12 * v12) / v12))
                                            if (!v19) {
                                                var v27 uint16 = 42
                                                var v28 uint64 = v2
                                                _ = len(v6)
                                                _ = v27
                                                _ = v28
                                            } else {
                                                v29 := GenS3{F0: 612, F1: 2147483647, F2: (fi18 - fi21)}
                                                var v30 uint32 = (^(v1))
                                                var v31 int8 = -100
                                                _ = v29
                                                _ = v30
                                                _ = v31
                                            }
                                            _ = v26
                                        }
                                    case -131:
                                        switch fi18 {
                                        case 128:
                                            v4 -= (v20 << ((255) & 63))
                                        case 263:
                                            fmt.Println(v2, v20, v19)
                                            v32 := GenS4{F0: 1}
                                            _ = v32
                                        case -2147483648:
                                            v33 := GenS5{F0: -2147483648, F1: 65535}
                                            _ = v33
                                        default:
                                            var v34 uint32 = 933
                                            _ = v34
                                        }
                                    default:
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    }
                                }
                                fmt.Println(v20, v4, v4)
                                v17 += 1.5
                                var v35 int8 = 0
                                _ = len(v6)
                                v36 := GenS5{F0: 2147483647, F1: v1}
                                switch (fi21 & (fi21 | (-(fi18)))) {
                                case 2147483647:
                                    switch -9223372036854775808 {
                                    case -170:
                                        var v37 float64 = 3.14159
                                        _ = v37
                                    case 128:
                                        if v19 {
                                            if (!false) {
                                                var v38 int16 = 369
                                                switch (^(fi21)) {
                                                case 32768:
                                                    if (v19 && (v19 && ((v12 + -2.25) > ((v12 / ((v12 / v12) - v12)) + v12)))) {
                                                        if v19 {
                                                            if len(v6) > 0 {
                                                                v6[0] = (fi18 ^ (fi21 * fi21))
                                                            }
                                                            var v39 int16 = v38
                                                            _ = v39
                                                        }
                                                        v2 = v2
                                                        var v40 int64 = -9223372036854775808
                                                        _ = v40
                                                    } else {
                                                        if true {
                                                            fmt.Println(v20, fi18)
                                                        }
                                                        v17 = (v17 + (v17 / v17))
                                                    }
                                                }
                                                _ = v38
                                            } else {
                                                v41 := GenS6{F0: 2, F1: 0, F2: ("foo bar" + ("" + "hello"))}
                                                _ = v41
                                            }
                                        }
                                    case -2147483648:
                                        v6 = append(v6, 496)
                                    }
                                default:
                                    switch -784 {
                                    case -2147483648:
                                        v3 -= v3
                                    case 4294967295:
                                        if ((v2 >> ((487) & 63)) > (^(v2))) {
                                            v1 += (v3 - v3)
                                            if (0 == 515) {
                                                switch "foo bar" {
                                                case "":
                                                    if ((!(0 < -125)) || false) {
                                                        switch fi18 {
                                                        case 2147483647:
                                                            v42 := GenS4{F0: -316}
                                                            _ = v42
                                                        case 32767:
                                                            v43 := GenS5{F0: 0, F1: (^(v1))}
                                                            _ = v43
                                                        default:
                                                            v2 = (v2 ^ 42)
                                                        }
                                                        var v44 float64 = v17
                                                        if v19 {
                                                            v6 = append(v6, fi21)
                                                            if ((12 < 9223372036854775807) && false) {
                                                                v45 := GenS7{F0: v1, F1: 18446744073709551615, F2: 0}
                                                                v19 = (v19 && v19)
                                                                _ = v45
                                                            }
                                                        } else {
                                                            v46 := GenS2{F0: 0}
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            v44 -= float64(float32(v17))
                                                            _ = v46
                                                        }
                                                        _ = v44
                                                    }
                                                }
                                                v12 -= float32(-2.25)
                                            }
                                        }
                                    case 127:
                                        _ = len(v6)
                                    default:
                                        var v47 int64 = 474
                                        _ = v47
                                    }
                                }
                                _ = fi21
                                _ = v35
                                _ = v36
                            }
                        }
                        _ = v20
                    } else {
                        switch "a" {
                        case "hello":
                            v48 := GenS8{F0: 2147483648, F1: fi18}
                            _ = v48
                        }
                    }
                } else {
                    var v49 uint32 = (v3 - uint32(v2))
                    fmt.Println(v19, v2)
                    _ = v49
                }
                v50 := GenS9{F0: 390, F1: -2147483648}
                _ = fi18
                _ = v19
                _ = v50
            }
        case "foo bar":
            var v51 int8 = -6
            _ = v51
        case "hello":
            for fi52 := 0; fi52 < 19; fi52++ {
                if (381 > fi52) {
                    v17 = -2.25
                }
                v4 -= (v4 | (v4 << ((4294967295) & 63)))
                _ = fi52
            }
        }
        _ = v17
    }
    v2 = 47
    fmt.Println(v12, v1, v3)
    for fi53 := 0; fi53 < 4; fi53++ {
        switch (fi53 >> ((1) & 63)) {
        case -2147483648:
            fmt.Println(v2, v4)
        default:
            switch ("\n\t" + "hello") {
            case "a":
                _ = len(v6)
            case "hello":
                var v54 uint8 = 42
                _ = v54
            case "":
                for fi55 := 0; fi55 < 11; fi55++ {
                    v56 := GenS9{F0: 2147483647, F1: -557}
                    var v57 float32 = 0.1
                    _ = fi55
                    _ = v56
                    _ = v57
                }
            default:
                var v58 bool = false
                _ = v58
            }
        }
        if false {
            v3 = 506
        }
        _ = fi53
    }
    for fi59 := 0; fi59 < 1; fi59++ {
        v2 += (v2 / ((1) | 1))
        if (81 >= 42) {
            var v60 int64 = 42
            _ = v60
        } else {
            if (("hello" + "") > "\n\t") {
                switch "hello" {
                case "\n\t":
                    v61 := GenS10{F0: v2, F1: 2147483647}
                    _ = v61
                }
            }
            if len(v6) > 0 {
                v6[0] = (fi59 / ((fi59) | 1))
            }
            v2 -= 455
        }
        _ = fi59
    }
    for fi62 := 0; fi62 < 11; fi62++ {
        switch ("foo bar" + "\n\t") {
        case "a":
            for fi63 := 0; fi63 < 7; fi63++ {
                switch "\n\t" {
                case "":
                    switch 701 {
                    case 128:
                        var v64 int = 759
                        _ = v64
                    }
                case "\n\t":
                    v65 := GenS1{F0: "", F1: v1}
                    _ = v65
                default:
                    var v66 int16 = -217
                    _ = v66
                }
                _ = fi63
            }
        case "foo bar":
            switch (fi62 | (fi62 | 127)) {
            case -2147483648:
                if (-648 <= 0) {
                    v67 := GenS11{F0: fi62, F1: v3}
                    v1 = (v1 | v1)
                    _ = v67
                } else {
                    var v68 float32 = 3.14159
                    v4 = (v4 * v4)
                    _ = v68
                }
            case 9223372036854775807:
                var v69 int = fi62
                _ = v69
            case -441:
                for fi70 := 0; fi70 < 15; fi70++ {
                    v71 := GenS12{F0: -352}
                    if false {
                        v72 := GenS7{F0: 2, F1: (v2 >> ((18446744073709551615) & 63)), F2: -354}
                        var v73 int8 = 0
                        _ = v72
                        _ = v73
                    }
                    switch (^(fi62)) {
                    case 128:
                        v2 -= v2
                    case 2147483647:
                        v74 := GenS13{F0: -99, F1: (^(fi62))}
                        _ = v74
                    }
                    if (128 > v4) {
                        var v75 float64 = (0.1 + (0.1 + (-2.25 * (float64(v12) + ((1e10 + (((3.14159 * float64((((1.5 / (1e10 + 0.0)) - 0.0) - ((3.14159 / v12) + 1e10)))) * 1.5) - float64(1e10))) / (-2.25 + 1e10))))))
                        switch "a" {
                        case "a":
                            var v76 bool = (100 <= 42)
                            _ = v76
                        case "foo bar":
                            v4 += v4
                        default:
                            v4 = uint8(v1)
                        }
                        _ = v75
                    } else {
                        if false {
                            var v77 uint32 = 2147483648
                            fmt.Println(v3)
                            v4 = (v4 >> ((255) & 63))
                            _ = v77
                        }
                        if len(v6) > 0 {
                            v6[0] = (-(fi62))
                        }
                    }
                    _ = fi70
                    _ = v71
                }
            }
        case "":
            v78 := GenS14{F0: "foo bar"}
            _ = v78
        default:
            fmt.Println(v12, v2, v2)
        }
        if (594 != 0) {
            fmt.Println(v2)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            if ((((fi62 % ((-9223372036854775808) | 1)) >= (fi62 ^ fi62)) || true) && (fi62 >= (-(fi62)))) {
                var v79 uint16 = 255
                v2 -= (v2 >> ((uint(v3)) & 63))
                v80 := GenS11{F0: fi62, F1: v1}
                _ = v79
                _ = v80
            } else {
                var v81 uint8 = v4
                switch "foo bar" {
                case "":
                    if false {
                        v81 = (v81 - (v4 >> ((234) & 63)))
                        v2 -= v2
                        var v82 float64 = 0.1
                        _ = v82
                    }
                case "\n\t":
                    if false {
                        v81 -= (v4 % (((v81 % (((v4 * v4)) | 1))) | 1))
                        fmt.Println(v12)
                        var v83 float64 = -2.25
                        _ = v83
                    } else {
                        v2 -= 42
                    }
                }
                _ = v81
            }
        } else {
            for fi84 := 0; fi84 < 8; fi84++ {
                switch fi84 {
                case 9223372036854775807:
                    if true {
                        v85 := GenS13{F0: -2147483648, F1: (fi62 >> ((478) & 63))}
                        _ = v85
                    } else {
                        fmt.Println(v1, v4, v2)
                        if len(v6) > 0 {
                            v6[0] = fi62
                        }
                    }
                case -271:
                    v1 += v3
                case 4294967295:
                    fmt.Println(fi62, fi62, fi84)
                }
                if true {
                    var v86 int64 = -440
                    _ = v86
                }
                var v87 int16 = -1
                _ = fi84
                _ = v87
            }
        }
        _ = fi62
    }
    switch -2147483648 {
    case -15:
        for fi88 := 0; fi88 < 13; fi88++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            for fi89 := 0; fi89 < 17; fi89++ {
                v2 = 42
                _ = fi89
            }
            _ = fi88
        }
    case 1:
        var v90 bool = true
        _ = v90
    case 4294967295:
        fmt.Println(v3, v4)
    }
    var v91 int16 = -32768
    switch "foo bar" {
    case "hello":
        v4 = v4
    case "a":
        var v92 bool = (((!(v4 <= uint8(v1))) || true) && ((^(v2)) <= (v2 / ((v2) | 1))))
        _ = v92
    case "\n\t":
        switch "" {
        case "foo bar":
            fmt.Println(v1, v3, v1)
        case "":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case "hello":
            if true {
                if len(v6) > 0 {
                    v6[0] = 1
                }
                v3 += v1
                if true {
                    var v93 uint64 = v2
                    for fi94 := 0; fi94 < 11; fi94++ {
                        if ((^(v2)) >= 18446744073709551615) {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        } else {
                            v95 := GenS15{F0: 42, F1: false}
                            var v96 int64 = 42
                            _ = v95
                            _ = v96
                        }
                        v12 -= v12
                        _ = fi94
                    }
                    _ = v93
                }
            }
        }
    default:
        for fi97 := 0; fi97 < 3; fi97++ {
            v3 -= 65535
            _ = fi97
        }
    }
    for fi98 := 0; fi98 < 4; fi98++ {
        for fi99 := 0; fi99 < 17; fi99++ {
            v100 := GenS14{F0: "hello"}
            _ = fi99
            _ = v100
        }
        if (!true) {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v101 uint32 = v1
            _ = v101
        } else {
            v91 = -886
        }
        if len(v6) > 0 {
            v6[0] = int(v1)
        }
        _ = fi98
    }
    var v102 uint = 1
    for fi103 := 0; fi103 < 7; fi103++ {
        fmt.Println(v1, v102, v102)
        _ = len(v6)
        _ = fi103
    }
    switch -2147483648 {
    case -1:
        for fi104 := 0; fi104 < 17; fi104++ {
            for fi105 := 0; fi105 < 16; fi105++ {
                v91 -= v91
                _ = fi105
            }
            if (0.1 >= v12) {
                fmt.Println(v102, v91, v12)
                var v106 uint8 = (^(v4))
                for fi107 := 0; fi107 < 3; fi107++ {
                    if (float64(v12) == (-2.25 + 0.1)) {
                        switch fi107 {
                        case 2147483647:
                            switch "a" {
                            case "hello":
                                var v108 int64 = 574
                                _ = v108
                            case "a":
                                var v109 int64 = -1
                                _ = v109
                            case "\n\t":
                                if (false || (v2 > (v2 >> ((v102) & 63)))) {
                                    if (0 <= 42) {
                                        switch (fi104 << ((uint(v3)) & 63)) {
                                        case 1:
                                            if ((!(true || (true && true))) && (!((fi104 | (fi104 - (fi104 ^ fi104))) > fi104))) {
                                                _ = len(v6)
                                                v106 -= v106
                                                v110 := GenS13{F0: -1, F1: (fi104 + 9223372036854775807)}
                                                _ = v110
                                            } else {
                                                v111 := GenS16{F0: -128, F1: v102, F2: (fi104 + int(v102))}
                                                v3 = v3
                                                v106 = (v106 + v106)
                                                _ = v111
                                            }
                                        case 127:
                                            var v112 uint16 = 431
                                            _ = v112
                                        case 42:
                                            if ((false && (!(true && false))) && false) {
                                                v113 := GenS13{F0: -4, F1: (fi104 * (fi107 & fi104))}
                                                switch "" {
                                                case "":
                                                    v114 := GenS17{F0: v91, F1: v1}
                                                    _ = v114
                                                case "\n\t":
                                                    v115 := GenS6{F0: 430, F1: (v91 | 1000), F2: ("\n\t" + "a")}
                                                    _ = v115
                                                case "foo bar":
                                                    v116 := GenS2{F0: 32767}
                                                    if false {
                                                        fmt.Println(v4)
                                                        v117 := GenS15{F0: -290, F1: false}
                                                        _ = v117
                                                    }
                                                    v12 = v12
                                                    _ = v116
                                                }
                                                var v118 bool = (true || (!(3.14159 > 0.0)))
                                                if v118 {
                                                    if (1.5 > (1e10 / 0.1)) {
                                                        switch (fi104 / ((fi107) | 1)) {
                                                        case 1:
                                                            v1 = (^(v1))
                                                        case 42:
                                                            fmt.Println(v102, v1, v102)
                                                        default:
                                                            var v119 uint = 14
                                                            _ = v119
                                                        }
                                                        if v118 {
                                                            var v120 int64 = 865
                                                            v121 := GenS17{F0: -32768, F1: (^(v3))}
                                                            var v122 int = fi107
                                                            var v123 int8 = 18
                                                            v91 += -721
                                                            _ = v120
                                                            _ = v121
                                                            _ = v122
                                                            _ = v123
                                                        }
                                                        switch (fi104 / (((fi104 + (fi104 + int(v102)))) | 1)) {
                                                        case 128:
                                                            if len(v6) > 0 {
                                                                v6[0] = (fi107 << ((v102) & 63))
                                                            }
                                                        case -1:
                                                            if ((v102 << ((580) & 63)) != v102) {
                                                                fmt.Println(v102)
                                                            } else {
                                                                v124 := GenS18{F0: v102}
                                                                v125 := GenS19{F0: (v102 * v102), F1: v102, F2: ("hello" + "foo bar")}
                                                                _ = v124
                                                                _ = v125
                                                            }
                                                        case -544:
                                                            var v126 int32 = 0
                                                            _ = v126
                                                        }
                                                    } else {
                                                        if len(v6) > 0 {
                                                            v6[0] = fi107
                                                        }
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                        var v127 bool = v118
                                                        _ = v127
                                                    }
                                                    var v128 bool = (!v118)
                                                    switch ("hello" + ("" + ("" + ("foo bar" + "")))) {
                                                    case "\n\t":
                                                        if true {
                                                            fmt.Println(v128, v102)
                                                        }
                                                    case "a":
                                                        var v129 string = ""
                                                        _ = v129
                                                    }
                                                    _ = v128
                                                }
                                                v130 := GenS11{F0: (fi107 - (fi104 / ((fi104) | 1))), F1: v1}
                                                _ = v113
                                                _ = v118
                                                _ = v130
                                            } else {
                                                v131 := GenS18{F0: (v102 << (((^(v102))) & 63))}
                                                var v132 float32 = 1e10
                                                if ((v102 < (v102 % ((18446744073709551615) | 1))) || false) {
                                                    var v133 string = "hello"
                                                    _ = len(v6)
                                                    if (2147483647 == -1) {
                                                        v3 -= (v3 << ((v102) & 63))
                                                        if (127 <= fi107) {
                                                            var v134 float32 = (((0.0 / 1e10) - float32(0.1)) - 0.1)
                                                            var v135 bool = true
                                                            v136 := GenS15{F0: 0, F1: (v3 == v1)}
                                                            _ = v134
                                                            _ = v135
                                                            _ = v136
                                                        } else {
                                                            _ = len(v6)
                                                            v1 -= (v1 - (v3 - v1))
                                                            v4 = v4
                                                        }
                                                        var v137 int8 = 42
                                                        _ = v137
                                                    } else {
                                                        var v138 int32 = 2147483647
                                                        if true {
                                                            var v139 int64 = -9223372036854775808
                                                            if true {
                                                                v140 := GenS5{F0: -822, F1: 823}
                                                                fmt.Println(v1, v2)
                                                                var v141 float32 = 0.0
                                                                _ = v140
                                                                _ = v141
                                                            } else {
                                                                v139 = 618
                                                                if (-2.25 > (0.0 - 0.0)) {
                                                                    var v142 int8 = -128
                                                                    _ = v142
                                                                }
                                                                v2 = (v2 | v2)
                                                            }
                                                            switch -227 {
                                                            case 32768:
                                                                var v143 uint = 263
                                                                _ = v143
                                                            }
                                                            _ = v139
                                                        } else {
                                                            var v144 int16 = (v91 ^ (v91 * 658))
                                                            v145 := GenS2{F0: 749}
                                                            fmt.Println(v1)
                                                            _ = v144
                                                            _ = v145
                                                        }
                                                        v102 = (^(v102))
                                                        _ = v138
                                                    }
                                                    _ = v133
                                                } else {
                                                    v6 = append(v6, (fi107 * fi104))
                                                    var v146 int64 = -9223372036854775808
                                                    if (true && (int32(v3) == 11)) {
                                                        fmt.Println(v12, v12, v146)
                                                        if (uint16(v102) < uint16(v2)) {
                                                            if len(v6) > 0 {
                                                                v6[0] = 723
                                                            }
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            if true {
                                                                _ = len(v6)
                                                                v147 := GenS18{F0: (v102 >> ((4294967295) & 63))}
                                                                _ = v147
                                                            } else {
                                                                var v148 bool = false
                                                                if (0.0 <= (v12 * v132)) {
                                                                    var v149 uint64 = v2
                                                                    _ = v149
                                                                } else {
                                                                    v148 = (-1 < -70)
                                                                }
                                                                fmt.Println(v12)
                                                                _ = v148
                                                            }
                                                        }
                                                        _ = len(v6)
                                                        var v150 int64 = v146
                                                        _ = v150
                                                    } else {
                                                        switch "hello" {
                                                        case "foo bar":
                                                            v2 = (^(v2))
                                                        case "a":
                                                            var v151 int16 = (v91 >> ((v102) & 63))
                                                            _ = v151
                                                        }
                                                        v3 += (v3 + 761)
                                                        if len(v6) > 0 {
                                                            v6[0] = (^(fi104))
                                                        }
                                                        if (true || true) {
                                                            v12 = v12
                                                        } else {
                                                            var v152 float32 = v132
                                                            if (889 <= 422) {
                                                                fmt.Println(v1, v1, v2)
                                                                if (9223372036854775807 > (fi104 >> ((uint(v1)) & 63))) {
                                                                    v153 := GenS13{F0: 1, F1: (fi107 | fi104)}
                                                                    if ((^(v102)) < 65535) {
                                                                        v3 += v3
                                                                        var v154 uint64 = uint64(v106)
                                                                        _ = v154
                                                                    } else {
                                                                        if ((v4 != (v4 & (v4 * v4))) || (true && ((!(-9223372036854775808 <= v146)) || false))) {
                                                                            _ = len(v6)
                                                                        } else {
                                                                            var v155 string = "foo bar"
                                                                            _ = v155
                                                                        }
                                                                        if false {
                                                                            var v156 uint8 = v106
                                                                            _ = v156
                                                                        } else {
                                                                            switch fi104 {
                                                                            case -25:
                                                                                v157 := GenS20{F0: 42, F1: -1, F2: v4}
                                                                                _ = v157
                                                                            case 32767:
                                                                                var v158 int = -328
                                                                                _ = v158
                                                                            case 282:
                                                                                v132 -= v152
                                                                            }
                                                                            v159 := GenS21{F0: fi104, F1: 384}
                                                                            _ = v159
                                                                        }
                                                                    }
                                                                    _ = v153
                                                                }
                                                                var v160 int8 = -1
                                                                _ = v160
                                                            } else {
                                                                var v161 int8 = 127
                                                                _ = len(v6)
                                                                v4 = (^(v106))
                                                                _ = v161
                                                            }
                                                            _ = v152
                                                        }
                                                    }
                                                    var v162 int32 = -1
                                                    switch fi107 {
                                                    case -427:
                                                        if true {
                                                            v146 = 9223372036854775807
                                                        }
                                                    case -2147483648:
                                                        v163 := GenS22{F0: (v146 - v146), F1: v91}
                                                        _ = v163
                                                    case -42:
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    default:
                                                        v91 = 984
                                                    }
                                                    v102 = (v102 % (((v102 & 0)) | 1))
                                                    _ = v146
                                                    _ = v162
                                                }
                                                _ = v131
                                                _ = v132
                                            }
                                        default:
                                            var v164 int32 = -482
                                            _ = v164
                                        }
                                        if (false || true) {
                                            v6 = append(v6, (-(fi104)))
                                        }
                                        v12 = (v12 + float32(1.5))
                                    }
                                    var v165 float64 = 0.1
                                    _ = v165
                                }
                            default:
                                v3 = 65535
                            }
                            var v166 uint64 = 2
                            switch (fi104 - -974) {
                            case 0:
                                var v167 int64 = 9223372036854775807
                                _ = v167
                            default:
                                v166 = 42
                            }
                            if (0 != 529) {
                                var v168 int8 = 127
                                _ = v168
                            } else {
                                v1 -= v3
                                if (!false) {
                                    if false {
                                        v4 = (^(v4))
                                    }
                                    v106 = (v4 + (v4 % ((v4) | 1)))
                                } else {
                                    var v169 uint = v102
                                    _ = v169
                                }
                            }
                            _ = v166
                        default:
                            _ = len(v6)
                        }
                        var v170 uint64 = (v2 >> (((v102 / ((v102) | 1))) & 63))
                        fmt.Println(v106, fi107, v3)
                        _ = len(v6)
                        _ = v170
                    } else {
                        fmt.Println(v91, v91)
                        v106 = (v4 * (^(v4)))
                        v12 = v12
                    }
                    switch (^(fi104)) {
                    case 1:
                        v91 += (-(v91))
                    case -1:
                        var v171 int = fi104
                        _ = v171
                    case -19:
                        v172 := GenS20{F0: -37, F1: -128, F2: v106}
                        _ = v172
                    }
                    _ = fi107
                }
                _ = v106
            }
            v173 := GenS20{F0: 1, F1: 0, F2: v4}
            _ = fi104
            _ = v173
        }
    case -2:
        var v174 int64 = 1
        _ = v174
    case 737:
        v175 := GenS23{F0: 0, F1: v1}
        _ = v175
    default:
        var v176 uint64 = (^(v2))
        _ = v176
    }
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v102
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v91
}
