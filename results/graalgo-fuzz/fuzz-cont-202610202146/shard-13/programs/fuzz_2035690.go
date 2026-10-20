package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int32; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 int16; F1 uint16; F2 string }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int8; F1 int64; F2 string }
type GenS9 struct { F0 uint16; F1 uint32 }
type GenS10 struct { F0 uint32; F1 uint }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int32; F1 int8 }
type GenS13 struct { F0 uint64; F1 uint8; F2 string }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 uint16; F1 int64 }
type GenS16 struct { F0 int16; F1 int16 }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS19 struct { F0 int16; F1 string; F2 int }
type GenS20 struct { F0 uint64; F1 int }
type GenS21 struct { F0 string; F1 uint }
type GenS22 struct { F0 bool }
func genFn0() bool {
    v5 := []int{}
    for fi6 := 0; fi6 < 13; fi6++ {
        fmt.Println(fi6)
        _ = fi6
    }
    var v7 bool = (255 == 814)
    fmt.Println(v7)
    switch "\n\t" {
    case "":
        v8 := GenS1{F0: 42, F1: 1, F2: 65535}
        _ = v8
    case "\n\t":
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 bool = (65535 == 255)
            if len(v5) > 0 {
                v5[0] = -830
            }
            _ = fi9
            _ = v10
        }
    }
    _ = v5
    _ = v7
    return (v7 || v7)
}
func genFn1(p0 int8, p1 int) uint16 {
    var v11 float32 = float32(0.0)
    v12 := []int{}
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    switch "a" {
    case "":
        switch (p1 - p1) {
        case 32768:
            switch (("hello" + ("hello" + ("" + "a"))) + ("\n\t" + "foo bar")) {
            case "":
                switch 42 {
                case 0:
                    var v13 float64 = (float64(v11) + 1e10)
                    _ = v13
                case 4294967295:
                    _ = len(v12)
                }
            case "hello":
                v14 := GenS2{F0: 42, F1: (-(p1))}
                _ = v14
            case "a":
                if (9223372036854775807 >= -977) {
                    var v15 float32 = 0.0
                    _ = v15
                } else {
                    v11 = v11
                }
            default:
                p1 += (p1 << ((4294967295) & 63))
            }
        }
    default:
        v16 := GenS1{F0: p0, F1: 2147483648, F2: 65535}
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 1
}

func main() {
    v1 := []uint32{}
    var v2 bool = (127 < -1)
    var v3 uint32 = 2
    var v4 uint32 = (v3 * (v3 & v3))
    switch -9223372036854775808 {
    case 980:
        if v2 {
            switch "a" {
            case "a":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
    case 728:
        for fi17 := 0; fi17 < 17; fi17++ {
            v18 := GenS1{F0: -128, F1: (v3 << ((320) & 63)), F2: 42}
            _ = fi17
            _ = v18
        }
    case -1:
        v19 := GenS1{F0: 0, F1: 55, F2: (v4 / (((v3 % (((v3 / (((^(v3))) | 1))) | 1))) | 1))}
        _ = v19
    }
    var v20 bool = (false || true)
    v3 -= 2147483647
    switch 32768 {
    case 9223372036854775807:
        var v21 float32 = 0.0
        _ = v21
    case 4294967295:
        var v22 int32 = -1
        _ = v22
    case -1:
        for fi23 := 0; fi23 < 17; fi23++ {
            v20 = ("" != ("\n\t" + ("\n\t" + "hello")))
            var v24 int32 = -1
            _ = fi23
            _ = v24
        }
    default:
        var v25 uint32 = (^(v4))
        _ = v25
    }
    if v20 {
        if true {
            v1 = append(v1, 42)
            _ = len(v1)
        }
    }
    if v20 {
        fmt.Println(v4, v3)
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        for fi27 := 0; fi27 < 14; fi27++ {
            if false {
                v4 += v3
                fmt.Println(v3, fi27)
                var v28 uint64 = 757
                switch "" {
                case "":
                    v1 = append(v1, (v3 / (((v3 % (((v3 | (v3 % ((v4) | 1)))) | 1))) | 1)))
                default:
                    fmt.Println(fi26, v20)
                    if (42 > -32768) {
                        _ = len(v1)
                    } else {
                        fmt.Println(v2)
                        _ = len(v1)
                    }
                    if true {
                        v1 = append(v1, 701)
                        if (fi26 != (fi27 * 128)) {
                            if (-2.25 < (((-2.25 + (((3.14159 / float64(float32((1e10 - 0.1)))) - 3.14159) / 1.0)) * (((-2.25 / 1e10) / -2.25) / 1e10)) + (-2.25 / 0.1))) {
                                var v29 int = fi27
                                switch ("hello" + ("" + "foo bar")) {
                                case "foo bar":
                                    var v30 uint8 = 186
                                    _ = v30
                                case "":
                                    var v31 int32 = -392
                                    _ = v31
                                case "\n\t":
                                    v2 = (v28 >= (^(v28)))
                                default:
                                    v32 := GenS2{F0: v3, F1: 236}
                                    _ = v32
                                }
                                switch "" {
                                case "a":
                                    v4 -= 1
                                case "hello":
                                    var v33 bool = (v20 && (true || true))
                                    _ = v33
                                default:
                                    if v20 {
                                        v34 := GenS1{F0: 127, F1: 1, F2: v4}
                                        var v35 int = (fi27 ^ v29)
                                        var v36 float64 = ((float64(1.5) + 1.5) * 0.1)
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                    } else {
                                        var v37 uint32 = 1
                                        switch ("a" + "") {
                                        case "a":
                                            if ((v4 ^ v3) <= v4) {
                                                v37 += (v3 / ((294) | 1))
                                                fmt.Println(v3, v2)
                                                if (642 == 32767) {
                                                    var v38 int64 = -1
                                                    v39 := GenS3{F0: 0, F1: 127, F2: 127}
                                                    var v40 uint8 = 175
                                                    _ = v38
                                                    _ = v39
                                                    _ = v40
                                                }
                                                v41 := GenS1{F0: 100, F1: v37, F2: (v37 - v3)}
                                                v42 := GenS3{F0: 208, F1: -1, F2: uint8(v28)}
                                                _ = v41
                                                _ = v42
                                            } else {
                                                var v43 uint32 = 0
                                                v28 = 18446744073709551615
                                                v3 += (v37 * v43)
                                                _ = v43
                                            }
                                        case "\n\t":
                                            fmt.Println(v3, v2, v28)
                                        default:
                                            v1 = append(v1, v4)
                                        }
                                        _ = v37
                                    }
                                }
                                _ = v29
                            } else {
                                var v44 int64 = -587
                                var v45 int64 = -1
                                switch (fi27 ^ (^(fi26))) {
                                case 32768:
                                    switch (("a" + ((("a" + "a") + "foo bar") + "foo bar")) + "") {
                                    case "":
                                        v2 = v2
                                    }
                                case 1:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case 127:
                                    v2 = (280 != -2147483648)
                                default:
                                    v1 = append(v1, (^(v3)))
                                }
                                _ = v44
                                _ = v45
                            }
                            var v46 uint16 = 255
                            _ = v46
                        }
                    } else {
                        var v47 float32 = ((0.0 + ((3.14159 - 1e10) - (((1.5 * 0.0) + 0.1) / 1.0))) * -2.25)
                        v28 -= 4294967295
                        v48 := GenS4{F0: 42, F1: 1}
                        switch fi27 {
                        case 4294967295:
                            v47 -= v47
                        default:
                            v4 -= (v3 / (((v4 >> ((0) & 63))) | 1))
                        }
                        _ = v47
                        _ = v48
                    }
                }
                _ = v28
            }
            if (!(-92 == 100)) {
                v3 = 518
                if false {
                    fmt.Println(v2)
                }
            } else {
                if v2 {
                    switch ("a" + "") {
                    case "":
                        switch (fi26 * (fi27 % ((fi27) | 1))) {
                        case -203:
                            v3 += (v4 << ((255) & 63))
                        case 9223372036854775807:
                            v49 := GenS5{F0: (v4 >> ((765) & 63)), F1: v2}
                            _ = v49
                        case -477:
                            switch 862 {
                            case 42:
                                if (!false) {
                                    var v50 uint32 = (v4 & (^(v3)))
                                    if false {
                                        _ = len(v1)
                                    }
                                    v3 = 1
                                    _ = v50
                                }
                            case -123:
                                var v51 uint64 = 301
                                _ = v51
                            }
                        }
                    case "hello":
                        switch ((("hello" + "foo bar") + "hello") + "") {
                        case "a":
                            fmt.Println(v4)
                        default:
                            fmt.Println(v3, fi27, fi27)
                        }
                    default:
                        v2 = ((!(v2 && true)) && (int16(fi26) == -1))
                    }
                    if v2 {
                        if v2 {
                            v3 += v3
                        }
                        v3 = (^(v4))
                        if (!v20) {
                            v52 := GenS6{F0: -10, F1: 0, F2: "a"}
                            switch -9223372036854775808 {
                            case 42:
                                v2 = (4294967295 == 4294967295)
                            default:
                                fmt.Println(v3, v20, v2)
                            }
                            if (v2 || (!v20)) {
                                var v53 uint8 = 1
                                _ = v53
                            } else {
                                v1 = append(v1, v3)
                                v3 -= 651
                                if v2 {
                                    if (false && ((138 < 0) && v2)) {
                                        var v54 uint16 = 32767
                                        if v20 {
                                            v55 := GenS3{F0: 0, F1: int8(v54), F2: 126}
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            if (-461 < 1) {
                                                var v56 int16 = 96
                                                var v57 int8 = 127
                                                v1 = append(v1, 65535)
                                                _ = v56
                                                _ = v57
                                            } else {
                                                v1 = append(v1, v4)
                                                if len(v1) > 0 {
                                                    v1[0] = (v3 % ((v3) | 1))
                                                }
                                                v58 := GenS3{F0: -2147483648, F1: 1, F2: 163}
                                                _ = v58
                                            }
                                            _ = v55
                                        } else {
                                            var v59 float64 = float64((3.14159 + ((0.1 * ((0.1 * float32(0.0)) + 0.0)) * 0.1)))
                                            v4 += v4
                                            _ = v59
                                        }
                                        _ = v54
                                    }
                                    v4 = (^(v4))
                                }
                            }
                            _ = v52
                        } else {
                            v60 := GenS5{F0: v4, F1: v20}
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch fi27 {
                            case 459:
                                switch "a" {
                                case "foo bar":
                                    if false {
                                        if (false && (282 < 4294967295)) {
                                            var v61 float64 = ((0.0 * -2.25) - 3.14159)
                                            var v62 float64 = (v61 / v61)
                                            if false {
                                                _ = len(v1)
                                                switch ("foo bar" + ("a" + "hello")) {
                                                case "":
                                                    v4 = 42
                                                default:
                                                    switch "a" {
                                                    case "\n\t":
                                                        switch fi26 {
                                                        case 4294967295:
                                                            if v20 {
                                                                _ = len(v1)
                                                                v63 := GenS7{F0: 212}
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                _ = v63
                                                            }
                                                        case -422:
                                                            v64 := GenS6{F0: int16(v3), F1: 32767, F2: ""}
                                                            _ = v64
                                                        case -9223372036854775808:
                                                            v65 := GenS8{F0: -2, F1: 42, F2: (("\n\t" + "hello") + "")}
                                                            _ = v65
                                                        }
                                                    case "foo bar":
                                                        fmt.Println(v2, v2)
                                                    default:
                                                        v2 = false
                                                    }
                                                }
                                            } else {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                var v66 int = -9223372036854775808
                                                _ = v66
                                            }
                                            _ = v61
                                            _ = v62
                                        } else {
                                            var v67 int32 = -2147483648
                                            var v68 uint16 = 737
                                            _ = v67
                                            _ = v68
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = v4
                                        }
                                        v2 = ((!false) && true)
                                        v4 -= (v3 % ((962) | 1))
                                        switch (fi26 % ((fi26) | 1)) {
                                        case 439:
                                            var v69 int64 = -1
                                            fmt.Println(v20)
                                            v4 = v3
                                            _ = v69
                                        case 0:
                                            v1 = append(v1, 2147483648)
                                        case 323:
                                            switch (-(fi26)) {
                                            case 128:
                                                var v70 float32 = ((0.0 / -2.25) * 0.0)
                                                _ = v70
                                            case 0:
                                                var v71 int8 = -115
                                                _ = v71
                                            default:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                            v72 := GenS9{F0: 0, F1: v4}
                                            _ = v72
                                        default:
                                            _ = len(v1)
                                            switch (("foo bar" + "a") + (("hello" + "hello") + "hello")) {
                                            case "foo bar":
                                                switch "a" {
                                                case "a":
                                                    if (!((!(v20 || v20)) || v20)) {
                                                        switch (fi27 & fi27) {
                                                        case 127:
                                                            v73 := GenS8{F0: 0, F1: -742, F2: "foo bar"}
                                                            _ = v73
                                                        case 4294967295:
                                                            if (v20 || v20) {
                                                                v74 := GenS10{F0: (v3 >> ((2) & 63)), F1: 255}
                                                                _ = len(v1)
                                                                var v75 string = "hello"
                                                                _ = v74
                                                                _ = v75
                                                            }
                                                        case 32768:
                                                            fmt.Println(v2, v3, v4)
                                                        default:
                                                            v4 += 1
                                                            var v76 uint8 = 42
                                                            var v77 string = "a"
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                        v78 := GenS11{F0: -27}
                                                        v79 := GenS9{F0: 255, F1: v3}
                                                        fmt.Println(v3)
                                                        _ = v78
                                                        _ = v79
                                                    } else {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        if false {
                                                            v3 -= v3
                                                        } else {
                                                            var v80 uint8 = 1
                                                            var v81 int16 = -1
                                                            v3 = v4
                                                            _ = v80
                                                            _ = v81
                                                        }
                                                        switch "\n\t" {
                                                        case "":
                                                            v82 := GenS12{F0: -1, F1: -128}
                                                            _ = v82
                                                        }
                                                    }
                                                case "foo bar":
                                                    v3 += (v4 + v3)
                                                default:
                                                    v1 = append(v1, (v3 / ((4294967295) | 1)))
                                                }
                                            case "":
                                                if v2 {
                                                    _ = len(v1)
                                                    v83 := GenS12{F0: -900, F1: 0}
                                                    _ = v83
                                                } else {
                                                    v84 := GenS13{F0: 42, F1: 42, F2: ("\n\t" + "foo bar")}
                                                    var v85 int8 = -111
                                                    v86 := GenS14{F0: 24}
                                                    _ = v84
                                                    _ = v85
                                                    _ = v86
                                                }
                                            case "\n\t":
                                                var v87 string = ("foo bar" + "a")
                                                _ = v87
                                            }
                                        }
                                    } else {
                                        switch ("\n\t" + "a") {
                                        case "":
                                            var v88 float32 = (float32(1e10) * -2.25)
                                            _ = v88
                                        case "\n\t":
                                            v89 := GenS1{F0: -59, F1: (v3 - 0), F2: 0}
                                            _ = v89
                                        case "hello":
                                            var v90 string = ""
                                            if len(v1) > 0 {
                                                v1[0] = v3
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            fmt.Println(v4)
                                            if false {
                                                var v91 uint64 = 42
                                                if false {
                                                    fmt.Println(v90)
                                                }
                                                _ = v91
                                            } else {
                                                switch v90 {
                                                case "":
                                                    v1 = append(v1, (v3 ^ 2147483648))
                                                case "foo bar":
                                                    if false {
                                                        v92 := GenS15{F0: genFn1(int8(fi27), fi27), F1: -1}
                                                        v93 := GenS4{F0: 635, F1: 1}
                                                        var v94 int = fi27
                                                        _ = v92
                                                        _ = v93
                                                        _ = v94
                                                    }
                                                }
                                            }
                                            _ = v90
                                        }
                                    }
                                case "hello":
                                    switch "a" {
                                    case "foo bar":
                                        var v95 uint32 = 1
                                        _ = v95
                                    case "hello":
                                        _ = len(v1)
                                    default:
                                        var v96 int16 = 32767
                                        _ = v96
                                    }
                                case "\n\t":
                                    _ = len(v1)
                                }
                            case 0:
                                switch (fi27 + (fi26 & (fi27 - fi26))) {
                                case -530:
                                    v20 = v20
                                case 32767:
                                    if len(v1) > 0 {
                                        v1[0] = 42
                                    }
                                }
                            default:
                                v97 := GenS16{F0: -1, F1: 1000}
                                _ = v97
                            }
                            _ = v60
                        }
                    } else {
                        v4 -= (v4 + (v3 - v3))
                        v2 = false
                        v98 := GenS10{F0: (v4 / (((v3 ^ (^(v3)))) | 1)), F1: 1}
                        _ = v98
                    }
                } else {
                    var v99 int16 = -1
                    if true {
                        switch fi26 {
                        case 34:
                            var v100 uint16 = 0
                            _ = v100
                        case 2147483647:
                            var v101 uint16 = 328
                            _ = v101
                        }
                        if len(v1) > 0 {
                            v1[0] = 0
                        }
                        fmt.Println(v4)
                    } else {
                        v102 := GenS14{F0: 0}
                        fmt.Println(v4, v4, v4)
                        fmt.Println(v99, fi27)
                        if len(v1) > 0 {
                            v1[0] = v4
                        }
                        _ = v102
                    }
                    _ = v99
                }
                v4 = (v3 % ((v3) | 1))
            }
            _ = fi27
        }
        if (v20 || true) {
            v103 := GenS17{F0: 1}
            if v2 {
                switch "\n\t" {
                case "foo bar":
                    switch -2147483648 {
                    case -294:
                        switch "" {
                        case "hello":
                            for fi104 := 0; fi104 < 14; fi104++ {
                                if false {
                                    if v20 {
                                        var v105 int64 = -1
                                        fmt.Println(v20)
                                        var v106 uint32 = 0
                                        _ = v105
                                        _ = v106
                                    }
                                }
                                v107 := GenS18{F0: 128, F1: -1, F2: 127}
                                _ = fi104
                                _ = v107
                            }
                        default:
                            switch fi26 {
                            case 0:
                                v20 = (1 >= -1)
                            }
                        }
                    default:
                        if (!(float64((0.1 / (-2.25 - float32(1e10)))) == (1e10 * (1e10 - 3.14159)))) {
                            v108 := GenS19{F0: 32767, F1: "hello", F2: -1}
                            if len(v1) > 0 {
                                v1[0] = (v4 & v3)
                            }
                            if true {
                                switch ("" + "") {
                                case "foo bar":
                                    v109 := GenS13{F0: 308, F1: 128, F2: (("" + "hello") + ("" + "a"))}
                                    _ = v109
                                case "":
                                    if len(v1) > 0 {
                                        v1[0] = v4
                                    }
                                }
                                var v110 string = "hello"
                                _ = v110
                            }
                            _ = v108
                        }
                    }
                case "\n\t":
                    fmt.Println(v2, v2)
                default:
                    v3 = (v4 - 1)
                }
                for fi111 := 0; fi111 < 12; fi111++ {
                    fmt.Println(fi26, v2)
                    _ = fi111
                }
            }
            _ = v103
        }
        var v112 int32 = 42
        _ = fi26
        _ = v112
    }
    v4 = v4
    fmt.Println(v4)
    if (42 < 4294967295) {
        v113 := GenS20{F0: 18446744073709551615, F1: -2147483648}
        _ = v113
    } else {
        var v114 uint16 = 2
        if v2 {
            v2 = v2
        }
        for fi115 := 0; fi115 < 6; fi115++ {
            var v116 float32 = (3.14159 * 1e10)
            v117 := GenS21{F0: ("\n\t" + "hello"), F1: 4294967295}
            _ = fi115
            _ = v116
            _ = v117
        }
        _ = v114
    }
    v118 := GenS21{F0: "", F1: 728}
    var v119 uint32 = v3
    var v120 int = -9223372036854775808
    v119 = 112
    v4 = v3
    v121 := GenS22{F0: false}
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v119:", v119)
    fmt.Println("v120:", v120)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v118
    _ = v119
    _ = v120
    _ = v121
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
