package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 int64; F2 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 int16 }
type GenS7 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int32; F1 string }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint8; F1 uint64; F2 int }
type GenS13 struct { F0 bool; F1 int64 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint64; F1 int16 }

func main() {
    var v1 int32 = 1
    var v2 uint32 = 2147483648
    v3 := make(map[uint32]int)
    v4 := make(map[int]int)
    v5 := []string{"\n\t", "foo bar", "foo bar", "a"}
    var v6 int8 = 127
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        if (true || ((("a" + "\n\t") + "") > ("a" + ("foo bar" + "\n\t")))) {
            for fi7 := 0; fi7 < 1; fi7++ {
                v2 += (^(v2))
                v8 := GenS1{F0: true, F1: 0}
                if false {
                    fmt.Println(v6)
                    if true {
                        var v9 uint32 = v2
                        _ = v9
                    } else {
                        var v10 int16 = 1
                        _ = v10
                    }
                }
                _ = fi7
                _ = v8
            }
        }
    case "hello":
        var v11 float64 = 0.1
        _ = v11
    case "\n\t":
        if (((v1 << ((1) & 63)) > v1) && false) {
            v2 -= 375
            fmt.Println(v6, v6, v1)
        }
    default:
        v12 := GenS1{F0: (1e10 != 0.0), F1: 42}
        _ = v12
    }
    _ = len(v5)
    v2 -= 2147483648
    v2 -= (v2 - (v2 << ((684) & 63)))
    var v13 string = "foo bar"
    v14 := GenS2{F0: 0, F1: -9223372036854775808}
    var v15 string = v13
    var v16 uint8 = 135
    v17 := GenS2{F0: (v16 | (v16 * (v16 % ((v16) | 1)))), F1: 645}
    var v18 int64 = -251
    if false {
        fmt.Println(v15, v1, v16)
        for fi19 := 0; fi19 < 18; fi19++ {
            v3[65535] = (fi19 ^ (fi19 ^ (fi19 ^ fi19)))
            _ = fi19
        }
    } else {
        for fi20 := 0; fi20 < 13; fi20++ {
            v1 = (v1 % ((v1) | 1))
            _ = fi20
        }
        for fi21 := 0; fi21 < 20; fi21++ {
            if false {
                v22 := GenS1{F0: (v18 != (v18 % (((^(v18))) | 1))), F1: v18}
                if ((v15 + ((v15 + v15) + v13)) >= v15) {
                    if false {
                        if (3.14159 > (1.5 - float32((3.14159 + 0.1)))) {
                            for fi23 := 0; fi23 < 6; fi23++ {
                                v24 := GenS3{F0: ((-2.25 - float32(float64(1.5))) <= 0.1)}
                                var v25 uint32 = v2
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                            v15 = ((v15 + v15) + v13)
                        } else {
                            var v26 float64 = 0.0
                            for fi27 := 0; fi27 < 8; fi27++ {
                                if (false || (v2 != (v2 & v2))) {
                                    if ((v6 | 1) < v6) {
                                        var v28 uint = 0
                                        _ = v28
                                    } else {
                                        v26 = v26
                                        v26 = float64(-2.25)
                                        delete(v3, 4294967295)
                                        v13 = "\n\t"
                                    }
                                    _ = len(v3)
                                    var v29 int32 = (v1 * (v1 * (v1 % ((v1) | 1))))
                                    if true {
                                        v1 -= (^(v1))
                                        fmt.Println(v15, v15)
                                        fmt.Println(fi27, v1)
                                    } else {
                                        fmt.Println(fi27, v16)
                                    }
                                    _ = v29
                                } else {
                                    v18 = (v18 >> ((0) & 63))
                                    switch ((v15 + v15) + v15) {
                                    case "":
                                        var v30 uint = 4294967295
                                        _ = v30
                                    case "foo bar":
                                        v26 -= ((v26 * (v26 + v26)) + 3.14159)
                                    case "a":
                                        if (true || true) {
                                            v15 = v15
                                            if ((-(v1)) > v1) {
                                                if false {
                                                    var v31 int = (fi27 / ((fi27) | 1))
                                                    v32 := GenS1{F0: false, F1: 9223372036854775807}
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    switch fi21 {
                                                    case 127:
                                                        if (true || true) {
                                                            if ((v16 | uint8(v2)) == (^(v16))) {
                                                                fmt.Println(v18)
                                                            }
                                                        } else {
                                                            v33 := GenS2{F0: v16, F1: (v18 - (v18 & (v18 * v18)))}
                                                            v13 = "foo bar"
                                                            v2 = v2
                                                            _ = v33
                                                        }
                                                    case 606:
                                                        v34 := GenS3{F0: ((18446744073709551615 == 65535) || true)}
                                                        _ = v34
                                                    }
                                                    switch ("a" + v13) {
                                                    case "hello":
                                                        delete(v4, 4294967295)
                                                    case "\n\t":
                                                        switch (^(fi21)) {
                                                        case -7:
                                                            v26 = 0.1
                                                        case 123:
                                                            if len(v5) > 0 {
                                                                v5[0] = "hello"
                                                            }
                                                            var v35 bool = false
                                                            v36 := GenS4{F0: v13, F1: v18, F2: 32767}
                                                            _ = v35
                                                            _ = v36
                                                        default:
                                                            var v37 int32 = (v1 + v1)
                                                            _ = v37
                                                        }
                                                    default:
                                                        v18 -= v18
                                                    }
                                                    _ = v31
                                                    _ = v32
                                                }
                                                v2 = 2
                                                if true {
                                                    if true {
                                                        v26 = (float64(0.0) * v26)
                                                        v38 := v3[(v2 >> ((uint(v2)) & 63))]
                                                        _ = v38
                                                    } else {
                                                        var v39 string = "a"
                                                        v16 = v16
                                                        var v40 int64 = 42
                                                        _ = v39
                                                        _ = v40
                                                    }
                                                    switch fi27 {
                                                    case -1:
                                                        var v41 uint16 = 0
                                                        _ = v41
                                                    default:
                                                        var v42 int16 = 0
                                                        v6 = v6
                                                        _ = v42
                                                    }
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                                v18 = (v18 << ((65535) & 63))
                                _ = fi27
                            }
                            if (true || true) {
                                fmt.Println(v18)
                                for fi43 := 0; fi43 < 20; fi43++ {
                                    v44 := GenS5{F0: fi43}
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    var v45 int64 = (v18 % ((v18) | 1))
                                    _ = fi43
                                    _ = v44
                                    _ = v45
                                }
                            } else {
                                v46 := GenS5{F0: fi21}
                                switch int(v18) {
                                case 501:
                                    v26 = v26
                                case -10:
                                    v16 -= v16
                                case 2147483647:
                                    var v47 float64 = 1e10
                                    _ = v47
                                }
                                var v48 uint = 2
                                _ = v46
                                _ = v48
                            }
                            _ = v26
                        }
                        if len(v5) > 0 {
                            v5[0] = v15
                        }
                    } else {
                        if true {
                            var v49 string = v15
                            _ = len(v3)
                            _ = v49
                        }
                        switch "foo bar" {
                        case "hello":
                            v18 = (-(v18))
                        case "a":
                            v50 := GenS2{F0: (^(v16)), F1: 9223372036854775807}
                            _ = v50
                        case "":
                            if ((v6 | v6) != 80) {
                                v5 = append(v5, "hello")
                                v2 = (v2 >> ((65535) & 63))
                            } else {
                                switch (fi21 & 0) {
                                case -823:
                                    if false {
                                        if (!false) {
                                            for fi51 := 0; fi51 < 10; fi51++ {
                                                v15 = (v15 + "hello")
                                                switch v13 {
                                                case "hello":
                                                    if (100 == v6) {
                                                        v4[(^(fi21))] = fi21
                                                        v52 := v4[fi51]
                                                        switch "hello" {
                                                        case "":
                                                            var v53 bool = (18446744073709551615 == 983)
                                                            _ = v53
                                                        }
                                                        if (false && true) {
                                                            v18 -= (v18 * (v18 / ((v18) | 1)))
                                                        }
                                                        _ = v52
                                                    } else {
                                                        var v54 int8 = (v6 - (v6 / (((v6 - int8(fi51))) | 1)))
                                                        _ = v54
                                                    }
                                                case "":
                                                    fmt.Println(fi21, v16, v1)
                                                case "a":
                                                    v55 := GenS2{F0: (v16 >> ((466) & 63)), F1: 9223372036854775807}
                                                    _ = v55
                                                default:
                                                    v56 := GenS2{F0: v16, F1: (v18 / (((v18 - int64(v16))) | 1))}
                                                    _ = v56
                                                }
                                                _ = fi51
                                            }
                                            var v57 int32 = -2147483648
                                            if (!(0 > 2)) {
                                                for fi58 := 0; fi58 < 7; fi58++ {
                                                    if (true && ((false || (915 <= 2)) || false)) {
                                                        v18 -= v18
                                                        if false {
                                                            delete(v3, (v2 + v2))
                                                            v59 := GenS6{F0: -32768, F1: 32767}
                                                            _ = v59
                                                        }
                                                        switch int(v16) {
                                                        case -2147483648:
                                                            v60 := GenS6{F0: 0, F1: 1}
                                                            var v61 float32 = (1.5 + 1e10)
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            v62 := GenS7{F0: v1, F1: 965, F2: 439}
                                                            v63 := GenS5{F0: (fi58 / (((fi58 | 32768)) | 1))}
                                                            _ = v60
                                                            _ = v61
                                                            _ = v62
                                                            _ = v63
                                                        }
                                                    }
                                                    v64 := GenS8{F0: (^(v2))}
                                                    switch fi21 {
                                                    case 467:
                                                        v15 = v15
                                                    default:
                                                        if false {
                                                            var v65 int64 = (v18 + v18)
                                                            _ = v65
                                                        }
                                                    }
                                                    _ = fi58
                                                    _ = v64
                                                }
                                                if (false || true) {
                                                    v66 := GenS1{F0: (1.5 <= (0.0 - -2.25)), F1: -1}
                                                    var v67 int = (fi21 ^ -1)
                                                    _ = v66
                                                    _ = v67
                                                }
                                                v4[fi21] = (fi21 - 128)
                                            } else {
                                                v18 += (v18 % ((v18) | 1))
                                            }
                                            _ = v57
                                        }
                                    }
                                case 128:
                                    if len(v5) > 0 {
                                        v5[0] = v15
                                    }
                                }
                                switch v13 {
                                case "\n\t":
                                    if true {
                                        v5 = append(v5, (v13 + "\n\t"))
                                    } else {
                                        v16 -= v16
                                        fmt.Println(v15, v2, v18)
                                        switch 0 {
                                        case 9223372036854775807:
                                            v68 := GenS2{F0: (v16 / ((255) | 1)), F1: v18}
                                            _ = v68
                                        case 1:
                                            var v69 bool = false
                                            _ = v69
                                        default:
                                            if true {
                                                fmt.Println(v1)
                                            }
                                        }
                                    }
                                }
                            }
                        default:
                            v4[(fi21 & (-(fi21)))] = 2147483647
                        }
                        v18 -= (v18 / ((v18) | 1))
                    }
                    var v70 bool = (!false)
                    _ = len(v3)
                    _ = v70
                } else {
                    if true {
                        v71 := GenS8{F0: 198}
                        if len(v5) > 0 {
                            v5[0] = (v13 + "\n\t")
                        }
                        v1 -= (v1 * (v1 | int32(v18)))
                        _ = v71
                    } else {
                        var v72 float32 = float32((0.0 - 0.0))
                        for fi73 := 0; fi73 < 16; fi73++ {
                            if (((3.14159 / float64((float32(1e10) * 3.14159))) + 1.5) <= float64(v72)) {
                                fmt.Println(fi73)
                                if true {
                                    if (false || false) {
                                        v15 = ""
                                        v15 = ((("\n\t" + v13) + v15) + v15)
                                        if (((v1 >> ((595) & 63)) > (-(v1))) && true) {
                                            var v74 int64 = 0
                                            fmt.Println(v16, v6, fi73)
                                            fmt.Println(v18)
                                            _ = v74
                                        } else {
                                            switch "hello" {
                                            case "":
                                                switch v13 {
                                                case "\n\t":
                                                    if true {
                                                        delete(v3, v2)
                                                        switch (fi73 << ((uint(v2)) & 63)) {
                                                        case 0:
                                                            v2 = 2
                                                        case 4294967295:
                                                            v75 := GenS9{F0: 18446744073709551615}
                                                            v76 := v3[(v2 * 2147483647)]
                                                            _ = v75
                                                            _ = v76
                                                        }
                                                        switch (-(fi73)) {
                                                        case -204:
                                                            v18 -= v18
                                                        case 4294967295:
                                                            v6 += v6
                                                        default:
                                                            if (true && (!(v72 >= (v72 + 0.0)))) {
                                                                v18 -= -570
                                                            }
                                                        }
                                                    }
                                                }
                                            case "a":
                                                v77 := v4[(fi21 - (fi21 - 32768))]
                                                _ = v77
                                            case "hello":
                                                fmt.Println(fi73, fi21)
                                            default:
                                                var v78 bool = (!true)
                                                var v79 string = (((v13 + "a") + v15) + v15)
                                                _ = v78
                                                _ = v79
                                            }
                                        }
                                    } else {
                                        var v80 int32 = int32(v18)
                                        fmt.Println(v2)
                                        v5 = append(v5, (v13 + v13))
                                        v16 -= 255
                                        _ = v80
                                    }
                                }
                                var v81 int = fi73
                                switch v13 {
                                case "\n\t":
                                    var v82 int16 = 484
                                    _ = v82
                                case "":
                                    v83 := GenS7{F0: v1, F1: 764, F2: 1}
                                    _ = v83
                                case "hello":
                                    var v84 string = ("" + v13)
                                    _ = v84
                                default:
                                    if len(v5) > 0 {
                                        v5[0] = v15
                                    }
                                }
                                _ = v81
                            } else {
                                _ = len(v5)
                                switch v13 {
                                case "":
                                    v6 += v6
                                case "foo bar":
                                    if (-239 == -32768) {
                                        var v85 float32 = v72
                                        v5 = append(v5, v13)
                                        v72 += (v85 * (-2.25 - v72))
                                        _ = v85
                                    }
                                case "a":
                                    var v86 float64 = 3.14159
                                    _ = v86
                                }
                            }
                            v6 += v6
                            if true {
                                var v87 uint8 = (v16 << ((1) & 63))
                                v88 := GenS1{F0: (true || ((v6 * (v6 ^ (^(v6)))) >= 0)), F1: v18}
                                fmt.Println(v15, v16)
                                _ = v87
                                _ = v88
                            } else {
                                var v89 float64 = float64((0.0 / float32(-2.25)))
                                v5 = append(v5, "")
                                _ = v89
                            }
                            switch fi73 {
                            case 981:
                                switch -2147483648 {
                                case 9223372036854775807:
                                    fmt.Println(v16, fi73)
                                }
                            case -259:
                                v72 = v72
                            }
                            v13 = ((((v15 + v15) + "a") + (v15 + (v13 + v13))) + (v15 + v13))
                            _ = fi73
                        }
                        _ = v72
                    }
                    switch v15 {
                    case "a":
                        v2 -= v2
                    case "\n\t":
                        for fi90 := 0; fi90 < 5; fi90++ {
                            var v91 string = v15
                            v13 = "foo bar"
                            var v92 string = "\n\t"
                            v91 = v92
                            _ = fi90
                            _ = v91
                            _ = v92
                        }
                    default:
                        switch v15 {
                        case "foo bar":
                            if (((((1e10 + 0.0) / (-2.25 - 1e10)) / (3.14159 * 3.14159)) + (0.1 - 3.14159)) < float32((1.5 / float64(1.5)))) {
                                var v93 int64 = 0
                                _ = v93
                            }
                        default:
                            fmt.Println(v2, v6, v18)
                        }
                    }
                    if false {
                        for fi94 := 0; fi94 < 2; fi94++ {
                            v15 = v15
                            var v95 float32 = (0.1 - 0.0)
                            v96 := GenS1{F0: true, F1: 42}
                            _ = fi94
                            _ = v95
                            _ = v96
                        }
                        if len(v5) > 0 {
                            v5[0] = (v15 + "hello")
                        }
                        v15 = ("" + v15)
                    } else {
                        v6 = v6
                        for fi97 := 0; fi97 < 19; fi97++ {
                            v5 = append(v5, v15)
                            if false {
                                switch "foo bar" {
                                case "":
                                    if true {
                                        var v98 uint16 = 32767
                                        if ((!false) && (false || (false && (1e10 == -2.25)))) {
                                            v99 := GenS10{F0: (-(v1)), F1: v15}
                                            _ = v99
                                        }
                                        switch 128 {
                                        case -543:
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        case 2147483647:
                                            fmt.Println(v6)
                                        default:
                                            _ = len(v5)
                                        }
                                        _ = v98
                                    } else {
                                        fmt.Println(v16, v2, fi97)
                                    }
                                case "foo bar":
                                    var v100 uint64 = 1
                                    _ = v100
                                case "hello":
                                    var v101 float64 = 0.0
                                    _ = v101
                                }
                            }
                            if false {
                                fmt.Println(v1)
                                if false {
                                    v6 = int8(v2)
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                } else {
                                    if ((v16 + (v16 * v16)) > v16) {
                                        if false {
                                            switch (fi97 % (((-(fi21))) | 1)) {
                                            case 127:
                                                var v102 int16 = 32767
                                                v13 = v15
                                                _ = v102
                                            }
                                            v4[2147483647] = fi21
                                            v1 = v1
                                        }
                                        if true {
                                            var v103 uint32 = 2147483647
                                            if (!false) {
                                                v13 = "a"
                                                switch 127 {
                                                case -1:
                                                    v6 = 21
                                                case 32768:
                                                    v6 = v6
                                                case -480:
                                                    v104 := v4[(fi97 / ((fi97) | 1))]
                                                    _ = v104
                                                default:
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                }
                                                var v105 string = v13
                                                _ = v105
                                            }
                                            _ = v103
                                        } else {
                                            v6 = (v6 + (v6 + (v6 & v6)))
                                            v18 = v18
                                        }
                                    }
                                    fmt.Println(v2)
                                    var v106 string = (v15 + v15)
                                    var v107 int8 = 127
                                    var v108 uint64 = 42
                                    switch (fi21 << ((348) & 63)) {
                                    case 487:
                                        var v109 float64 = float64(3.14159)
                                        _ = v109
                                    }
                                    _ = v106
                                    _ = v107
                                    _ = v108
                                }
                                v3[(v2 ^ (v2 | 0))] = 748
                                switch v13 {
                                case "hello":
                                    switch fi21 {
                                    case -9223372036854775808:
                                        v110 := GenS11{F0: v6}
                                        _ = v110
                                    case 32767:
                                        v111 := GenS9{F0: 820}
                                        _ = v111
                                    case 32768:
                                        fmt.Println(v18, v16)
                                    }
                                default:
                                    if (false && false) {
                                        var v112 int = (fi97 + fi21)
                                        var v113 uint16 = 32767
                                        v6 = v6
                                        _ = v112
                                        _ = v113
                                    }
                                }
                                v15 = (v13 + (v15 + ("foo bar" + (v13 + (v13 + (v13 + ("hello" + (v13 + "\n\t"))))))))
                                if (131 >= (v16 ^ v16)) {
                                    v16 = v16
                                    switch "foo bar" {
                                    case "hello":
                                        _ = len(v3)
                                    case "":
                                        var v114 uint64 = 4294967295
                                        switch v13 {
                                        case "hello":
                                            v1 = (v1 / (((^(v1))) | 1))
                                        default:
                                            v115 := v4[(fi97 << ((0) & 63))]
                                            _ = v115
                                        }
                                        var v116 float32 = 1.5
                                        switch fi97 {
                                        case 9223372036854775807:
                                            v116 = v116
                                        case 127:
                                            if (!false) {
                                                v114 -= 88
                                                v114 = 18446744073709551615
                                                var v117 float64 = 1.5
                                                _ = v117
                                            } else {
                                                v118 := GenS11{F0: (^(v6))}
                                                if false {
                                                    v119 := GenS5{F0: fi21}
                                                    _ = v119
                                                } else {
                                                    v5 = append(v5, (v15 + ("foo bar" + ((("foo bar" + v15) + ("\n\t" + ("" + v15))) + v13))))
                                                    switch "foo bar" {
                                                    case "hello":
                                                        if (v2 == v2) {
                                                            v16 -= 127
                                                            fmt.Println(v1)
                                                            var v120 float32 = 1.5
                                                            _ = v120
                                                        }
                                                    default:
                                                        v18 -= (-(v18))
                                                    }
                                                    if ((v18 - v18) <= (v18 << ((42) & 63))) {
                                                        _ = len(v3)
                                                    } else {
                                                        switch 4294967295 {
                                                        case -1:
                                                            var v121 string = ("\n\t" + v13)
                                                            _ = v121
                                                        case 0:
                                                            switch 0 {
                                                            case 935:
                                                                delete(v4, fi97)
                                                            default:
                                                                var v122 uint64 = (v114 << ((255) & 63))
                                                                _ = v122
                                                            }
                                                        case 42:
                                                            v123 := GenS4{F0: v13, F1: (v18 - v18), F2: 32767}
                                                            _ = v123
                                                        default:
                                                            _ = len(v5)
                                                        }
                                                        if true {
                                                            v114 += (v114 & 2)
                                                            if (0.1 > float32(0.0)) {
                                                                if (403 == 698) {
                                                                    if (((v13 + v15) >= ((v13 + v15) + v13)) && (2 <= 42)) {
                                                                        if len(v5) > 0 {
                                                                            v5[0] = (v13 + v15)
                                                                        }
                                                                    } else {
                                                                        v6 += (-(v6))
                                                                    }
                                                                } else {
                                                                    v124 := GenS2{F0: 1, F1: v18}
                                                                    var v125 int32 = (v1 | v1)
                                                                    _ = v124
                                                                    _ = v125
                                                                }
                                                                v126 := GenS1{F0: false, F1: 1}
                                                                fmt.Println(v18, v18)
                                                                v2 = v2
                                                                var v127 float64 = (3.14159 * 1e10)
                                                                _ = v126
                                                                _ = v127
                                                            } else {
                                                                var v128 uint32 = 0
                                                                if false {
                                                                    v116 = v116
                                                                    delete(v4, fi21)
                                                                }
                                                                _ = v128
                                                            }
                                                            fmt.Println(fi97, v6)
                                                            v129 := GenS6{F0: 1, F1: -394}
                                                            _ = v129
                                                        }
                                                        if true {
                                                            if false {
                                                                fmt.Println(v116, v13)
                                                                if (v13 < (v13 + (v13 + "\n\t"))) {
                                                                    fmt.Println(v1)
                                                                    if (!false) {
                                                                        v15 = v15
                                                                        v2 += v2
                                                                        v13 = (("a" + "\n\t") + v15)
                                                                    }
                                                                    v130 := GenS12{F0: v16, F1: (v114 << ((536) & 63)), F2: 2147483647}
                                                                    v1 = 2147483647
                                                                    var v131 int = (fi21 - (^(fi21)))
                                                                    var v132 uint8 = 236
                                                                    v116 += (1.5 - v116)
                                                                    var v133 int = fi97
                                                                    if (!(true || ((v114 == (^(v114))) && false))) {
                                                                        fmt.Println(v6, v16)
                                                                        fmt.Println(v133)
                                                                        var v134 uint64 = uint64(v132)
                                                                        _ = v134
                                                                    }
                                                                    v18 -= (v18 << ((255) & 63))
                                                                    _ = v130
                                                                    _ = v131
                                                                    _ = v132
                                                                    _ = v133
                                                                }
                                                            }
                                                        } else {
                                                            var v135 bool = ((v18 ^ int64(v1)) >= (v18 + v18))
                                                            var v136 int64 = (^(v18))
                                                            _ = v135
                                                            _ = v136
                                                        }
                                                        var v137 int = fi97
                                                        _ = v137
                                                    }
                                                }
                                                _ = v118
                                            }
                                        default:
                                            var v138 int16 = 483
                                            _ = v138
                                        }
                                        _ = v114
                                        _ = v116
                                    }
                                    v6 += v6
                                } else {
                                    if (!true) {
                                        var v139 uint = 2
                                        if false {
                                            fmt.Println(v13, v13)
                                            switch "" {
                                            case "hello":
                                                v1 = v1
                                            case "\n\t":
                                                v140 := GenS10{F0: v1, F1: "\n\t"}
                                                _ = v140
                                            case "foo bar":
                                                var v141 int8 = 100
                                                _ = v141
                                            default:
                                                v142 := GenS5{F0: fi21}
                                                _ = v142
                                            }
                                            var v143 float64 = 3.14159
                                            if false {
                                                switch fi21 {
                                                case 128:
                                                    v144 := GenS8{F0: (v2 / (((v2 + v2)) | 1))}
                                                    _ = v144
                                                default:
                                                    v18 = int64(v1)
                                                }
                                                v145 := GenS13{F0: ((^(v2)) <= (v2 ^ (v2 << (((v139 << ((v139) & 63))) & 63)))), F1: int64(v16)}
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                _ = v145
                                            }
                                            var v146 float64 = -2.25
                                            if (1 >= 2) {
                                                if true {
                                                    var v147 float32 = 1e10
                                                    v1 = 632
                                                    _ = v147
                                                }
                                            } else {
                                                var v148 uint = v139
                                                v149 := GenS7{F0: (v1 | v1), F1: 622, F2: 0}
                                                _ = v148
                                                _ = v149
                                            }
                                            switch (-(fi97)) {
                                            case 9223372036854775807:
                                                v139 -= 65535
                                            }
                                            _ = v143
                                            _ = v146
                                        }
                                        v150 := GenS11{F0: v6}
                                        _ = v139
                                        _ = v150
                                    } else {
                                        var v151 int8 = (v6 << ((549) & 63))
                                        _ = len(v5)
                                        var v152 int = fi21
                                        _ = v151
                                        _ = v152
                                    }
                                    fmt.Println(v2, v13, fi97)
                                    v5 = append(v5, (v15 + v15))
                                }
                            } else {
                                switch ("\n\t" + "\n\t") {
                                case "hello":
                                    fmt.Println(v6, v2)
                                case "a":
                                    switch (v13 + ((v15 + v15) + (((("" + "a") + v15) + "\n\t") + v15))) {
                                    case "hello":
                                        if (v1 != 2147483647) {
                                            fmt.Println(v2, v13)
                                        } else {
                                            if (!false) {
                                                if (42 != 4294967295) {
                                                    var v153 int64 = 558
                                                    var v154 bool = ((-2.25 / ((3.14159 * float32((1e10 / (1.5 - (3.14159 + (((-2.25 - 1e10) - 3.14159) / 1.5)))))) / -2.25)) > (float32(-2.25) - 3.14159))
                                                    _ = v153
                                                    _ = v154
                                                }
                                                v155 := GenS14{F0: 32767}
                                                var v156 float64 = float64(3.14159)
                                                _ = v155
                                                _ = v156
                                            } else {
                                                v1 -= 3
                                            }
                                        }
                                    default:
                                        v5 = append(v5, v15)
                                    }
                                }
                                if ((-1 > -1) && true) {
                                    if len(v5) > 0 {
                                        v5[0] = v15
                                    }
                                    var v157 int8 = v6
                                    v2 -= 2147483648
                                    _ = v157
                                } else {
                                    switch v15 {
                                    case "foo bar":
                                        var v158 uint8 = 42
                                        _ = v158
                                    case "\n\t":
                                        var v159 string = v13
                                        var v160 uint8 = (v16 & v16)
                                        _ = v159
                                        _ = v160
                                    case "hello":
                                        v18 = (v18 * int64(v1))
                                    default:
                                        v16 = (v16 + 236)
                                    }
                                    v161 := GenS1{F0: (0 != 2), F1: v18}
                                    _ = v161
                                }
                                if (148 != 1) {
                                    _ = len(v4)
                                    v162 := GenS15{F0: 118, F1: 1000}
                                    _ = v162
                                }
                            }
                            _ = fi97
                        }
                    }
                }
                var v163 uint16 = 255
                _ = v163
                _ = v22
            }
            v15 = v15
            _ = fi21
        }
        var v164 uint32 = (^(v2))
        _ = v164
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
