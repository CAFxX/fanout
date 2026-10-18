package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int32; F1 uint64; F2 int64 }
type GenS4 struct { F0 uint; F1 int; F2 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32; F1 uint32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint64; F1 string; F2 int }
type GenS11 struct { F0 uint32; F1 bool }
type GenS12 struct { F0 string; F1 bool; F2 int16 }
type GenS13 struct { F0 int16; F1 uint64 }
type GenS14 struct { F0 uint8; F1 bool }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 string; F1 int16; F2 uint32 }

func main() {
    var v1 float32 = (3.14159 / 1.0)
    var v2 float32 = (v1 - v1)
    var v3 float64 = 3.14159
    var v4 bool = true
    if true {
        v2 -= 3.14159
        for fi5 := 0; fi5 < 4; fi5++ {
            var v6 float32 = v2
            var v7 uint64 = 588
            _ = fi5
            _ = v6
            _ = v7
        }
        if (!(v4 && (0 < 2147483647))) {
            var v8 bool = (v4 || v4)
            v9 := GenS1{F0: 1}
            _ = v8
            _ = v9
        } else {
            v4 = (-1 != 0)
            fmt.Println(v2)
            v10 := []int{32767, 9223372036854775807}
            _ = v10
        }
    } else {
        v11 := []int{0, 128, 1}
        if (v4 && v4) {
            v2 += v2
            var v12 int64 = 9223372036854775807
            _ = v12
        } else {
            var v13 int = 32768
            _ = v13
        }
        if (422 == -32768) {
            v4 = true
            var v14 float32 = ((-2.25 / ((-2.25 * v1) / v2)) / 1.0)
            for fi15 := 0; fi15 < 18; fi15++ {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = fi15
            }
            _ = v14
        } else {
            for fi16 := 0; fi16 < 4; fi16++ {
                v2 += -2.25
                if true {
                    switch fi16 {
                    case -1:
                        var v17 string = (("hello" + "hello") + "\n\t")
                        _ = v17
                    }
                    v4 = (v4 || v4)
                }
                if len(v11) > 0 {
                    v11[0] = fi16
                }
                _ = fi16
            }
            var v18 string = "\n\t"
            _ = v18
        }
        _ = v11
    }
    v19 := GenS2{F0: -25, F1: 32767}
    v20 := []int{-2147483648, 32767, 9223372036854775807}
    if true {
        var v21 uint = 255
        _ = v21
    }
    v22 := GenS2{F0: 0, F1: 2}
    v23 := GenS2{F0: 42, F1: 42}
    v24 := GenS1{F0: 42}
    if true {
        var v25 uint32 = 967
        v26 := GenS3{F0: 2147483647, F1: 2, F2: 371}
        _ = v25
        _ = v26
    }
    switch ("a" + "foo bar") {
    case "hello":
        v20 = append(v20, 128)
    case "foo bar":
        if (!((v4 && (4294967295 <= 42)) && (0 != 216))) {
            v3 = v3
        }
    case "a":
        v2 -= 1e10
    default:
        for fi27 := 0; fi27 < 6; fi27++ {
            v28 := GenS4{F0: 4294967295, F1: (fi27 + fi27), F2: 698}
            fmt.Println(fi27)
            v3 = v3
            _ = fi27
            _ = v28
        }
    }
    switch "hello" {
    case "a":
        v3 = v3
    case "hello":
        switch "" {
        case "a":
            var v29 uint = 0
            _ = v29
        default:
            for fi30 := 0; fi30 < 2; fi30++ {
                for fi31 := 0; fi31 < 1; fi31++ {
                    var v32 string = "foo bar"
                    if (!(v4 && v4)) {
                        if v4 {
                            v33 := GenS1{F0: 2}
                            var v34 int8 = -1
                            _ = v33
                            _ = v34
                        }
                        var v35 int16 = 872
                        var v36 uint32 = 4294967295
                        _ = v35
                        _ = v36
                    }
                    v37 := GenS5{F0: -1}
                    v38 := GenS6{F0: 2147483648, F1: 42}
                    _ = fi31
                    _ = v32
                    _ = v37
                    _ = v38
                }
                v39 := GenS7{F0: 1}
                for fi40 := 0; fi40 < 12; fi40++ {
                    v41 := GenS6{F0: 348, F1: 4294967295}
                    _ = fi40
                    _ = v41
                }
                _ = fi30
                _ = v39
            }
        }
    case "foo bar":
        for fi42 := 0; fi42 < 15; fi42++ {
            switch "hello" {
            case "hello":
                fmt.Println(fi42)
            default:
                v43 := GenS6{F0: 2147483647, F1: 853}
                _ = v43
            }
            switch "foo bar" {
            case "":
                if false {
                    var v44 uint32 = 2147483647
                    for fi45 := 0; fi45 < 19; fi45++ {
                        switch "hello" {
                        case "hello":
                            var v46 int32 = 428
                            _ = v46
                        }
                        _ = fi45
                    }
                    switch fi42 {
                    case 1:
                        v47 := GenS8{F0: v44}
                        _ = v47
                    case 127:
                        v1 -= (v2 + v2)
                    case 32767:
                        fmt.Println(v1, fi42, v4)
                    default:
                        for fi48 := 0; fi48 < 11; fi48++ {
                            fmt.Println(v2, fi48)
                            v44 -= v44
                            var v49 int64 = 0
                            _ = fi48
                            _ = v49
                        }
                    }
                    _ = v44
                } else {
                    _ = len(v20)
                }
            case "foo bar":
                switch fi42 {
                case 1:
                    v1 = v2
                case -241:
                    for fi50 := 0; fi50 < 20; fi50++ {
                        if (1.5 == v3) {
                            switch "" {
                            case "foo bar":
                                v3 -= 3.14159
                            case "":
                                v3 -= float64(v1)
                            case "\n\t":
                                v51 := GenS2{F0: 1, F1: 115}
                                _ = v51
                            default:
                                v52 := GenS3{F0: 131, F1: 0, F2: -383}
                                _ = v52
                            }
                            v2 = 0.1
                        }
                        _ = len(v20)
                        var v53 string = "foo bar"
                        _ = fi50
                        _ = v53
                    }
                case 32767:
                    for fi54 := 0; fi54 < 3; fi54++ {
                        var v55 uint = 0
                        var v56 int = 53
                        _ = fi54
                        _ = v55
                        _ = v56
                    }
                }
            }
            if (815 != 65535) {
                switch "a" {
                case "a":
                    v3 = 0.1
                case "\n\t":
                    if true {
                        var v57 float32 = (float32(1e10) / 1.5)
                        if (!true) {
                            if v4 {
                                v58 := GenS7{F0: 255}
                                var v59 int32 = 1
                                if (127 == -128) {
                                    if (false || (!v4)) {
                                        v60 := GenS9{F0: 1}
                                        _ = v60
                                    }
                                }
                                _ = v58
                                _ = v59
                            } else {
                                fmt.Println(v57)
                                v57 = 3.14159
                            }
                            switch fi42 {
                            case 42:
                                fmt.Println(v3, v2)
                            case -9223372036854775808:
                                if (true || v4) {
                                    switch "" {
                                    case "":
                                        v61 := GenS10{F0: 2, F1: "\n\t", F2: (fi42 / ((781) | 1))}
                                        _ = v61
                                    case "hello":
                                        var v62 uint16 = 1
                                        _ = v62
                                    default:
                                        switch (-(fi42)) {
                                        case -1:
                                            v57 -= v2
                                        default:
                                            var v63 uint32 = 237
                                            _ = v63
                                        }
                                    }
                                } else {
                                    var v64 int16 = 343
                                    v1 += 1.5
                                    _ = v64
                                }
                            case 794:
                                var v65 int8 = -54
                                _ = v65
                            }
                        }
                        if false {
                            var v66 float32 = (((v2 + 0.0) * v2) + 1.5)
                            for fi67 := 0; fi67 < 6; fi67++ {
                                fmt.Println(v2, v57)
                                var v68 int64 = 42
                                if len(v20) > 0 {
                                    v20[0] = (^(fi42))
                                }
                                _ = fi67
                                _ = v68
                            }
                            _ = v66
                        } else {
                            var v69 string = ("" + "")
                            for fi70 := 0; fi70 < 7; fi70++ {
                                v3 -= (1e10 - float64(v57))
                                if false {
                                    v1 = -2.25
                                    if v4 {
                                        v20 = append(v20, fi42)
                                    } else {
                                        _ = len(v20)
                                        v71 := GenS4{F0: 65535, F1: 9223372036854775807, F2: 2}
                                        _ = v71
                                    }
                                    if true {
                                        var v72 int8 = 1
                                        _ = v72
                                    }
                                } else {
                                    var v73 uint = 4294967295
                                    v74 := GenS3{F0: -2147483648, F1: 1, F2: 167}
                                    v1 = float32(v3)
                                    _ = v73
                                    _ = v74
                                }
                                fmt.Println(fi70)
                                var v75 uint32 = 411
                                _ = fi70
                                _ = v75
                            }
                            _ = v69
                        }
                        _ = v57
                    } else {
                        if (-9223372036854775808 == -9223372036854775808) {
                            switch fi42 {
                            case 181:
                                switch (((("\n\t" + (("" + "a") + "a")) + ("hello" + "hello")) + "foo bar") + "hello") {
                                case "hello":
                                    fmt.Println(v3, v4, v1)
                                case "a":
                                    for fi76 := 0; fi76 < 11; fi76++ {
                                        v2 += (-2.25 + ((1e10 + 0.0) + (((-2.25 - v1) - (-2.25 + float32((v3 - (v3 - (v3 * -2.25)))))) * v2)))
                                        switch (fi76 + fi42) {
                                        case -346:
                                            switch 42 {
                                            case -232:
                                                switch -357 {
                                                case -2147483648:
                                                    var v77 int32 = 0
                                                    _ = v77
                                                case 357:
                                                    if v4 {
                                                        switch "" {
                                                        case "foo bar":
                                                            fmt.Println(fi76)
                                                        case "a":
                                                            fmt.Println(v2, v1)
                                                        }
                                                    }
                                                    if len(v20) >= 2 {
                                                        v20 = v20[1:len(v20)]
                                                    }
                                                    v78 := GenS1{F0: 8}
                                                    _ = v78
                                                case 32768:
                                                    var v79 uint64 = 1
                                                    _ = v79
                                                }
                                            case 42:
                                                if len(v20) >= 2 {
                                                    v20 = v20[1:len(v20)]
                                                }
                                            default:
                                                fmt.Println(v4, v3)
                                            }
                                        case -940:
                                            var v80 uint = 65535
                                            _ = v80
                                        }
                                        _ = fi76
                                    }
                                case "\n\t":
                                    for fi81 := 0; fi81 < 5; fi81++ {
                                        var v82 uint64 = 18446744073709551615
                                        var v83 int8 = 53
                                        _ = fi81
                                        _ = v82
                                        _ = v83
                                    }
                                default:
                                    if (42 != 313) {
                                        var v84 int64 = -758
                                        _ = len(v20)
                                        if len(v20) > 0 {
                                            v20[0] = 32767
                                        }
                                        _ = v84
                                    }
                                }
                            case -2147483648:
                                v2 -= 0.1
                            }
                            if ((("" + "hello") + "a") >= "a") {
                                for fi85 := 0; fi85 < 15; fi85++ {
                                    _ = len(v20)
                                    v3 += (v3 / ((v3 / ((v3 * v3) - v3)) / ((0.1 - (3.14159 + 3.14159)) / 1.0)))
                                    _ = fi85
                                }
                                for fi86 := 0; fi86 < 9; fi86++ {
                                    v2 += 1e10
                                    if v4 {
                                        v4 = (18446744073709551615 <= 526)
                                    } else {
                                        fmt.Println(fi42, fi86)
                                    }
                                    _ = fi86
                                }
                            } else {
                                switch (fi42 * (^(fi42))) {
                                case 2147483647:
                                    for fi87 := 0; fi87 < 11; fi87++ {
                                        var v88 string = "hello"
                                        if false {
                                            switch v88 {
                                            case "hello":
                                                v2 = (v2 + ((v2 - v2) * (v1 * v1)))
                                            default:
                                                var v89 int16 = 511
                                                _ = v89
                                            }
                                            var v90 uint16 = 42
                                            v91 := GenS11{F0: 0, F1: (1000 > 42)}
                                            if v4 {
                                                v3 = 0.0
                                                v90 += (^(v90))
                                                if (!(0 != 0)) {
                                                    var v92 float64 = float64(v1)
                                                    v88 = (((v88 + v88) + ("a" + "foo bar")) + v88)
                                                    v93 := GenS9{F0: 4294967295}
                                                    v88 = (v88 + v88)
                                                    _ = v92
                                                    _ = v93
                                                } else {
                                                    var v94 uint64 = 428
                                                    v95 := GenS1{F0: v94}
                                                    v96 := GenS12{F0: "\n\t", F1: v4, F2: -77}
                                                    _ = v94
                                                    _ = v95
                                                    _ = v96
                                                }
                                                v4 = true
                                                v20 = append(v20, fi87)
                                            }
                                            _ = v90
                                            _ = v91
                                        } else {
                                            switch v88 {
                                            case "\n\t":
                                                var v97 float64 = v3
                                                _ = v97
                                            case "a":
                                                if v4 {
                                                    v98 := GenS13{F0: -1, F1: 4294967295}
                                                    var v99 float64 = (v3 + v3)
                                                    _ = v98
                                                    _ = v99
                                                } else {
                                                    v20 = append(v20, fi87)
                                                    if (18446744073709551615 >= 330) {
                                                        if (1 < 1) {
                                                            fmt.Println(v3, v1, v3)
                                                        }
                                                    }
                                                }
                                            case "hello":
                                                if false {
                                                    switch fi42 {
                                                    case -27:
                                                        v100 := GenS13{F0: 1, F1: 131}
                                                        _ = v100
                                                    }
                                                    v4 = (65535 < 0)
                                                    if v4 {
                                                        v3 += (v3 - float64((1.5 + float32(float64(1.5)))))
                                                    } else {
                                                        var v101 uint32 = 0
                                                        v4 = (v4 || (886 <= -683))
                                                        var v102 int16 = 1
                                                        var v103 uint = 65535
                                                        _ = v101
                                                        _ = v102
                                                        _ = v103
                                                    }
                                                } else {
                                                    v4 = v4
                                                    v3 = -2.25
                                                    var v104 uint32 = 2147483647
                                                    if (162 != 468) {
                                                        var v105 int8 = -128
                                                        if true {
                                                            switch fi42 {
                                                            case -697:
                                                                v4 = false
                                                            }
                                                        }
                                                        v106 := GenS11{F0: 42, F1: ((v2 > 1e10) || (true && (!(1 == -32768))))}
                                                        fmt.Println(fi87, v105)
                                                        v88 = v88
                                                        fmt.Println(v3, v2, fi42)
                                                        _ = v105
                                                        _ = v106
                                                    } else {
                                                        fmt.Println(v2)
                                                        v107 := GenS14{F0: 128, F1: (65535 >= 835)}
                                                        v20 = append(v20, fi87)
                                                        _ = v107
                                                    }
                                                    if len(v20) >= 2 {
                                                        v20 = v20[1:len(v20)]
                                                    }
                                                    v1 -= v2
                                                    _ = v104
                                                }
                                            }
                                            if true {
                                                switch "a" {
                                                case "foo bar":
                                                    if (245 >= 10) {
                                                        _ = len(v20)
                                                        if len(v20) >= 2 {
                                                            v20 = v20[1:len(v20)]
                                                        }
                                                    }
                                                case "\n\t":
                                                    v1 -= v1
                                                case "":
                                                    v88 = ("a" + "foo bar")
                                                }
                                                if len(v20) > 0 {
                                                    v20[0] = (fi42 % (((fi42 & 127)) | 1))
                                                }
                                            }
                                            switch "" {
                                            case "hello":
                                                switch (fi87 >> ((1) & 63)) {
                                                case 493:
                                                    if len(v20) >= 2 {
                                                        v20 = v20[1:len(v20)]
                                                    }
                                                default:
                                                    v3 = v3
                                                }
                                            default:
                                                var v108 string = v88
                                                _ = v108
                                            }
                                            v109 := GenS15{F0: 4294967295}
                                            _ = v109
                                        }
                                        v110 := GenS2{F0: 127, F1: 0}
                                        _ = fi87
                                        _ = v110
                                        _ = v88
                                    }
                                case -674:
                                    var v111 float64 = v3
                                    _ = v111
                                }
                                if ("foo bar" > "\n\t") {
                                    for fi112 := 0; fi112 < 19; fi112++ {
                                        v113 := GenS16{F0: "\n\t", F1: 42, F2: 225}
                                        _ = fi112
                                        _ = v113
                                    }
                                    if (v4 && (v4 && true)) {
                                        var v114 uint16 = 0
                                        _ = v114
                                    }
                                    var v115 uint16 = 475
                                    _ = v115
                                }
                            }
                            if (65535 > 0) {
                                var v116 int = fi42
                                var v117 uint64 = 2
                                _ = v116
                                _ = v117
                            }
                        }
                        var v118 int32 = 622
                        if (v4 && v4) {
                            for fi119 := 0; fi119 < 2; fi119++ {
                                var v120 uint16 = 0
                                v120 = (v120 - v120)
                                _ = fi119
                                _ = v120
                            }
                            var v121 uint32 = 981
                            _ = v121
                        } else {
                            v4 = (!(float32(v3) >= v1))
                            if ((!v4) || ((!true) || true)) {
                                fmt.Println(v2, v118)
                                var v122 uint64 = 1
                                v4 = (v4 || (!(3.14159 >= (((v3 - float64((float32(v3) / 1.5))) * ((float64(v1) * (v3 - 3.14159)) + ((v3 + 1.5) * 3.14159))) * float64((float32(v3) * v1))))))
                                _ = v122
                            } else {
                                v20 = append(v20, fi42)
                                v4 = (v1 >= v2)
                            }
                            var v123 uint64 = 0
                            _ = v123
                        }
                        _ = v118
                    }
                case "":
                    v20 = append(v20, 1)
                default:
                    for fi124 := 0; fi124 < 19; fi124++ {
                        v4 = (v3 <= v3)
                        v125 := GenS10{F0: 42, F1: ("a" + ""), F2: (-(fi42))}
                        _ = fi124
                        _ = v125
                    }
                }
                var v126 bool = true
                v2 += v1
                _ = v126
            } else {
                if len(v20) > 0 {
                    v20[0] = 753
                }
            }
            _ = fi42
        }
    default:
        v127 := GenS13{F0: 1, F1: 731}
        _ = v127
    }
    _ = len(v20)
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    var v128 string = ("foo bar" + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v128:", v128)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v128
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
