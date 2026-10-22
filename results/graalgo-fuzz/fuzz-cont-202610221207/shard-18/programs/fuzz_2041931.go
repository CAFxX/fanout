package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 uint64; F1 uint; F2 string }
type GenS6 struct { F0 string; F1 string; F2 int64 }
type GenS7 struct { F0 int16; F1 int }
type GenS8 struct { F0 uint64; F1 int64 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int8; F1 int32; F2 uint32 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS15 struct { F0 int16; F1 int }
type GenS16 struct { F0 int64; F1 bool; F2 uint }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 int64 }
type GenS20 struct { F0 int; F1 bool; F2 uint64 }
type GenS21 struct { F0 int8; F1 bool }
type GenS22 struct { F0 uint16; F1 int; F2 uint32 }
type GenS23 struct { F0 uint8; F1 uint16 }
type GenS24 struct { F0 int64 }
type GenS25 struct { F0 int8; F1 int16 }
type GenS26 struct { F0 int16; F1 int16 }
func genFn0(p0 uint, p1 int, p2 uint32) uint16 {
    for fi4 := 0; fi4 < 11; fi4++ {
        for fi5 := 0; fi5 < 17; fi5++ {
            v6 := GenS1{F0: (^(p0))}
            p0 -= 42
            p2 -= (p2 % ((65535) | 1))
            p0 = p0
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 19; fi7++ {
            p2 = 2147483647
            var v8 uint32 = p2
            _ = fi7
            _ = v8
        }
        p1 += (fi4 + (p1 >> (((p0 % ((p0) | 1))) & 63)))
        _ = fi4
    }
    var v9 int16 = 523
    if false {
        var v10 float32 = float32(1e10)
        _ = v10
    } else {
        p0 = 18446744073709551615
        var v11 uint64 = 42
        for fi12 := 0; fi12 < 5; fi12++ {
            for fi13 := 0; fi13 < 1; fi13++ {
                v14 := GenS1{F0: p0}
                v15 := GenS1{F0: p0}
                p0 = (^(p0))
                _ = fi13
                _ = v14
                _ = v15
            }
            switch p1 {
            case 128:
                var v16 uint16 = 1
                _ = v16
            }
            _ = fi12
        }
        _ = v11
    }
    if true {
        var v17 uint32 = p2
        switch "" {
        case "\n\t":
            v9 -= -32768
        case "":
            for fi18 := 0; fi18 < 10; fi18++ {
                v19 := GenS1{F0: p0}
                v20 := GenS2{F0: -870, F1: 74, F2: 2}
                _ = fi18
                _ = v19
                _ = v20
            }
        }
        if true {
            for fi21 := 0; fi21 < 7; fi21++ {
                switch "hello" {
                case "a":
                    v9 -= (v9 * (v9 + 1000))
                case "":
                    var v22 uint8 = 1
                    _ = v22
                case "hello":
                    p1 += (p1 - -552)
                default:
                    p1 += (p1 % ((32768) | 1))
                }
                if (fi21 >= 320) {
                    v23 := []int{p1}
                    var v24 uint32 = p2
                    _ = v23
                    _ = v24
                }
                _ = fi21
            }
            p2 = v17
            var v25 uint16 = 42
            _ = v25
        } else {
            v26 := GenS3{F0: 456, F1: v9, F2: 765}
            v27 := GenS2{F0: 289, F1: 128, F2: v17}
            _ = v26
            _ = v27
        }
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 0
}
func genFn1(p0 int64, p1 int8, p2 uint16) uint64 {
    if ((3.14159 >= 1.5) && (36 >= 2)) {
        p2 = 465
        var v28 int = 128
        for fi29 := 0; fi29 < 11; fi29++ {
            for fi30 := 0; fi30 < 12; fi30++ {
                var v31 uint8 = 203
                v32 := GenS4{F0: (p0 <= 1), F1: int8(p0)}
                _ = fi30
                _ = v31
                _ = v32
            }
            var v33 int8 = (p1 & (^(p1)))
            p0 = p0
            _ = fi29
            _ = v33
        }
        _ = v28
    }
    v34 := []int{-2147483648, 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v34
    return 0
}

func main() {
    v1 := []string{"\n\t", (("hello" + "foo bar") + ("hello" + ("hello" + ("a" + "")))), "", ("foo bar" + "a")}
    v2 := make(map[int]string)
    var v3 uint16 = 0
    v35 := GenS4{F0: false, F1: -1}
    var v36 int16 = 1
    for fi37 := 0; fi37 < 10; fi37++ {
        var v38 int8 = 31
        _ = fi37
        _ = v38
    }
    v39 := GenS5{F0: 491, F1: 1, F2: "hello"}
    v40 := GenS6{F0: "foo bar", F1: "\n\t", F2: 42}
    v41 := GenS4{F0: (18446744073709551615 >= 18446744073709551615), F1: 1}
    _ = len(v2)
    for fi42 := 0; fi42 < 2; fi42++ {
        v3 = (v3 | 254)
        var v43 float64 = (1.5 * 1e10)
        _ = fi42
        _ = v43
    }
    var v44 uint16 = v3
    switch -9223372036854775808 {
    case -2147483648:
        v45 := GenS7{F0: 1, F1: 127}
        _ = v45
    case 128:
        if true {
            v36 = v36
            var v46 int32 = 2147483647
            _ = v46
        }
    default:
        if true {
            if false {
                for fi47 := 0; fi47 < 13; fi47++ {
                    v36 += v36
                    v44 = (^(v44))
                    switch (fi47 % (((fi47 + (fi47 % ((fi47) | 1)))) | 1)) {
                    case 127:
                        if true {
                            v3 = (v44 - v3)
                        } else {
                            var v48 int64 = -282
                            _ = v48
                        }
                    }
                    _ = fi47
                }
                v3 = (v44 + v44)
            } else {
                if false {
                    if len(v1) > 0 {
                        v1[0] = (((("foo bar" + "foo bar") + "a") + ("" + "\n\t")) + (("hello" + "hello") + "\n\t"))
                    }
                    var v49 int32 = 2147483647
                    var v50 float32 = 0.0
                    _ = v49
                    _ = v50
                }
            }
            var v51 uint32 = 2
            var v52 float32 = 1.5
            _ = v51
            _ = v52
        }
    }
    for fi53 := 0; fi53 < 6; fi53++ {
        if false {
            v36 += v36
            switch "a" {
            case "foo bar":
                if false {
                    switch "\n\t" {
                    case "foo bar":
                        var v54 int64 = 9223372036854775807
                        _ = v54
                    case "\n\t":
                        v55 := v2[(fi53 / ((fi53) | 1))]
                        _ = v55
                    case "":
                        fmt.Println(fi53, v3, v3)
                    default:
                        _ = len(v2)
                    }
                    for fi56 := 0; fi56 < 14; fi56++ {
                        _ = len(v2)
                        switch "" {
                        case "\n\t":
                            var v57 int32 = 1
                            _ = v57
                        case "":
                            var v58 int32 = 1
                            _ = v58
                        }
                        _ = fi56
                    }
                    if true {
                        _ = len(v2)
                        switch "hello" {
                        case "foo bar":
                            if (-911 <= 1) {
                                v59 := GenS8{F0: 192, F1: -690}
                                if (-111 >= (fi53 - fi53)) {
                                    switch fi53 {
                                    case 0:
                                        v36 += v36
                                    case 156:
                                        v60 := GenS5{F0: 304, F1: 132, F2: "a"}
                                        _ = v60
                                    case -9223372036854775808:
                                        for fi61 := 0; fi61 < 13; fi61++ {
                                            fmt.Println(v36)
                                            var v62 float32 = ((0.1 + -2.25) / 0.1)
                                            _ = fi61
                                            _ = v62
                                        }
                                    }
                                    if (215 > 1) {
                                        v1 = append(v1, ("hello" + ("a" + (("a" + "foo bar") + "hello"))))
                                    }
                                    _ = len(v2)
                                } else {
                                    for fi63 := 0; fi63 < 13; fi63++ {
                                        if (!true) {
                                            switch (-(fi63)) {
                                            case -870:
                                                v1 = append(v1, "hello")
                                            case -321:
                                                var v64 uint32 = 2147483647
                                                _ = v64
                                            case 9223372036854775807:
                                                v36 += 1
                                                if (255 > 1) {
                                                    v44 += (v3 | 0)
                                                    switch (("" + ("" + "")) + "\n\t") {
                                                    case "\n\t":
                                                        v36 -= (v36 % (((v36 & (v36 / (((-(v36))) | 1)))) | 1))
                                                    case "foo bar":
                                                        var v65 string = (("\n\t" + ("foo bar" + "")) + ("\n\t" + ""))
                                                        _ = v65
                                                    case "hello":
                                                        var v66 uint8 = 1
                                                        _ = v66
                                                    default:
                                                        if (1 < genFn1(431, 42, genFn0(1, 128, 0))) {
                                                            v67 := GenS9{F0: 0}
                                                            v68 := GenS3{F0: 0, F1: v36, F2: 923}
                                                            _ = v67
                                                            _ = v68
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(v44, fi53)
                                                }
                                                fmt.Println(v44, v36)
                                            }
                                            var v69 int64 = -1
                                            v70 := GenS10{F0: (v44 >> ((42) & 63))}
                                            _ = v69
                                            _ = v70
                                        } else {
                                            if true {
                                                if true {
                                                    if false {
                                                        var v71 uint = 42
                                                        v72 := GenS11{F0: (v3 - 42)}
                                                        switch (-(fi63)) {
                                                        case -9223372036854775808:
                                                            if false {
                                                                v36 = 1
                                                                v1 = append(v1, (("foo bar" + ("a" + "hello")) + ""))
                                                            } else {
                                                                _ = len(v1)
                                                                v73 := GenS12{F0: 127, F1: -655, F2: 353}
                                                                v44 = v3
                                                                _ = v73
                                                            }
                                                        }
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                    var v74 float32 = 3.14159
                                                    var v75 uint16 = 65535
                                                    _ = v74
                                                    _ = v75
                                                }
                                                if (!true) {
                                                    v36 -= (v36 / (((v36 / (((v36 * 342)) | 1))) | 1))
                                                }
                                            }
                                            v3 += v44
                                            if true {
                                                if (false && true) {
                                                    v76 := GenS7{F0: (v36 ^ 1000), F1: int(v44)}
                                                    v77 := GenS13{F0: 4294967295}
                                                    _ = v76
                                                    _ = v77
                                                }
                                            } else {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    switch "hello" {
                                                    case "a":
                                                        if true {
                                                            v78 := GenS14{F0: 2, F1: 520, F2: v36}
                                                            if false {
                                                                fmt.Println(v3)
                                                                v44 += genFn0(689, fi63, 77)
                                                                switch fi53 {
                                                                case 127:
                                                                    fmt.Println(v44, fi63)
                                                                }
                                                            } else {
                                                                v79 := GenS15{F0: v36, F1: fi53}
                                                                var v80 int16 = v36
                                                                if (2147483647 > 609) {
                                                                    switch (("" + "a") + (("\n\t" + "foo bar") + "")) {
                                                                    case "\n\t":
                                                                        var v81 int32 = -1
                                                                        _ = v81
                                                                    case "":
                                                                        if (float32(1e10) >= (float32((1e10 - 0.0)) + 0.1)) {
                                                                            v44 = (v44 - v3)
                                                                        }
                                                                    case "foo bar":
                                                                        var v82 uint64 = 2
                                                                        if false {
                                                                            v83 := GenS16{F0: 1, F1: ((0.1 < (-2.25 / 0.1)) || false), F2: 677}
                                                                            var v84 uint32 = 844
                                                                            switch "hello" {
                                                                            case "\n\t":
                                                                                var v85 int32 = 896
                                                                                _ = v85
                                                                            case "hello":
                                                                                if (0 <= 0) {
                                                                                    v86 := GenS17{F0: true}
                                                                                    var v87 int = 42
                                                                                    var v88 int16 = v80
                                                                                    _ = v86
                                                                                    _ = v87
                                                                                    _ = v88
                                                                                }
                                                                            }
                                                                            _ = v83
                                                                            _ = v84
                                                                        } else {
                                                                            var v89 string = ((("a" + "hello") + "\n\t") + ("a" + ""))
                                                                            if len(v1) > 0 {
                                                                                v1[0] = ((v89 + "a") + v89)
                                                                            }
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                v90 := v2[fi63]
                                                                                _ = v90
                                                                            case "":
                                                                                switch 9223372036854775807 {
                                                                                case 32767:
                                                                                    var v91 int32 = 219
                                                                                    _ = v91
                                                                                case -1:
                                                                                    _ = len(v1)
                                                                                case 887:
                                                                                    v92 := GenS8{F0: v82, F1: 591}
                                                                                    _ = v92
                                                                                default:
                                                                                    if true {
                                                                                        v82 = (v82 * genFn1(0, 127, (v44 & 266)))
                                                                                        switch (v89 + "\n\t") {
                                                                                        case "a":
                                                                                            if len(v1) >= 2 {
                                                                                                v1 = v1[1:len(v1)]
                                                                                            }
                                                                                        case "hello":
                                                                                            if true {
                                                                                                var v93 uint16 = v44
                                                                                                if true {
                                                                                                    v36 = v80
                                                                                                    if false {
                                                                                                        v36 += (^(v36))
                                                                                                        v94 := v2[(fi63 * fi63)]
                                                                                                        switch 127 {
                                                                                                        case 0:
                                                                                                            var v95 float32 = ((1e10 - 1e10) * ((0.1 * (1.5 - 1.5)) * 0.0))
                                                                                                            _ = v95
                                                                                                        }
                                                                                                        var v96 bool = true
                                                                                                        _ = v94
                                                                                                        _ = v96
                                                                                                    } else {
                                                                                                        var v97 float64 = 0.1
                                                                                                        v98 := GenS18{F0: 100}
                                                                                                        if (!false) {
                                                                                                            delete(v2, int(v82))
                                                                                                            if true {
                                                                                                                var v99 uint = 255
                                                                                                                _ = v99
                                                                                                            } else {
                                                                                                                var v100 uint = 42
                                                                                                                _ = v100
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v97
                                                                                                        _ = v98
                                                                                                    }
                                                                                                    var v101 int16 = (v80 >> ((821) & 63))
                                                                                                    _ = v101
                                                                                                } else {
                                                                                                    if (float32((0.0 / 3.14159)) == -2.25) {
                                                                                                        v2[fi63] = v89
                                                                                                        var v102 float32 = float32(0.0)
                                                                                                        v103 := GenS19{F0: -706}
                                                                                                        switch (-(fi63)) {
                                                                                                        case -1:
                                                                                                            var v104 bool = (true || (-1 < int32(v93)))
                                                                                                            _ = v104
                                                                                                        default:
                                                                                                            v105 := GenS13{F0: fi53}
                                                                                                            _ = v105
                                                                                                        }
                                                                                                        var v106 int16 = v80
                                                                                                        _ = v102
                                                                                                        _ = v103
                                                                                                        _ = v106
                                                                                                    }
                                                                                                    switch v89 {
                                                                                                    case "a":
                                                                                                        var v107 float32 = 3.14159
                                                                                                        _ = v107
                                                                                                    case "foo bar":
                                                                                                        v80 = -1
                                                                                                    }
                                                                                                }
                                                                                                var v108 float32 = -2.25
                                                                                                _ = v108
                                                                                                _ = v93
                                                                                            } else {
                                                                                                if true {
                                                                                                    var v109 float64 = 1e10
                                                                                                    _ = v109
                                                                                                } else {
                                                                                                    var v110 uint = 42
                                                                                                    v36 = -32768
                                                                                                    var v111 uint8 = 0
                                                                                                    fmt.Println(fi63, v89, fi63)
                                                                                                    _ = v110
                                                                                                    _ = v111
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            if (true && (!(4294967295 == uint(v82)))) {
                                                                                                if (true && true) {
                                                                                                    if (42 >= (^(v80))) {
                                                                                                        _ = len(v2)
                                                                                                        if len(v1) >= 2 {
                                                                                                            v1 = v1[1:len(v1)]
                                                                                                        }
                                                                                                    } else {
                                                                                                        if (!false) {
                                                                                                            fmt.Println(v36)
                                                                                                            if (3.14159 != float64((float32((0.0 + -2.25)) + 0.1))) {
                                                                                                                v1 = append(v1, ((v89 + (v89 + ((((v89 + ("hello" + v89)) + v89) + v89) + (v89 + "")))) + "\n\t"))
                                                                                                            }
                                                                                                        } else {
                                                                                                            if false {
                                                                                                                v89 = (v89 + "foo bar")
                                                                                                                v89 = ((v89 + v89) + v89)
                                                                                                            }
                                                                                                        }
                                                                                                        var v112 uint64 = (v82 / ((genFn1(42, int8(v82), (v44 / ((v44) | 1)))) | 1))
                                                                                                        if (273 <= uint(v82)) {
                                                                                                            if (false && false) {
                                                                                                                _ = len(v2)
                                                                                                                var v113 string = ("hello" + "hello")
                                                                                                                _ = v113
                                                                                                            } else {
                                                                                                                v114 := GenS17{F0: false}
                                                                                                                _ = v114
                                                                                                            }
                                                                                                            if (18446744073709551615 > 375) {
                                                                                                                var v115 float64 = 1e10
                                                                                                                _ = v115
                                                                                                            }
                                                                                                        } else {
                                                                                                            var v116 float64 = 1e10
                                                                                                            var v117 float32 = 0.0
                                                                                                            v118 := GenS20{F0: -276, F1: (2147483647 < 2147483647), F2: (v112 | (v112 / ((0) | 1)))}
                                                                                                            if true {
                                                                                                                var v119 int64 = -732
                                                                                                                if true {
                                                                                                                    if false {
                                                                                                                        switch (((v89 + ("" + v89)) + "\n\t") + (v89 + "hello")) {
                                                                                                                        case "":
                                                                                                                            v120 := GenS21{F0: 127, F1: true}
                                                                                                                            _ = v120
                                                                                                                        default:
                                                                                                                            fmt.Println(v112, v3)
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v112 = genFn1((v119 / ((v119) | 1)), 0, (v44 ^ (v44 ^ 706)))
                                                                                                                    v121 := GenS22{F0: v44, F1: -9223372036854775808, F2: 2147483648}
                                                                                                                    _ = v121
                                                                                                                } else {
                                                                                                                    fmt.Println(v116, fi53, v80)
                                                                                                                }
                                                                                                                v117 += v117
                                                                                                                _ = v119
                                                                                                            }
                                                                                                            _ = v116
                                                                                                            _ = v117
                                                                                                            _ = v118
                                                                                                        }
                                                                                                        _ = v112
                                                                                                    }
                                                                                                    if true {
                                                                                                        delete(v2, (fi63 + (^(fi53))))
                                                                                                        v80 = v80
                                                                                                    } else {
                                                                                                        v89 = v89
                                                                                                        if (uint(v44) < 4294967295) {
                                                                                                            v80 -= v36
                                                                                                        } else {
                                                                                                            v3 = 1
                                                                                                            v44 -= (v44 | (v44 << ((960) & 63)))
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                v80 -= (v36 / (((v36 | (^(v36)))) | 1))
                                                                                                v122 := GenS18{F0: -8}
                                                                                                _ = v122
                                                                                            } else {
                                                                                                var v123 int32 = 0
                                                                                                _ = v123
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case "hello":
                                                                                var v124 float64 = (1e10 * 1.5)
                                                                                _ = v124
                                                                            default:
                                                                                if true {
                                                                                    var v125 int = (fi53 | -152)
                                                                                    v126 := GenS7{F0: -1, F1: (-(fi63))}
                                                                                    _ = v125
                                                                                    _ = v126
                                                                                }
                                                                            }
                                                                            _ = v89
                                                                        }
                                                                        v80 += v80
                                                                        _ = v82
                                                                    }
                                                                    if len(v1) > 0 {
                                                                        v1[0] = "foo bar"
                                                                    }
                                                                }
                                                                _ = v79
                                                                _ = v80
                                                            }
                                                            v127 := GenS23{F0: 89, F1: genFn0(42, 436, 841)}
                                                            _ = v127
                                                            _ = v78
                                                        }
                                                    }
                                                case "":
                                                    v36 = v36
                                                }
                                                v128 := GenS11{F0: v3}
                                                _ = v128
                                            }
                                            switch "a" {
                                            case "":
                                                v129 := GenS9{F0: 65535}
                                                _ = v129
                                            case "hello":
                                                v2[(^(fi53))] = ("foo bar" + "")
                                            }
                                            v2[fi63] = "hello"
                                        }
                                        v2[(fi63 ^ (fi63 * (fi53 / ((fi63) | 1))))] = (("foo bar" + "\n\t") + "foo bar")
                                        var v130 uint = 102
                                        _ = fi63
                                        _ = v130
                                    }
                                    var v131 int = 9223372036854775807
                                    for fi132 := 0; fi132 < 4; fi132++ {
                                        fmt.Println(v131)
                                        if false {
                                            switch (fi132 ^ (v131 & fi53)) {
                                            case 0:
                                                var v133 int32 = -2147483648
                                                _ = v133
                                            case -190:
                                                delete(v2, (fi132 | (fi132 - 2147483647)))
                                            default:
                                                var v134 uint32 = 159
                                                _ = v134
                                            }
                                            fmt.Println(fi53, fi132)
                                            if len(v1) > 0 {
                                                v1[0] = "hello"
                                            }
                                        } else {
                                            v44 -= 1
                                            if true {
                                                if false {
                                                    v36 = (^(v36))
                                                    v131 += (v131 | fi53)
                                                    v135 := GenS24{F0: 126}
                                                    _ = v135
                                                }
                                                if true {
                                                    if (false && true) {
                                                        v136 := GenS25{F0: 1, F1: 1}
                                                        delete(v2, (v131 >> ((284) & 63)))
                                                        _ = v136
                                                    }
                                                    v44 = v3
                                                    var v137 int16 = 844
                                                    _ = v137
                                                } else {
                                                    v36 = 1
                                                    switch 42 {
                                                    case 0:
                                                        v138 := v2[(fi53 ^ fi53)]
                                                        _ = v138
                                                    }
                                                    if ((fi132 >= (v131 >> ((42) & 63))) || true) {
                                                        var v139 uint16 = (v3 % (((v3 / ((genFn0(18446744073709551615, fi132, 42)) | 1))) | 1))
                                                        v139 -= (v3 % ((v139) | 1))
                                                        _ = v139
                                                    }
                                                }
                                                v3 -= genFn0(65535, fi132, 0)
                                            } else {
                                                v3 -= 32767
                                                v140 := GenS26{F0: (v36 - v36), F1: v36}
                                                delete(v2, (v131 ^ v131))
                                                _ = v140
                                            }
                                        }
                                        fmt.Println(v3)
                                        _ = fi132
                                    }
                                    _ = v131
                                }
                                if false {
                                    if false {
                                        v44 -= (v44 ^ genFn0(uint(v3), (fi53 >> ((883) & 63)), 696))
                                    } else {
                                        var v141 int16 = (^(v36))
                                        _ = v141
                                    }
                                } else {
                                    fmt.Println(v44, v3, fi53)
                                    if (true && true) {
                                        fmt.Println(v3)
                                    } else {
                                        var v142 int64 = -777
                                        v142 += (v142 | (v142 - (v142 - (v142 / ((1) | 1)))))
                                        v44 = (v44 << ((2) & 63))
                                        _ = v142
                                    }
                                }
                                _ = v59
                            }
                        }
                        v3 = (^(v44))
                    } else {
                        for fi143 := 0; fi143 < 5; fi143++ {
                            fmt.Println(fi53, v3, v36)
                            v44 -= 0
                            _ = fi143
                        }
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                    }
                } else {
                    if len(v1) > 0 {
                        v1[0] = ("" + "foo bar")
                    }
                    var v144 uint = uint(v3)
                    _ = v144
                }
            }
            v44 += (v3 - 949)
        } else {
            for fi145 := 0; fi145 < 17; fi145++ {
                var v146 uint32 = 65535
                v147 := v2[(fi53 & (fi53 - 317))]
                fmt.Println(fi145, fi53)
                _ = fi145
                _ = v146
                _ = v147
            }
        }
        _ = fi53
    }
    for fi148 := 0; fi148 < 3; fi148++ {
        var v149 int64 = 9223372036854775807
        var v150 int32 = 42
        _ = fi148
        _ = v149
        _ = v150
    }
    fmt.Println(v3, v3)
    var v151 uint16 = genFn0(42, -9223372036854775808, 488)
    var v152 float32 = 0.1
    var v153 int16 = v36
    var v154 uint16 = v3
    for fi155 := 0; fi155 < 8; fi155++ {
        switch "" {
        case "hello":
            v1 = append(v1, "hello")
        }
        _ = fi155
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v151:", v151)
    fmt.Println("v152:", v152)
    fmt.Println("v153:", v153)
    fmt.Println("v154:", v154)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v151
    _ = v152
    _ = v153
    _ = v154
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v40
    _ = v41
    _ = v44
}
