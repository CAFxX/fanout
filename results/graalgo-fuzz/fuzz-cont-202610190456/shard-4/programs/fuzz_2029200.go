package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 string; F1 uint16; F2 bool }
type GenS6 struct { F0 int8; F1 int64 }
type GenS7 struct { F0 int64; F1 uint32; F2 int }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int64; F1 uint32 }
type GenS10 struct { F0 uint8; F1 int64 }
type GenS11 struct { F0 int16; F1 int64 }
type GenS12 struct { F0 int32; F1 uint16 }
type GenS13 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 int64; F1 int; F2 string }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 int64; F1 int; F2 int64 }
func genFn0(p0 bool, p1 uint16) uint {
    v7 := GenS1{F0: 42}
    switch "" {
    case "":
        p1 = 42
    case "foo bar":
        switch 128 {
        case -424:
            switch ("a" + ("" + "a")) {
            case "\n\t":
                p0 = true
            }
        default:
            var v8 int = 128
            _ = v8
        }
    case "a":
        v9 := []int{246, 2147483647, 364}
        _ = v9
    }
    var v10 uint16 = (p1 ^ 65535)
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return 0
}
func genFn1(p0 uint8, p1 string) int {
    v11 := []int{4294967295}
    var v12 uint8 = p0
    v12 = 1
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 2147483647
}

