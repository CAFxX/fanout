package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int; F1 int; F2 int }
type GenS5 struct { F0 int; F1 uint; F2 int8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint16; F1 int64 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint64; F1 int64 }

func main() {
    v1 := []int{-943, -1}
    var v2 int16 = -59
    var v3 float32 = 0.1
    var v4 uint64 = 2
    var v5 uint = 255
    if (float32(3.14159) < v3) {
        v5 = v5
    }
    if true {
        fmt.Println(v3, v4, v3)
    }
    v6 := GenS1{F0: 32767}
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(v4)
        _ = fi7
    }
    v2 = v2
    var v8 uint32 = 24
    switch "a" {
    case "":
        if false {
            switch ((("foo bar" + "a") + "\n\t") + ("a" + "")) {
            case "foo bar":
                v9 := GenS1{F0: 2}
                _ = v9
            case "":
                v3 = v3
            case "hello":
                v3 += (v3 / (v3 - v3))
            default:
                fmt.Println(v5)
            }
            switch 957 {
            case 128:
                var v10 uint = (^(v5))
                _ = v10
            default:
                v3 += (v3 + (v3 + 0.0))
            }
            switch ("a" + "a") {
            case "\n\t":
                v11 := GenS1{F0: 464}
                _ = v11
            case "hello":
                _ = len(v1)
            case "":
                for fi12 := 0; fi12 < 2; fi12++ {
                    switch "\n\t" {
                    case "\n\t":
                        v2 += 710
                    case "foo bar":
                        var v13 int64 = 138
                        _ = v13
                    case "":
                        var v14 bool = (false || false)
                        _ = v14
                    }
                    _ = fi12
                }
            default:
                if (3.14159 == 3.14159) {
                    for fi15 := 0; fi15 < 16; fi15++ {
                        if (v2 != v2) {
                            for fi16 := 0; fi16 < 5; fi16++ {
                                v3 += v3
                                _ = fi16
                            }
                        }
                        v17 := GenS2{F0: -192, F1: 1}
                        _ = fi15
                        _ = v17
                    }
                    fmt.Println(v4, v8)
                    var v18 int = 0
                    _ = v18
                }
            }
        } else {
            v19 := GenS3{F0: (v5 + v5), F1: 42}
            v1 = append(v1, -851)
            _ = v19
        }
    case "foo bar":
        switch int(v5) {
        case 42:
            if ((!(-787 < -9223372036854775808)) && true) {
                for fi20 := 0; fi20 < 2; fi20++ {
                    fmt.Println(v3, v3)
                    switch -2147483648 {
                    case 677:
                        if (-683 > 9223372036854775807) {
                            v5 -= v5
                            for fi21 := 0; fi21 < 17; fi21++ {
                                if (v4 < (v4 % ((uint64(v8)) | 1))) {
                                    v1 = append(v1, 505)
                                    if false {
                                        var v22 bool = false
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v22
                                    }
                                } else {
                                    if (!false) {
                                        v4 = v4
                                        var v23 float64 = ((0.1 * float64(v3)) / (0.1 - 1e10))
                                        _ = v23
                                    } else {
                                        if true {
                                            var v24 int16 = (-(v2))
                                            var v25 int64 = 199
                                            if ((42 != 191) || (0 != 100)) {
                                                if (("a" + "a") <= ("foo bar" + "")) {
                                                    v24 = -1
                                                    v2 -= -32768
                                                } else {
                                                    var v26 int64 = (v25 & v25)
                                                    v5 -= (v5 & v5)
                                                    _ = v26
                                                }
                                                var v27 int = (-(fi20))
                                                v24 = 60
                                                _ = v27
                                            } else {
                                                v28 := GenS2{F0: (v25 | (^(v25))), F1: 42}
                                                switch (-(fi20)) {
                                                case -1:
                                                    v29 := GenS3{F0: (v5 >> (((v5 / ((65535) | 1))) & 63)), F1: v5}
                                                    _ = v29
                                                case -2147483648:
                                                    switch 32767 {
                                                    case 32768:
                                                        switch "foo bar" {
                                                        case "":
                                                            v1 = append(v1, (^(fi21)))
                                                        case "\n\t":
                                                            if true {
                                                                var v30 string = "hello"
                                                                if true {
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        var v31 uint64 = v4
                                                                        var v32 int32 = 2147483647
                                                                        _ = v31
                                                                        _ = v32
                                                                    case "":
                                                                        v30 = (v30 + "a")
                                                                    }
                                                                    var v33 uint8 = uint8(v4)
                                                                    fmt.Println(v33, v2)
                                                                    _ = v33
                                                                } else {
                                                                    v1 = append(v1, fi20)
                                                                    if false {
                                                                        switch 612 {
                                                                        case 527:
                                                                            var v34 int8 = 23
                                                                            _ = v34
                                                                        case 2147483647:
                                                                            var v35 float32 = (float32(0.1) * v3)
                                                                            _ = v35
                                                                        case 593:
                                                                            v2 = (^(v24))
                                                                        default:
                                                                            switch ((("\n\t" + "\n\t") + v30) + "foo bar") {
                                                                            case "\n\t":
                                                                                v5 -= v5
                                                                            default:
                                                                                fmt.Println(v3)
                                                                                var v36 uint16 = 255
                                                                                fmt.Println(v24)
                                                                                _ = v36
                                                                            }
                                                                        }
                                                                        var v37 uint16 = 32767
                                                                        fmt.Println(fi21)
                                                                        _ = v37
                                                                    }
                                                                    if false {
                                                                        var v38 uint8 = 166
                                                                        _ = v38
                                                                    }
                                                                }
                                                                _ = v30
                                                            } else {
                                                                fmt.Println(v5)
                                                            }
                                                        }
                                                    case 2147483647:
                                                        var v39 int16 = -32768
                                                        _ = v39
                                                    }
                                                }
                                                _ = v28
                                            }
                                            v2 += (v24 + (v24 << (((v5 * (v5 >> ((v5) & 63)))) & 63)))
                                            var v40 uint32 = 42
                                            if true {
                                                switch 0 {
                                                case -2147483648:
                                                    if false {
                                                        fmt.Println(fi20, fi21)
                                                    } else {
                                                        var v41 uint8 = 0
                                                        _ = v41
                                                    }
                                                }
                                                if ((105 > 4) && ((v8 + (v8 / (((v40 & (v40 / ((v40) | 1)))) | 1))) != uint32(v4))) {
                                                    var v42 uint16 = 32767
                                                    _ = len(v1)
                                                    _ = v42
                                                } else {
                                                    if false {
                                                        var v43 int32 = -2147483648
                                                        v4 = (v4 * (v4 + v4))
                                                        if false {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            switch 128 {
                                                            case -144:
                                                                if true {
                                                                    v3 -= v3
                                                                }
                                                                v40 -= v8
                                                                var v44 int64 = (v25 | v25)
                                                                var v45 string = "a"
                                                                switch ("hello" + "hello") {
                                                                case "":
                                                                    if ((((v45 + ("foo bar" + "hello")) + v45) + "foo bar") <= v45) {
                                                                        var v46 uint64 = uint64(v8)
                                                                        v47 := GenS4{F0: (fi20 | fi21), F1: 127, F2: fi20}
                                                                        v46 = v4
                                                                        v48 := GenS5{F0: (fi20 + (-(fi20))), F1: v5, F2: 42}
                                                                        _ = v46
                                                                        _ = v47
                                                                        _ = v48
                                                                    }
                                                                default:
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        var v49 int32 = (v43 / ((687) | 1))
                                                                        _ = v49
                                                                    default:
                                                                        var v50 float32 = (float32(-2.25) - 0.1)
                                                                        _ = v50
                                                                    }
                                                                }
                                                                v8 = (v40 % ((v40) | 1))
                                                                switch (fi20 + int(v40)) {
                                                                case -715:
                                                                    switch v45 {
                                                                    case "foo bar":
                                                                        var v51 int8 = int8(v24)
                                                                        _ = v51
                                                                    case "hello":
                                                                        if false {
                                                                            var v52 int64 = 530
                                                                            v53 := GenS3{F0: (v5 >> ((65535) & 63)), F1: (^(v5))}
                                                                            _ = v52
                                                                            _ = v53
                                                                        }
                                                                    case "":
                                                                        var v54 int64 = v25
                                                                        v44 -= (v54 ^ v54)
                                                                        _ = v54
                                                                    }
                                                                }
                                                                var v55 float64 = ((-2.25 / (0.1 * 3.14159)) / ((3.14159 - -2.25) - 0.0))
                                                                _ = v44
                                                                _ = v45
                                                                _ = v55
                                                            case -1:
                                                                var v56 float64 = 0.1
                                                                _ = v56
                                                            default:
                                                                switch "hello" {
                                                                case "hello":
                                                                    v8 -= 38
                                                                case "a":
                                                                    if ((v24 & v2) >= (v2 >> ((v5) & 63))) {
                                                                        var v57 int8 = 0
                                                                        v8 = (v40 / (((v8 / (((^(v8))) | 1))) | 1))
                                                                        if true {
                                                                            if (("foo bar" != ("a" + (("\n\t" + "\n\t") + ""))) || (false && false)) {
                                                                                var v58 string = "a"
                                                                                v24 += (v2 & 1000)
                                                                                var v59 bool = (((v57 * (v57 >> ((4294967295) & 63))) >= v57) || (!false))
                                                                                _ = v58
                                                                                _ = v59
                                                                            }
                                                                            var v60 float64 = 1.5
                                                                            _ = v60
                                                                        }
                                                                        _ = v57
                                                                    }
                                                                default:
                                                                    fmt.Println(v5, v24)
                                                                }
                                                            }
                                                            switch (-(fi21)) {
                                                            case 42:
                                                                if (false && ((v43 - (v43 / ((-2147483648) | 1))) < v43)) {
                                                                    fmt.Println(v8, v3)
                                                                    fmt.Println(v24, v25, v25)
                                                                    fmt.Println(v43, v24, v25)
                                                                } else {
                                                                    var v61 float64 = float64(float32((((3.14159 + 0.0) - 1e10) * (1.5 / (1.5 / float64(v3))))))
                                                                    _ = v61
                                                                }
                                                            case 32768:
                                                                v62 := GenS6{F0: (v4 - (v4 + v4))}
                                                                _ = v62
                                                            case 758:
                                                                var v63 string = "hello"
                                                                _ = v63
                                                            default:
                                                                _ = len(v1)
                                                            }
                                                            if true {
                                                                v3 += ((v3 + v3) - 1e10)
                                                            }
                                                        } else {
                                                            fmt.Println(v25, v5)
                                                            if false {
                                                                var v64 int64 = (-(v25))
                                                                _ = v64
                                                            } else {
                                                                var v65 uint = (v5 << ((v5) & 63))
                                                                v66 := GenS3{F0: 0, F1: v5}
                                                                _ = v65
                                                                _ = v66
                                                            }
                                                            v67 := GenS7{F0: (fi21 / ((42) | 1))}
                                                            var v68 string = "foo bar"
                                                            _ = v67
                                                            _ = v68
                                                        }
                                                        if (true || (("\n\t" == "\n\t") && true)) {
                                                            if false {
                                                                var v69 uint = (v5 / ((v5) | 1))
                                                                _ = v69
                                                            }
                                                            switch (fi21 * fi20) {
                                                            case 32767:
                                                                v43 = 1
                                                            case 32768:
                                                                switch 7 {
                                                                case 128:
                                                                    v4 -= 4294967295
                                                                case 0:
                                                                    fmt.Println(fi20)
                                                                default:
                                                                    switch -882 {
                                                                    case -9223372036854775808:
                                                                        v24 += v2
                                                                    }
                                                                }
                                                            case 42:
                                                                var v70 uint32 = (v8 | v40)
                                                                _ = v70
                                                            }
                                                        } else {
                                                            if true {
                                                                if true {
                                                                    var v71 int = (fi20 + fi20)
                                                                    v72 := GenS8{F0: 42}
                                                                    v2 = (v2 << ((65535) & 63))
                                                                    var v73 bool = (!true)
                                                                    _ = v71
                                                                    _ = v72
                                                                    _ = v73
                                                                }
                                                                var v74 uint32 = (v8 - 668)
                                                                _ = v74
                                                            }
                                                        }
                                                        _ = v43
                                                    }
                                                    v4 += 0
                                                }
                                            }
                                            if ((v4 - (v4 - (v4 - 42))) > v4) {
                                                var v75 bool = true
                                                switch ("hello" + ((("\n\t" + "") + "foo bar") + (("foo bar" + "\n\t") + ("" + "hello")))) {
                                                case "foo bar":
                                                    var v76 int64 = int64(v24)
                                                    _ = v76
                                                case "\n\t":
                                                    var v77 uint32 = v8
                                                    v78 := GenS9{F0: 742, F1: (v25 / ((0) | 1))}
                                                    v40 = (v77 - v40)
                                                    fmt.Println(v5, v3, v40)
                                                    _ = v77
                                                    _ = v78
                                                case "hello":
                                                    if true {
                                                        v1 = append(v1, (fi20 & (fi20 - 42)))
                                                        var v79 int64 = (v25 >> ((v5) & 63))
                                                        if v75 {
                                                            var v80 uint32 = 42
                                                            _ = v80
                                                        } else {
                                                            var v81 string = "\n\t"
                                                            switch (fi20 | (^(fi20))) {
                                                            case 2147483647:
                                                                var v82 string = "\n\t"
                                                                v83 := GenS4{F0: fi21, F1: (^(fi21)), F2: (fi21 & -394)}
                                                                v82 = "a"
                                                                var v84 uint32 = (v40 & (v40 | v40))
                                                                var v85 uint16 = 0
                                                                var v86 float64 = 0.1
                                                                _ = v82
                                                                _ = v83
                                                                _ = v84
                                                                _ = v85
                                                                _ = v86
                                                            case -160:
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                            _ = v81
                                                        }
                                                        v25 = v79
                                                        v75 = (("\n\t" < "") || (-1 < -1))
                                                        _ = v79
                                                    }
                                                default:
                                                    v2 += v24
                                                }
                                                var v87 string = "foo bar"
                                                var v88 float64 = 1e10
                                                var v89 uint16 = 32767
                                                _ = v75
                                                _ = v87
                                                _ = v88
                                                _ = v89
                                            }
                                            if (uint16(v4) <= 65535) {
                                                v1 = append(v1, int(v8))
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                var v90 int32 = -2147483648
                                                _ = len(v1)
                                                v5 = 1
                                                _ = v90
                                            }
                                            _ = v24
                                            _ = v25
                                            _ = v40
                                        } else {
                                            if false {
                                                v1 = append(v1, (-(fi21)))
                                                var v91 float32 = 0.0
                                                var v92 float32 = ((((v91 + v3) + (v3 + ((v3 * float32(0.0)) * -2.25))) * 3.14159) + ((-2.25 / (v3 * float32(1.5))) + (v91 - (v3 / v91))))
                                                _ = v91
                                                _ = v92
                                            }
                                            if ("" >= ("\n\t" + ("hello" + "foo bar"))) {
                                                if false {
                                                    fmt.Println(v3)
                                                    v3 += v3
                                                    v93 := GenS2{F0: 42, F1: 0}
                                                    var v94 int32 = 1
                                                    _ = v93
                                                    _ = v94
                                                } else {
                                                    if true {
                                                        fmt.Println(v8, v8, fi21)
                                                        v3 += v3
                                                        var v95 bool = ((!(0 >= -2147483648)) && true)
                                                        _ = v95
                                                    }
                                                }
                                                fmt.Println(v4, fi21)
                                                v96 := GenS6{F0: (^(v4))}
                                                _ = v96
                                            } else {
                                                var v97 bool = false
                                                v98 := GenS8{F0: -128}
                                                if (!v97) {
                                                    var v99 uint8 = 127
                                                    switch (-(fi20)) {
                                                    case 127:
                                                        var v100 int32 = -1
                                                        _ = v100
                                                    default:
                                                        switch fi21 {
                                                        case 4294967295:
                                                            var v101 string = "hello"
                                                            _ = v101
                                                        default:
                                                            v102 := GenS10{F0: "foo bar"}
                                                            _ = v102
                                                        }
                                                    }
                                                    _ = len(v1)
                                                    _ = v99
                                                }
                                                _ = v97
                                                _ = v98
                                            }
                                        }
                                        if true {
                                            v8 += v8
                                        } else {
                                            v4 += 18446744073709551615
                                        }
                                    }
                                    var v103 int64 = 42
                                    if (float32(1e10) < ((v3 - v3) + float32(3.14159))) {
                                        _ = len(v1)
                                        v104 := GenS8{F0: -1}
                                        _ = v104
                                    } else {
                                        v5 -= (v5 - 255)
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = (fi21 & 32768)
                                    }
                                    _ = v103
                                }
                                v1 = append(v1, fi20)
                                var v105 int32 = 37
                                _ = len(v1)
                                _ = fi21
                                _ = v105
                            }
                            if len(v1) > 0 {
                                v1[0] = (fi20 * (fi20 ^ fi20))
                            }
                        }
                    case 946:
                        var v106 float32 = (v3 * v3)
                        _ = v106
                    default:
                        if len(v1) > 0 {
                            v1[0] = fi20
                        }
                    }
                    _ = fi20
                }
            } else {
                var v107 uint64 = (v4 << (((v5 % ((65535) | 1))) & 63))
                fmt.Println(v5, v107, v5)
                _ = v107
            }
        case 630:
            v108 := GenS11{F0: v5}
            _ = v108
        }
    }
    var v109 int32 = 2147483647
    v110 := GenS12{F0: 303, F1: 0}
    switch "foo bar" {
    case "":
        if (!(4294967295 != (v8 + v8))) {
            v8 += v8
        }
    }
    v1 = append(v1, 1)
    if len(v1) > 0 {
        v1[0] = 32767
    }
    for fi111 := 0; fi111 < 20; fi111++ {
        fmt.Println(v2, v8)
        if true {
            if true {
                v112 := GenS9{F0: uint16(v5), F1: 351}
                switch "" {
                case "hello":
                    _ = len(v1)
                case "\n\t":
                    v113 := GenS7{F0: 854}
                    _ = v113
                case "a":
                    v3 += v3
                }
                _ = v112
            }
        } else {
            switch fi111 {
            case 42:
                if true {
                    _ = len(v1)
                    for fi114 := 0; fi114 < 2; fi114++ {
                        if true {
                            switch ("" + "") {
                            case "":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            default:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            switch ("hello" + "a") {
                            case "\n\t":
                                fmt.Println(v3, v4, v8)
                            case "a":
                                var v115 int16 = (-(v2))
                                _ = v115
                            }
                        }
                        var v116 string = "hello"
                        _ = fi114
                        _ = v116
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            case -1:
                fmt.Println(v8, v5, v5)
            case 4294967295:
                if (-128 > 68) {
                    v117 := GenS10{F0: ("foo bar" + ("\n\t" + ""))}
                    var v118 uint = v5
                    if true {
                        var v119 uint8 = 129
                        fmt.Println(v119, v8)
                        _ = v119
                    }
                    _ = v117
                    _ = v118
                } else {
                    switch (fi111 ^ (fi111 & 1)) {
                    case 2147483647:
                        v4 = 0
                    }
                }
            default:
                v2 += (v2 + v2)
            }
            v8 -= (v8 ^ (^(v8)))
            var v120 bool = true
            _ = v120
        }
        _ = fi111
    }
    if true {
        var v121 uint16 = 32767
        v109 += (v109 / ((42) | 1))
        _ = v121
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi122 := 0; fi122 < 7; fi122++ {
        v4 = uint64(v8)
        _ = fi122
    }
    for fi123 := 0; fi123 < 20; fi123++ {
        fmt.Println(fi123, v4)
        _ = fi123
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v109:", v109)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v109
    _ = v110
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
