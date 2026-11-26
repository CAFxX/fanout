package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 string; F1 string; F2 int16 }
type GenS6 struct { F0 int; F1 int32; F2 bool }
type GenS7 struct { F0 int8; F1 int16; F2 uint64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int16; F1 uint64; F2 int16 }
type GenS11 struct { F0 string; F1 int16; F2 uint16 }
type GenS12 struct { F0 int16; F1 int; F2 uint16 }
type GenS13 struct { F0 int16; F1 uint32; F2 uint8 }
func genFn0() uint {
    var v6 uint8 = 96
    for fi7 := 0; fi7 < 15; fi7++ {
        if true {
            var v8 int32 = 42
            fmt.Println(v8, v6)
            _ = v8
        }
        for fi9 := 0; fi9 < 13; fi9++ {
            v6 = (v6 / ((v6) | 1))
            var v10 int16 = 42
            _ = fi9
            _ = v10
        }
        _ = fi7
    }
    if true {
        if false {
            v11 := GenS1{F0: 42, F1: 0}
            _ = v11
        } else {
            v6 = (v6 & v6)
        }
    } else {
        for fi12 := 0; fi12 < 7; fi12++ {
            v13 := []int{fi12}
            switch ("" + "\n\t") {
            case "\n\t":
                v14 := GenS1{F0: 2, F1: 2147483647}
                _ = v14
            case "foo bar":
                for fi15 := 0; fi15 < 7; fi15++ {
                    v6 = v6
                    if false {
                        if (fi15 <= -2147483648) {
                            if len(v13) > 0 {
                                v13[0] = (fi12 % ((fi12) | 1))
                            }
                            _ = len(v13)
                            v6 -= (^(v6))
                            v13 = append(v13, 1)
                        } else {
                            var v16 int32 = -1
                            if len(v13) > 0 {
                                v13[0] = 32767
                            }
                            v6 -= v6
                            v16 = v16
                            if (-96 < 1) {
                                if len(v13) > 0 {
                                    v13[0] = fi15
                                }
                                var v17 int8 = 100
                                _ = v17
                            } else {
                                var v18 int = (fi15 % ((fi12) | 1))
                                if len(v13) > 0 {
                                    v13[0] = 730
                                }
                                _ = v18
                            }
                            var v19 int16 = -1
                            _ = v16
                            _ = v19
                        }
                        var v20 uint16 = 65535
                        var v21 float64 = (0.1 + 0.0)
                        _ = v20
                        _ = v21
                    } else {
                        v6 += (v6 / (((^(v6))) | 1))
                        var v22 uint16 = 255
                        fmt.Println(fi12, v22, fi15)
                        v6 = v6
                        switch fi12 {
                        case -9223372036854775808:
                            v6 = v6
                        case -1:
                            var v23 float32 = 1.5
                            _ = v23
                        case 128:
                            if (631 == uint(v6)) {
                                var v24 int8 = 58
                                _ = v24
                            } else {
                                var v25 int16 = -32768
                                var v26 int8 = 0
                                _ = v25
                                _ = v26
                            }
                        default:
                            fmt.Println(v22, v22)
                        }
                        _ = v22
                    }
                    _ = fi15
                }
            case "":
                var v27 int8 = -47
                _ = v27
            default:
                if false {
                    v6 += (v6 ^ (v6 | 172))
                    v6 = (^(v6))
                    for fi28 := 0; fi28 < 16; fi28++ {
                        var v29 float32 = 0.1
                        if true {
                            v6 -= (v6 << ((42) & 63))
                        } else {
                            v29 = (v29 + (v29 - v29))
                            v29 = 0.0
                            var v30 uint8 = v6
                            if (-2.25 != (1.5 + -2.25)) {
                                if true {
                                    if false {
                                        var v31 int = 1
                                        fmt.Println(v29)
                                        if true {
                                            if true {
                                                v13 = append(v13, fi12)
                                                var v32 int8 = 39
                                                v32 = (^(v32))
                                                _ = v32
                                            } else {
                                                if (1.5 != 0.1) {
                                                    fmt.Println(fi12, fi12)
                                                    if len(v13) >= 2 {
                                                        v13 = v13[1:len(v13)]
                                                    }
                                                    v31 = 2147483647
                                                }
                                                v31 = v31
                                            }
                                        } else {
                                            switch ((("foo bar" + (("a" + (("" + "\n\t") + (("a" + "\n\t") + ("\n\t" + ("\n\t" + "\n\t"))))) + "\n\t")) + "") + "hello") {
                                            case "\n\t":
                                                v6 -= v30
                                            case "a":
                                                v13 = append(v13, fi28)
                                            default:
                                                if len(v13) > 0 {
                                                    v13[0] = (fi28 & -459)
                                                }
                                            }
                                            if (!false) {
                                                switch "foo bar" {
                                                case "hello":
                                                    if false {
                                                        _ = len(v13)
                                                        fmt.Println(v29, fi12)
                                                        v6 -= 248
                                                    } else {
                                                        var v33 int8 = 1
                                                        if (-1 <= -9223372036854775808) {
                                                            v31 = (fi28 | int(v33))
                                                            v34 := GenS2{F0: (v33 + v33)}
                                                            _ = v34
                                                        }
                                                        switch (fi28 * (-(fi12))) {
                                                        case 127:
                                                            if len(v13) > 0 {
                                                                v13[0] = (-(v31))
                                                            }
                                                        case 42:
                                                            if len(v13) > 0 {
                                                                v13[0] = fi12
                                                            }
                                                        case 0:
                                                            if len(v13) >= 2 {
                                                                v13 = v13[1:len(v13)]
                                                            }
                                                        default:
                                                            var v35 float64 = float64((1e10 - 0.0))
                                                            _ = v35
                                                        }
                                                        _ = v33
                                                    }
                                                case "a":
                                                    if (-1 >= -128) {
                                                        if false {
                                                            fmt.Println(fi28, v29, fi28)
                                                            v36 := GenS2{F0: 0}
                                                            v37 := GenS3{F0: 18446744073709551615, F1: 0}
                                                            _ = v36
                                                            _ = v37
                                                        } else {
                                                            switch "a" {
                                                            case "":
                                                                if (-1 > -91) {
                                                                    var v38 uint = 4294967295
                                                                    if len(v13) >= 2 {
                                                                        v13 = v13[1:len(v13)]
                                                                    }
                                                                    _ = v38
                                                                } else {
                                                                    v39 := GenS2{F0: 1}
                                                                    v30 += (v6 % ((128) | 1))
                                                                    _ = v39
                                                                }
                                                            }
                                                            switch "a" {
                                                            case "foo bar":
                                                                v30 = 2
                                                            default:
                                                                switch ("hello" + "hello") {
                                                                case "":
                                                                    if (!(240 == v30)) {
                                                                        v30 -= v30
                                                                    }
                                                                case "foo bar":
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        v31 = 128
                                                                    default:
                                                                        v40 := GenS4{F0: -102, F1: 42}
                                                                        _ = v40
                                                                    }
                                                                default:
                                                                    v30 = 2
                                                                }
                                                            }
                                                            var v41 uint = 42
                                                            _ = v41
                                                        }
                                                    } else {
                                                        v6 += v30
                                                        v13 = append(v13, (fi28 * fi28))
                                                        v6 -= v30
                                                    }
                                                }
                                                var v42 int = (v31 + 836)
                                                switch 32767 {
                                                case 581:
                                                    if len(v13) >= 2 {
                                                        v13 = v13[1:len(v13)]
                                                    }
                                                case -9223372036854775808:
                                                    _ = len(v13)
                                                }
                                                var v43 int32 = -968
                                                var v44 float64 = 0.1
                                                _ = v42
                                                _ = v43
                                                _ = v44
                                            }
                                            if true {
                                                fmt.Println(v30, fi28)
                                            } else {
                                                if ((42 != 0) && (1 < 439)) {
                                                    switch fi12 {
                                                    case 0:
                                                        switch "\n\t" {
                                                        case "hello":
                                                            fmt.Println(v29, v31, fi12)
                                                        case "\n\t":
                                                            v6 = v6
                                                        }
                                                    }
                                                    var v45 float64 = float64(v29)
                                                    var v46 float64 = v45
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    var v47 int8 = 1
                                                    v6 = 196
                                                    v30 += 127
                                                    _ = v47
                                                }
                                                switch "a" {
                                                case "":
                                                    var v48 int8 = 1
                                                    _ = v48
                                                case "hello":
                                                    v29 = -2.25
                                                case "a":
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        if false {
                                                            v31 = 7
                                                            v30 = (v30 & (v30 << ((217) & 63)))
                                                        }
                                                    case "hello":
                                                        if (false && true) {
                                                            v49 := GenS5{F0: "foo bar", F1: "\n\t", F2: 1000}
                                                            _ = v49
                                                        }
                                                    }
                                                default:
                                                    v6 = v6
                                                }
                                            }
                                        }
                                        v50 := GenS5{F0: (("a" + "foo bar") + ("foo bar" + "\n\t")), F1: "foo bar", F2: 0}
                                        var v51 int32 = -1
                                        _ = v31
                                        _ = v50
                                        _ = v51
                                    } else {
                                        var v52 int = (fi12 - 1)
                                        v52 -= (v52 | -506)
                                        v53 := GenS5{F0: ("foo bar" + ""), F1: "", F2: -291}
                                        _ = v52
                                        _ = v53
                                    }
                                } else {
                                    v54 := GenS3{F0: 0, F1: 20}
                                    _ = len(v13)
                                    _ = len(v13)
                                    _ = v54
                                }
                                var v55 uint32 = 680
                                if (!true) {
                                    v55 = 42
                                } else {
                                    if (((v30 - (v6 >> ((843) & 63))) != v30) || true) {
                                        var v56 bool = true
                                        _ = v56
                                    } else {
                                        var v57 int32 = 1
                                        var v58 uint64 = 860
                                        _ = v57
                                        _ = v58
                                    }
                                    v13 = append(v13, (fi28 / ((fi28) | 1)))
                                    if len(v13) > 0 {
                                        v13[0] = (fi28 & (-(fi28)))
                                    }
                                }
                                _ = v55
                            }
                            v13 = append(v13, (fi12 / ((-9223372036854775808) | 1)))
                            var v59 float64 = (-2.25 + 3.14159)
                            if (-2.25 != ((((-2.25 * (float32((float64(((v29 * v29) * v29)) + v59)) - -2.25)) - v29) * (v29 + v29)) - float32(float64((-2.25 / 3.14159))))) {
                                v13 = append(v13, -158)
                            }
                            if false {
                                v60 := GenS6{F0: (-(fi12)), F1: -2147483648, F2: false}
                                if true {
                                    v61 := GenS7{F0: 75, F1: -32768, F2: 1}
                                    if true {
                                        var v62 int8 = -59
                                        _ = v62
                                    } else {
                                        fmt.Println(v59)
                                        v30 = v30
                                    }
                                    var v63 int8 = 127
                                    _ = v61
                                    _ = v63
                                } else {
                                    if (100 == 0) {
                                        var v64 uint8 = (v30 ^ 1)
                                        _ = v64
                                    }
                                    var v65 uint64 = 988
                                    if (!(2 == 4294967295)) {
                                        switch (fi12 << ((255) & 63)) {
                                        case -9223372036854775808:
                                            var v66 int16 = -1
                                            _ = v66
                                        default:
                                            v65 = v65
                                        }
                                        v67 := GenS6{F0: 969, F1: -2147483648, F2: false}
                                        _ = v67
                                    }
                                    _ = v65
                                }
                                if ((false && (v30 == (v30 >> ((4294967295) & 63)))) || false) {
                                    switch (^(fi12)) {
                                    case 2147483647:
                                        if (!(4294967295 >= 1)) {
                                            v68 := GenS6{F0: 0, F1: -119, F2: true}
                                            var v69 int16 = 32767
                                            _ = v68
                                            _ = v69
                                        } else {
                                            switch (("a" + "a") + "a") {
                                            case "foo bar":
                                                if true {
                                                    v13 = append(v13, -2147483648)
                                                    v29 = v29
                                                } else {
                                                    var v70 uint8 = v30
                                                    v6 -= v30
                                                    _ = v70
                                                }
                                            default:
                                                v59 = v59
                                            }
                                            v59 = 0.1
                                        }
                                    default:
                                        if len(v13) > 0 {
                                            v13[0] = fi12
                                        }
                                    }
                                    v6 = v6
                                    var v71 bool = true
                                    _ = v71
                                }
                                _ = v60
                            }
                            _ = v30
                            _ = v59
                        }
                        if len(v13) > 0 {
                            v13[0] = 622
                        }
                        _ = fi28
                        _ = v29
                    }
                } else {
                    var v72 string = ""
                    v6 = (v6 % (((v6 * 127)) | 1))
                    for fi73 := 0; fi73 < 6; fi73++ {
                        switch fi12 {
                        case 42:
                            switch "" {
                            case "a":
                                var v74 int64 = -9223372036854775808
                                _ = v74
                            case "foo bar":
                                if ("" != ("foo bar" + v72)) {
                                    if len(v13) > 0 {
                                        v13[0] = fi12
                                    }
                                }
                            case "":
                                if true {
                                    var v75 bool = (385 != 941)
                                    v76 := GenS7{F0: 1, F1: -805, F2: 18446744073709551615}
                                    v72 = "hello"
                                    _ = v75
                                    _ = v76
                                } else {
                                    if (!false) {
                                        v77 := GenS4{F0: 42, F1: 811}
                                        _ = v77
                                    }
                                    if len(v13) > 0 {
                                        v13[0] = (fi73 % ((fi73) | 1))
                                    }
                                    switch v72 {
                                    case "hello":
                                        v78 := GenS7{F0: 42, F1: -32768, F2: 18446744073709551615}
                                        _ = v78
                                    default:
                                        v72 = ((v72 + "") + "a")
                                    }
                                    v79 := GenS8{F0: (^(fi73))}
                                    v6 = (v6 & (^(v6)))
                                    _ = v79
                                }
                            default:
                                switch (v72 + ("\n\t" + (v72 + v72))) {
                                case "\n\t":
                                    v6 = (^(v6))
                                default:
                                    switch (fi73 * fi12) {
                                    case -321:
                                        if false {
                                            v72 = "hello"
                                        }
                                    default:
                                        v80 := GenS4{F0: 1, F1: 2}
                                        _ = v80
                                    }
                                }
                            }
                        }
                        _ = fi73
                    }
                    _ = v72
                }
            }
            _ = fi12
            _ = v13
        }
        v81 := GenS4{F0: 1, F1: 1}
        if false {
            v6 += (v6 / ((v6) | 1))
        } else {
            fmt.Println(v6)
            v82 := []int{2147483647, -1, 32767}
            _ = v82
        }
        _ = v81
    }
    _ = v6
    return 255
}
func genFn1(p0 uint32) uint8 {
    fmt.Println(p0)
    v83 := GenS9{F0: 100}
    var v84 bool = (4294967295 <= genFn0())
    _ = p0
    _ = v83
    _ = v84
    return 105
}

