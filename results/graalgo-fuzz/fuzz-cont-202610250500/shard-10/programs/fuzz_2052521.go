package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 string }
type GenS2 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int; F1 int32; F2 int64 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int32; F1 uint32 }
func genFn0() uint64 {
    var v6 uint8 = 1
    for fi7 := 0; fi7 < 17; fi7++ {
        if true {
            var v8 int32 = 42
            _ = v8
        } else {
            var v9 bool = (!true)
            if v9 {
                if false {
                    v6 = 1
                    for fi10 := 0; fi10 < 15; fi10++ {
                        fmt.Println(v9, v6, v6)
                        v9 = v9
                        _ = fi10
                    }
                }
                var v11 int16 = -1
                _ = v11
            } else {
                if false {
                    v12 := GenS1{F0: 297, F1: 2147483647, F2: ((("a" + "foo bar") + ("" + "\n\t")) + "\n\t")}
                    for fi13 := 0; fi13 < 6; fi13++ {
                        fmt.Println(fi7, v9, v6)
                        fmt.Println(v9, v6, fi7)
                        _ = fi13
                    }
                    _ = v12
                }
                var v14 uint8 = v6
                _ = v14
            }
            v15 := []int{(fi7 >> ((1) & 63)), 994, -173}
            _ = v15
            _ = v9
        }
        _ = fi7
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        if false {
            v17 := GenS1{F0: 42, F1: 2, F2: "hello"}
            _ = v17
        }
        v6 = 57
        _ = fi16
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        if true {
            v6 = 76
            switch fi18 {
            case 32767:
                var v19 int = 2147483647
                _ = v19
            default:
                for fi20 := 0; fi20 < 13; fi20++ {
                    if false {
                        var v21 string = ("a" + "hello")
                        _ = v21
                    }
                    var v22 uint16 = 2
                    _ = fi20
                    _ = v22
                }
            }
        } else {
            v23 := []int{}
            var v24 int = fi18
            _ = v23
            _ = v24
        }
        _ = fi18
    }
    var v25 uint = 1
    _ = v25
    _ = v6
    return 2
}

