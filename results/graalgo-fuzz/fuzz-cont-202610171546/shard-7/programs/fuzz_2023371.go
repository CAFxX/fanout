package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int; F1 uint64; F2 uint }
type GenS7 struct { F0 int32; F1 int32; F2 uint8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint16; F1 uint8; F2 int }
type GenS10 struct { F0 int16; F1 int; F2 int }
type GenS11 struct { F0 string; F1 uint16; F2 uint32 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint; F1 bool; F2 int32 }
type GenS14 struct { F0 int8; F1 int32 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 uint; F1 uint32 }
type GenS17 struct { F0 int64; F1 bool; F2 int8 }
type GenS18 struct { F0 int16; F1 string }
type GenS19 struct { F0 int; F1 string; F2 int32 }

func main() {
    var v1 int8 = -1
    var v2 uint = 1
    var v3 int64 = 42
    v4 := make(map[string]int)
    var v5 int64 = (^(v3))
    v6 := make(map[int]string)
    for fi7 := 0; fi7 < 13; fi7++ {
        if (!true) {
            v1 += (v1 & v1)
            if true {
                v8 := []int{(fi7 >> ((v2) & 63))}
                if (1.5 == (0.1 + 3.14159)) {
                    v4["a"] = fi7
                } else {
                    if ("foo bar" != "foo bar") {
                        switch (fi7 + (-(fi7))) {
                        case 127:
                            fmt.Println(v1)
                        case -1:
                            for fi9 := 0; fi9 < 20; fi9++ {
                                if true {
                                    v10 := GenS1{F0: 32767, F1: v5, F2: (v3 << ((v2) & 63))}
                                    if (!(false || false)) {
                                        fmt.Println(v3, v1, v2)
                                    }
                                    _ = v10
                                }
                                switch (^(fi9)) {
                                case 4294967295:
                                    v11 := GenS1{F0: 42, F1: v5, F2: 1}
                                    _ = v11
                                default:
                                    switch fi9 {
                                    case 32767:
                                        v3 = (v5 + v5)
                                    case 2147483647:
                                        var v12 int16 = -32768
                                        _ = v12
                                    case 42:
                                        var v13 uint8 = 127
                                        var v14 uint64 = 2
                                        var v15 uint16 = 65535
                                        if (false || false) {
                                            switch "" {
                                            case "":
                                                if true {
                                                    var v16 int32 = 1
                                                    fmt.Println(v14, v1, v15)
                                                    _ = v16
                                                }
                                            case "hello":
                                                if false {
                                                    switch fi7 {
                                                    case -566:
                                                        fmt.Println(v3)
                                                    case 42:
                                                        switch "\n\t" {
                                                        case "a":
                                                            var v17 bool = (!false)
                                                            v18 := GenS2{F0: (v13 - 127)}
                                                            _ = v17
                                                            _ = v18
                                                        default:
                                                            v5 -= v3
                                                        }
                                                    case -281:
                                                        switch fi7 {
                                                        case -9223372036854775808:
                                                            if (((0 <= (v15 * (^(v15)))) && (false && false)) || true) {
                                                                if (true && (!true)) {
                                                                    v2 = (v2 / ((2) | 1))
                                                                    var v19 uint16 = (v15 ^ v15)
                                                                    switch (fi7 % ((128) | 1)) {
                                                                    case 42:
                                                                        v20 := GenS1{F0: -496, F1: (v5 << (((v2 % ((697) | 1))) & 63)), F2: 753}
                                                                        _ = v20
                                                                    case 968:
                                                                        switch fi7 {
                                                                        case 244:
                                                                            v5 += -1
                                                                        }
                                                                        var v21 int32 = 209
                                                                        if (!false) {
                                                                            fmt.Println(v3)
                                                                            v22 := GenS2{F0: (^(v13))}
                                                                            v23 := GenS3{F0: (v14 | 0)}
                                                                            _ = v22
                                                                            _ = v23
                                                                        } else {
                                                                            var v24 string = "\n\t"
                                                                            switch 110 {
                                                                            case 127:
                                                                                v25 := v6[fi9]
                                                                                _ = v25
                                                                            case 9223372036854775807:
                                                                                v26 := v4["hello"]
                                                                                _ = v26
                                                                            case 630:
                                                                                fmt.Println(v13, v14)
                                                                            }
                                                                            _ = v24
                                                                        }
                                                                        _ = v21
                                                                    }
                                                                    _ = v19
                                                                } else {
                                                                    var v27 bool = false
                                                                    var v28 int16 = 1
                                                                    fmt.Println(v27, v5, v27)
                                                                    _ = v27
                                                                    _ = v28
                                                                }
                                                            } else {
                                                                switch ("" + ("a" + "foo bar")) {
                                                                case "\n\t":
                                                                    if false {
                                                                        v29 := GenS3{F0: (v14 % ((v14) | 1))}
                                                                        if len(v8) > 0 {
                                                                            v8[0] = fi9
                                                                        }
                                                                        var v30 string = ""
                                                                        v15 -= uint16(v14)
                                                                        var v31 int32 = 2147483647
                                                                        _ = v29
                                                                        _ = v30
                                                                        _ = v31
                                                                    } else {
                                                                        fmt.Println(v2, fi7, v13)
                                                                        v14 += (v14 % ((uint64(v13)) | 1))
                                                                        v32 := GenS1{F0: 1, F1: 505, F2: (v3 ^ v5)}
                                                                        _ = v32
                                                                    }
                                                                case "":
                                                                    if (2147483647 < 1) {
                                                                        var v33 int64 = (v3 ^ v3)
                                                                        v34 := v4["hello"]
                                                                        var v35 string = ("hello" + "hello")
                                                                        _ = v33
                                                                        _ = v34
                                                                        _ = v35
                                                                    }
                                                                case "hello":
                                                                    if (!(!(v5 != v3))) {
                                                                        v2 += v2
                                                                        if len(v8) > 0 {
                                                                            v8[0] = 42
                                                                        }
                                                                        var v36 int8 = 16
                                                                        if len(v8) >= 2 {
                                                                            v8 = v8[1:len(v8)]
                                                                        }
                                                                        if len(v8) >= 2 {
                                                                            v8 = v8[1:len(v8)]
                                                                        }
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            if ((^(v2)) > v2) {
                                                                                switch 127 {
                                                                                case 128:
                                                                                    var v37 float32 = 3.14159
                                                                                    _ = v37
                                                                                case 2147483647:
                                                                                    var v38 uint32 = 42
                                                                                    _ = v38
                                                                                case -1:
                                                                                    switch ((("a" + "\n\t") + "\n\t") + "a") {
                                                                                    case "":
                                                                                        v39 := GenS1{F0: 42, F1: 42, F2: v3}
                                                                                        _ = v39
                                                                                    case "a":
                                                                                        fmt.Println(v13)
                                                                                    case "foo bar":
                                                                                        v1 += v1
                                                                                    default:
                                                                                        v40 := v6[(fi9 * fi7)]
                                                                                        _ = v40
                                                                                    }
                                                                                }
                                                                                v3 = v3
                                                                                v13 += (v13 / (((v13 % ((v13) | 1))) | 1))
                                                                            }
                                                                        case "a":
                                                                            var v41 uint = uint(v13)
                                                                            _ = v41
                                                                        default:
                                                                            switch (^(fi7)) {
                                                                            case 1:
                                                                                v36 = (v36 * 100)
                                                                            }
                                                                        }
                                                                        var v42 uint16 = (v15 * (v15 * v15))
                                                                        var v43 int64 = v3
                                                                        var v44 float64 = 1e10
                                                                        if (uint16(v13) == (v42 ^ v42)) {
                                                                            var v45 uint64 = v14
                                                                            v46 := GenS3{F0: uint64(v2)}
                                                                            if (((-2.25 + ((((-2.25 + float32(1.5)) - 1e10) * 3.14159) * -2.25)) + 1e10) > 1e10) {
                                                                                var v47 uint64 = v45
                                                                                v48 := GenS4{F0: (^(v14))}
                                                                                fmt.Println(v44, v2, fi9)
                                                                                _ = v47
                                                                                _ = v48
                                                                            } else {
                                                                                if (0 >= 42) {
                                                                                    fmt.Println(v36, v3, fi9)
                                                                                    v49 := GenS5{F0: (v13 >> (((v2 - v2)) & 63))}
                                                                                    _ = v49
                                                                                }
                                                                                v43 -= -525
                                                                            }
                                                                            _ = v45
                                                                            _ = v46
                                                                        } else {
                                                                            if (255 > 1) {
                                                                                _ = len(v8)
                                                                            } else {
                                                                                var v50 float64 = float64(1.5)
                                                                                _ = v50
                                                                            }
                                                                            var v51 int8 = v1
                                                                            v2 -= v2
                                                                            switch "foo bar" {
                                                                            case "foo bar":
                                                                                v51 = int8(fi9)
                                                                            case "a":
                                                                                switch (fi7 / (((-(fi7))) | 1)) {
                                                                                case 42:
                                                                                    if (v2 == v2) {
                                                                                        if true {
                                                                                            v43 = v43
                                                                                            delete(v6, (-(fi7)))
                                                                                            v52 := GenS1{F0: 1000, F1: (v43 - -1), F2: (v3 >> ((v2) & 63))}
                                                                                            if (((-1 >= 42) && true) || false) {
                                                                                                v3 += (v43 % ((-447) | 1))
                                                                                                if true {
                                                                                                    fmt.Println(v5, v36)
                                                                                                    v53 := GenS6{F0: fi9, F1: 1, F2: v2}
                                                                                                    var v54 uint8 = v13
                                                                                                    v15 -= (^(v15))
                                                                                                    switch -9223372036854775808 {
                                                                                                    case 42:
                                                                                                        switch "foo bar" {
                                                                                                        case "hello":
                                                                                                            v55 := GenS6{F0: fi7, F1: v14, F2: 2}
                                                                                                            _ = v55
                                                                                                        }
                                                                                                    case 32768:
                                                                                                        v56 := v4["hello"]
                                                                                                        _ = v56
                                                                                                    case -926:
                                                                                                        v15 -= v42
                                                                                                    default:
                                                                                                        var v57 bool = (383 > 535)
                                                                                                        _ = v57
                                                                                                    }
                                                                                                    _ = v53
                                                                                                    _ = v54
                                                                                                }
                                                                                                if ("hello" == ((("a" + ("" + ("a" + ""))) + "foo bar") + "a")) {
                                                                                                    v58 := GenS1{F0: 32767, F1: -711, F2: 0}
                                                                                                    var v59 int16 = 0
                                                                                                    _ = v58
                                                                                                    _ = v59
                                                                                                }
                                                                                                var v60 string = "\n\t"
                                                                                                if true {
                                                                                                    var v61 bool = ("foo bar" < (v60 + v60))
                                                                                                    _ = v61
                                                                                                }
                                                                                                _ = v60
                                                                                            }
                                                                                            _ = v52
                                                                                        }
                                                                                        v62 := GenS4{F0: v14}
                                                                                        var v63 int = (fi7 | fi9)
                                                                                        _ = v62
                                                                                        _ = v63
                                                                                    }
                                                                                case 486:
                                                                                    if true {
                                                                                        switch fi7 {
                                                                                        case 9223372036854775807:
                                                                                            var v64 int64 = 760
                                                                                            _ = v64
                                                                                        case 287:
                                                                                            if true {
                                                                                                switch -1 {
                                                                                                case 127:
                                                                                                    if false {
                                                                                                        v65 := v4[("\n\t" + ("\n\t" + "\n\t"))]
                                                                                                        var v66 int64 = (v43 >> (((v2 + v2)) & 63))
                                                                                                        _ = v65
                                                                                                        _ = v66
                                                                                                    }
                                                                                                case 9223372036854775807:
                                                                                                    switch "" {
                                                                                                    case "a":
                                                                                                        var v67 int32 = int32(v14)
                                                                                                        _ = v67
                                                                                                    case "":
                                                                                                        if len(v8) > 0 {
                                                                                                            v8[0] = fi7
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        if false {
                                                                                                            v51 -= v51
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    if true {
                                                                                                        v6[(fi9 >> ((v2) & 63))] = "foo bar"
                                                                                                    }
                                                                                                }
                                                                                                v68 := GenS5{F0: (v13 & v13)}
                                                                                                _ = v68
                                                                                            } else {
                                                                                                if (fi9 < 400) {
                                                                                                    _ = len(v8)
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            switch "foo bar" {
                                                                                            case "foo bar":
                                                                                                var v69 float32 = 3.14159
                                                                                                _ = v69
                                                                                            case "a":
                                                                                                v14 -= 2
                                                                                            }
                                                                                        }
                                                                                        _ = len(v6)
                                                                                    } else {
                                                                                        var v70 int8 = int8(v13)
                                                                                        v71 := GenS7{F0: 42, F1: 42, F2: (v13 << (((v2 - uint(v15))) & 63))}
                                                                                        _ = v70
                                                                                        _ = v71
                                                                                    }
                                                                                case 127:
                                                                                    v72 := GenS6{F0: -9223372036854775808, F1: 808, F2: 4294967295}
                                                                                    _ = v72
                                                                                }
                                                                            default:
                                                                                v73 := GenS8{F0: "hello"}
                                                                                _ = v73
                                                                            }
                                                                            _ = v51
                                                                        }
                                                                        _ = v36
                                                                        _ = v42
                                                                        _ = v43
                                                                        _ = v44
                                                                    }
                                                                }
                                                            }
                                                        case 42:
                                                            var v74 bool = (false && false)
                                                            _ = v74
                                                        case -953:
                                                            v75 := GenS8{F0: "\n\t"}
                                                            _ = v75
                                                        }
                                                    }
                                                    switch fi9 {
                                                    case 128:
                                                        fmt.Println(v13, v13)
                                                    case 551:
                                                        v5 = v5
                                                    case 2147483647:
                                                        var v76 float64 = (float64((1e10 + 3.14159)) * (((0.1 * 1.5) * 0.1) * 0.0))
                                                        _ = v76
                                                    default:
                                                        switch fi7 {
                                                        case 42:
                                                            switch (("" + "\n\t") + "a") {
                                                            case "hello":
                                                                v15 += (v15 - (v15 & 295))
                                                            case "foo bar":
                                                                delete(v6, fi7)
                                                            }
                                                        case 0:
                                                            v77 := GenS9{F0: 255, F1: (^(v13)), F2: (fi9 | fi7)}
                                                            _ = v77
                                                        case -687:
                                                            v78 := GenS7{F0: 2147483647, F1: 0, F2: v13}
                                                            _ = v78
                                                        }
                                                    }
                                                    var v79 float32 = 0.0
                                                    _ = v79
                                                } else {
                                                    var v80 int64 = 550
                                                    v81 := GenS5{F0: v13}
                                                    fmt.Println(v80)
                                                    _ = v80
                                                    _ = v81
                                                }
                                            default:
                                                delete(v4, "foo bar")
                                            }
                                            switch 4294967295 {
                                            case 9223372036854775807:
                                                fmt.Println(v13, fi9, v14)
                                            case 0:
                                                v15 = (v15 | (v15 >> ((0) & 63)))
                                            }
                                        }
                                        v82 := GenS10{F0: -1, F1: (-(fi7)), F2: (fi9 - (fi9 >> (((v2 >> ((uint(v14)) & 63))) & 63)))}
                                        _ = v13
                                        _ = v14
                                        _ = v15
                                        _ = v82
                                    default:
                                        fmt.Println(fi9)
                                    }
                                }
                                _ = fi9
                            }
                        default:
                            if (v2 > (v2 + 42)) {
                                fmt.Println(v3, v3, v5)
                                v5 = (v3 | v3)
                                fmt.Println(v1, fi7)
                            } else {
                                _ = len(v8)
                                fmt.Println(v2, v5, fi7)
                            }
                        }
                        v2 = (v2 * v2)
                        if true {
                            v83 := GenS9{F0: 0, F1: 2, F2: fi7}
                            if ((true || false) || (v2 != v2)) {
                                v84 := GenS11{F0: "\n\t", F1: 65535, F2: 230}
                                for fi85 := 0; fi85 < 3; fi85++ {
                                    var v86 int64 = 359
                                    v87 := GenS12{F0: "foo bar"}
                                    var v88 int8 = 1
                                    _ = fi85
                                    _ = v86
                                    _ = v87
                                    _ = v88
                                }
                                _ = v84
                            } else {
                                if true {
                                    fmt.Println(v2, v3)
                                }
                                var v89 int = (fi7 | fi7)
                                v8 = append(v8, (v89 & (v89 % ((fi7) | 1))))
                                _ = v89
                            }
                            _ = v83
                        } else {
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            fmt.Println(v2)
                        }
                    }
                    if true {
                        fmt.Println(v5, v1, v1)
                    }
                }
                fmt.Println(fi7)
                _ = v8
            } else {
                fmt.Println(v3)
            }
            v90 := GenS6{F0: fi7, F1: 740, F2: 255}
            _ = v90
        } else {
            if (((42 != 0) || true) && false) {
                fmt.Println(v1, fi7, v2)
                for fi91 := 0; fi91 < 9; fi91++ {
                    switch 2147483647 {
                    case 1:
                        var v92 uint64 = 2
                        _ = v92
                    case -1:
                        if true {
                            v93 := GenS4{F0: 293}
                            _ = v93
                        }
                    case -9223372036854775808:
                        fmt.Println(v1, fi7)
                    }
                    var v94 int64 = v3
                    _ = fi91
                    _ = v94
                }
                if true {
                    fmt.Println(v2)
                }
            }
            if ((fi7 * fi7) > (fi7 >> (((v2 | v2)) & 63))) {
                for fi95 := 0; fi95 < 17; fi95++ {
                    v96 := v6[(fi7 * 9223372036854775807)]
                    if (!(!true)) {
                        if true {
                            if (109 >= v1) {
                                switch fi95 {
                                case 393:
                                    switch (^(fi7)) {
                                    case 9223372036854775807:
                                        if false {
                                            var v97 bool = true
                                            v98 := GenS11{F0: v96, F1: 32767, F2: 2147483647}
                                            v99 := GenS13{F0: 734, F1: (1e10 <= (((1e10 / (1.5 * (1e10 - 0.1))) / 3.14159) * 0.0)), F2: 0}
                                            _ = v97
                                            _ = v98
                                            _ = v99
                                        } else {
                                            switch (v96 + v96) {
                                            case "":
                                                v100 := []int{225, (fi95 * (^(fi95)))}
                                                _ = v100
                                            case "\n\t":
                                                if ((v2 / (((v2 * (v2 ^ (v2 / ((v2) | 1))))) | 1)) == (v2 << (((v2 - v2)) & 63))) {
                                                    v101 := []int{}
                                                    v4[((v96 + "foo bar") + v96)] = fi95
                                                    v3 = 391
                                                    switch "" {
                                                    case "\n\t":
                                                        fmt.Println(fi7, fi7)
                                                    }
                                                    v5 -= (v5 ^ 42)
                                                    _ = v101
                                                }
                                            case "hello":
                                                v102 := GenS14{F0: v1, F1: 0}
                                                _ = v102
                                            default:
                                                if (32767 == 32767) {
                                                    switch v96 {
                                                    case "a":
                                                        delete(v4, ("foo bar" + "hello"))
                                                    case "hello":
                                                        var v103 int16 = -703
                                                        _ = v103
                                                    case "":
                                                        v3 = (v5 / ((1) | 1))
                                                    default:
                                                        v104 := v4[(v96 + v96)]
                                                        _ = v104
                                                    }
                                                    var v105 uint32 = 646
                                                    _ = v105
                                                } else {
                                                    v5 -= v3
                                                    v5 = v5
                                                }
                                            }
                                            v106 := []int{fi95, 485}
                                            switch 32767 {
                                            case -457:
                                                if (255 < 58) {
                                                    v1 -= 100
                                                    v107 := GenS15{F0: 18446744073709551615}
                                                    _ = v107
                                                }
                                            case 32767:
                                                v96 = v96
                                            case -711:
                                                v108 := GenS7{F0: 842, F1: 42, F2: 1}
                                                _ = v108
                                            default:
                                                switch ((("" + v96) + v96) + ("\n\t" + ("foo bar" + ("\n\t" + v96)))) {
                                                case "a":
                                                    if (!(!false)) {
                                                        var v109 bool = (false && true)
                                                        v109 = (v109 && (65535 != 74))
                                                        if true {
                                                            switch "" {
                                                            case "foo bar":
                                                                v110 := v4[("foo bar" + v96)]
                                                                _ = v110
                                                            case "hello":
                                                                var v111 uint32 = 65535
                                                                _ = v111
                                                            default:
                                                                switch (fi95 | (fi95 / ((-236) | 1))) {
                                                                case 9223372036854775807:
                                                                    var v112 int32 = 1
                                                                    _ = v112
                                                                case 387:
                                                                    v113 := GenS12{F0: v96}
                                                                    _ = v113
                                                                case 128:
                                                                    if len(v106) > 0 {
                                                                        v106[0] = (fi95 - 128)
                                                                    }
                                                                default:
                                                                    v114 := GenS16{F0: v2, F1: 65535}
                                                                    _ = v114
                                                                }
                                                            }
                                                            v109 = (fi95 == fi7)
                                                            var v115 string = (v96 + v96)
                                                            _ = v115
                                                        } else {
                                                            v116 := GenS14{F0: v1, F1: 42}
                                                            _ = v116
                                                        }
                                                        _ = v109
                                                    }
                                                case "\n\t":
                                                    v106 = append(v106, 9223372036854775807)
                                                }
                                            }
                                            _ = v106
                                        }
                                    case -2147483648:
                                        if false {
                                            if false {
                                                _ = len(v6)
                                                var v117 int8 = v1
                                                if (2 != 42) {
                                                    v118 := []int{94, (fi7 - -2147483648), (fi7 % ((fi95) | 1)), (fi7 % ((-277) | 1))}
                                                    _ = v118
                                                } else {
                                                    var v119 uint = 2
                                                    v120 := []int{756, (^(fi7)), fi95}
                                                    _ = v119
                                                    _ = v120
                                                }
                                                _ = v117
                                            } else {
                                                var v121 int = 4294967295
                                                v96 = "hello"
                                                if (false || (!true)) {
                                                    v5 -= (v3 & (-(v5)))
                                                    if (-1 > 1) {
                                                        if false {
                                                            var v122 int16 = 1
                                                            switch fi95 {
                                                            case 128:
                                                                v123 := []int{v121, 169}
                                                                _ = v123
                                                            case -2147483648:
                                                                v96 = v96
                                                            default:
                                                                var v124 int8 = (-(v1))
                                                                _ = v124
                                                            }
                                                            var v125 uint32 = 42
                                                            _ = v122
                                                            _ = v125
                                                        } else {
                                                            switch v96 {
                                                            case "hello":
                                                                v3 += -9223372036854775808
                                                            case "":
                                                                delete(v4, (v96 + "hello"))
                                                            case "a":
                                                                if (int(v2) == int(v2)) {
                                                                    var v126 uint8 = 42
                                                                    switch (v96 + "hello") {
                                                                    case "":
                                                                        v127 := GenS11{F0: v96, F1: 255, F2: 2147483647}
                                                                        _ = v127
                                                                    case "hello":
                                                                        if false {
                                                                            var v128 int = int(v2)
                                                                            v128 = 834
                                                                            v5 = (v3 * v5)
                                                                            v129 := GenS2{F0: (v126 << ((v2) & 63))}
                                                                            _ = v128
                                                                            _ = v129
                                                                        } else {
                                                                            if (898 >= 0) {
                                                                                v1 = v1
                                                                                v1 += int8(v126)
                                                                                switch "" {
                                                                                case "foo bar":
                                                                                    v130 := GenS13{F0: (v2 % ((v2) | 1)), F1: ((v121 >= 900) && true), F2: 597}
                                                                                    if (-1 <= 1) {
                                                                                        if (!false) {
                                                                                            switch fi7 {
                                                                                            case 32768:
                                                                                                v131 := []int{fi95, 128, (fi95 + (fi7 | v121))}
                                                                                                _ = v131
                                                                                            case 2147483647:
                                                                                                if true {
                                                                                                    v121 -= (fi7 | (-(v121)))
                                                                                                    var v132 int = int(v1)
                                                                                                    _ = v132
                                                                                                }
                                                                                            }
                                                                                            v126 += (v126 >> (((v2 >> ((v2) & 63))) & 63))
                                                                                        } else {
                                                                                            fmt.Println(v96, v96)
                                                                                        }
                                                                                        fmt.Println(v126, v126)
                                                                                        var v133 uint8 = 0
                                                                                        switch -9223372036854775808 {
                                                                                        case 128:
                                                                                            v96 = ("hello" + ("\n\t" + (((("\n\t" + (v96 + ("" + (((((v96 + "\n\t") + (v96 + ((v96 + ((((v96 + "") + (v96 + (v96 + (v96 + v96)))) + ("" + v96)) + (("foo bar" + v96) + v96))) + v96))) + v96) + "foo bar") + "\n\t")))) + (("\n\t" + "hello") + ("hello" + v96))) + "foo bar") + v96)))
                                                                                        case -849:
                                                                                            v134 := GenS11{F0: v96, F1: 2, F2: 4294967295}
                                                                                            _ = v134
                                                                                        case 127:
                                                                                            if (!false) {
                                                                                                fmt.Println(v5, v3)
                                                                                                v135 := GenS5{F0: v126}
                                                                                                v136 := GenS17{F0: v5, F1: true, F2: (v1 * 70)}
                                                                                                v96 = v96
                                                                                                v126 += v133
                                                                                                _ = v135
                                                                                                _ = v136
                                                                                            }
                                                                                        }
                                                                                        v121 -= 42
                                                                                        _ = v133
                                                                                    }
                                                                                    _ = v130
                                                                                default:
                                                                                    var v137 uint64 = 42
                                                                                    _ = v137
                                                                                }
                                                                            } else {
                                                                                v2 = 167
                                                                            }
                                                                        }
                                                                    case "a":
                                                                        v138 := []int{v121, v121, v121, (v121 / ((128) | 1))}
                                                                        _ = v138
                                                                    }
                                                                    _ = v126
                                                                }
                                                                v121 = (v121 << ((182) & 63))
                                                            }
                                                            var v139 float64 = (0.0 - 3.14159)
                                                            v139 -= 0.1
                                                            _ = v139
                                                        }
                                                    }
                                                }
                                                _ = v121
                                            }
                                            var v140 string = (("foo bar" + v96) + (v96 + ""))
                                            _ = v140
                                        } else {
                                            var v141 uint = v2
                                            v141 += 294
                                            v142 := GenS18{F0: -1, F1: ""}
                                            _ = v141
                                            _ = v142
                                        }
                                    default:
                                        if false {
                                            v2 += v2
                                        }
                                        v2 = v2
                                        if (true && ((!false) && true)) {
                                            v143 := []int{}
                                            switch fi95 {
                                            case 309:
                                                if (true || (4294967295 == 2147483647)) {
                                                    switch (fi95 * fi95) {
                                                    case 1:
                                                        var v144 int32 = 42
                                                        _ = v144
                                                    case 865:
                                                        var v145 bool = (v96 < "")
                                                        _ = v145
                                                    case -2147483648:
                                                        v96 = v96
                                                        v3 -= int64(v1)
                                                        v146 := GenS18{F0: -32768, F1: "a"}
                                                        _ = v146
                                                    }
                                                    v6[fi7] = "foo bar"
                                                } else {
                                                    if true {
                                                        fmt.Println(v3, v1, v96)
                                                        delete(v6, 128)
                                                    } else {
                                                        v5 -= 1
                                                        v147 := GenS19{F0: fi7, F1: v96, F2: 2147483647}
                                                        _ = v147
                                                    }
                                                }
                                            case 884:
                                                delete(v4, v96)
                                            case 731:
                                                if false {
                                                    v5 = v3
                                                    v148 := v4[((v96 + v96) + "a")]
                                                    _ = v148
                                                }
                                            }
                                            fmt.Println(v5)
                                            _ = v143
                                        } else {
                                            v149 := []int{(fi7 % ((int(v1)) | 1))}
                                            if len(v149) > 0 {
                                                v149[0] = (fi7 * (^(fi95)))
                                            }
                                            _ = v149
                                        }
                                    }
                                case -1:
                                    v1 = (v1 / (((v1 ^ v1)) | 1))
                                case -2147483648:
                                    v1 -= 42
                                default:
                                    v96 = v96
                                }
                                var v150 uint = v2
                                _ = v150
                            }
                            v151 := GenS7{F0: 42, F1: 0, F2: 255}
                            _ = v151
                        }
                    }
                    v5 -= v5
                    switch v96 {
                    case "hello":
                        var v152 int64 = (v3 + (v5 + int64(fi7)))
                        _ = v152
                    case "foo bar":
                        fmt.Println(v5, v2)
                    case "a":
                        var v153 float64 = 0.1
                        _ = v153
                    }
                    v154 := GenS15{F0: 764}
                    _ = fi95
                    _ = v154
                    _ = v96
                }
                var v155 int16 = -990
                _ = v155
            } else {
                v156 := GenS18{F0: 1000, F1: "\n\t"}
                _ = v156
            }
        }
        var v157 float64 = 0.0
        fmt.Println(fi7, v2)
        _ = fi7
        _ = v157
    }
    v158 := GenS2{F0: 9}
    v159 := []int{}
    v160 := GenS13{F0: 442, F1: ((0.1 * ((3.14159 * 0.1) / 3.14159)) > 0.1), F2: 524}
    v5 = -1
    v1 -= v1
    var v161 uint64 = 42
    var v162 string = (("\n\t" + (("foo bar" + "") + ("hello" + ""))) + "hello")
    v163 := GenS10{F0: 196, F1: -1, F2: 777}
    delete(v4, "hello")
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v159:", len(v159), v159)
    fmt.Println("v161:", v161)
    fmt.Println("v162:", v162)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v158
    _ = v159
    _ = v160
    _ = v161
    _ = v162
    _ = v163
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
