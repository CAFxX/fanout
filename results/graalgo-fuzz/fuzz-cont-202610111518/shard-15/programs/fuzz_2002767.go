package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 uint32; F1 string }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint16; F1 string }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 int64; F1 int8; F2 uint }
type GenS18 struct { F0 uint64 }
func genFn0(p0 uint8, p1 int64) int8 {
    if (!true) {
        v6 := []int{537, -2147483648, -9223372036854775808}
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 int64 = p1
            switch 1 {
            case -107:
                switch "" {
                case "foo bar":
                    v8 = (v8 ^ (v8 * v8))
                case "":
                    p1 -= v8
                }
            case -112:
                if false {
                    switch "hello" {
                    case "\n\t":
                        var v9 float32 = (-2.25 * 1e10)
                        _ = v9
                    case "foo bar":
                        v10 := GenS1{F0: ((620 >= 512) || false), F1: (false || false), F2: -1}
                        _ = v10
                    case "hello":
                        var v11 uint16 = 268
                        _ = v11
                    }
                    var v12 int64 = -19
                    _ = v12
                }
            case 525:
                for fi13 := 0; fi13 < 6; fi13++ {
                    v8 = (p1 / ((p1) | 1))
                    p0 = (p0 - (p0 - p0))
                    if (false && true) {
                        if (p0 != (p0 / ((p0) | 1))) {
                            if false {
                                if true {
                                    _ = len(v6)
                                    if false {
                                        var v14 uint = 255
                                        v14 += v14
                                        if (128 <= p0) {
                                            fmt.Println(p1, p0, fi13)
                                        } else {
                                            if len(v6) > 0 {
                                                v6[0] = fi7
                                            }
                                            if (fi13 >= fi7) {
                                                v15 := GenS2{F0: 1}
                                                _ = v15
                                            }
                                        }
                                        _ = v14
                                    }
                                    if true {
                                        if (false && false) {
                                            v16 := GenS3{F0: "foo bar"}
                                            p1 = 42
                                            _ = v16
                                        }
                                        p1 = (p1 << ((42) & 63))
                                    } else {
                                        fmt.Println(p1, fi7)
                                        var v17 uint64 = 1
                                        v18 := GenS4{F0: 662}
                                        _ = v17
                                        _ = v18
                                    }
                                }
                            } else {
                                p1 = (^(v8))
                            }
                            switch fi13 {
                            case -9223372036854775808:
                                if true {
                                    v19 := GenS2{F0: 65535}
                                    v20 := GenS5{F0: (p1 >> ((255) & 63)), F1: 2}
                                    _ = v19
                                    _ = v20
                                } else {
                                    switch ("foo bar" + "hello") {
                                    case "foo bar":
                                        if false {
                                            if false {
                                                v21 := GenS2{F0: 1}
                                                _ = v21
                                            } else {
                                                if (1 >= 32767) {
                                                    v22 := GenS6{F0: p0, F1: ""}
                                                    _ = v22
                                                }
                                                _ = len(v6)
                                                _ = len(v6)
                                            }
                                            if ("hello" == "foo bar") {
                                                var v23 int8 = -66
                                                _ = v23
                                            } else {
                                                var v24 uint64 = 1
                                                v25 := GenS3{F0: "hello"}
                                                if true {
                                                    var v26 int16 = 1
                                                    p0 -= p0
                                                    var v27 string = "\n\t"
                                                    _ = v26
                                                    _ = v27
                                                }
                                                _ = v24
                                                _ = v25
                                            }
                                            p0 = 128
                                        }
                                    case "":
                                        if true {
                                            p0 = 243
                                            v28 := GenS7{F0: 2, F1: "foo bar"}
                                            var v29 bool = false
                                            _ = v28
                                            _ = v29
                                        }
                                    }
                                    switch "\n\t" {
                                    case "foo bar":
                                        if (int(p0) >= fi7) {
                                            if ((42 != v8) && false) {
                                                switch "a" {
                                                case "a":
                                                    var v30 bool = (v8 == 42)
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    var v31 uint16 = 767
                                                    _ = v30
                                                    _ = v31
                                                default:
                                                    var v32 float32 = (((((-2.25 + 3.14159) / 3.14159) + float32(1e10)) / (3.14159 + -2.25)) * 1.5)
                                                    _ = v32
                                                }
                                                if (55 < 2147483647) {
                                                    if (false || (2147483647 <= 445)) {
                                                        if false {
                                                            var v33 uint8 = (p0 + 255)
                                                            switch "foo bar" {
                                                            case "hello":
                                                                var v34 int = fi13
                                                                _ = v34
                                                            }
                                                            _ = v33
                                                        } else {
                                                            p0 -= 143
                                                        }
                                                        fmt.Println(v8)
                                                    }
                                                    switch (fi13 + (fi13 / (((fi13 - (fi7 / (((^(fi7))) | 1)))) | 1))) {
                                                    case 2147483647:
                                                        switch (fi13 >> ((732) & 63)) {
                                                        case 127:
                                                            var v35 uint32 = 2147483647
                                                            _ = v35
                                                        default:
                                                            if ((1.5 / 1.0) < 0.1) {
                                                                fmt.Println(p0)
                                                            }
                                                        }
                                                    case 0:
                                                        v8 = v8
                                                    default:
                                                        var v36 uint32 = 0
                                                        _ = v36
                                                    }
                                                    switch "a" {
                                                    case "hello":
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    }
                                                } else {
                                                    if (false && (!(!(489 <= -1)))) {
                                                        v6 = append(v6, -129)
                                                    } else {
                                                        var v37 int = (fi13 >> ((472) & 63))
                                                        var v38 int16 = 1
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                    var v39 uint8 = 238
                                                    _ = v39
                                                }
                                            } else {
                                                if (!(true || true)) {
                                                    v40 := GenS2{F0: 466}
                                                    fmt.Println(v8, fi7, v8)
                                                    _ = len(v6)
                                                    p1 += (p1 * (v8 & (p1 << ((4294967295) & 63))))
                                                    _ = v40
                                                } else {
                                                    if ((493 != 489) || false) {
                                                        var v41 uint8 = 1
                                                        switch fi7 {
                                                        case -9223372036854775808:
                                                            fmt.Println(p1)
                                                        case 9223372036854775807:
                                                            var v42 int = -780
                                                            _ = v42
                                                        }
                                                        _ = v41
                                                    }
                                                }
                                            }
                                            p0 += (p0 * (p0 ^ (p0 * (p0 ^ 128))))
                                        }
                                    case "":
                                        if len(v6) > 0 {
                                            v6[0] = (fi13 | fi13)
                                        }
                                        switch -81 {
                                        case 127:
                                            if false {
                                                if (42 >= p0) {
                                                    switch fi7 {
                                                    case -2147483648:
                                                        var v43 int = (fi7 & fi7)
                                                        _ = v43
                                                    case 0:
                                                        var v44 uint16 = 32767
                                                        _ = v44
                                                    case 32767:
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            var v45 uint64 = 18446744073709551615
                                                            _ = v45
                                                        }
                                                    }
                                                }
                                            } else {
                                                p0 = 0
                                                if true {
                                                    p1 = (p1 % (((p1 | (v8 + p1))) | 1))
                                                    v46 := GenS7{F0: 220, F1: ("\n\t" + ("a" + "hello"))}
                                                    if ((1.5 + (1.5 * (0.1 / 1.0))) != 0.0) {
                                                        p0 -= p0
                                                        var v47 int64 = (v8 | (v8 / ((0) | 1)))
                                                        _ = v47
                                                    } else {
                                                        v48 := GenS8{F0: 443}
                                                        p0 = p0
                                                        fmt.Println(v8, v8)
                                                        _ = v48
                                                    }
                                                    _ = v46
                                                }
                                                fmt.Println(fi13, v8, v8)
                                            }
                                        case 42:
                                            if len(v6) > 0 {
                                                v6[0] = (fi13 + fi7)
                                            }
                                        default:
                                            p1 = p1
                                        }
                                        v6 = append(v6, fi13)
                                    default:
                                        var v49 int = (fi7 / ((425) | 1))
                                        _ = v49
                                    }
                                    if (true && true) {
                                        switch "hello" {
                                        case "foo bar":
                                            fmt.Println(fi13, fi13, p1)
                                        default:
                                            if (65535 < 4294967295) {
                                                if true {
                                                    v8 = -17
                                                    p0 -= 255
                                                } else {
                                                    if true {
                                                        switch 4294967295 {
                                                        case -577:
                                                            fmt.Println(fi13)
                                                        case 127:
                                                            fmt.Println(p1)
                                                        default:
                                                            _ = len(v6)
                                                        }
                                                        switch (fi13 / ((-336) | 1)) {
                                                        case 0:
                                                            fmt.Println(v8, p1, v8)
                                                        }
                                                        fmt.Println(fi7, fi13, p1)
                                                    } else {
                                                        switch ("" + ("\n\t" + "\n\t")) {
                                                        case "hello":
                                                            var v50 int = 4294967295
                                                            _ = v50
                                                        case "":
                                                            fmt.Println(p0, v8, fi7)
                                                        default:
                                                            if true {
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    v51 := GenS6{F0: (p0 + p0), F1: ""}
                                                                    _ = v51
                                                                case "foo bar":
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        p0 = p0
                                                                    default:
                                                                        p0 -= p0
                                                                    }
                                                                default:
                                                                    if (false && false) {
                                                                        var v52 int = int(p0)
                                                                        switch -1 {
                                                                        case -327:
                                                                            v53 := GenS7{F0: 2, F1: ""}
                                                                            _ = v53
                                                                        case -2147483648:
                                                                            var v54 bool = (!(true && false))
                                                                            _ = v54
                                                                        case 128:
                                                                            v52 -= 0
                                                                        }
                                                                        if (255 < 1) {
                                                                            var v55 uint = 340
                                                                            fmt.Println(p1, v8, fi13)
                                                                            _ = v55
                                                                        }
                                                                        _ = v52
                                                                    }
                                                                }
                                                                var v56 uint64 = 889
                                                                _ = v56
                                                            } else {
                                                                if true {
                                                                    v8 += (^(p1))
                                                                } else {
                                                                    if (("foo bar" + "hello") > "foo bar") {
                                                                        var v57 float32 = (0.0 * float32(0.1))
                                                                        var v58 uint16 = 492
                                                                        _ = v57
                                                                        _ = v58
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = len(v6)
                                                    }
                                                }
                                                var v59 bool = ((((1.5 / 1.5) <= 1e10) || (!true)) || (true && ((fi7 % ((fi13) | 1)) >= (fi13 & (fi7 << ((255) & 63))))))
                                                var v60 int = 32767
                                                _ = v59
                                                _ = v60
                                            }
                                        }
                                        v8 = (-(p1))
                                        var v61 bool = false
                                        _ = v61
                                    }
                                }
                            case 9223372036854775807:
                                var v62 int32 = -843
                                _ = v62
                            default:
                                v63 := GenS7{F0: 4294967295, F1: "\n\t"}
                                _ = v63
                            }
                        } else {
                            fmt.Println(v8)
                        }
                        var v64 uint8 = (p0 & (p0 | p0))
                        _ = v64
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        p1 = 0
                        if (!false) {
                            switch "\n\t" {
                            case "a":
                                var v65 int64 = -9223372036854775808
                                _ = v65
                            default:
                                p1 -= 42
                                v66 := GenS8{F0: 2147483647}
                                p0 += p0
                                if ((1e10 - -2.25) == (0.0 - (float64(1e10) + (((1.5 / ((1e10 + (0.0 / 0.1)) + (3.14159 - 3.14159))) / 1e10) * 3.14159)))) {
                                    if (v8 <= (p1 * 299)) {
                                        if ((fi7 * fi7) > (fi13 & -9223372036854775808)) {
                                            v67 := GenS9{F0: (p0 << ((796) & 63))}
                                            var v68 int64 = -540
                                            v69 := GenS4{F0: 42}
                                            p1 += 938
                                            _ = v67
                                            _ = v68
                                            _ = v69
                                        }
                                        v70 := GenS10{F0: -29}
                                        _ = v70
                                    }
                                    switch (fi13 - fi7) {
                                    case 2147483647:
                                        if (142 <= 1) {
                                            var v71 float64 = 3.14159
                                            v71 -= v71
                                            _ = v71
                                        } else {
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        }
                                    case 32768:
                                        fmt.Println(v8, v8, v8)
                                    }
                                    v72 := GenS8{F0: 1}
                                    p1 += 495
                                    if len(v6) > 0 {
                                        v6[0] = 32768
                                    }
                                    _ = v72
                                }
                                fmt.Println(v8, p0)
                                _ = v66
                            }
                            v6 = append(v6, fi13)
                            switch "a" {
                            case "\n\t":
                                switch "a" {
                                case "hello":
                                    var v73 uint32 = 42
                                    _ = v73
                                case "a":
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                default:
                                    if false {
                                        p1 = v8
                                        p0 = p0
                                    }
                                }
                            }
                            v74 := GenS4{F0: 18446744073709551615}
                            _ = len(v6)
                            v8 = (v8 + (p1 << ((2) & 63)))
                            _ = v74
                        }
                    }
                    _ = fi13
                }
            default:
                if true {
                    if true {
                        var v75 float64 = (1.5 - -2.25)
                        p1 = 0
                        _ = v75
                    } else {
                        for fi76 := 0; fi76 < 14; fi76++ {
                            if true {
                                if (false || (true || (-1 <= 1))) {
                                    if ((!(((true && true) && false) || true)) && (!true)) {
                                        _ = len(v6)
                                    }
                                }
                            } else {
                                p0 = p0
                            }
                            _ = fi76
                        }
                        fmt.Println(p1, p1, p1)
                        fmt.Println(v8, v8)
                    }
                    if false {
                        var v77 string = (("" + "\n\t") + "\n\t")
                        _ = v77
                    }
                    for fi78 := 0; fi78 < 1; fi78++ {
                        if false {
                            p0 += 0
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        } else {
                            switch "" {
                            case "":
                                p0 += (p0 / ((33) | 1))
                            case "\n\t":
                                v79 := GenS6{F0: 1, F1: "\n\t"}
                                _ = v79
                            case "foo bar":
                                fmt.Println(p1, p1)
                            }
                        }
                        fmt.Println(p1)
                        _ = fi78
                    }
                } else {
                    p0 -= p0
                    for fi80 := 0; fi80 < 18; fi80++ {
                        v81 := GenS1{F0: (fi80 > (-(fi7))), F1: true, F2: 413}
                        _ = len(v6)
                        v82 := GenS2{F0: 18}
                        v8 = v8
                        _ = fi80
                        _ = v81
                        _ = v82
                    }
                }
            }
            v8 -= (p1 | v8)
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        p1 += p1
        v83 := GenS11{F0: (514 >= 1), F1: 1, F2: 18446744073709551615}
        var v84 bool = false
        _ = v83
        _ = v84
    }
    for fi85 := 0; fi85 < 11; fi85++ {
        for fi86 := 0; fi86 < 8; fi86++ {
            if false {
                var v87 uint8 = p0
                p0 -= v87
                switch "" {
                case "\n\t":
                    if false {
                        var v88 int = (fi85 / ((fi86) | 1))
                        v88 = -211
                        _ = v88
                    }
                default:
                    var v89 float64 = ((float64(0.1) + 0.0) - (float64(1.5) * ((1e10 - ((0.0 * ((3.14159 - -2.25) / float64((((-2.25 * 3.14159) / (1e10 / 1.5)) / (float32(1.5) * (3.14159 * -2.25)))))) / 3.14159)) * 0.1)))
                    _ = v89
                }
                p1 = 279
                _ = v87
            }
            p1 = p1
            var v90 string = ("\n\t" + "foo bar")
            _ = fi86
            _ = v90
        }
        p1 = (p1 ^ -1)
        for fi91 := 0; fi91 < 9; fi91++ {
            var v92 int = 4294967295
            v93 := []int{fi91, (v92 | fi85), fi85, fi91}
            _ = fi91
            _ = v92
            _ = v93
        }
        _ = fi85
    }
    switch 127 {
    case 4294967295:
        var v94 int32 = 0
        _ = v94
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1() uint16 {
    var v95 float64 = 3.14159
    fmt.Println(v95)
    switch ("hello" + "foo bar") {
    case "foo bar":
        v95 = (v95 / (0.1 + v95))
    case "hello":
        var v96 uint = 275
        _ = v96
    default:
        var v97 float32 = float32(v95)
        _ = v97
    }
    _ = v95
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 0
    v3 := make(map[int]string)
    v4 := []int{}
    v5 := []string{"", "foo bar", "foo bar"}
    switch 4294967295 {
    case -9223372036854775808:
        var v98 bool = false
        _ = v98
    default:
        for fi99 := 0; fi99 < 5; fi99++ {
            fmt.Println(v2, v2)
            _ = fi99
        }
    }
    v100 := v3[128]
    var v101 uint16 = genFn1()
    for fi102 := 0; fi102 < 18; fi102++ {
        var v103 int16 = 1000
        v103 += -1
        switch fi102 {
        case -9223372036854775808:
            delete(v3, int(v101))
        case -2147483648:
            if false {
                if false {
                    if false {
                        v103 = (v103 & int16(v101))
                    } else {
                        if false {
                            var v104 uint32 = 0
                            if len(v5) > 0 {
                                v5[0] = v100
                            }
                            _ = v104
                        } else {
                            var v105 uint8 = 0
                            v101 = genFn1()
                            _ = v105
                        }
                        v103 = v103
                    }
                    var v106 int32 = 1
                    _ = v106
                }
            }
        default:
            if true {
                if false {
                    v1[v100] = v100
                    switch (^(fi102)) {
                    case -9223372036854775808:
                        switch (fi102 / ((fi102) | 1)) {
                        case 792:
                            v107 := GenS11{F0: false, F1: 620, F2: 2}
                            _ = v107
                        default:
                            var v108 int64 = -745
                            _ = v108
                        }
                    case 0:
                        if (!false) {
                            v100 = v100
                        } else {
                            switch "" {
                            case "":
                                v2 += 2
                            case "a":
                                switch ("\n\t" + ("\n\t" + "a")) {
                                case "\n\t":
                                    if (!((v101 * v101) > (^(v101)))) {
                                        var v109 float64 = (1.5 + 1.5)
                                        fmt.Println(fi102, v2, v100)
                                        var v110 bool = ((-2.25 * 3.14159) >= float32((v109 / v109)))
                                        _ = v109
                                        _ = v110
                                    }
                                case "hello":
                                    for fi111 := 0; fi111 < 7; fi111++ {
                                        if false {
                                            v2 += v2
                                            fmt.Println(v101)
                                            var v112 int = (-(fi102))
                                            switch "hello" {
                                            case "a":
                                                if false {
                                                    var v113 uint8 = 143
                                                    var v114 uint64 = 42
                                                    _ = v113
                                                    _ = v114
                                                }
                                            case "hello":
                                                fmt.Println(v101, fi102, fi102)
                                            case "foo bar":
                                                switch (fi111 << ((v2) & 63)) {
                                                case 4294967295:
                                                    var v115 bool = (!true)
                                                    _ = v115
                                                case 128:
                                                    if true {
                                                        if (845 != 2) {
                                                            v116 := GenS9{F0: 250}
                                                            if true {
                                                                if (int8(v103) <= 127) {
                                                                    v117 := GenS12{F0: (!false)}
                                                                    var v118 uint = (v2 << ((v2) & 63))
                                                                    var v119 int32 = -695
                                                                    if true {
                                                                        v112 += fi102
                                                                    }
                                                                    _ = v117
                                                                    _ = v118
                                                                    _ = v119
                                                                } else {
                                                                    fmt.Println(v101)
                                                                }
                                                                v3[fi111] = ("" + v100)
                                                                switch v100 {
                                                                case "a":
                                                                    if ((v101 % ((v101) | 1)) >= genFn1()) {
                                                                        switch v100 {
                                                                        case "\n\t":
                                                                            if false {
                                                                                var v120 uint64 = 327
                                                                                v121 := GenS9{F0: 127}
                                                                                var v122 uint = v2
                                                                                _ = v120
                                                                                _ = v121
                                                                                _ = v122
                                                                            } else {
                                                                                switch v100 {
                                                                                case "a":
                                                                                    var v123 float32 = (0.0 - (1.5 * (0.1 * (1.5 / float32(1e10)))))
                                                                                    _ = v123
                                                                                case "\n\t":
                                                                                    v2 -= v2
                                                                                case "foo bar":
                                                                                    if ((float64(0.0) * 3.14159) >= 0.1) {
                                                                                        var v124 float64 = 0.0
                                                                                        _ = v124
                                                                                    }
                                                                                }
                                                                                v100 = "foo bar"
                                                                            }
                                                                        }
                                                                    }
                                                                case "":
                                                                    var v125 uint64 = 4294967295
                                                                    _ = v125
                                                                }
                                                                var v126 uint16 = (v101 >> ((v2) & 63))
                                                                _ = v126
                                                            } else {
                                                                v127 := GenS12{F0: true}
                                                                _ = v127
                                                            }
                                                            var v128 int = fi102
                                                            _ = v116
                                                            _ = v128
                                                        }
                                                    }
                                                case 0:
                                                    var v129 float64 = (3.14159 + float64(3.14159))
                                                    _ = v129
                                                }
                                                var v130 int = -2147483648
                                                v130 += (fi111 - fi102)
                                                fmt.Println(fi111)
                                                var v131 float64 = (0.0 + (1e10 - ((0.1 / float64(1.5)) + (1.5 * 3.14159))))
                                                _ = v130
                                                _ = v131
                                            }
                                            var v132 string = v100
                                            v101 = 451
                                            v133 := GenS8{F0: 1}
                                            v101 = v101
                                            _ = v112
                                            _ = v132
                                            _ = v133
                                        } else {
                                            var v134 int8 = 0
                                            var v135 int8 = (v134 >> ((v2) & 63))
                                            if ((v134 * 0) == v135) {
                                                switch -1 {
                                                case -861:
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        if false {
                                                            v103 -= -32768
                                                            var v136 int8 = (v135 / (((v135 - -128)) | 1))
                                                            _ = v136
                                                        }
                                                    case "a":
                                                        if false {
                                                            var v137 string = v100
                                                            v138 := GenS13{F0: v101, F1: ((("\n\t" + v100) + (v100 + (v100 + v100))) + v137)}
                                                            _ = v137
                                                            _ = v138
                                                        }
                                                    default:
                                                        var v139 bool = false
                                                        _ = v139
                                                    }
                                                case 32768:
                                                    v140 := GenS14{F0: v2}
                                                    _ = v140
                                                default:
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    switch int(v101) {
                                                    case -421:
                                                        var v141 float64 = 1e10
                                                        _ = v141
                                                    case 4294967295:
                                                        var v142 float32 = 0.0
                                                        _ = v142
                                                    case -352:
                                                        v3[-1] = v100
                                                    default:
                                                        v2 -= (^(v2))
                                                    }
                                                }
                                                v143 := GenS3{F0: ("a" + "foo bar")}
                                                _ = v143
                                            }
                                            _ = v134
                                            _ = v135
                                        }
                                        _ = fi111
                                    }
                                default:
                                    var v144 int8 = -17
                                    _ = v144
                                }
                            }
                            v100 = ""
                            v103 -= (v103 << ((v2) & 63))
                        }
                    case 32767:
                        fmt.Println(fi102, v101, v2)
                    }
                    fmt.Println(v103)
                } else {
                    if false {
                        v100 = (v100 + ("hello" + v100))
                        if (806 <= 0) {
                            var v145 string = "hello"
                            _ = v145
                        } else {
                            v101 += (v101 ^ genFn1())
                        }
                        if (!(!true)) {
                            var v146 float64 = float64(0.1)
                            var v147 uint64 = 4294967295
                            _ = v146
                            _ = v147
                        }
                    }
                    v103 = 210
                }
                fmt.Println(v103, v103)
            }
        }
        _ = fi102
        _ = v103
    }
    switch -805 {
    case -791:
        for fi148 := 0; fi148 < 15; fi148++ {
            v100 = v100
            _ = fi148
        }
    case 127:
        v3[4294967295] = "\n\t"
    case 0:
        for fi149 := 0; fi149 < 9; fi149++ {
            v150 := GenS4{F0: 18446744073709551615}
            for fi151 := 0; fi151 < 18; fi151++ {
                if len(v5) > 0 {
                    v5[0] = "a"
                }
                var v152 uint16 = (v101 >> ((v2) & 63))
                var v153 uint32 = 1
                _ = fi151
                _ = v152
                _ = v153
            }
            var v154 uint = (v2 * v2)
            _ = fi149
            _ = v150
            _ = v154
        }
    default:
        _ = len(v4)
    }
    var v155 int32 = 2147483647
    v156 := v1[v100]
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi157 := 0; fi157 < 1; fi157++ {
        for fi158 := 0; fi158 < 2; fi158++ {
            var v159 uint8 = 42
            v100 = (v100 + v156)
            var v160 bool = ((0.0 - float32((1e10 + (-2.25 * ((1e10 + 0.0) * 1e10))))) < 0.1)
            _ = fi158
            _ = v159
            _ = v160
        }
        if len(v4) > 0 {
            v4[0] = fi157
        }
        _ = fi157
    }
    _ = len(v1)
    fmt.Println(v2, v155)
    switch 4294967295 {
    case 2147483647:
        fmt.Println(v100, v100, v101)
    case 4294967295:
        if (false || (false && false)) {
            v161 := GenS9{F0: 0}
            _ = v161
        } else {
            v162 := GenS2{F0: uint32(v2)}
            _ = v162
        }
    case 127:
        var v163 float64 = 3.14159
        _ = v163
    }
    if (0 >= 42) {
        if (!false) {
            if (!(-648 > -846)) {
                fmt.Println(v2)
                v164 := GenS12{F0: (genFn0(255, -9223372036854775808) == -1)}
                var v165 string = v156
                _ = v164
                _ = v165
            } else {
                var v166 int16 = 0
                if false {
                    if false {
                        fmt.Println(v155)
                        if (false && false) {
                            v166 = (v166 | (v166 << ((v2) & 63)))
                        } else {
                            var v167 uint32 = 1
                            _ = v167
                        }
                        fmt.Println(v100)
                    } else {
                        v155 += (v155 >> ((0) & 63))
                    }
                    var v168 int64 = -609
                    v169 := GenS15{F0: 211}
                    _ = v168
                    _ = v169
                } else {
                    fmt.Println(v166, v2, v156)
                    fmt.Println(v156, v155)
                    v156 = "hello"
                }
                if true {
                    var v170 bool = ((!true) || true)
                    _ = v170
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    if ((^(v166)) < -32768) {
                        v155 = (v155 << ((v2) & 63))
                        for fi171 := 0; fi171 < 7; fi171++ {
                            var v172 int = 1
                            v3[(^(fi171))] = v156
                            _ = fi171
                            _ = v172
                        }
                        switch v156 {
                        case "foo bar":
                            _ = len(v5)
                        case "a":
                            fmt.Println(v100)
                        default:
                            if false {
                                v101 += v101
                                v173 := GenS7{F0: 976, F1: "\n\t"}
                                switch -54 {
                                case -124:
                                    v2 -= 2
                                }
                                _ = v173
                            }
                        }
                    }
                    var v174 int32 = v155
                    _ = v174
                }
                _ = v166
            }
        } else {
            for fi175 := 0; fi175 < 11; fi175++ {
                switch "a" {
                case "\n\t":
                    var v176 int8 = -115
                    _ = v176
                default:
                    var v177 uint = v2
                    _ = v177
                }
                for fi178 := 0; fi178 < 16; fi178++ {
                    v2 = 0
                    _ = fi178
                }
                var v179 uint = (v2 % ((v2) | 1))
                _ = fi175
                _ = v179
            }
            v180 := GenS4{F0: 4294967295}
            _ = v180
        }
    } else {
        fmt.Println(v100, v155)
        switch "\n\t" {
        case "a":
            v100 = v100
        case "foo bar":
            v155 -= v155
        case "":
            v100 = "foo bar"
        default:
            v4 = append(v4, -9223372036854775808)
        }
    }
    v100 = v156
    fmt.Println(v155, v101)
    switch "hello" {
    case "a":
        for fi181 := 0; fi181 < 15; fi181++ {
            if (false || (0 > genFn0(37, 9223372036854775807))) {
                var v182 int64 = -1
                for fi183 := 0; fi183 < 19; fi183++ {
                    v184 := GenS15{F0: 4294967295}
                    _ = fi183
                    _ = v184
                }
                _ = v182
            } else {
                switch fi181 {
                case 32767:
                    for fi185 := 0; fi185 < 11; fi185++ {
                        _ = len(v4)
                        v155 = (v155 & (v155 * v155))
                        v101 += v101
                        _ = fi185
                    }
                case 42:
                    var v186 bool = (!true)
                    _ = v186
                case -2147483648:
                    if (-106 != 0) {
                        switch v100 {
                        case "hello":
                            v187 := GenS16{F0: 511}
                            _ = v187
                        case "":
                            var v188 uint32 = 42
                            _ = v188
                        case "a":
                            if (((((-2.25 - (float64(0.0) - float64(1.5))) / -2.25) * 0.1) == float64(0.1)) && true) {
                                v101 += genFn1()
                            } else {
                                v5 = append(v5, "hello")
                                var v189 uint16 = v101
                                if (false || (true || true)) {
                                    v190 := GenS5{F0: 315, F1: uint8(v2)}
                                    var v191 float32 = float32(1e10)
                                    _ = v190
                                    _ = v191
                                } else {
                                    for fi192 := 0; fi192 < 4; fi192++ {
                                        var v193 uint32 = 2147483647
                                        switch fi192 {
                                        case -9223372036854775808:
                                            var v194 int = fi192
                                            _ = v194
                                        default:
                                            v2 = v2
                                        }
                                        var v195 bool = (!(2 == (^(v101))))
                                        _ = fi192
                                        _ = v193
                                        _ = v195
                                    }
                                    v196 := GenS17{F0: int64(v155), F1: 127, F2: (v2 + v2)}
                                    _ = v196
                                }
                                _ = v189
                            }
                        }
                        v197 := GenS18{F0: 819}
                        _ = v197
                    } else {
                        v2 = (^(v2))
                    }
                }
                v2 += v2
            }
            v101 = (^(v101))
            _ = fi181
        }
    case "\n\t":
        _ = len(v5)
    case "hello":
        var v198 int8 = -1
        _ = v198
    default:
        var v199 int8 = -1
        _ = v199
    }
    fmt.Println(v156)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("v101:", v101)
    fmt.Println("v155:", v155)
    fmt.Println("v156:", v156)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v100
    _ = v101
    _ = v155
    _ = v156
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