func main() {
    var v1 int16 = 1000
    var v2 uint = 42
    v3 := make(map[int]string)
    v4 := make(map[string]int)
    var v5 uint64 = 502
    v6 := []string{"hello", "\n\t", (((("\n\t" + ("" + "hello")) + ("\n\t" + "a")) + ("" + "a")) + "\n\t")}
    var v13 uint64 = v5
    switch 128 {
    case -553:
        v14 := GenS2{F0: true, F1: (v1 / ((v1) | 1))}
        _ = v14
    default:
        if false {
            for fi15 := 0; fi15 < 14; fi15++ {
                fmt.Println(fi15)
                _ = fi15
            }
        } else {
            fmt.Println(v2, v1, v1)
            switch "a" {
            case "foo bar":
                v5 -= (v13 / ((v13) | 1))
            case "\n\t":
                var v16 float32 = (1.5 - (0.0 + (((((-2.25 - 3.14159) / 3.14159) * (((0.1 / float32(-2.25)) * (1e10 - (1e10 / 1.0))) - 0.1)) - 3.14159) + 3.14159)))
                _ = v16
            case "hello":
                var v17 int16 = v1
                _ = v17
            default:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            var v18 uint8 = 127
            _ = v18
        }
    }
    if ((1.5 * float64(1e10)) < float64(0.1)) {
        if (0.1 <= 1e10) {
            if len(v6) > 0 {
                v6[0] = "foo bar"
            }
            v4["foo bar"] = 32767
        }
        v19 := GenS3{F0: true, F1: v13, F2: 2147483648}
        _ = v19
    }
    fmt.Println(v13, v5, v5)
    v2 -= v2
    switch "hello" {
    case "":
        if (false && false) {
            v1 += -1
            v2 += genFn0((!(2 <= 128)), 32767)
            if false {
                fmt.Println(v2, v5)
            }
        }
    }
    v2 = (v2 % (((v2 / ((v2) | 1))) | 1))
    var v20 int = -57
    for fi21 := 0; fi21 < 16; fi21++ {
        v6 = append(v6, "a")
        var v22 uint16 = 856
        v23 := GenS2{F0: false, F1: 0}
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println(v2, v1, v1)
    for fi24 := 0; fi24 < 14; fi24++ {
        var v25 int16 = 1
        if len(v6) > 0 {
            v6[0] = "hello"
        }
        _ = fi24
        _ = v25
    }
    var v26 uint32 = 1
    switch -891 {
    case 127:
        _ = len(v4)
    case -1:
        v1 = (v1 + (v1 % ((v1) | 1)))
    case -9223372036854775808:
        for fi27 := 0; fi27 < 4; fi27++ {
            v28 := GenS1{F0: 63}
            v2 = uint(v26)
            if ((1000 > v1) || ("\n\t" < "foo bar")) {
                if (v20 != -577) {
                    switch (v20 * -423) {
                    case 336:
                        if (32767 != v1) {
                            var v29 uint32 = (v26 ^ (^(v26)))
                            _ = v29
                        }
                    default:
                        for fi30 := 0; fi30 < 13; fi30++ {
                            if true {
                                switch "a" {
                                case "a":
                                    fmt.Println(fi30)
                                default:
                                    switch "hello" {
                                    case "\n\t":
                                        fmt.Println(v1, v5)
                                    case "foo bar":
                                        fmt.Println(v13, fi27, v2)
                                    case "":
                                        fmt.Println(v20, v13, v2)
                                    default:
                                        var v31 int32 = -2147483648
                                        _ = v31
                                    }
                                }
                                switch (-(fi27)) {
                                case 9223372036854775807:
                                    var v32 uint64 = uint64(v2)
                                    _ = v32
                                case 272:
                                    v33 := GenS4{F0: uint(v5), F1: -58}
                                    _ = v33
                                case -9223372036854775808:
                                    fmt.Println(fi30, v26, v26)
                                default:
                                    if true {
                                        v1 = (v1 / ((-32768) | 1))
                                        switch "hello" {
                                        case "foo bar":
                                            v34 := GenS1{F0: -128}
                                            _ = v34
                                        case "":
                                            v35 := GenS5{F0: ("foo bar" + ("a" + "foo bar")), F1: 75, F2: false}
                                            _ = v35
                                        default:
                                            var v36 string = ""
                                            _ = v36
                                        }
                                    }
                                }
                                if true {
                                    var v37 int64 = 9223372036854775807
                                    if true {
                                        switch (-(fi27)) {
                                        case 761:
                                            v37 -= (v37 | v37)
                                        case 9223372036854775807:
                                            var v38 int32 = -120
                                            _ = v38
                                        case -428:
                                            var v39 int16 = v1
                                            _ = v39
                                        default:
                                            var v40 string = "foo bar"
                                            _ = v40
                                        }
                                    } else {
                                        if false {
                                            if (1.5 < 1e10) {
                                                if true {
                                                    fmt.Println(v1)
                                                }
                                            }
                                            var v41 string = ("\n\t" + "hello")
                                            _ = v41
                                        }
                                        var v42 uint64 = v13
                                        _ = v42
                                    }
                                    v13 = 42
                                    var v43 int8 = -1
                                    if false {
                                        if false {
                                            v44 := GenS4{F0: v2, F1: v1}
                                            _ = v44
                                        } else {
                                            var v45 uint32 = v26
                                            _ = v45
                                        }
                                    }
                                    var v46 uint8 = 255
                                    var v47 uint32 = v26
                                    fmt.Println(v5, v43, fi27)
                                    v26 = v26
                                    _ = v37
                                    _ = v43
                                    _ = v46
                                    _ = v47
                                }
                            } else {
                                _ = len(v6)
                            }
                            v48 := GenS5{F0: "hello", F1: 32767, F2: (-912 < 1)}
                            var v49 uint = (v2 % (((v2 << ((v2) & 63))) | 1))
                            if false {
                                var v50 uint8 = 2
                                _ = v50
                            }
                            _ = fi30
                            _ = v48
                            _ = v49
                        }
                    }
                } else {
                    if (18446744073709551615 == 382) {
                        switch "a" {
                        case "":
                            var v51 int32 = 800
                            _ = v51
                        case "\n\t":
                            switch (fi27 >> ((v2) & 63)) {
                            case 2147483647:
                                switch 0 {
                                case 127:
                                    if true {
                                        v52 := GenS6{F0: 100, F1: 658}
                                        for fi53 := 0; fi53 < 4; fi53++ {
                                            var v54 string = "hello"
                                            if true {
                                                fmt.Println(v54, v2, v2)
                                                v55 := v3[fi53]
                                                _ = v55
                                            }
                                            switch (fi27 % (((fi53 & genFn1(196, v54))) | 1)) {
                                            case -2147483648:
                                                _ = len(v3)
                                            case 880:
                                                if ((v26 << ((v2) & 63)) != v26) {
                                                    if true {
                                                        var v56 uint32 = 65535
                                                        if (int64(v26) > 1) {
                                                            if false {
                                                                switch (v20 + (^(fi27))) {
                                                                case 155:
                                                                    var v57 int32 = 2147483647
                                                                    _ = v57
                                                                default:
                                                                    var v58 float64 = (3.14159 - 0.1)
                                                                    _ = v58
                                                                }
                                                                var v59 int = 127
                                                                _ = len(v3)
                                                                _ = v59
                                                            } else {
                                                                v20 = (v20 - genFn1(0, ("\n\t" + "a")))
                                                            }
                                                            switch "foo bar" {
                                                            case "hello":
                                                                fmt.Println(v26, v20, v56)
                                                            default:
                                                                var v60 string = v54
                                                                if false {
                                                                    fmt.Println(v1, v56)
                                                                    fmt.Println(v56, v13, v1)
                                                                }
                                                                var v61 uint16 = 0
                                                                _ = v60
                                                                _ = v61
                                                            }
                                                        } else {
                                                            if ((v56 / ((2147483648) | 1)) > 2147483648) {
                                                                switch v54 {
                                                                case "a":
                                                                    delete(v3, 289)
                                                                default:
                                                                    var v62 int8 = 0
                                                                    _ = v62
                                                                }
                                                                v63 := GenS7{F0: 639, F1: (v26 * (v56 | 4294967295)), F2: genFn1(128, v54)}
                                                                _ = v63
                                                            }
                                                            var v64 uint = v2
                                                            v65 := GenS8{F0: v13}
                                                            _ = v64
                                                            _ = v65
                                                        }
                                                        v2 = (v2 & uint(v26))
                                                        _ = v56
                                                    }
                                                    if false {
                                                        v66 := GenS9{F0: int64(v1), F1: v26}
                                                        fmt.Println(fi27)
                                                        var v67 string = "a"
                                                        fmt.Println(v26, v2)
                                                        _ = v66
                                                        _ = v67
                                                    }
                                                    var v68 bool = false
                                                    _ = v68
                                                } else {
                                                    fmt.Println(v2, fi27, v1)
                                                    var v69 uint16 = 32767
                                                    var v70 int = -9223372036854775808
                                                    _ = v69
                                                    _ = v70
                                                }
                                            default:
                                                var v71 uint32 = (v26 & 818)
                                                _ = v71
                                            }
                                            _ = fi53
                                            _ = v54
                                        }
                                        fmt.Println(v2)
                                        _ = v52
                                    }
                                case -178:
                                    var v72 int64 = -1
                                    _ = v72
                                }
                            }
                        case "a":
                            var v73 int64 = -400
                            _ = v73
                        default:
                            v74 := GenS5{F0: "foo bar", F1: 785, F2: (!true)}
                            _ = v74
                        }
                    }
                    v75 := GenS3{F0: true, F1: v5, F2: (v26 & (^(v26)))}
                    v5 += 0
                    _ = v75
                }
            }
            _ = fi27
            _ = v28
        }
    default:
        if (!false) {
            var v76 uint = (v2 / ((820) | 1))
            fmt.Println(v5)
            _ = v76
        } else {
            fmt.Println(v1)
            for fi77 := 0; fi77 < 16; fi77++ {
                var v78 uint32 = (^(v26))
                var v79 int64 = 0
                v80 := GenS1{F0: 100}
                _ = fi77
                _ = v78
                _ = v79
                _ = v80
            }
            if true {
                for fi81 := 0; fi81 < 10; fi81++ {
                    for fi82 := 0; fi82 < 6; fi82++ {
                        var v83 int32 = 0
                        _ = len(v3)
                        if true {
                            v2 = (v2 - v2)
                            var v84 string = "foo bar"
                            switch (v84 + "foo bar") {
                            case "\n\t":
                                var v85 string = v84
                                _ = v85
                            case "hello":
                                fmt.Println(v1, v26)
                            }
                            v83 = (-(v83))
                            var v86 int32 = 1
                            _ = v84
                            _ = v86
                        }
                        switch (fi82 << ((genFn0((128 < 0), 42)) & 63)) {
                        case -277:
                            var v87 int8 = -70
                            _ = v87
                        case 32767:
                            var v88 bool = false
                            _ = v88
                        default:
                            v5 = (^(v13))
                            if false {
                                if true {
                                    v89 := GenS7{F0: 0, F1: (v26 >> ((v2) & 63)), F2: v20}
                                    if (!(v83 <= 246)) {
                                        if false {
                                            switch (fi81 / (((fi81 / ((78) | 1))) | 1)) {
                                            case -9223372036854775808:
                                                switch (fi82 >> ((v2) & 63)) {
                                                case 2147483647:
                                                    switch 9223372036854775807 {
                                                    case 967:
                                                        fmt.Println(v83, fi82, v83)
                                                    case 0:
                                                        switch "a" {
                                                        case "hello":
                                                            var v90 uint8 = 255
                                                            _ = v90
                                                        case "a":
                                                            _ = len(v3)
                                                        default:
                                                            switch (-(fi82)) {
                                                            case 4294967295:
                                                                v5 += (v13 % (((v5 + 1)) | 1))
                                                            case 2147483647:
                                                                if false {
                                                                    var v91 int64 = -9223372036854775808
                                                                    v92 := GenS10{F0: 127, F1: 0}
                                                                    _ = v91
                                                                    _ = v92
                                                                } else {
                                                                    v93 := GenS11{F0: v1, F1: -9223372036854775808}
                                                                    if (v1 < v1) {
                                                                        delete(v4, "a")
                                                                        fmt.Println(v2, fi82)
                                                                        if false {
                                                                            if true {
                                                                                var v94 int = -2147483648
                                                                                if (!true) {
                                                                                    if len(v6) >= 2 {
                                                                                        v6 = v6[1:len(v6)]
                                                                                    }
                                                                                    fmt.Println(v26, v1)
                                                                                    fmt.Println(v94, v2, v5)
                                                                                }
                                                                                _ = v94
                                                                            } else {
                                                                                if (true && false) {
                                                                                    v95 := GenS12{F0: v83, F1: 675}
                                                                                    _ = v95
                                                                                }
                                                                                v96 := GenS2{F0: false, F1: (v1 % (((v1 / ((v1) | 1))) | 1))}
                                                                                v83 -= (v83 ^ v83)
                                                                                _ = v96
                                                                            }
                                                                            v13 += uint64(v26)
                                                                            v13 += 0
                                                                            var v97 string = ("\n\t" + ((("\n\t" + "\n\t") + "") + "a"))
                                                                            v98 := GenS13{F0: (v26 & uint32(v2)), F1: 2, F2: 42}
                                                                            switch ((v97 + (v97 + (v97 + "foo bar"))) + "foo bar") {
                                                                            case "hello":
                                                                                v20 -= fi82
                                                                            default:
                                                                                v99 := GenS14{F0: v26}
                                                                                _ = v99
                                                                            }
                                                                            var v100 int64 = -1
                                                                            _ = v100
                                                                            _ = v97
                                                                            _ = v98
                                                                        }
                                                                    }
                                                                    v26 = (v26 ^ 527)
                                                                    _ = v93
                                                                }
                                                            }
                                                            v83 = v83
                                                            var v101 float32 = float32(0.0)
                                                            _ = v101
                                                        }
                                                    }
                                                case -908:
                                                    v1 += v1
                                                default:
                                                    if (!(2 < 42)) {
                                                        var v102 int = fi82
                                                        _ = v102
                                                    }
                                                }
                                                var v103 int64 = 42
                                                _ = v103
                                            }
                                        }
                                        if (v26 < v26) {
                                            var v104 int16 = int16(v83)
                                            switch "foo bar" {
                                            case "\n\t":
                                                v5 -= (^(v13))
                                            case "hello":
                                                v20 = fi82
                                            default:
                                                v6 = append(v6, (("\n\t" + "") + "a"))
                                            }
                                            _ = v104
                                        }
                                    }
                                    _ = v89
                                } else {
                                    if true {
                                        v6 = append(v6, "a")
                                    }
                                    v105 := GenS11{F0: (v1 | v1), F1: 521}
                                    _ = v105
                                }
                                if ((!false) && false) {
                                    if false {
                                        if false {
                                            v1 = 208
                                            v5 -= v13
                                            var v106 float32 = (float32(float64((1.5 + (3.14159 / 1e10)))) - 1e10)
                                            _ = v106
                                        }
                                    }
                                }
                                fmt.Println(fi82, v20, fi81)
                                v2 = (v2 | (v2 & (v2 / ((v2) | 1))))
                            }
                            switch "a" {
                            case "":
                                var v107 int = fi82
                                _ = v107
                            case "hello":
                                if (v2 < genFn0((!((v26 * v26) != 2147483647)), 65535)) {
                                    delete(v3, (fi82 + genFn1(42, "a")))
                                    if true {
                                        v2 = genFn0((true && true), 65535)
                                    }
                                }
                            default:
                                if false {
                                    switch (fi82 | 468) {
                                    case 128:
                                        v2 = (v2 >> ((uint(v26)) & 63))
                                    case 2147483647:
                                        var v108 int = genFn1(11, "a")
                                        _ = v108
                                    }
                                    v26 = (v26 ^ v26)
                                    switch (fi82 % ((v20) | 1)) {
                                    case -2147483648:
                                        v3[int(v26)] = "foo bar"
                                    case 4294967295:
                                        v5 = v5
                                    default:
                                        switch (("\n\t" + "hello") + "\n\t") {
                                        case "":
                                            v109 := GenS2{F0: true, F1: (v1 / ((1) | 1))}
                                            _ = v109
                                        case "\n\t":
                                            if false {
                                                v1 -= -878
                                                var v110 int = (fi81 / (((v20 ^ v20)) | 1))
                                                _ = v110
                                            } else {
                                                switch v20 {
                                                case 42:
                                                    v20 = (fi81 & 42)
                                                    switch (v20 & (fi82 << (((v2 % (((v2 + (v2 << ((v2) & 63)))) | 1))) & 63))) {
                                                    case -542:
                                                        var v111 int = (v20 + fi82)
                                                        _ = v111
                                                    case 2147483647:
                                                        v112 := GenS6{F0: 1, F1: 295}
                                                        _ = v112
                                                    default:
                                                        var v113 bool = (false || true)
                                                        _ = v113
                                                    }
                                                    v26 = (v26 / ((711) | 1))
                                                    fmt.Println(v83, v5)
                                                    if true {
                                                        delete(v4, "a")
                                                        var v114 int16 = v1
                                                        var v115 uint = uint(v26)
                                                        _ = v114
                                                        _ = v115
                                                    } else {
                                                        v83 = v83
                                                        if (!(v1 < (v1 >> ((v2) & 63)))) {
                                                            if (!(!false)) {
                                                                v116 := GenS15{F0: int64(v83), F1: v20, F2: ""}
                                                                _ = v116
                                                            }
                                                        } else {
                                                            var v117 string = "\n\t"
                                                            switch genFn1(0, (v117 + v117)) {
                                                            case 2147483647:
                                                                var v118 string = v117
                                                                _ = v118
                                                            case 821:
                                                                v2 = 255
                                                            }
                                                            if false {
                                                                var v119 uint16 = 508
                                                                v5 += 4294967295
                                                                var v120 uint8 = 0
                                                                _ = v119
                                                                _ = v120
                                                            }
                                                            _ = v117
                                                        }
                                                    }
                                                case 628:
                                                    fmt.Println(v20, fi81, v83)
                                                default:
                                                    v121 := GenS4{F0: v2, F1: v1}
                                                    _ = v121
                                                }
                                            }
                                        case "hello":
                                            v122 := GenS10{F0: 17, F1: 812}
                                            _ = v122
                                        default:
                                            v1 -= v1
                                        }
                                        fmt.Println(v26, v83, v1)
                                    }
                                }
                                fmt.Println(v20, fi82, v20)
                            }
                        }
                        _ = fi82
                        _ = v83
                    }
                    switch 32768 {
                    case 2147483647:
                        if (false && false) {
                            switch genFn1(36, "foo bar") {
                            case 0:
                                switch (-(v20)) {
                                case 0:
                                    v2 -= genFn0((4294967295 <= (v26 >> (((v2 | (v2 / ((v2) | 1)))) & 63))), 299)
                                default:
                                    v1 = (v1 & 1000)
                                }
                            default:
                                v123 := GenS14{F0: 0}
                                _ = v123
                            }
                        }
                    case -9223372036854775808:
                        v124 := GenS3{F0: true, F1: 2, F2: (v26 * 216)}
                        _ = v124
                    case -474:
                        v13 -= 1
                    }
                    _ = fi81
                }
                if (91 != 57) {
                    if false {
                        var v125 int = v20
                        var v126 float32 = -2.25
                        _ = v125
                        _ = v126
                    } else {
                        v5 -= (v5 >> ((881) & 63))
                    }
                }
            }
        }
    }
    if (1e10 > (((float64(0.1) / (0.1 - (1.5 + 0.1))) - (-2.25 * (1e10 + 1e10))) - 1.5)) {
        if (false || false) {
            var v127 int32 = -576
            if (!false) {
                var v128 uint = genFn0(true, 1)
                switch ("a" + "hello") {
                case "hello":
                    _ = len(v6)
                }
                fmt.Println(v128, v5, v1)
                _ = v128
            } else {
                v129 := GenS2{F0: ((-(v127)) == v127), F1: (v1 + (v1 - -1))}
                v1 -= v1
                if false {
                    var v130 int8 = 127
                    for fi131 := 0; fi131 < 18; fi131++ {
                        if (3.14159 == (1.5 / -2.25)) {
                            v127 += (v127 / (((v127 & v127)) | 1))
                            if (1 < 1) {
                                if true {
                                    v26 -= (^(v26))
                                    v20 -= v20
                                }
                            }
                            fmt.Println(v1)
                        } else {
                            v20 -= genFn1(0, "\n\t")
                        }
                        var v132 int = (^(v20))
                        _ = fi131
                        _ = v132
                    }
                    _ = v130
                }
                _ = v129
            }
            _ = v127
        } else {
            var v133 uint = uint(v5)
            v134 := GenS16{F0: v13}
            _ = v133
            _ = v134
        }
    }
    for fi135 := 0; fi135 < 20; fi135++ {
        var v136 int8 = -19
        for fi137 := 0; fi137 < 15; fi137++ {
            if ((v5 ^ 18446744073709551615) != v5) {
                var v138 uint = (v2 << ((uint(v26)) & 63))
                var v139 int32 = -557
                if (((v1 & (v1 >> ((v2) & 63))) < (v1 % (((-(v1))) | 1))) && true) {
                    fmt.Println(v2, fi135, v136)
                    v6 = append(v6, "hello")
                }
                _ = v138
                _ = v139
            } else {
                v140 := v4["foo bar"]
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v141 := GenS17{F0: -63, F1: (v140 >> ((65535) & 63)), F2: 0}
                _ = v140
                _ = v141
            }
            v2 += 4294967295
            v2 = genFn0((("\n\t" + ((("\n\t" + "a") + "hello") + ((("" + ((((("foo bar" + "") + ("" + "a")) + "hello") + "hello") + "")) + ("hello" + ("a" + ""))) + (("" + "") + "")))) < ("hello" + "foo bar")), 32767)
            _ = fi137
        }
        _ = fi135
        _ = v136
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
