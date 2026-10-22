package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint; F1 string; F2 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 uint64; F2 int32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint8; F1 uint8; F2 string }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 string; F1 int32 }
type GenS13 struct { F0 int64; F1 string }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 bool; F1 uint8 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 string }
type GenS18 struct { F0 uint16 }
type GenS19 struct { F0 int32; F1 int32; F2 int16 }
type GenS20 struct { F0 int; F1 int32 }
func genFn0() int8 {
    var v4 uint16 = 539
    if false {
        for fi5 := 0; fi5 < 13; fi5++ {
            v4 = (v4 + v4)
            v6 := []int{}
            _ = fi5
            _ = v6
        }
        switch "" {
        case "":
            v4 -= 65535
        case "foo bar":
            var v7 bool = (true && true)
            _ = v7
        }
    } else {
        for fi8 := 0; fi8 < 2; fi8++ {
            for fi9 := 0; fi9 < 8; fi9++ {
                fmt.Println(fi9)
                v10 := []int{-2147483648}
                switch fi9 {
                case 32768:
                    v11 := GenS1{F0: 106, F1: 255}
                    _ = v11
                case 592:
                    v4 -= v4
                case 127:
                    _ = len(v10)
                }
                if false {
                    var v12 int8 = 0
                    var v13 float64 = ((0.1 + (((-2.25 / 0.1) - (((-2.25 * (3.14159 / -2.25)) * float64(0.1)) / 1.5)) * -2.25)) - 3.14159)
                    var v14 bool = (42 > 100)
                    _ = v12
                    _ = v13
                    _ = v14
                }
                _ = fi9
                _ = v10
            }
            if true {
                if (!(0 != 100)) {
                    fmt.Println(fi8)
                    v4 = (^(v4))
                }
                if (v4 >= 32767) {
                    for fi15 := 0; fi15 < 4; fi15++ {
                        v4 += v4
                        v16 := GenS2{F0: (1 == 127)}
                        v17 := GenS2{F0: true}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    if (127 >= -26) {
                        v4 = v4
                    }
                    switch ("foo bar" + "a") {
                    case "":
                        if (false && false) {
                            var v18 uint = 2
                            _ = v18
                        }
                    case "foo bar":
                        v19 := GenS1{F0: 2, F1: 97}
                        _ = v19
                    case "\n\t":
                        switch fi8 {
                        case 227:
                            v20 := []int{}
                            _ = v20
                        case 202:
                            for fi21 := 0; fi21 < 16; fi21++ {
                                v4 = v4
                                _ = fi21
                            }
                        default:
                            var v22 uint64 = 2
                            _ = v22
                        }
                    default:
                        for fi23 := 0; fi23 < 12; fi23++ {
                            switch fi23 {
                            case 32768:
                                var v24 uint = 549
                                _ = v24
                            case -998:
                                var v25 int8 = 100
                                _ = v25
                            case 128:
                                v4 += (v4 | (v4 * (v4 / ((334) | 1))))
                            }
                            fmt.Println(v4)
                            switch "foo bar" {
                            case "":
                                v4 -= (v4 & 350)
                            default:
                                var v26 int64 = -9223372036854775808
                                _ = v26
                            }
                            _ = fi23
                        }
                    }
                } else {
                    v27 := []int{fi8, -9223372036854775808, 0, fi8}
                    for fi28 := 0; fi28 < 15; fi28++ {
                        fmt.Println(fi28, v4, fi8)
                        switch "hello" {
                        case "hello":
                            switch ("a" + "foo bar") {
                            case "\n\t":
                                v29 := GenS2{F0: false}
                                _ = v29
                            case "a":
                                v4 = (v4 % (((^(v4))) | 1))
                            }
                        case "a":
                            v4 = v4
                        case "\n\t":
                            var v30 uint32 = 332
                            _ = v30
                        default:
                            var v31 uint16 = v4
                            _ = v31
                        }
                        var v32 int = (fi28 + fi8)
                        _ = fi28
                        _ = v32
                    }
                    v4 += (v4 - (v4 % ((0) | 1)))
                    _ = v27
                }
                fmt.Println(fi8, v4)
            }
            v4 = (v4 + v4)
            _ = fi8
        }
    }
    _ = v4
    return 0
}
func genFn1(p0 int16) uint16 {
    for fi33 := 0; fi33 < 7; fi33++ {
        switch (fi33 + 128) {
        case 9223372036854775807:
            fmt.Println(p0, p0)
        }
        v34 := GenS1{F0: 127, F1: 128}
        _ = fi33
        _ = v34
    }
    v35 := GenS1{F0: 0, F1: 60}
    switch "" {
    case "foo bar":
        for fi36 := 0; fi36 < 17; fi36++ {
            var v37 int = fi36
            var v38 int = (fi36 * (v37 ^ (fi36 << ((567) & 63))))
            p0 = (p0 * (p0 << ((2) & 63)))
            _ = fi36
            _ = v37
            _ = v38
        }
    case "":
        var v39 uint64 = 42
        _ = v39
    default:
        p0 = (p0 | p0)
    }
    p0 += 42
    _ = p0
    _ = v35
    return 65535
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{}
    var v3 uint = 2
    for fi40 := 0; fi40 < 13; fi40++ {
        v41 := GenS3{F0: 1, F1: 234, F2: 2}
        for fi42 := 0; fi42 < 20; fi42++ {
            fmt.Println(fi40)
            _ = fi42
        }
        v43 := GenS3{F0: 463, F1: 255, F2: 4294967295}
        _ = fi40
        _ = v41
        _ = v43
    }
    var v44 int32 = -327
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ("" == "foo bar") {
        v45 := GenS3{F0: (^(v44)), F1: 2, F2: uint64(v3)}
        for fi46 := 0; fi46 < 9; fi46++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(fi46, v3)
            _ = fi46
        }
        var v47 uint = 2
        _ = v45
        _ = v47
    }
    for fi48 := 0; fi48 < 5; fi48++ {
        v3 += v3
        _ = fi48
    }
    fmt.Println(v44)
    v3 += v3
    for fi49 := 0; fi49 < 1; fi49++ {
        _ = len(v2)
        if true {
            for fi50 := 0; fi50 < 13; fi50++ {
                if (100 != 100) {
                    _ = len(v1)
                    v51 := GenS3{F0: -2147483648, F1: 32767, F2: 0}
                    v52 := GenS1{F0: 2, F1: 218}
                    _ = v51
                    _ = v52
                } else {
                    v53 := GenS3{F0: (v44 - (v44 & v44)), F1: 265, F2: 709}
                    if true {
                        if (-128 == genFn0()) {
                            if len(v2) > 0 {
                                v2[0] = "\n\t"
                            }
                            v3 += (v3 - (v3 + v3))
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v2 = append(v2, "foo bar")
                        }
                    }
                    _ = v53
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi50
            }
            if false {
                for fi54 := 0; fi54 < 11; fi54++ {
                    var v55 float32 = ((0.1 + float32(1.5)) * -2.25)
                    _ = fi54
                    _ = v55
                }
                for fi56 := 0; fi56 < 9; fi56++ {
                    var v57 int = (-(fi56))
                    var v58 int = fi49
                    _ = len(v2)
                    _ = fi56
                    _ = v57
                    _ = v58
                }
            } else {
                fmt.Println(v44, fi49, v44)
                v3 = v3
            }
            delete(v1, fi49)
        } else {
            _ = len(v1)
            if (-661 == -778) {
                if false {
                    v3 = 255
                    for fi59 := 0; fi59 < 16; fi59++ {
                        var v60 uint64 = 42
                        _ = fi59
                        _ = v60
                    }
                } else {
                    v44 = v44
                    v44 += -1
                }
                v61 := GenS4{F0: 18446744073709551615, F1: "hello", F2: (-(v44))}
                if false {
                    if len(v2) > 0 {
                        v2[0] = (("foo bar" + "\n\t") + "")
                    }
                }
                _ = v61
            }
        }
        _ = fi49
    }
    fmt.Println(v44, v3)
    for fi62 := 0; fi62 < 13; fi62++ {
        var v63 float64 = 3.14159
        if (((-1 > 99) && ((false || true) || false)) || ("hello" <= "a")) {
            for fi64 := 0; fi64 < 2; fi64++ {
                var v65 int64 = 9223372036854775807
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                _ = fi64
                _ = v65
            }
            delete(v1, (fi62 | fi62))
        } else {
            var v66 uint8 = 1
            if ("a" > "foo bar") {
                if false {
                    v67 := GenS5{F0: 119}
                    if (255 > 472) {
                        if false {
                            v68 := GenS3{F0: (-(v44)), F1: 568, F2: 18446744073709551615}
                            _ = v68
                        } else {
                            fmt.Println(v3, fi62, v63)
                            v69 := GenS3{F0: 2147483647, F1: 255, F2: 4294967295}
                            var v70 int64 = 1
                            _ = v69
                            _ = v70
                        }
                        v71 := GenS6{F0: 29, F1: 79, F2: v44}
                        v72 := GenS7{F0: true}
                        _ = v71
                        _ = v72
                    } else {
                        fmt.Println(v66)
                    }
                    _ = v67
                }
                var v73 float32 = 3.14159
                for fi74 := 0; fi74 < 16; fi74++ {
                    var v75 int32 = v44
                    switch "a" {
                    case "\n\t":
                        _ = len(v2)
                    case "a":
                        var v76 int16 = 42
                        _ = v76
                    case "hello":
                        v77 := v1[(fi74 << ((v3) & 63))]
                        _ = v77
                    default:
                        fmt.Println(v44)
                    }
                    var v78 uint64 = 300
                    _ = fi74
                    _ = v75
                    _ = v78
                }
                _ = v73
            }
            _ = v66
        }
        for fi79 := 0; fi79 < 6; fi79++ {
            switch ("foo bar" + "\n\t") {
            case "foo bar":
                if false {
                    var v80 uint = (v3 ^ (v3 % ((2) | 1)))
                    fmt.Println(v44)
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                    _ = v80
                }
            case "":
                v63 -= v63
            case "\n\t":
                var v81 int32 = (v44 & 211)
                _ = v81
            default:
                fmt.Println(v44, v3, fi62)
            }
            switch 520 {
            case 128:
                if (135 > 1) {
                    v82 := GenS3{F0: 0, F1: 42, F2: 1}
                    fmt.Println(v44, v3, v44)
                    _ = v82
                } else {
                    v83 := GenS4{F0: v3, F1: ("" + ("a" + "hello")), F2: (-(v44))}
                    var v84 uint64 = 18446744073709551615
                    if (((v63 <= float64(1.5)) || true) || (!((v3 << ((v3) & 63)) <= uint(v84)))) {
                        if (127 > 0) {
                            if true {
                                if ((true || (false && true)) && (65535 == 370)) {
                                    var v85 int32 = (v44 & (v44 * (v44 ^ v44)))
                                    if (-32768 > 1) {
                                        if (128 > 107) {
                                            switch (fi79 ^ fi79) {
                                            case -390:
                                                v86 := GenS4{F0: v3, F1: ("a" + "a"), F2: (v44 % (((-(v44))) | 1))}
                                                _ = v86
                                            case -2147483648:
                                                fmt.Println(fi79)
                                            case 32768:
                                                var v87 int32 = 1
                                                _ = v87
                                            }
                                            switch (fi79 % ((fi79) | 1)) {
                                            case -9223372036854775808:
                                                v85 = (v85 + -840)
                                            case 1:
                                                v3 = 182
                                            default:
                                                v88 := v1[fi62]
                                                _ = v88
                                            }
                                        } else {
                                            switch fi62 {
                                            case -2147483648:
                                                var v89 string = "\n\t"
                                                _ = v89
                                            case 32767:
                                                v90 := GenS8{F0: 731}
                                                _ = v90
                                            case 0:
                                                fmt.Println(fi79)
                                            }
                                            if false {
                                                if (!((v44 + v85) != (v85 / ((v44) | 1)))) {
                                                    v91 := GenS6{F0: 100, F1: v84, F2: (-(v44))}
                                                    var v92 string = ""
                                                    v84 -= v84
                                                    _ = v91
                                                    _ = v92
                                                }
                                                if true {
                                                    if true {
                                                        if true {
                                                            if len(v2) > 0 {
                                                                v2[0] = "hello"
                                                            }
                                                            v63 = 1e10
                                                        }
                                                    } else {
                                                        var v93 uint8 = 209
                                                        v1[-2147483648] = (fi79 - fi79)
                                                        var v94 int = -118
                                                        _ = v93
                                                        _ = v94
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v3 = v3
                                                }
                                            } else {
                                                switch fi79 {
                                                case 1:
                                                    if (true || ("" > "foo bar")) {
                                                        v44 = (v44 * 2147483647)
                                                        v95 := GenS3{F0: (v85 + v85), F1: 42, F2: (v84 - 42)}
                                                        _ = len(v2)
                                                        _ = v95
                                                    } else {
                                                        v96 := GenS5{F0: -63}
                                                        fmt.Println(v3, v63)
                                                        _ = len(v1)
                                                        _ = v96
                                                    }
                                                case 32767:
                                                    switch (("\n\t" + "\n\t") + ("" + ("\n\t" + "foo bar"))) {
                                                    case "a":
                                                        if (v63 <= ((0.0 - -2.25) * 0.1)) {
                                                            v44 = (v85 ^ (v44 + (-(v85))))
                                                            if (float32((0.1 - -2.25)) < (1e10 * (1e10 + -2.25))) {
                                                                fmt.Println(v63, fi62)
                                                                v97 := GenS3{F0: 580, F1: 524, F2: v84}
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    v98 := GenS9{F0: 1, F1: 128, F2: "foo bar"}
                                                                    _ = v98
                                                                case "":
                                                                    v84 = (v84 - v84)
                                                                case "foo bar":
                                                                    switch fi79 {
                                                                    case 301:
                                                                        if true {
                                                                            if true {
                                                                                v99 := v1[9223372036854775807]
                                                                                _ = v99
                                                                            } else {
                                                                                var v100 int16 = 32767
                                                                                v63 += v63
                                                                                if (true || (v3 <= (v3 | 255))) {
                                                                                    var v101 float64 = ((float64(float32(float64(float32(v63)))) + v63) - ((v63 * 3.14159) + float64(0.0)))
                                                                                    _ = v101
                                                                                } else {
                                                                                    v3 -= v3
                                                                                }
                                                                                _ = v100
                                                                            }
                                                                            if (v63 != v63) {
                                                                                if false {
                                                                                    if (!(false || (v85 > (-(v85))))) {
                                                                                        v102 := GenS8{F0: -466}
                                                                                        var v103 uint64 = v84
                                                                                        if false {
                                                                                            if false {
                                                                                                v3 += (^(v3))
                                                                                            }
                                                                                        }
                                                                                        _ = v102
                                                                                        _ = v103
                                                                                    }
                                                                                    if true {
                                                                                        v1[(fi79 << (((v3 * (^(v3)))) & 63))] = (fi79 * (fi79 >> ((v3) & 63)))
                                                                                        v104 := GenS10{F0: 0}
                                                                                        _ = v104
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                switch (fi62 << ((v3) & 63)) {
                                                                                case 9223372036854775807:
                                                                                    switch (("hello" + "hello") + ("\n\t" + "")) {
                                                                                    case "\n\t":
                                                                                        var v105 uint16 = 708
                                                                                        _ = v105
                                                                                    case "hello":
                                                                                        var v106 bool = ((v63 * v63) < (v63 * ((-2.25 - 1e10) * v63)))
                                                                                        _ = v106
                                                                                    }
                                                                                case 793:
                                                                                    var v107 bool = (!true)
                                                                                    _ = v107
                                                                                }
                                                                            }
                                                                        }
                                                                    case 32767:
                                                                        switch fi62 {
                                                                        case 4294967295:
                                                                            v84 += v84
                                                                        case 2147483647:
                                                                            switch (fi79 ^ 55) {
                                                                            case 128:
                                                                                switch "foo bar" {
                                                                                case "":
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                case "foo bar":
                                                                                    var v108 int32 = (v44 - v44)
                                                                                    _ = v108
                                                                                case "hello":
                                                                                    fmt.Println(v3, v63, v44)
                                                                                    var v109 float64 = v63
                                                                                    _ = v109
                                                                                }
                                                                            }
                                                                        case 169:
                                                                            if false {
                                                                                switch "a" {
                                                                                case "a":
                                                                                    if true {
                                                                                        var v110 int16 = -32768
                                                                                        v85 = 1
                                                                                        switch "hello" {
                                                                                        case "hello":
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                        default:
                                                                                            fmt.Println(v110)
                                                                                        }
                                                                                        _ = v110
                                                                                    }
                                                                                case "hello":
                                                                                    switch 32767 {
                                                                                    case 127:
                                                                                        switch (fi79 << (((v3 | (v3 % ((65535) | 1)))) & 63)) {
                                                                                        case 245:
                                                                                            var v111 uint8 = 2
                                                                                            if (true || (((!true) || (("\n\t" < ("hello" + (("a" + "foo bar") + ("foo bar" + (("a" + "") + ("\n\t" + "")))))) && (("\n\t" <= "") && false))) || (("hello" != "a") && (!(true || true))))) {
                                                                                                if false {
                                                                                                    if (false && false) {
                                                                                                        switch (("hello" + (("hello" + "a") + "\n\t")) + "hello") {
                                                                                                        case "hello":
                                                                                                            if (!(2 > 0)) {
                                                                                                                v112 := GenS9{F0: (v111 << (((v3 ^ (^(v3)))) & 63)), F1: (^(v111)), F2: "\n\t"}
                                                                                                                v2 = append(v2, ("foo bar" + "foo bar"))
                                                                                                                _ = v112
                                                                                                            }
                                                                                                        case "foo bar":
                                                                                                            var v113 int = 1
                                                                                                            _ = v113
                                                                                                        default:
                                                                                                            v44 = v85
                                                                                                        }
                                                                                                        var v114 int16 = 1
                                                                                                        if false {
                                                                                                            v1[(fi79 >> (((v3 + v3)) & 63))] = (fi79 / ((-811) | 1))
                                                                                                            if true {
                                                                                                                v115 := GenS8{F0: v114}
                                                                                                                v84 = 2
                                                                                                                _ = v115
                                                                                                            } else {
                                                                                                                switch "" {
                                                                                                                case "a":
                                                                                                                    if false {
                                                                                                                        v3 -= (v3 - v3)
                                                                                                                        var v116 uint8 = v111
                                                                                                                        if len(v2) > 0 {
                                                                                                                            v2[0] = "hello"
                                                                                                                        }
                                                                                                                        _ = v116
                                                                                                                    } else {
                                                                                                                        fmt.Println(v44)
                                                                                                                        v114 -= (v114 << ((714) & 63))
                                                                                                                        var v117 uint = 1
                                                                                                                        _ = v117
                                                                                                                    }
                                                                                                                case "foo bar":
                                                                                                                    v84 -= 418
                                                                                                                    if len(v2) >= 2 {
                                                                                                                        v2 = v2[1:len(v2)]
                                                                                                                    }
                                                                                                                case "":
                                                                                                                    v84 = 567
                                                                                                                default:
                                                                                                                    switch fi79 {
                                                                                                                    case -9223372036854775808:
                                                                                                                        v118 := GenS3{F0: v85, F1: 789, F2: (v84 ^ v84)}
                                                                                                                        _ = v118
                                                                                                                    case -536:
                                                                                                                        if true {
                                                                                                                            var v119 uint16 = 0
                                                                                                                            fmt.Println(v63)
                                                                                                                            var v120 int64 = 0
                                                                                                                            _ = v119
                                                                                                                            _ = v120
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        var v121 uint8 = 175
                                                                                                                        _ = v121
                                                                                                                    }
                                                                                                                }
                                                                                                                var v122 int16 = 42
                                                                                                                _ = v122
                                                                                                            }
                                                                                                            v44 += v44
                                                                                                        }
                                                                                                        _ = v114
                                                                                                    } else {
                                                                                                        if ((v84 - (v84 << ((v3) & 63))) != v84) {
                                                                                                            v3 = (v3 % ((v3) | 1))
                                                                                                            switch 32767 {
                                                                                                            case -2147483648:
                                                                                                                switch int(v44) {
                                                                                                                case 127:
                                                                                                                    v123 := GenS7{F0: false}
                                                                                                                    _ = v123
                                                                                                                }
                                                                                                                switch fi79 {
                                                                                                                case 32767:
                                                                                                                    var v124 string = "hello"
                                                                                                                    _ = v124
                                                                                                                case -1:
                                                                                                                    var v125 float64 = 3.14159
                                                                                                                    _ = v125
                                                                                                                case -9223372036854775808:
                                                                                                                    v126 := GenS6{F0: -128, F1: uint64(v3), F2: -2147483648}
                                                                                                                    _ = v126
                                                                                                                }
                                                                                                                if false {
                                                                                                                    fmt.Println(v84, v111, v111)
                                                                                                                    v127 := GenS2{F0: (!(false || true))}
                                                                                                                    if (0.0 != 1e10) {
                                                                                                                        v63 = (v63 - v63)
                                                                                                                        if (false && (false && false)) {
                                                                                                                            var v128 uint64 = v84
                                                                                                                            _ = v128
                                                                                                                        }
                                                                                                                        v85 += v85
                                                                                                                        _ = len(v2)
                                                                                                                    }
                                                                                                                    switch "a" {
                                                                                                                    case "\n\t":
                                                                                                                        switch fi79 {
                                                                                                                        case 2147483647:
                                                                                                                            v111 = v111
                                                                                                                        case -1:
                                                                                                                            switch (-(fi79)) {
                                                                                                                            case -1:
                                                                                                                                v129 := v1[fi79]
                                                                                                                                _ = v129
                                                                                                                            case -942:
                                                                                                                                var v130 float64 = v63
                                                                                                                                _ = v130
                                                                                                                            default:
                                                                                                                                _ = len(v2)
                                                                                                                            }
                                                                                                                        case 369:
                                                                                                                            if true {
                                                                                                                                var v131 bool = false
                                                                                                                                fmt.Println(fi79, v3, v111)
                                                                                                                                _ = v131
                                                                                                                            }
                                                                                                                        }
                                                                                                                    case "foo bar":
                                                                                                                        var v132 uint8 = v111
                                                                                                                        _ = v132
                                                                                                                    case "":
                                                                                                                        v3 -= 42
                                                                                                                    }
                                                                                                                    fmt.Println(v63)
                                                                                                                    _ = v127
                                                                                                                }
                                                                                                            case 0:
                                                                                                                if ((903 <= -247) && false) {
                                                                                                                    var v133 uint16 = 530
                                                                                                                    var v134 bool = true
                                                                                                                    _ = v133
                                                                                                                    _ = v134
                                                                                                                }
                                                                                                            }
                                                                                                        } else {
                                                                                                            v135 := GenS11{F0: uint32(v111)}
                                                                                                            switch "hello" {
                                                                                                            case "":
                                                                                                                if (2 == 0) {
                                                                                                                    v136 := v1[(^(fi62))]
                                                                                                                    var v137 int64 = -466
                                                                                                                    v85 += (v44 % ((454) | 1))
                                                                                                                    _ = v136
                                                                                                                    _ = v137
                                                                                                                } else {
                                                                                                                    if len(v2) > 0 {
                                                                                                                        v2[0] = ""
                                                                                                                    }
                                                                                                                }
                                                                                                            default:
                                                                                                                _ = len(v1)
                                                                                                                if false {
                                                                                                                    fmt.Println(fi62, v3, v63)
                                                                                                                    var v138 float64 = 0.1
                                                                                                                    switch ("" + ("" + "")) {
                                                                                                                    case "\n\t":
                                                                                                                        v3 += v3
                                                                                                                    case "hello":
                                                                                                                        var v139 int8 = -102
                                                                                                                        _ = v139
                                                                                                                    case "":
                                                                                                                        switch fi79 {
                                                                                                                        case 32768:
                                                                                                                            v140 := GenS6{F0: 127, F1: 2, F2: (v85 | (-(v85)))}
                                                                                                                            _ = v140
                                                                                                                        case 688:
                                                                                                                            v84 -= (v84 >> (((v3 & (v3 / ((v3) | 1)))) & 63))
                                                                                                                        case 4294967295:
                                                                                                                            fmt.Println(fi79, fi62)
                                                                                                                        default:
                                                                                                                            if true {
                                                                                                                                _ = len(v1)
                                                                                                                                switch "hello" {
                                                                                                                                case "":
                                                                                                                                    v141 := v1[fi62]
                                                                                                                                    _ = v141
                                                                                                                                case "\n\t":
                                                                                                                                    var v142 bool = true
                                                                                                                                    _ = v142
                                                                                                                                case "hello":
                                                                                                                                    var v143 int32 = v85
                                                                                                                                    _ = v143
                                                                                                                                default:
                                                                                                                                    var v144 float64 = (v63 + v138)
                                                                                                                                    _ = v144
                                                                                                                                }
                                                                                                                                v145 := GenS8{F0: 1000}
                                                                                                                                _ = v145
                                                                                                                            } else {
                                                                                                                                var v146 string = "a"
                                                                                                                                _ = v146
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v147 := GenS1{F0: 131, F1: (v111 << ((v3) & 63))}
                                                                                                                        switch 42 {
                                                                                                                        case -598:
                                                                                                                            var v148 int = int(v84)
                                                                                                                            _ = v148
                                                                                                                        case -9223372036854775808:
                                                                                                                            delete(v1, 32767)
                                                                                                                        case -884:
                                                                                                                            var v149 int64 = 890
                                                                                                                            _ = v149
                                                                                                                        }
                                                                                                                        _ = v147
                                                                                                                    }
                                                                                                                    v111 = v111
                                                                                                                    _ = v138
                                                                                                                }
                                                                                                            }
                                                                                                            var v150 uint16 = 42
                                                                                                            _ = v135
                                                                                                            _ = v150
                                                                                                        }
                                                                                                    }
                                                                                                    var v151 uint64 = v84
                                                                                                    delete(v1, 32767)
                                                                                                    fmt.Println(v151, v3, v63)
                                                                                                    if true {
                                                                                                        switch "\n\t" {
                                                                                                        case "hello":
                                                                                                            fmt.Println(v85)
                                                                                                        case "a":
                                                                                                            if false {
                                                                                                                if true {
                                                                                                                    if (42 <= 576) {
                                                                                                                        switch (fi62 * int(v3)) {
                                                                                                                        case -412:
                                                                                                                            var v152 string = ("hello" + ("hello" + "a"))
                                                                                                                            var v153 uint64 = (v84 | v84)
                                                                                                                            _ = v152
                                                                                                                            _ = v153
                                                                                                                        case 945:
                                                                                                                            if (!false) {
                                                                                                                                _ = len(v1)
                                                                                                                                v151 = v84
                                                                                                                            }
                                                                                                                        case 128:
                                                                                                                            fmt.Println(fi62)
                                                                                                                            v85 = 1
                                                                                                                        }
                                                                                                                        v154 := GenS7{F0: ((false && false) || true)}
                                                                                                                        var v155 uint8 = v111
                                                                                                                        _ = v154
                                                                                                                        _ = v155
                                                                                                                    }
                                                                                                                    fmt.Println(fi79, v111, v63)
                                                                                                                    v85 = (-(v44))
                                                                                                                }
                                                                                                                v1[(fi62 & fi79)] = (fi62 | fi79)
                                                                                                                v1[-9223372036854775808] = 4294967295
                                                                                                                var v156 int32 = 42
                                                                                                                v3 += 1
                                                                                                                _ = v156
                                                                                                            } else {
                                                                                                                v151 = (v151 ^ 2)
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    v157 := GenS3{F0: v44, F1: 2, F2: (^(v84))}
                                                                                                    _ = v151
                                                                                                    _ = v157
                                                                                                }
                                                                                                v158 := GenS6{F0: 103, F1: (v84 ^ 447), F2: v85}
                                                                                                _ = v158
                                                                                            } else {
                                                                                                var v159 int16 = 32767
                                                                                                v84 -= 42
                                                                                                _ = len(v1)
                                                                                                _ = v159
                                                                                            }
                                                                                            v63 += (0.1 * -2.25)
                                                                                            _ = v111
                                                                                        }
                                                                                    }
                                                                                case "foo bar":
                                                                                    var v160 uint64 = v84
                                                                                    _ = v160
                                                                                }
                                                                                v84 = v84
                                                                            }
                                                                        default:
                                                                            var v161 int64 = 1
                                                                            var v162 uint16 = 2
                                                                            v161 = (v161 * v161)
                                                                            _ = v161
                                                                            _ = v162
                                                                        }
                                                                    default:
                                                                        if (113 <= 141) {
                                                                            v84 = (^(v84))
                                                                        }
                                                                    }
                                                                }
                                                                var v163 uint64 = (v84 + 4294967295)
                                                                v63 = 1.5
                                                                _ = v163
                                                                _ = v97
                                                            } else {
                                                                v3 += uint(v84)
                                                                v63 = v63
                                                                var v164 uint16 = 32767
                                                                _ = v164
                                                            }
                                                        } else {
                                                            v165 := GenS5{F0: -1}
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            switch fi62 {
                                                            case 4294967295:
                                                                if (((float64(0.0) + 1.5) + v63) != (v63 / 1.5)) {
                                                                    v166 := GenS5{F0: int8(v3)}
                                                                    _ = v166
                                                                } else {
                                                                    var v167 uint16 = genFn1(1)
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        var v168 bool = (true && ((v85 >> ((255) & 63)) > v44))
                                                                        _ = v168
                                                                    case "":
                                                                        if (float32(((0.0 - 1e10) - float64(0.1))) <= 3.14159) {
                                                                            var v169 bool = ((fi79 % ((0) | 1)) < (fi79 - fi62))
                                                                            delete(v1, (fi79 << (((v3 % ((v3) | 1))) & 63)))
                                                                            switch fi62 {
                                                                            case 128:
                                                                                var v170 uint = 0
                                                                                _ = v170
                                                                            case 2147483647:
                                                                                v169 = v169
                                                                            case -9223372036854775808:
                                                                                switch (fi79 >> (((v3 & (^(v3)))) & 63)) {
                                                                                case -10:
                                                                                    v85 += (v44 / ((0) | 1))
                                                                                case 2147483647:
                                                                                    switch fi62 {
                                                                                    case -49:
                                                                                        var v171 string = "a"
                                                                                        _ = v171
                                                                                    default:
                                                                                        if true {
                                                                                            if (-128 > genFn0()) {
                                                                                                v44 = v85
                                                                                            } else {
                                                                                                var v172 uint16 = (v167 % (((^(v167))) | 1))
                                                                                                v173 := GenS3{F0: v44, F1: (^(v167)), F2: (^(v84))}
                                                                                                var v174 int = (fi62 % ((447) | 1))
                                                                                                _ = v172
                                                                                                _ = v173
                                                                                                _ = v174
                                                                                            }
                                                                                            switch "\n\t" {
                                                                                            case "a":
                                                                                                fmt.Println(v44, v44, v167)
                                                                                            case "\n\t":
                                                                                                var v175 bool = false
                                                                                                _ = v175
                                                                                            case "":
                                                                                                if ((v167 + (v167 - 896)) > v167) {
                                                                                                    v176 := GenS12{F0: ("hello" + "\n\t"), F1: v85}
                                                                                                    var v177 uint16 = v167
                                                                                                    switch "" {
                                                                                                    case "":
                                                                                                        if (v169 && ("\n\t" > "foo bar")) {
                                                                                                            var v178 string = "hello"
                                                                                                            fmt.Println(v84, v178, v167)
                                                                                                            _ = v178
                                                                                                        }
                                                                                                    case "\n\t":
                                                                                                        delete(v1, fi79)
                                                                                                    case "foo bar":
                                                                                                        if v169 {
                                                                                                            if v169 {
                                                                                                                var v179 uint8 = 127
                                                                                                                v44 -= 42
                                                                                                                _ = v179
                                                                                                            }
                                                                                                            fmt.Println(v3)
                                                                                                        } else {
                                                                                                            if v169 {
                                                                                                                fmt.Println(v3, v177)
                                                                                                                _ = len(v2)
                                                                                                                var v180 int64 = -9223372036854775808
                                                                                                                _ = v180
                                                                                                            } else {
                                                                                                                fmt.Println(v169, v85, v3)
                                                                                                                switch "" {
                                                                                                                case "":
                                                                                                                    v177 = (v177 - (v177 ^ v177))
                                                                                                                case "a":
                                                                                                                    v181 := GenS8{F0: 11}
                                                                                                                    _ = v181
                                                                                                                case "hello":
                                                                                                                    var v182 int16 = -194
                                                                                                                    _ = v182
                                                                                                                }
                                                                                                            }
                                                                                                            fmt.Println(v169, v85)
                                                                                                            switch (fi79 << ((v3) & 63)) {
                                                                                                            case -2147483648:
                                                                                                                if v169 {
                                                                                                                    var v183 uint = 2
                                                                                                                    if len(v2) > 0 {
                                                                                                                        v2[0] = (("hello" + "") + (((("foo bar" + "") + "a") + "a") + ("\n\t" + "")))
                                                                                                                    }
                                                                                                                    _ = v183
                                                                                                                } else {
                                                                                                                    switch "" {
                                                                                                                    case "a":
                                                                                                                        if v169 {
                                                                                                                            v184 := GenS11{F0: 569}
                                                                                                                            if (771 > 2147483647) {
                                                                                                                                switch (fi79 * fi62) {
                                                                                                                                case -2147483648:
                                                                                                                                    v167 = v177
                                                                                                                                case 4294967295:
                                                                                                                                    var v185 int8 = 100
                                                                                                                                    _ = v185
                                                                                                                                }
                                                                                                                                v44 += (v85 ^ (v44 + v85))
                                                                                                                            } else {
                                                                                                                                v186 := GenS11{F0: uint32(v3)}
                                                                                                                                _ = v186
                                                                                                                            }
                                                                                                                            var v187 uint64 = uint64(v177)
                                                                                                                            delete(v1, (fi62 | (fi62 << ((1) & 63))))
                                                                                                                            _ = v184
                                                                                                                            _ = v187
                                                                                                                        } else {
                                                                                                                            if (!(2147483647 < 2)) {
                                                                                                                                v2 = append(v2, "a")
                                                                                                                                if v169 {
                                                                                                                                    if (0.0 == 1.5) {
                                                                                                                                        switch "foo bar" {
                                                                                                                                        case "hello":
                                                                                                                                            switch (fi79 - (fi79 | fi62)) {
                                                                                                                                            case 1:
                                                                                                                                                fmt.Println(fi79, v169, fi79)
                                                                                                                                            case 670:
                                                                                                                                                v177 += v167
                                                                                                                                            default:
                                                                                                                                                v188 := GenS13{F0: 42, F1: ""}
                                                                                                                                                _ = v188
                                                                                                                                            }
                                                                                                                                        default:
                                                                                                                                            switch "a" {
                                                                                                                                            case "foo bar":
                                                                                                                                                v189 := GenS1{F0: 244, F1: uint8(v84)}
                                                                                                                                                _ = v189
                                                                                                                                            default:
                                                                                                                                                var v190 int8 = 42
                                                                                                                                                _ = v190
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    if false {
                                                                                                                                        v191 := GenS9{F0: 1, F1: 128, F2: ((("\n\t" + "") + "") + "foo bar")}
                                                                                                                                        v192 := GenS14{F0: -1}
                                                                                                                                        _ = v191
                                                                                                                                        _ = v192
                                                                                                                                    } else {
                                                                                                                                        if (!(!(!v169))) {
                                                                                                                                            var v193 uint8 = 2
                                                                                                                                            v63 = v63
                                                                                                                                            _ = len(v2)
                                                                                                                                            var v194 uint64 = (^(v84))
                                                                                                                                            _ = v193
                                                                                                                                            _ = v194
                                                                                                                                        } else {
                                                                                                                                            var v195 uint32 = uint32(v3)
                                                                                                                                            _ = v195
                                                                                                                                        }
                                                                                                                                        _ = len(v1)
                                                                                                                                        v196 := GenS15{F0: v169, F1: 128}
                                                                                                                                        _ = v196
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                _ = len(v1)
                                                                                                                            }
                                                                                                                            if (fi79 > fi79) {
                                                                                                                                fmt.Println(v167)
                                                                                                                                v167 = (^(v167))
                                                                                                                                switch "a" {
                                                                                                                                case "a":
                                                                                                                                    _ = len(v1)
                                                                                                                                case "hello":
                                                                                                                                    v84 -= 0
                                                                                                                                default:
                                                                                                                                    v177 += v177
                                                                                                                                }
                                                                                                                                if ((v169 && (genFn1(0) != v167)) && (false || (!(fi79 == (-(fi79)))))) {
                                                                                                                                    fmt.Println(v167)
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                var v197 uint64 = (^(v84))
                                                                                                                                _ = v197
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                    if (fi79 <= (fi62 / ((fi62) | 1))) {
                                                                                                                        var v198 uint = 772
                                                                                                                        v177 += genFn1(1000)
                                                                                                                        _ = v198
                                                                                                                    } else {
                                                                                                                        v199 := GenS12{F0: ("a" + (("a" + "") + ((("hello" + ("\n\t" + ("\n\t" + "hello"))) + (("\n\t" + "") + "hello")) + "\n\t"))), F1: (v44 << ((743) & 63))}
                                                                                                                        if false {
                                                                                                                            v177 = genFn1(1)
                                                                                                                            v200 := GenS7{F0: v169}
                                                                                                                            if v169 {
                                                                                                                                fmt.Println(v167, v167, fi79)
                                                                                                                                v167 = v167
                                                                                                                            }
                                                                                                                            _ = v200
                                                                                                                        }
                                                                                                                        _ = v199
                                                                                                                    }
                                                                                                                }
                                                                                                            case -957:
                                                                                                                if v169 {
                                                                                                                    var v201 uint32 = 65535
                                                                                                                    var v202 int64 = 395
                                                                                                                    _ = v201
                                                                                                                    _ = v202
                                                                                                                } else {
                                                                                                                    v203 := GenS16{F0: 127}
                                                                                                                    fmt.Println(fi79)
                                                                                                                    var v204 int32 = (v44 << ((509) & 63))
                                                                                                                    _ = v203
                                                                                                                    _ = v204
                                                                                                                }
                                                                                                            case 0:
                                                                                                                v3 = v3
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v176
                                                                                                    _ = v177
                                                                                                } else {
                                                                                                    var v205 uint64 = (v84 % ((uint64(v167)) | 1))
                                                                                                    var v206 bool = false
                                                                                                    _ = v205
                                                                                                    _ = v206
                                                                                                }
                                                                                            default:
                                                                                                var v207 uint64 = uint64(v167)
                                                                                                _ = v207
                                                                                            }
                                                                                            v167 += (v167 / (((^(v167))) | 1))
                                                                                        } else {
                                                                                            fmt.Println(v167, v169, v167)
                                                                                            v63 = v63
                                                                                            v85 += (-(v44))
                                                                                        }
                                                                                    }
                                                                                case -278:
                                                                                    v167 = uint16(v84)
                                                                                }
                                                                            default:
                                                                                var v208 uint32 = 0
                                                                                _ = v208
                                                                            }
                                                                            v167 += (v167 - 2)
                                                                            _ = v169
                                                                        } else {
                                                                            v3 -= 255
                                                                            v85 += v44
                                                                            v3 = v3
                                                                        }
                                                                    }
                                                                    _ = v167
                                                                }
                                                            }
                                                            v3 -= 1
                                                            fmt.Println(v85, fi62)
                                                            _ = v165
                                                        }
                                                    default:
                                                        v63 = (float64((3.14159 - 3.14159)) / v63)
                                                    }
                                                default:
                                                    var v209 int16 = 42
                                                    _ = v209
                                                }
                                                v63 += v63
                                                v210 := GenS17{F0: ""}
                                                _ = v210
                                            }
                                        }
                                    }
                                    _ = v85
                                }
                                var v211 uint = (^(v3))
                                _ = v211
                            }
                            fmt.Println(v44)
                            fmt.Println(v63, fi62)
                        } else {
                            if len(v2) > 0 {
                                v2[0] = ("a" + "hello")
                            }
                            var v212 int64 = -9223372036854775808
                            _ = v212
                        }
                    }
                    _ = v83
                    _ = v84
                }
            default:
                v44 -= (v44 ^ (^(v44)))
            }
            _ = fi79
        }
        _ = fi62
        _ = v63
    }
    v44 -= 42
    v213 := GenS17{F0: (((("a" + "hello") + (("\n\t" + "foo bar") + "a")) + (("\n\t" + ("foo bar" + ("" + "hello"))) + (("\n\t" + "a") + "foo bar"))) + "hello")}
    for fi214 := 0; fi214 < 8; fi214++ {
        for fi215 := 0; fi215 < 10; fi215++ {
            if (1.5 != float32(0.0)) {
                v3 += (v3 / (((v3 & (^(v3)))) | 1))
                switch "\n\t" {
                case "\n\t":
                    v216 := GenS4{F0: (v3 + 688), F1: "a", F2: v44}
                    _ = v216
                case "foo bar":
                    var v217 int64 = 0
                    _ = v217
                }
            }
            _ = fi215
        }
        var v218 string = "foo bar"
        for fi219 := 0; fi219 < 4; fi219++ {
            v220 := GenS13{F0: 401, F1: v218}
            fmt.Println(fi219, v44, fi219)
            if (("hello" + "hello") <= ("a" + v218)) {
                v3 = 1
            }
            fmt.Println(v3, v3)
            if false {
                var v221 uint8 = 128
                var v222 int8 = 42
                _ = v221
                _ = v222
            } else {
                if true {
                    v223 := GenS18{F0: 483}
                    v218 = ("foo bar" + v218)
                    if false {
                        v3 = v3
                    } else {
                        if len(v2) > 0 {
                            v2[0] = v218
                        }
                        var v224 bool = true
                        v44 += int32(fi214)
                        _ = v224
                    }
                    switch ("" + v218) {
                    case "foo bar":
                        v2 = append(v2, "a")
                    case "a":
                        var v225 int32 = v44
                        _ = v225
                    case "":
                        if (true && (47 < 968)) {
                            v44 = 0
                        }
                    }
                    _ = v223
                } else {
                    var v226 uint32 = 1
                    var v227 uint = (v3 % ((uint(v226)) | 1))
                    _ = v226
                    _ = v227
                }
                v228 := GenS6{F0: 12, F1: 574, F2: v44}
                var v229 uint8 = 127
                _ = v228
                _ = v229
            }
            _ = len(v1)
            fmt.Println(fi219, v3)
            _ = fi219
            _ = v220
        }
        _ = fi214
        _ = v218
    }
    for fi230 := 0; fi230 < 18; fi230++ {
        if false {
            switch (fi230 | (fi230 | int(v44))) {
            case -926:
                for fi231 := 0; fi231 < 11; fi231++ {
                    v232 := GenS19{F0: (v44 | v44), F1: (v44 | 581), F2: -796}
                    var v233 uint16 = 127
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    _ = fi231
                    _ = v232
                    _ = v233
                }
            default:
                for fi234 := 0; fi234 < 11; fi234++ {
                    if false {
                        var v235 int64 = -982
                        var v236 uint8 = 52
                        _ = v235
                        _ = v236
                    }
                    _ = fi234
                }
            }
            var v237 uint8 = 1
            _ = v237
        }
        switch "hello" {
        case "a":
            var v238 string = "hello"
            _ = v238
        default:
            var v239 int32 = (v44 & 602)
            _ = v239
        }
        for fi240 := 0; fi240 < 16; fi240++ {
            var v241 bool = false
            var v242 int = fi240
            _ = fi240
            _ = v241
            _ = v242
        }
        _ = fi230
    }
    if true {
        var v243 int32 = (v44 & (v44 << ((255) & 63)))
        switch "foo bar" {
        case "\n\t":
            var v244 uint8 = 127
            _ = v244
        case "foo bar":
            if false {
                for fi245 := 0; fi245 < 10; fi245++ {
                    v44 = v243
                    _ = fi245
                }
                v246 := GenS20{F0: 7, F1: (-(v243))}
                var v247 float64 = (((-2.25 - -2.25) / (float64((float32(3.14159) + -2.25)) * ((((0.1 / ((3.14159 / (1.5 - (float64((float32(-2.25) * -2.25)) - (0.1 / (0.1 * (float64(float32(0.0)) + 3.14159)))))) - 0.1)) * (-2.25 * (((3.14159 / 1e10) + 1e10) - 1.5))) / 1e10) * 3.14159))) * (1e10 * (1e10 - 0.0)))
                _ = v246
                _ = v247
            }
        default:
            _ = len(v2)
        }
        _ = v243
    }
    for fi248 := 0; fi248 < 8; fi248++ {
        for fi249 := 0; fi249 < 16; fi249++ {
            if true {
                var v250 int16 = -32768
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                _ = v250
            } else {
                var v251 int16 = 42
                var v252 uint32 = 468
                _ = v251
                _ = v252
            }
            v3 = 981
            v253 := GenS16{F0: -1}
            _ = fi249
            _ = v253
        }
        var v254 uint64 = 18446744073709551615
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        _ = fi248
        _ = v254
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v213
    _ = v3
    _ = v44
}
