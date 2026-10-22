package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32; F1 int16; F2 bool }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 string; F1 int32 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int32; F1 uint64; F2 uint }
type GenS12 struct { F0 bool; F1 int32; F2 uint64 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint; F1 int16; F2 int64 }
type GenS15 struct { F0 int64; F1 int32; F2 int8 }
type GenS16 struct { F0 string; F1 int8 }
type GenS17 struct { F0 int8; F1 uint8; F2 uint }
type GenS18 struct { F0 int32; F1 uint }
type GenS19 struct { F0 uint64; F1 uint8; F2 uint }
type GenS20 struct { F0 int32; F1 uint16 }

func main() {
    v1 := make(map[string]bool)
    var v2 bool = false
    v3 := []int{67}
    v4 := make(map[int]string)
    var v5 int64 = -505
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: -60}
        for fi8 := 0; fi8 < 1; fi8++ {
            v3 = append(v3, (-(fi8)))
            _ = fi8
        }
        v9 := GenS2{F0: 2}
        _ = fi6
        _ = v7
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 uint = 416
        _ = fi10
        _ = v11
    }
    v5 = v5
    var v12 uint8 = 2
    if ((0 > 0) || v2) {
        var v13 uint8 = (v12 - 0)
        for fi14 := 0; fi14 < 13; fi14++ {
            switch int(v13) {
            case -12:
                if len(v3) > 0 {
                    v3[0] = 0
                }
            }
            _ = fi14
        }
        _ = v13
    }
    v15 := GenS3{F0: (-(v5)), F1: 4294967295, F2: v12}
    v2 = false
    for fi16 := 0; fi16 < 3; fi16++ {
        if v2 {
            switch fi16 {
            case 2147483647:
                fmt.Println(v5)
            default:
                v17 := GenS2{F0: (v12 - 0)}
                _ = v17
            }
            fmt.Println(v2)
            if (501 > 65535) {
                var v18 int32 = 2147483647
                if (int16(v5) != 1000) {
                    if len(v3) > 0 {
                        v3[0] = int(v18)
                    }
                    v18 = 1
                    fmt.Println(v2)
                } else {
                    var v19 int16 = -32768
                    _ = v19
                }
                _ = v18
            } else {
                v4[-661] = ("hello" + "")
                v5 -= -9223372036854775808
                v12 -= v12
            }
        } else {
            if false {
                for fi20 := 0; fi20 < 2; fi20++ {
                    v5 -= v5
                    _ = fi20
                }
                delete(v1, ("foo bar" + "foo bar"))
                var v21 int16 = 32767
                _ = v21
            } else {
                switch "a" {
                case "hello":
                    fmt.Println(v2, fi16)
                case "\n\t":
                    v5 = v5
                }
                var v22 int8 = 127
                v2 = (v2 && v2)
                _ = v22
            }
            if (((0.1 + (0.1 + float32((3.14159 / 1e10)))) + 0.1) == -2.25) {
                for fi23 := 0; fi23 < 19; fi23++ {
                    fmt.Println(v2, fi16, v12)
                    if (2147483648 != 4294967295) {
                        if (1.5 > (0.0 * 1e10)) {
                            if ((127 >= 1) || ((v5 ^ (v5 / ((v5) | 1))) == (v5 << ((4294967295) & 63)))) {
                                if v2 {
                                    v5 = v5
                                    switch fi16 {
                                    case 2147483647:
                                        if true {
                                            var v24 uint8 = (v12 ^ (v12 >> ((255) & 63)))
                                            v25 := GenS4{F0: 42}
                                            _ = v24
                                            _ = v25
                                        } else {
                                            switch "" {
                                            case "foo bar":
                                                switch 604 {
                                                case 379:
                                                    switch fi23 {
                                                    case 131:
                                                        v2 = ((v5 << ((161) & 63)) < v5)
                                                    }
                                                case -723:
                                                    fmt.Println(fi23)
                                                case -236:
                                                    var v26 int64 = v5
                                                    _ = v26
                                                default:
                                                    switch ("foo bar" + "") {
                                                    case "a":
                                                        v2 = ((false || (42 != int16(v5))) && (0 == 0))
                                                    default:
                                                        delete(v1, ("hello" + ""))
                                                    }
                                                }
                                            case "a":
                                                switch fi23 {
                                                case 0:
                                                    if len(v3) > 0 {
                                                        v3[0] = fi23
                                                    }
                                                case -1:
                                                    var v27 uint8 = (v12 - 128)
                                                    _ = v27
                                                }
                                            case "":
                                                v2 = (!((697 != 18446744073709551615) || (42 < 255)))
                                            }
                                            switch "a" {
                                            case "hello":
                                                if len(v3) > 0 {
                                                    v3[0] = (-(fi16))
                                                }
                                            case "foo bar":
                                                delete(v4, (fi16 - (fi16 - (^(fi23)))))
                                            case "a":
                                                _ = len(v1)
                                            }
                                        }
                                    case -111:
                                        var v28 uint8 = v12
                                        _ = v28
                                    case 69:
                                        switch "\n\t" {
                                        case "":
                                            switch "\n\t" {
                                            case "foo bar":
                                                if len(v3) > 0 {
                                                    v3[0] = fi23
                                                }
                                                if false {
                                                    var v29 uint64 = 1
                                                    _ = v29
                                                } else {
                                                    v30 := v1["foo bar"]
                                                    var v31 uint64 = uint64(v12)
                                                    var v32 float32 = 1e10
                                                    _ = v30
                                                    _ = v31
                                                    _ = v32
                                                }
                                                fmt.Println(fi16)
                                            }
                                        case "foo bar":
                                            switch (fi16 % ((128) | 1)) {
                                            case -1:
                                                v33 := v4[(-(fi23))]
                                                _ = v33
                                            case -2147483648:
                                                if v2 {
                                                    v12 = v12
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    v5 -= (v5 & v5)
                                                    if len(v3) > 0 {
                                                        v3[0] = fi16
                                                    }
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = fi23
                                                    }
                                                    switch "foo bar" {
                                                    case "hello":
                                                        fmt.Println(fi23, v2, v5)
                                                    case "foo bar":
                                                        var v34 int8 = 1
                                                        _ = v34
                                                    case "\n\t":
                                                        if v2 {
                                                            if (v12 > v12) {
                                                                fmt.Println(v12, fi16)
                                                                var v35 int64 = v5
                                                                _ = v35
                                                            } else {
                                                                v36 := v4[fi16]
                                                                var v37 uint8 = (^(v12))
                                                                _ = v36
                                                                _ = v37
                                                            }
                                                            v38 := v4[fi16]
                                                            _ = v38
                                                        }
                                                    default:
                                                        var v39 uint = 18446744073709551615
                                                        _ = v39
                                                    }
                                                    if (-2.25 >= 0.0) {
                                                        v40 := GenS5{F0: 176}
                                                        if v2 {
                                                            v41 := GenS6{F0: 236, F1: 723, F2: false}
                                                            var v42 float32 = (float32((1.5 / 0.1)) + 1e10)
                                                            _ = v41
                                                            _ = v42
                                                        }
                                                        var v43 uint32 = 2147483647
                                                        _ = v40
                                                        _ = v43
                                                    }
                                                }
                                            case -849:
                                                fmt.Println(fi16)
                                            }
                                        }
                                    default:
                                        if ((float32(1e10) - ((0.1 - (0.0 + 1e10)) + 0.0)) <= 0.1) {
                                            if ((2147483647 < -2147483648) && false) {
                                                switch fi23 {
                                                case -2147483648:
                                                    if (!false) {
                                                        _ = len(v4)
                                                        v44 := v4[fi16]
                                                        v2 = ((255 == 810) && true)
                                                        _ = v44
                                                    }
                                                case 42:
                                                    if false {
                                                        if true {
                                                            fmt.Println(v2, fi23, fi23)
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                var v45 float32 = 0.1
                                                                v46 := GenS6{F0: 876, F1: 1000, F2: false}
                                                                _ = v45
                                                                _ = v46
                                                            case "a":
                                                                v47 := GenS7{F0: -20}
                                                                _ = v47
                                                            case "":
                                                                if v2 {
                                                                    if v2 {
                                                                        var v48 uint = 42
                                                                        v12 += v12
                                                                        var v49 float64 = (3.14159 * -2.25)
                                                                        _ = v48
                                                                        _ = v49
                                                                    }
                                                                }
                                                            }
                                                            var v50 uint16 = 368
                                                            switch ("" + "\n\t") {
                                                            case "foo bar":
                                                                var v51 uint8 = 255
                                                                _ = v51
                                                            case "\n\t":
                                                                fmt.Println(v50)
                                                            case "hello":
                                                                var v52 string = "\n\t"
                                                                _ = v52
                                                            default:
                                                                v50 = v50
                                                            }
                                                            if v2 {
                                                                if v2 {
                                                                    v53 := GenS6{F0: 2147483647, F1: 1, F2: (32767 <= 42)}
                                                                    switch (fi23 + (^(fi16))) {
                                                                    case 9223372036854775807:
                                                                        if (42 != 4294967295) {
                                                                            var v54 float32 = 3.14159
                                                                            v55 := GenS4{F0: 42}
                                                                            var v56 int8 = -1
                                                                            fmt.Println(fi16, fi23, v5)
                                                                            switch fi16 {
                                                                            case 128:
                                                                                v5 = v5
                                                                            }
                                                                            _ = v54
                                                                            _ = v55
                                                                            _ = v56
                                                                        } else {
                                                                            var v57 float64 = (float64(float32(1.5)) + ((1.5 / 0.1) / 1.0))
                                                                            v58 := v4[(fi23 | fi16)]
                                                                            v5 = v5
                                                                            _ = v57
                                                                            _ = v58
                                                                        }
                                                                        if false {
                                                                            if (v2 && (0 != -74)) {
                                                                                if ((-366 == 42) && v2) {
                                                                                    v12 = v12
                                                                                    v50 += uint16(v12)
                                                                                    var v59 int64 = (v5 & (v5 + v5))
                                                                                    _ = v59
                                                                                } else {
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                    v12 = (v12 << ((65535) & 63))
                                                                                }
                                                                                v50 = v50
                                                                            } else {
                                                                                _ = len(v3)
                                                                                v60 := GenS8{F0: fi16}
                                                                                _ = v60
                                                                            }
                                                                        } else {
                                                                            if v2 {
                                                                                var v61 float32 = (1.5 + 1e10)
                                                                                var v62 uint = 1
                                                                                _ = v61
                                                                                _ = v62
                                                                            }
                                                                            var v63 float64 = (0.0 + ((0.0 - ((0.0 * 1e10) / -2.25)) * -2.25))
                                                                            v12 -= 1
                                                                            _ = v63
                                                                        }
                                                                        switch (("" + ("hello" + ((((("hello" + (("hello" + "\n\t") + "a")) + "foo bar") + ("\n\t" + "\n\t")) + "hello") + ("\n\t" + "hello")))) + ("\n\t" + "\n\t")) {
                                                                        case "a":
                                                                            switch "foo bar" {
                                                                            case "\n\t":
                                                                                delete(v1, "")
                                                                            case "foo bar":
                                                                                var v64 int64 = -152
                                                                                _ = v64
                                                                            case "a":
                                                                                v65 := v4[(fi16 + fi16)]
                                                                                _ = v65
                                                                            default:
                                                                                _ = len(v4)
                                                                            }
                                                                        default:
                                                                            fmt.Println(v2, fi23, v50)
                                                                        }
                                                                    case 769:
                                                                        var v66 int8 = 127
                                                                        _ = v66
                                                                    }
                                                                    v50 = v50
                                                                    if (v2 && (0 >= 42)) {
                                                                        v12 = (v12 % ((v12) | 1))
                                                                        fmt.Println(fi23, v5, fi23)
                                                                        var v67 uint64 = 18446744073709551615
                                                                        _ = v67
                                                                    }
                                                                    var v68 uint8 = v12
                                                                    v69 := GenS5{F0: 1}
                                                                    var v70 int16 = 0
                                                                    _ = v53
                                                                    _ = v68
                                                                    _ = v69
                                                                    _ = v70
                                                                }
                                                                switch (fi23 - -1) {
                                                                case -558:
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    var v71 float32 = float32(1e10)
                                                                    if (true || false) {
                                                                        _ = len(v3)
                                                                    } else {
                                                                        v72 := GenS2{F0: 0}
                                                                        v50 = v50
                                                                        switch (("hello" + (("hello" + "a") + "\n\t")) + "hello") {
                                                                        case "":
                                                                            if len(v3) > 0 {
                                                                                v3[0] = fi23
                                                                            }
                                                                        case "foo bar":
                                                                            var v73 float32 = ((0.1 / (float32(((1.5 / 1.0) / 1.0)) + float32(-2.25))) + 1.5)
                                                                            _ = v73
                                                                        case "a":
                                                                            switch (("hello" + "") + ("hello" + ("foo bar" + "hello"))) {
                                                                            case "a":
                                                                                var v74 uint8 = 127
                                                                                _ = v74
                                                                            case "":
                                                                                v71 = v71
                                                                            case "\n\t":
                                                                                v75 := GenS4{F0: 972}
                                                                                _ = v75
                                                                            default:
                                                                                if false {
                                                                                    var v76 uint64 = 2
                                                                                    v77 := GenS3{F0: -1, F1: v76, F2: (v12 >> ((42) & 63))}
                                                                                    _ = v76
                                                                                    _ = v77
                                                                                } else {
                                                                                    v50 = v50
                                                                                    var v78 uint16 = 2
                                                                                    fmt.Println(fi23, v2, v2)
                                                                                    _ = v78
                                                                                }
                                                                            }
                                                                        default:
                                                                            if v2 {
                                                                                v50 = 2
                                                                            } else {
                                                                                v79 := v4[(fi16 | (fi23 & (-(fi23))))]
                                                                                switch (v79 + "\n\t") {
                                                                                case "foo bar":
                                                                                    switch fi23 {
                                                                                    case -405:
                                                                                        fmt.Println(fi16, v79)
                                                                                    case 128:
                                                                                        v50 -= (v50 * v50)
                                                                                    case -1:
                                                                                        if (1e10 > (1e10 + (1e10 * ((1e10 / -2.25) + -2.25)))) {
                                                                                            delete(v4, (fi16 >> ((59) & 63)))
                                                                                        } else {
                                                                                            if (!v2) {
                                                                                                fmt.Println(v50)
                                                                                                if false {
                                                                                                    v80 := GenS4{F0: -1}
                                                                                                    var v81 uint16 = 2
                                                                                                    v71 = 3.14159
                                                                                                    v79 = v79
                                                                                                    v81 += v50
                                                                                                    if false {
                                                                                                        if true {
                                                                                                            fmt.Println(v79)
                                                                                                        }
                                                                                                        var v82 uint32 = 2
                                                                                                        _ = v82
                                                                                                    } else {
                                                                                                        var v83 bool = ((((3.14159 + (v71 - float32(1e10))) + 0.1) / v71) > float32(((-2.25 * float64(1.5)) / 1.0)))
                                                                                                        v84 := GenS9{F0: ("a" + ""), F1: 533}
                                                                                                        _ = v83
                                                                                                        _ = v84
                                                                                                    }
                                                                                                    var v85 int32 = -2147483648
                                                                                                    v86 := GenS7{F0: 0}
                                                                                                    _ = v80
                                                                                                    _ = v81
                                                                                                    _ = v85
                                                                                                    _ = v86
                                                                                                } else {
                                                                                                    v50 -= 65535
                                                                                                    v2 = false
                                                                                                    if (false || (!v2)) {
                                                                                                        if true {
                                                                                                            var v87 uint32 = 1
                                                                                                            switch fi23 {
                                                                                                            case 221:
                                                                                                                var v88 int = fi23
                                                                                                                _ = v88
                                                                                                            case 32767:
                                                                                                                var v89 uint8 = (^(v12))
                                                                                                                _ = v89
                                                                                                            }
                                                                                                            _ = v87
                                                                                                        }
                                                                                                        switch (^(fi23)) {
                                                                                                        case -96:
                                                                                                            switch fi23 {
                                                                                                            case 0:
                                                                                                                if v2 {
                                                                                                                    if len(v3) >= 2 {
                                                                                                                        v3 = v3[1:len(v3)]
                                                                                                                    }
                                                                                                                    v71 += 1.5
                                                                                                                    switch "foo bar" {
                                                                                                                    case "hello":
                                                                                                                        v90 := GenS2{F0: (v12 & 128)}
                                                                                                                        _ = v90
                                                                                                                    case "\n\t":
                                                                                                                        if (!(!v2)) {
                                                                                                                            var v91 int = (fi16 - fi16)
                                                                                                                            v92 := GenS9{F0: "hello", F1: 2147483647}
                                                                                                                            if (-58 >= -106) {
                                                                                                                                if (!false) {
                                                                                                                                    fmt.Println(v79, v12)
                                                                                                                                    switch v91 {
                                                                                                                                    case 9223372036854775807:
                                                                                                                                        v93 := GenS5{F0: 2147483648}
                                                                                                                                        var v94 float32 = (-2.25 / (v71 * v71))
                                                                                                                                        var v95 string = v79
                                                                                                                                        _ = v93
                                                                                                                                        _ = v94
                                                                                                                                        _ = v95
                                                                                                                                    default:
                                                                                                                                        var v96 int64 = v5
                                                                                                                                        _ = v96
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            }
                                                                                                                            _ = v91
                                                                                                                            _ = v92
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        fmt.Println(v79, fi16)
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    var v97 int32 = -314
                                                                                                                    if true {
                                                                                                                        _ = len(v3)
                                                                                                                        v2 = v2
                                                                                                                        var v98 uint64 = 72
                                                                                                                        _ = v98
                                                                                                                    } else {
                                                                                                                        if len(v3) > 0 {
                                                                                                                            v3[0] = -788
                                                                                                                        }
                                                                                                                        v97 += (v97 / ((-2147483648) | 1))
                                                                                                                        var v99 float64 = -2.25
                                                                                                                        fmt.Println(v71, v99, v50)
                                                                                                                        _ = v99
                                                                                                                    }
                                                                                                                    v5 -= v5
                                                                                                                    _ = v97
                                                                                                                }
                                                                                                            case -2147483648:
                                                                                                                delete(v1, "hello")
                                                                                                            case 798:
                                                                                                                var v100 int = (fi23 / (((fi23 / (((fi23 ^ (fi23 - fi16))) | 1))) | 1))
                                                                                                                _ = v100
                                                                                                            }
                                                                                                            var v101 bool = v2
                                                                                                            _ = v101
                                                                                                        default:
                                                                                                            v102 := v1[v79]
                                                                                                            _ = v102
                                                                                                        }
                                                                                                        v12 += (v12 - 204)
                                                                                                    }
                                                                                                }
                                                                                                var v103 uint8 = 42
                                                                                                _ = v103
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                case "\n\t":
                                                                                    var v104 uint16 = (v50 % ((1) | 1))
                                                                                    _ = v104
                                                                                case "hello":
                                                                                    v79 = "\n\t"
                                                                                default:
                                                                                    if (v12 <= v12) {
                                                                                        v50 = (v50 ^ (v50 / ((v50) | 1)))
                                                                                        fmt.Println(fi16, v50)
                                                                                    }
                                                                                }
                                                                                switch (^(fi16)) {
                                                                                case 1:
                                                                                    var v105 uint8 = v12
                                                                                    _ = v105
                                                                                case 17:
                                                                                    fmt.Println(v12, v79, v79)
                                                                                }
                                                                                _ = v79
                                                                            }
                                                                        }
                                                                        _ = v72
                                                                    }
                                                                    _ = v71
                                                                case 32767:
                                                                    if (true || v2) {
                                                                        fmt.Println(v2)
                                                                        switch 4294967295 {
                                                                        case 42:
                                                                            v2 = v2
                                                                        default:
                                                                            delete(v1, "foo bar")
                                                                            _ = len(v1)
                                                                        }
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        var v106 uint16 = v50
                                                                        var v107 uint8 = v12
                                                                        _ = v106
                                                                        _ = v107
                                                                    } else {
                                                                        v108 := GenS3{F0: (v5 % (((v5 ^ (-(v5)))) | 1)), F1: 1, F2: 255}
                                                                        _ = v108
                                                                    }
                                                                default:
                                                                    v109 := GenS4{F0: 0}
                                                                    _ = v109
                                                                }
                                                                if v2 {
                                                                    v2 = ((false && (188 != (v12 | v12))) || false)
                                                                }
                                                                v3 = append(v3, fi16)
                                                                if v2 {
                                                                    var v110 uint8 = 255
                                                                    _ = v110
                                                                } else {
                                                                    v3 = append(v3, (^(fi23)))
                                                                    if ((!(fi16 == 1)) && v2) {
                                                                        var v111 uint32 = 4294967295
                                                                        var v112 int64 = 42
                                                                        v5 -= 1
                                                                        fmt.Println(v2)
                                                                        v2 = true
                                                                        _ = v111
                                                                        _ = v112
                                                                    } else {
                                                                        fmt.Println(fi16, fi16, fi16)
                                                                    }
                                                                    fmt.Println(fi16, v12, v50)
                                                                }
                                                            } else {
                                                                v5 = (^(v5))
                                                                fmt.Println(v12)
                                                                var v113 int32 = -1
                                                                _ = v113
                                                            }
                                                            _ = v50
                                                        } else {
                                                            switch 32767 {
                                                            case 0:
                                                                v114 := GenS4{F0: -55}
                                                                _ = v114
                                                            case -1:
                                                                switch "a" {
                                                                case "a":
                                                                    var v115 float32 = 1.5
                                                                    _ = v115
                                                                case "foo bar":
                                                                    if (0 <= 88) {
                                                                        v4[(fi16 + (fi16 / (((fi16 | fi16)) | 1)))] = "hello"
                                                                        v2 = v2
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        var v116 uint32 = 0
                                                        _ = v116
                                                    }
                                                default:
                                                    v12 = (v12 * v12)
                                                }
                                                v12 -= (v12 * 42)
                                                v117 := GenS10{F0: 521}
                                                _ = v117
                                            }
                                            var v118 float32 = (1.5 * 1.5)
                                            _ = v118
                                        }
                                    }
                                    v2 = (255 <= 4294967295)
                                }
                                switch (^(fi23)) {
                                case 42:
                                    var v119 int32 = -2147483648
                                    _ = v119
                                case 1:
                                    v120 := GenS8{F0: fi16}
                                    _ = v120
                                case -896:
                                    fmt.Println(fi16, fi16)
                                default:
                                    if v2 {
                                        fmt.Println(fi23)
                                        v3 = append(v3, 805)
                                    } else {
                                        var v121 float64 = 1.5
                                        _ = v121
                                    }
                                }
                            } else {
                                v122 := GenS8{F0: (-(fi16))}
                                var v123 int16 = -32768
                                _ = v122
                                _ = v123
                            }
                        } else {
                            var v124 string = ("" + ("a" + ""))
                            var v125 float32 = 1.5
                            _ = v124
                            _ = v125
                        }
                        var v126 int8 = 1
                        _ = v126
                    } else {
                        var v127 int = (fi23 + fi16)
                        var v128 uint32 = 2147483647
                        _ = len(v3)
                        _ = v127
                        _ = v128
                    }
                    _ = fi23
                }
            }
        }
        for fi129 := 0; fi129 < 4; fi129++ {
            var v130 uint = 18446744073709551615
            switch "foo bar" {
            case "foo bar":
                switch "" {
                case "hello":
                    v131 := GenS11{F0: -781, F1: 2, F2: v130}
                    _ = v131
                case "":
                    v132 := GenS5{F0: 0}
                    _ = v132
                }
            case "a":
                var v133 uint = (v130 - v130)
                _ = v133
            case "hello":
                if true {
                    v1["\n\t"] = (v2 && (!(-32768 != -462)))
                } else {
                    v12 = (v12 | v12)
                    switch "hello" {
                    case "":
                        delete(v1, ("hello" + ""))
                    case "foo bar":
                        var v134 uint32 = 904
                        _ = v134
                    default:
                        fmt.Println(fi129, v5)
                    }
                    var v135 string = ("" + ("a" + "hello"))
                    _ = v135
                }
            }
            if (!(81 <= 6)) {
                var v136 string = ("" + "a")
                fmt.Println(v12)
                _ = v136
            }
            _ = fi129
            _ = v130
        }
        var v137 int32 = 2147483647
        _ = fi16
        _ = v137
    }
    switch 127 {
    case 1:
        fmt.Println(v5)
    case -2147483648:
        v138 := GenS4{F0: 42}
        _ = v138
    case -780:
        if (1 > 742) {
            v5 -= v5
            for fi139 := 0; fi139 < 15; fi139++ {
                for fi140 := 0; fi140 < 19; fi140++ {
                    v5 += (v5 + -9223372036854775808)
                    if false {
                        if len(v3) > 0 {
                            v3[0] = (fi140 % ((fi140) | 1))
                        }
                    } else {
                        var v141 bool = ((v2 || v2) && (219 >= fi139))
                        v142 := GenS2{F0: (v12 % ((v12) | 1))}
                        var v143 int = 4294967295
                        _ = v141
                        _ = v142
                        _ = v143
                    }
                    if v2 {
                        v2 = (!true)
                    }
                    _ = fi140
                }
                v5 -= -580
                _ = fi139
            }
            switch ("foo bar" + "foo bar") {
            case "a":
                var v144 int = -2147483648
                _ = v144
            case "":
                v145 := GenS12{F0: v2, F1: -457, F2: 4294967295}
                _ = v145
            }
        } else {
            var v146 float64 = -2.25
            v12 -= (v12 << ((42) & 63))
            if (((uint(v12) == 186) || (128 != 42)) || (v2 && false)) {
                if len(v3) > 0 {
                    v3[0] = 734
                }
                fmt.Println(v5, v146, v5)
            } else {
                if ((v5 + (^(v5))) < (v5 | v5)) {
                    v12 = (v12 | v12)
                    for fi147 := 0; fi147 < 15; fi147++ {
                        v146 = v146
                        var v148 int32 = 0
                        v146 -= float64((float32(v146) + (3.14159 / 0.1)))
                        _ = fi147
                        _ = v148
                    }
                    for fi149 := 0; fi149 < 15; fi149++ {
                        var v150 float32 = 1.5
                        v146 += (v146 / 0.1)
                        var v151 int16 = 18
                        _ = fi149
                        _ = v150
                        _ = v151
                    }
                }
                var v152 string = "a"
                _ = v152
            }
            _ = v146
        }
    }
    for fi153 := 0; fi153 < 19; fi153++ {
        var v154 uint32 = 42
        v2 = (!true)
        _ = fi153
        _ = v154
    }
    fmt.Println(v5, v5)
    v155 := GenS13{F0: -1}
    fmt.Println(v2)
    v156 := GenS7{F0: 127}
    if (v2 && (v2 && (v2 || (-32768 > 1000)))) {
        for fi157 := 0; fi157 < 7; fi157++ {
            v158 := GenS14{F0: 65535, F1: -1, F2: v5}
            switch (fi157 % (((fi157 + (fi157 * fi157))) | 1)) {
            case 2147483647:
                switch "\n\t" {
                case "hello":
                    _ = len(v3)
                case "\n\t":
                    var v159 int = fi157
                    _ = v159
                case "a":
                    var v160 bool = true
                    _ = v160
                }
            case -1:
                fmt.Println(v12, v12)
            }
            for fi161 := 0; fi161 < 10; fi161++ {
                var v162 float32 = (float32(1e10) - 0.0)
                var v163 int32 = 42
                switch "foo bar" {
                case "hello":
                    var v164 uint = 65535
                    _ = v164
                }
                switch (fi157 ^ (fi161 | (^(fi157)))) {
                case 32768:
                    v162 -= -2.25
                case 4294967295:
                    if (true || v2) {
                        v162 -= v162
                        if v2 {
                            switch "a" {
                            case "\n\t":
                                v3 = append(v3, (fi161 / ((fi157) | 1)))
                            case "hello":
                                var v165 uint32 = 2147483647
                                _ = v165
                            case "a":
                                var v166 uint8 = v12
                                _ = v166
                            }
                            v12 = v12
                            if (v2 && (v2 && (368 >= 18446744073709551615))) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v167 := GenS15{F0: -1, F1: (v163 & 2147483647), F2: -1}
                                v168 := GenS3{F0: 42, F1: 2, F2: v12}
                                switch "hello" {
                                case "\n\t":
                                    var v169 int8 = 1
                                    _ = v169
                                case "a":
                                    switch "" {
                                    case "":
                                        fmt.Println(v162)
                                    case "a":
                                        if v2 {
                                            var v170 int = (fi161 / ((fi157) | 1))
                                            v5 += -9223372036854775808
                                            _ = v170
                                        } else {
                                            v171 := GenS16{F0: (("\n\t" + "foo bar") + ""), F1: 68}
                                            var v172 uint16 = 1000
                                            v172 = (v172 | (v172 ^ 0))
                                            _ = v171
                                            _ = v172
                                        }
                                    default:
                                        var v173 int16 = 42
                                        _ = v173
                                    }
                                case "hello":
                                    var v174 int32 = v163
                                    _ = v174
                                default:
                                    fmt.Println(fi157, v163, v2)
                                }
                                v163 = v163
                                _ = v167
                                _ = v168
                            }
                        } else {
                            var v175 uint32 = 42
                            var v176 int64 = 0
                            var v177 uint8 = (^(v12))
                            _ = v175
                            _ = v176
                            _ = v177
                        }
                        v4[fi161] = ""
                        fmt.Println(fi157, fi157, fi157)
                    }
                    _ = len(v1)
                    v2 = false
                }
                _ = fi161
                _ = v162
                _ = v163
            }
            _ = fi157
            _ = v158
        }
        if v2 {
            v2 = (1 != -1)
            v5 += (v5 % ((-9223372036854775808) | 1))
            if (!((!v2) && true)) {
                v12 = v12
                var v178 float64 = 0.0
                _ = v178
            } else {
                v179 := v1[("a" + ("foo bar" + "\n\t"))]
                if v179 {
                    for fi180 := 0; fi180 < 1; fi180++ {
                        if true {
                            switch "a" {
                            case "hello":
                                fmt.Println(v5, v2, v2)
                            case "foo bar":
                                v5 -= 345
                            case "a":
                                v181 := GenS15{F0: (v5 - 9223372036854775807), F1: 1, F2: int8(v12)}
                                _ = v181
                            default:
                                if (!((!((1 <= -1) || true)) && v2)) {
                                    v12 -= (v12 | (v12 * (v12 % ((127) | 1))))
                                    var v182 uint32 = 4294967295
                                    if (!v2) {
                                        v183 := GenS9{F0: ("hello" + "hello"), F1: 2147483647}
                                        v182 = 63
                                        v5 = -224
                                        _ = v183
                                    } else {
                                        v184 := v1[("" + "a")]
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        switch "\n\t" {
                                        case "":
                                            if v2 {
                                                switch int(v5) {
                                                case -792:
                                                    v185 := GenS17{F0: int8(fi180), F1: (v12 % ((18) | 1)), F2: 255}
                                                    _ = v185
                                                default:
                                                    v186 := GenS4{F0: -1}
                                                    _ = v186
                                                }
                                                v184 = v184
                                                delete(v4, fi180)
                                            }
                                        }
                                        _ = v184
                                    }
                                    _ = v182
                                }
                            }
                            v1["hello"] = (1 < 0)
                            var v187 int32 = -667
                            _ = v187
                        }
                        for fi188 := 0; fi188 < 15; fi188++ {
                            switch ("a" + ("foo bar" + "foo bar")) {
                            case "hello":
                                v179 = (!(18446744073709551615 < 4294967295))
                            default:
                                v189 := GenS18{F0: 42, F1: 268}
                                _ = v189
                            }
                            v190 := GenS19{F0: 1, F1: (v12 + v12), F2: 1}
                            _ = fi188
                            _ = v190
                        }
                        _ = fi180
                    }
                    switch 42 {
                    case 9223372036854775807:
                        for fi191 := 0; fi191 < 19; fi191++ {
                            switch fi191 {
                            case -9223372036854775808:
                                v192 := GenS7{F0: -53}
                                _ = v192
                            default:
                                var v193 uint32 = 1
                                _ = v193
                            }
                            _ = fi191
                        }
                    case 128:
                        fmt.Println(v179)
                    case -877:
                        fmt.Println(v2)
                    }
                } else {
                    v194 := GenS5{F0: 2}
                    v5 -= v5
                    v12 = 128
                    _ = v194
                }
                _ = v179
            }
        } else {
            if false {
                var v195 int16 = 42
                _ = v195
            }
            var v196 uint32 = 4294967295
            if v2 {
                for fi197 := 0; fi197 < 4; fi197++ {
                    fmt.Println(v196, v5, fi197)
                    for fi198 := 0; fi198 < 9; fi198++ {
                        if v2 {
                            v199 := GenS10{F0: -478}
                            var v200 float64 = (0.1 / -2.25)
                            v201 := GenS19{F0: 636, F1: (v12 / ((128) | 1)), F2: 65535}
                            var v202 uint64 = 630
                            var v203 int64 = 42
                            _ = v199
                            _ = v200
                            _ = v201
                            _ = v202
                            _ = v203
                        }
                        var v204 float32 = (1.5 / (1e10 - 3.14159))
                        var v205 string = "\n\t"
                        v204 = 0.1
                        v205 = v205
                        _ = fi198
                        _ = v204
                        _ = v205
                    }
                    if (-2147483648 <= 42) {
                        if false {
                            v206 := GenS20{F0: 0, F1: 2}
                            v207 := GenS10{F0: 0}
                            var v208 uint = 42
                            _ = v206
                            _ = v207
                            _ = v208
                        }
                    }
                    _ = fi197
                }
                var v209 float64 = 1e10
                _ = v209
            } else {
                if true {
                    v210 := GenS4{F0: 42}
                    _ = v210
                } else {
                    fmt.Println(v2)
                }
                var v211 int16 = 32767
                _ = v211
            }
            _ = v196
        }
    } else {
        v12 -= (^(v12))
        var v212 uint16 = 32767
        _ = v212
    }
    switch "\n\t" {
    case "foo bar":
        v5 += 9223372036854775807
    case "\n\t":
        fmt.Println(v2, v5, v5)
    case "":
        v5 -= v5
    }
    fmt.Println(v5, v12, v5)
    if (-1 <= -1) {
        v213 := GenS2{F0: (v12 << ((751) & 63))}
        _ = v213
    } else {
        var v214 int8 = 18
        _ = v214
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v155
    _ = v156
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
