package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint32; F2 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 bool; F2 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64; F1 string; F2 uint64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint16; F1 int; F2 string }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint8; F1 int8 }
type GenS13 struct { F0 int64; F1 int32; F2 int16 }
type GenS14 struct { F0 uint; F1 uint32 }
type GenS15 struct { F0 uint; F1 int }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 int; F1 bool; F2 uint64 }

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    var v4 int = 1
    var v5 float32 = (v1 + (v1 / v1))
    delete(v2, 0)
    var v6 bool = (949 <= 2147483647)
    var v7 uint8 = 27
    v8 := []int{v4, (^(v4)), v4}
    v9 := GenS1{F0: 0}
    if (112 <= 18446744073709551615) {
        v8 = append(v8, -517)
        v10 := GenS1{F0: 9223372036854775807}
        _ = v10
    }
    v7 -= v7
    fmt.Println(v5)
    switch (v4 ^ -2147483648) {
    case -2147483648:
        v5 = v5
    case 128:
        v6 = false
    default:
        v5 -= v1
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        if ((v7 * v7) >= 0) {
            v4 -= (v4 + (fi11 & (-(fi11))))
            if (v6 || v6) {
                delete(v2, 0)
            }
            _ = len(v2)
        } else {
            var v12 bool = v6
            var v13 int16 = 204
            _ = v12
            _ = v13
        }
        switch (("a" + "foo bar") + "foo bar") {
        case "":
            if v6 {
                v14 := GenS1{F0: 565}
                v15 := GenS2{F0: 24, F1: 42, F2: (v4 * v4)}
                _ = v14
                _ = v15
            }
        case "hello":
            for fi16 := 0; fi16 < 19; fi16++ {
                if v6 {
                    _ = len(v3)
                }
                switch "" {
                case "foo bar":
                    v17 := GenS1{F0: 9223372036854775807}
                    _ = v17
                case "":
                    var v18 int64 = 1
                    _ = v18
                }
                _ = len(v2)
                _ = fi16
            }
        case "a":
            v8 = append(v8, (-(v4)))
        }
        var v19 int8 = 0
        _ = fi11
        _ = v19
    }
    if ((v6 && (!v6)) || (1000 <= -32768)) {
        switch "\n\t" {
        case "foo bar":
            v5 += 3.14159
        case "hello":
            for fi20 := 0; fi20 < 17; fi20++ {
                if (float32(0.0) > v1) {
                    switch "a" {
                    case "\n\t":
                        for fi21 := 0; fi21 < 15; fi21++ {
                            _ = len(v2)
                            if (!((!v6) || (!((v7 % (((v7 >> ((0) & 63))) | 1)) > (v7 / (((v7 + (^(v7)))) | 1)))))) {
                                var v22 uint16 = 220
                                _ = v22
                            } else {
                                var v23 int64 = 9223372036854775807
                                if v6 {
                                    var v24 int = 32768
                                    switch -679 {
                                    case 127:
                                        v25 := GenS3{F0: 286}
                                        _ = v25
                                    }
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    _ = v24
                                }
                                _ = v23
                            }
                            fmt.Println(v1)
                            switch ((("a" + "") + ("" + "\n\t")) + ("" + ("\n\t" + "\n\t"))) {
                            case "a":
                                var v26 int64 = -694
                                _ = v26
                            case "":
                                v1 = float32(0.0)
                            case "\n\t":
                                var v27 bool = v6
                                if (1 > 1) {
                                    fmt.Println(v4)
                                } else {
                                    fmt.Println(v27, v6, fi21)
                                    v27 = (v6 && true)
                                }
                                fmt.Println(fi21, fi21)
                                if (!(!(true && v6))) {
                                    v7 = v7
                                    var v28 uint = 167
                                    switch "" {
                                    case "\n\t":
                                        var v29 bool = true
                                        _ = v29
                                    case "foo bar":
                                        if (v5 != v5) {
                                            var v30 int16 = 1
                                            delete(v2, 4294967295)
                                            var v31 uint8 = 226
                                            _ = v30
                                            _ = v31
                                        }
                                    default:
                                        if true {
                                            v28 += v28
                                            if (-2147483648 <= -1) {
                                                if len(v8) >= 2 {
                                                    v8 = v8[1:len(v8)]
                                                }
                                                switch (fi21 ^ 709) {
                                                case 127:
                                                    if (v27 && v27) {
                                                        v27 = v6
                                                        if (2 != 42) {
                                                            _ = len(v2)
                                                        } else {
                                                            var v32 uint8 = v7
                                                            if len(v8) > 0 {
                                                                v8[0] = fi20
                                                            }
                                                            _ = v32
                                                        }
                                                    } else {
                                                        var v33 uint = 275
                                                        var v34 float32 = (-2.25 * (v5 - -2.25))
                                                        switch fi21 {
                                                        case 806:
                                                            v35 := GenS4{F0: v7, F1: (!v27), F2: (-6 == -118)}
                                                            _ = v35
                                                        case -9223372036854775808:
                                                            fmt.Println(v4)
                                                        }
                                                        v6 = v6
                                                        _ = v33
                                                        _ = v34
                                                    }
                                                default:
                                                    if v27 {
                                                        switch (v4 ^ (fi21 - (v4 ^ (fi20 % (((fi20 | -9223372036854775808)) | 1))))) {
                                                        case -2147483648:
                                                            var v36 float32 = -2.25
                                                            _ = v36
                                                        }
                                                    }
                                                }
                                                v2[97] = fi20
                                                var v37 int32 = 42
                                                v38 := GenS5{F0: 9223372036854775807}
                                                _ = v37
                                                _ = v38
                                            }
                                        } else {
                                            fmt.Println(v1, v5, fi20)
                                            if (true && false) {
                                                if v6 {
                                                    v39 := GenS6{F0: 0, F1: "foo bar", F2: 4294967295}
                                                    v40 := GenS4{F0: (v7 >> ((1) & 63)), F1: v6, F2: (v27 || (v7 < v7))}
                                                    fmt.Println(fi20)
                                                    v7 = (^(v7))
                                                    _ = v39
                                                    _ = v40
                                                } else {
                                                    if true {
                                                        v41 := GenS7{F0: true}
                                                        if (v28 <= (^(v28))) {
                                                            if len(v8) > 0 {
                                                                v8[0] = fi21
                                                            }
                                                            fmt.Println(fi21)
                                                            switch fi20 {
                                                            case -2147483648:
                                                                if ((((((v1 - v5) / (v5 / v1)) / -2.25) + (v1 - (v1 * v1))) / float32((-2.25 + 0.1))) != v5) {
                                                                    var v42 float64 = ((-2.25 * -2.25) + float64(v1))
                                                                    v3[(("foo bar" + "") + ("hello" + ("" + "a")))] = 600
                                                                    _ = v42
                                                                } else {
                                                                    var v43 int32 = -341
                                                                    v44 := GenS8{F0: 1}
                                                                    v28 += (v28 >> ((uint(v7)) & 63))
                                                                    switch ((("foo bar" + "") + "foo bar") + ("hello" + "a")) {
                                                                    case "a":
                                                                        v45 := GenS7{F0: v27}
                                                                        _ = v45
                                                                    case "\n\t":
                                                                        v46 := GenS9{F0: 497, F1: int(v7), F2: "hello"}
                                                                        if (!false) {
                                                                            var v47 float64 = (0.1 + -2.25)
                                                                            var v48 string = ""
                                                                            switch (fi21 / ((int(v28)) | 1)) {
                                                                            case 0:
                                                                                if false {
                                                                                    v49 := GenS8{F0: 42}
                                                                                    _ = v49
                                                                                }
                                                                            }
                                                                            _ = v47
                                                                            _ = v48
                                                                        }
                                                                        _ = v46
                                                                    case "":
                                                                        var v50 int8 = -90
                                                                        _ = v50
                                                                    default:
                                                                        v6 = ((v7 / ((176) | 1)) >= 2)
                                                                    }
                                                                    _ = v43
                                                                    _ = v44
                                                                }
                                                            case -9223372036854775808:
                                                                switch "\n\t" {
                                                                case "a":
                                                                    fmt.Println(fi21, v6, v7)
                                                                case "\n\t":
                                                                    if ((!false) || (!false)) {
                                                                        if (18446744073709551615 < 1) {
                                                                            v4 = (-(fi20))
                                                                            v5 -= 0.1
                                                                            v51 := GenS8{F0: 572}
                                                                            _ = v51
                                                                        } else {
                                                                            var v52 string = "hello"
                                                                            _ = v52
                                                                        }
                                                                    } else {
                                                                        v53 := v3[("\n\t" + "\n\t")]
                                                                        _ = v53
                                                                    }
                                                                default:
                                                                    fmt.Println(fi20)
                                                                }
                                                            }
                                                            var v54 uint32 = 717
                                                            if false {
                                                                _ = len(v2)
                                                                if ((^(v7)) >= v7) {
                                                                    var v55 float64 = (0.1 / 1.0)
                                                                    var v56 string = (("hello" + (("a" + "") + ("\n\t" + "hello"))) + "")
                                                                    _ = v55
                                                                    _ = v56
                                                                } else {
                                                                    v28 = 42
                                                                    var v57 int = fi21
                                                                    var v58 int = v57
                                                                    switch v58 {
                                                                    case -245:
                                                                        v54 = (^(v54))
                                                                    case -9223372036854775808:
                                                                        fmt.Println(v1, v27, v57)
                                                                    case -288:
                                                                        switch -9223372036854775808 {
                                                                        case -1:
                                                                            switch -639 {
                                                                            case -1:
                                                                                if (!v27) {
                                                                                    v27 = v27
                                                                                    var v59 int32 = int32(v7)
                                                                                    if v27 {
                                                                                        var v60 uint = v28
                                                                                        v60 = v28
                                                                                        _ = v60
                                                                                    } else {
                                                                                        var v61 uint = (v28 & 264)
                                                                                        var v62 uint8 = 26
                                                                                        _ = v61
                                                                                        _ = v62
                                                                                    }
                                                                                    _ = v59
                                                                                }
                                                                            }
                                                                        default:
                                                                            v63 := GenS10{F0: (-142 != -728)}
                                                                            _ = v63
                                                                        }
                                                                    }
                                                                    _ = v57
                                                                    _ = v58
                                                                }
                                                                var v64 uint8 = (v7 & (^(v7)))
                                                                _ = v64
                                                            }
                                                            _ = v54
                                                        } else {
                                                            v5 -= v1
                                                            var v65 int64 = 86
                                                            v65 += v65
                                                            _ = v65
                                                        }
                                                        switch "a" {
                                                        case "\n\t":
                                                            v4 -= (fi20 & fi20)
                                                        default:
                                                            switch ("foo bar" + "a") {
                                                            case "a":
                                                                v66 := GenS11{F0: -2147483648}
                                                                _ = v66
                                                            case "":
                                                                v67 := v2[36]
                                                                _ = v67
                                                            default:
                                                                _ = len(v8)
                                                            }
                                                        }
                                                        var v68 uint = (^(v28))
                                                        if true {
                                                            var v69 uint32 = 42
                                                            switch "hello" {
                                                            case "hello":
                                                                if ((v69 + v69) >= (v69 & v69)) {
                                                                    _ = len(v8)
                                                                    fmt.Println(fi21)
                                                                    v70 := GenS12{F0: uint8(v68), F1: 42}
                                                                    _ = v70
                                                                }
                                                            default:
                                                                var v71 int16 = 640
                                                                _ = v71
                                                            }
                                                            _ = v69
                                                        }
                                                        _ = v41
                                                        _ = v68
                                                    }
                                                    v28 = (^(v28))
                                                    fmt.Println(v1, fi20)
                                                }
                                                var v72 uint16 = 255
                                                v2[2147483648] = (fi21 | 766)
                                                var v73 int16 = 42
                                                _ = v72
                                                _ = v73
                                            } else {
                                                switch ("" + ("foo bar" + "a")) {
                                                case "hello":
                                                    if (0 > 65535) {
                                                        var v74 string = ""
                                                        if (-38 >= -822) {
                                                            var v75 int64 = 9223372036854775807
                                                            if len(v8) > 0 {
                                                                v8[0] = (v4 % (((v4 << (((v28 - (v28 | 18446744073709551615))) & 63))) | 1))
                                                            }
                                                            _ = v75
                                                        }
                                                        switch (v74 + "hello") {
                                                        case "hello":
                                                            if (-97 == 0) {
                                                                v4 = int(v7)
                                                            }
                                                        case "foo bar":
                                                            if (false && true) {
                                                                fmt.Println(v5, v27, v6)
                                                                switch fi20 {
                                                                case 282:
                                                                    _ = len(v8)
                                                                case -2147483648:
                                                                    var v76 uint = v28
                                                                    _ = v76
                                                                }
                                                                v5 -= -2.25
                                                                var v77 uint8 = v7
                                                                _ = v77
                                                            } else {
                                                                v78 := v3[(v74 + "foo bar")]
                                                                var v79 int = 32768
                                                                var v80 int16 = int16(v28)
                                                                _ = v78
                                                                _ = v79
                                                                _ = v80
                                                            }
                                                        case "a":
                                                            v5 = v5
                                                        }
                                                        _ = v74
                                                    } else {
                                                        var v81 int = fi20
                                                        var v82 bool = (v6 && (true || true))
                                                        if v27 {
                                                            var v83 int16 = -810
                                                            _ = v83
                                                        }
                                                        var v84 uint16 = 421
                                                        _ = v81
                                                        _ = v82
                                                        _ = v84
                                                    }
                                                case "\n\t":
                                                    if v6 {
                                                        v85 := GenS3{F0: 105}
                                                        v5 -= (1e10 + (1e10 * float32((3.14159 + 1.5))))
                                                        switch fi20 {
                                                        case 1:
                                                            if (0.0 >= (3.14159 - 3.14159)) {
                                                                fmt.Println(fi20)
                                                                v86 := GenS3{F0: 1}
                                                                _ = v86
                                                            }
                                                            if v6 {
                                                                if len(v8) >= 2 {
                                                                    v8 = v8[1:len(v8)]
                                                                }
                                                                var v87 int64 = 1
                                                                fmt.Println(v5)
                                                                _ = v87
                                                            }
                                                            var v88 int8 = 8
                                                            if (!(592 > 1)) {
                                                                fmt.Println(v5, v28, v5)
                                                            }
                                                            delete(v2, 42)
                                                            _ = v88
                                                        case 0:
                                                            if v6 {
                                                                v27 = v6
                                                                if false {
                                                                    if (0 < 0) {
                                                                        v89 := v2[uint32(v28)]
                                                                        v90 := GenS2{F0: v28, F1: 2147483648, F2: fi21}
                                                                        v28 += v28
                                                                        _ = v89
                                                                        _ = v90
                                                                    } else {
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            v27 = false
                                                                        case "hello":
                                                                            v91 := GenS13{F0: 9223372036854775807, F1: -1, F2: 42}
                                                                            fmt.Println(v4)
                                                                            v92 := GenS4{F0: uint8(v28), F1: ((^(v7)) == 192), F2: false}
                                                                            _ = v91
                                                                            _ = v92
                                                                        }
                                                                    }
                                                                } else {
                                                                    v4 = (fi21 & (v4 * fi20))
                                                                    v6 = (1 > 293)
                                                                }
                                                                fmt.Println(v4)
                                                            }
                                                        case -785:
                                                            if ((fi20 % ((fi20) | 1)) > fi21) {
                                                                fmt.Println(v4, fi20, v7)
                                                                if (!(-32768 > -302)) {
                                                                    _ = len(v3)
                                                                    if v27 {
                                                                        if (!(3.14159 != 3.14159)) {
                                                                            if (true || false) {
                                                                                _ = len(v2)
                                                                                v93 := v2[4294967295]
                                                                                v27 = true
                                                                                _ = v93
                                                                            }
                                                                            v94 := GenS5{F0: 42}
                                                                            _ = v94
                                                                        }
                                                                        v27 = v27
                                                                    }
                                                                }
                                                                fmt.Println(v1, fi20, v5)
                                                            }
                                                        default:
                                                            switch ("\n\t" + ("foo bar" + "\n\t")) {
                                                            case "":
                                                                v28 += 0
                                                            case "\n\t":
                                                                var v95 bool = ((v28 >> ((1) & 63)) >= v28)
                                                                if (((("\n\t" + ("\n\t" + "\n\t")) + "\n\t") + "foo bar") == ("\n\t" + ("\n\t" + "a"))) {
                                                                    if (fi21 <= 42) {
                                                                        var v96 float32 = v5
                                                                        _ = v96
                                                                    } else {
                                                                        _ = len(v8)
                                                                    }
                                                                    var v97 int64 = -9223372036854775808
                                                                    fmt.Println(v95, v6)
                                                                    _ = v97
                                                                } else {
                                                                    switch "hello" {
                                                                    case "":
                                                                        if (v5 >= (float32(0.0) / v5)) {
                                                                            if (v7 > v7) {
                                                                                fmt.Println(fi21)
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    _ = len(v3)
                                                                                case "foo bar":
                                                                                    switch (("hello" + "hello") + "") {
                                                                                    case "hello":
                                                                                        if true {
                                                                                            if true {
                                                                                                v4 = 652
                                                                                                var v98 float32 = (float32((0.0 / 1.0)) + 1e10)
                                                                                                if true {
                                                                                                    if (42 == uint64(v7)) {
                                                                                                        delete(v3, "hello")
                                                                                                        v99 := GenS14{F0: v28, F1: 4294967295}
                                                                                                        _ = v99
                                                                                                    }
                                                                                                }
                                                                                                _ = v98
                                                                                            }
                                                                                        }
                                                                                    case "foo bar":
                                                                                        v4 = (fi20 - 0)
                                                                                        v5 += 3.14159
                                                                                        v5 -= v1
                                                                                    default:
                                                                                        fmt.Println(v28)
                                                                                    }
                                                                                case "a":
                                                                                    v100 := GenS2{F0: 65535, F1: uint32(v28), F2: 0}
                                                                                    _ = v100
                                                                                }
                                                                                v95 = v95
                                                                            } else {
                                                                                if (v6 && (32767 < 65535)) {
                                                                                    if (109 > 937) {
                                                                                        v27 = v27
                                                                                    } else {
                                                                                        if v95 {
                                                                                            fmt.Println(v28, v4, v1)
                                                                                            if (v95 || v95) {
                                                                                                v3["\n\t"] = 569
                                                                                            } else {
                                                                                                var v101 int32 = 551
                                                                                                _ = v101
                                                                                            }
                                                                                        }
                                                                                        v102 := GenS5{F0: -92}
                                                                                        fmt.Println(v95, v1, v95)
                                                                                        v103 := GenS15{F0: (^(v28)), F1: 264}
                                                                                        var v104 uint64 = 572
                                                                                        if (v1 >= (v5 - (float32(0.1) - v1))) {
                                                                                            var v105 float32 = 0.1
                                                                                            v105 -= float32(-2.25)
                                                                                            var v106 uint16 = 770
                                                                                            _ = v105
                                                                                            _ = v106
                                                                                        } else {
                                                                                            v7 = (v7 - v7)
                                                                                        }
                                                                                        _ = v102
                                                                                        _ = v103
                                                                                        _ = v104
                                                                                    }
                                                                                    if v95 {
                                                                                        if len(v8) >= 2 {
                                                                                            v8 = v8[1:len(v8)]
                                                                                        }
                                                                                        var v107 int = (^(fi20))
                                                                                        _ = v107
                                                                                    }
                                                                                    if v6 {
                                                                                        v28 += (v28 << ((v28) & 63))
                                                                                        fmt.Println(v6, v95)
                                                                                        var v108 uint8 = (^(v7))
                                                                                        _ = v108
                                                                                    }
                                                                                }
                                                                                var v109 uint16 = 1
                                                                                v110 := GenS15{F0: (^(v28)), F1: fi20}
                                                                                _ = v109
                                                                                _ = v110
                                                                            }
                                                                            v8 = append(v8, (fi20 % ((fi20) | 1)))
                                                                            v1 = 1.5
                                                                        }
                                                                    }
                                                                    switch (v4 >> ((v28) & 63)) {
                                                                    case -676:
                                                                        v28 = (v28 % ((609) | 1))
                                                                    case 9223372036854775807:
                                                                        v111 := GenS9{F0: 42, F1: fi21, F2: (("foo bar" + "hello") + ("a" + "a"))}
                                                                        v95 = (762 > 2)
                                                                        _ = v111
                                                                    }
                                                                    if (("a" + "a") != "\n\t") {
                                                                        v112 := GenS9{F0: 42, F1: v4, F2: ("\n\t" + "")}
                                                                        var v113 string = ((("hello" + "") + "hello") + "a")
                                                                        _ = v112
                                                                        _ = v113
                                                                    } else {
                                                                        if ((fi21 + fi20) != 127) {
                                                                            v4 -= (fi21 & fi20)
                                                                        }
                                                                        v114 := v3["a"]
                                                                        _ = v114
                                                                    }
                                                                }
                                                                _ = v95
                                                            }
                                                        }
                                                        _ = v85
                                                    } else {
                                                        fmt.Println(v28, v4, v27)
                                                        v115 := GenS3{F0: 1}
                                                        _ = v115
                                                    }
                                                }
                                                v116 := GenS4{F0: v7, F1: v6, F2: (!v27)}
                                                var v117 uint64 = 2
                                                _ = v116
                                                _ = v117
                                            }
                                        }
                                    }
                                    var v118 uint8 = 232
                                    v6 = ((v28 ^ v28) < (v28 % (((v28 + v28)) | 1)))
                                    _ = v118
                                    _ = v28
                                }
                                _ = v27
                            }
                            _ = fi21
                        }
                    case "hello":
                        for fi119 := 0; fi119 < 19; fi119++ {
                            var v120 uint16 = 255
                            _ = fi119
                            _ = v120
                        }
                    case "":
                        var v121 float32 = (float32((1.5 - 0.0)) / float32(1.5))
                        _ = v121
                    default:
                        var v122 uint64 = 893
                        _ = v122
                    }
                }
                _ = fi20
            }
        default:
            var v123 int64 = 9223372036854775807
            _ = v123
        }
    }
    if (("\n\t" + "hello") != ("" + (("foo bar" + "") + ""))) {
        var v124 uint32 = 2147483648
        var v125 uint64 = 113
        v126 := GenS5{F0: 9223372036854775807}
        _ = v124
        _ = v125
        _ = v126
    } else {
        v6 = v6
        if (!false) {
            if (0 <= 312) {
                for fi127 := 0; fi127 < 8; fi127++ {
                    v8 = append(v8, (fi127 >> ((360) & 63)))
                    _ = fi127
                }
            } else {
                v128 := GenS10{F0: (!false)}
                _ = v128
            }
            var v129 float32 = 1e10
            v130 := GenS8{F0: 3}
            _ = v129
            _ = v130
        }
        for fi131 := 0; fi131 < 17; fi131++ {
            if (1 >= -2147483648) {
                if v6 {
                    if v6 {
                        v132 := GenS5{F0: 42}
                        var v133 int64 = -1
                        var v134 float32 = (float32(-2.25) + ((v5 * v1) - 1e10))
                        _ = v132
                        _ = v133
                        _ = v134
                    } else {
                        for fi135 := 0; fi135 < 18; fi135++ {
                            v6 = (v6 && (1e10 <= (v1 / v5)))
                            var v136 int = 1
                            _ = fi135
                            _ = v136
                        }
                        fmt.Println(v1, v1, v4)
                    }
                } else {
                    var v137 bool = v6
                    _ = v137
                }
            }
            v138 := v3[("a" + "hello")]
            _ = fi131
            _ = v138
        }
    }
    fmt.Println(v1, v5)
    switch ("\n\t" + "") {
    case "\n\t":
        v5 += (((v1 / 3.14159) * (float32(3.14159) + ((((-2.25 * v1) + float32(0.1)) * 0.0) + 0.0))) + v1)
    case "":
        var v139 float32 = (0.1 / -2.25)
        _ = v139
    }
    var v140 int64 = -9223372036854775808
    v141 := GenS16{F0: v140}
    switch int(v7) {
    case 0:
        v4 += -1
    case -2147483648:
        switch (v4 << ((352) & 63)) {
        case -2147483648:
            var v142 uint = 4294967295
            _ = v142
        case 128:
            var v143 int8 = 0
            _ = v143
        default:
            v144 := GenS17{F0: (v4 + (v4 / ((v4) | 1))), F1: ((((((v5 / v5) + 0.0) * (float32(0.1) / v5)) / ((0.0 / v1) / (v1 + v1))) * (v1 + (v1 / v5))) <= (0.0 * v1)), F2: 18446744073709551615}
            _ = v144
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v140:", v140)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v140
    _ = v141
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
