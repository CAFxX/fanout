package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 string; F1 uint64; F2 int }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS10 struct { F0 int32; F1 int16 }
type GenS11 struct { F0 int; F1 int32 }
type GenS12 struct { F0 uint16; F1 int16; F2 int8 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int64; F1 uint16 }
type GenS15 struct { F0 int64; F1 uint16 }
type GenS16 struct { F0 string; F1 bool }
type GenS17 struct { F0 string; F1 bool }
type GenS18 struct { F0 int32; F1 uint; F2 int }
type GenS19 struct { F0 uint64; F1 string; F2 int64 }
func genFn0(p0 int16, p1 bool, p2 uint16) int32 {
    v4 := []int{9223372036854775807, 127, -286, -1}
    for fi5 := 0; fi5 < 16; fi5++ {
        for fi6 := 0; fi6 < 6; fi6++ {
            p0 -= p0
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            p1 = (false || p1)
            _ = fi6
        }
        switch 32767 {
        case 0:
            v7 := GenS1{F0: p1, F1: 0}
            _ = v7
        case -19:
            for fi8 := 0; fi8 < 8; fi8++ {
                var v9 uint32 = uint32(p2)
                if ((!("hello" >= "")) && ((255 >= 1) || (!p1))) {
                    fmt.Println(p0)
                    if true {
                        v10 := GenS1{F0: p1, F1: (v9 * (v9 % (((v9 * v9)) | 1)))}
                        _ = v10
                    }
                    var v11 string = ("foo bar" + "a")
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
        case 42:
            switch ((((("" + ("hello" + "\n\t")) + ("foo bar" + "a")) + ("" + ("foo bar" + ""))) + ("foo bar" + (("foo bar" + "foo bar") + "\n\t"))) + "\n\t") {
            case "foo bar":
                fmt.Println(p0)
            case "":
                var v12 int64 = 1
                _ = v12
            default:
                var v13 float64 = -2.25
                _ = v13
            }
        default:
            p2 += (p2 >> ((65535) & 63))
        }
        _ = fi5
    }
    if (p1 || (!((("foo bar" + ("foo bar" + ("hello" + "\n\t"))) + "\n\t") >= "a"))) {
        switch -202 {
        case 0:
            if false {
                if (float64(0.1) < 3.14159) {
                    var v14 int32 = 0
                    _ = v14
                }
                for fi15 := 0; fi15 < 14; fi15++ {
                    if (2 >= 18446744073709551615) {
                        fmt.Println(p2, fi15, p0)
                        p1 = ((!false) && ((0.0 / 0.1) > 0.0))
                        v4 = append(v4, fi15)
                    } else {
                        var v16 int8 = 1
                        var v17 uint32 = 1
                        _ = v16
                        _ = v17
                    }
                    fmt.Println(p1, p0, fi15)
                    _ = fi15
                }
                switch 82 {
                case 4294967295:
                    var v18 uint32 = 65535
                    _ = v18
                case -153:
                    p1 = (p0 != 83)
                case -700:
                    fmt.Println(p1, p1, p1)
                default:
                    var v19 bool = (-70 <= 0)
                    _ = v19
                }
            }
        }
        var v20 float64 = -2.25
        _ = v20
    }
    switch ("\n\t" + ("\n\t" + "hello")) {
    case "\n\t":
        v21 := GenS2{F0: 77, F1: int8(p2), F2: 255}
        _ = v21
    case "a":
        v22 := GenS2{F0: 128, F1: 1, F2: 2}
        _ = v22
    case "foo bar":
        _ = len(v4)
    default:
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
    }
    fmt.Println(p1, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -823
}
func genFn1(p0 uint, p1 int32, p2 uint32) uint16 {
    if (!(p1 < 0)) {
        p2 -= (p2 % ((p2) | 1))
        var v23 bool = (2147483647 < 42)
        _ = v23
    } else {
        v24 := GenS3{F0: 120}
        _ = v24
    }
    v25 := GenS4{F0: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    return 0
}

func main() {
    v1 := []string{("" + ""), "foo bar"}
    v2 := []int{-310, 4294967295, 2147483647, 9223372036854775807}
    var v3 uint = 65535
    var v26 uint16 = 65535
    for fi27 := 0; fi27 < 20; fi27++ {
        var v28 float64 = (1e10 - float64((1.5 * 1.5)))
        for fi29 := 0; fi29 < 20; fi29++ {
            var v30 int32 = -1
            v31 := GenS2{F0: 2, F1: 79, F2: 127}
            _ = fi29
            _ = v30
            _ = v31
        }
        v26 += v26
        _ = fi27
        _ = v28
    }
    v1 = append(v1, "hello")
    for fi32 := 0; fi32 < 12; fi32++ {
        if (genFn1(v3, 1, 493) < v26) {
            if false {
                if false {
                    v3 = (v3 << ((uint(v26)) & 63))
                }
                for fi33 := 0; fi33 < 3; fi33++ {
                    switch ("hello" + ("hello" + "foo bar")) {
                    case "hello":
                        var v34 uint8 = uint8(v3)
                        _ = v34
                    case "\n\t":
                        _ = len(v1)
                        v3 -= 2
                        var v35 int = (fi33 << ((18446744073709551615) & 63))
                        _ = v35
                    default:
                        if (!(fi33 >= (fi32 + -945))) {
                            v3 = (v3 - 4294967295)
                            if (36 >= 65535) {
                                var v36 uint64 = 4294967295
                                _ = v36
                            } else {
                                if ((2 < 2) || ((-2.25 - 1.5) > 1e10)) {
                                    v37 := GenS3{F0: 100}
                                    v38 := GenS2{F0: 0, F1: -89, F2: 85}
                                    _ = v37
                                    _ = v38
                                } else {
                                    v39 := GenS1{F0: true, F1: 2147483648}
                                    v3 += (v3 - (v3 % ((v3) | 1)))
                                    _ = v39
                                }
                            }
                        } else {
                            v3 -= v3
                            v26 = (v26 % ((2) | 1))
                        }
                    }
                    v3 += v3
                    switch fi33 {
                    case 667:
                        v2 = append(v2, 127)
                    case 2147483647:
                        v2 = append(v2, (fi32 ^ (fi32 >> (((v3 ^ (^(v3)))) & 63))))
                    }
                    _ = fi33
                }
                if false {
                    v26 += 2
                    var v40 int64 = 9223372036854775807
                    switch "" {
                    case "\n\t":
                        var v41 float32 = 3.14159
                        _ = v41
                    case "a":
                        v40 = 9223372036854775807
                    case "":
                        for fi42 := 0; fi42 < 3; fi42++ {
                            v43 := GenS2{F0: 6, F1: int8(v40), F2: 2}
                            var v44 int8 = 61
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                    default:
                        switch fi32 {
                        case -1:
                            var v45 uint = v3
                            _ = v45
                        case -9223372036854775808:
                            _ = len(v1)
                        default:
                            for fi46 := 0; fi46 < 6; fi46++ {
                                var v47 float32 = 0.1
                                _ = len(v2)
                                v40 = v40
                                _ = fi46
                                _ = v47
                            }
                        }
                    }
                    _ = v40
                }
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v48 string = "hello"
                _ = v48
            }
        } else {
            v49 := GenS5{F0: 279, F1: 0}
            _ = v49
        }
        _ = fi32
    }
    if (v26 > v26) {
        v50 := GenS6{F0: "", F1: 2, F2: 42}
        v26 -= 469
        v51 := GenS4{F0: 2}
        _ = v50
        _ = v51
    }
    switch "\n\t" {
    case "a":
        var v52 int = -9223372036854775808
        _ = v52
    case "hello":
        for fi53 := 0; fi53 < 18; fi53++ {
            if true {
                switch -457 {
                case 369:
                    v3 -= 4294967295
                case 4294967295:
                    switch (-(fi53)) {
                    case 32767:
                        _ = len(v1)
                    case -1:
                        fmt.Println(v3)
                    case -243:
                        switch "" {
                        case "foo bar":
                            switch "foo bar" {
                            case "\n\t":
                                for fi54 := 0; fi54 < 3; fi54++ {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    fmt.Println(fi54)
                                    _ = fi54
                                }
                            case "foo bar":
                                switch -1 {
                                case 0:
                                    var v55 float32 = 1.5
                                    _ = v55
                                default:
                                    if (true && false) {
                                        v26 = v26
                                    }
                                }
                            }
                        case "a":
                            if true {
                                v56 := GenS7{F0: -32768}
                                for fi57 := 0; fi57 < 18; fi57++ {
                                    switch 32768 {
                                    case -376:
                                        v58 := GenS1{F0: false, F1: 2147483648}
                                        _ = v58
                                    default:
                                        fmt.Println(fi57, fi53, v3)
                                    }
                                    v26 -= (v26 - v26)
                                    v26 += v26
                                    v59 := GenS8{F0: 42}
                                    _ = fi57
                                    _ = v59
                                }
                                _ = v56
                            } else {
                                var v60 uint32 = 987
                                _ = v60
                            }
                        default:
                            if false {
                                v26 -= (v26 | (v26 >> ((v3) & 63)))
                            } else {
                                if (219 <= 128) {
                                    for fi61 := 0; fi61 < 11; fi61++ {
                                        if (-2147483648 <= 895) {
                                            v3 -= v3
                                            var v62 uint64 = 2
                                            fmt.Println(v62, v62)
                                            _ = v62
                                        } else {
                                            fmt.Println(fi53, v3)
                                            if true {
                                                switch (("foo bar" + ("\n\t" + "a")) + "\n\t") {
                                                case "hello":
                                                    var v63 string = "\n\t"
                                                    _ = v63
                                                }
                                                var v64 uint32 = 4294967295
                                                _ = v64
                                            } else {
                                                var v65 int16 = 42
                                                v66 := GenS9{F0: 2147483648, F1: (v26 & (v26 | genFn1(uint(v26), 2147483647, 709))), F2: 0}
                                                _ = v65
                                                _ = v66
                                            }
                                            v67 := GenS6{F0: "", F1: uint64(v26), F2: (^(fi53))}
                                            _ = v67
                                        }
                                        var v68 uint32 = 2147483648
                                        if (-30 >= 93) {
                                            v69 := GenS3{F0: -10}
                                            _ = v69
                                        }
                                        v70 := GenS2{F0: 197, F1: 1, F2: 88}
                                        if len(v2) > 0 {
                                            v2[0] = fi53
                                        }
                                        _ = len(v2)
                                        _ = fi61
                                        _ = v68
                                        _ = v70
                                    }
                                    v71 := GenS2{F0: 127, F1: -125, F2: 81}
                                    switch ("hello" + ("\n\t" + "foo bar")) {
                                    case "a":
                                        for fi72 := 0; fi72 < 11; fi72++ {
                                            if (42 == 2) {
                                                v73 := GenS7{F0: 1}
                                                if len(v2) > 0 {
                                                    v2[0] = (fi72 ^ 4294967295)
                                                }
                                                if (!false) {
                                                    v26 += (v26 ^ v26)
                                                } else {
                                                    var v74 uint8 = 65
                                                    _ = v74
                                                }
                                                _ = v73
                                            } else {
                                                if false {
                                                    v3 = (v3 & v3)
                                                    fmt.Println(v26, fi72)
                                                    v75 := GenS1{F0: (!false), F1: 921}
                                                    _ = v75
                                                } else {
                                                    v76 := GenS8{F0: 1}
                                                    switch (fi53 % (((fi53 - fi72)) | 1)) {
                                                    case -80:
                                                        v3 -= 4294967295
                                                        v3 = (v3 - v3)
                                                    default:
                                                        if true {
                                                            v26 += v26
                                                            v77 := GenS5{F0: 32767, F1: 1}
                                                            _ = v77
                                                        }
                                                    }
                                                    _ = len(v2)
                                                    _ = v76
                                                }
                                                v78 := GenS10{F0: 42, F1: 1}
                                                if ("hello" < "a") {
                                                    var v79 int = (fi72 & fi53)
                                                    v80 := GenS7{F0: -299}
                                                    if (true || false) {
                                                        if true {
                                                            if (0 < -32768) {
                                                                if false {
                                                                    v3 = 42
                                                                }
                                                                if (!(2 <= 1)) {
                                                                    if (!(true || false)) {
                                                                        if len(v1) > 0 {
                                                                            v1[0] = ""
                                                                        }
                                                                        if true {
                                                                            v81 := GenS11{F0: -818, F1: 2147483647}
                                                                            switch "a" {
                                                                            case "":
                                                                                v3 -= v3
                                                                            case "foo bar":
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            }
                                                                            v26 -= (v26 << ((42) & 63))
                                                                            fmt.Println(v3, fi72)
                                                                            switch "" {
                                                                            case "hello":
                                                                                var v82 uint32 = 88
                                                                                _ = v82
                                                                            case "\n\t":
                                                                                v83 := GenS7{F0: 32767}
                                                                                _ = v83
                                                                            case "a":
                                                                                if false {
                                                                                    switch ("\n\t" + "a") {
                                                                                    case "":
                                                                                        v26 -= 42
                                                                                    }
                                                                                    var v84 int32 = 492
                                                                                    switch -990 {
                                                                                    case -2147483648:
                                                                                        switch 128 {
                                                                                        case -168:
                                                                                            if false {
                                                                                                if (1000 == 1000) {
                                                                                                    if (int8(v84) < 127) {
                                                                                                        var v85 uint16 = genFn1(v3, (v84 - -2147483648), 4294967295)
                                                                                                        v86 := GenS3{F0: 1}
                                                                                                        if false {
                                                                                                            switch fi53 {
                                                                                                            case -9223372036854775808:
                                                                                                                if (false && false) {
                                                                                                                    var v87 int = -9223372036854775808
                                                                                                                    if false {
                                                                                                                        fmt.Println(fi72, v79, v87)
                                                                                                                        v3 -= v3
                                                                                                                    }
                                                                                                                    _ = v87
                                                                                                                }
                                                                                                            case 32767:
                                                                                                                switch ("\n\t" + "hello") {
                                                                                                                case "":
                                                                                                                    v84 = (v84 * 561)
                                                                                                                case "\n\t":
                                                                                                                    var v88 float64 = 1.5
                                                                                                                    _ = v88
                                                                                                                case "a":
                                                                                                                    v89 := GenS12{F0: (v85 + (v85 << (((^(v3))) & 63))), F1: 707, F2: 100}
                                                                                                                    _ = v89
                                                                                                                default:
                                                                                                                    if false {
                                                                                                                        var v90 bool = false
                                                                                                                        _ = v90
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        if false {
                                                                                                            _ = len(v2)
                                                                                                            var v91 string = ("foo bar" + "hello")
                                                                                                            v26 = (^(v85))
                                                                                                            _ = v91
                                                                                                        }
                                                                                                        v26 -= genFn1(926, (v84 & (v84 % ((-556) | 1))), 201)
                                                                                                        var v92 int8 = -128
                                                                                                        var v93 int32 = v84
                                                                                                        _ = v85
                                                                                                        _ = v86
                                                                                                        _ = v92
                                                                                                        _ = v93
                                                                                                    }
                                                                                                    var v94 bool = true
                                                                                                    switch (^(fi53)) {
                                                                                                    case -2147483648:
                                                                                                        if v94 {
                                                                                                            var v95 int64 = 0
                                                                                                            _ = len(v2)
                                                                                                            _ = v95
                                                                                                        } else {
                                                                                                            var v96 int16 = 1000
                                                                                                            _ = v96
                                                                                                        }
                                                                                                    case 244:
                                                                                                        fmt.Println(v84)
                                                                                                    }
                                                                                                    _ = v94
                                                                                                } else {
                                                                                                    v3 += (v3 >> ((119) & 63))
                                                                                                    v97 := GenS5{F0: 32767, F1: -77}
                                                                                                    var v98 uint16 = (^(v26))
                                                                                                    _ = v97
                                                                                                    _ = v98
                                                                                                }
                                                                                                v99 := GenS13{F0: 0}
                                                                                                switch (-(v79)) {
                                                                                                case 4294967295:
                                                                                                    fmt.Println(v3)
                                                                                                case 358:
                                                                                                    v100 := GenS14{F0: 816, F1: (v26 | (v26 << ((740) & 63)))}
                                                                                                    _ = v100
                                                                                                case -932:
                                                                                                    var v101 uint32 = uint32(v3)
                                                                                                    _ = v101
                                                                                                }
                                                                                                _ = v99
                                                                                            } else {
                                                                                                v1 = append(v1, (("\n\t" + "foo bar") + (("hello" + "hello") + (("foo bar" + ("a" + "foo bar")) + "a"))))
                                                                                            }
                                                                                        case -2147483648:
                                                                                            fmt.Println(v26, v79, fi72)
                                                                                        }
                                                                                    case 32768:
                                                                                        if (true && (!(true && (((18446744073709551615 >= 18446744073709551615) && (-149 < 0)) || false)))) {
                                                                                            v1 = append(v1, "hello")
                                                                                            fmt.Println(fi53, v3, fi53)
                                                                                            var v102 uint32 = 2147483648
                                                                                            _ = v102
                                                                                        }
                                                                                    case 1:
                                                                                        fmt.Println(v26, fi53)
                                                                                    default:
                                                                                        switch fi53 {
                                                                                        case 4294967295:
                                                                                            v103 := GenS15{F0: 0, F1: genFn1(v3, (v84 / ((v84) | 1)), 5)}
                                                                                            _ = v103
                                                                                        case 144:
                                                                                            fmt.Println(v3, fi53, v3)
                                                                                        default:
                                                                                            fmt.Println(v3)
                                                                                        }
                                                                                    }
                                                                                    v3 -= v3
                                                                                    fmt.Println(fi53, v79, v79)
                                                                                    var v104 string = ("" + "")
                                                                                    _ = v104
                                                                                    _ = v84
                                                                                }
                                                                            }
                                                                            _ = v81
                                                                        } else {
                                                                            var v105 int64 = 0
                                                                            if (false && true) {
                                                                                switch (fi72 / ((fi53) | 1)) {
                                                                                case 0:
                                                                                    v106 := GenS16{F0: "\n\t", F1: true}
                                                                                    _ = v106
                                                                                case -9223372036854775808:
                                                                                    var v107 int16 = 0
                                                                                    _ = v107
                                                                                case -199:
                                                                                    v108 := GenS3{F0: 42}
                                                                                    _ = v108
                                                                                }
                                                                            } else {
                                                                                if (false || (32767 != -32768)) {
                                                                                    fmt.Println(v26, fi72, fi53)
                                                                                } else {
                                                                                    v105 -= -9223372036854775808
                                                                                }
                                                                                v79 = fi53
                                                                            }
                                                                            _ = v105
                                                                        }
                                                                        switch fi53 {
                                                                        case -1:
                                                                            var v109 uint = (v3 * 0)
                                                                            _ = v109
                                                                        default:
                                                                            v26 = (v26 + v26)
                                                                        }
                                                                    }
                                                                    if true {
                                                                        if false {
                                                                            var v110 int8 = 100
                                                                            var v111 int32 = -424
                                                                            v112 := GenS9{F0: 42, F1: v26, F2: 968}
                                                                            _ = v110
                                                                            _ = v111
                                                                            _ = v112
                                                                        }
                                                                    } else {
                                                                        var v113 int8 = -1
                                                                        var v114 int64 = 182
                                                                        if true {
                                                                            v3 += v3
                                                                        }
                                                                        _ = v113
                                                                        _ = v114
                                                                    }
                                                                }
                                                                v115 := GenS9{F0: 42, F1: v26, F2: 2}
                                                                _ = v115
                                                            } else {
                                                                switch (fi72 / ((2147483647) | 1)) {
                                                                case 32767:
                                                                    if (float64(float32(0.0)) < 0.0) {
                                                                        var v116 int = (fi53 / ((fi72) | 1))
                                                                        _ = v116
                                                                    }
                                                                case 127:
                                                                    v26 = 2
                                                                }
                                                                var v117 int8 = -74
                                                                var v118 int8 = (^(v117))
                                                                _ = v117
                                                                _ = v118
                                                            }
                                                            v26 -= (v26 / (((v26 / (((v26 >> ((v3) & 63))) | 1))) | 1))
                                                        } else {
                                                            if (true && false) {
                                                                if ((-2.25 + 0.0) == 1.5) {
                                                                    v79 = fi72
                                                                    var v119 uint32 = 2
                                                                    _ = v119
                                                                }
                                                                var v120 uint16 = (^(v26))
                                                                if (!false) {
                                                                    _ = len(v1)
                                                                }
                                                                _ = v120
                                                            }
                                                            if true {
                                                                v26 = (v26 + 0)
                                                                var v121 uint32 = 101
                                                                var v122 uint32 = v121
                                                                _ = v121
                                                                _ = v122
                                                            } else {
                                                                switch v79 {
                                                                case 2147483647:
                                                                    v123 := GenS5{F0: 57, F1: 0}
                                                                    _ = v123
                                                                case 9223372036854775807:
                                                                    var v124 int64 = -1
                                                                    _ = v124
                                                                case 32767:
                                                                    if len(v2) > 0 {
                                                                        v2[0] = -612
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        v125 := GenS17{F0: ((("" + ("" + "\n\t")) + "a") + ""), F1: (false || false)}
                                                        v3 = 0
                                                        _ = v125
                                                    }
                                                    var v126 float32 = (3.14159 * 0.1)
                                                    switch 32768 {
                                                    case 32768:
                                                        var v127 int64 = -580
                                                        _ = v127
                                                    case -31:
                                                        v26 = genFn1((v3 - (^(v3))), 42, 1)
                                                    }
                                                    var v128 bool = (!(213 < 4294967295))
                                                    v1 = append(v1, "a")
                                                    _ = v126
                                                    _ = v128
                                                    _ = v79
                                                    _ = v80
                                                } else {
                                                    var v129 float64 = (0.1 * (3.14159 + (3.14159 - 1e10)))
                                                    v3 = v3
                                                    var v130 bool = (!true)
                                                    _ = v129
                                                    _ = v130
                                                }
                                                _ = v78
                                            }
                                            if false {
                                                if len(v1) > 0 {
                                                    v1[0] = "a"
                                                }
                                            }
                                            var v131 string = ("" + ("foo bar" + ("hello" + "")))
                                            _ = fi72
                                            _ = v131
                                        }
                                    }
                                    _ = v71
                                }
                                v132 := GenS9{F0: 1, F1: (^(v26)), F2: 639}
                                _ = v132
                            }
                        }
                    }
                case -9223372036854775808:
                    if true {
                        for fi133 := 0; fi133 < 11; fi133++ {
                            switch fi53 {
                            case -690:
                                var v134 uint32 = 911
                                _ = v134
                            }
                            v3 += 568
                            _ = fi133
                        }
                    } else {
                        fmt.Println(v26)
                    }
                }
            } else {
                v135 := GenS18{F0: -2147483648, F1: (v3 & (v3 * (v3 & v3))), F2: fi53}
                _ = v135
            }
            if true {
                for fi136 := 0; fi136 < 1; fi136++ {
                    v26 += v26
                    var v137 int8 = 20
                    v26 -= (v26 >> ((v3) & 63))
                    _ = fi136
                    _ = v137
                }
            } else {
                v26 = (^(v26))
                var v138 uint64 = 0
                v26 = genFn1(0, 42, 773)
                _ = v138
            }
            v26 -= (v26 - genFn1(1, 42, 4294967295))
            _ = fi53
        }
    case "foo bar":
        v26 -= (v26 ^ genFn1(v3, -643, 4294967295))
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (!true) {
        var v139 int64 = -1
        var v140 uint32 = 2
        _ = v139
        _ = v140
    }
    for fi141 := 0; fi141 < 6; fi141++ {
        v3 += v3
        var v142 uint32 = 42
        _ = fi141
        _ = v142
    }
    for fi143 := 0; fi143 < 2; fi143++ {
        _ = len(v1)
        var v144 float32 = -2.25
        v26 = 49
        _ = fi143
        _ = v144
    }
    v145 := GenS19{F0: 1, F1: "", F2: -9223372036854775808}
    var v146 int8 = 100
    v3 = (v3 + (v3 / ((v3) | 1)))
    switch ("foo bar" + "a") {
    case "foo bar":
        v26 = v26
    }
    if false {
        fmt.Println(v146, v26, v3)
    }
    fmt.Println(v3)
    var v147 bool = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v146:", v146)
    fmt.Println("v147:", v147)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v145
    _ = v146
    _ = v147
    _ = v2
    _ = v26
    _ = v3
}
