package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 string; F2 uint }
type GenS5 struct { F0 int16; F1 uint }
type GenS6 struct { F0 uint; F1 uint64; F2 uint }
type GenS7 struct { F0 uint; F1 uint16 }
type GenS8 struct { F0 int32; F1 int; F2 uint }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint; F1 uint32 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int32; F1 bool }
type GenS13 struct { F0 uint32; F1 int8 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint16; F1 uint64 }
type GenS16 struct { F0 bool; F1 uint; F2 int64 }
type GenS17 struct { F0 uint8; F1 int }
type GenS18 struct { F0 bool; F1 int8 }
type GenS19 struct { F0 uint32; F1 int64 }
type GenS20 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS21 struct { F0 int32; F1 string }
type GenS22 struct { F0 int64; F1 bool }
type GenS23 struct { F0 int8; F1 uint8; F2 int }
type GenS24 struct { F0 int8; F1 int8; F2 uint64 }
type GenS25 struct { F0 int16 }
type GenS26 struct { F0 bool; F1 uint64 }
type GenS27 struct { F0 uint }
type GenS28 struct { F0 uint16; F1 string }
type GenS29 struct { F0 int64; F1 int64; F2 int64 }
type GenS30 struct { F0 int; F1 uint32 }
type GenS31 struct { F0 int64; F1 int; F2 string }
type GenS32 struct { F0 string; F1 string; F2 int }
type GenS33 struct { F0 int16; F1 uint64 }
type GenS34 struct { F0 int64; F1 uint }
type GenS35 struct { F0 bool }
type GenS36 struct { F0 bool; F1 string }
type GenS37 struct { F0 uint64; F1 int16 }
type GenS38 struct { F0 uint16; F1 uint16 }
type GenS39 struct { F0 int32; F1 int64; F2 int16 }
func genFn0(p0 uint8) int16 {
    fmt.Println(p0)
    var v6 int64 = 1
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int32 = -568
        v8 -= -850
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 1}
    _ = p0
    _ = v6
    _ = v9
    return -1
}
func genFn1() uint32 {
    if false {
        var v10 int32 = 0
        _ = v10
    } else {
        if true {
            var v11 uint = 0
            _ = v11
        } else {
            for fi12 := 0; fi12 < 13; fi12++ {
                fmt.Println(fi12)
                var v13 uint32 = 2
                v14 := GenS2{F0: "foo bar", F1: -9223372036854775808}
                _ = fi12
                _ = v13
                _ = v14
            }
            var v15 uint16 = 32767
            _ = v15
        }
        if (2 != 128) {
            switch "a" {
            case "hello":
                v16 := GenS3{F0: false}
                _ = v16
            case "":
                for fi17 := 0; fi17 < 17; fi17++ {
                    var v18 uint64 = 348
                    if true {
                        fmt.Println(v18, fi17)
                        if (0 > 2147483647) {
                            if true {
                                if (-9223372036854775808 > 351) {
                                    switch (fi17 + (fi17 & int(v18))) {
                                    case 9223372036854775807:
                                        var v19 float32 = (float32(-2.25) * float32(float64(3.14159)))
                                        _ = v19
                                    default:
                                        var v20 uint16 = 879
                                        _ = v20
                                    }
                                    v21 := []int{}
                                    _ = v21
                                }
                                fmt.Println(fi17)
                            }
                            var v22 int = fi17
                            if (127 >= -125) {
                                v22 += 128
                                var v23 int16 = 0
                                _ = v23
                            } else {
                                v18 += (v18 % (((v18 + v18)) | 1))
                                for fi24 := 0; fi24 < 12; fi24++ {
                                    v22 -= (fi17 - v22)
                                    if (("a" == ((("\n\t" + ("a" + "\n\t")) + "foo bar") + "\n\t")) && false) {
                                        v25 := GenS1{F0: -895}
                                        _ = v25
                                    }
                                    _ = fi24
                                }
                                v26 := GenS1{F0: -1}
                                _ = v26
                            }
                            _ = v22
                        } else {
                            v18 = 984
                            v27 := GenS2{F0: "", F1: 9223372036854775807}
                            fmt.Println(v18, v18)
                            _ = v27
                        }
                        var v28 uint32 = 42
                        _ = v28
                    } else {
                        var v29 int32 = 1
                        var v30 uint = 4294967295
                        switch (^(fi17)) {
                        case 331:
                            v30 = (v30 | v30)
                        case -2147483648:
                            v31 := GenS1{F0: (v29 >> (((v30 * 65535)) & 63))}
                            _ = v31
                        default:
                            v32 := GenS4{F0: 4294967295, F1: "", F2: (^(v30))}
                            _ = v32
                        }
                        _ = v29
                        _ = v30
                    }
                    var v33 int32 = 1
                    _ = fi17
                    _ = v18
                    _ = v33
                }
            }
            if true {
                v34 := []int{42, 32768}
                var v35 uint64 = 16
                var v36 float32 = ((0.1 * (0.0 + 1.5)) - -2.25)
                _ = v34
                _ = v35
                _ = v36
            } else {
                for fi37 := 0; fi37 < 6; fi37++ {
                    v38 := []int{fi37}
                    if false {
                        var v39 int8 = 127
                        fmt.Println(fi37, v39)
                        var v40 bool = true
                        _ = v39
                        _ = v40
                    } else {
                        if (0 >= 856) {
                            var v41 string = "a"
                            switch "\n\t" {
                            case "hello":
                                if false {
                                    v42 := GenS2{F0: v41, F1: 1}
                                    switch (fi37 ^ -935) {
                                    case -2147483648:
                                        var v43 uint64 = 324
                                        _ = v43
                                    case 42:
                                        if (!(148 >= 42)) {
                                            if false {
                                                switch "" {
                                                case "a":
                                                    if true {
                                                        var v44 int = fi37
                                                        var v45 string = v41
                                                        _ = v44
                                                        _ = v45
                                                    } else {
                                                        switch v41 {
                                                        case "hello":
                                                            switch (fi37 + (fi37 & (fi37 + 666))) {
                                                            case -772:
                                                                if len(v38) >= 2 {
                                                                    v38 = v38[1:len(v38)]
                                                                }
                                                            case 127:
                                                                if (!(0 > 2147483647)) {
                                                                    if true {
                                                                        var v46 int16 = 42
                                                                        _ = v46
                                                                    } else {
                                                                        v47 := GenS5{F0: -1, F1: 4294967295}
                                                                        if (325 >= 401) {
                                                                            for fi48 := 0; fi48 < 19; fi48++ {
                                                                                var v49 float32 = ((1.5 / 1e10) / 0.1)
                                                                                _ = fi48
                                                                                _ = v49
                                                                            }
                                                                            for fi50 := 0; fi50 < 16; fi50++ {
                                                                                switch fi50 {
                                                                                case -220:
                                                                                    if true {
                                                                                        _ = len(v38)
                                                                                        fmt.Println(fi37)
                                                                                    }
                                                                                case -971:
                                                                                    v41 = "hello"
                                                                                default:
                                                                                    if len(v38) >= 2 {
                                                                                        v38 = v38[1:len(v38)]
                                                                                    }
                                                                                    var v51 int8 = int8(fi50)
                                                                                    _ = v51
                                                                                }
                                                                                _ = fi50
                                                                            }
                                                                        }
                                                                        if true {
                                                                            var v52 uint16 = 629
                                                                            _ = v52
                                                                        }
                                                                        _ = v47
                                                                    }
                                                                    for fi53 := 0; fi53 < 12; fi53++ {
                                                                        if (0 <= 0) {
                                                                            var v54 uint16 = 169
                                                                            fmt.Println(fi37, fi37, fi37)
                                                                            _ = len(v38)
                                                                            _ = v54
                                                                        }
                                                                        v41 = v41
                                                                        fmt.Println(v41)
                                                                        var v55 int32 = 0
                                                                        _ = fi53
                                                                        _ = v55
                                                                    }
                                                                    v41 = ""
                                                                } else {
                                                                    if (((v41 + (("a" + "foo bar") + v41)) + v41) != (v41 + "foo bar")) {
                                                                        var v56 uint64 = 4294967295
                                                                        v41 = v41
                                                                        for fi57 := 0; fi57 < 10; fi57++ {
                                                                            fmt.Println(fi37)
                                                                            v38 = append(v38, (fi57 % (((fi57 & fi37)) | 1)))
                                                                            v41 = ("foo bar" + v41)
                                                                            _ = fi57
                                                                        }
                                                                        _ = v56
                                                                    }
                                                                    var v58 int16 = -479
                                                                    _ = v58
                                                                }
                                                            case -981:
                                                                var v59 int16 = 42
                                                                _ = v59
                                                            default:
                                                                v60 := GenS5{F0: 0, F1: 290}
                                                                _ = v60
                                                            }
                                                        case "a":
                                                            for fi61 := 0; fi61 < 6; fi61++ {
                                                                if len(v38) >= 2 {
                                                                    v38 = v38[1:len(v38)]
                                                                }
                                                                v41 = v41
                                                                _ = fi61
                                                            }
                                                        case "foo bar":
                                                            for fi62 := 0; fi62 < 15; fi62++ {
                                                                switch v41 {
                                                                case "a":
                                                                    var v63 bool = true
                                                                    _ = v63
                                                                case "\n\t":
                                                                    fmt.Println(v41, fi37)
                                                                }
                                                                v64 := GenS1{F0: 0}
                                                                _ = fi62
                                                                _ = v64
                                                            }
                                                        default:
                                                            v41 = v41
                                                        }
                                                        v65 := GenS6{F0: 4294967295, F1: 4294967295, F2: 4294967295}
                                                        v66 := GenS7{F0: 544, F1: 255}
                                                        _ = v65
                                                        _ = v66
                                                    }
                                                case "":
                                                    switch "hello" {
                                                    case "\n\t":
                                                        if false {
                                                            _ = len(v38)
                                                            fmt.Println(fi37)
                                                        } else {
                                                            if (0 < -1) {
                                                                for fi67 := 0; fi67 < 11; fi67++ {
                                                                    v38 = append(v38, fi37)
                                                                    if (2 < 443) {
                                                                        v41 = "foo bar"
                                                                        switch ("" + ((v41 + v41) + v41)) {
                                                                        case "\n\t":
                                                                            switch v41 {
                                                                            case "\n\t":
                                                                                if false {
                                                                                    var v68 uint = 959
                                                                                    _ = v68
                                                                                } else {
                                                                                    var v69 uint8 = 1
                                                                                    _ = v69
                                                                                }
                                                                            case "hello":
                                                                                switch v41 {
                                                                                case "a":
                                                                                    if len(v38) >= 2 {
                                                                                        v38 = v38[1:len(v38)]
                                                                                    }
                                                                                }
                                                                            }
                                                                            if (2 >= 2) {
                                                                                if (!false) {
                                                                                    var v70 float32 = 0.0
                                                                                    _ = v70
                                                                                } else {
                                                                                    v41 = v41
                                                                                    v38 = append(v38, 82)
                                                                                    switch v41 {
                                                                                    case "\n\t":
                                                                                        var v71 int16 = 529
                                                                                        _ = v71
                                                                                    default:
                                                                                        v72 := GenS6{F0: 2, F1: 1, F2: 1}
                                                                                        _ = v72
                                                                                    }
                                                                                    var v73 int64 = 238
                                                                                    if true {
                                                                                        if len(v38) > 0 {
                                                                                            v38[0] = (fi37 ^ -569)
                                                                                        }
                                                                                    }
                                                                                    _ = v73
                                                                                }
                                                                                if len(v38) > 0 {
                                                                                    v38[0] = fi67
                                                                                }
                                                                                if false {
                                                                                    var v74 uint16 = 32767
                                                                                    _ = v74
                                                                                } else {
                                                                                    switch fi67 {
                                                                                    case -2147483648:
                                                                                        v41 = v41
                                                                                    case 515:
                                                                                        switch v41 {
                                                                                        case "":
                                                                                            switch ((((v41 + v41) + (v41 + v41)) + ("hello" + v41)) + ("foo bar" + (((((v41 + (v41 + ((v41 + v41) + ((v41 + v41) + v41)))) + "a") + "\n\t") + "") + "a"))) {
                                                                                            case "\n\t":
                                                                                                fmt.Println(v41, fi67, fi67)
                                                                                            case "":
                                                                                                v38 = append(v38, (fi37 + (fi67 & -1)))
                                                                                            }
                                                                                        }
                                                                                    case 12:
                                                                                        v75 := GenS2{F0: ((v41 + v41) + v41), F1: 0}
                                                                                        _ = v75
                                                                                    default:
                                                                                        v41 = (v41 + ("a" + (("\n\t" + v41) + v41)))
                                                                                    }
                                                                                    v76 := GenS8{F0: -2147483648, F1: (fi67 % (((fi37 >> ((971) & 63))) | 1)), F2: 255}
                                                                                    _ = v76
                                                                                }
                                                                            }
                                                                            var v77 float32 = 1e10
                                                                            _ = v77
                                                                        default:
                                                                            fmt.Println(fi67, fi37, fi67)
                                                                        }
                                                                        var v78 uint64 = 620
                                                                        _ = v78
                                                                    }
                                                                    if (!true) {
                                                                        var v79 uint64 = 2
                                                                        _ = v79
                                                                    } else {
                                                                        var v80 int64 = -9223372036854775808
                                                                        _ = v80
                                                                    }
                                                                    v41 = (v41 + v41)
                                                                    _ = fi67
                                                                }
                                                                v81 := GenS9{F0: 0}
                                                                _ = v81
                                                            }
                                                            fmt.Println(fi37, v41)
                                                        }
                                                    case "hello":
                                                        var v82 string = ""
                                                        _ = v82
                                                    }
                                                case "foo bar":
                                                    v41 = ((v41 + ("a" + v41)) + (v41 + v41))
                                                }
                                                for fi83 := 0; fi83 < 13; fi83++ {
                                                    fmt.Println(fi37, fi83, v41)
                                                    var v84 int = fi37
                                                    if false {
                                                        if false {
                                                            var v85 bool = false
                                                            v86 := GenS7{F0: 65535, F1: 2}
                                                            if true {
                                                                switch ((v41 + v41) + v41) {
                                                                case "":
                                                                    if v85 {
                                                                        if (v85 || (false && false)) {
                                                                            switch ("foo bar" + v41) {
                                                                            case "a":
                                                                                v84 = (v84 & -9223372036854775808)
                                                                            case "\n\t":
                                                                                v84 -= -9223372036854775808
                                                                            case "hello":
                                                                                fmt.Println(v85)
                                                                            default:
                                                                                v84 += v84
                                                                                v87 := GenS10{F0: 616, F1: 248}
                                                                                _ = v87
                                                                            }
                                                                            if len(v38) > 0 {
                                                                                v38[0] = v84
                                                                            }
                                                                        }
                                                                        v88 := GenS4{F0: 42, F1: ((v41 + v41) + "foo bar"), F2: 4294967295}
                                                                        v89 := GenS4{F0: 1, F1: v41, F2: 381}
                                                                        _ = v88
                                                                        _ = v89
                                                                    } else {
                                                                        if (!v85) {
                                                                            v84 += (v84 - fi83)
                                                                            switch fi37 {
                                                                            case 32767:
                                                                                if (1 < 18446744073709551615) {
                                                                                    v90 := GenS4{F0: 2147483647, F1: v41, F2: 4294967295}
                                                                                    _ = v90
                                                                                } else {
                                                                                    v91 := GenS6{F0: 656, F1: 42, F2: 0}
                                                                                    var v92 int = (fi83 >> ((239) & 63))
                                                                                    v92 += (^(fi37))
                                                                                    _ = v91
                                                                                    _ = v92
                                                                                }
                                                                            case 807:
                                                                                var v93 uint8 = 168
                                                                                _ = v93
                                                                            }
                                                                            switch 1 {
                                                                            case 32768:
                                                                                v84 = 32768
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v85, v84)
                                                                        }
                                                                        v41 = "hello"
                                                                    }
                                                                case "\n\t":
                                                                    var v94 float32 = 1e10
                                                                    _ = v94
                                                                }
                                                            }
                                                            _ = v85
                                                            _ = v86
                                                        } else {
                                                            fmt.Println(fi83)
                                                            var v95 uint16 = 924
                                                            _ = v95
                                                        }
                                                        if len(v38) >= 2 {
                                                            v38 = v38[1:len(v38)]
                                                        }
                                                        fmt.Println(fi37)
                                                    }
                                                    var v96 float64 = (1.5 / 1e10)
                                                    v97 := GenS11{F0: v41}
                                                    _ = fi83
                                                    _ = v84
                                                    _ = v96
                                                    _ = v97
                                                }
                                                if true {
                                                    switch v41 {
                                                    case "a":
                                                        switch (fi37 >> ((1) & 63)) {
                                                        case 2147483647:
                                                            var v98 int = (fi37 - fi37)
                                                            _ = v98
                                                        case 9223372036854775807:
                                                            fmt.Println(fi37)
                                                        default:
                                                            var v99 uint32 = 42
                                                            _ = v99
                                                        }
                                                    case "\n\t":
                                                        var v100 string = v41
                                                        _ = v100
                                                    }
                                                    v41 = v41
                                                    var v101 uint = 4294967295
                                                    _ = v101
                                                }
                                            } else {
                                                if true {
                                                    var v102 float64 = 0.1
                                                    fmt.Println(v102, fi37, v41)
                                                    _ = v102
                                                }
                                                switch fi37 {
                                                case 32768:
                                                    for fi103 := 0; fi103 < 6; fi103++ {
                                                        switch 302 {
                                                        case 885:
                                                            v104 := GenS1{F0: 0}
                                                            _ = v104
                                                        case -104:
                                                            fmt.Println(v41, fi103)
                                                        case 32767:
                                                            var v105 uint64 = 1
                                                            _ = v105
                                                        }
                                                        if ((1 > 42) || (-2.25 != 1.5)) {
                                                            v106 := GenS12{F0: 2147483647, F1: false}
                                                            _ = v106
                                                        } else {
                                                            var v107 uint32 = 2
                                                            switch (-(fi103)) {
                                                            case 2147483647:
                                                                if (((^(v107)) <= 404) && (true || false)) {
                                                                    if true {
                                                                        var v108 int16 = 1
                                                                        var v109 uint64 = 4294967295
                                                                        v109 += 391
                                                                        _ = v108
                                                                        _ = v109
                                                                    }
                                                                    v110 := GenS11{F0: v41}
                                                                    var v111 int32 = 0
                                                                    v112 := GenS4{F0: (v107 % ((0) | 1)), F1: "hello", F2: 1}
                                                                    _ = len(v38)
                                                                    _ = v110
                                                                    _ = v111
                                                                    _ = v112
                                                                } else {
                                                                    v41 = "a"
                                                                    switch v41 {
                                                                    case "":
                                                                        if len(v38) > 0 {
                                                                            v38[0] = fi37
                                                                        }
                                                                    case "foo bar":
                                                                        v41 = "\n\t"
                                                                        switch "a" {
                                                                        case "a":
                                                                            _ = len(v38)
                                                                            switch ("foo bar" + (v41 + "hello")) {
                                                                            case "hello":
                                                                                if ((174 < 127) || (205 == 128)) {
                                                                                    v107 = v107
                                                                                    v107 -= 0
                                                                                    v41 = v41
                                                                                }
                                                                            case "\n\t":
                                                                                _ = len(v38)
                                                                            }
                                                                        case "":
                                                                            var v113 uint64 = 257
                                                                            _ = v113
                                                                        default:
                                                                            if false {
                                                                                v114 := GenS6{F0: 240, F1: 445, F2: 255}
                                                                                _ = v114
                                                                            } else {
                                                                                v107 = v107
                                                                            }
                                                                        }
                                                                        v41 = "foo bar"
                                                                        if (!true) {
                                                                            v115 := GenS12{F0: 1, F1: false}
                                                                            _ = v115
                                                                        } else {
                                                                            switch fi37 {
                                                                            case 1:
                                                                                switch fi37 {
                                                                                case 782:
                                                                                    if (!(!false)) {
                                                                                        fmt.Println(v41, fi37, fi103)
                                                                                        if (!true) {
                                                                                            _ = len(v38)
                                                                                        } else {
                                                                                            var v116 float32 = 3.14159
                                                                                            v116 -= 0.0
                                                                                            v117 := GenS13{F0: 421, F1: 127}
                                                                                            if len(v38) > 0 {
                                                                                                v38[0] = fi37
                                                                                            }
                                                                                            v118 := GenS14{F0: -159}
                                                                                            _ = v116
                                                                                            _ = v117
                                                                                            _ = v118
                                                                                        }
                                                                                        var v119 bool = true
                                                                                        _ = v119
                                                                                    } else {
                                                                                        if false {
                                                                                            var v120 uint8 = 17
                                                                                            v120 += (^(v120))
                                                                                            _ = v120
                                                                                        } else {
                                                                                            switch v41 {
                                                                                            case "hello":
                                                                                                if (401 != 42) {
                                                                                                    fmt.Println(v41, v107)
                                                                                                    if true {
                                                                                                        switch ("\n\t" + (((v41 + "") + v41) + v41)) {
                                                                                                        case "foo bar":
                                                                                                            var v121 int8 = -91
                                                                                                            _ = v121
                                                                                                        case "\n\t":
                                                                                                            var v122 int16 = -364
                                                                                                            _ = v122
                                                                                                        case "":
                                                                                                            if ((1.5 - (0.1 - (0.0 - 3.14159))) <= float64(0.1)) {
                                                                                                                switch fi103 {
                                                                                                                case 127:
                                                                                                                    var v123 string = ""
                                                                                                                    _ = v123
                                                                                                                case 1:
                                                                                                                    fmt.Println(fi37)
                                                                                                                default:
                                                                                                                    if false {
                                                                                                                        if ((fi103 > (fi37 >> ((uint(v107)) & 63))) || ((1 > 1) && (127 >= 0))) {
                                                                                                                            if len(v38) > 0 {
                                                                                                                                v38[0] = -2147483648
                                                                                                                            }
                                                                                                                            v41 = v41
                                                                                                                        } else {
                                                                                                                            if (1.5 <= ((3.14159 * -2.25) + 0.1)) {
                                                                                                                                switch v41 {
                                                                                                                                case "":
                                                                                                                                    switch (v41 + (("a" + ((v41 + ("\n\t" + "a")) + (v41 + ((v41 + v41) + (((v41 + "foo bar") + "hello") + "hello"))))) + (v41 + v41))) {
                                                                                                                                    case "\n\t":
                                                                                                                                        v38 = append(v38, -2147483648)
                                                                                                                                    }
                                                                                                                                    if len(v38) >= 2 {
                                                                                                                                        v38 = v38[1:len(v38)]
                                                                                                                                    }
                                                                                                                                case "hello":
                                                                                                                                    var v124 int = -602
                                                                                                                                    _ = v124
                                                                                                                                }
                                                                                                                                var v125 float32 = (0.1 - -2.25)
                                                                                                                                v126 := GenS13{F0: 65535, F1: -1}
                                                                                                                                _ = v125
                                                                                                                                _ = v126
                                                                                                                            }
                                                                                                                            fmt.Println(fi37, fi103, v41)
                                                                                                                        }
                                                                                                                        switch v41 {
                                                                                                                        case "foo bar":
                                                                                                                            v107 -= v107
                                                                                                                        case "\n\t":
                                                                                                                            if (!false) {
                                                                                                                                if (-32768 > 676) {
                                                                                                                                    v107 = v107
                                                                                                                                    v127 := GenS5{F0: 32767, F1: 4294967295}
                                                                                                                                    v41 = "foo bar"
                                                                                                                                    _ = v127
                                                                                                                                } else {
                                                                                                                                    v128 := GenS15{F0: 2, F1: 1}
                                                                                                                                    var v129 uint16 = 255
                                                                                                                                    var v130 uint16 = v129
                                                                                                                                    var v131 uint16 = 262
                                                                                                                                    if false {
                                                                                                                                        var v132 int8 = 42
                                                                                                                                        v130 -= 629
                                                                                                                                        switch v41 {
                                                                                                                                        case "foo bar":
                                                                                                                                            if len(v38) >= 2 {
                                                                                                                                                v38 = v38[1:len(v38)]
                                                                                                                                            }
                                                                                                                                        case "a":
                                                                                                                                            var v133 int32 = 0
                                                                                                                                            _ = v133
                                                                                                                                        }
                                                                                                                                        _ = v132
                                                                                                                                    } else {
                                                                                                                                        var v134 int8 = 1
                                                                                                                                        _ = v134
                                                                                                                                    }
                                                                                                                                    _ = v128
                                                                                                                                    _ = v129
                                                                                                                                    _ = v130
                                                                                                                                    _ = v131
                                                                                                                                }
                                                                                                                                v41 = "foo bar"
                                                                                                                            }
                                                                                                                        case "a":
                                                                                                                            fmt.Println(v107)
                                                                                                                            v135 := GenS5{F0: -825, F1: 65535}
                                                                                                                            _ = v135
                                                                                                                        }
                                                                                                                        var v136 int16 = -825
                                                                                                                        if true {
                                                                                                                            v38 = append(v38, (fi37 ^ fi37))
                                                                                                                            v137 := GenS6{F0: 0, F1: 18446744073709551615, F2: uint(v107)}
                                                                                                                            if len(v38) >= 2 {
                                                                                                                                v38 = v38[1:len(v38)]
                                                                                                                            }
                                                                                                                            _ = v137
                                                                                                                        } else {
                                                                                                                            v107 = (v107 >> ((430) & 63))
                                                                                                                            v138 := GenS16{F0: false, F1: 42, F2: -1}
                                                                                                                            var v139 int64 = 42
                                                                                                                            _ = v138
                                                                                                                            _ = v139
                                                                                                                        }
                                                                                                                        var v140 uint32 = (v107 ^ v107)
                                                                                                                        _ = v136
                                                                                                                        _ = v140
                                                                                                                    } else {
                                                                                                                        switch v41 {
                                                                                                                        case "hello":
                                                                                                                            switch v41 {
                                                                                                                            case "hello":
                                                                                                                                v38 = append(v38, (fi37 / (((fi37 << ((1) & 63))) | 1)))
                                                                                                                            }
                                                                                                                        }
                                                                                                                        var v141 string = (((v41 + v41) + (v41 + ("hello" + (v41 + v41)))) + (v41 + "a"))
                                                                                                                        _ = v141
                                                                                                                    }
                                                                                                                }
                                                                                                                v107 = v107
                                                                                                                v41 = "a"
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    v142 := GenS11{F0: v41}
                                                                                                    if (-1 != -2147483648) {
                                                                                                        var v143 float64 = (0.0 - 1e10)
                                                                                                        fmt.Println(fi103, v41, fi103)
                                                                                                        if false {
                                                                                                            var v144 int = (-(fi37))
                                                                                                            _ = v144
                                                                                                        }
                                                                                                        if len(v38) >= 2 {
                                                                                                            v38 = v38[1:len(v38)]
                                                                                                        }
                                                                                                        v41 = v41
                                                                                                        _ = v143
                                                                                                    }
                                                                                                    _ = v142
                                                                                                }
                                                                                                if (false || false) {
                                                                                                    if (true || (false || (fi37 >= -84))) {
                                                                                                        if len(v38) >= 2 {
                                                                                                            v38 = v38[1:len(v38)]
                                                                                                        }
                                                                                                        v41 = (("foo bar" + ("" + (v41 + ("a" + (v41 + ""))))) + (v41 + (((("foo bar" + v41) + "foo bar") + v41) + ("" + v41))))
                                                                                                        var v145 uint16 = 668
                                                                                                        _ = v145
                                                                                                    }
                                                                                                    v41 = v41
                                                                                                }
                                                                                                if len(v38) > 0 {
                                                                                                    v38[0] = (fi103 >> ((447) & 63))
                                                                                                }
                                                                                            case "a":
                                                                                                if ((fi37 / (((-(fi37))) | 1)) < (fi37 << ((41) & 63))) {
                                                                                                    if len(v38) > 0 {
                                                                                                        v38[0] = 4294967295
                                                                                                    }
                                                                                                    var v146 uint16 = uint16(v107)
                                                                                                    _ = v146
                                                                                                }
                                                                                                v41 = v41
                                                                                                var v147 uint8 = 0
                                                                                                switch (fi37 ^ (-(fi103))) {
                                                                                                case 183:
                                                                                                    if len(v38) > 0 {
                                                                                                        v38[0] = (-(fi103))
                                                                                                    }
                                                                                                }
                                                                                                v148 := GenS4{F0: (v107 / (((^(v107))) | 1)), F1: v41, F2: 42}
                                                                                                _ = v147
                                                                                                _ = v148
                                                                                            case "foo bar":
                                                                                                switch (fi103 - -744) {
                                                                                                case 128:
                                                                                                    if (false && (!false)) {
                                                                                                        fmt.Println(v107)
                                                                                                    }
                                                                                                    v149 := GenS17{F0: 1, F1: fi103}
                                                                                                    v150 := GenS18{F0: true, F1: 21}
                                                                                                    _ = v149
                                                                                                    _ = v150
                                                                                                case 1:
                                                                                                    if false {
                                                                                                        if (!(false || true)) {
                                                                                                            v107 += (v107 >> ((2) & 63))
                                                                                                            v41 = v41
                                                                                                        } else {
                                                                                                            var v151 int32 = -2147483648
                                                                                                            var v152 int64 = 0
                                                                                                            _ = v151
                                                                                                            _ = v152
                                                                                                        }
                                                                                                        if len(v38) >= 2 {
                                                                                                            v38 = v38[1:len(v38)]
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    v153 := GenS19{F0: (v107 + v107), F1: 42}
                                                                                                    _ = v153
                                                                                                }
                                                                                            }
                                                                                            v41 = "\n\t"
                                                                                            if ((false && true) || (42 < 2)) {
                                                                                                fmt.Println(v41, v107, v41)
                                                                                            } else {
                                                                                                if len(v38) >= 2 {
                                                                                                    v38 = v38[1:len(v38)]
                                                                                                }
                                                                                                if len(v38) >= 2 {
                                                                                                    v38 = v38[1:len(v38)]
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    if (false && false) {
                                                                                        v41 = v41
                                                                                        v41 = v41
                                                                                        v38 = append(v38, (fi37 / (((fi103 % ((592) | 1))) | 1)))
                                                                                    }
                                                                                    if true {
                                                                                        switch "\n\t" {
                                                                                        case "\n\t":
                                                                                            if (!true) {
                                                                                                v154 := GenS20{F0: 2, F1: true, F2: (v107 - 742)}
                                                                                                _ = v154
                                                                                            }
                                                                                        case "a":
                                                                                            switch (^(fi37)) {
                                                                                            case -847:
                                                                                                if true {
                                                                                                    v41 = v41
                                                                                                    if ((!false) || true) {
                                                                                                        if true {
                                                                                                            var v155 uint16 = 674
                                                                                                            _ = v155
                                                                                                        } else {
                                                                                                            var v156 uint64 = uint64(v107)
                                                                                                            var v157 bool = (!(!false))
                                                                                                            v107 = v107
                                                                                                            _ = v156
                                                                                                            _ = v157
                                                                                                        }
                                                                                                    } else {
                                                                                                        if (true && ((!true) && false)) {
                                                                                                            var v158 float64 = -2.25
                                                                                                            v107 += 4294967295
                                                                                                            fmt.Println(v107, v107, fi103)
                                                                                                            if len(v38) >= 2 {
                                                                                                                v38 = v38[1:len(v38)]
                                                                                                            }
                                                                                                            v159 := GenS21{F0: -2147483648, F1: v41}
                                                                                                            _ = v158
                                                                                                            _ = v159
                                                                                                        }
                                                                                                        if len(v38) > 0 {
                                                                                                            v38[0] = fi103
                                                                                                        }
                                                                                                    }
                                                                                                    if (!((908 >= 18446744073709551615) || true)) {
                                                                                                        var v160 int = -435
                                                                                                        _ = v160
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            fmt.Println(v107, v41, v107)
                                                                                                        }
                                                                                                        if (1 != -2147483648) {
                                                                                                            v41 = "hello"
                                                                                                            v161 := GenS7{F0: 255, F1: 255}
                                                                                                            fmt.Println(v107)
                                                                                                            _ = v161
                                                                                                        }
                                                                                                    }
                                                                                                } else {
                                                                                                    fmt.Println(v41, fi37)
                                                                                                    v162 := GenS6{F0: 4294967295, F1: 0, F2: 844}
                                                                                                    _ = v162
                                                                                                }
                                                                                            default:
                                                                                                var v163 uint8 = 128
                                                                                                _ = v163
                                                                                            }
                                                                                        case "foo bar":
                                                                                            if (!(((uint16(v107) <= 255) || true) && false)) {
                                                                                                v164 := GenS22{F0: int64(fi103), F1: false}
                                                                                                _ = v164
                                                                                            } else {
                                                                                                v107 -= (v107 << ((4294967295) & 63))
                                                                                                var v165 uint16 = 470
                                                                                                var v166 uint = 18446744073709551615
                                                                                                _ = v165
                                                                                                _ = v166
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        var v167 string = (v41 + (("" + v41) + (v41 + "\n\t")))
                                                                                        var v168 int64 = -9223372036854775808
                                                                                        _ = v167
                                                                                        _ = v168
                                                                                    }
                                                                                }
                                                                            case 9223372036854775807:
                                                                                if true {
                                                                                    _ = len(v38)
                                                                                    switch v41 {
                                                                                    case "hello":
                                                                                        if len(v38) >= 2 {
                                                                                            v38 = v38[1:len(v38)]
                                                                                        }
                                                                                    case "":
                                                                                        if len(v38) > 0 {
                                                                                            v38[0] = fi103
                                                                                        }
                                                                                    }
                                                                                    v107 += 980
                                                                                    var v169 float64 = 0.1
                                                                                    v107 = (v107 | v107)
                                                                                    v107 = 2147483647
                                                                                    _ = v169
                                                                                } else {
                                                                                    if (!false) {
                                                                                        var v170 int64 = 0
                                                                                        var v171 string = v41
                                                                                        var v172 float32 = (0.1 * 1.5)
                                                                                        _ = v170
                                                                                        _ = v171
                                                                                        _ = v172
                                                                                    }
                                                                                    var v173 int = -593
                                                                                    _ = v173
                                                                                }
                                                                            default:
                                                                                var v174 uint = 1
                                                                                _ = v174
                                                                            }
                                                                            _ = len(v38)
                                                                        }
                                                                        v41 = (v41 + (("foo bar" + ((v41 + v41) + v41)) + v41))
                                                                    }
                                                                }
                                                            case 42:
                                                                v41 = v41
                                                            default:
                                                                var v175 bool = false
                                                                _ = len(v38)
                                                                fmt.Println(fi103, v41)
                                                                fmt.Println(fi37, v41, fi103)
                                                                v41 = v41
                                                                _ = v175
                                                            }
                                                            v41 = (v41 + (v41 + ("hello" + v41)))
                                                            _ = v107
                                                        }
                                                        _ = fi103
                                                    }
                                                case -506:
                                                    for fi176 := 0; fi176 < 20; fi176++ {
                                                        fmt.Println(v41, fi37, fi176)
                                                        _ = fi176
                                                    }
                                                case 127:
                                                    _ = len(v38)
                                                }
                                                var v177 int64 = 9223372036854775807
                                                _ = v177
                                            }
                                            var v178 int16 = 1000
                                            for fi179 := 0; fi179 < 20; fi179++ {
                                                if false {
                                                    var v180 bool = (!(!(!(false || (9223372036854775807 >= 0)))))
                                                    _ = v180
                                                } else {
                                                    var v181 int16 = v178
                                                    _ = v181
                                                }
                                                _ = fi179
                                            }
                                            _ = v178
                                        } else {
                                            v182 := GenS9{F0: 220}
                                            _ = v182
                                        }
                                    case 4294967295:
                                        v38 = append(v38, (fi37 / ((-12) | 1)))
                                    }
                                    _ = v42
                                } else {
                                    var v183 int16 = -828
                                    var v184 int = fi37
                                    var v185 bool = (-1 <= 127)
                                    _ = v183
                                    _ = v184
                                    _ = v185
                                }
                            case "\n\t":
                                if true {
                                    for fi186 := 0; fi186 < 16; fi186++ {
                                        if (("foo bar" > "hello") || false) {
                                            v38 = append(v38, (-(fi186)))
                                            if (fi186 > 127) {
                                                if (32767 > 42) {
                                                    v41 = v41
                                                    var v187 uint64 = 18446744073709551615
                                                    v41 = v41
                                                    _ = v187
                                                }
                                                v188 := GenS23{F0: 1, F1: 32, F2: (fi186 | (^(fi186)))}
                                                _ = v188
                                            } else {
                                                _ = len(v38)
                                                if true {
                                                    if false {
                                                        v41 = ("hello" + v41)
                                                        switch fi186 {
                                                        case 32768:
                                                            if (127 < 1) {
                                                                var v189 float32 = (-2.25 * (1e10 - (0.0 - 1.5)))
                                                                _ = v189
                                                            }
                                                        case 1:
                                                            if false {
                                                                v41 = ((((((("\n\t" + v41) + ((v41 + (v41 + "\n\t")) + v41)) + ("hello" + v41)) + "a") + "foo bar") + v41) + ("a" + "foo bar"))
                                                                var v190 int8 = 126
                                                                _ = v190
                                                            } else {
                                                                if (0.0 == 3.14159) {
                                                                    var v191 int64 = 792
                                                                    switch (v41 + "hello") {
                                                                    case "\n\t":
                                                                        var v192 float64 = 1.5
                                                                        v192 += (v192 * v192)
                                                                        _ = v192
                                                                    case "a":
                                                                        var v193 uint32 = 2
                                                                        _ = v193
                                                                    case "hello":
                                                                        var v194 int16 = -1
                                                                        _ = v194
                                                                    }
                                                                    v195 := GenS5{F0: 42, F1: 361}
                                                                    switch fi186 {
                                                                    case 4294967295:
                                                                        v196 := GenS1{F0: -2147483648}
                                                                        _ = v196
                                                                    case 160:
                                                                        v41 = v41
                                                                    }
                                                                    if false {
                                                                        if (true && ((("hello" + v41) + (("" + v41) + (("a" + v41) + ("foo bar" + "")))) >= ("" + "hello"))) {
                                                                            v41 = "hello"
                                                                        } else {
                                                                            v197 := GenS24{F0: 100, F1: 10, F2: 446}
                                                                            var v198 uint16 = 0
                                                                            _ = v197
                                                                            _ = v198
                                                                        }
                                                                    }
                                                                    _ = v191
                                                                    _ = v195
                                                                }
                                                                v41 = (v41 + v41)
                                                            }
                                                            v199 := GenS25{F0: genFn0(0)}
                                                            v41 = "\n\t"
                                                            _ = v199
                                                        case -2147483648:
                                                            if len(v38) > 0 {
                                                                v38[0] = (fi186 / ((fi186) | 1))
                                                            }
                                                        }
                                                        if (!(!(-2.25 == (float64(0.1) / (-2.25 - 1.5))))) {
                                                            v200 := GenS26{F0: false, F1: 803}
                                                            v41 = ("" + "\n\t")
                                                            _ = v200
                                                        } else {
                                                            v41 = (((v41 + (v41 + v41)) + v41) + "foo bar")
                                                            v201 := GenS13{F0: 229, F1: 0}
                                                            _ = v201
                                                        }
                                                    } else {
                                                        if (("" <= (("a" + "") + ((("foo bar" + "\n\t") + "foo bar") + ("foo bar" + "\n\t")))) || (!(443 == 255))) {
                                                            var v202 float32 = float32(0.1)
                                                            var v203 int = fi37
                                                            _ = v202
                                                            _ = v203
                                                        } else {
                                                            var v204 float64 = float64(3.14159)
                                                            if false {
                                                                var v205 uint32 = 1
                                                                switch "hello" {
                                                                case "":
                                                                    v205 += (v205 * 2147483647)
                                                                case "foo bar":
                                                                    if false {
                                                                        if false {
                                                                            if false {
                                                                                if (((((v204 / v204) * float64((1e10 / 1e10))) + v204) + v204) == 0.0) {
                                                                                    v206 := GenS8{F0: -1, F1: (fi186 >> ((255) & 63)), F2: 4294967295}
                                                                                    _ = v206
                                                                                } else {
                                                                                    var v207 int8 = -1
                                                                                    v41 = v41
                                                                                    switch v41 {
                                                                                    case "\n\t":
                                                                                        if len(v38) > 0 {
                                                                                            v38[0] = fi37
                                                                                        }
                                                                                        var v208 uint = 65535
                                                                                        v209 := GenS16{F0: (0 == 224), F1: (v208 * uint(v205)), F2: 0}
                                                                                        _ = v208
                                                                                        _ = v209
                                                                                    }
                                                                                    _ = v207
                                                                                }
                                                                                v41 = v41
                                                                            } else {
                                                                                v204 = (v204 - (0.0 + v204))
                                                                                v41 = v41
                                                                                if len(v38) >= 2 {
                                                                                    v38 = v38[1:len(v38)]
                                                                                }
                                                                                var v210 int = fi186
                                                                                if (v210 >= 32768) {
                                                                                    var v211 int8 = 42
                                                                                    fmt.Println(v211, v205)
                                                                                    v211 += v211
                                                                                    _ = v211
                                                                                } else {
                                                                                    var v212 float32 = ((3.14159 / 3.14159) / 0.1)
                                                                                    v213 := GenS27{F0: 503}
                                                                                    var v214 string = v41
                                                                                    _ = v212
                                                                                    _ = v213
                                                                                    _ = v214
                                                                                }
                                                                                _ = v210
                                                                            }
                                                                            if true {
                                                                                if true {
                                                                                    _ = len(v38)
                                                                                    fmt.Println(fi186, v41, v204)
                                                                                    _ = len(v38)
                                                                                } else {
                                                                                    v205 += v205
                                                                                    v215 := GenS14{F0: 856}
                                                                                    _ = v215
                                                                                }
                                                                                var v216 uint64 = 18446744073709551615
                                                                                v217 := GenS11{F0: ((v41 + "\n\t") + v41)}
                                                                                if (42 <= -2147483648) {
                                                                                    var v218 int32 = -351
                                                                                    v216 = (v216 & 2)
                                                                                    _ = v218
                                                                                } else {
                                                                                    v219 := GenS28{F0: 2, F1: "\n\t"}
                                                                                    _ = len(v38)
                                                                                    switch 4294967295 {
                                                                                    case 32767:
                                                                                        var v220 int = (fi37 << ((767) & 63))
                                                                                        _ = v220
                                                                                    case -471:
                                                                                        var v221 uint64 = (v216 << ((255) & 63))
                                                                                        _ = v221
                                                                                    case 42:
                                                                                        if ((v204 > v204) && (false && (1 > genFn0(0)))) {
                                                                                            var v222 int8 = 100
                                                                                            _ = v222
                                                                                        } else {
                                                                                            fmt.Println(v41, v216, v41)
                                                                                        }
                                                                                    }
                                                                                    v223 := GenS29{F0: 0, F1: 1, F2: 905}
                                                                                    fmt.Println(v216, v216, v205)
                                                                                    _ = v219
                                                                                    _ = v223
                                                                                }
                                                                                _ = v216
                                                                                _ = v217
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v204, v204)
                                                                            if false {
                                                                                var v224 int16 = -1
                                                                                _ = v224
                                                                            } else {
                                                                                v205 = 2147483647
                                                                                var v225 int32 = 42
                                                                                _ = v225
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    if false {
                                                                        var v226 int32 = 1
                                                                        _ = v226
                                                                    } else {
                                                                        if true {
                                                                            var v227 int16 = 627
                                                                            _ = v227
                                                                        }
                                                                        var v228 uint16 = 65535
                                                                        _ = v228
                                                                    }
                                                                }
                                                                switch 0 {
                                                                case 32768:
                                                                    if len(v38) > 0 {
                                                                        v38[0] = (fi37 ^ fi37)
                                                                    }
                                                                }
                                                                _ = v205
                                                            } else {
                                                                fmt.Println(fi186, fi37, fi37)
                                                            }
                                                            v229 := GenS30{F0: fi186, F1: 4294967295}
                                                            _ = v204
                                                            _ = v229
                                                        }
                                                        if ((fi186 | 80) >= 42) {
                                                            v41 = (v41 + "")
                                                            v41 = ""
                                                            var v230 int8 = -46
                                                            _ = v230
                                                        }
                                                    }
                                                    v231 := GenS6{F0: 1, F1: 1, F2: 4294967295}
                                                    _ = v231
                                                } else {
                                                    var v232 string = v41
                                                    v41 = v232
                                                    _ = v232
                                                }
                                                var v233 int16 = 156
                                                _ = v233
                                            }
                                            if len(v38) > 0 {
                                                v38[0] = fi37
                                            }
                                        }
                                        v234 := GenS3{F0: false}
                                        if (!(((true && true) || (0 != 42)) && true)) {
                                            var v235 bool = true
                                            var v236 int8 = -1
                                            switch (v41 + (v41 + "\n\t")) {
                                            case "":
                                                switch -9223372036854775808 {
                                                case 42:
                                                    var v237 int8 = v236
                                                    _ = v237
                                                case -2147483648:
                                                    var v238 uint8 = 127
                                                    _ = v238
                                                default:
                                                    v239 := GenS27{F0: 1}
                                                    _ = v239
                                                }
                                            }
                                            _ = v235
                                            _ = v236
                                        } else {
                                            switch (fi37 % ((42) | 1)) {
                                            case 127:
                                                v240 := GenS31{F0: 0, F1: fi186, F2: v41}
                                                _ = v240
                                            case 2147483647:
                                                var v241 bool = (true || true)
                                                _ = v241
                                            case 669:
                                                if (496 <= 1) {
                                                    fmt.Println(fi186)
                                                } else {
                                                    v242 := GenS25{F0: 1000}
                                                    v41 = ""
                                                    fmt.Println(fi186, fi37, fi186)
                                                    v243 := GenS29{F0: 502, F1: -1, F2: -811}
                                                    _ = v242
                                                    _ = v243
                                                }
                                            default:
                                                v41 = v41
                                                var v244 int16 = 1
                                                _ = v244
                                            }
                                        }
                                        if true {
                                            if (42 != (^(fi37))) {
                                                if (188 != 936) {
                                                    if ((0.1 - 0.1) < 1.5) {
                                                        switch fi37 {
                                                        case 4294967295:
                                                            v245 := GenS32{F0: v41, F1: v41, F2: (fi186 & fi186)}
                                                            _ = v245
                                                        }
                                                    } else {
                                                        if len(v38) >= 2 {
                                                            v38 = v38[1:len(v38)]
                                                        }
                                                        if (34 < -128) {
                                                            _ = len(v38)
                                                            if true {
                                                                v41 = "foo bar"
                                                                v246 := GenS33{F0: -1, F1: 42}
                                                                fmt.Println(fi37)
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    fmt.Println(v41, fi186)
                                                                }
                                                                _ = v246
                                                            } else {
                                                                switch (fi37 | 128) {
                                                                case 307:
                                                                    fmt.Println(v41, fi37)
                                                                }
                                                                fmt.Println(v41, fi37, fi186)
                                                            }
                                                        }
                                                        if (false || false) {
                                                            v247 := GenS31{F0: -1, F1: fi186, F2: v41}
                                                            _ = v247
                                                        } else {
                                                            v41 = ""
                                                            _ = len(v38)
                                                        }
                                                    }
                                                    fmt.Println(fi37)
                                                    var v248 float64 = 0.0
                                                    v41 = v41
                                                    v249 := GenS11{F0: ""}
                                                    switch "foo bar" {
                                                    case "":
                                                        if (!(!true)) {
                                                            var v250 string = (v41 + (v41 + v41))
                                                            _ = v250
                                                        } else {
                                                            if (false && (246 != 255)) {
                                                                var v251 int8 = 55
                                                                if len(v38) >= 2 {
                                                                    v38 = v38[1:len(v38)]
                                                                }
                                                                _ = v251
                                                            }
                                                            var v252 int8 = 42
                                                            _ = v252
                                                        }
                                                    default:
                                                        var v253 uint16 = 610
                                                        _ = v253
                                                    }
                                                    var v254 int8 = 84
                                                    fmt.Println(v41, v254)
                                                    v248 = v248
                                                    _ = v248
                                                    _ = v249
                                                    _ = v254
                                                } else {
                                                    var v255 int32 = 0
                                                    _ = v255
                                                }
                                            }
                                        } else {
                                            var v256 int64 = 0
                                            v256 = (v256 ^ (-(v256)))
                                            fmt.Println(fi186, v256)
                                            _ = v256
                                        }
                                        var v257 int64 = 1
                                        fmt.Println(fi37, v41)
                                        v257 = 1
                                        _ = fi186
                                        _ = v234
                                        _ = v257
                                    }
                                } else {
                                    fmt.Println(v41, v41)
                                    _ = len(v38)
                                }
                            default:
                                v258 := GenS27{F0: 18446744073709551615}
                                _ = v258
                            }
                            _ = v41
                        } else {
                            if len(v38) > 0 {
                                v38[0] = -291
                            }
                            var v259 int8 = -33
                            _ = v259
                        }
                        if len(v38) >= 2 {
                            v38 = v38[1:len(v38)]
                        }
                    }
                    var v260 float64 = (-2.25 * (-2.25 * 3.14159))
                    _ = fi37
                    _ = v260
                    _ = v38
                }
                var v261 uint32 = 65535
                v261 += 163
                _ = v261
            }
        }
    }
    var v262 uint32 = 551
    var v263 int16 = -1
    _ = v262
    _ = v263
    return v262
}

func main() {
    v1 := []string{"hello"}
    v2 := []uint32{1, 2147483648, 975}
    var v3 int64 = 1
    var v4 uint8 = 128
    var v5 int64 = (v3 >> ((101) & 63))
    var v264 int = int(v5)
    v264 += v264
    switch "\n\t" {
    case "hello":
        if (((true || (-2.25 < (3.14159 * float32((float64(((((-2.25 / 0.1) + (0.0 / -2.25)) / 3.14159) + -2.25)) * (0.0 + 1e10)))))) || true) || true) {
            switch (v264 - (-(v264))) {
            case 32767:
                var v265 float64 = -2.25
                _ = v265
            case 610:
                if (284 != 973) {
                    if false {
                        if (255 == v4) {
                            for fi266 := 0; fi266 < 3; fi266++ {
                                v267 := GenS34{F0: v5, F1: 65535}
                                if true {
                                    var v268 int8 = 0
                                    _ = v268
                                } else {
                                    _ = len(v2)
                                    var v269 uint16 = 2
                                    _ = v269
                                }
                                if ((fi266 | v264) != fi266) {
                                    v270 := GenS35{F0: ((v5 * -1) == 9223372036854775807)}
                                    v4 -= (^(v4))
                                    var v271 string = "hello"
                                    _ = v270
                                    _ = v271
                                }
                                _ = fi266
                                _ = v267
                            }
                            var v272 int32 = -1
                            _ = v272
                        }
                        fmt.Println(v3, v4)
                    }
                    for fi273 := 0; fi273 < 9; fi273++ {
                        for fi274 := 0; fi274 < 4; fi274++ {
                            var v275 int32 = 2147483647
                            var v276 float32 = 0.0
                            v277 := GenS36{F0: false, F1: "a"}
                            _ = fi274
                            _ = v275
                            _ = v276
                            _ = v277
                        }
                        _ = fi273
                    }
                }
            default:
                var v278 uint32 = 710
                _ = v278
            }
            var v279 uint16 = 255
            if false {
                switch 1 {
                case 0:
                    v279 += 1
                }
                if len(v2) > 0 {
                    v2[0] = 2147483647
                }
                v280 := GenS37{F0: 0, F1: 318}
                _ = v280
            }
            _ = v279
        }
    default:
        var v281 bool = false
        _ = v281
    }
    if true {
        v3 = (-(v5))
        var v282 string = "\n\t"
        _ = v282
    }
    var v283 uint32 = genFn1()
    if (-566 <= genFn0((v4 % (((^(v4))) | 1)))) {
        v284 := GenS38{F0: 138, F1: 2}
        _ = v284
    } else {
        if (0.0 == ((-2.25 - (-2.25 * 0.1)) * 1e10)) {
            v3 -= (v3 % ((1) | 1))
            for fi285 := 0; fi285 < 1; fi285++ {
                v286 := GenS26{F0: (100 <= 42), F1: 900}
                var v287 uint16 = 2
                _ = fi285
                _ = v286
                _ = v287
            }
            v5 = 1
        } else {
            v264 += (-(v264))
            v288 := GenS39{F0: 42, F1: v3, F2: 42}
            _ = v288
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v289 := GenS5{F0: 1000, F1: 4294967295}
        _ = v289
    }
    for fi290 := 0; fi290 < 14; fi290++ {
        for fi291 := 0; fi291 < 18; fi291++ {
            var v292 int8 = 42
            _ = fi291
            _ = v292
        }
        _ = fi290
    }
    switch ("a" + "foo bar") {
    case "foo bar":
        v293 := GenS14{F0: 32767}
        _ = v293
    case "":
        if len(v2) > 0 {
            v2[0] = 2
        }
    case "hello":
        fmt.Println(v3)
    default:
        var v294 string = "\n\t"
        _ = v294
    }
    switch ("hello" + "foo bar") {
    case "foo bar":
        var v295 uint16 = 32767
        _ = v295
    default:
        if (v264 == v264) {
            var v296 uint64 = 162
            v1 = append(v1, ("a" + "\n\t"))
            _ = v296
        }
    }
    v297 := GenS1{F0: -544}
    v5 = (v5 - -1)
    switch v264 {
    case -2147483648:
        v3 += 283
    case -1:
        for fi298 := 0; fi298 < 19; fi298++ {
            fmt.Println(v5)
            _ = fi298
        }
    }
    for fi299 := 0; fi299 < 14; fi299++ {
        switch 128 {
        case 32768:
            v1 = append(v1, "hello")
        default:
            v4 -= (v4 >> ((261) & 63))
        }
        _ = fi299
    }
    v4 = 1
    if true {
        var v300 int32 = -287
        _ = v300
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v264:", v264)
    fmt.Println("v283:", v283)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v264
    _ = v283
    _ = v297
    _ = v3
    _ = v4
    _ = v5
}
