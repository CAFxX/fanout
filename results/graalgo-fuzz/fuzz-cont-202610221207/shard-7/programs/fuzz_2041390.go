package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 int8; F1 int8; F2 uint32 }
type GenS5 struct { F0 uint32; F1 bool; F2 bool }
type GenS6 struct { F0 int; F1 uint64 }
type GenS7 struct { F0 uint8; F1 int }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 string; F1 int8 }
type GenS10 struct { F0 bool; F1 uint8; F2 int32 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 string; F1 uint8; F2 uint }
type GenS13 struct { F0 int; F1 uint16; F2 uint16 }
type GenS14 struct { F0 string; F1 string; F2 uint64 }
type GenS15 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS16 struct { F0 uint; F1 uint8 }
type GenS17 struct { F0 int; F1 uint32 }
type GenS18 struct { F0 int16; F1 int; F2 uint32 }
type GenS19 struct { F0 uint32; F1 uint8 }
type GenS20 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS21 struct { F0 uint16; F1 uint16 }

func main() {
    v1 := []string{"hello"}
    var v2 float32 = 1e10
    v3 := make(map[uint32]int)
    var v4 uint32 = 914
    var v5 uint16 = 32767
    if false {
        v6 := GenS1{F0: "foo bar"}
        v7 := GenS1{F0: ("hello" + "")}
        _ = v6
        _ = v7
    } else {
        if ((^(v5)) >= 2) {
            var v8 uint8 = 2
            v9 := GenS2{F0: -474, F1: 42}
            _ = v8
            _ = v9
        }
    }
    var v10 int64 = 0
    if true {
        v11 := GenS2{F0: 32767, F1: 126}
        switch "hello" {
        case "foo bar":
            var v12 uint64 = 379
            _ = v12
        case "":
            if (2147483647 == -551) {
                var v13 uint = 620
                _ = v13
            } else {
                if false {
                    v5 += v5
                    fmt.Println(v10, v4, v4)
                }
                var v14 uint8 = 128
                fmt.Println(v5, v2)
                _ = v14
            }
        case "a":
            v1 = append(v1, "\n\t")
        }
        _ = v11
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 string = "hello"
            v2 += (v2 + (v2 - 0.0))
            _ = fi16
            _ = v17
        }
        v18 := GenS1{F0: "foo bar"}
        _ = fi15
        _ = v18
    }
    var v19 float64 = ((1.5 * 0.1) * 1.5)
    for fi20 := 0; fi20 < 11; fi20++ {
        switch "foo bar" {
        case "\n\t":
            v4 = 2147483648
        }
        _ = fi20
    }
    if (!(4294967295 >= 4294967295)) {
        v4 = v4
        for fi21 := 0; fi21 < 5; fi21++ {
            v22 := GenS1{F0: "a"}
            if (-547 <= 42) {
                var v23 uint32 = 2147483647
                var v24 int16 = 1000
                v25 := GenS1{F0: ("hello" + "")}
                _ = v23
                _ = v24
                _ = v25
            }
            _ = fi21
            _ = v22
        }
    } else {
        v2 = float32(((v19 - 3.14159) - (v19 / ((v19 * v19) / (0.0 + (float64(v2) - v19))))))
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        for fi27 := 0; fi27 < 1; fi27++ {
            v28 := GenS3{F0: -212, F1: 648}
            v10 = (v10 >> ((65535) & 63))
            var v29 uint32 = (v4 - (v4 / ((v4) | 1)))
            _ = fi27
            _ = v28
            _ = v29
        }
        var v30 string = (("a" + "a") + "hello")
        for fi31 := 0; fi31 < 13; fi31++ {
            v3[2147483647] = fi26
            _ = fi31
        }
        _ = fi26
        _ = v30
    }
    for fi32 := 0; fi32 < 10; fi32++ {
        var v33 bool = false
        var v34 uint8 = 1
        _ = fi32
        _ = v33
        _ = v34
    }
    fmt.Println(v2, v5)
    v19 -= (1e10 * v19)
    for fi35 := 0; fi35 < 7; fi35++ {
        v36 := GenS4{F0: 127, F1: -1, F2: v4}
        var v37 uint16 = v5
        switch "\n\t" {
        case "a":
            var v38 float64 = float64(-2.25)
            _ = v38
        case "":
            for fi39 := 0; fi39 < 7; fi39++ {
                if (0 != 18446744073709551615) {
                    v40 := GenS1{F0: "hello"}
                    _ = v40
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    switch "foo bar" {
                    case "":
                        var v41 uint16 = v5
                        _ = v41
                    case "\n\t":
                        if ((-2.25 - float64(float32(v19))) == v19) {
                            if false {
                                var v42 uint = 18446744073709551615
                                v37 = v37
                                if ("" <= ("hello" + "a")) {
                                    var v43 uint16 = 32767
                                    _ = v43
                                }
                                switch "hello" {
                                case "":
                                    v44 := GenS5{F0: uint32(v5), F1: (v19 == (v19 * v19)), F2: ("a" < "\n\t")}
                                    _ = v44
                                case "a":
                                    var v45 uint32 = v4
                                    _ = v45
                                case "\n\t":
                                    switch fi39 {
                                    case 864:
                                        v42 = v42
                                    case -1:
                                        if (0.1 < -2.25) {
                                            switch "\n\t" {
                                            case "foo bar":
                                                v37 += (v5 >> ((42) & 63))
                                            }
                                            v46 := GenS5{F0: 1, F1: (42 <= 159), F2: true}
                                            v3[4294967295] = 0
                                            _ = v46
                                        }
                                    default:
                                        v47 := GenS6{F0: 9223372036854775807, F1: 42}
                                        _ = v47
                                    }
                                }
                                if false {
                                    var v48 uint8 = 137
                                    v49 := GenS2{F0: 491, F1: -75}
                                    v19 = 0.1
                                    if (v2 >= v2) {
                                        if (!true) {
                                            v50 := GenS1{F0: "a"}
                                            var v51 bool = false
                                            _ = v50
                                            _ = v51
                                        } else {
                                            switch ("hello" + ("\n\t" + "foo bar")) {
                                            case "foo bar":
                                                v52 := GenS1{F0: ("foo bar" + ("" + "hello"))}
                                                _ = v52
                                            case "a":
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                            v2 = v2
                                            v53 := GenS1{F0: "\n\t"}
                                            _ = v53
                                        }
                                        var v54 int = (fi39 >> (((v42 << (((^(v42))) & 63))) & 63))
                                        _ = v54
                                    }
                                    _ = v48
                                    _ = v49
                                }
                                if false {
                                    v55 := GenS4{F0: 127, F1: -95, F2: uint32(v5)}
                                    fmt.Println(v4, v4, v37)
                                    switch fi35 {
                                    case -2147483648:
                                        v56 := GenS6{F0: 0, F1: 42}
                                        _ = v56
                                    case 0:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    default:
                                        v10 = 131
                                    }
                                    _ = v55
                                }
                                _ = v42
                            } else {
                                var v57 uint64 = 1
                                _ = v57
                            }
                            if (!false) {
                                if false {
                                    if ("a" >= "hello") {
                                        v5 = (v37 ^ (v37 % (((v5 | 32767)) | 1)))
                                        if true {
                                            v5 -= v5
                                            if false {
                                                var v58 int = (fi35 - (fi35 << ((65535) & 63)))
                                                _ = v58
                                            } else {
                                                var v59 int16 = -370
                                                v19 += v19
                                                _ = v59
                                            }
                                        }
                                    }
                                    var v60 uint32 = 42
                                    v10 -= 235
                                    _ = v60
                                } else {
                                    if true {
                                        switch "a" {
                                        case "hello":
                                            fmt.Println(v5)
                                        case "foo bar":
                                            fmt.Println(fi39, v2)
                                        default:
                                            v61 := GenS7{F0: 2, F1: 235}
                                            _ = v61
                                        }
                                        delete(v3, v4)
                                        if true {
                                            fmt.Println(v19)
                                            if true {
                                                if (-1 <= 42) {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        switch fi35 {
                                                        case 128:
                                                            var v62 int8 = 127
                                                            _ = v62
                                                        case -194:
                                                            v37 = v5
                                                        case 42:
                                                            var v63 float64 = ((v19 * (v19 + v19)) * 3.14159)
                                                            _ = v63
                                                        }
                                                        fmt.Println(v37, fi35, v4)
                                                    case "a":
                                                        switch 2147483647 {
                                                        case 32768:
                                                            var v64 uint8 = 0
                                                            _ = v64
                                                        }
                                                        v4 = v4
                                                    case "\n\t":
                                                        var v65 int = (fi39 % (((fi39 * fi35)) | 1))
                                                        _ = v65
                                                    default:
                                                        v1 = append(v1, "hello")
                                                        v19 += (v19 / 1.5)
                                                    }
                                                    v5 = v5
                                                    switch fi39 {
                                                    case 0:
                                                        v3[(v4 >> ((2) & 63))] = (fi35 / (((fi39 % ((fi35) | 1))) | 1))
                                                    }
                                                } else {
                                                    var v66 uint32 = (v4 / ((v4) | 1))
                                                    fmt.Println(v2, v4)
                                                    delete(v3, (v4 % ((2) | 1)))
                                                    v5 = (^(v37))
                                                    v1 = append(v1, "a")
                                                    _ = v66
                                                }
                                            } else {
                                                v19 += 0.0
                                                v10 = (v10 | v10)
                                            }
                                            var v67 float64 = (float64((1e10 / 3.14159)) - 0.0)
                                            _ = v67
                                        } else {
                                            v5 += 0
                                            v68 := GenS8{F0: 18446744073709551615}
                                            var v69 int = (fi39 & (-(fi39)))
                                            _ = v68
                                            _ = v69
                                        }
                                    } else {
                                        var v70 int16 = 1
                                        var v71 float64 = v19
                                        var v72 uint64 = 42
                                        _ = v70
                                        _ = v71
                                        _ = v72
                                    }
                                    delete(v3, (^(v4)))
                                    switch "foo bar" {
                                    case "a":
                                        fmt.Println(v37, v5)
                                    case "\n\t":
                                        switch "" {
                                        case "":
                                            if false {
                                                var v73 uint32 = v4
                                                v2 = v2
                                                v37 = v37
                                                _ = v73
                                            }
                                        case "a":
                                            v74 := GenS9{F0: "", F1: -128}
                                            _ = v74
                                        case "foo bar":
                                            if (18446744073709551615 <= 359) {
                                                var v75 int16 = 32767
                                                if true {
                                                    v19 -= 3.14159
                                                }
                                                _ = v75
                                            }
                                            v10 = (v10 + (v10 ^ (v10 & int64(fi35))))
                                        default:
                                            var v76 float64 = ((v19 * 3.14159) + (((v19 + (3.14159 + ((float64((-2.25 / ((((((float32(-2.25) / (v2 - 1e10)) - (v2 * 3.14159)) - 1e10) * v2) - v2) * 0.1))) / ((((1e10 + (0.0 + 0.1)) + (0.0 * v19)) * (1e10 * 3.14159)) / 3.14159)) + (float64(((((3.14159 / v2) * -2.25) + (-2.25 / (v2 + v2))) * (float32(0.0) - 0.0))) * v19)))) / v19) - (v19 + (v19 * -2.25))))
                                            _ = v76
                                        }
                                    case "":
                                        var v77 int8 = -108
                                        _ = v77
                                    }
                                }
                            }
                        } else {
                            v5 = uint16(v4)
                            v78 := GenS2{F0: 1000, F1: -16}
                            _ = v78
                        }
                    default:
                        v79 := GenS10{F0: (!(true || (!(true && false)))), F1: 42, F2: 42}
                        _ = v79
                    }
                    var v80 int32 = 432
                    _ = v80
                }
                _ = fi39
            }
        case "hello":
            switch "" {
            case "":
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
            case "a":
                if true {
                    v2 = float32(v19)
                    for fi81 := 0; fi81 < 16; fi81++ {
                        var v82 uint16 = v5
                        v83 := GenS11{F0: fi81}
                        _ = fi81
                        _ = v82
                        _ = v83
                    }
                    switch ("hello" + "") {
                    case "\n\t":
                        switch ("\n\t" + "a") {
                        case "a":
                            var v84 uint = uint(v4)
                            _ = v84
                        case "":
                            switch fi35 {
                            case -1:
                                var v85 int8 = 42
                                _ = v85
                            case 32768:
                                var v86 bool = (948 > 1)
                                _ = v86
                            }
                        }
                    case "hello":
                        if true {
                            var v87 int64 = 9223372036854775807
                            if (-32768 >= int16(v4)) {
                                var v88 int64 = (v10 - (v87 - (v10 & (v87 * 391))))
                                _ = v88
                            } else {
                                v2 += v2
                                switch ((("hello" + "") + "") + "") {
                                case "foo bar":
                                    v89 := GenS6{F0: (fi35 / ((fi35) | 1)), F1: 1}
                                    _ = v89
                                case "\n\t":
                                    for fi90 := 0; fi90 < 20; fi90++ {
                                        if false {
                                            v91 := GenS12{F0: (("hello" + (((((("" + "") + "\n\t") + "a") + "a") + "a") + "a")) + ""), F1: 255, F2: 255}
                                            v2 -= v2
                                            _ = v91
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        switch "hello" {
                                        case "\n\t":
                                            v92 := GenS11{F0: 2147483647}
                                            _ = v92
                                        case "foo bar":
                                            v87 += 525
                                        case "hello":
                                            switch 42 {
                                            case 42:
                                                var v93 int16 = -818
                                                _ = v93
                                            case 4294967295:
                                                var v94 int16 = -32768
                                                _ = v94
                                            }
                                        }
                                        var v95 int8 = -42
                                        v96 := GenS5{F0: 431, F1: (0 == 2), F2: (false || false)}
                                        _ = fi90
                                        _ = v95
                                        _ = v96
                                    }
                                default:
                                    if ((v37 - v5) != v5) {
                                        var v97 uint = 2
                                        _ = v97
                                    }
                                }
                                fmt.Println(v37, fi35, v4)
                            }
                            _ = v87
                        }
                    default:
                        v1 = append(v1, ("\n\t" + "foo bar"))
                    }
                } else {
                    if false {
                        for fi98 := 0; fi98 < 3; fi98++ {
                            switch (-(fi35)) {
                            case 0:
                                v37 = (v5 & (v37 >> ((255) & 63)))
                            default:
                                if (!true) {
                                    if false {
                                        v5 = v5
                                    } else {
                                        switch (fi98 - (fi98 + fi35)) {
                                        case 2147483647:
                                            if (42 == uint64(v37)) {
                                                delete(v3, 26)
                                            } else {
                                                if true {
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v37 = (v37 + v5)
                                                    case "a":
                                                        v3[(^(v4))] = (^(fi98))
                                                    default:
                                                        v99 := GenS13{F0: -289, F1: v5, F2: v37}
                                                        _ = v99
                                                    }
                                                    v37 += v5
                                                    v100 := GenS14{F0: "a", F1: "foo bar", F2: 2}
                                                    _ = v100
                                                }
                                                if (1 <= 42) {
                                                    v19 = 0.1
                                                    v101 := GenS15{F0: (v4 >> ((42) & 63)), F1: ((656 != 42) && (false && true)), F2: (^(v4))}
                                                    _ = v101
                                                }
                                            }
                                        case 128:
                                            v102 := GenS16{F0: 0, F1: 212}
                                            _ = v102
                                        }
                                        if ((v19 - v19) >= 1e10) {
                                            v2 -= -2.25
                                        }
                                    }
                                    v103 := GenS17{F0: (fi98 >> ((65535) & 63)), F1: 319}
                                    var v104 int8 = 100
                                    v105 := GenS15{F0: 42, F1: ((float32(-2.25) + float32(-2.25)) == v2), F2: 2}
                                    v2 -= v2
                                    switch ("\n\t" + ("foo bar" + "\n\t")) {
                                    case "":
                                        switch "a" {
                                        case "hello":
                                            v106 := v3[v4]
                                            _ = v106
                                        default:
                                            var v107 bool = (v4 <= (v4 << ((0) & 63)))
                                            _ = v107
                                        }
                                    case "foo bar":
                                        switch fi98 {
                                        case 2147483647:
                                            var v108 float64 = 1e10
                                            _ = v108
                                        default:
                                            fmt.Println(v104, v37, v4)
                                        }
                                    case "a":
                                        v109 := GenS13{F0: int(v5), F1: 1, F2: (v37 | 255)}
                                        _ = v109
                                    default:
                                        if false {
                                            if false {
                                                if true {
                                                    switch fi35 {
                                                    case 629:
                                                        var v110 uint64 = 1
                                                        _ = v110
                                                    case -553:
                                                        v4 += v4
                                                    default:
                                                        if ("" < "\n\t") {
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                var v111 uint = 255
                                                                _ = v111
                                                            case "foo bar":
                                                                if ((v10 * v10) != -9223372036854775808) {
                                                                    if (v2 >= v2) {
                                                                        var v112 float64 = v19
                                                                        v19 = float64(0.0)
                                                                        var v113 int64 = -753
                                                                        _ = v112
                                                                        _ = v113
                                                                    } else {
                                                                        if (false || false) {
                                                                            v114 := GenS11{F0: 128}
                                                                            _ = v114
                                                                        } else {
                                                                            var v115 float32 = v2
                                                                            if (!true) {
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                v115 -= (0.1 - ((((0.0 / 0.1) * (v2 * v115)) / v2) + v2))
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = ""
                                                                                }
                                                                            }
                                                                            _ = v115
                                                                        }
                                                                    }
                                                                    fmt.Println(v104, fi98, v4)
                                                                    var v116 float32 = v2
                                                                    v116 += ((3.14159 / ((v2 / (v116 + float32((v19 / -2.25)))) / 0.1)) + (-2.25 / 0.1))
                                                                    switch fi35 {
                                                                    case -709:
                                                                        v2 = (float32(-2.25) / (v116 / ((1e10 - 3.14159) - (v116 * (0.0 / float32(float64((3.14159 / -2.25))))))))
                                                                        var v117 int32 = -693
                                                                        _ = v117
                                                                    case 2147483647:
                                                                        v5 = v5
                                                                    case 32768:
                                                                        if false {
                                                                            switch ("" + "hello") {
                                                                            case "foo bar":
                                                                                if true {
                                                                                    var v118 int8 = v104
                                                                                    var v119 uint16 = 0
                                                                                    var v120 uint32 = v4
                                                                                    var v121 float32 = ((v2 + v2) * (v2 - (((v2 * 0.0) * ((v2 * v2) - v2)) * 0.1)))
                                                                                    v5 = 452
                                                                                    _ = v118
                                                                                    _ = v119
                                                                                    _ = v120
                                                                                    _ = v121
                                                                                } else {
                                                                                    var v122 int32 = 42
                                                                                    _ = v122
                                                                                }
                                                                            case "a":
                                                                                var v123 int32 = 1
                                                                                _ = v123
                                                                            default:
                                                                                var v124 uint8 = 1
                                                                                _ = v124
                                                                            }
                                                                            if true {
                                                                                v3[(v4 / ((44) | 1))] = (fi35 | (-(fi98)))
                                                                            }
                                                                        } else {
                                                                            v5 = v37
                                                                            v4 += v4
                                                                        }
                                                                    }
                                                                    _ = v116
                                                                } else {
                                                                    if true {
                                                                        v19 += (3.14159 * ((1e10 - (v19 + float64(1e10))) - v19))
                                                                        v5 = v37
                                                                        var v125 int32 = 0
                                                                        _ = v125
                                                                    } else {
                                                                        var v126 uint32 = (v4 + 0)
                                                                        v127 := GenS7{F0: 127, F1: fi35}
                                                                        switch (fi35 * 32768) {
                                                                        case -901:
                                                                            var v128 int = fi98
                                                                            _ = v128
                                                                        }
                                                                        _ = v126
                                                                        _ = v127
                                                                    }
                                                                    switch -321 {
                                                                    case 619:
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                    }
                                                                    if ((0.1 + v19) <= v19) {
                                                                        var v129 int64 = 0
                                                                        v130 := v3[v4]
                                                                        v10 = (v10 | v129)
                                                                        _ = v129
                                                                        _ = v130
                                                                    } else {
                                                                        if ((fi35 | (fi98 / (((-(fi35))) | 1))) != (fi35 & int(v37))) {
                                                                            var v131 int16 = int16(v4)
                                                                            _ = v131
                                                                        }
                                                                        v3[(^(v4))] = (-(fi98))
                                                                        if (42 > v5) {
                                                                            v37 = v37
                                                                            var v132 uint16 = v5
                                                                            var v133 uint64 = 1
                                                                            _ = v132
                                                                            _ = v133
                                                                        }
                                                                        v5 = v37
                                                                        fmt.Println(fi35, fi35)
                                                                        v5 = v5
                                                                    }
                                                                }
                                                            case "a":
                                                                if true {
                                                                    v104 -= v104
                                                                    fmt.Println(v104, v104)
                                                                    v134 := GenS15{F0: v4, F1: false, F2: (v4 & (v4 * v4))}
                                                                    v5 += 255
                                                                    _ = v134
                                                                } else {
                                                                    fmt.Println(v4, v104, v10)
                                                                }
                                                                v135 := GenS6{F0: fi35, F1: 4294967295}
                                                                if ((3.14159 - (((v2 / (((v2 - 0.0) + 0.0) - v2)) * (v2 - 1.5)) + (0.1 - v2))) >= 3.14159) {
                                                                    v2 = 0.0
                                                                } else {
                                                                    v136 := v3[42]
                                                                    fmt.Println(fi35, v104, fi35)
                                                                    _ = v136
                                                                }
                                                                _ = v135
                                                            }
                                                            switch (fi98 >> ((966) & 63)) {
                                                            case 1:
                                                                v4 = (^(v4))
                                                            case 32768:
                                                                switch "" {
                                                                case "\n\t":
                                                                    v137 := GenS11{F0: fi35}
                                                                    _ = v137
                                                                default:
                                                                    v138 := v3[v4]
                                                                    v139 := GenS5{F0: v4, F1: (!(v104 <= (v104 / ((int8(v37)) | 1)))), F2: false}
                                                                    if true {
                                                                        if false {
                                                                            var v140 uint = 1
                                                                            if (1 < 42) {
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                v3[v4] = (fi98 * (fi35 / ((v138) | 1)))
                                                                                _ = len(v3)
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = ("hello" + "a")
                                                                                }
                                                                            }
                                                                            v141 := GenS2{F0: 42, F1: -74}
                                                                            _ = v140
                                                                            _ = v141
                                                                        }
                                                                        v10 = v10
                                                                        if len(v1) > 0 {
                                                                            v1[0] = "hello"
                                                                        }
                                                                    } else {
                                                                        v142 := GenS18{F0: -659, F1: (^(fi35)), F2: v4}
                                                                        _ = v142
                                                                    }
                                                                    _ = v138
                                                                    _ = v139
                                                                }
                                                            default:
                                                                v143 := GenS19{F0: 4294967295, F1: 42}
                                                                _ = v143
                                                            }
                                                        } else {
                                                            var v144 int64 = 9223372036854775807
                                                            var v145 int16 = -1
                                                            var v146 int = 984
                                                            _ = v144
                                                            _ = v145
                                                            _ = v146
                                                        }
                                                    }
                                                    switch "hello" {
                                                    case "\n\t":
                                                        v104 = v104
                                                    default:
                                                        if ((0.0 != float32((v19 + v19))) && ((!true) && true)) {
                                                            v147 := GenS20{F0: (v37 % ((2) | 1)), F1: (2147483647 >= (fi35 >> ((0) & 63))), F2: (v5 * 65535)}
                                                            _ = v147
                                                        }
                                                    }
                                                } else {
                                                    _ = len(v1)
                                                }
                                                v104 = int8(fi98)
                                                var v148 bool = true
                                                _ = v148
                                            }
                                        }
                                        v104 -= -46
                                        if ((("foo bar" + (((("foo bar" + ("hello" + "foo bar")) + "foo bar") + ("\n\t" + "\n\t")) + ("\n\t" + "hello"))) + "a") <= "hello") {
                                            v37 -= (v37 | v37)
                                            switch "" {
                                            case "foo bar":
                                                v10 = (v10 / ((v10) | 1))
                                            default:
                                                if false {
                                                    v149 := GenS21{F0: v5, F1: 255}
                                                    v19 = v19
                                                    _ = v149
                                                } else {
                                                    var v150 float64 = ((float64((v2 / -2.25)) + v19) - (0.0 / v19))
                                                    _ = v150
                                                }
                                            }
                                        }
                                    }
                                    _ = v103
                                    _ = v104
                                    _ = v105
                                } else {
                                    fmt.Println(fi35)
                                    var v151 uint32 = (^(v4))
                                    _ = v151
                                }
                            }
                            _ = fi98
                        }
                        v1 = append(v1, "\n\t")
                    } else {
                        _ = len(v1)
                    }
                }
            }
        }
        _ = fi35
        _ = v36
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