func main() {
    var v1 uint = 42
    var v2 uint32 = 907
    v3 := []uint32{2147483648, v2}
    v4 := []int{42, 660}
    v5 := []string{"hello", "a", "foo bar", "hello"}
    var v85 uint = (v1 / ((695) | 1))
    var v86 uint8 = 127
    var v87 bool = true
    if ((0 >= 237) || ("hello" == "hello")) {
        fmt.Println(v86, v2, v86)
        for fi88 := 0; fi88 < 10; fi88++ {
            if false {
                for fi89 := 0; fi89 < 2; fi89++ {
                    if v87 {
                        if len(v5) > 0 {
                            v5[0] = "hello"
                        }
                        var v90 int16 = 42
                        _ = v90
                    } else {
                        v1 += 2
                        v91 := GenS10{F0: 0, F1: uint64(v2), F2: 1000}
                        _ = v91
                    }
                    _ = fi89
                }
                switch (fi88 >> ((genFn0()) & 63)) {
                case 9223372036854775807:
                    v92 := GenS11{F0: (("hello" + "foo bar") + (("foo bar" + "a") + "foo bar")), F1: -32768, F2: 255}
                    _ = v92
                case -80:
                    v85 = genFn0()
                case -507:
                    switch ("foo bar" + ("\n\t" + ("hello" + "a"))) {
                    case "foo bar":
                        var v93 int16 = -429
                        _ = v93
                    case "hello":
                        if false {
                            v87 = (78 == -128)
                            v2 = (v2 << ((42) & 63))
                        } else {
                            if v87 {
                                fmt.Println(v1, v2, v1)
                            }
                        }
                    }
                default:
                    fmt.Println(v1, fi88, fi88)
                }
                v5 = append(v5, "")
            } else {
                var v94 bool = v87
                _ = v94
            }
            v87 = (9223372036854775807 != 9223372036854775807)
            v95 := GenS6{F0: (fi88 + 4294967295), F1: -1, F2: (!((^(v1)) >= v85))}
            _ = fi88
            _ = v95
        }
        for fi96 := 0; fi96 < 18; fi96++ {
            fmt.Println(fi96, v86, fi96)
            _ = fi96
        }
    }
    for fi97 := 0; fi97 < 2; fi97++ {
        switch (fi97 % ((32768) | 1)) {
        case -196:
            if len(v3) > 0 {
                v3[0] = 65535
            }
        default:
            if (1 <= -590) {
                v2 -= (v2 / ((v2) | 1))
            }
        }
        var v98 string = "foo bar"
        switch fi97 {
        case 32768:
            var v99 int32 = -1
            _ = v99
        case 127:
            fmt.Println(v85)
        case 32767:
            v100 := GenS12{F0: -32768, F1: (fi97 & -9223372036854775808), F2: 42}
            _ = v100
        default:
            for fi101 := 0; fi101 < 9; fi101++ {
                if ((((-2.25 - 0.0) / 1.5) != 3.14159) || (!(104 == 24))) {
                    v1 -= 767
                    var v102 float32 = 0.1
                    _ = v102
                } else {
                    v85 += (v85 * (v85 | v1))
                    var v103 uint8 = (^(v86))
                    v104 := GenS13{F0: 1, F1: (v2 - v2), F2: 118}
                    _ = v103
                    _ = v104
                }
                fmt.Println(fi97, v87)
                v85 += v1
                var v105 bool = (4294967295 < 0)
                _ = fi101
                _ = v105
            }
        }
        _ = fi97
        _ = v98
    }
    var v106 string = "a"
    for fi107 := 0; fi107 < 2; fi107++ {
        v1 -= (v85 % ((uint(v86)) | 1))
        fmt.Println(v87, v86)
        _ = fi107
    }
    var v108 uint32 = 90
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (((-2.25 - 1.5) + (1e10 / 0.1)) != 0.1) {
        v108 = (v108 & (v108 << ((v1) & 63)))
        v109 := GenS6{F0: 32768, F1: 2147483647, F2: (((-32768 == -485) && (!(v87 || true))) || v87)}
        fmt.Println(v87, v2, v1)
        _ = v109
    } else {
        for fi110 := 0; fi110 < 4; fi110++ {
            var v111 string = v106
            if (-32768 >= 1000) {
                v2 = (v108 % ((v2) | 1))
                v86 = uint8(v1)
                switch ("" + v111) {
                case "hello":
                    _ = len(v4)
                default:
                    for fi112 := 0; fi112 < 13; fi112++ {
                        switch "" {
                        case "a":
                            var v113 int = (fi110 ^ -13)
                            _ = v113
                        case "":
                            if false {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        case "\n\t":
                            fmt.Println(fi110)
                            v3 = append(v3, 0)
                        }
                        v87 = false
                        _ = fi112
                    }
                }
            }
            if len(v4) > 0 {
                v4[0] = (fi110 - fi110)
            }
            _ = fi110
            _ = v111
        }
    }
    for fi114 := 0; fi114 < 12; fi114++ {
        var v115 int64 = 1
        v116 := GenS12{F0: -780, F1: fi114, F2: 0}
        if v87 {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            fmt.Println(v115)
        } else {
            var v117 uint = genFn0()
            for fi118 := 0; fi118 < 3; fi118++ {
                v119 := GenS8{F0: (-(fi118))}
                var v120 uint16 = uint16(v108)
                v121 := GenS12{F0: -32768, F1: fi114, F2: (v120 ^ v120)}
                var v122 uint64 = 933
                _ = fi118
                _ = v119
                _ = v120
                _ = v121
                _ = v122
            }
            var v123 int = 0
            _ = v117
            _ = v123
        }
        _ = fi114
        _ = v115
        _ = v116
    }
    fmt.Println(v2, v108)
    v108 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v106:", v106)
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v106
    _ = v108
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v85
    _ = v86
    _ = v87
}
