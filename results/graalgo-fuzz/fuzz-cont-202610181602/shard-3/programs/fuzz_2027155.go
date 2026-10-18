package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 uint8; F1 int; F2 int64 }
type GenS7 struct { F0 int64; F1 int64 }
type GenS8 struct { F0 int8; F1 int8 }
type GenS9 struct { F0 int8; F1 uint; F2 int32 }
type GenS10 struct { F0 uint; F1 int16 }
type GenS11 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS12 struct { F0 int }
func genFn0(p0 uint) uint32 {
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 int16 = 13
        _ = fi6
        _ = v7
    }
    p0 += p0
    v8 := []int{32767, 165, 0, -1}
    for fi9 := 0; fi9 < 6; fi9++ {
        if len(v8) > 0 {
            v8[0] = 128
        }
        if (false || (1000 == -882)) {
            p0 -= (p0 - 113)
            var v10 uint16 = 471
            v10 += v10
            _ = v10
        }
        _ = fi9
    }
    _ = p0
    _ = v8
    return 946
}
func genFn1(p0 int64) bool {
    v11 := GenS1{F0: 1}
    p0 = p0
    switch -1 {
    case -1:
        if (!true) {
            fmt.Println(p0)
            switch -2147483648 {
            case 0:
                v12 := []int{41}
                _ = v12
            case 570:
                p0 -= p0
            case 4294967295:
                v13 := GenS1{F0: 42}
                _ = v13
            }
        } else {
            var v14 int16 = 490
            var v15 int64 = p0
            for fi16 := 0; fi16 < 6; fi16++ {
                v15 = 0
                _ = fi16
            }
            _ = v14
            _ = v15
        }
    case -759:
        var v17 uint32 = 733
        _ = v17
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 bool = (4294967295 != 80)
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = v11
    return (-899 != -458)
}

func main() {
    var v1 string = "hello"
    var v2 float64 = (1.5 / 1.5)
    v3 := []string{"a"}
    v4 := []string{}
    var v5 int = 2147483647
    switch "hello" {
    case "a":
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 uint32 = 29
            var v22 int8 = -128
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    v23 := GenS2{F0: 1, F1: v1}
    v5 += (-(v5))
    v24 := GenS3{F0: 4294967295, F1: 1}
    if (9223372036854775807 == (v5 | 127)) {
        var v25 int64 = 1
        fmt.Println(v5, v25)
        if len(v4) > 0 {
            v4[0] = ((v1 + ("" + "a")) + "hello")
        }
        _ = v25
    } else {
        v26 := GenS4{F0: false}
        _ = v26
    }
    var v27 int = (v5 * v5)
    if (1 <= 2) {
        for fi28 := 0; fi28 < 3; fi28++ {
            var v29 float64 = (1e10 / 3.14159)
            v5 = (fi28 * (v5 >> ((407) & 63)))
            v1 = "\n\t"
            _ = fi28
            _ = v29
        }
        var v30 bool = (-9223372036854775808 > 9223372036854775807)
        _ = v30
    } else {
        v3 = append(v3, v1)
    }
    if true {
        v5 -= 1
        for fi31 := 0; fi31 < 17; fi31++ {
            fmt.Println(v1, fi31, v1)
            _ = fi31
        }
        v1 = v1
    }
    if (!true) {
        if true {
            if ("" >= "") {
                for fi32 := 0; fi32 < 1; fi32++ {
                    _ = len(v4)
                    _ = fi32
                }
                for fi33 := 0; fi33 < 3; fi33++ {
                    switch v1 {
                    case "hello":
                        if true {
                            var v34 bool = ((((v27 / ((v27) | 1)) <= v5) || (true && true)) || false)
                            fmt.Println(v2, v27)
                            for fi35 := 0; fi35 < 17; fi35++ {
                                if (91 <= 255) {
                                    if (false && (127 <= -36)) {
                                        if ((159 != 42) && true) {
                                            v36 := GenS2{F0: 42, F1: v1}
                                            v1 = ("hello" + (v1 + v1))
                                            if ((v1 >= "") || v34) {
                                                v37 := GenS5{F0: v1, F1: 241}
                                                _ = v37
                                            }
                                            _ = v36
                                        }
                                        fmt.Println(fi33, fi33, v5)
                                        var v38 float32 = 0.1
                                        _ = v38
                                    }
                                    v1 = (v1 + (v1 + "foo bar"))
                                    switch v1 {
                                    case "a":
                                        v1 = v1
                                    case "":
                                        var v39 string = ((v1 + (("\n\t" + ((((v1 + "a") + v1) + ((v1 + (v1 + "\n\t")) + v1)) + v1)) + v1)) + "\n\t")
                                        if len(v3) > 0 {
                                            v3[0] = "a"
                                        }
                                        var v40 int16 = 700
                                        _ = v39
                                        _ = v40
                                    case "\n\t":
                                        v1 = ((v1 + v1) + "\n\t")
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = "foo bar"
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = v1
                                    }
                                }
                                _ = fi35
                            }
                            _ = v34
                        }
                    case "\n\t":
                        var v41 int8 = -117
                        _ = v41
                    }
                    _ = fi33
                }
                v4 = append(v4, v1)
            }
        } else {
            for fi42 := 0; fi42 < 6; fi42++ {
                var v43 uint32 = 0
                fmt.Println(v43, v43, v1)
                v27 += v5
                _ = fi42
                _ = v43
            }
            if (float64(1e10) < v2) {
                v44 := GenS1{F0: 42}
                var v45 uint8 = 255
                _ = v44
                _ = v45
            }
        }
        switch "a" {
        case "":
            if (-635 > 9223372036854775807) {
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
                var v46 uint8 = 128
                _ = v46
            }
        }
        fmt.Println(v2)
    } else {
        if (!(v2 >= 1.5)) {
            for fi47 := 0; fi47 < 16; fi47++ {
                _ = len(v4)
                _ = fi47
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            var v48 uint16 = 652
            _ = v48
        }
        switch "foo bar" {
        case "":
            for fi49 := 0; fi49 < 10; fi49++ {
                v1 = ""
                switch ((((((v1 + v1) + "\n\t") + ("\n\t" + "a")) + (v1 + v1)) + ("\n\t" + v1)) + "hello") {
                case "a":
                    if false {
                        var v50 float32 = -2.25
                        var v51 int8 = -128
                        switch v5 {
                        case -1:
                            fmt.Println(v5)
                        case 1:
                            var v52 uint32 = 4294967295
                            _ = v52
                        case 2147483647:
                            _ = len(v3)
                        }
                        _ = v50
                        _ = v51
                    }
                case "foo bar":
                    if len(v4) > 0 {
                        v4[0] = v1
                    }
                case "\n\t":
                    switch v1 {
                    case "":
                        switch "" {
                        case "foo bar":
                            var v53 uint32 = 2
                            _ = v53
                        case "":
                            fmt.Println(v5, v1)
                        case "hello":
                            for fi54 := 0; fi54 < 20; fi54++ {
                                if (true || true) {
                                    fmt.Println(fi49, fi49, v5)
                                    v5 = (v27 / (((fi54 - v27)) | 1))
                                    var v55 string = "\n\t"
                                    _ = v55
                                } else {
                                    v2 -= (3.14159 * v2)
                                }
                                if (((!(v2 != (v2 + float64(0.1)))) && (3.14159 != ((((v2 * v2) * ((v2 - -2.25) * v2)) * v2) * (v2 + (v2 - (1.5 / 0.1)))))) && (((-1 != -9223372036854775808) && (-444 <= 595)) || (false && true))) {
                                    v3 = append(v3, "foo bar")
                                    if ("\n\t" == "") {
                                        v27 = (fi54 * v27)
                                        var v56 float32 = 0.0
                                        _ = v56
                                    }
                                    switch v1 {
                                    case "":
                                        v5 = (v5 << ((4294967295) & 63))
                                    case "hello":
                                        switch (v27 ^ 127) {
                                        case -22:
                                            v1 = v1
                                            v57 := GenS4{F0: true}
                                            _ = v57
                                        }
                                    case "\n\t":
                                        v58 := GenS4{F0: ((!false) || true)}
                                        _ = v58
                                    default:
                                        v2 += v2
                                    }
                                    fmt.Println(v2, v27, v27)
                                }
                                var v59 int8 = 1
                                _ = fi54
                                _ = v59
                            }
                        }
                    case "foo bar":
                        switch v1 {
                        case "hello":
                            if (42 >= 1) {
                                fmt.Println(v5, v1)
                                v60 := GenS5{F0: (v1 + "a"), F1: 2147483648}
                                switch v1 {
                                case "\n\t":
                                    fmt.Println(fi49, v2)
                                default:
                                    v2 += 1.5
                                }
                                _ = v60
                            }
                        case "\n\t":
                            switch v1 {
                            case "\n\t":
                                var v61 uint8 = 255
                                _ = v61
                            case "hello":
                                switch v1 {
                                case "":
                                    v2 += v2
                                case "hello":
                                    if true {
                                        var v62 uint32 = 2147483648
                                        _ = v62
                                    }
                                case "a":
                                    v63 := GenS6{F0: 128, F1: -1, F2: 42}
                                    _ = v63
                                }
                            case "foo bar":
                                v27 -= 221
                            default:
                                fmt.Println(v1, fi49)
                            }
                        }
                    }
                default:
                    v27 -= (-(v27))
                }
                _ = fi49
            }
        case "hello":
            switch 32767 {
            case 0:
                for fi64 := 0; fi64 < 14; fi64++ {
                    switch ((v1 + "") + (v1 + "")) {
                    case "":
                        fmt.Println(v5, v2)
                    default:
                        var v65 string = (v1 + v1)
                        _ = v65
                    }
                    v66 := GenS6{F0: 0, F1: (fi64 - (-(v5))), F2: 42}
                    switch (v1 + v1) {
                    case "":
                        v67 := GenS1{F0: 42}
                        _ = v67
                    default:
                        v27 += v5
                    }
                    _ = fi64
                    _ = v66
                }
            case 128:
                switch v1 {
                case "hello":
                    var v68 uint8 = 147
                    _ = v68
                case "a":
                    v1 = (v1 + v1)
                case "foo bar":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    switch (v5 | (v27 % ((v27) | 1))) {
                    case -9223372036854775808:
                        v3 = append(v3, "a")
                    case 2147483647:
                        v27 += (-(v5))
                    }
                }
            }
        default:
            var v69 uint32 = 4294967295
            _ = v69
        }
    }
    if len(v3) > 0 {
        v3[0] = v1
    }
    var v70 int64 = 0
    v27 -= (v5 ^ v27)
    for fi71 := 0; fi71 < 5; fi71++ {
        _ = len(v4)
        v72 := GenS7{F0: v70, F1: 0}
        var v73 uint16 = 65535
        _ = fi71
        _ = v72
        _ = v73
    }
    for fi74 := 0; fi74 < 18; fi74++ {
        for fi75 := 0; fi75 < 3; fi75++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v4 = append(v4, v1)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi75
        }
        v2 -= v2
        _ = fi74
    }
    v76 := GenS7{F0: 9223372036854775807, F1: 854}
    if ((true || true) || ((3.14159 * float32(0.1)) != 1e10)) {
        fmt.Println(v2, v70)
        fmt.Println(v1, v1)
        v4 = append(v4, (v1 + ("a" + v1)))
    }
    switch "\n\t" {
    case "a":
        var v77 uint8 = 38
        _ = v77
    case "foo bar":
        if ((32768 <= (v5 - v5)) || false) {
            v27 = (v5 + 4294967295)
            v5 += v5
            for fi78 := 0; fi78 < 10; fi78++ {
                for fi79 := 0; fi79 < 14; fi79++ {
                    if (147 >= 48) {
                        _ = len(v3)
                        if true {
                            if true {
                                var v80 int16 = 476
                                _ = v80
                            }
                            v70 = v70
                            fmt.Println(fi78)
                            v3 = append(v3, v1)
                            v5 += (fi78 >> ((255) & 63))
                        }
                        v2 += 0.0
                        var v81 int64 = int64(fi78)
                        _ = v81
                    } else {
                        v70 = (v70 << ((1) & 63))
                        fmt.Println(v70, fi78, v1)
                    }
                    if (fi78 >= (fi78 << ((259) & 63))) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v82 uint8 = 0
                        v83 := GenS7{F0: 0, F1: (v70 ^ (v70 | v70))}
                        _ = v82
                        _ = v83
                    } else {
                        if (!(4294967295 <= 42)) {
                            var v84 uint8 = 170
                            switch "hello" {
                            case "\n\t":
                                var v85 int64 = (-(v70))
                                _ = v85
                            }
                            var v86 int64 = -794
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v87 := GenS8{F0: 100, F1: -1}
                            v84 -= v84
                            _ = v84
                            _ = v86
                            _ = v87
                        }
                        var v88 int64 = (^(v70))
                        _ = v88
                    }
                    if (((("" + v1) + "a") + ("" + (v1 + v1))) < v1) {
                        if true {
                            switch int(v70) {
                            case -2147483648:
                                var v89 int16 = -32768
                                var v90 int16 = -760
                                v5 += 127
                                v5 = 1
                                _ = v89
                                _ = v90
                            }
                            fmt.Println(fi78, fi78)
                            fmt.Println(v5)
                        } else {
                            var v91 int = (fi79 % ((fi78) | 1))
                            v5 += v91
                            _ = v91
                        }
                        switch (("a" + "foo bar") + v1) {
                        case "":
                            var v92 bool = (!false)
                            _ = v92
                        case "foo bar":
                            switch (fi79 | -2147483648) {
                            case 0:
                                if (248 != 255) {
                                    v27 = (fi78 * (v5 + -1))
                                    v5 -= -429
                                    var v93 uint = 65535
                                    _ = v93
                                }
                            case 60:
                                if false {
                                    v94 := GenS7{F0: 42, F1: v70}
                                    v27 = fi79
                                    switch (((v1 + v1) + v1) + v1) {
                                    case "":
                                        var v95 uint32 = 0
                                        _ = v95
                                    case "foo bar":
                                        if (false && true) {
                                            v5 += 9223372036854775807
                                        } else {
                                            v70 = (v70 << ((0) & 63))
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                    case "hello":
                                        if ((float64(3.14159) < v2) || (!(!(366 == -32768)))) {
                                            var v96 int16 = -295
                                            fmt.Println(v27, v5)
                                            switch 32767 {
                                            case 1:
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            case -9223372036854775808:
                                                if (((-1 <= v96) || (!(!false))) && false) {
                                                    fmt.Println(fi78, fi78, v27)
                                                    var v97 float64 = 1.5
                                                    _ = v97
                                                }
                                            }
                                            _ = v96
                                        } else {
                                            v98 := GenS6{F0: 0, F1: 9223372036854775807, F2: (-(v70))}
                                            var v99 bool = (true || true)
                                            if v99 {
                                                _ = len(v4)
                                            } else {
                                                v27 -= (v5 ^ (v5 | fi78))
                                                if v99 {
                                                    _ = len(v4)
                                                    var v100 int64 = v70
                                                    var v101 string = v1
                                                    _ = v100
                                                    _ = v101
                                                } else {
                                                    switch (v1 + (v1 + v1)) {
                                                    case "a":
                                                        var v102 uint64 = 738
                                                        _ = v102
                                                    case "\n\t":
                                                        var v103 uint8 = 156
                                                        _ = v103
                                                    case "":
                                                        var v104 float64 = float64(-2.25)
                                                        _ = v104
                                                    default:
                                                        v27 = 42
                                                    }
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    var v105 uint16 = 255
                                                    _ = v105
                                                }
                                            }
                                            _ = v98
                                            _ = v99
                                        }
                                    }
                                    _ = v94
                                } else {
                                    if (true || ((2 < genFn0(42)) || (false && (!false)))) {
                                        v27 -= v27
                                        switch v1 {
                                        case "foo bar":
                                            var v106 int32 = -573
                                            var v107 uint8 = 49
                                            switch "a" {
                                            case "a":
                                                var v108 int32 = (v106 & (v106 / ((v106) | 1)))
                                                _ = v108
                                            case "foo bar":
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                var v109 int8 = 127
                                                if false {
                                                    v109 += (v109 * v109)
                                                    v110 := GenS4{F0: true}
                                                    _ = v110
                                                }
                                                _ = v109
                                            case "\n\t":
                                                var v111 int16 = 42
                                                _ = v111
                                            default:
                                                if false {
                                                    var v112 int16 = 1000
                                                    v1 = (((("hello" + "a") + "\n\t") + ("" + v1)) + "")
                                                    fmt.Println(v106, v2)
                                                    _ = v112
                                                }
                                            }
                                            _ = v106
                                            _ = v107
                                        case "\n\t":
                                            var v113 string = v1
                                            _ = v113
                                        case "a":
                                            v3 = append(v3, (v1 + v1))
                                        default:
                                            _ = len(v3)
                                        }
                                        v70 = -139
                                        v70 -= v70
                                    } else {
                                        fmt.Println(fi78, v5, v1)
                                        _ = len(v4)
                                    }
                                }
                            default:
                                v27 = fi78
                            }
                        case "\n\t":
                            fmt.Println(v5)
                        default:
                            fmt.Println(v27, v1, v70)
                        }
                        _ = len(v3)
                        v70 += v70
                        if true {
                            fmt.Println(fi79, v1, v5)
                            if (65535 < genFn0(418)) {
                                v2 = v2
                                if len(v4) > 0 {
                                    v4[0] = v1
                                }
                            } else {
                                v114 := GenS9{F0: -124, F1: 971, F2: 1}
                                _ = v114
                            }
                            v115 := GenS4{F0: (!(0 > 0))}
                            _ = v115
                        }
                    }
                    _ = fi79
                }
                if false {
                    if len(v4) > 0 {
                        v4[0] = "hello"
                    }
                    var v116 int16 = 0
                    for fi117 := 0; fi117 < 1; fi117++ {
                        var v118 uint = 255
                        if (!(v1 >= "foo bar")) {
                            var v119 int8 = 127
                            switch 127 {
                            case 32768:
                                v27 = -421
                            }
                            v119 = -1
                            _ = v119
                        } else {
                            v5 = (fi78 + (fi78 - (v5 << (((^(v118))) & 63))))
                            if false {
                                v70 += v70
                                var v120 int = v5
                                _ = v120
                            } else {
                                fmt.Println(v118)
                            }
                            v2 -= 1.5
                            v121 := GenS10{F0: v118, F1: v116}
                            v122 := GenS11{F0: 42, F1: 2147483647, F2: int8(v70)}
                            var v123 int8 = int8(v70)
                            if (("hello" == "") || false) {
                                if ((285 != 1) || true) {
                                    fmt.Println(fi117)
                                } else {
                                    v5 = (v27 + v27)
                                }
                                v124 := GenS11{F0: 255, F1: uint32(v118), F2: 42}
                                if true {
                                    v116 = v116
                                    if true {
                                        var v125 uint8 = 127
                                        v2 = v2
                                        v4 = append(v4, "\n\t")
                                        _ = v125
                                    }
                                    if (!false) {
                                        fmt.Println(v1, v116)
                                        if (((v118 <= v118) && (0 == 42)) && (((v118 ^ (^(v118))) < (v118 + v118)) || true)) {
                                            if (true && false) {
                                                var v126 int = (fi117 % ((v27) | 1))
                                                var v127 int8 = (v123 / ((v123) | 1))
                                                _ = v126
                                                _ = v127
                                            } else {
                                                var v128 string = "hello"
                                                _ = v128
                                            }
                                        } else {
                                            v2 = 1.5
                                            var v129 float32 = (3.14159 + 1e10)
                                            v2 = v2
                                            _ = v129
                                        }
                                    }
                                }
                                _ = v124
                            } else {
                                v130 := GenS7{F0: v70, F1: v70}
                                var v131 int = (v27 % ((v27) | 1))
                                _ = v130
                                _ = v131
                            }
                            v132 := GenS12{F0: 99}
                            v3 = append(v3, "a")
                            v3 = append(v3, (((v1 + v1) + ((v1 + v1) + "hello")) + (("foo bar" + "foo bar") + "foo bar")))
                            fmt.Println(v116, fi117, v5)
                            switch v1 {
                            case "hello":
                                var v133 int = -9223372036854775808
                                _ = v133
                            }
                            _ = v121
                            _ = v122
                            _ = v123
                            _ = v132
                        }
                        _ = fi117
                        _ = v118
                    }
                    _ = v116
                } else {
                    var v134 uint16 = 255
                    var v135 uint16 = v134
                    var v136 float64 = 1.5
                    _ = v134
                    _ = v135
                    _ = v136
                }
                _ = fi78
            }
        } else {
            v2 -= 0.1
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v5 += v5
        }
    case "\n\t":
        var v137 int64 = -821
        _ = v137
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v70
    _ = v76
}
