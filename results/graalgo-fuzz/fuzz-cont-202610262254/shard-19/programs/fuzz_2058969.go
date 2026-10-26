package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int; F1 uint16; F2 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 int16; F1 int32; F2 bool }
type GenS7 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS8 struct { F0 uint16; F1 bool; F2 string }
type GenS9 struct { F0 int64; F1 int16; F2 int8 }
type GenS10 struct { F0 int; F1 string; F2 int32 }
type GenS11 struct { F0 int; F1 uint16; F2 uint32 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int32; F1 uint8; F2 bool }
type GenS14 struct { F0 uint64; F1 int32; F2 string }
func genFn0(p0 bool, p1 bool, p2 int) uint8 {
    for fi4 := 0; fi4 < 8; fi4++ {
        for fi5 := 0; fi5 < 10; fi5++ {
            v6 := []int{(p2 >> ((4294967295) & 63))}
            switch "" {
            case "foo bar":
                p1 = p1
            case "a":
                if p0 {
                    if true {
                        p1 = p1
                        p0 = false
                        _ = len(v6)
                    }
                } else {
                    if (-847 == 155) {
                        var v7 uint16 = 42
                        p1 = (107 >= 4294967295)
                        switch (("foo bar" + "") + "\n\t") {
                        case "\n\t":
                            if ((!p0) || (277 >= -117)) {
                                switch (fi5 / (((p2 | p2)) | 1)) {
                                case 9223372036854775807:
                                    if true {
                                        v8 := GenS1{F0: "", F1: 2}
                                        var v9 float64 = 3.14159
                                        v10 := GenS2{F0: 32767, F1: 438}
                                        _ = v10
                                        _ = v8
                                        _ = v9
                                    } else {
                                        p1 = p0
                                        if ((^(v7)) > v7) {
                                            var v11 uint = 4294967295
                                            p1 = false
                                            _ = v11
                                        }
                                        fmt.Println(p1, v7)
                                    }
                                case 812:
                                    v12 := GenS1{F0: "a", F1: 255}
                                    _ = v12
                                case 811:
                                    var v13 uint64 = 283
                                    _ = v13
                                default:
                                    if false {
                                        switch 1 {
                                        case 1:
                                            var v14 float64 = -2.25
                                            _ = v14
                                        case -1:
                                            var v15 float64 = -2.25
                                            _ = v15
                                        case 42:
                                            switch ((("hello" + "foo bar") + "foo bar") + "hello") {
                                            case "foo bar":
                                                if p0 {
                                                    var v16 uint8 = 3
                                                    v6 = append(v6, (fi4 % (((fi5 & (^(p2)))) | 1)))
                                                    p2 -= p2
                                                    _ = v16
                                                } else {
                                                    fmt.Println(v7, fi5, p1)
                                                }
                                            case "":
                                                var v17 string = "foo bar"
                                                _ = v17
                                            case "hello":
                                                p2 -= -1
                                            default:
                                                switch "\n\t" {
                                                case "foo bar":
                                                    fmt.Println(fi4)
                                                case "a":
                                                    fmt.Println(p2)
                                                case "":
                                                    _ = len(v6)
                                                }
                                            }
                                        }
                                        switch "foo bar" {
                                        case "\n\t":
                                            _ = len(v6)
                                        default:
                                            var v18 uint16 = 42
                                            _ = v18
                                        }
                                        v19 := GenS1{F0: "", F1: 255}
                                        _ = v19
                                    }
                                }
                                p1 = p1
                                var v20 uint16 = 255
                                _ = v20
                            }
                        case "a":
                            var v21 uint32 = 0
                            switch p2 {
                            case 32768:
                                v6 = append(v6, -2147483648)
                            default:
                                if p0 {
                                    p0 = (255 <= 127)
                                    p0 = false
                                    v21 = 2
                                } else {
                                    v6 = append(v6, (p2 | (fi5 | -1)))
                                    var v22 uint64 = 1
                                    var v23 uint64 = v22
                                    _ = v22
                                    _ = v23
                                }
                            }
                            _ = v21
                        case "":
                            if ((42 < 32767) && (32767 != 32767)) {
                                var v24 bool = (2147483647 == 42)
                                var v25 bool = (v24 || (v24 && (!false)))
                                if v24 {
                                    p0 = ((1e10 * ((1e10 - 1e10) * 1.5)) < 0.1)
                                }
                                p2 = 32768
                                _ = v24
                                _ = v25
                            }
                        default:
                            p0 = p0
                        }
                        switch ("hello" + ("hello" + "")) {
                        case "hello":
                            var v26 int8 = 100
                            _ = v26
                        case "foo bar":
                            if p1 {
                                if len(v6) > 0 {
                                    v6[0] = p2
                                }
                                switch "hello" {
                                case "":
                                    _ = len(v6)
                                case "foo bar":
                                    v7 = 42
                                case "a":
                                    var v27 uint32 = 1
                                    _ = v27
                                }
                                v28 := GenS2{F0: v7, F1: v7}
                                _ = v28
                            }
                        default:
                            switch "" {
                            case "hello":
                                v29 := GenS1{F0: "\n\t", F1: 22}
                                _ = v29
                            case "foo bar":
                                p0 = (p0 && p0)
                            case "":
                                p1 = (31 != 2)
                            }
                        }
                        p0 = false
                        _ = v7
                    } else {
                        p1 = p1
                        var v30 uint8 = 128
                        switch "foo bar" {
                        case "\n\t":
                            p2 = 128
                            if true {
                                var v31 uint = 255
                                _ = v31
                            } else {
                                var v32 uint16 = 65535
                                v33 := GenS1{F0: "", F1: 0}
                                _ = v32
                                _ = v33
                            }
                            var v34 int = fi4
                            _ = v34
                        }
                        p0 = (65535 == 0)
                        fmt.Println(v30, fi4)
                        v30 = 31
                        _ = len(v6)
                        _ = v30
                    }
                    var v35 uint64 = 42
                    _ = v35
                }
            default:
                fmt.Println(p2, p2)
            }
            fmt.Println(fi4, fi4, fi4)
            _ = fi5
            _ = v6
        }
        v36 := GenS1{F0: "\n\t", F1: 0}
        _ = fi4
        _ = v36
    }
    v37 := []int{(-(p2)), -2147483648, (p2 | (p2 << ((279) & 63)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v37
    return 2
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[int]bool)
    v3 := make(map[string]uint32)
    v1 = 69
    v38 := GenS2{F0: (v1 + 0), F1: v1}
    switch ("hello" + ("hello" + "hello")) {
    case "\n\t":
        delete(v2, 2147483647)
    case "":
        v1 = v1
    }
    for fi39 := 0; fi39 < 6; fi39++ {
        var v40 string = "foo bar"
        var v41 int64 = -9223372036854775808
        fmt.Println(fi39)
        _ = fi39
        _ = v40
        _ = v41
    }
    switch 2147483647 {
    case 32767:
        for fi42 := 0; fi42 < 7; fi42++ {
            for fi43 := 0; fi43 < 10; fi43++ {
                v1 += v1
                fmt.Println(fi42, fi43)
                _ = fi43
            }
            for fi44 := 0; fi44 < 7; fi44++ {
                var v45 bool = true
                v46 := GenS1{F0: ("\n\t" + ("" + (("\n\t" + "foo bar") + "foo bar"))), F1: 65535}
                _ = fi44
                _ = v45
                _ = v46
            }
            var v47 int64 = -1
            _ = fi42
            _ = v47
        }
    default:
        if false {
            switch "a" {
            case "hello":
                if (true && (float32(3.14159) <= 1.5)) {
                    var v48 uint8 = 127
                    if true {
                        var v49 uint16 = v1
                        if true {
                            v50 := GenS3{F0: 2147483647, F1: (v1 | v49), F2: 1}
                            _ = v50
                        } else {
                            if false {
                                v51 := GenS4{F0: -284}
                                for fi52 := 0; fi52 < 4; fi52++ {
                                    delete(v2, 42)
                                    _ = fi52
                                }
                                v53 := GenS5{F0: genFn0(((false || false) || (!(false || true))), (!(518 != 0)), 945), F1: v1}
                                _ = v51
                                _ = v53
                            } else {
                                if false {
                                    for fi54 := 0; fi54 < 19; fi54++ {
                                        if (!false) {
                                            v1 = v1
                                            if true {
                                                v55 := GenS6{F0: 0, F1: int32(fi54), F2: (("foo bar" + "foo bar") < "")}
                                                _ = v55
                                            } else {
                                                v49 = v1
                                                fmt.Println(v49)
                                                switch "foo bar" {
                                                case "hello":
                                                    var v56 int64 = -689
                                                    _ = v56
                                                case "\n\t":
                                                    if false {
                                                        if false {
                                                            v49 -= v49
                                                            v48 = (v48 >> ((42) & 63))
                                                            v57 := []int{(fi54 % ((fi54) | 1)), (fi54 * -670), -2147483648, (fi54 << ((416) & 63))}
                                                            _ = v57
                                                        } else {
                                                            var v58 int = 32768
                                                            v1 += (v49 >> ((0) & 63))
                                                            var v59 int16 = 0
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                        var v60 int64 = 165
                                                        _ = v60
                                                    } else {
                                                        v1 -= (v49 >> ((0) & 63))
                                                        switch ("hello" + "a") {
                                                        case "hello":
                                                            for fi61 := 0; fi61 < 8; fi61++ {
                                                                v49 -= (^(v49))
                                                                v49 = (v49 ^ (v1 - v49))
                                                                if (((-2.25 - 1.5) * ((float64(1.5) * 0.1) * 1.5)) <= 1.5) {
                                                                    v62 := []int{32767, 127, -402}
                                                                    if len(v62) > 0 {
                                                                        v62[0] = (-(fi61))
                                                                    }
                                                                    _ = v62
                                                                }
                                                                if (0 <= -9223372036854775808) {
                                                                    switch ("a" + "\n\t") {
                                                                    case "hello":
                                                                        v48 = (v48 % ((uint8(v49)) | 1))
                                                                    case "\n\t":
                                                                        if (1 > (v48 % (((^(v48))) | 1))) {
                                                                            if false {
                                                                                v63 := GenS6{F0: 0, F1: -177, F2: (false && true)}
                                                                                delete(v3, "\n\t")
                                                                                _ = v63
                                                                            } else {
                                                                                fmt.Println(fi61, v1, v48)
                                                                                switch "foo bar" {
                                                                                case "foo bar":
                                                                                    v64 := []int{fi61, fi54, -1}
                                                                                    _ = v64
                                                                                case "":
                                                                                    v1 += (v1 / ((0) | 1))
                                                                                }
                                                                                fmt.Println(fi61, v1)
                                                                                var v65 string = "hello"
                                                                                v66 := v3[v65]
                                                                                _ = v65
                                                                                _ = v66
                                                                            }
                                                                            if false {
                                                                                v48 = v48
                                                                                v49 = 0
                                                                            }
                                                                            var v67 int8 = 100
                                                                            _ = v67
                                                                        }
                                                                    default:
                                                                        var v68 string = "\n\t"
                                                                        _ = v68
                                                                    }
                                                                    var v69 float64 = 3.14159
                                                                    fmt.Println(v48, v48, v1)
                                                                    _ = v69
                                                                } else {
                                                                    var v70 int32 = 0
                                                                    var v71 uint32 = 493
                                                                    _ = v70
                                                                    _ = v71
                                                                }
                                                                _ = fi61
                                                            }
                                                        }
                                                        v48 = 182
                                                    }
                                                }
                                            }
                                        }
                                        _ = fi54
                                    }
                                }
                                _ = len(v2)
                                v48 -= v48
                            }
                            var v72 int16 = -512
                            var v73 uint = 255
                            _ = v72
                            _ = v73
                        }
                        _ = v49
                    } else {
                        v48 += v48
                    }
                    for fi74 := 0; fi74 < 11; fi74++ {
                        v75 := v3[("\n\t" + "\n\t")]
                        _ = fi74
                        _ = v75
                    }
                    _ = v48
                } else {
                    switch "a" {
                    case "":
                        fmt.Println(v1)
                    case "hello":
                        switch "" {
                        case "a":
                            if (int16(v1) >= -675) {
                                switch -9223372036854775808 {
                                case 32767:
                                    if false {
                                        switch "foo bar" {
                                        case "a":
                                            for fi76 := 0; fi76 < 9; fi76++ {
                                                v2[-2147483648] = true
                                                _ = fi76
                                            }
                                        case "foo bar":
                                            var v77 uint8 = 28
                                            _ = v77
                                        case "":
                                            for fi78 := 0; fi78 < 2; fi78++ {
                                                if ((32767 < 0) || false) {
                                                    if true {
                                                        fmt.Println(v1, v1)
                                                        v1 = (v1 | v1)
                                                    } else {
                                                        v1 += 835
                                                        var v79 int16 = -1
                                                        switch (fi78 ^ 2147483647) {
                                                        case 32767:
                                                            v1 += (^(v1))
                                                        default:
                                                            v1 += 32767
                                                        }
                                                        _ = v79
                                                    }
                                                } else {
                                                    v80 := GenS2{F0: v1, F1: v1}
                                                    var v81 uint64 = 1
                                                    _ = v80
                                                    _ = v81
                                                }
                                                var v82 int8 = 42
                                                _ = fi78
                                                _ = v82
                                            }
                                        }
                                    }
                                case -586:
                                    for fi83 := 0; fi83 < 13; fi83++ {
                                        switch -9223372036854775808 {
                                        case 167:
                                            var v84 int8 = 0
                                            _ = v84
                                        case 128:
                                            switch "foo bar" {
                                            case "a":
                                                v85 := v2[(fi83 % (((fi83 >> ((42) & 63))) | 1))]
                                                _ = v85
                                            case "hello":
                                                var v86 uint16 = (v1 << ((65535) & 63))
                                                _ = v86
                                            case "\n\t":
                                                var v87 float32 = (0.1 * 1e10)
                                                _ = v87
                                            default:
                                                for fi88 := 0; fi88 < 14; fi88++ {
                                                    switch "\n\t" {
                                                    case "a":
                                                        v89 := GenS7{F0: 124, F1: 791, F2: (v1 & v1)}
                                                        _ = v89
                                                    case "":
                                                        v90 := GenS8{F0: (v1 ^ (v1 * (v1 + (v1 + 65535)))), F1: (!(false || (0 != -166))), F2: "a"}
                                                        _ = v90
                                                    default:
                                                        if (4294967295 < 1) {
                                                            fmt.Println(v1, fi83)
                                                        }
                                                    }
                                                    v91 := GenS8{F0: (v1 | 32767), F1: (357 <= 65535), F2: ("a" + "")}
                                                    _ = fi88
                                                    _ = v91
                                                }
                                            }
                                        case 127:
                                            var v92 float32 = (0.1 * (-2.25 * -2.25))
                                            _ = v92
                                        }
                                        var v93 int8 = -31
                                        _ = fi83
                                        _ = v93
                                    }
                                default:
                                    var v94 uint64 = 42
                                    _ = v94
                                }
                                if false {
                                    for fi95 := 0; fi95 < 8; fi95++ {
                                        v1 += 0
                                        var v96 uint8 = 42
                                        v97 := GenS6{F0: -304, F1: 0, F2: true}
                                        _ = fi95
                                        _ = v96
                                        _ = v97
                                    }
                                    v1 += v1
                                    switch (("a" + ("\n\t" + "hello")) + "a") {
                                    case "hello":
                                        for fi98 := 0; fi98 < 13; fi98++ {
                                            var v99 string = "foo bar"
                                            var v100 float64 = (float64(((1e10 + ((3.14159 * 1.5) + (1.5 * 0.1))) / 0.1)) + float64(-2.25))
                                            _ = fi98
                                            _ = v100
                                            _ = v99
                                        }
                                    case "foo bar":
                                        v101 := GenS8{F0: v1, F1: true, F2: "foo bar"}
                                        _ = v101
                                    case "":
                                        var v102 bool = false
                                        _ = v102
                                    }
                                }
                            } else {
                                for fi103 := 0; fi103 < 5; fi103++ {
                                    if false {
                                        v2[32768] = false
                                        for fi104 := 0; fi104 < 16; fi104++ {
                                            if true {
                                                v1 -= v1
                                                if true {
                                                    switch "hello" {
                                                    case "\n\t":
                                                        v1 = (v1 % ((42) | 1))
                                                        if (67 < 100) {
                                                            v1 -= 65535
                                                            v1 -= v1
                                                        } else {
                                                            if ((-2.25 < 0.0) && (false || (42 >= 0))) {
                                                                var v105 int16 = -509
                                                                v106 := []int{(fi104 - fi103), -935}
                                                                _ = v105
                                                                _ = v106
                                                            } else {
                                                                fmt.Println(fi103)
                                                                v1 = (^(v1))
                                                                if (988 < 1) {
                                                                    var v107 uint8 = 145
                                                                    if (!(1 < genFn0((1.5 == 1.5), ("a" > "a"), (fi103 & fi104)))) {
                                                                        fmt.Println(fi104, v1, v107)
                                                                        var v108 uint64 = 262
                                                                        v108 = (v108 % (((v108 / ((18446744073709551615) | 1))) | 1))
                                                                        _ = v108
                                                                    } else {
                                                                        v109 := GenS6{F0: 1, F1: 2147483647, F2: true}
                                                                        _ = v109
                                                                    }
                                                                    var v110 uint64 = uint64(v1)
                                                                    var v111 uint = 491
                                                                    _ = v107
                                                                    _ = v110
                                                                    _ = v111
                                                                } else {
                                                                    v112 := []int{(fi104 / (((fi103 ^ (fi103 % (((-(fi104))) | 1)))) | 1)), (fi103 / ((fi103) | 1))}
                                                                    v113 := GenS1{F0: ("\n\t" + ("a" + "hello")), F1: 18446744073709551615}
                                                                    _ = v112
                                                                    _ = v113
                                                                }
                                                            }
                                                        }
                                                        switch "\n\t" {
                                                        case "":
                                                            v1 = v1
                                                        }
                                                        v1 += 0
                                                        v114 := GenS5{F0: 0, F1: 65535}
                                                        v115 := GenS6{F0: 32767, F1: 1, F2: false}
                                                        _ = v114
                                                        _ = v115
                                                    case "hello":
                                                        delete(v2, fi104)
                                                    case "foo bar":
                                                        v1 = 255
                                                    }
                                                    var v116 int8 = 0
                                                    if false {
                                                        fmt.Println(fi104, v1)
                                                        fmt.Println(v1, v1, fi103)
                                                        if true {
                                                            if (!((0 == 1) && (1.5 >= 1.5))) {
                                                                fmt.Println(fi104, v116, v1)
                                                            } else {
                                                                v117 := GenS4{F0: fi104}
                                                                _ = v117
                                                            }
                                                        }
                                                    } else {
                                                        var v118 uint32 = 1
                                                        _ = v118
                                                    }
                                                    _ = v116
                                                } else {
                                                    v1 += v1
                                                }
                                                var v119 int64 = 0
                                                _ = v119
                                            }
                                            switch "foo bar" {
                                            case "foo bar":
                                                if true {
                                                    switch fi104 {
                                                    case 2147483647:
                                                        var v120 float32 = 3.14159
                                                        _ = v120
                                                    default:
                                                        var v121 uint64 = 42
                                                        _ = v121
                                                    }
                                                }
                                            }
                                            v1 = v1
                                            fmt.Println(fi104, v1, fi104)
                                            _ = fi104
                                        }
                                        var v122 bool = (!(true && ((false || (!(-169 >= 32767))) && false)))
                                        _ = v122
                                    }
                                    var v123 bool = false
                                    var v124 uint = 42
                                    _ = fi103
                                    _ = v123
                                    _ = v124
                                }
                                var v125 int32 = 943
                                _ = v125
                            }
                        case "hello":
                            if (42 > 127) {
                                v2[328] = false
                                if (false || (!false)) {
                                    switch -1 {
                                    case 42:
                                        v1 += v1
                                    case 128:
                                        if (((!false) && false) || (2 != 2147483647)) {
                                            for fi126 := 0; fi126 < 13; fi126++ {
                                                for fi127 := 0; fi127 < 5; fi127++ {
                                                    v128 := []int{fi126, (fi127 / (((^(fi127))) | 1)), fi127, -2147483648}
                                                    var v129 uint = 65535
                                                    fmt.Println(v129)
                                                    _ = fi127
                                                    _ = v128
                                                    _ = v129
                                                }
                                                _ = fi126
                                            }
                                            v1 = (v1 % (((v1 + (v1 % ((1) | 1)))) | 1))
                                            fmt.Println(v1)
                                        } else {
                                            v130 := GenS9{F0: -9223372036854775808, F1: 32767, F2: -67}
                                            for fi131 := 0; fi131 < 12; fi131++ {
                                                v132 := GenS10{F0: fi131, F1: ("foo bar" + ""), F2: -423}
                                                fmt.Println(fi131, fi131)
                                                _ = fi131
                                                _ = v132
                                            }
                                            v133 := v2[-753]
                                            _ = v130
                                            _ = v133
                                        }
                                    }
                                }
                                switch (("" + "a") + "") {
                                case "a":
                                    v134 := GenS11{F0: 4294967295, F1: v1, F2: 65535}
                                    _ = v134
                                case "hello":
                                    for fi135 := 0; fi135 < 8; fi135++ {
                                        v1 -= v1
                                        v3[""] = 1
                                        _ = fi135
                                    }
                                case "foo bar":
                                    fmt.Println(v1)
                                default:
                                    v136 := []int{}
                                    _ = v136
                                }
                            } else {
                                var v137 uint16 = v1
                                _ = v137
                            }
                        default:
                            var v138 int16 = 1000
                            _ = v138
                        }
                    default:
                        v1 = (^(v1))
                    }
                    for fi139 := 0; fi139 < 16; fi139++ {
                        var v140 int16 = 0
                        switch ("" + "foo bar") {
                        case "\n\t":
                            var v141 uint32 = 4294967295
                            _ = v141
                        case "":
                            var v142 float32 = 0.0
                            _ = v142
                        }
                        if true {
                            var v143 int64 = 0
                            for fi144 := 0; fi144 < 9; fi144++ {
                                switch 42 {
                                case 404:
                                    var v145 float64 = 1e10
                                    _ = v145
                                default:
                                    if true {
                                        switch "a" {
                                        case "":
                                            var v146 uint32 = 1
                                            _ = v146
                                        }
                                        if true {
                                            v140 += (v140 * v140)
                                            v147 := GenS12{F0: int32(fi144)}
                                            _ = v147
                                        }
                                        _ = len(v3)
                                    }
                                }
                                var v148 uint8 = 128
                                _ = fi144
                                _ = v148
                            }
                            var v149 uint16 = v1
                            _ = v143
                            _ = v149
                        } else {
                            switch (fi139 >> ((4294967295) & 63)) {
                            case 128:
                                v150 := []int{(fi139 - 128), fi139, fi139, (fi139 << ((4294967295) & 63))}
                                _ = v150
                            case -313:
                                var v151 uint8 = 83
                                _ = v151
                            case -502:
                                if true {
                                    for fi152 := 0; fi152 < 18; fi152++ {
                                        fmt.Println(fi152, fi139, fi139)
                                        v1 = v1
                                        var v153 uint32 = 270
                                        _ = fi152
                                        _ = v153
                                    }
                                }
                            }
                        }
                        _ = fi139
                        _ = v140
                    }
                    fmt.Println(v1)
                }
            }
        } else {
            v1 = (^(v1))
            if (54 <= 0) {
                var v154 int64 = 42
                _ = v154
            }
        }
    }
    var v155 uint8 = 127
    if (!((((1 >= (v1 - v1)) && false) && ((1.5 + 1e10) <= 3.14159)) || (false && ((1 != 0) && true)))) {
        fmt.Println(v155, v155)
        v1 -= v1
        fmt.Println(v155)
    } else {
        var v156 int32 = -1
        var v157 uint16 = v1
        v155 += genFn0(true, (3.14159 > 0.0), 1)
        _ = v156
        _ = v157
    }
    v158 := GenS13{F0: 42, F1: (v155 - (v155 | (^(v155)))), F2: true}
    v159 := GenS14{F0: 18446744073709551615, F1: 940, F2: "hello"}
    if false {
        v160 := []int{786, 127, 128, -1}
        switch 9223372036854775807 {
        case 32767:
            v1 = 32767
        case -1:
            v2[1] = true
        default:
            v155 -= v155
        }
        _ = v160
    } else {
        v155 += (v155 | 127)
        for fi161 := 0; fi161 < 13; fi161++ {
            v1 = (v1 >> ((1) & 63))
            switch (fi161 >> ((65535) & 63)) {
            case 4294967295:
                var v162 int = (fi161 - 4294967295)
                _ = v162
            case 128:
                if (!false) {
                    v155 = (^(v155))
                } else {
                    delete(v3, ("" + "hello"))
                }
            }
            v155 = 127
            _ = fi161
        }
    }
    v155 += 127
    v163 := GenS9{F0: 420, F1: -879, F2: 127}
    var v164 uint8 = v155
    fmt.Println("v1:", v1)
    fmt.Println("v155:", v155)
    fmt.Println("v164:", v164)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v155
    _ = v158
    _ = v159
    _ = v163
    _ = v164
    _ = v2
    _ = v3
    _ = v38
}
