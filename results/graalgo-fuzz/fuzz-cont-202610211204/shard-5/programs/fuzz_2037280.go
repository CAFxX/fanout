package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int; F1 uint64; F2 bool }
type GenS3 struct { F0 uint32; F1 int32; F2 bool }
type GenS4 struct { F0 int8; F1 int16; F2 int }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 bool; F1 uint64; F2 int8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int; F1 uint32; F2 int }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int16; F1 uint32 }
type GenS11 struct { F0 uint; F1 uint16; F2 int16 }
type GenS12 struct { F0 uint16; F1 string; F2 int16 }
type GenS13 struct { F0 uint; F1 string }
type GenS14 struct { F0 uint8; F1 string }
type GenS15 struct { F0 int8; F1 string }
type GenS16 struct { F0 int8 }

func main() {
    var v1 uint64 = 272
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    v4 := []string{"\n\t", (("hello" + "") + (("a" + "") + "foo bar"))}
    var v5 int64 = 0
    var v6 float32 = 1.5
    switch 32768 {
    case 4294967295:
        var v7 uint64 = 4294967295
        _ = v7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        if len(v4) > 0 {
            v4[0] = (("\n\t" + "a") + "hello")
        }
        v9 := GenS1{F0: (^(v1)), F1: 567}
        if false {
            v6 -= (v6 / v6)
        } else {
            if (0 > 956) {
                if false {
                    var v10 int64 = 9223372036854775807
                    v11 := GenS2{F0: fi8, F1: v1, F2: false}
                    _ = v10
                    _ = v11
                }
                var v12 int32 = -506
                _ = v12
            } else {
                v5 = (v5 % (((v5 / (((^(v5))) | 1))) | 1))
            }
        }
        _ = fi8
        _ = v9
    }
    var v13 uint16 = 993
    for fi14 := 0; fi14 < 16; fi14++ {
        v13 = v13
        fmt.Println(v1)
        if false {
            v15 := GenS1{F0: (^(v1)), F1: 0}
            v13 = 42
            if true {
                var v16 int = fi14
                _ = v16
            }
            _ = v15
        }
        _ = fi14
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 uint64 = 18446744073709551615
        delete(v2, (fi17 + (fi17 - 4294967295)))
        if (int8(v1) != 0) {
            delete(v3, 0)
        }
        _ = fi17
        _ = v18
    }
    v6 = ((-2.25 / ((1.5 + v6) * v6)) + (v6 - v6))
    _ = len(v2)
    v4 = append(v4, "hello")
    var v19 int32 = -1
    fmt.Println(v5)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v20 uint = 0
    fmt.Println(v5)
    v6 += v6
    for fi21 := 0; fi21 < 14; fi21++ {
        v19 -= v19
        fmt.Println(fi21)
        if (false || (((0.0 + 1e10) - (float64(-2.25) / (1.5 - 3.14159))) < (1e10 - 0.0))) {
            if (!("a" >= "\n\t")) {
                switch "\n\t" {
                case "":
                    var v22 int32 = (v19 | 42)
                    _ = v22
                case "\n\t":
                    for fi23 := 0; fi23 < 13; fi23++ {
                        var v24 uint64 = 18446744073709551615
                        switch "" {
                        case "foo bar":
                            if true {
                                switch "" {
                                case "hello":
                                    v1 -= 782
                                case "":
                                    if ((v6 != v6) || (!true)) {
                                        if true {
                                            v6 = 0.0
                                            var v25 int32 = (^(v19))
                                            fmt.Println(v24, v19, v19)
                                            var v26 int8 = 42
                                            v4 = append(v4, ("foo bar" + "\n\t"))
                                            v19 = v19
                                            fmt.Println(v5, v5)
                                            _ = v25
                                            _ = v26
                                        }
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    } else {
                                        if false {
                                            v27 := GenS1{F0: v1, F1: v19}
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            v28 := GenS2{F0: (fi21 / ((fi21) | 1)), F1: v24, F2: true}
                                            _ = v27
                                            _ = v28
                                        }
                                    }
                                    delete(v3, 9223372036854775807)
                                case "foo bar":
                                    if true {
                                        v29 := v2[(fi23 / ((-2147483648) | 1))]
                                        switch ("hello" + ("" + "hello")) {
                                        case "\n\t":
                                            var v30 float64 = 0.1
                                            _ = v30
                                        case "":
                                            switch "\n\t" {
                                            case "hello":
                                                if (true || ((v24 < uint64(v20)) || (v5 <= -741))) {
                                                    v3[fi21] = -84
                                                    v4 = append(v4, ("" + "\n\t"))
                                                    switch ((("" + ("" + "hello")) + ("hello" + "\n\t")) + "hello") {
                                                    case "\n\t":
                                                        v31 := GenS2{F0: (fi21 << ((0) & 63)), F1: (v24 % (((v1 | (v24 | v24))) | 1)), F2: false}
                                                        v32 := GenS2{F0: (fi23 / ((9223372036854775807) | 1)), F1: 4294967295, F2: false}
                                                        if true {
                                                            v19 = v19
                                                            v24 = v24
                                                            v1 += v1
                                                        }
                                                        v33 := GenS1{F0: v1, F1: (v19 + (-(v19)))}
                                                        _ = v31
                                                        _ = v32
                                                        _ = v33
                                                    case "a":
                                                        v24 = (v24 ^ (v1 % ((v24) | 1)))
                                                    case "":
                                                        v24 += (v24 / ((383) | 1))
                                                    default:
                                                        v20 += 343
                                                    }
                                                } else {
                                                    fmt.Println(v1)
                                                    v34 := GenS3{F0: (v29 | 2147483648), F1: v19, F2: false}
                                                    v24 = 2
                                                    v35 := GenS4{F0: int8(fi21), F1: -1, F2: -401}
                                                    _ = v34
                                                    _ = v35
                                                }
                                            }
                                        }
                                        var v36 uint16 = (v13 + (^(v13)))
                                        _ = v29
                                        _ = v36
                                    }
                                default:
                                    switch 127 {
                                    case -247:
                                        if (!true) {
                                            switch "foo bar" {
                                            case "\n\t":
                                                v37 := GenS4{F0: 1, F1: -1, F2: (fi21 << (((v20 >> ((v20) & 63))) & 63))}
                                                _ = v37
                                            default:
                                                var v38 bool = (true && false)
                                                _ = v38
                                            }
                                            switch "\n\t" {
                                            case "a":
                                                switch (fi23 % ((int(v19)) | 1)) {
                                                case 127:
                                                    var v39 int64 = v5
                                                    _ = v39
                                                case 4:
                                                    var v40 float64 = ((3.14159 + (0.1 * -2.25)) - 0.1)
                                                    _ = v40
                                                default:
                                                    v24 -= 1
                                                    var v41 uint32 = 2
                                                    _ = v41
                                                }
                                            case "\n\t":
                                                delete(v3, -1)
                                            default:
                                                var v42 int8 = -28
                                                v20 = 4294967295
                                                var v43 int8 = (v42 | v42)
                                                v44 := GenS5{F0: v5, F1: (^(v13))}
                                                _ = v42
                                                _ = v43
                                                _ = v44
                                            }
                                        }
                                    case -2147483648:
                                        fmt.Println(v6, v24, v6)
                                        var v45 bool = ("foo bar" <= ("\n\t" + "\n\t"))
                                        fmt.Println(v6)
                                        _ = v45
                                    }
                                }
                                v13 = v13
                                switch fi23 {
                                case -2147483648:
                                    v46 := GenS6{F0: false, F1: (v1 >> ((uint(v1)) & 63)), F2: 100}
                                    _ = v46
                                case 128:
                                    v20 -= 161
                                }
                            }
                        case "hello":
                            if false {
                                var v47 string = (("foo bar" + "\n\t") + "")
                                var v48 int64 = v5
                                _ = v47
                                _ = v48
                            } else {
                                var v49 int32 = 2147483647
                                var v50 float32 = v6
                                _ = v49
                                _ = v50
                            }
                            var v51 uint32 = 2147483647
                            fmt.Println(v5, fi21)
                            _ = v51
                        case "\n\t":
                            switch "foo bar" {
                            case "\n\t":
                                switch fi23 {
                                case 42:
                                    if (!true) {
                                        var v52 uint64 = (v24 >> ((0) & 63))
                                        _ = v52
                                    }
                                case -440:
                                    v53 := GenS2{F0: (fi23 % ((fi21) | 1)), F1: 2, F2: false}
                                    _ = v53
                                }
                            }
                        }
                        if false {
                            switch (fi21 + (fi21 | fi21)) {
                            case -108:
                                v54 := GenS6{F0: (!true), F1: v1, F2: 127}
                                _ = v54
                            }
                        } else {
                            v5 -= (v5 & (^(v5)))
                        }
                        v55 := GenS3{F0: 251, F1: (^(v19)), F2: true}
                        var v56 uint = v20
                        if false {
                            switch "a" {
                            case "":
                                v20 -= v56
                            case "\n\t":
                                v1 = 0
                            }
                        } else {
                            if (false || true) {
                                v19 = v19
                            } else {
                                if false {
                                    if (0 == -47) {
                                        v57 := GenS3{F0: 588, F1: int32(v20), F2: false}
                                        _ = v57
                                    } else {
                                        v1 = uint64(v13)
                                        var v58 uint32 = 100
                                        _ = v58
                                    }
                                    v59 := GenS6{F0: true, F1: 1, F2: -28}
                                    _ = v59
                                } else {
                                    var v60 bool = false
                                    _ = v60
                                }
                                if (!(!("a" == ("a" + "\n\t")))) {
                                    if (v1 > v24) {
                                        switch 127 {
                                        case 508:
                                            if len(v4) > 0 {
                                                v4[0] = "hello"
                                            }
                                        case -1:
                                            if ((fi21 | -406) < fi23) {
                                                if (v56 == (^(v20))) {
                                                    v61 := GenS7{F0: v1}
                                                    var v62 int16 = 1
                                                    _ = v61
                                                    _ = v62
                                                }
                                                if (false && true) {
                                                    _ = len(v4)
                                                    v24 = (v24 << ((1) & 63))
                                                    v20 = (^(v20))
                                                }
                                            } else {
                                                if ((4294967295 <= 66) || false) {
                                                    var v63 float64 = 3.14159
                                                    _ = v63
                                                }
                                            }
                                        case 9223372036854775807:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        default:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        switch "\n\t" {
                                        case "a":
                                            _ = len(v2)
                                        case "\n\t":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        default:
                                            var v64 uint32 = 771
                                            _ = v64
                                        }
                                    } else {
                                        v4 = append(v4, ("foo bar" + ("" + "hello")))
                                        fmt.Println(v24, v1, v24)
                                    }
                                    _ = len(v2)
                                    v65 := GenS8{F0: fi23, F1: 624, F2: (fi21 / ((-1) | 1))}
                                    _ = v65
                                }
                            }
                            switch "a" {
                            case "\n\t":
                                if (1e10 != 0.1) {
                                    var v66 uint16 = (v13 / (((^(v13))) | 1))
                                    if (v19 <= 1) {
                                        v56 += v20
                                        v5 = (^(v5))
                                        var v67 int16 = 32767
                                        _ = v67
                                    } else {
                                        v6 -= (v6 * (1e10 + v6))
                                        v19 = 985
                                    }
                                    v6 = 0.1
                                    _ = v66
                                } else {
                                    switch fi21 {
                                    case 128:
                                        var v68 string = "a"
                                        _ = v68
                                    case -9223372036854775808:
                                        switch 2147483647 {
                                        case 32767:
                                            v69 := GenS1{F0: v24, F1: (v19 / ((v19) | 1))}
                                            v20 = (v20 - v56)
                                            v5 = v5
                                            v20 += v56
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = v69
                                        }
                                    }
                                    fmt.Println(v6, v20, fi23)
                                }
                            default:
                                var v70 string = "foo bar"
                                _ = v70
                            }
                            fmt.Println(v13)
                        }
                        _ = fi23
                        _ = v24
                        _ = v55
                        _ = v56
                    }
                default:
                    fmt.Println(v20)
                }
            }
        } else {
            switch -1 {
            case 264:
                v71 := GenS1{F0: (v1 / (((v1 - (v1 | (v1 / ((v1) | 1))))) | 1)), F1: v19}
                _ = v71
            }
        }
        _ = fi21
    }
    if true {
        for fi72 := 0; fi72 < 9; fi72++ {
            v13 -= (^(v13))
            if (42 >= v5) {
                v73 := GenS6{F0: true, F1: (v1 % ((168) | 1)), F2: 100}
                _ = v73
            }
            var v74 uint32 = 0
            _ = fi72
            _ = v74
        }
        for fi75 := 0; fi75 < 16; fi75++ {
            var v76 uint = v20
            v13 += uint16(v1)
            var v77 int64 = v5
            _ = fi75
            _ = v76
            _ = v77
        }
        var v78 bool = true
        _ = v78
    } else {
        if false {
            _ = len(v3)
            if (v19 < (-(v19))) {
                if (!false) {
                    switch -427 {
                    case 128:
                        if false {
                            if (9223372036854775807 >= -2147483648) {
                                if true {
                                    switch "\n\t" {
                                    case "foo bar":
                                        switch "foo bar" {
                                        case "a":
                                            if true {
                                                v79 := GenS1{F0: (v1 % ((v1) | 1)), F1: v19}
                                                for fi80 := 0; fi80 < 7; fi80++ {
                                                    v19 = (-(v19))
                                                    _ = fi80
                                                }
                                                _ = v79
                                            } else {
                                                v13 -= v13
                                                fmt.Println(v5)
                                            }
                                        case "\n\t":
                                            v81 := v2[1]
                                            _ = v81
                                        case "hello":
                                            if len(v4) > 0 {
                                                v4[0] = "hello"
                                            }
                                        }
                                    case "":
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    default:
                                        v13 -= (v13 | (^(v13)))
                                    }
                                    fmt.Println(v1, v19, v20)
                                }
                            }
                            v6 += v6
                            var v82 int32 = -31
                            _ = v82
                        }
                    case -861:
                        for fi83 := 0; fi83 < 18; fi83++ {
                            v84 := GenS1{F0: v1, F1: v19}
                            v1 += (v1 % ((v1) | 1))
                            for fi85 := 0; fi85 < 15; fi85++ {
                                var v86 int8 = 127
                                var v87 uint16 = (^(v13))
                                _ = fi85
                                _ = v86
                                _ = v87
                            }
                            _ = fi83
                            _ = v84
                        }
                    case -759:
                        v19 += v19
                    }
                    for fi88 := 0; fi88 < 14; fi88++ {
                        v1 -= (v1 & v1)
                        if (!false) {
                            if ("" >= (("" + "\n\t") + "")) {
                                if true {
                                    if len(v4) > 0 {
                                        v4[0] = "\n\t"
                                    }
                                }
                            }
                            if (4294967295 == (fi88 / ((fi88) | 1))) {
                                if true {
                                    var v89 bool = (!true)
                                    _ = v89
                                }
                                if false {
                                    v13 = v13
                                    var v90 bool = false
                                    _ = v90
                                } else {
                                    for fi91 := 0; fi91 < 4; fi91++ {
                                        fmt.Println(v1, v5)
                                        _ = fi91
                                    }
                                    v13 = (^(v13))
                                    for fi92 := 0; fi92 < 11; fi92++ {
                                        v1 -= (v1 & v1)
                                        if ((!true) && false) {
                                            _ = len(v4)
                                            v93 := GenS2{F0: 42, F1: (v1 % ((v1) | 1)), F2: (((v19 | (v19 * v19)) >= (-(v19))) || ((fi92 >> (((v20 << ((v20) & 63))) & 63)) >= fi92))}
                                            if len(v4) > 0 {
                                                v4[0] = (("foo bar" + "hello") + ("\n\t" + (("hello" + "") + "foo bar")))
                                            }
                                            _ = v93
                                        } else {
                                            fmt.Println(v5, v19)
                                        }
                                        _ = fi92
                                    }
                                }
                            } else {
                                v13 -= v13
                            }
                        } else {
                            v1 = 1
                            if true {
                                v94 := GenS9{F0: -32768}
                                if false {
                                    var v95 int = -1
                                    fmt.Println(v1, v13, v19)
                                    v96 := GenS7{F0: (v1 ^ 776)}
                                    _ = v95
                                    _ = v96
                                } else {
                                    for fi97 := 0; fi97 < 9; fi97++ {
                                        if true {
                                            v6 -= v6
                                            v1 -= v1
                                            v6 += (v6 * ((v6 + v6) / (0.1 / v6)))
                                        } else {
                                            v13 = uint16(v20)
                                            fmt.Println(fi97, v1, v5)
                                            fmt.Println(v1, v13)
                                            var v98 int32 = (^(v19))
                                            v99 := GenS4{F0: 42, F1: int16(fi88), F2: (-(fi97))}
                                            _ = v98
                                            _ = v99
                                        }
                                        switch fi97 {
                                        case -709:
                                            fmt.Println(v5, fi88, v19)
                                        case 42:
                                            if false {
                                                var v100 int64 = (v5 ^ v5)
                                                fmt.Println(v5, v20, fi97)
                                                _ = v100
                                            } else {
                                                if true {
                                                    var v101 int16 = 961
                                                    _ = v101
                                                }
                                                v13 += (v13 - v13)
                                                fmt.Println(v13, v19)
                                                var v102 float64 = -2.25
                                                _ = v102
                                            }
                                        case -2147483648:
                                            v19 = (v19 >> ((uint(v13)) & 63))
                                        }
                                        v3[-625] = fi88
                                        _ = fi97
                                    }
                                }
                                if true {
                                    v1 += (v1 ^ v1)
                                    delete(v3, fi88)
                                    if true {
                                        if (1 > -32768) {
                                            v6 = 1e10
                                            var v103 uint64 = v1
                                            for fi104 := 0; fi104 < 16; fi104++ {
                                                v105 := GenS10{F0: -32768, F1: 1}
                                                v19 = (-(v19))
                                                var v106 int32 = (v19 / ((v19) | 1))
                                                v106 -= int32(v20)
                                                _ = fi104
                                                _ = v105
                                                _ = v106
                                            }
                                            _ = v103
                                        } else {
                                            for fi107 := 0; fi107 < 17; fi107++ {
                                                fmt.Println(v20, v1)
                                                v108 := GenS11{F0: (v20 & v20), F1: v13, F2: 0}
                                                _ = fi107
                                                _ = v108
                                            }
                                            v109 := GenS6{F0: (1.5 <= (0.0 / 1e10)), F1: (^(v1)), F2: -1}
                                            if true {
                                                _ = len(v2)
                                                if ("a" > "a") {
                                                    var v110 bool = ((fi88 ^ fi88) > (fi88 & fi88))
                                                    if false {
                                                        if len(v4) > 0 {
                                                            v4[0] = "foo bar"
                                                        }
                                                    } else {
                                                        if v110 {
                                                            v111 := GenS12{F0: (v13 % (((v13 * (v13 % (((v13 + (v13 * v13))) | 1)))) | 1)), F1: "foo bar", F2: -32768}
                                                            var v112 bool = ((false || true) && true)
                                                            for fi113 := 0; fi113 < 20; fi113++ {
                                                                v114 := GenS13{F0: v20, F1: ""}
                                                                var v115 string = "\n\t"
                                                                _ = fi113
                                                                _ = v114
                                                                _ = v115
                                                            }
                                                            _ = v111
                                                            _ = v112
                                                        }
                                                        _ = len(v3)
                                                        v19 -= (v19 - v19)
                                                    }
                                                    _ = v110
                                                }
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            } else {
                                                var v116 int8 = 70
                                                var v117 float32 = (v6 - v6)
                                                _ = v116
                                                _ = v117
                                            }
                                            _ = v109
                                        }
                                        fmt.Println(v1, fi88)
                                    } else {
                                        switch ("a" + "foo bar") {
                                        case "a":
                                            v118 := GenS14{F0: 1, F1: "\n\t"}
                                            _ = v118
                                        }
                                        var v119 float64 = (0.1 / 1.5)
                                        if ((^(v20)) != (v20 >> ((370) & 63))) {
                                            v6 = (1.5 * ((v6 - (v6 / float32(-2.25))) - (float32((v119 * v119)) - float32(float64(v6)))))
                                            fmt.Println(v119, v19)
                                        }
                                        _ = v119
                                    }
                                }
                                _ = v94
                            } else {
                                if true {
                                    var v120 uint = 916
                                    v120 = v120
                                    _ = v120
                                }
                            }
                            v121 := GenS7{F0: (v1 | (^(v1)))}
                            _ = v121
                        }
                        var v122 uint = 32
                        _ = fi88
                        _ = v122
                    }
                    for fi123 := 0; fi123 < 8; fi123++ {
                        switch (fi123 / ((fi123) | 1)) {
                        case 908:
                            if ((11 >= 79) || (1e10 == 1e10)) {
                                var v124 uint32 = 334
                                var v125 uint16 = v13
                                _ = v124
                                _ = v125
                            }
                        case 32767:
                            delete(v2, fi123)
                        }
                        var v126 bool = (false || (v20 != (v20 | v20)))
                        _ = fi123
                        _ = v126
                    }
                } else {
                    v6 -= ((1.5 + v6) - (0.0 + v6))
                    v20 -= 18446744073709551615
                }
            } else {
                v127 := GenS7{F0: 826}
                for fi128 := 0; fi128 < 7; fi128++ {
                    var v129 uint = v20
                    v19 = v19
                    _ = fi128
                    _ = v129
                }
                for fi130 := 0; fi130 < 1; fi130++ {
                    for fi131 := 0; fi131 < 16; fi131++ {
                        var v132 uint64 = 42
                        _ = fi131
                        _ = v132
                    }
                    v133 := GenS12{F0: (v13 / ((v13) | 1)), F1: "hello", F2: 32767}
                    for fi134 := 0; fi134 < 8; fi134++ {
                        v5 = v5
                        var v135 uint8 = 127
                        switch fi130 {
                        case 835:
                            if false {
                                if (((^(v135)) > v135) && false) {
                                    fmt.Println(fi134)
                                } else {
                                    v136 := GenS15{F0: 42, F1: "foo bar"}
                                    if len(v4) > 0 {
                                        v4[0] = ((("\n\t" + "a") + "a") + "hello")
                                    }
                                    switch "\n\t" {
                                    case "":
                                        var v137 uint32 = 42
                                        _ = v137
                                    case "\n\t":
                                        _ = len(v2)
                                    default:
                                        v20 += v20
                                    }
                                    _ = v136
                                }
                                v135 += v135
                            }
                        }
                        _ = fi134
                        _ = v135
                    }
                    _ = fi130
                    _ = v133
                }
                _ = v127
            }
        } else {
            for fi138 := 0; fi138 < 6; fi138++ {
                if ((v20 * (v20 & (v20 ^ (v20 * 2)))) < (v20 - (v20 + (v20 ^ (^(v20)))))) {
                    v20 = v20
                }
                var v139 uint64 = 4294967295
                if (v1 > (v1 ^ 456)) {
                    v140 := GenS16{F0: 25}
                    if ((v19 > v19) && (!false)) {
                        var v141 int8 = 127
                        var v142 uint16 = v13
                        _ = v141
                        _ = v142
                    }
                    if len(v4) > 0 {
                        v4[0] = ""
                    }
                    _ = v140
                }
                _ = fi138
                _ = v139
            }
            var v143 float64 = -2.25
            _ = v143
        }
    }
    v144 := GenS4{F0: 42, F1: 32767, F2: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v144
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
