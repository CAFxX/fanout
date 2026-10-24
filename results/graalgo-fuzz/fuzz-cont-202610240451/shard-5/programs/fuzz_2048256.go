package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int64; F2 bool }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32; F1 uint16; F2 int16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 bool; F1 int64; F2 int64 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 string; F1 uint; F2 int8 }
type GenS11 struct { F0 string; F1 uint }
type GenS12 struct { F0 int64; F1 int8; F2 uint16 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 uint8; F1 uint8; F2 uint64 }
func genFn0(p0 int32) uint8 {
    var v6 string = ("a" + "foo bar")
    switch v6 {
    case "foo bar":
        fmt.Println(v6)
    default:
        var v7 uint64 = 18446744073709551615
        _ = v7
    }
    _ = p0
    _ = v6
    return 42
}

func main() {
    var v1 float32 = -2.25
    v2 := []uint32{430}
    v3 := make(map[string]bool)
    var v4 uint8 = 128
    var v5 uint64 = 767
    var v8 uint64 = v5
    v4 -= (v4 % ((0) | 1))
    v8 += (v5 * uint64(v4))
    if false {
        for fi9 := 0; fi9 < 12; fi9++ {
            for fi10 := 0; fi10 < 6; fi10++ {
                v1 = (((v1 / 1.5) * v1) / ((0.1 / (v1 / (v1 - (v1 - 0.0)))) / (1.5 - 1e10)))
                v5 = 18446744073709551615
                _ = fi10
            }
            v4 -= 1
            for fi11 := 0; fi11 < 20; fi11++ {
                v4 = (^(v4))
                if false {
                    fmt.Println(fi9)
                }
                switch 42 {
                case -724:
                    var v12 uint16 = 255
                    _ = v12
                case 128:
                    _ = len(v3)
                case 748:
                    if true {
                        v4 -= 0
                        if false {
                            var v13 float64 = 3.14159
                            v1 = 0.1
                            fmt.Println(v4)
                            _ = v13
                        }
                        v1 += v1
                    }
                default:
                    v8 += (^(v8))
                }
                if true {
                    var v14 uint16 = 288
                    if (!(2147483647 >= 51)) {
                        if true {
                            v4 -= v4
                        }
                        fmt.Println(v5, v5)
                        switch (("foo bar" + "a") + (("a" + "hello") + (("foo bar" + "hello") + ("foo bar" + "")))) {
                        case "hello":
                            fmt.Println(fi9)
                        }
                        var v15 float32 = (((v1 * v1) + (((0.1 / v1) * ((v1 * v1) / (-2.25 / 1e10))) + v1)) - (v1 / ((float32(0.0) * v1) / (v1 / v1))))
                        var v16 int32 = -183
                        v1 = v15
                        _ = v15
                        _ = v16
                    } else {
                        v17 := GenS1{F0: uint(v4)}
                        var v18 bool = (true && true)
                        var v19 bool = true
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    v2 = append(v2, 0)
                    if (3.14159 > ((float64(((v1 / (v1 * v1)) / float32(3.14159))) / ((float64(v1) - 0.1) + (1e10 / 1.5))) / ((0.0 / float64(v1)) + ((((1e10 * (float64(-2.25) * (float64((v1 + (((v1 - v1) * 0.0) * 1e10))) * (1.5 * ((3.14159 * 1.5) - 0.1))))) * 1.5) * -2.25) / (0.0 - -2.25))))) {
                        fmt.Println(v4, fi9, v8)
                        var v20 string = "hello"
                        _ = v20
                    } else {
                        switch "foo bar" {
                        case "hello":
                            switch (("a" + ((("a" + "\n\t") + (("foo bar" + "foo bar") + "\n\t")) + (((("" + "\n\t") + "") + (("hello" + "a") + "foo bar")) + ""))) + "\n\t") {
                            case "\n\t":
                                if len(v2) > 0 {
                                    v2[0] = 284
                                }
                            case "a":
                                var v21 float32 = (1e10 - v1)
                                _ = v21
                            case "hello":
                                var v22 int8 = -86
                                _ = v22
                            default:
                                var v23 uint8 = (v4 - (^(v4)))
                                _ = v23
                            }
                        case "\n\t":
                            fmt.Println(v5, v4)
                        }
                    }
                    _ = v14
                }
                _ = fi11
            }
            _ = fi9
        }
        if (-1 >= 0) {
            v24 := v3[""]
            _ = v24
        }
    }
    v25 := GenS1{F0: 809}
    switch "\n\t" {
    case "hello":
        if (!false) {
            if (!true) {
                v4 = genFn0(150)
            } else {
                v5 = (v5 & 4294967295)
            }
            v1 = ((float32(0.0) / float32(((3.14159 - (-2.25 / float64((-2.25 + ((v1 / v1) / v1))))) + ((0.0 / 1.0) / ((1.5 + (0.1 * (3.14159 * -2.25))) - 1.5))))) * (v1 * v1))
            switch 2147483647 {
            case 889:
                switch "" {
                case "a":
                    v26 := v3[("foo bar" + "\n\t")]
                    _ = v26
                case "\n\t":
                    var v27 int16 = -32768
                    _ = v27
                case "foo bar":
                    v2 = append(v2, 2147483647)
                default:
                    v28 := GenS1{F0: 2}
                    _ = v28
                }
            case 9223372036854775807:
                for fi29 := 0; fi29 < 15; fi29++ {
                    switch fi29 {
                    case 1:
                        if true {
                            switch ("foo bar" + "foo bar") {
                            case "hello":
                                if (true || false) {
                                    for fi30 := 0; fi30 < 13; fi30++ {
                                        v4 = 42
                                        _ = fi30
                                    }
                                    for fi31 := 0; fi31 < 13; fi31++ {
                                        if (!("hello" != ("a" + ("\n\t" + "a")))) {
                                            v5 = v5
                                            if (float64(((3.14159 * v1) + (float32(1.5) / (((float32((3.14159 - 0.1)) + ((((v1 + float32(-2.25)) - ((v1 / v1) * float32((0.1 + float64(v1))))) - (v1 - v1)) - (1.5 - v1))) - v1) + v1)))) >= 0.1) {
                                                fmt.Println(v1)
                                                switch fi29 {
                                                case -239:
                                                    v32 := GenS1{F0: 0}
                                                    _ = v32
                                                default:
                                                    _ = len(v3)
                                                }
                                            }
                                            v8 += (v8 + 1)
                                        }
                                        _ = fi31
                                    }
                                } else {
                                    if false {
                                        if (42 >= -2147483648) {
                                            v33 := GenS2{F0: 2147483648, F1: 336, F2: (1.5 > (0.1 * -2.25))}
                                            _ = v33
                                        } else {
                                            v5 = 42
                                        }
                                        fmt.Println(fi29, v4)
                                        var v34 uint32 = 329
                                        _ = v34
                                    }
                                    v35 := GenS2{F0: 2147483647, F1: 1, F2: (!true)}
                                    fmt.Println(v8, v5, v8)
                                    _ = v35
                                }
                            default:
                                var v36 float32 = ((v1 + float32((1e10 + -2.25))) * (v1 - (3.14159 / v1)))
                                _ = v36
                            }
                        }
                    case 656:
                        fmt.Println(v5, v1)
                    default:
                        switch "hello" {
                        case "hello":
                            if (false && true) {
                                v37 := GenS1{F0: 1}
                                _ = v37
                            } else {
                                _ = len(v3)
                                var v38 int64 = 42
                                _ = v38
                            }
                        }
                    }
                    for fi39 := 0; fi39 < 4; fi39++ {
                        if (!false) {
                            switch "a" {
                            case "a":
                                switch -1 {
                                case -136:
                                    switch (-(fi39)) {
                                    case 32768:
                                        if (496 != 396) {
                                            fmt.Println(v4)
                                            v1 = (float32((((0.0 / 0.1) * 1.5) * float64(float32(-2.25)))) - v1)
                                        } else {
                                            v4 = 219
                                            var v40 int16 = -786
                                            v41 := v3["\n\t"]
                                            _ = v40
                                            _ = v41
                                        }
                                    }
                                case 127:
                                    v8 = v8
                                case -9223372036854775808:
                                    _ = len(v3)
                                default:
                                    v42 := GenS3{F0: -156, F1: 259}
                                    var v43 int64 = -9223372036854775808
                                    if false {
                                        v3["hello"] = ((v43 >> ((2) & 63)) != v43)
                                    } else {
                                        fmt.Println(v8)
                                    }
                                    _ = v42
                                    _ = v43
                                }
                            case "hello":
                                var v44 uint16 = 65535
                                _ = v44
                            default:
                                switch (fi39 % (((fi39 * fi39)) | 1)) {
                                case -111:
                                    switch -135 {
                                    case -1:
                                        v45 := GenS2{F0: 42, F1: 9223372036854775807, F2: (-1 > 127)}
                                        _ = v45
                                    case -9223372036854775808:
                                        v4 = (^(v4))
                                    default:
                                        var v46 uint64 = v8
                                        _ = v46
                                    }
                                case 429:
                                    switch "" {
                                    case "hello":
                                        if ((3.14159 / 3.14159) <= (float64(-2.25) * (0.1 + float64(1.5)))) {
                                            v4 -= 42
                                            v47 := GenS2{F0: 57, F1: 0, F2: true}
                                            _ = v47
                                        } else {
                                            if true {
                                                v48 := GenS4{F0: ("" + "foo bar")}
                                                _ = v48
                                            }
                                            v1 -= float32(0.1)
                                            v49 := v3["foo bar"]
                                            _ = v49
                                        }
                                    }
                                    switch "foo bar" {
                                    case "":
                                        v50 := GenS5{F0: 31}
                                        if true {
                                            fmt.Println(fi39)
                                            var v51 uint = 184
                                            switch 32767 {
                                            case 32768:
                                                v52 := GenS6{F0: 0, F1: 310, F2: 125}
                                                _ = v52
                                            default:
                                                var v53 float32 = v1
                                                _ = v53
                                            }
                                            _ = v51
                                        }
                                        _ = v50
                                    case "foo bar":
                                        v8 = v5
                                    case "a":
                                        v8 -= v8
                                    }
                                    v1 = float32(0.1)
                                    switch "\n\t" {
                                    case "a":
                                        v4 = 127
                                    default:
                                        fmt.Println(v4)
                                    }
                                case -424:
                                    v1 = float32(3.14159)
                                }
                            }
                            fmt.Println(v5, fi39, v4)
                            var v54 uint32 = uint32(v8)
                            v4 += v4
                            delete(v3, ("hello" + ""))
                            _ = v54
                        }
                        _ = fi39
                    }
                    _ = fi29
                }
            default:
                switch 2147483647 {
                case 32768:
                    _ = len(v2)
                case 42:
                    var v55 float32 = 0.0
                    _ = v55
                }
            }
        }
    }
    v1 += (v1 + (v1 + v1))
    if (42 > -715) {
        v4 += 1
        if true {
            v56 := GenS6{F0: 0, F1: 65535, F2: 1000}
            _ = v56
        } else {
            fmt.Println(v5)
        }
        v3[("" + "a")] = (1 >= 2)
    } else {
        switch -1 {
        case 32768:
            switch "" {
            case "hello":
                if false {
                    var v57 uint32 = 2147483647
                    switch ("" + "") {
                    case "foo bar":
                        if (true || true) {
                            var v58 uint8 = (v4 | v4)
                            for fi59 := 0; fi59 < 12; fi59++ {
                                fmt.Println(v58)
                                v58 -= v58
                                _ = fi59
                            }
                            _ = v58
                        }
                    case "":
                        switch "hello" {
                        case "hello":
                            var v60 int = -149
                            _ = v60
                        default:
                            var v61 uint32 = (v57 << ((515) & 63))
                            _ = v61
                        }
                    default:
                        v1 += v1
                    }
                    _ = v57
                }
            case "foo bar":
                v5 += v5
            }
        case 2147483647:
            v4 = (v4 | v4)
        case 127:
            v8 = v8
        default:
            if true {
                v4 -= 182
                var v62 bool = true
                _ = v62
            } else {
                for fi63 := 0; fi63 < 7; fi63++ {
                    for fi64 := 0; fi64 < 7; fi64++ {
                        if (!(!false)) {
                            if (!(true && false)) {
                                switch "foo bar" {
                                case "a":
                                    var v65 int64 = 0
                                    _ = v65
                                default:
                                    fmt.Println(v8, fi63, v8)
                                }
                                var v66 int32 = -1
                                if (3.14159 >= (1e10 + 0.1)) {
                                    v67 := GenS6{F0: (v66 * (v66 / (((v66 >> ((255) & 63))) | 1))), F1: 65535, F2: 32767}
                                    _ = v67
                                }
                                fmt.Println(fi63, fi64)
                                _ = v66
                            } else {
                                var v68 int = (fi63 - (-(fi64)))
                                fmt.Println(fi64)
                                _ = v68
                            }
                            if false {
                                if len(v2) > 0 {
                                    v2[0] = 4294967295
                                }
                                v1 += ((((v1 + ((0.0 / v1) * (v1 + v1))) * float32((((1e10 - 0.1) / (((3.14159 / -2.25) - 1.5) - 1.5)) * ((-2.25 - (1.5 * 1e10)) - -2.25)))) * v1) / (v1 - (v1 + (v1 + 3.14159))))
                                var v69 uint32 = 0
                                _ = v69
                            } else {
                                var v70 uint64 = v5
                                if (255 <= 0) {
                                    switch "" {
                                    case "\n\t":
                                        if (genFn0(-2147483648) > v4) {
                                            var v71 int8 = 1
                                            var v72 int = (fi64 % ((fi63) | 1))
                                            if (!false) {
                                                v5 = 0
                                                if (42 != 4294967295) {
                                                    var v73 uint32 = 42
                                                    v71 -= v71
                                                    _ = v73
                                                }
                                                v71 -= v71
                                            } else {
                                                v5 = v8
                                                if true {
                                                    v3[(("a" + ((("" + "foo bar") + "hello") + "")) + (((("a" + "hello") + "\n\t") + "foo bar") + "a"))] = (true || (-1 > -893))
                                                    var v74 string = (("" + "a") + "hello")
                                                    switch (fi63 | (fi63 << ((862) & 63))) {
                                                    case -770:
                                                        var v75 int = v72
                                                        _ = v75
                                                    default:
                                                        v1 = 1.5
                                                    }
                                                    v76 := GenS7{F0: -2147483648}
                                                    _ = v74
                                                    _ = v76
                                                } else {
                                                    var v77 uint64 = v70
                                                    fmt.Println(v8, v8)
                                                    switch "\n\t" {
                                                    case "":
                                                        if (v77 != 629) {
                                                            v2 = append(v2, 2147483648)
                                                            fmt.Println(v1)
                                                            v77 = v5
                                                            switch fi63 {
                                                            case 2147483647:
                                                                v71 = v71
                                                                v8 += v77
                                                            case 9223372036854775807:
                                                                v8 = (v8 ^ v70)
                                                            default:
                                                                var v78 float32 = ((3.14159 + v1) * (float32((0.0 + 1.5)) / v1))
                                                                _ = v78
                                                            }
                                                            var v79 int32 = -2147483648
                                                            _ = v79
                                                        } else {
                                                            fmt.Println(v5)
                                                            v71 = (v71 / (((^(v71))) | 1))
                                                            v80 := v3[""]
                                                            v72 -= v72
                                                            _ = v80
                                                        }
                                                    case "foo bar":
                                                        if true {
                                                            v81 := GenS1{F0: 4294967295}
                                                            var v82 int8 = v71
                                                            _ = v81
                                                            _ = v82
                                                        } else {
                                                            switch (-(fi63)) {
                                                            case 32767:
                                                                var v83 uint64 = uint64(v4)
                                                                _ = v83
                                                            case 127:
                                                                if (((v1 / v1) / v1) == v1) {
                                                                    switch (("" + ("" + "a")) + ((("foo bar" + "") + "a") + "foo bar")) {
                                                                    case "\n\t":
                                                                        var v84 int32 = 1
                                                                        _ = v84
                                                                    case "":
                                                                        v5 = v8
                                                                    }
                                                                    var v85 float32 = (-2.25 * ((-2.25 / (float32(1e10) * (0.0 - 1e10))) / -2.25))
                                                                    _ = v85
                                                                } else {
                                                                    v86 := GenS5{F0: (^(v4))}
                                                                    v71 += v71
                                                                    _ = v86
                                                                }
                                                            }
                                                            var v87 int32 = -494
                                                            _ = v87
                                                        }
                                                    case "hello":
                                                        v88 := v3[("hello" + ("\n\t" + "hello"))]
                                                        _ = v88
                                                    default:
                                                        var v89 uint = 255
                                                        _ = v89
                                                    }
                                                    _ = v77
                                                }
                                                var v90 int64 = -536
                                                _ = v90
                                            }
                                            _ = v71
                                            _ = v72
                                        } else {
                                            _ = len(v2)
                                            v1 -= (v1 * v1)
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                    case "":
                                        var v91 int8 = 127
                                        _ = v91
                                    default:
                                        var v92 uint = 1
                                        _ = v92
                                    }
                                    _ = len(v2)
                                }
                                _ = len(v3)
                                _ = v70
                            }
                            var v93 uint32 = 2147483647
                            if false {
                                v1 = v1
                                var v94 int8 = 42
                                if true {
                                    v94 -= v94
                                    if true {
                                        delete(v3, "hello")
                                        if (!(("foo bar" + "\n\t") != (("a" + "") + ("foo bar" + "\n\t")))) {
                                            var v95 uint = 71
                                            _ = v95
                                        }
                                        switch "\n\t" {
                                        case "hello":
                                            var v96 int16 = -1
                                            _ = v96
                                        }
                                        switch fi63 {
                                        case 865:
                                            v97 := GenS3{F0: 710, F1: 652}
                                            _ = v97
                                        case -2147483648:
                                            v98 := GenS4{F0: ((("" + ("hello" + "hello")) + "hello") + "foo bar")}
                                            _ = v98
                                        case 4294967295:
                                            var v99 uint64 = (^(v5))
                                            _ = v99
                                        default:
                                            v1 = (v1 - 3.14159)
                                        }
                                        v100 := GenS4{F0: "a"}
                                        var v101 int16 = -257
                                        var v102 int = 2147483647
                                        var v103 int8 = (v94 ^ -19)
                                        _ = v100
                                        _ = v101
                                        _ = v102
                                        _ = v103
                                    }
                                    v104 := GenS4{F0: "hello"}
                                    _ = v104
                                }
                                _ = v94
                            } else {
                                if (0 > -1) {
                                    v105 := GenS4{F0: "foo bar"}
                                    switch 127 {
                                    case -939:
                                        if (42 != -1) {
                                            switch "\n\t" {
                                            case "a":
                                                _ = len(v3)
                                            default:
                                                if ((3.14159 / v1) < (((v1 * v1) / 0.1) / 0.1)) {
                                                    v106 := GenS7{F0: 2147483647}
                                                    _ = v106
                                                } else {
                                                    if (!false) {
                                                        v107 := GenS8{F0: ((1.5 * ((((((-2.25 * (-2.25 + float64((((((v1 / ((1e10 / (v1 - ((v1 * float32(-2.25)) * v1))) + 0.0)) - v1) + v1) - v1) / float32(3.14159))))) * (-2.25 / float64((v1 / 1e10)))) + -2.25) * (0.1 - float64(v1))) * -2.25) * 1.5)) > 0.1), F1: 622, F2: 0}
                                                        _ = v107
                                                    } else {
                                                        v93 -= v93
                                                    }
                                                    v3[((((("foo bar" + ("foo bar" + ("hello" + "foo bar"))) + "a") + "hello") + "hello") + "\n\t")] = (v8 == (v5 << ((42) & 63)))
                                                    v108 := v3["\n\t"]
                                                    _ = v108
                                                }
                                            }
                                            v109 := GenS9{F0: false}
                                            switch (-(fi64)) {
                                            case 32767:
                                                switch fi64 {
                                                case 799:
                                                    var v110 uint = 29
                                                    _ = v110
                                                case -2147483648:
                                                    fmt.Println(fi63)
                                                case 4294967295:
                                                    if false {
                                                        if false {
                                                            v111 := GenS4{F0: ""}
                                                            v112 := GenS10{F0: "hello", F1: 47, F2: 42}
                                                            v5 += 156
                                                            _ = v111
                                                            _ = v112
                                                        } else {
                                                            if true {
                                                                v113 := GenS9{F0: false}
                                                                if (!false) {
                                                                    v5 = (v8 % ((v5) | 1))
                                                                } else {
                                                                    v114 := GenS11{F0: ("\n\t" + "foo bar"), F1: 119}
                                                                    var v115 int8 = 1
                                                                    _ = v114
                                                                    _ = v115
                                                                }
                                                                v93 = 2147483647
                                                                _ = v113
                                                            }
                                                            v116 := GenS6{F0: 953, F1: 0, F2: -762}
                                                            _ = v116
                                                        }
                                                    } else {
                                                        var v117 uint = 18446744073709551615
                                                        var v118 string = ("\n\t" + "a")
                                                        var v119 int16 = 1000
                                                        _ = v117
                                                        _ = v118
                                                        _ = v119
                                                    }
                                                default:
                                                    var v120 float64 = (1e10 * (-2.25 + 0.1))
                                                    _ = v120
                                                }
                                            case -716:
                                                var v121 string = ("\n\t" + ("hello" + ("\n\t" + "\n\t")))
                                                _ = v121
                                            default:
                                                var v122 float64 = 3.14159
                                                _ = v122
                                            }
                                            _ = v109
                                        }
                                    }
                                    if (v1 == v1) {
                                        var v123 int64 = 0
                                        var v124 int = 32767
                                        _ = v123
                                        _ = v124
                                    } else {
                                        switch (fi64 & (fi64 ^ (fi64 << ((864) & 63)))) {
                                        case 2147483647:
                                            if true {
                                                if true {
                                                    v125 := GenS11{F0: ("" + ("foo bar" + "hello")), F1: 4294967295}
                                                    if false {
                                                        var v126 float64 = (float64((-2.25 + 1e10)) + float64(-2.25))
                                                        if (!((fi64 % ((fi64) | 1)) > fi63)) {
                                                            fmt.Println(v8)
                                                            var v127 uint16 = uint16(v4)
                                                            v4 += 0
                                                            _ = v127
                                                        }
                                                        _ = v126
                                                    }
                                                    var v128 uint = 255
                                                    _ = v125
                                                    _ = v128
                                                } else {
                                                    v129 := GenS10{F0: "a", F1: uint(v5), F2: -128}
                                                    _ = v129
                                                }
                                                if true {
                                                    fmt.Println(v8)
                                                    if ((v4 * (v4 << ((723) & 63))) < 213) {
                                                        if (!(0 > (v8 - (v8 << ((42) & 63))))) {
                                                            v130 := v3["a"]
                                                            v131 := v3[("hello" + "foo bar")]
                                                            _ = v130
                                                            _ = v131
                                                        } else {
                                                            fmt.Println(fi63)
                                                            v93 = (v93 + v93)
                                                            v8 -= (v5 >> ((65535) & 63))
                                                            var v132 int64 = 184
                                                            _ = v132
                                                        }
                                                        v4 = (v4 & v4)
                                                        v133 := GenS12{F0: -418, F1: 1, F2: 65535}
                                                        var v134 int32 = 0
                                                        fmt.Println(fi64, v5, v8)
                                                        var v135 uint8 = v4
                                                        _ = v133
                                                        _ = v134
                                                        _ = v135
                                                    } else {
                                                        v4 = 0
                                                        v136 := GenS6{F0: 807, F1: 42, F2: 1000}
                                                        _ = v136
                                                    }
                                                    fmt.Println(v8, fi64, v93)
                                                }
                                                if ((v93 / ((v93) | 1)) >= uint32(v8)) {
                                                    fmt.Println(v93)
                                                }
                                            } else {
                                                fmt.Println(v1, v4)
                                                var v137 int = (fi64 >> ((0) & 63))
                                                _ = len(v2)
                                                _ = v137
                                            }
                                        case 9223372036854775807:
                                            fmt.Println(fi64, fi63)
                                        case 32767:
                                            v5 = (v5 / ((uint64(v93)) | 1))
                                        default:
                                            v1 += (v1 / (v1 - 1e10))
                                        }
                                    }
                                    var v138 int32 = 0
                                    _ = v105
                                    _ = v138
                                }
                                var v139 int8 = 100
                                var v140 bool = true
                                _ = v139
                                _ = v140
                            }
                            fmt.Println(fi64, v5, v4)
                            var v141 int32 = -45
                            fmt.Println(v5, v93, v141)
                            fmt.Println(fi64)
                            var v142 string = ("" + (("foo bar" + "\n\t") + "hello"))
                            _ = v141
                            _ = v142
                            _ = v93
                        }
                        _ = fi64
                    }
                    _ = fi63
                }
                delete(v3, (("foo bar" + "") + "hello"))
            }
        }
    }
    switch "a" {
    case "":
        var v143 int64 = -763
        _ = v143
    default:
        v5 -= v5
    }
    v144 := GenS5{F0: genFn0(713)}
    var v145 int32 = 1
    var v146 float64 = float64((v1 * ((3.14159 / (float32(float64(float32(0.0))) - 1.5)) + float32(1e10))))
    if len(v2) > 0 {
        v2[0] = 2
    }
    for fi147 := 0; fi147 < 1; fi147++ {
        for fi148 := 0; fi148 < 16; fi148++ {
            if ("" != "\n\t") {
                var v149 float32 = v1
                _ = len(v2)
                var v150 int64 = -9223372036854775808
                _ = v149
                _ = v150
            } else {
                if (!true) {
                    var v151 int32 = (v145 << ((255) & 63))
                    _ = v151
                }
                v146 -= 1e10
            }
            switch 2147483647 {
            case -1:
                switch "" {
                case "a":
                    _ = len(v2)
                }
            case 602:
                v152 := GenS6{F0: -2147483648, F1: 42, F2: -32768}
                _ = v152
            case 42:
                v153 := v3["\n\t"]
                _ = v153
            }
            v2 = append(v2, 65535)
            var v154 uint16 = 2
            _ = fi148
            _ = v154
        }
        _ = fi147
    }
    for fi155 := 0; fi155 < 13; fi155++ {
        if ((v4 == v4) || false) {
            v145 = -1
            v156 := GenS4{F0: "\n\t"}
            _ = v156
        } else {
            if false {
                fmt.Println(v1, fi155, v5)
                v8 -= v5
                var v157 uint16 = 42
                _ = v157
            } else {
                v3[("foo bar" + ("" + "hello"))] = true
            }
            v158 := GenS13{F0: 2}
            _ = v158
        }
        var v159 uint = 65535
        switch "hello" {
        case "\n\t":
            var v160 float64 = (1e10 + (1.5 * (0.1 - v146)))
            _ = v160
        case "a":
            if false {
                var v161 uint8 = genFn0(v145)
                _ = v161
            } else {
                v162 := GenS14{F0: (v4 | 128), F1: v4, F2: v5}
                _ = v162
            }
        }
        _ = fi155
        _ = v159
    }
    var v163 uint16 = 42
    v164 := GenS9{F0: true}
    v4 = (v4 - genFn0(2147483647))
    fmt.Println("v1:", v1)
    fmt.Println("v145:", v145)
    fmt.Println("v146:", v146)
    fmt.Println("v163:", v163)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v144
    _ = v145
    _ = v146
    _ = v163
    _ = v164
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
