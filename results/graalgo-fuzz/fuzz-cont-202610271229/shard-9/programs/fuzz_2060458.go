package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64; F1 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint8; F1 uint8 }
type GenS10 struct { F0 int64; F1 uint64; F2 bool }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 string; F1 uint; F2 string }
type GenS13 struct { F0 int; F1 int16; F2 int16 }
type GenS14 struct { F0 int64; F1 int8 }
type GenS15 struct { F0 int16; F1 int64; F2 int }
type GenS16 struct { F0 int32; F1 int16; F2 int }
type GenS17 struct { F0 uint16; F1 uint16; F2 string }
type GenS18 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 uint8; F1 string }
type GenS21 struct { F0 bool; F1 uint }

func main() {
    var v1 int64 = -1
    v2 := []int{}
    var v3 int64 = v1
    v4 := make(map[string]string)
    var v5 uint = 4294967295
    var v6 string = ""
    v3 = 42
    v7 := GenS1{F0: 2147483648, F1: int16(v5)}
    if (!false) {
        var v8 uint16 = 2
        switch v6 {
        case "":
            if (true && true) {
                switch 9223372036854775807 {
                case 32768:
                    var v9 int64 = v1
                    _ = v9
                case -9223372036854775808:
                    for fi10 := 0; fi10 < 5; fi10++ {
                        for fi11 := 0; fi11 < 8; fi11++ {
                            var v12 int = (fi10 >> ((v5) & 63))
                            _ = fi11
                            _ = v12
                        }
                        _ = fi10
                    }
                default:
                    switch (v6 + "a") {
                    case "foo bar":
                        v13 := GenS2{F0: (v5 & 255), F1: -18}
                        _ = v13
                    case "\n\t":
                        var v14 int64 = (-(v3))
                        _ = v14
                    case "a":
                        switch ((v6 + "hello") + (((v6 + (v6 + "a")) + v6) + v6)) {
                        case "a":
                            if (1.5 != float64((0.0 - (float32(((0.0 / float64(1.5)) + -2.25)) + 3.14159)))) {
                                var v15 string = v6
                                _ = v15
                            }
                        default:
                            if len(v2) > 0 {
                                v2[0] = 713
                            }
                        }
                    default:
                        switch v6 {
                        case "\n\t":
                            switch ("a" + ((v6 + ((v6 + "foo bar") + ("\n\t" + v6))) + (((v6 + (v6 + "hello")) + v6) + "\n\t"))) {
                            case "":
                                if len(v2) > 0 {
                                    v2[0] = int(v1)
                                }
                            }
                        case "a":
                            v16 := GenS3{F0: 1}
                            _ = v16
                        case "hello":
                            v17 := v4["\n\t"]
                            _ = v17
                        }
                    }
                }
                var v18 float64 = (float64(1e10) - 1e10)
                if ((4294967295 <= 4294967295) && (true && true)) {
                    v19 := GenS2{F0: (v5 | 4294967295), F1: 100}
                    _ = v19
                } else {
                    var v20 string = (v6 + v6)
                    _ = v20
                }
                _ = v18
            } else {
                if false {
                    v4[((v6 + "hello") + "foo bar")] = ((v6 + v6) + v6)
                    v21 := GenS1{F0: 42, F1: -709}
                    var v22 uint8 = 0
                    _ = v21
                    _ = v22
                } else {
                    for fi23 := 0; fi23 < 13; fi23++ {
                        if false {
                            delete(v4, ("hello" + v6))
                            v24 := GenS1{F0: 2147483648, F1: 1}
                            if ((!(!true)) || false) {
                                if (-106 != 100) {
                                    _ = len(v4)
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v3 -= v3
                            }
                            _ = v24
                        } else {
                            var v25 uint = (^(v5))
                            if (!(!((-32768 != 1) || (true || (!(439 > 42)))))) {
                                switch (fi23 << (((^(v25))) & 63)) {
                                case 297:
                                    if true {
                                        var v26 int8 = 100
                                        var v27 float32 = (3.14159 + (-2.25 * float32((0.0 + 3.14159))))
                                        fmt.Println(v27)
                                        _ = v26
                                        _ = v27
                                    }
                                case 989:
                                    if true {
                                        for fi28 := 0; fi28 < 12; fi28++ {
                                            if false {
                                                switch ("\n\t" + "foo bar") {
                                                case "hello":
                                                    v1 -= v1
                                                }
                                            } else {
                                                if (true && true) {
                                                    var v29 uint64 = 379
                                                    v30 := GenS2{F0: v5, F1: 42}
                                                    _ = v29
                                                    _ = v30
                                                } else {
                                                    v5 += v25
                                                    var v31 int32 = int32(v3)
                                                    _ = v31
                                                }
                                            }
                                            _ = fi28
                                        }
                                    } else {
                                        for fi32 := 0; fi32 < 18; fi32++ {
                                            fmt.Println(v1, v1, v6)
                                            var v33 string = v6
                                            _ = fi32
                                            _ = v33
                                        }
                                        for fi34 := 0; fi34 < 12; fi34++ {
                                            delete(v4, v6)
                                            if false {
                                                var v35 int32 = -1
                                                _ = len(v4)
                                                var v36 uint = (v5 ^ (v5 & (v25 | uint(v8))))
                                                _ = v35
                                                _ = v36
                                            } else {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v37 uint32 = 2147483647
                                                _ = v37
                                            }
                                            if (-128 <= 62) {
                                                if true {
                                                    var v38 uint16 = 32767
                                                    v4[(v6 + ("foo bar" + (v6 + v6)))] = v6
                                                    v39 := GenS3{F0: 42}
                                                    if false {
                                                        if false {
                                                            var v40 uint64 = 965
                                                            v41 := GenS4{F0: fi23, F1: 126}
                                                            _ = v40
                                                            _ = v41
                                                        } else {
                                                            var v42 uint8 = 128
                                                            v5 = (v5 >> ((v25) & 63))
                                                            v43 := GenS1{F0: 162, F1: -1}
                                                            v44 := GenS3{F0: 32767}
                                                            v42 -= v42
                                                            var v45 float64 = 3.14159
                                                            _ = v42
                                                            _ = v43
                                                            _ = v44
                                                            _ = v45
                                                        }
                                                        if (uint32(v38) != 65535) {
                                                            if false {
                                                                var v46 uint = v25
                                                                _ = v46
                                                            } else {
                                                                if len(v2) > 0 {
                                                                    v2[0] = (fi34 & (fi34 << ((255) & 63)))
                                                                }
                                                                fmt.Println(fi34, v6, v3)
                                                                v38 = (v38 % (((v38 / (((v38 | 32767)) | 1))) | 1))
                                                            }
                                                            fmt.Println(v1)
                                                            fmt.Println(fi34)
                                                            switch "\n\t" {
                                                            case "hello":
                                                                v47 := GenS1{F0: 42, F1: 307}
                                                                _ = v47
                                                            case "foo bar":
                                                                if (true && true) {
                                                                    v25 = (v25 / (((v25 >> (((^(v5))) & 63))) | 1))
                                                                    var v48 int64 = int64(v8)
                                                                    v48 = (v48 | (v1 & -9223372036854775808))
                                                                    if false {
                                                                        _ = len(v2)
                                                                        var v49 uint16 = v8
                                                                        if true {
                                                                            _ = len(v2)
                                                                        } else {
                                                                            var v50 uint = (^(v25))
                                                                            var v51 int64 = v1
                                                                            var v52 int = 753
                                                                            if ((0.1 + 0.1) == 0.0) {
                                                                                v5 -= (v25 | (v25 & 65535))
                                                                                v53 := v4[v6]
                                                                                _ = v53
                                                                            }
                                                                            _ = v50
                                                                            _ = v51
                                                                            _ = v52
                                                                        }
                                                                        _ = v49
                                                                    }
                                                                    fmt.Println(v3)
                                                                    v1 = v3
                                                                    fmt.Println(v5, v8, v38)
                                                                    fmt.Println(v6)
                                                                    switch fi23 {
                                                                    case 9223372036854775807:
                                                                        v54 := GenS3{F0: -304}
                                                                        _ = v54
                                                                    case 1:
                                                                        switch v6 {
                                                                        case "\n\t":
                                                                            var v55 uint = v25
                                                                            _ = v55
                                                                        case "":
                                                                            if true {
                                                                                if ("a" >= v6) {
                                                                                    fmt.Println(v38, fi34, v38)
                                                                                    if false {
                                                                                        v56 := GenS2{F0: (v25 % ((v5) | 1)), F1: 0}
                                                                                        _ = v56
                                                                                    }
                                                                                } else {
                                                                                    switch v6 {
                                                                                    case "a":
                                                                                        if true {
                                                                                            switch (fi34 << (((^(v5))) & 63)) {
                                                                                            case 127:
                                                                                                var v57 uint = (v25 + 199)
                                                                                                _ = v57
                                                                                            case 42:
                                                                                                switch 478 {
                                                                                                case 0:
                                                                                                    switch ("foo bar" + v6) {
                                                                                                    case "":
                                                                                                        if false {
                                                                                                            switch "\n\t" {
                                                                                                            case "":
                                                                                                                v8 = v8
                                                                                                            default:
                                                                                                                v5 -= (v25 | (v5 ^ (v25 - (^(v25)))))
                                                                                                            }
                                                                                                            var v58 float64 = (1e10 / (-2.25 * 1e10))
                                                                                                            _ = v58
                                                                                                        }
                                                                                                    case "a":
                                                                                                        var v59 float32 = 0.1
                                                                                                        _ = v59
                                                                                                    case "\n\t":
                                                                                                        var v60 int32 = -1
                                                                                                        _ = v60
                                                                                                    }
                                                                                                }
                                                                                            case 32768:
                                                                                                v25 = (v25 + v25)
                                                                                            default:
                                                                                                var v61 bool = (false && ((v25 + v25) <= v5))
                                                                                                _ = v61
                                                                                            }
                                                                                            delete(v4, (v6 + v6))
                                                                                        } else {
                                                                                            switch (v6 + "foo bar") {
                                                                                            case "a":
                                                                                                fmt.Println(v5, v48, v5)
                                                                                            case "foo bar":
                                                                                                switch v6 {
                                                                                                case "\n\t":
                                                                                                    v62 := GenS5{F0: (v3 - v48)}
                                                                                                    _ = v62
                                                                                                case "foo bar":
                                                                                                    v8 -= (^(v38))
                                                                                                case "a":
                                                                                                    v63 := v4[("foo bar" + v6)]
                                                                                                    _ = v63
                                                                                                default:
                                                                                                    if true {
                                                                                                        var v64 uint64 = 524
                                                                                                        var v65 string = (v6 + "")
                                                                                                        if (!(86 <= 1)) {
                                                                                                            v66 := v4["a"]
                                                                                                            v48 -= -466
                                                                                                            _ = v66
                                                                                                        } else {
                                                                                                            if false {
                                                                                                                v67 := GenS4{F0: (-(fi23)), F1: -1}
                                                                                                                v38 += (^(v38))
                                                                                                                v68 := GenS6{F0: 42, F1: v5}
                                                                                                                _ = v67
                                                                                                                _ = v68
                                                                                                            }
                                                                                                            if ((v25 | v25) <= v5) {
                                                                                                                fmt.Println(v6, v48)
                                                                                                                fmt.Println(v38, v25, v1)
                                                                                                            }
                                                                                                            v69 := GenS4{F0: int(v3), F1: 0}
                                                                                                            if (true || false) {
                                                                                                                var v70 bool = true
                                                                                                                switch fi34 {
                                                                                                                case 32767:
                                                                                                                    var v71 int32 = -312
                                                                                                                    _ = v71
                                                                                                                case 128:
                                                                                                                    _ = len(v4)
                                                                                                                case 806:
                                                                                                                    var v72 uint8 = 127
                                                                                                                    _ = v72
                                                                                                                }
                                                                                                                _ = v70
                                                                                                            } else {
                                                                                                                if ((v25 | uint(v38)) == v5) {
                                                                                                                    v2 = append(v2, fi34)
                                                                                                                } else {
                                                                                                                    v48 = v48
                                                                                                                    v8 = (v8 % ((v38) | 1))
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v69
                                                                                                        }
                                                                                                        _ = v64
                                                                                                        _ = v65
                                                                                                    } else {
                                                                                                        v73 := GenS7{F0: 32767}
                                                                                                        var v74 uint16 = (v38 ^ (^(v38)))
                                                                                                        delete(v4, v6)
                                                                                                        _ = v73
                                                                                                        _ = v74
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                v38 -= (v38 << ((553) & 63))
                                                                                            }
                                                                                            fmt.Println(v6, v6)
                                                                                        }
                                                                                    default:
                                                                                        if (v38 >= 685) {
                                                                                            fmt.Println(v6)
                                                                                        } else {
                                                                                            var v75 float32 = 3.14159
                                                                                            var v76 uint32 = 1
                                                                                            fmt.Println(v1, v76)
                                                                                            if true {
                                                                                                if (false && true) {
                                                                                                    var v77 int64 = 1
                                                                                                    v78 := GenS8{F0: 765}
                                                                                                    _ = v77
                                                                                                    _ = v78
                                                                                                } else {
                                                                                                    v79 := GenS9{F0: 127, F1: 128}
                                                                                                    var v80 int32 = 1
                                                                                                    var v81 uint = (v25 - (v5 + 4294967295))
                                                                                                    var v82 int8 = -52
                                                                                                    _ = v79
                                                                                                    _ = v80
                                                                                                    _ = v81
                                                                                                    _ = v82
                                                                                                }
                                                                                                if true {
                                                                                                    switch "foo bar" {
                                                                                                    case "a":
                                                                                                        if (true && false) {
                                                                                                            v83 := v4[(v6 + (v6 + "foo bar"))]
                                                                                                            v84 := GenS5{F0: -1}
                                                                                                            _ = v83
                                                                                                            _ = v84
                                                                                                        } else {
                                                                                                            v25 = (^(v25))
                                                                                                            delete(v4, v6)
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                var v85 int64 = v1
                                                                                                _ = v85
                                                                                            } else {
                                                                                                var v86 uint16 = v8
                                                                                                if true {
                                                                                                    if (true && false) {
                                                                                                        v5 += v25
                                                                                                        if false {
                                                                                                            if false {
                                                                                                                var v87 uint8 = uint8(v86)
                                                                                                                _ = v87
                                                                                                            } else {
                                                                                                                var v88 int16 = -1
                                                                                                                fmt.Println(v86)
                                                                                                                v89 := GenS10{F0: (^(v1)), F1: uint64(v8), F2: (587 <= -1)}
                                                                                                                _ = v88
                                                                                                                _ = v89
                                                                                                            }
                                                                                                        } else {
                                                                                                            if true {
                                                                                                                _ = len(v4)
                                                                                                                var v90 int = (fi23 * 127)
                                                                                                                v48 = (-(v3))
                                                                                                                _ = v90
                                                                                                            } else {
                                                                                                                fmt.Println(v76, v75)
                                                                                                            }
                                                                                                            switch v6 {
                                                                                                            case "a":
                                                                                                                switch "" {
                                                                                                                case "\n\t":
                                                                                                                    if ((true || (!(v38 <= (v8 ^ (v8 - v38))))) || true) {
                                                                                                                        v3 -= v1
                                                                                                                        var v91 float64 = (3.14159 + (0.0 * (-2.25 / 1.0)))
                                                                                                                        _ = v91
                                                                                                                    }
                                                                                                                    v3 = v48
                                                                                                                }
                                                                                                                switch (fi23 + 138) {
                                                                                                                case -2147483648:
                                                                                                                    delete(v4, "a")
                                                                                                                case 9223372036854775807:
                                                                                                                    var v92 int8 = 85
                                                                                                                    _ = v92
                                                                                                                default:
                                                                                                                    v75 = (v75 / v75)
                                                                                                                }
                                                                                                                var v93 int = fi23
                                                                                                                if (false || (32 <= -128)) {
                                                                                                                    fmt.Println(v8, v93, v3)
                                                                                                                }
                                                                                                                _ = v93
                                                                                                            case "":
                                                                                                                fmt.Println(v1)
                                                                                                            default:
                                                                                                                if true {
                                                                                                                    var v94 uint = 4294967295
                                                                                                                    _ = v94
                                                                                                                } else {
                                                                                                                    v95 := GenS7{F0: v8}
                                                                                                                    _ = v95
                                                                                                                }
                                                                                                                var v96 uint16 = 0
                                                                                                                if false {
                                                                                                                    switch "a" {
                                                                                                                    case "\n\t":
                                                                                                                        switch (fi23 >> ((v25) & 63)) {
                                                                                                                        case 9223372036854775807:
                                                                                                                            var v97 uint8 = 128
                                                                                                                            _ = v97
                                                                                                                        }
                                                                                                                    case "a":
                                                                                                                        v38 -= v38
                                                                                                                    }
                                                                                                                    fmt.Println(v3)
                                                                                                                    delete(v4, "foo bar")
                                                                                                                    if false {
                                                                                                                        fmt.Println(v8)
                                                                                                                    }
                                                                                                                    fmt.Println(v6, v38, v3)
                                                                                                                    v75 += v75
                                                                                                                } else {
                                                                                                                    v8 -= 42
                                                                                                                    v8 = v38
                                                                                                                    switch v6 {
                                                                                                                    case "\n\t":
                                                                                                                        if (0.0 != 1.5) {
                                                                                                                            switch (-(fi23)) {
                                                                                                                            case -910:
                                                                                                                                if (true || false) {
                                                                                                                                    v4[(v6 + "foo bar")] = ((((("hello" + v6) + "") + "hello") + (v6 + (v6 + v6))) + v6)
                                                                                                                                } else {
                                                                                                                                    var v98 bool = (v75 <= -2.25)
                                                                                                                                    _ = v98
                                                                                                                                }
                                                                                                                            case 4294967295:
                                                                                                                                v99 := GenS5{F0: -1}
                                                                                                                                _ = v99
                                                                                                                            case 1:
                                                                                                                                var v100 int64 = -864
                                                                                                                                _ = v100
                                                                                                                            }
                                                                                                                            var v101 int16 = -980
                                                                                                                            _ = v101
                                                                                                                        }
                                                                                                                    case "a":
                                                                                                                        switch (((("a" + v6) + "foo bar") + v6) + "") {
                                                                                                                        case "\n\t":
                                                                                                                            if true {
                                                                                                                                v102 := GenS11{F0: (0 == 63)}
                                                                                                                                v103 := GenS12{F0: v6, F1: v5, F2: v6}
                                                                                                                                fmt.Println(v6)
                                                                                                                                _ = v102
                                                                                                                                _ = v103
                                                                                                                            } else {
                                                                                                                                switch "" {
                                                                                                                                case "hello":
                                                                                                                                    switch "" {
                                                                                                                                    case "hello":
                                                                                                                                        v86 = v38
                                                                                                                                    case "foo bar":
                                                                                                                                        switch (("hello" + v6) + v6) {
                                                                                                                                        case "":
                                                                                                                                            v25 = (v25 ^ v25)
                                                                                                                                        default:
                                                                                                                                            v76 += (v76 & (v76 * (v76 >> (((^(v25))) & 63))))
                                                                                                                                            v6 = v6
                                                                                                                                            fmt.Println(v3, v1)
                                                                                                                                        }
                                                                                                                                    case "\n\t":
                                                                                                                                        fmt.Println(v6)
                                                                                                                                    default:
                                                                                                                                        switch ("" + (("foo bar" + v6) + (v6 + v6))) {
                                                                                                                                        case "":
                                                                                                                                            fmt.Println(v75)
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                case "a":
                                                                                                                                    var v104 int8 = 100
                                                                                                                                    _ = v104
                                                                                                                                case "\n\t":
                                                                                                                                    var v105 float32 = v75
                                                                                                                                    _ = v105
                                                                                                                                default:
                                                                                                                                    v106 := v4[v6]
                                                                                                                                    _ = v106
                                                                                                                                }
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            v107 := GenS13{F0: int(v3), F1: 0, F2: -1}
                                                                                                                            _ = v107
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v96
                                                                                                            }
                                                                                                            switch -1 {
                                                                                                            case -2147483648:
                                                                                                                if (!false) {
                                                                                                                    fmt.Println(v6, v48, v86)
                                                                                                                }
                                                                                                            }
                                                                                                            if false {
                                                                                                                v108 := GenS2{F0: v25, F1: int8(v25)}
                                                                                                                _ = v108
                                                                                                            }
                                                                                                        }
                                                                                                        if len(v2) >= 2 {
                                                                                                            v2 = v2[1:len(v2)]
                                                                                                        }
                                                                                                        v2 = append(v2, 132)
                                                                                                    }
                                                                                                }
                                                                                                if true {
                                                                                                    if (((^(v25)) == (v25 - v5)) || (int8(fi34) < 90)) {
                                                                                                        v109 := GenS3{F0: 42}
                                                                                                        _ = v109
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            v48 += v48
                                                                                                            if true {
                                                                                                                v6 = v6
                                                                                                                v48 = v1
                                                                                                            }
                                                                                                        }
                                                                                                        v5 = (v5 & v5)
                                                                                                        var v110 int32 = int32(v38)
                                                                                                        if len(v2) >= 2 {
                                                                                                            v2 = v2[1:len(v2)]
                                                                                                        }
                                                                                                        switch (fi34 + 0) {
                                                                                                        case -315:
                                                                                                            var v111 uint8 = 2
                                                                                                            _ = v111
                                                                                                        default:
                                                                                                            v2 = append(v2, fi34)
                                                                                                        }
                                                                                                        _ = v110
                                                                                                    }
                                                                                                }
                                                                                                fmt.Println(v75, v5)
                                                                                                _ = v86
                                                                                            }
                                                                                            _ = v75
                                                                                            _ = v76
                                                                                        }
                                                                                        v5 -= 648
                                                                                        var v112 uint16 = (v38 + 2)
                                                                                        _ = v112
                                                                                    }
                                                                                    fmt.Println(v38, v48)
                                                                                }
                                                                            }
                                                                        case "hello":
                                                                            var v113 uint16 = 141
                                                                            _ = v113
                                                                        }
                                                                    case -2147483648:
                                                                        var v114 int32 = 42
                                                                        _ = v114
                                                                    default:
                                                                        v115 := GenS6{F0: (v1 % ((v48) | 1)), F1: v5}
                                                                        _ = v115
                                                                    }
                                                                    v3 -= int64(v25)
                                                                    _ = v48
                                                                }
                                                            default:
                                                                var v116 uint16 = v8
                                                                _ = v116
                                                            }
                                                            var v117 bool = (1.5 < float32((3.14159 - 1.5)))
                                                            _ = v117
                                                        }
                                                        var v118 float64 = (1.5 / 0.1)
                                                        _ = v118
                                                    } else {
                                                        var v119 int16 = 321
                                                        v120 := GenS6{F0: (v3 * v1), F1: 18446744073709551615}
                                                        var v121 string = ""
                                                        _ = v119
                                                        _ = v120
                                                        _ = v121
                                                    }
                                                    fmt.Println(fi23, v38, v1)
                                                    switch (fi23 - (fi34 << ((v5) & 63))) {
                                                    case 463:
                                                        v38 = (v38 | 2)
                                                    case 502:
                                                        v122 := GenS8{F0: 291}
                                                        _ = v122
                                                    case 4294967295:
                                                        var v123 string = v6
                                                        _ = v123
                                                    default:
                                                        v124 := GenS14{F0: v1, F1: 1}
                                                        _ = v124
                                                    }
                                                    if true {
                                                        var v125 string = ((v6 + v6) + v6)
                                                        var v126 float32 = ((float32(3.14159) + float32((1.5 + float64((0.1 / 1e10))))) + (float32((0.1 + 1.5)) * -2.25))
                                                        _ = v125
                                                        _ = v126
                                                    } else {
                                                        var v127 string = ((((v6 + ((v6 + (v6 + (v6 + ""))) + "hello")) + "\n\t") + v6) + ("a" + v6))
                                                        var v128 bool = (42 > 300)
                                                        if (!((!(true || (v8 == (v8 | v8)))) && false)) {
                                                            fmt.Println(v127)
                                                            v129 := GenS6{F0: v1, F1: (v5 & (v5 + v25))}
                                                            v130 := v4[v127]
                                                            _ = v129
                                                            _ = v130
                                                        } else {
                                                            v5 -= 597
                                                            switch "" {
                                                            case "":
                                                                var v131 uint64 = 18446744073709551615
                                                                _ = v131
                                                            default:
                                                                var v132 int64 = v1
                                                                _ = v132
                                                            }
                                                            v133 := GenS15{F0: 631, F1: (-(v1)), F2: fi34}
                                                            if v128 {
                                                                var v134 uint64 = 196
                                                                _ = v134
                                                            }
                                                            v6 = ((v127 + (v6 + v6)) + v6)
                                                            _ = v133
                                                        }
                                                        if true {
                                                            v135 := GenS4{F0: (fi23 - fi23), F1: 1}
                                                            var v136 uint32 = 2147483648
                                                            var v137 string = ("a" + "hello")
                                                            _ = v135
                                                            _ = v136
                                                            _ = v137
                                                        }
                                                        v38 = v8
                                                        fmt.Println(v8, v1)
                                                        var v138 int32 = 1
                                                        _ = v127
                                                        _ = v128
                                                        _ = v138
                                                    }
                                                    _ = v38
                                                    _ = v39
                                                } else {
                                                    if (18446744073709551615 != 2) {
                                                        var v139 int = fi34
                                                        switch v6 {
                                                        case "a":
                                                            if ((1 != 42) && true) {
                                                                _ = len(v2)
                                                                v8 = 912
                                                                _ = len(v2)
                                                            }
                                                        case "\n\t":
                                                            switch "hello" {
                                                            case "foo bar":
                                                                var v140 uint = 4294967295
                                                                _ = v140
                                                            case "\n\t":
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            }
                                                        default:
                                                            v6 = v6
                                                        }
                                                        var v141 uint8 = 76
                                                        _ = v139
                                                        _ = v141
                                                    } else {
                                                        v142 := GenS10{F0: 1, F1: 141, F2: false}
                                                        v8 = v8
                                                        v2 = append(v2, (fi34 ^ (fi23 >> (((v5 ^ v25)) & 63))))
                                                        _ = v142
                                                    }
                                                    if true {
                                                        switch v6 {
                                                        case "hello":
                                                            fmt.Println(v6, fi23)
                                                        }
                                                        if (!false) {
                                                            var v143 float64 = 1.5
                                                            var v144 uint64 = 554
                                                            if false {
                                                                switch fi34 {
                                                                case 32767:
                                                                    fmt.Println(v143, v3, v143)
                                                                case -2147483648:
                                                                    var v145 bool = (!(1 != -1))
                                                                    _ = v145
                                                                case 32768:
                                                                    fmt.Println(fi23, v25)
                                                                    fmt.Println(v5, fi34)
                                                                default:
                                                                    var v146 float64 = float64(0.1)
                                                                    _ = v146
                                                                }
                                                                if ((((v144 * 603) != (^(v144))) || (!false)) && false) {
                                                                    if (65535 > 1) {
                                                                        v147 := GenS10{F0: v1, F1: (v144 / ((v144) | 1)), F2: (v143 < 0.1)}
                                                                        v6 = "foo bar"
                                                                        _ = v147
                                                                    }
                                                                    if true {
                                                                        var v148 int64 = (v1 * (v3 >> (((^(v5))) & 63)))
                                                                        _ = v148
                                                                    }
                                                                } else {
                                                                    v144 = (v144 % (((v144 + 124)) | 1))
                                                                    var v149 int64 = -1
                                                                    fmt.Println(v149, v143, v144)
                                                                    v6 = "hello"
                                                                    var v150 int8 = 121
                                                                    switch (v6 + ((v6 + v6) + v6)) {
                                                                    case "hello":
                                                                        var v151 uint16 = v8
                                                                        _ = v151
                                                                    default:
                                                                        var v152 int64 = (^(v1))
                                                                        _ = v152
                                                                    }
                                                                    _ = v149
                                                                    _ = v150
                                                                }
                                                                v153 := GenS12{F0: "", F1: (v25 | 302), F2: v6}
                                                                var v154 int8 = 52
                                                                v143 = 0.0
                                                                if true {
                                                                    var v155 float64 = v143
                                                                    if true {
                                                                        fmt.Println(v6, fi23)
                                                                        if false {
                                                                            v155 = v143
                                                                        } else {
                                                                            if (float64(float32((0.0 * v143))) < (float64(1.5) + 0.0)) {
                                                                                fmt.Println(fi34, v5)
                                                                            } else {
                                                                                if false {
                                                                                    if false {
                                                                                        switch (fi23 ^ fi34) {
                                                                                        case 32767:
                                                                                            switch fi34 {
                                                                                            case 127:
                                                                                                v155 += float64(1.5)
                                                                                            case 650:
                                                                                                v156 := GenS16{F0: -2147483648, F1: int16(fi34), F2: (fi23 + fi23)}
                                                                                                _ = v156
                                                                                            }
                                                                                        }
                                                                                        v5 -= 608
                                                                                        if (!false) {
                                                                                            if (!true) {
                                                                                                if true {
                                                                                                    var v157 int8 = v154
                                                                                                    _ = v157
                                                                                                } else {
                                                                                                    switch v6 {
                                                                                                    case "\n\t":
                                                                                                        switch fi34 {
                                                                                                        case 42:
                                                                                                            switch fi23 {
                                                                                                            case 127:
                                                                                                                v158 := GenS9{F0: 127, F1: 189}
                                                                                                                fmt.Println(v155, v155)
                                                                                                                _ = v158
                                                                                                            default:
                                                                                                                var v159 uint = uint(v144)
                                                                                                                _ = v159
                                                                                                            }
                                                                                                        case 4294967295:
                                                                                                            if (274 > 1) {
                                                                                                                if len(v2) > 0 {
                                                                                                                    v2[0] = 1
                                                                                                                }
                                                                                                                var v160 int32 = -2147483648
                                                                                                                fmt.Println(v8, fi34)
                                                                                                                _ = v160
                                                                                                            } else {
                                                                                                                v154 -= -1
                                                                                                                v3 = -362
                                                                                                                switch "a" {
                                                                                                                case "\n\t":
                                                                                                                    var v161 int64 = v1
                                                                                                                    _ = v161
                                                                                                                case "":
                                                                                                                    v4[v6] = v6
                                                                                                                case "hello":
                                                                                                                    v6 = ("a" + "hello")
                                                                                                                }
                                                                                                            }
                                                                                                        default:
                                                                                                            v162 := GenS16{F0: -359, F1: 1, F2: fi23}
                                                                                                            _ = v162
                                                                                                        }
                                                                                                    case "foo bar":
                                                                                                        delete(v4, v6)
                                                                                                    case "a":
                                                                                                        var v163 int16 = int16(v144)
                                                                                                        _ = v163
                                                                                                    }
                                                                                                    fmt.Println(fi23, fi34, v8)
                                                                                                    v164 := v4[(((v6 + "foo bar") + (v6 + (v6 + (v6 + v6)))) + ((("" + ("\n\t" + "foo bar")) + v6) + ((v6 + (v6 + "hello")) + v6)))]
                                                                                                    _ = v164
                                                                                                }
                                                                                                var v165 bool = true
                                                                                                _ = v165
                                                                                            } else {
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        v166 := GenS17{F0: v8, F1: v8, F2: (("a" + v6) + ((v6 + v6) + ((v6 + v6) + v6)))}
                                                                        _ = v166
                                                                    } else {
                                                                        var v167 uint64 = 0
                                                                        var v168 int32 = 771
                                                                        _ = v167
                                                                        _ = v168
                                                                    }
                                                                    v6 = ("hello" + ((("hello" + v6) + "a") + "\n\t"))
                                                                    _ = v155
                                                                }
                                                                _ = v153
                                                                _ = v154
                                                            } else {
                                                                var v169 int = 4294967295
                                                                _ = v169
                                                            }
                                                            _ = v143
                                                            _ = v144
                                                        } else {
                                                            fmt.Println(v25, v3)
                                                            if true {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            } else {
                                                                fmt.Println(fi34)
                                                                if ((^(fi23)) == fi34) {
                                                                    var v170 bool = ((true && false) || (1e10 != (((1.5 + 0.1) + (3.14159 - float32((3.14159 + (3.14159 / 1.5))))) / 0.1)))
                                                                    _ = v170
                                                                }
                                                                var v171 uint16 = v8
                                                                _ = v171
                                                            }
                                                            v4[("a" + (v6 + v6))] = v6
                                                            var v172 uint16 = v8
                                                            fmt.Println(fi34, v25)
                                                            _ = v172
                                                        }
                                                        _ = len(v2)
                                                    }
                                                }
                                            } else {
                                                switch v6 {
                                                case "a":
                                                    v3 += (v3 * -1)
                                                case "hello":
                                                    v8 = (^(v8))
                                                }
                                                if false {
                                                    var v173 float32 = 1e10
                                                    v174 := GenS9{F0: 218, F1: 2}
                                                    fmt.Println(v6)
                                                    _ = v173
                                                    _ = v174
                                                } else {
                                                    var v175 uint16 = v8
                                                    delete(v4, v6)
                                                    if true {
                                                        var v176 int64 = (v3 / ((-9223372036854775808) | 1))
                                                        _ = v176
                                                    }
                                                    _ = v175
                                                }
                                            }
                                            _ = fi34
                                        }
                                    }
                                case 32767:
                                    v177 := v4["a"]
                                    _ = v177
                                }
                                if true {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v8 = v8
                                    if false {
                                        for fi178 := 0; fi178 < 7; fi178++ {
                                            v1 = (-(v3))
                                            v179 := GenS3{F0: 69}
                                            v180 := GenS18{F0: 4294967295, F1: 495, F2: 1}
                                            _ = fi178
                                            _ = v179
                                            _ = v180
                                        }
                                        v181 := GenS19{F0: (v25 & (v5 - (v25 * v5)))}
                                        if false {
                                            var v182 uint16 = (v8 | (v8 & v8))
                                            if true {
                                                for fi183 := 0; fi183 < 5; fi183++ {
                                                    if false {
                                                        v184 := v4[v6]
                                                        var v185 float32 = (1e10 * 1.5)
                                                        if (42 < -104) {
                                                            v8 = 65535
                                                        } else {
                                                            v6 = (v6 + v6)
                                                            fmt.Println(v8, v182)
                                                            v184 = (v6 + "foo bar")
                                                        }
                                                        _ = v184
                                                        _ = v185
                                                    }
                                                    switch v6 {
                                                    case "hello":
                                                        v2 = append(v2, (fi183 - 2147483647))
                                                    case "":
                                                        if true {
                                                            delete(v4, (("foo bar" + v6) + v6))
                                                            delete(v4, ((("foo bar" + v6) + ((v6 + v6) + v6)) + (v6 + v6)))
                                                            v186 := v4[(("a" + v6) + ((v6 + "foo bar") + v6))]
                                                            _ = v186
                                                        } else {
                                                            var v187 uint32 = 905
                                                            _ = v187
                                                        }
                                                        if (false && true) {
                                                            if len(v2) > 0 {
                                                                v2[0] = (fi23 - (-(fi183)))
                                                            }
                                                        }
                                                    case "a":
                                                        v2 = append(v2, (fi23 >> ((v25) & 63)))
                                                    }
                                                    v188 := GenS10{F0: -1, F1: 4294967295, F2: (127 == -1)}
                                                    v25 = v25
                                                    _ = fi183
                                                    _ = v188
                                                }
                                                v189 := v4[((v6 + ("" + ("\n\t" + ("" + v6)))) + v6)]
                                                if true {
                                                    var v190 float32 = 0.1
                                                    var v191 uint16 = v8
                                                    if false {
                                                        v182 -= 218
                                                        if len(v2) > 0 {
                                                            v2[0] = (fi23 - (fi23 - 1))
                                                        }
                                                    } else {
                                                        fmt.Println(v1, v6)
                                                        v192 := GenS11{F0: (18446744073709551615 != 0)}
                                                        _ = v192
                                                    }
                                                    _ = v190
                                                    _ = v191
                                                } else {
                                                    if (fi23 >= fi23) {
                                                        v189 = (v6 + ((v189 + v189) + v6))
                                                        v25 = (^(v25))
                                                    } else {
                                                        var v193 int8 = -97
                                                        _ = v193
                                                    }
                                                }
                                                _ = v189
                                            } else {
                                                if ((!true) && true) {
                                                    switch v6 {
                                                    case "\n\t":
                                                        v194 := GenS2{F0: (v25 ^ (v5 >> (((v25 + 254)) & 63))), F1: -128}
                                                        _ = v194
                                                    }
                                                    v195 := GenS16{F0: -454, F1: 781, F2: 819}
                                                    _ = v195
                                                } else {
                                                    v8 = v182
                                                    v5 += (v5 + (v25 >> (((v5 & (v25 / ((42) | 1)))) & 63)))
                                                }
                                                var v196 uint8 = uint8(v25)
                                                _ = v196
                                            }
                                            _ = v182
                                        } else {
                                            var v197 uint64 = 18446744073709551615
                                            v5 = v5
                                            _ = v197
                                        }
                                        _ = v181
                                    }
                                } else {
                                    v198 := GenS20{F0: 119, F1: "\n\t"}
                                    var v199 uint32 = 0
                                    _ = v198
                                    _ = v199
                                }
                                if len(v2) > 0 {
                                    v2[0] = (fi23 + (fi23 & fi23))
                                }
                            }
                            v200 := v4["\n\t"]
                            _ = v200
                            _ = v25
                        }
                        v6 = v6
                        _ = fi23
                    }
                    v201 := GenS21{F0: false, F1: v5}
                    if (v6 <= ((v6 + v6) + ("" + ((v6 + (("foo bar" + (((v6 + v6) + v6) + ("a" + ((v6 + v6) + v6)))) + v6)) + ((v6 + v6) + "\n\t"))))) {
                        var v202 string = "\n\t"
                        if (true || false) {
                            fmt.Println(v202, v1)
                            v203 := v4[v202]
                            _ = v203
                        }
                        _ = v202
                    }
                    _ = v201
                }
                var v204 int64 = (v1 - (v3 * -9223372036854775808))
                if (float32(((0.1 * float64(0.0)) * float64((3.14159 + -2.25)))) == float32((0.0 * 0.1))) {
                    var v205 uint32 = 2147483647
                    v206 := GenS16{F0: 0, F1: -1, F2: 2147483647}
                    fmt.Println(v1, v3)
                    _ = v205
                    _ = v206
                }
                _ = v204
            }
        }
        _ = v8
    }
    switch -589 {
    case -507:
        for fi207 := 0; fi207 < 16; fi207++ {
            var v208 int8 = -12
            _ = fi207
            _ = v208
        }
    case -607:
        v209 := GenS10{F0: 9223372036854775807, F1: 67, F2: false}
        _ = v209
    case 2147483647:
        for fi210 := 0; fi210 < 4; fi210++ {
            v1 -= -1
            var v211 float32 = -2.25
            _ = fi210
            _ = v211
        }
    }
    for fi212 := 0; fi212 < 14; fi212++ {
        if true {
            for fi213 := 0; fi213 < 18; fi213++ {
                v3 -= v1
                var v214 int64 = v3
                _ = fi213
                _ = v214
            }
            switch (-(fi212)) {
            case -9223372036854775808:
                switch (fi212 << (((v5 - (v5 - 432))) & 63)) {
                case -2147483648:
                    var v215 bool = ((v1 > v3) && true)
                    _ = v215
                case -610:
                    for fi216 := 0; fi216 < 4; fi216++ {
                        if true {
                            var v217 int32 = 1
                            v1 = v1
                            var v218 float32 = float32(-2.25)
                            v217 -= int32(v3)
                            _ = v217
                            _ = v218
                        }
                        _ = fi216
                    }
                }
            case 32767:
                v6 = "a"
            default:
                if true {
                    var v219 float64 = 1.5
                    for fi220 := 0; fi220 < 15; fi220++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v221 int = (fi220 & fi212)
                        var v222 uint32 = 474
                        _ = fi220
                        _ = v221
                        _ = v222
                    }
                    switch fi212 {
                    case 128:
                        _ = len(v4)
                    case 32768:
                        v219 = ((0.0 + 0.0) + (0.1 + v219))
                    }
                    _ = v219
                }
            }
        } else {
            var v223 uint64 = 42
            for fi224 := 0; fi224 < 1; fi224++ {
                if (1.5 < 3.14159) {
                    if false {
                        fmt.Println(v5)
                        v225 := v4[(((v6 + (("\n\t" + "") + ("hello" + v6))) + v6) + v6)]
                        var v226 uint64 = (^(v223))
                        _ = len(v4)
                        _ = v225
                        _ = v226
                    }
                } else {
                    switch v6 {
                    case "foo bar":
                        if len(v2) > 0 {
                            v2[0] = (fi212 * fi212)
                        }
                    }
                    var v227 uint32 = 65535
                    v228 := GenS15{F0: 1, F1: (v3 << (((v5 % ((v5) | 1))) & 63)), F2: 683}
                    v229 := GenS1{F0: (v227 - v227), F1: -905}
                    _ = v227
                    _ = v228
                    _ = v229
                }
                fmt.Println(v6)
                _ = fi224
            }
            _ = v223
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi212
    }
    v230 := GenS1{F0: 4294967295, F1: 198}
    v5 = (v5 - 18446744073709551615)
    fmt.Println(v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v230
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