func main() {
    v1 := []string{"foo bar", "foo bar", "a"}
    var v2 float64 = float64(0.0)
    var v3 uint = 824
    v4 := make(map[int]string)
    var v5 float32 = 0.0
    if (2 != 2147483648) {
        _ = len(v4)
        if (((!(float32(v2) == (v5 - v5))) && (31 == -77)) || (("a" + "hello") > "hello")) {
            var v26 bool = (!false)
            if (!(396 >= 42)) {
                switch "a" {
                case "a":
                    fmt.Println(v2, v26, v3)
                case "hello":
                    for fi27 := 0; fi27 < 10; fi27++ {
                        v28 := GenS1{F0: -1, F1: 42, F2: "a"}
                        v29 := GenS1{F0: 1, F1: 2147483647, F2: ("\n\t" + "foo bar")}
                        v30 := GenS1{F0: -2147483648, F1: 2147483648, F2: ""}
                        _ = fi27
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                default:
                    v3 = 18446744073709551615
                }
            }
            var v31 uint64 = 18446744073709551615
            _ = v26
            _ = v31
        } else {
            fmt.Println(v5, v2, v3)
            v32 := GenS1{F0: -115, F1: 42, F2: ""}
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
            _ = v32
        }
        var v33 uint16 = 983
        _ = v33
    }
    switch ((("a" + ("foo bar" + ("foo bar" + (("\n\t" + "\n\t") + ("hello" + "a"))))) + ("foo bar" + "\n\t")) + "") {
    case "a":
        v5 += 1.5
    case "foo bar":
        v3 = (v3 << (((v3 & (v3 | (v3 * 0)))) & 63))
    case "":
        switch ("foo bar" + "\n\t") {
        case "foo bar":
            if (-100 <= 42) {
                var v34 float64 = (0.1 - v2)
                var v35 float64 = v34
                var v36 int8 = -89
                _ = v34
                _ = v35
                _ = v36
            } else {
                switch "a" {
                case "a":
                    v37 := GenS2{F0: 25, F1: -542, F2: 158}
                    _ = v37
                }
                v5 -= (v5 / 0.1)
                for fi38 := 0; fi38 < 4; fi38++ {
                    var v39 uint64 = 966
                    v39 = (^(v39))
                    if (121 >= 42) {
                        v5 -= v5
                    }
                    _ = fi38
                    _ = v39
                }
            }
        case "\n\t":
            fmt.Println(v5)
        default:
            v2 = 3.14159
        }
    default:
        fmt.Println(v5)
    }
    v5 -= 0.0
    fmt.Println(v3, v3)
    for fi40 := 0; fi40 < 11; fi40++ {
        fmt.Println(fi40, fi40)
        var v41 uint32 = 0
        fmt.Println(fi40)
        _ = fi40
        _ = v41
    }
    if true {
        switch -1 {
        case -112:
            for fi42 := 0; fi42 < 19; fi42++ {
                switch "a" {
                case "hello":
                    if (42 >= 127) {
                        var v43 float64 = 0.1
                        var v44 uint = (v3 + v3)
                        var v45 int16 = 1
                        _ = v43
                        _ = v44
                        _ = v45
                    } else {
                        v2 = 3.14159
                        for fi46 := 0; fi46 < 3; fi46++ {
                            switch 4294967295 {
                            case 128:
                                v47 := v4[(^(fi42))]
                                _ = v47
                            case -813:
                                v3 += (v3 - v3)
                            case -1:
                                var v48 string = "\n\t"
                                _ = v48
                            default:
                                v2 = v2
                            }
                            v49 := GenS2{F0: 255, F1: -9223372036854775808, F2: 1}
                            v3 = v3
                            if (-32768 == 0) {
                                switch "hello" {
                                case "hello":
                                    fmt.Println(v3)
                                case "foo bar":
                                    var v50 uint8 = 127
                                    _ = v50
                                case "\n\t":
                                    var v51 uint8 = 2
                                    _ = v51
                                }
                            }
                            _ = fi46
                            _ = v49
                        }
                        v52 := GenS1{F0: 0, F1: 0, F2: "hello"}
                        _ = v52
                    }
                case "foo bar":
                    if false {
                        if true {
                            var v53 uint8 = 255
                            for fi54 := 0; fi54 < 1; fi54++ {
                                fmt.Println(v53, v2, v3)
                                if true {
                                    var v55 float32 = v5
                                    v2 -= ((v2 / 1.0) - 0.1)
                                    _ = v55
                                }
                                fmt.Println(v5, v53, v5)
                                _ = fi54
                            }
                            switch 42 {
                            case -802:
                                var v56 bool = true
                                _ = v56
                            }
                            _ = v53
                        } else {
                            fmt.Println(v2, fi42)
                        }
                        v57 := GenS3{F0: -32768}
                        _ = v57
                    } else {
                        switch fi42 {
                        case -9223372036854775808:
                            _ = len(v1)
                        case -2147483648:
                            if true {
                                var v58 int64 = 1
                                _ = v58
                            } else {
                                for fi59 := 0; fi59 < 17; fi59++ {
                                    var v60 float32 = 1.5
                                    v5 += float32((0.0 / v2))
                                    _ = fi59
                                    _ = v60
                                }
                                if true {
                                    switch ("a" + "hello") {
                                    case "":
                                        switch fi42 {
                                        case 32767:
                                            if (287 == 174) {
                                                switch (("a" + "foo bar") + ((("foo bar" + ("" + "hello")) + "hello") + "")) {
                                                case "\n\t":
                                                    for fi61 := 0; fi61 < 19; fi61++ {
                                                        v2 -= (((((0.1 + float64(((v5 / v5) + v5))) - 0.1) - 1e10) - v2) / v2)
                                                        if (!(2 <= 65535)) {
                                                            var v62 int = fi61
                                                            _ = v62
                                                        } else {
                                                            v63 := v4[fi61]
                                                            _ = v63
                                                        }
                                                        _ = fi61
                                                    }
                                                case "a":
                                                    switch fi42 {
                                                    case 0:
                                                        for fi64 := 0; fi64 < 2; fi64++ {
                                                            _ = len(v4)
                                                            if true {
                                                                if (!true) {
                                                                    if true {
                                                                        var v65 uint = (^(v3))
                                                                        if false {
                                                                            _ = len(v4)
                                                                            if false {
                                                                                if true {
                                                                                    v3 = 0
                                                                                    fmt.Println(v65, fi64, v65)
                                                                                } else {
                                                                                    var v66 uint16 = 32767
                                                                                    var v67 float64 = v2
                                                                                    _ = v66
                                                                                    _ = v67
                                                                                }
                                                                            } else {
                                                                                v65 = v3
                                                                                v68 := v4[(fi42 << (((v65 << (((v65 % ((v3) | 1))) & 63))) & 63))]
                                                                                _ = v68
                                                                            }
                                                                        }
                                                                        _ = v65
                                                                    }
                                                                    if (false || true) {
                                                                        v69 := v4[(fi64 >> ((v3) & 63))]
                                                                        v70 := GenS4{F0: 1}
                                                                        fmt.Println(fi42, v69)
                                                                        _ = v69
                                                                        _ = v70
                                                                    } else {
                                                                        if true {
                                                                            if (42 < 503) {
                                                                                fmt.Println(fi64, fi42)
                                                                                fmt.Println(fi64)
                                                                                _ = len(v1)
                                                                            }
                                                                        } else {
                                                                            if true {
                                                                                v71 := GenS4{F0: 0}
                                                                                v72 := GenS5{F0: 65535, F1: -1, F2: 126}
                                                                                var v73 uint64 = 42
                                                                                if false {
                                                                                    switch 467 {
                                                                                    case 4294967295:
                                                                                        v74 := GenS1{F0: -1, F1: 2, F2: "hello"}
                                                                                        _ = v74
                                                                                    case -1:
                                                                                        switch ("\n\t" + "") {
                                                                                        case "\n\t":
                                                                                            switch "" {
                                                                                            case "\n\t":
                                                                                                v5 -= -2.25
                                                                                            case "hello":
                                                                                                v3 -= v3
                                                                                            default:
                                                                                                v3 -= v3
                                                                                            }
                                                                                        }
                                                                                    case -9223372036854775808:
                                                                                        v3 += v3
                                                                                    default:
                                                                                        v75 := GenS3{F0: -1}
                                                                                        _ = v75
                                                                                    }
                                                                                    v2 = 1e10
                                                                                    if (!false) {
                                                                                        var v76 string = "hello"
                                                                                        v77 := GenS6{F0: 32767, F1: 1}
                                                                                        v73 = (v73 >> ((65535) & 63))
                                                                                        var v78 uint8 = 1
                                                                                        _ = v76
                                                                                        _ = v77
                                                                                        _ = v78
                                                                                    } else {
                                                                                        if (true || true) {
                                                                                            v4[fi42] = ""
                                                                                        }
                                                                                        v79 := GenS2{F0: 65535, F1: 0, F2: 2}
                                                                                        var v80 int8 = -103
                                                                                        _ = v79
                                                                                        _ = v80
                                                                                    }
                                                                                    switch "a" {
                                                                                    case "foo bar":
                                                                                        if (!(true && false)) {
                                                                                            fmt.Println(v5, fi42, v2)
                                                                                            switch "" {
                                                                                            case "foo bar":
                                                                                                v81 := GenS7{F0: (^(v73))}
                                                                                                _ = v81
                                                                                            case "hello":
                                                                                                v82 := GenS8{F0: 698}
                                                                                                _ = v82
                                                                                            default:
                                                                                                v3 -= (v3 + (v3 * v3))
                                                                                            }
                                                                                        }
                                                                                    case "a":
                                                                                        switch fi64 {
                                                                                        case 32768:
                                                                                            _ = len(v1)
                                                                                        }
                                                                                    case "hello":
                                                                                        var v83 int16 = -32768
                                                                                        _ = v83
                                                                                    default:
                                                                                        var v84 uint16 = 121
                                                                                        _ = v84
                                                                                    }
                                                                                } else {
                                                                                    v3 = v3
                                                                                }
                                                                                _ = v71
                                                                                _ = v72
                                                                                _ = v73
                                                                            }
                                                                            if false {
                                                                                var v85 bool = ((v3 - 0) != v3)
                                                                                if v85 {
                                                                                    v4[(fi64 ^ (fi64 >> (((v3 & v3)) & 63)))] = "hello"
                                                                                } else {
                                                                                    switch 18 {
                                                                                    case 0:
                                                                                        if false {
                                                                                            if (false || false) {
                                                                                                v5 = (v5 * v5)
                                                                                                v2 = ((1e10 - v2) / (1.5 / 1e10))
                                                                                            } else {
                                                                                                switch "" {
                                                                                                case "foo bar":
                                                                                                    v85 = true
                                                                                                default:
                                                                                                    delete(v4, 702)
                                                                                                }
                                                                                            }
                                                                                            switch "a" {
                                                                                            case "\n\t":
                                                                                                v4[fi42] = "\n\t"
                                                                                            case "foo bar":
                                                                                                var v86 bool = true
                                                                                                _ = v86
                                                                                            }
                                                                                            var v87 int32 = 1
                                                                                            _ = v87
                                                                                        }
                                                                                    default:
                                                                                        v88 := GenS9{F0: 2}
                                                                                        _ = v88
                                                                                    }
                                                                                    v2 = v2
                                                                                }
                                                                                v5 -= v5
                                                                                _ = v85
                                                                            } else {
                                                                                switch "foo bar" {
                                                                                case "foo bar":
                                                                                    switch "" {
                                                                                    case "a":
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = ("\n\t" + ("foo bar" + "\n\t"))
                                                                                        }
                                                                                    case "hello":
                                                                                        v89 := v4[(fi64 & -639)]
                                                                                        _ = v89
                                                                                    default:
                                                                                        if false {
                                                                                            v90 := GenS7{F0: 0}
                                                                                            if false {
                                                                                                v91 := GenS7{F0: 111}
                                                                                                switch fi64 {
                                                                                                case 4294967295:
                                                                                                    v2 = v2
                                                                                                case 1:
                                                                                                    v3 = v3
                                                                                                }
                                                                                                fmt.Println(v5)
                                                                                                _ = v91
                                                                                            }
                                                                                            v3 += (v3 / ((v3) | 1))
                                                                                            _ = v90
                                                                                        }
                                                                                    }
                                                                                case "":
                                                                                    var v92 uint32 = 42
                                                                                    _ = v92
                                                                                case "a":
                                                                                    fmt.Println(fi42)
                                                                                default:
                                                                                    v93 := GenS10{F0: 128, F1: 1, F2: 1}
                                                                                    _ = v93
                                                                                }
                                                                                if false {
                                                                                    fmt.Println(v3)
                                                                                }
                                                                            }
                                                                            v94 := GenS6{F0: 1, F1: -9223372036854775808}
                                                                            _ = v94
                                                                        }
                                                                        fmt.Println(fi64)
                                                                        switch "hello" {
                                                                        case "a":
                                                                            v4[fi64] = ("\n\t" + (("hello" + "\n\t") + "\n\t"))
                                                                        default:
                                                                            if (4294967295 != 2147483648) {
                                                                                v5 = float32(float64((float32(v2) / v5)))
                                                                                switch (fi64 >> (((v3 / ((332) | 1))) & 63)) {
                                                                                case 42:
                                                                                    var v95 int64 = -9223372036854775808
                                                                                    _ = v95
                                                                                }
                                                                                switch (fi42 * (fi42 | fi42)) {
                                                                                case -2147483648:
                                                                                    v96 := GenS8{F0: 0}
                                                                                    _ = v96
                                                                                default:
                                                                                    var v97 uint8 = 0
                                                                                    _ = v97
                                                                                }
                                                                            }
                                                                        }
                                                                        var v98 float64 = 1e10
                                                                        if false {
                                                                            var v99 int16 = 0
                                                                            if true {
                                                                                var v100 uint8 = 127
                                                                                var v101 uint32 = 709
                                                                                _ = v100
                                                                                _ = v101
                                                                            }
                                                                            _ = v99
                                                                        } else {
                                                                            fmt.Println(v5, fi42, fi64)
                                                                            var v102 uint8 = 128
                                                                            _ = v102
                                                                        }
                                                                        var v103 string = "hello"
                                                                        v104 := GenS11{F0: genFn0()}
                                                                        _ = len(v4)
                                                                        fmt.Println(v2)
                                                                        _ = v103
                                                                        _ = v104
                                                                        _ = v98
                                                                    }
                                                                    v5 += v5
                                                                    v2 = v2
                                                                } else {
                                                                    fmt.Println(v3)
                                                                }
                                                                v3 = v3
                                                                fmt.Println(v3, fi64, v5)
                                                                var v105 int16 = 42
                                                                var v106 float32 = 1.5
                                                                _ = v105
                                                                _ = v106
                                                            }
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            v5 = 1.5
                                                            switch -433 {
                                                            case 843:
                                                                switch -864 {
                                                                case 2147483647:
                                                                    fmt.Println(v3)
                                                                case 471:
                                                                    var v107 int64 = -230
                                                                    _ = v107
                                                                case 4294967295:
                                                                    var v108 uint8 = 42
                                                                    _ = v108
                                                                }
                                                            }
                                                            _ = fi64
                                                        }
                                                    case -2147483648:
                                                        fmt.Println(fi42)
                                                    default:
                                                        v5 = 1e10
                                                    }
                                                default:
                                                    var v109 int32 = 2147483647
                                                    _ = v109
                                                }
                                                v110 := GenS8{F0: 42}
                                                for fi111 := 0; fi111 < 17; fi111++ {
                                                    if ((v2 + v2) > v2) {
                                                        var v112 bool = ((!(-9223372036854775808 <= 0)) && (541 >= (fi42 - fi111)))
                                                        v112 = (v5 <= v5)
                                                        v4[(fi42 << (((^(v3))) & 63))] = ("hello" + ("hello" + ("a" + "")))
                                                        _ = v112
                                                    } else {
                                                        var v113 int64 = -1
                                                        v5 = 0.0
                                                        _ = v113
                                                    }
                                                    if (true && ((420 >= 298) || true)) {
                                                        var v114 string = ("" + ("hello" + "a"))
                                                        fmt.Println(fi111, v114)
                                                        v5 = ((v5 - (3.14159 / v5)) / v5)
                                                        v114 = "\n\t"
                                                        _ = v114
                                                    }
                                                    if true {
                                                        var v115 int16 = 521
                                                        fmt.Println(v5, v2)
                                                        _ = v115
                                                    }
                                                    _ = fi111
                                                }
                                                _ = v110
                                            } else {
                                                var v116 int16 = 1
                                                v3 = 0
                                                _ = v116
                                            }
                                        case 0:
                                            var v117 int = fi42
                                            _ = v117
                                        case -1:
                                            for fi118 := 0; fi118 < 12; fi118++ {
                                                if (0 != -1) {
                                                    if ((0 == 2147483647) && (false || (-1 == 42))) {
                                                        var v119 float64 = ((1.5 * 0.1) / v2)
                                                        v5 += ((v5 / 1e10) + v5)
                                                        _ = v119
                                                    }
                                                }
                                                v5 -= (float32(3.14159) + 1.5)
                                                var v120 int32 = 2147483647
                                                var v121 float32 = (v5 - v5)
                                                var v122 int64 = -1
                                                _ = fi118
                                                _ = v120
                                                _ = v121
                                                _ = v122
                                            }
                                        }
                                    case "hello":
                                        v3 = (v3 ^ v3)
                                    case "a":
                                        v3 += (v3 * (v3 % ((2) | 1)))
                                    default:
                                        if ((v5 - (3.14159 + (v5 + v5))) != v5) {
                                            v3 = 18446744073709551615
                                        } else {
                                            v5 = 1.5
                                            v123 := v4[(fi42 | 1)]
                                            _ = v123
                                        }
                                    }
                                } else {
                                    if false {
                                        v2 = v2
                                        var v124 uint32 = 1
                                        _ = v124
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            }
                        }
                    }
                }
                _ = fi42
            }
        case -9223372036854775808:
            for fi125 := 0; fi125 < 10; fi125++ {
                var v126 uint16 = 991
                if false {
                    fmt.Println(v5, fi125)
                    v5 = (-2.25 + (((v5 - (0.0 / 1.0)) + v5) * (v5 * v5)))
                } else {
                    var v127 int64 = 9223372036854775807
                    _ = v127
                }
                v5 -= v5
                _ = fi125
                _ = v126
            }
        case 1:
            if (65535 < 220) {
                v128 := GenS5{F0: 998, F1: 0, F2: 1}
                v129 := GenS3{F0: -1}
                fmt.Println(v5, v2)
                _ = v128
                _ = v129
            } else {
                switch "a" {
                case "a":
                    var v130 uint8 = 42
                    _ = v130
                case "foo bar":
                    var v131 int8 = 42
                    _ = v131
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        default:
            v5 -= (v5 - 1e10)
        }
        v5 -= float32(1.5)
        var v132 float32 = ((v5 + (v5 + float32(v2))) / v5)
        _ = v132
    } else {
        v5 = (v5 * v5)
        v5 += 1e10
        v133 := v4[9223372036854775807]
        _ = v133
    }
    var v134 int = 17
    v3 += 1
    for fi135 := 0; fi135 < 5; fi135++ {
        v136 := GenS12{F0: -2147483648, F1: 1}
        fmt.Println(v2, v134)
        _ = fi135
        _ = v136
    }
    _ = len(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v134:", v134)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v134
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
