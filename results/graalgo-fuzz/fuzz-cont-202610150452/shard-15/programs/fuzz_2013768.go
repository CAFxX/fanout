package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 int32; F1 int16; F2 int }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string; F1 uint64 }
type GenS9 struct { F0 int16; F1 uint8 }
type GenS10 struct { F0 bool; F1 uint }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint64; F1 string }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 string; F1 uint64; F2 int32 }
type GenS15 struct { F0 uint16; F1 int8; F2 bool }
type GenS16 struct { F0 uint16; F1 int8 }
type GenS17 struct { F0 string; F1 uint16 }
type GenS18 struct { F0 int; F1 uint8 }
type GenS19 struct { F0 int32; F1 bool; F2 uint8 }
type GenS20 struct { F0 string; F1 int64 }
type GenS21 struct { F0 bool; F1 bool }
type GenS22 struct { F0 uint32; F1 uint32 }
func genFn0() int8 {
    var v6 int8 = 85
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v6)
    case "hello":
        v6 += (v6 >> ((255) & 63))
    default:
        v7 := []int{4294967295, -956, -2147483648}
        _ = v7
    }
    var v8 float32 = 0.1
    v9 := GenS1{F0: 9223372036854775807}
    _ = v6
    _ = v8
    _ = v9
    return v6
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 2147483647
    var v3 int64 = 9223372036854775807
    v4 := make(map[uint32]bool)
    v5 := []int{0, -2147483648}
    for fi10 := 0; fi10 < 14; fi10++ {
        if (v2 != v2) {
            v3 = -532
        }
        var v11 uint16 = 1
        v2 += v2
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        for fi13 := 0; fi13 < 1; fi13++ {
            v2 = (v2 / ((v2) | 1))
            v14 := GenS2{F0: (1.5 != 3.14159), F1: 42}
            if len(v5) > 0 {
                v5[0] = fi13
            }
            fmt.Println(fi12, fi13, v3)
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    var v15 uint = 255
    switch 481 {
    case 2147483647:
        if (370 < 42) {
            v16 := GenS1{F0: -2147483648}
            _ = v16
        }
    case 0:
        _ = len(v5)
    case 42:
        v15 += (^(v15))
    default:
        v3 = 887
    }
    switch ("hello" + (((("hello" + (("hello" + "\n\t") + "a")) + "foo bar") + "a") + "")) {
    case "\n\t":
        if false {
            if (!(true || ((!(false || true)) && (true && true)))) {
                v5 = append(v5, -953)
            } else {
                v17 := GenS1{F0: 32768}
                var v18 uint64 = 562
                var v19 uint16 = uint16(v15)
                _ = v17
                _ = v18
                _ = v19
            }
            switch 32767 {
            case 4294967295:
                switch 382 {
                case 42:
                    for fi20 := 0; fi20 < 3; fi20++ {
                        switch 679 {
                        case 32768:
                            v3 = -9223372036854775808
                        }
                        switch "\n\t" {
                        case "":
                            for fi21 := 0; fi21 < 7; fi21++ {
                                if false {
                                    fmt.Println(fi20, v2)
                                    var v22 string = ("" + "hello")
                                    var v23 int16 = -1
                                    switch 637 {
                                    case 0:
                                        v15 = (^(v15))
                                    case -1:
                                        var v24 int8 = 42
                                        _ = v24
                                    case -163:
                                        _ = len(v5)
                                    }
                                    if (true && true) {
                                        v25 := GenS2{F0: true, F1: 2}
                                        if len(v5) > 0 {
                                            v5[0] = fi21
                                        }
                                        if ((v23 != v23) && ((-(v2)) != (v2 + (v2 & (v2 >> (((v15 >> ((v15) & 63))) & 63)))))) {
                                            v22 = "hello"
                                            var v26 bool = ((v3 - v3) <= v3)
                                            _ = v26
                                        } else {
                                            switch (((v22 + "a") + "hello") + "hello") {
                                            case "a":
                                                fmt.Println(fi21)
                                            case "foo bar":
                                                fmt.Println(fi21)
                                            default:
                                                fmt.Println(v15, v2)
                                            }
                                            switch v22 {
                                            case "a":
                                                v22 = ("hello" + (v22 + (v22 + "hello")))
                                            case "":
                                                v22 = "foo bar"
                                            case "hello":
                                                if false {
                                                    var v27 int64 = (v3 >> ((v15) & 63))
                                                    var v28 int64 = v27
                                                    _ = v27
                                                    _ = v28
                                                }
                                            default:
                                                var v29 string = v22
                                                _ = v29
                                            }
                                        }
                                        _ = v25
                                    } else {
                                        var v30 uint32 = 980
                                        var v31 string = v22
                                        _ = v30
                                        _ = v31
                                    }
                                    var v32 string = v22
                                    _ = v22
                                    _ = v23
                                    _ = v32
                                }
                                _ = fi21
                            }
                        case "hello":
                            for fi33 := 0; fi33 < 8; fi33++ {
                                switch (("a" + ("hello" + "hello")) + ("hello" + "")) {
                                case "":
                                    if (255 == 2) {
                                        if true {
                                            var v34 int8 = -9
                                            _ = v34
                                        }
                                        if (false && false) {
                                            v2 += (-(v2))
                                            var v35 int16 = -1
                                            if true {
                                                if false {
                                                    fmt.Println(v15, fi33)
                                                    v15 = v15
                                                    v15 += (v15 - (v15 * 18446744073709551615))
                                                } else {
                                                    v3 += 9223372036854775807
                                                }
                                            }
                                            var v36 int32 = (^(v2))
                                            _ = v35
                                            _ = v36
                                        } else {
                                            switch (fi33 - (^(fi33))) {
                                            case 42:
                                                if true {
                                                    fmt.Println(fi33)
                                                    var v37 float32 = ((0.1 - 1e10) + 0.0)
                                                    fmt.Println(v3, v37, fi20)
                                                    _ = v37
                                                }
                                            case 1:
                                                var v38 uint = (v15 >> (((^(v15))) & 63))
                                                _ = v38
                                            case 9223372036854775807:
                                                v39 := GenS2{F0: true, F1: 2}
                                                _ = v39
                                            default:
                                                if false {
                                                    var v40 int8 = 104
                                                    if (1 >= 255) {
                                                        var v41 int32 = (-(v2))
                                                        switch "foo bar" {
                                                        case "":
                                                            v41 = v2
                                                        default:
                                                            var v42 uint8 = 2
                                                            if ((v40 ^ (v40 + v40)) > (v40 * 127)) {
                                                                if len(v5) > 0 {
                                                                    v5[0] = fi33
                                                                }
                                                                fmt.Println(v42, v41)
                                                                if ((v42 / ((199) | 1)) >= v42) {
                                                                    switch int(v3) {
                                                                    case -355:
                                                                        if (1.5 < 0.1) {
                                                                            fmt.Println(fi33, v15)
                                                                            v43 := GenS3{F0: "", F1: 2147483647}
                                                                            var v44 bool = true
                                                                            var v45 uint = 4294967295
                                                                            _ = v43
                                                                            _ = v44
                                                                            _ = v45
                                                                        }
                                                                    case -395:
                                                                        v46 := v1["foo bar"]
                                                                        _ = v46
                                                                    case -1:
                                                                        var v47 int32 = v41
                                                                        _ = v47
                                                                    }
                                                                    var v48 uint32 = 2147483647
                                                                    _ = v48
                                                                }
                                                            } else {
                                                                var v49 int64 = (v3 ^ v3)
                                                                _ = v49
                                                            }
                                                            _ = v42
                                                        }
                                                        fmt.Println(v2)
                                                        _ = v41
                                                    } else {
                                                        if (!((float32(3.14159) / 1.0) <= (0.1 / 1.5))) {
                                                            switch fi33 {
                                                            case -1:
                                                                v50 := GenS4{F0: (v2 - (v2 + -1)), F1: 42, F2: (fi20 + (-(fi33)))}
                                                                _ = v50
                                                            case -9223372036854775808:
                                                                if false {
                                                                    var v51 int8 = (-(v40))
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        v52 := GenS5{F0: (v2 ^ (v2 | (v2 * (-(v2))))), F1: (v40 | (v51 ^ v51))}
                                                                        _ = v52
                                                                    case "foo bar":
                                                                        v40 = (v51 % (((v51 ^ (v40 | -55))) | 1))
                                                                    case "hello":
                                                                        fmt.Println(v3, v2, v51)
                                                                    default:
                                                                        var v53 bool = ((0 == v2) && ((32767 <= -17) && false))
                                                                        var v54 bool = false
                                                                        _ = v53
                                                                        _ = v54
                                                                    }
                                                                    _ = v51
                                                                } else {
                                                                    fmt.Println(fi20, v40)
                                                                    var v55 float32 = (float32(3.14159) - 3.14159)
                                                                    if ((^(v15)) > 729) {
                                                                        v56 := GenS6{F0: ""}
                                                                        _ = v56
                                                                    } else {
                                                                        v2 = -2147483648
                                                                        if false {
                                                                            var v57 int16 = -32768
                                                                            fmt.Println(fi20, fi20)
                                                                            _ = v57
                                                                        } else {
                                                                            v3 = v3
                                                                            switch fi33 {
                                                                            case 32767:
                                                                                var v58 string = ("" + "\n\t")
                                                                                fmt.Println(v15, fi20)
                                                                                var v59 uint64 = 1
                                                                                _ = v58
                                                                                _ = v59
                                                                            case 2147483647:
                                                                                v60 := v1["foo bar"]
                                                                                _ = v60
                                                                            case 993:
                                                                                if ((!false) && true) {
                                                                                    v61 := GenS3{F0: "", F1: 128}
                                                                                    switch "\n\t" {
                                                                                    case "foo bar":
                                                                                        var v62 int16 = int16(v40)
                                                                                        _ = v62
                                                                                    case "a":
                                                                                        _ = len(v5)
                                                                                    default:
                                                                                        var v63 int32 = -2147483648
                                                                                        _ = v63
                                                                                    }
                                                                                    _ = v61
                                                                                }
                                                                            default:
                                                                                var v64 int8 = -43
                                                                                _ = v64
                                                                            }
                                                                        }
                                                                    }
                                                                    var v65 int32 = v2
                                                                    var v66 float64 = 1e10
                                                                    _ = v55
                                                                    _ = v65
                                                                    _ = v66
                                                                }
                                                            case 9223372036854775807:
                                                                if (((3.14159 + float64((((1.5 + 0.0) - -2.25) * -2.25))) * 3.14159) >= 1e10) {
                                                                    var v67 string = "foo bar"
                                                                    v3 -= (v3 - v3)
                                                                    v68 := GenS2{F0: false, F1: 18446744073709551615}
                                                                    if (-873 <= 1) {
                                                                        var v69 int8 = v40
                                                                        _ = v69
                                                                    }
                                                                    _ = v67
                                                                    _ = v68
                                                                }
                                                            default:
                                                                v70 := GenS4{F0: v2, F1: 42, F2: (fi33 & fi20)}
                                                                _ = v70
                                                            }
                                                            v71 := v1[("foo bar" + ("foo bar" + "hello"))]
                                                            var v72 int = 1
                                                            _ = v71
                                                            _ = v72
                                                        }
                                                        switch "" {
                                                        case "foo bar":
                                                            switch ("foo bar" + (((("hello" + "hello") + ("a" + ("a" + "a"))) + (("foo bar" + "") + "a")) + "\n\t")) {
                                                            case "a":
                                                                switch (-(fi20)) {
                                                                case -1:
                                                                    v2 = v2
                                                                case 127:
                                                                    if false {
                                                                        var v73 uint32 = 1
                                                                        if len(v5) > 0 {
                                                                            v5[0] = fi33
                                                                        }
                                                                        _ = v73
                                                                    } else {
                                                                        if (2 <= 609) {
                                                                            var v74 int32 = (-(v2))
                                                                            var v75 int8 = genFn0()
                                                                            fmt.Println(fi20, v40, v2)
                                                                            _ = v74
                                                                            _ = v75
                                                                        }
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                var v76 uint32 = 2
                                                                _ = v76
                                                            default:
                                                                v2 = v2
                                                            }
                                                        default:
                                                            var v77 uint8 = 112
                                                            _ = v77
                                                        }
                                                    }
                                                    _ = v40
                                                }
                                            }
                                        }
                                        v15 = 42
                                    } else {
                                        delete(v1, ((("foo bar" + ("hello" + "\n\t")) + "a") + ""))
                                        v2 = v2
                                        var v78 string = "foo bar"
                                        _ = v78
                                    }
                                case "a":
                                    v15 += v15
                                default:
                                    switch (fi33 - -205) {
                                    case 127:
                                        var v79 bool = (!true)
                                        _ = v79
                                    case -1:
                                        v3 -= v3
                                    case 0:
                                        switch fi20 {
                                        case -9223372036854775808:
                                            if true {
                                                if false {
                                                    var v80 uint = 1
                                                    _ = v80
                                                }
                                            } else {
                                                v81 := GenS7{F0: 255}
                                                switch fi33 {
                                                case 2147483647:
                                                    if (!(!((!false) && ((true && false) && (true && false))))) {
                                                        var v82 uint32 = 2147483648
                                                        v83 := GenS1{F0: fi20}
                                                        _ = v82
                                                        _ = v83
                                                    } else {
                                                        v2 += int32(v3)
                                                        switch 2147483647 {
                                                        case 9223372036854775807:
                                                            if (2 < 0) {
                                                                var v84 int16 = 0
                                                                _ = v84
                                                            } else {
                                                                switch 127 {
                                                                case -9223372036854775808:
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        if false {
                                                                            _ = len(v1)
                                                                            if true {
                                                                                fmt.Println(v2, fi33)
                                                                                switch ("foo bar" + "a") {
                                                                                case "hello":
                                                                                    switch "foo bar" {
                                                                                    case "foo bar":
                                                                                        if true {
                                                                                            var v85 uint64 = 18446744073709551615
                                                                                            var v86 uint32 = 0
                                                                                            var v87 uint64 = v85
                                                                                            _ = v85
                                                                                            _ = v86
                                                                                            _ = v87
                                                                                        } else {
                                                                                            var v88 int64 = (^(v3))
                                                                                            switch ((("foo bar" + "a") + "\n\t") + "foo bar") {
                                                                                            case "foo bar":
                                                                                                var v89 uint8 = 2
                                                                                                _ = v89
                                                                                            case "":
                                                                                                if ("hello" < "hello") {
                                                                                                    if len(v5) > 0 {
                                                                                                        v5[0] = (fi20 / ((fi20) | 1))
                                                                                                    }
                                                                                                    var v90 uint8 = 0
                                                                                                    var v91 int32 = v2
                                                                                                    _ = v90
                                                                                                    _ = v91
                                                                                                }
                                                                                            }
                                                                                            _ = v88
                                                                                        }
                                                                                    case "":
                                                                                        v92 := GenS8{F0: ("a" + (("" + "a") + ("a" + ("a" + "\n\t")))), F1: 20}
                                                                                        _ = v92
                                                                                    case "a":
                                                                                        v3 -= (^(v3))
                                                                                    default:
                                                                                        v3 = (v3 << ((255) & 63))
                                                                                    }
                                                                                case "":
                                                                                    var v93 float64 = 1.5
                                                                                    _ = v93
                                                                                case "a":
                                                                                    v3 = v3
                                                                                default:
                                                                                    var v94 uint16 = 0
                                                                                    _ = v94
                                                                                }
                                                                            }
                                                                            v95 := v4[4294967295]
                                                                            if len(v5) > 0 {
                                                                                v5[0] = -2147483648
                                                                            }
                                                                            v3 = v3
                                                                            if (((1e10 + -2.25) + 3.14159) == (0.1 + -2.25)) {
                                                                                var v96 float32 = 0.0
                                                                                var v97 string = "foo bar"
                                                                                _ = v96
                                                                                _ = v97
                                                                            }
                                                                            _ = v95
                                                                        }
                                                                        v98 := GenS9{F0: -1, F1: 42}
                                                                        _ = v98
                                                                    default:
                                                                        fmt.Println(fi33, v15, v15)
                                                                    }
                                                                }
                                                            }
                                                        case 32767:
                                                            if (855 == uint64(v15)) {
                                                                if true {
                                                                    if ((^(v3)) < (v3 - v3)) {
                                                                        fmt.Println(v15, fi20, v15)
                                                                        v99 := GenS2{F0: true, F1: uint64(v15)}
                                                                        fmt.Println(v2)
                                                                        _ = v99
                                                                    }
                                                                    var v100 uint8 = 127
                                                                    _ = v100
                                                                }
                                                                switch (fi20 << ((18446744073709551615) & 63)) {
                                                                case 32768:
                                                                    if true {
                                                                        var v101 uint8 = 1
                                                                        var v102 int8 = -1
                                                                        v2 = v2
                                                                        _ = v101
                                                                        _ = v102
                                                                    } else {
                                                                        var v103 int = fi20
                                                                        switch "a" {
                                                                        case "hello":
                                                                            v4[1] = (2147483647 > fi33)
                                                                        }
                                                                        v15 -= v15
                                                                        _ = v103
                                                                    }
                                                                case 377:
                                                                    if (!true) {
                                                                        v3 -= v3
                                                                    } else {
                                                                        var v104 int64 = (v3 | 686)
                                                                        if (128 > 127) {
                                                                            fmt.Println(v15)
                                                                            if (true || (0 == 0)) {
                                                                                fmt.Println(fi33, fi33)
                                                                            } else {
                                                                                _ = len(v4)
                                                                                var v105 int8 = 125
                                                                                _ = v105
                                                                            }
                                                                        } else {
                                                                            switch ("a" + "a") {
                                                                            case "a":
                                                                                v106 := GenS8{F0: "hello", F1: 2}
                                                                                _ = v106
                                                                            case "":
                                                                                v104 -= (v104 | (v104 / (((v104 / ((349) | 1))) | 1)))
                                                                                var v107 int16 = 32767
                                                                                v108 := GenS5{F0: (-(v2)), F1: -15}
                                                                                _ = v107
                                                                                _ = v108
                                                                            default:
                                                                                v109 := GenS10{F0: false, F1: v15}
                                                                                _ = v109
                                                                            }
                                                                            switch fi33 {
                                                                            case -1:
                                                                                if true {
                                                                                    var v110 uint = 18446744073709551615
                                                                                    v3 = v3
                                                                                    _ = v110
                                                                                } else {
                                                                                    if (!true) {
                                                                                        fmt.Println(v2, fi33, v104)
                                                                                    } else {
                                                                                        switch fi33 {
                                                                                        case -749:
                                                                                            if (42 >= 1000) {
                                                                                                v111 := v4[991]
                                                                                                v111 = (!((!(fi20 <= -2147483648)) || true))
                                                                                                v5 = append(v5, 349)
                                                                                                fmt.Println(v104)
                                                                                                _ = v111
                                                                                            }
                                                                                        case 4294967295:
                                                                                            var v112 int64 = v104
                                                                                            _ = v112
                                                                                        default:
                                                                                            var v113 int = (^(fi33))
                                                                                            _ = v113
                                                                                        }
                                                                                        var v114 float32 = 0.1
                                                                                        _ = v114
                                                                                    }
                                                                                    v2 = 863
                                                                                    v115 := GenS10{F0: (true && false), F1: 288}
                                                                                    _ = v115
                                                                                }
                                                                            case 2147483647:
                                                                                var v116 uint16 = 608
                                                                                _ = v116
                                                                            case -9223372036854775808:
                                                                                v15 = v15
                                                                            }
                                                                        }
                                                                        if len(v5) > 0 {
                                                                            v5[0] = fi20
                                                                        }
                                                                        _ = v104
                                                                    }
                                                                case 127:
                                                                    if ((float64(((0.1 + -2.25) + (3.14159 * 3.14159))) == 0.1) && false) {
                                                                        delete(v1, "\n\t")
                                                                        var v117 uint16 = 741
                                                                        _ = v117
                                                                    } else {
                                                                        var v118 int16 = 1
                                                                        switch (^(fi33)) {
                                                                        case 170:
                                                                            switch (fi20 / (((fi20 / (((fi33 / ((127) | 1))) | 1))) | 1)) {
                                                                            case 42:
                                                                                v3 = v3
                                                                            }
                                                                        case -1:
                                                                            var v119 uint8 = 1
                                                                            _ = v119
                                                                        default:
                                                                            switch 32768 {
                                                                            case 1:
                                                                                switch 556 {
                                                                                case 369:
                                                                                    if true {
                                                                                        v120 := GenS11{F0: (v3 - v3)}
                                                                                        _ = v120
                                                                                    } else {
                                                                                        v118 = v118
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v118
                                                                    }
                                                                }
                                                                var v121 uint16 = 32767
                                                                _ = v121
                                                            } else {
                                                                v2 += v2
                                                            }
                                                        default:
                                                            var v122 uint8 = 0
                                                            _ = v122
                                                        }
                                                        if true {
                                                            var v123 string = ""
                                                            v124 := GenS12{F0: 416, F1: "\n\t"}
                                                            v3 += v3
                                                            _ = v123
                                                            _ = v124
                                                        }
                                                    }
                                                default:
                                                    if false {
                                                        v125 := GenS1{F0: int(v3)}
                                                        switch fi20 {
                                                        case 768:
                                                            var v126 int = fi20
                                                            _ = v126
                                                        case 32768:
                                                            var v127 uint64 = 0
                                                            v127 -= (v127 | 2)
                                                            _ = v127
                                                        }
                                                        delete(v4, 2147483648)
                                                        fmt.Println(v2)
                                                        v2 += (v2 ^ 276)
                                                        _ = v125
                                                    } else {
                                                        v5 = append(v5, (fi33 & 9223372036854775807))
                                                        var v128 float64 = ((0.1 - (0.0 - -2.25)) * (3.14159 / (1.5 / (-2.25 - float64(1e10)))))
                                                        _ = v128
                                                    }
                                                }
                                                _ = v81
                                            }
                                        case 42:
                                            if (false || (fi33 > 128)) {
                                                fmt.Println(v3)
                                                v2 -= v2
                                            } else {
                                                var v129 uint = 2
                                                _ = v129
                                            }
                                        default:
                                            v3 -= 9223372036854775807
                                        }
                                    default:
                                        fmt.Println(v3, v2, fi20)
                                        switch "a" {
                                        case "":
                                            var v130 int8 = 0
                                            _ = v130
                                        case "foo bar":
                                            var v131 int = (fi33 - fi20)
                                            v4[904] = ((^(v2)) != (v2 / ((int32(v15)) | 1)))
                                            _ = v131
                                        case "hello":
                                            if false {
                                                v132 := v1[""]
                                                delete(v1, ("a" + ("hello" + ("" + "foo bar"))))
                                                v1[("\n\t" + "foo bar")] = (fi20 - (fi20 ^ (fi33 >> ((2) & 63))))
                                                _ = v132
                                            }
                                        }
                                        if true {
                                            v133 := GenS13{F0: (!true)}
                                            var v134 string = "\n\t"
                                            _ = v133
                                            _ = v134
                                        } else {
                                            v3 = -1
                                            v135 := GenS14{F0: "a", F1: 1, F2: v2}
                                            _ = v135
                                        }
                                    }
                                }
                                switch (fi33 >> (((v15 % ((v15) | 1))) & 63)) {
                                case 32768:
                                    if false {
                                        var v136 uint = (v15 ^ 746)
                                        fmt.Println(v3, fi33)
                                        if ((42 < 672) && true) {
                                            if (42 <= 32767) {
                                                var v137 float64 = 3.14159
                                                v3 += (v3 + v3)
                                                if false {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    switch (fi33 - fi20) {
                                                    case 42:
                                                        v1[("foo bar" + "foo bar")] = (^(fi20))
                                                    default:
                                                        var v138 int8 = 42
                                                        _ = v138
                                                    }
                                                    var v139 bool = true
                                                    _ = v139
                                                } else {
                                                    v3 -= 251
                                                    var v140 uint32 = 65535
                                                    _ = len(v5)
                                                    _ = v140
                                                }
                                                _ = v137
                                            } else {
                                                v141 := GenS15{F0: 163, F1: -128, F2: (1 >= -41)}
                                                _ = v141
                                            }
                                            v3 += v3
                                            var v142 string = "a"
                                            if true {
                                                fmt.Println(fi20)
                                            } else {
                                                var v143 float64 = float64(1.5)
                                                v136 += v15
                                                _ = v143
                                            }
                                            if ((((-2.25 / ((0.0 * ((((float64(-2.25) / 1.0) + 3.14159) * 1e10) * -2.25)) - (1e10 + -2.25))) + 0.1) >= (((-2.25 * 0.0) + 0.1) - (3.14159 + 3.14159))) || true) {
                                                var v144 int64 = (v3 | v3)
                                                v145 := GenS4{F0: int32(fi20), F1: 42, F2: (fi20 << ((v136) & 63))}
                                                _ = v144
                                                _ = v145
                                            } else {
                                                v146 := GenS16{F0: 255, F1: 0}
                                                var v147 string = (v142 + (v142 + "\n\t"))
                                                _ = v146
                                                _ = v147
                                            }
                                            _ = v142
                                        }
                                        var v148 uint16 = 295
                                        _ = v136
                                        _ = v148
                                    }
                                }
                                v15 = v15
                                _ = fi33
                            }
                        case "foo bar":
                            v15 = (v15 ^ v15)
                        default:
                            switch -62 {
                            case -431:
                                v2 = v2
                            case 723:
                                for fi149 := 0; fi149 < 2; fi149++ {
                                    var v150 uint64 = 42
                                    v2 += (v2 - v2)
                                    var v151 int16 = -256
                                    _ = fi149
                                    _ = v150
                                    _ = v151
                                }
                            case 32767:
                                if (("foo bar" == "\n\t") && true) {
                                    if (94 <= v2) {
                                        v2 = v2
                                        v4[2147483647] = true
                                    } else {
                                        if ("" < ("" + (("" + ("foo bar" + "a")) + "\n\t"))) {
                                            if (!(!true)) {
                                                var v152 int32 = v2
                                                _ = v152
                                            } else {
                                                fmt.Println(fi20)
                                                var v153 bool = false
                                                _ = v153
                                            }
                                            for fi154 := 0; fi154 < 5; fi154++ {
                                                v155 := GenS15{F0: 65535, F1: int8(fi20), F2: false}
                                                _ = fi154
                                                _ = v155
                                            }
                                            v156 := GenS17{F0: "\n\t", F1: 525}
                                            _ = v156
                                        }
                                    }
                                    switch fi20 {
                                    case -36:
                                        fmt.Println(v3, v2)
                                    case 206:
                                        v15 = v15
                                    case 32768:
                                        fmt.Println(v15, fi20, v3)
                                    }
                                }
                            default:
                                v1[("hello" + "foo bar")] = fi20
                            }
                        }
                        for fi157 := 0; fi157 < 16; fi157++ {
                            v15 = (v15 ^ v15)
                            var v158 int = fi20
                            v3 = v3
                            var v159 uint8 = 1
                            if (false && ((0 >= (v3 & 9223372036854775807)) || (-2.25 == float32(0.1)))) {
                                switch -9223372036854775808 {
                                case -107:
                                    fmt.Println(fi157, v15)
                                    v158 -= int(v3)
                                    switch "a" {
                                    case "hello":
                                        v160 := GenS18{F0: (^(v158)), F1: 0}
                                        _ = v160
                                    case "":
                                        if (!((("" + (("\n\t" + "hello") + "a")) + "a") == ("a" + "a"))) {
                                            fmt.Println(v2)
                                        }
                                    case "foo bar":
                                        fmt.Println(v2, v15)
                                    default:
                                        v15 += 255
                                    }
                                case -943:
                                    var v161 int16 = 32767
                                    v162 := GenS18{F0: v158, F1: v159}
                                    var v163 uint8 = v159
                                    var v164 uint = (v15 % (((v15 / ((v15) | 1))) | 1))
                                    if (!false) {
                                        v5 = append(v5, -2147483648)
                                        fmt.Println(v164, v164, v2)
                                        v165 := GenS5{F0: v2, F1: 1}
                                        _ = v165
                                    } else {
                                        switch "a" {
                                        case "a":
                                            v166 := GenS19{F0: v2, F1: (646 == (v2 ^ v2)), F2: v163}
                                            _ = v166
                                        case "hello":
                                            v161 += (-(v161))
                                        }
                                        if true {
                                            fmt.Println(fi20, v158, v159)
                                        } else {
                                            if (!(255 >= 42)) {
                                                if true {
                                                    var v167 int = fi157
                                                    if false {
                                                        fmt.Println(v164, fi20, v15)
                                                        v168 := GenS2{F0: ((!((1.5 <= 0.0) && true)) && ((!(-128 != genFn0())) && ("hello" >= "foo bar"))), F1: 1}
                                                        _ = v168
                                                    }
                                                    _ = v167
                                                } else {
                                                    v3 -= (v3 - (v3 / (((v3 << ((v15) & 63))) | 1)))
                                                    fmt.Println(v2)
                                                }
                                                if true {
                                                    fmt.Println(v164, v2, v3)
                                                    switch (fi20 / (((fi157 - (^(fi20)))) | 1)) {
                                                    case 1:
                                                        if (442 <= 2147483647) {
                                                            var v169 int64 = v3
                                                            var v170 int16 = 1000
                                                            _ = v169
                                                            _ = v170
                                                        }
                                                    case 42:
                                                        v2 += (v2 >> ((679) & 63))
                                                    }
                                                    var v171 int16 = v161
                                                    _ = v171
                                                }
                                            } else {
                                                switch (fi157 * v158) {
                                                case -1:
                                                    _ = len(v5)
                                                case -9223372036854775808:
                                                    var v172 uint64 = 367
                                                    _ = v172
                                                }
                                                var v173 uint16 = 118
                                                _ = v173
                                            }
                                        }
                                    }
                                    var v174 int16 = (-(v161))
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = v161
                                    _ = v162
                                    _ = v163
                                    _ = v164
                                    _ = v174
                                case 295:
                                    var v175 bool = ((!true) && (false || (42 != 17)))
                                    _ = v175
                                default:
                                    if len(v5) > 0 {
                                        v5[0] = (^(fi20))
                                    }
                                }
                                var v176 int = v158
                                if (!((v15 | (v15 - 255)) < v15)) {
                                    var v177 bool = true
                                    fmt.Println(v15, v15, v3)
                                    delete(v1, "hello")
                                    v178 := GenS6{F0: ""}
                                    v177 = v177
                                    _ = v177
                                    _ = v178
                                } else {
                                    v2 = -925
                                    var v179 uint16 = 544
                                    var v180 int64 = (v3 - v3)
                                    var v181 string = ("hello" + "foo bar")
                                    if true {
                                        v182 := GenS11{F0: v3}
                                        _ = v182
                                    }
                                    v183 := GenS20{F0: v181, F1: -1}
                                    if (!((v180 | 42) == (v3 | (v180 ^ (^(v180)))))) {
                                        v184 := v4[281]
                                        if v184 {
                                            switch ((v181 + (v181 + ((("hello" + ((((v181 + (v181 + ((("" + v181) + ((((((v181 + v181) + (v181 + "a")) + "hello") + ((v181 + "\n\t") + "hello")) + (v181 + ("hello" + (v181 + v181)))) + "\n\t")) + "\n\t"))) + v181) + v181) + v181)) + "foo bar") + (v181 + ((v181 + v181) + "hello"))))) + "\n\t") {
                                            case "":
                                                v3 = (v3 & 1)
                                            default:
                                                v185 := GenS21{F0: ((0.0 * (1e10 / 0.1)) == 1e10), F1: (v15 <= v15)}
                                                _ = v185
                                            }
                                            if v184 {
                                                v186 := GenS11{F0: (v3 & -1)}
                                                _ = v186
                                            }
                                            _ = len(v1)
                                        }
                                        _ = v184
                                    } else {
                                        var v187 int64 = (v3 ^ v180)
                                        var v188 int8 = 100
                                        fmt.Println(v181, v187, v159)
                                        _ = v187
                                        _ = v188
                                    }
                                    var v189 uint32 = 2
                                    _ = v179
                                    _ = v180
                                    _ = v181
                                    _ = v183
                                    _ = v189
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                if ("a" < ("hello" + "\n\t")) {
                                    fmt.Println(v158)
                                    var v190 int32 = (v2 ^ 249)
                                    _ = v190
                                } else {
                                    var v191 uint = v15
                                    _ = len(v5)
                                    _ = v191
                                }
                                v3 = v3
                                v176 = 0
                                _ = v176
                            } else {
                                if ((((1.5 * float64((0.1 * ((3.14159 + (1e10 - 3.14159)) + (0.1 - 0.1))))) + -2.25) * float64(0.1)) <= (0.0 * 0.1)) {
                                    var v192 string = ("foo bar" + "\n\t")
                                    _ = v192
                                }
                                _ = len(v1)
                                switch ("" + "\n\t") {
                                case "\n\t":
                                    v193 := GenS22{F0: 4294967295, F1: 1}
                                    _ = v193
                                case "foo bar":
                                    fmt.Println(fi157, v158, fi20)
                                default:
                                    var v194 string = ""
                                    _ = v194
                                }
                            }
                            delete(v1, "foo bar")
                            fmt.Println(v158)
                            _ = fi157
                            _ = v158
                            _ = v159
                        }
                        _ = fi20
                    }
                case 4294967295:
                    _ = len(v1)
                }
            default:
                var v195 int8 = 1
                _ = v195
            }
            v2 = v2
        } else {
            if len(v5) > 0 {
                v5[0] = -336
            }
        }
    }
    v196 := GenS17{F0: "", F1: 1}
    fmt.Println(v2, v15)
    var v197 bool = false
    fmt.Println(v15, v3)
    var v198 string = "\n\t"
    v15 = (v15 << (((v15 * v15)) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v197:", v197)
    fmt.Println("v198:", v198)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v196
    _ = v197
    _ = v198
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
